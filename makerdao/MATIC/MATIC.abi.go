// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MATIC

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

// MATICABI is the input ABI used to generate the binding from.
const MATICABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// MATIC is an auto generated Go binding around an Ethereum contract.
type MATIC struct {
	MATICCaller     // Read-only binding to the contract
	MATICTransactor // Write-only binding to the contract
	MATICFilterer   // Log filterer for contract events
}

// MATICCaller is an auto generated read-only Go binding around an Ethereum contract.
type MATICCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MATICTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MATICTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MATICFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MATICFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MATICSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MATICSession struct {
	Contract     *MATIC            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MATICCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MATICCallerSession struct {
	Contract *MATICCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MATICTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MATICTransactorSession struct {
	Contract     *MATICTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MATICRaw is an auto generated low-level Go binding around an Ethereum contract.
type MATICRaw struct {
	Contract *MATIC // Generic contract binding to access the raw methods on
}

// MATICCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MATICCallerRaw struct {
	Contract *MATICCaller // Generic read-only contract binding to access the raw methods on
}

// MATICTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MATICTransactorRaw struct {
	Contract *MATICTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMATIC creates a new instance of MATIC, bound to a specific deployed contract.
func NewMATIC(address common.Address, backend bind.ContractBackend) (*MATIC, error) {
	contract, err := bindMATIC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MATIC{MATICCaller: MATICCaller{contract: contract}, MATICTransactor: MATICTransactor{contract: contract}, MATICFilterer: MATICFilterer{contract: contract}}, nil
}

// NewMATICCaller creates a new read-only instance of MATIC, bound to a specific deployed contract.
func NewMATICCaller(address common.Address, caller bind.ContractCaller) (*MATICCaller, error) {
	contract, err := bindMATIC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MATICCaller{contract: contract}, nil
}

// NewMATICTransactor creates a new write-only instance of MATIC, bound to a specific deployed contract.
func NewMATICTransactor(address common.Address, transactor bind.ContractTransactor) (*MATICTransactor, error) {
	contract, err := bindMATIC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MATICTransactor{contract: contract}, nil
}

// NewMATICFilterer creates a new log filterer instance of MATIC, bound to a specific deployed contract.
func NewMATICFilterer(address common.Address, filterer bind.ContractFilterer) (*MATICFilterer, error) {
	contract, err := bindMATIC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MATICFilterer{contract: contract}, nil
}

// bindMATIC binds a generic wrapper to an already deployed contract.
func bindMATIC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MATICABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MATIC *MATICRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MATIC.Contract.MATICCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MATIC *MATICRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MATIC.Contract.MATICTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MATIC *MATICRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MATIC.Contract.MATICTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MATIC *MATICCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MATIC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MATIC *MATICTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MATIC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MATIC *MATICTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MATIC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MATIC *MATICCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MATIC *MATICSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MATIC.Contract.Allowance(&_MATIC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MATIC *MATICCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MATIC.Contract.Allowance(&_MATIC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MATIC *MATICCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MATIC *MATICSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MATIC.Contract.BalanceOf(&_MATIC.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MATIC *MATICCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MATIC.Contract.BalanceOf(&_MATIC.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MATIC *MATICCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MATIC *MATICSession) Decimals() (uint8, error) {
	return _MATIC.Contract.Decimals(&_MATIC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MATIC *MATICCallerSession) Decimals() (uint8, error) {
	return _MATIC.Contract.Decimals(&_MATIC.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MATIC *MATICCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MATIC *MATICSession) IsPauser(account common.Address) (bool, error) {
	return _MATIC.Contract.IsPauser(&_MATIC.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MATIC *MATICCallerSession) IsPauser(account common.Address) (bool, error) {
	return _MATIC.Contract.IsPauser(&_MATIC.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MATIC *MATICCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MATIC *MATICSession) Name() (string, error) {
	return _MATIC.Contract.Name(&_MATIC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MATIC *MATICCallerSession) Name() (string, error) {
	return _MATIC.Contract.Name(&_MATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MATIC *MATICCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MATIC *MATICSession) Paused() (bool, error) {
	return _MATIC.Contract.Paused(&_MATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MATIC *MATICCallerSession) Paused() (bool, error) {
	return _MATIC.Contract.Paused(&_MATIC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MATIC *MATICCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MATIC *MATICSession) Symbol() (string, error) {
	return _MATIC.Contract.Symbol(&_MATIC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MATIC *MATICCallerSession) Symbol() (string, error) {
	return _MATIC.Contract.Symbol(&_MATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MATIC *MATICCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MATIC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MATIC *MATICSession) TotalSupply() (*big.Int, error) {
	return _MATIC.Contract.TotalSupply(&_MATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MATIC *MATICCallerSession) TotalSupply() (*big.Int, error) {
	return _MATIC.Contract.TotalSupply(&_MATIC.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MATIC *MATICTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MATIC *MATICSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MATIC.Contract.AddPauser(&_MATIC.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MATIC *MATICTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MATIC.Contract.AddPauser(&_MATIC.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MATIC *MATICTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MATIC *MATICSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.Approve(&_MATIC.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MATIC *MATICTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.Approve(&_MATIC.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_MATIC *MATICTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_MATIC *MATICSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.DecreaseAllowance(&_MATIC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_MATIC *MATICTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.DecreaseAllowance(&_MATIC.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_MATIC *MATICTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_MATIC *MATICSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.IncreaseAllowance(&_MATIC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_MATIC *MATICTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.IncreaseAllowance(&_MATIC.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MATIC *MATICTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MATIC *MATICSession) Pause() (*types.Transaction, error) {
	return _MATIC.Contract.Pause(&_MATIC.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MATIC *MATICTransactorSession) Pause() (*types.Transaction, error) {
	return _MATIC.Contract.Pause(&_MATIC.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MATIC *MATICTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MATIC *MATICSession) RenouncePauser() (*types.Transaction, error) {
	return _MATIC.Contract.RenouncePauser(&_MATIC.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MATIC *MATICTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _MATIC.Contract.RenouncePauser(&_MATIC.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MATIC *MATICTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MATIC *MATICSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.Transfer(&_MATIC.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MATIC *MATICTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.Transfer(&_MATIC.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MATIC *MATICTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MATIC *MATICSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.TransferFrom(&_MATIC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MATIC *MATICTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MATIC.Contract.TransferFrom(&_MATIC.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MATIC *MATICTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MATIC.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MATIC *MATICSession) Unpause() (*types.Transaction, error) {
	return _MATIC.Contract.Unpause(&_MATIC.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MATIC *MATICTransactorSession) Unpause() (*types.Transaction, error) {
	return _MATIC.Contract.Unpause(&_MATIC.TransactOpts)
}

// MATICApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MATIC contract.
type MATICApprovalIterator struct {
	Event *MATICApproval // Event containing the contract specifics and raw log

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
func (it *MATICApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICApproval)
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
		it.Event = new(MATICApproval)
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
func (it *MATICApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICApproval represents a Approval event raised by the MATIC contract.
type MATICApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MATIC *MATICFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MATICApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MATICApprovalIterator{contract: _MATIC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MATIC *MATICFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MATICApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICApproval)
				if err := _MATIC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MATIC *MATICFilterer) ParseApproval(log types.Log) (*MATICApproval, error) {
	event := new(MATICApproval)
	if err := _MATIC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MATICPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MATIC contract.
type MATICPausedIterator struct {
	Event *MATICPaused // Event containing the contract specifics and raw log

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
func (it *MATICPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICPaused)
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
		it.Event = new(MATICPaused)
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
func (it *MATICPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICPaused represents a Paused event raised by the MATIC contract.
type MATICPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MATIC *MATICFilterer) FilterPaused(opts *bind.FilterOpts) (*MATICPausedIterator, error) {

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MATICPausedIterator{contract: _MATIC.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MATIC *MATICFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MATICPaused) (event.Subscription, error) {

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICPaused)
				if err := _MATIC.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MATIC *MATICFilterer) ParsePaused(log types.Log) (*MATICPaused, error) {
	event := new(MATICPaused)
	if err := _MATIC.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MATICPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the MATIC contract.
type MATICPauserAddedIterator struct {
	Event *MATICPauserAdded // Event containing the contract specifics and raw log

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
func (it *MATICPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICPauserAdded)
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
		it.Event = new(MATICPauserAdded)
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
func (it *MATICPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICPauserAdded represents a PauserAdded event raised by the MATIC contract.
type MATICPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_MATIC *MATICFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*MATICPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MATICPauserAddedIterator{contract: _MATIC.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_MATIC *MATICFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *MATICPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICPauserAdded)
				if err := _MATIC.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_MATIC *MATICFilterer) ParsePauserAdded(log types.Log) (*MATICPauserAdded, error) {
	event := new(MATICPauserAdded)
	if err := _MATIC.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MATICPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the MATIC contract.
type MATICPauserRemovedIterator struct {
	Event *MATICPauserRemoved // Event containing the contract specifics and raw log

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
func (it *MATICPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICPauserRemoved)
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
		it.Event = new(MATICPauserRemoved)
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
func (it *MATICPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICPauserRemoved represents a PauserRemoved event raised by the MATIC contract.
type MATICPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_MATIC *MATICFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*MATICPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MATICPauserRemovedIterator{contract: _MATIC.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_MATIC *MATICFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *MATICPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICPauserRemoved)
				if err := _MATIC.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_MATIC *MATICFilterer) ParsePauserRemoved(log types.Log) (*MATICPauserRemoved, error) {
	event := new(MATICPauserRemoved)
	if err := _MATIC.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MATICTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MATIC contract.
type MATICTransferIterator struct {
	Event *MATICTransfer // Event containing the contract specifics and raw log

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
func (it *MATICTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICTransfer)
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
		it.Event = new(MATICTransfer)
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
func (it *MATICTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICTransfer represents a Transfer event raised by the MATIC contract.
type MATICTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MATIC *MATICFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MATICTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MATICTransferIterator{contract: _MATIC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MATIC *MATICFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MATICTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICTransfer)
				if err := _MATIC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MATIC *MATICFilterer) ParseTransfer(log types.Log) (*MATICTransfer, error) {
	event := new(MATICTransfer)
	if err := _MATIC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MATICUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MATIC contract.
type MATICUnpausedIterator struct {
	Event *MATICUnpaused // Event containing the contract specifics and raw log

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
func (it *MATICUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MATICUnpaused)
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
		it.Event = new(MATICUnpaused)
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
func (it *MATICUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MATICUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MATICUnpaused represents a Unpaused event raised by the MATIC contract.
type MATICUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MATIC *MATICFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MATICUnpausedIterator, error) {

	logs, sub, err := _MATIC.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MATICUnpausedIterator{contract: _MATIC.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MATIC *MATICFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MATICUnpaused) (event.Subscription, error) {

	logs, sub, err := _MATIC.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MATICUnpaused)
				if err := _MATIC.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MATIC *MATICFilterer) ParseUnpaused(log types.Log) (*MATICUnpaused, error) {
	event := new(MATICUnpaused)
	if err := _MATIC.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
