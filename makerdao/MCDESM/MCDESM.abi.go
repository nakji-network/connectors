// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDESM

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

// MCDESMABI is the input ABI used to generate the binding from.
const MCDESMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxy_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"min_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pause\",\"type\":\"address\"}],\"name\":\"DenyProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Fire\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Sum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"denyProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"internalType\":\"contractEndLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokesGovernanceAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ret\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDESM is an auto generated Go binding around an Ethereum contract.
type MCDESM struct {
	MCDESMCaller     // Read-only binding to the contract
	MCDESMTransactor // Write-only binding to the contract
	MCDESMFilterer   // Log filterer for contract events
}

// MCDESMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDESMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDESMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDESMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDESMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDESMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDESMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDESMSession struct {
	Contract     *MCDESM           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDESMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDESMCallerSession struct {
	Contract *MCDESMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDESMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDESMTransactorSession struct {
	Contract     *MCDESMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDESMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDESMRaw struct {
	Contract *MCDESM // Generic contract binding to access the raw methods on
}

// MCDESMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDESMCallerRaw struct {
	Contract *MCDESMCaller // Generic read-only contract binding to access the raw methods on
}

// MCDESMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDESMTransactorRaw struct {
	Contract *MCDESMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDESM creates a new instance of MCDESM, bound to a specific deployed contract.
func NewMCDESM(address common.Address, backend bind.ContractBackend) (*MCDESM, error) {
	contract, err := bindMCDESM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDESM{MCDESMCaller: MCDESMCaller{contract: contract}, MCDESMTransactor: MCDESMTransactor{contract: contract}, MCDESMFilterer: MCDESMFilterer{contract: contract}}, nil
}

// NewMCDESMCaller creates a new read-only instance of MCDESM, bound to a specific deployed contract.
func NewMCDESMCaller(address common.Address, caller bind.ContractCaller) (*MCDESMCaller, error) {
	contract, err := bindMCDESM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDESMCaller{contract: contract}, nil
}

// NewMCDESMTransactor creates a new write-only instance of MCDESM, bound to a specific deployed contract.
func NewMCDESMTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDESMTransactor, error) {
	contract, err := bindMCDESM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDESMTransactor{contract: contract}, nil
}

// NewMCDESMFilterer creates a new log filterer instance of MCDESM, bound to a specific deployed contract.
func NewMCDESMFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDESMFilterer, error) {
	contract, err := bindMCDESM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDESMFilterer{contract: contract}, nil
}

// bindMCDESM binds a generic wrapper to an already deployed contract.
func bindMCDESM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDESMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDESM *MCDESMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDESM.Contract.MCDESMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDESM *MCDESMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDESM.Contract.MCDESMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDESM *MCDESMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDESM.Contract.MCDESMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDESM *MCDESMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDESM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDESM *MCDESMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDESM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDESM *MCDESMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDESM.Contract.contract.Transact(opts, method, params...)
}

// Sum is a free data retrieval call binding the contract method 0x37be827d.
//
// Solidity: function Sum() view returns(uint256)
func (_MCDESM *MCDESMCaller) Sum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "Sum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sum is a free data retrieval call binding the contract method 0x37be827d.
//
// Solidity: function Sum() view returns(uint256)
func (_MCDESM *MCDESMSession) Sum() (*big.Int, error) {
	return _MCDESM.Contract.Sum(&_MCDESM.CallOpts)
}

// Sum is a free data retrieval call binding the contract method 0x37be827d.
//
// Solidity: function Sum() view returns(uint256)
func (_MCDESM *MCDESMCallerSession) Sum() (*big.Int, error) {
	return _MCDESM.Contract.Sum(&_MCDESM.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDESM *MCDESMCaller) End(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "end")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDESM *MCDESMSession) End() (common.Address, error) {
	return _MCDESM.Contract.End(&_MCDESM.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDESM *MCDESMCallerSession) End() (common.Address, error) {
	return _MCDESM.Contract.End(&_MCDESM.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDESM *MCDESMCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDESM *MCDESMSession) Gem() (common.Address, error) {
	return _MCDESM.Contract.Gem(&_MCDESM.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDESM *MCDESMCallerSession) Gem() (common.Address, error) {
	return _MCDESM.Contract.Gem(&_MCDESM.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDESM *MCDESMCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDESM *MCDESMSession) Live() (*big.Int, error) {
	return _MCDESM.Contract.Live(&_MCDESM.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDESM *MCDESMCallerSession) Live() (*big.Int, error) {
	return _MCDESM.Contract.Live(&_MCDESM.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_MCDESM *MCDESMCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_MCDESM *MCDESMSession) Min() (*big.Int, error) {
	return _MCDESM.Contract.Min(&_MCDESM.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_MCDESM *MCDESMCallerSession) Min() (*big.Int, error) {
	return _MCDESM.Contract.Min(&_MCDESM.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDESM *MCDESMCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDESM *MCDESMSession) Proxy() (common.Address, error) {
	return _MCDESM.Contract.Proxy(&_MCDESM.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MCDESM *MCDESMCallerSession) Proxy() (common.Address, error) {
	return _MCDESM.Contract.Proxy(&_MCDESM.CallOpts)
}

// RevokesGovernanceAccess is a free data retrieval call binding the contract method 0x14c7bbd5.
//
// Solidity: function revokesGovernanceAccess() view returns(bool ret)
func (_MCDESM *MCDESMCaller) RevokesGovernanceAccess(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "revokesGovernanceAccess")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RevokesGovernanceAccess is a free data retrieval call binding the contract method 0x14c7bbd5.
//
// Solidity: function revokesGovernanceAccess() view returns(bool ret)
func (_MCDESM *MCDESMSession) RevokesGovernanceAccess() (bool, error) {
	return _MCDESM.Contract.RevokesGovernanceAccess(&_MCDESM.CallOpts)
}

// RevokesGovernanceAccess is a free data retrieval call binding the contract method 0x14c7bbd5.
//
// Solidity: function revokesGovernanceAccess() view returns(bool ret)
func (_MCDESM *MCDESMCallerSession) RevokesGovernanceAccess() (bool, error) {
	return _MCDESM.Contract.RevokesGovernanceAccess(&_MCDESM.CallOpts)
}

// Sum2 is a free data retrieval call binding the contract method 0x7e459c60.
//
// Solidity: function sum(address ) view returns(uint256)
func (_MCDESM *MCDESMCaller) Sum2(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "sum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sum2 is a free data retrieval call binding the contract method 0x7e459c60.
//
// Solidity: function sum(address ) view returns(uint256)
func (_MCDESM *MCDESMSession) Sum2(arg0 common.Address) (*big.Int, error) {
	return _MCDESM.Contract.Sum2(&_MCDESM.CallOpts, arg0)
}

// Sum2 is a free data retrieval call binding the contract method 0x7e459c60.
//
// Solidity: function sum(address ) view returns(uint256)
func (_MCDESM *MCDESMCallerSession) Sum2(arg0 common.Address) (*big.Int, error) {
	return _MCDESM.Contract.Sum2(&_MCDESM.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDESM *MCDESMCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDESM.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDESM *MCDESMSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDESM.Contract.Wards(&_MCDESM.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDESM *MCDESMCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDESM.Contract.Wards(&_MCDESM.CallOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MCDESM *MCDESMTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MCDESM *MCDESMSession) Burn() (*types.Transaction, error) {
	return _MCDESM.Contract.Burn(&_MCDESM.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_MCDESM *MCDESMTransactorSession) Burn() (*types.Transaction, error) {
	return _MCDESM.Contract.Burn(&_MCDESM.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDESM *MCDESMTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDESM *MCDESMSession) Cage() (*types.Transaction, error) {
	return _MCDESM.Contract.Cage(&_MCDESM.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDESM *MCDESMTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDESM.Contract.Cage(&_MCDESM.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDESM *MCDESMTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDESM *MCDESMSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.Deny(&_MCDESM.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDESM *MCDESMTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.Deny(&_MCDESM.TransactOpts, usr)
}

// DenyProxy is a paid mutator transaction binding the contract method 0x0715940e.
//
// Solidity: function denyProxy(address target) returns()
func (_MCDESM *MCDESMTransactor) DenyProxy(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "denyProxy", target)
}

// DenyProxy is a paid mutator transaction binding the contract method 0x0715940e.
//
// Solidity: function denyProxy(address target) returns()
func (_MCDESM *MCDESMSession) DenyProxy(target common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.DenyProxy(&_MCDESM.TransactOpts, target)
}

// DenyProxy is a paid mutator transaction binding the contract method 0x0715940e.
//
// Solidity: function denyProxy(address target) returns()
func (_MCDESM *MCDESMTransactorSession) DenyProxy(target common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.DenyProxy(&_MCDESM.TransactOpts, target)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDESM *MCDESMTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDESM *MCDESMSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDESM.Contract.File(&_MCDESM.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDESM *MCDESMTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDESM.Contract.File(&_MCDESM.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDESM *MCDESMTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDESM *MCDESMSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.File0(&_MCDESM.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDESM *MCDESMTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.File0(&_MCDESM.TransactOpts, what, data)
}

// Fire is a paid mutator transaction binding the contract method 0x457094cc.
//
// Solidity: function fire() returns()
func (_MCDESM *MCDESMTransactor) Fire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "fire")
}

// Fire is a paid mutator transaction binding the contract method 0x457094cc.
//
// Solidity: function fire() returns()
func (_MCDESM *MCDESMSession) Fire() (*types.Transaction, error) {
	return _MCDESM.Contract.Fire(&_MCDESM.TransactOpts)
}

// Fire is a paid mutator transaction binding the contract method 0x457094cc.
//
// Solidity: function fire() returns()
func (_MCDESM *MCDESMTransactorSession) Fire() (*types.Transaction, error) {
	return _MCDESM.Contract.Fire(&_MCDESM.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDESM *MCDESMTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDESM *MCDESMSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _MCDESM.Contract.Join(&_MCDESM.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDESM *MCDESMTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _MCDESM.Contract.Join(&_MCDESM.TransactOpts, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDESM *MCDESMTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDESM.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDESM *MCDESMSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.Rely(&_MCDESM.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDESM *MCDESMTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDESM.Contract.Rely(&_MCDESM.TransactOpts, usr)
}

// MCDESMDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDESM contract.
type MCDESMDenyIterator struct {
	Event *MCDESMDeny // Event containing the contract specifics and raw log

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
func (it *MCDESMDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMDeny)
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
		it.Event = new(MCDESMDeny)
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
func (it *MCDESMDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMDeny represents a Deny event raised by the MCDESM contract.
type MCDESMDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDESM *MCDESMFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDESMDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMDenyIterator{contract: _MCDESM.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDESM *MCDESMFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDESMDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMDeny)
				if err := _MCDESM.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDESM *MCDESMFilterer) ParseDeny(log types.Log) (*MCDESMDeny, error) {
	event := new(MCDESMDeny)
	if err := _MCDESM.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMDenyProxyIterator is returned from FilterDenyProxy and is used to iterate over the raw logs and unpacked data for DenyProxy events raised by the MCDESM contract.
type MCDESMDenyProxyIterator struct {
	Event *MCDESMDenyProxy // Event containing the contract specifics and raw log

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
func (it *MCDESMDenyProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMDenyProxy)
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
		it.Event = new(MCDESMDenyProxy)
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
func (it *MCDESMDenyProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMDenyProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMDenyProxy represents a DenyProxy event raised by the MCDESM contract.
type MCDESMDenyProxy struct {
	Base  common.Address
	Pause common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDenyProxy is a free log retrieval operation binding the contract event 0x74269424d52f84560e5738dd2769b61b0e04d5c87c87ca043a1e005e0722a678.
//
// Solidity: event DenyProxy(address indexed base, address indexed pause)
func (_MCDESM *MCDESMFilterer) FilterDenyProxy(opts *bind.FilterOpts, base []common.Address, pause []common.Address) (*MCDESMDenyProxyIterator, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var pauseRule []interface{}
	for _, pauseItem := range pause {
		pauseRule = append(pauseRule, pauseItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "DenyProxy", baseRule, pauseRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMDenyProxyIterator{contract: _MCDESM.contract, event: "DenyProxy", logs: logs, sub: sub}, nil
}

// WatchDenyProxy is a free log subscription operation binding the contract event 0x74269424d52f84560e5738dd2769b61b0e04d5c87c87ca043a1e005e0722a678.
//
// Solidity: event DenyProxy(address indexed base, address indexed pause)
func (_MCDESM *MCDESMFilterer) WatchDenyProxy(opts *bind.WatchOpts, sink chan<- *MCDESMDenyProxy, base []common.Address, pause []common.Address) (event.Subscription, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var pauseRule []interface{}
	for _, pauseItem := range pause {
		pauseRule = append(pauseRule, pauseItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "DenyProxy", baseRule, pauseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMDenyProxy)
				if err := _MCDESM.contract.UnpackLog(event, "DenyProxy", log); err != nil {
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

// ParseDenyProxy is a log parse operation binding the contract event 0x74269424d52f84560e5738dd2769b61b0e04d5c87c87ca043a1e005e0722a678.
//
// Solidity: event DenyProxy(address indexed base, address indexed pause)
func (_MCDESM *MCDESMFilterer) ParseDenyProxy(log types.Log) (*MCDESMDenyProxy, error) {
	event := new(MCDESMDenyProxy)
	if err := _MCDESM.contract.UnpackLog(event, "DenyProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDESM contract.
type MCDESMFileIterator struct {
	Event *MCDESMFile // Event containing the contract specifics and raw log

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
func (it *MCDESMFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMFile)
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
		it.Event = new(MCDESMFile)
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
func (it *MCDESMFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMFile represents a File event raised by the MCDESM contract.
type MCDESMFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDESM *MCDESMFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDESMFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMFileIterator{contract: _MCDESM.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDESM *MCDESMFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDESMFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMFile)
				if err := _MCDESM.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDESM *MCDESMFilterer) ParseFile(log types.Log) (*MCDESMFile, error) {
	event := new(MCDESMFile)
	if err := _MCDESM.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the MCDESM contract.
type MCDESMFile0Iterator struct {
	Event *MCDESMFile0 // Event containing the contract specifics and raw log

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
func (it *MCDESMFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMFile0)
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
		it.Event = new(MCDESMFile0)
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
func (it *MCDESMFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMFile0 represents a File0 event raised by the MCDESM contract.
type MCDESMFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDESM *MCDESMFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*MCDESMFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMFile0Iterator{contract: _MCDESM.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDESM *MCDESMFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *MCDESMFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMFile0)
				if err := _MCDESM.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDESM *MCDESMFilterer) ParseFile0(log types.Log) (*MCDESMFile0, error) {
	event := new(MCDESMFile0)
	if err := _MCDESM.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMFireIterator is returned from FilterFire and is used to iterate over the raw logs and unpacked data for Fire events raised by the MCDESM contract.
type MCDESMFireIterator struct {
	Event *MCDESMFire // Event containing the contract specifics and raw log

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
func (it *MCDESMFireIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMFire)
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
		it.Event = new(MCDESMFire)
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
func (it *MCDESMFireIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMFireIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMFire represents a Fire event raised by the MCDESM contract.
type MCDESMFire struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFire is a free log retrieval operation binding the contract event 0x14c539ef1ff6ef515371448cd46419fd75ebb8698b7f6644c4f66b03b0327085.
//
// Solidity: event Fire()
func (_MCDESM *MCDESMFilterer) FilterFire(opts *bind.FilterOpts) (*MCDESMFireIterator, error) {

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "Fire")
	if err != nil {
		return nil, err
	}
	return &MCDESMFireIterator{contract: _MCDESM.contract, event: "Fire", logs: logs, sub: sub}, nil
}

// WatchFire is a free log subscription operation binding the contract event 0x14c539ef1ff6ef515371448cd46419fd75ebb8698b7f6644c4f66b03b0327085.
//
// Solidity: event Fire()
func (_MCDESM *MCDESMFilterer) WatchFire(opts *bind.WatchOpts, sink chan<- *MCDESMFire) (event.Subscription, error) {

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "Fire")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMFire)
				if err := _MCDESM.contract.UnpackLog(event, "Fire", log); err != nil {
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

// ParseFire is a log parse operation binding the contract event 0x14c539ef1ff6ef515371448cd46419fd75ebb8698b7f6644c4f66b03b0327085.
//
// Solidity: event Fire()
func (_MCDESM *MCDESMFilterer) ParseFire(log types.Log) (*MCDESMFire, error) {
	event := new(MCDESMFire)
	if err := _MCDESM.contract.UnpackLog(event, "Fire", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the MCDESM contract.
type MCDESMJoinIterator struct {
	Event *MCDESMJoin // Event containing the contract specifics and raw log

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
func (it *MCDESMJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMJoin)
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
		it.Event = new(MCDESMJoin)
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
func (it *MCDESMJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMJoin represents a Join event raised by the MCDESM contract.
type MCDESMJoin struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_MCDESM *MCDESMFilterer) FilterJoin(opts *bind.FilterOpts, usr []common.Address) (*MCDESMJoinIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMJoinIterator{contract: _MCDESM.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_MCDESM *MCDESMFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *MCDESMJoin, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMJoin)
				if err := _MCDESM.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_MCDESM *MCDESMFilterer) ParseJoin(log types.Log) (*MCDESMJoin, error) {
	event := new(MCDESMJoin)
	if err := _MCDESM.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDESMRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDESM contract.
type MCDESMRelyIterator struct {
	Event *MCDESMRely // Event containing the contract specifics and raw log

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
func (it *MCDESMRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDESMRely)
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
		it.Event = new(MCDESMRely)
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
func (it *MCDESMRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDESMRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDESMRely represents a Rely event raised by the MCDESM contract.
type MCDESMRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDESM *MCDESMFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDESMRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDESMRelyIterator{contract: _MCDESM.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDESM *MCDESMFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDESMRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDESM.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDESMRely)
				if err := _MCDESM.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDESM *MCDESMFilterer) ParseRely(log types.Log) (*MCDESMRely, error) {
	event := new(MCDESMRely)
	if err := _MCDESM.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
