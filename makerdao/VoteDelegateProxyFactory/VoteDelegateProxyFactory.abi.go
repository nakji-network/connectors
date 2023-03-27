// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VoteDelegateProxyFactory

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

// VOTEDELEGATEPROXYFACTORYABI is the input ABI used to generate the binding from.
const VOTEDELEGATEPROXYFACTORYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chief\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polling\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voteDelegate\",\"type\":\"address\"}],\"name\":\"CreateVoteDelegate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"chief\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voteDelegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"isDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polling\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VOTEDELEGATEPROXYFACTORY is an auto generated Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORY struct {
	VOTEDELEGATEPROXYFACTORYCaller     // Read-only binding to the contract
	VOTEDELEGATEPROXYFACTORYTransactor // Write-only binding to the contract
	VOTEDELEGATEPROXYFACTORYFilterer   // Log filterer for contract events
}

// VOTEDELEGATEPROXYFACTORYCaller is an auto generated read-only Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEDELEGATEPROXYFACTORYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEDELEGATEPROXYFACTORYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VOTEDELEGATEPROXYFACTORYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VOTEDELEGATEPROXYFACTORYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VOTEDELEGATEPROXYFACTORYSession struct {
	Contract     *VOTEDELEGATEPROXYFACTORY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VOTEDELEGATEPROXYFACTORYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VOTEDELEGATEPROXYFACTORYCallerSession struct {
	Contract *VOTEDELEGATEPROXYFACTORYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VOTEDELEGATEPROXYFACTORYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VOTEDELEGATEPROXYFACTORYTransactorSession struct {
	Contract     *VOTEDELEGATEPROXYFACTORYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VOTEDELEGATEPROXYFACTORYRaw is an auto generated low-level Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORYRaw struct {
	Contract *VOTEDELEGATEPROXYFACTORY // Generic contract binding to access the raw methods on
}

// VOTEDELEGATEPROXYFACTORYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORYCallerRaw struct {
	Contract *VOTEDELEGATEPROXYFACTORYCaller // Generic read-only contract binding to access the raw methods on
}

// VOTEDELEGATEPROXYFACTORYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VOTEDELEGATEPROXYFACTORYTransactorRaw struct {
	Contract *VOTEDELEGATEPROXYFACTORYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVOTEDELEGATEPROXYFACTORY creates a new instance of VOTEDELEGATEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEDELEGATEPROXYFACTORY(address common.Address, backend bind.ContractBackend) (*VOTEDELEGATEPROXYFACTORY, error) {
	contract, err := bindVOTEDELEGATEPROXYFACTORY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VOTEDELEGATEPROXYFACTORY{VOTEDELEGATEPROXYFACTORYCaller: VOTEDELEGATEPROXYFACTORYCaller{contract: contract}, VOTEDELEGATEPROXYFACTORYTransactor: VOTEDELEGATEPROXYFACTORYTransactor{contract: contract}, VOTEDELEGATEPROXYFACTORYFilterer: VOTEDELEGATEPROXYFACTORYFilterer{contract: contract}}, nil
}

// NewVOTEDELEGATEPROXYFACTORYCaller creates a new read-only instance of VOTEDELEGATEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEDELEGATEPROXYFACTORYCaller(address common.Address, caller bind.ContractCaller) (*VOTEDELEGATEPROXYFACTORYCaller, error) {
	contract, err := bindVOTEDELEGATEPROXYFACTORY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VOTEDELEGATEPROXYFACTORYCaller{contract: contract}, nil
}

// NewVOTEDELEGATEPROXYFACTORYTransactor creates a new write-only instance of VOTEDELEGATEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEDELEGATEPROXYFACTORYTransactor(address common.Address, transactor bind.ContractTransactor) (*VOTEDELEGATEPROXYFACTORYTransactor, error) {
	contract, err := bindVOTEDELEGATEPROXYFACTORY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VOTEDELEGATEPROXYFACTORYTransactor{contract: contract}, nil
}

// NewVOTEDELEGATEPROXYFACTORYFilterer creates a new log filterer instance of VOTEDELEGATEPROXYFACTORY, bound to a specific deployed contract.
func NewVOTEDELEGATEPROXYFACTORYFilterer(address common.Address, filterer bind.ContractFilterer) (*VOTEDELEGATEPROXYFACTORYFilterer, error) {
	contract, err := bindVOTEDELEGATEPROXYFACTORY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VOTEDELEGATEPROXYFACTORYFilterer{contract: contract}, nil
}

// bindVOTEDELEGATEPROXYFACTORY binds a generic wrapper to an already deployed contract.
func bindVOTEDELEGATEPROXYFACTORY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VOTEDELEGATEPROXYFACTORYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VOTEDELEGATEPROXYFACTORY.Contract.VOTEDELEGATEPROXYFACTORYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.VOTEDELEGATEPROXYFACTORYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.VOTEDELEGATEPROXYFACTORYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VOTEDELEGATEPROXYFACTORY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.contract.Transact(opts, method, params...)
}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCaller) Chief(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VOTEDELEGATEPROXYFACTORY.contract.Call(opts, &out, "chief")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYSession) Chief() (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Chief(&_VOTEDELEGATEPROXYFACTORY.CallOpts)
}

// Chief is a free data retrieval call binding the contract method 0xffd864d3.
//
// Solidity: function chief() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCallerSession) Chief() (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Chief(&_VOTEDELEGATEPROXYFACTORY.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCaller) Delegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VOTEDELEGATEPROXYFACTORY.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Delegates(&_VOTEDELEGATEPROXYFACTORY.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCallerSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Delegates(&_VOTEDELEGATEPROXYFACTORY.CallOpts, arg0)
}

// IsDelegate is a free data retrieval call binding the contract method 0x07779627.
//
// Solidity: function isDelegate(address guy) view returns(bool)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCaller) IsDelegate(opts *bind.CallOpts, guy common.Address) (bool, error) {
	var out []interface{}
	err := _VOTEDELEGATEPROXYFACTORY.contract.Call(opts, &out, "isDelegate", guy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegate is a free data retrieval call binding the contract method 0x07779627.
//
// Solidity: function isDelegate(address guy) view returns(bool)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYSession) IsDelegate(guy common.Address) (bool, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.IsDelegate(&_VOTEDELEGATEPROXYFACTORY.CallOpts, guy)
}

// IsDelegate is a free data retrieval call binding the contract method 0x07779627.
//
// Solidity: function isDelegate(address guy) view returns(bool)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCallerSession) IsDelegate(guy common.Address) (bool, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.IsDelegate(&_VOTEDELEGATEPROXYFACTORY.CallOpts, guy)
}

// Polling is a free data retrieval call binding the contract method 0x54717496.
//
// Solidity: function polling() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCaller) Polling(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VOTEDELEGATEPROXYFACTORY.contract.Call(opts, &out, "polling")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Polling is a free data retrieval call binding the contract method 0x54717496.
//
// Solidity: function polling() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYSession) Polling() (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Polling(&_VOTEDELEGATEPROXYFACTORY.CallOpts)
}

// Polling is a free data retrieval call binding the contract method 0x54717496.
//
// Solidity: function polling() view returns(address)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYCallerSession) Polling() (common.Address, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Polling(&_VOTEDELEGATEPROXYFACTORY.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYSession) Create() (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Create(&_VOTEDELEGATEPROXYFACTORY.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYTransactorSession) Create() (*types.Transaction, error) {
	return _VOTEDELEGATEPROXYFACTORY.Contract.Create(&_VOTEDELEGATEPROXYFACTORY.TransactOpts)
}

// VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator is returned from FilterCreateVoteDelegate and is used to iterate over the raw logs and unpacked data for CreateVoteDelegate events raised by the VOTEDELEGATEPROXYFACTORY contract.
type VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator struct {
	Event *VOTEDELEGATEPROXYFACTORYCreateVoteDelegate // Event containing the contract specifics and raw log

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
func (it *VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VOTEDELEGATEPROXYFACTORYCreateVoteDelegate)
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
		it.Event = new(VOTEDELEGATEPROXYFACTORYCreateVoteDelegate)
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
func (it *VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VOTEDELEGATEPROXYFACTORYCreateVoteDelegate represents a CreateVoteDelegate event raised by the VOTEDELEGATEPROXYFACTORY contract.
type VOTEDELEGATEPROXYFACTORYCreateVoteDelegate struct {
	Delegate     common.Address
	VoteDelegate common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateVoteDelegate is a free log retrieval operation binding the contract event 0x2187b96b95fffefab01016c852705bc8ec76d1ea17dd5bffef25fd7136633644.
//
// Solidity: event CreateVoteDelegate(address indexed delegate, address indexed voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYFilterer) FilterCreateVoteDelegate(opts *bind.FilterOpts, delegate []common.Address, voteDelegate []common.Address) (*VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var voteDelegateRule []interface{}
	for _, voteDelegateItem := range voteDelegate {
		voteDelegateRule = append(voteDelegateRule, voteDelegateItem)
	}

	logs, sub, err := _VOTEDELEGATEPROXYFACTORY.contract.FilterLogs(opts, "CreateVoteDelegate", delegateRule, voteDelegateRule)
	if err != nil {
		return nil, err
	}
	return &VOTEDELEGATEPROXYFACTORYCreateVoteDelegateIterator{contract: _VOTEDELEGATEPROXYFACTORY.contract, event: "CreateVoteDelegate", logs: logs, sub: sub}, nil
}

// WatchCreateVoteDelegate is a free log subscription operation binding the contract event 0x2187b96b95fffefab01016c852705bc8ec76d1ea17dd5bffef25fd7136633644.
//
// Solidity: event CreateVoteDelegate(address indexed delegate, address indexed voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYFilterer) WatchCreateVoteDelegate(opts *bind.WatchOpts, sink chan<- *VOTEDELEGATEPROXYFACTORYCreateVoteDelegate, delegate []common.Address, voteDelegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var voteDelegateRule []interface{}
	for _, voteDelegateItem := range voteDelegate {
		voteDelegateRule = append(voteDelegateRule, voteDelegateItem)
	}

	logs, sub, err := _VOTEDELEGATEPROXYFACTORY.contract.WatchLogs(opts, "CreateVoteDelegate", delegateRule, voteDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VOTEDELEGATEPROXYFACTORYCreateVoteDelegate)
				if err := _VOTEDELEGATEPROXYFACTORY.contract.UnpackLog(event, "CreateVoteDelegate", log); err != nil {
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

// ParseCreateVoteDelegate is a log parse operation binding the contract event 0x2187b96b95fffefab01016c852705bc8ec76d1ea17dd5bffef25fd7136633644.
//
// Solidity: event CreateVoteDelegate(address indexed delegate, address indexed voteDelegate)
func (_VOTEDELEGATEPROXYFACTORY *VOTEDELEGATEPROXYFACTORYFilterer) ParseCreateVoteDelegate(log types.Log) (*VOTEDELEGATEPROXYFACTORYCreateVoteDelegate, error) {
	event := new(VOTEDELEGATEPROXYFACTORYCreateVoteDelegate)
	if err := _VOTEDELEGATEPROXYFACTORY.contract.UnpackLog(event, "CreateVoteDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
