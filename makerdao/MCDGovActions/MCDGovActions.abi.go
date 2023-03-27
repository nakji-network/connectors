// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDGovActions

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

// MCDGOVACTIONSABI is the input ABI used to generate the binding from.
const MCDGOVACTIONSABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"}],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"dripAndFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"dripAndFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"setAuthorityAndDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDGOVACTIONS is an auto generated Go binding around an Ethereum contract.
type MCDGOVACTIONS struct {
	MCDGOVACTIONSCaller     // Read-only binding to the contract
	MCDGOVACTIONSTransactor // Write-only binding to the contract
	MCDGOVACTIONSFilterer   // Log filterer for contract events
}

// MCDGOVACTIONSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDGOVACTIONSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDGOVACTIONSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDGOVACTIONSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDGOVACTIONSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDGOVACTIONSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDGOVACTIONSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDGOVACTIONSSession struct {
	Contract     *MCDGOVACTIONS    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDGOVACTIONSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDGOVACTIONSCallerSession struct {
	Contract *MCDGOVACTIONSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MCDGOVACTIONSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDGOVACTIONSTransactorSession struct {
	Contract     *MCDGOVACTIONSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MCDGOVACTIONSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDGOVACTIONSRaw struct {
	Contract *MCDGOVACTIONS // Generic contract binding to access the raw methods on
}

// MCDGOVACTIONSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDGOVACTIONSCallerRaw struct {
	Contract *MCDGOVACTIONSCaller // Generic read-only contract binding to access the raw methods on
}

// MCDGOVACTIONSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDGOVACTIONSTransactorRaw struct {
	Contract *MCDGOVACTIONSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDGOVACTIONS creates a new instance of MCDGOVACTIONS, bound to a specific deployed contract.
func NewMCDGOVACTIONS(address common.Address, backend bind.ContractBackend) (*MCDGOVACTIONS, error) {
	contract, err := bindMCDGOVACTIONS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDGOVACTIONS{MCDGOVACTIONSCaller: MCDGOVACTIONSCaller{contract: contract}, MCDGOVACTIONSTransactor: MCDGOVACTIONSTransactor{contract: contract}, MCDGOVACTIONSFilterer: MCDGOVACTIONSFilterer{contract: contract}}, nil
}

// NewMCDGOVACTIONSCaller creates a new read-only instance of MCDGOVACTIONS, bound to a specific deployed contract.
func NewMCDGOVACTIONSCaller(address common.Address, caller bind.ContractCaller) (*MCDGOVACTIONSCaller, error) {
	contract, err := bindMCDGOVACTIONS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDGOVACTIONSCaller{contract: contract}, nil
}

// NewMCDGOVACTIONSTransactor creates a new write-only instance of MCDGOVACTIONS, bound to a specific deployed contract.
func NewMCDGOVACTIONSTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDGOVACTIONSTransactor, error) {
	contract, err := bindMCDGOVACTIONS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDGOVACTIONSTransactor{contract: contract}, nil
}

// NewMCDGOVACTIONSFilterer creates a new log filterer instance of MCDGOVACTIONS, bound to a specific deployed contract.
func NewMCDGOVACTIONSFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDGOVACTIONSFilterer, error) {
	contract, err := bindMCDGOVACTIONS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDGOVACTIONSFilterer{contract: contract}, nil
}

// bindMCDGOVACTIONS binds a generic wrapper to an already deployed contract.
func bindMCDGOVACTIONS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDGOVACTIONSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDGOVACTIONS *MCDGOVACTIONSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDGOVACTIONS.Contract.MCDGOVACTIONSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDGOVACTIONS *MCDGOVACTIONSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.MCDGOVACTIONSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDGOVACTIONS *MCDGOVACTIONSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.MCDGOVACTIONSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDGOVACTIONS *MCDGOVACTIONSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDGOVACTIONS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.contract.Transact(opts, method, params...)
}

// Cage is a paid mutator transaction binding the contract method 0xb169542c.
//
// Solidity: function cage(address end) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) Cage(opts *bind.TransactOpts, end common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "cage", end)
}

// Cage is a paid mutator transaction binding the contract method 0xb169542c.
//
// Solidity: function cage(address end) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) Cage(end common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Cage(&_MCDGOVACTIONS.TransactOpts, end)
}

// Cage is a paid mutator transaction binding the contract method 0xb169542c.
//
// Solidity: function cage(address end) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) Cage(end common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Cage(&_MCDGOVACTIONS.TransactOpts, end)
}

// Deny is a paid mutator transaction binding the contract method 0xe43c09f4.
//
// Solidity: function deny(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) Deny(opts *bind.TransactOpts, who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "deny", who, to)
}

// Deny is a paid mutator transaction binding the contract method 0xe43c09f4.
//
// Solidity: function deny(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) Deny(who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Deny(&_MCDGOVACTIONS.TransactOpts, who, to)
}

// Deny is a paid mutator transaction binding the contract method 0xe43c09f4.
//
// Solidity: function deny(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) Deny(who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Deny(&_MCDGOVACTIONS.TransactOpts, who, to)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x86c09ded.
//
// Solidity: function dripAndFile(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) DripAndFile(opts *bind.TransactOpts, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "dripAndFile", who, ilk, what, data)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x86c09ded.
//
// Solidity: function dripAndFile(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) DripAndFile(who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.DripAndFile(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// DripAndFile is a paid mutator transaction binding the contract method 0x86c09ded.
//
// Solidity: function dripAndFile(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) DripAndFile(who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.DripAndFile(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xe3c277dd.
//
// Solidity: function dripAndFile(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) DripAndFile0(opts *bind.TransactOpts, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "dripAndFile0", who, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xe3c277dd.
//
// Solidity: function dripAndFile(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) DripAndFile0(who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.DripAndFile0(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// DripAndFile0 is a paid mutator transaction binding the contract method 0xe3c277dd.
//
// Solidity: function dripAndFile(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) DripAndFile0(who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.DripAndFile0(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// File is a paid mutator transaction binding the contract method 0x2d9f080e.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) File(opts *bind.TransactOpts, who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "file", who, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x2d9f080e.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) File(who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x2d9f080e.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) File(who common.Address, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xb7beac59.
//
// Solidity: function file(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) File0(opts *bind.TransactOpts, who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "file0", who, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xb7beac59.
//
// Solidity: function file(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) File0(who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File0(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xb7beac59.
//
// Solidity: function file(address who, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) File0(who common.Address, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File0(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xf1f169e7.
//
// Solidity: function file(address who, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) File1(opts *bind.TransactOpts, who common.Address, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "file1", who, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xf1f169e7.
//
// Solidity: function file(address who, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) File1(who common.Address, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File1(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xf1f169e7.
//
// Solidity: function file(address who, bytes32 what, address data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) File1(who common.Address, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File1(&_MCDGOVACTIONS.TransactOpts, who, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xfa6b6f6d.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) File2(opts *bind.TransactOpts, who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "file2", who, ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xfa6b6f6d.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) File2(who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File2(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xfa6b6f6d.
//
// Solidity: function file(address who, bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) File2(who common.Address, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.File2(&_MCDGOVACTIONS.TransactOpts, who, ilk, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address who, bytes32 ilk) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) Init(opts *bind.TransactOpts, who common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "init", who, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address who, bytes32 ilk) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) Init(who common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Init(&_MCDGOVACTIONS.TransactOpts, who, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address who, bytes32 ilk) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) Init(who common.Address, ilk [32]byte) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Init(&_MCDGOVACTIONS.TransactOpts, who, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0xd95270f1.
//
// Solidity: function rely(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) Rely(opts *bind.TransactOpts, who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "rely", who, to)
}

// Rely is a paid mutator transaction binding the contract method 0xd95270f1.
//
// Solidity: function rely(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) Rely(who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Rely(&_MCDGOVACTIONS.TransactOpts, who, to)
}

// Rely is a paid mutator transaction binding the contract method 0xd95270f1.
//
// Solidity: function rely(address who, address to) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) Rely(who common.Address, to common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.Rely(&_MCDGOVACTIONS.TransactOpts, who, to)
}

// SetAuthority is a paid mutator transaction binding the contract method 0xe7796f33.
//
// Solidity: function setAuthority(address pause, address newAuthority) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) SetAuthority(opts *bind.TransactOpts, pause common.Address, newAuthority common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "setAuthority", pause, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0xe7796f33.
//
// Solidity: function setAuthority(address pause, address newAuthority) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) SetAuthority(pause common.Address, newAuthority common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetAuthority(&_MCDGOVACTIONS.TransactOpts, pause, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0xe7796f33.
//
// Solidity: function setAuthority(address pause, address newAuthority) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) SetAuthority(pause common.Address, newAuthority common.Address) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetAuthority(&_MCDGOVACTIONS.TransactOpts, pause, newAuthority)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0xd9c6f876.
//
// Solidity: function setAuthorityAndDelay(address pause, address newAuthority, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) SetAuthorityAndDelay(opts *bind.TransactOpts, pause common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "setAuthorityAndDelay", pause, newAuthority, newDelay)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0xd9c6f876.
//
// Solidity: function setAuthorityAndDelay(address pause, address newAuthority, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) SetAuthorityAndDelay(pause common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetAuthorityAndDelay(&_MCDGOVACTIONS.TransactOpts, pause, newAuthority, newDelay)
}

// SetAuthorityAndDelay is a paid mutator transaction binding the contract method 0xd9c6f876.
//
// Solidity: function setAuthorityAndDelay(address pause, address newAuthority, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) SetAuthorityAndDelay(pause common.Address, newAuthority common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetAuthorityAndDelay(&_MCDGOVACTIONS.TransactOpts, pause, newAuthority, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x6d21e4e2.
//
// Solidity: function setDelay(address pause, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactor) SetDelay(opts *bind.TransactOpts, pause common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.contract.Transact(opts, "setDelay", pause, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x6d21e4e2.
//
// Solidity: function setDelay(address pause, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSSession) SetDelay(pause common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetDelay(&_MCDGOVACTIONS.TransactOpts, pause, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x6d21e4e2.
//
// Solidity: function setDelay(address pause, uint256 newDelay) returns()
func (_MCDGOVACTIONS *MCDGOVACTIONSTransactorSession) SetDelay(pause common.Address, newDelay *big.Int) (*types.Transaction, error) {
	return _MCDGOVACTIONS.Contract.SetDelay(&_MCDGOVACTIONS.TransactOpts, pause, newDelay)
}
