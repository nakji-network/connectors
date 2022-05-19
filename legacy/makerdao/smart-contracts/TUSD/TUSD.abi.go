// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TUSD

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

// TUSDABI is the input ABI used to generate the binding from.
const TUSDABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proxyOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingProxyOwner\",\"outputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimProxyOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferProxyOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"ProxyOwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// TUSD is an auto generated Go binding around an Ethereum contract.
type TUSD struct {
	TUSDCaller     // Read-only binding to the contract
	TUSDTransactor // Write-only binding to the contract
	TUSDFilterer   // Log filterer for contract events
}

// TUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type TUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TUSDSession struct {
	Contract     *TUSD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TUSDCallerSession struct {
	Contract *TUSDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TUSDTransactorSession struct {
	Contract     *TUSDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type TUSDRaw struct {
	Contract *TUSD // Generic contract binding to access the raw methods on
}

// TUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TUSDCallerRaw struct {
	Contract *TUSDCaller // Generic read-only contract binding to access the raw methods on
}

// TUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TUSDTransactorRaw struct {
	Contract *TUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTUSD creates a new instance of TUSD, bound to a specific deployed contract.
func NewTUSD(address common.Address, backend bind.ContractBackend) (*TUSD, error) {
	contract, err := bindTUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TUSD{TUSDCaller: TUSDCaller{contract: contract}, TUSDTransactor: TUSDTransactor{contract: contract}, TUSDFilterer: TUSDFilterer{contract: contract}}, nil
}

// NewTUSDCaller creates a new read-only instance of TUSD, bound to a specific deployed contract.
func NewTUSDCaller(address common.Address, caller bind.ContractCaller) (*TUSDCaller, error) {
	contract, err := bindTUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TUSDCaller{contract: contract}, nil
}

// NewTUSDTransactor creates a new write-only instance of TUSD, bound to a specific deployed contract.
func NewTUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*TUSDTransactor, error) {
	contract, err := bindTUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TUSDTransactor{contract: contract}, nil
}

// NewTUSDFilterer creates a new log filterer instance of TUSD, bound to a specific deployed contract.
func NewTUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*TUSDFilterer, error) {
	contract, err := bindTUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TUSDFilterer{contract: contract}, nil
}

// bindTUSD binds a generic wrapper to an already deployed contract.
func bindTUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TUSD *TUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TUSD.Contract.TUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TUSD *TUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TUSD.Contract.TUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TUSD *TUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TUSD.Contract.TUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TUSD *TUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TUSD *TUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TUSD *TUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TUSD.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_TUSD *TUSDCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TUSD.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_TUSD *TUSDSession) Implementation() (common.Address, error) {
	return _TUSD.Contract.Implementation(&_TUSD.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_TUSD *TUSDCallerSession) Implementation() (common.Address, error) {
	return _TUSD.Contract.Implementation(&_TUSD.CallOpts)
}

// PendingProxyOwner is a free data retrieval call binding the contract method 0x0add8140.
//
// Solidity: function pendingProxyOwner() view returns(address pendingOwner)
func (_TUSD *TUSDCaller) PendingProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TUSD.contract.Call(opts, &out, "pendingProxyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxyOwner is a free data retrieval call binding the contract method 0x0add8140.
//
// Solidity: function pendingProxyOwner() view returns(address pendingOwner)
func (_TUSD *TUSDSession) PendingProxyOwner() (common.Address, error) {
	return _TUSD.Contract.PendingProxyOwner(&_TUSD.CallOpts)
}

// PendingProxyOwner is a free data retrieval call binding the contract method 0x0add8140.
//
// Solidity: function pendingProxyOwner() view returns(address pendingOwner)
func (_TUSD *TUSDCallerSession) PendingProxyOwner() (common.Address, error) {
	return _TUSD.Contract.PendingProxyOwner(&_TUSD.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address owner)
func (_TUSD *TUSDCaller) ProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TUSD.contract.Call(opts, &out, "proxyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address owner)
func (_TUSD *TUSDSession) ProxyOwner() (common.Address, error) {
	return _TUSD.Contract.ProxyOwner(&_TUSD.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address owner)
func (_TUSD *TUSDCallerSession) ProxyOwner() (common.Address, error) {
	return _TUSD.Contract.ProxyOwner(&_TUSD.CallOpts)
}

// ClaimProxyOwnership is a paid mutator transaction binding the contract method 0x9965b3d6.
//
// Solidity: function claimProxyOwnership() returns()
func (_TUSD *TUSDTransactor) ClaimProxyOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TUSD.contract.Transact(opts, "claimProxyOwnership")
}

// ClaimProxyOwnership is a paid mutator transaction binding the contract method 0x9965b3d6.
//
// Solidity: function claimProxyOwnership() returns()
func (_TUSD *TUSDSession) ClaimProxyOwnership() (*types.Transaction, error) {
	return _TUSD.Contract.ClaimProxyOwnership(&_TUSD.TransactOpts)
}

// ClaimProxyOwnership is a paid mutator transaction binding the contract method 0x9965b3d6.
//
// Solidity: function claimProxyOwnership() returns()
func (_TUSD *TUSDTransactorSession) ClaimProxyOwnership() (*types.Transaction, error) {
	return _TUSD.Contract.ClaimProxyOwnership(&_TUSD.TransactOpts)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_TUSD *TUSDTransactor) TransferProxyOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TUSD.contract.Transact(opts, "transferProxyOwnership", newOwner)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_TUSD *TUSDSession) TransferProxyOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TUSD.Contract.TransferProxyOwnership(&_TUSD.TransactOpts, newOwner)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_TUSD *TUSDTransactorSession) TransferProxyOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TUSD.Contract.TransferProxyOwnership(&_TUSD.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address implementation) returns()
func (_TUSD *TUSDTransactor) UpgradeTo(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _TUSD.contract.Transact(opts, "upgradeTo", implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address implementation) returns()
func (_TUSD *TUSDSession) UpgradeTo(implementation common.Address) (*types.Transaction, error) {
	return _TUSD.Contract.UpgradeTo(&_TUSD.TransactOpts, implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address implementation) returns()
func (_TUSD *TUSDTransactorSession) UpgradeTo(implementation common.Address) (*types.Transaction, error) {
	return _TUSD.Contract.UpgradeTo(&_TUSD.TransactOpts, implementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TUSD *TUSDTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TUSD.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TUSD *TUSDSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TUSD.Contract.Fallback(&_TUSD.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TUSD *TUSDTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TUSD.Contract.Fallback(&_TUSD.TransactOpts, calldata)
}

// TUSDNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the TUSD contract.
type TUSDNewPendingOwnerIterator struct {
	Event *TUSDNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *TUSDNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TUSDNewPendingOwner)
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
		it.Event = new(TUSDNewPendingOwner)
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
func (it *TUSDNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TUSDNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TUSDNewPendingOwner represents a NewPendingOwner event raised by the TUSD contract.
type TUSDNewPendingOwner struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address currentOwner, address pendingOwner)
func (_TUSD *TUSDFilterer) FilterNewPendingOwner(opts *bind.FilterOpts) (*TUSDNewPendingOwnerIterator, error) {

	logs, sub, err := _TUSD.contract.FilterLogs(opts, "NewPendingOwner")
	if err != nil {
		return nil, err
	}
	return &TUSDNewPendingOwnerIterator{contract: _TUSD.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address currentOwner, address pendingOwner)
func (_TUSD *TUSDFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *TUSDNewPendingOwner) (event.Subscription, error) {

	logs, sub, err := _TUSD.contract.WatchLogs(opts, "NewPendingOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TUSDNewPendingOwner)
				if err := _TUSD.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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

// ParseNewPendingOwner is a log parse operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address currentOwner, address pendingOwner)
func (_TUSD *TUSDFilterer) ParseNewPendingOwner(log types.Log) (*TUSDNewPendingOwner, error) {
	event := new(TUSDNewPendingOwner)
	if err := _TUSD.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TUSDProxyOwnershipTransferredIterator is returned from FilterProxyOwnershipTransferred and is used to iterate over the raw logs and unpacked data for ProxyOwnershipTransferred events raised by the TUSD contract.
type TUSDProxyOwnershipTransferredIterator struct {
	Event *TUSDProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TUSDProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TUSDProxyOwnershipTransferred)
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
		it.Event = new(TUSDProxyOwnershipTransferred)
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
func (it *TUSDProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TUSDProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TUSDProxyOwnershipTransferred represents a ProxyOwnershipTransferred event raised by the TUSD contract.
type TUSDProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProxyOwnershipTransferred is a free log retrieval operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TUSD *TUSDFilterer) FilterProxyOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TUSDProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TUSD.contract.FilterLogs(opts, "ProxyOwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TUSDProxyOwnershipTransferredIterator{contract: _TUSD.contract, event: "ProxyOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchProxyOwnershipTransferred is a free log subscription operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TUSD *TUSDFilterer) WatchProxyOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TUSDProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TUSD.contract.WatchLogs(opts, "ProxyOwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TUSDProxyOwnershipTransferred)
				if err := _TUSD.contract.UnpackLog(event, "ProxyOwnershipTransferred", log); err != nil {
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

// ParseProxyOwnershipTransferred is a log parse operation binding the contract event 0x5a3e66efaa1e445ebd894728a69d6959842ea1e97bd79b892797106e270efcd9.
//
// Solidity: event ProxyOwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TUSD *TUSDFilterer) ParseProxyOwnershipTransferred(log types.Log) (*TUSDProxyOwnershipTransferred, error) {
	event := new(TUSDProxyOwnershipTransferred)
	if err := _TUSD.contract.UnpackLog(event, "ProxyOwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TUSDUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TUSD contract.
type TUSDUpgradedIterator struct {
	Event *TUSDUpgraded // Event containing the contract specifics and raw log

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
func (it *TUSDUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TUSDUpgraded)
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
		it.Event = new(TUSDUpgraded)
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
func (it *TUSDUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TUSDUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TUSDUpgraded represents a Upgraded event raised by the TUSD contract.
type TUSDUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TUSD *TUSDFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TUSDUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TUSD.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TUSDUpgradedIterator{contract: _TUSD.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TUSD *TUSDFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TUSDUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TUSD.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TUSDUpgraded)
				if err := _TUSD.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
// Solidity: event Upgraded(address indexed implementation)
func (_TUSD *TUSDFilterer) ParseUpgraded(log types.Log) (*TUSDUpgraded, error) {
	event := new(TUSDUpgraded)
	if err := _TUSD.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
