// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PAX

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

// PAXABI is the input ABI used to generate the binding from.
const PAXABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// PAX is an auto generated Go binding around an Ethereum contract.
type PAX struct {
	PAXCaller     // Read-only binding to the contract
	PAXTransactor // Write-only binding to the contract
	PAXFilterer   // Log filterer for contract events
}

// PAXCaller is an auto generated read-only Go binding around an Ethereum contract.
type PAXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PAXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PAXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PAXSession struct {
	Contract     *PAX              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PAXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PAXCallerSession struct {
	Contract *PAXCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PAXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PAXTransactorSession struct {
	Contract     *PAXTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PAXRaw is an auto generated low-level Go binding around an Ethereum contract.
type PAXRaw struct {
	Contract *PAX // Generic contract binding to access the raw methods on
}

// PAXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PAXCallerRaw struct {
	Contract *PAXCaller // Generic read-only contract binding to access the raw methods on
}

// PAXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PAXTransactorRaw struct {
	Contract *PAXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPAX creates a new instance of PAX, bound to a specific deployed contract.
func NewPAX(address common.Address, backend bind.ContractBackend) (*PAX, error) {
	contract, err := bindPAX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PAX{PAXCaller: PAXCaller{contract: contract}, PAXTransactor: PAXTransactor{contract: contract}, PAXFilterer: PAXFilterer{contract: contract}}, nil
}

// NewPAXCaller creates a new read-only instance of PAX, bound to a specific deployed contract.
func NewPAXCaller(address common.Address, caller bind.ContractCaller) (*PAXCaller, error) {
	contract, err := bindPAX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PAXCaller{contract: contract}, nil
}

// NewPAXTransactor creates a new write-only instance of PAX, bound to a specific deployed contract.
func NewPAXTransactor(address common.Address, transactor bind.ContractTransactor) (*PAXTransactor, error) {
	contract, err := bindPAX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PAXTransactor{contract: contract}, nil
}

// NewPAXFilterer creates a new log filterer instance of PAX, bound to a specific deployed contract.
func NewPAXFilterer(address common.Address, filterer bind.ContractFilterer) (*PAXFilterer, error) {
	contract, err := bindPAX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PAXFilterer{contract: contract}, nil
}

// bindPAX binds a generic wrapper to an already deployed contract.
func bindPAX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PAXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PAX *PAXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PAX.Contract.PAXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PAX *PAXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PAX.Contract.PAXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PAX *PAXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PAX.Contract.PAXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PAX *PAXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PAX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PAX *PAXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PAX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PAX *PAXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PAX.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAX *PAXCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PAX.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAX *PAXSession) Admin() (common.Address, error) {
	return _PAX.Contract.Admin(&_PAX.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAX *PAXCallerSession) Admin() (common.Address, error) {
	return _PAX.Contract.Admin(&_PAX.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAX *PAXCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PAX.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAX *PAXSession) Implementation() (common.Address, error) {
	return _PAX.Contract.Implementation(&_PAX.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAX *PAXCallerSession) Implementation() (common.Address, error) {
	return _PAX.Contract.Implementation(&_PAX.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAX *PAXTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _PAX.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAX *PAXSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PAX.Contract.ChangeAdmin(&_PAX.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAX *PAXTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PAX.Contract.ChangeAdmin(&_PAX.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAX *PAXTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PAX.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAX *PAXSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PAX.Contract.UpgradeTo(&_PAX.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAX *PAXTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PAX.Contract.UpgradeTo(&_PAX.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAX *PAXTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAX.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAX *PAXSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAX.Contract.UpgradeToAndCall(&_PAX.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAX *PAXTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAX.Contract.UpgradeToAndCall(&_PAX.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAX *PAXTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PAX.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAX *PAXSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PAX.Contract.Fallback(&_PAX.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAX *PAXTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PAX.Contract.Fallback(&_PAX.TransactOpts, calldata)
}

// PAXAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PAX contract.
type PAXAdminChangedIterator struct {
	Event *PAXAdminChanged // Event containing the contract specifics and raw log

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
func (it *PAXAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PAXAdminChanged)
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
		it.Event = new(PAXAdminChanged)
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
func (it *PAXAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PAXAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PAXAdminChanged represents a AdminChanged event raised by the PAX contract.
type PAXAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PAX *PAXFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PAXAdminChangedIterator, error) {

	logs, sub, err := _PAX.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PAXAdminChangedIterator{contract: _PAX.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PAX *PAXFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PAXAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PAX.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PAXAdminChanged)
				if err := _PAX.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_PAX *PAXFilterer) ParseAdminChanged(log types.Log) (*PAXAdminChanged, error) {
	event := new(PAXAdminChanged)
	if err := _PAX.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PAXUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PAX contract.
type PAXUpgradedIterator struct {
	Event *PAXUpgraded // Event containing the contract specifics and raw log

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
func (it *PAXUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PAXUpgraded)
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
		it.Event = new(PAXUpgraded)
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
func (it *PAXUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PAXUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PAXUpgraded represents a Upgraded event raised by the PAX contract.
type PAXUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_PAX *PAXFilterer) FilterUpgraded(opts *bind.FilterOpts) (*PAXUpgradedIterator, error) {

	logs, sub, err := _PAX.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &PAXUpgradedIterator{contract: _PAX.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_PAX *PAXFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PAXUpgraded) (event.Subscription, error) {

	logs, sub, err := _PAX.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PAXUpgraded)
				if err := _PAX.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
// Solidity: event Upgraded(address implementation)
func (_PAX *PAXFilterer) ParseUpgraded(log types.Log) (*PAXUpgraded, error) {
	event := new(PAXUpgraded)
	if err := _PAX.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
