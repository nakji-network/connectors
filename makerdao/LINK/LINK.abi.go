// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LINK

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

// LINKABI is the input ABI used to generate the binding from.
const LINKABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// LINK is an auto generated Go binding around an Ethereum contract.
type LINK struct {
	LINKCaller     // Read-only binding to the contract
	LINKTransactor // Write-only binding to the contract
	LINKFilterer   // Log filterer for contract events
}

// LINKCaller is an auto generated read-only Go binding around an Ethereum contract.
type LINKCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LINKTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LINKTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LINKFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LINKFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LINKSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LINKSession struct {
	Contract     *LINK             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LINKCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LINKCallerSession struct {
	Contract *LINKCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LINKTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LINKTransactorSession struct {
	Contract     *LINKTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LINKRaw is an auto generated low-level Go binding around an Ethereum contract.
type LINKRaw struct {
	Contract *LINK // Generic contract binding to access the raw methods on
}

// LINKCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LINKCallerRaw struct {
	Contract *LINKCaller // Generic read-only contract binding to access the raw methods on
}

// LINKTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LINKTransactorRaw struct {
	Contract *LINKTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLINK creates a new instance of LINK, bound to a specific deployed contract.
func NewLINK(address common.Address, backend bind.ContractBackend) (*LINK, error) {
	contract, err := bindLINK(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LINK{LINKCaller: LINKCaller{contract: contract}, LINKTransactor: LINKTransactor{contract: contract}, LINKFilterer: LINKFilterer{contract: contract}}, nil
}

// NewLINKCaller creates a new read-only instance of LINK, bound to a specific deployed contract.
func NewLINKCaller(address common.Address, caller bind.ContractCaller) (*LINKCaller, error) {
	contract, err := bindLINK(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LINKCaller{contract: contract}, nil
}

// NewLINKTransactor creates a new write-only instance of LINK, bound to a specific deployed contract.
func NewLINKTransactor(address common.Address, transactor bind.ContractTransactor) (*LINKTransactor, error) {
	contract, err := bindLINK(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LINKTransactor{contract: contract}, nil
}

// NewLINKFilterer creates a new log filterer instance of LINK, bound to a specific deployed contract.
func NewLINKFilterer(address common.Address, filterer bind.ContractFilterer) (*LINKFilterer, error) {
	contract, err := bindLINK(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LINKFilterer{contract: contract}, nil
}

// bindLINK binds a generic wrapper to an already deployed contract.
func bindLINK(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LINKABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LINK *LINKRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LINK.Contract.LINKCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LINK *LINKRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LINK.Contract.LINKTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LINK *LINKRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LINK.Contract.LINKTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LINK *LINKCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LINK.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LINK *LINKTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LINK.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LINK *LINKTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LINK.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_LINK *LINKCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_LINK *LINKSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LINK.Contract.Allowance(&_LINK.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_LINK *LINKCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LINK.Contract.Allowance(&_LINK.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_LINK *LINKCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_LINK *LINKSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LINK.Contract.BalanceOf(&_LINK.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_LINK *LINKCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LINK.Contract.BalanceOf(&_LINK.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LINK *LINKCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LINK *LINKSession) Decimals() (uint8, error) {
	return _LINK.Contract.Decimals(&_LINK.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LINK *LINKCallerSession) Decimals() (uint8, error) {
	return _LINK.Contract.Decimals(&_LINK.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LINK *LINKCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LINK *LINKSession) Name() (string, error) {
	return _LINK.Contract.Name(&_LINK.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LINK *LINKCallerSession) Name() (string, error) {
	return _LINK.Contract.Name(&_LINK.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LINK *LINKCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LINK *LINKSession) Symbol() (string, error) {
	return _LINK.Contract.Symbol(&_LINK.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LINK *LINKCallerSession) Symbol() (string, error) {
	return _LINK.Contract.Symbol(&_LINK.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LINK *LINKCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LINK.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LINK *LINKSession) TotalSupply() (*big.Int, error) {
	return _LINK.Contract.TotalSupply(&_LINK.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LINK *LINKCallerSession) TotalSupply() (*big.Int, error) {
	return _LINK.Contract.TotalSupply(&_LINK.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LINK *LINKTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LINK *LINKSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.Approve(&_LINK.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LINK *LINKTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.Approve(&_LINK.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_LINK *LINKTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_LINK *LINKSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.DecreaseApproval(&_LINK.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_LINK *LINKTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.DecreaseApproval(&_LINK.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_LINK *LINKTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_LINK *LINKSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.IncreaseApproval(&_LINK.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_LINK *LINKTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.IncreaseApproval(&_LINK.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_LINK *LINKTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_LINK *LINKSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.Transfer(&_LINK.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_LINK *LINKTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.Transfer(&_LINK.TransactOpts, _to, _value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool success)
func (_LINK *LINKTransactor) TransferAndCall(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "transferAndCall", _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool success)
func (_LINK *LINKSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _LINK.Contract.TransferAndCall(&_LINK.TransactOpts, _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool success)
func (_LINK *LINKTransactorSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _LINK.Contract.TransferAndCall(&_LINK.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LINK *LINKTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LINK *LINKSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.TransferFrom(&_LINK.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LINK *LINKTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LINK.Contract.TransferFrom(&_LINK.TransactOpts, _from, _to, _value)
}

// LINKApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LINK contract.
type LINKApprovalIterator struct {
	Event *LINKApproval // Event containing the contract specifics and raw log

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
func (it *LINKApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LINKApproval)
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
		it.Event = new(LINKApproval)
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
func (it *LINKApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LINKApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LINKApproval represents a Approval event raised by the LINK contract.
type LINKApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LINK *LINKFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LINKApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LINK.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LINKApprovalIterator{contract: _LINK.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LINK *LINKFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LINKApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LINK.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LINKApproval)
				if err := _LINK.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LINK *LINKFilterer) ParseApproval(log types.Log) (*LINKApproval, error) {
	event := new(LINKApproval)
	if err := _LINK.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LINKTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LINK contract.
type LINKTransferIterator struct {
	Event *LINKTransfer // Event containing the contract specifics and raw log

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
func (it *LINKTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LINKTransfer)
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
		it.Event = new(LINKTransfer)
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
func (it *LINKTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LINKTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LINKTransfer represents a Transfer event raised by the LINK contract.
type LINKTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_LINK *LINKFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LINKTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LINK.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LINKTransferIterator{contract: _LINK.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_LINK *LINKFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LINKTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LINK.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LINKTransfer)
				if err := _LINK.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_LINK *LINKFilterer) ParseTransfer(log types.Log) (*LINKTransfer, error) {
	event := new(LINKTransfer)
	if err := _LINK.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
