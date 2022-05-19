// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OPTIMISM_GOV_RELAY

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

// OPTIMISMGOVRELAYABI is the input ABI used to generate the binding from.
const OPTIMISMGOVRELAYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2GovernanceRelay\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2GovernanceRelay\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"l2gas\",\"type\":\"uint32\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISMGOVRELAY is an auto generated Go binding around an Ethereum contract.
type OPTIMISMGOVRELAY struct {
	OPTIMISMGOVRELAYCaller     // Read-only binding to the contract
	OPTIMISMGOVRELAYTransactor // Write-only binding to the contract
	OPTIMISMGOVRELAYFilterer   // Log filterer for contract events
}

// OPTIMISMGOVRELAYCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISMGOVRELAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMGOVRELAYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISMGOVRELAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMGOVRELAYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISMGOVRELAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMGOVRELAYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISMGOVRELAYSession struct {
	Contract     *OPTIMISMGOVRELAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OPTIMISMGOVRELAYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISMGOVRELAYCallerSession struct {
	Contract *OPTIMISMGOVRELAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OPTIMISMGOVRELAYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISMGOVRELAYTransactorSession struct {
	Contract     *OPTIMISMGOVRELAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OPTIMISMGOVRELAYRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISMGOVRELAYRaw struct {
	Contract *OPTIMISMGOVRELAY // Generic contract binding to access the raw methods on
}

// OPTIMISMGOVRELAYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISMGOVRELAYCallerRaw struct {
	Contract *OPTIMISMGOVRELAYCaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISMGOVRELAYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISMGOVRELAYTransactorRaw struct {
	Contract *OPTIMISMGOVRELAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISMGOVRELAY creates a new instance of OPTIMISMGOVRELAY, bound to a specific deployed contract.
func NewOPTIMISMGOVRELAY(address common.Address, backend bind.ContractBackend) (*OPTIMISMGOVRELAY, error) {
	contract, err := bindOPTIMISMGOVRELAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAY{OPTIMISMGOVRELAYCaller: OPTIMISMGOVRELAYCaller{contract: contract}, OPTIMISMGOVRELAYTransactor: OPTIMISMGOVRELAYTransactor{contract: contract}, OPTIMISMGOVRELAYFilterer: OPTIMISMGOVRELAYFilterer{contract: contract}}, nil
}

// NewOPTIMISMGOVRELAYCaller creates a new read-only instance of OPTIMISMGOVRELAY, bound to a specific deployed contract.
func NewOPTIMISMGOVRELAYCaller(address common.Address, caller bind.ContractCaller) (*OPTIMISMGOVRELAYCaller, error) {
	contract, err := bindOPTIMISMGOVRELAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAYCaller{contract: contract}, nil
}

// NewOPTIMISMGOVRELAYTransactor creates a new write-only instance of OPTIMISMGOVRELAY, bound to a specific deployed contract.
func NewOPTIMISMGOVRELAYTransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISMGOVRELAYTransactor, error) {
	contract, err := bindOPTIMISMGOVRELAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAYTransactor{contract: contract}, nil
}

// NewOPTIMISMGOVRELAYFilterer creates a new log filterer instance of OPTIMISMGOVRELAY, bound to a specific deployed contract.
func NewOPTIMISMGOVRELAYFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISMGOVRELAYFilterer, error) {
	contract, err := bindOPTIMISMGOVRELAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAYFilterer{contract: contract}, nil
}

// bindOPTIMISMGOVRELAY binds a generic wrapper to an already deployed contract.
func bindOPTIMISMGOVRELAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISMGOVRELAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMGOVRELAY.Contract.OPTIMISMGOVRELAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.OPTIMISMGOVRELAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.OPTIMISMGOVRELAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMGOVRELAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.contract.Transact(opts, method, params...)
}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCaller) L2GovernanceRelay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMGOVRELAY.contract.Call(opts, &out, "l2GovernanceRelay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) L2GovernanceRelay() (common.Address, error) {
	return _OPTIMISMGOVRELAY.Contract.L2GovernanceRelay(&_OPTIMISMGOVRELAY.CallOpts)
}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCallerSession) L2GovernanceRelay() (common.Address, error) {
	return _OPTIMISMGOVRELAY.Contract.L2GovernanceRelay(&_OPTIMISMGOVRELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMGOVRELAY.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) Messenger() (common.Address, error) {
	return _OPTIMISMGOVRELAY.Contract.Messenger(&_OPTIMISMGOVRELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCallerSession) Messenger() (common.Address, error) {
	return _OPTIMISMGOVRELAY.Contract.Messenger(&_OPTIMISMGOVRELAY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISMGOVRELAY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMGOVRELAY.Contract.Wards(&_OPTIMISMGOVRELAY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMGOVRELAY.Contract.Wards(&_OPTIMISMGOVRELAY.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Deny(&_OPTIMISMGOVRELAY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Deny(&_OPTIMISMGOVRELAY.TransactOpts, usr)
}

// Relay is a paid mutator transaction binding the contract method 0x55444002.
//
// Solidity: function relay(address target, bytes targetData, uint32 l2gas) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactor) Relay(opts *bind.TransactOpts, target common.Address, targetData []byte, l2gas uint32) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.contract.Transact(opts, "relay", target, targetData, l2gas)
}

// Relay is a paid mutator transaction binding the contract method 0x55444002.
//
// Solidity: function relay(address target, bytes targetData, uint32 l2gas) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) Relay(target common.Address, targetData []byte, l2gas uint32) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Relay(&_OPTIMISMGOVRELAY.TransactOpts, target, targetData, l2gas)
}

// Relay is a paid mutator transaction binding the contract method 0x55444002.
//
// Solidity: function relay(address target, bytes targetData, uint32 l2gas) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactorSession) Relay(target common.Address, targetData []byte, l2gas uint32) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Relay(&_OPTIMISMGOVRELAY.TransactOpts, target, targetData, l2gas)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Rely(&_OPTIMISMGOVRELAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMGOVRELAY.Contract.Rely(&_OPTIMISMGOVRELAY.TransactOpts, usr)
}

// OPTIMISMGOVRELAYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISMGOVRELAY contract.
type OPTIMISMGOVRELAYDenyIterator struct {
	Event *OPTIMISMGOVRELAYDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISMGOVRELAYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMGOVRELAYDeny)
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
		it.Event = new(OPTIMISMGOVRELAYDeny)
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
func (it *OPTIMISMGOVRELAYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMGOVRELAYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMGOVRELAYDeny represents a Deny event raised by the OPTIMISMGOVRELAY contract.
type OPTIMISMGOVRELAYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMGOVRELAYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMGOVRELAY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAYDenyIterator{contract: _OPTIMISMGOVRELAY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISMGOVRELAYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMGOVRELAY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMGOVRELAYDeny)
				if err := _OPTIMISMGOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) ParseDeny(log types.Log) (*OPTIMISMGOVRELAYDeny, error) {
	event := new(OPTIMISMGOVRELAYDeny)
	if err := _OPTIMISMGOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMGOVRELAYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISMGOVRELAY contract.
type OPTIMISMGOVRELAYRelyIterator struct {
	Event *OPTIMISMGOVRELAYRely // Event containing the contract specifics and raw log

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
func (it *OPTIMISMGOVRELAYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMGOVRELAYRely)
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
		it.Event = new(OPTIMISMGOVRELAYRely)
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
func (it *OPTIMISMGOVRELAYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMGOVRELAYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMGOVRELAYRely represents a Rely event raised by the OPTIMISMGOVRELAY contract.
type OPTIMISMGOVRELAYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMGOVRELAYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMGOVRELAY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMGOVRELAYRelyIterator{contract: _OPTIMISMGOVRELAY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISMGOVRELAYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMGOVRELAY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMGOVRELAYRely)
				if err := _OPTIMISMGOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISMGOVRELAY *OPTIMISMGOVRELAYFilterer) ParseRely(log types.Log) (*OPTIMISMGOVRELAYRely, error) {
	event := new(OPTIMISMGOVRELAYRely)
	if err := _OPTIMISMGOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
