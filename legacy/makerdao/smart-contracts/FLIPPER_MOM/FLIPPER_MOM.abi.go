// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FLIPPER_MOM

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

// FLIPPERMOMABI is the input ABI used to generate the binding from.
const FLIPPERMOMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cat\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAuthority\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"SetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FLIPPERMOM is an auto generated Go binding around an Ethereum contract.
type FLIPPERMOM struct {
	FLIPPERMOMCaller     // Read-only binding to the contract
	FLIPPERMOMTransactor // Write-only binding to the contract
	FLIPPERMOMFilterer   // Log filterer for contract events
}

// FLIPPERMOMCaller is an auto generated read-only Go binding around an Ethereum contract.
type FLIPPERMOMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPPERMOMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FLIPPERMOMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPPERMOMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FLIPPERMOMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPPERMOMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FLIPPERMOMSession struct {
	Contract     *FLIPPERMOM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FLIPPERMOMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FLIPPERMOMCallerSession struct {
	Contract *FLIPPERMOMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FLIPPERMOMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FLIPPERMOMTransactorSession struct {
	Contract     *FLIPPERMOMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FLIPPERMOMRaw is an auto generated low-level Go binding around an Ethereum contract.
type FLIPPERMOMRaw struct {
	Contract *FLIPPERMOM // Generic contract binding to access the raw methods on
}

// FLIPPERMOMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FLIPPERMOMCallerRaw struct {
	Contract *FLIPPERMOMCaller // Generic read-only contract binding to access the raw methods on
}

// FLIPPERMOMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FLIPPERMOMTransactorRaw struct {
	Contract *FLIPPERMOMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFLIPPERMOM creates a new instance of FLIPPERMOM, bound to a specific deployed contract.
func NewFLIPPERMOM(address common.Address, backend bind.ContractBackend) (*FLIPPERMOM, error) {
	contract, err := bindFLIPPERMOM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOM{FLIPPERMOMCaller: FLIPPERMOMCaller{contract: contract}, FLIPPERMOMTransactor: FLIPPERMOMTransactor{contract: contract}, FLIPPERMOMFilterer: FLIPPERMOMFilterer{contract: contract}}, nil
}

// NewFLIPPERMOMCaller creates a new read-only instance of FLIPPERMOM, bound to a specific deployed contract.
func NewFLIPPERMOMCaller(address common.Address, caller bind.ContractCaller) (*FLIPPERMOMCaller, error) {
	contract, err := bindFLIPPERMOM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMCaller{contract: contract}, nil
}

// NewFLIPPERMOMTransactor creates a new write-only instance of FLIPPERMOM, bound to a specific deployed contract.
func NewFLIPPERMOMTransactor(address common.Address, transactor bind.ContractTransactor) (*FLIPPERMOMTransactor, error) {
	contract, err := bindFLIPPERMOM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMTransactor{contract: contract}, nil
}

// NewFLIPPERMOMFilterer creates a new log filterer instance of FLIPPERMOM, bound to a specific deployed contract.
func NewFLIPPERMOMFilterer(address common.Address, filterer bind.ContractFilterer) (*FLIPPERMOMFilterer, error) {
	contract, err := bindFLIPPERMOM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMFilterer{contract: contract}, nil
}

// bindFLIPPERMOM binds a generic wrapper to an already deployed contract.
func bindFLIPPERMOM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FLIPPERMOMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLIPPERMOM *FLIPPERMOMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLIPPERMOM.Contract.FLIPPERMOMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLIPPERMOM *FLIPPERMOMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.FLIPPERMOMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLIPPERMOM *FLIPPERMOMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.FLIPPERMOMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLIPPERMOM *FLIPPERMOMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLIPPERMOM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLIPPERMOM *FLIPPERMOMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLIPPERMOM *FLIPPERMOMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FLIPPERMOM.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMSession) Authority() (common.Address, error) {
	return _FLIPPERMOM.Contract.Authority(&_FLIPPERMOM.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCallerSession) Authority() (common.Address, error) {
	return _FLIPPERMOM.Contract.Authority(&_FLIPPERMOM.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCaller) Cat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FLIPPERMOM.contract.Call(opts, &out, "cat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMSession) Cat() (common.Address, error) {
	return _FLIPPERMOM.Contract.Cat(&_FLIPPERMOM.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCallerSession) Cat() (common.Address, error) {
	return _FLIPPERMOM.Contract.Cat(&_FLIPPERMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FLIPPERMOM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMSession) Owner() (common.Address, error) {
	return _FLIPPERMOM.Contract.Owner(&_FLIPPERMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FLIPPERMOM *FLIPPERMOMCallerSession) Owner() (common.Address, error) {
	return _FLIPPERMOM.Contract.Owner(&_FLIPPERMOM.CallOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactor) Deny(opts *bind.TransactOpts, flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.contract.Transact(opts, "deny", flip)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMSession) Deny(flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.Deny(&_FLIPPERMOM.TransactOpts, flip)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactorSession) Deny(flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.Deny(&_FLIPPERMOM.TransactOpts, flip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactor) Rely(opts *bind.TransactOpts, flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.contract.Transact(opts, "rely", flip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMSession) Rely(flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.Rely(&_FLIPPERMOM.TransactOpts, flip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address flip) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactorSession) Rely(flip common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.Rely(&_FLIPPERMOM.TransactOpts, flip)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_FLIPPERMOM *FLIPPERMOMSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.SetAuthority(&_FLIPPERMOM.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.SetAuthority(&_FLIPPERMOM.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FLIPPERMOM *FLIPPERMOMSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.SetOwner(&_FLIPPERMOM.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FLIPPERMOM *FLIPPERMOMTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _FLIPPERMOM.Contract.SetOwner(&_FLIPPERMOM.TransactOpts, owner_)
}

// FLIPPERMOMDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the FLIPPERMOM contract.
type FLIPPERMOMDenyIterator struct {
	Event *FLIPPERMOMDeny // Event containing the contract specifics and raw log

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
func (it *FLIPPERMOMDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FLIPPERMOMDeny)
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
		it.Event = new(FLIPPERMOMDeny)
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
func (it *FLIPPERMOMDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FLIPPERMOMDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FLIPPERMOMDeny represents a Deny event raised by the FLIPPERMOM contract.
type FLIPPERMOMDeny struct {
	Flip common.Address
	Cat  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x60dd05bd58bdeef1594e95d6b8e3c6bbca8a6ecfc2df9fb42b0870b7d788e316.
//
// Solidity: event Deny(address flip, address cat)
func (_FLIPPERMOM *FLIPPERMOMFilterer) FilterDeny(opts *bind.FilterOpts) (*FLIPPERMOMDenyIterator, error) {

	logs, sub, err := _FLIPPERMOM.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMDenyIterator{contract: _FLIPPERMOM.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x60dd05bd58bdeef1594e95d6b8e3c6bbca8a6ecfc2df9fb42b0870b7d788e316.
//
// Solidity: event Deny(address flip, address cat)
func (_FLIPPERMOM *FLIPPERMOMFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *FLIPPERMOMDeny) (event.Subscription, error) {

	logs, sub, err := _FLIPPERMOM.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FLIPPERMOMDeny)
				if err := _FLIPPERMOM.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x60dd05bd58bdeef1594e95d6b8e3c6bbca8a6ecfc2df9fb42b0870b7d788e316.
//
// Solidity: event Deny(address flip, address cat)
func (_FLIPPERMOM *FLIPPERMOMFilterer) ParseDeny(log types.Log) (*FLIPPERMOMDeny, error) {
	event := new(FLIPPERMOMDeny)
	if err := _FLIPPERMOM.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FLIPPERMOMRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the FLIPPERMOM contract.
type FLIPPERMOMRelyIterator struct {
	Event *FLIPPERMOMRely // Event containing the contract specifics and raw log

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
func (it *FLIPPERMOMRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FLIPPERMOMRely)
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
		it.Event = new(FLIPPERMOMRely)
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
func (it *FLIPPERMOMRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FLIPPERMOMRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FLIPPERMOMRely represents a Rely event raised by the FLIPPERMOM contract.
type FLIPPERMOMRely struct {
	Flip common.Address
	Usr  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0x08374d5a3379fefca0a4640028fd6eea4a6e36073bab3f235124a97c36c5af68.
//
// Solidity: event Rely(address flip, address usr)
func (_FLIPPERMOM *FLIPPERMOMFilterer) FilterRely(opts *bind.FilterOpts) (*FLIPPERMOMRelyIterator, error) {

	logs, sub, err := _FLIPPERMOM.contract.FilterLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMRelyIterator{contract: _FLIPPERMOM.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0x08374d5a3379fefca0a4640028fd6eea4a6e36073bab3f235124a97c36c5af68.
//
// Solidity: event Rely(address flip, address usr)
func (_FLIPPERMOM *FLIPPERMOMFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *FLIPPERMOMRely) (event.Subscription, error) {

	logs, sub, err := _FLIPPERMOM.contract.WatchLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FLIPPERMOMRely)
				if err := _FLIPPERMOM.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0x08374d5a3379fefca0a4640028fd6eea4a6e36073bab3f235124a97c36c5af68.
//
// Solidity: event Rely(address flip, address usr)
func (_FLIPPERMOM *FLIPPERMOMFilterer) ParseRely(log types.Log) (*FLIPPERMOMRely, error) {
	event := new(FLIPPERMOMRely)
	if err := _FLIPPERMOM.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FLIPPERMOMSetAuthorityIterator is returned from FilterSetAuthority and is used to iterate over the raw logs and unpacked data for SetAuthority events raised by the FLIPPERMOM contract.
type FLIPPERMOMSetAuthorityIterator struct {
	Event *FLIPPERMOMSetAuthority // Event containing the contract specifics and raw log

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
func (it *FLIPPERMOMSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FLIPPERMOMSetAuthority)
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
		it.Event = new(FLIPPERMOMSetAuthority)
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
func (it *FLIPPERMOMSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FLIPPERMOMSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FLIPPERMOMSetAuthority represents a SetAuthority event raised by the FLIPPERMOM contract.
type FLIPPERMOMSetAuthority struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAuthority is a free log retrieval operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address oldAuthority, address newAuthority)
func (_FLIPPERMOM *FLIPPERMOMFilterer) FilterSetAuthority(opts *bind.FilterOpts) (*FLIPPERMOMSetAuthorityIterator, error) {

	logs, sub, err := _FLIPPERMOM.contract.FilterLogs(opts, "SetAuthority")
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMSetAuthorityIterator{contract: _FLIPPERMOM.contract, event: "SetAuthority", logs: logs, sub: sub}, nil
}

// WatchSetAuthority is a free log subscription operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address oldAuthority, address newAuthority)
func (_FLIPPERMOM *FLIPPERMOMFilterer) WatchSetAuthority(opts *bind.WatchOpts, sink chan<- *FLIPPERMOMSetAuthority) (event.Subscription, error) {

	logs, sub, err := _FLIPPERMOM.contract.WatchLogs(opts, "SetAuthority")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FLIPPERMOMSetAuthority)
				if err := _FLIPPERMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
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
// Solidity: event SetAuthority(address oldAuthority, address newAuthority)
func (_FLIPPERMOM *FLIPPERMOMFilterer) ParseSetAuthority(log types.Log) (*FLIPPERMOMSetAuthority, error) {
	event := new(FLIPPERMOMSetAuthority)
	if err := _FLIPPERMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FLIPPERMOMSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the FLIPPERMOM contract.
type FLIPPERMOMSetOwnerIterator struct {
	Event *FLIPPERMOMSetOwner // Event containing the contract specifics and raw log

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
func (it *FLIPPERMOMSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FLIPPERMOMSetOwner)
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
		it.Event = new(FLIPPERMOMSetOwner)
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
func (it *FLIPPERMOMSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FLIPPERMOMSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FLIPPERMOMSetOwner represents a SetOwner event raised by the FLIPPERMOM contract.
type FLIPPERMOMSetOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address oldOwner, address newOwner)
func (_FLIPPERMOM *FLIPPERMOMFilterer) FilterSetOwner(opts *bind.FilterOpts) (*FLIPPERMOMSetOwnerIterator, error) {

	logs, sub, err := _FLIPPERMOM.contract.FilterLogs(opts, "SetOwner")
	if err != nil {
		return nil, err
	}
	return &FLIPPERMOMSetOwnerIterator{contract: _FLIPPERMOM.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address oldOwner, address newOwner)
func (_FLIPPERMOM *FLIPPERMOMFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *FLIPPERMOMSetOwner) (event.Subscription, error) {

	logs, sub, err := _FLIPPERMOM.contract.WatchLogs(opts, "SetOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FLIPPERMOMSetOwner)
				if err := _FLIPPERMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
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
// Solidity: event SetOwner(address oldOwner, address newOwner)
func (_FLIPPERMOM *FLIPPERMOMFilterer) ParseSetOwner(log types.Log) (*FLIPPERMOMSetOwner, error) {
	event := new(FLIPPERMOMSetOwner)
	if err := _FLIPPERMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
