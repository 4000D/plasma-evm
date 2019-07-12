package core

import (
	"errors"
	"math/big"
	"strings"

	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/params"
)

const (
	Minute = 60 * 1
	Day    = 60 * 60 * 24

	// from contracts/stamina/contract/stamina.go
	StaminaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RECOVER_EPOCH_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"getWithdrawal\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint128\"},{\"name\":\"requestBlockNumber\",\"type\":\"uint128\"},{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"setDelegator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getTotalDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"name\":\"recoveryEpochLength\",\"type\":\"uint256\"},{\"name\":\"withdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getLastRecoveryBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getNumRecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getNumWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"recovered\",\"type\":\"bool\"}],\"name\":\"StaminaAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StaminaSubtracted\",\"type\":\"event\"}]"
)

var DefaultStamina = big.NewInt(1 * params.Ether)

var (
	errUpdateStamina = errors.New("failed to update stamina")

	StaminaContractAddressHex  = "0x000000000000000000000000000000000000dead"
	StaminaContractAddress     = common.HexToAddress(StaminaContractAddressHex)
	StaminaContractDeployedBin = "0x6080604052600436106101115763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ebb172a81146101165780631556d8ac1461013d578063158ef93e146101525780633900e4ec1461017b5780633ccfd60b1461019c5780635be4f765146101b15780637b929c271461021a57806383cd9cc31461022f578063857184d1146102505780638cd8db8a14610271578063937aaef1146102915780639b4e735f146102b2578063b69ad63b146102ef578063bcac973614610310578063c35082a914610334578063d1c0c0421461035b578063d898ae1c1461037f578063da95ebf7146103a0578063e1e158a5146103c4578063f340fa01146103d9575b600080fd5b34801561012257600080fd5b5061012b6103ed565b60408051918252519081900360200190f35b34801561014957600080fd5b5061012b6103f3565b34801561015e57600080fd5b506101676103f9565b604080519115158252519081900360200190f35b34801561018757600080fd5b5061012b600160a060020a0360043516610402565b3480156101a857600080fd5b5061016761041d565b3480156101bd57600080fd5b506101d5600160a060020a0360043516602435610633565b604080516fffffffffffffffffffffffffffffffff9586168152939094166020840152600160a060020a03909116828401521515606082015290519081900360800190f35b34801561022657600080fd5b50610167610716565b34801561023b57600080fd5b50610167600160a060020a036004351661071f565b34801561025c57600080fd5b5061012b600160a060020a03600435166107b2565b34801561027d57600080fd5b5061028f6004356024356044356107cd565b005b34801561029d57600080fd5b5061012b600160a060020a036004351661082e565b3480156102be57600080fd5b506102d3600160a060020a0360043516610849565b60408051600160a060020a039092168252519081900360200190f35b3480156102fb57600080fd5b5061012b600160a060020a0360043516610867565b34801561031c57600080fd5b50610167600160a060020a0360043516602435610882565b34801561034057600080fd5b5061012b600160a060020a0360043581169060243516610a1a565b34801561036757600080fd5b50610167600160a060020a0360043516602435610a45565b34801561038b57600080fd5b5061012b600160a060020a0360043516610ae8565b3480156103ac57600080fd5b50610167600160a060020a0360043516602435610b03565b3480156103d057600080fd5b5061012b610d24565b610167600160a060020a0360043516610d2a565b600b5481565b600a5481565b60085460ff1681565b600160a060020a031660009081526001602052604090205490565b33600090815260056020526040812080548290819081908190811061044157600080fd5b3360009081526006602052604090205493508315801561048c575084600081548110151561046b57fe5b906000526020600020906002020160010160149054906101000a900460ff16155b1561049a57600092506104c0565b8315156104b9578454600211156104b057600080fd5b600192506104c0565b8360010192505b845483106104cd57600080fd5b3360009081526005602052604090208054849081106104e857fe5b906000526020600020906002020191508160010160149054906101000a900460ff1615151561051657600080fd5b600b548254437001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff16909101111561055157600080fd5b50805460018201805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905533600081815260066020526040808220869055516fffffffffffffffffffffffffffffffff9093169283156108fc0291849190818181858888f193505050501580156105db573d6000803e3d6000fd5b50600182015460408051838152602081018690528151600160a060020a039093169233927f91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2928290030190a360019550505050505090565b600080600080610641610e73565b61064a87610ae8565b861061065557600080fd5b600160a060020a038716600090815260056020526040902080548790811061067957fe5b600091825260209182902060408051608081018252600290930290910180546fffffffffffffffffffffffffffffffff80821680865270010000000000000000000000000000000090920416948401859052600190910154600160a060020a03811692840183905260ff7401000000000000000000000000000000000000000090910416151560609093018390529a929950975095509350505050565b600c5460ff1681565b600854600090819060ff16151561073557600080fd5b50600160a060020a038281166000818152602081815260409182902080543373ffffffffffffffffffffffffffffffffffffffff19821681179092558351951680865291850152815190937f5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c69292908290030190a250600192915050565b600160a060020a031660009081526002602052604090205490565b60085460ff16156107dd57600080fd5b600083116107ea57600080fd5b600082116107f757600080fd5b6000811161080457600080fd5b60028202811161081357600080fd5b600992909255600a55600b556008805460ff19166001179055565b600160a060020a031660009081526004602052604090205490565b600160a060020a039081166000908152602081905260409020541690565b600160a060020a031660009081526007602052604090205490565b600c5460009081908190819060ff168061089a575033155b15156108a557600080fd5b600a54600160a060020a0387166000908152600460205260409020544391011161094957600160a060020a038616600081815260026020908152604080832054600180845282852091909155600483528184204390556007835281842080548201905581519384529183019190915280517f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f68099281900390910190a260019350610a11565b600160a060020a0386166000908152600260209081526040808320546001909252909120549093509150848201821061098157600080fd5b50808401828111156109ad57600160a060020a03861660009081526001602052604090208390556109c9565b600160a060020a03861660009081526001602052604090208190555b60408051868152600060208201528151600160a060020a038916927f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809928290030190a2600193505b50505092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600c54600090819060ff1680610a59575033155b1515610a6457600080fd5b50600160a060020a0383166000908152600160205260409020548281038111610a8c57600080fd5b600160a060020a0384166000818152600160209081526040918290208685039055815186815291517f66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f9281900390910190a25060019392505050565b600160a060020a031660009081526005602052604090205490565b6000806000806000806000600860009054906101000a900460ff161515610b2957600080fd5b60008811610b3657600080fd5b600160a060020a0389166000818152600260209081526040808320543384526003835281842094845293825280832054600190925282205492985096509094508511610b8157600080fd5b8786038611610b8f57600080fd5b8785038511610b9d57600080fd5b600160a060020a03891660008181526002602090815260408083208c8b0390553383526003825280832093835292905220888603905587841115610bfd57600160a060020a03891660009081526001602052604090208885039055610c17565b600160a060020a0389166000908152600160205260408120555b336000908152600560205260409020805490935091508282610c3c8260018301610e9a565b81548110610c4657fe5b60009182526020918290206002919091020180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff8b8116919091178116700100000000000000000000000000000000439283160217825560018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038e16908117909155604080518d81529485019290925283820186905290519193509133917f3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69916060908290030190a350600198975050505050505050565b60095481565b60085460009081908190819060ff161515610d4457600080fd5b600954341015610d5357600080fd5b505050600160a060020a03821660008181526002602090815260408083205433845260038352818420948452938252808320546001909252909120543483018310610d9d57600080fd5b3482018210610dab57600080fd5b3481018110610db957600080fd5b600160a060020a038516600081815260026020908152604080832034888101909155338452600383528184209484529382528083208685019055600182528083209385019093556004905220541515610e2857600160a060020a03851660009081526004602052604090204390555b604080513481529051600160a060020a0387169133917f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a79181900360200190a3506001949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b815481835581811115610ec657600202816002028360005260206000209182019101610ec69190610ecb565b505050565b610f0991905b80821115610f05576000815560018101805474ffffffffffffffffffffffffffffffffffffffffff19169055600201610ed1565b5090565b905600a165627a7a72305820b06bd3333e2e53776453551e59e68eaca02b9467d31ef476ff0cd67025335b4d0029"

	// Storage key for state variables of 'stamina' config.
	InitializedKey        = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000008")
	MinDepositKey         = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000009")
	RecoverEpochLengthKey = common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000000a")
	WithdrawalDelayKey    = common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000000b")

	delegateePosition = common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000")
	staminaPosition   = common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")

	blockchainAccount = accountWrapper{common.HexToAddress("0x00")}
	staminaAccount    = accountWrapper{StaminaContractAddress}
	staminaABI, _     = abi.JSON(strings.NewReader(StaminaABI))
)

type accountWrapper struct {
	address common.Address
}

func (a accountWrapper) Address() common.Address {
	return a.address
}

func GetDelegatee(evm *vm.EVM, from common.Address) (common.Address, error) {
	data, err := staminaABI.Pack("getDelegatee", from)
	if err != nil {
		return common.Address{}, err
	}

	ret, _, err := evm.StaticCall(blockchainAccount, StaminaContractAddress, data, 1000000)

	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(ret), nil
}

func GetStamina(evm *vm.EVM, delegatee common.Address) (*big.Int, error) {
	data, err := staminaABI.Pack("getStamina", delegatee)
	if err != nil {
		return big.NewInt(0), err
	}

	ret, _, err := evm.StaticCall(blockchainAccount, StaminaContractAddress, data, 1000000)

	if err != nil {
		return big.NewInt(0), err
	}

	stamina := new(big.Int)
	stamina.SetBytes(ret)

	return stamina, nil
}

func AddStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := staminaABI.Pack("addStamina", delegatee, gas)
	if err != nil {
		return err
	}

	res, _, err := evm.Call(blockchainAccount, StaminaContractAddress, data, 1000000, big.NewInt(0))

	if new(big.Int).SetBytes(res).Cmp(common.Big1) != 0 {
		return errUpdateStamina
	}

	return err
}

func SubtractStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := staminaABI.Pack("subtractStamina", delegatee, gas)
	if err != nil {
		return err
	}

	res, _, err := evm.Call(blockchainAccount, StaminaContractAddress, data, 1000000, big.NewInt(0))

	if new(big.Int).SetBytes(res).Cmp(common.Big1) != 0 {
		return errUpdateStamina
	}

	return err
}

// StaminaConfig are the configuration parameters of stamina.
type StaminaConfig struct {
	Initialized        bool
	MinDeposit         *big.Int
	RecoverEpochLength *big.Int
	WithdrawalDelay    *big.Int
}

// DefaultStaminaConfig contains the default configurations for the stamina.
var DefaultStaminaConfig = &StaminaConfig{
	Initialized:        true,
	MinDeposit:         big.NewInt(0.5 * params.Ether),
	RecoverEpochLength: big.NewInt(7 * Day / Minute),
	WithdrawalDelay:    big.NewInt((7 * Day / Minute) * 3),
}

func GetStaminaConfig(bc *BlockChain) *StaminaConfig {
	statedb, _ := bc.State()
	initialized := statedb.GetState(StaminaContractAddress, InitializedKey)
	minDeposit := statedb.GetState(StaminaContractAddress, MinDepositKey)
	recoverEpochLength := statedb.GetState(StaminaContractAddress, RecoverEpochLengthKey)
	withdrawalDelay := statedb.GetState(StaminaContractAddress, WithdrawalDelayKey)

	return &StaminaConfig{
		Initialized:        common.ByteToBool(initialized.Bytes()[31]),
		MinDeposit:         minDeposit.Big(),
		RecoverEpochLength: recoverEpochLength.Big(),
		WithdrawalDelay:    withdrawalDelay.Big(),
	}
}

// '000000000000000000000000' + operator address + stamina state variable position
func GetStaminaKey(operator common.Address) common.Hash {
	key := append(common.LeftPadBytes(operator.Bytes(), 32), staminaPosition...)
	return crypto.Keccak256Hash(key)
}

// '000000000000000000000000' + operator address + delegatee state variable position
func GetOperatorAsDelegatorKey(operator common.Address) common.Hash {
	key := append(common.LeftPadBytes(operator.Bytes(), 32), delegateePosition...)
	return crypto.Keccak256Hash(key)
}
