// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OptimismEscrow

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

// OPTIMISMESCROWABI is the input ABI used to generate the binding from.
const OPTIMISMESCROWABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISMESCROW is an auto generated Go binding around an Ethereum contract.
type OPTIMISMESCROW struct {
	OPTIMISMESCROWCaller     // Read-only binding to the contract
	OPTIMISMESCROWTransactor // Write-only binding to the contract
	OPTIMISMESCROWFilterer   // Log filterer for contract events
}

// OPTIMISMESCROWCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISMESCROWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMESCROWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISMESCROWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMESCROWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISMESCROWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMESCROWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISMESCROWSession struct {
	Contract     *OPTIMISMESCROW   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OPTIMISMESCROWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISMESCROWCallerSession struct {
	Contract *OPTIMISMESCROWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OPTIMISMESCROWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISMESCROWTransactorSession struct {
	Contract     *OPTIMISMESCROWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OPTIMISMESCROWRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISMESCROWRaw struct {
	Contract *OPTIMISMESCROW // Generic contract binding to access the raw methods on
}

// OPTIMISMESCROWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISMESCROWCallerRaw struct {
	Contract *OPTIMISMESCROWCaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISMESCROWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISMESCROWTransactorRaw struct {
	Contract *OPTIMISMESCROWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISMESCROW creates a new instance of OPTIMISMESCROW, bound to a specific deployed contract.
func NewOPTIMISMESCROW(address common.Address, backend bind.ContractBackend) (*OPTIMISMESCROW, error) {
	contract, err := bindOPTIMISMESCROW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROW{OPTIMISMESCROWCaller: OPTIMISMESCROWCaller{contract: contract}, OPTIMISMESCROWTransactor: OPTIMISMESCROWTransactor{contract: contract}, OPTIMISMESCROWFilterer: OPTIMISMESCROWFilterer{contract: contract}}, nil
}

// NewOPTIMISMESCROWCaller creates a new read-only instance of OPTIMISMESCROW, bound to a specific deployed contract.
func NewOPTIMISMESCROWCaller(address common.Address, caller bind.ContractCaller) (*OPTIMISMESCROWCaller, error) {
	contract, err := bindOPTIMISMESCROW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWCaller{contract: contract}, nil
}

// NewOPTIMISMESCROWTransactor creates a new write-only instance of OPTIMISMESCROW, bound to a specific deployed contract.
func NewOPTIMISMESCROWTransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISMESCROWTransactor, error) {
	contract, err := bindOPTIMISMESCROW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWTransactor{contract: contract}, nil
}

// NewOPTIMISMESCROWFilterer creates a new log filterer instance of OPTIMISMESCROW, bound to a specific deployed contract.
func NewOPTIMISMESCROWFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISMESCROWFilterer, error) {
	contract, err := bindOPTIMISMESCROW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWFilterer{contract: contract}, nil
}

// bindOPTIMISMESCROW binds a generic wrapper to an already deployed contract.
func bindOPTIMISMESCROW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISMESCROWABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMESCROW *OPTIMISMESCROWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMESCROW.Contract.OPTIMISMESCROWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMESCROW *OPTIMISMESCROWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.OPTIMISMESCROWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMESCROW *OPTIMISMESCROWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.OPTIMISMESCROWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMESCROW *OPTIMISMESCROWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMESCROW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.contract.Transact(opts, method, params...)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMESCROW *OPTIMISMESCROWCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISMESCROW.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMESCROW *OPTIMISMESCROWSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMESCROW.Contract.Wards(&_OPTIMISMESCROW.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMESCROW *OPTIMISMESCROWCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMESCROW.Contract.Wards(&_OPTIMISMESCROW.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISMESCROW.contract.Transact(opts, "approve", token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Approve(&_OPTIMISMESCROW.TransactOpts, token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactorSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Approve(&_OPTIMISMESCROW.TransactOpts, token, spender, value)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Deny(&_OPTIMISMESCROW.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Deny(&_OPTIMISMESCROW.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Rely(&_OPTIMISMESCROW.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMESCROW *OPTIMISMESCROWTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMESCROW.Contract.Rely(&_OPTIMISMESCROW.TransactOpts, usr)
}

// OPTIMISMESCROWApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWApproveIterator struct {
	Event *OPTIMISMESCROWApprove // Event containing the contract specifics and raw log

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
func (it *OPTIMISMESCROWApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMESCROWApprove)
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
		it.Event = new(OPTIMISMESCROWApprove)
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
func (it *OPTIMISMESCROWApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMESCROWApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMESCROWApprove represents a Approve event raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWApprove struct {
	Token   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) FilterApprove(opts *bind.FilterOpts, token []common.Address, spender []common.Address) (*OPTIMISMESCROWApproveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.FilterLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWApproveIterator{contract: _OPTIMISMESCROW.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *OPTIMISMESCROWApprove, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.WatchLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMESCROWApprove)
				if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Approve", log); err != nil {
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
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) ParseApprove(log types.Log) (*OPTIMISMESCROWApprove, error) {
	event := new(OPTIMISMESCROWApprove)
	if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMESCROWDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWDenyIterator struct {
	Event *OPTIMISMESCROWDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISMESCROWDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMESCROWDeny)
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
		it.Event = new(OPTIMISMESCROWDeny)
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
func (it *OPTIMISMESCROWDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMESCROWDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMESCROWDeny represents a Deny event raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMESCROWDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWDenyIterator{contract: _OPTIMISMESCROW.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISMESCROWDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMESCROWDeny)
				if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) ParseDeny(log types.Log) (*OPTIMISMESCROWDeny, error) {
	event := new(OPTIMISMESCROWDeny)
	if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMESCROWRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWRelyIterator struct {
	Event *OPTIMISMESCROWRely // Event containing the contract specifics and raw log

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
func (it *OPTIMISMESCROWRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMESCROWRely)
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
		it.Event = new(OPTIMISMESCROWRely)
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
func (it *OPTIMISMESCROWRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMESCROWRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMESCROWRely represents a Rely event raised by the OPTIMISMESCROW contract.
type OPTIMISMESCROWRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMESCROWRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMESCROWRelyIterator{contract: _OPTIMISMESCROW.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISMESCROWRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMESCROW.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMESCROWRely)
				if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISMESCROW *OPTIMISMESCROWFilterer) ParseRely(log types.Log) (*OPTIMISMESCROWRely, error) {
	event := new(OPTIMISMESCROWRely)
	if err := _OPTIMISMESCROW.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
