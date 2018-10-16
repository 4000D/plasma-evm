pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./SafeMath.sol";
import "./Math.sol";
import "./RLP.sol";
import "./RLPEncode.sol";
import "../patricia_tree/PatriciaTree.sol";
import {RequestableContractI} from "../RequestableContractI.sol";

library Data {
  using SafeMath for *;
  using Math for *;
  using RLP for *;
  using RLPEncode for *;

  // signature of function applyRequestInChildChain(bool isExit,uint256 requestId,address requestor,bytes32 trieKey,bytes32 trieValue)
  bytes4 public constant APPLY_IN_CHILDCHAIN_SIGNATURE = 0xe904e3d9;

  // signature of function applyRequestInRootChain(bool isExit,uint256 requestId,address requestor,bytes32 trieKey,bytes32 trieValue)
  bytes4 public constant APPLY_IN_ROOTCHAIN_SIGNATURE = 0xd9afd3a9;

  address public constant NA = address(0);
  uint public constant NA_TX_GAS_PRICE = 1;
  uint public constant NA_TX_GAS_LIMIT = 100000;

  struct Epoch {
    uint64 requestStart;      // first request id
    uint64 requestEnd;        // last request id
    uint64 startBlockNumber;  // first block number of the epoch
    uint64 endBlockNumber;    // last block number of the epoch
    uint64 forkedBlockNumber; // forked block number due to URB or challenge
                              // last finalized block is forkedBlockNumber - 1

    uint64 limit;             // the maximum number of request transactions in
                              // a request block

    uint64 timestamp;         // timestamp when the epoch is initialized.
                              // required for URB / ORB

    bool initialized;         // true if epoch is initialized
    bool isRequest;           // true in case of URB / ORB
    bool userActivated;       // true in case of URB
    bool challenged;          // true if a block in the epoch is challenged
    bool challenging;         // true if a block in the epoch is being challenged
    bool finalized;           // true if it is successfully finalized
  }

  function getNumBlocks(Epoch _e) internal returns (uint) {
    if (_e.startBlockNumber == _e.endBlockNumber) return 0;
    return _e.endBlockNumber - _e.startBlockNumber + 1;
  }

  /**
   * @notice This returns the request block number if the request is included
   *         in an epoch. Otherwise, returns 0.
   */
  function getBlockNumber(Epoch _e, uint _requestId) internal returns (uint) {
    if (!_e.isRequest
      || _e.limit == 0
      || _e.requestStart < _requestId
      || _e.requestEnd > _requestId) {
        return 0;
    }

    return uint(_e.startBlockNumber)
      .add(uint(_requestId - _e.requestStart + 1).divCeil(_e.limit));
  }


  function getRequestRange(Epoch _e, uint _blockNumber, uint _limit) internal returns (uint requestStart, uint requestEnd) {
    require(_e.isRequest);
    require(_blockNumber >= _e.startBlockNumber && _blockNumber <= _e.endBlockNumber);

    if (_blockNumber == _e.endBlockNumber) {
      requestStart = _e.requestStart + (getNumBlocks(_e) - 1) * _limit;
      requestEnd = _e.requestEnd;
      return;
    }

    requestStart = _e.requestStart + (_blockNumber - _e.startBlockNumber) * _limit;
    requestEnd = requestStart + _limit;
    return;
  }

  struct PlasmaBlock {
    bytes32 statesRoot;
    bytes32 transactionsRoot;
    bytes32 intermediateStatesRoot;
    uint64 forkNumber;
    uint64 epochNumber;
    uint64 timestamp;
    uint64 requestBlockId;    // id of RequestBlock[]
    bool isRequest;           // true in case of URB & ORB
    bool userActivated;       // true in case of URB
    bool challenged;          // true if it is challenged
    bool challenging;         // true if it is being challenged
    bool finalized;           // true if it is successfully finalized
  }

  struct Request {
    uint64 timestamp;
    bool isExit;
    bool applied;             // true if request is applied
    bool finalized;           // true if request is finalized
    bool challenged;
    address requestor;
    address to;               // requestable contract in root chain
    bytes32 trieKey;
    bytes32 trieValue;
  }

  function applyRequestInRootChain(
    Request memory self,
    uint _requestId
  )
    internal
    returns (bool)
  {
    // TODO: ignore it is reverted?
    return RequestableContractI(self.to).applyRequestInRootChain(
      self.isExit,
      _requestId,
      self.requestor,
      self.trieKey,
      self.trieValue
    );
  }

  function toChildChainRequest(
    Request memory self,
    address _to
  )
    internal
    returns (Request memory out)
  {
    out.isExit = self.isExit;
    out.requestor = self.requestor;
    out.trieKey = self.trieKey;
    out.trieValue = self.trieValue;

    out.to = _to;
  }

  /**
   * @notice return tx.data
   */
  function getData(
    Request memory self,
    uint _requestId,
    bool _rootchain
  )
    internal
    returns (bytes memory out)
  {
    bytes8 funcSig = _rootchain ? APPLY_IN_ROOTCHAIN_SIGNATURE : APPLY_IN_CHILDCHAIN_SIGNATURE;

    out = abi.encodePacked(
      funcSig,
      self.isExit,
      _requestId,
      self.requestor,
      self.trieKey,
      self.trieValue
    );
  }

  /**
   * @notice convert Request to TX
   */
  function toTX(
    Request memory self,
    uint _requestId,
    bool _rootchain
  )
    internal
    returns (TX memory out)
  {
    out.gasPrice = NA_TX_GAS_PRICE;
    out.gasLimit = uint64(NA_TX_GAS_LIMIT);
    out.data = getData(self, _requestId, _rootchain);
  }

  struct RequestBlock {
    uint64 requestStart;      // first request id
    uint64 requestEnd;        // last request id
    address trie;             // patricia tree contract address
    bytes32 transactionsRoot; // duplicated?
  }

  function init(RequestBlock storage self) internal {
    if (self.trie == address(0)) {
      self.trie = address(new PatriciaTree());
    }
  }

  function addRequest(
    RequestBlock storage self,
    Request memory _request, // should be child chain request
    uint _requestId
  ) internal {
    require(self.trie != address(0));

    bytes32 requestHash = hash(toTX(_request, _requestId, false));
    uint txIndex = self.requestStart.sub(_requestId);

    bytes memory key;
    bytes memory value;

    assembly {
      mstore(add(key, 0x20), txIndex)
      mstore(add(value, 0x20), requestHash)
    }

    PatriciaTree(self.trie).insert(key, value);
    self.transactionsRoot = PatriciaTree(self.trie).getRootHash();
  }

  /*
   * TX for Ethereum transaction
   */
  struct TX {
    uint64 nonce;
    uint256 gasPrice;
    uint64 gasLimit;
    address to;
    uint256 value;
    bytes data;
    uint256 v;
    uint256 r;
    uint256 s;
  }

  function isNATX(TX memory self) internal returns (bool) {
    return self.v == 0 && self.r == 0 && self.s == 0;
  }

  function fromBytes(bytes memory self) internal returns (TX memory out) {
    RLP.RLPItem[] memory packArr = self.toRLPItem().toList(9);

    out.nonce = uint64(packArr[0].toUint());
    out.gasPrice = packArr[1].toUint();
    out.gasLimit = uint64(packArr[2].toUint());
    out.to = packArr[3].toAddress();
    out.value = packArr[4].toUint();
    out.data = packArr[5].toBytes();
    out.v = packArr[6].toUint();
    out.r = packArr[7].toUint();
    out.s = packArr[8].toUint();
  }

  function toTX(
    uint64 _nonce,
    uint256 _gasPrice,
    uint64 _gasLimit,
    address _to,
    uint256 _value,
    bytes _data,
    uint256 _v,
    uint256 _r,
    uint256 _s
  )
    internal
    pure
    returns (TX memory out)
  {
    out.nonce = _nonce;
    out.gasPrice = _gasPrice;
    out.gasLimit = _gasLimit;
    out.to = _to;
    out.value = _value;
    out.data = _data;
    out.v = _v;
    out.r = _r;
    out.s = _s;
  }

  function hash(TX memory _tx) internal pure returns (bytes32) {
    bytes[] memory packArr = new bytes[](9);

    packArr[0] = _tx.nonce.encodeUint();
    packArr[1] = _tx.gasPrice.encodeUint();
    packArr[2] = _tx.gasLimit.encodeUint();
    packArr[3] = _tx.to.encodeAddress();
    packArr[4] = _tx.value.encodeUint();
    packArr[5] = _tx.data.encodeBytes();
    packArr[6] = _tx.v.encodeUint();
    packArr[7] = _tx.r.encodeUint();
    packArr[8] = _tx.s.encodeUint();

    return keccak256(packArr.encodeList());
  }
}