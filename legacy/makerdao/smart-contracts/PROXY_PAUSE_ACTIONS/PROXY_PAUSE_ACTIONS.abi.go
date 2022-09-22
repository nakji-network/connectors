// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_PAUSE_ACTIONS

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

// PROXYPAUSEACTIONSABI is the input ABI used to generate the binding from.
const PROXYPAUSEACTIONSABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"dripAndFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"dripAndFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actions\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"setAuthorityAndDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PROXYPAUSEACTIONS is an auto generated Go binding around an Ethereum contract.
type PROXYPAUSEACTIONS struct {
	PROXYPAUSEACTIONSCaller     // Read-only binding to the contract
	PROXYPAUSEACTIONSTransactor // Write-only binding to the contract
	PROXYPAUSEACTIONSFilterer   // Log filterer for contract events
}

// PROXYPAUSEACTIONSCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYPAUSEACTIONSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYPAUSEACTIONSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYPAUSEACTIONSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYPAUSEACTIONSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYPAUSEACTIONSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYPAUSEACTIONSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYPAUSEACTIONSSession struct {
	Contract     *PROXYPAUSEACTIONS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PROXYPAUSEACTIONSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYPAUSEACTIONSCallerSession struct {
	Contract *PROXYPAUSEACTIONSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PROXYPAUSEACTIONSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYPAUSEACTIONSTransactorSession struct {
	Contract     *PROXYPAUSEACTIONSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PROXYPAUSEACTIONSRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYPAUSEACTIONSRaw struct {
	Contract *PROXYPAUSEACTIONS // Generic contract binding to access the raw methods on
}

// PROXYPAUSEACTIONSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYPAUSEACTIONSCallerRaw struct {
	Contract *PROXYPAUSEACTIONSCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYPAUSEACTIONSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYPAUSEACTIONSTransactorRaw struct {
	Contract *PROXYPAUSEACTIONSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYPAUSEACTIONS creates a new instance of PROXYPAUSEACTIONS, bound to a specific deployed contract.
func NewPROXYPAUSEACTIONS(address common.Address, backend bind.ContractBackend) (*PROXYPAUSEACTIONS, error) {
	contract, err := bindPROXYPAUSEACTIONS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYPAUSEACTIONS{PROXYPAUSEACTIONSCaller: PROXYPAUSEACTIONSCaller{contract: contract}, PROXYPAUSEACTIONSTransactor: PROXYPAUSEACTIONSTransactor{contract: contract}, PROXYPAUSEACTIONSFilterer: PROXYPAUSEACTIONSFilterer{contract: contract}}, nil
}

// NewPROXYPAUSEACTIONSCaller creates a new read-only instance of PROXYPAUSEACTIONS, bound to a specific deployed contract.
func NewPROXYPAUSEACTIONSCaller(address common.Address, caller bind.ContractCaller) (*PROXYPAUSEACTIONSCaller, error) {
	contract, err := bindPROXYPAUSEACTIONS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYPAUSEACTIONSCaller{contract: contract}, nil
}

// NewPROXYPAUSEACTIONSTransactor creates a new write-only instance of PROXYPAUSEACTIONS, bound to a specific deployed contract.
func NewPROXYPAUSEACTIONSTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYPAUSEACTIONSTransactor, error) {
	contract, err := bindPROXYPAUSEACTIONS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYPAUSEACTIONSTransactor{contract: contract}, nil
}

// NewPROXYPAUSEACTIONSFilterer creates a new log filterer instance of PROXYPAUSEACTIONS, bound to a specific deployed contract.
func NewPROXYPAUSEACTIONSFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYPAUSEACTIONSFilterer, error) {
	contract, err := bindPROXYPAUSEACTIONS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYPAUSEACTIONSFilterer{contract: contract}, nil
}

// bindPROXYPAUSEACTIONS binds a generic wrapper to an already deployed contract.
func bindPROXYPAUSEACTIONS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYPAUSEACTIONSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYPAUSEACTIONS.Contract.PROXYPAUSEACTIONSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.PROXYPAUSEACTIONSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.PROXYPAUSEACTIONSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYPAUSEACTIONS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.contract.Transact(opts, method, params...)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x67fe0502.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) DripAndFile(opts *bind.TransactOpts, pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "dripAndFile", pause, actions, who, ilk, what, data)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x67fe0502.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) DripAndFile(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.DripAndFile(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x67fe0502.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) DripAndFile(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.DripAndFile(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xc035afa5.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) DripAndFile0(opts *bind.TransactOpts, pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "dripAndFile0", pause, actions, who, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xc035afa5.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) DripAndFile0(pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.DripAndFile0(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xc035afa5.
//
// Solidity: function dripAndFile(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) DripAndFile0(pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.DripAndFile0(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, what, data)
}

// File is a paid mutator transaction binding the contract method 0x4df1249f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) File(opts *bind.TransactOpts, pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "file", pause, actions, who, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x4df1249f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) File(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x4df1249f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) File(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x61be647f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) File0(opts *bind.TransactOpts, pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "file0", pause, actions, who, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x61be647f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) File0(pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File0(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x61be647f.
//
// Solidity: function file(address pause, address actions, address who, bytes32 what, uint256 data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) File0(pause common.Address, actions common.Address, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File0(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0x726e0ace.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, address data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) File1(opts *bind.TransactOpts, pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "file1", pause, actions, who, ilk, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0x726e0ace.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, address data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) File1(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File1(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0x726e0ace.
//
// Solidity: function file(address pause, address actions, address who, bytes32 ilk, bytes32 what, address data) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) File1(pause common.Address, actions common.Address, who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.File1(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, who, ilk, what, data)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0x46d34994.
//
// Solidity: function setAuthorityAndDelay(address pause, address actions, address newAuthority, uint256 newDelay) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactor) SetAuthorityAndDelay(opts *bind.TransactOpts, pause common.Address, actions common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.contract.Transact(opts, "setAuthorityAndDelay", pause, actions, newAuthority, newDelay)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0x46d34994.
//
// Solidity: function setAuthorityAndDelay(address pause, address actions, address newAuthority, uint256 newDelay) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSSession) SetAuthorityAndDelay(pause common.Address, actions common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.SetAuthorityAndDelay(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, newAuthority, newDelay)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0x46d34994.
//
// Solidity: function setAuthorityAndDelay(address pause, address actions, address newAuthority, uint256 newDelay) returns()
func (_PROXYPAUSEACTIONS *PROXYPAUSEACTIONSTransactorSession) SetAuthorityAndDelay(pause common.Address, actions common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _PROXYPAUSEACTIONS.Contract.SetAuthorityAndDelay(&_PROXYPAUSEACTIONS.TransactOpts, pause, actions, newAuthority, newDelay)
}
