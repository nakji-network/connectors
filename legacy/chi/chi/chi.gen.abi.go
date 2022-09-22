// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChiMetaData contains all meta data concerning the Chi contract.
var ChiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"computeAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeFromUpTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeUpTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChiABI is the input ABI used to generate the binding from.
// Deprecated: Use ChiMetaData.ABI instead.
var ChiABI = ChiMetaData.ABI

// Chi is an auto generated Go binding around an Ethereum contract.
type Chi struct {
	ChiCaller     // Read-only binding to the contract
	ChiTransactor // Write-only binding to the contract
	ChiFilterer   // Log filterer for contract events
}

// ChiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiSession struct {
	Contract     *Chi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiCallerSession struct {
	Contract *ChiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChiTransactorSession struct {
	Contract     *ChiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChiRaw struct {
	Contract *Chi // Generic contract binding to access the raw methods on
}

// ChiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChiCallerRaw struct {
	Contract *ChiCaller // Generic read-only contract binding to access the raw methods on
}

// ChiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChiTransactorRaw struct {
	Contract *ChiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChi creates a new instance of Chi, bound to a specific deployed contract.
func NewChi(address common.Address, backend bind.ContractBackend) (*Chi, error) {
	contract, err := bindChi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chi{ChiCaller: ChiCaller{contract: contract}, ChiTransactor: ChiTransactor{contract: contract}, ChiFilterer: ChiFilterer{contract: contract}}, nil
}

// NewChiCaller creates a new read-only instance of Chi, bound to a specific deployed contract.
func NewChiCaller(address common.Address, caller bind.ContractCaller) (*ChiCaller, error) {
	contract, err := bindChi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChiCaller{contract: contract}, nil
}

// NewChiTransactor creates a new write-only instance of Chi, bound to a specific deployed contract.
func NewChiTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiTransactor, error) {
	contract, err := bindChi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChiTransactor{contract: contract}, nil
}

// NewChiFilterer creates a new log filterer instance of Chi, bound to a specific deployed contract.
func NewChiFilterer(address common.Address, filterer bind.ContractFilterer) (*ChiFilterer, error) {
	contract, err := bindChi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChiFilterer{contract: contract}, nil
}

// bindChi binds a generic wrapper to an already deployed contract.
func bindChi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chi *ChiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chi.Contract.ChiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chi *ChiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chi.Contract.ChiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chi *ChiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chi.Contract.ChiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chi *ChiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chi *ChiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chi *ChiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chi.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Chi *ChiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Chi *ChiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Chi.Contract.Allowance(&_Chi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Chi *ChiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Chi.Contract.Allowance(&_Chi.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Chi *ChiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Chi *ChiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Chi.Contract.BalanceOf(&_Chi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Chi *ChiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Chi.Contract.BalanceOf(&_Chi.CallOpts, account)
}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) view returns(address)
func (_Chi *ChiCaller) ComputeAddress2(opts *bind.CallOpts, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "computeAddress2", salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) view returns(address)
func (_Chi *ChiSession) ComputeAddress2(salt *big.Int) (common.Address, error) {
	return _Chi.Contract.ComputeAddress2(&_Chi.CallOpts, salt)
}

// ComputeAddress2 is a free data retrieval call binding the contract method 0xb0ac19a0.
//
// Solidity: function computeAddress2(uint256 salt) view returns(address)
func (_Chi *ChiCallerSession) ComputeAddress2(salt *big.Int) (common.Address, error) {
	return _Chi.Contract.ComputeAddress2(&_Chi.CallOpts, salt)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chi *ChiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chi *ChiSession) Decimals() (uint8, error) {
	return _Chi.Contract.Decimals(&_Chi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chi *ChiCallerSession) Decimals() (uint8, error) {
	return _Chi.Contract.Decimals(&_Chi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chi *ChiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chi *ChiSession) Name() (string, error) {
	return _Chi.Contract.Name(&_Chi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chi *ChiCallerSession) Name() (string, error) {
	return _Chi.Contract.Name(&_Chi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chi *ChiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chi *ChiSession) Symbol() (string, error) {
	return _Chi.Contract.Symbol(&_Chi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chi *ChiCallerSession) Symbol() (string, error) {
	return _Chi.Contract.Symbol(&_Chi.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Chi *ChiCaller) TotalBurned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "totalBurned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Chi *ChiSession) TotalBurned() (*big.Int, error) {
	return _Chi.Contract.TotalBurned(&_Chi.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Chi *ChiCallerSession) TotalBurned() (*big.Int, error) {
	return _Chi.Contract.TotalBurned(&_Chi.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Chi *ChiCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Chi *ChiSession) TotalMinted() (*big.Int, error) {
	return _Chi.Contract.TotalMinted(&_Chi.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Chi *ChiCallerSession) TotalMinted() (*big.Int, error) {
	return _Chi.Contract.TotalMinted(&_Chi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chi *ChiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chi *ChiSession) TotalSupply() (*big.Int, error) {
	return _Chi.Contract.TotalSupply(&_Chi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chi *ChiCallerSession) TotalSupply() (*big.Int, error) {
	return _Chi.Contract.TotalSupply(&_Chi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Chi *ChiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Chi *ChiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Approve(&_Chi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Chi *ChiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Approve(&_Chi.TransactOpts, spender, amount)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_Chi *ChiTransactor) Free(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "free", value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_Chi *ChiSession) Free(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Free(&_Chi.TransactOpts, value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(uint256)
func (_Chi *ChiTransactorSession) Free(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Free(&_Chi.TransactOpts, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_Chi *ChiTransactor) FreeFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "freeFrom", from, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_Chi *ChiSession) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeFrom(&_Chi.TransactOpts, from, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(uint256)
func (_Chi *ChiTransactorSession) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeFrom(&_Chi.TransactOpts, from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_Chi *ChiTransactor) FreeFromUpTo(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "freeFromUpTo", from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_Chi *ChiSession) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeFromUpTo(&_Chi.TransactOpts, from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256)
func (_Chi *ChiTransactorSession) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeFromUpTo(&_Chi.TransactOpts, from, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_Chi *ChiTransactor) FreeUpTo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "freeUpTo", value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_Chi *ChiSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeUpTo(&_Chi.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256)
func (_Chi *ChiTransactorSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.FreeUpTo(&_Chi.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_Chi *ChiTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_Chi *ChiSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Mint(&_Chi.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_Chi *ChiTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Mint(&_Chi.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Chi *ChiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Chi *ChiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Transfer(&_Chi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Chi *ChiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.Transfer(&_Chi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Chi *ChiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Chi *ChiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.TransferFrom(&_Chi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Chi *ChiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Chi.Contract.TransferFrom(&_Chi.TransactOpts, sender, recipient, amount)
}

// ChiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Chi contract.
type ChiApprovalIterator struct {
	Event *ChiApproval // Event containing the contract specifics and raw log

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
func (it *ChiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChiApproval)
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
		it.Event = new(ChiApproval)
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
func (it *ChiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChiApproval represents a Approval event raised by the Chi contract.
type ChiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Chi *ChiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Chi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChiApprovalIterator{contract: _Chi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Chi *ChiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Chi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChiApproval)
				if err := _Chi.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Chi *ChiFilterer) ParseApproval(log types.Log) (*ChiApproval, error) {
	event := new(ChiApproval)
	if err := _Chi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Chi contract.
type ChiTransferIterator struct {
	Event *ChiTransfer // Event containing the contract specifics and raw log

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
func (it *ChiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChiTransfer)
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
		it.Event = new(ChiTransfer)
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
func (it *ChiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChiTransfer represents a Transfer event raised by the Chi contract.
type ChiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Chi *ChiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Chi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChiTransferIterator{contract: _Chi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Chi *ChiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Chi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChiTransfer)
				if err := _Chi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Chi *ChiFilterer) ParseTransfer(log types.Log) (*ChiTransfer, error) {
	event := new(ChiTransfer)
	if err := _Chi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
