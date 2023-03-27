// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Changelog

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

// CHANGELOGABI is the input ABI used to generate the binding from.
const CHANGELOGABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"RemoveAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"UpdateAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfs\",\"type\":\"string\"}],\"name\":\"UpdateIPFS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sha256sum\",\"type\":\"string\"}],\"name\":\"UpdateSha256sum\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateVersion\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"removeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"}],\"name\":\"setIPFS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sha256sum\",\"type\":\"string\"}],\"name\":\"setSha256sum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sha256sum\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CHANGELOG is an auto generated Go binding around an Ethereum contract.
type CHANGELOG struct {
	CHANGELOGCaller     // Read-only binding to the contract
	CHANGELOGTransactor // Write-only binding to the contract
	CHANGELOGFilterer   // Log filterer for contract events
}

// CHANGELOGCaller is an auto generated read-only Go binding around an Ethereum contract.
type CHANGELOGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CHANGELOGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CHANGELOGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CHANGELOGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CHANGELOGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CHANGELOGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CHANGELOGSession struct {
	Contract     *CHANGELOG        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CHANGELOGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CHANGELOGCallerSession struct {
	Contract *CHANGELOGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CHANGELOGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CHANGELOGTransactorSession struct {
	Contract     *CHANGELOGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CHANGELOGRaw is an auto generated low-level Go binding around an Ethereum contract.
type CHANGELOGRaw struct {
	Contract *CHANGELOG // Generic contract binding to access the raw methods on
}

// CHANGELOGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CHANGELOGCallerRaw struct {
	Contract *CHANGELOGCaller // Generic read-only contract binding to access the raw methods on
}

// CHANGELOGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CHANGELOGTransactorRaw struct {
	Contract *CHANGELOGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCHANGELOG creates a new instance of CHANGELOG, bound to a specific deployed contract.
func NewCHANGELOG(address common.Address, backend bind.ContractBackend) (*CHANGELOG, error) {
	contract, err := bindCHANGELOG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CHANGELOG{CHANGELOGCaller: CHANGELOGCaller{contract: contract}, CHANGELOGTransactor: CHANGELOGTransactor{contract: contract}, CHANGELOGFilterer: CHANGELOGFilterer{contract: contract}}, nil
}

// NewCHANGELOGCaller creates a new read-only instance of CHANGELOG, bound to a specific deployed contract.
func NewCHANGELOGCaller(address common.Address, caller bind.ContractCaller) (*CHANGELOGCaller, error) {
	contract, err := bindCHANGELOG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CHANGELOGCaller{contract: contract}, nil
}

// NewCHANGELOGTransactor creates a new write-only instance of CHANGELOG, bound to a specific deployed contract.
func NewCHANGELOGTransactor(address common.Address, transactor bind.ContractTransactor) (*CHANGELOGTransactor, error) {
	contract, err := bindCHANGELOG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CHANGELOGTransactor{contract: contract}, nil
}

// NewCHANGELOGFilterer creates a new log filterer instance of CHANGELOG, bound to a specific deployed contract.
func NewCHANGELOGFilterer(address common.Address, filterer bind.ContractFilterer) (*CHANGELOGFilterer, error) {
	contract, err := bindCHANGELOG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CHANGELOGFilterer{contract: contract}, nil
}

// bindCHANGELOG binds a generic wrapper to an already deployed contract.
func bindCHANGELOG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CHANGELOGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CHANGELOG *CHANGELOGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CHANGELOG.Contract.CHANGELOGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CHANGELOG *CHANGELOGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CHANGELOG.Contract.CHANGELOGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CHANGELOG *CHANGELOGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CHANGELOG.Contract.CHANGELOGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CHANGELOG *CHANGELOGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CHANGELOG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CHANGELOG *CHANGELOGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CHANGELOG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CHANGELOG *CHANGELOGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CHANGELOG.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CHANGELOG *CHANGELOGCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CHANGELOG *CHANGELOGSession) Count() (*big.Int, error) {
	return _CHANGELOG.Contract.Count(&_CHANGELOG.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CHANGELOG *CHANGELOGCallerSession) Count() (*big.Int, error) {
	return _CHANGELOG.Contract.Count(&_CHANGELOG.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32, address)
func (_CHANGELOG *CHANGELOGCaller) Get(opts *bind.CallOpts, _index *big.Int) ([32]byte, common.Address, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "get", _index)

	if err != nil {
		return *new([32]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32, address)
func (_CHANGELOG *CHANGELOGSession) Get(_index *big.Int) ([32]byte, common.Address, error) {
	return _CHANGELOG.Contract.Get(&_CHANGELOG.CallOpts, _index)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _index) view returns(bytes32, address)
func (_CHANGELOG *CHANGELOGCallerSession) Get(_index *big.Int) ([32]byte, common.Address, error) {
	return _CHANGELOG.Contract.Get(&_CHANGELOG.CallOpts, _index)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address addr)
func (_CHANGELOG *CHANGELOGCaller) GetAddress(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "getAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address addr)
func (_CHANGELOG *CHANGELOGSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _CHANGELOG.Contract.GetAddress(&_CHANGELOG.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address addr)
func (_CHANGELOG *CHANGELOGCallerSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _CHANGELOG.Contract.GetAddress(&_CHANGELOG.CallOpts, _key)
}

// Ipfs is a free data retrieval call binding the contract method 0xd7959cf9.
//
// Solidity: function ipfs() view returns(string)
func (_CHANGELOG *CHANGELOGCaller) Ipfs(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "ipfs")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Ipfs is a free data retrieval call binding the contract method 0xd7959cf9.
//
// Solidity: function ipfs() view returns(string)
func (_CHANGELOG *CHANGELOGSession) Ipfs() (string, error) {
	return _CHANGELOG.Contract.Ipfs(&_CHANGELOG.CallOpts)
}

// Ipfs is a free data retrieval call binding the contract method 0xd7959cf9.
//
// Solidity: function ipfs() view returns(string)
func (_CHANGELOG *CHANGELOGCallerSession) Ipfs() (string, error) {
	return _CHANGELOG.Contract.Ipfs(&_CHANGELOG.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(bytes32)
func (_CHANGELOG *CHANGELOGCaller) Keys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "keys", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(bytes32)
func (_CHANGELOG *CHANGELOGSession) Keys(arg0 *big.Int) ([32]byte, error) {
	return _CHANGELOG.Contract.Keys(&_CHANGELOG.CallOpts, arg0)
}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(bytes32)
func (_CHANGELOG *CHANGELOGCallerSession) Keys(arg0 *big.Int) ([32]byte, error) {
	return _CHANGELOG.Contract.Keys(&_CHANGELOG.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_CHANGELOG *CHANGELOGCaller) List(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_CHANGELOG *CHANGELOGSession) List() ([][32]byte, error) {
	return _CHANGELOG.Contract.List(&_CHANGELOG.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_CHANGELOG *CHANGELOGCallerSession) List() ([][32]byte, error) {
	return _CHANGELOG.Contract.List(&_CHANGELOG.CallOpts)
}

// Sha256sum is a free data retrieval call binding the contract method 0xf3295b93.
//
// Solidity: function sha256sum() view returns(string)
func (_CHANGELOG *CHANGELOGCaller) Sha256sum(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "sha256sum")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Sha256sum is a free data retrieval call binding the contract method 0xf3295b93.
//
// Solidity: function sha256sum() view returns(string)
func (_CHANGELOG *CHANGELOGSession) Sha256sum() (string, error) {
	return _CHANGELOG.Contract.Sha256sum(&_CHANGELOG.CallOpts)
}

// Sha256sum is a free data retrieval call binding the contract method 0xf3295b93.
//
// Solidity: function sha256sum() view returns(string)
func (_CHANGELOG *CHANGELOGCallerSession) Sha256sum() (string, error) {
	return _CHANGELOG.Contract.Sha256sum(&_CHANGELOG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CHANGELOG *CHANGELOGCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CHANGELOG *CHANGELOGSession) Version() (string, error) {
	return _CHANGELOG.Contract.Version(&_CHANGELOG.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CHANGELOG *CHANGELOGCallerSession) Version() (string, error) {
	return _CHANGELOG.Contract.Version(&_CHANGELOG.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_CHANGELOG *CHANGELOGCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CHANGELOG.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_CHANGELOG *CHANGELOGSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _CHANGELOG.Contract.Wards(&_CHANGELOG.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_CHANGELOG *CHANGELOGCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _CHANGELOG.Contract.Wards(&_CHANGELOG.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_CHANGELOG *CHANGELOGTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_CHANGELOG *CHANGELOGSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.Deny(&_CHANGELOG.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.Deny(&_CHANGELOG.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_CHANGELOG *CHANGELOGTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_CHANGELOG *CHANGELOGSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.Rely(&_CHANGELOG.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.Rely(&_CHANGELOG.TransactOpts, usr)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 _key) returns()
func (_CHANGELOG *CHANGELOGTransactor) RemoveAddress(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "removeAddress", _key)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 _key) returns()
func (_CHANGELOG *CHANGELOGSession) RemoveAddress(_key [32]byte) (*types.Transaction, error) {
	return _CHANGELOG.Contract.RemoveAddress(&_CHANGELOG.TransactOpts, _key)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 _key) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) RemoveAddress(_key [32]byte) (*types.Transaction, error) {
	return _CHANGELOG.Contract.RemoveAddress(&_CHANGELOG.TransactOpts, _key)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _addr) returns()
func (_CHANGELOG *CHANGELOGTransactor) SetAddress(opts *bind.TransactOpts, _key [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "setAddress", _key, _addr)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _addr) returns()
func (_CHANGELOG *CHANGELOGSession) SetAddress(_key [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetAddress(&_CHANGELOG.TransactOpts, _key, _addr)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _key, address _addr) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) SetAddress(_key [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetAddress(&_CHANGELOG.TransactOpts, _key, _addr)
}

// SetIPFS is a paid mutator transaction binding the contract method 0xd19ac77a.
//
// Solidity: function setIPFS(string _ipfs) returns()
func (_CHANGELOG *CHANGELOGTransactor) SetIPFS(opts *bind.TransactOpts, _ipfs string) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "setIPFS", _ipfs)
}

// SetIPFS is a paid mutator transaction binding the contract method 0xd19ac77a.
//
// Solidity: function setIPFS(string _ipfs) returns()
func (_CHANGELOG *CHANGELOGSession) SetIPFS(_ipfs string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetIPFS(&_CHANGELOG.TransactOpts, _ipfs)
}

// SetIPFS is a paid mutator transaction binding the contract method 0xd19ac77a.
//
// Solidity: function setIPFS(string _ipfs) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) SetIPFS(_ipfs string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetIPFS(&_CHANGELOG.TransactOpts, _ipfs)
}

// SetSha256sum is a paid mutator transaction binding the contract method 0x884a631f.
//
// Solidity: function setSha256sum(string _sha256sum) returns()
func (_CHANGELOG *CHANGELOGTransactor) SetSha256sum(opts *bind.TransactOpts, _sha256sum string) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "setSha256sum", _sha256sum)
}

// SetSha256sum is a paid mutator transaction binding the contract method 0x884a631f.
//
// Solidity: function setSha256sum(string _sha256sum) returns()
func (_CHANGELOG *CHANGELOGSession) SetSha256sum(_sha256sum string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetSha256sum(&_CHANGELOG.TransactOpts, _sha256sum)
}

// SetSha256sum is a paid mutator transaction binding the contract method 0x884a631f.
//
// Solidity: function setSha256sum(string _sha256sum) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) SetSha256sum(_sha256sum string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetSha256sum(&_CHANGELOG.TransactOpts, _sha256sum)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_CHANGELOG *CHANGELOGTransactor) SetVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _CHANGELOG.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_CHANGELOG *CHANGELOGSession) SetVersion(_version string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetVersion(&_CHANGELOG.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_CHANGELOG *CHANGELOGTransactorSession) SetVersion(_version string) (*types.Transaction, error) {
	return _CHANGELOG.Contract.SetVersion(&_CHANGELOG.TransactOpts, _version)
}

// CHANGELOGDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the CHANGELOG contract.
type CHANGELOGDenyIterator struct {
	Event *CHANGELOGDeny // Event containing the contract specifics and raw log

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
func (it *CHANGELOGDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGDeny)
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
		it.Event = new(CHANGELOGDeny)
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
func (it *CHANGELOGDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGDeny represents a Deny event raised by the CHANGELOG contract.
type CHANGELOGDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_CHANGELOG *CHANGELOGFilterer) FilterDeny(opts *bind.FilterOpts) (*CHANGELOGDenyIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGDenyIterator{contract: _CHANGELOG.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_CHANGELOG *CHANGELOGFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *CHANGELOGDeny) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGDeny)
				if err := _CHANGELOG.contract.UnpackLog(event, "Deny", log); err != nil {
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
// Solidity: event Deny(address usr)
func (_CHANGELOG *CHANGELOGFilterer) ParseDeny(log types.Log) (*CHANGELOGDeny, error) {
	event := new(CHANGELOGDeny)
	if err := _CHANGELOG.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the CHANGELOG contract.
type CHANGELOGRelyIterator struct {
	Event *CHANGELOGRely // Event containing the contract specifics and raw log

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
func (it *CHANGELOGRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGRely)
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
		it.Event = new(CHANGELOGRely)
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
func (it *CHANGELOGRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGRely represents a Rely event raised by the CHANGELOG contract.
type CHANGELOGRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_CHANGELOG *CHANGELOGFilterer) FilterRely(opts *bind.FilterOpts) (*CHANGELOGRelyIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGRelyIterator{contract: _CHANGELOG.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_CHANGELOG *CHANGELOGFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *CHANGELOGRely) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGRely)
				if err := _CHANGELOG.contract.UnpackLog(event, "Rely", log); err != nil {
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
// Solidity: event Rely(address usr)
func (_CHANGELOG *CHANGELOGFilterer) ParseRely(log types.Log) (*CHANGELOGRely, error) {
	event := new(CHANGELOGRely)
	if err := _CHANGELOG.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGRemoveAddressIterator is returned from FilterRemoveAddress and is used to iterate over the raw logs and unpacked data for RemoveAddress events raised by the CHANGELOG contract.
type CHANGELOGRemoveAddressIterator struct {
	Event *CHANGELOGRemoveAddress // Event containing the contract specifics and raw log

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
func (it *CHANGELOGRemoveAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGRemoveAddress)
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
		it.Event = new(CHANGELOGRemoveAddress)
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
func (it *CHANGELOGRemoveAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGRemoveAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGRemoveAddress represents a RemoveAddress event raised by the CHANGELOG contract.
type CHANGELOGRemoveAddress struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveAddress is a free log retrieval operation binding the contract event 0x236084b20b5721ddf3e40ed45b570186612a08edec64abae6894b60f315c54bc.
//
// Solidity: event RemoveAddress(bytes32 key)
func (_CHANGELOG *CHANGELOGFilterer) FilterRemoveAddress(opts *bind.FilterOpts) (*CHANGELOGRemoveAddressIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "RemoveAddress")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGRemoveAddressIterator{contract: _CHANGELOG.contract, event: "RemoveAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveAddress is a free log subscription operation binding the contract event 0x236084b20b5721ddf3e40ed45b570186612a08edec64abae6894b60f315c54bc.
//
// Solidity: event RemoveAddress(bytes32 key)
func (_CHANGELOG *CHANGELOGFilterer) WatchRemoveAddress(opts *bind.WatchOpts, sink chan<- *CHANGELOGRemoveAddress) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "RemoveAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGRemoveAddress)
				if err := _CHANGELOG.contract.UnpackLog(event, "RemoveAddress", log); err != nil {
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

// ParseRemoveAddress is a log parse operation binding the contract event 0x236084b20b5721ddf3e40ed45b570186612a08edec64abae6894b60f315c54bc.
//
// Solidity: event RemoveAddress(bytes32 key)
func (_CHANGELOG *CHANGELOGFilterer) ParseRemoveAddress(log types.Log) (*CHANGELOGRemoveAddress, error) {
	event := new(CHANGELOGRemoveAddress)
	if err := _CHANGELOG.contract.UnpackLog(event, "RemoveAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGUpdateAddressIterator is returned from FilterUpdateAddress and is used to iterate over the raw logs and unpacked data for UpdateAddress events raised by the CHANGELOG contract.
type CHANGELOGUpdateAddressIterator struct {
	Event *CHANGELOGUpdateAddress // Event containing the contract specifics and raw log

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
func (it *CHANGELOGUpdateAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGUpdateAddress)
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
		it.Event = new(CHANGELOGUpdateAddress)
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
func (it *CHANGELOGUpdateAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGUpdateAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGUpdateAddress represents a UpdateAddress event raised by the CHANGELOG contract.
type CHANGELOGUpdateAddress struct {
	Key  [32]byte
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAddress is a free log retrieval operation binding the contract event 0xcf0ecd10d1399d98978051da48095f368d6f1fa3292d5fc09135b92aa2a4d733.
//
// Solidity: event UpdateAddress(bytes32 key, address addr)
func (_CHANGELOG *CHANGELOGFilterer) FilterUpdateAddress(opts *bind.FilterOpts) (*CHANGELOGUpdateAddressIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "UpdateAddress")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGUpdateAddressIterator{contract: _CHANGELOG.contract, event: "UpdateAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateAddress is a free log subscription operation binding the contract event 0xcf0ecd10d1399d98978051da48095f368d6f1fa3292d5fc09135b92aa2a4d733.
//
// Solidity: event UpdateAddress(bytes32 key, address addr)
func (_CHANGELOG *CHANGELOGFilterer) WatchUpdateAddress(opts *bind.WatchOpts, sink chan<- *CHANGELOGUpdateAddress) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "UpdateAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGUpdateAddress)
				if err := _CHANGELOG.contract.UnpackLog(event, "UpdateAddress", log); err != nil {
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

// ParseUpdateAddress is a log parse operation binding the contract event 0xcf0ecd10d1399d98978051da48095f368d6f1fa3292d5fc09135b92aa2a4d733.
//
// Solidity: event UpdateAddress(bytes32 key, address addr)
func (_CHANGELOG *CHANGELOGFilterer) ParseUpdateAddress(log types.Log) (*CHANGELOGUpdateAddress, error) {
	event := new(CHANGELOGUpdateAddress)
	if err := _CHANGELOG.contract.UnpackLog(event, "UpdateAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGUpdateIPFSIterator is returned from FilterUpdateIPFS and is used to iterate over the raw logs and unpacked data for UpdateIPFS events raised by the CHANGELOG contract.
type CHANGELOGUpdateIPFSIterator struct {
	Event *CHANGELOGUpdateIPFS // Event containing the contract specifics and raw log

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
func (it *CHANGELOGUpdateIPFSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGUpdateIPFS)
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
		it.Event = new(CHANGELOGUpdateIPFS)
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
func (it *CHANGELOGUpdateIPFSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGUpdateIPFSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGUpdateIPFS represents a UpdateIPFS event raised by the CHANGELOG contract.
type CHANGELOGUpdateIPFS struct {
	Ipfs string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateIPFS is a free log retrieval operation binding the contract event 0xc7d51f843aa4e748744967527c18b55884fbbd978794939a53e6de466659169e.
//
// Solidity: event UpdateIPFS(string ipfs)
func (_CHANGELOG *CHANGELOGFilterer) FilterUpdateIPFS(opts *bind.FilterOpts) (*CHANGELOGUpdateIPFSIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "UpdateIPFS")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGUpdateIPFSIterator{contract: _CHANGELOG.contract, event: "UpdateIPFS", logs: logs, sub: sub}, nil
}

// WatchUpdateIPFS is a free log subscription operation binding the contract event 0xc7d51f843aa4e748744967527c18b55884fbbd978794939a53e6de466659169e.
//
// Solidity: event UpdateIPFS(string ipfs)
func (_CHANGELOG *CHANGELOGFilterer) WatchUpdateIPFS(opts *bind.WatchOpts, sink chan<- *CHANGELOGUpdateIPFS) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "UpdateIPFS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGUpdateIPFS)
				if err := _CHANGELOG.contract.UnpackLog(event, "UpdateIPFS", log); err != nil {
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

// ParseUpdateIPFS is a log parse operation binding the contract event 0xc7d51f843aa4e748744967527c18b55884fbbd978794939a53e6de466659169e.
//
// Solidity: event UpdateIPFS(string ipfs)
func (_CHANGELOG *CHANGELOGFilterer) ParseUpdateIPFS(log types.Log) (*CHANGELOGUpdateIPFS, error) {
	event := new(CHANGELOGUpdateIPFS)
	if err := _CHANGELOG.contract.UnpackLog(event, "UpdateIPFS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGUpdateSha256sumIterator is returned from FilterUpdateSha256sum and is used to iterate over the raw logs and unpacked data for UpdateSha256sum events raised by the CHANGELOG contract.
type CHANGELOGUpdateSha256sumIterator struct {
	Event *CHANGELOGUpdateSha256sum // Event containing the contract specifics and raw log

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
func (it *CHANGELOGUpdateSha256sumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGUpdateSha256sum)
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
		it.Event = new(CHANGELOGUpdateSha256sum)
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
func (it *CHANGELOGUpdateSha256sumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGUpdateSha256sumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGUpdateSha256sum represents a UpdateSha256sum event raised by the CHANGELOG contract.
type CHANGELOGUpdateSha256sum struct {
	Sha256sum string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateSha256sum is a free log retrieval operation binding the contract event 0x6ce764b6ccabce70dd110fb75895ead7988d54d58fef7e8ee13a32d04c21f3b9.
//
// Solidity: event UpdateSha256sum(string sha256sum)
func (_CHANGELOG *CHANGELOGFilterer) FilterUpdateSha256sum(opts *bind.FilterOpts) (*CHANGELOGUpdateSha256sumIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "UpdateSha256sum")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGUpdateSha256sumIterator{contract: _CHANGELOG.contract, event: "UpdateSha256sum", logs: logs, sub: sub}, nil
}

// WatchUpdateSha256sum is a free log subscription operation binding the contract event 0x6ce764b6ccabce70dd110fb75895ead7988d54d58fef7e8ee13a32d04c21f3b9.
//
// Solidity: event UpdateSha256sum(string sha256sum)
func (_CHANGELOG *CHANGELOGFilterer) WatchUpdateSha256sum(opts *bind.WatchOpts, sink chan<- *CHANGELOGUpdateSha256sum) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "UpdateSha256sum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGUpdateSha256sum)
				if err := _CHANGELOG.contract.UnpackLog(event, "UpdateSha256sum", log); err != nil {
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

// ParseUpdateSha256sum is a log parse operation binding the contract event 0x6ce764b6ccabce70dd110fb75895ead7988d54d58fef7e8ee13a32d04c21f3b9.
//
// Solidity: event UpdateSha256sum(string sha256sum)
func (_CHANGELOG *CHANGELOGFilterer) ParseUpdateSha256sum(log types.Log) (*CHANGELOGUpdateSha256sum, error) {
	event := new(CHANGELOGUpdateSha256sum)
	if err := _CHANGELOG.contract.UnpackLog(event, "UpdateSha256sum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CHANGELOGUpdateVersionIterator is returned from FilterUpdateVersion and is used to iterate over the raw logs and unpacked data for UpdateVersion events raised by the CHANGELOG contract.
type CHANGELOGUpdateVersionIterator struct {
	Event *CHANGELOGUpdateVersion // Event containing the contract specifics and raw log

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
func (it *CHANGELOGUpdateVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CHANGELOGUpdateVersion)
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
		it.Event = new(CHANGELOGUpdateVersion)
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
func (it *CHANGELOGUpdateVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CHANGELOGUpdateVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CHANGELOGUpdateVersion represents a UpdateVersion event raised by the CHANGELOG contract.
type CHANGELOGUpdateVersion struct {
	Version string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateVersion is a free log retrieval operation binding the contract event 0xc96e830455c6cd482ced2b3637de7e7f3c87ed0f17eff0999ba5e262ea44c67b.
//
// Solidity: event UpdateVersion(string version)
func (_CHANGELOG *CHANGELOGFilterer) FilterUpdateVersion(opts *bind.FilterOpts) (*CHANGELOGUpdateVersionIterator, error) {

	logs, sub, err := _CHANGELOG.contract.FilterLogs(opts, "UpdateVersion")
	if err != nil {
		return nil, err
	}
	return &CHANGELOGUpdateVersionIterator{contract: _CHANGELOG.contract, event: "UpdateVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateVersion is a free log subscription operation binding the contract event 0xc96e830455c6cd482ced2b3637de7e7f3c87ed0f17eff0999ba5e262ea44c67b.
//
// Solidity: event UpdateVersion(string version)
func (_CHANGELOG *CHANGELOGFilterer) WatchUpdateVersion(opts *bind.WatchOpts, sink chan<- *CHANGELOGUpdateVersion) (event.Subscription, error) {

	logs, sub, err := _CHANGELOG.contract.WatchLogs(opts, "UpdateVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CHANGELOGUpdateVersion)
				if err := _CHANGELOG.contract.UnpackLog(event, "UpdateVersion", log); err != nil {
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

// ParseUpdateVersion is a log parse operation binding the contract event 0xc96e830455c6cd482ced2b3637de7e7f3c87ed0f17eff0999ba5e262ea44c67b.
//
// Solidity: event UpdateVersion(string version)
func (_CHANGELOG *CHANGELOGFilterer) ParseUpdateVersion(log types.Log) (*CHANGELOGUpdateVersion, error) {
	event := new(CHANGELOGUpdateVersion)
	if err := _CHANGELOG.contract.UnpackLog(event, "UpdateVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
