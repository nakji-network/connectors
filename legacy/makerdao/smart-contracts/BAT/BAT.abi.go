// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BAT

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

// BATABI is the input ABI used to generate the binding from.
const BATABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"batFundDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"batFund\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenExchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethFundDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createTokens\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ethFundDeposit\",\"type\":\"address\"},{\"name\":\"_batFundDeposit\",\"type\":\"address\"},{\"name\":\"_fundingStartBlock\",\"type\":\"uint256\"},{\"name\":\"_fundingEndBlock\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"CreateBAT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BAT is an auto generated Go binding around an Ethereum contract.
type BAT struct {
	BATCaller     // Read-only binding to the contract
	BATTransactor // Write-only binding to the contract
	BATFilterer   // Log filterer for contract events
}

// BATCaller is an auto generated read-only Go binding around an Ethereum contract.
type BATCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BATTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BATTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BATFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BATFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BATSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BATSession struct {
	Contract     *BAT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BATCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BATCallerSession struct {
	Contract *BATCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BATTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BATTransactorSession struct {
	Contract     *BATTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BATRaw is an auto generated low-level Go binding around an Ethereum contract.
type BATRaw struct {
	Contract *BAT // Generic contract binding to access the raw methods on
}

// BATCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BATCallerRaw struct {
	Contract *BATCaller // Generic read-only contract binding to access the raw methods on
}

// BATTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BATTransactorRaw struct {
	Contract *BATTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBAT creates a new instance of BAT, bound to a specific deployed contract.
func NewBAT(address common.Address, backend bind.ContractBackend) (*BAT, error) {
	contract, err := bindBAT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BAT{BATCaller: BATCaller{contract: contract}, BATTransactor: BATTransactor{contract: contract}, BATFilterer: BATFilterer{contract: contract}}, nil
}

// NewBATCaller creates a new read-only instance of BAT, bound to a specific deployed contract.
func NewBATCaller(address common.Address, caller bind.ContractCaller) (*BATCaller, error) {
	contract, err := bindBAT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BATCaller{contract: contract}, nil
}

// NewBATTransactor creates a new write-only instance of BAT, bound to a specific deployed contract.
func NewBATTransactor(address common.Address, transactor bind.ContractTransactor) (*BATTransactor, error) {
	contract, err := bindBAT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BATTransactor{contract: contract}, nil
}

// NewBATFilterer creates a new log filterer instance of BAT, bound to a specific deployed contract.
func NewBATFilterer(address common.Address, filterer bind.ContractFilterer) (*BATFilterer, error) {
	contract, err := bindBAT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BATFilterer{contract: contract}, nil
}

// bindBAT binds a generic wrapper to an already deployed contract.
func bindBAT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BATABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BAT *BATRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BAT.Contract.BATCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BAT *BATRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAT.Contract.BATTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BAT *BATRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BAT.Contract.BATTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BAT *BATCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BAT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BAT *BATTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BAT *BATTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BAT.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_BAT *BATCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_BAT *BATSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BAT.Contract.Allowance(&_BAT.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_BAT *BATCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BAT.Contract.Allowance(&_BAT.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_BAT *BATCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_BAT *BATSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BAT.Contract.BalanceOf(&_BAT.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_BAT *BATCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BAT.Contract.BalanceOf(&_BAT.CallOpts, _owner)
}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_BAT *BATCaller) BatFund(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "batFund")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_BAT *BATSession) BatFund() (*big.Int, error) {
	return _BAT.Contract.BatFund(&_BAT.CallOpts)
}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_BAT *BATCallerSession) BatFund() (*big.Int, error) {
	return _BAT.Contract.BatFund(&_BAT.CallOpts)
}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_BAT *BATCaller) BatFundDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "batFundDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_BAT *BATSession) BatFundDeposit() (common.Address, error) {
	return _BAT.Contract.BatFundDeposit(&_BAT.CallOpts)
}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_BAT *BATCallerSession) BatFundDeposit() (common.Address, error) {
	return _BAT.Contract.BatFundDeposit(&_BAT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_BAT *BATCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_BAT *BATSession) Decimals() (*big.Int, error) {
	return _BAT.Contract.Decimals(&_BAT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_BAT *BATCallerSession) Decimals() (*big.Int, error) {
	return _BAT.Contract.Decimals(&_BAT.CallOpts)
}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_BAT *BATCaller) EthFundDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "ethFundDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_BAT *BATSession) EthFundDeposit() (common.Address, error) {
	return _BAT.Contract.EthFundDeposit(&_BAT.CallOpts)
}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_BAT *BATCallerSession) EthFundDeposit() (common.Address, error) {
	return _BAT.Contract.EthFundDeposit(&_BAT.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_BAT *BATCaller) FundingEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "fundingEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_BAT *BATSession) FundingEndBlock() (*big.Int, error) {
	return _BAT.Contract.FundingEndBlock(&_BAT.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_BAT *BATCallerSession) FundingEndBlock() (*big.Int, error) {
	return _BAT.Contract.FundingEndBlock(&_BAT.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_BAT *BATCaller) FundingStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "fundingStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_BAT *BATSession) FundingStartBlock() (*big.Int, error) {
	return _BAT.Contract.FundingStartBlock(&_BAT.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_BAT *BATCallerSession) FundingStartBlock() (*big.Int, error) {
	return _BAT.Contract.FundingStartBlock(&_BAT.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_BAT *BATCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_BAT *BATSession) IsFinalized() (bool, error) {
	return _BAT.Contract.IsFinalized(&_BAT.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_BAT *BATCallerSession) IsFinalized() (bool, error) {
	return _BAT.Contract.IsFinalized(&_BAT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_BAT *BATCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_BAT *BATSession) Name() (string, error) {
	return _BAT.Contract.Name(&_BAT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_BAT *BATCallerSession) Name() (string, error) {
	return _BAT.Contract.Name(&_BAT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_BAT *BATCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_BAT *BATSession) Symbol() (string, error) {
	return _BAT.Contract.Symbol(&_BAT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_BAT *BATCallerSession) Symbol() (string, error) {
	return _BAT.Contract.Symbol(&_BAT.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_BAT *BATCaller) TokenCreationCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "tokenCreationCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_BAT *BATSession) TokenCreationCap() (*big.Int, error) {
	return _BAT.Contract.TokenCreationCap(&_BAT.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_BAT *BATCallerSession) TokenCreationCap() (*big.Int, error) {
	return _BAT.Contract.TokenCreationCap(&_BAT.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_BAT *BATCaller) TokenCreationMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "tokenCreationMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_BAT *BATSession) TokenCreationMin() (*big.Int, error) {
	return _BAT.Contract.TokenCreationMin(&_BAT.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_BAT *BATCallerSession) TokenCreationMin() (*big.Int, error) {
	return _BAT.Contract.TokenCreationMin(&_BAT.CallOpts)
}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_BAT *BATCaller) TokenExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "tokenExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_BAT *BATSession) TokenExchangeRate() (*big.Int, error) {
	return _BAT.Contract.TokenExchangeRate(&_BAT.CallOpts)
}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_BAT *BATCallerSession) TokenExchangeRate() (*big.Int, error) {
	return _BAT.Contract.TokenExchangeRate(&_BAT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_BAT *BATCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_BAT *BATSession) TotalSupply() (*big.Int, error) {
	return _BAT.Contract.TotalSupply(&_BAT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_BAT *BATCallerSession) TotalSupply() (*big.Int, error) {
	return _BAT.Contract.TotalSupply(&_BAT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_BAT *BATCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BAT.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_BAT *BATSession) Version() (string, error) {
	return _BAT.Contract.Version(&_BAT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_BAT *BATCallerSession) Version() (string, error) {
	return _BAT.Contract.Version(&_BAT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BAT *BATTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BAT *BATSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.Approve(&_BAT.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BAT *BATTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.Approve(&_BAT.TransactOpts, _spender, _value)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_BAT *BATTransactor) CreateTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "createTokens")
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_BAT *BATSession) CreateTokens() (*types.Transaction, error) {
	return _BAT.Contract.CreateTokens(&_BAT.TransactOpts)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_BAT *BATTransactorSession) CreateTokens() (*types.Transaction, error) {
	return _BAT.Contract.CreateTokens(&_BAT.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BAT *BATTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BAT *BATSession) Finalize() (*types.Transaction, error) {
	return _BAT.Contract.Finalize(&_BAT.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BAT *BATTransactorSession) Finalize() (*types.Transaction, error) {
	return _BAT.Contract.Finalize(&_BAT.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BAT *BATTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BAT *BATSession) Refund() (*types.Transaction, error) {
	return _BAT.Contract.Refund(&_BAT.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BAT *BATTransactorSession) Refund() (*types.Transaction, error) {
	return _BAT.Contract.Refund(&_BAT.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BAT *BATTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BAT *BATSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.Transfer(&_BAT.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BAT *BATTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.Transfer(&_BAT.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BAT *BATTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BAT *BATSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.TransferFrom(&_BAT.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BAT *BATTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BAT.Contract.TransferFrom(&_BAT.TransactOpts, _from, _to, _value)
}

// BATApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BAT contract.
type BATApprovalIterator struct {
	Event *BATApproval // Event containing the contract specifics and raw log

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
func (it *BATApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BATApproval)
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
		it.Event = new(BATApproval)
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
func (it *BATApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BATApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BATApproval represents a Approval event raised by the BAT contract.
type BATApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_BAT *BATFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*BATApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BAT.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &BATApprovalIterator{contract: _BAT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_BAT *BATFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BATApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BAT.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BATApproval)
				if err := _BAT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_BAT *BATFilterer) ParseApproval(log types.Log) (*BATApproval, error) {
	event := new(BATApproval)
	if err := _BAT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BATCreateBATIterator is returned from FilterCreateBAT and is used to iterate over the raw logs and unpacked data for CreateBAT events raised by the BAT contract.
type BATCreateBATIterator struct {
	Event *BATCreateBAT // Event containing the contract specifics and raw log

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
func (it *BATCreateBATIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BATCreateBAT)
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
		it.Event = new(BATCreateBAT)
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
func (it *BATCreateBATIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BATCreateBATIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BATCreateBAT represents a CreateBAT event raised by the BAT contract.
type BATCreateBAT struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateBAT is a free log retrieval operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) FilterCreateBAT(opts *bind.FilterOpts, _to []common.Address) (*BATCreateBATIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.FilterLogs(opts, "CreateBAT", _toRule)
	if err != nil {
		return nil, err
	}
	return &BATCreateBATIterator{contract: _BAT.contract, event: "CreateBAT", logs: logs, sub: sub}, nil
}

// WatchCreateBAT is a free log subscription operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) WatchCreateBAT(opts *bind.WatchOpts, sink chan<- *BATCreateBAT, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.WatchLogs(opts, "CreateBAT", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BATCreateBAT)
				if err := _BAT.contract.UnpackLog(event, "CreateBAT", log); err != nil {
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

// ParseCreateBAT is a log parse operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) ParseCreateBAT(log types.Log) (*BATCreateBAT, error) {
	event := new(BATCreateBAT)
	if err := _BAT.contract.UnpackLog(event, "CreateBAT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BATLogRefundIterator is returned from FilterLogRefund and is used to iterate over the raw logs and unpacked data for LogRefund events raised by the BAT contract.
type BATLogRefundIterator struct {
	Event *BATLogRefund // Event containing the contract specifics and raw log

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
func (it *BATLogRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BATLogRefund)
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
		it.Event = new(BATLogRefund)
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
func (it *BATLogRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BATLogRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BATLogRefund represents a LogRefund event raised by the BAT contract.
type BATLogRefund struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogRefund is a free log retrieval operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) FilterLogRefund(opts *bind.FilterOpts, _to []common.Address) (*BATLogRefundIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.FilterLogs(opts, "LogRefund", _toRule)
	if err != nil {
		return nil, err
	}
	return &BATLogRefundIterator{contract: _BAT.contract, event: "LogRefund", logs: logs, sub: sub}, nil
}

// WatchLogRefund is a free log subscription operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) WatchLogRefund(opts *bind.WatchOpts, sink chan<- *BATLogRefund, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.WatchLogs(opts, "LogRefund", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BATLogRefund)
				if err := _BAT.contract.UnpackLog(event, "LogRefund", log); err != nil {
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

// ParseLogRefund is a log parse operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_BAT *BATFilterer) ParseLogRefund(log types.Log) (*BATLogRefund, error) {
	event := new(BATLogRefund)
	if err := _BAT.contract.UnpackLog(event, "LogRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BATTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BAT contract.
type BATTransferIterator struct {
	Event *BATTransfer // Event containing the contract specifics and raw log

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
func (it *BATTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BATTransfer)
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
		it.Event = new(BATTransfer)
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
func (it *BATTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BATTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BATTransfer represents a Transfer event raised by the BAT contract.
type BATTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_BAT *BATFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*BATTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BATTransferIterator{contract: _BAT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_BAT *BATFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BATTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BAT.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BATTransfer)
				if err := _BAT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_BAT *BATFilterer) ParseTransfer(log types.Log) (*BATTransfer, error) {
	event := new(BATTransfer)
	if err := _BAT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
