// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RWAUrn

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

// RWAURNABI is the input ABI used to generate the binding from.
const RWAURNABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputConduit_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Draw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Free\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Hope\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Nope\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Quit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Wipe\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gemJoin\",\"outputs\":[{\"internalType\":\"contractGemJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"jug\",\"outputs\":[{\"internalType\":\"contractJugAbstract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"outputConduit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"quit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatAbstract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RWAURN is an auto generated Go binding around an Ethereum contract.
type RWAURN struct {
	RWAURNCaller     // Read-only binding to the contract
	RWAURNTransactor // Write-only binding to the contract
	RWAURNFilterer   // Log filterer for contract events
}

// RWAURNCaller is an auto generated read-only Go binding around an Ethereum contract.
type RWAURNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWAURNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RWAURNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWAURNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RWAURNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWAURNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RWAURNSession struct {
	Contract     *RWAURN           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWAURNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RWAURNCallerSession struct {
	Contract *RWAURNCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RWAURNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RWAURNTransactorSession struct {
	Contract     *RWAURNTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWAURNRaw is an auto generated low-level Go binding around an Ethereum contract.
type RWAURNRaw struct {
	Contract *RWAURN // Generic contract binding to access the raw methods on
}

// RWAURNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RWAURNCallerRaw struct {
	Contract *RWAURNCaller // Generic read-only contract binding to access the raw methods on
}

// RWAURNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RWAURNTransactorRaw struct {
	Contract *RWAURNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRWAURN creates a new instance of RWAURN, bound to a specific deployed contract.
func NewRWAURN(address common.Address, backend bind.ContractBackend) (*RWAURN, error) {
	contract, err := bindRWAURN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RWAURN{RWAURNCaller: RWAURNCaller{contract: contract}, RWAURNTransactor: RWAURNTransactor{contract: contract}, RWAURNFilterer: RWAURNFilterer{contract: contract}}, nil
}

// NewRWAURNCaller creates a new read-only instance of RWAURN, bound to a specific deployed contract.
func NewRWAURNCaller(address common.Address, caller bind.ContractCaller) (*RWAURNCaller, error) {
	contract, err := bindRWAURN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RWAURNCaller{contract: contract}, nil
}

// NewRWAURNTransactor creates a new write-only instance of RWAURN, bound to a specific deployed contract.
func NewRWAURNTransactor(address common.Address, transactor bind.ContractTransactor) (*RWAURNTransactor, error) {
	contract, err := bindRWAURN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RWAURNTransactor{contract: contract}, nil
}

// NewRWAURNFilterer creates a new log filterer instance of RWAURN, bound to a specific deployed contract.
func NewRWAURNFilterer(address common.Address, filterer bind.ContractFilterer) (*RWAURNFilterer, error) {
	contract, err := bindRWAURN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RWAURNFilterer{contract: contract}, nil
}

// bindRWAURN binds a generic wrapper to an already deployed contract.
func bindRWAURN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RWAURNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWAURN *RWAURNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWAURN.Contract.RWAURNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWAURN *RWAURNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWAURN.Contract.RWAURNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWAURN *RWAURNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWAURN.Contract.RWAURNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWAURN *RWAURNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWAURN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWAURN *RWAURNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWAURN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWAURN *RWAURNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWAURN.Contract.contract.Transact(opts, method, params...)
}

// Can is a free data retrieval call binding the contract method 0xbc206b0a.
//
// Solidity: function can(address ) view returns(uint256)
func (_RWAURN *RWAURNCaller) Can(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "can", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0xbc206b0a.
//
// Solidity: function can(address ) view returns(uint256)
func (_RWAURN *RWAURNSession) Can(arg0 common.Address) (*big.Int, error) {
	return _RWAURN.Contract.Can(&_RWAURN.CallOpts, arg0)
}

// Can is a free data retrieval call binding the contract method 0xbc206b0a.
//
// Solidity: function can(address ) view returns(uint256)
func (_RWAURN *RWAURNCallerSession) Can(arg0 common.Address) (*big.Int, error) {
	return _RWAURN.Contract.Can(&_RWAURN.CallOpts, arg0)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWAURN *RWAURNCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWAURN *RWAURNSession) DaiJoin() (common.Address, error) {
	return _RWAURN.Contract.DaiJoin(&_RWAURN.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWAURN *RWAURNCallerSession) DaiJoin() (common.Address, error) {
	return _RWAURN.Contract.DaiJoin(&_RWAURN.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_RWAURN *RWAURNCaller) GemJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "gemJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_RWAURN *RWAURNSession) GemJoin() (common.Address, error) {
	return _RWAURN.Contract.GemJoin(&_RWAURN.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_RWAURN *RWAURNCallerSession) GemJoin() (common.Address, error) {
	return _RWAURN.Contract.GemJoin(&_RWAURN.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_RWAURN *RWAURNCaller) Jug(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "jug")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_RWAURN *RWAURNSession) Jug() (common.Address, error) {
	return _RWAURN.Contract.Jug(&_RWAURN.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_RWAURN *RWAURNCallerSession) Jug() (common.Address, error) {
	return _RWAURN.Contract.Jug(&_RWAURN.CallOpts)
}

// OutputConduit is a free data retrieval call binding the contract method 0x7692535f.
//
// Solidity: function outputConduit() view returns(address)
func (_RWAURN *RWAURNCaller) OutputConduit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "outputConduit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutputConduit is a free data retrieval call binding the contract method 0x7692535f.
//
// Solidity: function outputConduit() view returns(address)
func (_RWAURN *RWAURNSession) OutputConduit() (common.Address, error) {
	return _RWAURN.Contract.OutputConduit(&_RWAURN.CallOpts)
}

// OutputConduit is a free data retrieval call binding the contract method 0x7692535f.
//
// Solidity: function outputConduit() view returns(address)
func (_RWAURN *RWAURNCallerSession) OutputConduit() (common.Address, error) {
	return _RWAURN.Contract.OutputConduit(&_RWAURN.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWAURN *RWAURNCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWAURN *RWAURNSession) Vat() (common.Address, error) {
	return _RWAURN.Contract.Vat(&_RWAURN.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWAURN *RWAURNCallerSession) Vat() (common.Address, error) {
	return _RWAURN.Contract.Vat(&_RWAURN.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWAURN *RWAURNCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RWAURN.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWAURN *RWAURNSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _RWAURN.Contract.Wards(&_RWAURN.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWAURN *RWAURNCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _RWAURN.Contract.Wards(&_RWAURN.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWAURN *RWAURNTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWAURN *RWAURNSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Deny(&_RWAURN.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWAURN *RWAURNTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Deny(&_RWAURN.TransactOpts, usr)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWAURN *RWAURNTransactor) Draw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "draw", wad)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWAURN *RWAURNSession) Draw(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Draw(&_RWAURN.TransactOpts, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWAURN *RWAURNTransactorSession) Draw(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Draw(&_RWAURN.TransactOpts, wad)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWAURN *RWAURNTransactor) File(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWAURN *RWAURNSession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.File(&_RWAURN.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWAURN *RWAURNTransactorSession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.File(&_RWAURN.TransactOpts, what, data)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWAURN *RWAURNTransactor) Free(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "free", wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWAURN *RWAURNSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Free(&_RWAURN.TransactOpts, wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWAURN *RWAURNTransactorSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Free(&_RWAURN.TransactOpts, wad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_RWAURN *RWAURNTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_RWAURN *RWAURNSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Hope(&_RWAURN.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_RWAURN *RWAURNTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Hope(&_RWAURN.TransactOpts, usr)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWAURN *RWAURNTransactor) Lock(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "lock", wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWAURN *RWAURNSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Lock(&_RWAURN.TransactOpts, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWAURN *RWAURNTransactorSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Lock(&_RWAURN.TransactOpts, wad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_RWAURN *RWAURNTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_RWAURN *RWAURNSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Nope(&_RWAURN.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_RWAURN *RWAURNTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Nope(&_RWAURN.TransactOpts, usr)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWAURN *RWAURNTransactor) Quit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "quit")
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWAURN *RWAURNSession) Quit() (*types.Transaction, error) {
	return _RWAURN.Contract.Quit(&_RWAURN.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWAURN *RWAURNTransactorSession) Quit() (*types.Transaction, error) {
	return _RWAURN.Contract.Quit(&_RWAURN.TransactOpts)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWAURN *RWAURNTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWAURN *RWAURNSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Rely(&_RWAURN.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWAURN *RWAURNTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _RWAURN.Contract.Rely(&_RWAURN.TransactOpts, usr)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWAURN *RWAURNTransactor) Wipe(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.contract.Transact(opts, "wipe", wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWAURN *RWAURNSession) Wipe(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Wipe(&_RWAURN.TransactOpts, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWAURN *RWAURNTransactorSession) Wipe(wad *big.Int) (*types.Transaction, error) {
	return _RWAURN.Contract.Wipe(&_RWAURN.TransactOpts, wad)
}

// RWAURNDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the RWAURN contract.
type RWAURNDenyIterator struct {
	Event *RWAURNDeny // Event containing the contract specifics and raw log

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
func (it *RWAURNDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNDeny)
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
		it.Event = new(RWAURNDeny)
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
func (it *RWAURNDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNDeny represents a Deny event raised by the RWAURN contract.
type RWAURNDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_RWAURN *RWAURNFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*RWAURNDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNDenyIterator{contract: _RWAURN.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_RWAURN *RWAURNFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *RWAURNDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNDeny)
				if err := _RWAURN.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_RWAURN *RWAURNFilterer) ParseDeny(log types.Log) (*RWAURNDeny, error) {
	event := new(RWAURNDeny)
	if err := _RWAURN.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNDrawIterator is returned from FilterDraw and is used to iterate over the raw logs and unpacked data for Draw events raised by the RWAURN contract.
type RWAURNDrawIterator struct {
	Event *RWAURNDraw // Event containing the contract specifics and raw log

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
func (it *RWAURNDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNDraw)
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
		it.Event = new(RWAURNDraw)
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
func (it *RWAURNDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNDraw represents a Draw event raised by the RWAURN contract.
type RWAURNDraw struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDraw is a free log retrieval operation binding the contract event 0x7ffa12f2611233f19bb229a71c5d8224cb37373555ab6754b65aef59ea26831d.
//
// Solidity: event Draw(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) FilterDraw(opts *bind.FilterOpts, usr []common.Address) (*RWAURNDrawIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Draw", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNDrawIterator{contract: _RWAURN.contract, event: "Draw", logs: logs, sub: sub}, nil
}

// WatchDraw is a free log subscription operation binding the contract event 0x7ffa12f2611233f19bb229a71c5d8224cb37373555ab6754b65aef59ea26831d.
//
// Solidity: event Draw(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) WatchDraw(opts *bind.WatchOpts, sink chan<- *RWAURNDraw, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Draw", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNDraw)
				if err := _RWAURN.contract.UnpackLog(event, "Draw", log); err != nil {
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

// ParseDraw is a log parse operation binding the contract event 0x7ffa12f2611233f19bb229a71c5d8224cb37373555ab6754b65aef59ea26831d.
//
// Solidity: event Draw(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) ParseDraw(log types.Log) (*RWAURNDraw, error) {
	event := new(RWAURNDraw)
	if err := _RWAURN.contract.UnpackLog(event, "Draw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the RWAURN contract.
type RWAURNFileIterator struct {
	Event *RWAURNFile // Event containing the contract specifics and raw log

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
func (it *RWAURNFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNFile)
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
		it.Event = new(RWAURNFile)
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
func (it *RWAURNFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNFile represents a File event raised by the RWAURN contract.
type RWAURNFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_RWAURN *RWAURNFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*RWAURNFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNFileIterator{contract: _RWAURN.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_RWAURN *RWAURNFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *RWAURNFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNFile)
				if err := _RWAURN.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_RWAURN *RWAURNFilterer) ParseFile(log types.Log) (*RWAURNFile, error) {
	event := new(RWAURNFile)
	if err := _RWAURN.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNFreeIterator is returned from FilterFree and is used to iterate over the raw logs and unpacked data for Free events raised by the RWAURN contract.
type RWAURNFreeIterator struct {
	Event *RWAURNFree // Event containing the contract specifics and raw log

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
func (it *RWAURNFreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNFree)
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
		it.Event = new(RWAURNFree)
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
func (it *RWAURNFreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNFreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNFree represents a Free event raised by the RWAURN contract.
type RWAURNFree struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFree is a free log retrieval operation binding the contract event 0xce6c5af8fd109993cb40da4d5dc9e4dd8e61bc2e48f1e3901472141e4f56f293.
//
// Solidity: event Free(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) FilterFree(opts *bind.FilterOpts, usr []common.Address) (*RWAURNFreeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Free", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNFreeIterator{contract: _RWAURN.contract, event: "Free", logs: logs, sub: sub}, nil
}

// WatchFree is a free log subscription operation binding the contract event 0xce6c5af8fd109993cb40da4d5dc9e4dd8e61bc2e48f1e3901472141e4f56f293.
//
// Solidity: event Free(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) WatchFree(opts *bind.WatchOpts, sink chan<- *RWAURNFree, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Free", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNFree)
				if err := _RWAURN.contract.UnpackLog(event, "Free", log); err != nil {
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

// ParseFree is a log parse operation binding the contract event 0xce6c5af8fd109993cb40da4d5dc9e4dd8e61bc2e48f1e3901472141e4f56f293.
//
// Solidity: event Free(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) ParseFree(log types.Log) (*RWAURNFree, error) {
	event := new(RWAURNFree)
	if err := _RWAURN.contract.UnpackLog(event, "Free", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNHopeIterator is returned from FilterHope and is used to iterate over the raw logs and unpacked data for Hope events raised by the RWAURN contract.
type RWAURNHopeIterator struct {
	Event *RWAURNHope // Event containing the contract specifics and raw log

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
func (it *RWAURNHopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNHope)
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
		it.Event = new(RWAURNHope)
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
func (it *RWAURNHopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNHopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNHope represents a Hope event raised by the RWAURN contract.
type RWAURNHope struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHope is a free log retrieval operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_RWAURN *RWAURNFilterer) FilterHope(opts *bind.FilterOpts, usr []common.Address) (*RWAURNHopeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Hope", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNHopeIterator{contract: _RWAURN.contract, event: "Hope", logs: logs, sub: sub}, nil
}

// WatchHope is a free log subscription operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_RWAURN *RWAURNFilterer) WatchHope(opts *bind.WatchOpts, sink chan<- *RWAURNHope, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Hope", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNHope)
				if err := _RWAURN.contract.UnpackLog(event, "Hope", log); err != nil {
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

// ParseHope is a log parse operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_RWAURN *RWAURNFilterer) ParseHope(log types.Log) (*RWAURNHope, error) {
	event := new(RWAURNHope)
	if err := _RWAURN.contract.UnpackLog(event, "Hope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the RWAURN contract.
type RWAURNLockIterator struct {
	Event *RWAURNLock // Event containing the contract specifics and raw log

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
func (it *RWAURNLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNLock)
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
		it.Event = new(RWAURNLock)
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
func (it *RWAURNLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNLock represents a Lock event raised by the RWAURN contract.
type RWAURNLock struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) FilterLock(opts *bind.FilterOpts, usr []common.Address) (*RWAURNLockIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Lock", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNLockIterator{contract: _RWAURN.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *RWAURNLock, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Lock", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNLock)
				if err := _RWAURN.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) ParseLock(log types.Log) (*RWAURNLock, error) {
	event := new(RWAURNLock)
	if err := _RWAURN.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNNopeIterator is returned from FilterNope and is used to iterate over the raw logs and unpacked data for Nope events raised by the RWAURN contract.
type RWAURNNopeIterator struct {
	Event *RWAURNNope // Event containing the contract specifics and raw log

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
func (it *RWAURNNopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNNope)
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
		it.Event = new(RWAURNNope)
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
func (it *RWAURNNopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNNopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNNope represents a Nope event raised by the RWAURN contract.
type RWAURNNope struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNope is a free log retrieval operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_RWAURN *RWAURNFilterer) FilterNope(opts *bind.FilterOpts, usr []common.Address) (*RWAURNNopeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Nope", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNNopeIterator{contract: _RWAURN.contract, event: "Nope", logs: logs, sub: sub}, nil
}

// WatchNope is a free log subscription operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_RWAURN *RWAURNFilterer) WatchNope(opts *bind.WatchOpts, sink chan<- *RWAURNNope, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Nope", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNNope)
				if err := _RWAURN.contract.UnpackLog(event, "Nope", log); err != nil {
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

// ParseNope is a log parse operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_RWAURN *RWAURNFilterer) ParseNope(log types.Log) (*RWAURNNope, error) {
	event := new(RWAURNNope)
	if err := _RWAURN.contract.UnpackLog(event, "Nope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNQuitIterator is returned from FilterQuit and is used to iterate over the raw logs and unpacked data for Quit events raised by the RWAURN contract.
type RWAURNQuitIterator struct {
	Event *RWAURNQuit // Event containing the contract specifics and raw log

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
func (it *RWAURNQuitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNQuit)
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
		it.Event = new(RWAURNQuit)
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
func (it *RWAURNQuitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNQuitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNQuit represents a Quit event raised by the RWAURN contract.
type RWAURNQuit struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuit is a free log retrieval operation binding the contract event 0xc81bfec1ac9d038698c3b15fc900dafbff3af4b9f26062f895dd08a676ec78ae.
//
// Solidity: event Quit(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) FilterQuit(opts *bind.FilterOpts, usr []common.Address) (*RWAURNQuitIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Quit", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNQuitIterator{contract: _RWAURN.contract, event: "Quit", logs: logs, sub: sub}, nil
}

// WatchQuit is a free log subscription operation binding the contract event 0xc81bfec1ac9d038698c3b15fc900dafbff3af4b9f26062f895dd08a676ec78ae.
//
// Solidity: event Quit(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) WatchQuit(opts *bind.WatchOpts, sink chan<- *RWAURNQuit, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Quit", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNQuit)
				if err := _RWAURN.contract.UnpackLog(event, "Quit", log); err != nil {
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

// ParseQuit is a log parse operation binding the contract event 0xc81bfec1ac9d038698c3b15fc900dafbff3af4b9f26062f895dd08a676ec78ae.
//
// Solidity: event Quit(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) ParseQuit(log types.Log) (*RWAURNQuit, error) {
	event := new(RWAURNQuit)
	if err := _RWAURN.contract.UnpackLog(event, "Quit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the RWAURN contract.
type RWAURNRelyIterator struct {
	Event *RWAURNRely // Event containing the contract specifics and raw log

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
func (it *RWAURNRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNRely)
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
		it.Event = new(RWAURNRely)
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
func (it *RWAURNRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNRely represents a Rely event raised by the RWAURN contract.
type RWAURNRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_RWAURN *RWAURNFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*RWAURNRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNRelyIterator{contract: _RWAURN.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_RWAURN *RWAURNFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *RWAURNRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNRely)
				if err := _RWAURN.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_RWAURN *RWAURNFilterer) ParseRely(log types.Log) (*RWAURNRely, error) {
	event := new(RWAURNRely)
	if err := _RWAURN.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWAURNWipeIterator is returned from FilterWipe and is used to iterate over the raw logs and unpacked data for Wipe events raised by the RWAURN contract.
type RWAURNWipeIterator struct {
	Event *RWAURNWipe // Event containing the contract specifics and raw log

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
func (it *RWAURNWipeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWAURNWipe)
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
		it.Event = new(RWAURNWipe)
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
func (it *RWAURNWipeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWAURNWipeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWAURNWipe represents a Wipe event raised by the RWAURN contract.
type RWAURNWipe struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWipe is a free log retrieval operation binding the contract event 0x2d2c7da251295f4d722a8ddaf337627952c957ce21b2757c852e47fe81b3a2af.
//
// Solidity: event Wipe(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) FilterWipe(opts *bind.FilterOpts, usr []common.Address) (*RWAURNWipeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.FilterLogs(opts, "Wipe", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWAURNWipeIterator{contract: _RWAURN.contract, event: "Wipe", logs: logs, sub: sub}, nil
}

// WatchWipe is a free log subscription operation binding the contract event 0x2d2c7da251295f4d722a8ddaf337627952c957ce21b2757c852e47fe81b3a2af.
//
// Solidity: event Wipe(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) WatchWipe(opts *bind.WatchOpts, sink chan<- *RWAURNWipe, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWAURN.contract.WatchLogs(opts, "Wipe", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWAURNWipe)
				if err := _RWAURN.contract.UnpackLog(event, "Wipe", log); err != nil {
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

// ParseWipe is a log parse operation binding the contract event 0x2d2c7da251295f4d722a8ddaf337627952c957ce21b2757c852e47fe81b3a2af.
//
// Solidity: event Wipe(address indexed usr, uint256 wad)
func (_RWAURN *RWAURNFilterer) ParseWipe(log types.Log) (*RWAURNWipe, error) {
	event := new(RWAURNWipe)
	if err := _RWAURN.contract.UnpackLog(event, "Wipe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
