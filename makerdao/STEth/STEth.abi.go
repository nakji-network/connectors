// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package STEth

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

// STETHABI is the input ABI used to generate the binding from.
const STETHABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proxyType\",\"outputs\":[{\"name\":\"proxyTypeId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDepositable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kernel\",\"type\":\"address\"},{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_initializePayload\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProxyDeposit\",\"type\":\"event\"}]"

// STETH is an auto generated Go binding around an Ethereum contract.
type STETH struct {
	STETHCaller     // Read-only binding to the contract
	STETHTransactor // Write-only binding to the contract
	STETHFilterer   // Log filterer for contract events
}

// STETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type STETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// STETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type STETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// STETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type STETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// STETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type STETHSession struct {
	Contract     *STETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// STETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type STETHCallerSession struct {
	Contract *STETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// STETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type STETHTransactorSession struct {
	Contract     *STETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// STETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type STETHRaw struct {
	Contract *STETH // Generic contract binding to access the raw methods on
}

// STETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type STETHCallerRaw struct {
	Contract *STETHCaller // Generic read-only contract binding to access the raw methods on
}

// STETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type STETHTransactorRaw struct {
	Contract *STETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSTETH creates a new instance of STETH, bound to a specific deployed contract.
func NewSTETH(address common.Address, backend bind.ContractBackend) (*STETH, error) {
	contract, err := bindSTETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &STETH{STETHCaller: STETHCaller{contract: contract}, STETHTransactor: STETHTransactor{contract: contract}, STETHFilterer: STETHFilterer{contract: contract}}, nil
}

// NewSTETHCaller creates a new read-only instance of STETH, bound to a specific deployed contract.
func NewSTETHCaller(address common.Address, caller bind.ContractCaller) (*STETHCaller, error) {
	contract, err := bindSTETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &STETHCaller{contract: contract}, nil
}

// NewSTETHTransactor creates a new write-only instance of STETH, bound to a specific deployed contract.
func NewSTETHTransactor(address common.Address, transactor bind.ContractTransactor) (*STETHTransactor, error) {
	contract, err := bindSTETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &STETHTransactor{contract: contract}, nil
}

// NewSTETHFilterer creates a new log filterer instance of STETH, bound to a specific deployed contract.
func NewSTETHFilterer(address common.Address, filterer bind.ContractFilterer) (*STETHFilterer, error) {
	contract, err := bindSTETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &STETHFilterer{contract: contract}, nil
}

// bindSTETH binds a generic wrapper to an already deployed contract.
func bindSTETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(STETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_STETH *STETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _STETH.Contract.STETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_STETH *STETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _STETH.Contract.STETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_STETH *STETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _STETH.Contract.STETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_STETH *STETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _STETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_STETH *STETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _STETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_STETH *STETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _STETH.Contract.contract.Transact(opts, method, params...)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_STETH *STETHCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _STETH.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_STETH *STETHSession) AppId() ([32]byte, error) {
	return _STETH.Contract.AppId(&_STETH.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_STETH *STETHCallerSession) AppId() ([32]byte, error) {
	return _STETH.Contract.AppId(&_STETH.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_STETH *STETHCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _STETH.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_STETH *STETHSession) Implementation() (common.Address, error) {
	return _STETH.Contract.Implementation(&_STETH.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_STETH *STETHCallerSession) Implementation() (common.Address, error) {
	return _STETH.Contract.Implementation(&_STETH.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_STETH *STETHCaller) IsDepositable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _STETH.contract.Call(opts, &out, "isDepositable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_STETH *STETHSession) IsDepositable() (bool, error) {
	return _STETH.Contract.IsDepositable(&_STETH.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_STETH *STETHCallerSession) IsDepositable() (bool, error) {
	return _STETH.Contract.IsDepositable(&_STETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_STETH *STETHCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _STETH.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_STETH *STETHSession) Kernel() (common.Address, error) {
	return _STETH.Contract.Kernel(&_STETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_STETH *STETHCallerSession) Kernel() (common.Address, error) {
	return _STETH.Contract.Kernel(&_STETH.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_STETH *STETHCaller) ProxyType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _STETH.contract.Call(opts, &out, "proxyType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_STETH *STETHSession) ProxyType() (*big.Int, error) {
	return _STETH.Contract.ProxyType(&_STETH.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_STETH *STETHCallerSession) ProxyType() (*big.Int, error) {
	return _STETH.Contract.ProxyType(&_STETH.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_STETH *STETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _STETH.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_STETH *STETHSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _STETH.Contract.Fallback(&_STETH.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_STETH *STETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _STETH.Contract.Fallback(&_STETH.TransactOpts, calldata)
}

// STETHProxyDepositIterator is returned from FilterProxyDeposit and is used to iterate over the raw logs and unpacked data for ProxyDeposit events raised by the STETH contract.
type STETHProxyDepositIterator struct {
	Event *STETHProxyDeposit // Event containing the contract specifics and raw log

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
func (it *STETHProxyDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(STETHProxyDeposit)
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
		it.Event = new(STETHProxyDeposit)
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
func (it *STETHProxyDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *STETHProxyDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// STETHProxyDeposit represents a ProxyDeposit event raised by the STETH contract.
type STETHProxyDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProxyDeposit is a free log retrieval operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_STETH *STETHFilterer) FilterProxyDeposit(opts *bind.FilterOpts) (*STETHProxyDepositIterator, error) {

	logs, sub, err := _STETH.contract.FilterLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return &STETHProxyDepositIterator{contract: _STETH.contract, event: "ProxyDeposit", logs: logs, sub: sub}, nil
}

// WatchProxyDeposit is a free log subscription operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_STETH *STETHFilterer) WatchProxyDeposit(opts *bind.WatchOpts, sink chan<- *STETHProxyDeposit) (event.Subscription, error) {

	logs, sub, err := _STETH.contract.WatchLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(STETHProxyDeposit)
				if err := _STETH.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
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

// ParseProxyDeposit is a log parse operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_STETH *STETHFilterer) ParseProxyDeposit(log types.Log) (*STETHProxyDeposit, error) {
	event := new(STETHProxyDeposit)
	if err := _STETH.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
