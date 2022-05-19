// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GOV_GUARD

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

// GOVGUARDABI is the input ABI used to generate the binding from.
const GOVGUARDABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"LogDeny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"LogRely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRoot\",\"type\":\"address\"}],\"name\":\"LogSetRoot\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GOVGUARD is an auto generated Go binding around an Ethereum contract.
type GOVGUARD struct {
	GOVGUARDCaller     // Read-only binding to the contract
	GOVGUARDTransactor // Write-only binding to the contract
	GOVGUARDFilterer   // Log filterer for contract events
}

// GOVGUARDCaller is an auto generated read-only Go binding around an Ethereum contract.
type GOVGUARDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOVGUARDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GOVGUARDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOVGUARDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GOVGUARDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOVGUARDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GOVGUARDSession struct {
	Contract     *GOVGUARD         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GOVGUARDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GOVGUARDCallerSession struct {
	Contract *GOVGUARDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GOVGUARDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GOVGUARDTransactorSession struct {
	Contract     *GOVGUARDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GOVGUARDRaw is an auto generated low-level Go binding around an Ethereum contract.
type GOVGUARDRaw struct {
	Contract *GOVGUARD // Generic contract binding to access the raw methods on
}

// GOVGUARDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GOVGUARDCallerRaw struct {
	Contract *GOVGUARDCaller // Generic read-only contract binding to access the raw methods on
}

// GOVGUARDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GOVGUARDTransactorRaw struct {
	Contract *GOVGUARDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGOVGUARD creates a new instance of GOVGUARD, bound to a specific deployed contract.
func NewGOVGUARD(address common.Address, backend bind.ContractBackend) (*GOVGUARD, error) {
	contract, err := bindGOVGUARD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GOVGUARD{GOVGUARDCaller: GOVGUARDCaller{contract: contract}, GOVGUARDTransactor: GOVGUARDTransactor{contract: contract}, GOVGUARDFilterer: GOVGUARDFilterer{contract: contract}}, nil
}

// NewGOVGUARDCaller creates a new read-only instance of GOVGUARD, bound to a specific deployed contract.
func NewGOVGUARDCaller(address common.Address, caller bind.ContractCaller) (*GOVGUARDCaller, error) {
	contract, err := bindGOVGUARD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDCaller{contract: contract}, nil
}

// NewGOVGUARDTransactor creates a new write-only instance of GOVGUARD, bound to a specific deployed contract.
func NewGOVGUARDTransactor(address common.Address, transactor bind.ContractTransactor) (*GOVGUARDTransactor, error) {
	contract, err := bindGOVGUARD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDTransactor{contract: contract}, nil
}

// NewGOVGUARDFilterer creates a new log filterer instance of GOVGUARD, bound to a specific deployed contract.
func NewGOVGUARDFilterer(address common.Address, filterer bind.ContractFilterer) (*GOVGUARDFilterer, error) {
	contract, err := bindGOVGUARD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDFilterer{contract: contract}, nil
}

// bindGOVGUARD binds a generic wrapper to an already deployed contract.
func bindGOVGUARD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GOVGUARDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GOVGUARD *GOVGUARDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GOVGUARD.Contract.GOVGUARDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GOVGUARD *GOVGUARDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GOVGUARD.Contract.GOVGUARDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GOVGUARD *GOVGUARDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GOVGUARD.Contract.GOVGUARDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GOVGUARD *GOVGUARDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GOVGUARD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GOVGUARD *GOVGUARDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GOVGUARD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GOVGUARD *GOVGUARDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GOVGUARD.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src, address , bytes4 sig) view returns(bool)
func (_GOVGUARD *GOVGUARDCaller) CanCall(opts *bind.CallOpts, src common.Address, arg1 common.Address, sig [4]byte) (bool, error) {
	var out []interface{}
	err := _GOVGUARD.contract.Call(opts, &out, "canCall", src, arg1, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src, address , bytes4 sig) view returns(bool)
func (_GOVGUARD *GOVGUARDSession) CanCall(src common.Address, arg1 common.Address, sig [4]byte) (bool, error) {
	return _GOVGUARD.Contract.CanCall(&_GOVGUARD.CallOpts, src, arg1, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src, address , bytes4 sig) view returns(bool)
func (_GOVGUARD *GOVGUARDCallerSession) CanCall(src common.Address, arg1 common.Address, sig [4]byte) (bool, error) {
	return _GOVGUARD.Contract.CanCall(&_GOVGUARD.CallOpts, src, arg1, sig)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_GOVGUARD *GOVGUARDCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GOVGUARD.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_GOVGUARD *GOVGUARDSession) Root() (common.Address, error) {
	return _GOVGUARD.Contract.Root(&_GOVGUARD.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_GOVGUARD *GOVGUARDCallerSession) Root() (common.Address, error) {
	return _GOVGUARD.Contract.Root(&_GOVGUARD.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GOVGUARD *GOVGUARDCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GOVGUARD.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GOVGUARD *GOVGUARDSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _GOVGUARD.Contract.Wards(&_GOVGUARD.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GOVGUARD *GOVGUARDCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _GOVGUARD.Contract.Wards(&_GOVGUARD.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GOVGUARD *GOVGUARDSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.Deny(&_GOVGUARD.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.Deny(&_GOVGUARD.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GOVGUARD *GOVGUARDSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.Rely(&_GOVGUARD.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.Rely(&_GOVGUARD.TransactOpts, usr)
}

// SetRoot is a paid mutator transaction binding the contract method 0x003ba1ed.
//
// Solidity: function setRoot(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactor) SetRoot(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.contract.Transact(opts, "setRoot", usr)
}

// SetRoot is a paid mutator transaction binding the contract method 0x003ba1ed.
//
// Solidity: function setRoot(address usr) returns()
func (_GOVGUARD *GOVGUARDSession) SetRoot(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.SetRoot(&_GOVGUARD.TransactOpts, usr)
}

// SetRoot is a paid mutator transaction binding the contract method 0x003ba1ed.
//
// Solidity: function setRoot(address usr) returns()
func (_GOVGUARD *GOVGUARDTransactorSession) SetRoot(usr common.Address) (*types.Transaction, error) {
	return _GOVGUARD.Contract.SetRoot(&_GOVGUARD.TransactOpts, usr)
}

// GOVGUARDLogDenyIterator is returned from FilterLogDeny and is used to iterate over the raw logs and unpacked data for LogDeny events raised by the GOVGUARD contract.
type GOVGUARDLogDenyIterator struct {
	Event *GOVGUARDLogDeny // Event containing the contract specifics and raw log

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
func (it *GOVGUARDLogDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GOVGUARDLogDeny)
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
		it.Event = new(GOVGUARDLogDeny)
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
func (it *GOVGUARDLogDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GOVGUARDLogDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GOVGUARDLogDeny represents a LogDeny event raised by the GOVGUARD contract.
type GOVGUARDLogDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogDeny is a free log retrieval operation binding the contract event 0x2569dba00999b9a749665fecc99518b00d82fa6179a90d094a04f4c50b952ba1.
//
// Solidity: event LogDeny(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) FilterLogDeny(opts *bind.FilterOpts, usr []common.Address) (*GOVGUARDLogDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _GOVGUARD.contract.FilterLogs(opts, "LogDeny", usrRule)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDLogDenyIterator{contract: _GOVGUARD.contract, event: "LogDeny", logs: logs, sub: sub}, nil
}

// WatchLogDeny is a free log subscription operation binding the contract event 0x2569dba00999b9a749665fecc99518b00d82fa6179a90d094a04f4c50b952ba1.
//
// Solidity: event LogDeny(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) WatchLogDeny(opts *bind.WatchOpts, sink chan<- *GOVGUARDLogDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _GOVGUARD.contract.WatchLogs(opts, "LogDeny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GOVGUARDLogDeny)
				if err := _GOVGUARD.contract.UnpackLog(event, "LogDeny", log); err != nil {
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

// ParseLogDeny is a log parse operation binding the contract event 0x2569dba00999b9a749665fecc99518b00d82fa6179a90d094a04f4c50b952ba1.
//
// Solidity: event LogDeny(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) ParseLogDeny(log types.Log) (*GOVGUARDLogDeny, error) {
	event := new(GOVGUARDLogDeny)
	if err := _GOVGUARD.contract.UnpackLog(event, "LogDeny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GOVGUARDLogRelyIterator is returned from FilterLogRely and is used to iterate over the raw logs and unpacked data for LogRely events raised by the GOVGUARD contract.
type GOVGUARDLogRelyIterator struct {
	Event *GOVGUARDLogRely // Event containing the contract specifics and raw log

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
func (it *GOVGUARDLogRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GOVGUARDLogRely)
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
		it.Event = new(GOVGUARDLogRely)
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
func (it *GOVGUARDLogRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GOVGUARDLogRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GOVGUARDLogRely represents a LogRely event raised by the GOVGUARD contract.
type GOVGUARDLogRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogRely is a free log retrieval operation binding the contract event 0xbcea0e0a8da02ca32f79ec827797ef9d9b799f0f530a7015479294f8605f26dc.
//
// Solidity: event LogRely(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) FilterLogRely(opts *bind.FilterOpts, usr []common.Address) (*GOVGUARDLogRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _GOVGUARD.contract.FilterLogs(opts, "LogRely", usrRule)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDLogRelyIterator{contract: _GOVGUARD.contract, event: "LogRely", logs: logs, sub: sub}, nil
}

// WatchLogRely is a free log subscription operation binding the contract event 0xbcea0e0a8da02ca32f79ec827797ef9d9b799f0f530a7015479294f8605f26dc.
//
// Solidity: event LogRely(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) WatchLogRely(opts *bind.WatchOpts, sink chan<- *GOVGUARDLogRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _GOVGUARD.contract.WatchLogs(opts, "LogRely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GOVGUARDLogRely)
				if err := _GOVGUARD.contract.UnpackLog(event, "LogRely", log); err != nil {
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

// ParseLogRely is a log parse operation binding the contract event 0xbcea0e0a8da02ca32f79ec827797ef9d9b799f0f530a7015479294f8605f26dc.
//
// Solidity: event LogRely(address indexed usr)
func (_GOVGUARD *GOVGUARDFilterer) ParseLogRely(log types.Log) (*GOVGUARDLogRely, error) {
	event := new(GOVGUARDLogRely)
	if err := _GOVGUARD.contract.UnpackLog(event, "LogRely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GOVGUARDLogSetRootIterator is returned from FilterLogSetRoot and is used to iterate over the raw logs and unpacked data for LogSetRoot events raised by the GOVGUARD contract.
type GOVGUARDLogSetRootIterator struct {
	Event *GOVGUARDLogSetRoot // Event containing the contract specifics and raw log

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
func (it *GOVGUARDLogSetRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GOVGUARDLogSetRoot)
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
		it.Event = new(GOVGUARDLogSetRoot)
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
func (it *GOVGUARDLogSetRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GOVGUARDLogSetRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GOVGUARDLogSetRoot represents a LogSetRoot event raised by the GOVGUARD contract.
type GOVGUARDLogSetRoot struct {
	NewRoot common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogSetRoot is a free log retrieval operation binding the contract event 0x590b975c6c0ad07291052d35e354944ebed986091d40026a932b3b6d122e753c.
//
// Solidity: event LogSetRoot(address indexed newRoot)
func (_GOVGUARD *GOVGUARDFilterer) FilterLogSetRoot(opts *bind.FilterOpts, newRoot []common.Address) (*GOVGUARDLogSetRootIterator, error) {

	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _GOVGUARD.contract.FilterLogs(opts, "LogSetRoot", newRootRule)
	if err != nil {
		return nil, err
	}
	return &GOVGUARDLogSetRootIterator{contract: _GOVGUARD.contract, event: "LogSetRoot", logs: logs, sub: sub}, nil
}

// WatchLogSetRoot is a free log subscription operation binding the contract event 0x590b975c6c0ad07291052d35e354944ebed986091d40026a932b3b6d122e753c.
//
// Solidity: event LogSetRoot(address indexed newRoot)
func (_GOVGUARD *GOVGUARDFilterer) WatchLogSetRoot(opts *bind.WatchOpts, sink chan<- *GOVGUARDLogSetRoot, newRoot []common.Address) (event.Subscription, error) {

	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _GOVGUARD.contract.WatchLogs(opts, "LogSetRoot", newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GOVGUARDLogSetRoot)
				if err := _GOVGUARD.contract.UnpackLog(event, "LogSetRoot", log); err != nil {
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

// ParseLogSetRoot is a log parse operation binding the contract event 0x590b975c6c0ad07291052d35e354944ebed986091d40026a932b3b6d122e753c.
//
// Solidity: event LogSetRoot(address indexed newRoot)
func (_GOVGUARD *GOVGUARDFilterer) ParseLogSetRoot(log types.Log) (*GOVGUARDLogSetRoot, error) {
	event := new(GOVGUARDLogSetRoot)
	if err := _GOVGUARD.contract.UnpackLog(event, "LogSetRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
