// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PaxUSD

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

// PAXUSDABI is the input ABI used to generate the binding from.
const PAXUSDABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// PAXUSD is an auto generated Go binding around an Ethereum contract.
type PAXUSD struct {
	PAXUSDCaller     // Read-only binding to the contract
	PAXUSDTransactor // Write-only binding to the contract
	PAXUSDFilterer   // Log filterer for contract events
}

// PAXUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type PAXUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PAXUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PAXUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PAXUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PAXUSDSession struct {
	Contract     *PAXUSD           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PAXUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PAXUSDCallerSession struct {
	Contract *PAXUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PAXUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PAXUSDTransactorSession struct {
	Contract     *PAXUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PAXUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type PAXUSDRaw struct {
	Contract *PAXUSD // Generic contract binding to access the raw methods on
}

// PAXUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PAXUSDCallerRaw struct {
	Contract *PAXUSDCaller // Generic read-only contract binding to access the raw methods on
}

// PAXUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PAXUSDTransactorRaw struct {
	Contract *PAXUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPAXUSD creates a new instance of PAXUSD, bound to a specific deployed contract.
func NewPAXUSD(address common.Address, backend bind.ContractBackend) (*PAXUSD, error) {
	contract, err := bindPAXUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PAXUSD{PAXUSDCaller: PAXUSDCaller{contract: contract}, PAXUSDTransactor: PAXUSDTransactor{contract: contract}, PAXUSDFilterer: PAXUSDFilterer{contract: contract}}, nil
}

// NewPAXUSDCaller creates a new read-only instance of PAXUSD, bound to a specific deployed contract.
func NewPAXUSDCaller(address common.Address, caller bind.ContractCaller) (*PAXUSDCaller, error) {
	contract, err := bindPAXUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PAXUSDCaller{contract: contract}, nil
}

// NewPAXUSDTransactor creates a new write-only instance of PAXUSD, bound to a specific deployed contract.
func NewPAXUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*PAXUSDTransactor, error) {
	contract, err := bindPAXUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PAXUSDTransactor{contract: contract}, nil
}

// NewPAXUSDFilterer creates a new log filterer instance of PAXUSD, bound to a specific deployed contract.
func NewPAXUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*PAXUSDFilterer, error) {
	contract, err := bindPAXUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PAXUSDFilterer{contract: contract}, nil
}

// bindPAXUSD binds a generic wrapper to an already deployed contract.
func bindPAXUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PAXUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PAXUSD *PAXUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PAXUSD.Contract.PAXUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PAXUSD *PAXUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PAXUSD.Contract.PAXUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PAXUSD *PAXUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PAXUSD.Contract.PAXUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PAXUSD *PAXUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PAXUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PAXUSD *PAXUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PAXUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PAXUSD *PAXUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PAXUSD.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAXUSD *PAXUSDCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PAXUSD.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAXUSD *PAXUSDSession) Admin() (common.Address, error) {
	return _PAXUSD.Contract.Admin(&_PAXUSD.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PAXUSD *PAXUSDCallerSession) Admin() (common.Address, error) {
	return _PAXUSD.Contract.Admin(&_PAXUSD.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAXUSD *PAXUSDCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PAXUSD.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAXUSD *PAXUSDSession) Implementation() (common.Address, error) {
	return _PAXUSD.Contract.Implementation(&_PAXUSD.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PAXUSD *PAXUSDCallerSession) Implementation() (common.Address, error) {
	return _PAXUSD.Contract.Implementation(&_PAXUSD.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAXUSD *PAXUSDTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _PAXUSD.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAXUSD *PAXUSDSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PAXUSD.Contract.ChangeAdmin(&_PAXUSD.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_PAXUSD *PAXUSDTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PAXUSD.Contract.ChangeAdmin(&_PAXUSD.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAXUSD *PAXUSDTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PAXUSD.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAXUSD *PAXUSDSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PAXUSD.Contract.UpgradeTo(&_PAXUSD.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PAXUSD *PAXUSDTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PAXUSD.Contract.UpgradeTo(&_PAXUSD.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAXUSD *PAXUSDTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAXUSD.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAXUSD *PAXUSDSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAXUSD.Contract.UpgradeToAndCall(&_PAXUSD.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PAXUSD *PAXUSDTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PAXUSD.Contract.UpgradeToAndCall(&_PAXUSD.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAXUSD *PAXUSDTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PAXUSD.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAXUSD *PAXUSDSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PAXUSD.Contract.Fallback(&_PAXUSD.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PAXUSD *PAXUSDTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PAXUSD.Contract.Fallback(&_PAXUSD.TransactOpts, calldata)
}

// PAXUSDAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PAXUSD contract.
type PAXUSDAdminChangedIterator struct {
	Event *PAXUSDAdminChanged // Event containing the contract specifics and raw log

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
func (it *PAXUSDAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PAXUSDAdminChanged)
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
		it.Event = new(PAXUSDAdminChanged)
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
func (it *PAXUSDAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PAXUSDAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PAXUSDAdminChanged represents a AdminChanged event raised by the PAXUSD contract.
type PAXUSDAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PAXUSD *PAXUSDFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PAXUSDAdminChangedIterator, error) {

	logs, sub, err := _PAXUSD.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PAXUSDAdminChangedIterator{contract: _PAXUSD.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PAXUSD *PAXUSDFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PAXUSDAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PAXUSD.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PAXUSDAdminChanged)
				if err := _PAXUSD.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_PAXUSD *PAXUSDFilterer) ParseAdminChanged(log types.Log) (*PAXUSDAdminChanged, error) {
	event := new(PAXUSDAdminChanged)
	if err := _PAXUSD.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PAXUSDUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PAXUSD contract.
type PAXUSDUpgradedIterator struct {
	Event *PAXUSDUpgraded // Event containing the contract specifics and raw log

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
func (it *PAXUSDUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PAXUSDUpgraded)
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
		it.Event = new(PAXUSDUpgraded)
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
func (it *PAXUSDUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PAXUSDUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PAXUSDUpgraded represents a Upgraded event raised by the PAXUSD contract.
type PAXUSDUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_PAXUSD *PAXUSDFilterer) FilterUpgraded(opts *bind.FilterOpts) (*PAXUSDUpgradedIterator, error) {

	logs, sub, err := _PAXUSD.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &PAXUSDUpgradedIterator{contract: _PAXUSD.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_PAXUSD *PAXUSDFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PAXUSDUpgraded) (event.Subscription, error) {

	logs, sub, err := _PAXUSD.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PAXUSDUpgraded)
				if err := _PAXUSD.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PAXUSD *PAXUSDFilterer) ParseUpgraded(log types.Log) (*PAXUSDUpgraded, error) {
	event := new(PAXUSDUpgraded)
	if err := _PAXUSD.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
