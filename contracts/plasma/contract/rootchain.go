// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/Onther-Tech/plasma-evm"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/event"
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
const AddressBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e2c15d31406954a71a6485749836de23d45c0bf8fb01c435f387d7b23a9999af0029`

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BitsABI is the input ABI used to generate the binding from.
const BitsABI = "[]"

// BitsBin is the compiled bytecode used for deploying new contracts.
const BitsBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a723058207a644ef047535603cff62c8178055146e8a5c41033b2f6e6f04d09fa09d6d5976c6578706572696d656e74616cf50037`

// DeployBits deploys a new Ethereum contract, binding an instance of Bits to it.
func DeployBits(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bits, error) {
	parsed, err := abi.JSON(strings.NewReader(BitsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BitsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bits{BitsCaller: BitsCaller{contract: contract}, BitsTransactor: BitsTransactor{contract: contract}, BitsFilterer: BitsFilterer{contract: contract}}, nil
}

// Bits is an auto generated Go binding around an Ethereum contract.
type Bits struct {
	BitsCaller     // Read-only binding to the contract
	BitsTransactor // Write-only binding to the contract
	BitsFilterer   // Log filterer for contract events
}

// BitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitsSession struct {
	Contract     *Bits             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitsCallerSession struct {
	Contract *BitsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitsTransactorSession struct {
	Contract     *BitsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitsRaw struct {
	Contract *Bits // Generic contract binding to access the raw methods on
}

// BitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitsCallerRaw struct {
	Contract *BitsCaller // Generic read-only contract binding to access the raw methods on
}

// BitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitsTransactorRaw struct {
	Contract *BitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBits creates a new instance of Bits, bound to a specific deployed contract.
func NewBits(address common.Address, backend bind.ContractBackend) (*Bits, error) {
	contract, err := bindBits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bits{BitsCaller: BitsCaller{contract: contract}, BitsTransactor: BitsTransactor{contract: contract}, BitsFilterer: BitsFilterer{contract: contract}}, nil
}

// NewBitsCaller creates a new read-only instance of Bits, bound to a specific deployed contract.
func NewBitsCaller(address common.Address, caller bind.ContractCaller) (*BitsCaller, error) {
	contract, err := bindBits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitsCaller{contract: contract}, nil
}

// NewBitsTransactor creates a new write-only instance of Bits, bound to a specific deployed contract.
func NewBitsTransactor(address common.Address, transactor bind.ContractTransactor) (*BitsTransactor, error) {
	contract, err := bindBits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitsTransactor{contract: contract}, nil
}

// NewBitsFilterer creates a new log filterer instance of Bits, bound to a specific deployed contract.
func NewBitsFilterer(address common.Address, filterer bind.ContractFilterer) (*BitsFilterer, error) {
	contract, err := bindBits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitsFilterer{contract: contract}, nil
}

// bindBits binds a generic wrapper to an already deployed contract.
func bindBits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bits *BitsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bits.Contract.BitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bits *BitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bits.Contract.BitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bits *BitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bits.Contract.BitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bits *BitsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bits *BitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bits *BitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bits.Contract.contract.Transact(opts, method, params...)
}

// DataABI is the input ABI used to generate the binding from.
const DataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_ROOTCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_CHILDCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x610205610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100835763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631927ac58811461008857806390e84f56146100a6578063a7b6ae28146100ae578063a89ca766146100c3578063ab73ff05146100cb575b600080fd5b6100906100e0565b60405161009d919061017c565b60405180910390f35b6100906100e5565b6100b66100ec565b60405161009d919061016e565b6100b6610110565b6100d3610134565b60405161009d919061015a565b600181565b620186a081565b7fd9afd3a90000000000000000000000000000000000000000000000000000000081565b7fe904e3d90000000000000000000000000000000000000000000000000000000081565b600081565b6101428161018a565b82525050565b610142816101a3565b610142816101c8565b602081016101688284610139565b92915050565b602081016101688284610148565b602081016101688284610151565b73ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b905600a265627a7a723058209a74e877106cbef7143bc700e8f8251280b0e47e5523988e43e3c3f326fcf13e6c6578706572696d656e74616cf50037`

// DeployData deploys a new Ethereum contract, binding an instance of Data to it.
func DeployData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Data, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// Data is an auto generated Go binding around an Ethereum contract.
type Data struct {
	DataCaller     // Read-only binding to the contract
	DataTransactor // Write-only binding to the contract
	DataFilterer   // Log filterer for contract events
}

// DataCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSession struct {
	Contract     *Data             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCallerSession struct {
	Contract *DataCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTransactorSession struct {
	Contract     *DataTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRaw struct {
	Contract *Data // Generic contract binding to access the raw methods on
}

// DataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCallerRaw struct {
	Contract *DataCaller // Generic read-only contract binding to access the raw methods on
}

// DataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTransactorRaw struct {
	Contract *DataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewData creates a new instance of Data, bound to a specific deployed contract.
func NewData(address common.Address, backend bind.ContractBackend) (*Data, error) {
	contract, err := bindData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// NewDataCaller creates a new read-only instance of Data, bound to a specific deployed contract.
func NewDataCaller(address common.Address, caller bind.ContractCaller) (*DataCaller, error) {
	contract, err := bindData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCaller{contract: contract}, nil
}

// NewDataTransactor creates a new write-only instance of Data, bound to a specific deployed contract.
func NewDataTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTransactor, error) {
	contract, err := bindData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTransactor{contract: contract}, nil
}

// NewDataFilterer creates a new log filterer instance of Data, bound to a specific deployed contract.
func NewDataFilterer(address common.Address, filterer bind.ContractFilterer) (*DataFilterer, error) {
	contract, err := bindData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataFilterer{contract: contract}, nil
}

// bindData binds a generic wrapper to an already deployed contract.
func bindData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.DataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.contract.Transact(opts, method, params...)
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCaller) APPLYINCHILDCHAINSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "APPLY_IN_CHILDCHAIN_SIGNATURE")
	return *ret0, err
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataSession) APPLYINCHILDCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINCHILDCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCallerSession) APPLYINCHILDCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINCHILDCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCaller) APPLYINROOTCHAINSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "APPLY_IN_ROOTCHAIN_SIGNATURE")
	return *ret0, err
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataSession) APPLYINROOTCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINROOTCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCallerSession) APPLYINROOTCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINROOTCHAINSIGNATURE(&_Data.CallOpts)
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataCaller) NA(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA")
	return *ret0, err
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataSession) NA() (common.Address, error) {
	return _Data.Contract.NA(&_Data.CallOpts)
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataCallerSession) NA() (common.Address, error) {
	return _Data.Contract.NA(&_Data.CallOpts)
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataCaller) NATXGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA_TX_GAS_LIMIT")
	return *ret0, err
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataSession) NATXGASLIMIT() (*big.Int, error) {
	return _Data.Contract.NATXGASLIMIT(&_Data.CallOpts)
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataCallerSession) NATXGASLIMIT() (*big.Int, error) {
	return _Data.Contract.NATXGASLIMIT(&_Data.CallOpts)
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataCaller) NATXGASPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA_TX_GAS_PRICE")
	return *ret0, err
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataSession) NATXGASPRICE() (*big.Int, error) {
	return _Data.Contract.NATXGASPRICE(&_Data.CallOpts)
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataCallerSession) NATXGASPRICE() (*big.Int, error) {
	return _Data.Contract.NATXGASPRICE(&_Data.CallOpts)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206821729e7a1c1a61076ce529d53dbbe3faa6286298adb57eced5b9334578de930029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// PatriciaTreeABI is the input ABI used to generate the binding from.
const PatriciaTreeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[2]\"},{\"name\":\"labelDatas\",\"type\":\"bytes32[2]\"},{\"name\":\"labelLengths\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"name\":\"h\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"labelData\",\"type\":\"bytes32\"},{\"name\":\"labelLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaTreeBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeBin = `0x608060405234801561001057600080fd5b5061124f806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c610097366004610f92565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004610e82565b61018d565b6040516100cd939291906110f5565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004610f5d565b610278565b6040516100cd929190611161565b34801561011057600080fd5b506101196104d7565b6040516100cd919061112b565b34801561013257600080fd5b5061013b6104dd565b6040516100cd93929190611139565b34801561015657600080fd5b5061016a610165366004610ea0565b610527565b6040516100cd919061111d565b6101896000838363ffffffff61067c16565b5050565b610195610cc2565b61019d610cc2565b6101a5610cc2565b6101ad610cdd565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60008054606090151561028a57600080fd5b610292610cf6565b60408051908101604052808580519060200120815260200161010081525090506102ba610d0d565b506040805180820182526001548152815180830190925260025482526003546020838101919091528101919091526102f0610d28565b6000806102fb610cf6565b610303610cf6565b600061030d610cf6565b6020880151610323908a9063ffffffff61073b16565b6020808b0151810151908301519296509094501461033d57fe5b6020830151151561034d57610454565b60208401519590950160ff81900360020a9a909a17996001019461037083610769565b8951600090815260046020526040902091935091506103d69060018490036002811061039857fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526107d0565b6001860195889061010081106103e857fe5b6020908102919091019190915288516000908152600490915260409020826002811061041057fe5b60408051808201825260039290920292909201805482528251808401909352600181015483526002015460208381019190915281019190915290985096508761030d565b60008511156104c95784604051908082528060200260200182016040528015610487578160200160208202803883390190505b50995060005b858110156104c757878161010081106104a257fe5b60200201518b828151811015156104b557fe5b6020908102909101015260010161048d565b505b505050505050505050915091565b60005490565b60008060006104ea610d0d565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610531610cf6565b6040805190810160405280878051906020012081526020016101008152509050610559610d0d565b85516020870120815260005b851561065257600061057687610812565b60ff1690508060019060020a02198716965061059e8160ff038561086690919063ffffffff16565b602085018190529094506000906105b490610769565b602086015290506105c3610cc2565b6105cc856107d0565b8183600281106105d857fe5b602002015287518890600019868203019081106105f157fe5b90602001906020020151818360010360028110151561060c57fe5b6020908102919091019190915281519181015160408051808401949094528381019190915280518084038201815260609093019052815191012084525050600101610565565b5060208101829052610663816107d0565b881461066e57600080fd5b506001979650505050505050565b610684610cf6565b506040805180820190915282516020808501919091208252610100818301528251908301206106b1610d0d565b855415156106c8576020810183905281815261070b565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261070890879085856108da565b90505b610714816107d0565b86558051600187015560209081015180516002880155015160039095019490945550505050565b610743610cf6565b61074b610cf6565b61075e846107598686610b14565b610866565b915091509250929050565b6000610773610cf6565b602083015160001061078457600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b80516020918201518083015190516040805180860194909452838101929092526060808401919091528151808403909101815260809092019052805191012090565b600081151561082057600080fd5b8160805b600160ff82161061085f5760001960ff821660020a0182161515610852579182019160ff811660020a909104905b600260ff90911604610824565b5050919050565b61086e610cf6565b610876610cf6565b8360200151831115801561088c57506101008311155b151561089457fe5b602082018390528215156108ab57600082526108bd565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b6108e2610d0d565b6020808501518101519084015110156108f757fe5b6108ff610cf6565b610907610cf6565b6000610911610cf6565b61091f87896020015161073b565b602081015191955093506000901515610939575085610aef565b6020808a01518101519086015110610a6057602084015160011061095957fe5b610961610cdd565b8951600090815260048c0160209081526040808320815160608101909252909290918391908201908390600290835b828210156109de576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610990565b505050508152505090506109f185610769565b82519195509350610a15908c908660028110610a0957fe5b6020020151858b6108da565b81518560028110610a2257fe5b602090810291909101919091528a51600090815260048d019091526040812090610a4c8282610d49565b5050610a588b82610b84565b915050610aef565b610a6984610769565b9093509150610a76610cdd565b604080518082019091528881526020810184905281518560028110610a9757fe5b602002018190525060408051908101604052808b600001518152602001610ac98c602001518960200151600101610bea565b90528151600186900360028110610adc57fe5b6020020152610aeb8b82610b84565b9150505b604080519081016040528082815260200186815250955050505050505b949350505050565b6000808260200151846020015110610b30578260200151610b36565b83602001515b9050801515610b49576000915050610b7e565b825184511861010082900360020a60000316801515610b6a57509050610b7e565b610b7381610c1e565b60ff0360ff16925050505b92915050565b600080610b9083610c6c565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610bf2610cf6565b6020830151821115610c0357600080fd5b60208084015183900390820152915160029190910a02815290565b6000811515610c2c57600080fd5b8160805b600160ff82161061085f5760ff811660020a600019810102821615610c5f579182019160ff811660020a909104905b600260ff90911604610c30565b8051600090610c8190825b60200201516107d0565b8251610c8e906001610c77565b6040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610cf1610d73565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610cf1610cf6565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b610d8b610d0d565b815260200190600190039081610d835790505090565b6000601f82018313610db257600080fd5b8135610dc5610dc0826111a8565b611181565b91508181835260208401935060208101905083856020840282011115610dea57600080fd5b60005b83811015610e165781610e008882610e20565b8452506020928301929190910190600101610ded565b5050505092915050565b6000610e2c82356111f1565b9392505050565b6000601f82018313610e4457600080fd5b8135610e52610dc0826111c9565b91508082526020830160208301858383011115610e6e57600080fd5b610e79838284611209565b50505092915050565b600060208284031215610e9457600080fd5b6000610b0c8484610e20565b600080600080600060a08688031215610eb857600080fd5b6000610ec48888610e20565b955050602086013567ffffffffffffffff811115610ee157600080fd5b610eed88828901610e33565b945050604086013567ffffffffffffffff811115610f0a57600080fd5b610f1688828901610e33565b9350506060610f2788828901610e20565b925050608086013567ffffffffffffffff811115610f4457600080fd5b610f5088828901610da1565b9150509295509295909350565b600060208284031215610f6f57600080fd5b813567ffffffffffffffff811115610f8657600080fd5b610b0c84828501610e33565b60008060408385031215610fa557600080fd5b823567ffffffffffffffff811115610fbc57600080fd5b610fc885828601610e33565b925050602083013567ffffffffffffffff811115610fe557600080fd5b610ff185828601610e33565b9150509250929050565b611004816111fa565b61100d826111f1565b60005b8281101561103d576110238583516110ec565b61102c826111f4565b602095909501949150600101611010565b5050505050565b600061104f82611200565b808452602084019350611061836111f4565b60005b82811015611091576110778683516110ec565b611080826111f4565b602096909601959150600101611064565b5093949350505050565b6110a4816111fa565b6110ad826111f1565b60005b8281101561103d576110c38583516110ec565b6110cc826111f4565b6020959095019491506001016110b0565b6110e681611204565b82525050565b6110e6816111f1565b60c081016111038286610ffb565b6111106040830185610ffb565b610b0c608083018461109b565b60208101610b7e82846110dd565b60208101610b7e82846110ec565b6060810161114782866110ec565b61115460208301856110ec565b610b0c60408301846110ec565b6040810161116f82856110ec565b8181036020830152610b0c8184611044565b60405181810167ffffffffffffffff811182821017156111a057600080fd5b604052919050565b600067ffffffffffffffff8211156111bf57600080fd5b5060209081020190565b600067ffffffffffffffff8211156111e057600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a7230582018b914ecdc6c0bae868236caa7c87914f895736b958199d65b8e37a35f3a8b3c6c6578706572696d656e74616cf50037`

// DeployPatriciaTree deploys a new Ethereum contract, binding an instance of PatriciaTree to it.
func DeployPatriciaTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTree, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTree{PatriciaTreeCaller: PatriciaTreeCaller{contract: contract}, PatriciaTreeTransactor: PatriciaTreeTransactor{contract: contract}, PatriciaTreeFilterer: PatriciaTreeFilterer{contract: contract}}, nil
}

// PatriciaTree is an auto generated Go binding around an Ethereum contract.
type PatriciaTree struct {
	PatriciaTreeCaller     // Read-only binding to the contract
	PatriciaTreeTransactor // Write-only binding to the contract
	PatriciaTreeFilterer   // Log filterer for contract events
}

// PatriciaTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeSession struct {
	Contract     *PatriciaTree     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeCallerSession struct {
	Contract *PatriciaTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PatriciaTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeTransactorSession struct {
	Contract     *PatriciaTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PatriciaTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeRaw struct {
	Contract *PatriciaTree // Generic contract binding to access the raw methods on
}

// PatriciaTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeCallerRaw struct {
	Contract *PatriciaTreeCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeTransactorRaw struct {
	Contract *PatriciaTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTree creates a new instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTree(address common.Address, backend bind.ContractBackend) (*PatriciaTree, error) {
	contract, err := bindPatriciaTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTree{PatriciaTreeCaller: PatriciaTreeCaller{contract: contract}, PatriciaTreeTransactor: PatriciaTreeTransactor{contract: contract}, PatriciaTreeFilterer: PatriciaTreeFilterer{contract: contract}}, nil
}

// NewPatriciaTreeCaller creates a new read-only instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeCaller, error) {
	contract, err := bindPatriciaTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeCaller{contract: contract}, nil
}

// NewPatriciaTreeTransactor creates a new write-only instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeTransactor, error) {
	contract, err := bindPatriciaTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeTransactor{contract: contract}, nil
}

// NewPatriciaTreeFilterer creates a new log filterer instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeFilterer, error) {
	contract, err := bindPatriciaTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFilterer{contract: contract}, nil
}

// bindPatriciaTree binds a generic wrapper to an already deployed contract.
func bindPatriciaTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTree *PatriciaTreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTree.Contract.PatriciaTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTree *PatriciaTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTree.Contract.PatriciaTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTree *PatriciaTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTree.Contract.PatriciaTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTree *PatriciaTreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTree *PatriciaTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTree *PatriciaTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTree.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeCaller) GetNode(opts *bind.CallOpts, hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	ret := new(struct {
		Nodes        [2][32]byte
		LabelDatas   [2][32]byte
		LabelLengths [2]*big.Int
	})
	out := ret
	err := _PatriciaTree.contract.Call(opts, out, "getNode", hash)
	return *ret, err
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTree.Contract.GetNode(&_PatriciaTree.CallOpts, hash)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeCallerSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTree.Contract.GetNode(&_PatriciaTree.CallOpts, hash)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeCaller) GetProof(opts *bind.CallOpts, key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	ret := new(struct {
		BranchMask *big.Int
		Siblings   [][32]byte
	})
	out := ret
	err := _PatriciaTree.contract.Call(opts, out, "getProof", key)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTree.Contract.GetProof(&_PatriciaTree.CallOpts, key)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeCallerSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTree.Contract.GetProof(&_PatriciaTree.CallOpts, key)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeCaller) GetRootEdge(opts *bind.CallOpts) (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	ret := new(struct {
		Node        [32]byte
		LabelData   [32]byte
		LabelLength *big.Int
	})
	out := ret
	err := _PatriciaTree.contract.Call(opts, out, "getRootEdge")
	return *ret, err
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTree.Contract.GetRootEdge(&_PatriciaTree.CallOpts)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeCallerSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTree.Contract.GetRootEdge(&_PatriciaTree.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeCaller) GetRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PatriciaTree.contract.Call(opts, out, "getRootHash")
	return *ret0, err
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTree.Contract.GetRootHash(&_PatriciaTree.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeCallerSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTree.Contract.GetRootHash(&_PatriciaTree.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeCaller) VerifyProof(opts *bind.CallOpts, rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PatriciaTree.contract.Call(opts, out, "verifyProof", rootHash, key, value, branchMask, siblings)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTree.Contract.VerifyProof(&_PatriciaTree.CallOpts, rootHash, key, value, branchMask, siblings)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeCallerSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTree.Contract.VerifyProof(&_PatriciaTree.CallOpts, rootHash, key, value, branchMask, siblings)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeTransactor) Insert(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.contract.Transact(opts, "insert", key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.Contract.Insert(&_PatriciaTree.TransactOpts, key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeTransactorSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.Contract.Insert(&_PatriciaTree.TransactOpts, key, value)
}

// PatriciaTreeDataABI is the input ABI used to generate the binding from.
const PatriciaTreeDataABI = "[]"

// PatriciaTreeDataBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeDataBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a72305820c1bcc829b65d2265560e832a095ceb3b9129c0717854dd5a99bc47e5216bd9526c6578706572696d656e74616cf50037`

// DeployPatriciaTreeData deploys a new Ethereum contract, binding an instance of PatriciaTreeData to it.
func DeployPatriciaTreeData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTreeData, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTreeData{PatriciaTreeDataCaller: PatriciaTreeDataCaller{contract: contract}, PatriciaTreeDataTransactor: PatriciaTreeDataTransactor{contract: contract}, PatriciaTreeDataFilterer: PatriciaTreeDataFilterer{contract: contract}}, nil
}

// PatriciaTreeData is an auto generated Go binding around an Ethereum contract.
type PatriciaTreeData struct {
	PatriciaTreeDataCaller     // Read-only binding to the contract
	PatriciaTreeDataTransactor // Write-only binding to the contract
	PatriciaTreeDataFilterer   // Log filterer for contract events
}

// PatriciaTreeDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeDataSession struct {
	Contract     *PatriciaTreeData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeDataCallerSession struct {
	Contract *PatriciaTreeDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PatriciaTreeDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeDataTransactorSession struct {
	Contract     *PatriciaTreeDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PatriciaTreeDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeDataRaw struct {
	Contract *PatriciaTreeData // Generic contract binding to access the raw methods on
}

// PatriciaTreeDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeDataCallerRaw struct {
	Contract *PatriciaTreeDataCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeDataTransactorRaw struct {
	Contract *PatriciaTreeDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTreeData creates a new instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeData(address common.Address, backend bind.ContractBackend) (*PatriciaTreeData, error) {
	contract, err := bindPatriciaTreeData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeData{PatriciaTreeDataCaller: PatriciaTreeDataCaller{contract: contract}, PatriciaTreeDataTransactor: PatriciaTreeDataTransactor{contract: contract}, PatriciaTreeDataFilterer: PatriciaTreeDataFilterer{contract: contract}}, nil
}

// NewPatriciaTreeDataCaller creates a new read-only instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeDataCaller, error) {
	contract, err := bindPatriciaTreeData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataCaller{contract: contract}, nil
}

// NewPatriciaTreeDataTransactor creates a new write-only instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeDataTransactor, error) {
	contract, err := bindPatriciaTreeData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataTransactor{contract: contract}, nil
}

// NewPatriciaTreeDataFilterer creates a new log filterer instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeDataFilterer, error) {
	contract, err := bindPatriciaTreeData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataFilterer{contract: contract}, nil
}

// bindPatriciaTreeData binds a generic wrapper to an already deployed contract.
func bindPatriciaTreeData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeData.Contract.PatriciaTreeDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.PatriciaTreeDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.PatriciaTreeDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeData *PatriciaTreeDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeData *PatriciaTreeDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeData *PatriciaTreeDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.contract.Transact(opts, method, params...)
}

// PatriciaTreeFaceABI is the input ABI used to generate the binding from.
const PatriciaTreeFaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[2]\"},{\"name\":\"labelDatas\",\"type\":\"bytes32[2]\"},{\"name\":\"labelLengths\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"name\":\"h\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"labelData\",\"type\":\"bytes32\"},{\"name\":\"labelLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaTreeFaceBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeFaceBin = `0x`

// DeployPatriciaTreeFace deploys a new Ethereum contract, binding an instance of PatriciaTreeFace to it.
func DeployPatriciaTreeFace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTreeFace, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeFaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeFaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTreeFace{PatriciaTreeFaceCaller: PatriciaTreeFaceCaller{contract: contract}, PatriciaTreeFaceTransactor: PatriciaTreeFaceTransactor{contract: contract}, PatriciaTreeFaceFilterer: PatriciaTreeFaceFilterer{contract: contract}}, nil
}

// PatriciaTreeFace is an auto generated Go binding around an Ethereum contract.
type PatriciaTreeFace struct {
	PatriciaTreeFaceCaller     // Read-only binding to the contract
	PatriciaTreeFaceTransactor // Write-only binding to the contract
	PatriciaTreeFaceFilterer   // Log filterer for contract events
}

// PatriciaTreeFaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeFaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeFaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeFaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeFaceSession struct {
	Contract     *PatriciaTreeFace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeFaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeFaceCallerSession struct {
	Contract *PatriciaTreeFaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PatriciaTreeFaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeFaceTransactorSession struct {
	Contract     *PatriciaTreeFaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PatriciaTreeFaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeFaceRaw struct {
	Contract *PatriciaTreeFace // Generic contract binding to access the raw methods on
}

// PatriciaTreeFaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeFaceCallerRaw struct {
	Contract *PatriciaTreeFaceCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeFaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeFaceTransactorRaw struct {
	Contract *PatriciaTreeFaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTreeFace creates a new instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFace(address common.Address, backend bind.ContractBackend) (*PatriciaTreeFace, error) {
	contract, err := bindPatriciaTreeFace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFace{PatriciaTreeFaceCaller: PatriciaTreeFaceCaller{contract: contract}, PatriciaTreeFaceTransactor: PatriciaTreeFaceTransactor{contract: contract}, PatriciaTreeFaceFilterer: PatriciaTreeFaceFilterer{contract: contract}}, nil
}

// NewPatriciaTreeFaceCaller creates a new read-only instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeFaceCaller, error) {
	contract, err := bindPatriciaTreeFace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceCaller{contract: contract}, nil
}

// NewPatriciaTreeFaceTransactor creates a new write-only instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeFaceTransactor, error) {
	contract, err := bindPatriciaTreeFace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceTransactor{contract: contract}, nil
}

// NewPatriciaTreeFaceFilterer creates a new log filterer instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeFaceFilterer, error) {
	contract, err := bindPatriciaTreeFace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceFilterer{contract: contract}, nil
}

// bindPatriciaTreeFace binds a generic wrapper to an already deployed contract.
func bindPatriciaTreeFace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeFaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeFace *PatriciaTreeFaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeFace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetNode(opts *bind.CallOpts, hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	ret := new(struct {
		Nodes        [2][32]byte
		LabelDatas   [2][32]byte
		LabelLengths [2]*big.Int
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getNode", hash)
	return *ret, err
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetNode(&_PatriciaTreeFace.CallOpts, hash)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetNode(&_PatriciaTreeFace.CallOpts, hash)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetProof(opts *bind.CallOpts, key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	ret := new(struct {
		BranchMask *big.Int
		Siblings   [][32]byte
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getProof", key)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTreeFace.Contract.GetProof(&_PatriciaTreeFace.CallOpts, key)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTreeFace.Contract.GetProof(&_PatriciaTreeFace.CallOpts, key)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetRootEdge(opts *bind.CallOpts) (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	ret := new(struct {
		Node        [32]byte
		LabelData   [32]byte
		LabelLength *big.Int
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getRootEdge")
	return *ret, err
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetRootEdge(&_PatriciaTreeFace.CallOpts)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetRootEdge(&_PatriciaTreeFace.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PatriciaTreeFace.contract.Call(opts, out, "getRootHash")
	return *ret0, err
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTreeFace.Contract.GetRootHash(&_PatriciaTreeFace.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTreeFace.Contract.GetRootHash(&_PatriciaTreeFace.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) VerifyProof(opts *bind.CallOpts, rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PatriciaTreeFace.contract.Call(opts, out, "verifyProof", rootHash, key, value, branchMask, siblings)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTreeFace.Contract.VerifyProof(&_PatriciaTreeFace.CallOpts, rootHash, key, value, branchMask, siblings)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTreeFace.Contract.VerifyProof(&_PatriciaTreeFace.CallOpts, rootHash, key, value, branchMask, siblings)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceTransactor) Insert(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.contract.Transact(opts, "insert", key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.Insert(&_PatriciaTreeFace.TransactOpts, key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.Insert(&_PatriciaTreeFace.TransactOpts, key, value)
}

// RLPABI is the input ABI used to generate the binding from.
const RLPABI = "[]"

// RLPBin is the compiled bytecode used for deploying new contracts.
const RLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820230548a37307bfa6c80e09c03d291d975ca5f35fca5a2faeb37542bf111860bd0029`

// DeployRLP deploys a new Ethereum contract, binding an instance of RLP to it.
func DeployRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLP, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
	RLPFilterer   // Log filterer for contract events
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPSession struct {
	Contract     *RLP              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPCallerSession struct {
	Contract *RLPCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPTransactorSession struct {
	Contract     *RLPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPRaw struct {
	Contract *RLP // Generic contract binding to access the raw methods on
}

// RLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPCallerRaw struct {
	Contract *RLPCaller // Generic read-only contract binding to access the raw methods on
}

// RLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPTransactorRaw struct {
	Contract *RLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLP creates a new instance of RLP, bound to a specific deployed contract.
func NewRLP(address common.Address, backend bind.ContractBackend) (*RLP, error) {
	contract, err := bindRLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// NewRLPFilterer creates a new log filterer instance of RLP, bound to a specific deployed contract.
func NewRLPFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPFilterer, error) {
	contract, err := bindRLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPFilterer{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.RLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
const RLPEncodeBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f2ceb9bea42b35ab26e1fabebadc09fd1ed9cd72d112da05517f68206ee7825e0029`

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// RequestableContractIABI is the input ABI used to generate the binding from.
const RequestableContractIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RequestableContractIBin is the compiled bytecode used for deploying new contracts.
const RequestableContractIBin = `0x`

// DeployRequestableContractI deploys a new Ethereum contract, binding an instance of RequestableContractI to it.
func DeployRequestableContractI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestableContractI, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableContractIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestableContractIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestableContractI{RequestableContractICaller: RequestableContractICaller{contract: contract}, RequestableContractITransactor: RequestableContractITransactor{contract: contract}, RequestableContractIFilterer: RequestableContractIFilterer{contract: contract}}, nil
}

// RequestableContractI is an auto generated Go binding around an Ethereum contract.
type RequestableContractI struct {
	RequestableContractICaller     // Read-only binding to the contract
	RequestableContractITransactor // Write-only binding to the contract
	RequestableContractIFilterer   // Log filterer for contract events
}

// RequestableContractICaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableContractICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractITransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableContractITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableContractIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableContractISession struct {
	Contract     *RequestableContractI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RequestableContractICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableContractICallerSession struct {
	Contract *RequestableContractICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RequestableContractITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableContractITransactorSession struct {
	Contract     *RequestableContractITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RequestableContractIRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableContractIRaw struct {
	Contract *RequestableContractI // Generic contract binding to access the raw methods on
}

// RequestableContractICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableContractICallerRaw struct {
	Contract *RequestableContractICaller // Generic read-only contract binding to access the raw methods on
}

// RequestableContractITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableContractITransactorRaw struct {
	Contract *RequestableContractITransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableContractI creates a new instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractI(address common.Address, backend bind.ContractBackend) (*RequestableContractI, error) {
	contract, err := bindRequestableContractI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableContractI{RequestableContractICaller: RequestableContractICaller{contract: contract}, RequestableContractITransactor: RequestableContractITransactor{contract: contract}, RequestableContractIFilterer: RequestableContractIFilterer{contract: contract}}, nil
}

// NewRequestableContractICaller creates a new read-only instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractICaller(address common.Address, caller bind.ContractCaller) (*RequestableContractICaller, error) {
	contract, err := bindRequestableContractI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableContractICaller{contract: contract}, nil
}

// NewRequestableContractITransactor creates a new write-only instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractITransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableContractITransactor, error) {
	contract, err := bindRequestableContractI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableContractITransactor{contract: contract}, nil
}

// NewRequestableContractIFilterer creates a new log filterer instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractIFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableContractIFilterer, error) {
	contract, err := bindRequestableContractI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableContractIFilterer{contract: contract}, nil
}

// bindRequestableContractI binds a generic wrapper to an already deployed contract.
func bindRequestableContractI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableContractIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableContractI *RequestableContractIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableContractI.Contract.RequestableContractICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableContractI *RequestableContractIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableContractI.Contract.RequestableContractITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableContractI *RequestableContractIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableContractI.Contract.RequestableContractITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableContractI *RequestableContractICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableContractI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableContractI *RequestableContractITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableContractI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableContractI *RequestableContractITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableContractI.Contract.contract.Transact(opts, method, params...)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractISession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInChildChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInChildChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractISession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInRootChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInRootChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"forkedBlockNumber\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startEnter\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRBEpochLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"applyRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_txByte\",\"type\":\"bytes\"},{\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"revertBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastFinalizedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORBEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestApplied\",\"outputs\":[{\"name\":\"applied\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"makeERU\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"applied\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"},{\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumORBs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"mapRequestableContract\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"applied\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"forkNumber\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"intermediateStatesRoot\",\"type\":\"bytes32\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_NRBEpochLength\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"StateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NRBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ORBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"URBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"firstBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x60806040526000805460ff191660011790553480156200001e57600080fd5b5060405160a08062006375833981018060405260a08110156200004057600080fd5b50805160208083015160408085015160608601516080909601516000805460ff19168715151761010060a860020a031916610100338102919091178255600986905560018054835260078852858320838052885294822094850184905560028501899055600385018390557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb588054600160c060020a03167801000000000000000000000000000000000000000000000000426001604060020a0316021790557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb5980547f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c790985264010000000061ffff1990981690911764ff000000001916871790559596939591949092916200017c918391906200019b810204565b6200018f64010000000062000208810204565b50505050505062000456565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b600480546001908101918290558054600090815260086020908152604080832085845290915290209181146200028957506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0378010000000000000000000000000000000000000000000000009091048116909101165b60028201805461010061ff00199091161790558154608060020a60c060020a0319167001000000000000000000000000000000006001604060020a038381169190910291909117808455600954600160c060020a0391821678010000000000000000000000000000000000000000000000009185016000190184168202178555600185018054909216429093160291909117905560008054819060a860020a60ff02191675010000000000000000000000000000000000000000008202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6917501000000000000000000000000000000000000000000900460ff1690808260028111156200039d57fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a037001000000000000000000000000000000008304811660208501527801000000000000000000000000000000000000000000000000909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b615f0f80620004666000396000f3006080604052600436106102715763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed811461027657806308c4fff01461029d5780630a88bd04146102b2578063164bc2ae14610363578063183d2d1c14610378578063192adc5b1461038d5780631f261d59146103a257806324839704146103b757806325119293146103fd5780632e7ab9481461042f578063570ca73514610459578063584e349e1461048a57806361d29e831461049f5780636299fb24146104b457806365d724bc146105e257806375395a58146105f7578063760247921461060c578063766718081461063c57806376e61c36146106515780637b929c271461067b57806383bcbcff146106905780638b5172d0146106ba5780638eb288ca146106cf57806393248d02146106e4578063935212221461070d57806394be3aa51461027657806397be455d1461072257806399bd36001461074b578063ae08ffe914610760578063b17fa6e914610792578063b2ae9ba814610276578063b41e69dd146107a7578063b443f3cc146107e6578063b540adba1461087e578063c0e8606414610893578063c19d93fb14610909578063c2bc88fa14610942578063cb5d742f14610957578063ce8a2bc214610992578063d691acd814610276578063da0185f8146109bc578063de0ce17d146109ef578063e67123c414610a04578063e6925d0814610a2d578063ea0c73f614610a42578063ea7f22a814610a57578063f1f3c46c14610a81578063f28a7afa14610aab578063f3aba34414610add578063f4f31de414610b10578063f4f911db14610b3a578063fb788a2714610be1575b600080fd5b34801561028257600080fd5b5061028b610bf6565b60408051918252519081900360200190f35b3480156102a957600080fd5b5061028b610c02565b3480156102be57600080fd5b506102e2600480360360408110156102d557600080fd5b5080359060200135610c07565b604080516001604060020a039e8f1681529c8e1660208e01529a8d168c8c0152988c1660608c0152968b1660808b0152948a1660a08a015292891660c0890152971660e08701529515156101008601529415156101208501529315156101408401529215156101608301529115156101808201529051908190036101a00190f35b34801561036f57600080fd5b5061028b610ca0565b34801561038457600080fd5b5061028b610ca6565b34801561039957600080fd5b5061028b610cac565b3480156103ae57600080fd5b5061028b610cb8565b6103e9600480360360608110156103cd57600080fd5b50600160a060020a038135169060208101359060400135610cbe565b604080519115158252519081900360200190f35b6103e96004803603606081101561041357600080fd5b50600160a060020a038135169060208101359060400135610d6b565b34801561043b57600080fd5b5061028b6004803603602081101561045257600080fd5b5035610e00565b34801561046557600080fd5b5061046e610e12565b60408051600160a060020a039092168252519081900360200190f35b34801561049657600080fd5b5061028b610e26565b3480156104ab57600080fd5b506103e9610e2c565b3480156104c057600080fd5b506105e0600480360360a08110156104d757600080fd5b813591908101906040810160208201356401000000008111156104f957600080fd5b82018360208201111561050b57600080fd5b8035906020019184600183028401116401000000008311171561052d57600080fd5b91939092909160208101903564010000000081111561054b57600080fd5b82018360208201111561055d57600080fd5b8035906020019184600183028401116401000000008311171561057f57600080fd5b919390928235926040810190602001356401000000008111156105a157600080fd5b8201836020820111156105b357600080fd5b803590602001918460208302840111640100000000831117156105d557600080fd5b50909250905061142d565b005b3480156105ee57600080fd5b5061028b6116b7565b34801561060357600080fd5b506103e96116bd565b34801561061857600080fd5b506105e06004803603604081101561062f57600080fd5b50803590602001356116d8565b34801561064857600080fd5b5061028b6116dc565b34801561065d57600080fd5b5061028b6004803603602081101561067457600080fd5b50356116e2565b34801561068757600080fd5b506103e96116f4565b34801561069c57600080fd5b5061028b600480360360208110156106b357600080fd5b50356116fd565b3480156106c657600080fd5b5061028b61170f565b3480156106db57600080fd5b5061028b61171b565b6103e9600480360360608110156106fa57600080fd5b5080359060208101359060400135611722565b34801561071957600080fd5b5061028b61183b565b6103e96004803603606081101561073857600080fd5b5080359060208101359060400135611841565b34801561075757600080fd5b506103e9611a4e565b34801561076c57600080fd5b506103e96004803603604081101561078357600080fd5b50803590602001351515611a53565b34801561079e57600080fd5b5061028b611aa0565b3480156107b357600080fd5b506103e9600480360360608110156107ca57600080fd5b50600160a060020a038135169060208101359060400135611aa5565b3480156107f257600080fd5b506108106004803603602081101561080957600080fd5b5035611b51565b604080516001604060020a03909b168b5298151560208b015296151589890152941515606089015292151560808801526001608060020a0390911660a0870152600160a060020a0390811660c08701521660e085015261010084015261012083015251908190036101400190f35b34801561088a57600080fd5b5061028b611bf3565b34801561089f57600080fd5b506108bd600480360360208110156108b657600080fd5b5035611bf9565b6040805196151587526001604060020a03958616602088015293851686850152919093166060850152600160a060020a03909216608084015260a0830191909152519081900360c00190f35b34801561091557600080fd5b5061091e611c6c565b6040518082600281111561092e57fe5b60ff16815260200191505060405180910390f35b34801561094e57600080fd5b5061028b611c7c565b34801561096357600080fd5b506103e96004803603604081101561097a57600080fd5b50600160a060020a0381358116916020013516611c82565b34801561099e57600080fd5b506103e9600480360360208110156109b557600080fd5b5035611d20565b3480156109c857600080fd5b5061046e600480360360208110156109df57600080fd5b5035600160a060020a0316611d2c565b3480156109fb57600080fd5b5061046e611d47565b6103e960048036036060811015610a1a57600080fd5b5080359060208101359060400135611d4c565b348015610a3957600080fd5b506103e9611ec2565b348015610a4e57600080fd5b5061028b611f4d565b348015610a6357600080fd5b506108bd60048036036020811015610a7a57600080fd5b5035611f53565b348015610a8d57600080fd5b5061028b60048036036020811015610aa457600080fd5b5035611f61565b348015610ab757600080fd5b506103e960048036036040811015610ace57600080fd5b50803590602001351515611f73565b348015610ae957600080fd5b506103e960048036036020811015610b0057600080fd5b5035600160a060020a0316611fc5565b348015610b1c57600080fd5b5061081060048036036020811015610b3357600080fd5b503561203d565b348015610b4657600080fd5b50610b6a60048036036040811015610b5d57600080fd5b508035906020013561204b565b604080516001604060020a039d8e1681529b8d1660208d0152998c168b8b015297909a1660608a0152608089019590955260a088019390935260c0870191909152151560e0860152151561010085015215156101208401529215156101408301529115156101608201529051908190036101800190f35b348015610bed57600080fd5b5061028b6120d7565b67016345785d8a000081565b600181565b60086020908152600092835260408084209091529082529020805460018201546002909201546001604060020a0380831693604060020a808504831694608060020a80820485169560c060020a92839004861695858116959485048116949283048116939092049091169060ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168d565b600e5481565b60015481565b670c7d713b49da000081565b60115481565b600067016345785d8a000034811115610cd657600080fd5b600080610cec600a600c898989600160006120dd565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a08101899052600160c0820152600060e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a15060019695505050505050565b6000806000610d82600a600c8888886000806120dd565b60408051838152336020820152600160a060020a038a1681830152606081018390526080810189905260a08101889052600060c0820181905260e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a150600195945050505050565b60026020526000908152604090205481565b6000546101009004600160a060020a031681565b60095481565b600e54600f5460008281526005602052604081205490929183918290821115610e5457600080fd5b60008481526007602090815260408083208584528252808320805488855260088452828520604060020a9091046001604060020a03168086529352922060048301549195509060ff161515610f51575b600086815260086020908152604080832088845290915290206002015462010000900460ff161515610f0b576000868152600860209081526040808320888452909152902060020154610100900460ff161515610f0057600080fd5b846001019450610ea4565b546000868152600860209081526040808320888452825280832089845260078352818420608060020a9095046001604060020a0316808552949092529091209194509091505b6004820154640100000000900460ff161515610f6c57600080fd5b6004820154610100900460ff16156111d657601154600b549093508310610f9257600080fd5b6000600b84815481101515610fa357fe5b906000526020600020906005020190506000600d8460000160109054906101000a90046001604060020a03166001604060020a0316815481101515610fe457fe5b6000918252602090912060039091020180549091507101000000000000000000000000000000000090046001604060020a031685141561106557600183015460006001604060020a03909116118015611050575060018301546000196001604060020a03918216011686145b1561105d5760018801600e555b60018601600f555b600185016011558154604060020a900460ff16801561108d57508154604860020a900460ff16155b1561116b578154604860020a69ff00000000000000000019909116811780845560408051610140810182526001604060020a038316815260ff604060020a840481161515602083015293830484161515918101919091526a010000000000000000000082048316151560608201526b010000000000000000000000820490921615156080830152606060020a90046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e082015260038301546101008201526004830154610120820152611169908661268b565b505b81546aff0000000000000000000019166a0100000000000000000000178255604080518681526001602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a160019850505050505050505061142a565b601054600a5490935083106111ea57600080fd5b6000600a848154811015156111fb57fe5b906000526020600020906005020190506000600c8460000160109054906101000a90046001604060020a03166001604060020a031681548110151561123c57fe5b6000918252602090912060039091020180549091507101000000000000000000000000000000000090046001604060020a03168514156112bd57600183015460006001604060020a039091161180156112a8575060018301546000196001604060020a03918216011686145b156112b55760018801600e555b60018601600f555b600185016010558154604060020a900460ff1680156112e557508154604860020a900460ff16155b156113c3578154604860020a69ff00000000000000000019909116811780845560408051610140810182526001604060020a038316815260ff604060020a840481161515602083015293830484161515918101919091526a010000000000000000000082048316151560608201526b010000000000000000000000820490921615156080830152606060020a90046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e0820152600383015461010082015260048301546101208201526113c1908661268b565b505b81546aff0000000000000000000019166a0100000000000000000000178255604080518681526000602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a16001985050505050505050505b90565b60015460009081526007602090815260408083208b84529091529020600481015460ff161561145b57600080fd5b8054426001604060020a0360c060020a909204919091166001011161147f57600080fd5b6004810154600090610100900460ff16156114d9578154600d80549091608060020a90046001604060020a03169081106114b557fe5b6000918252602090912060016003909202010154600160a060020a0316905061151a565b8154600c80549091608060020a90046001604060020a03169081106114fa57fe5b6000918252602090912060016003909202010154600160a060020a031690505b611522614a48565b61156188888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506127ce92505050565b905061156c81612917565b151561157757600080fd5b81600160a060020a031663f7e498f684600201548c8c8c8c8c8c8c6040518963ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018089815260200180602001806020018681526020018060200184810384528b8b82818152602001925080828437600083820152601f01601f191690910185810384528981526020019050898980828437600083820152601f01601f19169091018581038352868152602090810191508790870280828437600081840152601f19601f8201169050808301925050509b50505050505050505050505060206040518083038186803b15801561167357600080fd5b505afa158015611687573d6000803e3d6000fd5b505050506040513d602081101561169d57600080fd5b505115156116aa57600080fd5b5050505050505050505050565b60105481565b60006116c7612940565b15156116d257600080fd5b50600190565b5050565b60045481565b60066020526000908152604090205481565b60005460ff1681565b60036020526000908152604090205481565b6702c68af0bb14000081565b620186a081565b600080546101009004600160a060020a0316331461173f57600080fd5b60008060005460a860020a900460ff16600281111561175a57fe5b1461176457600080fd5b67016345785d8a00003481111561177a57600080fd5b611782612940565b50600154600090815260086020908152604080832060045484529091529020600281015462010000900460ff16156117b957600080fd5b60006117cb8888886000806000612ab2565b600154604080519182526020820183905280519293507fd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f192918290030190a1815460c060020a90046001604060020a031681141561182b5761182b612d06565b6001945050505b50509392505050565b6103e881565b600060028060005460a860020a900460ff16600281111561185e57fe5b1461186857600080fd5b670c7d713b49da00003481111561187e57600080fd5b6001805460009081526007602090815260408083206005835281842054845290915281206004015460ff1615916118bc908990899089908087612ab2565b90508015611a40576001805460009081526008602090815260408083206004548452825280832081516101a08101835281546001604060020a038082168352604060020a808304821696840196909652608060020a80830482169584019590955260c060020a9182900481166060840152968301548088166080840152948504871660a0830152928404861660c08201529190920490931660e0840152600281015460ff808216151561010080870191909152820481161515610120860152620100008204811615156101408601526301000000820481161515610160860152640100000000909104161515610180840152916119b89061317e565b825460018054600090815260056020526040902054929350608060020a9091046001604060020a031690910301808214156119f5576119f56131b1565b600154604080519182526020820186905280517f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d59281900390910190a1600197505050505050611832565b506000979650505050505050565b600190565b60008115611a6e57600b805484908110611a6957fe5b506000525b600a805484908110611a7c57fe5b6000918252602090912060059091020154604860020a900460ff1690505b92915050565b600381565b60006702c68af0bb14000034811115611abd57600080fd5b600080611ad2600b600d8989896001806120dd565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a08101899052600160c0820181905260e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a15060019695505050505050565b600a805482908110611b5f57fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001604060020a0384169550604060020a840460ff90811695604860020a86048216956a010000000000000000000081048316956b010000000000000000000000820490931694606060020a9091046001608060020a031693600160a060020a039384169390911691908a565b600a5490565b600d805482908110611c0757fe5b600091825260209091206003909102018054600182015460029092015460ff821693506001604060020a036101008304811693604860020a840482169371010000000000000000000000000000000000900490911691600160a060020a039091169086565b60005460a860020a900460ff1681565b610e1081565b600080546101009004600160a060020a03163314611c9f57600080fd5b611cb183600160a060020a031661339d565b1515611cbc57600080fd5b600160a060020a038381166000908152601260205260409020541615611ce157600080fd5b50600160a060020a039182166000908152601260205260409020805473ffffffffffffffffffffffffffffffffffffffff191691909216179055600190565b6001548114155b919050565b601260205260009081526040902054600160a060020a031681565b600081565b600080546101009004600160a060020a03163314611d6957600080fd5b60018060005460a860020a900460ff166002811115611d8457fe5b14611d8e57600080fd5b67016345785d8a000034811115611da457600080fd5b611dac612940565b50600154600090815260086020908152604080832060045484529091529020600281015462010000900460ff161515611de457600080fd5b6000611df68888886001600080612ab2565b600154604080519182526020820183905280519293507f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d592918290030190a160005460ff161515611ea1576001546000908152600760209081526040808320848452909152812054600c80549091608060020a90046001604060020a0316908110611e7d57fe5b90600052602060002090600302019050806002015488141515611e9f57600080fd5b505b815460c060020a90046001604060020a031681141561182b5761182b6131b1565b600080546101009004600160a060020a03163314611edf57600080fd5b60028060005460a860020a900460ff166002811115611efa57fe5b1415611f0557600080fd5b6000805475ff00000000000000000000000000000000000000000019167502000000000000000000000000000000000000000000179055611f446133a5565b600191505b5090565b600c5490565b600c805482908110611c0757fe5b60056020526000908152604090205481565b60008115611f8e57600b805484908110611f8957fe5b506000525b600a805484908110611f9c57fe5b60009182526020909120600590910201546a0100000000000000000000900460ff169392505050565b6000611fd03361339d565b1515611fdb57600080fd5b33600090815260126020526040902054600160a060020a031615611ffe57600080fd5b503360009081526012602052604090208054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600b805482908110611b5f57fe5b6007602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401546001604060020a0380851695604060020a8604821695608060020a810483169560c060020a90910490921693919260ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168c565b600f5481565b60008083156121255782612108576121033467016345785d8a000063ffffffff61357716565b612120565b612120346702c68af0bb14000063ffffffff61357716565b612127565b345b9050600160a060020a0387163314801561214057508015155b801561214a575085155b8015612154575084155b8061218d5750600160a060020a03878116600090815260126020526040902054161580159061218257508515155b801561218d57508415155b151561219857600080fd5b88546121a78a60018301614ab0565b9150600089838154811015156121b957fe5b600091825260209091206005909102016001810180543373ffffffffffffffffffffffffffffffffffffffff19918216179091558154600283018054909216600160a060020a038c161790915560038201899055600482018890557fffffffff00000000000000000000000000000000ffffffffffffffffffffffff16606060020a6001608060020a038516021767ffffffffffffffff1916426001604060020a03161768ff00000000000000001916604060020a87151590810291909117825590915061237757805469ff0000000000000000001916604860020a178155600160a060020a038816331461237757604080516101408101825282546001604060020a0381168252604060020a810460ff90811615156020840152604860020a820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001820154600160a060020a0390811660c083015260028301541660e08201526003820154610100820152600482015461012082015261236c908461268b565b151561237757600080fd5b8854600090151561239b5789546123918b60018301614adc565b50600090506123a3565b508854600019015b60008a828154811015156123b357fe5b906000526020600020906003020190506123cc8161358c565b805460ff168061240f5750805460016001604060020a03604860020a830481167101000000000000000000000000000000000090930481169290920301166103e8145b15612473578a548b906124258260018301614adc565b8154811061242f57fe5b60009182526020909120600390910201805470ffffffffffffffff0000000000000000001916604860020a6001604060020a0388160217815590506124738161358c565b805478ffffffffffffffff00000000000000000000000000000000001916710100000000000000000000000000000000006001604060020a0387160217815533600160a060020a038b16141561259e57604080516101408101825284546001604060020a0381168252604060020a810460ff90811615156020840152604860020a820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e0820152600384015461010082015260048401546101208201526125999061258b908c6135f8565b82908763ffffffff61365a16565b61267c565b600160a060020a038a81166000908152601260209081526040918290205482516101408101845287546001604060020a0381168252604060020a810460ff908116151594830194909452604860020a810484161515948201949094526a010000000000000000000084048316151560608201526b010000000000000000000000840490921615156080830152606060020a9092046001608060020a031660a08201526001860154831660c08201526002860154831660e08201526003860154610100820152600486015461012082015261267c9261258b92166135f8565b50505097509795505050505050565b60008260e00151600160a060020a03168360c00151600160a060020a03161415612703578260e00151600160a060020a03166108fc8460a001516001608060020a03169081150290604051600060405180830381858888f193505050501580156126f9573d6000803e3d6000fd5b5060019050611a9a565b60e083015160208085015160c0860151610100870151610120880151604080517fd9afd3a9000000000000000000000000000000000000000000000000000000008152941515600486015260248501899052600160a060020a039384166044860152606485019290925260848401525193169263d9afd3a99260a4808401939192918290030181600087803b15801561279b57600080fd5b505af11580156127af573d6000803e3d6000fd5b505050506040513d60208110156127c557600080fd5b50519392505050565b6127d6614a48565b60606127f260096127e6856138b8565b9063ffffffff6138db16565b905061281581600081518110151561280657fe5b90602001906020020151613969565b6001604060020a031682528051612833908290600190811061280657fe5b6020830152805161284b908290600290811061280657fe5b6001604060020a03166040830152805161287b908290600390811061286c57fe5b9060200190602002015161398d565b600160a060020a03166060830152805161289c908290600490811061280657fe5b608083015280516128c390829060059081106128b457fe5b906020019060200201516139a9565b60a083015280516128db908290600690811061280657fe5b60c083015280516128f3908290600790811061280657fe5b60e0830152805161290b908290600890811061280657fe5b61010083015250919050565b60008160c00151600014801561292f575060e0820151155b8015611a9a57505061010001511590565b6000600260005460a860020a900460ff16600281111561295c57fe5b141561296a5750600061142a565b6001805460009081526006602090815260408083205460059092529091205491019081111561299d57600091505061142a565b6001546000908152600760209081526040808320848452909152902060048101546301000000900460ff16156129d85760009250505061142a565b600481015460ff1615612a25578054426001604060020a0360c060020a909204919091166001011115612a105760009250505061142a565b612a1a8183613a01565b60019250505061142a565b805460016001604060020a03604060020a90920482160116612a4681613a6e565b15612a72578154612a6690604060020a90046001604060020a0316613bae565b6001935050505061142a565b8154426001604060020a0360c060020a909204919091166003011115612a9e576000935050505061142a565b612aa88284613a01565b6001935050505090565b6000838015612abe5750825b8015612ac75750815b15612c2d576000612ae360018054613d0f90919063ffffffff16565b6000818152600260208181526040808420546008835281852081865290925290922090810154929350909162010000900460ff168015612b3e57506001810154426001604060020a0360c060020a90920491909116610e1001105b15612b8857506000828152600860209081526040808320938352928152828220828155600181018390556002908101805464ffffffffff1916905593825292909252812055612cfc565b6001838155600084815260086020818152604080842060028352818520548552825280842054600019890185529282528084208785528252928390209093018054608060020a9092046001604060020a031667ffffffffffffffff1990921682179055815186815292830181905281519096507f18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a89281900390910190a1505050612c52565b60018054600090815260056020526040902054612c4f9163ffffffff613d0f16565b90505b60018054600090815260076020908152604080832085845282528083208085018c9055600281018b9055600381018a905580546001604060020a0342811660c060020a02600160c060020a039092169190911780835560048054909216604060020a026fffffffffffffffff0000000000000000199091161782550180548815156101000261ff00198b151560ff1990931692909217919091161790559254825260059052208190555b9695505050505050565b6004805460018181019283905580546000908152600860209081526040808320958352858252808320948352949052929092205481546001604060020a0360c060020a92839004811685018116608060020a0277ffffffffffffffff000000000000000000000000000000001990921691909117835560028301805461ff001962ff000019909116620100001716610100179055928201805442909416909102600160c060020a0390931692909217909155612dc181613d28565b600281015460ff1615612e45578054600160c060020a0377ffffffffffffffff00000000000000000000000000000000196fffffffffffffffff00000000000000001983166001604060020a03938416604060020a0217908116608060020a918290048416600019018416820217918216910490911660c060020a02178155612ec9565b600a5481546fffffffffffffffff00000000000000001916604060020a6000199092016001604060020a03908116830291909117808455600192612e9892828116919092048216038301166103e8614156565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a039091161781555b600080546001919075ff000000000000000000000000000000000000000000191660a860020a8302179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff169080826002811115612f3557fe5b60ff16815260200191505060405180910390a160045481546002830154604080519384526001604060020a03608060020a84048116602086015260c060020a84048116858301528084166060860152604060020a909304909216608084015260ff16151560a0830152600160c0830152600060e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a1600281015460ff1615612ff257612fed6131b1565b61317b565b604080516101a08101825282546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018601548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600282015460ff8082161515610100808501919091528204811615156101208401526201000082048116151561014084015263010000008204811615156101608401526401000000009091041615156101808201526000906130c89061317e565b905060005b81816001604060020a03161015613178576001805460009081526007602081815260408084208854608060020a908190046001604060020a03908116890181168752918452828620600401805460ff19168817905589870154875487529484528286208a5482900483168901831687529093529320805477ffffffffffffffff000000000000000000000000000000001916604060020a9093048416860190931602179055016130cd565b50505b50565b60008161010001511561319357506000611d27565b81604001518260600151036001016001604060020a03169050919050565b6004805460019081019182905580546000908152600860209081526040808320858452909152902091811461321c57506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0360c060020a9091048116909101165b60028201805461010061ff0019909116179055815477ffffffffffffffff000000000000000000000000000000001916608060020a6001604060020a038381169190910291909117808455600954600160c060020a0391821660c060020a9185016000190184168202178555600185018054909216429093160291909117905560008054819075ff000000000000000000000000000000000000000000191660a860020a8202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561330657fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a03608060020a83048116602085015260c060020a909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b6000903b1190565b60018054600081815260066020908152604080832054600783528184208185528352818420805495870180865260028552838620604060020a9097046001604060020a03169687905560088552838620878752909452919093209294909391929083141561342157805467ffffffffffffffff1916815561346f565b6001805460009081526008602090815260408083206002835281842054845290915290205482546001604060020a03604060020a90920482169092011667ffffffffffffffff199091161781555b60028101805463ff0000001962ff0000199091166201000017166301000000179055600181810180546001604060020a03421660c060020a02600160c060020a03909116179055600b546134c89163ffffffff61357716565b81546fffffffffffffffff00000000000000001916604060020a6001604060020a0392831681029190911777ffffffffffffffff000000000000000000000000000000001916608060020a60018981018516919091029190911780855590926135419282821692048116919091038301166103e8614156565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a0390911617905550505050565b60008282111561358657600080fd5b50900390565b6001810154600160a060020a0316151561317b576135a8614b08565b604051809103906000f0801580156135c4573d6000803e3d6000fd5b50600191909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b613600614b18565b60208084015115159082015260c080840151600160a060020a039081169183019190915260a0808501516001608060020a03169083015261010080850151908301526101209384015193820193909352911660e082015290565b6001830154600160a060020a0316151561367357600080fd5b604080516020808252818301909252606091602082018180388339505085549192506000916136b491508490604860020a90046001604060020a0316613577565b905080602083015260606136d26136cd86866000614190565b6141bf565b6001870154604080517f20ba5b6000000000000000000000000000000000000000000000000000000000815260048101918252865160448201528651939450600160a060020a03909216926320ba5b60928792869290918291602482019160640190602087019080838360005b8381101561375757818101518382015260200161373f565b50505050905090810190601f1680156137845780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156137b757818101518382015260200161379f565b50505050905090810190601f1680156137e45780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561380557600080fd5b505af1158015613819573d6000803e3d6000fd5b5050506001870154604080517f80759f1f0000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692506380759f1f916004808301926020929190829003018186803b15801561387c57600080fd5b505afa158015613890573d6000803e3d6000fd5b505050506040513d60208110156138a657600080fd5b50516002909601959095555050505050565b6138c0614b6c565b50805160408051808201909152602092830181529182015290565b60608160405190808252806020026020018201604052801561391757816020015b613904614b6c565b8152602001906001900390816138fc5790505b509050613922614b83565b61392b84614375565b905060005b83811015613961576139418261439a565b838281518110151561394f57fe5b60209081029091010152600101613930565b505092915050565b6000806000613977846143cc565b90516020919091036101000a9004949350505050565b600080613999836143cc565b5051606060020a90049392505050565b60208101516060908015156139be5750611d27565b806040519080825280601f01601f1916602001820160405280156139e9576020820181803883390190505b5091506139fb8360000151838361442b565b50919050565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b6000600454821115613a8257506000611d27565b60015460009081526008602090815260408083208584529091529020600281015462010000900460ff161515613abc576000915050611d27565b600281015460ff1615613aec576001908101544260c060020a9091046001604060020a0316909101119050611d27565b6001546000908152600560205260409020548154608060020a90046001604060020a03161115613b20576000915050611d27565b60015460009081526007602090815260408083208454608060020a90046001604060020a0316845290915290206004810154640100000000900460ff1615613b6d57600192505050611d27565b60048101546301000000900460ff1615613b8c57600092505050611d27565b5442600160c060020a9092046001604060020a03169190910111159392505050565b6001546000908152600860209081526040808320848452909152812080549091608060020a9091046001604060020a0316905b825460c060020a90046001604060020a03168211613c6a576001546000908152600760209081526040808320858452909152902060048101546301000000900460ff1680613c395750600481015462010000900460ff165b15613c48576001915050613c6a565b600401805464ff00000000191664010000000017905560019190910190613be1565b600081613c775782613c7c565b600183035b8454909150608060020a90046001604060020a03168110613d085760018054600090815260066020908152604091829020849055915486548251918252928101889052608060020a9092046001604060020a03168282015260608201839052517f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5916080908290030190a15b5050505050565b600082820183811015613d2157600080fd5b9392505050565b600a541515613d455760028101805460ff1916600117905561317b565b600154600090815260086020908152604080832060045460011901845290915290208054600a54604060020a9091046001604060020a03166000199091011415613d995760028201805460ff191660011790555b60015460021415613daa575061317b565b600282015460ff1615613f415780548254604060020a9091046001604060020a031667ffffffffffffffff19909116178255600281015460ff1615613e275760018181015490830180546fffffffffffffffff00000000000000001916604060020a928390046001604060020a0316909202919091179055613f3c565b604080516101a08101825282546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018601548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600282015460ff808216151561010080850191909152820481161515610120840152620100008204811615156101408401526301000000820481161515610160840152640100000000909104161515610180820152613efa9061317e565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b614112565b6001546000908152600360205260409020541515613f7357600454600154600090815260036020526040902055614112565b600281015460ff1615613fdd578054825467ffffffffffffffff19166001604060020a03604060020a9283900481169190911784556001808401549085018054918490049092169092026fffffffffffffffff000000000000000019909216919091179055614112565b8054825460016001604060020a03604060020a9384900481168201811667ffffffffffffffff19909316929092178555604080516101a0810182528554808516825285810485166020830152608060020a80820486169383019390935260c060020a9081900485166060830152928601548085166080830152948504841660a0820152908404831660c082015292041660e0820152600282015460ff8082161515610100848101919091528204811615156101208401526201000082048116151561014084015263010000008204811615156101608401526401000000009091041615156101808201526140d09061317e565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b600c54600010156116d857600c805460019190600019810190811061413357fe5b60009182526020909120600390910201805460ff19169115159190911790555050565b6000818381151561416357fe5b061561417d57818381151561417457fe5b04600101613d21565b818381151561418857fe5b049392505050565b614198614a48565b60016020820152620186a060408201526141b3848484614469565b60a08201529392505050565b6040805160098082526101408201909252606091829190816020015b60608152602001906001900390816141db5750508351909150614206906001604060020a0316614554565b81600081518110151561421557fe5b9060200190602002018190525061422f8360200151614554565b81600181518110151561423e57fe5b60209081029091010152604083015161425f906001604060020a0316614554565b81600281518110151561426e57fe5b60209081029091010152606083015161428f90600160a060020a0316614567565b81600381518110151561429e57fe5b6020908102909101015260808301516142b690614554565b8160048151811015156142c557fe5b6020908102909101015260a08301516142dd90614597565b8160058151811015156142ec57fe5b6020908102909101015260c083015161430490614554565b81600681518110151561431357fe5b6020908102909101015260e083015161432b90614554565b81600781518110151561433a57fe5b6020908102909101015261010083015161435390614554565b81600881518110151561436257fe5b60209081029091010152613d218161461a565b61437d614b83565b60006143888361467b565b83519383529092016020820152919050565b6143a2614b6c565b602082015160006143b2826146cc565b828452602080850182905292019390910192909252919050565b805180516000918291821a9060808210156143ee579250600191506144269050565b60b882101561440c5760018560200151039250806001019350614423565b602085015182820160b51901945082900360b60192505b50505b915091565b6020601f820104836020840160005b838110156144565760208102838101519083015260010161443a565b5060008551602001860152505050505050565b6060600082614498577fe904e3d9000000000000000000000000000000000000000000000000000000006144ba565b7fd9afd3a9000000000000000000000000000000000000000000000000000000005b60208681015160c088015161010089015161012090990151604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19969096169486019490945291151560f860020a0260288501526029840197909752600160a060020a03909616606060020a026049830152605d820196909652607d8082019590955285518082039095018552609d01909452509092915050565b6060611a9a614562836146f8565b614597565b60408051741400000000000000000000000000000000000000008318601482015260348101909152606090613d21815b6060815160011480156145f6575081517f7f0000000000000000000000000000000000000000000000000000000000000090839060009081106145d657fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b15614602575080611d27565b611a9a6146148351608060ff16614821565b83614949565b604080516000808252602082019092526060915b83518110156146625761465882858381518110151561464957fe5b90602001906020020151614949565b915060010161462e565b50613d21614675825160c060ff16614821565b82614949565b8051805160009190821a90608082101561469a57600092505050611d27565b60b88210806146b5575060c082101580156146b5575060f882105b156146c557600192505050611d27565b5050919050565b8051600090811a60808110156146e557600191506139fb565b60b88110156139fb57607e190192915050565b6040805160208082528183019092526060916000918391602082018180388339019050509050836020820152600091505b602082101561478757808281518110151561474057fe5b60209101015160f860020a90819004027fff00000000000000000000000000000000000000000000000000000000000000161561477c57614787565b600190910190614729565b6060826020036040519080825280601f01601f1916602001820160405280156147b7576020820181803883390190505b50905060005b81518110156148185782516001850194849181106147d757fe5b90602001015160f860020a900460f860020a0282828151811015156147f857fe5b906020010190600160f860020a031916908160001a9053506001016147bd565b50949350505050565b6060604060020a831061489557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e70757420746f6f206c6f6e67000000000000000000000000000000000000604482015290519081900360640190fd5b604080516001808252818301909252606091602082018180388339019050509050603784116148f55782840160f860020a028160008151811015156148d657fe5b906020010190600160f860020a031916908160001a9053509050611a9a565b6060614900856146f8565b90508381510160370160f860020a0282600081518110151561491e57fe5b906020010190600160f860020a031916908160001a9053506149408282614949565b95945050505050565b60608082518451016040519080825280601f01601f19166020018201604052801561497b576020820181803883390190505b5090506000805b85518110156149de57858181518110151561499957fe5b90602001015160f860020a900460f860020a0283838151811015156149ba57fe5b906020010190600160f860020a031916908160001a90535060019182019101614982565b5060005b8451811015614a3e5784818151811015156149f957fe5b90602001015160f860020a900460f860020a028383815181101515614a1a57fe5b906020010190600160f860020a031916908160001a905350600191820191016149e2565b5090949350505050565b6101206040519081016040528060006001604060020a031681526020016000815260200160006001604060020a031681526020016000600160a060020a0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b815481835581811115613178576005028160050283600052602060002091820191016131789190614ba4565b815481835581811115613178576003028160030283600052602060002091820191016131789190614c18565b60405161126f80614c7583390190565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081019190915290565b604080518082019091526000808252602082015290565b606060405190810160405280614b97614b6c565b8152602001600081525090565b61142a91905b80821115611f495780547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff1990811690915560028201805490911690556000600382018190556004820155600501614baa565b61142a91905b80821115611f4957805478ffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff1916905560006002820155600301614c1e5600608060405234801561001057600080fd5b5061124f806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c610097366004610f92565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004610e82565b61018d565b6040516100cd939291906110f5565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004610f5d565b610278565b6040516100cd929190611161565b34801561011057600080fd5b506101196104d7565b6040516100cd919061112b565b34801561013257600080fd5b5061013b6104dd565b6040516100cd93929190611139565b34801561015657600080fd5b5061016a610165366004610ea0565b610527565b6040516100cd919061111d565b6101896000838363ffffffff61067c16565b5050565b610195610cc2565b61019d610cc2565b6101a5610cc2565b6101ad610cdd565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60008054606090151561028a57600080fd5b610292610cf6565b60408051908101604052808580519060200120815260200161010081525090506102ba610d0d565b506040805180820182526001548152815180830190925260025482526003546020838101919091528101919091526102f0610d28565b6000806102fb610cf6565b610303610cf6565b600061030d610cf6565b6020880151610323908a9063ffffffff61073b16565b6020808b0151810151908301519296509094501461033d57fe5b6020830151151561034d57610454565b60208401519590950160ff81900360020a9a909a17996001019461037083610769565b8951600090815260046020526040902091935091506103d69060018490036002811061039857fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526107d0565b6001860195889061010081106103e857fe5b6020908102919091019190915288516000908152600490915260409020826002811061041057fe5b60408051808201825260039290920292909201805482528251808401909352600181015483526002015460208381019190915281019190915290985096508761030d565b60008511156104c95784604051908082528060200260200182016040528015610487578160200160208202803883390190505b50995060005b858110156104c757878161010081106104a257fe5b60200201518b828151811015156104b557fe5b6020908102909101015260010161048d565b505b505050505050505050915091565b60005490565b60008060006104ea610d0d565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610531610cf6565b6040805190810160405280878051906020012081526020016101008152509050610559610d0d565b85516020870120815260005b851561065257600061057687610812565b60ff1690508060019060020a02198716965061059e8160ff038561086690919063ffffffff16565b602085018190529094506000906105b490610769565b602086015290506105c3610cc2565b6105cc856107d0565b8183600281106105d857fe5b602002015287518890600019868203019081106105f157fe5b90602001906020020151818360010360028110151561060c57fe5b6020908102919091019190915281519181015160408051808401949094528381019190915280518084038201815260609093019052815191012084525050600101610565565b5060208101829052610663816107d0565b881461066e57600080fd5b506001979650505050505050565b610684610cf6565b506040805180820190915282516020808501919091208252610100818301528251908301206106b1610d0d565b855415156106c8576020810183905281815261070b565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261070890879085856108da565b90505b610714816107d0565b86558051600187015560209081015180516002880155015160039095019490945550505050565b610743610cf6565b61074b610cf6565b61075e846107598686610b14565b610866565b915091509250929050565b6000610773610cf6565b602083015160001061078457600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b80516020918201518083015190516040805180860194909452838101929092526060808401919091528151808403909101815260809092019052805191012090565b600081151561082057600080fd5b8160805b600160ff82161061085f5760001960ff821660020a0182161515610852579182019160ff811660020a909104905b600260ff90911604610824565b5050919050565b61086e610cf6565b610876610cf6565b8360200151831115801561088c57506101008311155b151561089457fe5b602082018390528215156108ab57600082526108bd565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b6108e2610d0d565b6020808501518101519084015110156108f757fe5b6108ff610cf6565b610907610cf6565b6000610911610cf6565b61091f87896020015161073b565b602081015191955093506000901515610939575085610aef565b6020808a01518101519086015110610a6057602084015160011061095957fe5b610961610cdd565b8951600090815260048c0160209081526040808320815160608101909252909290918391908201908390600290835b828210156109de576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610990565b505050508152505090506109f185610769565b82519195509350610a15908c908660028110610a0957fe5b6020020151858b6108da565b81518560028110610a2257fe5b602090810291909101919091528a51600090815260048d019091526040812090610a4c8282610d49565b5050610a588b82610b84565b915050610aef565b610a6984610769565b9093509150610a76610cdd565b604080518082019091528881526020810184905281518560028110610a9757fe5b602002018190525060408051908101604052808b600001518152602001610ac98c602001518960200151600101610bea565b90528151600186900360028110610adc57fe5b6020020152610aeb8b82610b84565b9150505b604080519081016040528082815260200186815250955050505050505b949350505050565b6000808260200151846020015110610b30578260200151610b36565b83602001515b9050801515610b49576000915050610b7e565b825184511861010082900360020a60000316801515610b6a57509050610b7e565b610b7381610c1e565b60ff0360ff16925050505b92915050565b600080610b9083610c6c565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610bf2610cf6565b6020830151821115610c0357600080fd5b60208084015183900390820152915160029190910a02815290565b6000811515610c2c57600080fd5b8160805b600160ff82161061085f5760ff811660020a600019810102821615610c5f579182019160ff811660020a909104905b600260ff90911604610c30565b8051600090610c8190825b60200201516107d0565b8251610c8e906001610c77565b6040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610cf1610d73565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610cf1610cf6565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b610d8b610d0d565b815260200190600190039081610d835790505090565b6000601f82018313610db257600080fd5b8135610dc5610dc0826111a8565b611181565b91508181835260208401935060208101905083856020840282011115610dea57600080fd5b60005b83811015610e165781610e008882610e20565b8452506020928301929190910190600101610ded565b5050505092915050565b6000610e2c82356111f1565b9392505050565b6000601f82018313610e4457600080fd5b8135610e52610dc0826111c9565b91508082526020830160208301858383011115610e6e57600080fd5b610e79838284611209565b50505092915050565b600060208284031215610e9457600080fd5b6000610b0c8484610e20565b600080600080600060a08688031215610eb857600080fd5b6000610ec48888610e20565b955050602086013567ffffffffffffffff811115610ee157600080fd5b610eed88828901610e33565b945050604086013567ffffffffffffffff811115610f0a57600080fd5b610f1688828901610e33565b9350506060610f2788828901610e20565b925050608086013567ffffffffffffffff811115610f4457600080fd5b610f5088828901610da1565b9150509295509295909350565b600060208284031215610f6f57600080fd5b813567ffffffffffffffff811115610f8657600080fd5b610b0c84828501610e33565b60008060408385031215610fa557600080fd5b823567ffffffffffffffff811115610fbc57600080fd5b610fc885828601610e33565b925050602083013567ffffffffffffffff811115610fe557600080fd5b610ff185828601610e33565b9150509250929050565b611004816111fa565b61100d826111f1565b60005b8281101561103d576110238583516110ec565b61102c826111f4565b602095909501949150600101611010565b5050505050565b600061104f82611200565b808452602084019350611061836111f4565b60005b82811015611091576110778683516110ec565b611080826111f4565b602096909601959150600101611064565b5093949350505050565b6110a4816111fa565b6110ad826111f1565b60005b8281101561103d576110c38583516110ec565b6110cc826111f4565b6020959095019491506001016110b0565b6110e681611204565b82525050565b6110e6816111f1565b60c081016111038286610ffb565b6111106040830185610ffb565b610b0c608083018461109b565b60208101610b7e82846110dd565b60208101610b7e82846110ec565b6060810161114782866110ec565b61115460208301856110ec565b610b0c60408301846110ec565b6040810161116f82856110ec565b8181036020830152610b0c8184611044565b60405181810167ffffffffffffffff811182821017156111a057600080fd5b604052919050565b600067ffffffffffffffff8211156111bf57600080fd5b5060209081020190565b600067ffffffffffffffff8211156111e057600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a7230582018b914ecdc6c0bae868236caa7c87914f895736b958199d65b8e37a35f3a8b3c6c6578706572696d656e74616cf50037a165627a7a72305820b4503c5d28e45144cd6367f02eff6518e18bf66da9d41c1db28f9b3374e2a5c90029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _NRBEpochLength *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend, _development, _NRBEpochLength, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainSession) COSTERO() (*big.Int, error) {
	return _RootChain.Contract.COSTERO(&_RootChain.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTERO() (*big.Int, error) {
	return _RootChain.Contract.COSTERO(&_RootChain.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainSession) COSTERU() (*big.Int, error) {
	return _RootChain.Contract.COSTERU(&_RootChain.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTERU() (*big.Int, error) {
	return _RootChain.Contract.COSTERU(&_RootChain.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTNRB() (*big.Int, error) {
	return _RootChain.Contract.COSTNRB(&_RootChain.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTNRB() (*big.Int, error) {
	return _RootChain.Contract.COSTNRB(&_RootChain.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTORB() (*big.Int, error) {
	return _RootChain.Contract.COSTORB(&_RootChain.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTORB() (*big.Int, error) {
	return _RootChain.Contract.COSTORB(&_RootChain.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTURB() (*big.Int, error) {
	return _RootChain.Contract.COSTURB(&_RootChain.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTURB() (*big.Int, error) {
	return _RootChain.Contract.COSTURB(&_RootChain.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChain.Contract.COSTURBPREPARE(&_RootChain.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChain.Contract.COSTURBPREPARE(&_RootChain.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChain.Contract.CPCOMPUTATION(&_RootChain.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChain.Contract.CPCOMPUTATION(&_RootChain.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChain.Contract.CPWITHHOLDING(&_RootChain.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChain.Contract.CPWITHHOLDING(&_RootChain.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		Applied    bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCallerSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		Applied    bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCallerSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainCaller) MAXREQUESTS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "MAX_REQUESTS")
	return *ret0, err
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainCallerSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainCaller) NRBEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "NRBEpochLength")
	return *ret0, err
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainSession) NRBEpochLength() (*big.Int, error) {
	return _RootChain.Contract.NRBEpochLength(&_RootChain.CallOpts)
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainCallerSession) NRBEpochLength() (*big.Int, error) {
	return _RootChain.Contract.NRBEpochLength(&_RootChain.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainSession) NULLADDRESS() (common.Address, error) {
	return _RootChain.Contract.NULLADDRESS(&_RootChain.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainCallerSession) NULLADDRESS() (common.Address, error) {
	return _RootChain.Contract.NULLADDRESS(&_RootChain.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		EpochNumber      uint64
		RequestStart     uint64
		RequestEnd       uint64
		Trie             common.Address
		TransactionsRoot [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChain.Contract.PREPARETIMEOUT(&_RootChain.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChain.Contract.PREPARETIMEOUT(&_RootChain.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainSession) REQUESTGAS() (*big.Int, error) {
	return _RootChain.Contract.REQUESTGAS(&_RootChain.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainCallerSession) REQUESTGAS() (*big.Int, error) {
	return _RootChain.Contract.REQUESTGAS(&_RootChain.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		EpochNumber      uint64
		RequestStart     uint64
		RequestEnd       uint64
		Trie             common.Address
		TransactionsRoot [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	ret := new(struct {
		ForkNumber             uint64
		EpochNumber            uint64
		RequestBlockId         uint64
		Timestamp              uint64
		StatesRoot             [32]byte
		TransactionsRoot       [32]byte
		IntermediateStatesRoot [32]byte
		IsRequest              bool
		UserActivated          bool
		Challenged             bool
		Challenging            bool
		Finalized              bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "blocks", arg0, arg1)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	return _RootChain.Contract.Blocks(&_RootChain.CallOpts, arg0, arg1)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCallerSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	return _RootChain.Contract.Blocks(&_RootChain.CallOpts, arg0, arg1)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentEpoch() (*big.Int, error) {
	return _RootChain.Contract.CurrentEpoch(&_RootChain.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentEpoch() (*big.Int, error) {
	return _RootChain.Contract.CurrentEpoch(&_RootChain.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentFork() (*big.Int, error) {
	return _RootChain.Contract.CurrentFork(&_RootChain.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentFork() (*big.Int, error) {
	return _RootChain.Contract.CurrentFork(&_RootChain.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainSession) Development() (bool, error) {
	return _RootChain.Contract.Development(&_RootChain.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainCallerSession) Development() (bool, error) {
	return _RootChain.Contract.Development(&_RootChain.CallOpts)
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	ret := new(struct {
		RequestStart        uint64
		RequestEnd          uint64
		StartBlockNumber    uint64
		EndBlockNumber      uint64
		ForkedBlockNumber   uint64
		FirstRequestBlockId uint64
		Limit               uint64
		Timestamp           uint64
		IsEmpty             bool
		Initialized         bool
		IsRequest           bool
		UserActivated       bool
		Finalized           bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "epochs", arg0, arg1)
	return *ret, err
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainSession) Epochs(arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	return _RootChain.Contract.Epochs(&_RootChain.CallOpts, arg0, arg1)
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCallerSession) Epochs(arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	return _RootChain.Contract.Epochs(&_RootChain.CallOpts, arg0, arg1)
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstEpoch", arg0)
	return *ret0, err
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) FirstEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstEpoch(&_RootChain.CallOpts, arg0)
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstEpoch(&_RootChain.CallOpts, arg0)
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstFilledORBEpochNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstFilledORBEpochNumber", arg0)
	return *ret0, err
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) FirstFilledORBEpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORBEpochNumber(&_RootChain.CallOpts, arg0)
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstFilledORBEpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORBEpochNumber(&_RootChain.CallOpts, arg0)
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainCaller) GetNumEROs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getNumEROs")
	return *ret0, err
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainSession) GetNumEROs() (*big.Int, error) {
	return _RootChain.Contract.GetNumEROs(&_RootChain.CallOpts)
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetNumEROs() (*big.Int, error) {
	return _RootChain.Contract.GetNumEROs(&_RootChain.CallOpts)
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainCaller) GetNumORBs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getNumORBs")
	return *ret0, err
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainSession) GetNumORBs() (*big.Int, error) {
	return _RootChain.Contract.GetNumORBs(&_RootChain.CallOpts)
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetNumORBs() (*big.Int, error) {
	return _RootChain.Contract.GetNumORBs(&_RootChain.CallOpts)
}

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainCaller) GetRequestApplied(opts *bind.CallOpts, _requestId *big.Int, _userActivated bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getRequestApplied", _requestId, _userActivated)
	return *ret0, err
}

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainSession) GetRequestApplied(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestApplied(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainCallerSession) GetRequestApplied(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestApplied(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainCaller) GetRequestFinalized(opts *bind.CallOpts, _requestId *big.Int, _userActivated bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getRequestFinalized", _requestId, _userActivated)
	return *ret0, err
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainCallerSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) HighestBlockNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "highestBlockNumber", arg0)
	return *ret0, err
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) HighestBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.HighestBlockNumber(&_RootChain.CallOpts, arg0)
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) HighestBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.HighestBlockNumber(&_RootChain.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedBlockNumber(&_RootChain.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedBlockNumber(&_RootChain.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedERO")
	return *ret0, err
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedERO() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERO(&_RootChain.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedERO() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERO(&_RootChain.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedERU")
	return *ret0, err
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedERU() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERU(&_RootChain.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedERU() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERU(&_RootChain.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedForkNumber(&_RootChain.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedForkNumber(&_RootChain.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) LastFinalizedBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastFinalizedBlock", arg0)
	return *ret0, err
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) LastFinalizedBlock(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastFinalizedBlock(&_RootChain.CallOpts, arg0)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastFinalizedBlock(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastFinalizedBlock(&_RootChain.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCallerSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainSession) State() (uint8, error) {
	return _RootChain.Contract.State(&_RootChain.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainCallerSession) State() (uint8, error) {
	return _RootChain.Contract.State(&_RootChain.CallOpts)
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainTransactor) ApplyRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "applyRequest")
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainSession) ApplyRequest() (*types.Transaction, error) {
	return _RootChain.Contract.ApplyRequest(&_RootChain.TransactOpts)
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainTransactorSession) ApplyRequest() (*types.Transaction, error) {
	return _RootChain.Contract.ApplyRequest(&_RootChain.TransactOpts)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainTransactor) ChallengeNullAddress(opts *bind.TransactOpts, _blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeNullAddress", _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainTransactorSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainTransactor) FinalizeBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeBlock")
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainTransactorSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainTransactor) FinalizeRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeRequest")
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainTransactorSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainTransactor) Forked(opts *bind.TransactOpts, _forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "forked", _forkNumber)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainSession) Forked(_forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Forked(&_RootChain.TransactOpts, _forkNumber)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainTransactorSession) Forked(_forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Forked(&_RootChain.TransactOpts, _forkNumber)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) MakeERU(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "makeERU", _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainTransactor) MapRequestableContract(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "mapRequestableContract", _target)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainSession) MapRequestableContract(_target common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContract(&_RootChain.TransactOpts, _target)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainTransactorSession) MapRequestableContract(_target common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContract(&_RootChain.TransactOpts, _target)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainTransactor) MapRequestableContractByOperator(opts *bind.TransactOpts, _rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "mapRequestableContractByOperator", _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainTransactorSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainTransactor) PrepareToSubmitURB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "prepareToSubmitURB")
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainTransactorSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainTransactor) RevertBlock(opts *bind.TransactOpts, _forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "revertBlock", _forkNumber, _blockNumber)
}

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainSession) RevertBlock(_forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RevertBlock(&_RootChain.TransactOpts, _forkNumber, _blockNumber)
}

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainTransactorSession) RevertBlock(_forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RevertBlock(&_RootChain.TransactOpts, _forkNumber, _blockNumber)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartEnter(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startEnter", _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitNRB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitNRB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitORB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitORB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitURB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitURB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// RootChainBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the RootChain contract.
type RootChainBlockFinalizedIterator struct {
	Event *RootChainBlockFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBlockFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainBlockFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBlockFinalized represents a BlockFinalized event raised by the RootChain contract.
type RootChainBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainBlockFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBlockFinalizedIterator{contract: _RootChain.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *RootChainBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBlockFinalized)
				if err := _RootChain.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the RootChain contract.
type RootChainERUCreatedIterator struct {
	Event *RootChainERUCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainERUCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainERUCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainERUCreated represents a ERUCreated event raised by the RootChain contract.
type RootChainERUCreated struct {
	RequestId *big.Int
	Requestor common.Address
	To        common.Address
	TrieKey   [32]byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainERUCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainERUCreatedIterator{contract: _RootChain.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *RootChainERUCreated) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainERUCreated)
				if err := _RootChain.contract.UnpackLog(event, "ERUCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the RootChain contract.
type RootChainEpochFinalizedIterator struct {
	Event *RootChainEpochFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainEpochFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochFinalized represents a EpochFinalized event raised by the RootChain contract.
type RootChainEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	FirstBlockNumber *big.Int
	LastBlockNumber  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, firstBlockNumber uint256, lastBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEpochFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFinalizedIterator{contract: _RootChain.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, firstBlockNumber uint256, lastBlockNumber uint256)
func (_RootChain *RootChainFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *RootChainEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochFinalized)
				if err := _RootChain.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the RootChain contract.
type RootChainEpochPreparedIterator struct {
	Event *RootChainEpochPrepared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochPrepared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainEpochPrepared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochPrepared represents a EpochPrepared event raised by the RootChain contract.
type RootChainEpochPrepared struct {
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	RequestStart     *big.Int
	RequestEnd       *big.Int
	EpochIsEmpty     bool
	IsRequest        bool
	UserActivated    bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochPrepared is a free log retrieval operation binding the contract event 0x99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2.
//
// Solidity: e EpochPrepared(epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEpochPreparedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochPreparedIterator{contract: _RootChain.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2.
//
// Solidity: e EpochPrepared(epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChain *RootChainFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *RootChainEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochPrepared)
				if err := _RootChain.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the RootChain contract.
type RootChainForkedIterator struct {
	Event *RootChainForked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainForked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainForked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainForked represents a Forked event raised by the RootChain contract.
type RootChainForked struct {
	NewFork           *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a8.
//
// Solidity: e Forked(newFork uint256, forkedBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainForkedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainForkedIterator{contract: _RootChain.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a8.
//
// Solidity: e Forked(newFork uint256, forkedBlockNumber uint256)
func (_RootChain *RootChainFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *RootChainForked) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainForked)
				if err := _RootChain.contract.UnpackLog(event, "Forked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainNRBSubmittedIterator is returned from FilterNRBSubmitted and is used to iterate over the raw logs and unpacked data for NRBSubmitted events raised by the RootChain contract.
type RootChainNRBSubmittedIterator struct {
	Event *RootChainNRBSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainNRBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainNRBSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainNRBSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainNRBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainNRBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainNRBSubmitted represents a NRBSubmitted event raised by the RootChain contract.
type RootChainNRBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNRBSubmitted is a free log retrieval operation binding the contract event 0xd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f1.
//
// Solidity: e NRBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterNRBSubmitted(opts *bind.FilterOpts) (*RootChainNRBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "NRBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainNRBSubmittedIterator{contract: _RootChain.contract, event: "NRBSubmitted", logs: logs, sub: sub}, nil
}

// WatchNRBSubmitted is a free log subscription operation binding the contract event 0xd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f1.
//
// Solidity: e NRBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchNRBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainNRBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "NRBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainNRBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "NRBSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainORBSubmittedIterator is returned from FilterORBSubmitted and is used to iterate over the raw logs and unpacked data for ORBSubmitted events raised by the RootChain contract.
type RootChainORBSubmittedIterator struct {
	Event *RootChainORBSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainORBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainORBSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainORBSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainORBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainORBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainORBSubmitted represents a ORBSubmitted event raised by the RootChain contract.
type RootChainORBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterORBSubmitted is a free log retrieval operation binding the contract event 0x041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d5.
//
// Solidity: e ORBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterORBSubmitted(opts *bind.FilterOpts) (*RootChainORBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ORBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainORBSubmittedIterator{contract: _RootChain.contract, event: "ORBSubmitted", logs: logs, sub: sub}, nil
}

// WatchORBSubmitted is a free log subscription operation binding the contract event 0x041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d5.
//
// Solidity: e ORBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchORBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainORBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ORBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainORBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "ORBSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RootChain contract.
type RootChainRequestCreatedIterator struct {
	Event *RootChainRequestCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainRequestCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestCreated represents a RequestCreated event raised by the RootChain contract.
type RootChainRequestCreated struct {
	RequestId     *big.Int
	Requestor     common.Address
	To            common.Address
	WeiAmount     *big.Int
	TrieKey       [32]byte
	TrieValue     [32]byte
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f0.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isExit bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainRequestCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestCreatedIterator{contract: _RootChain.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f0.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isExit bool, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RootChainRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestCreated)
				if err := _RootChain.contract.UnpackLog(event, "RequestCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the RootChain contract.
type RootChainRequestFinalizedIterator struct {
	Event *RootChainRequestFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainRequestFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestFinalized represents a RequestFinalized event raised by the RootChain contract.
type RootChainRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainRequestFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestFinalizedIterator{contract: _RootChain.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *RootChainRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestFinalized)
				if err := _RootChain.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the RootChain contract.
type RootChainSessionTimeoutIterator struct {
	Event *RootChainSessionTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSessionTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainSessionTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSessionTimeout represents a SessionTimeout event raised by the RootChain contract.
type RootChainSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChain *RootChainFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainSessionTimeoutIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainSessionTimeoutIterator{contract: _RootChain.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChain *RootChainFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *RootChainSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSessionTimeout)
				if err := _RootChain.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainStateChangedIterator is returned from FilterStateChanged and is used to iterate over the raw logs and unpacked data for StateChanged events raised by the RootChain contract.
type RootChainStateChangedIterator struct {
	Event *RootChainStateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainStateChanged represents a StateChanged event raised by the RootChain contract.
type RootChainStateChanged struct {
	State uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateChanged is a free log retrieval operation binding the contract event 0x551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6.
//
// Solidity: e StateChanged(state uint8)
func (_RootChain *RootChainFilterer) FilterStateChanged(opts *bind.FilterOpts) (*RootChainStateChangedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return &RootChainStateChangedIterator{contract: _RootChain.contract, event: "StateChanged", logs: logs, sub: sub}, nil
}

// WatchStateChanged is a free log subscription operation binding the contract event 0x551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6.
//
// Solidity: e StateChanged(state uint8)
func (_RootChain *RootChainFilterer) WatchStateChanged(opts *bind.WatchOpts, sink chan<- *RootChainStateChanged) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainStateChanged)
				if err := _RootChain.contract.UnpackLog(event, "StateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainURBSubmittedIterator is returned from FilterURBSubmitted and is used to iterate over the raw logs and unpacked data for URBSubmitted events raised by the RootChain contract.
type RootChainURBSubmittedIterator struct {
	Event *RootChainURBSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainURBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainURBSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainURBSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainURBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainURBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainURBSubmitted represents a URBSubmitted event raised by the RootChain contract.
type RootChainURBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterURBSubmitted is a free log retrieval operation binding the contract event 0xd8803c978d8d54db576ad0fe876c746dc02c25d13bd7847815860b723304c34f.
//
// Solidity: e URBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterURBSubmitted(opts *bind.FilterOpts) (*RootChainURBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "URBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainURBSubmittedIterator{contract: _RootChain.contract, event: "URBSubmitted", logs: logs, sub: sub}, nil
}

// WatchURBSubmitted is a free log subscription operation binding the contract event 0xd8803c978d8d54db576ad0fe876c746dc02c25d13bd7847815860b723304c34f.
//
// Solidity: e URBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchURBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainURBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "URBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainURBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "URBSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582055a16021c31a45d118cd984993fd2642aaad49aa15fcbff9f296fa9598670d900029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
