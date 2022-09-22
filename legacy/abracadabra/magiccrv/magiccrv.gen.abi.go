// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package magiccrv

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MagiccrvMetaData contains all meta data concerning the Magiccrv contract.
var MagiccrvMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractICurveVoter\",\"name\":\"_curveVoter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CRV\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveVoter\",\"outputs\":[{\"internalType\":\"contractICurveVoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"mintFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCRVTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MagiccrvABI is the input ABI used to generate the binding from.
// Deprecated: Use MagiccrvMetaData.ABI instead.
var MagiccrvABI = MagiccrvMetaData.ABI

// Magiccrv is an auto generated Go binding around an Ethereum contract.
type Magiccrv struct {
	MagiccrvCaller     // Read-only binding to the contract
	MagiccrvTransactor // Write-only binding to the contract
	MagiccrvFilterer   // Log filterer for contract events
}

// MagiccrvCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagiccrvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagiccrvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagiccrvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagiccrvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagiccrvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagiccrvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagiccrvSession struct {
	Contract     *Magiccrv         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MagiccrvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagiccrvCallerSession struct {
	Contract *MagiccrvCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MagiccrvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagiccrvTransactorSession struct {
	Contract     *MagiccrvTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MagiccrvRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagiccrvRaw struct {
	Contract *Magiccrv // Generic contract binding to access the raw methods on
}

// MagiccrvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagiccrvCallerRaw struct {
	Contract *MagiccrvCaller // Generic read-only contract binding to access the raw methods on
}

// MagiccrvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagiccrvTransactorRaw struct {
	Contract *MagiccrvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagiccrv creates a new instance of Magiccrv, bound to a specific deployed contract.
func NewMagiccrv(address common.Address, backend bind.ContractBackend) (*Magiccrv, error) {
	contract, err := bindMagiccrv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Magiccrv{MagiccrvCaller: MagiccrvCaller{contract: contract}, MagiccrvTransactor: MagiccrvTransactor{contract: contract}, MagiccrvFilterer: MagiccrvFilterer{contract: contract}}, nil
}

// NewMagiccrvCaller creates a new read-only instance of Magiccrv, bound to a specific deployed contract.
func NewMagiccrvCaller(address common.Address, caller bind.ContractCaller) (*MagiccrvCaller, error) {
	contract, err := bindMagiccrv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagiccrvCaller{contract: contract}, nil
}

// NewMagiccrvTransactor creates a new write-only instance of Magiccrv, bound to a specific deployed contract.
func NewMagiccrvTransactor(address common.Address, transactor bind.ContractTransactor) (*MagiccrvTransactor, error) {
	contract, err := bindMagiccrv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagiccrvTransactor{contract: contract}, nil
}

// NewMagiccrvFilterer creates a new log filterer instance of Magiccrv, bound to a specific deployed contract.
func NewMagiccrvFilterer(address common.Address, filterer bind.ContractFilterer) (*MagiccrvFilterer, error) {
	contract, err := bindMagiccrv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagiccrvFilterer{contract: contract}, nil
}

// bindMagiccrv binds a generic wrapper to an already deployed contract.
func bindMagiccrv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MagiccrvABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Magiccrv *MagiccrvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Magiccrv.Contract.MagiccrvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Magiccrv *MagiccrvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiccrv.Contract.MagiccrvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Magiccrv *MagiccrvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Magiccrv.Contract.MagiccrvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Magiccrv *MagiccrvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Magiccrv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Magiccrv *MagiccrvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiccrv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Magiccrv *MagiccrvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Magiccrv.Contract.contract.Transact(opts, method, params...)
}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_Magiccrv *MagiccrvCaller) CRV(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "CRV")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_Magiccrv *MagiccrvSession) CRV() (common.Address, error) {
	return _Magiccrv.Contract.CRV(&_Magiccrv.CallOpts)
}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_Magiccrv *MagiccrvCallerSession) CRV() (common.Address, error) {
	return _Magiccrv.Contract.CRV(&_Magiccrv.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Magiccrv *MagiccrvCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Magiccrv *MagiccrvSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Magiccrv.Contract.DOMAINSEPARATOR(&_Magiccrv.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Magiccrv *MagiccrvCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Magiccrv.Contract.DOMAINSEPARATOR(&_Magiccrv.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Magiccrv *MagiccrvCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Magiccrv *MagiccrvSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Magiccrv.Contract.PERMITTYPEHASH(&_Magiccrv.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Magiccrv *MagiccrvCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Magiccrv.Contract.PERMITTYPEHASH(&_Magiccrv.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Magiccrv *MagiccrvCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Magiccrv *MagiccrvSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.Allowance(&_Magiccrv.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Magiccrv *MagiccrvCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.Allowance(&_Magiccrv.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Magiccrv *MagiccrvCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Magiccrv *MagiccrvSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.BalanceOf(&_Magiccrv.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Magiccrv *MagiccrvCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.BalanceOf(&_Magiccrv.CallOpts, arg0)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_Magiccrv *MagiccrvCaller) CurveVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "curveVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_Magiccrv *MagiccrvSession) CurveVoter() (common.Address, error) {
	return _Magiccrv.Contract.CurveVoter(&_Magiccrv.CallOpts)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_Magiccrv *MagiccrvCallerSession) CurveVoter() (common.Address, error) {
	return _Magiccrv.Contract.CurveVoter(&_Magiccrv.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Magiccrv *MagiccrvCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Magiccrv *MagiccrvSession) Decimals() (uint8, error) {
	return _Magiccrv.Contract.Decimals(&_Magiccrv.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Magiccrv *MagiccrvCallerSession) Decimals() (uint8, error) {
	return _Magiccrv.Contract.Decimals(&_Magiccrv.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiccrv *MagiccrvCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiccrv *MagiccrvSession) Name() (string, error) {
	return _Magiccrv.Contract.Name(&_Magiccrv.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiccrv *MagiccrvCallerSession) Name() (string, error) {
	return _Magiccrv.Contract.Name(&_Magiccrv.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Magiccrv *MagiccrvCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Magiccrv *MagiccrvSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.Nonces(&_Magiccrv.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Magiccrv *MagiccrvCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Magiccrv.Contract.Nonces(&_Magiccrv.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiccrv *MagiccrvCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiccrv *MagiccrvSession) Symbol() (string, error) {
	return _Magiccrv.Contract.Symbol(&_Magiccrv.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiccrv *MagiccrvCallerSession) Symbol() (string, error) {
	return _Magiccrv.Contract.Symbol(&_Magiccrv.CallOpts)
}

// TotalCRVTokens is a free data retrieval call binding the contract method 0x51f3b948.
//
// Solidity: function totalCRVTokens() view returns(uint256)
func (_Magiccrv *MagiccrvCaller) TotalCRVTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "totalCRVTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCRVTokens is a free data retrieval call binding the contract method 0x51f3b948.
//
// Solidity: function totalCRVTokens() view returns(uint256)
func (_Magiccrv *MagiccrvSession) TotalCRVTokens() (*big.Int, error) {
	return _Magiccrv.Contract.TotalCRVTokens(&_Magiccrv.CallOpts)
}

// TotalCRVTokens is a free data retrieval call binding the contract method 0x51f3b948.
//
// Solidity: function totalCRVTokens() view returns(uint256)
func (_Magiccrv *MagiccrvCallerSession) TotalCRVTokens() (*big.Int, error) {
	return _Magiccrv.Contract.TotalCRVTokens(&_Magiccrv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiccrv *MagiccrvCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiccrv.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiccrv *MagiccrvSession) TotalSupply() (*big.Int, error) {
	return _Magiccrv.Contract.TotalSupply(&_Magiccrv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiccrv *MagiccrvCallerSession) TotalSupply() (*big.Int, error) {
	return _Magiccrv.Contract.TotalSupply(&_Magiccrv.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Approve(&_Magiccrv.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Approve(&_Magiccrv.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_Magiccrv *MagiccrvTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_Magiccrv *MagiccrvSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Mint(&_Magiccrv.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_Magiccrv *MagiccrvTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Mint(&_Magiccrv.TransactOpts, amount)
}

// MintFor is a paid mutator transaction binding the contract method 0xad62f1ca.
//
// Solidity: function mintFor(uint256 amount, address recipient) returns(uint256)
func (_Magiccrv *MagiccrvTransactor) MintFor(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "mintFor", amount, recipient)
}

// MintFor is a paid mutator transaction binding the contract method 0xad62f1ca.
//
// Solidity: function mintFor(uint256 amount, address recipient) returns(uint256)
func (_Magiccrv *MagiccrvSession) MintFor(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Magiccrv.Contract.MintFor(&_Magiccrv.TransactOpts, amount, recipient)
}

// MintFor is a paid mutator transaction binding the contract method 0xad62f1ca.
//
// Solidity: function mintFor(uint256 amount, address recipient) returns(uint256)
func (_Magiccrv *MagiccrvTransactorSession) MintFor(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Magiccrv.Contract.MintFor(&_Magiccrv.TransactOpts, amount, recipient)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Magiccrv *MagiccrvTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Magiccrv *MagiccrvSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Magiccrv.Contract.Permit(&_Magiccrv.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Magiccrv *MagiccrvTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Magiccrv.Contract.Permit(&_Magiccrv.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Transfer(&_Magiccrv.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.Transfer(&_Magiccrv.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.TransferFrom(&_Magiccrv.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Magiccrv *MagiccrvTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Magiccrv.Contract.TransferFrom(&_Magiccrv.TransactOpts, from, to, amount)
}

// MagiccrvApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Magiccrv contract.
type MagiccrvApprovalIterator struct {
	Event *MagiccrvApproval // Event containing the contract specifics and raw log

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
func (it *MagiccrvApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagiccrvApproval)
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
		it.Event = new(MagiccrvApproval)
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
func (it *MagiccrvApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagiccrvApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagiccrvApproval represents a Approval event raised by the Magiccrv contract.
type MagiccrvApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MagiccrvApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Magiccrv.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MagiccrvApprovalIterator{contract: _Magiccrv.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MagiccrvApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Magiccrv.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagiccrvApproval)
				if err := _Magiccrv.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) ParseApproval(log types.Log) (*MagiccrvApproval, error) {
	event := new(MagiccrvApproval)
	if err := _Magiccrv.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagiccrvTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Magiccrv contract.
type MagiccrvTransferIterator struct {
	Event *MagiccrvTransfer // Event containing the contract specifics and raw log

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
func (it *MagiccrvTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagiccrvTransfer)
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
		it.Event = new(MagiccrvTransfer)
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
func (it *MagiccrvTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagiccrvTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagiccrvTransfer represents a Transfer event raised by the Magiccrv contract.
type MagiccrvTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MagiccrvTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Magiccrv.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MagiccrvTransferIterator{contract: _Magiccrv.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MagiccrvTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Magiccrv.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagiccrvTransfer)
				if err := _Magiccrv.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Magiccrv *MagiccrvFilterer) ParseTransfer(log types.Log) (*MagiccrvTransfer, error) {
	event := new(MagiccrvTransfer)
	if err := _Magiccrv.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
