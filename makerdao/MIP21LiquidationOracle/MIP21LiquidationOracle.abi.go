// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MIP21LiquidationOracle

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

// MIP21LIQUIDATIONORACLEABI is the input ABI used to generate the binding from.
const MIP21LIQUIDATIONORACLEABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"Bump\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"Cull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Cure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"tau\",\"type\":\"uint48\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Tell\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"bump\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"cull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"cure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"good\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tau\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"toc\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"tau\",\"type\":\"uint48\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"tell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatAbstract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MIP21LIQUIDATIONORACLE is an auto generated Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLE struct {
	MIP21LIQUIDATIONORACLECaller     // Read-only binding to the contract
	MIP21LIQUIDATIONORACLETransactor // Write-only binding to the contract
	MIP21LIQUIDATIONORACLEFilterer   // Log filterer for contract events
}

// MIP21LIQUIDATIONORACLECaller is an auto generated read-only Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIP21LIQUIDATIONORACLETransactor is an auto generated write-only Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIP21LIQUIDATIONORACLEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MIP21LIQUIDATIONORACLEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIP21LIQUIDATIONORACLESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MIP21LIQUIDATIONORACLESession struct {
	Contract     *MIP21LIQUIDATIONORACLE // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MIP21LIQUIDATIONORACLECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MIP21LIQUIDATIONORACLECallerSession struct {
	Contract *MIP21LIQUIDATIONORACLECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MIP21LIQUIDATIONORACLETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MIP21LIQUIDATIONORACLETransactorSession struct {
	Contract     *MIP21LIQUIDATIONORACLETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MIP21LIQUIDATIONORACLERaw is an auto generated low-level Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLERaw struct {
	Contract *MIP21LIQUIDATIONORACLE // Generic contract binding to access the raw methods on
}

// MIP21LIQUIDATIONORACLECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLECallerRaw struct {
	Contract *MIP21LIQUIDATIONORACLECaller // Generic read-only contract binding to access the raw methods on
}

// MIP21LIQUIDATIONORACLETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MIP21LIQUIDATIONORACLETransactorRaw struct {
	Contract *MIP21LIQUIDATIONORACLETransactor // Generic write-only contract binding to access the raw methods on
}

// NewMIP21LIQUIDATIONORACLE creates a new instance of MIP21LIQUIDATIONORACLE, bound to a specific deployed contract.
func NewMIP21LIQUIDATIONORACLE(address common.Address, backend bind.ContractBackend) (*MIP21LIQUIDATIONORACLE, error) {
	contract, err := bindMIP21LIQUIDATIONORACLE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLE{MIP21LIQUIDATIONORACLECaller: MIP21LIQUIDATIONORACLECaller{contract: contract}, MIP21LIQUIDATIONORACLETransactor: MIP21LIQUIDATIONORACLETransactor{contract: contract}, MIP21LIQUIDATIONORACLEFilterer: MIP21LIQUIDATIONORACLEFilterer{contract: contract}}, nil
}

// NewMIP21LIQUIDATIONORACLECaller creates a new read-only instance of MIP21LIQUIDATIONORACLE, bound to a specific deployed contract.
func NewMIP21LIQUIDATIONORACLECaller(address common.Address, caller bind.ContractCaller) (*MIP21LIQUIDATIONORACLECaller, error) {
	contract, err := bindMIP21LIQUIDATIONORACLE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLECaller{contract: contract}, nil
}

// NewMIP21LIQUIDATIONORACLETransactor creates a new write-only instance of MIP21LIQUIDATIONORACLE, bound to a specific deployed contract.
func NewMIP21LIQUIDATIONORACLETransactor(address common.Address, transactor bind.ContractTransactor) (*MIP21LIQUIDATIONORACLETransactor, error) {
	contract, err := bindMIP21LIQUIDATIONORACLE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLETransactor{contract: contract}, nil
}

// NewMIP21LIQUIDATIONORACLEFilterer creates a new log filterer instance of MIP21LIQUIDATIONORACLE, bound to a specific deployed contract.
func NewMIP21LIQUIDATIONORACLEFilterer(address common.Address, filterer bind.ContractFilterer) (*MIP21LIQUIDATIONORACLEFilterer, error) {
	contract, err := bindMIP21LIQUIDATIONORACLE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLEFilterer{contract: contract}, nil
}

// bindMIP21LIQUIDATIONORACLE binds a generic wrapper to an already deployed contract.
func bindMIP21LIQUIDATIONORACLE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MIP21LIQUIDATIONORACLEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIP21LIQUIDATIONORACLE.Contract.MIP21LIQUIDATIONORACLECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.MIP21LIQUIDATIONORACLETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.MIP21LIQUIDATIONORACLETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIP21LIQUIDATIONORACLE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.contract.Transact(opts, method, params...)
}

// Good is a free data retrieval call binding the contract method 0xf1fad330.
//
// Solidity: function good(bytes32 ilk) view returns(bool)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECaller) Good(opts *bind.CallOpts, ilk [32]byte) (bool, error) {
	var out []interface{}
	err := _MIP21LIQUIDATIONORACLE.contract.Call(opts, &out, "good", ilk)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Good is a free data retrieval call binding the contract method 0xf1fad330.
//
// Solidity: function good(bytes32 ilk) view returns(bool)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Good(ilk [32]byte) (bool, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Good(&_MIP21LIQUIDATIONORACLE.CallOpts, ilk)
}

// Good is a free data retrieval call binding the contract method 0xf1fad330.
//
// Solidity: function good(bytes32 ilk) view returns(bool)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerSession) Good(ilk [32]byte) (bool, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Good(&_MIP21LIQUIDATIONORACLE.CallOpts, ilk)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(string doc, address pip, uint48 tau, uint48 toc)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Doc string
	Pip common.Address
	Tau *big.Int
	Toc *big.Int
}, error) {
	var out []interface{}
	err := _MIP21LIQUIDATIONORACLE.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Doc string
		Pip common.Address
		Tau *big.Int
		Toc *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Doc = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Pip = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Tau = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Toc = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(string doc, address pip, uint48 tau, uint48 toc)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Ilks(arg0 [32]byte) (struct {
	Doc string
	Pip common.Address
	Tau *big.Int
	Toc *big.Int
}, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Ilks(&_MIP21LIQUIDATIONORACLE.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(string doc, address pip, uint48 tau, uint48 toc)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerSession) Ilks(arg0 [32]byte) (struct {
	Doc string
	Pip common.Address
	Tau *big.Int
	Toc *big.Int
}, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Ilks(&_MIP21LIQUIDATIONORACLE.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MIP21LIQUIDATIONORACLE.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Vat() (common.Address, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Vat(&_MIP21LIQUIDATIONORACLE.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerSession) Vat() (common.Address, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Vat(&_MIP21LIQUIDATIONORACLE.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MIP21LIQUIDATIONORACLE.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Vow() (common.Address, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Vow(&_MIP21LIQUIDATIONORACLE.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerSession) Vow() (common.Address, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Vow(&_MIP21LIQUIDATIONORACLE.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MIP21LIQUIDATIONORACLE.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Wards(&_MIP21LIQUIDATIONORACLE.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLECallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Wards(&_MIP21LIQUIDATIONORACLE.CallOpts, arg0)
}

// Bump is a paid mutator transaction binding the contract method 0xe9eb09f6.
//
// Solidity: function bump(bytes32 ilk, uint256 val) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Bump(opts *bind.TransactOpts, ilk [32]byte, val *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "bump", ilk, val)
}

// Bump is a paid mutator transaction binding the contract method 0xe9eb09f6.
//
// Solidity: function bump(bytes32 ilk, uint256 val) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Bump(ilk [32]byte, val *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Bump(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, val)
}

// Bump is a paid mutator transaction binding the contract method 0xe9eb09f6.
//
// Solidity: function bump(bytes32 ilk, uint256 val) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Bump(ilk [32]byte, val *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Bump(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, val)
}

// Cull is a paid mutator transaction binding the contract method 0x0d9bb087.
//
// Solidity: function cull(bytes32 ilk, address urn) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Cull(opts *bind.TransactOpts, ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "cull", ilk, urn)
}

// Cull is a paid mutator transaction binding the contract method 0x0d9bb087.
//
// Solidity: function cull(bytes32 ilk, address urn) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Cull(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Cull(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, urn)
}

// Cull is a paid mutator transaction binding the contract method 0x0d9bb087.
//
// Solidity: function cull(bytes32 ilk, address urn) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Cull(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Cull(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, urn)
}

// Cure is a paid mutator transaction binding the contract method 0x9fcdeba6.
//
// Solidity: function cure(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Cure(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "cure", ilk)
}

// Cure is a paid mutator transaction binding the contract method 0x9fcdeba6.
//
// Solidity: function cure(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Cure(ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Cure(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk)
}

// Cure is a paid mutator transaction binding the contract method 0x9fcdeba6.
//
// Solidity: function cure(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Cure(ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Cure(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Deny(&_MIP21LIQUIDATIONORACLE.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Deny(&_MIP21LIQUIDATIONORACLE.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) File(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.File(&_MIP21LIQUIDATIONORACLE.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.File(&_MIP21LIQUIDATIONORACLE.TransactOpts, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x3304e282.
//
// Solidity: function init(bytes32 ilk, uint256 val, string doc, uint48 tau) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Init(opts *bind.TransactOpts, ilk [32]byte, val *big.Int, doc string, tau *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "init", ilk, val, doc, tau)
}

// Init is a paid mutator transaction binding the contract method 0x3304e282.
//
// Solidity: function init(bytes32 ilk, uint256 val, string doc, uint48 tau) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Init(ilk [32]byte, val *big.Int, doc string, tau *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Init(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, val, doc, tau)
}

// Init is a paid mutator transaction binding the contract method 0x3304e282.
//
// Solidity: function init(bytes32 ilk, uint256 val, string doc, uint48 tau) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Init(ilk [32]byte, val *big.Int, doc string, tau *big.Int) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Init(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk, val, doc, tau)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Rely(&_MIP21LIQUIDATIONORACLE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Rely(&_MIP21LIQUIDATIONORACLE.TransactOpts, usr)
}

// Tell is a paid mutator transaction binding the contract method 0xc881986c.
//
// Solidity: function tell(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactor) Tell(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.contract.Transact(opts, "tell", ilk)
}

// Tell is a paid mutator transaction binding the contract method 0xc881986c.
//
// Solidity: function tell(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLESession) Tell(ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Tell(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk)
}

// Tell is a paid mutator transaction binding the contract method 0xc881986c.
//
// Solidity: function tell(bytes32 ilk) returns()
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLETransactorSession) Tell(ilk [32]byte) (*types.Transaction, error) {
	return _MIP21LIQUIDATIONORACLE.Contract.Tell(&_MIP21LIQUIDATIONORACLE.TransactOpts, ilk)
}

// MIP21LIQUIDATIONORACLEBumpIterator is returned from FilterBump and is used to iterate over the raw logs and unpacked data for Bump events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEBumpIterator struct {
	Event *MIP21LIQUIDATIONORACLEBump // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLEBumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLEBump)
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
		it.Event = new(MIP21LIQUIDATIONORACLEBump)
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
func (it *MIP21LIQUIDATIONORACLEBumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLEBumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLEBump represents a Bump event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEBump struct {
	Ilk [32]byte
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBump is a free log retrieval operation binding the contract event 0x410d2bd57d18419e442f5e911a3bc3073f2cdce69a60532dada09a762daf3f2c.
//
// Solidity: event Bump(bytes32 indexed ilk, uint256 val)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterBump(opts *bind.FilterOpts, ilk [][32]byte) (*MIP21LIQUIDATIONORACLEBumpIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Bump", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLEBumpIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Bump", logs: logs, sub: sub}, nil
}

// WatchBump is a free log subscription operation binding the contract event 0x410d2bd57d18419e442f5e911a3bc3073f2cdce69a60532dada09a762daf3f2c.
//
// Solidity: event Bump(bytes32 indexed ilk, uint256 val)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchBump(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLEBump, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Bump", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLEBump)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Bump", log); err != nil {
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

// ParseBump is a log parse operation binding the contract event 0x410d2bd57d18419e442f5e911a3bc3073f2cdce69a60532dada09a762daf3f2c.
//
// Solidity: event Bump(bytes32 indexed ilk, uint256 val)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseBump(log types.Log) (*MIP21LIQUIDATIONORACLEBump, error) {
	event := new(MIP21LIQUIDATIONORACLEBump)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Bump", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLECullIterator is returned from FilterCull and is used to iterate over the raw logs and unpacked data for Cull events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLECullIterator struct {
	Event *MIP21LIQUIDATIONORACLECull // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLECullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLECull)
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
		it.Event = new(MIP21LIQUIDATIONORACLECull)
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
func (it *MIP21LIQUIDATIONORACLECullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLECullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLECull represents a Cull event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLECull struct {
	Ilk [32]byte
	Urn common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCull is a free log retrieval operation binding the contract event 0x9b39535ec34e14a93328be1f7e61f415eec9809b3d2a902b4c3dedb42f0c870c.
//
// Solidity: event Cull(bytes32 indexed ilk, address indexed urn)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterCull(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address) (*MIP21LIQUIDATIONORACLECullIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Cull", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLECullIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Cull", logs: logs, sub: sub}, nil
}

// WatchCull is a free log subscription operation binding the contract event 0x9b39535ec34e14a93328be1f7e61f415eec9809b3d2a902b4c3dedb42f0c870c.
//
// Solidity: event Cull(bytes32 indexed ilk, address indexed urn)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchCull(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLECull, ilk [][32]byte, urn []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Cull", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLECull)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Cull", log); err != nil {
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

// ParseCull is a log parse operation binding the contract event 0x9b39535ec34e14a93328be1f7e61f415eec9809b3d2a902b4c3dedb42f0c870c.
//
// Solidity: event Cull(bytes32 indexed ilk, address indexed urn)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseCull(log types.Log) (*MIP21LIQUIDATIONORACLECull, error) {
	event := new(MIP21LIQUIDATIONORACLECull)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Cull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLECureIterator is returned from FilterCure and is used to iterate over the raw logs and unpacked data for Cure events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLECureIterator struct {
	Event *MIP21LIQUIDATIONORACLECure // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLECureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLECure)
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
		it.Event = new(MIP21LIQUIDATIONORACLECure)
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
func (it *MIP21LIQUIDATIONORACLECureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLECureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLECure represents a Cure event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLECure struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCure is a free log retrieval operation binding the contract event 0x59cc51b6b7601a171c376803cec950d480385ea3486ba0c34e0b7bdb44833472.
//
// Solidity: event Cure(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterCure(opts *bind.FilterOpts, ilk [][32]byte) (*MIP21LIQUIDATIONORACLECureIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Cure", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLECureIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Cure", logs: logs, sub: sub}, nil
}

// WatchCure is a free log subscription operation binding the contract event 0x59cc51b6b7601a171c376803cec950d480385ea3486ba0c34e0b7bdb44833472.
//
// Solidity: event Cure(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchCure(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLECure, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Cure", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLECure)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Cure", log); err != nil {
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

// ParseCure is a log parse operation binding the contract event 0x59cc51b6b7601a171c376803cec950d480385ea3486ba0c34e0b7bdb44833472.
//
// Solidity: event Cure(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseCure(log types.Log) (*MIP21LIQUIDATIONORACLECure, error) {
	event := new(MIP21LIQUIDATIONORACLECure)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Cure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLEDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEDenyIterator struct {
	Event *MIP21LIQUIDATIONORACLEDeny // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLEDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLEDeny)
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
		it.Event = new(MIP21LIQUIDATIONORACLEDeny)
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
func (it *MIP21LIQUIDATIONORACLEDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLEDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLEDeny represents a Deny event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MIP21LIQUIDATIONORACLEDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLEDenyIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLEDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLEDeny)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseDeny(log types.Log) (*MIP21LIQUIDATIONORACLEDeny, error) {
	event := new(MIP21LIQUIDATIONORACLEDeny)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLEFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEFileIterator struct {
	Event *MIP21LIQUIDATIONORACLEFile // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLEFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLEFile)
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
		it.Event = new(MIP21LIQUIDATIONORACLEFile)
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
func (it *MIP21LIQUIDATIONORACLEFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLEFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLEFile represents a File event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MIP21LIQUIDATIONORACLEFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLEFileIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLEFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLEFile)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseFile(log types.Log) (*MIP21LIQUIDATIONORACLEFile, error) {
	event := new(MIP21LIQUIDATIONORACLEFile)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLEInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEInitIterator struct {
	Event *MIP21LIQUIDATIONORACLEInit // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLEInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLEInit)
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
		it.Event = new(MIP21LIQUIDATIONORACLEInit)
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
func (it *MIP21LIQUIDATIONORACLEInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLEInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLEInit represents a Init event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLEInit struct {
	Ilk [32]byte
	Val *big.Int
	Doc string
	Tau *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x8a99605e1142f58396e9f93017f6597e6f0cef6d47f77cc422033d4f80331ae4.
//
// Solidity: event Init(bytes32 indexed ilk, uint256 val, string doc, uint48 tau)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterInit(opts *bind.FilterOpts, ilk [][32]byte) (*MIP21LIQUIDATIONORACLEInitIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLEInitIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x8a99605e1142f58396e9f93017f6597e6f0cef6d47f77cc422033d4f80331ae4.
//
// Solidity: event Init(bytes32 indexed ilk, uint256 val, string doc, uint48 tau)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLEInit, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLEInit)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x8a99605e1142f58396e9f93017f6597e6f0cef6d47f77cc422033d4f80331ae4.
//
// Solidity: event Init(bytes32 indexed ilk, uint256 val, string doc, uint48 tau)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseInit(log types.Log) (*MIP21LIQUIDATIONORACLEInit, error) {
	event := new(MIP21LIQUIDATIONORACLEInit)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLERelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLERelyIterator struct {
	Event *MIP21LIQUIDATIONORACLERely // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLERelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLERely)
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
		it.Event = new(MIP21LIQUIDATIONORACLERely)
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
func (it *MIP21LIQUIDATIONORACLERelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLERelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLERely represents a Rely event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLERely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MIP21LIQUIDATIONORACLERelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLERelyIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLERely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLERely)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseRely(log types.Log) (*MIP21LIQUIDATIONORACLERely, error) {
	event := new(MIP21LIQUIDATIONORACLERely)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MIP21LIQUIDATIONORACLETellIterator is returned from FilterTell and is used to iterate over the raw logs and unpacked data for Tell events raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLETellIterator struct {
	Event *MIP21LIQUIDATIONORACLETell // Event containing the contract specifics and raw log

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
func (it *MIP21LIQUIDATIONORACLETellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MIP21LIQUIDATIONORACLETell)
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
		it.Event = new(MIP21LIQUIDATIONORACLETell)
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
func (it *MIP21LIQUIDATIONORACLETellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MIP21LIQUIDATIONORACLETellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MIP21LIQUIDATIONORACLETell represents a Tell event raised by the MIP21LIQUIDATIONORACLE contract.
type MIP21LIQUIDATIONORACLETell struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTell is a free log retrieval operation binding the contract event 0xc6849e4b55ae3f3ef2563004433f7b6b01e6ead3ef575d400d4647edc7ede129.
//
// Solidity: event Tell(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) FilterTell(opts *bind.FilterOpts, ilk [][32]byte) (*MIP21LIQUIDATIONORACLETellIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.FilterLogs(opts, "Tell", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MIP21LIQUIDATIONORACLETellIterator{contract: _MIP21LIQUIDATIONORACLE.contract, event: "Tell", logs: logs, sub: sub}, nil
}

// WatchTell is a free log subscription operation binding the contract event 0xc6849e4b55ae3f3ef2563004433f7b6b01e6ead3ef575d400d4647edc7ede129.
//
// Solidity: event Tell(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) WatchTell(opts *bind.WatchOpts, sink chan<- *MIP21LIQUIDATIONORACLETell, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MIP21LIQUIDATIONORACLE.contract.WatchLogs(opts, "Tell", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MIP21LIQUIDATIONORACLETell)
				if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Tell", log); err != nil {
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

// ParseTell is a log parse operation binding the contract event 0xc6849e4b55ae3f3ef2563004433f7b6b01e6ead3ef575d400d4647edc7ede129.
//
// Solidity: event Tell(bytes32 indexed ilk)
func (_MIP21LIQUIDATIONORACLE *MIP21LIQUIDATIONORACLEFilterer) ParseTell(log types.Log) (*MIP21LIQUIDATIONORACLETell, error) {
	event := new(MIP21LIQUIDATIONORACLETell)
	if err := _MIP21LIQUIDATIONORACLE.contract.UnpackLog(event, "Tell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
