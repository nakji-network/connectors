// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OptimismL2Dai

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

// OPTIMISML2DAIABI is the input ABI used to generate the binding from.
const OPTIMISML2DAIABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deploymentChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISML2DAI is an auto generated Go binding around an Ethereum contract.
type OPTIMISML2DAI struct {
	OPTIMISML2DAICaller     // Read-only binding to the contract
	OPTIMISML2DAITransactor // Write-only binding to the contract
	OPTIMISML2DAIFilterer   // Log filterer for contract events
}

// OPTIMISML2DAICaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISML2DAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISML2DAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISML2DAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISML2DAISession struct {
	Contract     *OPTIMISML2DAI    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OPTIMISML2DAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISML2DAICallerSession struct {
	Contract *OPTIMISML2DAICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OPTIMISML2DAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISML2DAITransactorSession struct {
	Contract     *OPTIMISML2DAITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OPTIMISML2DAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISML2DAIRaw struct {
	Contract *OPTIMISML2DAI // Generic contract binding to access the raw methods on
}

// OPTIMISML2DAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISML2DAICallerRaw struct {
	Contract *OPTIMISML2DAICaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISML2DAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISML2DAITransactorRaw struct {
	Contract *OPTIMISML2DAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISML2DAI creates a new instance of OPTIMISML2DAI, bound to a specific deployed contract.
func NewOPTIMISML2DAI(address common.Address, backend bind.ContractBackend) (*OPTIMISML2DAI, error) {
	contract, err := bindOPTIMISML2DAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAI{OPTIMISML2DAICaller: OPTIMISML2DAICaller{contract: contract}, OPTIMISML2DAITransactor: OPTIMISML2DAITransactor{contract: contract}, OPTIMISML2DAIFilterer: OPTIMISML2DAIFilterer{contract: contract}}, nil
}

// NewOPTIMISML2DAICaller creates a new read-only instance of OPTIMISML2DAI, bound to a specific deployed contract.
func NewOPTIMISML2DAICaller(address common.Address, caller bind.ContractCaller) (*OPTIMISML2DAICaller, error) {
	contract, err := bindOPTIMISML2DAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAICaller{contract: contract}, nil
}

// NewOPTIMISML2DAITransactor creates a new write-only instance of OPTIMISML2DAI, bound to a specific deployed contract.
func NewOPTIMISML2DAITransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISML2DAITransactor, error) {
	contract, err := bindOPTIMISML2DAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITransactor{contract: contract}, nil
}

// NewOPTIMISML2DAIFilterer creates a new log filterer instance of OPTIMISML2DAI, bound to a specific deployed contract.
func NewOPTIMISML2DAIFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISML2DAIFilterer, error) {
	contract, err := bindOPTIMISML2DAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAIFilterer{contract: contract}, nil
}

// bindOPTIMISML2DAI binds a generic wrapper to an already deployed contract.
func bindOPTIMISML2DAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISML2DAIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2DAI *OPTIMISML2DAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2DAI.Contract.OPTIMISML2DAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2DAI *OPTIMISML2DAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.OPTIMISML2DAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2DAI *OPTIMISML2DAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.OPTIMISML2DAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2DAI *OPTIMISML2DAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2DAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) DOMAINSEPARATOR() ([32]byte, error) {
	return _OPTIMISML2DAI.Contract.DOMAINSEPARATOR(&_OPTIMISML2DAI.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _OPTIMISML2DAI.Contract.DOMAINSEPARATOR(&_OPTIMISML2DAI.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) PERMITTYPEHASH() ([32]byte, error) {
	return _OPTIMISML2DAI.Contract.PERMITTYPEHASH(&_OPTIMISML2DAI.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _OPTIMISML2DAI.Contract.PERMITTYPEHASH(&_OPTIMISML2DAI.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Allowance(&_OPTIMISML2DAI.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Allowance(&_OPTIMISML2DAI.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.BalanceOf(&_OPTIMISML2DAI.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.BalanceOf(&_OPTIMISML2DAI.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Decimals() (uint8, error) {
	return _OPTIMISML2DAI.Contract.Decimals(&_OPTIMISML2DAI.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Decimals() (uint8, error) {
	return _OPTIMISML2DAI.Contract.Decimals(&_OPTIMISML2DAI.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) DeploymentChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "deploymentChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) DeploymentChainId() (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.DeploymentChainId(&_OPTIMISML2DAI.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) DeploymentChainId() (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.DeploymentChainId(&_OPTIMISML2DAI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Name() (string, error) {
	return _OPTIMISML2DAI.Contract.Name(&_OPTIMISML2DAI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Name() (string, error) {
	return _OPTIMISML2DAI.Contract.Name(&_OPTIMISML2DAI.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Nonces(&_OPTIMISML2DAI.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Nonces(&_OPTIMISML2DAI.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Symbol() (string, error) {
	return _OPTIMISML2DAI.Contract.Symbol(&_OPTIMISML2DAI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Symbol() (string, error) {
	return _OPTIMISML2DAI.Contract.Symbol(&_OPTIMISML2DAI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) TotalSupply() (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.TotalSupply(&_OPTIMISML2DAI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) TotalSupply() (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.TotalSupply(&_OPTIMISML2DAI.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Version() (string, error) {
	return _OPTIMISML2DAI.Contract.Version(&_OPTIMISML2DAI.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Version() (string, error) {
	return _OPTIMISML2DAI.Contract.Version(&_OPTIMISML2DAI.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAI.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Wards(&_OPTIMISML2DAI.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAI *OPTIMISML2DAICallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAI.Contract.Wards(&_OPTIMISML2DAI.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Approve(&_OPTIMISML2DAI.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Approve(&_OPTIMISML2DAI.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Burn(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "burn", from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Burn(&_OPTIMISML2DAI.TransactOpts, from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Burn(&_OPTIMISML2DAI.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.DecreaseAllowance(&_OPTIMISML2DAI.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.DecreaseAllowance(&_OPTIMISML2DAI.TransactOpts, spender, subtractedValue)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Deny(&_OPTIMISML2DAI.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Deny(&_OPTIMISML2DAI.TransactOpts, usr)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.IncreaseAllowance(&_OPTIMISML2DAI.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.IncreaseAllowance(&_OPTIMISML2DAI.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Mint(&_OPTIMISML2DAI.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Mint(&_OPTIMISML2DAI.TransactOpts, to, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Permit(&_OPTIMISML2DAI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Permit(&_OPTIMISML2DAI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Rely(&_OPTIMISML2DAI.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Rely(&_OPTIMISML2DAI.TransactOpts, usr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Transfer(&_OPTIMISML2DAI.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.Transfer(&_OPTIMISML2DAI.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAISession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.TransferFrom(&_OPTIMISML2DAI.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OPTIMISML2DAI *OPTIMISML2DAITransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAI.Contract.TransferFrom(&_OPTIMISML2DAI.TransactOpts, from, to, value)
}

// OPTIMISML2DAIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIApprovalIterator struct {
	Event *OPTIMISML2DAIApproval // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAIApproval)
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
		it.Event = new(OPTIMISML2DAIApproval)
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
func (it *OPTIMISML2DAIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAIApproval represents a Approval event raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OPTIMISML2DAIApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAIApprovalIterator{contract: _OPTIMISML2DAI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAIApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAIApproval)
				if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) ParseApproval(log types.Log) (*OPTIMISML2DAIApproval, error) {
	event := new(OPTIMISML2DAIApproval)
	if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2DAIDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIDenyIterator struct {
	Event *OPTIMISML2DAIDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAIDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAIDeny)
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
		it.Event = new(OPTIMISML2DAIDeny)
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
func (it *OPTIMISML2DAIDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAIDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAIDeny represents a Deny event raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2DAIDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAIDenyIterator{contract: _OPTIMISML2DAI.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAIDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAIDeny)
				if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) ParseDeny(log types.Log) (*OPTIMISML2DAIDeny, error) {
	event := new(OPTIMISML2DAIDeny)
	if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2DAIRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIRelyIterator struct {
	Event *OPTIMISML2DAIRely // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAIRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAIRely)
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
		it.Event = new(OPTIMISML2DAIRely)
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
func (it *OPTIMISML2DAIRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAIRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAIRely represents a Rely event raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAIRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2DAIRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAIRelyIterator{contract: _OPTIMISML2DAI.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAIRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAIRely)
				if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) ParseRely(log types.Log) (*OPTIMISML2DAIRely, error) {
	event := new(OPTIMISML2DAIRely)
	if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2DAITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAITransferIterator struct {
	Event *OPTIMISML2DAITransfer // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAITransfer)
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
		it.Event = new(OPTIMISML2DAITransfer)
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
func (it *OPTIMISML2DAITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAITransfer represents a Transfer event raised by the OPTIMISML2DAI contract.
type OPTIMISML2DAITransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OPTIMISML2DAITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITransferIterator{contract: _OPTIMISML2DAI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAITransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OPTIMISML2DAI.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAITransfer)
				if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OPTIMISML2DAI *OPTIMISML2DAIFilterer) ParseTransfer(log types.Log) (*OPTIMISML2DAITransfer, error) {
	event := new(OPTIMISML2DAITransfer)
	if err := _OPTIMISML2DAI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
