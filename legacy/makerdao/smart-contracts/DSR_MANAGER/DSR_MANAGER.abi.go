// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DSR_MANAGER

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

// DSRMANAGERABI is the input ABI used to generate the binding from.
const DSRMANAGERABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pot_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"daiBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"exitAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pieOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pot\",\"outputs\":[{\"internalType\":\"contractPotLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DSRMANAGER is an auto generated Go binding around an Ethereum contract.
type DSRMANAGER struct {
	DSRMANAGERCaller     // Read-only binding to the contract
	DSRMANAGERTransactor // Write-only binding to the contract
	DSRMANAGERFilterer   // Log filterer for contract events
}

// DSRMANAGERCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSRMANAGERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRMANAGERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSRMANAGERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRMANAGERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSRMANAGERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRMANAGERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSRMANAGERSession struct {
	Contract     *DSRMANAGER       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSRMANAGERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSRMANAGERCallerSession struct {
	Contract *DSRMANAGERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DSRMANAGERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSRMANAGERTransactorSession struct {
	Contract     *DSRMANAGERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DSRMANAGERRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSRMANAGERRaw struct {
	Contract *DSRMANAGER // Generic contract binding to access the raw methods on
}

// DSRMANAGERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSRMANAGERCallerRaw struct {
	Contract *DSRMANAGERCaller // Generic read-only contract binding to access the raw methods on
}

// DSRMANAGERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSRMANAGERTransactorRaw struct {
	Contract *DSRMANAGERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSRMANAGER creates a new instance of DSRMANAGER, bound to a specific deployed contract.
func NewDSRMANAGER(address common.Address, backend bind.ContractBackend) (*DSRMANAGER, error) {
	contract, err := bindDSRMANAGER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGER{DSRMANAGERCaller: DSRMANAGERCaller{contract: contract}, DSRMANAGERTransactor: DSRMANAGERTransactor{contract: contract}, DSRMANAGERFilterer: DSRMANAGERFilterer{contract: contract}}, nil
}

// NewDSRMANAGERCaller creates a new read-only instance of DSRMANAGER, bound to a specific deployed contract.
func NewDSRMANAGERCaller(address common.Address, caller bind.ContractCaller) (*DSRMANAGERCaller, error) {
	contract, err := bindDSRMANAGER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGERCaller{contract: contract}, nil
}

// NewDSRMANAGERTransactor creates a new write-only instance of DSRMANAGER, bound to a specific deployed contract.
func NewDSRMANAGERTransactor(address common.Address, transactor bind.ContractTransactor) (*DSRMANAGERTransactor, error) {
	contract, err := bindDSRMANAGER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGERTransactor{contract: contract}, nil
}

// NewDSRMANAGERFilterer creates a new log filterer instance of DSRMANAGER, bound to a specific deployed contract.
func NewDSRMANAGERFilterer(address common.Address, filterer bind.ContractFilterer) (*DSRMANAGERFilterer, error) {
	contract, err := bindDSRMANAGER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGERFilterer{contract: contract}, nil
}

// bindDSRMANAGER binds a generic wrapper to an already deployed contract.
func bindDSRMANAGER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSRMANAGERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSRMANAGER *DSRMANAGERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSRMANAGER.Contract.DSRMANAGERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSRMANAGER *DSRMANAGERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.DSRMANAGERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSRMANAGER *DSRMANAGERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.DSRMANAGERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSRMANAGER *DSRMANAGERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSRMANAGER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSRMANAGER *DSRMANAGERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSRMANAGER *DSRMANAGERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DSRMANAGER *DSRMANAGERCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DSRMANAGER.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DSRMANAGER *DSRMANAGERSession) Dai() (common.Address, error) {
	return _DSRMANAGER.Contract.Dai(&_DSRMANAGER.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DSRMANAGER *DSRMANAGERCallerSession) Dai() (common.Address, error) {
	return _DSRMANAGER.Contract.Dai(&_DSRMANAGER.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DSRMANAGER *DSRMANAGERCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DSRMANAGER.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DSRMANAGER *DSRMANAGERSession) DaiJoin() (common.Address, error) {
	return _DSRMANAGER.Contract.DaiJoin(&_DSRMANAGER.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DSRMANAGER *DSRMANAGERCallerSession) DaiJoin() (common.Address, error) {
	return _DSRMANAGER.Contract.DaiJoin(&_DSRMANAGER.CallOpts)
}

// PieOf is a free data retrieval call binding the contract method 0x88787f2b.
//
// Solidity: function pieOf(address ) view returns(uint256)
func (_DSRMANAGER *DSRMANAGERCaller) PieOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSRMANAGER.contract.Call(opts, &out, "pieOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PieOf is a free data retrieval call binding the contract method 0x88787f2b.
//
// Solidity: function pieOf(address ) view returns(uint256)
func (_DSRMANAGER *DSRMANAGERSession) PieOf(arg0 common.Address) (*big.Int, error) {
	return _DSRMANAGER.Contract.PieOf(&_DSRMANAGER.CallOpts, arg0)
}

// PieOf is a free data retrieval call binding the contract method 0x88787f2b.
//
// Solidity: function pieOf(address ) view returns(uint256)
func (_DSRMANAGER *DSRMANAGERCallerSession) PieOf(arg0 common.Address) (*big.Int, error) {
	return _DSRMANAGER.Contract.PieOf(&_DSRMANAGER.CallOpts, arg0)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_DSRMANAGER *DSRMANAGERCaller) Pot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DSRMANAGER.contract.Call(opts, &out, "pot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_DSRMANAGER *DSRMANAGERSession) Pot() (common.Address, error) {
	return _DSRMANAGER.Contract.Pot(&_DSRMANAGER.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_DSRMANAGER *DSRMANAGERCallerSession) Pot() (common.Address, error) {
	return _DSRMANAGER.Contract.Pot(&_DSRMANAGER.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_DSRMANAGER *DSRMANAGERCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRMANAGER.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_DSRMANAGER *DSRMANAGERSession) Supply() (*big.Int, error) {
	return _DSRMANAGER.Contract.Supply(&_DSRMANAGER.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_DSRMANAGER *DSRMANAGERCallerSession) Supply() (*big.Int, error) {
	return _DSRMANAGER.Contract.Supply(&_DSRMANAGER.CallOpts)
}

// DaiBalance is a paid mutator transaction binding the contract method 0xd7f7098f.
//
// Solidity: function daiBalance(address usr) returns(uint256 wad)
func (_DSRMANAGER *DSRMANAGERTransactor) DaiBalance(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.contract.Transact(opts, "daiBalance", usr)
}

// DaiBalance is a paid mutator transaction binding the contract method 0xd7f7098f.
//
// Solidity: function daiBalance(address usr) returns(uint256 wad)
func (_DSRMANAGER *DSRMANAGERSession) DaiBalance(usr common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.DaiBalance(&_DSRMANAGER.TransactOpts, usr)
}

// DaiBalance is a paid mutator transaction binding the contract method 0xd7f7098f.
//
// Solidity: function daiBalance(address usr) returns(uint256 wad)
func (_DSRMANAGER *DSRMANAGERTransactorSession) DaiBalance(usr common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.DaiBalance(&_DSRMANAGER.TransactOpts, usr)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERTransactor) Exit(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.contract.Transact(opts, "exit", dst, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERSession) Exit(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.Exit(&_DSRMANAGER.TransactOpts, dst, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERTransactorSession) Exit(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.Exit(&_DSRMANAGER.TransactOpts, dst, wad)
}

// ExitAll is a paid mutator transaction binding the contract method 0xeb0dff66.
//
// Solidity: function exitAll(address dst) returns()
func (_DSRMANAGER *DSRMANAGERTransactor) ExitAll(opts *bind.TransactOpts, dst common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.contract.Transact(opts, "exitAll", dst)
}

// ExitAll is a paid mutator transaction binding the contract method 0xeb0dff66.
//
// Solidity: function exitAll(address dst) returns()
func (_DSRMANAGER *DSRMANAGERSession) ExitAll(dst common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.ExitAll(&_DSRMANAGER.TransactOpts, dst)
}

// ExitAll is a paid mutator transaction binding the contract method 0xeb0dff66.
//
// Solidity: function exitAll(address dst) returns()
func (_DSRMANAGER *DSRMANAGERTransactorSession) ExitAll(dst common.Address) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.ExitAll(&_DSRMANAGER.TransactOpts, dst)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERTransactor) Join(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.contract.Transact(opts, "join", dst, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERSession) Join(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.Join(&_DSRMANAGER.TransactOpts, dst, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address dst, uint256 wad) returns()
func (_DSRMANAGER *DSRMANAGERTransactorSession) Join(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSRMANAGER.Contract.Join(&_DSRMANAGER.TransactOpts, dst, wad)
}

// DSRMANAGERExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the DSRMANAGER contract.
type DSRMANAGERExitIterator struct {
	Event *DSRMANAGERExit // Event containing the contract specifics and raw log

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
func (it *DSRMANAGERExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRMANAGERExit)
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
		it.Event = new(DSRMANAGERExit)
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
func (it *DSRMANAGERExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRMANAGERExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRMANAGERExit represents a Exit event raised by the DSRMANAGER contract.
type DSRMANAGERExit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) FilterExit(opts *bind.FilterOpts, dst []common.Address) (*DSRMANAGERExitIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSRMANAGER.contract.FilterLogs(opts, "Exit", dstRule)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGERExitIterator{contract: _DSRMANAGER.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *DSRMANAGERExit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSRMANAGER.contract.WatchLogs(opts, "Exit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRMANAGERExit)
				if err := _DSRMANAGER.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) ParseExit(log types.Log) (*DSRMANAGERExit, error) {
	event := new(DSRMANAGERExit)
	if err := _DSRMANAGER.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSRMANAGERJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the DSRMANAGER contract.
type DSRMANAGERJoinIterator struct {
	Event *DSRMANAGERJoin // Event containing the contract specifics and raw log

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
func (it *DSRMANAGERJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRMANAGERJoin)
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
		it.Event = new(DSRMANAGERJoin)
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
func (it *DSRMANAGERJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRMANAGERJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRMANAGERJoin represents a Join event raised by the DSRMANAGER contract.
type DSRMANAGERJoin struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) FilterJoin(opts *bind.FilterOpts, dst []common.Address) (*DSRMANAGERJoinIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSRMANAGER.contract.FilterLogs(opts, "Join", dstRule)
	if err != nil {
		return nil, err
	}
	return &DSRMANAGERJoinIterator{contract: _DSRMANAGER.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *DSRMANAGERJoin, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSRMANAGER.contract.WatchLogs(opts, "Join", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRMANAGERJoin)
				if err := _DSRMANAGER.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed dst, uint256 wad)
func (_DSRMANAGER *DSRMANAGERFilterer) ParseJoin(log types.Log) (*DSRMANAGERJoin, error) {
	event := new(DSRMANAGERJoin)
	if err := _DSRMANAGER.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
