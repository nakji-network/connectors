// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CRVV1ETHSTETH

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

// CRVV1ETHSTETHABI is the input ABI used to generate the binding from.
const CRVV1ETHSTETHABI = "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":261},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74713},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111355},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37794},{\"name\":\"increaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_added_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39038},{\"name\":\"decreaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_subtracted_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39062},{\"name\":\"mint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":75652},{\"name\":\"burnFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":75670},{\"name\":\"set_minter\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36458},{\"name\":\"set_name\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":178219},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7763},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6816},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1636},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1881},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1481},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1511}]"

// CRVV1ETHSTETH is an auto generated Go binding around an Ethereum contract.
type CRVV1ETHSTETH struct {
	CRVV1ETHSTETHCaller     // Read-only binding to the contract
	CRVV1ETHSTETHTransactor // Write-only binding to the contract
	CRVV1ETHSTETHFilterer   // Log filterer for contract events
}

// CRVV1ETHSTETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type CRVV1ETHSTETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRVV1ETHSTETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CRVV1ETHSTETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRVV1ETHSTETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CRVV1ETHSTETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRVV1ETHSTETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CRVV1ETHSTETHSession struct {
	Contract     *CRVV1ETHSTETH    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CRVV1ETHSTETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CRVV1ETHSTETHCallerSession struct {
	Contract *CRVV1ETHSTETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CRVV1ETHSTETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CRVV1ETHSTETHTransactorSession struct {
	Contract     *CRVV1ETHSTETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CRVV1ETHSTETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type CRVV1ETHSTETHRaw struct {
	Contract *CRVV1ETHSTETH // Generic contract binding to access the raw methods on
}

// CRVV1ETHSTETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CRVV1ETHSTETHCallerRaw struct {
	Contract *CRVV1ETHSTETHCaller // Generic read-only contract binding to access the raw methods on
}

// CRVV1ETHSTETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CRVV1ETHSTETHTransactorRaw struct {
	Contract *CRVV1ETHSTETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCRVV1ETHSTETH creates a new instance of CRVV1ETHSTETH, bound to a specific deployed contract.
func NewCRVV1ETHSTETH(address common.Address, backend bind.ContractBackend) (*CRVV1ETHSTETH, error) {
	contract, err := bindCRVV1ETHSTETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETH{CRVV1ETHSTETHCaller: CRVV1ETHSTETHCaller{contract: contract}, CRVV1ETHSTETHTransactor: CRVV1ETHSTETHTransactor{contract: contract}, CRVV1ETHSTETHFilterer: CRVV1ETHSTETHFilterer{contract: contract}}, nil
}

// NewCRVV1ETHSTETHCaller creates a new read-only instance of CRVV1ETHSTETH, bound to a specific deployed contract.
func NewCRVV1ETHSTETHCaller(address common.Address, caller bind.ContractCaller) (*CRVV1ETHSTETHCaller, error) {
	contract, err := bindCRVV1ETHSTETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETHCaller{contract: contract}, nil
}

// NewCRVV1ETHSTETHTransactor creates a new write-only instance of CRVV1ETHSTETH, bound to a specific deployed contract.
func NewCRVV1ETHSTETHTransactor(address common.Address, transactor bind.ContractTransactor) (*CRVV1ETHSTETHTransactor, error) {
	contract, err := bindCRVV1ETHSTETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETHTransactor{contract: contract}, nil
}

// NewCRVV1ETHSTETHFilterer creates a new log filterer instance of CRVV1ETHSTETH, bound to a specific deployed contract.
func NewCRVV1ETHSTETHFilterer(address common.Address, filterer bind.ContractFilterer) (*CRVV1ETHSTETHFilterer, error) {
	contract, err := bindCRVV1ETHSTETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETHFilterer{contract: contract}, nil
}

// bindCRVV1ETHSTETH binds a generic wrapper to an already deployed contract.
func bindCRVV1ETHSTETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CRVV1ETHSTETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CRVV1ETHSTETH.Contract.CRVV1ETHSTETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.CRVV1ETHSTETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.CRVV1ETHSTETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CRVV1ETHSTETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.Allowance(&_CRVV1ETHSTETH.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.Allowance(&_CRVV1ETHSTETH.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.BalanceOf(&_CRVV1ETHSTETH.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.BalanceOf(&_CRVV1ETHSTETH.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Decimals() (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.Decimals(&_CRVV1ETHSTETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) Decimals() (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.Decimals(&_CRVV1ETHSTETH.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Minter() (common.Address, error) {
	return _CRVV1ETHSTETH.Contract.Minter(&_CRVV1ETHSTETH.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) Minter() (common.Address, error) {
	return _CRVV1ETHSTETH.Contract.Minter(&_CRVV1ETHSTETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Name() (string, error) {
	return _CRVV1ETHSTETH.Contract.Name(&_CRVV1ETHSTETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) Name() (string, error) {
	return _CRVV1ETHSTETH.Contract.Name(&_CRVV1ETHSTETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Symbol() (string, error) {
	return _CRVV1ETHSTETH.Contract.Symbol(&_CRVV1ETHSTETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) Symbol() (string, error) {
	return _CRVV1ETHSTETH.Contract.Symbol(&_CRVV1ETHSTETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CRVV1ETHSTETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) TotalSupply() (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.TotalSupply(&_CRVV1ETHSTETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHCallerSession) TotalSupply() (*big.Int, error) {
	return _CRVV1ETHSTETH.Contract.TotalSupply(&_CRVV1ETHSTETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Approve(&_CRVV1ETHSTETH.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Approve(&_CRVV1ETHSTETH.TransactOpts, _spender, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) BurnFrom(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "burnFrom", _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.BurnFrom(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.BurnFrom(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.DecreaseAllowance(&_CRVV1ETHSTETH.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.DecreaseAllowance(&_CRVV1ETHSTETH.TransactOpts, _spender, _subtracted_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.IncreaseAllowance(&_CRVV1ETHSTETH.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.IncreaseAllowance(&_CRVV1ETHSTETH.TransactOpts, _spender, _added_value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Mint(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Mint(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.SetMinter(&_CRVV1ETHSTETH.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.SetMinter(&_CRVV1ETHSTETH.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.SetName(&_CRVV1ETHSTETH.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.SetName(&_CRVV1ETHSTETH.TransactOpts, _name, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Transfer(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.Transfer(&_CRVV1ETHSTETH.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.TransferFrom(&_CRVV1ETHSTETH.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CRVV1ETHSTETH.Contract.TransferFrom(&_CRVV1ETHSTETH.TransactOpts, _from, _to, _value)
}

// CRVV1ETHSTETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CRVV1ETHSTETH contract.
type CRVV1ETHSTETHApprovalIterator struct {
	Event *CRVV1ETHSTETHApproval // Event containing the contract specifics and raw log

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
func (it *CRVV1ETHSTETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CRVV1ETHSTETHApproval)
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
		it.Event = new(CRVV1ETHSTETHApproval)
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
func (it *CRVV1ETHSTETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CRVV1ETHSTETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CRVV1ETHSTETHApproval represents a Approval event raised by the CRVV1ETHSTETH contract.
type CRVV1ETHSTETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CRVV1ETHSTETHApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CRVV1ETHSTETH.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETHApprovalIterator{contract: _CRVV1ETHSTETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CRVV1ETHSTETHApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CRVV1ETHSTETH.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CRVV1ETHSTETHApproval)
				if err := _CRVV1ETHSTETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) ParseApproval(log types.Log) (*CRVV1ETHSTETHApproval, error) {
	event := new(CRVV1ETHSTETHApproval)
	if err := _CRVV1ETHSTETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CRVV1ETHSTETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CRVV1ETHSTETH contract.
type CRVV1ETHSTETHTransferIterator struct {
	Event *CRVV1ETHSTETHTransfer // Event containing the contract specifics and raw log

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
func (it *CRVV1ETHSTETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CRVV1ETHSTETHTransfer)
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
		it.Event = new(CRVV1ETHSTETHTransfer)
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
func (it *CRVV1ETHSTETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CRVV1ETHSTETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CRVV1ETHSTETHTransfer represents a Transfer event raised by the CRVV1ETHSTETH contract.
type CRVV1ETHSTETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CRVV1ETHSTETHTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CRVV1ETHSTETH.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CRVV1ETHSTETHTransferIterator{contract: _CRVV1ETHSTETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CRVV1ETHSTETHTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CRVV1ETHSTETH.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CRVV1ETHSTETHTransfer)
				if err := _CRVV1ETHSTETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CRVV1ETHSTETH *CRVV1ETHSTETHFilterer) ParseTransfer(log types.Log) (*CRVV1ETHSTETHTransfer, error) {
	event := new(CRVV1ETHSTETHTransfer)
	if err := _CRVV1ETHSTETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
