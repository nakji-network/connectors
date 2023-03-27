// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RenBTC

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

// RENBTCABI is the input ABI used to generate the binding from.
const RENBTCABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// RENBTC is an auto generated Go binding around an Ethereum contract.
type RENBTC struct {
	RENBTCCaller     // Read-only binding to the contract
	RENBTCTransactor // Write-only binding to the contract
	RENBTCFilterer   // Log filterer for contract events
}

// RENBTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type RENBTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RENBTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RENBTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RENBTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RENBTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RENBTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RENBTCSession struct {
	Contract     *RENBTC           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RENBTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RENBTCCallerSession struct {
	Contract *RENBTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RENBTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RENBTCTransactorSession struct {
	Contract     *RENBTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RENBTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type RENBTCRaw struct {
	Contract *RENBTC // Generic contract binding to access the raw methods on
}

// RENBTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RENBTCCallerRaw struct {
	Contract *RENBTCCaller // Generic read-only contract binding to access the raw methods on
}

// RENBTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RENBTCTransactorRaw struct {
	Contract *RENBTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRENBTC creates a new instance of RENBTC, bound to a specific deployed contract.
func NewRENBTC(address common.Address, backend bind.ContractBackend) (*RENBTC, error) {
	contract, err := bindRENBTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RENBTC{RENBTCCaller: RENBTCCaller{contract: contract}, RENBTCTransactor: RENBTCTransactor{contract: contract}, RENBTCFilterer: RENBTCFilterer{contract: contract}}, nil
}

// NewRENBTCCaller creates a new read-only instance of RENBTC, bound to a specific deployed contract.
func NewRENBTCCaller(address common.Address, caller bind.ContractCaller) (*RENBTCCaller, error) {
	contract, err := bindRENBTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RENBTCCaller{contract: contract}, nil
}

// NewRENBTCTransactor creates a new write-only instance of RENBTC, bound to a specific deployed contract.
func NewRENBTCTransactor(address common.Address, transactor bind.ContractTransactor) (*RENBTCTransactor, error) {
	contract, err := bindRENBTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RENBTCTransactor{contract: contract}, nil
}

// NewRENBTCFilterer creates a new log filterer instance of RENBTC, bound to a specific deployed contract.
func NewRENBTCFilterer(address common.Address, filterer bind.ContractFilterer) (*RENBTCFilterer, error) {
	contract, err := bindRENBTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RENBTCFilterer{contract: contract}, nil
}

// bindRENBTC binds a generic wrapper to an already deployed contract.
func bindRENBTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RENBTCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RENBTC *RENBTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RENBTC.Contract.RENBTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RENBTC *RENBTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RENBTC.Contract.RENBTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RENBTC *RENBTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RENBTC.Contract.RENBTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RENBTC *RENBTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RENBTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RENBTC *RENBTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RENBTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RENBTC *RENBTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RENBTC.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_RENBTC *RENBTCTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_RENBTC *RENBTCSession) Admin() (*types.Transaction, error) {
	return _RENBTC.Contract.Admin(&_RENBTC.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_RENBTC *RENBTCTransactorSession) Admin() (*types.Transaction, error) {
	return _RENBTC.Contract.Admin(&_RENBTC.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_RENBTC *RENBTCTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_RENBTC *RENBTCSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _RENBTC.Contract.ChangeAdmin(&_RENBTC.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_RENBTC *RENBTCTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _RENBTC.Contract.ChangeAdmin(&_RENBTC.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_RENBTC *RENBTCTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_RENBTC *RENBTCSession) Implementation() (*types.Transaction, error) {
	return _RENBTC.Contract.Implementation(&_RENBTC.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_RENBTC *RENBTCTransactorSession) Implementation() (*types.Transaction, error) {
	return _RENBTC.Contract.Implementation(&_RENBTC.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_RENBTC *RENBTCTransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "initialize", _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_RENBTC *RENBTCSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Initialize(&_RENBTC.TransactOpts, _logic, _admin, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address _logic, address _admin, bytes _data) payable returns()
func (_RENBTC *RENBTCTransactorSession) Initialize(_logic common.Address, _admin common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Initialize(&_RENBTC.TransactOpts, _logic, _admin, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_RENBTC *RENBTCTransactor) Initialize0(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "initialize0", _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_RENBTC *RENBTCSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Initialize0(&_RENBTC.TransactOpts, _logic, _data)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_RENBTC *RENBTCTransactorSession) Initialize0(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Initialize0(&_RENBTC.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RENBTC *RENBTCTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RENBTC *RENBTCSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RENBTC.Contract.UpgradeTo(&_RENBTC.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RENBTC *RENBTCTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RENBTC.Contract.UpgradeTo(&_RENBTC.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RENBTC *RENBTCTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RENBTC.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RENBTC *RENBTCSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.UpgradeToAndCall(&_RENBTC.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RENBTC *RENBTCTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.UpgradeToAndCall(&_RENBTC.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RENBTC *RENBTCTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RENBTC.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RENBTC *RENBTCSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Fallback(&_RENBTC.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RENBTC *RENBTCTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RENBTC.Contract.Fallback(&_RENBTC.TransactOpts, calldata)
}

// RENBTCAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RENBTC contract.
type RENBTCAdminChangedIterator struct {
	Event *RENBTCAdminChanged // Event containing the contract specifics and raw log

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
func (it *RENBTCAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RENBTCAdminChanged)
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
		it.Event = new(RENBTCAdminChanged)
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
func (it *RENBTCAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RENBTCAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RENBTCAdminChanged represents a AdminChanged event raised by the RENBTC contract.
type RENBTCAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RENBTC *RENBTCFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RENBTCAdminChangedIterator, error) {

	logs, sub, err := _RENBTC.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RENBTCAdminChangedIterator{contract: _RENBTC.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RENBTC *RENBTCFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RENBTCAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RENBTC.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RENBTCAdminChanged)
				if err := _RENBTC.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_RENBTC *RENBTCFilterer) ParseAdminChanged(log types.Log) (*RENBTCAdminChanged, error) {
	event := new(RENBTCAdminChanged)
	if err := _RENBTC.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RENBTCUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RENBTC contract.
type RENBTCUpgradedIterator struct {
	Event *RENBTCUpgraded // Event containing the contract specifics and raw log

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
func (it *RENBTCUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RENBTCUpgraded)
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
		it.Event = new(RENBTCUpgraded)
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
func (it *RENBTCUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RENBTCUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RENBTCUpgraded represents a Upgraded event raised by the RENBTC contract.
type RENBTCUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RENBTC *RENBTCFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RENBTCUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RENBTC.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RENBTCUpgradedIterator{contract: _RENBTC.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RENBTC *RENBTCFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RENBTCUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RENBTC.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RENBTCUpgraded)
				if err := _RENBTC.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_RENBTC *RENBTCFilterer) ParseUpgraded(log types.Log) (*RENBTCUpgraded, error) {
	event := new(RENBTCUpgraded)
	if err := _RENBTC.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
