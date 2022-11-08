// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WooCrossChainRouterV1

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

// WooCrossChainRouterV1MetaData contains all meta data concerning the WooCrossChainRouterV1 contract.
var WooCrossChainRouterV1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"realToToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"name\":\"WooCrossSwapOnDstChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"realQuoteAmount\",\"type\":\"uint256\"}],\"name\":\"WooCrossSwapOnSrcChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeSlippage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcMinQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstMinToAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"crossSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dstGasForCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inCaseNativeTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stuckToken\",\"type\":\"address\"}],\"name\":\"inCaseTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wooPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stargateRouter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"refId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstMinToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"quoteLayerZeroFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"quotePoolIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bridgeSlippage\",\"type\":\"uint256\"}],\"name\":\"setBridgeSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstGasForCall\",\"type\":\"uint256\"}],\"name\":\"setDstGasForCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_quotePoolId\",\"type\":\"uint256\"}],\"name\":\"setQuotePoolId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stargateRouter\",\"type\":\"address\"}],\"name\":\"setStargateRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_wooCrossRouter\",\"type\":\"address\"}],\"name\":\"setWooCrossChainRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wooPool\",\"type\":\"address\"}],\"name\":\"setWooPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sgReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargateRouter\",\"outputs\":[{\"internalType\":\"contractIStargateRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"wooCrossRouters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooPool\",\"outputs\":[{\"internalType\":\"contractIWooPP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WooCrossChainRouterV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WooCrossChainRouterV1MetaData.ABI instead.
var WooCrossChainRouterV1ABI = WooCrossChainRouterV1MetaData.ABI

// WooCrossChainRouterV1 is an auto generated Go binding around an Ethereum contract.
type WooCrossChainRouterV1 struct {
	WooCrossChainRouterV1Caller     // Read-only binding to the contract
	WooCrossChainRouterV1Transactor // Write-only binding to the contract
	WooCrossChainRouterV1Filterer   // Log filterer for contract events
}

// WooCrossChainRouterV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type WooCrossChainRouterV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooCrossChainRouterV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WooCrossChainRouterV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooCrossChainRouterV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WooCrossChainRouterV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooCrossChainRouterV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WooCrossChainRouterV1Session struct {
	Contract     *WooCrossChainRouterV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WooCrossChainRouterV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WooCrossChainRouterV1CallerSession struct {
	Contract *WooCrossChainRouterV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WooCrossChainRouterV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WooCrossChainRouterV1TransactorSession struct {
	Contract     *WooCrossChainRouterV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WooCrossChainRouterV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type WooCrossChainRouterV1Raw struct {
	Contract *WooCrossChainRouterV1 // Generic contract binding to access the raw methods on
}

// WooCrossChainRouterV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WooCrossChainRouterV1CallerRaw struct {
	Contract *WooCrossChainRouterV1Caller // Generic read-only contract binding to access the raw methods on
}

// WooCrossChainRouterV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WooCrossChainRouterV1TransactorRaw struct {
	Contract *WooCrossChainRouterV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWooCrossChainRouterV1 creates a new instance of WooCrossChainRouterV1, bound to a specific deployed contract.
func NewWooCrossChainRouterV1(address common.Address, backend bind.ContractBackend) (*WooCrossChainRouterV1, error) {
	contract, err := bindWooCrossChainRouterV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1{WooCrossChainRouterV1Caller: WooCrossChainRouterV1Caller{contract: contract}, WooCrossChainRouterV1Transactor: WooCrossChainRouterV1Transactor{contract: contract}, WooCrossChainRouterV1Filterer: WooCrossChainRouterV1Filterer{contract: contract}}, nil
}

// NewWooCrossChainRouterV1Caller creates a new read-only instance of WooCrossChainRouterV1, bound to a specific deployed contract.
func NewWooCrossChainRouterV1Caller(address common.Address, caller bind.ContractCaller) (*WooCrossChainRouterV1Caller, error) {
	contract, err := bindWooCrossChainRouterV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1Caller{contract: contract}, nil
}

// NewWooCrossChainRouterV1Transactor creates a new write-only instance of WooCrossChainRouterV1, bound to a specific deployed contract.
func NewWooCrossChainRouterV1Transactor(address common.Address, transactor bind.ContractTransactor) (*WooCrossChainRouterV1Transactor, error) {
	contract, err := bindWooCrossChainRouterV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1Transactor{contract: contract}, nil
}

// NewWooCrossChainRouterV1Filterer creates a new log filterer instance of WooCrossChainRouterV1, bound to a specific deployed contract.
func NewWooCrossChainRouterV1Filterer(address common.Address, filterer bind.ContractFilterer) (*WooCrossChainRouterV1Filterer, error) {
	contract, err := bindWooCrossChainRouterV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1Filterer{contract: contract}, nil
}

// bindWooCrossChainRouterV1 binds a generic wrapper to an already deployed contract.
func bindWooCrossChainRouterV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WooCrossChainRouterV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooCrossChainRouterV1.Contract.WooCrossChainRouterV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.WooCrossChainRouterV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.WooCrossChainRouterV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooCrossChainRouterV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) WETH() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WETH(&_WooCrossChainRouterV1.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) WETH() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WETH(&_WooCrossChainRouterV1.CallOpts)
}

// BridgeSlippage is a free data retrieval call binding the contract method 0x83af5550.
//
// Solidity: function bridgeSlippage() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) BridgeSlippage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "bridgeSlippage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeSlippage is a free data retrieval call binding the contract method 0x83af5550.
//
// Solidity: function bridgeSlippage() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) BridgeSlippage() (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.BridgeSlippage(&_WooCrossChainRouterV1.CallOpts)
}

// BridgeSlippage is a free data retrieval call binding the contract method 0x83af5550.
//
// Solidity: function bridgeSlippage() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) BridgeSlippage() (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.BridgeSlippage(&_WooCrossChainRouterV1.CallOpts)
}

// DstGasForCall is a free data retrieval call binding the contract method 0x2bb8e169.
//
// Solidity: function dstGasForCall() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) DstGasForCall(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "dstGasForCall")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DstGasForCall is a free data retrieval call binding the contract method 0x2bb8e169.
//
// Solidity: function dstGasForCall() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) DstGasForCall() (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.DstGasForCall(&_WooCrossChainRouterV1.CallOpts)
}

// DstGasForCall is a free data retrieval call binding the contract method 0x2bb8e169.
//
// Solidity: function dstGasForCall() view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) DstGasForCall() (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.DstGasForCall(&_WooCrossChainRouterV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) Owner() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.Owner(&_WooCrossChainRouterV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) Owner() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.Owner(&_WooCrossChainRouterV1.CallOpts)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0xb8f1e0bc.
//
// Solidity: function quoteLayerZeroFee(uint16 dstChainId, address toToken, uint256 refId, uint256 dstMinToAmount, address to) view returns(uint256, uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) QuoteLayerZeroFee(opts *bind.CallOpts, dstChainId uint16, toToken common.Address, refId *big.Int, dstMinToAmount *big.Int, to common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "quoteLayerZeroFee", dstChainId, toToken, refId, dstMinToAmount, to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0xb8f1e0bc.
//
// Solidity: function quoteLayerZeroFee(uint16 dstChainId, address toToken, uint256 refId, uint256 dstMinToAmount, address to) view returns(uint256, uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) QuoteLayerZeroFee(dstChainId uint16, toToken common.Address, refId *big.Int, dstMinToAmount *big.Int, to common.Address) (*big.Int, *big.Int, error) {
	return _WooCrossChainRouterV1.Contract.QuoteLayerZeroFee(&_WooCrossChainRouterV1.CallOpts, dstChainId, toToken, refId, dstMinToAmount, to)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0xb8f1e0bc.
//
// Solidity: function quoteLayerZeroFee(uint16 dstChainId, address toToken, uint256 refId, uint256 dstMinToAmount, address to) view returns(uint256, uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) QuoteLayerZeroFee(dstChainId uint16, toToken common.Address, refId *big.Int, dstMinToAmount *big.Int, to common.Address) (*big.Int, *big.Int, error) {
	return _WooCrossChainRouterV1.Contract.QuoteLayerZeroFee(&_WooCrossChainRouterV1.CallOpts, dstChainId, toToken, refId, dstMinToAmount, to)
}

// QuotePoolIds is a free data retrieval call binding the contract method 0x2b51d7eb.
//
// Solidity: function quotePoolIds(uint16 ) view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) QuotePoolIds(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "quotePoolIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotePoolIds is a free data retrieval call binding the contract method 0x2b51d7eb.
//
// Solidity: function quotePoolIds(uint16 ) view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) QuotePoolIds(arg0 uint16) (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.QuotePoolIds(&_WooCrossChainRouterV1.CallOpts, arg0)
}

// QuotePoolIds is a free data retrieval call binding the contract method 0x2b51d7eb.
//
// Solidity: function quotePoolIds(uint16 ) view returns(uint256)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) QuotePoolIds(arg0 uint16) (*big.Int, error) {
	return _WooCrossChainRouterV1.Contract.QuotePoolIds(&_WooCrossChainRouterV1.CallOpts, arg0)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) QuoteToken() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.QuoteToken(&_WooCrossChainRouterV1.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) QuoteToken() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.QuoteToken(&_WooCrossChainRouterV1.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) StargateRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "stargateRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) StargateRouter() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.StargateRouter(&_WooCrossChainRouterV1.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) StargateRouter() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.StargateRouter(&_WooCrossChainRouterV1.CallOpts)
}

// WooCrossRouters is a free data retrieval call binding the contract method 0xf6af6957.
//
// Solidity: function wooCrossRouters(uint16 ) view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) WooCrossRouters(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "wooCrossRouters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooCrossRouters is a free data retrieval call binding the contract method 0xf6af6957.
//
// Solidity: function wooCrossRouters(uint16 ) view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) WooCrossRouters(arg0 uint16) (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WooCrossRouters(&_WooCrossChainRouterV1.CallOpts, arg0)
}

// WooCrossRouters is a free data retrieval call binding the contract method 0xf6af6957.
//
// Solidity: function wooCrossRouters(uint16 ) view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) WooCrossRouters(arg0 uint16) (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WooCrossRouters(&_WooCrossChainRouterV1.CallOpts, arg0)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Caller) WooPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooCrossChainRouterV1.contract.Call(opts, &out, "wooPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) WooPool() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WooPool(&_WooCrossChainRouterV1.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1CallerSession) WooPool() (common.Address, error) {
	return _WooCrossChainRouterV1.Contract.WooPool(&_WooCrossChainRouterV1.CallOpts)
}

// CrossSwap is a paid mutator transaction binding the contract method 0x4c904106.
//
// Solidity: function crossSwap(uint256 refId_, address fromToken, address toToken, uint256 fromAmount, uint256 srcMinQuoteAmount, uint256 dstMinToAmount, uint16 srcChainId, uint16 dstChainId, address to) payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) CrossSwap(opts *bind.TransactOpts, refId_ *big.Int, fromToken common.Address, toToken common.Address, fromAmount *big.Int, srcMinQuoteAmount *big.Int, dstMinToAmount *big.Int, srcChainId uint16, dstChainId uint16, to common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "crossSwap", refId_, fromToken, toToken, fromAmount, srcMinQuoteAmount, dstMinToAmount, srcChainId, dstChainId, to)
}

// CrossSwap is a paid mutator transaction binding the contract method 0x4c904106.
//
// Solidity: function crossSwap(uint256 refId_, address fromToken, address toToken, uint256 fromAmount, uint256 srcMinQuoteAmount, uint256 dstMinToAmount, uint16 srcChainId, uint16 dstChainId, address to) payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) CrossSwap(refId_ *big.Int, fromToken common.Address, toToken common.Address, fromAmount *big.Int, srcMinQuoteAmount *big.Int, dstMinToAmount *big.Int, srcChainId uint16, dstChainId uint16, to common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.CrossSwap(&_WooCrossChainRouterV1.TransactOpts, refId_, fromToken, toToken, fromAmount, srcMinQuoteAmount, dstMinToAmount, srcChainId, dstChainId, to)
}

// CrossSwap is a paid mutator transaction binding the contract method 0x4c904106.
//
// Solidity: function crossSwap(uint256 refId_, address fromToken, address toToken, uint256 fromAmount, uint256 srcMinQuoteAmount, uint256 dstMinToAmount, uint16 srcChainId, uint16 dstChainId, address to) payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) CrossSwap(refId_ *big.Int, fromToken common.Address, toToken common.Address, fromAmount *big.Int, srcMinQuoteAmount *big.Int, dstMinToAmount *big.Int, srcChainId uint16, dstChainId uint16, to common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.CrossSwap(&_WooCrossChainRouterV1.TransactOpts, refId_, fromToken, toToken, fromAmount, srcMinQuoteAmount, dstMinToAmount, srcChainId, dstChainId, to)
}

// InCaseNativeTokensGetStuck is a paid mutator transaction binding the contract method 0x85857419.
//
// Solidity: function inCaseNativeTokensGetStuck() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) InCaseNativeTokensGetStuck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "inCaseNativeTokensGetStuck")
}

// InCaseNativeTokensGetStuck is a paid mutator transaction binding the contract method 0x85857419.
//
// Solidity: function inCaseNativeTokensGetStuck() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) InCaseNativeTokensGetStuck() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.InCaseNativeTokensGetStuck(&_WooCrossChainRouterV1.TransactOpts)
}

// InCaseNativeTokensGetStuck is a paid mutator transaction binding the contract method 0x85857419.
//
// Solidity: function inCaseNativeTokensGetStuck() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) InCaseNativeTokensGetStuck() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.InCaseNativeTokensGetStuck(&_WooCrossChainRouterV1.TransactOpts)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address stuckToken) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) InCaseTokensGetStuck(opts *bind.TransactOpts, stuckToken common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "inCaseTokensGetStuck", stuckToken)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address stuckToken) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) InCaseTokensGetStuck(stuckToken common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.InCaseTokensGetStuck(&_WooCrossChainRouterV1.TransactOpts, stuckToken)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address stuckToken) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) InCaseTokensGetStuck(stuckToken common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.InCaseTokensGetStuck(&_WooCrossChainRouterV1.TransactOpts, stuckToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _weth, address _wooPool, address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) Initialize(opts *bind.TransactOpts, _weth common.Address, _wooPool common.Address, _stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "initialize", _weth, _wooPool, _stargateRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _weth, address _wooPool, address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) Initialize(_weth common.Address, _wooPool common.Address, _stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.Initialize(&_WooCrossChainRouterV1.TransactOpts, _weth, _wooPool, _stargateRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _weth, address _wooPool, address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) Initialize(_weth common.Address, _wooPool common.Address, _stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.Initialize(&_WooCrossChainRouterV1.TransactOpts, _weth, _wooPool, _stargateRouter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.RenounceOwnership(&_WooCrossChainRouterV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.RenounceOwnership(&_WooCrossChainRouterV1.TransactOpts)
}

// SetBridgeSlippage is a paid mutator transaction binding the contract method 0x403a01e5.
//
// Solidity: function setBridgeSlippage(uint256 _bridgeSlippage) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetBridgeSlippage(opts *bind.TransactOpts, _bridgeSlippage *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setBridgeSlippage", _bridgeSlippage)
}

// SetBridgeSlippage is a paid mutator transaction binding the contract method 0x403a01e5.
//
// Solidity: function setBridgeSlippage(uint256 _bridgeSlippage) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetBridgeSlippage(_bridgeSlippage *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetBridgeSlippage(&_WooCrossChainRouterV1.TransactOpts, _bridgeSlippage)
}

// SetBridgeSlippage is a paid mutator transaction binding the contract method 0x403a01e5.
//
// Solidity: function setBridgeSlippage(uint256 _bridgeSlippage) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetBridgeSlippage(_bridgeSlippage *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetBridgeSlippage(&_WooCrossChainRouterV1.TransactOpts, _bridgeSlippage)
}

// SetDstGasForCall is a paid mutator transaction binding the contract method 0x991f2188.
//
// Solidity: function setDstGasForCall(uint256 _dstGasForCall) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetDstGasForCall(opts *bind.TransactOpts, _dstGasForCall *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setDstGasForCall", _dstGasForCall)
}

// SetDstGasForCall is a paid mutator transaction binding the contract method 0x991f2188.
//
// Solidity: function setDstGasForCall(uint256 _dstGasForCall) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetDstGasForCall(_dstGasForCall *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetDstGasForCall(&_WooCrossChainRouterV1.TransactOpts, _dstGasForCall)
}

// SetDstGasForCall is a paid mutator transaction binding the contract method 0x991f2188.
//
// Solidity: function setDstGasForCall(uint256 _dstGasForCall) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetDstGasForCall(_dstGasForCall *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetDstGasForCall(&_WooCrossChainRouterV1.TransactOpts, _dstGasForCall)
}

// SetQuotePoolId is a paid mutator transaction binding the contract method 0x97724a86.
//
// Solidity: function setQuotePoolId(uint16 _chainId, uint256 _quotePoolId) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetQuotePoolId(opts *bind.TransactOpts, _chainId uint16, _quotePoolId *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setQuotePoolId", _chainId, _quotePoolId)
}

// SetQuotePoolId is a paid mutator transaction binding the contract method 0x97724a86.
//
// Solidity: function setQuotePoolId(uint16 _chainId, uint256 _quotePoolId) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetQuotePoolId(_chainId uint16, _quotePoolId *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetQuotePoolId(&_WooCrossChainRouterV1.TransactOpts, _chainId, _quotePoolId)
}

// SetQuotePoolId is a paid mutator transaction binding the contract method 0x97724a86.
//
// Solidity: function setQuotePoolId(uint16 _chainId, uint256 _quotePoolId) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetQuotePoolId(_chainId uint16, _quotePoolId *big.Int) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetQuotePoolId(&_WooCrossChainRouterV1.TransactOpts, _chainId, _quotePoolId)
}

// SetStargateRouter is a paid mutator transaction binding the contract method 0x51b78b47.
//
// Solidity: function setStargateRouter(address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetStargateRouter(opts *bind.TransactOpts, _stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setStargateRouter", _stargateRouter)
}

// SetStargateRouter is a paid mutator transaction binding the contract method 0x51b78b47.
//
// Solidity: function setStargateRouter(address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetStargateRouter(_stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetStargateRouter(&_WooCrossChainRouterV1.TransactOpts, _stargateRouter)
}

// SetStargateRouter is a paid mutator transaction binding the contract method 0x51b78b47.
//
// Solidity: function setStargateRouter(address _stargateRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetStargateRouter(_stargateRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetStargateRouter(&_WooCrossChainRouterV1.TransactOpts, _stargateRouter)
}

// SetWooCrossChainRouter is a paid mutator transaction binding the contract method 0xb5997a5b.
//
// Solidity: function setWooCrossChainRouter(uint16 _chainId, address _wooCrossRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetWooCrossChainRouter(opts *bind.TransactOpts, _chainId uint16, _wooCrossRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setWooCrossChainRouter", _chainId, _wooCrossRouter)
}

// SetWooCrossChainRouter is a paid mutator transaction binding the contract method 0xb5997a5b.
//
// Solidity: function setWooCrossChainRouter(uint16 _chainId, address _wooCrossRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetWooCrossChainRouter(_chainId uint16, _wooCrossRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetWooCrossChainRouter(&_WooCrossChainRouterV1.TransactOpts, _chainId, _wooCrossRouter)
}

// SetWooCrossChainRouter is a paid mutator transaction binding the contract method 0xb5997a5b.
//
// Solidity: function setWooCrossChainRouter(uint16 _chainId, address _wooCrossRouter) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetWooCrossChainRouter(_chainId uint16, _wooCrossRouter common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetWooCrossChainRouter(&_WooCrossChainRouterV1.TransactOpts, _chainId, _wooCrossRouter)
}

// SetWooPool is a paid mutator transaction binding the contract method 0xcbe127f2.
//
// Solidity: function setWooPool(address _wooPool) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SetWooPool(opts *bind.TransactOpts, _wooPool common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "setWooPool", _wooPool)
}

// SetWooPool is a paid mutator transaction binding the contract method 0xcbe127f2.
//
// Solidity: function setWooPool(address _wooPool) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SetWooPool(_wooPool common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetWooPool(&_WooCrossChainRouterV1.TransactOpts, _wooPool)
}

// SetWooPool is a paid mutator transaction binding the contract method 0xcbe127f2.
//
// Solidity: function setWooPool(address _wooPool) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SetWooPool(_wooPool common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SetWooPool(&_WooCrossChainRouterV1.TransactOpts, _wooPool)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 _chainId, bytes _srcAddress, uint256 _nonce, address _token, uint256 amountLD, bytes payload) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) SgReceive(opts *bind.TransactOpts, _chainId uint16, _srcAddress []byte, _nonce *big.Int, _token common.Address, amountLD *big.Int, payload []byte) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "sgReceive", _chainId, _srcAddress, _nonce, _token, amountLD, payload)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 _chainId, bytes _srcAddress, uint256 _nonce, address _token, uint256 amountLD, bytes payload) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) SgReceive(_chainId uint16, _srcAddress []byte, _nonce *big.Int, _token common.Address, amountLD *big.Int, payload []byte) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SgReceive(&_WooCrossChainRouterV1.TransactOpts, _chainId, _srcAddress, _nonce, _token, amountLD, payload)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 _chainId, bytes _srcAddress, uint256 _nonce, address _token, uint256 amountLD, bytes payload) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) SgReceive(_chainId uint16, _srcAddress []byte, _nonce *big.Int, _token common.Address, amountLD *big.Int, payload []byte) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.SgReceive(&_WooCrossChainRouterV1.TransactOpts, _chainId, _srcAddress, _nonce, _token, amountLD, payload)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.TransferOwnership(&_WooCrossChainRouterV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.TransferOwnership(&_WooCrossChainRouterV1.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooCrossChainRouterV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Session) Receive() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.Receive(&_WooCrossChainRouterV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1TransactorSession) Receive() (*types.Transaction, error) {
	return _WooCrossChainRouterV1.Contract.Receive(&_WooCrossChainRouterV1.TransactOpts)
}

// WooCrossChainRouterV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1OwnershipTransferredIterator struct {
	Event *WooCrossChainRouterV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WooCrossChainRouterV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooCrossChainRouterV1OwnershipTransferred)
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
		it.Event = new(WooCrossChainRouterV1OwnershipTransferred)
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
func (it *WooCrossChainRouterV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooCrossChainRouterV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooCrossChainRouterV1OwnershipTransferred represents a OwnershipTransferred event raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooCrossChainRouterV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1OwnershipTransferredIterator{contract: _WooCrossChainRouterV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WooCrossChainRouterV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooCrossChainRouterV1OwnershipTransferred)
				if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) ParseOwnershipTransferred(log types.Log) (*WooCrossChainRouterV1OwnershipTransferred, error) {
	event := new(WooCrossChainRouterV1OwnershipTransferred)
	if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooCrossChainRouterV1WooCrossSwapOnDstChainIterator is returned from FilterWooCrossSwapOnDstChain and is used to iterate over the raw logs and unpacked data for WooCrossSwapOnDstChain events raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1WooCrossSwapOnDstChainIterator struct {
	Event *WooCrossChainRouterV1WooCrossSwapOnDstChain // Event containing the contract specifics and raw log

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
func (it *WooCrossChainRouterV1WooCrossSwapOnDstChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
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
		it.Event = new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
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
func (it *WooCrossChainRouterV1WooCrossSwapOnDstChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooCrossChainRouterV1WooCrossSwapOnDstChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooCrossChainRouterV1WooCrossSwapOnDstChain represents a WooCrossSwapOnDstChain event raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1WooCrossSwapOnDstChain struct {
	RefId         *big.Int
	Sender        common.Address
	To            common.Address
	BridgedToken  common.Address
	BridgedAmount *big.Int
	ToToken       common.Address
	RealToToken   common.Address
	MinToAmount   *big.Int
	RealToAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWooCrossSwapOnDstChain is a free log retrieval operation binding the contract event 0x12ec1cd5a97a783f66bac513e496864dd4a0f398d181f887a6bb6df6bb9330fc.
//
// Solidity: event WooCrossSwapOnDstChain(uint256 indexed refId, address indexed sender, address indexed to, address bridgedToken, uint256 bridgedAmount, address toToken, address realToToken, uint256 minToAmount, uint256 realToAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) FilterWooCrossSwapOnDstChain(opts *bind.FilterOpts, refId []*big.Int, sender []common.Address, to []common.Address) (*WooCrossChainRouterV1WooCrossSwapOnDstChainIterator, error) {

	var refIdRule []interface{}
	for _, refIdItem := range refId {
		refIdRule = append(refIdRule, refIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.FilterLogs(opts, "WooCrossSwapOnDstChain", refIdRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1WooCrossSwapOnDstChainIterator{contract: _WooCrossChainRouterV1.contract, event: "WooCrossSwapOnDstChain", logs: logs, sub: sub}, nil
}

// WatchWooCrossSwapOnDstChain is a free log subscription operation binding the contract event 0x12ec1cd5a97a783f66bac513e496864dd4a0f398d181f887a6bb6df6bb9330fc.
//
// Solidity: event WooCrossSwapOnDstChain(uint256 indexed refId, address indexed sender, address indexed to, address bridgedToken, uint256 bridgedAmount, address toToken, address realToToken, uint256 minToAmount, uint256 realToAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) WatchWooCrossSwapOnDstChain(opts *bind.WatchOpts, sink chan<- *WooCrossChainRouterV1WooCrossSwapOnDstChain, refId []*big.Int, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var refIdRule []interface{}
	for _, refIdItem := range refId {
		refIdRule = append(refIdRule, refIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.WatchLogs(opts, "WooCrossSwapOnDstChain", refIdRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
				if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "WooCrossSwapOnDstChain", log); err != nil {
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

// ParseWooCrossSwapOnDstChain is a log parse operation binding the contract event 0x12ec1cd5a97a783f66bac513e496864dd4a0f398d181f887a6bb6df6bb9330fc.
//
// Solidity: event WooCrossSwapOnDstChain(uint256 indexed refId, address indexed sender, address indexed to, address bridgedToken, uint256 bridgedAmount, address toToken, address realToToken, uint256 minToAmount, uint256 realToAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) ParseWooCrossSwapOnDstChain(log types.Log) (*WooCrossChainRouterV1WooCrossSwapOnDstChain, error) {
	event := new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
	if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "WooCrossSwapOnDstChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator is returned from FilterWooCrossSwapOnSrcChain and is used to iterate over the raw logs and unpacked data for WooCrossSwapOnSrcChain events raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator struct {
	Event *WooCrossChainRouterV1WooCrossSwapOnSrcChain // Event containing the contract specifics and raw log

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
func (it *WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
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
		it.Event = new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
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
func (it *WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooCrossChainRouterV1WooCrossSwapOnSrcChain represents a WooCrossSwapOnSrcChain event raised by the WooCrossChainRouterV1 contract.
type WooCrossChainRouterV1WooCrossSwapOnSrcChain struct {
	RefId           *big.Int
	Sender          common.Address
	To              common.Address
	FromToken       common.Address
	FromAmount      *big.Int
	MinQuoteAmount  *big.Int
	RealQuoteAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWooCrossSwapOnSrcChain is a free log retrieval operation binding the contract event 0x84f8431fa975655da1378bee00f1e50b540c722eadd17490117d753a896a1611.
//
// Solidity: event WooCrossSwapOnSrcChain(uint256 indexed refId, address indexed sender, address indexed to, address fromToken, uint256 fromAmount, uint256 minQuoteAmount, uint256 realQuoteAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) FilterWooCrossSwapOnSrcChain(opts *bind.FilterOpts, refId []*big.Int, sender []common.Address, to []common.Address) (*WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator, error) {

	var refIdRule []interface{}
	for _, refIdItem := range refId {
		refIdRule = append(refIdRule, refIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.FilterLogs(opts, "WooCrossSwapOnSrcChain", refIdRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooCrossChainRouterV1WooCrossSwapOnSrcChainIterator{contract: _WooCrossChainRouterV1.contract, event: "WooCrossSwapOnSrcChain", logs: logs, sub: sub}, nil
}

// WatchWooCrossSwapOnSrcChain is a free log subscription operation binding the contract event 0x84f8431fa975655da1378bee00f1e50b540c722eadd17490117d753a896a1611.
//
// Solidity: event WooCrossSwapOnSrcChain(uint256 indexed refId, address indexed sender, address indexed to, address fromToken, uint256 fromAmount, uint256 minQuoteAmount, uint256 realQuoteAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) WatchWooCrossSwapOnSrcChain(opts *bind.WatchOpts, sink chan<- *WooCrossChainRouterV1WooCrossSwapOnSrcChain, refId []*big.Int, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var refIdRule []interface{}
	for _, refIdItem := range refId {
		refIdRule = append(refIdRule, refIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooCrossChainRouterV1.contract.WatchLogs(opts, "WooCrossSwapOnSrcChain", refIdRule, senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
				if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "WooCrossSwapOnSrcChain", log); err != nil {
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

// ParseWooCrossSwapOnSrcChain is a log parse operation binding the contract event 0x84f8431fa975655da1378bee00f1e50b540c722eadd17490117d753a896a1611.
//
// Solidity: event WooCrossSwapOnSrcChain(uint256 indexed refId, address indexed sender, address indexed to, address fromToken, uint256 fromAmount, uint256 minQuoteAmount, uint256 realQuoteAmount)
func (_WooCrossChainRouterV1 *WooCrossChainRouterV1Filterer) ParseWooCrossSwapOnSrcChain(log types.Log) (*WooCrossChainRouterV1WooCrossSwapOnSrcChain, error) {
	event := new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
	if err := _WooCrossChainRouterV1.contract.UnpackLog(event, "WooCrossSwapOnSrcChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
