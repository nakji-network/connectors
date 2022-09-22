// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GUSD

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

// GUSDABI is the input ABI used to generate the binding from.
const GUSDABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emitTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"requestImplChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emitApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// GUSD is an auto generated Go binding around an Ethereum contract.
type GUSD struct {
	GUSDCaller     // Read-only binding to the contract
	GUSDTransactor // Write-only binding to the contract
	GUSDFilterer   // Log filterer for contract events
}

// GUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type GUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GUSDSession struct {
	Contract     *GUSD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GUSDCallerSession struct {
	Contract *GUSDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GUSDTransactorSession struct {
	Contract     *GUSDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type GUSDRaw struct {
	Contract *GUSD // Generic contract binding to access the raw methods on
}

// GUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GUSDCallerRaw struct {
	Contract *GUSDCaller // Generic read-only contract binding to access the raw methods on
}

// GUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GUSDTransactorRaw struct {
	Contract *GUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGUSD creates a new instance of GUSD, bound to a specific deployed contract.
func NewGUSD(address common.Address, backend bind.ContractBackend) (*GUSD, error) {
	contract, err := bindGUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GUSD{GUSDCaller: GUSDCaller{contract: contract}, GUSDTransactor: GUSDTransactor{contract: contract}, GUSDFilterer: GUSDFilterer{contract: contract}}, nil
}

// NewGUSDCaller creates a new read-only instance of GUSD, bound to a specific deployed contract.
func NewGUSDCaller(address common.Address, caller bind.ContractCaller) (*GUSDCaller, error) {
	contract, err := bindGUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GUSDCaller{contract: contract}, nil
}

// NewGUSDTransactor creates a new write-only instance of GUSD, bound to a specific deployed contract.
func NewGUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*GUSDTransactor, error) {
	contract, err := bindGUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GUSDTransactor{contract: contract}, nil
}

// NewGUSDFilterer creates a new log filterer instance of GUSD, bound to a specific deployed contract.
func NewGUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*GUSDFilterer, error) {
	contract, err := bindGUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GUSDFilterer{contract: contract}, nil
}

// bindGUSD binds a generic wrapper to an already deployed contract.
func bindGUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GUSD *GUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GUSD.Contract.GUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GUSD *GUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GUSD.Contract.GUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GUSD *GUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GUSD.Contract.GUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GUSD *GUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GUSD *GUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GUSD *GUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GUSD.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_GUSD *GUSDCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_GUSD *GUSDSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GUSD.Contract.Allowance(&_GUSD.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_GUSD *GUSDCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GUSD.Contract.Allowance(&_GUSD.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_GUSD *GUSDCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_GUSD *GUSDSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _GUSD.Contract.BalanceOf(&_GUSD.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_GUSD *GUSDCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _GUSD.Contract.BalanceOf(&_GUSD.CallOpts, _owner)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() view returns(address)
func (_GUSD *GUSDCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "custodian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() view returns(address)
func (_GUSD *GUSDSession) Custodian() (common.Address, error) {
	return _GUSD.Contract.Custodian(&_GUSD.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() view returns(address)
func (_GUSD *GUSDCallerSession) Custodian() (common.Address, error) {
	return _GUSD.Contract.Custodian(&_GUSD.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "custodianChangeReqs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _GUSD.Contract.CustodianChangeReqs(&_GUSD.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _GUSD.Contract.CustodianChangeReqs(&_GUSD.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GUSD *GUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GUSD *GUSDSession) Decimals() (uint8, error) {
	return _GUSD.Contract.Decimals(&_GUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GUSD *GUSDCallerSession) Decimals() (uint8, error) {
	return _GUSD.Contract.Decimals(&_GUSD.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() view returns(address)
func (_GUSD *GUSDCaller) Erc20Impl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "erc20Impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() view returns(address)
func (_GUSD *GUSDSession) Erc20Impl() (common.Address, error) {
	return _GUSD.Contract.Erc20Impl(&_GUSD.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() view returns(address)
func (_GUSD *GUSDCallerSession) Erc20Impl() (common.Address, error) {
	return _GUSD.Contract.Erc20Impl(&_GUSD.CallOpts)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDCaller) ImplChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "implChangeReqs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _GUSD.Contract.ImplChangeReqs(&_GUSD.CallOpts, arg0)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs(bytes32 ) view returns(address proposedNew)
func (_GUSD *GUSDCallerSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _GUSD.Contract.ImplChangeReqs(&_GUSD.CallOpts, arg0)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() view returns(uint256)
func (_GUSD *GUSDCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "lockRequestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() view returns(uint256)
func (_GUSD *GUSDSession) LockRequestCount() (*big.Int, error) {
	return _GUSD.Contract.LockRequestCount(&_GUSD.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() view returns(uint256)
func (_GUSD *GUSDCallerSession) LockRequestCount() (*big.Int, error) {
	return _GUSD.Contract.LockRequestCount(&_GUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GUSD *GUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GUSD *GUSDSession) Name() (string, error) {
	return _GUSD.Contract.Name(&_GUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GUSD *GUSDCallerSession) Name() (string, error) {
	return _GUSD.Contract.Name(&_GUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GUSD *GUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GUSD *GUSDSession) Symbol() (string, error) {
	return _GUSD.Contract.Symbol(&_GUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GUSD *GUSDCallerSession) Symbol() (string, error) {
	return _GUSD.Contract.Symbol(&_GUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GUSD *GUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GUSD.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GUSD *GUSDSession) TotalSupply() (*big.Int, error) {
	return _GUSD.Contract.TotalSupply(&_GUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GUSD *GUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _GUSD.Contract.TotalSupply(&_GUSD.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_GUSD *GUSDSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.Approve(&_GUSD.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.Approve(&_GUSD.TransactOpts, _spender, _value)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(bytes32 _lockId) returns()
func (_GUSD *GUSDTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(bytes32 _lockId) returns()
func (_GUSD *GUSDSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.Contract.ConfirmCustodianChange(&_GUSD.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(bytes32 _lockId) returns()
func (_GUSD *GUSDTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.Contract.ConfirmCustodianChange(&_GUSD.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(bytes32 _lockId) returns()
func (_GUSD *GUSDTransactor) ConfirmImplChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "confirmImplChange", _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(bytes32 _lockId) returns()
func (_GUSD *GUSDSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.Contract.ConfirmImplChange(&_GUSD.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(bytes32 _lockId) returns()
func (_GUSD *GUSDTransactorSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _GUSD.Contract.ConfirmImplChange(&_GUSD.TransactOpts, _lockId)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_GUSD *GUSDTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_GUSD *GUSDSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.DecreaseApproval(&_GUSD.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_GUSD *GUSDTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.DecreaseApproval(&_GUSD.TransactOpts, _spender, _subtractedValue)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(address _owner, address _spender, uint256 _value) returns()
func (_GUSD *GUSDTransactor) EmitApproval(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "emitApproval", _owner, _spender, _value)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(address _owner, address _spender, uint256 _value) returns()
func (_GUSD *GUSDSession) EmitApproval(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.EmitApproval(&_GUSD.TransactOpts, _owner, _spender, _value)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(address _owner, address _spender, uint256 _value) returns()
func (_GUSD *GUSDTransactorSession) EmitApproval(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.EmitApproval(&_GUSD.TransactOpts, _owner, _spender, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(address _from, address _to, uint256 _value) returns()
func (_GUSD *GUSDTransactor) EmitTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "emitTransfer", _from, _to, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(address _from, address _to, uint256 _value) returns()
func (_GUSD *GUSDSession) EmitTransfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.EmitTransfer(&_GUSD.TransactOpts, _from, _to, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(address _from, address _to, uint256 _value) returns()
func (_GUSD *GUSDTransactorSession) EmitTransfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.EmitTransfer(&_GUSD.TransactOpts, _from, _to, _value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_GUSD *GUSDTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_GUSD *GUSDSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.IncreaseApproval(&_GUSD.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_GUSD *GUSDTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.IncreaseApproval(&_GUSD.TransactOpts, _spender, _addedValue)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(address _proposedCustodian) returns(bytes32 lockId)
func (_GUSD *GUSDTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(address _proposedCustodian) returns(bytes32 lockId)
func (_GUSD *GUSDSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _GUSD.Contract.RequestCustodianChange(&_GUSD.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(address _proposedCustodian) returns(bytes32 lockId)
func (_GUSD *GUSDTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _GUSD.Contract.RequestCustodianChange(&_GUSD.TransactOpts, _proposedCustodian)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(address _proposedImpl) returns(bytes32 lockId)
func (_GUSD *GUSDTransactor) RequestImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "requestImplChange", _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(address _proposedImpl) returns(bytes32 lockId)
func (_GUSD *GUSDSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _GUSD.Contract.RequestImplChange(&_GUSD.TransactOpts, _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(address _proposedImpl) returns(bytes32 lockId)
func (_GUSD *GUSDTransactorSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _GUSD.Contract.RequestImplChange(&_GUSD.TransactOpts, _proposedImpl)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.Transfer(&_GUSD.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.Transfer(&_GUSD.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.TransferFrom(&_GUSD.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_GUSD *GUSDTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GUSD.Contract.TransferFrom(&_GUSD.TransactOpts, _from, _to, _value)
}

// GUSDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GUSD contract.
type GUSDApprovalIterator struct {
	Event *GUSDApproval // Event containing the contract specifics and raw log

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
func (it *GUSDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDApproval)
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
		it.Event = new(GUSDApproval)
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
func (it *GUSDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDApproval represents a Approval event raised by the GUSD contract.
type GUSDApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_GUSD *GUSDFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*GUSDApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &GUSDApprovalIterator{contract: _GUSD.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_GUSD *GUSDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GUSDApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDApproval)
				if err := _GUSD.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_GUSD *GUSDFilterer) ParseApproval(log types.Log) (*GUSDApproval, error) {
	event := new(GUSDApproval)
	if err := _GUSD.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUSDCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the GUSD contract.
type GUSDCustodianChangeConfirmedIterator struct {
	Event *GUSDCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *GUSDCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDCustodianChangeConfirmed)
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
		it.Event = new(GUSDCustodianChangeConfirmed)
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
func (it *GUSDCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the GUSD contract.
type GUSDCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: event CustodianChangeConfirmed(bytes32 _lockId, address _newCustodian)
func (_GUSD *GUSDFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*GUSDCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &GUSDCustodianChangeConfirmedIterator{contract: _GUSD.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: event CustodianChangeConfirmed(bytes32 _lockId, address _newCustodian)
func (_GUSD *GUSDFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *GUSDCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDCustodianChangeConfirmed)
				if err := _GUSD.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// ParseCustodianChangeConfirmed is a log parse operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: event CustodianChangeConfirmed(bytes32 _lockId, address _newCustodian)
func (_GUSD *GUSDFilterer) ParseCustodianChangeConfirmed(log types.Log) (*GUSDCustodianChangeConfirmed, error) {
	event := new(GUSDCustodianChangeConfirmed)
	if err := _GUSD.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUSDCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the GUSD contract.
type GUSDCustodianChangeRequestedIterator struct {
	Event *GUSDCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *GUSDCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDCustodianChangeRequested)
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
		it.Event = new(GUSDCustodianChangeRequested)
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
func (it *GUSDCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDCustodianChangeRequested represents a CustodianChangeRequested event raised by the GUSD contract.
type GUSDCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: event CustodianChangeRequested(bytes32 _lockId, address _msgSender, address _proposedCustodian)
func (_GUSD *GUSDFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*GUSDCustodianChangeRequestedIterator, error) {

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &GUSDCustodianChangeRequestedIterator{contract: _GUSD.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: event CustodianChangeRequested(bytes32 _lockId, address _msgSender, address _proposedCustodian)
func (_GUSD *GUSDFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *GUSDCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDCustodianChangeRequested)
				if err := _GUSD.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ParseCustodianChangeRequested is a log parse operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: event CustodianChangeRequested(bytes32 _lockId, address _msgSender, address _proposedCustodian)
func (_GUSD *GUSDFilterer) ParseCustodianChangeRequested(log types.Log) (*GUSDCustodianChangeRequested, error) {
	event := new(GUSDCustodianChangeRequested)
	if err := _GUSD.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUSDImplChangeConfirmedIterator is returned from FilterImplChangeConfirmed and is used to iterate over the raw logs and unpacked data for ImplChangeConfirmed events raised by the GUSD contract.
type GUSDImplChangeConfirmedIterator struct {
	Event *GUSDImplChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *GUSDImplChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDImplChangeConfirmed)
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
		it.Event = new(GUSDImplChangeConfirmed)
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
func (it *GUSDImplChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDImplChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDImplChangeConfirmed represents a ImplChangeConfirmed event raised by the GUSD contract.
type GUSDImplChangeConfirmed struct {
	LockId  [32]byte
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplChangeConfirmed is a free log retrieval operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: event ImplChangeConfirmed(bytes32 _lockId, address _newImpl)
func (_GUSD *GUSDFilterer) FilterImplChangeConfirmed(opts *bind.FilterOpts) (*GUSDImplChangeConfirmedIterator, error) {

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &GUSDImplChangeConfirmedIterator{contract: _GUSD.contract, event: "ImplChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchImplChangeConfirmed is a free log subscription operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: event ImplChangeConfirmed(bytes32 _lockId, address _newImpl)
func (_GUSD *GUSDFilterer) WatchImplChangeConfirmed(opts *bind.WatchOpts, sink chan<- *GUSDImplChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDImplChangeConfirmed)
				if err := _GUSD.contract.UnpackLog(event, "ImplChangeConfirmed", log); err != nil {
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

// ParseImplChangeConfirmed is a log parse operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: event ImplChangeConfirmed(bytes32 _lockId, address _newImpl)
func (_GUSD *GUSDFilterer) ParseImplChangeConfirmed(log types.Log) (*GUSDImplChangeConfirmed, error) {
	event := new(GUSDImplChangeConfirmed)
	if err := _GUSD.contract.UnpackLog(event, "ImplChangeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUSDImplChangeRequestedIterator is returned from FilterImplChangeRequested and is used to iterate over the raw logs and unpacked data for ImplChangeRequested events raised by the GUSD contract.
type GUSDImplChangeRequestedIterator struct {
	Event *GUSDImplChangeRequested // Event containing the contract specifics and raw log

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
func (it *GUSDImplChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDImplChangeRequested)
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
		it.Event = new(GUSDImplChangeRequested)
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
func (it *GUSDImplChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDImplChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDImplChangeRequested represents a ImplChangeRequested event raised by the GUSD contract.
type GUSDImplChangeRequested struct {
	LockId       [32]byte
	MsgSender    common.Address
	ProposedImpl common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImplChangeRequested is a free log retrieval operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: event ImplChangeRequested(bytes32 _lockId, address _msgSender, address _proposedImpl)
func (_GUSD *GUSDFilterer) FilterImplChangeRequested(opts *bind.FilterOpts) (*GUSDImplChangeRequestedIterator, error) {

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return &GUSDImplChangeRequestedIterator{contract: _GUSD.contract, event: "ImplChangeRequested", logs: logs, sub: sub}, nil
}

// WatchImplChangeRequested is a free log subscription operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: event ImplChangeRequested(bytes32 _lockId, address _msgSender, address _proposedImpl)
func (_GUSD *GUSDFilterer) WatchImplChangeRequested(opts *bind.WatchOpts, sink chan<- *GUSDImplChangeRequested) (event.Subscription, error) {

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDImplChangeRequested)
				if err := _GUSD.contract.UnpackLog(event, "ImplChangeRequested", log); err != nil {
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

// ParseImplChangeRequested is a log parse operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: event ImplChangeRequested(bytes32 _lockId, address _msgSender, address _proposedImpl)
func (_GUSD *GUSDFilterer) ParseImplChangeRequested(log types.Log) (*GUSDImplChangeRequested, error) {
	event := new(GUSDImplChangeRequested)
	if err := _GUSD.contract.UnpackLog(event, "ImplChangeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUSDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GUSD contract.
type GUSDTransferIterator struct {
	Event *GUSDTransfer // Event containing the contract specifics and raw log

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
func (it *GUSDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUSDTransfer)
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
		it.Event = new(GUSDTransfer)
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
func (it *GUSDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUSDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUSDTransfer represents a Transfer event raised by the GUSD contract.
type GUSDTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_GUSD *GUSDFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*GUSDTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _GUSD.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &GUSDTransferIterator{contract: _GUSD.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_GUSD *GUSDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GUSDTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _GUSD.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUSDTransfer)
				if err := _GUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_GUSD *GUSDFilterer) ParseTransfer(log types.Log) (*GUSDTransfer, error) {
	event := new(GUSDTransfer)
	if err := _GUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
