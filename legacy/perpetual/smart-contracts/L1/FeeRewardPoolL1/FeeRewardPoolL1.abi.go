// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FeeRewardPoolL1

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

// FeeRewardPoolL1ABI is the input ABI used to generate the binding from.
const FeeRewardPoolL1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// FeeRewardPoolL1 is an auto generated Go binding around an Ethereum contract.
type FeeRewardPoolL1 struct {
	FeeRewardPoolL1Caller     // Read-only binding to the contract
	FeeRewardPoolL1Transactor // Write-only binding to the contract
	FeeRewardPoolL1Filterer   // Log filterer for contract events
}

// FeeRewardPoolL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeeRewardPoolL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRewardPoolL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeRewardPoolL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRewardPoolL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeRewardPoolL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRewardPoolL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeRewardPoolL1Session struct {
	Contract     *FeeRewardPoolL1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeRewardPoolL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeRewardPoolL1CallerSession struct {
	Contract *FeeRewardPoolL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FeeRewardPoolL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeRewardPoolL1TransactorSession struct {
	Contract     *FeeRewardPoolL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FeeRewardPoolL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeeRewardPoolL1Raw struct {
	Contract *FeeRewardPoolL1 // Generic contract binding to access the raw methods on
}

// FeeRewardPoolL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeRewardPoolL1CallerRaw struct {
	Contract *FeeRewardPoolL1Caller // Generic read-only contract binding to access the raw methods on
}

// FeeRewardPoolL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeRewardPoolL1TransactorRaw struct {
	Contract *FeeRewardPoolL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeRewardPoolL1 creates a new instance of FeeRewardPoolL1, bound to a specific deployed contract.
func NewFeeRewardPoolL1(address common.Address, backend bind.ContractBackend) (*FeeRewardPoolL1, error) {
	contract, err := bindFeeRewardPoolL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1{FeeRewardPoolL1Caller: FeeRewardPoolL1Caller{contract: contract}, FeeRewardPoolL1Transactor: FeeRewardPoolL1Transactor{contract: contract}, FeeRewardPoolL1Filterer: FeeRewardPoolL1Filterer{contract: contract}}, nil
}

// NewFeeRewardPoolL1Caller creates a new read-only instance of FeeRewardPoolL1, bound to a specific deployed contract.
func NewFeeRewardPoolL1Caller(address common.Address, caller bind.ContractCaller) (*FeeRewardPoolL1Caller, error) {
	contract, err := bindFeeRewardPoolL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1Caller{contract: contract}, nil
}

// NewFeeRewardPoolL1Transactor creates a new write-only instance of FeeRewardPoolL1, bound to a specific deployed contract.
func NewFeeRewardPoolL1Transactor(address common.Address, transactor bind.ContractTransactor) (*FeeRewardPoolL1Transactor, error) {
	contract, err := bindFeeRewardPoolL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1Transactor{contract: contract}, nil
}

// NewFeeRewardPoolL1Filterer creates a new log filterer instance of FeeRewardPoolL1, bound to a specific deployed contract.
func NewFeeRewardPoolL1Filterer(address common.Address, filterer bind.ContractFilterer) (*FeeRewardPoolL1Filterer, error) {
	contract, err := bindFeeRewardPoolL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1Filterer{contract: contract}, nil
}

// bindFeeRewardPoolL1 binds a generic wrapper to an already deployed contract.
func bindFeeRewardPoolL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeRewardPoolL1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeRewardPoolL1 *FeeRewardPoolL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRewardPoolL1.Contract.FeeRewardPoolL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeRewardPoolL1 *FeeRewardPoolL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.FeeRewardPoolL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeRewardPoolL1 *FeeRewardPoolL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.FeeRewardPoolL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeRewardPoolL1 *FeeRewardPoolL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRewardPoolL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) Admin() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Admin(&_FeeRewardPoolL1.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) Admin() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Admin(&_FeeRewardPoolL1.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.ChangeAdmin(&_FeeRewardPoolL1.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.ChangeAdmin(&_FeeRewardPoolL1.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) Implementation() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Implementation(&_FeeRewardPoolL1.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) Implementation() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Implementation(&_FeeRewardPoolL1.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.UpgradeTo(&_FeeRewardPoolL1.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.UpgradeTo(&_FeeRewardPoolL1.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.UpgradeToAndCall(&_FeeRewardPoolL1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.UpgradeToAndCall(&_FeeRewardPoolL1.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Fallback(&_FeeRewardPoolL1.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Fallback(&_FeeRewardPoolL1.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRewardPoolL1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1Session) Receive() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Receive(&_FeeRewardPoolL1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeRewardPoolL1 *FeeRewardPoolL1TransactorSession) Receive() (*types.Transaction, error) {
	return _FeeRewardPoolL1.Contract.Receive(&_FeeRewardPoolL1.TransactOpts)
}

// FeeRewardPoolL1AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FeeRewardPoolL1 contract.
type FeeRewardPoolL1AdminChangedIterator struct {
	Event *FeeRewardPoolL1AdminChanged // Event containing the contract specifics and raw log

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
func (it *FeeRewardPoolL1AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRewardPoolL1AdminChanged)
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
		it.Event = new(FeeRewardPoolL1AdminChanged)
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
func (it *FeeRewardPoolL1AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRewardPoolL1AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRewardPoolL1AdminChanged represents a AdminChanged event raised by the FeeRewardPoolL1 contract.
type FeeRewardPoolL1AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*FeeRewardPoolL1AdminChangedIterator, error) {

	logs, sub, err := _FeeRewardPoolL1.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1AdminChangedIterator{contract: _FeeRewardPoolL1.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeRewardPoolL1AdminChanged) (event.Subscription, error) {

	logs, sub, err := _FeeRewardPoolL1.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRewardPoolL1AdminChanged)
				if err := _FeeRewardPoolL1.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) ParseAdminChanged(log types.Log) (*FeeRewardPoolL1AdminChanged, error) {
	event := new(FeeRewardPoolL1AdminChanged)
	if err := _FeeRewardPoolL1.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeRewardPoolL1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FeeRewardPoolL1 contract.
type FeeRewardPoolL1UpgradedIterator struct {
	Event *FeeRewardPoolL1Upgraded // Event containing the contract specifics and raw log

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
func (it *FeeRewardPoolL1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRewardPoolL1Upgraded)
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
		it.Event = new(FeeRewardPoolL1Upgraded)
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
func (it *FeeRewardPoolL1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRewardPoolL1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRewardPoolL1Upgraded represents a Upgraded event raised by the FeeRewardPoolL1 contract.
type FeeRewardPoolL1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FeeRewardPoolL1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeRewardPoolL1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FeeRewardPoolL1UpgradedIterator{contract: _FeeRewardPoolL1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FeeRewardPoolL1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeRewardPoolL1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRewardPoolL1Upgraded)
				if err := _FeeRewardPoolL1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FeeRewardPoolL1 *FeeRewardPoolL1Filterer) ParseUpgraded(log types.Log) (*FeeRewardPoolL1Upgraded, error) {
	event := new(FeeRewardPoolL1Upgraded)
	if err := _FeeRewardPoolL1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
