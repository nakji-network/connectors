// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDJug

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

// MCDJUGABI is the input ABI used to generate the binding from.
const MCDJUGABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rho\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDJUG is an auto generated Go binding around an Ethereum contract.
type MCDJUG struct {
	MCDJUGCaller     // Read-only binding to the contract
	MCDJUGTransactor // Write-only binding to the contract
	MCDJUGFilterer   // Log filterer for contract events
}

// MCDJUGCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDJUGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDJUGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDJUGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDJUGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDJUGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDJUGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDJUGSession struct {
	Contract     *MCDJUG           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDJUGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDJUGCallerSession struct {
	Contract *MCDJUGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDJUGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDJUGTransactorSession struct {
	Contract     *MCDJUGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDJUGRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDJUGRaw struct {
	Contract *MCDJUG // Generic contract binding to access the raw methods on
}

// MCDJUGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDJUGCallerRaw struct {
	Contract *MCDJUGCaller // Generic read-only contract binding to access the raw methods on
}

// MCDJUGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDJUGTransactorRaw struct {
	Contract *MCDJUGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDJUG creates a new instance of MCDJUG, bound to a specific deployed contract.
func NewMCDJUG(address common.Address, backend bind.ContractBackend) (*MCDJUG, error) {
	contract, err := bindMCDJUG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDJUG{MCDJUGCaller: MCDJUGCaller{contract: contract}, MCDJUGTransactor: MCDJUGTransactor{contract: contract}, MCDJUGFilterer: MCDJUGFilterer{contract: contract}}, nil
}

// NewMCDJUGCaller creates a new read-only instance of MCDJUG, bound to a specific deployed contract.
func NewMCDJUGCaller(address common.Address, caller bind.ContractCaller) (*MCDJUGCaller, error) {
	contract, err := bindMCDJUG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDJUGCaller{contract: contract}, nil
}

// NewMCDJUGTransactor creates a new write-only instance of MCDJUG, bound to a specific deployed contract.
func NewMCDJUGTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDJUGTransactor, error) {
	contract, err := bindMCDJUG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDJUGTransactor{contract: contract}, nil
}

// NewMCDJUGFilterer creates a new log filterer instance of MCDJUG, bound to a specific deployed contract.
func NewMCDJUGFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDJUGFilterer, error) {
	contract, err := bindMCDJUG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDJUGFilterer{contract: contract}, nil
}

// bindMCDJUG binds a generic wrapper to an already deployed contract.
func bindMCDJUG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDJUGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDJUG *MCDJUGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDJUG.Contract.MCDJUGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDJUG *MCDJUGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDJUG.Contract.MCDJUGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDJUG *MCDJUGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDJUG.Contract.MCDJUGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDJUG *MCDJUGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDJUG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDJUG *MCDJUGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDJUG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDJUG *MCDJUGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDJUG.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_MCDJUG *MCDJUGCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDJUG.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_MCDJUG *MCDJUGSession) Base() (*big.Int, error) {
	return _MCDJUG.Contract.Base(&_MCDJUG.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_MCDJUG *MCDJUGCallerSession) Base() (*big.Int, error) {
	return _MCDJUG.Contract.Base(&_MCDJUG.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_MCDJUG *MCDJUGCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	var out []interface{}
	err := _MCDJUG.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Duty *big.Int
		Rho  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Duty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rho = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_MCDJUG *MCDJUGSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _MCDJUG.Contract.Ilks(&_MCDJUG.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_MCDJUG *MCDJUGCallerSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _MCDJUG.Contract.Ilks(&_MCDJUG.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDJUG *MCDJUGCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDJUG.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDJUG *MCDJUGSession) Vat() (common.Address, error) {
	return _MCDJUG.Contract.Vat(&_MCDJUG.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDJUG *MCDJUGCallerSession) Vat() (common.Address, error) {
	return _MCDJUG.Contract.Vat(&_MCDJUG.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDJUG *MCDJUGCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDJUG.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDJUG *MCDJUGSession) Vow() (common.Address, error) {
	return _MCDJUG.Contract.Vow(&_MCDJUG.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDJUG *MCDJUGCallerSession) Vow() (common.Address, error) {
	return _MCDJUG.Contract.Vow(&_MCDJUG.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDJUG *MCDJUGCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDJUG.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDJUG *MCDJUGSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDJUG.Contract.Wards(&_MCDJUG.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDJUG *MCDJUGCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDJUG.Contract.Wards(&_MCDJUG.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDJUG *MCDJUGTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDJUG *MCDJUGSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.Deny(&_MCDJUG.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDJUG *MCDJUGTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.Deny(&_MCDJUG.TransactOpts, usr)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_MCDJUG *MCDJUGTransactor) Drip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "drip", ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_MCDJUG *MCDJUGSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.Contract.Drip(&_MCDJUG.TransactOpts, ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_MCDJUG *MCDJUGTransactorSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.Contract.Drip(&_MCDJUG.TransactOpts, ilk)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.Contract.File(&_MCDJUG.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.Contract.File(&_MCDJUG.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.Contract.File0(&_MCDJUG.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDJUG *MCDJUGTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDJUG.Contract.File0(&_MCDJUG.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDJUG *MCDJUGTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDJUG *MCDJUGSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.File1(&_MCDJUG.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDJUG *MCDJUGTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.File1(&_MCDJUG.TransactOpts, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MCDJUG *MCDJUGTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MCDJUG *MCDJUGSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.Contract.Init(&_MCDJUG.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MCDJUG *MCDJUGTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _MCDJUG.Contract.Init(&_MCDJUG.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDJUG *MCDJUGTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDJUG *MCDJUGSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.Rely(&_MCDJUG.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDJUG *MCDJUGTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDJUG.Contract.Rely(&_MCDJUG.TransactOpts, usr)
}
