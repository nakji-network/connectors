// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ClientBridge

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

// ClientBridgeABI is the input ABI used to generate the binding from.
const ClientBridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ClientBridge is an auto generated Go binding around an Ethereum contract.
type ClientBridge struct {
	ClientBridgeCaller     // Read-only binding to the contract
	ClientBridgeTransactor // Write-only binding to the contract
	ClientBridgeFilterer   // Log filterer for contract events
}

// ClientBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientBridgeSession struct {
	Contract     *ClientBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientBridgeCallerSession struct {
	Contract *ClientBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClientBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientBridgeTransactorSession struct {
	Contract     *ClientBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClientBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientBridgeRaw struct {
	Contract *ClientBridge // Generic contract binding to access the raw methods on
}

// ClientBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientBridgeCallerRaw struct {
	Contract *ClientBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ClientBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientBridgeTransactorRaw struct {
	Contract *ClientBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClientBridge creates a new instance of ClientBridge, bound to a specific deployed contract.
func NewClientBridge(address common.Address, backend bind.ContractBackend) (*ClientBridge, error) {
	contract, err := bindClientBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClientBridge{ClientBridgeCaller: ClientBridgeCaller{contract: contract}, ClientBridgeTransactor: ClientBridgeTransactor{contract: contract}, ClientBridgeFilterer: ClientBridgeFilterer{contract: contract}}, nil
}

// NewClientBridgeCaller creates a new read-only instance of ClientBridge, bound to a specific deployed contract.
func NewClientBridgeCaller(address common.Address, caller bind.ContractCaller) (*ClientBridgeCaller, error) {
	contract, err := bindClientBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientBridgeCaller{contract: contract}, nil
}

// NewClientBridgeTransactor creates a new write-only instance of ClientBridge, bound to a specific deployed contract.
func NewClientBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientBridgeTransactor, error) {
	contract, err := bindClientBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientBridgeTransactor{contract: contract}, nil
}

// NewClientBridgeFilterer creates a new log filterer instance of ClientBridge, bound to a specific deployed contract.
func NewClientBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientBridgeFilterer, error) {
	contract, err := bindClientBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientBridgeFilterer{contract: contract}, nil
}

// bindClientBridge binds a generic wrapper to an already deployed contract.
func bindClientBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClientBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientBridge *ClientBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientBridge.Contract.ClientBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientBridge *ClientBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientBridge.Contract.ClientBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientBridge *ClientBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientBridge.Contract.ClientBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientBridge *ClientBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientBridge *ClientBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientBridge *ClientBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientBridge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ClientBridge *ClientBridgeTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientBridge.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ClientBridge *ClientBridgeSession) Admin() (*types.Transaction, error) {
	return _ClientBridge.Contract.Admin(&_ClientBridge.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ClientBridge *ClientBridgeTransactorSession) Admin() (*types.Transaction, error) {
	return _ClientBridge.Contract.Admin(&_ClientBridge.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ClientBridge *ClientBridgeTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ClientBridge.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ClientBridge *ClientBridgeSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ClientBridge.Contract.ChangeAdmin(&_ClientBridge.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ClientBridge *ClientBridgeTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ClientBridge.Contract.ChangeAdmin(&_ClientBridge.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ClientBridge *ClientBridgeTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientBridge.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ClientBridge *ClientBridgeSession) Implementation() (*types.Transaction, error) {
	return _ClientBridge.Contract.Implementation(&_ClientBridge.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ClientBridge *ClientBridgeTransactorSession) Implementation() (*types.Transaction, error) {
	return _ClientBridge.Contract.Implementation(&_ClientBridge.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ClientBridge *ClientBridgeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ClientBridge.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ClientBridge *ClientBridgeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ClientBridge.Contract.UpgradeTo(&_ClientBridge.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ClientBridge *ClientBridgeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ClientBridge.Contract.UpgradeTo(&_ClientBridge.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ClientBridge *ClientBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ClientBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ClientBridge *ClientBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ClientBridge.Contract.UpgradeToAndCall(&_ClientBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ClientBridge *ClientBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ClientBridge.Contract.UpgradeToAndCall(&_ClientBridge.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ClientBridge *ClientBridgeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ClientBridge.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ClientBridge *ClientBridgeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ClientBridge.Contract.Fallback(&_ClientBridge.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ClientBridge *ClientBridgeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ClientBridge.Contract.Fallback(&_ClientBridge.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClientBridge *ClientBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClientBridge *ClientBridgeSession) Receive() (*types.Transaction, error) {
	return _ClientBridge.Contract.Receive(&_ClientBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClientBridge *ClientBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _ClientBridge.Contract.Receive(&_ClientBridge.TransactOpts)
}

// ClientBridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ClientBridge contract.
type ClientBridgeAdminChangedIterator struct {
	Event *ClientBridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *ClientBridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientBridgeAdminChanged)
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
		it.Event = new(ClientBridgeAdminChanged)
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
func (it *ClientBridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientBridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientBridgeAdminChanged represents a AdminChanged event raised by the ClientBridge contract.
type ClientBridgeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ClientBridge *ClientBridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ClientBridgeAdminChangedIterator, error) {

	logs, sub, err := _ClientBridge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ClientBridgeAdminChangedIterator{contract: _ClientBridge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ClientBridge *ClientBridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ClientBridgeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ClientBridge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientBridgeAdminChanged)
				if err := _ClientBridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_ClientBridge *ClientBridgeFilterer) ParseAdminChanged(log types.Log) (*ClientBridgeAdminChanged, error) {
	event := new(ClientBridgeAdminChanged)
	if err := _ClientBridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClientBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ClientBridge contract.
type ClientBridgeUpgradedIterator struct {
	Event *ClientBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *ClientBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientBridgeUpgraded)
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
		it.Event = new(ClientBridgeUpgraded)
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
func (it *ClientBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientBridgeUpgraded represents a Upgraded event raised by the ClientBridge contract.
type ClientBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ClientBridge *ClientBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ClientBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ClientBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ClientBridgeUpgradedIterator{contract: _ClientBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ClientBridge *ClientBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ClientBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ClientBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientBridgeUpgraded)
				if err := _ClientBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ClientBridge *ClientBridgeFilterer) ParseUpgraded(log types.Log) (*ClientBridgeUpgraded, error) {
	event := new(ClientBridgeUpgraded)
	if err := _ClientBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
