// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDCat

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

// MCDCATABI is the input ABI used to generate the binding from.
const MCDCATABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Bite\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"bite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"box\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"claw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chop\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dunk\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"litter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVowLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDCAT is an auto generated Go binding around an Ethereum contract.
type MCDCAT struct {
	MCDCATCaller     // Read-only binding to the contract
	MCDCATTransactor // Write-only binding to the contract
	MCDCATFilterer   // Log filterer for contract events
}

// MCDCATCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDCATCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCATTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDCATTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCATFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDCATFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCATSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDCATSession struct {
	Contract     *MCDCAT           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDCATCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDCATCallerSession struct {
	Contract *MCDCATCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDCATTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDCATTransactorSession struct {
	Contract     *MCDCATTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDCATRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDCATRaw struct {
	Contract *MCDCAT // Generic contract binding to access the raw methods on
}

// MCDCATCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDCATCallerRaw struct {
	Contract *MCDCATCaller // Generic read-only contract binding to access the raw methods on
}

// MCDCATTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDCATTransactorRaw struct {
	Contract *MCDCATTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDCAT creates a new instance of MCDCAT, bound to a specific deployed contract.
func NewMCDCAT(address common.Address, backend bind.ContractBackend) (*MCDCAT, error) {
	contract, err := bindMCDCAT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDCAT{MCDCATCaller: MCDCATCaller{contract: contract}, MCDCATTransactor: MCDCATTransactor{contract: contract}, MCDCATFilterer: MCDCATFilterer{contract: contract}}, nil
}

// NewMCDCATCaller creates a new read-only instance of MCDCAT, bound to a specific deployed contract.
func NewMCDCATCaller(address common.Address, caller bind.ContractCaller) (*MCDCATCaller, error) {
	contract, err := bindMCDCAT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCATCaller{contract: contract}, nil
}

// NewMCDCATTransactor creates a new write-only instance of MCDCAT, bound to a specific deployed contract.
func NewMCDCATTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDCATTransactor, error) {
	contract, err := bindMCDCAT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCATTransactor{contract: contract}, nil
}

// NewMCDCATFilterer creates a new log filterer instance of MCDCAT, bound to a specific deployed contract.
func NewMCDCATFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDCATFilterer, error) {
	contract, err := bindMCDCAT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDCATFilterer{contract: contract}, nil
}

// bindMCDCAT binds a generic wrapper to an already deployed contract.
func bindMCDCAT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDCATABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCAT *MCDCATRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCAT.Contract.MCDCATCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCAT *MCDCATRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCAT.Contract.MCDCATTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCAT *MCDCATRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCAT.Contract.MCDCATTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCAT *MCDCATCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCAT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCAT *MCDCATTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCAT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCAT *MCDCATTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCAT.Contract.contract.Transact(opts, method, params...)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(uint256)
func (_MCDCAT *MCDCATCaller) Box(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "box")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(uint256)
func (_MCDCAT *MCDCATSession) Box() (*big.Int, error) {
	return _MCDCAT.Contract.Box(&_MCDCAT.CallOpts)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(uint256)
func (_MCDCAT *MCDCATCallerSession) Box() (*big.Int, error) {
	return _MCDCAT.Contract.Box(&_MCDCAT.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, uint256 chop, uint256 dunk)
func (_MCDCAT *MCDCATCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Dunk *big.Int
}, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Flip common.Address
		Chop *big.Int
		Dunk *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Flip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Chop = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Dunk = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, uint256 chop, uint256 dunk)
func (_MCDCAT *MCDCATSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Dunk *big.Int
}, error) {
	return _MCDCAT.Contract.Ilks(&_MCDCAT.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, uint256 chop, uint256 dunk)
func (_MCDCAT *MCDCATCallerSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Chop *big.Int
	Dunk *big.Int
}, error) {
	return _MCDCAT.Contract.Ilks(&_MCDCAT.CallOpts, arg0)
}

// Litter is a free data retrieval call binding the contract method 0xa4fe8caf.
//
// Solidity: function litter() view returns(uint256)
func (_MCDCAT *MCDCATCaller) Litter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "litter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Litter is a free data retrieval call binding the contract method 0xa4fe8caf.
//
// Solidity: function litter() view returns(uint256)
func (_MCDCAT *MCDCATSession) Litter() (*big.Int, error) {
	return _MCDCAT.Contract.Litter(&_MCDCAT.CallOpts)
}

// Litter is a free data retrieval call binding the contract method 0xa4fe8caf.
//
// Solidity: function litter() view returns(uint256)
func (_MCDCAT *MCDCATCallerSession) Litter() (*big.Int, error) {
	return _MCDCAT.Contract.Litter(&_MCDCAT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDCAT *MCDCATCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDCAT *MCDCATSession) Live() (*big.Int, error) {
	return _MCDCAT.Contract.Live(&_MCDCAT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDCAT *MCDCATCallerSession) Live() (*big.Int, error) {
	return _MCDCAT.Contract.Live(&_MCDCAT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCAT *MCDCATCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCAT *MCDCATSession) Vat() (common.Address, error) {
	return _MCDCAT.Contract.Vat(&_MCDCAT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCAT *MCDCATCallerSession) Vat() (common.Address, error) {
	return _MCDCAT.Contract.Vat(&_MCDCAT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDCAT *MCDCATCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDCAT *MCDCATSession) Vow() (common.Address, error) {
	return _MCDCAT.Contract.Vow(&_MCDCAT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDCAT *MCDCATCallerSession) Vow() (common.Address, error) {
	return _MCDCAT.Contract.Vow(&_MCDCAT.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCAT *MCDCATCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDCAT.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCAT *MCDCATSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDCAT.Contract.Wards(&_MCDCAT.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDCAT *MCDCATCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDCAT.Contract.Wards(&_MCDCAT.CallOpts, arg0)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_MCDCAT *MCDCATTransactor) Bite(opts *bind.TransactOpts, ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "bite", ilk, urn)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_MCDCAT *MCDCATSession) Bite(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Bite(&_MCDCAT.TransactOpts, ilk, urn)
}

// Bite is a paid mutator transaction binding the contract method 0x45cf2230.
//
// Solidity: function bite(bytes32 ilk, address urn) returns(uint256 id)
func (_MCDCAT *MCDCATTransactorSession) Bite(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Bite(&_MCDCAT.TransactOpts, ilk, urn)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDCAT *MCDCATTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDCAT *MCDCATSession) Cage() (*types.Transaction, error) {
	return _MCDCAT.Contract.Cage(&_MCDCAT.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDCAT *MCDCATTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDCAT.Contract.Cage(&_MCDCAT.TransactOpts)
}

// Claw is a paid mutator transaction binding the contract method 0xe66d279b.
//
// Solidity: function claw(uint256 rad) returns()
func (_MCDCAT *MCDCATTransactor) Claw(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "claw", rad)
}

// Claw is a paid mutator transaction binding the contract method 0xe66d279b.
//
// Solidity: function claw(uint256 rad) returns()
func (_MCDCAT *MCDCATSession) Claw(rad *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.Claw(&_MCDCAT.TransactOpts, rad)
}

// Claw is a paid mutator transaction binding the contract method 0xe66d279b.
//
// Solidity: function claw(uint256 rad) returns()
func (_MCDCAT *MCDCATTransactorSession) Claw(rad *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.Claw(&_MCDCAT.TransactOpts, rad)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCAT *MCDCATTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCAT *MCDCATSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Deny(&_MCDCAT.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDCAT *MCDCATTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Deny(&_MCDCAT.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.File(&_MCDCAT.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.File(&_MCDCAT.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.File0(&_MCDCAT.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDCAT *MCDCATTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDCAT.Contract.File0(&_MCDCAT.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDCAT *MCDCATTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDCAT *MCDCATSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.File1(&_MCDCAT.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDCAT *MCDCATTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.File1(&_MCDCAT.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_MCDCAT *MCDCATTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "file2", ilk, what, flip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_MCDCAT *MCDCATSession) File2(ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.File2(&_MCDCAT.TransactOpts, ilk, what, flip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address flip) returns()
func (_MCDCAT *MCDCATTransactorSession) File2(ilk [32]byte, what [32]byte, flip common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.File2(&_MCDCAT.TransactOpts, ilk, what, flip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCAT *MCDCATTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCAT *MCDCATSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Rely(&_MCDCAT.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDCAT *MCDCATTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDCAT.Contract.Rely(&_MCDCAT.TransactOpts, usr)
}

// MCDCATBiteIterator is returned from FilterBite and is used to iterate over the raw logs and unpacked data for Bite events raised by the MCDCAT contract.
type MCDCATBiteIterator struct {
	Event *MCDCATBite // Event containing the contract specifics and raw log

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
func (it *MCDCATBiteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCATBite)
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
		it.Event = new(MCDCATBite)
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
func (it *MCDCATBiteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCATBiteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCATBite represents a Bite event raised by the MCDCAT contract.
type MCDCATBite struct {
	Ilk  [32]byte
	Urn  common.Address
	Ink  *big.Int
	Art  *big.Int
	Tab  *big.Int
	Flip common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBite is a free log retrieval operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_MCDCAT *MCDCATFilterer) FilterBite(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address) (*MCDCATBiteIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MCDCAT.contract.FilterLogs(opts, "Bite", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return &MCDCATBiteIterator{contract: _MCDCAT.contract, event: "Bite", logs: logs, sub: sub}, nil
}

// WatchBite is a free log subscription operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_MCDCAT *MCDCATFilterer) WatchBite(opts *bind.WatchOpts, sink chan<- *MCDCATBite, ilk [][32]byte, urn []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MCDCAT.contract.WatchLogs(opts, "Bite", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCATBite)
				if err := _MCDCAT.contract.UnpackLog(event, "Bite", log); err != nil {
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

// ParseBite is a log parse operation binding the contract event 0xa716da86bc1fb6d43d1493373f34d7a418b619681cd7b90f7ea667ba1489be28.
//
// Solidity: event Bite(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 tab, address flip, uint256 id)
func (_MCDCAT *MCDCATFilterer) ParseBite(log types.Log) (*MCDCATBite, error) {
	event := new(MCDCATBite)
	if err := _MCDCAT.contract.UnpackLog(event, "Bite", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
