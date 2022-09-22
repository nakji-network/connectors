// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joebar

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

// JoebarMetaData contains all meta data concerning the Joebar contract.
var JoebarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_joe\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joe\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_share\",\"type\":\"uint256\"}],\"name\":\"leave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// JoebarABI is the input ABI used to generate the binding from.
// Deprecated: Use JoebarMetaData.ABI instead.
var JoebarABI = JoebarMetaData.ABI

// Joebar is an auto generated Go binding around an Ethereum contract.
type Joebar struct {
	JoebarCaller     // Read-only binding to the contract
	JoebarTransactor // Write-only binding to the contract
	JoebarFilterer   // Log filterer for contract events
}

// JoebarCaller is an auto generated read-only Go binding around an Ethereum contract.
type JoebarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoebarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JoebarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoebarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JoebarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoebarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JoebarSession struct {
	Contract     *Joebar           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoebarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JoebarCallerSession struct {
	Contract *JoebarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JoebarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JoebarTransactorSession struct {
	Contract     *JoebarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoebarRaw is an auto generated low-level Go binding around an Ethereum contract.
type JoebarRaw struct {
	Contract *Joebar // Generic contract binding to access the raw methods on
}

// JoebarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JoebarCallerRaw struct {
	Contract *JoebarCaller // Generic read-only contract binding to access the raw methods on
}

// JoebarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JoebarTransactorRaw struct {
	Contract *JoebarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJoebar creates a new instance of Joebar, bound to a specific deployed contract.
func NewJoebar(address common.Address, backend bind.ContractBackend) (*Joebar, error) {
	contract, err := bindJoebar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Joebar{JoebarCaller: JoebarCaller{contract: contract}, JoebarTransactor: JoebarTransactor{contract: contract}, JoebarFilterer: JoebarFilterer{contract: contract}}, nil
}

// NewJoebarCaller creates a new read-only instance of Joebar, bound to a specific deployed contract.
func NewJoebarCaller(address common.Address, caller bind.ContractCaller) (*JoebarCaller, error) {
	contract, err := bindJoebar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JoebarCaller{contract: contract}, nil
}

// NewJoebarTransactor creates a new write-only instance of Joebar, bound to a specific deployed contract.
func NewJoebarTransactor(address common.Address, transactor bind.ContractTransactor) (*JoebarTransactor, error) {
	contract, err := bindJoebar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JoebarTransactor{contract: contract}, nil
}

// NewJoebarFilterer creates a new log filterer instance of Joebar, bound to a specific deployed contract.
func NewJoebarFilterer(address common.Address, filterer bind.ContractFilterer) (*JoebarFilterer, error) {
	contract, err := bindJoebar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JoebarFilterer{contract: contract}, nil
}

// bindJoebar binds a generic wrapper to an already deployed contract.
func bindJoebar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JoebarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joebar *JoebarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joebar.Contract.JoebarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joebar *JoebarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joebar.Contract.JoebarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joebar *JoebarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joebar.Contract.JoebarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joebar *JoebarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joebar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joebar *JoebarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joebar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joebar *JoebarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joebar.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Joebar *JoebarCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Joebar *JoebarSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Joebar.Contract.Allowance(&_Joebar.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Joebar *JoebarCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Joebar.Contract.Allowance(&_Joebar.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Joebar *JoebarCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Joebar *JoebarSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Joebar.Contract.BalanceOf(&_Joebar.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Joebar *JoebarCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Joebar.Contract.BalanceOf(&_Joebar.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Joebar *JoebarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Joebar *JoebarSession) Decimals() (uint8, error) {
	return _Joebar.Contract.Decimals(&_Joebar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Joebar *JoebarCallerSession) Decimals() (uint8, error) {
	return _Joebar.Contract.Decimals(&_Joebar.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Joebar *JoebarCaller) Joe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "joe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Joebar *JoebarSession) Joe() (common.Address, error) {
	return _Joebar.Contract.Joe(&_Joebar.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Joebar *JoebarCallerSession) Joe() (common.Address, error) {
	return _Joebar.Contract.Joe(&_Joebar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Joebar *JoebarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Joebar *JoebarSession) Name() (string, error) {
	return _Joebar.Contract.Name(&_Joebar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Joebar *JoebarCallerSession) Name() (string, error) {
	return _Joebar.Contract.Name(&_Joebar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Joebar *JoebarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Joebar *JoebarSession) Symbol() (string, error) {
	return _Joebar.Contract.Symbol(&_Joebar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Joebar *JoebarCallerSession) Symbol() (string, error) {
	return _Joebar.Contract.Symbol(&_Joebar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Joebar *JoebarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Joebar.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Joebar *JoebarSession) TotalSupply() (*big.Int, error) {
	return _Joebar.Contract.TotalSupply(&_Joebar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Joebar *JoebarCallerSession) TotalSupply() (*big.Int, error) {
	return _Joebar.Contract.TotalSupply(&_Joebar.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Joebar *JoebarSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Approve(&_Joebar.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Approve(&_Joebar.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Joebar *JoebarTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Joebar *JoebarSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.DecreaseAllowance(&_Joebar.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Joebar *JoebarTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.DecreaseAllowance(&_Joebar.TransactOpts, spender, subtractedValue)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Joebar *JoebarTransactor) Enter(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "enter", _amount)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Joebar *JoebarSession) Enter(_amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Enter(&_Joebar.TransactOpts, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Joebar *JoebarTransactorSession) Enter(_amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Enter(&_Joebar.TransactOpts, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Joebar *JoebarTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Joebar *JoebarSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.IncreaseAllowance(&_Joebar.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Joebar *JoebarTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.IncreaseAllowance(&_Joebar.TransactOpts, spender, addedValue)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Joebar *JoebarTransactor) Leave(opts *bind.TransactOpts, _share *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "leave", _share)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Joebar *JoebarSession) Leave(_share *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Leave(&_Joebar.TransactOpts, _share)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Joebar *JoebarTransactorSession) Leave(_share *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Leave(&_Joebar.TransactOpts, _share)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Transfer(&_Joebar.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.Transfer(&_Joebar.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.TransferFrom(&_Joebar.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Joebar *JoebarTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Joebar.Contract.TransferFrom(&_Joebar.TransactOpts, sender, recipient, amount)
}

// JoebarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Joebar contract.
type JoebarApprovalIterator struct {
	Event *JoebarApproval // Event containing the contract specifics and raw log

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
func (it *JoebarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoebarApproval)
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
		it.Event = new(JoebarApproval)
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
func (it *JoebarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoebarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoebarApproval represents a Approval event raised by the Joebar contract.
type JoebarApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Joebar *JoebarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JoebarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Joebar.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JoebarApprovalIterator{contract: _Joebar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Joebar *JoebarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JoebarApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Joebar.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoebarApproval)
				if err := _Joebar.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Joebar *JoebarFilterer) ParseApproval(log types.Log) (*JoebarApproval, error) {
	event := new(JoebarApproval)
	if err := _Joebar.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoebarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Joebar contract.
type JoebarTransferIterator struct {
	Event *JoebarTransfer // Event containing the contract specifics and raw log

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
func (it *JoebarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoebarTransfer)
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
		it.Event = new(JoebarTransfer)
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
func (it *JoebarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoebarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoebarTransfer represents a Transfer event raised by the Joebar contract.
type JoebarTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Joebar *JoebarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JoebarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Joebar.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JoebarTransferIterator{contract: _Joebar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Joebar *JoebarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JoebarTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Joebar.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoebarTransfer)
				if err := _Joebar.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Joebar *JoebarFilterer) ParseTransfer(log types.Log) (*JoebarTransfer, error) {
	event := new(JoebarTransfer)
	if err := _Joebar.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
