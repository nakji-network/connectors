// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ILK_REGISTRY

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

// ILKREGISTRYABI is the input ABI used to generate the binding from.
const ILKREGISTRYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dog_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spot_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"AddIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"NameError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"RemoveIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"SymbolError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"UpdateIlk\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cat\",\"outputs\":[{\"internalType\":\"contractCatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilkData\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"pos\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"dec\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"class\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dec\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spot\",\"outputs\":[{\"internalType\":\"contractSpotLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"xlip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILKREGISTRY is an auto generated Go binding around an Ethereum contract.
type ILKREGISTRY struct {
	ILKREGISTRYCaller     // Read-only binding to the contract
	ILKREGISTRYTransactor // Write-only binding to the contract
	ILKREGISTRYFilterer   // Log filterer for contract events
}

// ILKREGISTRYCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILKREGISTRYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILKREGISTRYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILKREGISTRYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILKREGISTRYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILKREGISTRYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILKREGISTRYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILKREGISTRYSession struct {
	Contract     *ILKREGISTRY      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILKREGISTRYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILKREGISTRYCallerSession struct {
	Contract *ILKREGISTRYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ILKREGISTRYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILKREGISTRYTransactorSession struct {
	Contract     *ILKREGISTRYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ILKREGISTRYRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILKREGISTRYRaw struct {
	Contract *ILKREGISTRY // Generic contract binding to access the raw methods on
}

// ILKREGISTRYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILKREGISTRYCallerRaw struct {
	Contract *ILKREGISTRYCaller // Generic read-only contract binding to access the raw methods on
}

// ILKREGISTRYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILKREGISTRYTransactorRaw struct {
	Contract *ILKREGISTRYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILKREGISTRY creates a new instance of ILKREGISTRY, bound to a specific deployed contract.
func NewILKREGISTRY(address common.Address, backend bind.ContractBackend) (*ILKREGISTRY, error) {
	contract, err := bindILKREGISTRY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRY{ILKREGISTRYCaller: ILKREGISTRYCaller{contract: contract}, ILKREGISTRYTransactor: ILKREGISTRYTransactor{contract: contract}, ILKREGISTRYFilterer: ILKREGISTRYFilterer{contract: contract}}, nil
}

// NewILKREGISTRYCaller creates a new read-only instance of ILKREGISTRY, bound to a specific deployed contract.
func NewILKREGISTRYCaller(address common.Address, caller bind.ContractCaller) (*ILKREGISTRYCaller, error) {
	contract, err := bindILKREGISTRY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYCaller{contract: contract}, nil
}

// NewILKREGISTRYTransactor creates a new write-only instance of ILKREGISTRY, bound to a specific deployed contract.
func NewILKREGISTRYTransactor(address common.Address, transactor bind.ContractTransactor) (*ILKREGISTRYTransactor, error) {
	contract, err := bindILKREGISTRY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYTransactor{contract: contract}, nil
}

// NewILKREGISTRYFilterer creates a new log filterer instance of ILKREGISTRY, bound to a specific deployed contract.
func NewILKREGISTRYFilterer(address common.Address, filterer bind.ContractFilterer) (*ILKREGISTRYFilterer, error) {
	contract, err := bindILKREGISTRY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYFilterer{contract: contract}, nil
}

// bindILKREGISTRY binds a generic wrapper to an already deployed contract.
func bindILKREGISTRY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILKREGISTRYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILKREGISTRY *ILKREGISTRYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILKREGISTRY.Contract.ILKREGISTRYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILKREGISTRY *ILKREGISTRYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.ILKREGISTRYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILKREGISTRY *ILKREGISTRYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.ILKREGISTRYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILKREGISTRY *ILKREGISTRYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILKREGISTRY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILKREGISTRY *ILKREGISTRYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILKREGISTRY *ILKREGISTRYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.contract.Transact(opts, method, params...)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Cat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "cat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Cat() (common.Address, error) {
	return _ILKREGISTRY.Contract.Cat(&_ILKREGISTRY.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Cat() (common.Address, error) {
	return _ILKREGISTRY.Contract.Cat(&_ILKREGISTRY.CallOpts)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCaller) Class(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "class", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYSession) Class(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Class(&_ILKREGISTRY.CallOpts, ilk)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Class(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Class(&_ILKREGISTRY.CallOpts, ilk)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYSession) Count() (*big.Int, error) {
	return _ILKREGISTRY.Contract.Count(&_ILKREGISTRY.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Count() (*big.Int, error) {
	return _ILKREGISTRY.Contract.Count(&_ILKREGISTRY.CallOpts)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCaller) Dec(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "dec", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYSession) Dec(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Dec(&_ILKREGISTRY.CallOpts, ilk)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Dec(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Dec(&_ILKREGISTRY.CallOpts, ilk)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Dog() (common.Address, error) {
	return _ILKREGISTRY.Contract.Dog(&_ILKREGISTRY.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Dog() (common.Address, error) {
	return _ILKREGISTRY.Contract.Dog(&_ILKREGISTRY.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Gem(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "gem", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Gem(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Gem(&_ILKREGISTRY.CallOpts, ilk)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Gem(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Gem(&_ILKREGISTRY.CallOpts, ilk)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_ILKREGISTRY *ILKREGISTRYCaller) Get(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "get", pos)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_ILKREGISTRY *ILKREGISTRYSession) Get(pos *big.Int) ([32]byte, error) {
	return _ILKREGISTRY.Contract.Get(&_ILKREGISTRY.CallOpts, pos)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Get(pos *big.Int) ([32]byte, error) {
	return _ILKREGISTRY.Contract.Get(&_ILKREGISTRY.CallOpts, pos)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_ILKREGISTRY *ILKREGISTRYCaller) IlkData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "ilkData", arg0)

	outstruct := new(struct {
		Pos    *big.Int
		Join   common.Address
		Gem    common.Address
		Dec    uint8
		Class  *big.Int
		Pip    common.Address
		Xlip   common.Address
		Name   string
		Symbol string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pos = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Join = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gem = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Dec = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Class = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[8], new(string)).(*string)

	return *outstruct, err

}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_ILKREGISTRY *ILKREGISTRYSession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _ILKREGISTRY.Contract.IlkData(&_ILKREGISTRY.CallOpts, arg0)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _ILKREGISTRY.Contract.IlkData(&_ILKREGISTRY.CallOpts, arg0)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_ILKREGISTRY *ILKREGISTRYCaller) Info(opts *bind.CallOpts, ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "info", ilk)

	outstruct := new(struct {
		Name   string
		Symbol string
		Class  *big.Int
		Dec    *big.Int
		Gem    common.Address
		Pip    common.Address
		Join   common.Address
		Xlip   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Class = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dec = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Gem = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Join = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_ILKREGISTRY *ILKREGISTRYSession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _ILKREGISTRY.Contract.Info(&_ILKREGISTRY.CallOpts, ilk)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _ILKREGISTRY.Contract.Info(&_ILKREGISTRY.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Join(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "join", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Join(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Join(&_ILKREGISTRY.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Join(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Join(&_ILKREGISTRY.CallOpts, ilk)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYCaller) List(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYSession) List() ([][32]byte, error) {
	return _ILKREGISTRY.Contract.List(&_ILKREGISTRY.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYCallerSession) List() ([][32]byte, error) {
	return _ILKREGISTRY.Contract.List(&_ILKREGISTRY.CallOpts)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYCaller) List0(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "list0", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYSession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _ILKREGISTRY.Contract.List0(&_ILKREGISTRY.CallOpts, start, end)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_ILKREGISTRY *ILKREGISTRYCallerSession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _ILKREGISTRY.Contract.List0(&_ILKREGISTRY.CallOpts, start, end)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYCaller) Name(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "name", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYSession) Name(ilk [32]byte) (string, error) {
	return _ILKREGISTRY.Contract.Name(&_ILKREGISTRY.CallOpts, ilk)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Name(ilk [32]byte) (string, error) {
	return _ILKREGISTRY.Contract.Name(&_ILKREGISTRY.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Pip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "pip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Pip(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Pip(&_ILKREGISTRY.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Pip(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Pip(&_ILKREGISTRY.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCaller) Pos(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "pos", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYSession) Pos(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Pos(&_ILKREGISTRY.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Pos(ilk [32]byte) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Pos(&_ILKREGISTRY.CallOpts, ilk)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Spot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "spot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Spot() (common.Address, error) {
	return _ILKREGISTRY.Contract.Spot(&_ILKREGISTRY.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Spot() (common.Address, error) {
	return _ILKREGISTRY.Contract.Spot(&_ILKREGISTRY.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYCaller) Symbol(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "symbol", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYSession) Symbol(ilk [32]byte) (string, error) {
	return _ILKREGISTRY.Contract.Symbol(&_ILKREGISTRY.CallOpts, ilk)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Symbol(ilk [32]byte) (string, error) {
	return _ILKREGISTRY.Contract.Symbol(&_ILKREGISTRY.CallOpts, ilk)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Vat() (common.Address, error) {
	return _ILKREGISTRY.Contract.Vat(&_ILKREGISTRY.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Vat() (common.Address, error) {
	return _ILKREGISTRY.Contract.Vat(&_ILKREGISTRY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Wards(&_ILKREGISTRY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ILKREGISTRY.Contract.Wards(&_ILKREGISTRY.CallOpts, arg0)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCaller) Xlip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ILKREGISTRY.contract.Call(opts, &out, "xlip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYSession) Xlip(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Xlip(&_ILKREGISTRY.CallOpts, ilk)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_ILKREGISTRY *ILKREGISTRYCallerSession) Xlip(ilk [32]byte) (common.Address, error) {
	return _ILKREGISTRY.Contract.Xlip(&_ILKREGISTRY.CallOpts, ilk)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Add(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "add", adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Add(adapter common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Add(&_ILKREGISTRY.TransactOpts, adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Add(adapter common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Add(&_ILKREGISTRY.TransactOpts, adapter)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Deny(&_ILKREGISTRY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Deny(&_ILKREGISTRY.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) File0(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "file0", ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File0(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File0(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File1(&_ILKREGISTRY.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File1(&_ILKREGISTRY.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "file2", ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File2(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.File2(&_ILKREGISTRY.TransactOpts, ilk, what, data)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Put(opts *bind.TransactOpts, _ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "put", _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Put(&_ILKREGISTRY.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Put(&_ILKREGISTRY.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Rely(&_ILKREGISTRY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Rely(&_ILKREGISTRY.TransactOpts, usr)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Remove(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "remove", ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Remove(&_ILKREGISTRY.TransactOpts, ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Remove(&_ILKREGISTRY.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) RemoveAuth(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "removeAuth", ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.RemoveAuth(&_ILKREGISTRY.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.RemoveAuth(&_ILKREGISTRY.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactor) Update(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.contract.Transact(opts, "update", ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYSession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Update(&_ILKREGISTRY.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_ILKREGISTRY *ILKREGISTRYTransactorSession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _ILKREGISTRY.Contract.Update(&_ILKREGISTRY.TransactOpts, ilk)
}

// ILKREGISTRYAddIlkIterator is returned from FilterAddIlk and is used to iterate over the raw logs and unpacked data for AddIlk events raised by the ILKREGISTRY contract.
type ILKREGISTRYAddIlkIterator struct {
	Event *ILKREGISTRYAddIlk // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYAddIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYAddIlk)
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
		it.Event = new(ILKREGISTRYAddIlk)
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
func (it *ILKREGISTRYAddIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYAddIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYAddIlk represents a AddIlk event raised by the ILKREGISTRY contract.
type ILKREGISTRYAddIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddIlk is a free log retrieval operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterAddIlk(opts *bind.FilterOpts) (*ILKREGISTRYAddIlkIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "AddIlk")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYAddIlkIterator{contract: _ILKREGISTRY.contract, event: "AddIlk", logs: logs, sub: sub}, nil
}

// WatchAddIlk is a free log subscription operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchAddIlk(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYAddIlk) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "AddIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYAddIlk)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "AddIlk", log); err != nil {
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

// ParseAddIlk is a log parse operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseAddIlk(log types.Log) (*ILKREGISTRYAddIlk, error) {
	event := new(ILKREGISTRYAddIlk)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "AddIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ILKREGISTRY contract.
type ILKREGISTRYDenyIterator struct {
	Event *ILKREGISTRYDeny // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYDeny)
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
		it.Event = new(ILKREGISTRYDeny)
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
func (it *ILKREGISTRYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYDeny represents a Deny event raised by the ILKREGISTRY contract.
type ILKREGISTRYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterDeny(opts *bind.FilterOpts) (*ILKREGISTRYDenyIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYDenyIterator{contract: _ILKREGISTRY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYDeny) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYDeny)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseDeny(log types.Log) (*ILKREGISTRYDeny, error) {
	event := new(ILKREGISTRYDeny)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the ILKREGISTRY contract.
type ILKREGISTRYFileIterator struct {
	Event *ILKREGISTRYFile // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYFile)
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
		it.Event = new(ILKREGISTRYFile)
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
func (it *ILKREGISTRYFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYFile represents a File event raised by the ILKREGISTRY contract.
type ILKREGISTRYFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterFile(opts *bind.FilterOpts) (*ILKREGISTRYFileIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "File")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYFileIterator{contract: _ILKREGISTRY.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYFile) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "File")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYFile)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "File", log); err != nil {
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
// Solidity: event File(bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseFile(log types.Log) (*ILKREGISTRYFile, error) {
	event := new(ILKREGISTRYFile)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the ILKREGISTRY contract.
type ILKREGISTRYFile0Iterator struct {
	Event *ILKREGISTRYFile0 // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYFile0)
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
		it.Event = new(ILKREGISTRYFile0)
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
func (it *ILKREGISTRYFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYFile0 represents a File0 event raised by the ILKREGISTRY contract.
type ILKREGISTRYFile0 struct {
	Ilk  [32]byte
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterFile0(opts *bind.FilterOpts) (*ILKREGISTRYFile0Iterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "File0")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYFile0Iterator{contract: _ILKREGISTRY.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYFile0) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "File0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYFile0)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseFile0(log types.Log) (*ILKREGISTRYFile0, error) {
	event := new(ILKREGISTRYFile0)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the ILKREGISTRY contract.
type ILKREGISTRYFile1Iterator struct {
	Event *ILKREGISTRYFile1 // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYFile1)
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
		it.Event = new(ILKREGISTRYFile1)
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
func (it *ILKREGISTRYFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYFile1 represents a File1 event raised by the ILKREGISTRY contract.
type ILKREGISTRYFile1 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterFile1(opts *bind.FilterOpts) (*ILKREGISTRYFile1Iterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYFile1Iterator{contract: _ILKREGISTRY.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYFile1) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYFile1)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "File1", log); err != nil {
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
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseFile1(log types.Log) (*ILKREGISTRYFile1, error) {
	event := new(ILKREGISTRYFile1)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYFile2Iterator is returned from FilterFile2 and is used to iterate over the raw logs and unpacked data for File2 events raised by the ILKREGISTRY contract.
type ILKREGISTRYFile2Iterator struct {
	Event *ILKREGISTRYFile2 // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYFile2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYFile2)
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
		it.Event = new(ILKREGISTRYFile2)
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
func (it *ILKREGISTRYFile2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYFile2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYFile2 represents a File2 event raised by the ILKREGISTRY contract.
type ILKREGISTRYFile2 struct {
	Ilk  [32]byte
	What [32]byte
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile2 is a free log retrieval operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterFile2(opts *bind.FilterOpts) (*ILKREGISTRYFile2Iterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "File2")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYFile2Iterator{contract: _ILKREGISTRY.contract, event: "File2", logs: logs, sub: sub}, nil
}

// WatchFile2 is a free log subscription operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchFile2(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYFile2) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "File2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYFile2)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "File2", log); err != nil {
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

// ParseFile2 is a log parse operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseFile2(log types.Log) (*ILKREGISTRYFile2, error) {
	event := new(ILKREGISTRYFile2)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "File2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYNameErrorIterator is returned from FilterNameError and is used to iterate over the raw logs and unpacked data for NameError events raised by the ILKREGISTRY contract.
type ILKREGISTRYNameErrorIterator struct {
	Event *ILKREGISTRYNameError // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYNameErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYNameError)
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
		it.Event = new(ILKREGISTRYNameError)
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
func (it *ILKREGISTRYNameErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYNameErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYNameError represents a NameError event raised by the ILKREGISTRY contract.
type ILKREGISTRYNameError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNameError is a free log retrieval operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterNameError(opts *bind.FilterOpts) (*ILKREGISTRYNameErrorIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "NameError")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYNameErrorIterator{contract: _ILKREGISTRY.contract, event: "NameError", logs: logs, sub: sub}, nil
}

// WatchNameError is a free log subscription operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchNameError(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYNameError) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "NameError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYNameError)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "NameError", log); err != nil {
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

// ParseNameError is a log parse operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseNameError(log types.Log) (*ILKREGISTRYNameError, error) {
	event := new(ILKREGISTRYNameError)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "NameError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ILKREGISTRY contract.
type ILKREGISTRYRelyIterator struct {
	Event *ILKREGISTRYRely // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYRely)
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
		it.Event = new(ILKREGISTRYRely)
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
func (it *ILKREGISTRYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYRely represents a Rely event raised by the ILKREGISTRY contract.
type ILKREGISTRYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterRely(opts *bind.FilterOpts) (*ILKREGISTRYRelyIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYRelyIterator{contract: _ILKREGISTRY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYRely) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYRely)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseRely(log types.Log) (*ILKREGISTRYRely, error) {
	event := new(ILKREGISTRYRely)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYRemoveIlkIterator is returned from FilterRemoveIlk and is used to iterate over the raw logs and unpacked data for RemoveIlk events raised by the ILKREGISTRY contract.
type ILKREGISTRYRemoveIlkIterator struct {
	Event *ILKREGISTRYRemoveIlk // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYRemoveIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYRemoveIlk)
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
		it.Event = new(ILKREGISTRYRemoveIlk)
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
func (it *ILKREGISTRYRemoveIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYRemoveIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYRemoveIlk represents a RemoveIlk event raised by the ILKREGISTRY contract.
type ILKREGISTRYRemoveIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveIlk is a free log retrieval operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterRemoveIlk(opts *bind.FilterOpts) (*ILKREGISTRYRemoveIlkIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "RemoveIlk")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYRemoveIlkIterator{contract: _ILKREGISTRY.contract, event: "RemoveIlk", logs: logs, sub: sub}, nil
}

// WatchRemoveIlk is a free log subscription operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchRemoveIlk(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYRemoveIlk) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "RemoveIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYRemoveIlk)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
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

// ParseRemoveIlk is a log parse operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseRemoveIlk(log types.Log) (*ILKREGISTRYRemoveIlk, error) {
	event := new(ILKREGISTRYRemoveIlk)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYSymbolErrorIterator is returned from FilterSymbolError and is used to iterate over the raw logs and unpacked data for SymbolError events raised by the ILKREGISTRY contract.
type ILKREGISTRYSymbolErrorIterator struct {
	Event *ILKREGISTRYSymbolError // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYSymbolErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYSymbolError)
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
		it.Event = new(ILKREGISTRYSymbolError)
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
func (it *ILKREGISTRYSymbolErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYSymbolErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYSymbolError represents a SymbolError event raised by the ILKREGISTRY contract.
type ILKREGISTRYSymbolError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSymbolError is a free log retrieval operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterSymbolError(opts *bind.FilterOpts) (*ILKREGISTRYSymbolErrorIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "SymbolError")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYSymbolErrorIterator{contract: _ILKREGISTRY.contract, event: "SymbolError", logs: logs, sub: sub}, nil
}

// WatchSymbolError is a free log subscription operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchSymbolError(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYSymbolError) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "SymbolError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYSymbolError)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "SymbolError", log); err != nil {
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

// ParseSymbolError is a log parse operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseSymbolError(log types.Log) (*ILKREGISTRYSymbolError, error) {
	event := new(ILKREGISTRYSymbolError)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "SymbolError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILKREGISTRYUpdateIlkIterator is returned from FilterUpdateIlk and is used to iterate over the raw logs and unpacked data for UpdateIlk events raised by the ILKREGISTRY contract.
type ILKREGISTRYUpdateIlkIterator struct {
	Event *ILKREGISTRYUpdateIlk // Event containing the contract specifics and raw log

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
func (it *ILKREGISTRYUpdateIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILKREGISTRYUpdateIlk)
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
		it.Event = new(ILKREGISTRYUpdateIlk)
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
func (it *ILKREGISTRYUpdateIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILKREGISTRYUpdateIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILKREGISTRYUpdateIlk represents a UpdateIlk event raised by the ILKREGISTRY contract.
type ILKREGISTRYUpdateIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateIlk is a free log retrieval operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) FilterUpdateIlk(opts *bind.FilterOpts) (*ILKREGISTRYUpdateIlkIterator, error) {

	logs, sub, err := _ILKREGISTRY.contract.FilterLogs(opts, "UpdateIlk")
	if err != nil {
		return nil, err
	}
	return &ILKREGISTRYUpdateIlkIterator{contract: _ILKREGISTRY.contract, event: "UpdateIlk", logs: logs, sub: sub}, nil
}

// WatchUpdateIlk is a free log subscription operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) WatchUpdateIlk(opts *bind.WatchOpts, sink chan<- *ILKREGISTRYUpdateIlk) (event.Subscription, error) {

	logs, sub, err := _ILKREGISTRY.contract.WatchLogs(opts, "UpdateIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILKREGISTRYUpdateIlk)
				if err := _ILKREGISTRY.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
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

// ParseUpdateIlk is a log parse operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_ILKREGISTRY *ILKREGISTRYFilterer) ParseUpdateIlk(log types.Log) (*ILKREGISTRYUpdateIlk, error) {
	event := new(ILKREGISTRYUpdateIlk)
	if err := _ILKREGISTRY.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
