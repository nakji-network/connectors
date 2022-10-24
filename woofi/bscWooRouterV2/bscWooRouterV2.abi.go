// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bscWooRouterV2

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

// BscWooRouterV2MetaData contains all meta data concerning the BscWooRouterV2 contract.
var BscWooRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"WooPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIWooRouterV2.SwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooRouterSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"externalSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"querySwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueNativeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realQuoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realBaseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooPool\",\"outputs\":[{\"internalType\":\"contractIWooPP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BscWooRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BscWooRouterV2MetaData.ABI instead.
var BscWooRouterV2ABI = BscWooRouterV2MetaData.ABI

// BscWooRouterV2 is an auto generated Go binding around an Ethereum contract.
type BscWooRouterV2 struct {
	BscWooRouterV2Caller     // Read-only binding to the contract
	BscWooRouterV2Transactor // Write-only binding to the contract
	BscWooRouterV2Filterer   // Log filterer for contract events
}

// BscWooRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BscWooRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BscWooRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscWooRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscWooRouterV2Session struct {
	Contract     *BscWooRouterV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscWooRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscWooRouterV2CallerSession struct {
	Contract *BscWooRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BscWooRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscWooRouterV2TransactorSession struct {
	Contract     *BscWooRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BscWooRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BscWooRouterV2Raw struct {
	Contract *BscWooRouterV2 // Generic contract binding to access the raw methods on
}

// BscWooRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscWooRouterV2CallerRaw struct {
	Contract *BscWooRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// BscWooRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscWooRouterV2TransactorRaw struct {
	Contract *BscWooRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBscWooRouterV2 creates a new instance of BscWooRouterV2, bound to a specific deployed contract.
func NewBscWooRouterV2(address common.Address, backend bind.ContractBackend) (*BscWooRouterV2, error) {
	contract, err := bindBscWooRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2{BscWooRouterV2Caller: BscWooRouterV2Caller{contract: contract}, BscWooRouterV2Transactor: BscWooRouterV2Transactor{contract: contract}, BscWooRouterV2Filterer: BscWooRouterV2Filterer{contract: contract}}, nil
}

// NewBscWooRouterV2Caller creates a new read-only instance of BscWooRouterV2, bound to a specific deployed contract.
func NewBscWooRouterV2Caller(address common.Address, caller bind.ContractCaller) (*BscWooRouterV2Caller, error) {
	contract, err := bindBscWooRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2Caller{contract: contract}, nil
}

// NewBscWooRouterV2Transactor creates a new write-only instance of BscWooRouterV2, bound to a specific deployed contract.
func NewBscWooRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BscWooRouterV2Transactor, error) {
	contract, err := bindBscWooRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2Transactor{contract: contract}, nil
}

// NewBscWooRouterV2Filterer creates a new log filterer instance of BscWooRouterV2, bound to a specific deployed contract.
func NewBscWooRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BscWooRouterV2Filterer, error) {
	contract, err := bindBscWooRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2Filterer{contract: contract}, nil
}

// bindBscWooRouterV2 binds a generic wrapper to an already deployed contract.
func bindBscWooRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscWooRouterV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWooRouterV2 *BscWooRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWooRouterV2.Contract.BscWooRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWooRouterV2 *BscWooRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.BscWooRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWooRouterV2 *BscWooRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.BscWooRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWooRouterV2 *BscWooRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWooRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWooRouterV2 *BscWooRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWooRouterV2 *BscWooRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Session) WETH() (common.Address, error) {
	return _BscWooRouterV2.Contract.WETH(&_BscWooRouterV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) WETH() (common.Address, error) {
	return _BscWooRouterV2.Contract.WETH(&_BscWooRouterV2.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_BscWooRouterV2 *BscWooRouterV2Caller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_BscWooRouterV2 *BscWooRouterV2Session) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _BscWooRouterV2.Contract.IsWhitelisted(&_BscWooRouterV2.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _BscWooRouterV2.Contract.IsWhitelisted(&_BscWooRouterV2.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Session) Owner() (common.Address, error) {
	return _BscWooRouterV2.Contract.Owner(&_BscWooRouterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) Owner() (common.Address, error) {
	return _BscWooRouterV2.Contract.Owner(&_BscWooRouterV2.CallOpts)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2Caller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySellBase(&_BscWooRouterV2.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySellBase(&_BscWooRouterV2.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooRouterV2 *BscWooRouterV2Caller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySellQuote(&_BscWooRouterV2.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySellQuote(&_BscWooRouterV2.CallOpts, baseToken, quoteAmount)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_BscWooRouterV2 *BscWooRouterV2Caller) QuerySwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "querySwap", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySwap(&_BscWooRouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _BscWooRouterV2.Contract.QuerySwap(&_BscWooRouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Session) QuoteToken() (common.Address, error) {
	return _BscWooRouterV2.Contract.QuoteToken(&_BscWooRouterV2.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) QuoteToken() (common.Address, error) {
	return _BscWooRouterV2.Contract.QuoteToken(&_BscWooRouterV2.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Caller) WooPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooRouterV2.contract.Call(opts, &out, "wooPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2Session) WooPool() (common.Address, error) {
	return _BscWooRouterV2.Contract.WooPool(&_BscWooRouterV2.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_BscWooRouterV2 *BscWooRouterV2CallerSession) WooPool() (common.Address, error) {
	return _BscWooRouterV2.Contract.WooPool(&_BscWooRouterV2.CallOpts)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2Transactor) ExternalSwap(opts *bind.TransactOpts, approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "externalSwap", approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.ExternalSwap(&_BscWooRouterV2.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.ExternalSwap(&_BscWooRouterV2.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RenounceOwnership(&_BscWooRouterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RenounceOwnership(&_BscWooRouterV2.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RescueFunds(&_BscWooRouterV2.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RescueFunds(&_BscWooRouterV2.TransactOpts, token, amount)
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) RescueNativeFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "rescueNativeFunds")
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) RescueNativeFunds() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RescueNativeFunds(&_BscWooRouterV2.TransactOpts)
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) RescueNativeFunds() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.RescueNativeFunds(&_BscWooRouterV2.TransactOpts)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2Transactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SellBase(&_BscWooRouterV2.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SellBase(&_BscWooRouterV2.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_BscWooRouterV2 *BscWooRouterV2Transactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SellQuote(&_BscWooRouterV2.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SellQuote(&_BscWooRouterV2.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) SetPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "setPool", newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SetPool(&_BscWooRouterV2.TransactOpts, newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SetPool(&_BscWooRouterV2.TransactOpts, newPool)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) SetWhitelisted(opts *bind.TransactOpts, target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "setWhitelisted", target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SetWhitelisted(&_BscWooRouterV2.TransactOpts, target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.SetWhitelisted(&_BscWooRouterV2.TransactOpts, target, whitelisted)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2Transactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2Session) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.Swap(&_BscWooRouterV2.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.Swap(&_BscWooRouterV2.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.TransferOwnership(&_BscWooRouterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.TransferOwnership(&_BscWooRouterV2.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BscWooRouterV2 *BscWooRouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooRouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BscWooRouterV2 *BscWooRouterV2Session) Receive() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.Receive(&_BscWooRouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BscWooRouterV2 *BscWooRouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _BscWooRouterV2.Contract.Receive(&_BscWooRouterV2.TransactOpts)
}

// BscWooRouterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscWooRouterV2 contract.
type BscWooRouterV2OwnershipTransferredIterator struct {
	Event *BscWooRouterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscWooRouterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooRouterV2OwnershipTransferred)
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
		it.Event = new(BscWooRouterV2OwnershipTransferred)
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
func (it *BscWooRouterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooRouterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooRouterV2OwnershipTransferred represents a OwnershipTransferred event raised by the BscWooRouterV2 contract.
type BscWooRouterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscWooRouterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooRouterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2OwnershipTransferredIterator{contract: _BscWooRouterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscWooRouterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooRouterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooRouterV2OwnershipTransferred)
				if err := _BscWooRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscWooRouterV2 *BscWooRouterV2Filterer) ParseOwnershipTransferred(log types.Log) (*BscWooRouterV2OwnershipTransferred, error) {
	event := new(BscWooRouterV2OwnershipTransferred)
	if err := _BscWooRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooRouterV2WooPoolChangedIterator is returned from FilterWooPoolChanged and is used to iterate over the raw logs and unpacked data for WooPoolChanged events raised by the BscWooRouterV2 contract.
type BscWooRouterV2WooPoolChangedIterator struct {
	Event *BscWooRouterV2WooPoolChanged // Event containing the contract specifics and raw log

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
func (it *BscWooRouterV2WooPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooRouterV2WooPoolChanged)
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
		it.Event = new(BscWooRouterV2WooPoolChanged)
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
func (it *BscWooRouterV2WooPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooRouterV2WooPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooRouterV2WooPoolChanged represents a WooPoolChanged event raised by the BscWooRouterV2 contract.
type BscWooRouterV2WooPoolChanged struct {
	NewPool common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWooPoolChanged is a free log retrieval operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) FilterWooPoolChanged(opts *bind.FilterOpts) (*BscWooRouterV2WooPoolChangedIterator, error) {

	logs, sub, err := _BscWooRouterV2.contract.FilterLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2WooPoolChangedIterator{contract: _BscWooRouterV2.contract, event: "WooPoolChanged", logs: logs, sub: sub}, nil
}

// WatchWooPoolChanged is a free log subscription operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) WatchWooPoolChanged(opts *bind.WatchOpts, sink chan<- *BscWooRouterV2WooPoolChanged) (event.Subscription, error) {

	logs, sub, err := _BscWooRouterV2.contract.WatchLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooRouterV2WooPoolChanged)
				if err := _BscWooRouterV2.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
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

// ParseWooPoolChanged is a log parse operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) ParseWooPoolChanged(log types.Log) (*BscWooRouterV2WooPoolChanged, error) {
	event := new(BscWooRouterV2WooPoolChanged)
	if err := _BscWooRouterV2.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooRouterV2WooRouterSwapIterator is returned from FilterWooRouterSwap and is used to iterate over the raw logs and unpacked data for WooRouterSwap events raised by the BscWooRouterV2 contract.
type BscWooRouterV2WooRouterSwapIterator struct {
	Event *BscWooRouterV2WooRouterSwap // Event containing the contract specifics and raw log

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
func (it *BscWooRouterV2WooRouterSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooRouterV2WooRouterSwap)
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
		it.Event = new(BscWooRouterV2WooRouterSwap)
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
func (it *BscWooRouterV2WooRouterSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooRouterV2WooRouterSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooRouterV2WooRouterSwap represents a WooRouterSwap event raised by the BscWooRouterV2 contract.
type BscWooRouterV2WooRouterSwap struct {
	SwapType   uint8
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	From       common.Address
	To         common.Address
	RebateTo   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWooRouterSwap is a free log retrieval operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) FilterWooRouterSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*BscWooRouterV2WooRouterSwapIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWooRouterV2.contract.FilterLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BscWooRouterV2WooRouterSwapIterator{contract: _BscWooRouterV2.contract, event: "WooRouterSwap", logs: logs, sub: sub}, nil
}

// WatchWooRouterSwap is a free log subscription operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) WatchWooRouterSwap(opts *bind.WatchOpts, sink chan<- *BscWooRouterV2WooRouterSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWooRouterV2.contract.WatchLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooRouterV2WooRouterSwap)
				if err := _BscWooRouterV2.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
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

// ParseWooRouterSwap is a log parse operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_BscWooRouterV2 *BscWooRouterV2Filterer) ParseWooRouterSwap(log types.Log) (*BscWooRouterV2WooRouterSwap, error) {
	event := new(BscWooRouterV2WooRouterSwap)
	if err := _BscWooRouterV2.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
