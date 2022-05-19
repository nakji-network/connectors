// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_ACTIONS_DSR

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

// PROXYACTIONSDSRABI is the input ABI used to generate the binding from.
const PROXYACTIONSDSRABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"apt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiJoin_join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pot\",\"type\":\"address\"}],\"name\":\"exitAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PROXYACTIONSDSR is an auto generated Go binding around an Ethereum contract.
type PROXYACTIONSDSR struct {
	PROXYACTIONSDSRCaller     // Read-only binding to the contract
	PROXYACTIONSDSRTransactor // Write-only binding to the contract
	PROXYACTIONSDSRFilterer   // Log filterer for contract events
}

// PROXYACTIONSDSRCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYACTIONSDSRCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSDSRTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYACTIONSDSRTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSDSRFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYACTIONSDSRFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSDSRSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYACTIONSDSRSession struct {
	Contract     *PROXYACTIONSDSR  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PROXYACTIONSDSRCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYACTIONSDSRCallerSession struct {
	Contract *PROXYACTIONSDSRCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PROXYACTIONSDSRTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYACTIONSDSRTransactorSession struct {
	Contract     *PROXYACTIONSDSRTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PROXYACTIONSDSRRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYACTIONSDSRRaw struct {
	Contract *PROXYACTIONSDSR // Generic contract binding to access the raw methods on
}

// PROXYACTIONSDSRCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYACTIONSDSRCallerRaw struct {
	Contract *PROXYACTIONSDSRCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYACTIONSDSRTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYACTIONSDSRTransactorRaw struct {
	Contract *PROXYACTIONSDSRTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYACTIONSDSR creates a new instance of PROXYACTIONSDSR, bound to a specific deployed contract.
func NewPROXYACTIONSDSR(address common.Address, backend bind.ContractBackend) (*PROXYACTIONSDSR, error) {
	contract, err := bindPROXYACTIONSDSR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSDSR{PROXYACTIONSDSRCaller: PROXYACTIONSDSRCaller{contract: contract}, PROXYACTIONSDSRTransactor: PROXYACTIONSDSRTransactor{contract: contract}, PROXYACTIONSDSRFilterer: PROXYACTIONSDSRFilterer{contract: contract}}, nil
}

// NewPROXYACTIONSDSRCaller creates a new read-only instance of PROXYACTIONSDSR, bound to a specific deployed contract.
func NewPROXYACTIONSDSRCaller(address common.Address, caller bind.ContractCaller) (*PROXYACTIONSDSRCaller, error) {
	contract, err := bindPROXYACTIONSDSR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSDSRCaller{contract: contract}, nil
}

// NewPROXYACTIONSDSRTransactor creates a new write-only instance of PROXYACTIONSDSR, bound to a specific deployed contract.
func NewPROXYACTIONSDSRTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYACTIONSDSRTransactor, error) {
	contract, err := bindPROXYACTIONSDSR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSDSRTransactor{contract: contract}, nil
}

// NewPROXYACTIONSDSRFilterer creates a new log filterer instance of PROXYACTIONSDSR, bound to a specific deployed contract.
func NewPROXYACTIONSDSRFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYACTIONSDSRFilterer, error) {
	contract, err := bindPROXYACTIONSDSR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSDSRFilterer{contract: contract}, nil
}

// bindPROXYACTIONSDSR binds a generic wrapper to an already deployed contract.
func bindPROXYACTIONSDSR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYACTIONSDSRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSDSR.Contract.PROXYACTIONSDSRCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.PROXYACTIONSDSRTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.PROXYACTIONSDSRTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSDSR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.contract.Transact(opts, method, params...)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactor) DaiJoinJoin(opts *bind.TransactOpts, apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.contract.Transact(opts, "daiJoin_join", apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.DaiJoinJoin(&_PROXYACTIONSDSR.TransactOpts, apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.DaiJoinJoin(&_PROXYACTIONSDSR.TransactOpts, apt, urn, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactor) Exit(opts *bind.TransactOpts, daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.contract.Transact(opts, "exit", daiJoin, pot, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRSession) Exit(daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.Exit(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorSession) Exit(daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.Exit(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot, wad)
}

// ExitAll is a paid mutator transaction binding the contract method 0xc77843b3.
//
// Solidity: function exitAll(address daiJoin, address pot) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactor) ExitAll(opts *bind.TransactOpts, daiJoin common.Address, pot common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.contract.Transact(opts, "exitAll", daiJoin, pot)
}

// ExitAll is a paid mutator transaction binding the contract method 0xc77843b3.
//
// Solidity: function exitAll(address daiJoin, address pot) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRSession) ExitAll(daiJoin common.Address, pot common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.ExitAll(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot)
}

// ExitAll is a paid mutator transaction binding the contract method 0xc77843b3.
//
// Solidity: function exitAll(address daiJoin, address pot) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorSession) ExitAll(daiJoin common.Address, pot common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.ExitAll(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactor) Join(opts *bind.TransactOpts, daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.contract.Transact(opts, "join", daiJoin, pot, wad)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRSession) Join(daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.Join(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot, wad)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address daiJoin, address pot, uint256 wad) returns()
func (_PROXYACTIONSDSR *PROXYACTIONSDSRTransactorSession) Join(daiJoin common.Address, pot common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSDSR.Contract.Join(&_PROXYACTIONSDSR.TransactOpts, daiJoin, pot, wad)
}
