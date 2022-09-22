// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_PAUSE

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

// MCDPAUSEABI is the input ABI used to generate the binding from.
const MCDPAUSEABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tag\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"drop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tag\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"exec\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"plans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tag\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"plot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractDSPauseProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDPAUSE is an auto generated Go binding around an Ethereum contract.
type MCDPAUSE struct {
	MCDPAUSECaller     // Read-only binding to the contract
	MCDPAUSETransactor // Write-only binding to the contract
	MCDPAUSEFilterer   // Log filterer for contract events
}

// MCDPAUSECaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPAUSECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSETransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPAUSETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPAUSEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPAUSESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPAUSESession struct {
	Contract     *MCDPAUSE         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPAUSECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPAUSECallerSession struct {
	Contract *MCDPAUSECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MCDPAUSETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPAUSETransactorSession struct {
	Contract     *MCDPAUSETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MCDPAUSERaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPAUSERaw struct {
	Contract *MCDPAUSE // Generic contract binding to access the raw methods on
}

// MCDPAUSECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPAUSECallerRaw struct {
	Contract *MCDPAUSECaller // Generic read-only contract binding to access the raw methods on
}

// MCDPAUSETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPAUSETransactorRaw struct {
	Contract *MCDPAUSETransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPAUSE creates a new instance of MCDPAUSE, bound to a specific deployed contract.
func NewMCDPAUSE(address common.Address, backend bind.ContractBackend) (*MCDPAUSE, error) {
	contract, err := bindMCDPAUSE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSE{MCDPAUSECaller: MCDPAUSECaller{contract: contract}, MCDPAUSETransactor: MCDPAUSETransactor{contract: contract}, MCDPAUSEFilterer: MCDPAUSEFilterer{contract: contract}}, nil
}

// NewMCDPAUSECaller creates a new read-only instance of MCDPAUSE, bound to a specific deployed contract.
func NewMCDPAUSECaller(address common.Address, caller bind.ContractCaller) (*MCDPAUSECaller, error) {
	contract, err := bindMCDPAUSE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSECaller{contract: contract}, nil
}

// NewMCDPAUSETransactor creates a new write-only instance of MCDPAUSE, bound to a specific deployed contract.
func NewMCDPAUSETransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPAUSETransactor, error) {
	contract, err := bindMCDPAUSE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSETransactor{contract: contract}, nil
}

// NewMCDPAUSEFilterer creates a new log filterer instance of MCDPAUSE, bound to a specific deployed contract.
func NewMCDPAUSEFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPAUSEFilterer, error) {
	contract, err := bindMCDPAUSE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSEFilterer{contract: contract}, nil
}

// bindMCDPAUSE binds a generic wrapper to an already deployed contract.
func bindMCDPAUSE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPAUSEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPAUSE *MCDPAUSERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPAUSE.Contract.MCDPAUSECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPAUSE *MCDPAUSERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.MCDPAUSETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPAUSE *MCDPAUSERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.MCDPAUSETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPAUSE *MCDPAUSECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPAUSE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPAUSE *MCDPAUSETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPAUSE *MCDPAUSETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDPAUSE *MCDPAUSECaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPAUSE.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDPAUSE *MCDPAUSESession) Authority() (common.Address, error) {
	return _MCDPAUSE.Contract.Authority(&_MCDPAUSE.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDPAUSE *MCDPAUSECallerSession) Authority() (common.Address, error) {
	return _MCDPAUSE.Contract.Authority(&_MCDPAUSE.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_MCDPAUSE *MCDPAUSECaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPAUSE.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_MCDPAUSE *MCDPAUSESession) Delay() (*big.Int, error) {
	return _MCDPAUSE.Contract.Delay(&_MCDPAUSE.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_MCDPAUSE *MCDPAUSECallerSession) Delay() (*big.Int, error) {
	return _MCDPAUSE.Contract.Delay(&_MCDPAUSE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSE *MCDPAUSECaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPAUSE.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSE *MCDPAUSESession) Owner() (common.Address, error) {
	return _MCDPAUSE.Contract.Owner(&_MCDPAUSE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDPAUSE *MCDPAUSECallerSession) Owner() (common.Address, error) {
	return _MCDPAUSE.Contract.Owner(&_MCDPAUSE.CallOpts)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(bool)
func (_MCDPAUSE *MCDPAUSECaller) Plans(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MCDPAUSE.contract.Call(opts, &out, "plans", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(bool)
func (_MCDPAUSE *MCDPAUSESession) Plans(arg0 [32]byte) (bool, error) {
	return _MCDPAUSE.Contract.Plans(&_MCDPAUSE.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(bool)
func (_MCDPAUSE *MCDPAUSECallerSession) Plans(arg0 [32]byte) (bool, error) {
	return _MCDPAUSE.Contract.Plans(&_MCDPAUSE.CallOpts, arg0)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDPAUSE *MCDPAUSECaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPAUSE.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDPAUSE *MCDPAUSESession) Proxy() (common.Address, error) {
	return _MCDPAUSE.Contract.Proxy(&_MCDPAUSE.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDPAUSE *MCDPAUSECallerSession) Proxy() (common.Address, error) {
	return _MCDPAUSE.Contract.Proxy(&_MCDPAUSE.CallOpts)
}

// Drop is a paid mutator transaction binding the contract method 0x162c7de3.
//
// Solidity: function drop(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSETransactor) Drop(opts *bind.TransactOpts, usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "drop", usr, tag, fax, eta)
}

// Drop is a paid mutator transaction binding the contract method 0x162c7de3.
//
// Solidity: function drop(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSESession) Drop(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Drop(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// Drop is a paid mutator transaction binding the contract method 0x162c7de3.
//
// Solidity: function drop(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSETransactorSession) Drop(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Drop(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// Exec is a paid mutator transaction binding the contract method 0x168ccd67.
//
// Solidity: function exec(address usr, bytes32 tag, bytes fax, uint256 eta) returns(bytes out)
func (_MCDPAUSE *MCDPAUSETransactor) Exec(opts *bind.TransactOpts, usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "exec", usr, tag, fax, eta)
}

// Exec is a paid mutator transaction binding the contract method 0x168ccd67.
//
// Solidity: function exec(address usr, bytes32 tag, bytes fax, uint256 eta) returns(bytes out)
func (_MCDPAUSE *MCDPAUSESession) Exec(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Exec(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// Exec is a paid mutator transaction binding the contract method 0x168ccd67.
//
// Solidity: function exec(address usr, bytes32 tag, bytes fax, uint256 eta) returns(bytes out)
func (_MCDPAUSE *MCDPAUSETransactorSession) Exec(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Exec(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// Plot is a paid mutator transaction binding the contract method 0x46d2fbbb.
//
// Solidity: function plot(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSETransactor) Plot(opts *bind.TransactOpts, usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "plot", usr, tag, fax, eta)
}

// Plot is a paid mutator transaction binding the contract method 0x46d2fbbb.
//
// Solidity: function plot(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSESession) Plot(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Plot(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// Plot is a paid mutator transaction binding the contract method 0x46d2fbbb.
//
// Solidity: function plot(address usr, bytes32 tag, bytes fax, uint256 eta) returns()
func (_MCDPAUSE *MCDPAUSETransactorSession) Plot(usr common.Address, tag [32]byte, fax []byte, eta *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.Plot(&_MCDPAUSE.TransactOpts, usr, tag, fax, eta)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDPAUSE *MCDPAUSETransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDPAUSE *MCDPAUSESession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetAuthority(&_MCDPAUSE.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDPAUSE *MCDPAUSETransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetAuthority(&_MCDPAUSE.TransactOpts, authority_)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_MCDPAUSE *MCDPAUSETransactor) SetDelay(opts *bind.TransactOpts, delay_ *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "setDelay", delay_)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_MCDPAUSE *MCDPAUSESession) SetDelay(delay_ *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetDelay(&_MCDPAUSE.TransactOpts, delay_)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_MCDPAUSE *MCDPAUSETransactorSession) SetDelay(delay_ *big.Int) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetDelay(&_MCDPAUSE.TransactOpts, delay_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDPAUSE *MCDPAUSETransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDPAUSE *MCDPAUSESession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetOwner(&_MCDPAUSE.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDPAUSE *MCDPAUSETransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDPAUSE.Contract.SetOwner(&_MCDPAUSE.TransactOpts, owner_)
}

// MCDPAUSELogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the MCDPAUSE contract.
type MCDPAUSELogSetAuthorityIterator struct {
	Event *MCDPAUSELogSetAuthority // Event containing the contract specifics and raw log

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
func (it *MCDPAUSELogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPAUSELogSetAuthority)
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
		it.Event = new(MCDPAUSELogSetAuthority)
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
func (it *MCDPAUSELogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPAUSELogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPAUSELogSetAuthority represents a LogSetAuthority event raised by the MCDPAUSE contract.
type MCDPAUSELogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDPAUSE *MCDPAUSEFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MCDPAUSELogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDPAUSE.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSELogSetAuthorityIterator{contract: _MCDPAUSE.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDPAUSE *MCDPAUSEFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MCDPAUSELogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDPAUSE.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPAUSELogSetAuthority)
				if err := _MCDPAUSE.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDPAUSE *MCDPAUSEFilterer) ParseLogSetAuthority(log types.Log) (*MCDPAUSELogSetAuthority, error) {
	event := new(MCDPAUSELogSetAuthority)
	if err := _MCDPAUSE.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPAUSELogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the MCDPAUSE contract.
type MCDPAUSELogSetOwnerIterator struct {
	Event *MCDPAUSELogSetOwner // Event containing the contract specifics and raw log

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
func (it *MCDPAUSELogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPAUSELogSetOwner)
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
		it.Event = new(MCDPAUSELogSetOwner)
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
func (it *MCDPAUSELogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPAUSELogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPAUSELogSetOwner represents a LogSetOwner event raised by the MCDPAUSE contract.
type MCDPAUSELogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDPAUSE *MCDPAUSEFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MCDPAUSELogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPAUSE.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPAUSELogSetOwnerIterator{contract: _MCDPAUSE.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDPAUSE *MCDPAUSEFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MCDPAUSELogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPAUSE.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPAUSELogSetOwner)
				if err := _MCDPAUSE.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDPAUSE *MCDPAUSEFilterer) ParseLogSetOwner(log types.Log) (*MCDPAUSELogSetOwner, error) {
	event := new(MCDPAUSELogSetOwner)
	if err := _MCDPAUSE.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
