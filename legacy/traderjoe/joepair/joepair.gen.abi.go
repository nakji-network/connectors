// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joepair

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20JoeMetaData contains all meta data concerning the IERC20Joe contract.
var IERC20JoeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20JoeABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20JoeMetaData.ABI instead.
var IERC20JoeABI = IERC20JoeMetaData.ABI

// Deprecated: Use IERC20JoeMetaData.Sigs instead.
// IERC20JoeFuncSigs maps the 4-byte function signature to its string representation.
var IERC20JoeFuncSigs = IERC20JoeMetaData.Sigs

// IERC20Joe is an auto generated Go binding around an Ethereum contract.
type IERC20Joe struct {
	IERC20JoeCaller     // Read-only binding to the contract
	IERC20JoeTransactor // Write-only binding to the contract
	IERC20JoeFilterer   // Log filterer for contract events
}

// IERC20JoeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20JoeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20JoeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20JoeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20JoeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20JoeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20JoeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20JoeSession struct {
	Contract     *IERC20Joe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20JoeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20JoeCallerSession struct {
	Contract *IERC20JoeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IERC20JoeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20JoeTransactorSession struct {
	Contract     *IERC20JoeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC20JoeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20JoeRaw struct {
	Contract *IERC20Joe // Generic contract binding to access the raw methods on
}

// IERC20JoeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20JoeCallerRaw struct {
	Contract *IERC20JoeCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20JoeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20JoeTransactorRaw struct {
	Contract *IERC20JoeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Joe creates a new instance of IERC20Joe, bound to a specific deployed contract.
func NewIERC20Joe(address common.Address, backend bind.ContractBackend) (*IERC20Joe, error) {
	contract, err := bindIERC20Joe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Joe{IERC20JoeCaller: IERC20JoeCaller{contract: contract}, IERC20JoeTransactor: IERC20JoeTransactor{contract: contract}, IERC20JoeFilterer: IERC20JoeFilterer{contract: contract}}, nil
}

// NewIERC20JoeCaller creates a new read-only instance of IERC20Joe, bound to a specific deployed contract.
func NewIERC20JoeCaller(address common.Address, caller bind.ContractCaller) (*IERC20JoeCaller, error) {
	contract, err := bindIERC20Joe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20JoeCaller{contract: contract}, nil
}

// NewIERC20JoeTransactor creates a new write-only instance of IERC20Joe, bound to a specific deployed contract.
func NewIERC20JoeTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20JoeTransactor, error) {
	contract, err := bindIERC20Joe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20JoeTransactor{contract: contract}, nil
}

// NewIERC20JoeFilterer creates a new log filterer instance of IERC20Joe, bound to a specific deployed contract.
func NewIERC20JoeFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20JoeFilterer, error) {
	contract, err := bindIERC20Joe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20JoeFilterer{contract: contract}, nil
}

// bindIERC20Joe binds a generic wrapper to an already deployed contract.
func bindIERC20Joe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20JoeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Joe *IERC20JoeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Joe.Contract.IERC20JoeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Joe *IERC20JoeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Joe.Contract.IERC20JoeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Joe *IERC20JoeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Joe.Contract.IERC20JoeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Joe *IERC20JoeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Joe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Joe *IERC20JoeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Joe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Joe *IERC20JoeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Joe.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Joe *IERC20JoeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Joe *IERC20JoeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Joe.Contract.Allowance(&_IERC20Joe.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Joe *IERC20JoeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Joe.Contract.Allowance(&_IERC20Joe.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20Joe *IERC20JoeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20Joe *IERC20JoeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20Joe.Contract.BalanceOf(&_IERC20Joe.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20Joe *IERC20JoeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20Joe.Contract.BalanceOf(&_IERC20Joe.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Joe *IERC20JoeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Joe *IERC20JoeSession) Decimals() (uint8, error) {
	return _IERC20Joe.Contract.Decimals(&_IERC20Joe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Joe *IERC20JoeCallerSession) Decimals() (uint8, error) {
	return _IERC20Joe.Contract.Decimals(&_IERC20Joe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Joe *IERC20JoeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Joe *IERC20JoeSession) Name() (string, error) {
	return _IERC20Joe.Contract.Name(&_IERC20Joe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Joe *IERC20JoeCallerSession) Name() (string, error) {
	return _IERC20Joe.Contract.Name(&_IERC20Joe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Joe *IERC20JoeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Joe *IERC20JoeSession) Symbol() (string, error) {
	return _IERC20Joe.Contract.Symbol(&_IERC20Joe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Joe *IERC20JoeCallerSession) Symbol() (string, error) {
	return _IERC20Joe.Contract.Symbol(&_IERC20Joe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Joe *IERC20JoeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Joe.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Joe *IERC20JoeSession) TotalSupply() (*big.Int, error) {
	return _IERC20Joe.Contract.TotalSupply(&_IERC20Joe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Joe *IERC20JoeCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Joe.Contract.TotalSupply(&_IERC20Joe.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.Approve(&_IERC20Joe.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.Approve(&_IERC20Joe.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.Transfer(&_IERC20Joe.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.Transfer(&_IERC20Joe.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.TransferFrom(&_IERC20Joe.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Joe *IERC20JoeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Joe.Contract.TransferFrom(&_IERC20Joe.TransactOpts, from, to, value)
}

// IERC20JoeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Joe contract.
type IERC20JoeApprovalIterator struct {
	Event *IERC20JoeApproval // Event containing the contract specifics and raw log

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
func (it *IERC20JoeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20JoeApproval)
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
		it.Event = new(IERC20JoeApproval)
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
func (it *IERC20JoeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20JoeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20JoeApproval represents a Approval event raised by the IERC20Joe contract.
type IERC20JoeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20JoeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Joe.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20JoeApprovalIterator{contract: _IERC20Joe.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20JoeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Joe.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20JoeApproval)
				if err := _IERC20Joe.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) ParseApproval(log types.Log) (*IERC20JoeApproval, error) {
	event := new(IERC20JoeApproval)
	if err := _IERC20Joe.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20JoeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Joe contract.
type IERC20JoeTransferIterator struct {
	Event *IERC20JoeTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20JoeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20JoeTransfer)
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
		it.Event = new(IERC20JoeTransfer)
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
func (it *IERC20JoeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20JoeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20JoeTransfer represents a Transfer event raised by the IERC20Joe contract.
type IERC20JoeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20JoeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Joe.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20JoeTransferIterator{contract: _IERC20Joe.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20JoeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Joe.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20JoeTransfer)
				if err := _IERC20Joe.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Joe *IERC20JoeFilterer) ParseTransfer(log types.Log) (*IERC20JoeTransfer, error) {
	event := new(IERC20JoeTransfer)
	if err := _IERC20Joe.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IJoeCalleeMetaData contains all meta data concerning the IJoeCallee contract.
var IJoeCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"joeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ee22dd87": "joeCall(address,uint256,uint256,bytes)",
	},
}

// IJoeCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use IJoeCalleeMetaData.ABI instead.
var IJoeCalleeABI = IJoeCalleeMetaData.ABI

// Deprecated: Use IJoeCalleeMetaData.Sigs instead.
// IJoeCalleeFuncSigs maps the 4-byte function signature to its string representation.
var IJoeCalleeFuncSigs = IJoeCalleeMetaData.Sigs

// IJoeCallee is an auto generated Go binding around an Ethereum contract.
type IJoeCallee struct {
	IJoeCalleeCaller     // Read-only binding to the contract
	IJoeCalleeTransactor // Write-only binding to the contract
	IJoeCalleeFilterer   // Log filterer for contract events
}

// IJoeCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJoeCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJoeCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJoeCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJoeCalleeSession struct {
	Contract     *IJoeCallee       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJoeCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJoeCalleeCallerSession struct {
	Contract *IJoeCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IJoeCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJoeCalleeTransactorSession struct {
	Contract     *IJoeCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IJoeCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJoeCalleeRaw struct {
	Contract *IJoeCallee // Generic contract binding to access the raw methods on
}

// IJoeCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJoeCalleeCallerRaw struct {
	Contract *IJoeCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// IJoeCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJoeCalleeTransactorRaw struct {
	Contract *IJoeCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJoeCallee creates a new instance of IJoeCallee, bound to a specific deployed contract.
func NewIJoeCallee(address common.Address, backend bind.ContractBackend) (*IJoeCallee, error) {
	contract, err := bindIJoeCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJoeCallee{IJoeCalleeCaller: IJoeCalleeCaller{contract: contract}, IJoeCalleeTransactor: IJoeCalleeTransactor{contract: contract}, IJoeCalleeFilterer: IJoeCalleeFilterer{contract: contract}}, nil
}

// NewIJoeCalleeCaller creates a new read-only instance of IJoeCallee, bound to a specific deployed contract.
func NewIJoeCalleeCaller(address common.Address, caller bind.ContractCaller) (*IJoeCalleeCaller, error) {
	contract, err := bindIJoeCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJoeCalleeCaller{contract: contract}, nil
}

// NewIJoeCalleeTransactor creates a new write-only instance of IJoeCallee, bound to a specific deployed contract.
func NewIJoeCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*IJoeCalleeTransactor, error) {
	contract, err := bindIJoeCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJoeCalleeTransactor{contract: contract}, nil
}

// NewIJoeCalleeFilterer creates a new log filterer instance of IJoeCallee, bound to a specific deployed contract.
func NewIJoeCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*IJoeCalleeFilterer, error) {
	contract, err := bindIJoeCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJoeCalleeFilterer{contract: contract}, nil
}

// bindIJoeCallee binds a generic wrapper to an already deployed contract.
func bindIJoeCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IJoeCalleeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJoeCallee *IJoeCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJoeCallee.Contract.IJoeCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJoeCallee *IJoeCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJoeCallee.Contract.IJoeCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJoeCallee *IJoeCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJoeCallee.Contract.IJoeCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJoeCallee *IJoeCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJoeCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJoeCallee *IJoeCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJoeCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJoeCallee *IJoeCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJoeCallee.Contract.contract.Transact(opts, method, params...)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IJoeCallee *IJoeCalleeTransactor) JoeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IJoeCallee.contract.Transact(opts, "joeCall", sender, amount0, amount1, data)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IJoeCallee *IJoeCalleeSession) JoeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IJoeCallee.Contract.JoeCall(&_IJoeCallee.TransactOpts, sender, amount0, amount1, data)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IJoeCallee *IJoeCalleeTransactorSession) JoeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IJoeCallee.Contract.JoeCall(&_IJoeCallee.TransactOpts, sender, amount0, amount1, data)
}

// IJoeFactoryMetaData contains all meta data concerning the IJoeFactory contract.
var IJoeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e3dd18b": "allPairs(uint256)",
		"574f2ba3": "allPairsLength()",
		"c9c65396": "createPair(address,address)",
		"017e7e58": "feeTo()",
		"094b7415": "feeToSetter()",
		"e6a43905": "getPair(address,address)",
		"7cd07e47": "migrator()",
		"f46901ed": "setFeeTo(address)",
		"a2e74af6": "setFeeToSetter(address)",
		"23cf3118": "setMigrator(address)",
	},
}

// IJoeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IJoeFactoryMetaData.ABI instead.
var IJoeFactoryABI = IJoeFactoryMetaData.ABI

// Deprecated: Use IJoeFactoryMetaData.Sigs instead.
// IJoeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IJoeFactoryFuncSigs = IJoeFactoryMetaData.Sigs

// IJoeFactory is an auto generated Go binding around an Ethereum contract.
type IJoeFactory struct {
	IJoeFactoryCaller     // Read-only binding to the contract
	IJoeFactoryTransactor // Write-only binding to the contract
	IJoeFactoryFilterer   // Log filterer for contract events
}

// IJoeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJoeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJoeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJoeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJoeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJoeFactorySession struct {
	Contract     *IJoeFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJoeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJoeFactoryCallerSession struct {
	Contract *IJoeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IJoeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJoeFactoryTransactorSession struct {
	Contract     *IJoeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IJoeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJoeFactoryRaw struct {
	Contract *IJoeFactory // Generic contract binding to access the raw methods on
}

// IJoeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJoeFactoryCallerRaw struct {
	Contract *IJoeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IJoeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJoeFactoryTransactorRaw struct {
	Contract *IJoeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJoeFactory creates a new instance of IJoeFactory, bound to a specific deployed contract.
func NewIJoeFactory(address common.Address, backend bind.ContractBackend) (*IJoeFactory, error) {
	contract, err := bindIJoeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJoeFactory{IJoeFactoryCaller: IJoeFactoryCaller{contract: contract}, IJoeFactoryTransactor: IJoeFactoryTransactor{contract: contract}, IJoeFactoryFilterer: IJoeFactoryFilterer{contract: contract}}, nil
}

// NewIJoeFactoryCaller creates a new read-only instance of IJoeFactory, bound to a specific deployed contract.
func NewIJoeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IJoeFactoryCaller, error) {
	contract, err := bindIJoeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJoeFactoryCaller{contract: contract}, nil
}

// NewIJoeFactoryTransactor creates a new write-only instance of IJoeFactory, bound to a specific deployed contract.
func NewIJoeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IJoeFactoryTransactor, error) {
	contract, err := bindIJoeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJoeFactoryTransactor{contract: contract}, nil
}

// NewIJoeFactoryFilterer creates a new log filterer instance of IJoeFactory, bound to a specific deployed contract.
func NewIJoeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IJoeFactoryFilterer, error) {
	contract, err := bindIJoeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJoeFactoryFilterer{contract: contract}, nil
}

// bindIJoeFactory binds a generic wrapper to an already deployed contract.
func bindIJoeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IJoeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJoeFactory *IJoeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJoeFactory.Contract.IJoeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJoeFactory *IJoeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJoeFactory.Contract.IJoeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJoeFactory *IJoeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJoeFactory.Contract.IJoeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJoeFactory *IJoeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJoeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJoeFactory *IJoeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJoeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJoeFactory *IJoeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJoeFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IJoeFactory *IJoeFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IJoeFactory *IJoeFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IJoeFactory.Contract.AllPairs(&_IJoeFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IJoeFactory *IJoeFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IJoeFactory.Contract.AllPairs(&_IJoeFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IJoeFactory *IJoeFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IJoeFactory *IJoeFactorySession) AllPairsLength() (*big.Int, error) {
	return _IJoeFactory.Contract.AllPairsLength(&_IJoeFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IJoeFactory *IJoeFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IJoeFactory.Contract.AllPairsLength(&_IJoeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IJoeFactory *IJoeFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IJoeFactory *IJoeFactorySession) FeeTo() (common.Address, error) {
	return _IJoeFactory.Contract.FeeTo(&_IJoeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IJoeFactory *IJoeFactoryCallerSession) FeeTo() (common.Address, error) {
	return _IJoeFactory.Contract.FeeTo(&_IJoeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IJoeFactory *IJoeFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IJoeFactory *IJoeFactorySession) FeeToSetter() (common.Address, error) {
	return _IJoeFactory.Contract.FeeToSetter(&_IJoeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IJoeFactory *IJoeFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _IJoeFactory.Contract.FeeToSetter(&_IJoeFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IJoeFactory *IJoeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IJoeFactory *IJoeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IJoeFactory.Contract.GetPair(&_IJoeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IJoeFactory *IJoeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IJoeFactory.Contract.GetPair(&_IJoeFactory.CallOpts, tokenA, tokenB)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IJoeFactory *IJoeFactoryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IJoeFactory.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IJoeFactory *IJoeFactorySession) Migrator() (common.Address, error) {
	return _IJoeFactory.Contract.Migrator(&_IJoeFactory.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_IJoeFactory *IJoeFactoryCallerSession) Migrator() (common.Address, error) {
	return _IJoeFactory.Contract.Migrator(&_IJoeFactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IJoeFactory *IJoeFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IJoeFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IJoeFactory *IJoeFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.CreatePair(&_IJoeFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IJoeFactory *IJoeFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.CreatePair(&_IJoeFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IJoeFactory *IJoeFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetFeeTo(&_IJoeFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetFeeTo(&_IJoeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IJoeFactory *IJoeFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetFeeToSetter(&_IJoeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetFeeToSetter(&_IJoeFactory.TransactOpts, arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactor) SetMigrator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.contract.Transact(opts, "setMigrator", arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_IJoeFactory *IJoeFactorySession) SetMigrator(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetMigrator(&_IJoeFactory.TransactOpts, arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_IJoeFactory *IJoeFactoryTransactorSession) SetMigrator(arg0 common.Address) (*types.Transaction, error) {
	return _IJoeFactory.Contract.SetMigrator(&_IJoeFactory.TransactOpts, arg0)
}

// IJoeFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IJoeFactory contract.
type IJoeFactoryPairCreatedIterator struct {
	Event *IJoeFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IJoeFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IJoeFactoryPairCreated)
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
		it.Event = new(IJoeFactoryPairCreated)
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
func (it *IJoeFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IJoeFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IJoeFactoryPairCreated represents a PairCreated event raised by the IJoeFactory contract.
type IJoeFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IJoeFactory *IJoeFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*IJoeFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IJoeFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &IJoeFactoryPairCreatedIterator{contract: _IJoeFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IJoeFactory *IJoeFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IJoeFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IJoeFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IJoeFactoryPairCreated)
				if err := _IJoeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IJoeFactory *IJoeFactoryFilterer) ParsePairCreated(log types.Log) (*IJoeFactoryPairCreated, error) {
	event := new(IJoeFactoryPairCreated)
	if err := _IJoeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMigratorMetaData contains all meta data concerning the IMigrator contract.
var IMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"desiredLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40dc0e37": "desiredLiquidity()",
	},
}

// IMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use IMigratorMetaData.ABI instead.
var IMigratorABI = IMigratorMetaData.ABI

// Deprecated: Use IMigratorMetaData.Sigs instead.
// IMigratorFuncSigs maps the 4-byte function signature to its string representation.
var IMigratorFuncSigs = IMigratorMetaData.Sigs

// IMigrator is an auto generated Go binding around an Ethereum contract.
type IMigrator struct {
	IMigratorCaller     // Read-only binding to the contract
	IMigratorTransactor // Write-only binding to the contract
	IMigratorFilterer   // Log filterer for contract events
}

// IMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMigratorSession struct {
	Contract     *IMigrator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMigratorCallerSession struct {
	Contract *IMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMigratorTransactorSession struct {
	Contract     *IMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMigratorRaw struct {
	Contract *IMigrator // Generic contract binding to access the raw methods on
}

// IMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMigratorCallerRaw struct {
	Contract *IMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// IMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMigratorTransactorRaw struct {
	Contract *IMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMigrator creates a new instance of IMigrator, bound to a specific deployed contract.
func NewIMigrator(address common.Address, backend bind.ContractBackend) (*IMigrator, error) {
	contract, err := bindIMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMigrator{IMigratorCaller: IMigratorCaller{contract: contract}, IMigratorTransactor: IMigratorTransactor{contract: contract}, IMigratorFilterer: IMigratorFilterer{contract: contract}}, nil
}

// NewIMigratorCaller creates a new read-only instance of IMigrator, bound to a specific deployed contract.
func NewIMigratorCaller(address common.Address, caller bind.ContractCaller) (*IMigratorCaller, error) {
	contract, err := bindIMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMigratorCaller{contract: contract}, nil
}

// NewIMigratorTransactor creates a new write-only instance of IMigrator, bound to a specific deployed contract.
func NewIMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*IMigratorTransactor, error) {
	contract, err := bindIMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMigratorTransactor{contract: contract}, nil
}

// NewIMigratorFilterer creates a new log filterer instance of IMigrator, bound to a specific deployed contract.
func NewIMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*IMigratorFilterer, error) {
	contract, err := bindIMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMigratorFilterer{contract: contract}, nil
}

// bindIMigrator binds a generic wrapper to an already deployed contract.
func bindIMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMigratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMigrator *IMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMigrator.Contract.IMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMigrator *IMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMigrator.Contract.IMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMigrator *IMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMigrator.Contract.IMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMigrator *IMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMigrator *IMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMigrator *IMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMigrator.Contract.contract.Transact(opts, method, params...)
}

// DesiredLiquidity is a free data retrieval call binding the contract method 0x40dc0e37.
//
// Solidity: function desiredLiquidity() view returns(uint256)
func (_IMigrator *IMigratorCaller) DesiredLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMigrator.contract.Call(opts, &out, "desiredLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DesiredLiquidity is a free data retrieval call binding the contract method 0x40dc0e37.
//
// Solidity: function desiredLiquidity() view returns(uint256)
func (_IMigrator *IMigratorSession) DesiredLiquidity() (*big.Int, error) {
	return _IMigrator.Contract.DesiredLiquidity(&_IMigrator.CallOpts)
}

// DesiredLiquidity is a free data retrieval call binding the contract method 0x40dc0e37.
//
// Solidity: function desiredLiquidity() view returns(uint256)
func (_IMigrator *IMigratorCallerSession) DesiredLiquidity() (*big.Int, error) {
	return _IMigrator.Contract.DesiredLiquidity(&_IMigrator.CallOpts)
}

// JoeERC20MetaData contains all meta data concerning the JoeERC20 contract.
var JoeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"30adf81f": "PERMIT_TYPEHASH()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604080518082018252600c81526b2537b2902628102a37b5b2b760a11b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527fd4eb8415c62493390808cc42ba6f3ba97fa30eb773771605674d83c7162115e7818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528351808303909101815260c0909101909252815191012060035561085f806100f36000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80633644e5151161008c57806395d89b411161006657806395d89b411461025b578063a9059cbb14610263578063d505accf1461028f578063dd62ed3e146102e2576100cf565b80633644e5151461020757806370a082311461020f5780637ecebe0014610235576100cf565b806306fdde03146100d4578063095ea7b31461015157806318160ddd1461019157806323b872dd146101ab57806330adf81f146101e1578063313ce567146101e9575b600080fd5b6100dc610310565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101165781810151838201526020016100fe565b50505050905090810190601f1680156101435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017d6004803603604081101561016757600080fd5b506001600160a01b038135169060200135610338565b604080519115158252519081900360200190f35b61019961034f565b60408051918252519081900360200190f35b61017d600480360360608110156101c157600080fd5b506001600160a01b03813581169160208101359091169060400135610355565b6101996103e9565b6101f161040d565b6040805160ff9092168252519081900360200190f35b610199610412565b6101996004803603602081101561022557600080fd5b50356001600160a01b0316610418565b6101996004803603602081101561024b57600080fd5b50356001600160a01b031661042a565b6100dc61043c565b61017d6004803603604081101561027957600080fd5b506001600160a01b03813516906020013561045b565b6102e0600480360360e08110156102a557600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135610468565b005b610199600480360360408110156102f857600080fd5b506001600160a01b038135811691602001351661065d565b6040518060400160405280600c81526020016b2537b2902628102a37b5b2b760a11b81525081565b600061034533848461067a565b5060015b92915050565b60005481565b6001600160a01b0383166000908152600260209081526040808320338452909152812054600019146103d4576001600160a01b03841660009081526002602090815260408083203384529091529020546103af90836106dc565b6001600160a01b03851660009081526002602090815260408083203384529091529020555b6103df84848461072c565b5060019392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60035481565b60016020526000908152604090205481565b60046020526000908152604090205481565b6040518060400160405280600381526020016204a4c560ec1b81525081565b600061034533848461072c565b428410156104ac576040805162461bcd60e51b815260206004820152600c60248201526b129bd94e881156141254915160a21b604482015290519081900360640190fd5b6003546001600160a01b0380891660008181526004602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa1580156105c7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906105fd5750886001600160a01b0316816001600160a01b0316145b610647576040805162461bcd60e51b81526020600482015260166024820152754a6f653a20494e56414c49445f5349474e415455524560501b604482015290519081900360640190fd5b61065289898961067a565b505050505050505050565b600260209081526000928352604080842090915290825290205481565b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b80820382811115610349576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6001600160a01b03831660009081526001602052604090205461074f90826106dc565b6001600160a01b03808516600090815260016020526040808220939093559084168152205461077e90826107da565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b80820182811015610349576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfea2646970667358221220e107a7bb8ab5930ad43890b7ebbaaaf97aa3bab563f7a816b234a4898954a88664736f6c634300060c0033",
}

// JoeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use JoeERC20MetaData.ABI instead.
var JoeERC20ABI = JoeERC20MetaData.ABI

// Deprecated: Use JoeERC20MetaData.Sigs instead.
// JoeERC20FuncSigs maps the 4-byte function signature to its string representation.
var JoeERC20FuncSigs = JoeERC20MetaData.Sigs

// JoeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JoeERC20MetaData.Bin instead.
var JoeERC20Bin = JoeERC20MetaData.Bin

// DeployJoeERC20 deploys a new Ethereum contract, binding an instance of JoeERC20 to it.
func DeployJoeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JoeERC20, error) {
	parsed, err := JoeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JoeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JoeERC20{JoeERC20Caller: JoeERC20Caller{contract: contract}, JoeERC20Transactor: JoeERC20Transactor{contract: contract}, JoeERC20Filterer: JoeERC20Filterer{contract: contract}}, nil
}

// JoeERC20 is an auto generated Go binding around an Ethereum contract.
type JoeERC20 struct {
	JoeERC20Caller     // Read-only binding to the contract
	JoeERC20Transactor // Write-only binding to the contract
	JoeERC20Filterer   // Log filterer for contract events
}

// JoeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type JoeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type JoeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JoeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JoeERC20Session struct {
	Contract     *JoeERC20         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JoeERC20CallerSession struct {
	Contract *JoeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// JoeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JoeERC20TransactorSession struct {
	Contract     *JoeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// JoeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type JoeERC20Raw struct {
	Contract *JoeERC20 // Generic contract binding to access the raw methods on
}

// JoeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JoeERC20CallerRaw struct {
	Contract *JoeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// JoeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JoeERC20TransactorRaw struct {
	Contract *JoeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewJoeERC20 creates a new instance of JoeERC20, bound to a specific deployed contract.
func NewJoeERC20(address common.Address, backend bind.ContractBackend) (*JoeERC20, error) {
	contract, err := bindJoeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JoeERC20{JoeERC20Caller: JoeERC20Caller{contract: contract}, JoeERC20Transactor: JoeERC20Transactor{contract: contract}, JoeERC20Filterer: JoeERC20Filterer{contract: contract}}, nil
}

// NewJoeERC20Caller creates a new read-only instance of JoeERC20, bound to a specific deployed contract.
func NewJoeERC20Caller(address common.Address, caller bind.ContractCaller) (*JoeERC20Caller, error) {
	contract, err := bindJoeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JoeERC20Caller{contract: contract}, nil
}

// NewJoeERC20Transactor creates a new write-only instance of JoeERC20, bound to a specific deployed contract.
func NewJoeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*JoeERC20Transactor, error) {
	contract, err := bindJoeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JoeERC20Transactor{contract: contract}, nil
}

// NewJoeERC20Filterer creates a new log filterer instance of JoeERC20, bound to a specific deployed contract.
func NewJoeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*JoeERC20Filterer, error) {
	contract, err := bindJoeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JoeERC20Filterer{contract: contract}, nil
}

// bindJoeERC20 binds a generic wrapper to an already deployed contract.
func bindJoeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JoeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JoeERC20 *JoeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JoeERC20.Contract.JoeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JoeERC20 *JoeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JoeERC20.Contract.JoeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JoeERC20 *JoeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JoeERC20.Contract.JoeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JoeERC20 *JoeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JoeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JoeERC20 *JoeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JoeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JoeERC20 *JoeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JoeERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoeERC20 *JoeERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoeERC20 *JoeERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _JoeERC20.Contract.DOMAINSEPARATOR(&_JoeERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoeERC20 *JoeERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JoeERC20.Contract.DOMAINSEPARATOR(&_JoeERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoeERC20 *JoeERC20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoeERC20 *JoeERC20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _JoeERC20.Contract.PERMITTYPEHASH(&_JoeERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoeERC20 *JoeERC20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _JoeERC20.Contract.PERMITTYPEHASH(&_JoeERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.Allowance(&_JoeERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoeERC20 *JoeERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.Allowance(&_JoeERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.BalanceOf(&_JoeERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.BalanceOf(&_JoeERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoeERC20 *JoeERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoeERC20 *JoeERC20Session) Decimals() (uint8, error) {
	return _JoeERC20.Contract.Decimals(&_JoeERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoeERC20 *JoeERC20CallerSession) Decimals() (uint8, error) {
	return _JoeERC20.Contract.Decimals(&_JoeERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoeERC20 *JoeERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoeERC20 *JoeERC20Session) Name() (string, error) {
	return _JoeERC20.Contract.Name(&_JoeERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoeERC20 *JoeERC20CallerSession) Name() (string, error) {
	return _JoeERC20.Contract.Name(&_JoeERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.Nonces(&_JoeERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoeERC20 *JoeERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _JoeERC20.Contract.Nonces(&_JoeERC20.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoeERC20 *JoeERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoeERC20 *JoeERC20Session) Symbol() (string, error) {
	return _JoeERC20.Contract.Symbol(&_JoeERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoeERC20 *JoeERC20CallerSession) Symbol() (string, error) {
	return _JoeERC20.Contract.Symbol(&_JoeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoeERC20 *JoeERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoeERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoeERC20 *JoeERC20Session) TotalSupply() (*big.Int, error) {
	return _JoeERC20.Contract.TotalSupply(&_JoeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoeERC20 *JoeERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _JoeERC20.Contract.TotalSupply(&_JoeERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.Approve(&_JoeERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.Approve(&_JoeERC20.TransactOpts, spender, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoeERC20 *JoeERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoeERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoeERC20 *JoeERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoeERC20.Contract.Permit(&_JoeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoeERC20 *JoeERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoeERC20.Contract.Permit(&_JoeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.Transfer(&_JoeERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.Transfer(&_JoeERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.TransferFrom(&_JoeERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoeERC20 *JoeERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoeERC20.Contract.TransferFrom(&_JoeERC20.TransactOpts, from, to, value)
}

// JoeERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the JoeERC20 contract.
type JoeERC20ApprovalIterator struct {
	Event *JoeERC20Approval // Event containing the contract specifics and raw log

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
func (it *JoeERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoeERC20Approval)
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
		it.Event = new(JoeERC20Approval)
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
func (it *JoeERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoeERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoeERC20Approval represents a Approval event raised by the JoeERC20 contract.
type JoeERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JoeERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JoeERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JoeERC20ApprovalIterator{contract: _JoeERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JoeERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JoeERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoeERC20Approval)
				if err := _JoeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) ParseApproval(log types.Log) (*JoeERC20Approval, error) {
	event := new(JoeERC20Approval)
	if err := _JoeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoeERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the JoeERC20 contract.
type JoeERC20TransferIterator struct {
	Event *JoeERC20Transfer // Event containing the contract specifics and raw log

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
func (it *JoeERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoeERC20Transfer)
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
		it.Event = new(JoeERC20Transfer)
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
func (it *JoeERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoeERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoeERC20Transfer represents a Transfer event raised by the JoeERC20 contract.
type JoeERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JoeERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoeERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JoeERC20TransferIterator{contract: _JoeERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JoeERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoeERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoeERC20Transfer)
				if err := _JoeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoeERC20 *JoeERC20Filterer) ParseTransfer(log types.Log) (*JoeERC20Transfer, error) {
	event := new(JoeERC20Transfer)
	if err := _JoeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairMetaData contains all meta data concerning the JoePair contract.
var JoePairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"ba9a7a56": "MINIMUM_LIQUIDITY()",
		"30adf81f": "PERMIT_TYPEHASH()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"89afcb44": "burn(address)",
		"313ce567": "decimals()",
		"c45a0155": "factory()",
		"0902f1ac": "getReserves()",
		"485cc955": "initialize(address,address)",
		"7464fc3d": "kLast()",
		"6a627842": "mint(address)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"5909c0d5": "price0CumulativeLast()",
		"5a3d5493": "price1CumulativeLast()",
		"bc25cf77": "skim(address)",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"95d89b41": "symbol()",
		"fff6cae9": "sync()",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x60806040526001600c5534801561001557600080fd5b50604080518082018252600c81526b2537b2902628102a37b5b2b760a11b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527fd4eb8415c62493390808cc42ba6f3ba97fa30eb773771605674d83c7162115e7818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528351808303909101815260c09091019092528151910120600355600580546001600160a01b031916331790556122f18061010a6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80636a627842116100f9578063ba9a7a5611610097578063d21220a711610071578063d21220a714610534578063d505accf1461053c578063dd62ed3e1461058d578063fff6cae9146105bb576101a9565b8063ba9a7a56146104fe578063bc25cf7714610506578063c45a01551461052c576101a9565b80637ecebe00116100d35780637ecebe001461046557806389afcb441461048b57806395d89b41146104ca578063a9059cbb146104d2576101a9565b80636a6278421461041157806370a08231146104375780637464fc3d1461045d576101a9565b806323b872dd116101665780633644e515116101405780633644e515146103cb578063485cc955146103d35780635909c0d5146104015780635a3d549314610409576101a9565b806323b872dd1461036f57806330adf81f146103a5578063313ce567146103ad576101a9565b8063022c0d9f146101ae57806306fdde031461023c5780630902f1ac146102b9578063095ea7b3146102f15780630dfe16811461033157806318160ddd14610355575b600080fd5b61023a600480360360808110156101c457600080fd5b8135916020810135916001600160a01b0360408301351691908101906080810160608201356401000000008111156101fb57600080fd5b82018360208201111561020d57600080fd5b8035906020019184600183028401116401000000008311171561022f57600080fd5b5090925090506105c3565b005b610244610afb565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561027e578181015183820152602001610266565b50505050905090810190601f1680156102ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c1610b23565b604080516001600160701b03948516815292909316602083015263ffffffff168183015290519081900360600190f35b61031d6004803603604081101561030757600080fd5b506001600160a01b038135169060200135610b4d565b604080519115158252519081900360200190f35b610339610b64565b604080516001600160a01b039092168252519081900360200190f35b61035d610b73565b60408051918252519081900360200190f35b61031d6004803603606081101561038557600080fd5b506001600160a01b03813581169160208101359091169060400135610b79565b61035d610c0d565b6103b5610c31565b6040805160ff9092168252519081900360200190f35b61035d610c36565b61023a600480360360408110156103e957600080fd5b506001600160a01b0381358116916020013516610c3c565b61035d610cba565b61035d610cc0565b61035d6004803603602081101561042757600080fd5b50356001600160a01b0316610cc6565b61035d6004803603602081101561044d57600080fd5b50356001600160a01b031661113c565b61035d61114e565b61035d6004803603602081101561047b57600080fd5b50356001600160a01b0316611154565b6104b1600480360360208110156104a157600080fd5b50356001600160a01b0316611166565b6040805192835260208301919091528051918290030190f35b6102446114f4565b61031d600480360360408110156104e857600080fd5b506001600160a01b038135169060200135611513565b61035d611520565b61023a6004803603602081101561051c57600080fd5b50356001600160a01b0316611526565b610339611692565b6103396116a1565b61023a600480360360e081101561055257600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c001356116b0565b61035d600480360360408110156105a357600080fd5b506001600160a01b03813581169160200135166118a5565b61023a6118c2565b600c54600114610608576040805162461bcd60e51b815260206004820152600b60248201526a129bd94e881313d0d2d15160aa1b604482015290519081900360640190fd5b6000600c558415158061061b5750600084115b61066c576040805162461bcd60e51b815260206004820152601f60248201527f4a6f653a20494e53554646494349454e545f4f55545055545f414d4f554e5400604482015290519081900360640190fd5b600080610677610b23565b5091509150816001600160701b03168710801561069c5750806001600160701b031686105b6106ed576040805162461bcd60e51b815260206004820152601b60248201527f4a6f653a20494e53554646494349454e545f4c49515549444954590000000000604482015290519081900360640190fd5b60065460075460009182916001600160a01b0391821691908116908916821480159061072b5750806001600160a01b0316896001600160a01b031614155b61076e576040805162461bcd60e51b815260206004820152600f60248201526e4a6f653a20494e56414c49445f544f60881b604482015290519081900360640190fd5b8a1561077f5761077f828a8d611a1e565b891561079057610790818a8c611a1e565b861561084257886001600160a01b031663ee22dd87338d8d8c8c6040518663ffffffff1660e01b815260040180866001600160a01b03168152602001858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561082957600080fd5b505af115801561083d573d6000803e3d6000fd5b505050505b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561088857600080fd5b505afa15801561089c573d6000803e3d6000fd5b505050506040513d60208110156108b257600080fd5b5051604080516370a0823160e01b815230600482015290519195506001600160a01b038316916370a0823191602480820192602092909190829003018186803b1580156108fe57600080fd5b505afa158015610912573d6000803e3d6000fd5b505050506040513d602081101561092857600080fd5b5051925060009150506001600160701b0385168a9003831161094b57600061095a565b89856001600160701b03160383035b9050600089856001600160701b0316038311610977576000610986565b89856001600160701b03160383035b905060008211806109975750600081115b6109e8576040805162461bcd60e51b815260206004820152601e60248201527f4a6f653a20494e53554646494349454e545f494e5055545f414d4f554e540000604482015290519081900360640190fd5b6000610a0a6109f8846003611baf565b610a04876103e8611baf565b90611c12565b90506000610a1c6109f8846003611baf565b9050610a41620f4240610a3b6001600160701b038b8116908b16611baf565b90611baf565b610a4b8383611baf565b1015610a87576040805162461bcd60e51b81526020600482015260066024820152654a6f653a204b60d01b604482015290519081900360640190fd5b5050610a9584848888611c62565b60408051838152602081018390528082018d9052606081018c905290516001600160a01b038b169133917fd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d8229181900360800190a350506001600c55505050505050505050565b6040518060400160405280600c81526020016b2537b2902628102a37b5b2b760a11b81525081565b6008546001600160701b0380821692600160701b830490911691600160e01b900463ffffffff1690565b6000610b5a338484611e1b565b5060015b92915050565b6006546001600160a01b031681565b60005481565b6001600160a01b038316600090815260026020908152604080832033845290915281205460001914610bf8576001600160a01b0384166000908152600260209081526040808320338452909152902054610bd39083611c12565b6001600160a01b03851660009081526002602090815260408083203384529091529020555b610c03848484611e7d565b5060019392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60035481565b6005546001600160a01b03163314610c8c576040805162461bcd60e51b815260206004820152600e60248201526d2537b29d102327a92124a22222a760911b604482015290519081900360640190fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055565b60095481565b600a5481565b6000600c54600114610d0d576040805162461bcd60e51b815260206004820152600b60248201526a129bd94e881313d0d2d15160aa1b604482015290519081900360640190fd5b6000600c81905580610d1d610b23565b50600654604080516370a0823160e01b815230600482015290519395509193506000926001600160a01b03909116916370a08231916024808301926020929190829003018186803b158015610d7157600080fd5b505afa158015610d85573d6000803e3d6000fd5b505050506040513d6020811015610d9b57600080fd5b5051600754604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b158015610dee57600080fd5b505afa158015610e02573d6000803e3d6000fd5b505050506040513d6020811015610e1857600080fd5b505190506000610e31836001600160701b038716611c12565b90506000610e48836001600160701b038716611c12565b90506000610e568787611f2b565b6000549091508061102d5760055460408051637cd07e4760e01b815290516000926001600160a01b031691637cd07e47916004808301926020929190829003018186803b158015610ea657600080fd5b505afa158015610eba573d6000803e3d6000fd5b505050506040513d6020811015610ed057600080fd5b50519050336001600160a01b0382161415610fab57806001600160a01b03166340dc0e376040518163ffffffff1660e01b815260040160206040518083038186803b158015610f1e57600080fd5b505afa158015610f32573d6000803e3d6000fd5b505050506040513d6020811015610f4857600080fd5b505199508915801590610f5d57506000198a14155b610fa6576040805162461bcd60e51b81526020600482015260156024820152744261642064657369726564206c697175696469747960581b604482015290519081900360640190fd5b611027565b6001600160a01b03811615611000576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103737ba103430bb329036b4b3b930ba37b960511b604482015290519081900360640190fd5b6110186103e8610a046110138888611baf565b61206b565b995061102760006103e86120bd565b50611070565b61106d6001600160701b0389166110448684611baf565b8161104b57fe5b046001600160701b0389166110608685611baf565b8161106757fe5b04612147565b98505b600089116110af5760405162461bcd60e51b815260040180806020018281038252602281526020018061229a6022913960400191505060405180910390fd5b6110b98a8a6120bd565b6110c586868a8a611c62565b81156110ef576008546110eb906001600160701b0380821691600160701b900416611baf565b600b555b6040805185815260208101859052815133927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f928290030190a250506001600c5550949695505050505050565b60016020526000908152604090205481565b600b5481565b60046020526000908152604090205481565b600080600c546001146111ae576040805162461bcd60e51b815260206004820152600b60248201526a129bd94e881313d0d2d15160aa1b604482015290519081900360640190fd5b6000600c819055806111be610b23565b50600654600754604080516370a0823160e01b815230600482015290519496509294506001600160a01b039182169391169160009184916370a08231916024808301926020929190829003018186803b15801561121a57600080fd5b505afa15801561122e573d6000803e3d6000fd5b505050506040513d602081101561124457600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b038516916370a08231916024808301926020929190829003018186803b15801561129257600080fd5b505afa1580156112a6573d6000803e3d6000fd5b505050506040513d60208110156112bc57600080fd5b5051306000908152600160205260408120549192506112db8888611f2b565b600054909150806112ec8487611baf565b816112f357fe5b049a50806113018486611baf565b8161130857fe5b04995060008b11801561131b575060008a115b6113565760405162461bcd60e51b81526004018080602001828103825260228152602001806122786022913960400191505060405180910390fd5b611360308461215f565b61136b878d8d611a1e565b611376868d8c611a1e565b604080516370a0823160e01b815230600482015290516001600160a01b038916916370a08231916024808301926020929190829003018186803b1580156113bc57600080fd5b505afa1580156113d0573d6000803e3d6000fd5b505050506040513d60208110156113e657600080fd5b5051604080516370a0823160e01b815230600482015290519196506001600160a01b038816916370a0823191602480820192602092909190829003018186803b15801561143257600080fd5b505afa158015611446573d6000803e3d6000fd5b505050506040513d602081101561145c57600080fd5b5051935061146c85858b8b611c62565b811561149657600854611492906001600160701b0380821691600160701b900416611baf565b600b555b604080518c8152602081018c905281516001600160a01b038f169233927fdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496929081900390910190a35050505050505050506001600c81905550915091565b6040518060400160405280600381526020016204a4c560ec1b81525081565b6000610b5a338484611e7d565b6103e881565b600c5460011461156b576040805162461bcd60e51b815260206004820152600b60248201526a129bd94e881313d0d2d15160aa1b604482015290519081900360640190fd5b6000600c55600654600754600854604080516370a0823160e01b815230600482015290516001600160a01b039485169490931692611614928592879261160f926001600160701b03169185916370a0823191602480820192602092909190829003018186803b1580156115dd57600080fd5b505afa1580156115f1573d6000803e3d6000fd5b505050506040513d602081101561160757600080fd5b505190611c12565b611a1e565b611688818461160f6008600e9054906101000a90046001600160701b03166001600160701b0316856001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156115dd57600080fd5b50506001600c5550565b6005546001600160a01b031681565b6007546001600160a01b031681565b428410156116f4576040805162461bcd60e51b815260206004820152600c60248201526b129bd94e881156141254915160a21b604482015290519081900360640190fd5b6003546001600160a01b0380891660008181526004602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa15801561180f573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906118455750886001600160a01b0316816001600160a01b0316145b61188f576040805162461bcd60e51b81526020600482015260166024820152754a6f653a20494e56414c49445f5349474e415455524560501b604482015290519081900360640190fd5b61189a898989611e1b565b505050505050505050565b600260209081526000928352604080842090915290825290205481565b600c54600114611907576040805162461bcd60e51b815260206004820152600b60248201526a129bd94e881313d0d2d15160aa1b604482015290519081900360640190fd5b6000600c55600654604080516370a0823160e01b81523060048201529051611a17926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561195857600080fd5b505afa15801561196c573d6000803e3d6000fd5b505050506040513d602081101561198257600080fd5b5051600754604080516370a0823160e01b815230600482015290516001600160a01b03909216916370a0823191602480820192602092909190829003018186803b1580156119cf57600080fd5b505afa1580156119e3573d6000803e3d6000fd5b505050506040513d60208110156119f957600080fd5b50516008546001600160701b0380821691600160701b900416611c62565b6001600c55565b604080518082018252601981527f7472616e7366657228616464726573732c75696e74323536290000000000000060209182015281516001600160a01b0385811660248301526044808301869052845180840390910181526064909201845291810180516001600160e01b031663a9059cbb60e01b1781529251815160009460609489169392918291908083835b60208310611acb5780518252601f199092019160209182019101611aac565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611b2d576040519150601f19603f3d011682016040523d82523d6000602084013e611b32565b606091505b5091509150818015611b60575080511580611b605750808060200190516020811015611b5d57600080fd5b50515b611ba8576040805162461bcd60e51b8152602060048201526014602482015273129bd94e881514905394d1915497d1905253115160621b604482015290519081900360640190fd5b5050505050565b6000811580611bca57505080820282828281611bc757fe5b04145b610b5e576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b80820382811115610b5e576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6001600160701b038411801590611c8057506001600160701b038311155b611cc1576040805162461bcd60e51b815260206004820152600d60248201526c4a6f653a204f564552464c4f5760981b604482015290519081900360640190fd5b60085463ffffffff42811691600160e01b90048116820390811615801590611cf157506001600160701b03841615155b8015611d0557506001600160701b03831615155b15611d70578063ffffffff16611d2d85611d1e866121f1565b6001600160e01b031690612203565b600980546001600160e01b03929092169290920201905563ffffffff8116611d5884611d1e876121f1565b600a80546001600160e01b0392909216929092020190555b600880546dffffffffffffffffffffffffffff19166001600160701b03888116919091176dffffffffffffffffffffffffffff60701b1916600160701b8883168102919091176001600160e01b0316600160e01b63ffffffff871602179283905560408051848416815291909304909116602082015281517f1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1929181900390910190a1505050505050565b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316600090815260016020526040902054611ea09082611c12565b6001600160a01b038085166000908152600160205260408082209390935590841681522054611ecf9082612228565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600080600560009054906101000a90046001600160a01b03166001600160a01b031663017e7e586040518163ffffffff1660e01b815260040160206040518083038186803b158015611f7c57600080fd5b505afa158015611f90573d6000803e3d6000fd5b505050506040513d6020811015611fa657600080fd5b5051600b546001600160a01b038216158015945091925090612057578015612052576000611fe36110136001600160701b03888116908816611baf565b90506000611ff08361206b565b90508082111561204f5760006120126120098484611c12565b60005490611baf565b9050600061202b83612025866005611baf565b90612228565b9050600081838161203857fe5b049050801561204b5761204b87826120bd565b5050505b50505b612063565b8015612063576000600b555b505092915050565b600060038211156120ae575080600160028204015b818110156120a85780915060028182858161209757fe5b0401816120a057fe5b049050612080565b506120b8565b81156120b8575060015b919050565b6000546120ca9082612228565b60009081556001600160a01b0383168152600160205260409020546120ef9082612228565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60008183106121565781612158565b825b9392505050565b6001600160a01b0382166000908152600160205260409020546121829082611c12565b6001600160a01b038316600090815260016020526040812091909155546121a99082611c12565b60009081556040805183815290516001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a35050565b6001600160701b0316600160701b0290565b60006001600160701b0382166001600160e01b0384168161222057fe5b049392505050565b80820182811015610b5e576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfe4a6f653a20494e53554646494349454e545f4c49515549444954595f4255524e45444a6f653a20494e53554646494349454e545f4c49515549444954595f4d494e544544a26469706673582212200df250e1b89fc63fdee72249a7e862d78a0a343843a6a412160203cf32686d9b64736f6c634300060c0033",
}

// JoePairABI is the input ABI used to generate the binding from.
// Deprecated: Use JoePairMetaData.ABI instead.
var JoePairABI = JoePairMetaData.ABI

// Deprecated: Use JoePairMetaData.Sigs instead.
// JoePairFuncSigs maps the 4-byte function signature to its string representation.
var JoePairFuncSigs = JoePairMetaData.Sigs

// JoePairBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JoePairMetaData.Bin instead.
var JoePairBin = JoePairMetaData.Bin

// DeployJoePair deploys a new Ethereum contract, binding an instance of JoePair to it.
func DeployJoePair(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JoePair, error) {
	parsed, err := JoePairMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JoePairBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JoePair{JoePairCaller: JoePairCaller{contract: contract}, JoePairTransactor: JoePairTransactor{contract: contract}, JoePairFilterer: JoePairFilterer{contract: contract}}, nil
}

// JoePair is an auto generated Go binding around an Ethereum contract.
type JoePair struct {
	JoePairCaller     // Read-only binding to the contract
	JoePairTransactor // Write-only binding to the contract
	JoePairFilterer   // Log filterer for contract events
}

// JoePairCaller is an auto generated read-only Go binding around an Ethereum contract.
type JoePairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoePairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JoePairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoePairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JoePairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoePairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JoePairSession struct {
	Contract     *JoePair          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoePairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JoePairCallerSession struct {
	Contract *JoePairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// JoePairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JoePairTransactorSession struct {
	Contract     *JoePairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JoePairRaw is an auto generated low-level Go binding around an Ethereum contract.
type JoePairRaw struct {
	Contract *JoePair // Generic contract binding to access the raw methods on
}

// JoePairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JoePairCallerRaw struct {
	Contract *JoePairCaller // Generic read-only contract binding to access the raw methods on
}

// JoePairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JoePairTransactorRaw struct {
	Contract *JoePairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJoePair creates a new instance of JoePair, bound to a specific deployed contract.
func NewJoePair(address common.Address, backend bind.ContractBackend) (*JoePair, error) {
	contract, err := bindJoePair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JoePair{JoePairCaller: JoePairCaller{contract: contract}, JoePairTransactor: JoePairTransactor{contract: contract}, JoePairFilterer: JoePairFilterer{contract: contract}}, nil
}

// NewJoePairCaller creates a new read-only instance of JoePair, bound to a specific deployed contract.
func NewJoePairCaller(address common.Address, caller bind.ContractCaller) (*JoePairCaller, error) {
	contract, err := bindJoePair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JoePairCaller{contract: contract}, nil
}

// NewJoePairTransactor creates a new write-only instance of JoePair, bound to a specific deployed contract.
func NewJoePairTransactor(address common.Address, transactor bind.ContractTransactor) (*JoePairTransactor, error) {
	contract, err := bindJoePair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JoePairTransactor{contract: contract}, nil
}

// NewJoePairFilterer creates a new log filterer instance of JoePair, bound to a specific deployed contract.
func NewJoePairFilterer(address common.Address, filterer bind.ContractFilterer) (*JoePairFilterer, error) {
	contract, err := bindJoePair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JoePairFilterer{contract: contract}, nil
}

// bindJoePair binds a generic wrapper to an already deployed contract.
func bindJoePair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JoePairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JoePair *JoePairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JoePair.Contract.JoePairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JoePair *JoePairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JoePair.Contract.JoePairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JoePair *JoePairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JoePair.Contract.JoePairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JoePair *JoePairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JoePair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JoePair *JoePairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JoePair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JoePair *JoePairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JoePair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoePair *JoePairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoePair *JoePairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JoePair.Contract.DOMAINSEPARATOR(&_JoePair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JoePair *JoePairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JoePair.Contract.DOMAINSEPARATOR(&_JoePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_JoePair *JoePairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_JoePair *JoePairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _JoePair.Contract.MINIMUMLIQUIDITY(&_JoePair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_JoePair *JoePairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _JoePair.Contract.MINIMUMLIQUIDITY(&_JoePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoePair *JoePairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoePair *JoePairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _JoePair.Contract.PERMITTYPEHASH(&_JoePair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_JoePair *JoePairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _JoePair.Contract.PERMITTYPEHASH(&_JoePair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoePair *JoePairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoePair *JoePairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JoePair.Contract.Allowance(&_JoePair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_JoePair *JoePairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JoePair.Contract.Allowance(&_JoePair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoePair *JoePairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoePair *JoePairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _JoePair.Contract.BalanceOf(&_JoePair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_JoePair *JoePairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _JoePair.Contract.BalanceOf(&_JoePair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoePair *JoePairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoePair *JoePairSession) Decimals() (uint8, error) {
	return _JoePair.Contract.Decimals(&_JoePair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JoePair *JoePairCallerSession) Decimals() (uint8, error) {
	return _JoePair.Contract.Decimals(&_JoePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_JoePair *JoePairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_JoePair *JoePairSession) Factory() (common.Address, error) {
	return _JoePair.Contract.Factory(&_JoePair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_JoePair *JoePairCallerSession) Factory() (common.Address, error) {
	return _JoePair.Contract.Factory(&_JoePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_JoePair *JoePairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_JoePair *JoePairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _JoePair.Contract.GetReserves(&_JoePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_JoePair *JoePairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _JoePair.Contract.GetReserves(&_JoePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_JoePair *JoePairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_JoePair *JoePairSession) KLast() (*big.Int, error) {
	return _JoePair.Contract.KLast(&_JoePair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_JoePair *JoePairCallerSession) KLast() (*big.Int, error) {
	return _JoePair.Contract.KLast(&_JoePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoePair *JoePairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoePair *JoePairSession) Name() (string, error) {
	return _JoePair.Contract.Name(&_JoePair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JoePair *JoePairCallerSession) Name() (string, error) {
	return _JoePair.Contract.Name(&_JoePair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoePair *JoePairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoePair *JoePairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _JoePair.Contract.Nonces(&_JoePair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_JoePair *JoePairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _JoePair.Contract.Nonces(&_JoePair.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_JoePair *JoePairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_JoePair *JoePairSession) Price0CumulativeLast() (*big.Int, error) {
	return _JoePair.Contract.Price0CumulativeLast(&_JoePair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_JoePair *JoePairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _JoePair.Contract.Price0CumulativeLast(&_JoePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_JoePair *JoePairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_JoePair *JoePairSession) Price1CumulativeLast() (*big.Int, error) {
	return _JoePair.Contract.Price1CumulativeLast(&_JoePair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_JoePair *JoePairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _JoePair.Contract.Price1CumulativeLast(&_JoePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoePair *JoePairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoePair *JoePairSession) Symbol() (string, error) {
	return _JoePair.Contract.Symbol(&_JoePair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JoePair *JoePairCallerSession) Symbol() (string, error) {
	return _JoePair.Contract.Symbol(&_JoePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_JoePair *JoePairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_JoePair *JoePairSession) Token0() (common.Address, error) {
	return _JoePair.Contract.Token0(&_JoePair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_JoePair *JoePairCallerSession) Token0() (common.Address, error) {
	return _JoePair.Contract.Token0(&_JoePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_JoePair *JoePairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_JoePair *JoePairSession) Token1() (common.Address, error) {
	return _JoePair.Contract.Token1(&_JoePair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_JoePair *JoePairCallerSession) Token1() (common.Address, error) {
	return _JoePair.Contract.Token1(&_JoePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoePair *JoePairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JoePair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoePair *JoePairSession) TotalSupply() (*big.Int, error) {
	return _JoePair.Contract.TotalSupply(&_JoePair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JoePair *JoePairCallerSession) TotalSupply() (*big.Int, error) {
	return _JoePair.Contract.TotalSupply(&_JoePair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoePair *JoePairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoePair *JoePairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.Approve(&_JoePair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JoePair *JoePairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.Approve(&_JoePair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_JoePair *JoePairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_JoePair *JoePairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Burn(&_JoePair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_JoePair *JoePairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Burn(&_JoePair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_JoePair *JoePairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_JoePair *JoePairSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Initialize(&_JoePair.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_JoePair *JoePairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Initialize(&_JoePair.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_JoePair *JoePairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_JoePair *JoePairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Mint(&_JoePair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_JoePair *JoePairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Mint(&_JoePair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoePair *JoePairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoePair *JoePairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoePair.Contract.Permit(&_JoePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JoePair *JoePairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JoePair.Contract.Permit(&_JoePair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_JoePair *JoePairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_JoePair *JoePairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Skim(&_JoePair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_JoePair *JoePairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _JoePair.Contract.Skim(&_JoePair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_JoePair *JoePairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_JoePair *JoePairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _JoePair.Contract.Swap(&_JoePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_JoePair *JoePairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _JoePair.Contract.Swap(&_JoePair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_JoePair *JoePairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_JoePair *JoePairSession) Sync() (*types.Transaction, error) {
	return _JoePair.Contract.Sync(&_JoePair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_JoePair *JoePairTransactorSession) Sync() (*types.Transaction, error) {
	return _JoePair.Contract.Sync(&_JoePair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoePair *JoePairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoePair *JoePairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.Transfer(&_JoePair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JoePair *JoePairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.Transfer(&_JoePair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoePair *JoePairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoePair *JoePairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.TransferFrom(&_JoePair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JoePair *JoePairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JoePair.Contract.TransferFrom(&_JoePair.TransactOpts, from, to, value)
}

// JoePairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the JoePair contract.
type JoePairApprovalIterator struct {
	Event *JoePairApproval // Event containing the contract specifics and raw log

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
func (it *JoePairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairApproval)
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
		it.Event = new(JoePairApproval)
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
func (it *JoePairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairApproval represents a Approval event raised by the JoePair contract.
type JoePairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoePair *JoePairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JoePairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JoePairApprovalIterator{contract: _JoePair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoePair *JoePairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JoePairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairApproval)
				if err := _JoePair.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JoePair *JoePairFilterer) ParseApproval(log types.Log) (*JoePairApproval, error) {
	event := new(JoePairApproval)
	if err := _JoePair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the JoePair contract.
type JoePairBurnIterator struct {
	Event *JoePairBurn // Event containing the contract specifics and raw log

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
func (it *JoePairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairBurn)
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
		it.Event = new(JoePairBurn)
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
func (it *JoePairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairBurn represents a Burn event raised by the JoePair contract.
type JoePairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_JoePair *JoePairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*JoePairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JoePairBurnIterator{contract: _JoePair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_JoePair *JoePairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *JoePairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairBurn)
				if err := _JoePair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_JoePair *JoePairFilterer) ParseBurn(log types.Log) (*JoePairBurn, error) {
	event := new(JoePairBurn)
	if err := _JoePair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the JoePair contract.
type JoePairMintIterator struct {
	Event *JoePairMint // Event containing the contract specifics and raw log

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
func (it *JoePairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairMint)
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
		it.Event = new(JoePairMint)
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
func (it *JoePairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairMint represents a Mint event raised by the JoePair contract.
type JoePairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_JoePair *JoePairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*JoePairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &JoePairMintIterator{contract: _JoePair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_JoePair *JoePairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *JoePairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairMint)
				if err := _JoePair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_JoePair *JoePairFilterer) ParseMint(log types.Log) (*JoePairMint, error) {
	event := new(JoePairMint)
	if err := _JoePair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the JoePair contract.
type JoePairSwapIterator struct {
	Event *JoePairSwap // Event containing the contract specifics and raw log

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
func (it *JoePairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairSwap)
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
		it.Event = new(JoePairSwap)
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
func (it *JoePairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairSwap represents a Swap event raised by the JoePair contract.
type JoePairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_JoePair *JoePairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*JoePairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JoePairSwapIterator{contract: _JoePair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_JoePair *JoePairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *JoePairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairSwap)
				if err := _JoePair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_JoePair *JoePairFilterer) ParseSwap(log types.Log) (*JoePairSwap, error) {
	event := new(JoePairSwap)
	if err := _JoePair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the JoePair contract.
type JoePairSyncIterator struct {
	Event *JoePairSync // Event containing the contract specifics and raw log

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
func (it *JoePairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairSync)
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
		it.Event = new(JoePairSync)
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
func (it *JoePairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairSync represents a Sync event raised by the JoePair contract.
type JoePairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_JoePair *JoePairFilterer) FilterSync(opts *bind.FilterOpts) (*JoePairSyncIterator, error) {

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &JoePairSyncIterator{contract: _JoePair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_JoePair *JoePairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *JoePairSync) (event.Subscription, error) {

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairSync)
				if err := _JoePair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_JoePair *JoePairFilterer) ParseSync(log types.Log) (*JoePairSync, error) {
	event := new(JoePairSync)
	if err := _JoePair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JoePairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the JoePair contract.
type JoePairTransferIterator struct {
	Event *JoePairTransfer // Event containing the contract specifics and raw log

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
func (it *JoePairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairTransfer)
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
		it.Event = new(JoePairTransfer)
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
func (it *JoePairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairTransfer represents a Transfer event raised by the JoePair contract.
type JoePairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoePair *JoePairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JoePairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JoePairTransferIterator{contract: _JoePair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoePair *JoePairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JoePairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JoePair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairTransfer)
				if err := _JoePair.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JoePair *JoePairFilterer) ParseTransfer(log types.Log) (*JoePairTransfer, error) {
	event := new(JoePairTransfer)
	if err := _JoePair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209dc8e4d00dc401f6690383b6c42f6311c616878b7b3d439f8fcefeb557330f3f64736f6c634300060c0033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// SafeMathJoeMetaData contains all meta data concerning the SafeMathJoe contract.
var SafeMathJoeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220242e35970891fe266a697584f46267e11ea73d1e8855a14e8541e0100694f7c464736f6c634300060c0033",
}

// SafeMathJoeABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathJoeMetaData.ABI instead.
var SafeMathJoeABI = SafeMathJoeMetaData.ABI

// SafeMathJoeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathJoeMetaData.Bin instead.
var SafeMathJoeBin = SafeMathJoeMetaData.Bin

// DeploySafeMathJoe deploys a new Ethereum contract, binding an instance of SafeMathJoe to it.
func DeploySafeMathJoe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathJoe, error) {
	parsed, err := SafeMathJoeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathJoeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathJoe{SafeMathJoeCaller: SafeMathJoeCaller{contract: contract}, SafeMathJoeTransactor: SafeMathJoeTransactor{contract: contract}, SafeMathJoeFilterer: SafeMathJoeFilterer{contract: contract}}, nil
}

// SafeMathJoe is an auto generated Go binding around an Ethereum contract.
type SafeMathJoe struct {
	SafeMathJoeCaller     // Read-only binding to the contract
	SafeMathJoeTransactor // Write-only binding to the contract
	SafeMathJoeFilterer   // Log filterer for contract events
}

// SafeMathJoeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathJoeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathJoeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathJoeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathJoeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathJoeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathJoeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathJoeSession struct {
	Contract     *SafeMathJoe      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathJoeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathJoeCallerSession struct {
	Contract *SafeMathJoeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SafeMathJoeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathJoeTransactorSession struct {
	Contract     *SafeMathJoeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SafeMathJoeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathJoeRaw struct {
	Contract *SafeMathJoe // Generic contract binding to access the raw methods on
}

// SafeMathJoeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathJoeCallerRaw struct {
	Contract *SafeMathJoeCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathJoeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathJoeTransactorRaw struct {
	Contract *SafeMathJoeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathJoe creates a new instance of SafeMathJoe, bound to a specific deployed contract.
func NewSafeMathJoe(address common.Address, backend bind.ContractBackend) (*SafeMathJoe, error) {
	contract, err := bindSafeMathJoe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathJoe{SafeMathJoeCaller: SafeMathJoeCaller{contract: contract}, SafeMathJoeTransactor: SafeMathJoeTransactor{contract: contract}, SafeMathJoeFilterer: SafeMathJoeFilterer{contract: contract}}, nil
}

// NewSafeMathJoeCaller creates a new read-only instance of SafeMathJoe, bound to a specific deployed contract.
func NewSafeMathJoeCaller(address common.Address, caller bind.ContractCaller) (*SafeMathJoeCaller, error) {
	contract, err := bindSafeMathJoe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathJoeCaller{contract: contract}, nil
}

// NewSafeMathJoeTransactor creates a new write-only instance of SafeMathJoe, bound to a specific deployed contract.
func NewSafeMathJoeTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathJoeTransactor, error) {
	contract, err := bindSafeMathJoe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathJoeTransactor{contract: contract}, nil
}

// NewSafeMathJoeFilterer creates a new log filterer instance of SafeMathJoe, bound to a specific deployed contract.
func NewSafeMathJoeFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathJoeFilterer, error) {
	contract, err := bindSafeMathJoe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathJoeFilterer{contract: contract}, nil
}

// bindSafeMathJoe binds a generic wrapper to an already deployed contract.
func bindSafeMathJoe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathJoeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathJoe *SafeMathJoeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathJoe.Contract.SafeMathJoeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathJoe *SafeMathJoeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathJoe.Contract.SafeMathJoeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathJoe *SafeMathJoeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathJoe.Contract.SafeMathJoeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathJoe *SafeMathJoeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathJoe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathJoe *SafeMathJoeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathJoe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathJoe *SafeMathJoeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathJoe.Contract.contract.Transact(opts, method, params...)
}

// UQ112x112MetaData contains all meta data concerning the UQ112x112 contract.
var UQ112x112MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203ed383f0d0c52f6887f09912940493107f1f1576ce33baa1ce3f794968ad547764736f6c634300060c0033",
}

// UQ112x112ABI is the input ABI used to generate the binding from.
// Deprecated: Use UQ112x112MetaData.ABI instead.
var UQ112x112ABI = UQ112x112MetaData.ABI

// UQ112x112Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UQ112x112MetaData.Bin instead.
var UQ112x112Bin = UQ112x112MetaData.Bin

// DeployUQ112x112 deploys a new Ethereum contract, binding an instance of UQ112x112 to it.
func DeployUQ112x112(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UQ112x112, error) {
	parsed, err := UQ112x112MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UQ112x112Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UQ112x112{UQ112x112Caller: UQ112x112Caller{contract: contract}, UQ112x112Transactor: UQ112x112Transactor{contract: contract}, UQ112x112Filterer: UQ112x112Filterer{contract: contract}}, nil
}

// UQ112x112 is an auto generated Go binding around an Ethereum contract.
type UQ112x112 struct {
	UQ112x112Caller     // Read-only binding to the contract
	UQ112x112Transactor // Write-only binding to the contract
	UQ112x112Filterer   // Log filterer for contract events
}

// UQ112x112Caller is an auto generated read-only Go binding around an Ethereum contract.
type UQ112x112Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UQ112x112Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UQ112x112Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UQ112x112Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UQ112x112Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UQ112x112Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UQ112x112Session struct {
	Contract     *UQ112x112        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UQ112x112CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UQ112x112CallerSession struct {
	Contract *UQ112x112Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UQ112x112TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UQ112x112TransactorSession struct {
	Contract     *UQ112x112Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UQ112x112Raw is an auto generated low-level Go binding around an Ethereum contract.
type UQ112x112Raw struct {
	Contract *UQ112x112 // Generic contract binding to access the raw methods on
}

// UQ112x112CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UQ112x112CallerRaw struct {
	Contract *UQ112x112Caller // Generic read-only contract binding to access the raw methods on
}

// UQ112x112TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UQ112x112TransactorRaw struct {
	Contract *UQ112x112Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUQ112x112 creates a new instance of UQ112x112, bound to a specific deployed contract.
func NewUQ112x112(address common.Address, backend bind.ContractBackend) (*UQ112x112, error) {
	contract, err := bindUQ112x112(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UQ112x112{UQ112x112Caller: UQ112x112Caller{contract: contract}, UQ112x112Transactor: UQ112x112Transactor{contract: contract}, UQ112x112Filterer: UQ112x112Filterer{contract: contract}}, nil
}

// NewUQ112x112Caller creates a new read-only instance of UQ112x112, bound to a specific deployed contract.
func NewUQ112x112Caller(address common.Address, caller bind.ContractCaller) (*UQ112x112Caller, error) {
	contract, err := bindUQ112x112(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UQ112x112Caller{contract: contract}, nil
}

// NewUQ112x112Transactor creates a new write-only instance of UQ112x112, bound to a specific deployed contract.
func NewUQ112x112Transactor(address common.Address, transactor bind.ContractTransactor) (*UQ112x112Transactor, error) {
	contract, err := bindUQ112x112(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UQ112x112Transactor{contract: contract}, nil
}

// NewUQ112x112Filterer creates a new log filterer instance of UQ112x112, bound to a specific deployed contract.
func NewUQ112x112Filterer(address common.Address, filterer bind.ContractFilterer) (*UQ112x112Filterer, error) {
	contract, err := bindUQ112x112(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UQ112x112Filterer{contract: contract}, nil
}

// bindUQ112x112 binds a generic wrapper to an already deployed contract.
func bindUQ112x112(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UQ112x112ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UQ112x112 *UQ112x112Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UQ112x112.Contract.UQ112x112Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UQ112x112 *UQ112x112Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UQ112x112.Contract.UQ112x112Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UQ112x112 *UQ112x112Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UQ112x112.Contract.UQ112x112Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UQ112x112 *UQ112x112CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UQ112x112.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UQ112x112 *UQ112x112TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UQ112x112.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UQ112x112 *UQ112x112TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UQ112x112.Contract.contract.Transact(opts, method, params...)
}
