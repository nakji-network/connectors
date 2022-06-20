// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_PAUSE_PROXY

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

// MCDPAUSEPROXYABI is the input ABI used to generate the binding from.
const MCDPAUSEPROXYABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"exec\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDPAUSEPROXY is an auto generated Go binding around an Ethereum contract.
type MCDPAUSEPROXY struct {
	MCDPAUSEPROXYCaller     // Read-only binding to the contract
	MCDPAUSEPROXYTransactor // Write-only binding to the contract
	MCDPAUSEPROXYFilterer   // Log filterer for contract events
}

// MCDPAUSEPROXYCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPAUSEPROXYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSEPROXYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPAUSEPROXYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSEPROXYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPAUSEPROXYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSEPROXYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPAUSEPROXYSession struct {
	Contract     *MCDPAUSEPROXY    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPAUSEPROXYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPAUSEPROXYCallerSession struct {
	Contract *MCDPAUSEPROXYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MCDPAUSEPROXYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPAUSEPROXYTransactorSession struct {
	Contract     *MCDPAUSEPROXYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MCDPAUSEPROXYRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPAUSEPROXYRaw struct {
	Contract *MCDPAUSEPROXY // Generic contract binding to access the raw methods on
}

// MCDPAUSEPROXYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPAUSEPROXYCallerRaw struct {
	Contract *MCDPAUSEPROXYCaller // Generic read-only contract binding to access the raw methods on
}

// MCDPAUSEPROXYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPAUSEPROXYTransactorRaw struct {
	Contract *MCDPAUSEPROXYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPAUSEPROXY creates a new instance of MCDPAUSEPROXY, bound to a specific deployed contract.
func NewMCDPAUSEPROXY(address common.Address, backend bind.ContractBackend) (*MCDPAUSEPROXY, error) {
	contract, err := bindMCDPAUSEPROXY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSEPROXY{MCDPAUSEPROXYCaller: MCDPAUSEPROXYCaller{contract: contract}, MCDPAUSEPROXYTransactor: MCDPAUSEPROXYTransactor{contract: contract}, MCDPAUSEPROXYFilterer: MCDPAUSEPROXYFilterer{contract: contract}}, nil
}

// NewMCDPAUSEPROXYCaller creates a new read-only instance of MCDPAUSEPROXY, bound to a specific deployed contract.
func NewMCDPAUSEPROXYCaller(address common.Address, caller bind.ContractCaller) (*MCDPAUSEPROXYCaller, error) {
	contract, err := bindMCDPAUSEPROXY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSEPROXYCaller{contract: contract}, nil
}

// NewMCDPAUSEPROXYTransactor creates a new write-only instance of MCDPAUSEPROXY, bound to a specific deployed contract.
func NewMCDPAUSEPROXYTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPAUSEPROXYTransactor, error) {
	contract, err := bindMCDPAUSEPROXY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSEPROXYTransactor{contract: contract}, nil
}

// NewMCDPAUSEPROXYFilterer creates a new log filterer instance of MCDPAUSEPROXY, bound to a specific deployed contract.
func NewMCDPAUSEPROXYFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPAUSEPROXYFilterer, error) {
	contract, err := bindMCDPAUSEPROXY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSEPROXYFilterer{contract: contract}, nil
}

// bindMCDPAUSEPROXY binds a generic wrapper to an already deployed contract.
func bindMCDPAUSEPROXY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPAUSEPROXYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPAUSEPROXY.Contract.MCDPAUSEPROXYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.MCDPAUSEPROXYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.MCDPAUSEPROXYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPAUSEPROXY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPAUSEPROXY *MCDPAUSEPROXYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPAUSEPROXY.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYSession) Owner() (common.Address, error) {
	return _MCDPAUSEPROXY.Contract.Owner(&_MCDPAUSEPROXY.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYCallerSession) Owner() (common.Address, error) {
	return _MCDPAUSEPROXY.Contract.Owner(&_MCDPAUSEPROXY.CallOpts)
}

// Exec is a paid mutator transaction binding the contract method 0xbe6002c2.
//
// Solidity: function exec(address usr, bytes fax) returns(bytes out)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYTransactor) Exec(opts *bind.TransactOpts, usr common.Address, fax []byte) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.contract.Transact(opts, "exec", usr, fax)
}

// Exec is a paid mutator transaction binding the contract method 0xbe6002c2.
//
// Solidity: function exec(address usr, bytes fax) returns(bytes out)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYSession) Exec(usr common.Address, fax []byte) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.Exec(&_MCDPAUSEPROXY.TransactOpts, usr, fax)
}

// Exec is a paid mutator transaction binding the contract method 0xbe6002c2.
//
// Solidity: function exec(address usr, bytes fax) returns(bytes out)
func (_MCDPAUSEPROXY *MCDPAUSEPROXYTransactorSession) Exec(usr common.Address, fax []byte) (*types.Transaction, error) {
	return _MCDPAUSEPROXY.Contract.Exec(&_MCDPAUSEPROXY.TransactOpts, usr, fax)
}
