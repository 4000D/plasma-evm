package pls

import (
	"bytes"
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/params"
)

const MAX_EPOCH_EVENTS = 0

var (
	baseCallOpt               = &bind.CallOpts{Pending: false, Context: context.Background()}
	requestableContractABI, _ = abi.JSON(strings.NewReader(rootchain.RequestableContractIABI))
	rootchainContractABI, _   = abi.JSON(strings.NewReader(rootchain.RootChainABI))

	//TODO: sholud delete this after fixing rcm.backend.NetworkId
	rootchainNetworkId = big.NewInt(1337)
)

type invalidExit struct {
	forkNumber  *big.Int
	blockNumber *big.Int
	receipt     *types.Receipt
	index       int64
	proof       []common.Hash
}

type invalidExits []*invalidExit

type RootChainManager struct {
	config *Config
	stopFn func()

	txPool     *core.TxPool
	blockchain *core.BlockChain

	backend           *ethclient.Client
	rootchainContract *rootchain.RootChain

	eventMux       *event.TypeMux
	accountManager *accounts.Manager

	miner *miner.Miner
	env   *miner.EpochEnvironment

	contractParams *rootchainParameters

	// channels
	quit             chan struct{}
	epochPreparedCh  chan *rootchain.RootChainEpochPrepared
	blockFinalizedCh chan *rootchain.RootChainBlockFinalized

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
}

func (rcm *RootChainManager) RootchainContract() *rootchain.RootChain { return rcm.rootchainContract }
func (rcm *RootChainManager) NRBEpochLength() (*big.Int, error) {
	return rcm.rootchainContract.NRBEpochLength(baseCallOpt)
}

func NewRootChainManager(
	config *Config,
	stopFn func(),
	txPool *core.TxPool,
	blockchain *core.BlockChain,
	backend *ethclient.Client,
	rootchainContract *rootchain.RootChain,
	eventMux *event.TypeMux,
	accountManager *accounts.Manager,
	miner *miner.Miner,
	env *miner.EpochEnvironment,
) (*RootChainManager, error) {
	rcm := &RootChainManager{

		config:            config,
		stopFn:            stopFn,
		txPool:            txPool,
		blockchain:        blockchain,
		backend:           backend,
		rootchainContract: rootchainContract,
		eventMux:          eventMux,
		accountManager:    accountManager,
		miner:             miner,
		env:               env,
		contractParams:    newRootchainParameters(rootchainContract, backend),
		quit:              make(chan struct{}),
		epochPreparedCh:   make(chan *rootchain.RootChainEpochPrepared, MAX_EPOCH_EVENTS),
		blockFinalizedCh:  make(chan *rootchain.RootChainBlockFinalized),
	}

	epochLength, err := rcm.NRBEpochLength()
	if err != nil {
		return nil, err
	}

	miner.SetNRBepochLength(epochLength)

	return rcm, nil
}

func (rcm *RootChainManager) Start() error {
	if err := rcm.run(); err != nil {
		return err
	}

	go rcm.pingBackend()

	return nil
}

func (rcm *RootChainManager) Stop() error {
	rcm.backend.Close()
	close(rcm.quit)
	return nil
}

func (rcm *RootChainManager) run() error {
	go rcm.runHandlers()
	go rcm.runSubmitter()
	go rcm.runDetector()

	if err := rcm.watchEvents(); err != nil {
		return err
	}

	return nil
}

// watchEvents watchs RootChain contract events
func (rcm *RootChainManager) watchEvents() error {
	filterer, err := rootchain.NewRootChainFilterer(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		return err
	}

	// rootchain block#1
	startBlockNumber := uint64(1)

	filterOpts := &bind.FilterOpts{
		Start:   startBlockNumber,
		End:     nil,
		Context: context.Background(),
	}

	// iterate previous events
	// TODO: the events fired while syncing should be dealt with in different way.
	iterator, err := filterer.FilterEpochPrepared(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating EpochPrepared event")

	for iterator.Next() {
		e := iterator.Event
		if e != nil {
			rcm.handleEpochPrepared(e)
		}
	}

	iterator2, err := filterer.FilterBlockFinalized(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating BlockFinalized event")

	for iterator2.Next() {
		e := iterator2.Event
		if e != nil {
			rcm.handleBlockFinalzied(e)
		}
	}

	// watch events from now
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   &startBlockNumber, // read events from rootchain block#1
	}

	epochPrepareWatchCh := make(chan *rootchain.RootChainEpochPrepared)
	epochPrepareSub, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareWatchCh)
	if err != nil {
		return err
	}

	log.Info("Watching EpochPrepared event", "startBlockNumber", startBlockNumber)

	blockFinalizedWatchCh := make(chan *rootchain.RootChainBlockFinalized)
	blockFinalizedSub, err := filterer.WatchBlockFinalized(watchOpts, blockFinalizedWatchCh)
	if err != nil {
		return err
	}

	log.Info("watching BlockFinalized event", "startBlockNumber", startBlockNumber)

	go func() {
		for {
			select {
			case e := <-epochPrepareWatchCh:
				if e != nil {
					rcm.epochPreparedCh <- e
				}

			case err := <-epochPrepareSub.Err():
				log.Error("EpochPrepared event subscription error", "err", err)
				rcm.stopFn()
				return

			case e := <-blockFinalizedWatchCh:
				if e != nil {
					rcm.blockFinalizedCh <- e
				}

			case err := <-blockFinalizedSub.Err():
				log.Error("BlockFinalized event subscription error", "err", err)
				rcm.stopFn()
				return

			case <-rcm.quit:
				return
			}
		}
	}()

	return nil
}

func (rcm *RootChainManager) runSubmitter() {
	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	w, err := rcm.accountManager.Find(rcm.config.Operator)
	if err != nil {
		log.Error("Failed to get operator wallet", "err", err)
	}

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			rcm.lock.Lock()

			blockInfo := ev.Data.(core.NewMinedBlockEvent)
			Nonce := rcm.contractParams.getNonce(rcm.backend)

			//TODO: rcm.backend.NetworkID does not work as intended. It should return 1337, not 1. And it should moved to rcm.config.RootchainNetworkId.
			networkID, err := rcm.backend.NetworkID(context.Background())
			log.Info("network Id", "id", networkID)
			if err != nil {
				log.Error("NetworkId error", "err", err)
			}

			// send request block to root chain contract
			if !rcm.env.IsRequest {
				input, err := rootchainContractABI.Pack(
					"submitNRB",
					blockInfo.Block.Header().Root,
					blockInfo.Block.Header().TxHash,
					blockInfo.Block.Header().ReceiptHash,
				)
				if err != nil {
					log.Error("Failed to pack submitNRB", "err", err)
				}
				submitTx := types.NewTransaction(Nonce, rcm.config.RootChainContract, rcm.contractParams.costNRB, params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)

				signedTx, err := w.SignTx(rcm.config.Operator, submitTx, rootchainNetworkId)
				if err != nil {
					log.Error("Failed to sign submitTx", "err", err)
				}

				err = rcm.backend.SendTransaction(context.Background(), signedTx)
				if err != nil {
					log.Error("Failed to send submitTx", "err", err)
				} else {
					// TODO: check TX is not reverted
					log.Info("NRB is submitted", "blockNumber", blockInfo.Block.NumberU64(), "hash", signedTx.Hash().Hex())
				}

				// send non-request block to root chain contract
			} else {
				input, err := rootchainContractABI.Pack(
					"submitORB",
					blockInfo.Block.Header().Root,
					blockInfo.Block.Header().TxHash,
					blockInfo.Block.Header().ReceiptHash,
				)
				if err != nil {
					log.Error("Failed to pack submitORB", "err", err)
				}
				submitTx := types.NewTransaction(Nonce, rcm.config.RootChainContract, rcm.contractParams.costORB, params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)

				signedTx, err := w.SignTx(rcm.config.Operator, submitTx, rootchainNetworkId)

				if err != nil {
					log.Error("Failed to sign submitTx", "err", err)
				}

				err = rcm.backend.SendTransaction(context.Background(), signedTx)
				if err != nil {
					log.Error("Failed to send submitTx", "err", err)
				} else {
					// TODO: check TX is not reverted
					log.Info("ORB is submitted", "blockNumber", blockInfo.Block.NumberU64(), "hash", signedTx.Hash().Hex())
				}
			}
			rcm.contractParams.incNonce()
			rcm.lock.Unlock()
		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) runHandlers() {
	for {
		select {
		case e := <-rcm.epochPreparedCh:
			if err := rcm.handleEpochPrepared(e); err != nil {
				log.Error("Failed to handle epoch prepared", "err", err)
			}
		case e := <-rcm.blockFinalizedCh:
			if err := rcm.handleBlockFinalzied(e); err != nil {
				log.Error("Failed to handle block finazlied", "err", err)
			}
		case <-rcm.quit:
			return
		}
	}
}

// handleEpochPrepared handles EpochPrepared event from RootChain contract after
// plasma chain is *SYNCED*.
func (rcm *RootChainManager) handleEpochPrepared(ev *rootchain.RootChainEpochPrepared) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	e := *ev

	log.Info("RootChain epoch prepared", "epochNumber", e.EpochNumber, "isRequest", e.IsRequest, "userActivated", e.UserActivated, "isEmpty", e.EpochIsEmpty)
	go rcm.eventMux.Post(miner.EpochPrepared{Payload: &e})
	// prepare request tx for ORBs
	if e.IsRequest && !e.EpochIsEmpty {
		events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
		defer events.Unsubscribe()

		numORBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
		numORBs = new(big.Int).Add(numORBs, big.NewInt(1))

		bodies := make([]types.Transactions, 0, numORBs.Uint64()) // [][]types.Transaction

		log.Debug("Num Orbs", "epochNumber", e.EpochNumber, "numORBs", numORBs, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)

		for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {
			currentFork, err := rcm.rootchainContract.CurrentFork(baseCallOpt)
			if err != nil {
				return err
			}

			pb, err := rcm.rootchainContract.Blocks(baseCallOpt, currentFork, blockNumber)
			if err != nil {
				return err
			}

			orb, err := rcm.rootchainContract.ORBs(baseCallOpt, big.NewInt(int64(pb.RequestBlockId)))
			if err != nil {
				return err
			}

			numRequests := orb.RequestEnd - orb.RequestStart + 1
			log.Debug("Fetching ORB", "requestBlockId", pb.RequestBlockId, "numRequests", numRequests)

			body := make(types.Transactions, 0, numRequests)

			for requestId := orb.RequestStart; requestId <= orb.RequestEnd; {
				request, err := rcm.rootchainContract.EROs(baseCallOpt, big.NewInt(int64(requestId)))
				if err != nil {
					return err
				}

				log.Debug("Request fetched", "requestId", requestId, "hash", common.Bytes2Hex(request.Hash[:]))

				var to common.Address
				var input []byte

				if request.IsTransfer {
					to = request.Requestor
				} else {
					to, _ = rcm.rootchainContract.RequestableContracts(baseCallOpt, request.To)
					input, err = requestableContractABI.Pack("applyRequestInChildChain",
						request.IsExit,
						big.NewInt(int64(requestId)),
						request.Requestor,
						request.TrieKey,
						request.TrieValue,
					)
					if err != nil {
						log.Error("Failed to pack applyRequestInChildChain", "err", err)
					}
				}

				requestTx := types.NewTransaction(0, to, request.Value, params.RequestTxGasLimit, params.RequestTxGasPrice, input)

				eroBytes, err := rcm.rootchainContract.GetEROBytes(baseCallOpt, big.NewInt(int64(requestId)))
				if err != nil {
					log.Error("Failed to get ERO bytes", "err", err)
				}

				if !bytes.Equal(eroBytes, requestTx.GetRlp()) {
					log.Error("ERO TX and request tx are different", "requestId", requestId, "eroBytes", common.Bytes2Hex(eroBytes), "requestTx.GetRlp()", common.Bytes2Hex(requestTx.GetRlp()))
				}

				body = append(body, requestTx)

				requestId += 1
			}

			log.Info("Request txs fetched", "blockNumber", blockNumber, "requestBlockId", pb.RequestBlockId, "body", body)

			bodies = append(bodies, body)

			blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
		}

		var numMinedORBs uint64 = 0

		for numMinedORBs < numORBs.Uint64() {
			rcm.txPool.EnqueueReqeustTxs(bodies[numMinedORBs])

			log.Info("Waiting new block mined event...")

			<-events.Chan()

			numMinedORBs += 1
		}
	}

	return nil
}

func (rcm *RootChainManager) handleBlockFinalzied(ev *rootchain.RootChainBlockFinalized) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	e := *ev

	log.Info("RootChain block finalized", "forkNumber", e.ForkNumber, "blockNubmer", e.BlockNumber)

	callerOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}

	w, err := rcm.accountManager.Find(rcm.config.Operator)
	if err != nil {
		log.Error("Failed to get operator wallet", "err", err)
	}

	block, err := rcm.rootchainContract.Blocks(callerOpts, e.ForkNumber, e.BlockNumber)
	if err != nil {
		return err
	}

	if block.IsRequest {
		fork := e.ForkNumber.Uint64()
		num := e.BlockNumber.Uint64()
		hash := rcm.blockchain.GetBlockByNumber(num).Hash()
		receipts := rcm.blockchain.GetReceiptsByHash(hash)
		ierc := rcm.blockchain.GetInvalidExitReceipts(fork, num)

		var proofs []byte

		for index, receipt := range ierc {
			proof := types.GetMerkleProof(receipts, int(index))
			for j := 0; j < len(proof); j++ {
				bytesOfproof := proof[j].Bytes()
				proofs = append(proofs, bytesOfproof...)
			}

			// TODO: ChallengeExit receipt check
			input, err := rootchainContractABI.Pack("challengeExit", e.ForkNumber, e.BlockNumber, big.NewInt(int64(index)), receipt.GetRlp(), proofs)
			if err != nil {
				log.Error("Failed to pack challengeExit", "error", err)
			}

			nonce := rcm.contractParams.getNonce(rcm.backend)
			challengeTx := types.NewTransaction(nonce, rcm.config.RootChainContract, big.NewInt(0), params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)
			signedTx, err := w.SignTx(rcm.config.Operator, challengeTx, rootchainNetworkId)
			if err != nil {
				log.Error("Failed to sign challengeTx", "err", err)
			}

			err = rcm.backend.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Error("Failed to send challengeTx", "err", err)
			} else {
				log.Info("challengeExit is submitted", "exit request number", index, "hash", signedTx.Hash().Hex())
			}
		}
	}

	return nil
}

func (rcm *RootChainManager) runDetector() {
	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	caller, err := rootchain.NewRootChainCaller(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		log.Warn("failed to make new root chain caller", "error", err)
		return
	}

	// TODO: check callOpts first if caller doesn't work.
	callerOpts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	for {
		select {
		case ev := <-events.Chan():
			rcm.lock.Lock()

			if rcm.env.IsRequest {
				forkNumber, err := caller.CurrentFork(callerOpts)

				if err != nil {
					log.Warn("failed to get current fork number", "error", err)
				}

				block := ev.Data.(core.NewMinedBlockEvent).Block
				receipts := rcm.blockchain.GetReceiptsByHash(block.Hash())

				// TODO: should check if the request[i] is enter or exit request. Undo request will make posterior enter request.
				// receipt index related invalid exit in a block
				var indices []uint64
				for i := 0; i < len(receipts); i++ {
					if receipts[i].Status == types.ReceiptStatusFailed {
						indices = append(indices, uint64(i))
						log.Info("invalid exit detected", "fork number", block.CurrentFork(), "fork number", forkNumber, "block number", block.Number)
					}
				}
				rcm.blockchain.SetInvalidExitReceipts(block.CurrentFork(), block.Hash(), block.NumberU64(), indices)
			}
			rcm.lock.Unlock()

		case <-rcm.quit:
			return
		}
	}
}

// pingBackend checks rootchain backend is alive.
func (rcm *RootChainManager) pingBackend() {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			if _, err := rcm.backend.SyncProgress(context.Background()); err != nil {
				log.Error("Rootchain provider doesn't respond", "err", err)
				ticker.Stop()
				rcm.stopFn()
				return
			}
		case <-rcm.quit:
			ticker.Stop()
			return
		}
	}
}

type rootchainParameters struct {
	// contract parameters
	costERO        *big.Int
	costERU        *big.Int
	costURBPrepare *big.Int
	costURB        *big.Int
	costORB        *big.Int
	costNRB        *big.Int
	maxRequests    *big.Int
	requestGas     *big.Int
	currentEpoch   *big.Int
	currentFork    *big.Int

	// operator tx parameters
	nonce uint64

	lastUpdateTime time.Time

	lock sync.Mutex
}

func newRootchainParameters(rootchainContract *rootchain.RootChain, backend *ethclient.Client) *rootchainParameters {
	rParams := &rootchainParameters{}

	rParams.getCostERO(rootchainContract)
	rParams.getCostERU(rootchainContract)
	rParams.getCostURBPrepare(rootchainContract)
	rParams.getCostURB(rootchainContract)
	rParams.getCostORB(rootchainContract)
	rParams.getCostNRB(rootchainContract)
	rParams.getMaxRequests(rootchainContract)
	rParams.getRequestGas(rootchainContract)
	rParams.getCurrentEpoch(rootchainContract)
	rParams.getCurrentFork(rootchainContract)

	rParams.getNonce(backend)

	return rParams
}

func (rp *rootchainParameters) getCostERU(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costERU, _ = rootchainContract.COSTERU(baseCallOpt)
	return rp.costERU
}
func (rp *rootchainParameters) getCostERO(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costERO, _ = rootchainContract.COSTERO(baseCallOpt)
	return rp.costERO
}
func (rp *rootchainParameters) getCostURBPrepare(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costURBPrepare, _ = rootchainContract.COSTURBPREPARE(baseCallOpt)
	return rp.costURBPrepare
}
func (rp *rootchainParameters) getCostURB(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costURB, _ = rootchainContract.COSTURB(baseCallOpt)
	return rp.costURB
}
func (rp *rootchainParameters) getCostORB(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costORB, _ = rootchainContract.COSTORB(baseCallOpt)
	return rp.costORB
}
func (rp *rootchainParameters) getCostNRB(rootchainContract *rootchain.RootChain) *big.Int {
	rp.costNRB, _ = rootchainContract.COSTNRB(baseCallOpt)
	return rp.costNRB
}
func (rp *rootchainParameters) getMaxRequests(rootchainContract *rootchain.RootChain) *big.Int {
	rp.maxRequests, _ = rootchainContract.MAXREQUESTS(baseCallOpt)
	return rp.maxRequests
}
func (rp *rootchainParameters) getRequestGas(rootchainContract *rootchain.RootChain) *big.Int {
	rp.requestGas, _ = rootchainContract.REQUESTGAS(baseCallOpt)
	return rp.requestGas
}
func (rp *rootchainParameters) getCurrentEpoch(rootchainContract *rootchain.RootChain) *big.Int {
	rp.currentEpoch, _ = rootchainContract.CurrentEpoch(baseCallOpt)
	return rp.currentEpoch
}
func (rp *rootchainParameters) getCurrentFork(rootchainContract *rootchain.RootChain) *big.Int {
	rp.currentFork, _ = rootchainContract.CurrentFork(baseCallOpt)
	return rp.currentFork
}
func (rp *rootchainParameters) getNonce(backend *ethclient.Client) uint64 {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	lastUpdateTime := rp.lastUpdateTime
	lastUpdateTime = lastUpdateTime.Add(2 * time.Second)

	now := time.Now()
	if now.Before(lastUpdateTime) {
		timer := time.NewTimer(lastUpdateTime.Sub(now))
		<-timer.C
	}

	rp.lastUpdateTime = time.Now()

	nonce, _ := backend.NonceAt(context.Background(), params.Operator, nil)
	if rp.nonce < nonce {
		rp.nonce = nonce
	}
	return rp.nonce
}
func (rp *rootchainParameters) incNonce() {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	rp.nonce += 1
}
