// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_DEPLOYER

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

// PROXYDEPLOYERABI is the input ABI used to generate the binding from.
const PROXYDEPLOYERABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"response\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"response\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"name\":\"setCache\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// PROXYDEPLOYER is an auto generated Go binding around an Ethereum contract.
type PROXYDEPLOYER struct {
	PROXYDEPLOYERCaller     // Read-only binding to the contract
	PROXYDEPLOYERTransactor // Write-only binding to the contract
	PROXYDEPLOYERFilterer   // Log filterer for contract events
}

// PROXYDEPLOYERCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYDEPLOYERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYDEPLOYERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYDEPLOYERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYDEPLOYERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYDEPLOYERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYDEPLOYERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYDEPLOYERSession struct {
	Contract     *PROXYDEPLOYER    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PROXYDEPLOYERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYDEPLOYERCallerSession struct {
	Contract *PROXYDEPLOYERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PROXYDEPLOYERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYDEPLOYERTransactorSession struct {
	Contract     *PROXYDEPLOYERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PROXYDEPLOYERRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYDEPLOYERRaw struct {
	Contract *PROXYDEPLOYER // Generic contract binding to access the raw methods on
}

// PROXYDEPLOYERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYDEPLOYERCallerRaw struct {
	Contract *PROXYDEPLOYERCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYDEPLOYERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYDEPLOYERTransactorRaw struct {
	Contract *PROXYDEPLOYERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYDEPLOYER creates a new instance of PROXYDEPLOYER, bound to a specific deployed contract.
func NewPROXYDEPLOYER(address common.Address, backend bind.ContractBackend) (*PROXYDEPLOYER, error) {
	contract, err := bindPROXYDEPLOYER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYER{PROXYDEPLOYERCaller: PROXYDEPLOYERCaller{contract: contract}, PROXYDEPLOYERTransactor: PROXYDEPLOYERTransactor{contract: contract}, PROXYDEPLOYERFilterer: PROXYDEPLOYERFilterer{contract: contract}}, nil
}

// NewPROXYDEPLOYERCaller creates a new read-only instance of PROXYDEPLOYER, bound to a specific deployed contract.
func NewPROXYDEPLOYERCaller(address common.Address, caller bind.ContractCaller) (*PROXYDEPLOYERCaller, error) {
	contract, err := bindPROXYDEPLOYER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYERCaller{contract: contract}, nil
}

// NewPROXYDEPLOYERTransactor creates a new write-only instance of PROXYDEPLOYER, bound to a specific deployed contract.
func NewPROXYDEPLOYERTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYDEPLOYERTransactor, error) {
	contract, err := bindPROXYDEPLOYER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYERTransactor{contract: contract}, nil
}

// NewPROXYDEPLOYERFilterer creates a new log filterer instance of PROXYDEPLOYER, bound to a specific deployed contract.
func NewPROXYDEPLOYERFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYDEPLOYERFilterer, error) {
	contract, err := bindPROXYDEPLOYER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYERFilterer{contract: contract}, nil
}

// bindPROXYDEPLOYER binds a generic wrapper to an already deployed contract.
func bindPROXYDEPLOYER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYDEPLOYERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYDEPLOYER *PROXYDEPLOYERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYDEPLOYER.Contract.PROXYDEPLOYERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYDEPLOYER *PROXYDEPLOYERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.PROXYDEPLOYERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYDEPLOYER *PROXYDEPLOYERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.PROXYDEPLOYERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYDEPLOYER *PROXYDEPLOYERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYDEPLOYER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYDEPLOYER.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Authority() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Authority(&_PROXYDEPLOYER.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCallerSession) Authority() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Authority(&_PROXYDEPLOYER.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYDEPLOYER.contract.Call(opts, &out, "cache")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Cache() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Cache(&_PROXYDEPLOYER.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCallerSession) Cache() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Cache(&_PROXYDEPLOYER.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYDEPLOYER.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Owner() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Owner(&_PROXYDEPLOYER.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PROXYDEPLOYER *PROXYDEPLOYERCallerSession) Owner() (common.Address, error) {
	return _PROXYDEPLOYER.Contract.Owner(&_PROXYDEPLOYER.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.Transact(opts, "execute", _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Execute(&_PROXYDEPLOYER.TransactOpts, _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Execute(&_PROXYDEPLOYER.TransactOpts, _target, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) Execute0(opts *bind.TransactOpts, _code []byte, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.Transact(opts, "execute0", _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Execute0(&_PROXYDEPLOYER.TransactOpts, _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes32 response)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Execute0(&_PROXYDEPLOYER.TransactOpts, _code, _data)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetAuthority(&_PROXYDEPLOYER.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetAuthority(&_PROXYDEPLOYER.TransactOpts, authority_)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) SetCache(opts *bind.TransactOpts, _cacheAddr common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.Transact(opts, "setCache", _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetCache(&_PROXYDEPLOYER.TransactOpts, _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetCache(&_PROXYDEPLOYER.TransactOpts, _cacheAddr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetOwner(&_PROXYDEPLOYER.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.SetOwner(&_PROXYDEPLOYER.TransactOpts, owner_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Fallback(&_PROXYDEPLOYER.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PROXYDEPLOYER *PROXYDEPLOYERTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PROXYDEPLOYER.Contract.Fallback(&_PROXYDEPLOYER.TransactOpts, calldata)
}

// PROXYDEPLOYERLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the PROXYDEPLOYER contract.
type PROXYDEPLOYERLogSetAuthorityIterator struct {
	Event *PROXYDEPLOYERLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *PROXYDEPLOYERLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PROXYDEPLOYERLogSetAuthority)
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
		it.Event = new(PROXYDEPLOYERLogSetAuthority)
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
func (it *PROXYDEPLOYERLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PROXYDEPLOYERLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PROXYDEPLOYERLogSetAuthority represents a LogSetAuthority event raised by the PROXYDEPLOYER contract.
type PROXYDEPLOYERLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*PROXYDEPLOYERLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _PROXYDEPLOYER.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYERLogSetAuthorityIterator{contract: _PROXYDEPLOYER.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *PROXYDEPLOYERLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _PROXYDEPLOYER.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PROXYDEPLOYERLogSetAuthority)
				if err := _PROXYDEPLOYER.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) ParseLogSetAuthority(log types.Log) (*PROXYDEPLOYERLogSetAuthority, error) {
	event := new(PROXYDEPLOYERLogSetAuthority)
	if err := _PROXYDEPLOYER.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PROXYDEPLOYERLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the PROXYDEPLOYER contract.
type PROXYDEPLOYERLogSetOwnerIterator struct {
	Event *PROXYDEPLOYERLogSetOwner // Event containing the contract specifics and raw log

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
func (it *PROXYDEPLOYERLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PROXYDEPLOYERLogSetOwner)
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
		it.Event = new(PROXYDEPLOYERLogSetOwner)
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
func (it *PROXYDEPLOYERLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PROXYDEPLOYERLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PROXYDEPLOYERLogSetOwner represents a LogSetOwner event raised by the PROXYDEPLOYER contract.
type PROXYDEPLOYERLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*PROXYDEPLOYERLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PROXYDEPLOYER.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PROXYDEPLOYERLogSetOwnerIterator{contract: _PROXYDEPLOYER.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *PROXYDEPLOYERLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PROXYDEPLOYER.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PROXYDEPLOYERLogSetOwner)
				if err := _PROXYDEPLOYER.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PROXYDEPLOYER *PROXYDEPLOYERFilterer) ParseLogSetOwner(log types.Log) (*PROXYDEPLOYERLogSetOwner, error) {
	event := new(PROXYDEPLOYERLogSetOwner)
	if err := _PROXYDEPLOYER.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
