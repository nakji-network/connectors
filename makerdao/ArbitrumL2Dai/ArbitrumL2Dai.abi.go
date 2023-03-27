// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ArbitrumL2Dai

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

// ARBITRUML2DAIABI is the input ABI used to generate the binding from.
const ARBITRUML2DAIABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deploymentChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARBITRUML2DAI is an auto generated Go binding around an Ethereum contract.
type ARBITRUML2DAI struct {
	ARBITRUML2DAICaller     // Read-only binding to the contract
	ARBITRUML2DAITransactor // Write-only binding to the contract
	ARBITRUML2DAIFilterer   // Log filterer for contract events
}

// ARBITRUML2DAICaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUML2DAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUML2DAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUML2DAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2DAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUML2DAISession struct {
	Contract     *ARBITRUML2DAI    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARBITRUML2DAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUML2DAICallerSession struct {
	Contract *ARBITRUML2DAICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ARBITRUML2DAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUML2DAITransactorSession struct {
	Contract     *ARBITRUML2DAITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ARBITRUML2DAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUML2DAIRaw struct {
	Contract *ARBITRUML2DAI // Generic contract binding to access the raw methods on
}

// ARBITRUML2DAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUML2DAICallerRaw struct {
	Contract *ARBITRUML2DAICaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUML2DAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUML2DAITransactorRaw struct {
	Contract *ARBITRUML2DAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUML2DAI creates a new instance of ARBITRUML2DAI, bound to a specific deployed contract.
func NewARBITRUML2DAI(address common.Address, backend bind.ContractBackend) (*ARBITRUML2DAI, error) {
	contract, err := bindARBITRUML2DAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAI{ARBITRUML2DAICaller: ARBITRUML2DAICaller{contract: contract}, ARBITRUML2DAITransactor: ARBITRUML2DAITransactor{contract: contract}, ARBITRUML2DAIFilterer: ARBITRUML2DAIFilterer{contract: contract}}, nil
}

// NewARBITRUML2DAICaller creates a new read-only instance of ARBITRUML2DAI, bound to a specific deployed contract.
func NewARBITRUML2DAICaller(address common.Address, caller bind.ContractCaller) (*ARBITRUML2DAICaller, error) {
	contract, err := bindARBITRUML2DAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAICaller{contract: contract}, nil
}

// NewARBITRUML2DAITransactor creates a new write-only instance of ARBITRUML2DAI, bound to a specific deployed contract.
func NewARBITRUML2DAITransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUML2DAITransactor, error) {
	contract, err := bindARBITRUML2DAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAITransactor{contract: contract}, nil
}

// NewARBITRUML2DAIFilterer creates a new log filterer instance of ARBITRUML2DAI, bound to a specific deployed contract.
func NewARBITRUML2DAIFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUML2DAIFilterer, error) {
	contract, err := bindARBITRUML2DAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIFilterer{contract: contract}, nil
}

// bindARBITRUML2DAI binds a generic wrapper to an already deployed contract.
func bindARBITRUML2DAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUML2DAIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2DAI *ARBITRUML2DAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2DAI.Contract.ARBITRUML2DAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2DAI *ARBITRUML2DAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.ARBITRUML2DAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2DAI *ARBITRUML2DAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.ARBITRUML2DAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2DAI *ARBITRUML2DAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2DAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ARBITRUML2DAI.Contract.DOMAINSEPARATOR(&_ARBITRUML2DAI.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ARBITRUML2DAI.Contract.DOMAINSEPARATOR(&_ARBITRUML2DAI.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) PERMITTYPEHASH() ([32]byte, error) {
	return _ARBITRUML2DAI.Contract.PERMITTYPEHASH(&_ARBITRUML2DAI.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _ARBITRUML2DAI.Contract.PERMITTYPEHASH(&_ARBITRUML2DAI.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Allowance(&_ARBITRUML2DAI.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Allowance(&_ARBITRUML2DAI.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.BalanceOf(&_ARBITRUML2DAI.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.BalanceOf(&_ARBITRUML2DAI.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Decimals() (uint8, error) {
	return _ARBITRUML2DAI.Contract.Decimals(&_ARBITRUML2DAI.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Decimals() (uint8, error) {
	return _ARBITRUML2DAI.Contract.Decimals(&_ARBITRUML2DAI.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) DeploymentChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "deploymentChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) DeploymentChainId() (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.DeploymentChainId(&_ARBITRUML2DAI.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) DeploymentChainId() (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.DeploymentChainId(&_ARBITRUML2DAI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Name() (string, error) {
	return _ARBITRUML2DAI.Contract.Name(&_ARBITRUML2DAI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Name() (string, error) {
	return _ARBITRUML2DAI.Contract.Name(&_ARBITRUML2DAI.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Nonces(&_ARBITRUML2DAI.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Nonces(&_ARBITRUML2DAI.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Symbol() (string, error) {
	return _ARBITRUML2DAI.Contract.Symbol(&_ARBITRUML2DAI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Symbol() (string, error) {
	return _ARBITRUML2DAI.Contract.Symbol(&_ARBITRUML2DAI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) TotalSupply() (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.TotalSupply(&_ARBITRUML2DAI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) TotalSupply() (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.TotalSupply(&_ARBITRUML2DAI.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Version() (string, error) {
	return _ARBITRUML2DAI.Contract.Version(&_ARBITRUML2DAI.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Version() (string, error) {
	return _ARBITRUML2DAI.Contract.Version(&_ARBITRUML2DAI.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2DAI.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Wards(&_ARBITRUML2DAI.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2DAI *ARBITRUML2DAICallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2DAI.Contract.Wards(&_ARBITRUML2DAI.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Approve(&_ARBITRUML2DAI.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Approve(&_ARBITRUML2DAI.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Burn(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "burn", from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Burn(&_ARBITRUML2DAI.TransactOpts, from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Burn(&_ARBITRUML2DAI.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.DecreaseAllowance(&_ARBITRUML2DAI.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.DecreaseAllowance(&_ARBITRUML2DAI.TransactOpts, spender, subtractedValue)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Deny(&_ARBITRUML2DAI.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Deny(&_ARBITRUML2DAI.TransactOpts, usr)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.IncreaseAllowance(&_ARBITRUML2DAI.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.IncreaseAllowance(&_ARBITRUML2DAI.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Mint(&_ARBITRUML2DAI.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Mint(&_ARBITRUML2DAI.TransactOpts, to, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Permit(&_ARBITRUML2DAI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Permit(&_ARBITRUML2DAI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Rely(&_ARBITRUML2DAI.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Rely(&_ARBITRUML2DAI.TransactOpts, usr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Transfer(&_ARBITRUML2DAI.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.Transfer(&_ARBITRUML2DAI.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAISession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.TransferFrom(&_ARBITRUML2DAI.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ARBITRUML2DAI *ARBITRUML2DAITransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ARBITRUML2DAI.Contract.TransferFrom(&_ARBITRUML2DAI.TransactOpts, from, to, value)
}

// ARBITRUML2DAIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIApprovalIterator struct {
	Event *ARBITRUML2DAIApproval // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIApproval)
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
		it.Event = new(ARBITRUML2DAIApproval)
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
func (it *ARBITRUML2DAIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIApproval represents a Approval event raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ARBITRUML2DAIApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIApprovalIterator{contract: _ARBITRUML2DAI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIApproval)
				if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) ParseApproval(log types.Log) (*ARBITRUML2DAIApproval, error) {
	event := new(ARBITRUML2DAIApproval)
	if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2DAIDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIDenyIterator struct {
	Event *ARBITRUML2DAIDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIDeny)
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
		it.Event = new(ARBITRUML2DAIDeny)
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
func (it *ARBITRUML2DAIDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIDeny represents a Deny event raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2DAIDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIDenyIterator{contract: _ARBITRUML2DAI.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIDeny)
				if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) ParseDeny(log types.Log) (*ARBITRUML2DAIDeny, error) {
	event := new(ARBITRUML2DAIDeny)
	if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2DAIRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIRelyIterator struct {
	Event *ARBITRUML2DAIRely // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAIRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAIRely)
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
		it.Event = new(ARBITRUML2DAIRely)
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
func (it *ARBITRUML2DAIRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAIRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAIRely represents a Rely event raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAIRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2DAIRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAIRelyIterator{contract: _ARBITRUML2DAI.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAIRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAIRely)
				if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) ParseRely(log types.Log) (*ARBITRUML2DAIRely, error) {
	event := new(ARBITRUML2DAIRely)
	if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2DAITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAITransferIterator struct {
	Event *ARBITRUML2DAITransfer // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2DAITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2DAITransfer)
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
		it.Event = new(ARBITRUML2DAITransfer)
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
func (it *ARBITRUML2DAITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2DAITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2DAITransfer represents a Transfer event raised by the ARBITRUML2DAI contract.
type ARBITRUML2DAITransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ARBITRUML2DAITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2DAITransferIterator{contract: _ARBITRUML2DAI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ARBITRUML2DAITransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ARBITRUML2DAI.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2DAITransfer)
				if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ARBITRUML2DAI *ARBITRUML2DAIFilterer) ParseTransfer(log types.Log) (*ARBITRUML2DAITransfer, error) {
	event := new(ARBITRUML2DAITransfer)
	if err := _ARBITRUML2DAI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
