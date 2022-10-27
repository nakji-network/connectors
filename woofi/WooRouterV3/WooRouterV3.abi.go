// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WooRouterV3

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

// WooRouterV3MetaData contains all meta data concerning the WooRouterV3 contract.
var WooRouterV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"WooPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIWooRouterV2.SwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooRouterSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"externalSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"querySwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueNativeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realQuoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realBaseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooPool\",\"outputs\":[{\"internalType\":\"contractIWooPP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WooRouterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use WooRouterV3MetaData.ABI instead.
var WooRouterV3ABI = WooRouterV3MetaData.ABI

// WooRouterV3 is an auto generated Go binding around an Ethereum contract.
type WooRouterV3 struct {
	WooRouterV3Caller     // Read-only binding to the contract
	WooRouterV3Transactor // Write-only binding to the contract
	WooRouterV3Filterer   // Log filterer for contract events
}

// WooRouterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type WooRouterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooRouterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WooRouterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooRouterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WooRouterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooRouterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WooRouterV3Session struct {
	Contract     *WooRouterV3      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WooRouterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WooRouterV3CallerSession struct {
	Contract *WooRouterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WooRouterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WooRouterV3TransactorSession struct {
	Contract     *WooRouterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WooRouterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type WooRouterV3Raw struct {
	Contract *WooRouterV3 // Generic contract binding to access the raw methods on
}

// WooRouterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WooRouterV3CallerRaw struct {
	Contract *WooRouterV3Caller // Generic read-only contract binding to access the raw methods on
}

// WooRouterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WooRouterV3TransactorRaw struct {
	Contract *WooRouterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWooRouterV3 creates a new instance of WooRouterV3, bound to a specific deployed contract.
func NewWooRouterV3(address common.Address, backend bind.ContractBackend) (*WooRouterV3, error) {
	contract, err := bindWooRouterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3{WooRouterV3Caller: WooRouterV3Caller{contract: contract}, WooRouterV3Transactor: WooRouterV3Transactor{contract: contract}, WooRouterV3Filterer: WooRouterV3Filterer{contract: contract}}, nil
}

// NewWooRouterV3Caller creates a new read-only instance of WooRouterV3, bound to a specific deployed contract.
func NewWooRouterV3Caller(address common.Address, caller bind.ContractCaller) (*WooRouterV3Caller, error) {
	contract, err := bindWooRouterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3Caller{contract: contract}, nil
}

// NewWooRouterV3Transactor creates a new write-only instance of WooRouterV3, bound to a specific deployed contract.
func NewWooRouterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*WooRouterV3Transactor, error) {
	contract, err := bindWooRouterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3Transactor{contract: contract}, nil
}

// NewWooRouterV3Filterer creates a new log filterer instance of WooRouterV3, bound to a specific deployed contract.
func NewWooRouterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*WooRouterV3Filterer, error) {
	contract, err := bindWooRouterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3Filterer{contract: contract}, nil
}

// bindWooRouterV3 binds a generic wrapper to an already deployed contract.
func bindWooRouterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WooRouterV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooRouterV3 *WooRouterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooRouterV3.Contract.WooRouterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooRouterV3 *WooRouterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooRouterV3.Contract.WooRouterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooRouterV3 *WooRouterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooRouterV3.Contract.WooRouterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooRouterV3 *WooRouterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooRouterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooRouterV3 *WooRouterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooRouterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooRouterV3 *WooRouterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooRouterV3.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooRouterV3 *WooRouterV3Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooRouterV3 *WooRouterV3Session) WETH() (common.Address, error) {
	return _WooRouterV3.Contract.WETH(&_WooRouterV3.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooRouterV3 *WooRouterV3CallerSession) WETH() (common.Address, error) {
	return _WooRouterV3.Contract.WETH(&_WooRouterV3.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_WooRouterV3 *WooRouterV3Caller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_WooRouterV3 *WooRouterV3Session) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _WooRouterV3.Contract.IsWhitelisted(&_WooRouterV3.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_WooRouterV3 *WooRouterV3CallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _WooRouterV3.Contract.IsWhitelisted(&_WooRouterV3.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooRouterV3 *WooRouterV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooRouterV3 *WooRouterV3Session) Owner() (common.Address, error) {
	return _WooRouterV3.Contract.Owner(&_WooRouterV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooRouterV3 *WooRouterV3CallerSession) Owner() (common.Address, error) {
	return _WooRouterV3.Contract.Owner(&_WooRouterV3.CallOpts)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooRouterV3 *WooRouterV3Caller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooRouterV3 *WooRouterV3Session) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySellBase(&_WooRouterV3.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooRouterV3 *WooRouterV3CallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySellBase(&_WooRouterV3.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooRouterV3 *WooRouterV3Caller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooRouterV3 *WooRouterV3Session) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySellQuote(&_WooRouterV3.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooRouterV3 *WooRouterV3CallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySellQuote(&_WooRouterV3.CallOpts, baseToken, quoteAmount)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooRouterV3 *WooRouterV3Caller) QuerySwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "querySwap", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooRouterV3 *WooRouterV3Session) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySwap(&_WooRouterV3.CallOpts, fromToken, toToken, fromAmount)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_WooRouterV3 *WooRouterV3CallerSession) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _WooRouterV3.Contract.QuerySwap(&_WooRouterV3.CallOpts, fromToken, toToken, fromAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooRouterV3 *WooRouterV3Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooRouterV3 *WooRouterV3Session) QuoteToken() (common.Address, error) {
	return _WooRouterV3.Contract.QuoteToken(&_WooRouterV3.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooRouterV3 *WooRouterV3CallerSession) QuoteToken() (common.Address, error) {
	return _WooRouterV3.Contract.QuoteToken(&_WooRouterV3.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooRouterV3 *WooRouterV3Caller) WooPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooRouterV3.contract.Call(opts, &out, "wooPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooRouterV3 *WooRouterV3Session) WooPool() (common.Address, error) {
	return _WooRouterV3.Contract.WooPool(&_WooRouterV3.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooRouterV3 *WooRouterV3CallerSession) WooPool() (common.Address, error) {
	return _WooRouterV3.Contract.WooPool(&_WooRouterV3.CallOpts)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3Transactor) ExternalSwap(opts *bind.TransactOpts, approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "externalSwap", approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3Session) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _WooRouterV3.Contract.ExternalSwap(&_WooRouterV3.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3TransactorSession) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _WooRouterV3.Contract.ExternalSwap(&_WooRouterV3.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooRouterV3 *WooRouterV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooRouterV3 *WooRouterV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _WooRouterV3.Contract.RenounceOwnership(&_WooRouterV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WooRouterV3.Contract.RenounceOwnership(&_WooRouterV3.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_WooRouterV3 *WooRouterV3Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_WooRouterV3 *WooRouterV3Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooRouterV3.Contract.RescueFunds(&_WooRouterV3.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooRouterV3.Contract.RescueFunds(&_WooRouterV3.TransactOpts, token, amount)
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_WooRouterV3 *WooRouterV3Transactor) RescueNativeFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "rescueNativeFunds")
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_WooRouterV3 *WooRouterV3Session) RescueNativeFunds() (*types.Transaction, error) {
	return _WooRouterV3.Contract.RescueNativeFunds(&_WooRouterV3.TransactOpts)
}

// RescueNativeFunds is a paid mutator transaction binding the contract method 0x063c27f8.
//
// Solidity: function rescueNativeFunds() returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) RescueNativeFunds() (*types.Transaction, error) {
	return _WooRouterV3.Contract.RescueNativeFunds(&_WooRouterV3.TransactOpts)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_WooRouterV3 *WooRouterV3Transactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_WooRouterV3 *WooRouterV3Session) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SellBase(&_WooRouterV3.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 realQuoteAmount)
func (_WooRouterV3 *WooRouterV3TransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SellBase(&_WooRouterV3.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_WooRouterV3 *WooRouterV3Transactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_WooRouterV3 *WooRouterV3Session) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SellQuote(&_WooRouterV3.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 realBaseAmount)
func (_WooRouterV3 *WooRouterV3TransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SellQuote(&_WooRouterV3.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_WooRouterV3 *WooRouterV3Transactor) SetPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "setPool", newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_WooRouterV3 *WooRouterV3Session) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SetPool(&_WooRouterV3.TransactOpts, newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SetPool(&_WooRouterV3.TransactOpts, newPool)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_WooRouterV3 *WooRouterV3Transactor) SetWhitelisted(opts *bind.TransactOpts, target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "setWhitelisted", target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_WooRouterV3 *WooRouterV3Session) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SetWhitelisted(&_WooRouterV3.TransactOpts, target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _WooRouterV3.Contract.SetWhitelisted(&_WooRouterV3.TransactOpts, target, whitelisted)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3Transactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3Session) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.Swap(&_WooRouterV3.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_WooRouterV3 *WooRouterV3TransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.Swap(&_WooRouterV3.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooRouterV3 *WooRouterV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooRouterV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooRouterV3 *WooRouterV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.TransferOwnership(&_WooRouterV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooRouterV3.Contract.TransferOwnership(&_WooRouterV3.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooRouterV3 *WooRouterV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooRouterV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooRouterV3 *WooRouterV3Session) Receive() (*types.Transaction, error) {
	return _WooRouterV3.Contract.Receive(&_WooRouterV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooRouterV3 *WooRouterV3TransactorSession) Receive() (*types.Transaction, error) {
	return _WooRouterV3.Contract.Receive(&_WooRouterV3.TransactOpts)
}

// WooRouterV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WooRouterV3 contract.
type WooRouterV3OwnershipTransferredIterator struct {
	Event *WooRouterV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WooRouterV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooRouterV3OwnershipTransferred)
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
		it.Event = new(WooRouterV3OwnershipTransferred)
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
func (it *WooRouterV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooRouterV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooRouterV3OwnershipTransferred represents a OwnershipTransferred event raised by the WooRouterV3 contract.
type WooRouterV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooRouterV3 *WooRouterV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooRouterV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooRouterV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3OwnershipTransferredIterator{contract: _WooRouterV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooRouterV3 *WooRouterV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WooRouterV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooRouterV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooRouterV3OwnershipTransferred)
				if err := _WooRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WooRouterV3 *WooRouterV3Filterer) ParseOwnershipTransferred(log types.Log) (*WooRouterV3OwnershipTransferred, error) {
	event := new(WooRouterV3OwnershipTransferred)
	if err := _WooRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooRouterV3WooPoolChangedIterator is returned from FilterWooPoolChanged and is used to iterate over the raw logs and unpacked data for WooPoolChanged events raised by the WooRouterV3 contract.
type WooRouterV3WooPoolChangedIterator struct {
	Event *WooRouterV3WooPoolChanged // Event containing the contract specifics and raw log

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
func (it *WooRouterV3WooPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooRouterV3WooPoolChanged)
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
		it.Event = new(WooRouterV3WooPoolChanged)
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
func (it *WooRouterV3WooPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooRouterV3WooPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooRouterV3WooPoolChanged represents a WooPoolChanged event raised by the WooRouterV3 contract.
type WooRouterV3WooPoolChanged struct {
	NewPool common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWooPoolChanged is a free log retrieval operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_WooRouterV3 *WooRouterV3Filterer) FilterWooPoolChanged(opts *bind.FilterOpts) (*WooRouterV3WooPoolChangedIterator, error) {

	logs, sub, err := _WooRouterV3.contract.FilterLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return &WooRouterV3WooPoolChangedIterator{contract: _WooRouterV3.contract, event: "WooPoolChanged", logs: logs, sub: sub}, nil
}

// WatchWooPoolChanged is a free log subscription operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_WooRouterV3 *WooRouterV3Filterer) WatchWooPoolChanged(opts *bind.WatchOpts, sink chan<- *WooRouterV3WooPoolChanged) (event.Subscription, error) {

	logs, sub, err := _WooRouterV3.contract.WatchLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooRouterV3WooPoolChanged)
				if err := _WooRouterV3.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
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
func (_WooRouterV3 *WooRouterV3Filterer) ParseWooPoolChanged(log types.Log) (*WooRouterV3WooPoolChanged, error) {
	event := new(WooRouterV3WooPoolChanged)
	if err := _WooRouterV3.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooRouterV3WooRouterSwapIterator is returned from FilterWooRouterSwap and is used to iterate over the raw logs and unpacked data for WooRouterSwap events raised by the WooRouterV3 contract.
type WooRouterV3WooRouterSwapIterator struct {
	Event *WooRouterV3WooRouterSwap // Event containing the contract specifics and raw log

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
func (it *WooRouterV3WooRouterSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooRouterV3WooRouterSwap)
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
		it.Event = new(WooRouterV3WooRouterSwap)
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
func (it *WooRouterV3WooRouterSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooRouterV3WooRouterSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooRouterV3WooRouterSwap represents a WooRouterSwap event raised by the WooRouterV3 contract.
type WooRouterV3WooRouterSwap struct {
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
func (_WooRouterV3 *WooRouterV3Filterer) FilterWooRouterSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*WooRouterV3WooRouterSwapIterator, error) {

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

	logs, sub, err := _WooRouterV3.contract.FilterLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooRouterV3WooRouterSwapIterator{contract: _WooRouterV3.contract, event: "WooRouterSwap", logs: logs, sub: sub}, nil
}

// WatchWooRouterSwap is a free log subscription operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_WooRouterV3 *WooRouterV3Filterer) WatchWooRouterSwap(opts *bind.WatchOpts, sink chan<- *WooRouterV3WooRouterSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WooRouterV3.contract.WatchLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooRouterV3WooRouterSwap)
				if err := _WooRouterV3.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
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
func (_WooRouterV3 *WooRouterV3Filterer) ParseWooRouterSwap(log types.Log) (*WooRouterV3WooRouterSwap, error) {
	event := new(WooRouterV3WooRouterSwap)
	if err := _WooRouterV3.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
