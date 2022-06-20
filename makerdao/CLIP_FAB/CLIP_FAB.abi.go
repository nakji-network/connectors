// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CLIP_FAB

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

// CLIPFABABI is the input ABI used to generate the binding from.
const CLIPFABABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vat\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spotter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dog\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"newClip\",\"outputs\":[{\"internalType\":\"contractClipper\",\"name\":\"clip\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CLIPFAB is an auto generated Go binding around an Ethereum contract.
type CLIPFAB struct {
	CLIPFABCaller     // Read-only binding to the contract
	CLIPFABTransactor // Write-only binding to the contract
	CLIPFABFilterer   // Log filterer for contract events
}

// CLIPFABCaller is an auto generated read-only Go binding around an Ethereum contract.
type CLIPFABCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPFABTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CLIPFABTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPFABFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CLIPFABFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLIPFABSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CLIPFABSession struct {
	Contract     *CLIPFAB          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CLIPFABCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CLIPFABCallerSession struct {
	Contract *CLIPFABCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CLIPFABTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CLIPFABTransactorSession struct {
	Contract     *CLIPFABTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CLIPFABRaw is an auto generated low-level Go binding around an Ethereum contract.
type CLIPFABRaw struct {
	Contract *CLIPFAB // Generic contract binding to access the raw methods on
}

// CLIPFABCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CLIPFABCallerRaw struct {
	Contract *CLIPFABCaller // Generic read-only contract binding to access the raw methods on
}

// CLIPFABTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CLIPFABTransactorRaw struct {
	Contract *CLIPFABTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCLIPFAB creates a new instance of CLIPFAB, bound to a specific deployed contract.
func NewCLIPFAB(address common.Address, backend bind.ContractBackend) (*CLIPFAB, error) {
	contract, err := bindCLIPFAB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CLIPFAB{CLIPFABCaller: CLIPFABCaller{contract: contract}, CLIPFABTransactor: CLIPFABTransactor{contract: contract}, CLIPFABFilterer: CLIPFABFilterer{contract: contract}}, nil
}

// NewCLIPFABCaller creates a new read-only instance of CLIPFAB, bound to a specific deployed contract.
func NewCLIPFABCaller(address common.Address, caller bind.ContractCaller) (*CLIPFABCaller, error) {
	contract, err := bindCLIPFAB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CLIPFABCaller{contract: contract}, nil
}

// NewCLIPFABTransactor creates a new write-only instance of CLIPFAB, bound to a specific deployed contract.
func NewCLIPFABTransactor(address common.Address, transactor bind.ContractTransactor) (*CLIPFABTransactor, error) {
	contract, err := bindCLIPFAB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CLIPFABTransactor{contract: contract}, nil
}

// NewCLIPFABFilterer creates a new log filterer instance of CLIPFAB, bound to a specific deployed contract.
func NewCLIPFABFilterer(address common.Address, filterer bind.ContractFilterer) (*CLIPFABFilterer, error) {
	contract, err := bindCLIPFAB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CLIPFABFilterer{contract: contract}, nil
}

// bindCLIPFAB binds a generic wrapper to an already deployed contract.
func bindCLIPFAB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CLIPFABABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLIPFAB *CLIPFABRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLIPFAB.Contract.CLIPFABCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLIPFAB *CLIPFABRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLIPFAB.Contract.CLIPFABTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLIPFAB *CLIPFABRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLIPFAB.Contract.CLIPFABTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLIPFAB *CLIPFABCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLIPFAB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLIPFAB *CLIPFABTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLIPFAB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLIPFAB *CLIPFABTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLIPFAB.Contract.contract.Transact(opts, method, params...)
}

// NewClip is a paid mutator transaction binding the contract method 0xbf5f804e.
//
// Solidity: function newClip(address owner, address vat, address spotter, address dog, bytes32 ilk) returns(address clip)
func (_CLIPFAB *CLIPFABTransactor) NewClip(opts *bind.TransactOpts, owner common.Address, vat common.Address, spotter common.Address, dog common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _CLIPFAB.contract.Transact(opts, "newClip", owner, vat, spotter, dog, ilk)
}

// NewClip is a paid mutator transaction binding the contract method 0xbf5f804e.
//
// Solidity: function newClip(address owner, address vat, address spotter, address dog, bytes32 ilk) returns(address clip)
func (_CLIPFAB *CLIPFABSession) NewClip(owner common.Address, vat common.Address, spotter common.Address, dog common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _CLIPFAB.Contract.NewClip(&_CLIPFAB.TransactOpts, owner, vat, spotter, dog, ilk)
}

// NewClip is a paid mutator transaction binding the contract method 0xbf5f804e.
//
// Solidity: function newClip(address owner, address vat, address spotter, address dog, bytes32 ilk) returns(address clip)
func (_CLIPFAB *CLIPFABTransactorSession) NewClip(owner common.Address, vat common.Address, spotter common.Address, dog common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _CLIPFAB.Contract.NewClip(&_CLIPFAB.TransactOpts, owner, vat, spotter, dog, ilk)
}
