// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InsuranceFund

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

// Decimaldecimal is an auto generated low-level Go binding around an user-defined struct.
type Decimaldecimal struct {
	D *big.Int
}

// InsuranceFundABI is the input ABI used to generate the binding from.
const InsuranceFundABI = "[{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShutdownAllAmms\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemoved\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"withdrawer\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addAmm\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"candidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIExchangeWrapper\"}],\"name\":\"exchange\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"contractIAmm[]\"}],\"name\":\"getAllAmms\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getQuoteTokenLength\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIInflationMonitor\"}],\"name\":\"inflationMonitor\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isExistedAmm\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIMinter\"}],\"name\":\"minter\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIERC20\"}],\"name\":\"perpToken\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIERC20\"}],\"name\":\"quoteTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeAmm\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"contractIERC20\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setBeneficiary\",\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"_exchange\",\"internalType\":\"contractIExchangeWrapper\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setInflationMonitor\",\"inputs\":[{\"type\":\"address\",\"name\":\"_inflationMonitor\",\"internalType\":\"contractIInflationMonitor\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setMinter\",\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\",\"internalType\":\"contractIMinter\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"shutdownAllAmm\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateOwner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"_quoteToken\",\"internalType\":\"contractIERC20\"},{\"type\":\"tuple\",\"name\":\"_amount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]}]"

// InsuranceFund is an auto generated Go binding around an Ethereum contract.
type InsuranceFund struct {
	InsuranceFundCaller     // Read-only binding to the contract
	InsuranceFundTransactor // Write-only binding to the contract
	InsuranceFundFilterer   // Log filterer for contract events
}

// InsuranceFundCaller is an auto generated read-only Go binding around an Ethereum contract.
type InsuranceFundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceFundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InsuranceFundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceFundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InsuranceFundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceFundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InsuranceFundSession struct {
	Contract     *InsuranceFund    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InsuranceFundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InsuranceFundCallerSession struct {
	Contract *InsuranceFundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InsuranceFundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InsuranceFundTransactorSession struct {
	Contract     *InsuranceFundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InsuranceFundRaw is an auto generated low-level Go binding around an Ethereum contract.
type InsuranceFundRaw struct {
	Contract *InsuranceFund // Generic contract binding to access the raw methods on
}

// InsuranceFundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InsuranceFundCallerRaw struct {
	Contract *InsuranceFundCaller // Generic read-only contract binding to access the raw methods on
}

// InsuranceFundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InsuranceFundTransactorRaw struct {
	Contract *InsuranceFundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsuranceFund creates a new instance of InsuranceFund, bound to a specific deployed contract.
func NewInsuranceFund(address common.Address, backend bind.ContractBackend) (*InsuranceFund, error) {
	contract, err := bindInsuranceFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InsuranceFund{InsuranceFundCaller: InsuranceFundCaller{contract: contract}, InsuranceFundTransactor: InsuranceFundTransactor{contract: contract}, InsuranceFundFilterer: InsuranceFundFilterer{contract: contract}}, nil
}

// NewInsuranceFundCaller creates a new read-only instance of InsuranceFund, bound to a specific deployed contract.
func NewInsuranceFundCaller(address common.Address, caller bind.ContractCaller) (*InsuranceFundCaller, error) {
	contract, err := bindInsuranceFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InsuranceFundCaller{contract: contract}, nil
}

// NewInsuranceFundTransactor creates a new write-only instance of InsuranceFund, bound to a specific deployed contract.
func NewInsuranceFundTransactor(address common.Address, transactor bind.ContractTransactor) (*InsuranceFundTransactor, error) {
	contract, err := bindInsuranceFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InsuranceFundTransactor{contract: contract}, nil
}

// NewInsuranceFundFilterer creates a new log filterer instance of InsuranceFund, bound to a specific deployed contract.
func NewInsuranceFundFilterer(address common.Address, filterer bind.ContractFilterer) (*InsuranceFundFilterer, error) {
	contract, err := bindInsuranceFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InsuranceFundFilterer{contract: contract}, nil
}

// bindInsuranceFund binds a generic wrapper to an already deployed contract.
func bindInsuranceFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsuranceFundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsuranceFund *InsuranceFundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InsuranceFund.Contract.InsuranceFundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsuranceFund *InsuranceFundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.Contract.InsuranceFundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsuranceFund *InsuranceFundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsuranceFund.Contract.InsuranceFundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsuranceFund *InsuranceFundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InsuranceFund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsuranceFund *InsuranceFundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsuranceFund *InsuranceFundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsuranceFund.Contract.contract.Transact(opts, method, params...)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_InsuranceFund *InsuranceFundSession) Candidate() (common.Address, error) {
	return _InsuranceFund.Contract.Candidate(&_InsuranceFund.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) Candidate() (common.Address, error) {
	return _InsuranceFund.Contract.Candidate(&_InsuranceFund.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_InsuranceFund *InsuranceFundSession) Exchange() (common.Address, error) {
	return _InsuranceFund.Contract.Exchange(&_InsuranceFund.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) Exchange() (common.Address, error) {
	return _InsuranceFund.Contract.Exchange(&_InsuranceFund.CallOpts)
}

// GetAllAmms is a free data retrieval call binding the contract method 0xb7538c62.
//
// Solidity: function getAllAmms() view returns(address[])
func (_InsuranceFund *InsuranceFundCaller) GetAllAmms(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "getAllAmms")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAmms is a free data retrieval call binding the contract method 0xb7538c62.
//
// Solidity: function getAllAmms() view returns(address[])
func (_InsuranceFund *InsuranceFundSession) GetAllAmms() ([]common.Address, error) {
	return _InsuranceFund.Contract.GetAllAmms(&_InsuranceFund.CallOpts)
}

// GetAllAmms is a free data retrieval call binding the contract method 0xb7538c62.
//
// Solidity: function getAllAmms() view returns(address[])
func (_InsuranceFund *InsuranceFundCallerSession) GetAllAmms() ([]common.Address, error) {
	return _InsuranceFund.Contract.GetAllAmms(&_InsuranceFund.CallOpts)
}

// GetQuoteTokenLength is a free data retrieval call binding the contract method 0xbeb0597a.
//
// Solidity: function getQuoteTokenLength() view returns(uint256)
func (_InsuranceFund *InsuranceFundCaller) GetQuoteTokenLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "getQuoteTokenLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteTokenLength is a free data retrieval call binding the contract method 0xbeb0597a.
//
// Solidity: function getQuoteTokenLength() view returns(uint256)
func (_InsuranceFund *InsuranceFundSession) GetQuoteTokenLength() (*big.Int, error) {
	return _InsuranceFund.Contract.GetQuoteTokenLength(&_InsuranceFund.CallOpts)
}

// GetQuoteTokenLength is a free data retrieval call binding the contract method 0xbeb0597a.
//
// Solidity: function getQuoteTokenLength() view returns(uint256)
func (_InsuranceFund *InsuranceFundCallerSession) GetQuoteTokenLength() (*big.Int, error) {
	return _InsuranceFund.Contract.GetQuoteTokenLength(&_InsuranceFund.CallOpts)
}

// InflationMonitor is a free data retrieval call binding the contract method 0xca58b198.
//
// Solidity: function inflationMonitor() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) InflationMonitor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "inflationMonitor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InflationMonitor is a free data retrieval call binding the contract method 0xca58b198.
//
// Solidity: function inflationMonitor() view returns(address)
func (_InsuranceFund *InsuranceFundSession) InflationMonitor() (common.Address, error) {
	return _InsuranceFund.Contract.InflationMonitor(&_InsuranceFund.CallOpts)
}

// InflationMonitor is a free data retrieval call binding the contract method 0xca58b198.
//
// Solidity: function inflationMonitor() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) InflationMonitor() (common.Address, error) {
	return _InsuranceFund.Contract.InflationMonitor(&_InsuranceFund.CallOpts)
}

// IsExistedAmm is a free data retrieval call binding the contract method 0x0774a784.
//
// Solidity: function isExistedAmm(address _amm) view returns(bool)
func (_InsuranceFund *InsuranceFundCaller) IsExistedAmm(opts *bind.CallOpts, _amm common.Address) (bool, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "isExistedAmm", _amm)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistedAmm is a free data retrieval call binding the contract method 0x0774a784.
//
// Solidity: function isExistedAmm(address _amm) view returns(bool)
func (_InsuranceFund *InsuranceFundSession) IsExistedAmm(_amm common.Address) (bool, error) {
	return _InsuranceFund.Contract.IsExistedAmm(&_InsuranceFund.CallOpts, _amm)
}

// IsExistedAmm is a free data retrieval call binding the contract method 0x0774a784.
//
// Solidity: function isExistedAmm(address _amm) view returns(bool)
func (_InsuranceFund *InsuranceFundCallerSession) IsExistedAmm(_amm common.Address) (bool, error) {
	return _InsuranceFund.Contract.IsExistedAmm(&_InsuranceFund.CallOpts, _amm)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_InsuranceFund *InsuranceFundSession) Minter() (common.Address, error) {
	return _InsuranceFund.Contract.Minter(&_InsuranceFund.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) Minter() (common.Address, error) {
	return _InsuranceFund.Contract.Minter(&_InsuranceFund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceFund *InsuranceFundSession) Owner() (common.Address, error) {
	return _InsuranceFund.Contract.Owner(&_InsuranceFund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) Owner() (common.Address, error) {
	return _InsuranceFund.Contract.Owner(&_InsuranceFund.CallOpts)
}

// PerpToken is a free data retrieval call binding the contract method 0x9cadb3a0.
//
// Solidity: function perpToken() view returns(address)
func (_InsuranceFund *InsuranceFundCaller) PerpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "perpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerpToken is a free data retrieval call binding the contract method 0x9cadb3a0.
//
// Solidity: function perpToken() view returns(address)
func (_InsuranceFund *InsuranceFundSession) PerpToken() (common.Address, error) {
	return _InsuranceFund.Contract.PerpToken(&_InsuranceFund.CallOpts)
}

// PerpToken is a free data retrieval call binding the contract method 0x9cadb3a0.
//
// Solidity: function perpToken() view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) PerpToken() (common.Address, error) {
	return _InsuranceFund.Contract.PerpToken(&_InsuranceFund.CallOpts)
}

// QuoteTokens is a free data retrieval call binding the contract method 0x2c9115c7.
//
// Solidity: function quoteTokens(uint256 ) view returns(address)
func (_InsuranceFund *InsuranceFundCaller) QuoteTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InsuranceFund.contract.Call(opts, &out, "quoteTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteTokens is a free data retrieval call binding the contract method 0x2c9115c7.
//
// Solidity: function quoteTokens(uint256 ) view returns(address)
func (_InsuranceFund *InsuranceFundSession) QuoteTokens(arg0 *big.Int) (common.Address, error) {
	return _InsuranceFund.Contract.QuoteTokens(&_InsuranceFund.CallOpts, arg0)
}

// QuoteTokens is a free data retrieval call binding the contract method 0x2c9115c7.
//
// Solidity: function quoteTokens(uint256 ) view returns(address)
func (_InsuranceFund *InsuranceFundCallerSession) QuoteTokens(arg0 *big.Int) (common.Address, error) {
	return _InsuranceFund.Contract.QuoteTokens(&_InsuranceFund.CallOpts, arg0)
}

// AddAmm is a paid mutator transaction binding the contract method 0xd926de1c.
//
// Solidity: function addAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundTransactor) AddAmm(opts *bind.TransactOpts, _amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "addAmm", _amm)
}

// AddAmm is a paid mutator transaction binding the contract method 0xd926de1c.
//
// Solidity: function addAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundSession) AddAmm(_amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.AddAmm(&_InsuranceFund.TransactOpts, _amm)
}

// AddAmm is a paid mutator transaction binding the contract method 0xd926de1c.
//
// Solidity: function addAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) AddAmm(_amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.AddAmm(&_InsuranceFund.TransactOpts, _amm)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_InsuranceFund *InsuranceFundTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_InsuranceFund *InsuranceFundSession) Initialize() (*types.Transaction, error) {
	return _InsuranceFund.Contract.Initialize(&_InsuranceFund.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_InsuranceFund *InsuranceFundTransactorSession) Initialize() (*types.Transaction, error) {
	return _InsuranceFund.Contract.Initialize(&_InsuranceFund.TransactOpts)
}

// RemoveAmm is a paid mutator transaction binding the contract method 0xcf230979.
//
// Solidity: function removeAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundTransactor) RemoveAmm(opts *bind.TransactOpts, _amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "removeAmm", _amm)
}

// RemoveAmm is a paid mutator transaction binding the contract method 0xcf230979.
//
// Solidity: function removeAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundSession) RemoveAmm(_amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.RemoveAmm(&_InsuranceFund.TransactOpts, _amm)
}

// RemoveAmm is a paid mutator transaction binding the contract method 0xcf230979.
//
// Solidity: function removeAmm(address _amm) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) RemoveAmm(_amm common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.RemoveAmm(&_InsuranceFund.TransactOpts, _amm)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_InsuranceFund *InsuranceFundTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "removeToken", _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_InsuranceFund *InsuranceFundSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.RemoveToken(&_InsuranceFund.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.RemoveToken(&_InsuranceFund.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InsuranceFund *InsuranceFundTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InsuranceFund *InsuranceFundSession) RenounceOwnership() (*types.Transaction, error) {
	return _InsuranceFund.Contract.RenounceOwnership(&_InsuranceFund.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InsuranceFund *InsuranceFundTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InsuranceFund.Contract.RenounceOwnership(&_InsuranceFund.TransactOpts)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InsuranceFund *InsuranceFundTransactor) SetBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "setBeneficiary", _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InsuranceFund *InsuranceFundSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetBeneficiary(&_InsuranceFund.TransactOpts, _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetBeneficiary(&_InsuranceFund.TransactOpts, _beneficiary)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address _exchange) returns()
func (_InsuranceFund *InsuranceFundTransactor) SetExchange(opts *bind.TransactOpts, _exchange common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "setExchange", _exchange)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address _exchange) returns()
func (_InsuranceFund *InsuranceFundSession) SetExchange(_exchange common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetExchange(&_InsuranceFund.TransactOpts, _exchange)
}

// SetExchange is a paid mutator transaction binding the contract method 0x67b1f5df.
//
// Solidity: function setExchange(address _exchange) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) SetExchange(_exchange common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetExchange(&_InsuranceFund.TransactOpts, _exchange)
}

// SetInflationMonitor is a paid mutator transaction binding the contract method 0xc68ade37.
//
// Solidity: function setInflationMonitor(address _inflationMonitor) returns()
func (_InsuranceFund *InsuranceFundTransactor) SetInflationMonitor(opts *bind.TransactOpts, _inflationMonitor common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "setInflationMonitor", _inflationMonitor)
}

// SetInflationMonitor is a paid mutator transaction binding the contract method 0xc68ade37.
//
// Solidity: function setInflationMonitor(address _inflationMonitor) returns()
func (_InsuranceFund *InsuranceFundSession) SetInflationMonitor(_inflationMonitor common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetInflationMonitor(&_InsuranceFund.TransactOpts, _inflationMonitor)
}

// SetInflationMonitor is a paid mutator transaction binding the contract method 0xc68ade37.
//
// Solidity: function setInflationMonitor(address _inflationMonitor) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) SetInflationMonitor(_inflationMonitor common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetInflationMonitor(&_InsuranceFund.TransactOpts, _inflationMonitor)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_InsuranceFund *InsuranceFundTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_InsuranceFund *InsuranceFundSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetMinter(&_InsuranceFund.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetMinter(&_InsuranceFund.TransactOpts, _minter)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_InsuranceFund *InsuranceFundTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_InsuranceFund *InsuranceFundSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetOwner(&_InsuranceFund.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _InsuranceFund.Contract.SetOwner(&_InsuranceFund.TransactOpts, newOwner)
}

// ShutdownAllAmm is a paid mutator transaction binding the contract method 0x9dd96eb2.
//
// Solidity: function shutdownAllAmm() returns()
func (_InsuranceFund *InsuranceFundTransactor) ShutdownAllAmm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "shutdownAllAmm")
}

// ShutdownAllAmm is a paid mutator transaction binding the contract method 0x9dd96eb2.
//
// Solidity: function shutdownAllAmm() returns()
func (_InsuranceFund *InsuranceFundSession) ShutdownAllAmm() (*types.Transaction, error) {
	return _InsuranceFund.Contract.ShutdownAllAmm(&_InsuranceFund.TransactOpts)
}

// ShutdownAllAmm is a paid mutator transaction binding the contract method 0x9dd96eb2.
//
// Solidity: function shutdownAllAmm() returns()
func (_InsuranceFund *InsuranceFundTransactorSession) ShutdownAllAmm() (*types.Transaction, error) {
	return _InsuranceFund.Contract.ShutdownAllAmm(&_InsuranceFund.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_InsuranceFund *InsuranceFundTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_InsuranceFund *InsuranceFundSession) UpdateOwner() (*types.Transaction, error) {
	return _InsuranceFund.Contract.UpdateOwner(&_InsuranceFund.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_InsuranceFund *InsuranceFundTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _InsuranceFund.Contract.UpdateOwner(&_InsuranceFund.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x252978b8.
//
// Solidity: function withdraw(address _quoteToken, (uint256) _amount) returns()
func (_InsuranceFund *InsuranceFundTransactor) Withdraw(opts *bind.TransactOpts, _quoteToken common.Address, _amount Decimaldecimal) (*types.Transaction, error) {
	return _InsuranceFund.contract.Transact(opts, "withdraw", _quoteToken, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x252978b8.
//
// Solidity: function withdraw(address _quoteToken, (uint256) _amount) returns()
func (_InsuranceFund *InsuranceFundSession) Withdraw(_quoteToken common.Address, _amount Decimaldecimal) (*types.Transaction, error) {
	return _InsuranceFund.Contract.Withdraw(&_InsuranceFund.TransactOpts, _quoteToken, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x252978b8.
//
// Solidity: function withdraw(address _quoteToken, (uint256) _amount) returns()
func (_InsuranceFund *InsuranceFundTransactorSession) Withdraw(_quoteToken common.Address, _amount Decimaldecimal) (*types.Transaction, error) {
	return _InsuranceFund.Contract.Withdraw(&_InsuranceFund.TransactOpts, _quoteToken, _amount)
}

// InsuranceFundOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InsuranceFund contract.
type InsuranceFundOwnershipTransferredIterator struct {
	Event *InsuranceFundOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InsuranceFundOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceFundOwnershipTransferred)
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
		it.Event = new(InsuranceFundOwnershipTransferred)
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
func (it *InsuranceFundOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceFundOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceFundOwnershipTransferred represents a OwnershipTransferred event raised by the InsuranceFund contract.
type InsuranceFundOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InsuranceFund *InsuranceFundFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InsuranceFundOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InsuranceFund.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InsuranceFundOwnershipTransferredIterator{contract: _InsuranceFund.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InsuranceFund *InsuranceFundFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InsuranceFundOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InsuranceFund.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceFundOwnershipTransferred)
				if err := _InsuranceFund.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InsuranceFund *InsuranceFundFilterer) ParseOwnershipTransferred(log types.Log) (*InsuranceFundOwnershipTransferred, error) {
	event := new(InsuranceFundOwnershipTransferred)
	if err := _InsuranceFund.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InsuranceFundShutdownAllAmmsIterator is returned from FilterShutdownAllAmms and is used to iterate over the raw logs and unpacked data for ShutdownAllAmms events raised by the InsuranceFund contract.
type InsuranceFundShutdownAllAmmsIterator struct {
	Event *InsuranceFundShutdownAllAmms // Event containing the contract specifics and raw log

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
func (it *InsuranceFundShutdownAllAmmsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceFundShutdownAllAmms)
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
		it.Event = new(InsuranceFundShutdownAllAmms)
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
func (it *InsuranceFundShutdownAllAmmsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceFundShutdownAllAmmsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceFundShutdownAllAmms represents a ShutdownAllAmms event raised by the InsuranceFund contract.
type InsuranceFundShutdownAllAmms struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShutdownAllAmms is a free log retrieval operation binding the contract event 0xfc9f9c6cbb93f675af09b9bb43859333114dfa5e4c5abd35297e153f24348101.
//
// Solidity: event ShutdownAllAmms(uint256 blockNumber)
func (_InsuranceFund *InsuranceFundFilterer) FilterShutdownAllAmms(opts *bind.FilterOpts) (*InsuranceFundShutdownAllAmmsIterator, error) {

	logs, sub, err := _InsuranceFund.contract.FilterLogs(opts, "ShutdownAllAmms")
	if err != nil {
		return nil, err
	}
	return &InsuranceFundShutdownAllAmmsIterator{contract: _InsuranceFund.contract, event: "ShutdownAllAmms", logs: logs, sub: sub}, nil
}

// WatchShutdownAllAmms is a free log subscription operation binding the contract event 0xfc9f9c6cbb93f675af09b9bb43859333114dfa5e4c5abd35297e153f24348101.
//
// Solidity: event ShutdownAllAmms(uint256 blockNumber)
func (_InsuranceFund *InsuranceFundFilterer) WatchShutdownAllAmms(opts *bind.WatchOpts, sink chan<- *InsuranceFundShutdownAllAmms) (event.Subscription, error) {

	logs, sub, err := _InsuranceFund.contract.WatchLogs(opts, "ShutdownAllAmms")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceFundShutdownAllAmms)
				if err := _InsuranceFund.contract.UnpackLog(event, "ShutdownAllAmms", log); err != nil {
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

// ParseShutdownAllAmms is a log parse operation binding the contract event 0xfc9f9c6cbb93f675af09b9bb43859333114dfa5e4c5abd35297e153f24348101.
//
// Solidity: event ShutdownAllAmms(uint256 blockNumber)
func (_InsuranceFund *InsuranceFundFilterer) ParseShutdownAllAmms(log types.Log) (*InsuranceFundShutdownAllAmms, error) {
	event := new(InsuranceFundShutdownAllAmms)
	if err := _InsuranceFund.contract.UnpackLog(event, "ShutdownAllAmms", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InsuranceFundTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the InsuranceFund contract.
type InsuranceFundTokenAddedIterator struct {
	Event *InsuranceFundTokenAdded // Event containing the contract specifics and raw log

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
func (it *InsuranceFundTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceFundTokenAdded)
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
		it.Event = new(InsuranceFundTokenAdded)
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
func (it *InsuranceFundTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceFundTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceFundTokenAdded represents a TokenAdded event raised by the InsuranceFund contract.
type InsuranceFundTokenAdded struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*InsuranceFundTokenAddedIterator, error) {

	logs, sub, err := _InsuranceFund.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &InsuranceFundTokenAddedIterator{contract: _InsuranceFund.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *InsuranceFundTokenAdded) (event.Subscription, error) {

	logs, sub, err := _InsuranceFund.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceFundTokenAdded)
				if err := _InsuranceFund.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) ParseTokenAdded(log types.Log) (*InsuranceFundTokenAdded, error) {
	event := new(InsuranceFundTokenAdded)
	if err := _InsuranceFund.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InsuranceFundTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the InsuranceFund contract.
type InsuranceFundTokenRemovedIterator struct {
	Event *InsuranceFundTokenRemoved // Event containing the contract specifics and raw log

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
func (it *InsuranceFundTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceFundTokenRemoved)
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
		it.Event = new(InsuranceFundTokenRemoved)
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
func (it *InsuranceFundTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceFundTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceFundTokenRemoved represents a TokenRemoved event raised by the InsuranceFund contract.
type InsuranceFundTokenRemoved struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*InsuranceFundTokenRemovedIterator, error) {

	logs, sub, err := _InsuranceFund.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &InsuranceFundTokenRemovedIterator{contract: _InsuranceFund.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *InsuranceFundTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _InsuranceFund.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceFundTokenRemoved)
				if err := _InsuranceFund.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address tokenAddress)
func (_InsuranceFund *InsuranceFundFilterer) ParseTokenRemoved(log types.Log) (*InsuranceFundTokenRemoved, error) {
	event := new(InsuranceFundTokenRemoved)
	if err := _InsuranceFund.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InsuranceFundWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the InsuranceFund contract.
type InsuranceFundWithdrawnIterator struct {
	Event *InsuranceFundWithdrawn // Event containing the contract specifics and raw log

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
func (it *InsuranceFundWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceFundWithdrawn)
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
		it.Event = new(InsuranceFundWithdrawn)
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
func (it *InsuranceFundWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceFundWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceFundWithdrawn represents a Withdrawn event raised by the InsuranceFund contract.
type InsuranceFundWithdrawn struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address withdrawer, uint256 amount)
func (_InsuranceFund *InsuranceFundFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*InsuranceFundWithdrawnIterator, error) {

	logs, sub, err := _InsuranceFund.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &InsuranceFundWithdrawnIterator{contract: _InsuranceFund.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address withdrawer, uint256 amount)
func (_InsuranceFund *InsuranceFundFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *InsuranceFundWithdrawn) (event.Subscription, error) {

	logs, sub, err := _InsuranceFund.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceFundWithdrawn)
				if err := _InsuranceFund.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address withdrawer, uint256 amount)
func (_InsuranceFund *InsuranceFundFilterer) ParseWithdrawn(log types.Log) (*InsuranceFundWithdrawn, error) {
	event := new(InsuranceFundWithdrawn)
	if err := _InsuranceFund.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
