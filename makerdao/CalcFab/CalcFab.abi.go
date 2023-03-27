// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CalcFab

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

// CALCFABABI is the input ABI used to generate the binding from.
const CALCFABABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"newExponentialDecrease\",\"outputs\":[{\"internalType\":\"contractExponentialDecrease\",\"name\":\"calc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"newLinearDecrease\",\"outputs\":[{\"internalType\":\"contractLinearDecrease\",\"name\":\"calc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"newStairstepExponentialDecrease\",\"outputs\":[{\"internalType\":\"contractStairstepExponentialDecrease\",\"name\":\"calc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CALCFAB is an auto generated Go binding around an Ethereum contract.
type CALCFAB struct {
	CALCFABCaller     // Read-only binding to the contract
	CALCFABTransactor // Write-only binding to the contract
	CALCFABFilterer   // Log filterer for contract events
}

// CALCFABCaller is an auto generated read-only Go binding around an Ethereum contract.
type CALCFABCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CALCFABTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CALCFABTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CALCFABFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CALCFABFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CALCFABSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CALCFABSession struct {
	Contract     *CALCFAB          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CALCFABCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CALCFABCallerSession struct {
	Contract *CALCFABCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CALCFABTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CALCFABTransactorSession struct {
	Contract     *CALCFABTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CALCFABRaw is an auto generated low-level Go binding around an Ethereum contract.
type CALCFABRaw struct {
	Contract *CALCFAB // Generic contract binding to access the raw methods on
}

// CALCFABCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CALCFABCallerRaw struct {
	Contract *CALCFABCaller // Generic read-only contract binding to access the raw methods on
}

// CALCFABTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CALCFABTransactorRaw struct {
	Contract *CALCFABTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCALCFAB creates a new instance of CALCFAB, bound to a specific deployed contract.
func NewCALCFAB(address common.Address, backend bind.ContractBackend) (*CALCFAB, error) {
	contract, err := bindCALCFAB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CALCFAB{CALCFABCaller: CALCFABCaller{contract: contract}, CALCFABTransactor: CALCFABTransactor{contract: contract}, CALCFABFilterer: CALCFABFilterer{contract: contract}}, nil
}

// NewCALCFABCaller creates a new read-only instance of CALCFAB, bound to a specific deployed contract.
func NewCALCFABCaller(address common.Address, caller bind.ContractCaller) (*CALCFABCaller, error) {
	contract, err := bindCALCFAB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CALCFABCaller{contract: contract}, nil
}

// NewCALCFABTransactor creates a new write-only instance of CALCFAB, bound to a specific deployed contract.
func NewCALCFABTransactor(address common.Address, transactor bind.ContractTransactor) (*CALCFABTransactor, error) {
	contract, err := bindCALCFAB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CALCFABTransactor{contract: contract}, nil
}

// NewCALCFABFilterer creates a new log filterer instance of CALCFAB, bound to a specific deployed contract.
func NewCALCFABFilterer(address common.Address, filterer bind.ContractFilterer) (*CALCFABFilterer, error) {
	contract, err := bindCALCFAB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CALCFABFilterer{contract: contract}, nil
}

// bindCALCFAB binds a generic wrapper to an already deployed contract.
func bindCALCFAB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CALCFABABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CALCFAB *CALCFABRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CALCFAB.Contract.CALCFABCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CALCFAB *CALCFABRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CALCFAB.Contract.CALCFABTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CALCFAB *CALCFABRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CALCFAB.Contract.CALCFABTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CALCFAB *CALCFABCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CALCFAB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CALCFAB *CALCFABTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CALCFAB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CALCFAB *CALCFABTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CALCFAB.Contract.contract.Transact(opts, method, params...)
}

// NewExponentialDecrease is a paid mutator transaction binding the contract method 0x1888b914.
//
// Solidity: function newExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactor) NewExponentialDecrease(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.contract.Transact(opts, "newExponentialDecrease", owner)
}

// NewExponentialDecrease is a paid mutator transaction binding the contract method 0x1888b914.
//
// Solidity: function newExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABSession) NewExponentialDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewExponentialDecrease(&_CALCFAB.TransactOpts, owner)
}

// NewExponentialDecrease is a paid mutator transaction binding the contract method 0x1888b914.
//
// Solidity: function newExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactorSession) NewExponentialDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewExponentialDecrease(&_CALCFAB.TransactOpts, owner)
}

// NewLinearDecrease is a paid mutator transaction binding the contract method 0xdb4d6271.
//
// Solidity: function newLinearDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactor) NewLinearDecrease(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.contract.Transact(opts, "newLinearDecrease", owner)
}

// NewLinearDecrease is a paid mutator transaction binding the contract method 0xdb4d6271.
//
// Solidity: function newLinearDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABSession) NewLinearDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewLinearDecrease(&_CALCFAB.TransactOpts, owner)
}

// NewLinearDecrease is a paid mutator transaction binding the contract method 0xdb4d6271.
//
// Solidity: function newLinearDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactorSession) NewLinearDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewLinearDecrease(&_CALCFAB.TransactOpts, owner)
}

// NewStairstepExponentialDecrease is a paid mutator transaction binding the contract method 0xf9c2a759.
//
// Solidity: function newStairstepExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactor) NewStairstepExponentialDecrease(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.contract.Transact(opts, "newStairstepExponentialDecrease", owner)
}

// NewStairstepExponentialDecrease is a paid mutator transaction binding the contract method 0xf9c2a759.
//
// Solidity: function newStairstepExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABSession) NewStairstepExponentialDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewStairstepExponentialDecrease(&_CALCFAB.TransactOpts, owner)
}

// NewStairstepExponentialDecrease is a paid mutator transaction binding the contract method 0xf9c2a759.
//
// Solidity: function newStairstepExponentialDecrease(address owner) returns(address calc)
func (_CALCFAB *CALCFABTransactorSession) NewStairstepExponentialDecrease(owner common.Address) (*types.Transaction, error) {
	return _CALCFAB.Contract.NewStairstepExponentialDecrease(&_CALCFAB.TransactOpts, owner)
}
