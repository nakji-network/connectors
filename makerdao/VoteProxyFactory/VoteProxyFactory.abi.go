// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VoteProxyFactory

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

// VOTEPROXYFACTORYABI is the input ABI used to generate the binding from.
const VOTEPROXYFACTORYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"chief_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cold\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hot\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voteProxy\",\"type\":\"address\"}],\"name\":\"LinkConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cold\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hot\",\"type\":\"address\"}],\"name\":\"LinkRequested\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cold\",\"type\":\"address\"}],\"name\":\"approveLink\",\"outputs\":[{\"internalType\":\"contractVoteProxy\",\"name\":\"voteProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"breakLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chief\",\"outputs\":[{\"internalType\":\"contractChiefLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coldMap\",\"outputs\":[{\"internalType\":\"contractVoteProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"hasProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hotMap\",\"outputs\":[{\"internalType\":\"contractVoteProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"hot\",\"type\":\"address\"}],\"name\":\"initiateLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"linkRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"linkSelf\",\"outputs\":[{\"internalType\":\"contractVoteProxy\",\"name\":\"voteProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VOTEPROXYFACTORY is an auto generated Go binding around an Ethereum contract.
type VOTEPROXYFACTORY struct {
	VOTEPROXYFACTORYCaller     // Read-only binding to the contract
	VOTEPROXYFACTORYTransactor // Write-only binding to the contract
	VOTEPROXYFACTORYFilterer   // Log filterer for contract events
}

// VOTEPROXYFACTORYCaller is an auto generated read-only Go binding around an Ethereum contract.
type VOTEPROXYFACTORYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEPROXYFACTORYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VOTEPROXYFACTORYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEPROXYFACTORYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VOTEPROXYFACTORYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEPROXYFACTORYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VOTEPROXYFACTORYSession struct {
	Contract     *VOTEPROXYFACTORY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VOTEPROXYFACTORYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VOTEPROXYFACTORYCallerSession struct {
	Contract *VOTEPROXYFACTORYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VOTEPROXYFACTORYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VOTEPROXYFACTORYTransactorSession struct {
	Contract     *VOTEPROXYFACTORYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VOTEPROXYFACTORYRaw is an auto generated low-level Go binding around an Ethereum contract.
type VOTEPROXYFACTORYRaw struct {
	Contract *VOTEPROXYFACTORY // Generic contract binding to access the raw methods on
}

// VOTEPROXYFACTORYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VOTEPROXYFACTORYCallerRaw struct {
	Contract *VOTEPROXYFACTORYCaller // Generic read-only contract binding to access the raw methods on
}

// VOTEPROXYFACTORYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VOTEPROXYFACTORYTransactorRaw struct {
	Contract *VOTEPROXYFACTORYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVOTEPROXYFACTORY creates a new instance of VOTEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEPROXYFACTORY(address common.Address, backend bind.ContractBackend) (*VOTEPROXYFACTORY, error) {
	contract, err := bindVOTEPROXYFACTORY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORY{VOTEPROXYFACTORYCaller: VOTEPROXYFACTORYCaller{contract: contract}, VOTEPROXYFACTORYTransactor: VOTEPROXYFACTORYTransactor{contract: contract}, VOTEPROXYFACTORYFilterer: VOTEPROXYFACTORYFilterer{contract: contract}}, nil
}

// NewVOTEPROXYFACTORYCaller creates a new read-only instance of VOTEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEPROXYFACTORYCaller(address common.Address, caller bind.ContractCaller) (*VOTEPROXYFACTORYCaller, error) {
	contract, err := bindVOTEPROXYFACTORY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORYCaller{contract: contract}, nil
}

// NewVOTEPROXYFACTORYTransactor creates a new write-only instance of VOTEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEPROXYFACTORYTransactor(address common.Address, transactor bind.ContractTransactor) (*VOTEPROXYFACTORYTransactor, error) {
	contract, err := bindVOTEPROXYFACTORY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORYTransactor{contract: contract}, nil
}

// NewVOTEPROXYFACTORYFilterer creates a new log filterer instance of VOTEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEPROXYFACTORYFilterer(address common.Address, filterer bind.ContractFilterer) (*VOTEPROXYFACTORYFilterer, error) {
	contract, err := bindVOTEPROXYFACTORY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORYFilterer{contract: contract}, nil
}

// bindVOTEPROXYFACTORY binds a generic wrapper to an already deployed contract.
func bindVOTEPROXYFACTORY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VOTEPROXYFACTORYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VOTEPROXYFACTORY.Contract.VOTEPROXYFACTORYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.VOTEPROXYFACTORYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.VOTEPROXYFACTORYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VOTEPROXYFACTORY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.contract.Transact(opts, method, params...)
}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCaller) Chief(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VOTEPROXYFACTORY.contract.Call(opts, &out, "chief")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) Chief() (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.Chief(&_VOTEPROXYFACTORY.CallOpts)
}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerSession) Chief() (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.Chief(&_VOTEPROXYFACTORY.CallOpts)
}

// ColdMap is a free data retrieval call binding the contract method 0xf14f4f96.
//
// Solidity: function coldMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCaller) ColdMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VOTEPROXYFACTORY.contract.Call(opts, &out, "coldMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ColdMap is a free data retrieval call binding the contract method 0xf14f4f96.
//
// Solidity: function coldMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) ColdMap(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.ColdMap(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// ColdMap is a free data retrieval call binding the contract method 0xf14f4f96.
//
// Solidity: function coldMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerSession) ColdMap(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.ColdMap(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// HasProxy is a free data retrieval call binding the contract method 0x5ab0bc73.
//
// Solidity: function hasProxy(address guy) view returns(bool)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCaller) HasProxy(opts *bind.CallOpts, guy common.Address) (bool, error) {
	var out []interface{}
	err := _VOTEPROXYFACTORY.contract.Call(opts, &out, "hasProxy", guy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasProxy is a free data retrieval call binding the contract method 0x5ab0bc73.
//
// Solidity: function hasProxy(address guy) view returns(bool)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) HasProxy(guy common.Address) (bool, error) {
	return _VOTEPROXYFACTORY.Contract.HasProxy(&_VOTEPROXYFACTORY.CallOpts, guy)
}

// HasProxy is a free data retrieval call binding the contract method 0x5ab0bc73.
//
// Solidity: function hasProxy(address guy) view returns(bool)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerSession) HasProxy(guy common.Address) (bool, error) {
	return _VOTEPROXYFACTORY.Contract.HasProxy(&_VOTEPROXYFACTORY.CallOpts, guy)
}

// HotMap is a free data retrieval call binding the contract method 0x6ef40889.
//
// Solidity: function hotMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCaller) HotMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VOTEPROXYFACTORY.contract.Call(opts, &out, "hotMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HotMap is a free data retrieval call binding the contract method 0x6ef40889.
//
// Solidity: function hotMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) HotMap(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.HotMap(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// HotMap is a free data retrieval call binding the contract method 0x6ef40889.
//
// Solidity: function hotMap(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerSession) HotMap(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.HotMap(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// LinkRequests is a free data retrieval call binding the contract method 0x7817ea66.
//
// Solidity: function linkRequests(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCaller) LinkRequests(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VOTEPROXYFACTORY.contract.Call(opts, &out, "linkRequests", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkRequests is a free data retrieval call binding the contract method 0x7817ea66.
//
// Solidity: function linkRequests(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) LinkRequests(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.LinkRequests(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// LinkRequests is a free data retrieval call binding the contract method 0x7817ea66.
//
// Solidity: function linkRequests(address ) view returns(address)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYCallerSession) LinkRequests(arg0 common.Address) (common.Address, error) {
	return _VOTEPROXYFACTORY.Contract.LinkRequests(&_VOTEPROXYFACTORY.CallOpts, arg0)
}

// ApproveLink is a paid mutator transaction binding the contract method 0xf363147a.
//
// Solidity: function approveLink(address cold) returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactor) ApproveLink(opts *bind.TransactOpts, cold common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.contract.Transact(opts, "approveLink", cold)
}

// ApproveLink is a paid mutator transaction binding the contract method 0xf363147a.
//
// Solidity: function approveLink(address cold) returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) ApproveLink(cold common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.ApproveLink(&_VOTEPROXYFACTORY.TransactOpts, cold)
}

// ApproveLink is a paid mutator transaction binding the contract method 0xf363147a.
//
// Solidity: function approveLink(address cold) returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorSession) ApproveLink(cold common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.ApproveLink(&_VOTEPROXYFACTORY.TransactOpts, cold)
}

// BreakLink is a paid mutator transaction binding the contract method 0xe83447d0.
//
// Solidity: function breakLink() returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactor) BreakLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.contract.Transact(opts, "breakLink")
}

// BreakLink is a paid mutator transaction binding the contract method 0xe83447d0.
//
// Solidity: function breakLink() returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) BreakLink() (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.BreakLink(&_VOTEPROXYFACTORY.TransactOpts)
}

// BreakLink is a paid mutator transaction binding the contract method 0xe83447d0.
//
// Solidity: function breakLink() returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorSession) BreakLink() (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.BreakLink(&_VOTEPROXYFACTORY.TransactOpts)
}

// InitiateLink is a paid mutator transaction binding the contract method 0x2d04ebd4.
//
// Solidity: function initiateLink(address hot) returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactor) InitiateLink(opts *bind.TransactOpts, hot common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.contract.Transact(opts, "initiateLink", hot)
}

// InitiateLink is a paid mutator transaction binding the contract method 0x2d04ebd4.
//
// Solidity: function initiateLink(address hot) returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) InitiateLink(hot common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.InitiateLink(&_VOTEPROXYFACTORY.TransactOpts, hot)
}

// InitiateLink is a paid mutator transaction binding the contract method 0x2d04ebd4.
//
// Solidity: function initiateLink(address hot) returns()
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorSession) InitiateLink(hot common.Address) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.InitiateLink(&_VOTEPROXYFACTORY.TransactOpts, hot)
}

// LinkSelf is a paid mutator transaction binding the contract method 0x09a14f2e.
//
// Solidity: function linkSelf() returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactor) LinkSelf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.contract.Transact(opts, "linkSelf")
}

// LinkSelf is a paid mutator transaction binding the contract method 0x09a14f2e.
//
// Solidity: function linkSelf() returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYSession) LinkSelf() (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.LinkSelf(&_VOTEPROXYFACTORY.TransactOpts)
}

// LinkSelf is a paid mutator transaction binding the contract method 0x09a14f2e.
//
// Solidity: function linkSelf() returns(address voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYTransactorSession) LinkSelf() (*types.Transaction, error) {
	return _VOTEPROXYFACTORY.Contract.LinkSelf(&_VOTEPROXYFACTORY.TransactOpts)
}

// VOTEPROXYFACTORYLinkConfirmedIterator is returned from FilterLinkConfirmed and is used to iterate over the raw logs and unpacked data for LinkConfirmed events raised by the VOTEPROXYFACTORY contract.
type VOTEPROXYFACTORYLinkConfirmedIterator struct {
	Event *VOTEPROXYFACTORYLinkConfirmed // Event containing the contract specifics and raw log

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
func (it *VOTEPROXYFACTORYLinkConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VOTEPROXYFACTORYLinkConfirmed)
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
		it.Event = new(VOTEPROXYFACTORYLinkConfirmed)
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
func (it *VOTEPROXYFACTORYLinkConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VOTEPROXYFACTORYLinkConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VOTEPROXYFACTORYLinkConfirmed represents a LinkConfirmed event raised by the VOTEPROXYFACTORY contract.
type VOTEPROXYFACTORYLinkConfirmed struct {
	Cold      common.Address
	Hot       common.Address
	VoteProxy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLinkConfirmed is a free log retrieval operation binding the contract event 0xf001c2d12c2288935c811b4977748cb3e5e3c485d08a1fb1984023cb2452d463.
//
// Solidity: event LinkConfirmed(address indexed cold, address indexed hot, address indexed voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) FilterLinkConfirmed(opts *bind.FilterOpts, cold []common.Address, hot []common.Address, voteProxy []common.Address) (*VOTEPROXYFACTORYLinkConfirmedIterator, error) {

	var coldRule []interface{}
	for _, coldItem := range cold {
		coldRule = append(coldRule, coldItem)
	}
	var hotRule []interface{}
	for _, hotItem := range hot {
		hotRule = append(hotRule, hotItem)
	}
	var voteProxyRule []interface{}
	for _, voteProxyItem := range voteProxy {
		voteProxyRule = append(voteProxyRule, voteProxyItem)
	}

	logs, sub, err := _VOTEPROXYFACTORY.contract.FilterLogs(opts, "LinkConfirmed", coldRule, hotRule, voteProxyRule)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORYLinkConfirmedIterator{contract: _VOTEPROXYFACTORY.contract, event: "LinkConfirmed", logs: logs, sub: sub}, nil
}

// WatchLinkConfirmed is a free log subscription operation binding the contract event 0xf001c2d12c2288935c811b4977748cb3e5e3c485d08a1fb1984023cb2452d463.
//
// Solidity: event LinkConfirmed(address indexed cold, address indexed hot, address indexed voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) WatchLinkConfirmed(opts *bind.WatchOpts, sink chan<- *VOTEPROXYFACTORYLinkConfirmed, cold []common.Address, hot []common.Address, voteProxy []common.Address) (event.Subscription, error) {

	var coldRule []interface{}
	for _, coldItem := range cold {
		coldRule = append(coldRule, coldItem)
	}
	var hotRule []interface{}
	for _, hotItem := range hot {
		hotRule = append(hotRule, hotItem)
	}
	var voteProxyRule []interface{}
	for _, voteProxyItem := range voteProxy {
		voteProxyRule = append(voteProxyRule, voteProxyItem)
	}

	logs, sub, err := _VOTEPROXYFACTORY.contract.WatchLogs(opts, "LinkConfirmed", coldRule, hotRule, voteProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VOTEPROXYFACTORYLinkConfirmed)
				if err := _VOTEPROXYFACTORY.contract.UnpackLog(event, "LinkConfirmed", log); err != nil {
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

// ParseLinkConfirmed is a log parse operation binding the contract event 0xf001c2d12c2288935c811b4977748cb3e5e3c485d08a1fb1984023cb2452d463.
//
// Solidity: event LinkConfirmed(address indexed cold, address indexed hot, address indexed voteProxy)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) ParseLinkConfirmed(log types.Log) (*VOTEPROXYFACTORYLinkConfirmed, error) {
	event := new(VOTEPROXYFACTORYLinkConfirmed)
	if err := _VOTEPROXYFACTORY.contract.UnpackLog(event, "LinkConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VOTEPROXYFACTORYLinkRequestedIterator is returned from FilterLinkRequested and is used to iterate over the raw logs and unpacked data for LinkRequested events raised by the VOTEPROXYFACTORY contract.
type VOTEPROXYFACTORYLinkRequestedIterator struct {
	Event *VOTEPROXYFACTORYLinkRequested // Event containing the contract specifics and raw log

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
func (it *VOTEPROXYFACTORYLinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VOTEPROXYFACTORYLinkRequested)
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
		it.Event = new(VOTEPROXYFACTORYLinkRequested)
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
func (it *VOTEPROXYFACTORYLinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VOTEPROXYFACTORYLinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VOTEPROXYFACTORYLinkRequested represents a LinkRequested event raised by the VOTEPROXYFACTORY contract.
type VOTEPROXYFACTORYLinkRequested struct {
	Cold common.Address
	Hot  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLinkRequested is a free log retrieval operation binding the contract event 0x508cad8883b1f632644173100ee83c013494290baf3ace13ec40588d1aed4505.
//
// Solidity: event LinkRequested(address indexed cold, address indexed hot)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) FilterLinkRequested(opts *bind.FilterOpts, cold []common.Address, hot []common.Address) (*VOTEPROXYFACTORYLinkRequestedIterator, error) {

	var coldRule []interface{}
	for _, coldItem := range cold {
		coldRule = append(coldRule, coldItem)
	}
	var hotRule []interface{}
	for _, hotItem := range hot {
		hotRule = append(hotRule, hotItem)
	}

	logs, sub, err := _VOTEPROXYFACTORY.contract.FilterLogs(opts, "LinkRequested", coldRule, hotRule)
	if err != nil {
		return nil, err
	}
	return &VOTEPROXYFACTORYLinkRequestedIterator{contract: _VOTEPROXYFACTORY.contract, event: "LinkRequested", logs: logs, sub: sub}, nil
}

// WatchLinkRequested is a free log subscription operation binding the contract event 0x508cad8883b1f632644173100ee83c013494290baf3ace13ec40588d1aed4505.
//
// Solidity: event LinkRequested(address indexed cold, address indexed hot)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) WatchLinkRequested(opts *bind.WatchOpts, sink chan<- *VOTEPROXYFACTORYLinkRequested, cold []common.Address, hot []common.Address) (event.Subscription, error) {

	var coldRule []interface{}
	for _, coldItem := range cold {
		coldRule = append(coldRule, coldItem)
	}
	var hotRule []interface{}
	for _, hotItem := range hot {
		hotRule = append(hotRule, hotItem)
	}

	logs, sub, err := _VOTEPROXYFACTORY.contract.WatchLogs(opts, "LinkRequested", coldRule, hotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VOTEPROXYFACTORYLinkRequested)
				if err := _VOTEPROXYFACTORY.contract.UnpackLog(event, "LinkRequested", log); err != nil {
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

// ParseLinkRequested is a log parse operation binding the contract event 0x508cad8883b1f632644173100ee83c013494290baf3ace13ec40588d1aed4505.
//
// Solidity: event LinkRequested(address indexed cold, address indexed hot)
func (_VOTEPROXYFACTORY *VOTEPROXYFACTORYFilterer) ParseLinkRequested(log types.Log) (*VOTEPROXYFACTORYLinkRequested, error) {
	event := new(VOTEPROXYFACTORYLinkRequested)
	if err := _VOTEPROXYFACTORY.contract.UnpackLog(event, "LinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
