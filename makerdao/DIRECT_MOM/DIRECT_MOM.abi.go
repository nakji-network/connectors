// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DIRECT_MOM

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

// DIRECTMOMABI is the input ABI used to generate the binding from.
const DIRECTMOMABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"Disable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAuthority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"SetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DIRECTMOM is an auto generated Go binding around an Ethereum contract.
type DIRECTMOM struct {
	DIRECTMOMCaller     // Read-only binding to the contract
	DIRECTMOMTransactor // Write-only binding to the contract
	DIRECTMOMFilterer   // Log filterer for contract events
}

// DIRECTMOMCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIRECTMOMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIRECTMOMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIRECTMOMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIRECTMOMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIRECTMOMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIRECTMOMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIRECTMOMSession struct {
	Contract     *DIRECTMOM        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIRECTMOMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIRECTMOMCallerSession struct {
	Contract *DIRECTMOMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DIRECTMOMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIRECTMOMTransactorSession struct {
	Contract     *DIRECTMOMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DIRECTMOMRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIRECTMOMRaw struct {
	Contract *DIRECTMOM // Generic contract binding to access the raw methods on
}

// DIRECTMOMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIRECTMOMCallerRaw struct {
	Contract *DIRECTMOMCaller // Generic read-only contract binding to access the raw methods on
}

// DIRECTMOMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIRECTMOMTransactorRaw struct {
	Contract *DIRECTMOMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIRECTMOM creates a new instance of DIRECTMOM, bound to a specific deployed contract.
func NewDIRECTMOM(address common.Address, backend bind.ContractBackend) (*DIRECTMOM, error) {
	contract, err := bindDIRECTMOM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOM{DIRECTMOMCaller: DIRECTMOMCaller{contract: contract}, DIRECTMOMTransactor: DIRECTMOMTransactor{contract: contract}, DIRECTMOMFilterer: DIRECTMOMFilterer{contract: contract}}, nil
}

// NewDIRECTMOMCaller creates a new read-only instance of DIRECTMOM, bound to a specific deployed contract.
func NewDIRECTMOMCaller(address common.Address, caller bind.ContractCaller) (*DIRECTMOMCaller, error) {
	contract, err := bindDIRECTMOM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMCaller{contract: contract}, nil
}

// NewDIRECTMOMTransactor creates a new write-only instance of DIRECTMOM, bound to a specific deployed contract.
func NewDIRECTMOMTransactor(address common.Address, transactor bind.ContractTransactor) (*DIRECTMOMTransactor, error) {
	contract, err := bindDIRECTMOM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMTransactor{contract: contract}, nil
}

// NewDIRECTMOMFilterer creates a new log filterer instance of DIRECTMOM, bound to a specific deployed contract.
func NewDIRECTMOMFilterer(address common.Address, filterer bind.ContractFilterer) (*DIRECTMOMFilterer, error) {
	contract, err := bindDIRECTMOM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMFilterer{contract: contract}, nil
}

// bindDIRECTMOM binds a generic wrapper to an already deployed contract.
func bindDIRECTMOM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIRECTMOMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIRECTMOM *DIRECTMOMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIRECTMOM.Contract.DIRECTMOMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIRECTMOM *DIRECTMOMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.DIRECTMOMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIRECTMOM *DIRECTMOMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.DIRECTMOMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIRECTMOM *DIRECTMOMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIRECTMOM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIRECTMOM *DIRECTMOMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIRECTMOM *DIRECTMOMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_DIRECTMOM *DIRECTMOMCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DIRECTMOM.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_DIRECTMOM *DIRECTMOMSession) Authority() (common.Address, error) {
	return _DIRECTMOM.Contract.Authority(&_DIRECTMOM.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_DIRECTMOM *DIRECTMOMCallerSession) Authority() (common.Address, error) {
	return _DIRECTMOM.Contract.Authority(&_DIRECTMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIRECTMOM *DIRECTMOMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DIRECTMOM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIRECTMOM *DIRECTMOMSession) Owner() (common.Address, error) {
	return _DIRECTMOM.Contract.Owner(&_DIRECTMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIRECTMOM *DIRECTMOMCallerSession) Owner() (common.Address, error) {
	return _DIRECTMOM.Contract.Owner(&_DIRECTMOM.CallOpts)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address who) returns()
func (_DIRECTMOM *DIRECTMOMTransactor) Disable(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.contract.Transact(opts, "disable", who)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address who) returns()
func (_DIRECTMOM *DIRECTMOMSession) Disable(who common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.Disable(&_DIRECTMOM.TransactOpts, who)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address who) returns()
func (_DIRECTMOM *DIRECTMOMTransactorSession) Disable(who common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.Disable(&_DIRECTMOM.TransactOpts, who)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DIRECTMOM *DIRECTMOMTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DIRECTMOM *DIRECTMOMSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.SetAuthority(&_DIRECTMOM.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DIRECTMOM *DIRECTMOMTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.SetAuthority(&_DIRECTMOM.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DIRECTMOM *DIRECTMOMTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DIRECTMOM *DIRECTMOMSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.SetOwner(&_DIRECTMOM.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DIRECTMOM *DIRECTMOMTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DIRECTMOM.Contract.SetOwner(&_DIRECTMOM.TransactOpts, owner_)
}

// DIRECTMOMDisableIterator is returned from FilterDisable and is used to iterate over the raw logs and unpacked data for Disable events raised by the DIRECTMOM contract.
type DIRECTMOMDisableIterator struct {
	Event *DIRECTMOMDisable // Event containing the contract specifics and raw log

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
func (it *DIRECTMOMDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIRECTMOMDisable)
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
		it.Event = new(DIRECTMOMDisable)
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
func (it *DIRECTMOMDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIRECTMOMDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIRECTMOMDisable represents a Disable event raised by the DIRECTMOM contract.
type DIRECTMOMDisable struct {
	Who common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisable is a free log retrieval operation binding the contract event 0x1d6f80aff2255e9cb1aee4abaf819aec89a6ae3fce7eb53b852640eb8052609c.
//
// Solidity: event Disable(address indexed who)
func (_DIRECTMOM *DIRECTMOMFilterer) FilterDisable(opts *bind.FilterOpts, who []common.Address) (*DIRECTMOMDisableIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DIRECTMOM.contract.FilterLogs(opts, "Disable", whoRule)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMDisableIterator{contract: _DIRECTMOM.contract, event: "Disable", logs: logs, sub: sub}, nil
}

// WatchDisable is a free log subscription operation binding the contract event 0x1d6f80aff2255e9cb1aee4abaf819aec89a6ae3fce7eb53b852640eb8052609c.
//
// Solidity: event Disable(address indexed who)
func (_DIRECTMOM *DIRECTMOMFilterer) WatchDisable(opts *bind.WatchOpts, sink chan<- *DIRECTMOMDisable, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DIRECTMOM.contract.WatchLogs(opts, "Disable", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIRECTMOMDisable)
				if err := _DIRECTMOM.contract.UnpackLog(event, "Disable", log); err != nil {
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

// ParseDisable is a log parse operation binding the contract event 0x1d6f80aff2255e9cb1aee4abaf819aec89a6ae3fce7eb53b852640eb8052609c.
//
// Solidity: event Disable(address indexed who)
func (_DIRECTMOM *DIRECTMOMFilterer) ParseDisable(log types.Log) (*DIRECTMOMDisable, error) {
	event := new(DIRECTMOMDisable)
	if err := _DIRECTMOM.contract.UnpackLog(event, "Disable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIRECTMOMSetAuthorityIterator is returned from FilterSetAuthority and is used to iterate over the raw logs and unpacked data for SetAuthority events raised by the DIRECTMOM contract.
type DIRECTMOMSetAuthorityIterator struct {
	Event *DIRECTMOMSetAuthority // Event containing the contract specifics and raw log

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
func (it *DIRECTMOMSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIRECTMOMSetAuthority)
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
		it.Event = new(DIRECTMOMSetAuthority)
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
func (it *DIRECTMOMSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIRECTMOMSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIRECTMOMSetAuthority represents a SetAuthority event raised by the DIRECTMOM contract.
type DIRECTMOMSetAuthority struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAuthority is a free log retrieval operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address indexed oldAuthority, address indexed newAuthority)
func (_DIRECTMOM *DIRECTMOMFilterer) FilterSetAuthority(opts *bind.FilterOpts, oldAuthority []common.Address, newAuthority []common.Address) (*DIRECTMOMSetAuthorityIterator, error) {

	var oldAuthorityRule []interface{}
	for _, oldAuthorityItem := range oldAuthority {
		oldAuthorityRule = append(oldAuthorityRule, oldAuthorityItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _DIRECTMOM.contract.FilterLogs(opts, "SetAuthority", oldAuthorityRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMSetAuthorityIterator{contract: _DIRECTMOM.contract, event: "SetAuthority", logs: logs, sub: sub}, nil
}

// WatchSetAuthority is a free log subscription operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address indexed oldAuthority, address indexed newAuthority)
func (_DIRECTMOM *DIRECTMOMFilterer) WatchSetAuthority(opts *bind.WatchOpts, sink chan<- *DIRECTMOMSetAuthority, oldAuthority []common.Address, newAuthority []common.Address) (event.Subscription, error) {

	var oldAuthorityRule []interface{}
	for _, oldAuthorityItem := range oldAuthority {
		oldAuthorityRule = append(oldAuthorityRule, oldAuthorityItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _DIRECTMOM.contract.WatchLogs(opts, "SetAuthority", oldAuthorityRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIRECTMOMSetAuthority)
				if err := _DIRECTMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
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

// ParseSetAuthority is a log parse operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address indexed oldAuthority, address indexed newAuthority)
func (_DIRECTMOM *DIRECTMOMFilterer) ParseSetAuthority(log types.Log) (*DIRECTMOMSetAuthority, error) {
	event := new(DIRECTMOMSetAuthority)
	if err := _DIRECTMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIRECTMOMSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the DIRECTMOM contract.
type DIRECTMOMSetOwnerIterator struct {
	Event *DIRECTMOMSetOwner // Event containing the contract specifics and raw log

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
func (it *DIRECTMOMSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIRECTMOMSetOwner)
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
		it.Event = new(DIRECTMOMSetOwner)
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
func (it *DIRECTMOMSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIRECTMOMSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIRECTMOMSetOwner represents a SetOwner event raised by the DIRECTMOM contract.
type DIRECTMOMSetOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address indexed oldOwner, address indexed newOwner)
func (_DIRECTMOM *DIRECTMOMFilterer) FilterSetOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*DIRECTMOMSetOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIRECTMOM.contract.FilterLogs(opts, "SetOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DIRECTMOMSetOwnerIterator{contract: _DIRECTMOM.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address indexed oldOwner, address indexed newOwner)
func (_DIRECTMOM *DIRECTMOMFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *DIRECTMOMSetOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIRECTMOM.contract.WatchLogs(opts, "SetOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIRECTMOMSetOwner)
				if err := _DIRECTMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
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

// ParseSetOwner is a log parse operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address indexed oldOwner, address indexed newOwner)
func (_DIRECTMOM *DIRECTMOMFilterer) ParseSetOwner(log types.Log) (*DIRECTMOMSetOwner, error) {
	event := new(DIRECTMOMSetOwner)
	if err := _DIRECTMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
