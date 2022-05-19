// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_FACTORY

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

// PROXYFACTORYABI is the input ABI used to generate the binding from.
const PROXYFACTORYABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cache\",\"type\":\"address\"}],\"name\":\"Created\",\"type\":\"event\"}]"

// PROXYFACTORY is an auto generated Go binding around an Ethereum contract.
type PROXYFACTORY struct {
	PROXYFACTORYCaller     // Read-only binding to the contract
	PROXYFACTORYTransactor // Write-only binding to the contract
	PROXYFACTORYFilterer   // Log filterer for contract events
}

// PROXYFACTORYCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYFACTORYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYFACTORYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYFACTORYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYFACTORYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYFACTORYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYFACTORYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYFACTORYSession struct {
	Contract     *PROXYFACTORY     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PROXYFACTORYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYFACTORYCallerSession struct {
	Contract *PROXYFACTORYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PROXYFACTORYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYFACTORYTransactorSession struct {
	Contract     *PROXYFACTORYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PROXYFACTORYRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYFACTORYRaw struct {
	Contract *PROXYFACTORY // Generic contract binding to access the raw methods on
}

// PROXYFACTORYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYFACTORYCallerRaw struct {
	Contract *PROXYFACTORYCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYFACTORYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYFACTORYTransactorRaw struct {
	Contract *PROXYFACTORYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYFACTORY creates a new instance of PROXYFACTORY, bound to a specific deployed contract.
func NewPROXYFACTORY(address common.Address, backend bind.ContractBackend) (*PROXYFACTORY, error) {
	contract, err := bindPROXYFACTORY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYFACTORY{PROXYFACTORYCaller: PROXYFACTORYCaller{contract: contract}, PROXYFACTORYTransactor: PROXYFACTORYTransactor{contract: contract}, PROXYFACTORYFilterer: PROXYFACTORYFilterer{contract: contract}}, nil
}

// NewPROXYFACTORYCaller creates a new read-only instance of PROXYFACTORY, bound to a specific deployed contract.
func NewPROXYFACTORYCaller(address common.Address, caller bind.ContractCaller) (*PROXYFACTORYCaller, error) {
	contract, err := bindPROXYFACTORY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYFACTORYCaller{contract: contract}, nil
}

// NewPROXYFACTORYTransactor creates a new write-only instance of PROXYFACTORY, bound to a specific deployed contract.
func NewPROXYFACTORYTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYFACTORYTransactor, error) {
	contract, err := bindPROXYFACTORY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYFACTORYTransactor{contract: contract}, nil
}

// NewPROXYFACTORYFilterer creates a new log filterer instance of PROXYFACTORY, bound to a specific deployed contract.
func NewPROXYFACTORYFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYFACTORYFilterer, error) {
	contract, err := bindPROXYFACTORY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYFACTORYFilterer{contract: contract}, nil
}

// bindPROXYFACTORY binds a generic wrapper to an already deployed contract.
func bindPROXYFACTORY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYFACTORYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYFACTORY *PROXYFACTORYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYFACTORY.Contract.PROXYFACTORYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYFACTORY *PROXYFACTORYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.PROXYFACTORYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYFACTORY *PROXYFACTORYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.PROXYFACTORYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYFACTORY *PROXYFACTORYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYFACTORY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYFACTORY *PROXYFACTORYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYFACTORY *PROXYFACTORYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.contract.Transact(opts, method, params...)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYFACTORY *PROXYFACTORYCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYFACTORY.contract.Call(opts, &out, "cache")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYFACTORY *PROXYFACTORYSession) Cache() (common.Address, error) {
	return _PROXYFACTORY.Contract.Cache(&_PROXYFACTORY.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYFACTORY *PROXYFACTORYCallerSession) Cache() (common.Address, error) {
	return _PROXYFACTORY.Contract.Cache(&_PROXYFACTORY.CallOpts)
}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) view returns(bool)
func (_PROXYFACTORY *PROXYFACTORYCaller) IsProxy(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PROXYFACTORY.contract.Call(opts, &out, "isProxy", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) view returns(bool)
func (_PROXYFACTORY *PROXYFACTORYSession) IsProxy(arg0 common.Address) (bool, error) {
	return _PROXYFACTORY.Contract.IsProxy(&_PROXYFACTORY.CallOpts, arg0)
}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) view returns(bool)
func (_PROXYFACTORY *PROXYFACTORYCallerSession) IsProxy(arg0 common.Address) (bool, error) {
	return _PROXYFACTORY.Contract.IsProxy(&_PROXYFACTORY.CallOpts, arg0)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYTransactor) Build(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYFACTORY.contract.Transact(opts, "build")
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYSession) Build() (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.Build(&_PROXYFACTORY.TransactOpts)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYTransactorSession) Build() (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.Build(&_PROXYFACTORY.TransactOpts)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYTransactor) Build0(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _PROXYFACTORY.contract.Transact(opts, "build0", owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYSession) Build0(owner common.Address) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.Build0(&_PROXYFACTORY.TransactOpts, owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_PROXYFACTORY *PROXYFACTORYTransactorSession) Build0(owner common.Address) (*types.Transaction, error) {
	return _PROXYFACTORY.Contract.Build0(&_PROXYFACTORY.TransactOpts, owner)
}

// PROXYFACTORYCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the PROXYFACTORY contract.
type PROXYFACTORYCreatedIterator struct {
	Event *PROXYFACTORYCreated // Event containing the contract specifics and raw log

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
func (it *PROXYFACTORYCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PROXYFACTORYCreated)
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
		it.Event = new(PROXYFACTORYCreated)
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
func (it *PROXYFACTORYCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PROXYFACTORYCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PROXYFACTORYCreated represents a Created event raised by the PROXYFACTORY contract.
type PROXYFACTORYCreated struct {
	Sender common.Address
	Owner  common.Address
	Proxy  common.Address
	Cache  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b.
//
// Solidity: event Created(address indexed sender, address indexed owner, address proxy, address cache)
func (_PROXYFACTORY *PROXYFACTORYFilterer) FilterCreated(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*PROXYFACTORYCreatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PROXYFACTORY.contract.FilterLogs(opts, "Created", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PROXYFACTORYCreatedIterator{contract: _PROXYFACTORY.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b.
//
// Solidity: event Created(address indexed sender, address indexed owner, address proxy, address cache)
func (_PROXYFACTORY *PROXYFACTORYFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *PROXYFACTORYCreated, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PROXYFACTORY.contract.WatchLogs(opts, "Created", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PROXYFACTORYCreated)
				if err := _PROXYFACTORY.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b.
//
// Solidity: event Created(address indexed sender, address indexed owner, address proxy, address cache)
func (_PROXYFACTORY *PROXYFACTORYFilterer) ParseCreated(log types.Log) (*PROXYFACTORYCreated, error) {
	event := new(PROXYFACTORYCreated)
	if err := _PROXYFACTORY.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
