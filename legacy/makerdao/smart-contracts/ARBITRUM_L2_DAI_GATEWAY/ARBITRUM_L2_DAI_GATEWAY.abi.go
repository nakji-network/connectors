// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ARBITRUM_L2_DAI_GATEWAY

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

// ARBITRUML2DAIGATEWAYABI is the input ABI used to generate the binding from.
const ARBITRUML2DAIGATEWAYABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARBITRUML2DAIGATEWAY is an auto generated Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAY struct {
	ARBITRUML2DAIGATEWAYCaller     // Read-only binding to the contract
	ARBITRUML2DAIGATEWAYTransactor // Write-only binding to the contract
	ARBITRUML2DAIGATEWAYFilterer   // Log filterer for contract events
}

// ARBITRUML2DAIGATEWAYCaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAIGATEWAYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAIGATEWAYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUML2DAIGATEWAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAIGATEWAYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUML2DAIGATEWAYSession struct {
	Contract     *ARBITRUML2DAIGATEWAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ARBITRUML2DAIGATEWAYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUML2DAIGATEWAYCallerSession struct {
	Contract *ARBITRUML2DAIGATEWAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ARBITRUML2DAIGATEWAYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUML2DAIGATEWAYTransactorSession struct {
	Contract     *ARBITRUML2DAIGATEWAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ARBITRUML2DAIGATEWAYRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAYRaw struct {
	Contract *ARBITRUML2DAIGATEWAY // Generic contract binding to access the raw methods on
}

// ARBITRUML2DAIGATEWAYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAYCallerRaw struct {
	Contract *ARBITRUML2DAIGATEWAYCaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUML2DAIGATEWAYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUML2DAIGATEWAYTransactorRaw struct {
	Contract *ARBITRUML2DAIGATEWAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUML2DAIGATEWAY creates a new instance of ARBITRUML2DAIGATEWAY, bound to a specific deployed contract.
func NewARBITRUML2DAIGATEWAY(address common.Address, backend bind.ContractBackend) (*ARBITRUML2DAIGATEWAY, error) {
	contract, err := bindARBITRUML2DAIGATEWAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAY{ARBITRUML2DAIGATEWAYCaller: ARBITRUML2DAIGATEWAYCaller{contract: contract}, ARBITRUML2DAIGATEWAYTransactor: ARBITRUML2DAIGATEWAYTransactor{contract: contract}, ARBITRUML2DAIGATEWAYFilterer: ARBITRUML2DAIGATEWAYFilterer{contract: contract}}, nil
}

// NewARBITRUML2DAIGATEWAYCaller creates a new read-only instance of ARBITRUML2DAIGATEWAY, bound to a specific deployed contract.
func NewARBITRUML2DAIGATEWAYCaller(address common.Address, caller bind.ContractCaller) (*ARBITRUML2DAIGATEWAYCaller, error) {
	contract, err := bindARBITRUML2DAIGATEWAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYCaller{contract: contract}, nil
}

// NewARBITRUML2DAIGATEWAYTransactor creates a new write-only instance of ARBITRUML2DAIGATEWAY, bound to a specific deployed contract.
func NewARBITRUML2DAIGATEWAYTransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUML2DAIGATEWAYTransactor, error) {
	contract, err := bindARBITRUML2DAIGATEWAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYTransactor{contract: contract}, nil
}

// NewARBITRUML2DAIGATEWAYFilterer creates a new log filterer instance of ARBITRUML2DAIGATEWAY, bound to a specific deployed contract.
func NewARBITRUML2DAIGATEWAYFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUML2DAIGATEWAYFilterer, error) {
	contract, err := bindARBITRUML2DAIGATEWAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYFilterer{contract: contract}, nil
}

// bindARBITRUML2DAIGATEWAY binds a generic wrapper to an already deployed contract.
func bindARBITRUML2DAIGATEWAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUML2DAIGATEWAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2DAIGATEWAY.Contract.ARBITRUML2DAIGATEWAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.ARBITRUML2DAIGATEWAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.ARBITRUML2DAIGATEWAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2DAIGATEWAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.contract.Transact(opts, method, params...)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAIGATEWAY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Wards(&_ARBITRUML2DAIGATEWAY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Wards(&_ARBITRUML2DAIGATEWAY.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.contract.Transact(opts, "approve", token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Approve(&_ARBITRUML2DAIGATEWAY.TransactOpts, token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactorSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Approve(&_ARBITRUML2DAIGATEWAY.TransactOpts, token, spender, value)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Deny(&_ARBITRUML2DAIGATEWAY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Deny(&_ARBITRUML2DAIGATEWAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Rely(&_ARBITRUML2DAIGATEWAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAIGATEWAY.Contract.Rely(&_ARBITRUML2DAIGATEWAY.TransactOpts, usr)
}

// ARBITRUML2DAIGATEWAYApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYApproveIterator struct {
	Event *ARBITRUML2DAIGATEWAYApprove // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIGATEWAYApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIGATEWAYApprove)
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
		it.Event = new(ARBITRUML2DAIGATEWAYApprove)
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
func (it *ARBITRUML2DAIGATEWAYApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIGATEWAYApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIGATEWAYApprove represents a Approve event raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYApprove struct {
	Token   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) FilterApprove(opts *bind.FilterOpts, token []common.Address, spender []common.Address) (*ARBITRUML2DAIGATEWAYApproveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.FilterLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYApproveIterator{contract: _ARBITRUML2DAIGATEWAY.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIGATEWAYApprove, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.WatchLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIGATEWAYApprove)
				if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Approve", log); err != nil {
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
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) ParseApprove(log types.Log) (*ARBITRUML2DAIGATEWAYApprove, error) {
	event := new(ARBITRUML2DAIGATEWAYApprove)
	if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2DAIGATEWAYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYDenyIterator struct {
	Event *ARBITRUML2DAIGATEWAYDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIGATEWAYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIGATEWAYDeny)
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
		it.Event = new(ARBITRUML2DAIGATEWAYDeny)
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
func (it *ARBITRUML2DAIGATEWAYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIGATEWAYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIGATEWAYDeny represents a Deny event raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2DAIGATEWAYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYDenyIterator{contract: _ARBITRUML2DAIGATEWAY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIGATEWAYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIGATEWAYDeny)
				if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) ParseDeny(log types.Log) (*ARBITRUML2DAIGATEWAYDeny, error) {
	event := new(ARBITRUML2DAIGATEWAYDeny)
	if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2DAIGATEWAYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYRelyIterator struct {
	Event *ARBITRUML2DAIGATEWAYRely // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIGATEWAYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIGATEWAYRely)
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
		it.Event = new(ARBITRUML2DAIGATEWAYRely)
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
func (it *ARBITRUML2DAIGATEWAYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIGATEWAYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIGATEWAYRely represents a Rely event raised by the ARBITRUML2DAIGATEWAY contract.
type ARBITRUML2DAIGATEWAYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2DAIGATEWAYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIGATEWAYRelyIterator{contract: _ARBITRUML2DAIGATEWAY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIGATEWAYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAIGATEWAY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIGATEWAYRely)
				if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_ARBITRUML2DAIGATEWAY *ARBITRUML2DAIGATEWAYFilterer) ParseRely(log types.Log) (*ARBITRUML2DAIGATEWAYRely, error) {
	event := new(ARBITRUML2DAIGATEWAYRely)
	if err := _ARBITRUML2DAIGATEWAY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
