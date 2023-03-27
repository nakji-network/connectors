// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ArbitrumEscrow

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

// ARBITRUMESCROWABI is the input ABI used to generate the binding from.
const ARBITRUMESCROWABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARBITRUMESCROW is an auto generated Go binding around an Ethereum contract.
type ARBITRUMESCROW struct {
	ARBITRUMESCROWCaller     // Read-only binding to the contract
	ARBITRUMESCROWTransactor // Write-only binding to the contract
	ARBITRUMESCROWFilterer   // Log filterer for contract events
}

// ARBITRUMESCROWCaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUMESCROWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMESCROWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUMESCROWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMESCROWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUMESCROWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMESCROWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUMESCROWSession struct {
	Contract     *ARBITRUMESCROW   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARBITRUMESCROWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUMESCROWCallerSession struct {
	Contract *ARBITRUMESCROWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ARBITRUMESCROWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUMESCROWTransactorSession struct {
	Contract     *ARBITRUMESCROWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ARBITRUMESCROWRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUMESCROWRaw struct {
	Contract *ARBITRUMESCROW // Generic contract binding to access the raw methods on
}

// ARBITRUMESCROWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUMESCROWCallerRaw struct {
	Contract *ARBITRUMESCROWCaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUMESCROWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUMESCROWTransactorRaw struct {
	Contract *ARBITRUMESCROWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUMESCROW creates a new instance of ARBITRUMESCROW, bound to a specific deployed contract.
func NewARBITRUMESCROW(address common.Address, backend bind.ContractBackend) (*ARBITRUMESCROW, error) {
	contract, err := bindARBITRUMESCROW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROW{ARBITRUMESCROWCaller: ARBITRUMESCROWCaller{contract: contract}, ARBITRUMESCROWTransactor: ARBITRUMESCROWTransactor{contract: contract}, ARBITRUMESCROWFilterer: ARBITRUMESCROWFilterer{contract: contract}}, nil
}

// NewARBITRUMESCROWCaller creates a new read-only instance of ARBITRUMESCROW, bound to a specific deployed contract.
func NewARBITRUMESCROWCaller(address common.Address, caller bind.ContractCaller) (*ARBITRUMESCROWCaller, error) {
	contract, err := bindARBITRUMESCROW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWCaller{contract: contract}, nil
}

// NewARBITRUMESCROWTransactor creates a new write-only instance of ARBITRUMESCROW, bound to a specific deployed contract.
func NewARBITRUMESCROWTransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUMESCROWTransactor, error) {
	contract, err := bindARBITRUMESCROW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWTransactor{contract: contract}, nil
}

// NewARBITRUMESCROWFilterer creates a new log filterer instance of ARBITRUMESCROW, bound to a specific deployed contract.
func NewARBITRUMESCROWFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUMESCROWFilterer, error) {
	contract, err := bindARBITRUMESCROW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWFilterer{contract: contract}, nil
}

// bindARBITRUMESCROW binds a generic wrapper to an already deployed contract.
func bindARBITRUMESCROW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUMESCROWABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMESCROW *ARBITRUMESCROWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMESCROW.Contract.ARBITRUMESCROWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMESCROW *ARBITRUMESCROWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.ARBITRUMESCROWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMESCROW *ARBITRUMESCROWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.ARBITRUMESCROWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMESCROW *ARBITRUMESCROWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMESCROW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.contract.Transact(opts, method, params...)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMESCROW *ARBITRUMESCROWCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUMESCROW.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMESCROW *ARBITRUMESCROWSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMESCROW.Contract.Wards(&_ARBITRUMESCROW.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMESCROW *ARBITRUMESCROWCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMESCROW.Contract.Wards(&_ARBITRUMESCROW.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUMESCROW.contract.Transact(opts, "approve", token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Approve(&_ARBITRUMESCROW.TransactOpts, token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactorSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Approve(&_ARBITRUMESCROW.TransactOpts, token, spender, value)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Deny(&_ARBITRUMESCROW.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Deny(&_ARBITRUMESCROW.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Rely(&_ARBITRUMESCROW.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMESCROW *ARBITRUMESCROWTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMESCROW.Contract.Rely(&_ARBITRUMESCROW.TransactOpts, usr)
}

// ARBITRUMESCROWApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWApproveIterator struct {
	Event *ARBITRUMESCROWApprove // Event containing the contract specifics and raw log

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
func (it *ARBITRUMESCROWApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMESCROWApprove)
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
		it.Event = new(ARBITRUMESCROWApprove)
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
func (it *ARBITRUMESCROWApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMESCROWApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMESCROWApprove represents a Approve event raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWApprove struct {
	Token   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) FilterApprove(opts *bind.FilterOpts, token []common.Address, spender []common.Address) (*ARBITRUMESCROWApproveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.FilterLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWApproveIterator{contract: _ARBITRUMESCROW.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *ARBITRUMESCROWApprove, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.WatchLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMESCROWApprove)
				if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Approve", log); err != nil {
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

// ParseApprove is a log parse operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) ParseApprove(log types.Log) (*ARBITRUMESCROWApprove, error) {
	event := new(ARBITRUMESCROWApprove)
	if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMESCROWDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWDenyIterator struct {
	Event *ARBITRUMESCROWDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUMESCROWDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMESCROWDeny)
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
		it.Event = new(ARBITRUMESCROWDeny)
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
func (it *ARBITRUMESCROWDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMESCROWDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMESCROWDeny represents a Deny event raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMESCROWDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWDenyIterator{contract: _ARBITRUMESCROW.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUMESCROWDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMESCROWDeny)
				if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) ParseDeny(log types.Log) (*ARBITRUMESCROWDeny, error) {
	event := new(ARBITRUMESCROWDeny)
	if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMESCROWRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWRelyIterator struct {
	Event *ARBITRUMESCROWRely // Event containing the contract specifics and raw log

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
func (it *ARBITRUMESCROWRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMESCROWRely)
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
		it.Event = new(ARBITRUMESCROWRely)
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
func (it *ARBITRUMESCROWRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMESCROWRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMESCROWRely represents a Rely event raised by the ARBITRUMESCROW contract.
type ARBITRUMESCROWRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMESCROWRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMESCROWRelyIterator{contract: _ARBITRUMESCROW.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUMESCROWRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMESCROW.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMESCROWRely)
				if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_ARBITRUMESCROW *ARBITRUMESCROWFilterer) ParseRely(log types.Log) (*ARBITRUMESCROWRely, error) {
	event := new(ARBITRUMESCROWRely)
	if err := _ARBITRUMESCROW.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
