// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KNC

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

// KNCABI is the input ABI used to generate the binding from.
const KNCABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenSaleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyERC20Drain\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleEndTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenTotalAmount\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\"},{\"name\":\"admin\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// KNC is an auto generated Go binding around an Ethereum contract.
type KNC struct {
	KNCCaller     // Read-only binding to the contract
	KNCTransactor // Write-only binding to the contract
	KNCFilterer   // Log filterer for contract events
}

// KNCCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNCSession struct {
	Contract     *KNC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNCCallerSession struct {
	Contract *KNCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KNCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNCTransactorSession struct {
	Contract     *KNCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNCRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNCRaw struct {
	Contract *KNC // Generic contract binding to access the raw methods on
}

// KNCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNCCallerRaw struct {
	Contract *KNCCaller // Generic read-only contract binding to access the raw methods on
}

// KNCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNCTransactorRaw struct {
	Contract *KNCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNC creates a new instance of KNC, bound to a specific deployed contract.
func NewKNC(address common.Address, backend bind.ContractBackend) (*KNC, error) {
	contract, err := bindKNC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNC{KNCCaller: KNCCaller{contract: contract}, KNCTransactor: KNCTransactor{contract: contract}, KNCFilterer: KNCFilterer{contract: contract}}, nil
}

// NewKNCCaller creates a new read-only instance of KNC, bound to a specific deployed contract.
func NewKNCCaller(address common.Address, caller bind.ContractCaller) (*KNCCaller, error) {
	contract, err := bindKNC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNCCaller{contract: contract}, nil
}

// NewKNCTransactor creates a new write-only instance of KNC, bound to a specific deployed contract.
func NewKNCTransactor(address common.Address, transactor bind.ContractTransactor) (*KNCTransactor, error) {
	contract, err := bindKNC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNCTransactor{contract: contract}, nil
}

// NewKNCFilterer creates a new log filterer instance of KNC, bound to a specific deployed contract.
func NewKNCFilterer(address common.Address, filterer bind.ContractFilterer) (*KNCFilterer, error) {
	contract, err := bindKNC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNCFilterer{contract: contract}, nil
}

// bindKNC binds a generic wrapper to an already deployed contract.
func bindKNC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNC *KNCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KNC.Contract.KNCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNC *KNCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNC.Contract.KNCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNC *KNCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNC.Contract.KNCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNC *KNCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KNC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNC *KNCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNC *KNCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_KNC *KNCCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_KNC *KNCSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _KNC.Contract.Allowance(&_KNC.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_KNC *KNCCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _KNC.Contract.Allowance(&_KNC.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_KNC *KNCCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_KNC *KNCSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KNC.Contract.BalanceOf(&_KNC.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_KNC *KNCCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KNC.Contract.BalanceOf(&_KNC.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_KNC *KNCCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_KNC *KNCSession) Decimals() (*big.Int, error) {
	return _KNC.Contract.Decimals(&_KNC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_KNC *KNCCallerSession) Decimals() (*big.Int, error) {
	return _KNC.Contract.Decimals(&_KNC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_KNC *KNCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_KNC *KNCSession) Name() (string, error) {
	return _KNC.Contract.Name(&_KNC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_KNC *KNCCallerSession) Name() (string, error) {
	return _KNC.Contract.Name(&_KNC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_KNC *KNCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_KNC *KNCSession) Owner() (common.Address, error) {
	return _KNC.Contract.Owner(&_KNC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_KNC *KNCCallerSession) Owner() (common.Address, error) {
	return _KNC.Contract.Owner(&_KNC.CallOpts)
}

// SaleEndTime is a free data retrieval call binding the contract method 0xed338ff1.
//
// Solidity: function saleEndTime() returns(uint256)
func (_KNC *KNCCaller) SaleEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "saleEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleEndTime is a free data retrieval call binding the contract method 0xed338ff1.
//
// Solidity: function saleEndTime() returns(uint256)
func (_KNC *KNCSession) SaleEndTime() (*big.Int, error) {
	return _KNC.Contract.SaleEndTime(&_KNC.CallOpts)
}

// SaleEndTime is a free data retrieval call binding the contract method 0xed338ff1.
//
// Solidity: function saleEndTime() returns(uint256)
func (_KNC *KNCCallerSession) SaleEndTime() (*big.Int, error) {
	return _KNC.Contract.SaleEndTime(&_KNC.CallOpts)
}

// SaleStartTime is a free data retrieval call binding the contract method 0x1cbaee2d.
//
// Solidity: function saleStartTime() returns(uint256)
func (_KNC *KNCCaller) SaleStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "saleStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleStartTime is a free data retrieval call binding the contract method 0x1cbaee2d.
//
// Solidity: function saleStartTime() returns(uint256)
func (_KNC *KNCSession) SaleStartTime() (*big.Int, error) {
	return _KNC.Contract.SaleStartTime(&_KNC.CallOpts)
}

// SaleStartTime is a free data retrieval call binding the contract method 0x1cbaee2d.
//
// Solidity: function saleStartTime() returns(uint256)
func (_KNC *KNCCallerSession) SaleStartTime() (*big.Int, error) {
	return _KNC.Contract.SaleStartTime(&_KNC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_KNC *KNCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_KNC *KNCSession) Symbol() (string, error) {
	return _KNC.Contract.Symbol(&_KNC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_KNC *KNCCallerSession) Symbol() (string, error) {
	return _KNC.Contract.Symbol(&_KNC.CallOpts)
}

// TokenSaleContract is a free data retrieval call binding the contract method 0x5d5aa277.
//
// Solidity: function tokenSaleContract() returns(address)
func (_KNC *KNCCaller) TokenSaleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "tokenSaleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSaleContract is a free data retrieval call binding the contract method 0x5d5aa277.
//
// Solidity: function tokenSaleContract() returns(address)
func (_KNC *KNCSession) TokenSaleContract() (common.Address, error) {
	return _KNC.Contract.TokenSaleContract(&_KNC.CallOpts)
}

// TokenSaleContract is a free data retrieval call binding the contract method 0x5d5aa277.
//
// Solidity: function tokenSaleContract() returns(address)
func (_KNC *KNCCallerSession) TokenSaleContract() (common.Address, error) {
	return _KNC.Contract.TokenSaleContract(&_KNC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_KNC *KNCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KNC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_KNC *KNCSession) TotalSupply() (*big.Int, error) {
	return _KNC.Contract.TotalSupply(&_KNC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_KNC *KNCCallerSession) TotalSupply() (*big.Int, error) {
	return _KNC.Contract.TotalSupply(&_KNC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_KNC *KNCTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_KNC *KNCSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Approve(&_KNC.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_KNC *KNCTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Approve(&_KNC.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_KNC *KNCTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_KNC *KNCSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Burn(&_KNC.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_KNC *KNCTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Burn(&_KNC.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool)
func (_KNC *KNCTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool)
func (_KNC *KNCSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.BurnFrom(&_KNC.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool)
func (_KNC *KNCTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.BurnFrom(&_KNC.TransactOpts, _from, _value)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(address token, uint256 amount) returns()
func (_KNC *KNCTransactor) EmergencyERC20Drain(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "emergencyERC20Drain", token, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(address token, uint256 amount) returns()
func (_KNC *KNCSession) EmergencyERC20Drain(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.EmergencyERC20Drain(&_KNC.TransactOpts, token, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(address token, uint256 amount) returns()
func (_KNC *KNCTransactorSession) EmergencyERC20Drain(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.EmergencyERC20Drain(&_KNC.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KNC *KNCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KNC *KNCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Transfer(&_KNC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KNC *KNCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.Transfer(&_KNC.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KNC *KNCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KNC *KNCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.TransferFrom(&_KNC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KNC *KNCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KNC.Contract.TransferFrom(&_KNC.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KNC *KNCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KNC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KNC *KNCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KNC.Contract.TransferOwnership(&_KNC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KNC *KNCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KNC.Contract.TransferOwnership(&_KNC.TransactOpts, newOwner)
}

// KNCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KNC contract.
type KNCApprovalIterator struct {
	Event *KNCApproval // Event containing the contract specifics and raw log

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
func (it *KNCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNCApproval)
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
		it.Event = new(KNCApproval)
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
func (it *KNCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNCApproval represents a Approval event raised by the KNC contract.
type KNCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_KNC *KNCFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*KNCApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _KNC.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &KNCApprovalIterator{contract: _KNC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_KNC *KNCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KNCApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _KNC.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNCApproval)
				if err := _KNC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KNC *KNCFilterer) ParseApproval(log types.Log) (*KNCApproval, error) {
	event := new(KNCApproval)
	if err := _KNC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KNCBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the KNC contract.
type KNCBurnIterator struct {
	Event *KNCBurn // Event containing the contract specifics and raw log

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
func (it *KNCBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNCBurn)
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
		it.Event = new(KNCBurn)
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
func (it *KNCBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNCBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNCBurn represents a Burn event raised by the KNC contract.
type KNCBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _burner, uint256 _value)
func (_KNC *KNCFilterer) FilterBurn(opts *bind.FilterOpts, _burner []common.Address) (*KNCBurnIterator, error) {

	var _burnerRule []interface{}
	for _, _burnerItem := range _burner {
		_burnerRule = append(_burnerRule, _burnerItem)
	}

	logs, sub, err := _KNC.contract.FilterLogs(opts, "Burn", _burnerRule)
	if err != nil {
		return nil, err
	}
	return &KNCBurnIterator{contract: _KNC.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _burner, uint256 _value)
func (_KNC *KNCFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *KNCBurn, _burner []common.Address) (event.Subscription, error) {

	var _burnerRule []interface{}
	for _, _burnerItem := range _burner {
		_burnerRule = append(_burnerRule, _burnerItem)
	}

	logs, sub, err := _KNC.contract.WatchLogs(opts, "Burn", _burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNCBurn)
				if err := _KNC.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _burner, uint256 _value)
func (_KNC *KNCFilterer) ParseBurn(log types.Log) (*KNCBurn, error) {
	event := new(KNCBurn)
	if err := _KNC.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KNCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KNC contract.
type KNCTransferIterator struct {
	Event *KNCTransfer // Event containing the contract specifics and raw log

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
func (it *KNCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNCTransfer)
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
		it.Event = new(KNCTransfer)
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
func (it *KNCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNCTransfer represents a Transfer event raised by the KNC contract.
type KNCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_KNC *KNCFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*KNCTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KNC.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &KNCTransferIterator{contract: _KNC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_KNC *KNCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KNCTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KNC.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNCTransfer)
				if err := _KNC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KNC *KNCFilterer) ParseTransfer(log types.Log) (*KNCTransfer, error) {
	event := new(KNCTransfer)
	if err := _KNC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
