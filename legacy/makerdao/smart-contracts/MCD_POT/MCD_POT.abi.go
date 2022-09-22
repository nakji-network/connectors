// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_POT

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

// MCDPOTABI is the input ABI used to generate the binding from.
const MCDPOTABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tmp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dsr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rho\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDPOT is an auto generated Go binding around an Ethereum contract.
type MCDPOT struct {
	MCDPOTCaller     // Read-only binding to the contract
	MCDPOTTransactor // Write-only binding to the contract
	MCDPOTFilterer   // Log filterer for contract events
}

// MCDPOTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPOTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPOTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPOTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPOTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPOTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPOTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPOTSession struct {
	Contract     *MCDPOT           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPOTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPOTCallerSession struct {
	Contract *MCDPOTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDPOTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPOTTransactorSession struct {
	Contract     *MCDPOTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPOTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPOTRaw struct {
	Contract *MCDPOT // Generic contract binding to access the raw methods on
}

// MCDPOTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPOTCallerRaw struct {
	Contract *MCDPOTCaller // Generic read-only contract binding to access the raw methods on
}

// MCDPOTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPOTTransactorRaw struct {
	Contract *MCDPOTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPOT creates a new instance of MCDPOT, bound to a specific deployed contract.
func NewMCDPOT(address common.Address, backend bind.ContractBackend) (*MCDPOT, error) {
	contract, err := bindMCDPOT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPOT{MCDPOTCaller: MCDPOTCaller{contract: contract}, MCDPOTTransactor: MCDPOTTransactor{contract: contract}, MCDPOTFilterer: MCDPOTFilterer{contract: contract}}, nil
}

// NewMCDPOTCaller creates a new read-only instance of MCDPOT, bound to a specific deployed contract.
func NewMCDPOTCaller(address common.Address, caller bind.ContractCaller) (*MCDPOTCaller, error) {
	contract, err := bindMCDPOT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPOTCaller{contract: contract}, nil
}

// NewMCDPOTTransactor creates a new write-only instance of MCDPOT, bound to a specific deployed contract.
func NewMCDPOTTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPOTTransactor, error) {
	contract, err := bindMCDPOT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPOTTransactor{contract: contract}, nil
}

// NewMCDPOTFilterer creates a new log filterer instance of MCDPOT, bound to a specific deployed contract.
func NewMCDPOTFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPOTFilterer, error) {
	contract, err := bindMCDPOT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPOTFilterer{contract: contract}, nil
}

// bindMCDPOT binds a generic wrapper to an already deployed contract.
func bindMCDPOT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPOTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPOT *MCDPOTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPOT.Contract.MCDPOTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPOT *MCDPOTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPOT.Contract.MCDPOTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPOT *MCDPOTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPOT.Contract.MCDPOTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPOT *MCDPOTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPOT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPOT *MCDPOTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPOT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPOT *MCDPOTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPOT.Contract.contract.Transact(opts, method, params...)
}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Pie(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "Pie")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_MCDPOT *MCDPOTSession) Pie() (*big.Int, error) {
	return _MCDPOT.Contract.Pie(&_MCDPOT.CallOpts)
}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Pie() (*big.Int, error) {
	return _MCDPOT.Contract.Pie(&_MCDPOT.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Chi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "chi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_MCDPOT *MCDPOTSession) Chi() (*big.Int, error) {
	return _MCDPOT.Contract.Chi(&_MCDPOT.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Chi() (*big.Int, error) {
	return _MCDPOT.Contract.Chi(&_MCDPOT.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Dsr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "dsr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_MCDPOT *MCDPOTSession) Dsr() (*big.Int, error) {
	return _MCDPOT.Contract.Dsr(&_MCDPOT.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Dsr() (*big.Int, error) {
	return _MCDPOT.Contract.Dsr(&_MCDPOT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDPOT *MCDPOTSession) Live() (*big.Int, error) {
	return _MCDPOT.Contract.Live(&_MCDPOT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Live() (*big.Int, error) {
	return _MCDPOT.Contract.Live(&_MCDPOT.CallOpts)
}

// Pie2 is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Pie2(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "pie", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pie2 is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) view returns(uint256)
func (_MCDPOT *MCDPOTSession) Pie2(arg0 common.Address) (*big.Int, error) {
	return _MCDPOT.Contract.Pie2(&_MCDPOT.CallOpts, arg0)
}

// Pie2 is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Pie2(arg0 common.Address) (*big.Int, error) {
	return _MCDPOT.Contract.Pie2(&_MCDPOT.CallOpts, arg0)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Rho(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "rho")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_MCDPOT *MCDPOTSession) Rho() (*big.Int, error) {
	return _MCDPOT.Contract.Rho(&_MCDPOT.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Rho() (*big.Int, error) {
	return _MCDPOT.Contract.Rho(&_MCDPOT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPOT *MCDPOTCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPOT *MCDPOTSession) Vat() (common.Address, error) {
	return _MCDPOT.Contract.Vat(&_MCDPOT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPOT *MCDPOTCallerSession) Vat() (common.Address, error) {
	return _MCDPOT.Contract.Vat(&_MCDPOT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPOT *MCDPOTCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPOT *MCDPOTSession) Vow() (common.Address, error) {
	return _MCDPOT.Contract.Vow(&_MCDPOT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPOT *MCDPOTCallerSession) Vow() (common.Address, error) {
	return _MCDPOT.Contract.Vow(&_MCDPOT.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPOT *MCDPOTCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDPOT.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPOT *MCDPOTSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPOT.Contract.Wards(&_MCDPOT.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPOT *MCDPOTCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPOT.Contract.Wards(&_MCDPOT.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDPOT *MCDPOTTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDPOT *MCDPOTSession) Cage() (*types.Transaction, error) {
	return _MCDPOT.Contract.Cage(&_MCDPOT.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDPOT *MCDPOTTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDPOT.Contract.Cage(&_MCDPOT.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDPOT *MCDPOTTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDPOT *MCDPOTSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.Deny(&_MCDPOT.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDPOT *MCDPOTTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.Deny(&_MCDPOT.TransactOpts, guy)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_MCDPOT *MCDPOTTransactor) Drip(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "drip")
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_MCDPOT *MCDPOTSession) Drip() (*types.Transaction, error) {
	return _MCDPOT.Contract.Drip(&_MCDPOT.TransactOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_MCDPOT *MCDPOTTransactorSession) Drip() (*types.Transaction, error) {
	return _MCDPOT.Contract.Drip(&_MCDPOT.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_MCDPOT *MCDPOTTransactor) Exit(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "exit", wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_MCDPOT *MCDPOTSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.Exit(&_MCDPOT.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_MCDPOT *MCDPOTTransactorSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.Exit(&_MCDPOT.TransactOpts, wad)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPOT *MCDPOTTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPOT *MCDPOTSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.File(&_MCDPOT.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPOT *MCDPOTTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.File(&_MCDPOT.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_MCDPOT *MCDPOTTransactor) File0(opts *bind.TransactOpts, what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "file0", what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_MCDPOT *MCDPOTSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.File0(&_MCDPOT.TransactOpts, what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_MCDPOT *MCDPOTTransactorSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.File0(&_MCDPOT.TransactOpts, what, addr)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDPOT *MCDPOTTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDPOT *MCDPOTSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.Join(&_MCDPOT.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_MCDPOT *MCDPOTTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _MCDPOT.Contract.Join(&_MCDPOT.TransactOpts, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDPOT *MCDPOTTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDPOT *MCDPOTSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.Rely(&_MCDPOT.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDPOT *MCDPOTTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _MCDPOT.Contract.Rely(&_MCDPOT.TransactOpts, guy)
}
