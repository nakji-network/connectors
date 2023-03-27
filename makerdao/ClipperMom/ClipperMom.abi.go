// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ClipperMom

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

// CLIPPERMOMABI is the input ABI used to generate the binding from.
const CLIPPERMOMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spotter_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAuthority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"SetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"SetBreaker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPriceTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotter\",\"outputs\":[{\"internalType\":\"contractSpotterLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"}],\"name\":\"tripBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CLIPPERMOM is an auto generated Go binding around an Ethereum contract.
type CLIPPERMOM struct {
	CLIPPERMOMCaller     // Read-only binding to the contract
	CLIPPERMOMTransactor // Write-only binding to the contract
	CLIPPERMOMFilterer   // Log filterer for contract events
}

// CLIPPERMOMCaller is an auto generated read-only Go binding around an Ethereum contract.
type CLIPPERMOMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPPERMOMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CLIPPERMOMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPPERMOMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CLIPPERMOMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPPERMOMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CLIPPERMOMSession struct {
	Contract     *CLIPPERMOM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CLIPPERMOMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CLIPPERMOMCallerSession struct {
	Contract *CLIPPERMOMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CLIPPERMOMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CLIPPERMOMTransactorSession struct {
	Contract     *CLIPPERMOMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CLIPPERMOMRaw is an auto generated low-level Go binding around an Ethereum contract.
type CLIPPERMOMRaw struct {
	Contract *CLIPPERMOM // Generic contract binding to access the raw methods on
}

// CLIPPERMOMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CLIPPERMOMCallerRaw struct {
	Contract *CLIPPERMOMCaller // Generic read-only contract binding to access the raw methods on
}

// CLIPPERMOMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CLIPPERMOMTransactorRaw struct {
	Contract *CLIPPERMOMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCLIPPERMOM creates a new instance of CLIPPERMOM, bound to a specific deployed contract.
func NewCLIPPERMOM(address common.Address, backend bind.ContractBackend) (*CLIPPERMOM, error) {
	contract, err := bindCLIPPERMOM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOM{CLIPPERMOMCaller: CLIPPERMOMCaller{contract: contract}, CLIPPERMOMTransactor: CLIPPERMOMTransactor{contract: contract}, CLIPPERMOMFilterer: CLIPPERMOMFilterer{contract: contract}}, nil
}

// NewCLIPPERMOMCaller creates a new read-only instance of CLIPPERMOM, bound to a specific deployed contract.
func NewCLIPPERMOMCaller(address common.Address, caller bind.ContractCaller) (*CLIPPERMOMCaller, error) {
	contract, err := bindCLIPPERMOM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMCaller{contract: contract}, nil
}

// NewCLIPPERMOMTransactor creates a new write-only instance of CLIPPERMOM, bound to a specific deployed contract.
func NewCLIPPERMOMTransactor(address common.Address, transactor bind.ContractTransactor) (*CLIPPERMOMTransactor, error) {
	contract, err := bindCLIPPERMOM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMTransactor{contract: contract}, nil
}

// NewCLIPPERMOMFilterer creates a new log filterer instance of CLIPPERMOM, bound to a specific deployed contract.
func NewCLIPPERMOMFilterer(address common.Address, filterer bind.ContractFilterer) (*CLIPPERMOMFilterer, error) {
	contract, err := bindCLIPPERMOM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMFilterer{contract: contract}, nil
}

// bindCLIPPERMOM binds a generic wrapper to an already deployed contract.
func bindCLIPPERMOM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CLIPPERMOMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLIPPERMOM *CLIPPERMOMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLIPPERMOM.Contract.CLIPPERMOMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLIPPERMOM *CLIPPERMOMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.CLIPPERMOMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLIPPERMOM *CLIPPERMOMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.CLIPPERMOMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLIPPERMOM *CLIPPERMOMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLIPPERMOM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLIPPERMOM *CLIPPERMOMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLIPPERMOM *CLIPPERMOMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLIPPERMOM.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMSession) Authority() (common.Address, error) {
	return _CLIPPERMOM.Contract.Authority(&_CLIPPERMOM.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCallerSession) Authority() (common.Address, error) {
	return _CLIPPERMOM.Contract.Authority(&_CLIPPERMOM.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CLIPPERMOM.contract.Call(opts, &out, "locked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMSession) Locked(arg0 common.Address) (*big.Int, error) {
	return _CLIPPERMOM.Contract.Locked(&_CLIPPERMOM.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMCallerSession) Locked(arg0 common.Address) (*big.Int, error) {
	return _CLIPPERMOM.Contract.Locked(&_CLIPPERMOM.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLIPPERMOM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMSession) Owner() (common.Address, error) {
	return _CLIPPERMOM.Contract.Owner(&_CLIPPERMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCallerSession) Owner() (common.Address, error) {
	return _CLIPPERMOM.Contract.Owner(&_CLIPPERMOM.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCaller) Spotter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CLIPPERMOM.contract.Call(opts, &out, "spotter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMSession) Spotter() (common.Address, error) {
	return _CLIPPERMOM.Contract.Spotter(&_CLIPPERMOM.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_CLIPPERMOM *CLIPPERMOMCallerSession) Spotter() (common.Address, error) {
	return _CLIPPERMOM.Contract.Spotter(&_CLIPPERMOM.CallOpts)
}

// Tolerance is a free data retrieval call binding the contract method 0xa70885c1.
//
// Solidity: function tolerance(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMCaller) Tolerance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CLIPPERMOM.contract.Call(opts, &out, "tolerance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tolerance is a free data retrieval call binding the contract method 0xa70885c1.
//
// Solidity: function tolerance(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMSession) Tolerance(arg0 common.Address) (*big.Int, error) {
	return _CLIPPERMOM.Contract.Tolerance(&_CLIPPERMOM.CallOpts, arg0)
}

// Tolerance is a free data retrieval call binding the contract method 0xa70885c1.
//
// Solidity: function tolerance(address ) view returns(uint256)
func (_CLIPPERMOM *CLIPPERMOMCallerSession) Tolerance(arg0 common.Address) (*big.Int, error) {
	return _CLIPPERMOM.Contract.Tolerance(&_CLIPPERMOM.CallOpts, arg0)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_CLIPPERMOM *CLIPPERMOMSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetAuthority(&_CLIPPERMOM.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetAuthority(&_CLIPPERMOM.TransactOpts, authority_)
}

// SetBreaker is a paid mutator transaction binding the contract method 0x2145a7f3.
//
// Solidity: function setBreaker(address clip, uint256 level, uint256 delay) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactor) SetBreaker(opts *bind.TransactOpts, clip common.Address, level *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.contract.Transact(opts, "setBreaker", clip, level, delay)
}

// SetBreaker is a paid mutator transaction binding the contract method 0x2145a7f3.
//
// Solidity: function setBreaker(address clip, uint256 level, uint256 delay) returns()
func (_CLIPPERMOM *CLIPPERMOMSession) SetBreaker(clip common.Address, level *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetBreaker(&_CLIPPERMOM.TransactOpts, clip, level, delay)
}

// SetBreaker is a paid mutator transaction binding the contract method 0x2145a7f3.
//
// Solidity: function setBreaker(address clip, uint256 level, uint256 delay) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactorSession) SetBreaker(clip common.Address, level *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetBreaker(&_CLIPPERMOM.TransactOpts, clip, level, delay)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_CLIPPERMOM *CLIPPERMOMSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetOwner(&_CLIPPERMOM.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetOwner(&_CLIPPERMOM.TransactOpts, owner_)
}

// SetPriceTolerance is a paid mutator transaction binding the contract method 0xbf90f418.
//
// Solidity: function setPriceTolerance(address clip, uint256 value) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactor) SetPriceTolerance(opts *bind.TransactOpts, clip common.Address, value *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.contract.Transact(opts, "setPriceTolerance", clip, value)
}

// SetPriceTolerance is a paid mutator transaction binding the contract method 0xbf90f418.
//
// Solidity: function setPriceTolerance(address clip, uint256 value) returns()
func (_CLIPPERMOM *CLIPPERMOMSession) SetPriceTolerance(clip common.Address, value *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetPriceTolerance(&_CLIPPERMOM.TransactOpts, clip, value)
}

// SetPriceTolerance is a paid mutator transaction binding the contract method 0xbf90f418.
//
// Solidity: function setPriceTolerance(address clip, uint256 value) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactorSession) SetPriceTolerance(clip common.Address, value *big.Int) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.SetPriceTolerance(&_CLIPPERMOM.TransactOpts, clip, value)
}

// TripBreaker is a paid mutator transaction binding the contract method 0xaa85b638.
//
// Solidity: function tripBreaker(address clip) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactor) TripBreaker(opts *bind.TransactOpts, clip common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.contract.Transact(opts, "tripBreaker", clip)
}

// TripBreaker is a paid mutator transaction binding the contract method 0xaa85b638.
//
// Solidity: function tripBreaker(address clip) returns()
func (_CLIPPERMOM *CLIPPERMOMSession) TripBreaker(clip common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.TripBreaker(&_CLIPPERMOM.TransactOpts, clip)
}

// TripBreaker is a paid mutator transaction binding the contract method 0xaa85b638.
//
// Solidity: function tripBreaker(address clip) returns()
func (_CLIPPERMOM *CLIPPERMOMTransactorSession) TripBreaker(clip common.Address) (*types.Transaction, error) {
	return _CLIPPERMOM.Contract.TripBreaker(&_CLIPPERMOM.TransactOpts, clip)
}

// CLIPPERMOMSetAuthorityIterator is returned from FilterSetAuthority and is used to iterate over the raw logs and unpacked data for SetAuthority events raised by the CLIPPERMOM contract.
type CLIPPERMOMSetAuthorityIterator struct {
	Event *CLIPPERMOMSetAuthority // Event containing the contract specifics and raw log

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
func (it *CLIPPERMOMSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLIPPERMOMSetAuthority)
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
		it.Event = new(CLIPPERMOMSetAuthority)
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
func (it *CLIPPERMOMSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CLIPPERMOMSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CLIPPERMOMSetAuthority represents a SetAuthority event raised by the CLIPPERMOM contract.
type CLIPPERMOMSetAuthority struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAuthority is a free log retrieval operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address indexed oldAuthority, address indexed newAuthority)
func (_CLIPPERMOM *CLIPPERMOMFilterer) FilterSetAuthority(opts *bind.FilterOpts, oldAuthority []common.Address, newAuthority []common.Address) (*CLIPPERMOMSetAuthorityIterator, error) {

	var oldAuthorityRule []interface{}
	for _, oldAuthorityItem := range oldAuthority {
		oldAuthorityRule = append(oldAuthorityRule, oldAuthorityItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.FilterLogs(opts, "SetAuthority", oldAuthorityRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMSetAuthorityIterator{contract: _CLIPPERMOM.contract, event: "SetAuthority", logs: logs, sub: sub}, nil
}

// WatchSetAuthority is a free log subscription operation binding the contract event 0x2ee176e1c4a74440b01c211156e3e4de7b868cd621c5abfe477e01918e7ef671.
//
// Solidity: event SetAuthority(address indexed oldAuthority, address indexed newAuthority)
func (_CLIPPERMOM *CLIPPERMOMFilterer) WatchSetAuthority(opts *bind.WatchOpts, sink chan<- *CLIPPERMOMSetAuthority, oldAuthority []common.Address, newAuthority []common.Address) (event.Subscription, error) {

	var oldAuthorityRule []interface{}
	for _, oldAuthorityItem := range oldAuthority {
		oldAuthorityRule = append(oldAuthorityRule, oldAuthorityItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.WatchLogs(opts, "SetAuthority", oldAuthorityRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CLIPPERMOMSetAuthority)
				if err := _CLIPPERMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
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
func (_CLIPPERMOM *CLIPPERMOMFilterer) ParseSetAuthority(log types.Log) (*CLIPPERMOMSetAuthority, error) {
	event := new(CLIPPERMOMSetAuthority)
	if err := _CLIPPERMOM.contract.UnpackLog(event, "SetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CLIPPERMOMSetBreakerIterator is returned from FilterSetBreaker and is used to iterate over the raw logs and unpacked data for SetBreaker events raised by the CLIPPERMOM contract.
type CLIPPERMOMSetBreakerIterator struct {
	Event *CLIPPERMOMSetBreaker // Event containing the contract specifics and raw log

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
func (it *CLIPPERMOMSetBreakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLIPPERMOMSetBreaker)
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
		it.Event = new(CLIPPERMOMSetBreaker)
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
func (it *CLIPPERMOMSetBreakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CLIPPERMOMSetBreakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CLIPPERMOMSetBreaker represents a SetBreaker event raised by the CLIPPERMOM contract.
type CLIPPERMOMSetBreaker struct {
	Clip  common.Address
	Level *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBreaker is a free log retrieval operation binding the contract event 0x01a1de957c616c7c99b8f1a1158174e9bd174302795cdbcfbfe0bc760d418c6f.
//
// Solidity: event SetBreaker(address indexed clip, uint256 level)
func (_CLIPPERMOM *CLIPPERMOMFilterer) FilterSetBreaker(opts *bind.FilterOpts, clip []common.Address) (*CLIPPERMOMSetBreakerIterator, error) {

	var clipRule []interface{}
	for _, clipItem := range clip {
		clipRule = append(clipRule, clipItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.FilterLogs(opts, "SetBreaker", clipRule)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMSetBreakerIterator{contract: _CLIPPERMOM.contract, event: "SetBreaker", logs: logs, sub: sub}, nil
}

// WatchSetBreaker is a free log subscription operation binding the contract event 0x01a1de957c616c7c99b8f1a1158174e9bd174302795cdbcfbfe0bc760d418c6f.
//
// Solidity: event SetBreaker(address indexed clip, uint256 level)
func (_CLIPPERMOM *CLIPPERMOMFilterer) WatchSetBreaker(opts *bind.WatchOpts, sink chan<- *CLIPPERMOMSetBreaker, clip []common.Address) (event.Subscription, error) {

	var clipRule []interface{}
	for _, clipItem := range clip {
		clipRule = append(clipRule, clipItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.WatchLogs(opts, "SetBreaker", clipRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CLIPPERMOMSetBreaker)
				if err := _CLIPPERMOM.contract.UnpackLog(event, "SetBreaker", log); err != nil {
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

// ParseSetBreaker is a log parse operation binding the contract event 0x01a1de957c616c7c99b8f1a1158174e9bd174302795cdbcfbfe0bc760d418c6f.
//
// Solidity: event SetBreaker(address indexed clip, uint256 level)
func (_CLIPPERMOM *CLIPPERMOMFilterer) ParseSetBreaker(log types.Log) (*CLIPPERMOMSetBreaker, error) {
	event := new(CLIPPERMOMSetBreaker)
	if err := _CLIPPERMOM.contract.UnpackLog(event, "SetBreaker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CLIPPERMOMSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the CLIPPERMOM contract.
type CLIPPERMOMSetOwnerIterator struct {
	Event *CLIPPERMOMSetOwner // Event containing the contract specifics and raw log

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
func (it *CLIPPERMOMSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CLIPPERMOMSetOwner)
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
		it.Event = new(CLIPPERMOMSetOwner)
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
func (it *CLIPPERMOMSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CLIPPERMOMSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CLIPPERMOMSetOwner represents a SetOwner event raised by the CLIPPERMOM contract.
type CLIPPERMOMSetOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address indexed oldOwner, address indexed newOwner)
func (_CLIPPERMOM *CLIPPERMOMFilterer) FilterSetOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CLIPPERMOMSetOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.FilterLogs(opts, "SetOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CLIPPERMOMSetOwnerIterator{contract: _CLIPPERMOM.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0xcbf985117192c8f614a58aaf97226bb80a754772f5f6edf06f87c675f2e6c663.
//
// Solidity: event SetOwner(address indexed oldOwner, address indexed newOwner)
func (_CLIPPERMOM *CLIPPERMOMFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *CLIPPERMOMSetOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CLIPPERMOM.contract.WatchLogs(opts, "SetOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CLIPPERMOMSetOwner)
				if err := _CLIPPERMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
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
func (_CLIPPERMOM *CLIPPERMOMFilterer) ParseSetOwner(log types.Log) (*CLIPPERMOMSetOwner, error) {
	event := new(CLIPPERMOMSetOwner)
	if err := _CLIPPERMOM.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
