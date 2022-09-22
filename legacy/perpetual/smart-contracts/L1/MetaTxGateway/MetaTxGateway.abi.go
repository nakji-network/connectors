// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MetaTxGateway

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

// MetaTxGatewayABI is the input ABI used to generate the binding from.
const MetaTxGatewayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// MetaTxGateway is an auto generated Go binding around an Ethereum contract.
type MetaTxGateway struct {
	MetaTxGatewayCaller     // Read-only binding to the contract
	MetaTxGatewayTransactor // Write-only binding to the contract
	MetaTxGatewayFilterer   // Log filterer for contract events
}

// MetaTxGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaTxGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaTxGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaTxGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaTxGatewaySession struct {
	Contract     *MetaTxGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaTxGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaTxGatewayCallerSession struct {
	Contract *MetaTxGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MetaTxGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaTxGatewayTransactorSession struct {
	Contract     *MetaTxGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MetaTxGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaTxGatewayRaw struct {
	Contract *MetaTxGateway // Generic contract binding to access the raw methods on
}

// MetaTxGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaTxGatewayCallerRaw struct {
	Contract *MetaTxGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// MetaTxGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaTxGatewayTransactorRaw struct {
	Contract *MetaTxGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaTxGateway creates a new instance of MetaTxGateway, bound to a specific deployed contract.
func NewMetaTxGateway(address common.Address, backend bind.ContractBackend) (*MetaTxGateway, error) {
	contract, err := bindMetaTxGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaTxGateway{MetaTxGatewayCaller: MetaTxGatewayCaller{contract: contract}, MetaTxGatewayTransactor: MetaTxGatewayTransactor{contract: contract}, MetaTxGatewayFilterer: MetaTxGatewayFilterer{contract: contract}}, nil
}

// NewMetaTxGatewayCaller creates a new read-only instance of MetaTxGateway, bound to a specific deployed contract.
func NewMetaTxGatewayCaller(address common.Address, caller bind.ContractCaller) (*MetaTxGatewayCaller, error) {
	contract, err := bindMetaTxGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTxGatewayCaller{contract: contract}, nil
}

// NewMetaTxGatewayTransactor creates a new write-only instance of MetaTxGateway, bound to a specific deployed contract.
func NewMetaTxGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaTxGatewayTransactor, error) {
	contract, err := bindMetaTxGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTxGatewayTransactor{contract: contract}, nil
}

// NewMetaTxGatewayFilterer creates a new log filterer instance of MetaTxGateway, bound to a specific deployed contract.
func NewMetaTxGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaTxGatewayFilterer, error) {
	contract, err := bindMetaTxGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaTxGatewayFilterer{contract: contract}, nil
}

// bindMetaTxGateway binds a generic wrapper to an already deployed contract.
func bindMetaTxGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaTxGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTxGateway *MetaTxGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaTxGateway.Contract.MetaTxGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTxGateway *MetaTxGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.MetaTxGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTxGateway *MetaTxGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.MetaTxGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTxGateway *MetaTxGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaTxGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTxGateway *MetaTxGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTxGateway *MetaTxGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_MetaTxGateway *MetaTxGatewayTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxGateway.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_MetaTxGateway *MetaTxGatewaySession) Admin() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Admin(&_MetaTxGateway.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_MetaTxGateway *MetaTxGatewayTransactorSession) Admin() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Admin(&_MetaTxGateway.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_MetaTxGateway *MetaTxGatewayTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_MetaTxGateway *MetaTxGatewaySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.ChangeAdmin(&_MetaTxGateway.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_MetaTxGateway *MetaTxGatewayTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.ChangeAdmin(&_MetaTxGateway.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_MetaTxGateway *MetaTxGatewayTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxGateway.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_MetaTxGateway *MetaTxGatewaySession) Implementation() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Implementation(&_MetaTxGateway.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_MetaTxGateway *MetaTxGatewayTransactorSession) Implementation() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Implementation(&_MetaTxGateway.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MetaTxGateway *MetaTxGatewayTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MetaTxGateway *MetaTxGatewaySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.UpgradeTo(&_MetaTxGateway.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MetaTxGateway *MetaTxGatewayTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.UpgradeTo(&_MetaTxGateway.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MetaTxGateway.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MetaTxGateway *MetaTxGatewaySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.UpgradeToAndCall(&_MetaTxGateway.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.UpgradeToAndCall(&_MetaTxGateway.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MetaTxGateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaTxGateway *MetaTxGatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Fallback(&_MetaTxGateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Fallback(&_MetaTxGateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaTxGateway *MetaTxGatewaySession) Receive() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Receive(&_MetaTxGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaTxGateway *MetaTxGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _MetaTxGateway.Contract.Receive(&_MetaTxGateway.TransactOpts)
}

// MetaTxGatewayAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the MetaTxGateway contract.
type MetaTxGatewayAdminChangedIterator struct {
	Event *MetaTxGatewayAdminChanged // Event containing the contract specifics and raw log

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
func (it *MetaTxGatewayAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaTxGatewayAdminChanged)
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
		it.Event = new(MetaTxGatewayAdminChanged)
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
func (it *MetaTxGatewayAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaTxGatewayAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaTxGatewayAdminChanged represents a AdminChanged event raised by the MetaTxGateway contract.
type MetaTxGatewayAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MetaTxGateway *MetaTxGatewayFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MetaTxGatewayAdminChangedIterator, error) {

	logs, sub, err := _MetaTxGateway.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MetaTxGatewayAdminChangedIterator{contract: _MetaTxGateway.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MetaTxGateway *MetaTxGatewayFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MetaTxGatewayAdminChanged) (event.Subscription, error) {

	logs, sub, err := _MetaTxGateway.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaTxGatewayAdminChanged)
				if err := _MetaTxGateway.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_MetaTxGateway *MetaTxGatewayFilterer) ParseAdminChanged(log types.Log) (*MetaTxGatewayAdminChanged, error) {
	event := new(MetaTxGatewayAdminChanged)
	if err := _MetaTxGateway.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaTxGatewayUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MetaTxGateway contract.
type MetaTxGatewayUpgradedIterator struct {
	Event *MetaTxGatewayUpgraded // Event containing the contract specifics and raw log

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
func (it *MetaTxGatewayUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaTxGatewayUpgraded)
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
		it.Event = new(MetaTxGatewayUpgraded)
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
func (it *MetaTxGatewayUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaTxGatewayUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaTxGatewayUpgraded represents a Upgraded event raised by the MetaTxGateway contract.
type MetaTxGatewayUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MetaTxGateway *MetaTxGatewayFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MetaTxGatewayUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MetaTxGateway.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MetaTxGatewayUpgradedIterator{contract: _MetaTxGateway.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MetaTxGateway *MetaTxGatewayFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MetaTxGatewayUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MetaTxGateway.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaTxGatewayUpgraded)
				if err := _MetaTxGateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MetaTxGateway *MetaTxGatewayFilterer) ParseUpgraded(log types.Log) (*MetaTxGatewayUpgraded, error) {
	event := new(MetaTxGatewayUpgraded)
	if err := _MetaTxGateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
