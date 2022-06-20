// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FLIP_FAB

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

// FLIPFABABI is the input ABI used to generate the binding from.
const FLIPFABABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cat\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"newFlip\",\"outputs\":[{\"internalType\":\"contractFlipper\",\"name\":\"flip\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FLIPFAB is an auto generated Go binding around an Ethereum contract.
type FLIPFAB struct {
	FLIPFABCaller     // Read-only binding to the contract
	FLIPFABTransactor // Write-only binding to the contract
	FLIPFABFilterer   // Log filterer for contract events
}

// FLIPFABCaller is an auto generated read-only Go binding around an Ethereum contract.
type FLIPFABCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPFABTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FLIPFABTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPFABFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FLIPFABFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLIPFABSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FLIPFABSession struct {
	Contract     *FLIPFAB          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FLIPFABCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FLIPFABCallerSession struct {
	Contract *FLIPFABCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FLIPFABTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FLIPFABTransactorSession struct {
	Contract     *FLIPFABTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FLIPFABRaw is an auto generated low-level Go binding around an Ethereum contract.
type FLIPFABRaw struct {
	Contract *FLIPFAB // Generic contract binding to access the raw methods on
}

// FLIPFABCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FLIPFABCallerRaw struct {
	Contract *FLIPFABCaller // Generic read-only contract binding to access the raw methods on
}

// FLIPFABTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FLIPFABTransactorRaw struct {
	Contract *FLIPFABTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFLIPFAB creates a new instance of FLIPFAB, bound to a specific deployed contract.
func NewFLIPFAB(address common.Address, backend bind.ContractBackend) (*FLIPFAB, error) {
	contract, err := bindFLIPFAB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FLIPFAB{FLIPFABCaller: FLIPFABCaller{contract: contract}, FLIPFABTransactor: FLIPFABTransactor{contract: contract}, FLIPFABFilterer: FLIPFABFilterer{contract: contract}}, nil
}

// NewFLIPFABCaller creates a new read-only instance of FLIPFAB, bound to a specific deployed contract.
func NewFLIPFABCaller(address common.Address, caller bind.ContractCaller) (*FLIPFABCaller, error) {
	contract, err := bindFLIPFAB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FLIPFABCaller{contract: contract}, nil
}

// NewFLIPFABTransactor creates a new write-only instance of FLIPFAB, bound to a specific deployed contract.
func NewFLIPFABTransactor(address common.Address, transactor bind.ContractTransactor) (*FLIPFABTransactor, error) {
	contract, err := bindFLIPFAB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FLIPFABTransactor{contract: contract}, nil
}

// NewFLIPFABFilterer creates a new log filterer instance of FLIPFAB, bound to a specific deployed contract.
func NewFLIPFABFilterer(address common.Address, filterer bind.ContractFilterer) (*FLIPFABFilterer, error) {
	contract, err := bindFLIPFAB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FLIPFABFilterer{contract: contract}, nil
}

// bindFLIPFAB binds a generic wrapper to an already deployed contract.
func bindFLIPFAB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FLIPFABABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLIPFAB *FLIPFABRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLIPFAB.Contract.FLIPFABCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLIPFAB *FLIPFABRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLIPFAB.Contract.FLIPFABTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLIPFAB *FLIPFABRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLIPFAB.Contract.FLIPFABTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLIPFAB *FLIPFABCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLIPFAB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLIPFAB *FLIPFABTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLIPFAB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLIPFAB *FLIPFABTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLIPFAB.Contract.contract.Transact(opts, method, params...)
}

// NewFlip is a paid mutator transaction binding the contract method 0x04466bbe.
//
// Solidity: function newFlip(address vat, address cat, bytes32 ilk) returns(address flip)
func (_FLIPFAB *FLIPFABTransactor) NewFlip(opts *bind.TransactOpts, vat common.Address, cat common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _FLIPFAB.contract.Transact(opts, "newFlip", vat, cat, ilk)
}

// NewFlip is a paid mutator transaction binding the contract method 0x04466bbe.
//
// Solidity: function newFlip(address vat, address cat, bytes32 ilk) returns(address flip)
func (_FLIPFAB *FLIPFABSession) NewFlip(vat common.Address, cat common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _FLIPFAB.Contract.NewFlip(&_FLIPFAB.TransactOpts, vat, cat, ilk)
}

// NewFlip is a paid mutator transaction binding the contract method 0x04466bbe.
//
// Solidity: function newFlip(address vat, address cat, bytes32 ilk) returns(address flip)
func (_FLIPFAB *FLIPFABTransactorSession) NewFlip(vat common.Address, cat common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _FLIPFAB.Contract.NewFlip(&_FLIPFAB.TransactOpts, vat, cat, ilk)
}
