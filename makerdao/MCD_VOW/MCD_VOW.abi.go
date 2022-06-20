// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_VOW

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

// MCDVOWABI is the input ABI used to generate the binding from.
const MCDVOWABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flapper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flopper_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"fess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"flap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flapper\",\"outputs\":[{\"internalType\":\"contractFlapLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"}],\"name\":\"flog\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"flop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flopper\",\"outputs\":[{\"internalType\":\"contractFlopLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"kiss\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDVOW is an auto generated Go binding around an Ethereum contract.
type MCDVOW struct {
	MCDVOWCaller     // Read-only binding to the contract
	MCDVOWTransactor // Write-only binding to the contract
	MCDVOWFilterer   // Log filterer for contract events
}

// MCDVOWCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDVOWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVOWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDVOWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVOWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDVOWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVOWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDVOWSession struct {
	Contract     *MCDVOW           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDVOWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDVOWCallerSession struct {
	Contract *MCDVOWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDVOWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDVOWTransactorSession struct {
	Contract     *MCDVOWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDVOWRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDVOWRaw struct {
	Contract *MCDVOW // Generic contract binding to access the raw methods on
}

// MCDVOWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDVOWCallerRaw struct {
	Contract *MCDVOWCaller // Generic read-only contract binding to access the raw methods on
}

// MCDVOWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDVOWTransactorRaw struct {
	Contract *MCDVOWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDVOW creates a new instance of MCDVOW, bound to a specific deployed contract.
func NewMCDVOW(address common.Address, backend bind.ContractBackend) (*MCDVOW, error) {
	contract, err := bindMCDVOW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDVOW{MCDVOWCaller: MCDVOWCaller{contract: contract}, MCDVOWTransactor: MCDVOWTransactor{contract: contract}, MCDVOWFilterer: MCDVOWFilterer{contract: contract}}, nil
}

// NewMCDVOWCaller creates a new read-only instance of MCDVOW, bound to a specific deployed contract.
func NewMCDVOWCaller(address common.Address, caller bind.ContractCaller) (*MCDVOWCaller, error) {
	contract, err := bindMCDVOW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVOWCaller{contract: contract}, nil
}

// NewMCDVOWTransactor creates a new write-only instance of MCDVOW, bound to a specific deployed contract.
func NewMCDVOWTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDVOWTransactor, error) {
	contract, err := bindMCDVOW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVOWTransactor{contract: contract}, nil
}

// NewMCDVOWFilterer creates a new log filterer instance of MCDVOW, bound to a specific deployed contract.
func NewMCDVOWFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDVOWFilterer, error) {
	contract, err := bindMCDVOW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDVOWFilterer{contract: contract}, nil
}

// bindMCDVOW binds a generic wrapper to an already deployed contract.
func bindMCDVOW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDVOWABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVOW *MCDVOWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVOW.Contract.MCDVOWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVOW *MCDVOWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVOW.Contract.MCDVOWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVOW *MCDVOWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVOW.Contract.MCDVOWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVOW *MCDVOWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVOW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVOW *MCDVOWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVOW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVOW *MCDVOWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVOW.Contract.contract.Transact(opts, method, params...)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Ash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "Ash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Ash() (*big.Int, error) {
	return _MCDVOW.Contract.Ash(&_MCDVOW.CallOpts)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Ash() (*big.Int, error) {
	return _MCDVOW.Contract.Ash(&_MCDVOW.CallOpts)
}

// Sin2 is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Sin2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "Sin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin2 is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Sin2() (*big.Int, error) {
	return _MCDVOW.Contract.Sin2(&_MCDVOW.CallOpts)
}

// Sin2 is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Sin2() (*big.Int, error) {
	return _MCDVOW.Contract.Sin2(&_MCDVOW.CallOpts)
}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Bump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "bump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Bump() (*big.Int, error) {
	return _MCDVOW.Contract.Bump(&_MCDVOW.CallOpts)
}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Bump() (*big.Int, error) {
	return _MCDVOW.Contract.Bump(&_MCDVOW.CallOpts)
}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Dump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "dump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Dump() (*big.Int, error) {
	return _MCDVOW.Contract.Dump(&_MCDVOW.CallOpts)
}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Dump() (*big.Int, error) {
	return _MCDVOW.Contract.Dump(&_MCDVOW.CallOpts)
}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_MCDVOW *MCDVOWCaller) Flapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "flapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_MCDVOW *MCDVOWSession) Flapper() (common.Address, error) {
	return _MCDVOW.Contract.Flapper(&_MCDVOW.CallOpts)
}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_MCDVOW *MCDVOWCallerSession) Flapper() (common.Address, error) {
	return _MCDVOW.Contract.Flapper(&_MCDVOW.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_MCDVOW *MCDVOWCaller) Flopper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "flopper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_MCDVOW *MCDVOWSession) Flopper() (common.Address, error) {
	return _MCDVOW.Contract.Flopper(&_MCDVOW.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_MCDVOW *MCDVOWCallerSession) Flopper() (common.Address, error) {
	return _MCDVOW.Contract.Flopper(&_MCDVOW.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Hump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "hump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Hump() (*big.Int, error) {
	return _MCDVOW.Contract.Hump(&_MCDVOW.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Hump() (*big.Int, error) {
	return _MCDVOW.Contract.Hump(&_MCDVOW.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Live() (*big.Int, error) {
	return _MCDVOW.Contract.Live(&_MCDVOW.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Live() (*big.Int, error) {
	return _MCDVOW.Contract.Live(&_MCDVOW.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Sin(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_MCDVOW *MCDVOWSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _MCDVOW.Contract.Sin(&_MCDVOW.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _MCDVOW.Contract.Sin(&_MCDVOW.CallOpts, arg0)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Sump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "sump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Sump() (*big.Int, error) {
	return _MCDVOW.Contract.Sump(&_MCDVOW.CallOpts)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Sump() (*big.Int, error) {
	return _MCDVOW.Contract.Sump(&_MCDVOW.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVOW *MCDVOWCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVOW *MCDVOWSession) Vat() (common.Address, error) {
	return _MCDVOW.Contract.Vat(&_MCDVOW.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVOW *MCDVOWCallerSession) Vat() (common.Address, error) {
	return _MCDVOW.Contract.Vat(&_MCDVOW.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Wait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "wait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDVOW *MCDVOWSession) Wait() (*big.Int, error) {
	return _MCDVOW.Contract.Wait(&_MCDVOW.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Wait() (*big.Int, error) {
	return _MCDVOW.Contract.Wait(&_MCDVOW.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVOW *MCDVOWCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDVOW.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVOW *MCDVOWSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVOW.Contract.Wards(&_MCDVOW.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVOW *MCDVOWCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVOW.Contract.Wards(&_MCDVOW.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDVOW *MCDVOWTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDVOW *MCDVOWSession) Cage() (*types.Transaction, error) {
	return _MCDVOW.Contract.Cage(&_MCDVOW.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDVOW *MCDVOWTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDVOW.Contract.Cage(&_MCDVOW.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVOW *MCDVOWTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVOW *MCDVOWSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.Deny(&_MCDVOW.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVOW *MCDVOWTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.Deny(&_MCDVOW.TransactOpts, usr)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_MCDVOW *MCDVOWTransactor) Fess(opts *bind.TransactOpts, tab *big.Int) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "fess", tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_MCDVOW *MCDVOWSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Fess(&_MCDVOW.TransactOpts, tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_MCDVOW *MCDVOWTransactorSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Fess(&_MCDVOW.TransactOpts, tab)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVOW *MCDVOWTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVOW *MCDVOWSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.File(&_MCDVOW.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVOW *MCDVOWTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.File(&_MCDVOW.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDVOW *MCDVOWTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDVOW *MCDVOWSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.File0(&_MCDVOW.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDVOW *MCDVOWTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.File0(&_MCDVOW.TransactOpts, what, data)
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_MCDVOW *MCDVOWTransactor) Flap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "flap")
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_MCDVOW *MCDVOWSession) Flap() (*types.Transaction, error) {
	return _MCDVOW.Contract.Flap(&_MCDVOW.TransactOpts)
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_MCDVOW *MCDVOWTransactorSession) Flap() (*types.Transaction, error) {
	return _MCDVOW.Contract.Flap(&_MCDVOW.TransactOpts)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_MCDVOW *MCDVOWTransactor) Flog(opts *bind.TransactOpts, era *big.Int) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "flog", era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_MCDVOW *MCDVOWSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Flog(&_MCDVOW.TransactOpts, era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_MCDVOW *MCDVOWTransactorSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Flog(&_MCDVOW.TransactOpts, era)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_MCDVOW *MCDVOWTransactor) Flop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "flop")
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_MCDVOW *MCDVOWSession) Flop() (*types.Transaction, error) {
	return _MCDVOW.Contract.Flop(&_MCDVOW.TransactOpts)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_MCDVOW *MCDVOWTransactorSession) Flop() (*types.Transaction, error) {
	return _MCDVOW.Contract.Flop(&_MCDVOW.TransactOpts)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MCDVOW *MCDVOWTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MCDVOW *MCDVOWSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Heal(&_MCDVOW.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MCDVOW *MCDVOWTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Heal(&_MCDVOW.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_MCDVOW *MCDVOWTransactor) Kiss(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "kiss", rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_MCDVOW *MCDVOWSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Kiss(&_MCDVOW.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_MCDVOW *MCDVOWTransactorSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _MCDVOW.Contract.Kiss(&_MCDVOW.TransactOpts, rad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVOW *MCDVOWTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVOW *MCDVOWSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.Rely(&_MCDVOW.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVOW *MCDVOWTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVOW.Contract.Rely(&_MCDVOW.TransactOpts, usr)
}
