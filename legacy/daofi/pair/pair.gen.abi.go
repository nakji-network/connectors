// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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

// IDAOfiV1PairABI is the input ABI used to generate the binding from.
const IDAOfiV1PairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountQuote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountQuote\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountQuote\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PLATFORM_FEE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountQuoteIn\",\"type\":\"uint256\"}],\"name\":\"getBaseOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBaseOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnerFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feesBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feesBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBaseIn\",\"type\":\"uint256\"}],\"name\":\"getQuoteOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountQuoteOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_formula\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_slopeNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_n\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_fee\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveRatio\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setPairOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slopeNumerator\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawPlatformFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDAOfiV1Pair is an auto generated Go binding around an Ethereum contract.
type IDAOfiV1Pair struct {
	IDAOfiV1PairCaller     // Read-only binding to the contract
	IDAOfiV1PairTransactor // Write-only binding to the contract
	IDAOfiV1PairFilterer   // Log filterer for contract events
}

// IDAOfiV1PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDAOfiV1PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDAOfiV1PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDAOfiV1PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDAOfiV1PairSession struct {
	Contract     *IDAOfiV1Pair     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAOfiV1PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDAOfiV1PairCallerSession struct {
	Contract *IDAOfiV1PairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IDAOfiV1PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDAOfiV1PairTransactorSession struct {
	Contract     *IDAOfiV1PairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDAOfiV1PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDAOfiV1PairRaw struct {
	Contract *IDAOfiV1Pair // Generic contract binding to access the raw methods on
}

// IDAOfiV1PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDAOfiV1PairCallerRaw struct {
	Contract *IDAOfiV1PairCaller // Generic read-only contract binding to access the raw methods on
}

// IDAOfiV1PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDAOfiV1PairTransactorRaw struct {
	Contract *IDAOfiV1PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDAOfiV1Pair creates a new instance of IDAOfiV1Pair, bound to a specific deployed contract.
func NewIDAOfiV1Pair(address common.Address, backend bind.ContractBackend) (*IDAOfiV1Pair, error) {
	contract, err := bindIDAOfiV1Pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1Pair{IDAOfiV1PairCaller: IDAOfiV1PairCaller{contract: contract}, IDAOfiV1PairTransactor: IDAOfiV1PairTransactor{contract: contract}, IDAOfiV1PairFilterer: IDAOfiV1PairFilterer{contract: contract}}, nil
}

// NewIDAOfiV1PairCaller creates a new read-only instance of IDAOfiV1Pair, bound to a specific deployed contract.
func NewIDAOfiV1PairCaller(address common.Address, caller bind.ContractCaller) (*IDAOfiV1PairCaller, error) {
	contract, err := bindIDAOfiV1Pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairCaller{contract: contract}, nil
}

// NewIDAOfiV1PairTransactor creates a new write-only instance of IDAOfiV1Pair, bound to a specific deployed contract.
func NewIDAOfiV1PairTransactor(address common.Address, transactor bind.ContractTransactor) (*IDAOfiV1PairTransactor, error) {
	contract, err := bindIDAOfiV1Pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairTransactor{contract: contract}, nil
}

// NewIDAOfiV1PairFilterer creates a new log filterer instance of IDAOfiV1Pair, bound to a specific deployed contract.
func NewIDAOfiV1PairFilterer(address common.Address, filterer bind.ContractFilterer) (*IDAOfiV1PairFilterer, error) {
	contract, err := bindIDAOfiV1Pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairFilterer{contract: contract}, nil
}

// bindIDAOfiV1Pair binds a generic wrapper to an already deployed contract.
func bindIDAOfiV1Pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDAOfiV1PairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOfiV1Pair *IDAOfiV1PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOfiV1Pair.Contract.IDAOfiV1PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOfiV1Pair *IDAOfiV1PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.IDAOfiV1PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOfiV1Pair *IDAOfiV1PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.IDAOfiV1PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOfiV1Pair *IDAOfiV1PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOfiV1Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.contract.Transact(opts, method, params...)
}

// PLATFORMFEE is a free data retrieval call binding the contract method 0x34fbc9a1.
//
// Solidity: function PLATFORM_FEE() view returns(uint8)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) PLATFORMFEE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "PLATFORM_FEE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PLATFORMFEE is a free data retrieval call binding the contract method 0x34fbc9a1.
//
// Solidity: function PLATFORM_FEE() view returns(uint8)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) PLATFORMFEE() (uint8, error) {
	return _IDAOfiV1Pair.Contract.PLATFORMFEE(&_IDAOfiV1Pair.CallOpts)
}

// PLATFORMFEE is a free data retrieval call binding the contract method 0x34fbc9a1.
//
// Solidity: function PLATFORM_FEE() view returns(uint8)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) PLATFORMFEE() (uint8, error) {
	return _IDAOfiV1Pair.Contract.PLATFORMFEE(&_IDAOfiV1Pair.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) BaseToken() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.BaseToken(&_IDAOfiV1Pair.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) BaseToken() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.BaseToken(&_IDAOfiV1Pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Factory() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.Factory(&_IDAOfiV1Pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) Factory() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.Factory(&_IDAOfiV1Pair.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) Fee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Fee() (uint32, error) {
	return _IDAOfiV1Pair.Contract.Fee(&_IDAOfiV1Pair.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) Fee() (uint32, error) {
	return _IDAOfiV1Pair.Contract.Fee(&_IDAOfiV1Pair.CallOpts)
}

// GetBaseOut is a free data retrieval call binding the contract method 0xe6d7059a.
//
// Solidity: function getBaseOut(uint256 amountQuoteIn) view returns(uint256 amountBaseOut)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) GetBaseOut(opts *bind.CallOpts, amountQuoteIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "getBaseOut", amountQuoteIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseOut is a free data retrieval call binding the contract method 0xe6d7059a.
//
// Solidity: function getBaseOut(uint256 amountQuoteIn) view returns(uint256 amountBaseOut)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) GetBaseOut(amountQuoteIn *big.Int) (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.GetBaseOut(&_IDAOfiV1Pair.CallOpts, amountQuoteIn)
}

// GetBaseOut is a free data retrieval call binding the contract method 0xe6d7059a.
//
// Solidity: function getBaseOut(uint256 amountQuoteIn) view returns(uint256 amountBaseOut)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) GetBaseOut(amountQuoteIn *big.Int) (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.GetBaseOut(&_IDAOfiV1Pair.CallOpts, amountQuoteIn)
}

// GetOwnerFees is a free data retrieval call binding the contract method 0x9dade07d.
//
// Solidity: function getOwnerFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) GetOwnerFees(opts *bind.CallOpts) (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "getOwnerFees")

	outstruct := new(struct {
		FeesBase  *big.Int
		FeesQuote *big.Int
	})

	outstruct.FeesBase = out[0].(*big.Int)
	outstruct.FeesQuote = out[1].(*big.Int)

	return *outstruct, err

}

// GetOwnerFees is a free data retrieval call binding the contract method 0x9dade07d.
//
// Solidity: function getOwnerFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) GetOwnerFees() (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetOwnerFees(&_IDAOfiV1Pair.CallOpts)
}

// GetOwnerFees is a free data retrieval call binding the contract method 0x9dade07d.
//
// Solidity: function getOwnerFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) GetOwnerFees() (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetOwnerFees(&_IDAOfiV1Pair.CallOpts)
}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) GetPlatformFees(opts *bind.CallOpts) (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "getPlatformFees")

	outstruct := new(struct {
		FeesBase  *big.Int
		FeesQuote *big.Int
	})

	outstruct.FeesBase = out[0].(*big.Int)
	outstruct.FeesQuote = out[1].(*big.Int)

	return *outstruct, err

}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) GetPlatformFees() (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetPlatformFees(&_IDAOfiV1Pair.CallOpts)
}

// GetPlatformFees is a free data retrieval call binding the contract method 0x55237200.
//
// Solidity: function getPlatformFees() view returns(uint256 feesBase, uint256 feesQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) GetPlatformFees() (struct {
	FeesBase  *big.Int
	FeesQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetPlatformFees(&_IDAOfiV1Pair.CallOpts)
}

// GetQuoteOut is a free data retrieval call binding the contract method 0x7b31ae64.
//
// Solidity: function getQuoteOut(uint256 amountBaseIn) view returns(uint256 amountQuoteOut)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) GetQuoteOut(opts *bind.CallOpts, amountBaseIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "getQuoteOut", amountBaseIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteOut is a free data retrieval call binding the contract method 0x7b31ae64.
//
// Solidity: function getQuoteOut(uint256 amountBaseIn) view returns(uint256 amountQuoteOut)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) GetQuoteOut(amountBaseIn *big.Int) (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.GetQuoteOut(&_IDAOfiV1Pair.CallOpts, amountBaseIn)
}

// GetQuoteOut is a free data retrieval call binding the contract method 0x7b31ae64.
//
// Solidity: function getQuoteOut(uint256 amountBaseIn) view returns(uint256 amountQuoteOut)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) GetQuoteOut(amountBaseIn *big.Int) (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.GetQuoteOut(&_IDAOfiV1Pair.CallOpts, amountBaseIn)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveBase, uint256 reserveQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		ReserveBase  *big.Int
		ReserveQuote *big.Int
	})

	outstruct.ReserveBase = out[0].(*big.Int)
	outstruct.ReserveQuote = out[1].(*big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveBase, uint256 reserveQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) GetReserves() (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetReserves(&_IDAOfiV1Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveBase, uint256 reserveQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) GetReserves() (struct {
	ReserveBase  *big.Int
	ReserveQuote *big.Int
}, error) {
	return _IDAOfiV1Pair.Contract.GetReserves(&_IDAOfiV1Pair.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) N(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) N() (uint32, error) {
	return _IDAOfiV1Pair.Contract.N(&_IDAOfiV1Pair.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) N() (uint32, error) {
	return _IDAOfiV1Pair.Contract.N(&_IDAOfiV1Pair.CallOpts)
}

// PairOwner is a free data retrieval call binding the contract method 0x2ec2fe3d.
//
// Solidity: function pairOwner() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) PairOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "pairOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairOwner is a free data retrieval call binding the contract method 0x2ec2fe3d.
//
// Solidity: function pairOwner() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) PairOwner() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.PairOwner(&_IDAOfiV1Pair.CallOpts)
}

// PairOwner is a free data retrieval call binding the contract method 0x2ec2fe3d.
//
// Solidity: function pairOwner() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) PairOwner() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.PairOwner(&_IDAOfiV1Pair.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Price() (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.Price(&_IDAOfiV1Pair.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) Price() (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.Price(&_IDAOfiV1Pair.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) QuoteToken() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.QuoteToken(&_IDAOfiV1Pair.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) QuoteToken() (common.Address, error) {
	return _IDAOfiV1Pair.Contract.QuoteToken(&_IDAOfiV1Pair.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) ReserveRatio(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "reserveRatio")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) ReserveRatio() (uint32, error) {
	return _IDAOfiV1Pair.Contract.ReserveRatio(&_IDAOfiV1Pair.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) ReserveRatio() (uint32, error) {
	return _IDAOfiV1Pair.Contract.ReserveRatio(&_IDAOfiV1Pair.CallOpts)
}

// SlopeNumerator is a free data retrieval call binding the contract method 0x21800fc3.
//
// Solidity: function slopeNumerator() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) SlopeNumerator(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "slopeNumerator")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlopeNumerator is a free data retrieval call binding the contract method 0x21800fc3.
//
// Solidity: function slopeNumerator() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) SlopeNumerator() (uint32, error) {
	return _IDAOfiV1Pair.Contract.SlopeNumerator(&_IDAOfiV1Pair.CallOpts)
}

// SlopeNumerator is a free data retrieval call binding the contract method 0x21800fc3.
//
// Solidity: function slopeNumerator() view returns(uint32)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) SlopeNumerator() (uint32, error) {
	return _IDAOfiV1Pair.Contract.SlopeNumerator(&_IDAOfiV1Pair.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDAOfiV1Pair.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Supply() (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.Supply(&_IDAOfiV1Pair.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_IDAOfiV1Pair *IDAOfiV1PairCallerSession) Supply() (*big.Int, error) {
	return _IDAOfiV1Pair.Contract.Supply(&_IDAOfiV1Pair.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address to) returns(uint256 amountBase)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) Deposit(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "deposit", to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address to) returns(uint256 amountBase)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Deposit(to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Deposit(&_IDAOfiV1Pair.TransactOpts, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address to) returns(uint256 amountBase)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) Deposit(to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Deposit(&_IDAOfiV1Pair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0xf940f446.
//
// Solidity: function initialize(address _formula, address _router, address _baseToken, address _quoteToken, address _pairOwner, uint32 _slopeNumerator, uint32 _n, uint32 _fee) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) Initialize(opts *bind.TransactOpts, _formula common.Address, _router common.Address, _baseToken common.Address, _quoteToken common.Address, _pairOwner common.Address, _slopeNumerator uint32, _n uint32, _fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "initialize", _formula, _router, _baseToken, _quoteToken, _pairOwner, _slopeNumerator, _n, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf940f446.
//
// Solidity: function initialize(address _formula, address _router, address _baseToken, address _quoteToken, address _pairOwner, uint32 _slopeNumerator, uint32 _n, uint32 _fee) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Initialize(_formula common.Address, _router common.Address, _baseToken common.Address, _quoteToken common.Address, _pairOwner common.Address, _slopeNumerator uint32, _n uint32, _fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Initialize(&_IDAOfiV1Pair.TransactOpts, _formula, _router, _baseToken, _quoteToken, _pairOwner, _slopeNumerator, _n, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf940f446.
//
// Solidity: function initialize(address _formula, address _router, address _baseToken, address _quoteToken, address _pairOwner, uint32 _slopeNumerator, uint32 _n, uint32 _fee) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) Initialize(_formula common.Address, _router common.Address, _baseToken common.Address, _quoteToken common.Address, _pairOwner common.Address, _slopeNumerator uint32, _n uint32, _fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Initialize(&_IDAOfiV1Pair.TransactOpts, _formula, _router, _baseToken, _quoteToken, _pairOwner, _slopeNumerator, _n, _fee)
}

// SetPairOwner is a paid mutator transaction binding the contract method 0x2dd2a921.
//
// Solidity: function setPairOwner(address ) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) SetPairOwner(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "setPairOwner", arg0)
}

// SetPairOwner is a paid mutator transaction binding the contract method 0x2dd2a921.
//
// Solidity: function setPairOwner(address ) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairSession) SetPairOwner(arg0 common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.SetPairOwner(&_IDAOfiV1Pair.TransactOpts, arg0)
}

// SetPairOwner is a paid mutator transaction binding the contract method 0x2dd2a921.
//
// Solidity: function setPairOwner(address ) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) SetPairOwner(arg0 common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.SetPairOwner(&_IDAOfiV1Pair.TransactOpts, arg0)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address to) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) Swap(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int, to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "swap", tokenIn, tokenOut, amountIn, amountOut, to)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address to) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Swap(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int, to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Swap(&_IDAOfiV1Pair.TransactOpts, tokenIn, tokenOut, amountIn, amountOut, to)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address to) returns()
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) Swap(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int, to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Swap(&_IDAOfiV1Pair.TransactOpts, tokenIn, tokenOut, amountIn, amountOut, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Withdraw(&_IDAOfiV1Pair.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.Withdraw(&_IDAOfiV1Pair.TransactOpts, to)
}

// WithdrawPlatformFees is a paid mutator transaction binding the contract method 0xd0b7830b.
//
// Solidity: function withdrawPlatformFees() returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactor) WithdrawPlatformFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOfiV1Pair.contract.Transact(opts, "withdrawPlatformFees")
}

// WithdrawPlatformFees is a paid mutator transaction binding the contract method 0xd0b7830b.
//
// Solidity: function withdrawPlatformFees() returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairSession) WithdrawPlatformFees() (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.WithdrawPlatformFees(&_IDAOfiV1Pair.TransactOpts)
}

// WithdrawPlatformFees is a paid mutator transaction binding the contract method 0xd0b7830b.
//
// Solidity: function withdrawPlatformFees() returns(uint256 amountBase, uint256 amountQuote)
func (_IDAOfiV1Pair *IDAOfiV1PairTransactorSession) WithdrawPlatformFees() (*types.Transaction, error) {
	return _IDAOfiV1Pair.Contract.WithdrawPlatformFees(&_IDAOfiV1Pair.TransactOpts)
}

// IDAOfiV1PairDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairDepositIterator struct {
	Event *IDAOfiV1PairDeposit // Event containing the contract specifics and raw log

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
func (it *IDAOfiV1PairDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOfiV1PairDeposit)
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
		it.Event = new(IDAOfiV1PairDeposit)
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
func (it *IDAOfiV1PairDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOfiV1PairDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOfiV1PairDeposit represents a Deposit event raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairDeposit struct {
	Sender      common.Address
	AmountBase  *big.Int
	AmountQuote *big.Int
	Output      *big.Int
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xede963da60ce3b657af455d654d08ff95549f13a4acb219b68756f68755e28e3.
//
// Solidity: event Deposit(address indexed sender, uint256 amountBase, uint256 amountQuote, uint256 output, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IDAOfiV1PairDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairDepositIterator{contract: _IDAOfiV1Pair.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xede963da60ce3b657af455d654d08ff95549f13a4acb219b68756f68755e28e3.
//
// Solidity: event Deposit(address indexed sender, uint256 amountBase, uint256 amountQuote, uint256 output, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IDAOfiV1PairDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOfiV1PairDeposit)
				if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xede963da60ce3b657af455d654d08ff95549f13a4acb219b68756f68755e28e3.
//
// Solidity: event Deposit(address indexed sender, uint256 amountBase, uint256 amountQuote, uint256 output, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) ParseDeposit(log types.Log) (*IDAOfiV1PairDeposit, error) {
	event := new(IDAOfiV1PairDeposit)
	if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDAOfiV1PairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairSwapIterator struct {
	Event *IDAOfiV1PairSwap // Event containing the contract specifics and raw log

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
func (it *IDAOfiV1PairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOfiV1PairSwap)
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
		it.Event = new(IDAOfiV1PairSwap)
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
func (it *IDAOfiV1PairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOfiV1PairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOfiV1PairSwap represents a Swap event raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairSwap struct {
	Pair      common.Address
	Sender    common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xaed33853adff9bbae920f24cf42e39d86728a74691e17fea513aa1cbc109c17a.
//
// Solidity: event Swap(address indexed pair, address indexed sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) FilterSwap(opts *bind.FilterOpts, pair []common.Address, sender []common.Address, to []common.Address) (*IDAOfiV1PairSwapIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.FilterLogs(opts, "Swap", pairRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairSwapIterator{contract: _IDAOfiV1Pair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xaed33853adff9bbae920f24cf42e39d86728a74691e17fea513aa1cbc109c17a.
//
// Solidity: event Swap(address indexed pair, address indexed sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IDAOfiV1PairSwap, pair []common.Address, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.WatchLogs(opts, "Swap", pairRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOfiV1PairSwap)
				if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xaed33853adff9bbae920f24cf42e39d86728a74691e17fea513aa1cbc109c17a.
//
// Solidity: event Swap(address indexed pair, address indexed sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) ParseSwap(log types.Log) (*IDAOfiV1PairSwap, error) {
	event := new(IDAOfiV1PairSwap)
	if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDAOfiV1PairWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairWithdrawIterator struct {
	Event *IDAOfiV1PairWithdraw // Event containing the contract specifics and raw log

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
func (it *IDAOfiV1PairWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOfiV1PairWithdraw)
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
		it.Event = new(IDAOfiV1PairWithdraw)
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
func (it *IDAOfiV1PairWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOfiV1PairWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOfiV1PairWithdraw represents a Withdraw event raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairWithdraw struct {
	Sender      common.Address
	AmountBase  *big.Int
	AmountQuote *big.Int
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IDAOfiV1PairWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairWithdrawIterator{contract: _IDAOfiV1Pair.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IDAOfiV1PairWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOfiV1PairWithdraw)
				if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) ParseWithdraw(log types.Log) (*IDAOfiV1PairWithdraw, error) {
	event := new(IDAOfiV1PairWithdraw)
	if err := _IDAOfiV1Pair.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDAOfiV1PairWithdrawFeesIterator is returned from FilterWithdrawFees and is used to iterate over the raw logs and unpacked data for WithdrawFees events raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairWithdrawFeesIterator struct {
	Event *IDAOfiV1PairWithdrawFees // Event containing the contract specifics and raw log

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
func (it *IDAOfiV1PairWithdrawFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOfiV1PairWithdrawFees)
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
		it.Event = new(IDAOfiV1PairWithdrawFees)
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
func (it *IDAOfiV1PairWithdrawFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOfiV1PairWithdrawFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOfiV1PairWithdrawFees represents a WithdrawFees event raised by the IDAOfiV1Pair contract.
type IDAOfiV1PairWithdrawFees struct {
	Sender      common.Address
	AmountBase  *big.Int
	AmountQuote *big.Int
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFees is a free log retrieval operation binding the contract event 0xfefe98d2f91f55619f0b737e514b21b1425ad73d0ed4684bf1b51547575fe884.
//
// Solidity: event WithdrawFees(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) FilterWithdrawFees(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IDAOfiV1PairWithdrawFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.FilterLogs(opts, "WithdrawFees", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1PairWithdrawFeesIterator{contract: _IDAOfiV1Pair.contract, event: "WithdrawFees", logs: logs, sub: sub}, nil
}

// WatchWithdrawFees is a free log subscription operation binding the contract event 0xfefe98d2f91f55619f0b737e514b21b1425ad73d0ed4684bf1b51547575fe884.
//
// Solidity: event WithdrawFees(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) WatchWithdrawFees(opts *bind.WatchOpts, sink chan<- *IDAOfiV1PairWithdrawFees, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IDAOfiV1Pair.contract.WatchLogs(opts, "WithdrawFees", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOfiV1PairWithdrawFees)
				if err := _IDAOfiV1Pair.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
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

// ParseWithdrawFees is a log parse operation binding the contract event 0xfefe98d2f91f55619f0b737e514b21b1425ad73d0ed4684bf1b51547575fe884.
//
// Solidity: event WithdrawFees(address indexed sender, uint256 amountBase, uint256 amountQuote, address indexed to)
func (_IDAOfiV1Pair *IDAOfiV1PairFilterer) ParseWithdrawFees(log types.Log) (*IDAOfiV1PairWithdrawFees, error) {
	event := new(IDAOfiV1PairWithdrawFees)
	if err := _IDAOfiV1Pair.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
		return nil, err
	}
	return event, nil
}
