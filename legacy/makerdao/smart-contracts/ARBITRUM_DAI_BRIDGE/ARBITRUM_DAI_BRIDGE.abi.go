// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ARBITRUM_DAI_BRIDGE

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ARBITRUMDAIBRIDGEABI is the input ABI used to generate the binding from.
const ARBITRUMDAIBRIDGEABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Escrow\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TxToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"outboundCalldata\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"contractIInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Dai\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Dai\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARBITRUMDAIBRIDGE is an auto generated Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGE struct {
	ARBITRUMDAIBRIDGECaller     // Read-only binding to the contract
	ARBITRUMDAIBRIDGETransactor // Write-only binding to the contract
	ARBITRUMDAIBRIDGEFilterer   // Log filterer for contract events
}

// ARBITRUMDAIBRIDGECaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMDAIBRIDGETransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMDAIBRIDGEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUMDAIBRIDGEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMDAIBRIDGESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUMDAIBRIDGESession struct {
	Contract     *ARBITRUMDAIBRIDGE // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ARBITRUMDAIBRIDGECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUMDAIBRIDGECallerSession struct {
	Contract *ARBITRUMDAIBRIDGECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ARBITRUMDAIBRIDGETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUMDAIBRIDGETransactorSession struct {
	Contract     *ARBITRUMDAIBRIDGETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ARBITRUMDAIBRIDGERaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGERaw struct {
	Contract *ARBITRUMDAIBRIDGE // Generic contract binding to access the raw methods on
}

// ARBITRUMDAIBRIDGECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGECallerRaw struct {
	Contract *ARBITRUMDAIBRIDGECaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUMDAIBRIDGETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUMDAIBRIDGETransactorRaw struct {
	Contract *ARBITRUMDAIBRIDGETransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUMDAIBRIDGE creates a new instance of ARBITRUMDAIBRIDGE, bound to a specific deployed contract.
func NewARBITRUMDAIBRIDGE(address common.Address, backend bind.ContractBackend) (*ARBITRUMDAIBRIDGE, error) {
	contract, err := bindARBITRUMDAIBRIDGE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGE{ARBITRUMDAIBRIDGECaller: ARBITRUMDAIBRIDGECaller{contract: contract}, ARBITRUMDAIBRIDGETransactor: ARBITRUMDAIBRIDGETransactor{contract: contract}, ARBITRUMDAIBRIDGEFilterer: ARBITRUMDAIBRIDGEFilterer{contract: contract}}, nil
}

// NewARBITRUMDAIBRIDGECaller creates a new read-only instance of ARBITRUMDAIBRIDGE, bound to a specific deployed contract.
func NewARBITRUMDAIBRIDGECaller(address common.Address, caller bind.ContractCaller) (*ARBITRUMDAIBRIDGECaller, error) {
	contract, err := bindARBITRUMDAIBRIDGE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGECaller{contract: contract}, nil
}

// NewARBITRUMDAIBRIDGETransactor creates a new write-only instance of ARBITRUMDAIBRIDGE, bound to a specific deployed contract.
func NewARBITRUMDAIBRIDGETransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUMDAIBRIDGETransactor, error) {
	contract, err := bindARBITRUMDAIBRIDGE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGETransactor{contract: contract}, nil
}

// NewARBITRUMDAIBRIDGEFilterer creates a new log filterer instance of ARBITRUMDAIBRIDGE, bound to a specific deployed contract.
func NewARBITRUMDAIBRIDGEFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUMDAIBRIDGEFilterer, error) {
	contract, err := bindARBITRUMDAIBRIDGE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGEFilterer{contract: contract}, nil
}

// bindARBITRUMDAIBRIDGE binds a generic wrapper to an already deployed contract.
func bindARBITRUMDAIBRIDGE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUMDAIBRIDGEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMDAIBRIDGE.Contract.ARBITRUMDAIBRIDGECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.ARBITRUMDAIBRIDGETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.ARBITRUMDAIBRIDGETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMDAIBRIDGE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1Token) view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "calculateL2TokenAddress", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1Token) view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) CalculateL2TokenAddress(l1Token common.Address) (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.CalculateL2TokenAddress(&_ARBITRUMDAIBRIDGE.CallOpts, l1Token)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1Token) view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) CalculateL2TokenAddress(l1Token common.Address) (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.CalculateL2TokenAddress(&_ARBITRUMDAIBRIDGE.CallOpts, l1Token)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) CounterpartGateway() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.CounterpartGateway(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) CounterpartGateway() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.CounterpartGateway(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address l1Token, address from, address to, uint256 amount, bytes data) pure returns(bytes outboundCalldata)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) GetOutboundCalldata(opts *bind.CallOpts, l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) ([]byte, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "getOutboundCalldata", l1Token, from, to, amount, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address l1Token, address from, address to, uint256 amount, bytes data) pure returns(bytes outboundCalldata)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) GetOutboundCalldata(l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) ([]byte, error) {
	return _ARBITRUMDAIBRIDGE.Contract.GetOutboundCalldata(&_ARBITRUMDAIBRIDGE.CallOpts, l1Token, from, to, amount, data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address l1Token, address from, address to, uint256 amount, bytes data) pure returns(bytes outboundCalldata)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) GetOutboundCalldata(l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) ([]byte, error) {
	return _ARBITRUMDAIBRIDGE.Contract.GetOutboundCalldata(&_ARBITRUMDAIBRIDGE.CallOpts, l1Token, from, to, amount, data)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) Inbox() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Inbox(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) Inbox() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Inbox(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) IsOpen() (*big.Int, error) {
	return _ARBITRUMDAIBRIDGE.Contract.IsOpen(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) IsOpen() (*big.Int, error) {
	return _ARBITRUMDAIBRIDGE.Contract.IsOpen(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Dai is a free data retrieval call binding the contract method 0x59784bb7.
//
// Solidity: function l1Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) L1Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "l1Dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Dai is a free data retrieval call binding the contract method 0x59784bb7.
//
// Solidity: function l1Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) L1Dai() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Dai(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Dai is a free data retrieval call binding the contract method 0x59784bb7.
//
// Solidity: function l1Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) L1Dai() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Dai(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Escrow is a free data retrieval call binding the contract method 0x000cc9e6.
//
// Solidity: function l1Escrow() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) L1Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "l1Escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Escrow is a free data retrieval call binding the contract method 0x000cc9e6.
//
// Solidity: function l1Escrow() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) L1Escrow() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Escrow(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Escrow is a free data retrieval call binding the contract method 0x000cc9e6.
//
// Solidity: function l1Escrow() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) L1Escrow() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Escrow(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Router is a free data retrieval call binding the contract method 0x407395e0.
//
// Solidity: function l1Router() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) L1Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "l1Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Router is a free data retrieval call binding the contract method 0x407395e0.
//
// Solidity: function l1Router() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) L1Router() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Router(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L1Router is a free data retrieval call binding the contract method 0x407395e0.
//
// Solidity: function l1Router() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) L1Router() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L1Router(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L2Counterpart is a free data retrieval call binding the contract method 0x8e7c8efe.
//
// Solidity: function l2Counterpart() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) L2Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "l2Counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Counterpart is a free data retrieval call binding the contract method 0x8e7c8efe.
//
// Solidity: function l2Counterpart() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) L2Counterpart() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L2Counterpart(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L2Counterpart is a free data retrieval call binding the contract method 0x8e7c8efe.
//
// Solidity: function l2Counterpart() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) L2Counterpart() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L2Counterpart(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L2Dai is a free data retrieval call binding the contract method 0x422e67f4.
//
// Solidity: function l2Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) L2Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "l2Dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Dai is a free data retrieval call binding the contract method 0x422e67f4.
//
// Solidity: function l2Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) L2Dai() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L2Dai(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// L2Dai is a free data retrieval call binding the contract method 0x422e67f4.
//
// Solidity: function l2Dai() view returns(address)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) L2Dai() (common.Address, error) {
	return _ARBITRUMDAIBRIDGE.Contract.L2Dai(&_ARBITRUMDAIBRIDGE.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUMDAIBRIDGE.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Wards(&_ARBITRUMDAIBRIDGE.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGECallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Wards(&_ARBITRUMDAIBRIDGE.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) Close() (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Close(&_ARBITRUMDAIBRIDGE.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorSession) Close() (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Close(&_ARBITRUMDAIBRIDGE.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Deny(&_ARBITRUMDAIBRIDGE.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Deny(&_ARBITRUMDAIBRIDGE.TransactOpts, usr)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address l1Token, address from, address to, uint256 amount, bytes data) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.contract.Transact(opts, "finalizeInboundTransfer", l1Token, from, to, amount, data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address l1Token, address from, address to, uint256 amount, bytes data) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) FinalizeInboundTransfer(l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.FinalizeInboundTransfer(&_ARBITRUMDAIBRIDGE.TransactOpts, l1Token, from, to, amount, data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address l1Token, address from, address to, uint256 amount, bytes data) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorSession) FinalizeInboundTransfer(l1Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.FinalizeInboundTransfer(&_ARBITRUMDAIBRIDGE.TransactOpts, l1Token, from, to, amount, data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address l1Token, address to, uint256 amount, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(bytes)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactor) OutboundTransfer(opts *bind.TransactOpts, l1Token common.Address, to common.Address, amount *big.Int, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.contract.Transact(opts, "outboundTransfer", l1Token, to, amount, maxGas, gasPriceBid, data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address l1Token, address to, uint256 amount, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(bytes)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) OutboundTransfer(l1Token common.Address, to common.Address, amount *big.Int, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.OutboundTransfer(&_ARBITRUMDAIBRIDGE.TransactOpts, l1Token, to, amount, maxGas, gasPriceBid, data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address l1Token, address to, uint256 amount, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(bytes)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorSession) OutboundTransfer(l1Token common.Address, to common.Address, amount *big.Int, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.OutboundTransfer(&_ARBITRUMDAIBRIDGE.TransactOpts, l1Token, to, amount, maxGas, gasPriceBid, data)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGESession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Rely(&_ARBITRUMDAIBRIDGE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGETransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMDAIBRIDGE.Contract.Rely(&_ARBITRUMDAIBRIDGE.TransactOpts, usr)
}

// ARBITRUMDAIBRIDGEClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEClosedIterator struct {
	Event *ARBITRUMDAIBRIDGEClosed // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGEClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGEClosed)
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
		it.Event = new(ARBITRUMDAIBRIDGEClosed)
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
func (it *ARBITRUMDAIBRIDGEClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGEClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGEClosed represents a Closed event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterClosed(opts *bind.FilterOpts) (*ARBITRUMDAIBRIDGEClosedIterator, error) {

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGEClosedIterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGEClosed) (event.Subscription, error) {

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGEClosed)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseClosed(log types.Log) (*ARBITRUMDAIBRIDGEClosed, error) {
	event := new(ARBITRUMDAIBRIDGEClosed)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMDAIBRIDGEDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEDenyIterator struct {
	Event *ARBITRUMDAIBRIDGEDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGEDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGEDeny)
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
		it.Event = new(ARBITRUMDAIBRIDGEDeny)
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
func (it *ARBITRUMDAIBRIDGEDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGEDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGEDeny represents a Deny event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMDAIBRIDGEDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGEDenyIterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGEDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGEDeny)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseDeny(log types.Log) (*ARBITRUMDAIBRIDGEDeny, error) {
	event := new(ARBITRUMDAIBRIDGEDeny)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMDAIBRIDGEDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEDepositInitiatedIterator struct {
	Event *ARBITRUMDAIBRIDGEDepositInitiated // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGEDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGEDepositInitiated)
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
		it.Event = new(ARBITRUMDAIBRIDGEDepositInitiated)
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
func (it *ARBITRUMDAIBRIDGEDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGEDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGEDepositInitiated represents a DepositInitiated event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEDepositInitiated struct {
	L1Token        common.Address
	From           common.Address
	To             common.Address
	SequenceNumber *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed from, address indexed to, uint256 indexed sequenceNumber, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterDepositInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address, sequenceNumber []*big.Int) (*ARBITRUMDAIBRIDGEDepositInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "DepositInitiated", fromRule, toRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGEDepositInitiatedIterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed from, address indexed to, uint256 indexed sequenceNumber, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGEDepositInitiated, from []common.Address, to []common.Address, sequenceNumber []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "DepositInitiated", fromRule, toRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGEDepositInitiated)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed from, address indexed to, uint256 indexed sequenceNumber, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseDepositInitiated(log types.Log) (*ARBITRUMDAIBRIDGEDepositInitiated, error) {
	event := new(ARBITRUMDAIBRIDGEDepositInitiated)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMDAIBRIDGERelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGERelyIterator struct {
	Event *ARBITRUMDAIBRIDGERely // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGERelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGERely)
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
		it.Event = new(ARBITRUMDAIBRIDGERely)
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
func (it *ARBITRUMDAIBRIDGERelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGERelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGERely represents a Rely event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGERely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMDAIBRIDGERelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGERelyIterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGERely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGERely)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseRely(log types.Log) (*ARBITRUMDAIBRIDGERely, error) {
	event := new(ARBITRUMDAIBRIDGERely)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMDAIBRIDGETxToL2Iterator is returned from FilterTxToL2 and is used to iterate over the raw logs and unpacked data for TxToL2 events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGETxToL2Iterator struct {
	Event *ARBITRUMDAIBRIDGETxToL2 // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGETxToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGETxToL2)
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
		it.Event = new(ARBITRUMDAIBRIDGETxToL2)
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
func (it *ARBITRUMDAIBRIDGETxToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGETxToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGETxToL2 represents a TxToL2 event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGETxToL2 struct {
	From   common.Address
	To     common.Address
	SeqNum *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxToL2 is a free log retrieval operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterTxToL2(opts *bind.FilterOpts, from []common.Address, to []common.Address, seqNum []*big.Int) (*ARBITRUMDAIBRIDGETxToL2Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "TxToL2", fromRule, toRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGETxToL2Iterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "TxToL2", logs: logs, sub: sub}, nil
}

// WatchTxToL2 is a free log subscription operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchTxToL2(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGETxToL2, from []common.Address, to []common.Address, seqNum []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "TxToL2", fromRule, toRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGETxToL2)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "TxToL2", log); err != nil {
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

// ParseTxToL2 is a log parse operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseTxToL2(log types.Log) (*ARBITRUMDAIBRIDGETxToL2, error) {
	event := new(ARBITRUMDAIBRIDGETxToL2)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "TxToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator struct {
	Event *ARBITRUMDAIBRIDGEWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMDAIBRIDGEWithdrawalFinalized)
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
		it.Event = new(ARBITRUMDAIBRIDGEWithdrawalFinalized)
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
func (it *ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMDAIBRIDGEWithdrawalFinalized represents a WithdrawalFinalized event raised by the ARBITRUMDAIBRIDGE contract.
type ARBITRUMDAIBRIDGEWithdrawalFinalized struct {
	L1Token common.Address
	From    common.Address
	To      common.Address
	ExitNum *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed from, address indexed to, uint256 indexed exitNum, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address, exitNum []*big.Int) (*ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var exitNumRule []interface{}
	for _, exitNumItem := range exitNum {
		exitNumRule = append(exitNumRule, exitNumItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.FilterLogs(opts, "WithdrawalFinalized", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMDAIBRIDGEWithdrawalFinalizedIterator{contract: _ARBITRUMDAIBRIDGE.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed from, address indexed to, uint256 indexed exitNum, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *ARBITRUMDAIBRIDGEWithdrawalFinalized, from []common.Address, to []common.Address, exitNum []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var exitNumRule []interface{}
	for _, exitNumItem := range exitNum {
		exitNumRule = append(exitNumRule, exitNumItem)
	}

	logs, sub, err := _ARBITRUMDAIBRIDGE.contract.WatchLogs(opts, "WithdrawalFinalized", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMDAIBRIDGEWithdrawalFinalized)
				if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed from, address indexed to, uint256 indexed exitNum, uint256 amount)
func (_ARBITRUMDAIBRIDGE *ARBITRUMDAIBRIDGEFilterer) ParseWithdrawalFinalized(log types.Log) (*ARBITRUMDAIBRIDGEWithdrawalFinalized, error) {
	event := new(ARBITRUMDAIBRIDGEWithdrawalFinalized)
	if err := _ARBITRUMDAIBRIDGE.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
