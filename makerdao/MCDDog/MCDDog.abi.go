// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDDog

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

// MCDDOGABI is the input ABI used to generate the binding from.
const MCDDOGABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"due\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Bark\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Digs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Dirt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"}],\"name\":\"bark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"chop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"digs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chop\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dirt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVowLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDDOG is an auto generated Go binding around an Ethereum contract.
type MCDDOG struct {
	MCDDOGCaller     // Read-only binding to the contract
	MCDDOGTransactor // Write-only binding to the contract
	MCDDOGFilterer   // Log filterer for contract events
}

// MCDDOGCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDDOGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDOGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDDOGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDOGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDDOGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDOGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDDOGSession struct {
	Contract     *MCDDOG           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDDOGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDDOGCallerSession struct {
	Contract *MCDDOGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDDOGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDDOGTransactorSession struct {
	Contract     *MCDDOGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDDOGRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDDOGRaw struct {
	Contract *MCDDOG // Generic contract binding to access the raw methods on
}

// MCDDOGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDDOGCallerRaw struct {
	Contract *MCDDOGCaller // Generic read-only contract binding to access the raw methods on
}

// MCDDOGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDDOGTransactorRaw struct {
	Contract *MCDDOGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDDOG creates a new instance of MCDDOG, bound to a specific deployed contract.
func NewMCDDOG(address common.Address, backend bind.ContractBackend) (*MCDDOG, error) {
	contract, err := bindMCDDOG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDDOG{MCDDOGCaller: MCDDOGCaller{contract: contract}, MCDDOGTransactor: MCDDOGTransactor{contract: contract}, MCDDOGFilterer: MCDDOGFilterer{contract: contract}}, nil
}

// NewMCDDOGCaller creates a new read-only instance of MCDDOG, bound to a specific deployed contract.
func NewMCDDOGCaller(address common.Address, caller bind.ContractCaller) (*MCDDOGCaller, error) {
	contract, err := bindMCDDOG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDDOGCaller{contract: contract}, nil
}

// NewMCDDOGTransactor creates a new write-only instance of MCDDOG, bound to a specific deployed contract.
func NewMCDDOGTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDDOGTransactor, error) {
	contract, err := bindMCDDOG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDDOGTransactor{contract: contract}, nil
}

// NewMCDDOGFilterer creates a new log filterer instance of MCDDOG, bound to a specific deployed contract.
func NewMCDDOGFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDDOGFilterer, error) {
	contract, err := bindMCDDOG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDDOGFilterer{contract: contract}, nil
}

// bindMCDDOG binds a generic wrapper to an already deployed contract.
func bindMCDDOG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDDOGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDDOG *MCDDOGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDDOG.Contract.MCDDOGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDDOG *MCDDOGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDOG.Contract.MCDDOGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDDOG *MCDDOGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDDOG.Contract.MCDDOGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDDOG *MCDDOGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDDOG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDDOG *MCDDOGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDOG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDDOG *MCDDOGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDDOG.Contract.contract.Transact(opts, method, params...)
}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_MCDDOG *MCDDOGCaller) Dirt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "Dirt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_MCDDOG *MCDDOGSession) Dirt() (*big.Int, error) {
	return _MCDDOG.Contract.Dirt(&_MCDDOG.CallOpts)
}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_MCDDOG *MCDDOGCallerSession) Dirt() (*big.Int, error) {
	return _MCDDOG.Contract.Dirt(&_MCDDOG.CallOpts)
}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_MCDDOG *MCDDOGCaller) Hole(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "Hole")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_MCDDOG *MCDDOGSession) Hole() (*big.Int, error) {
	return _MCDDOG.Contract.Hole(&_MCDDOG.CallOpts)
}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_MCDDOG *MCDDOGCallerSession) Hole() (*big.Int, error) {
	return _MCDDOG.Contract.Hole(&_MCDDOG.CallOpts)
}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_MCDDOG *MCDDOGCaller) Chop(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "chop", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_MCDDOG *MCDDOGSession) Chop(ilk [32]byte) (*big.Int, error) {
	return _MCDDOG.Contract.Chop(&_MCDDOG.CallOpts, ilk)
}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_MCDDOG *MCDDOGCallerSession) Chop(ilk [32]byte) (*big.Int, error) {
	return _MCDDOG.Contract.Chop(&_MCDDOG.CallOpts, ilk)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_MCDDOG *MCDDOGCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Clip common.Address
		Chop *big.Int
		Hole *big.Int
		Dirt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Clip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Chop = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Hole = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dirt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_MCDDOG *MCDDOGSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	return _MCDDOG.Contract.Ilks(&_MCDDOG.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_MCDDOG *MCDDOGCallerSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	return _MCDDOG.Contract.Ilks(&_MCDDOG.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDDOG *MCDDOGCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDDOG *MCDDOGSession) Live() (*big.Int, error) {
	return _MCDDOG.Contract.Live(&_MCDDOG.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDDOG *MCDDOGCallerSession) Live() (*big.Int, error) {
	return _MCDDOG.Contract.Live(&_MCDDOG.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDOG *MCDDOGCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDOG *MCDDOGSession) Vat() (common.Address, error) {
	return _MCDDOG.Contract.Vat(&_MCDDOG.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDOG *MCDDOGCallerSession) Vat() (common.Address, error) {
	return _MCDDOG.Contract.Vat(&_MCDDOG.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDOG *MCDDOGCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDOG *MCDDOGSession) Vow() (common.Address, error) {
	return _MCDDOG.Contract.Vow(&_MCDDOG.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDOG *MCDDOGCallerSession) Vow() (common.Address, error) {
	return _MCDDOG.Contract.Vow(&_MCDDOG.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDDOG *MCDDOGCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDDOG.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDDOG *MCDDOGSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDDOG.Contract.Wards(&_MCDDOG.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDDOG *MCDDOGCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDDOG.Contract.Wards(&_MCDDOG.CallOpts, arg0)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_MCDDOG *MCDDOGTransactor) Bark(opts *bind.TransactOpts, ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "bark", ilk, urn, kpr)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_MCDDOG *MCDDOGSession) Bark(ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Bark(&_MCDDOG.TransactOpts, ilk, urn, kpr)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_MCDDOG *MCDDOGTransactorSession) Bark(ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Bark(&_MCDDOG.TransactOpts, ilk, urn, kpr)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDDOG *MCDDOGTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDDOG *MCDDOGSession) Cage() (*types.Transaction, error) {
	return _MCDDOG.Contract.Cage(&_MCDDOG.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDDOG *MCDDOGTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDDOG.Contract.Cage(&_MCDDOG.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDDOG *MCDDOGTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDDOG *MCDDOGSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Deny(&_MCDDOG.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDDOG *MCDDOGTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Deny(&_MCDDOG.TransactOpts, usr)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_MCDDOG *MCDDOGTransactor) Digs(opts *bind.TransactOpts, ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "digs", ilk, rad)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_MCDDOG *MCDDOGSession) Digs(ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.Digs(&_MCDDOG.TransactOpts, ilk, rad)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_MCDDOG *MCDDOGTransactorSession) Digs(ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.Digs(&_MCDDOG.TransactOpts, ilk, rad)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.File(&_MCDDOG.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.File(&_MCDDOG.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.File0(&_MCDDOG.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDDOG *MCDDOGTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDDOG.Contract.File0(&_MCDDOG.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDDOG *MCDDOGTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDDOG *MCDDOGSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.File1(&_MCDDOG.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDDOG *MCDDOGTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.File1(&_MCDDOG.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_MCDDOG *MCDDOGTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "file2", ilk, what, clip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_MCDDOG *MCDDOGSession) File2(ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.File2(&_MCDDOG.TransactOpts, ilk, what, clip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_MCDDOG *MCDDOGTransactorSession) File2(ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.File2(&_MCDDOG.TransactOpts, ilk, what, clip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDDOG *MCDDOGTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDDOG *MCDDOGSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Rely(&_MCDDOG.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDDOG *MCDDOGTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDDOG.Contract.Rely(&_MCDDOG.TransactOpts, usr)
}

// MCDDOGBarkIterator is returned from FilterBark and is used to iterate over the raw logs and unpacked data for Bark events raised by the MCDDOG contract.
type MCDDOGBarkIterator struct {
	Event *MCDDOGBark // Event containing the contract specifics and raw log

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
func (it *MCDDOGBarkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGBark)
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
		it.Event = new(MCDDOGBark)
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
func (it *MCDDOGBarkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGBarkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGBark represents a Bark event raised by the MCDDOG contract.
type MCDDOGBark struct {
	Ilk  [32]byte
	Urn  common.Address
	Ink  *big.Int
	Art  *big.Int
	Due  *big.Int
	Clip common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBark is a free log retrieval operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_MCDDOG *MCDDOGFilterer) FilterBark(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address, id []*big.Int) (*MCDDOGBarkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "Bark", ilkRule, urnRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGBarkIterator{contract: _MCDDOG.contract, event: "Bark", logs: logs, sub: sub}, nil
}

// WatchBark is a free log subscription operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_MCDDOG *MCDDOGFilterer) WatchBark(opts *bind.WatchOpts, sink chan<- *MCDDOGBark, ilk [][32]byte, urn []common.Address, id []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "Bark", ilkRule, urnRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGBark)
				if err := _MCDDOG.contract.UnpackLog(event, "Bark", log); err != nil {
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

// ParseBark is a log parse operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_MCDDOG *MCDDOGFilterer) ParseBark(log types.Log) (*MCDDOGBark, error) {
	event := new(MCDDOGBark)
	if err := _MCDDOG.contract.UnpackLog(event, "Bark", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the MCDDOG contract.
type MCDDOGCageIterator struct {
	Event *MCDDOGCage // Event containing the contract specifics and raw log

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
func (it *MCDDOGCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGCage)
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
		it.Event = new(MCDDOGCage)
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
func (it *MCDDOGCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGCage represents a Cage event raised by the MCDDOG contract.
type MCDDOGCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_MCDDOG *MCDDOGFilterer) FilterCage(opts *bind.FilterOpts) (*MCDDOGCageIterator, error) {

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &MCDDOGCageIterator{contract: _MCDDOG.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_MCDDOG *MCDDOGFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *MCDDOGCage) (event.Subscription, error) {

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGCage)
				if err := _MCDDOG.contract.UnpackLog(event, "Cage", log); err != nil {
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

// ParseCage is a log parse operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_MCDDOG *MCDDOGFilterer) ParseCage(log types.Log) (*MCDDOGCage, error) {
	event := new(MCDDOGCage)
	if err := _MCDDOG.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDDOG contract.
type MCDDOGDenyIterator struct {
	Event *MCDDOGDeny // Event containing the contract specifics and raw log

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
func (it *MCDDOGDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGDeny)
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
		it.Event = new(MCDDOGDeny)
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
func (it *MCDDOGDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGDeny represents a Deny event raised by the MCDDOG contract.
type MCDDOGDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDDOG *MCDDOGFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDDOGDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGDenyIterator{contract: _MCDDOG.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDDOG *MCDDOGFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDDOGDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGDeny)
				if err := _MCDDOG.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDDOG *MCDDOGFilterer) ParseDeny(log types.Log) (*MCDDOGDeny, error) {
	event := new(MCDDOGDeny)
	if err := _MCDDOG.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGDigsIterator is returned from FilterDigs and is used to iterate over the raw logs and unpacked data for Digs events raised by the MCDDOG contract.
type MCDDOGDigsIterator struct {
	Event *MCDDOGDigs // Event containing the contract specifics and raw log

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
func (it *MCDDOGDigsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGDigs)
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
		it.Event = new(MCDDOGDigs)
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
func (it *MCDDOGDigsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGDigsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGDigs represents a Digs event raised by the MCDDOG contract.
type MCDDOGDigs struct {
	Ilk [32]byte
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDigs is a free log retrieval operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_MCDDOG *MCDDOGFilterer) FilterDigs(opts *bind.FilterOpts, ilk [][32]byte) (*MCDDOGDigsIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "Digs", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGDigsIterator{contract: _MCDDOG.contract, event: "Digs", logs: logs, sub: sub}, nil
}

// WatchDigs is a free log subscription operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_MCDDOG *MCDDOGFilterer) WatchDigs(opts *bind.WatchOpts, sink chan<- *MCDDOGDigs, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "Digs", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGDigs)
				if err := _MCDDOG.contract.UnpackLog(event, "Digs", log); err != nil {
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

// ParseDigs is a log parse operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_MCDDOG *MCDDOGFilterer) ParseDigs(log types.Log) (*MCDDOGDigs, error) {
	event := new(MCDDOGDigs)
	if err := _MCDDOG.contract.UnpackLog(event, "Digs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDDOG contract.
type MCDDOGFileIterator struct {
	Event *MCDDOGFile // Event containing the contract specifics and raw log

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
func (it *MCDDOGFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGFile)
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
		it.Event = new(MCDDOGFile)
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
func (it *MCDDOGFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGFile represents a File event raised by the MCDDOG contract.
type MCDDOGFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDDOG *MCDDOGFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDDOGFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGFileIterator{contract: _MCDDOG.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDDOG *MCDDOGFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDDOGFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGFile)
				if err := _MCDDOG.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDDOG *MCDDOGFilterer) ParseFile(log types.Log) (*MCDDOGFile, error) {
	event := new(MCDDOGFile)
	if err := _MCDDOG.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the MCDDOG contract.
type MCDDOGFile0Iterator struct {
	Event *MCDDOGFile0 // Event containing the contract specifics and raw log

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
func (it *MCDDOGFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGFile0)
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
		it.Event = new(MCDDOGFile0)
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
func (it *MCDDOGFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGFile0 represents a File0 event raised by the MCDDOG contract.
type MCDDOGFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDDOG *MCDDOGFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*MCDDOGFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGFile0Iterator{contract: _MCDDOG.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDDOG *MCDDOGFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *MCDDOGFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGFile0)
				if err := _MCDDOG.contract.UnpackLog(event, "File0", log); err != nil {
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
func (_MCDDOG *MCDDOGFilterer) ParseFile0(log types.Log) (*MCDDOGFile0, error) {
	event := new(MCDDOGFile0)
	if err := _MCDDOG.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the MCDDOG contract.
type MCDDOGFile1Iterator struct {
	Event *MCDDOGFile1 // Event containing the contract specifics and raw log

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
func (it *MCDDOGFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGFile1)
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
		it.Event = new(MCDDOGFile1)
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
func (it *MCDDOGFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGFile1 represents a File1 event raised by the MCDDOG contract.
type MCDDOGFile1 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_MCDDOG *MCDDOGFilterer) FilterFile1(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*MCDDOGFile1Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGFile1Iterator{contract: _MCDDOG.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_MCDDOG *MCDDOGFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *MCDDOGFile1, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGFile1)
				if err := _MCDDOG.contract.UnpackLog(event, "File1", log); err != nil {
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

// ParseFile1 is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_MCDDOG *MCDDOGFilterer) ParseFile1(log types.Log) (*MCDDOGFile1, error) {
	event := new(MCDDOGFile1)
	if err := _MCDDOG.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGFile2Iterator is returned from FilterFile2 and is used to iterate over the raw logs and unpacked data for File2 events raised by the MCDDOG contract.
type MCDDOGFile2Iterator struct {
	Event *MCDDOGFile2 // Event containing the contract specifics and raw log

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
func (it *MCDDOGFile2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGFile2)
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
		it.Event = new(MCDDOGFile2)
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
func (it *MCDDOGFile2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGFile2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGFile2 represents a File2 event raised by the MCDDOG contract.
type MCDDOGFile2 struct {
	Ilk  [32]byte
	What [32]byte
	Clip common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile2 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_MCDDOG *MCDDOGFilterer) FilterFile2(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*MCDDOGFile2Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGFile2Iterator{contract: _MCDDOG.contract, event: "File2", logs: logs, sub: sub}, nil
}

// WatchFile2 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_MCDDOG *MCDDOGFilterer) WatchFile2(opts *bind.WatchOpts, sink chan<- *MCDDOGFile2, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGFile2)
				if err := _MCDDOG.contract.UnpackLog(event, "File2", log); err != nil {
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

// ParseFile2 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_MCDDOG *MCDDOGFilterer) ParseFile2(log types.Log) (*MCDDOGFile2, error) {
	event := new(MCDDOGFile2)
	if err := _MCDDOG.contract.UnpackLog(event, "File2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDOGRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDDOG contract.
type MCDDOGRelyIterator struct {
	Event *MCDDOGRely // Event containing the contract specifics and raw log

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
func (it *MCDDOGRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDOGRely)
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
		it.Event = new(MCDDOGRely)
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
func (it *MCDDOGRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDOGRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDOGRely represents a Rely event raised by the MCDDOG contract.
type MCDDOGRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDDOG *MCDDOGFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDDOGRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDDOG.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDDOGRelyIterator{contract: _MCDDOG.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDDOG *MCDDOGFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDDOGRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDDOG.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDOGRely)
				if err := _MCDDOG.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDDOG *MCDDOGFilterer) ParseRely(log types.Log) (*MCDDOGRely, error) {
	event := new(MCDDOGRely)
	if err := _MCDDOG.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
