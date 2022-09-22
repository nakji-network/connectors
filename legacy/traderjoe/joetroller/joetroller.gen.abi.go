// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joetroller

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

// JoetrollerMetaData contains all meta data concerning the Joetroller contract.
var JoetrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"NewPendingImplementation\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"_setPendingImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JoetrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use JoetrollerMetaData.ABI instead.
var JoetrollerABI = JoetrollerMetaData.ABI

// Joetroller is an auto generated Go binding around an Ethereum contract.
type Joetroller struct {
	JoetrollerCaller     // Read-only binding to the contract
	JoetrollerTransactor // Write-only binding to the contract
	JoetrollerFilterer   // Log filterer for contract events
}

// JoetrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JoetrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoetrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JoetrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoetrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JoetrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoetrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JoetrollerSession struct {
	Contract     *Joetroller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoetrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JoetrollerCallerSession struct {
	Contract *JoetrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JoetrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JoetrollerTransactorSession struct {
	Contract     *JoetrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JoetrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JoetrollerRaw struct {
	Contract *Joetroller // Generic contract binding to access the raw methods on
}

// JoetrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JoetrollerCallerRaw struct {
	Contract *JoetrollerCaller // Generic read-only contract binding to access the raw methods on
}

// JoetrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JoetrollerTransactorRaw struct {
	Contract *JoetrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJoetroller creates a new instance of Joetroller, bound to a specific deployed contract.
func NewJoetroller(address common.Address, backend bind.ContractBackend) (*Joetroller, error) {
	contract, err := bindJoetroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Joetroller{JoetrollerCaller: JoetrollerCaller{contract: contract}, JoetrollerTransactor: JoetrollerTransactor{contract: contract}, JoetrollerFilterer: JoetrollerFilterer{contract: contract}}, nil
}

// NewJoetrollerCaller creates a new read-only instance of Joetroller, bound to a specific deployed contract.
func NewJoetrollerCaller(address common.Address, caller bind.ContractCaller) (*JoetrollerCaller, error) {
	contract, err := bindJoetroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JoetrollerCaller{contract: contract}, nil
}

// NewJoetrollerTransactor creates a new write-only instance of Joetroller, bound to a specific deployed contract.
func NewJoetrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*JoetrollerTransactor, error) {
	contract, err := bindJoetroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JoetrollerTransactor{contract: contract}, nil
}

// NewJoetrollerFilterer creates a new log filterer instance of Joetroller, bound to a specific deployed contract.
func NewJoetrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*JoetrollerFilterer, error) {
	contract, err := bindJoetroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JoetrollerFilterer{contract: contract}, nil
}

// bindJoetroller binds a generic wrapper to an already deployed contract.
func bindJoetroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JoetrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joetroller *JoetrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joetroller.Contract.JoetrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joetroller *JoetrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joetroller.Contract.JoetrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joetroller *JoetrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joetroller.Contract.JoetrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joetroller *JoetrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joetroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joetroller *JoetrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joetroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joetroller *JoetrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joetroller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Joetroller *JoetrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joetroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Joetroller *JoetrollerSession) Admin() (common.Address, error) {
	return _Joetroller.Contract.Admin(&_Joetroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Joetroller *JoetrollerCallerSession) Admin() (common.Address, error) {
	return _Joetroller.Contract.Admin(&_Joetroller.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Joetroller *JoetrollerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joetroller.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Joetroller *JoetrollerSession) Implementation() (common.Address, error) {
	return _Joetroller.Contract.Implementation(&_Joetroller.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Joetroller *JoetrollerCallerSession) Implementation() (common.Address, error) {
	return _Joetroller.Contract.Implementation(&_Joetroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Joetroller *JoetrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joetroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Joetroller *JoetrollerSession) PendingAdmin() (common.Address, error) {
	return _Joetroller.Contract.PendingAdmin(&_Joetroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Joetroller *JoetrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Joetroller.Contract.PendingAdmin(&_Joetroller.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_Joetroller *JoetrollerCaller) PendingImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joetroller.contract.Call(opts, &out, "pendingImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_Joetroller *JoetrollerSession) PendingImplementation() (common.Address, error) {
	return _Joetroller.Contract.PendingImplementation(&_Joetroller.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_Joetroller *JoetrollerCallerSession) PendingImplementation() (common.Address, error) {
	return _Joetroller.Contract.PendingImplementation(&_Joetroller.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Joetroller *JoetrollerTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joetroller.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Joetroller *JoetrollerSession) AcceptAdmin() (*types.Transaction, error) {
	return _Joetroller.Contract.AcceptAdmin(&_Joetroller.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Joetroller *JoetrollerTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Joetroller.Contract.AcceptAdmin(&_Joetroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Joetroller *JoetrollerTransactor) AcceptImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joetroller.contract.Transact(opts, "_acceptImplementation")
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Joetroller *JoetrollerSession) AcceptImplementation() (*types.Transaction, error) {
	return _Joetroller.Contract.AcceptImplementation(&_Joetroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Joetroller *JoetrollerTransactorSession) AcceptImplementation() (*types.Transaction, error) {
	return _Joetroller.Contract.AcceptImplementation(&_Joetroller.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Joetroller *JoetrollerTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Joetroller.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Joetroller *JoetrollerSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Joetroller.Contract.SetPendingAdmin(&_Joetroller.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Joetroller *JoetrollerTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Joetroller.Contract.SetPendingAdmin(&_Joetroller.TransactOpts, newPendingAdmin)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Joetroller *JoetrollerTransactor) SetPendingImplementation(opts *bind.TransactOpts, newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Joetroller.contract.Transact(opts, "_setPendingImplementation", newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Joetroller *JoetrollerSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Joetroller.Contract.SetPendingImplementation(&_Joetroller.TransactOpts, newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Joetroller *JoetrollerTransactorSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Joetroller.Contract.SetPendingImplementation(&_Joetroller.TransactOpts, newPendingImplementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Joetroller *JoetrollerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Joetroller.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Joetroller *JoetrollerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Joetroller.Contract.Fallback(&_Joetroller.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Joetroller *JoetrollerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Joetroller.Contract.Fallback(&_Joetroller.TransactOpts, calldata)
}

// JoetrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Joetroller contract.
type JoetrollerFailureIterator struct {
	Event *JoetrollerFailure // Event containing the contract specifics and raw log

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
func (it *JoetrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoetrollerFailure)
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
		it.Event = new(JoetrollerFailure)
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
func (it *JoetrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoetrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoetrollerFailure represents a Failure event raised by the Joetroller contract.
type JoetrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Joetroller *JoetrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*JoetrollerFailureIterator, error) {

	logs, sub, err := _Joetroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &JoetrollerFailureIterator{contract: _Joetroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Joetroller *JoetrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *JoetrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Joetroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoetrollerFailure)
				if err := _Joetroller.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Joetroller *JoetrollerFilterer) ParseFailure(log types.Log) (*JoetrollerFailure, error) {
	event := new(JoetrollerFailure)
	if err := _Joetroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoetrollerNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Joetroller contract.
type JoetrollerNewAdminIterator struct {
	Event *JoetrollerNewAdmin // Event containing the contract specifics and raw log

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
func (it *JoetrollerNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoetrollerNewAdmin)
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
		it.Event = new(JoetrollerNewAdmin)
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
func (it *JoetrollerNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoetrollerNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoetrollerNewAdmin represents a NewAdmin event raised by the Joetroller contract.
type JoetrollerNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Joetroller *JoetrollerFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*JoetrollerNewAdminIterator, error) {

	logs, sub, err := _Joetroller.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &JoetrollerNewAdminIterator{contract: _Joetroller.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Joetroller *JoetrollerFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *JoetrollerNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Joetroller.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoetrollerNewAdmin)
				if err := _Joetroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Joetroller *JoetrollerFilterer) ParseNewAdmin(log types.Log) (*JoetrollerNewAdmin, error) {
	event := new(JoetrollerNewAdmin)
	if err := _Joetroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoetrollerNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Joetroller contract.
type JoetrollerNewImplementationIterator struct {
	Event *JoetrollerNewImplementation // Event containing the contract specifics and raw log

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
func (it *JoetrollerNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoetrollerNewImplementation)
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
		it.Event = new(JoetrollerNewImplementation)
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
func (it *JoetrollerNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoetrollerNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoetrollerNewImplementation represents a NewImplementation event raised by the Joetroller contract.
type JoetrollerNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Joetroller *JoetrollerFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*JoetrollerNewImplementationIterator, error) {

	logs, sub, err := _Joetroller.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &JoetrollerNewImplementationIterator{contract: _Joetroller.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Joetroller *JoetrollerFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *JoetrollerNewImplementation) (event.Subscription, error) {

	logs, sub, err := _Joetroller.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoetrollerNewImplementation)
				if err := _Joetroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Joetroller *JoetrollerFilterer) ParseNewImplementation(log types.Log) (*JoetrollerNewImplementation, error) {
	event := new(JoetrollerNewImplementation)
	if err := _Joetroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoetrollerNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Joetroller contract.
type JoetrollerNewPendingAdminIterator struct {
	Event *JoetrollerNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *JoetrollerNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoetrollerNewPendingAdmin)
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
		it.Event = new(JoetrollerNewPendingAdmin)
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
func (it *JoetrollerNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoetrollerNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoetrollerNewPendingAdmin represents a NewPendingAdmin event raised by the Joetroller contract.
type JoetrollerNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Joetroller *JoetrollerFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*JoetrollerNewPendingAdminIterator, error) {

	logs, sub, err := _Joetroller.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &JoetrollerNewPendingAdminIterator{contract: _Joetroller.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Joetroller *JoetrollerFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *JoetrollerNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Joetroller.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoetrollerNewPendingAdmin)
				if err := _Joetroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Joetroller *JoetrollerFilterer) ParseNewPendingAdmin(log types.Log) (*JoetrollerNewPendingAdmin, error) {
	event := new(JoetrollerNewPendingAdmin)
	if err := _Joetroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoetrollerNewPendingImplementationIterator is returned from FilterNewPendingImplementation and is used to iterate over the raw logs and unpacked data for NewPendingImplementation events raised by the Joetroller contract.
type JoetrollerNewPendingImplementationIterator struct {
	Event *JoetrollerNewPendingImplementation // Event containing the contract specifics and raw log

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
func (it *JoetrollerNewPendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoetrollerNewPendingImplementation)
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
		it.Event = new(JoetrollerNewPendingImplementation)
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
func (it *JoetrollerNewPendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoetrollerNewPendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoetrollerNewPendingImplementation represents a NewPendingImplementation event raised by the Joetroller contract.
type JoetrollerNewPendingImplementation struct {
	OldPendingImplementation common.Address
	NewPendingImplementation common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewPendingImplementation is a free log retrieval operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Joetroller *JoetrollerFilterer) FilterNewPendingImplementation(opts *bind.FilterOpts) (*JoetrollerNewPendingImplementationIterator, error) {

	logs, sub, err := _Joetroller.contract.FilterLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return &JoetrollerNewPendingImplementationIterator{contract: _Joetroller.contract, event: "NewPendingImplementation", logs: logs, sub: sub}, nil
}

// WatchNewPendingImplementation is a free log subscription operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Joetroller *JoetrollerFilterer) WatchNewPendingImplementation(opts *bind.WatchOpts, sink chan<- *JoetrollerNewPendingImplementation) (event.Subscription, error) {

	logs, sub, err := _Joetroller.contract.WatchLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoetrollerNewPendingImplementation)
				if err := _Joetroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
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

// ParseNewPendingImplementation is a log parse operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Joetroller *JoetrollerFilterer) ParseNewPendingImplementation(log types.Log) (*JoetrollerNewPendingImplementation, error) {
	event := new(JoetrollerNewPendingImplementation)
	if err := _Joetroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
