// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AAVE

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

// AAVEABI is the input ABI used to generate the binding from.
const AAVEABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// AAVE is an auto generated Go binding around an Ethereum contract.
type AAVE struct {
	AAVECaller     // Read-only binding to the contract
	AAVETransactor // Write-only binding to the contract
	AAVEFilterer   // Log filterer for contract events
}

// AAVECaller is an auto generated read-only Go binding around an Ethereum contract.
type AAVECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAVETransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAVETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAVEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAVEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAVESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAVESession struct {
	Contract     *AAVE             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAVECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAVECallerSession struct {
	Contract *AAVECaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AAVETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAVETransactorSession struct {
	Contract     *AAVETransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAVERaw is an auto generated low-level Go binding around an Ethereum contract.
type AAVERaw struct {
	Contract *AAVE // Generic contract binding to access the raw methods on
}

// AAVECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAVECallerRaw struct {
	Contract *AAVECaller // Generic read-only contract binding to access the raw methods on
}

// AAVETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAVETransactorRaw struct {
	Contract *AAVETransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAVE creates a new instance of AAVE, bound to a specific deployed contract.
func NewAAVE(address common.Address, backend bind.ContractBackend) (*AAVE, error) {
	contract, err := bindAAVE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAVE{AAVECaller: AAVECaller{contract: contract}, AAVETransactor: AAVETransactor{contract: contract}, AAVEFilterer: AAVEFilterer{contract: contract}}, nil
}

// NewAAVECaller creates a new read-only instance of AAVE, bound to a specific deployed contract.
func NewAAVECaller(address common.Address, caller bind.ContractCaller) (*AAVECaller, error) {
	contract, err := bindAAVE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAVECaller{contract: contract}, nil
}

// NewAAVETransactor creates a new write-only instance of AAVE, bound to a specific deployed contract.
func NewAAVETransactor(address common.Address, transactor bind.ContractTransactor) (*AAVETransactor, error) {
	contract, err := bindAAVE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAVETransactor{contract: contract}, nil
}

// NewAAVEFilterer creates a new log filterer instance of AAVE, bound to a specific deployed contract.
func NewAAVEFilterer(address common.Address, filterer bind.ContractFilterer) (*AAVEFilterer, error) {
	contract, err := bindAAVE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAVEFilterer{contract: contract}, nil
}

// bindAAVE binds a generic wrapper to an already deployed contract.
func bindAAVE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AAVEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAVE *AAVERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAVE.Contract.AAVECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAVE *AAVERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAVE.Contract.AAVETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAVE *AAVERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAVE.Contract.AAVETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAVE *AAVECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAVE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAVE *AAVETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAVE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAVE *AAVETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAVE.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AAVE *AAVETransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AAVE *AAVESession) Admin() (*types.Transaction, error) {
	return _AAVE.Contract.Admin(&_AAVE.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AAVE *AAVETransactorSession) Admin() (*types.Transaction, error) {
	return _AAVE.Contract.Admin(&_AAVE.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AAVE *AAVETransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AAVE *AAVESession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AAVE.Contract.ChangeAdmin(&_AAVE.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AAVE *AAVETransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AAVE.Contract.ChangeAdmin(&_AAVE.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AAVE *AAVETransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AAVE *AAVESession) Implementation() (*types.Transaction, error) {
	return _AAVE.Contract.Implementation(&_AAVE.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AAVE *AAVETransactorSession) Implementation() (*types.Transaction, error) {
	return _AAVE.Contract.Implementation(&_AAVE.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_AAVE *AAVETransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "initialize", _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_AAVE *AAVESession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Initialize(&_AAVE.TransactOpts, _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_AAVE *AAVETransactorSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Initialize(&_AAVE.TransactOpts, _logic, _admin, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AAVE *AAVETransactor) Initialize0(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "initialize0", _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AAVE *AAVESession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Initialize0(&_AAVE.TransactOpts, _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AAVE *AAVETransactorSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Initialize0(&_AAVE.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AAVE *AAVETransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AAVE *AAVESession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AAVE.Contract.UpgradeTo(&_AAVE.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AAVE *AAVETransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AAVE.Contract.UpgradeTo(&_AAVE.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAVE *AAVETransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAVE.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAVE *AAVESession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.UpgradeToAndCall(&_AAVE.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAVE *AAVETransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAVE.Contract.UpgradeToAndCall(&_AAVE.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AAVE *AAVETransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AAVE.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AAVE *AAVESession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Fallback(&_AAVE.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AAVE *AAVETransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AAVE.Contract.Fallback(&_AAVE.TransactOpts, calldata)
}

// AAVEAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AAVE contract.
type AAVEAdminChangedIterator struct {
	Event *AAVEAdminChanged // Event containing the contract specifics and raw log

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
func (it *AAVEAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAVEAdminChanged)
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
		it.Event = new(AAVEAdminChanged)
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
func (it *AAVEAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAVEAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAVEAdminChanged represents a AdminChanged event raised by the AAVE contract.
type AAVEAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AAVE *AAVEFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AAVEAdminChangedIterator, error) {

	logs, sub, err := _AAVE.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AAVEAdminChangedIterator{contract: _AAVE.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AAVE *AAVEFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AAVEAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AAVE.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAVEAdminChanged)
				if err := _AAVE.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AAVE *AAVEFilterer) ParseAdminChanged(log types.Log) (*AAVEAdminChanged, error) {
	event := new(AAVEAdminChanged)
	if err := _AAVE.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAVEUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AAVE contract.
type AAVEUpgradedIterator struct {
	Event *AAVEUpgraded // Event containing the contract specifics and raw log

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
func (it *AAVEUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAVEUpgraded)
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
		it.Event = new(AAVEUpgraded)
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
func (it *AAVEUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAVEUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAVEUpgraded represents a Upgraded event raised by the AAVE contract.
type AAVEUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAVE *AAVEFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AAVEUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAVE.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AAVEUpgradedIterator{contract: _AAVE.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAVE *AAVEFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AAVEUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAVE.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAVEUpgraded)
				if err := _AAVE.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AAVE *AAVEFilterer) ParseUpgraded(log types.Log) (*AAVEUpgraded, error) {
	event := new(AAVEUpgraded)
	if err := _AAVE.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
