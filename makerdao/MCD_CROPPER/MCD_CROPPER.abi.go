// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_CROPPER

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

// MCDCROPPERABI is the input ABI used to generate the binding from.
const MCDCROPPERABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetImplementation\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDCROPPER is an auto generated Go binding around an Ethereum contract.
type MCDCROPPER struct {
	MCDCROPPERCaller     // Read-only binding to the contract
	MCDCROPPERTransactor // Write-only binding to the contract
	MCDCROPPERFilterer   // Log filterer for contract events
}

// MCDCROPPERCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDCROPPERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDCROPPERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDCROPPERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDCROPPERSession struct {
	Contract     *MCDCROPPER       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDCROPPERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDCROPPERCallerSession struct {
	Contract *MCDCROPPERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MCDCROPPERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDCROPPERTransactorSession struct {
	Contract     *MCDCROPPERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MCDCROPPERRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDCROPPERRaw struct {
	Contract *MCDCROPPER // Generic contract binding to access the raw methods on
}

// MCDCROPPERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDCROPPERCallerRaw struct {
	Contract *MCDCROPPERCaller // Generic read-only contract binding to access the raw methods on
}

// MCDCROPPERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDCROPPERTransactorRaw struct {
	Contract *MCDCROPPERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDCROPPER creates a new instance of MCDCROPPER, bound to a specific deployed contract.
func NewMCDCROPPER(address common.Address, backend bind.ContractBackend) (*MCDCROPPER, error) {
	contract, err := bindMCDCROPPER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPER{MCDCROPPERCaller: MCDCROPPERCaller{contract: contract}, MCDCROPPERTransactor: MCDCROPPERTransactor{contract: contract}, MCDCROPPERFilterer: MCDCROPPERFilterer{contract: contract}}, nil
}

// NewMCDCROPPERCaller creates a new read-only instance of MCDCROPPER, bound to a specific deployed contract.
func NewMCDCROPPERCaller(address common.Address, caller bind.ContractCaller) (*MCDCROPPERCaller, error) {
	contract, err := bindMCDCROPPER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERCaller{contract: contract}, nil
}

// NewMCDCROPPERTransactor creates a new write-only instance of MCDCROPPER, bound to a specific deployed contract.
func NewMCDCROPPERTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDCROPPERTransactor, error) {
	contract, err := bindMCDCROPPER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERTransactor{contract: contract}, nil
}

// NewMCDCROPPERFilterer creates a new log filterer instance of MCDCROPPER, bound to a specific deployed contract.
func NewMCDCROPPERFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDCROPPERFilterer, error) {
	contract, err := bindMCDCROPPER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERFilterer{contract: contract}, nil
}

// bindMCDCROPPER binds a generic wrapper to an already deployed contract.
func bindMCDCROPPER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDCROPPERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCROPPER *MCDCROPPERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCROPPER.Contract.MCDCROPPERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCROPPER *MCDCROPPERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.MCDCROPPERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCROPPER *MCDCROPPERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.MCDCROPPERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCROPPER *MCDCROPPERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCROPPER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCROPPER *MCDCROPPERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCROPPER *MCDCROPPERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MCDCROPPER *MCDCROPPERCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDCROPPER.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MCDCROPPER *MCDCROPPERSession) Implementation() (common.Address, error) {
	return _MCDCROPPER.Contract.Implementation(&_MCDCROPPER.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MCDCROPPER *MCDCROPPERCallerSession) Implementation() (common.Address, error) {
	return _MCDCROPPER.Contract.Implementation(&_MCDCROPPER.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCROPPER *MCDCROPPERCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDCROPPER.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCROPPER *MCDCROPPERSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDCROPPER.Contract.Wards(&_MCDCROPPER.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCROPPER *MCDCROPPERCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDCROPPER.Contract.Wards(&_MCDCROPPER.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCROPPER *MCDCROPPERTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCROPPER *MCDCROPPERSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Deny(&_MCDCROPPER.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCROPPER *MCDCROPPERTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Deny(&_MCDCROPPER.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCROPPER *MCDCROPPERTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCROPPER *MCDCROPPERSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Rely(&_MCDCROPPER.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCROPPER *MCDCROPPERTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Rely(&_MCDCROPPER.TransactOpts, usr)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_MCDCROPPER *MCDCROPPERTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.contract.Transact(opts, "setImplementation", implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_MCDCROPPER *MCDCROPPERSession) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.SetImplementation(&_MCDCROPPER.TransactOpts, implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_MCDCROPPER *MCDCROPPERTransactorSession) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.SetImplementation(&_MCDCROPPER.TransactOpts, implementation_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MCDCROPPER *MCDCROPPERTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MCDCROPPER.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MCDCROPPER *MCDCROPPERSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Fallback(&_MCDCROPPER.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MCDCROPPER *MCDCROPPERTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MCDCROPPER.Contract.Fallback(&_MCDCROPPER.TransactOpts, calldata)
}

// MCDCROPPERDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDCROPPER contract.
type MCDCROPPERDenyIterator struct {
	Event *MCDCROPPERDeny // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERDeny)
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
		it.Event = new(MCDCROPPERDeny)
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
func (it *MCDCROPPERDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERDeny represents a Deny event raised by the MCDCROPPER contract.
type MCDCROPPERDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDCROPPER *MCDCROPPERFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDCROPPERDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDCROPPER.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERDenyIterator{contract: _MCDCROPPER.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDCROPPER *MCDCROPPERFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDCROPPERDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDCROPPER.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERDeny)
				if err := _MCDCROPPER.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDCROPPER *MCDCROPPERFilterer) ParseDeny(log types.Log) (*MCDCROPPERDeny, error) {
	event := new(MCDCROPPERDeny)
	if err := _MCDCROPPER.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDCROPPERRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDCROPPER contract.
type MCDCROPPERRelyIterator struct {
	Event *MCDCROPPERRely // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERRely)
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
		it.Event = new(MCDCROPPERRely)
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
func (it *MCDCROPPERRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERRely represents a Rely event raised by the MCDCROPPER contract.
type MCDCROPPERRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDCROPPER *MCDCROPPERFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDCROPPERRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDCROPPER.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERRelyIterator{contract: _MCDCROPPER.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDCROPPER *MCDCROPPERFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDCROPPERRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDCROPPER.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERRely)
				if err := _MCDCROPPER.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDCROPPER *MCDCROPPERFilterer) ParseRely(log types.Log) (*MCDCROPPERRely, error) {
	event := new(MCDCROPPERRely)
	if err := _MCDCROPPER.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDCROPPERSetImplementationIterator is returned from FilterSetImplementation and is used to iterate over the raw logs and unpacked data for SetImplementation events raised by the MCDCROPPER contract.
type MCDCROPPERSetImplementationIterator struct {
	Event *MCDCROPPERSetImplementation // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERSetImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERSetImplementation)
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
		it.Event = new(MCDCROPPERSetImplementation)
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
func (it *MCDCROPPERSetImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERSetImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERSetImplementation represents a SetImplementation event raised by the MCDCROPPER contract.
type MCDCROPPERSetImplementation struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetImplementation is a free log retrieval operation binding the contract event 0xddebe6de740fe0dd01cc33ffa314d11c6ac6acbbe50b80513c4c360ae7aa4f04.
//
// Solidity: event SetImplementation(address indexed arg0)
func (_MCDCROPPER *MCDCROPPERFilterer) FilterSetImplementation(opts *bind.FilterOpts, arg0 []common.Address) (*MCDCROPPERSetImplementationIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MCDCROPPER.contract.FilterLogs(opts, "SetImplementation", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERSetImplementationIterator{contract: _MCDCROPPER.contract, event: "SetImplementation", logs: logs, sub: sub}, nil
}

// WatchSetImplementation is a free log subscription operation binding the contract event 0xddebe6de740fe0dd01cc33ffa314d11c6ac6acbbe50b80513c4c360ae7aa4f04.
//
// Solidity: event SetImplementation(address indexed arg0)
func (_MCDCROPPER *MCDCROPPERFilterer) WatchSetImplementation(opts *bind.WatchOpts, sink chan<- *MCDCROPPERSetImplementation, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MCDCROPPER.contract.WatchLogs(opts, "SetImplementation", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERSetImplementation)
				if err := _MCDCROPPER.contract.UnpackLog(event, "SetImplementation", log); err != nil {
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

// ParseSetImplementation is a log parse operation binding the contract event 0xddebe6de740fe0dd01cc33ffa314d11c6ac6acbbe50b80513c4c360ae7aa4f04.
//
// Solidity: event SetImplementation(address indexed arg0)
func (_MCDCROPPER *MCDCROPPERFilterer) ParseSetImplementation(log types.Log) (*MCDCROPPERSetImplementation, error) {
	event := new(MCDCROPPERSetImplementation)
	if err := _MCDCROPPER.contract.UnpackLog(event, "SetImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
