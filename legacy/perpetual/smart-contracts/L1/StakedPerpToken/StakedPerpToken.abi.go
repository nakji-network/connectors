// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StakedPerpToken

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

// StakedPerpTokenABI is the input ABI used to generate the binding from.
const StakedPerpTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// StakedPerpToken is an auto generated Go binding around an Ethereum contract.
type StakedPerpToken struct {
	StakedPerpTokenCaller     // Read-only binding to the contract
	StakedPerpTokenTransactor // Write-only binding to the contract
	StakedPerpTokenFilterer   // Log filterer for contract events
}

// StakedPerpTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedPerpTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedPerpTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedPerpTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedPerpTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedPerpTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedPerpTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedPerpTokenSession struct {
	Contract     *StakedPerpToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedPerpTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedPerpTokenCallerSession struct {
	Contract *StakedPerpTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakedPerpTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedPerpTokenTransactorSession struct {
	Contract     *StakedPerpTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakedPerpTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedPerpTokenRaw struct {
	Contract *StakedPerpToken // Generic contract binding to access the raw methods on
}

// StakedPerpTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedPerpTokenCallerRaw struct {
	Contract *StakedPerpTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StakedPerpTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedPerpTokenTransactorRaw struct {
	Contract *StakedPerpTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedPerpToken creates a new instance of StakedPerpToken, bound to a specific deployed contract.
func NewStakedPerpToken(address common.Address, backend bind.ContractBackend) (*StakedPerpToken, error) {
	contract, err := bindStakedPerpToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedPerpToken{StakedPerpTokenCaller: StakedPerpTokenCaller{contract: contract}, StakedPerpTokenTransactor: StakedPerpTokenTransactor{contract: contract}, StakedPerpTokenFilterer: StakedPerpTokenFilterer{contract: contract}}, nil
}

// NewStakedPerpTokenCaller creates a new read-only instance of StakedPerpToken, bound to a specific deployed contract.
func NewStakedPerpTokenCaller(address common.Address, caller bind.ContractCaller) (*StakedPerpTokenCaller, error) {
	contract, err := bindStakedPerpToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedPerpTokenCaller{contract: contract}, nil
}

// NewStakedPerpTokenTransactor creates a new write-only instance of StakedPerpToken, bound to a specific deployed contract.
func NewStakedPerpTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedPerpTokenTransactor, error) {
	contract, err := bindStakedPerpToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedPerpTokenTransactor{contract: contract}, nil
}

// NewStakedPerpTokenFilterer creates a new log filterer instance of StakedPerpToken, bound to a specific deployed contract.
func NewStakedPerpTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedPerpTokenFilterer, error) {
	contract, err := bindStakedPerpToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedPerpTokenFilterer{contract: contract}, nil
}

// bindStakedPerpToken binds a generic wrapper to an already deployed contract.
func bindStakedPerpToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakedPerpTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedPerpToken *StakedPerpTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedPerpToken.Contract.StakedPerpTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedPerpToken *StakedPerpTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.StakedPerpTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedPerpToken *StakedPerpTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.StakedPerpTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedPerpToken *StakedPerpTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedPerpToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedPerpToken *StakedPerpTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedPerpToken *StakedPerpTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_StakedPerpToken *StakedPerpTokenTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedPerpToken.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_StakedPerpToken *StakedPerpTokenSession) Admin() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Admin(&_StakedPerpToken.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_StakedPerpToken *StakedPerpTokenTransactorSession) Admin() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Admin(&_StakedPerpToken.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_StakedPerpToken *StakedPerpTokenTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_StakedPerpToken *StakedPerpTokenSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.ChangeAdmin(&_StakedPerpToken.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_StakedPerpToken *StakedPerpTokenTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.ChangeAdmin(&_StakedPerpToken.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_StakedPerpToken *StakedPerpTokenTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedPerpToken.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_StakedPerpToken *StakedPerpTokenSession) Implementation() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Implementation(&_StakedPerpToken.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_StakedPerpToken *StakedPerpTokenTransactorSession) Implementation() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Implementation(&_StakedPerpToken.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakedPerpToken *StakedPerpTokenTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakedPerpToken *StakedPerpTokenSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.UpgradeTo(&_StakedPerpToken.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakedPerpToken *StakedPerpTokenTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.UpgradeTo(&_StakedPerpToken.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakedPerpToken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakedPerpToken *StakedPerpTokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.UpgradeToAndCall(&_StakedPerpToken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.UpgradeToAndCall(&_StakedPerpToken.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakedPerpToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedPerpToken *StakedPerpTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Fallback(&_StakedPerpToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Fallback(&_StakedPerpToken.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedPerpToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedPerpToken *StakedPerpTokenSession) Receive() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Receive(&_StakedPerpToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedPerpToken *StakedPerpTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _StakedPerpToken.Contract.Receive(&_StakedPerpToken.TransactOpts)
}

// StakedPerpTokenAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StakedPerpToken contract.
type StakedPerpTokenAdminChangedIterator struct {
	Event *StakedPerpTokenAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakedPerpTokenAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedPerpTokenAdminChanged)
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
		it.Event = new(StakedPerpTokenAdminChanged)
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
func (it *StakedPerpTokenAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedPerpTokenAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedPerpTokenAdminChanged represents a AdminChanged event raised by the StakedPerpToken contract.
type StakedPerpTokenAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakedPerpToken *StakedPerpTokenFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakedPerpTokenAdminChangedIterator, error) {

	logs, sub, err := _StakedPerpToken.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakedPerpTokenAdminChangedIterator{contract: _StakedPerpToken.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakedPerpToken *StakedPerpTokenFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakedPerpTokenAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StakedPerpToken.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedPerpTokenAdminChanged)
				if err := _StakedPerpToken.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakedPerpToken *StakedPerpTokenFilterer) ParseAdminChanged(log types.Log) (*StakedPerpTokenAdminChanged, error) {
	event := new(StakedPerpTokenAdminChanged)
	if err := _StakedPerpToken.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedPerpTokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakedPerpToken contract.
type StakedPerpTokenUpgradedIterator struct {
	Event *StakedPerpTokenUpgraded // Event containing the contract specifics and raw log

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
func (it *StakedPerpTokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedPerpTokenUpgraded)
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
		it.Event = new(StakedPerpTokenUpgraded)
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
func (it *StakedPerpTokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedPerpTokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedPerpTokenUpgraded represents a Upgraded event raised by the StakedPerpToken contract.
type StakedPerpTokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakedPerpToken *StakedPerpTokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakedPerpTokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakedPerpToken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakedPerpTokenUpgradedIterator{contract: _StakedPerpToken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakedPerpToken *StakedPerpTokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakedPerpTokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakedPerpToken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedPerpTokenUpgraded)
				if err := _StakedPerpToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakedPerpToken *StakedPerpTokenFilterer) ParseUpgraded(log types.Log) (*StakedPerpTokenUpgraded, error) {
	event := new(StakedPerpTokenUpgraded)
	if err := _StakedPerpToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
