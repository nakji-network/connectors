// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProxyActionsEnd

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

// PROXYACTIONSENDABI is the input ABI used to generate the binding from.
const PROXYACTIONSENDABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"cashETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"cashGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"apt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiJoin_join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"freeETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"freeGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pack\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PROXYACTIONSEND is an auto generated Go binding around an Ethereum contract.
type PROXYACTIONSEND struct {
	PROXYACTIONSENDCaller     // Read-only binding to the contract
	PROXYACTIONSENDTransactor // Write-only binding to the contract
	PROXYACTIONSENDFilterer   // Log filterer for contract events
}

// PROXYACTIONSENDCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYACTIONSENDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYACTIONSENDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYACTIONSENDSession struct {
	Contract     *PROXYACTIONSEND  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PROXYACTIONSENDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYACTIONSENDCallerSession struct {
	Contract *PROXYACTIONSENDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PROXYACTIONSENDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYACTIONSENDTransactorSession struct {
	Contract     *PROXYACTIONSENDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PROXYACTIONSENDRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYACTIONSENDRaw struct {
	Contract *PROXYACTIONSEND // Generic contract binding to access the raw methods on
}

// PROXYACTIONSENDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCallerRaw struct {
	Contract *PROXYACTIONSENDCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYACTIONSENDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYACTIONSENDTransactorRaw struct {
	Contract *PROXYACTIONSENDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYACTIONSEND creates a new instance of PROXYACTIONSEND, bound to a specific deployed contract.
func NewPROXYACTIONSEND(address common.Address, backend bind.ContractBackend) (*PROXYACTIONSEND, error) {
	contract, err := bindPROXYACTIONSEND(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSEND{PROXYACTIONSENDCaller: PROXYACTIONSENDCaller{contract: contract}, PROXYACTIONSENDTransactor: PROXYACTIONSENDTransactor{contract: contract}, PROXYACTIONSENDFilterer: PROXYACTIONSENDFilterer{contract: contract}}, nil
}

// NewPROXYACTIONSENDCaller creates a new read-only instance of PROXYACTIONSEND, bound to a specific deployed contract.
func NewPROXYACTIONSENDCaller(address common.Address, caller bind.ContractCaller) (*PROXYACTIONSENDCaller, error) {
	contract, err := bindPROXYACTIONSEND(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDCaller{contract: contract}, nil
}

// NewPROXYACTIONSENDTransactor creates a new write-only instance of PROXYACTIONSEND, bound to a specific deployed contract.
func NewPROXYACTIONSENDTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYACTIONSENDTransactor, error) {
	contract, err := bindPROXYACTIONSEND(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDTransactor{contract: contract}, nil
}

// NewPROXYACTIONSENDFilterer creates a new log filterer instance of PROXYACTIONSEND, bound to a specific deployed contract.
func NewPROXYACTIONSENDFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYACTIONSENDFilterer, error) {
	contract, err := bindPROXYACTIONSEND(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDFilterer{contract: contract}, nil
}

// bindPROXYACTIONSEND binds a generic wrapper to an already deployed contract.
func bindPROXYACTIONSEND(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYACTIONSENDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSEND *PROXYACTIONSENDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSEND.Contract.PROXYACTIONSENDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSEND *PROXYACTIONSENDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.PROXYACTIONSENDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSEND *PROXYACTIONSENDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.PROXYACTIONSENDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSEND *PROXYACTIONSENDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSEND.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.contract.Transact(opts, method, params...)
}

// CashETH is a paid mutator transaction binding the contract method 0x66007a9b.
//
// Solidity: function cashETH(address ethJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) CashETH(opts *bind.TransactOpts, ethJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "cashETH", ethJoin, end, ilk, wad)
}

// CashETH is a paid mutator transaction binding the contract method 0x66007a9b.
//
// Solidity: function cashETH(address ethJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) CashETH(ethJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.CashETH(&_PROXYACTIONSEND.TransactOpts, ethJoin, end, ilk, wad)
}

// CashETH is a paid mutator transaction binding the contract method 0x66007a9b.
//
// Solidity: function cashETH(address ethJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) CashETH(ethJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.CashETH(&_PROXYACTIONSEND.TransactOpts, ethJoin, end, ilk, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x938759df.
//
// Solidity: function cashGem(address gemJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) CashGem(opts *bind.TransactOpts, gemJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "cashGem", gemJoin, end, ilk, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x938759df.
//
// Solidity: function cashGem(address gemJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) CashGem(gemJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.CashGem(&_PROXYACTIONSEND.TransactOpts, gemJoin, end, ilk, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x938759df.
//
// Solidity: function cashGem(address gemJoin, address end, bytes32 ilk, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) CashGem(gemJoin common.Address, end common.Address, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.CashGem(&_PROXYACTIONSEND.TransactOpts, gemJoin, end, ilk, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) DaiJoinJoin(opts *bind.TransactOpts, apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "daiJoin_join", apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.DaiJoinJoin(&_PROXYACTIONSEND.TransactOpts, apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.DaiJoinJoin(&_PROXYACTIONSEND.TransactOpts, apt, urn, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0xce168fbe.
//
// Solidity: function freeETH(address manager, address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) FreeETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "freeETH", manager, ethJoin, end, cdp)
}

// FreeETH is a paid mutator transaction binding the contract method 0xce168fbe.
//
// Solidity: function freeETH(address manager, address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) FreeETH(manager common.Address, ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.FreeETH(&_PROXYACTIONSEND.TransactOpts, manager, ethJoin, end, cdp)
}

// FreeETH is a paid mutator transaction binding the contract method 0xce168fbe.
//
// Solidity: function freeETH(address manager, address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) FreeETH(manager common.Address, ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.FreeETH(&_PROXYACTIONSEND.TransactOpts, manager, ethJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0xfc03dbaf.
//
// Solidity: function freeGem(address manager, address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) FreeGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "freeGem", manager, gemJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0xfc03dbaf.
//
// Solidity: function freeGem(address manager, address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) FreeGem(manager common.Address, gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.FreeGem(&_PROXYACTIONSEND.TransactOpts, manager, gemJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0xfc03dbaf.
//
// Solidity: function freeGem(address manager, address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) FreeGem(manager common.Address, gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.FreeGem(&_PROXYACTIONSEND.TransactOpts, manager, gemJoin, end, cdp)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactor) Pack(opts *bind.TransactOpts, daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.contract.Transact(opts, "pack", daiJoin, end, wad)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDSession) Pack(daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.Pack(&_PROXYACTIONSEND.TransactOpts, daiJoin, end, wad)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSEND *PROXYACTIONSENDTransactorSession) Pack(daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSEND.Contract.Pack(&_PROXYACTIONSEND.TransactOpts, daiJoin, end, wad)
}
