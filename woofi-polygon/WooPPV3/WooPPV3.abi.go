// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WooPPV3

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

// WooPPV3MetaData contains all meta data concerning the WooPPV3 contract.
var WooPPV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newQuoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"ParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRewardManager\",\"type\":\"address\"}],\"name\":\"RewardManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"StrategistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"WooGuardianUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"WooracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractIWooFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStrategist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairsInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"removeBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPairsInfo\",\"type\":\"string\"}],\"name\":\"setPairsInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"setWooGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"setWooracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"threshold\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastResetTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"R\",\"type\":\"uint64\"},{\"internalType\":\"uint112\",\"name\":\"target\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAllToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooGuardian\",\"outputs\":[{\"internalType\":\"contractIWooGuardian\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WooPPV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use WooPPV3MetaData.ABI instead.
var WooPPV3ABI = WooPPV3MetaData.ABI

// WooPPV3 is an auto generated Go binding around an Ethereum contract.
type WooPPV3 struct {
	WooPPV3Caller     // Read-only binding to the contract
	WooPPV3Transactor // Write-only binding to the contract
	WooPPV3Filterer   // Log filterer for contract events
}

// WooPPV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type WooPPV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WooPPV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WooPPV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WooPPV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WooPPV3Session struct {
	Contract     *WooPPV3          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WooPPV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WooPPV3CallerSession struct {
	Contract *WooPPV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WooPPV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WooPPV3TransactorSession struct {
	Contract     *WooPPV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WooPPV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type WooPPV3Raw struct {
	Contract *WooPPV3 // Generic contract binding to access the raw methods on
}

// WooPPV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WooPPV3CallerRaw struct {
	Contract *WooPPV3Caller // Generic read-only contract binding to access the raw methods on
}

// WooPPV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WooPPV3TransactorRaw struct {
	Contract *WooPPV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWooPPV3 creates a new instance of WooPPV3, bound to a specific deployed contract.
func NewWooPPV3(address common.Address, backend bind.ContractBackend) (*WooPPV3, error) {
	contract, err := bindWooPPV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WooPPV3{WooPPV3Caller: WooPPV3Caller{contract: contract}, WooPPV3Transactor: WooPPV3Transactor{contract: contract}, WooPPV3Filterer: WooPPV3Filterer{contract: contract}}, nil
}

// NewWooPPV3Caller creates a new read-only instance of WooPPV3, bound to a specific deployed contract.
func NewWooPPV3Caller(address common.Address, caller bind.ContractCaller) (*WooPPV3Caller, error) {
	contract, err := bindWooPPV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WooPPV3Caller{contract: contract}, nil
}

// NewWooPPV3Transactor creates a new write-only instance of WooPPV3, bound to a specific deployed contract.
func NewWooPPV3Transactor(address common.Address, transactor bind.ContractTransactor) (*WooPPV3Transactor, error) {
	contract, err := bindWooPPV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WooPPV3Transactor{contract: contract}, nil
}

// NewWooPPV3Filterer creates a new log filterer instance of WooPPV3, bound to a specific deployed contract.
func NewWooPPV3Filterer(address common.Address, filterer bind.ContractFilterer) (*WooPPV3Filterer, error) {
	contract, err := bindWooPPV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WooPPV3Filterer{contract: contract}, nil
}

// bindWooPPV3 binds a generic wrapper to an already deployed contract.
func bindWooPPV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WooPPV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooPPV3 *WooPPV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooPPV3.Contract.WooPPV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooPPV3 *WooPPV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV3.Contract.WooPPV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooPPV3 *WooPPV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooPPV3.Contract.WooPPV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WooPPV3 *WooPPV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WooPPV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WooPPV3 *WooPPV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WooPPV3 *WooPPV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WooPPV3.Contract.contract.Transact(opts, method, params...)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3Caller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3Session) NEWOWNER() (common.Address, error) {
	return _WooPPV3.Contract.NEWOWNER(&_WooPPV3.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) NEWOWNER() (common.Address, error) {
	return _WooPPV3.Contract.NEWOWNER(&_WooPPV3.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3Caller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3Session) OWNER() (common.Address, error) {
	return _WooPPV3.Contract.OWNER(&_WooPPV3.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) OWNER() (common.Address, error) {
	return _WooPPV3.Contract.OWNER(&_WooPPV3.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_WooPPV3 *WooPPV3Caller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_WooPPV3 *WooPPV3Session) FeeManager() (common.Address, error) {
	return _WooPPV3.Contract.FeeManager(&_WooPPV3.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) FeeManager() (common.Address, error) {
	return _WooPPV3.Contract.FeeManager(&_WooPPV3.CallOpts)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_WooPPV3 *WooPPV3Caller) IsStrategist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "isStrategist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_WooPPV3 *WooPPV3Session) IsStrategist(arg0 common.Address) (bool, error) {
	return _WooPPV3.Contract.IsStrategist(&_WooPPV3.CallOpts, arg0)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_WooPPV3 *WooPPV3CallerSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _WooPPV3.Contract.IsStrategist(&_WooPPV3.CallOpts, arg0)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_WooPPV3 *WooPPV3Caller) PairsInfo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "pairsInfo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_WooPPV3 *WooPPV3Session) PairsInfo() (string, error) {
	return _WooPPV3.Contract.PairsInfo(&_WooPPV3.CallOpts)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_WooPPV3 *WooPPV3CallerSession) PairsInfo() (string, error) {
	return _WooPPV3.Contract.PairsInfo(&_WooPPV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV3 *WooPPV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV3 *WooPPV3Session) Paused() (bool, error) {
	return _WooPPV3.Contract.Paused(&_WooPPV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WooPPV3 *WooPPV3CallerSession) Paused() (bool, error) {
	return _WooPPV3.Contract.Paused(&_WooPPV3.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV3 *WooPPV3Caller) PoolSize(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "poolSize", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV3 *WooPPV3Session) PoolSize(token common.Address) (*big.Int, error) {
	return _WooPPV3.Contract.PoolSize(&_WooPPV3.CallOpts, token)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_WooPPV3 *WooPPV3CallerSession) PoolSize(token common.Address) (*big.Int, error) {
	return _WooPPV3.Contract.PoolSize(&_WooPPV3.CallOpts, token)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3Caller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3Session) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _WooPPV3.Contract.QuerySellBase(&_WooPPV3.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3CallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _WooPPV3.Contract.QuerySellBase(&_WooPPV3.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3Caller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3Session) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _WooPPV3.Contract.QuerySellQuote(&_WooPPV3.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3CallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _WooPPV3.Contract.QuerySellQuote(&_WooPPV3.CallOpts, baseToken, quoteAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV3 *WooPPV3Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV3 *WooPPV3Session) QuoteToken() (common.Address, error) {
	return _WooPPV3.Contract.QuoteToken(&_WooPPV3.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) QuoteToken() (common.Address, error) {
	return _WooPPV3.Contract.QuoteToken(&_WooPPV3.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_WooPPV3 *WooPPV3Caller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "tokenInfo", arg0)

	outstruct := new(struct {
		Reserve            *big.Int
		Threshold          *big.Int
		LastResetTimestamp uint32
		R                  uint64
		Target             *big.Int
		IsValid            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Threshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastResetTimestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.R = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Target = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_WooPPV3 *WooPPV3Session) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _WooPPV3.Contract.TokenInfo(&_WooPPV3.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_WooPPV3 *WooPPV3CallerSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _WooPPV3.Contract.TokenInfo(&_WooPPV3.CallOpts, arg0)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_WooPPV3 *WooPPV3Caller) WooGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "wooGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_WooPPV3 *WooPPV3Session) WooGuardian() (common.Address, error) {
	return _WooPPV3.Contract.WooGuardian(&_WooPPV3.CallOpts)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) WooGuardian() (common.Address, error) {
	return _WooPPV3.Contract.WooGuardian(&_WooPPV3.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV3 *WooPPV3Caller) Wooracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WooPPV3.contract.Call(opts, &out, "wooracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV3 *WooPPV3Session) Wooracle() (common.Address, error) {
	return _WooPPV3.Contract.Wooracle(&_WooPPV3.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_WooPPV3 *WooPPV3CallerSession) Wooracle() (common.Address, error) {
	return _WooPPV3.Contract.Wooracle(&_WooPPV3.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_WooPPV3 *WooPPV3Transactor) AddBaseToken(opts *bind.TransactOpts, baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "addBaseToken", baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_WooPPV3 *WooPPV3Session) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.AddBaseToken(&_WooPPV3.TransactOpts, baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_WooPPV3 *WooPPV3TransactorSession) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.AddBaseToken(&_WooPPV3.TransactOpts, baseToken, threshold, R)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_WooPPV3 *WooPPV3Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_WooPPV3 *WooPPV3Session) ClaimOwnership() (*types.Transaction, error) {
	return _WooPPV3.Contract.ClaimOwnership(&_WooPPV3.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_WooPPV3 *WooPPV3TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _WooPPV3.Contract.ClaimOwnership(&_WooPPV3.TransactOpts)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_WooPPV3 *WooPPV3Transactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_WooPPV3 *WooPPV3Session) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.InitOwner(&_WooPPV3.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_WooPPV3 *WooPPV3TransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.InitOwner(&_WooPPV3.TransactOpts, newOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV3 *WooPPV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV3 *WooPPV3Session) Pause() (*types.Transaction, error) {
	return _WooPPV3.Contract.Pause(&_WooPPV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WooPPV3 *WooPPV3TransactorSession) Pause() (*types.Transaction, error) {
	return _WooPPV3.Contract.Pause(&_WooPPV3.TransactOpts)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_WooPPV3 *WooPPV3Transactor) RemoveBaseToken(opts *bind.TransactOpts, baseToken common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "removeBaseToken", baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_WooPPV3 *WooPPV3Session) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.RemoveBaseToken(&_WooPPV3.TransactOpts, baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_WooPPV3 *WooPPV3TransactorSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.RemoveBaseToken(&_WooPPV3.TransactOpts, baseToken)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3Transactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3Session) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SellBase(&_WooPPV3.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_WooPPV3 *WooPPV3TransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SellBase(&_WooPPV3.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3Transactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3Session) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SellQuote(&_WooPPV3.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_WooPPV3 *WooPPV3TransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SellQuote(&_WooPPV3.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_WooPPV3 *WooPPV3Transactor) SetFeeManager(opts *bind.TransactOpts, newFeeManager common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "setFeeManager", newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_WooPPV3 *WooPPV3Session) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetFeeManager(&_WooPPV3.TransactOpts, newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_WooPPV3 *WooPPV3TransactorSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetFeeManager(&_WooPPV3.TransactOpts, newFeeManager)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_WooPPV3 *WooPPV3Transactor) SetPairsInfo(opts *bind.TransactOpts, newPairsInfo string) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "setPairsInfo", newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_WooPPV3 *WooPPV3Session) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetPairsInfo(&_WooPPV3.TransactOpts, newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_WooPPV3 *WooPPV3TransactorSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetPairsInfo(&_WooPPV3.TransactOpts, newPairsInfo)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_WooPPV3 *WooPPV3Transactor) SetStrategist(opts *bind.TransactOpts, strategist common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "setStrategist", strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_WooPPV3 *WooPPV3Session) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetStrategist(&_WooPPV3.TransactOpts, strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_WooPPV3 *WooPPV3TransactorSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetStrategist(&_WooPPV3.TransactOpts, strategist, flag)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_WooPPV3 *WooPPV3Transactor) SetWooGuardian(opts *bind.TransactOpts, newWooGuardian common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "setWooGuardian", newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_WooPPV3 *WooPPV3Session) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetWooGuardian(&_WooPPV3.TransactOpts, newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_WooPPV3 *WooPPV3TransactorSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetWooGuardian(&_WooPPV3.TransactOpts, newWooGuardian)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_WooPPV3 *WooPPV3Transactor) SetWooracle(opts *bind.TransactOpts, newWooracle common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "setWooracle", newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_WooPPV3 *WooPPV3Session) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetWooracle(&_WooPPV3.TransactOpts, newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_WooPPV3 *WooPPV3TransactorSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.SetWooracle(&_WooPPV3.TransactOpts, newWooracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV3 *WooPPV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV3 *WooPPV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.TransferOwnership(&_WooPPV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WooPPV3 *WooPPV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.TransferOwnership(&_WooPPV3.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_WooPPV3 *WooPPV3Transactor) TuneParameters(opts *bind.TransactOpts, token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "tuneParameters", token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_WooPPV3 *WooPPV3Session) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.TuneParameters(&_WooPPV3.TransactOpts, token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_WooPPV3 *WooPPV3TransactorSession) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.TuneParameters(&_WooPPV3.TransactOpts, token, newThreshold, newR)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV3 *WooPPV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV3 *WooPPV3Session) Unpause() (*types.Transaction, error) {
	return _WooPPV3.Contract.Unpause(&_WooPPV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WooPPV3 *WooPPV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _WooPPV3.Contract.Unpause(&_WooPPV3.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_WooPPV3 *WooPPV3Transactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_WooPPV3 *WooPPV3Session) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.Withdraw(&_WooPPV3.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_WooPPV3 *WooPPV3TransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WooPPV3.Contract.Withdraw(&_WooPPV3.TransactOpts, token, to, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_WooPPV3 *WooPPV3Transactor) WithdrawAll(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "withdrawAll", token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_WooPPV3 *WooPPV3Session) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.WithdrawAll(&_WooPPV3.TransactOpts, token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_WooPPV3 *WooPPV3TransactorSession) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.WithdrawAll(&_WooPPV3.TransactOpts, token, to)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_WooPPV3 *WooPPV3Transactor) WithdrawAllToOwner(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WooPPV3.contract.Transact(opts, "withdrawAllToOwner", token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_WooPPV3 *WooPPV3Session) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.WithdrawAllToOwner(&_WooPPV3.TransactOpts, token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_WooPPV3 *WooPPV3TransactorSession) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _WooPPV3.Contract.WithdrawAllToOwner(&_WooPPV3.TransactOpts, token)
}

// WooPPV3FeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the WooPPV3 contract.
type WooPPV3FeeManagerUpdatedIterator struct {
	Event *WooPPV3FeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3FeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3FeeManagerUpdated)
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
		it.Event = new(WooPPV3FeeManagerUpdated)
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
func (it *WooPPV3FeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3FeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3FeeManagerUpdated represents a FeeManagerUpdated event raised by the WooPPV3 contract.
type WooPPV3FeeManagerUpdated struct {
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_WooPPV3 *WooPPV3Filterer) FilterFeeManagerUpdated(opts *bind.FilterOpts, newFeeManager []common.Address) (*WooPPV3FeeManagerUpdatedIterator, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3FeeManagerUpdatedIterator{contract: _WooPPV3.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_WooPPV3 *WooPPV3Filterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3FeeManagerUpdated, newFeeManager []common.Address) (event.Subscription, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3FeeManagerUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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

// ParseFeeManagerUpdated is a log parse operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_WooPPV3 *WooPPV3Filterer) ParseFeeManagerUpdated(log types.Log) (*WooPPV3FeeManagerUpdated, error) {
	event := new(WooPPV3FeeManagerUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3OwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the WooPPV3 contract.
type WooPPV3OwnershipTransferPreparedIterator struct {
	Event *WooPPV3OwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *WooPPV3OwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3OwnershipTransferPrepared)
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
		it.Event = new(WooPPV3OwnershipTransferPrepared)
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
func (it *WooPPV3OwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3OwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3OwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the WooPPV3 contract.
type WooPPV3OwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_WooPPV3 *WooPPV3Filterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooPPV3OwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3OwnershipTransferPreparedIterator{contract: _WooPPV3.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_WooPPV3 *WooPPV3Filterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *WooPPV3OwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3OwnershipTransferPrepared)
				if err := _WooPPV3.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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

// ParseOwnershipTransferPrepared is a log parse operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_WooPPV3 *WooPPV3Filterer) ParseOwnershipTransferPrepared(log types.Log) (*WooPPV3OwnershipTransferPrepared, error) {
	event := new(WooPPV3OwnershipTransferPrepared)
	if err := _WooPPV3.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WooPPV3 contract.
type WooPPV3OwnershipTransferredIterator struct {
	Event *WooPPV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WooPPV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3OwnershipTransferred)
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
		it.Event = new(WooPPV3OwnershipTransferred)
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
func (it *WooPPV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3OwnershipTransferred represents a OwnershipTransferred event raised by the WooPPV3 contract.
type WooPPV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooPPV3 *WooPPV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WooPPV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3OwnershipTransferredIterator{contract: _WooPPV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WooPPV3 *WooPPV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WooPPV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3OwnershipTransferred)
				if err := _WooPPV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WooPPV3 *WooPPV3Filterer) ParseOwnershipTransferred(log types.Log) (*WooPPV3OwnershipTransferred, error) {
	event := new(WooPPV3OwnershipTransferred)
	if err := _WooPPV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3ParametersUpdatedIterator is returned from FilterParametersUpdated and is used to iterate over the raw logs and unpacked data for ParametersUpdated events raised by the WooPPV3 contract.
type WooPPV3ParametersUpdatedIterator struct {
	Event *WooPPV3ParametersUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3ParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3ParametersUpdated)
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
		it.Event = new(WooPPV3ParametersUpdated)
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
func (it *WooPPV3ParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3ParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3ParametersUpdated represents a ParametersUpdated event raised by the WooPPV3 contract.
type WooPPV3ParametersUpdated struct {
	BaseToken    common.Address
	NewThreshold *big.Int
	NewR         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParametersUpdated is a free log retrieval operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_WooPPV3 *WooPPV3Filterer) FilterParametersUpdated(opts *bind.FilterOpts, baseToken []common.Address) (*WooPPV3ParametersUpdatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3ParametersUpdatedIterator{contract: _WooPPV3.contract, event: "ParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchParametersUpdated is a free log subscription operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_WooPPV3 *WooPPV3Filterer) WatchParametersUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3ParametersUpdated, baseToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3ParametersUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
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

// ParseParametersUpdated is a log parse operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_WooPPV3 *WooPPV3Filterer) ParseParametersUpdated(log types.Log) (*WooPPV3ParametersUpdated, error) {
	event := new(WooPPV3ParametersUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WooPPV3 contract.
type WooPPV3PausedIterator struct {
	Event *WooPPV3Paused // Event containing the contract specifics and raw log

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
func (it *WooPPV3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3Paused)
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
		it.Event = new(WooPPV3Paused)
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
func (it *WooPPV3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3Paused represents a Paused event raised by the WooPPV3 contract.
type WooPPV3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV3 *WooPPV3Filterer) FilterPaused(opts *bind.FilterOpts) (*WooPPV3PausedIterator, error) {

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WooPPV3PausedIterator{contract: _WooPPV3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV3 *WooPPV3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WooPPV3Paused) (event.Subscription, error) {

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3Paused)
				if err := _WooPPV3.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WooPPV3 *WooPPV3Filterer) ParsePaused(log types.Log) (*WooPPV3Paused, error) {
	event := new(WooPPV3Paused)
	if err := _WooPPV3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3RewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the WooPPV3 contract.
type WooPPV3RewardManagerUpdatedIterator struct {
	Event *WooPPV3RewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3RewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3RewardManagerUpdated)
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
		it.Event = new(WooPPV3RewardManagerUpdated)
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
func (it *WooPPV3RewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3RewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3RewardManagerUpdated represents a RewardManagerUpdated event raised by the WooPPV3 contract.
type WooPPV3RewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_WooPPV3 *WooPPV3Filterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*WooPPV3RewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3RewardManagerUpdatedIterator{contract: _WooPPV3.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_WooPPV3 *WooPPV3Filterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3RewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3RewardManagerUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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

// ParseRewardManagerUpdated is a log parse operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_WooPPV3 *WooPPV3Filterer) ParseRewardManagerUpdated(log types.Log) (*WooPPV3RewardManagerUpdated, error) {
	event := new(WooPPV3RewardManagerUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3StrategistUpdatedIterator is returned from FilterStrategistUpdated and is used to iterate over the raw logs and unpacked data for StrategistUpdated events raised by the WooPPV3 contract.
type WooPPV3StrategistUpdatedIterator struct {
	Event *WooPPV3StrategistUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3StrategistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3StrategistUpdated)
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
		it.Event = new(WooPPV3StrategistUpdated)
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
func (it *WooPPV3StrategistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3StrategistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3StrategistUpdated represents a StrategistUpdated event raised by the WooPPV3 contract.
type WooPPV3StrategistUpdated struct {
	Strategist common.Address
	Flag       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategistUpdated is a free log retrieval operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_WooPPV3 *WooPPV3Filterer) FilterStrategistUpdated(opts *bind.FilterOpts, strategist []common.Address) (*WooPPV3StrategistUpdatedIterator, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3StrategistUpdatedIterator{contract: _WooPPV3.contract, event: "StrategistUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategistUpdated is a free log subscription operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_WooPPV3 *WooPPV3Filterer) WatchStrategistUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3StrategistUpdated, strategist []common.Address) (event.Subscription, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3StrategistUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
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

// ParseStrategistUpdated is a log parse operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_WooPPV3 *WooPPV3Filterer) ParseStrategistUpdated(log types.Log) (*WooPPV3StrategistUpdated, error) {
	event := new(WooPPV3StrategistUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WooPPV3 contract.
type WooPPV3UnpausedIterator struct {
	Event *WooPPV3Unpaused // Event containing the contract specifics and raw log

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
func (it *WooPPV3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3Unpaused)
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
		it.Event = new(WooPPV3Unpaused)
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
func (it *WooPPV3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3Unpaused represents a Unpaused event raised by the WooPPV3 contract.
type WooPPV3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV3 *WooPPV3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*WooPPV3UnpausedIterator, error) {

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WooPPV3UnpausedIterator{contract: _WooPPV3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV3 *WooPPV3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WooPPV3Unpaused) (event.Subscription, error) {

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3Unpaused)
				if err := _WooPPV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WooPPV3 *WooPPV3Filterer) ParseUnpaused(log types.Log) (*WooPPV3Unpaused, error) {
	event := new(WooPPV3Unpaused)
	if err := _WooPPV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the WooPPV3 contract.
type WooPPV3WithdrawIterator struct {
	Event *WooPPV3Withdraw // Event containing the contract specifics and raw log

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
func (it *WooPPV3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3Withdraw)
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
		it.Event = new(WooPPV3Withdraw)
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
func (it *WooPPV3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3Withdraw represents a Withdraw event raised by the WooPPV3 contract.
type WooPPV3Withdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_WooPPV3 *WooPPV3Filterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*WooPPV3WithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3WithdrawIterator{contract: _WooPPV3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_WooPPV3 *WooPPV3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WooPPV3Withdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3Withdraw)
				if err := _WooPPV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_WooPPV3 *WooPPV3Filterer) ParseWithdraw(log types.Log) (*WooPPV3Withdraw, error) {
	event := new(WooPPV3Withdraw)
	if err := _WooPPV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3WooGuardianUpdatedIterator is returned from FilterWooGuardianUpdated and is used to iterate over the raw logs and unpacked data for WooGuardianUpdated events raised by the WooPPV3 contract.
type WooPPV3WooGuardianUpdatedIterator struct {
	Event *WooPPV3WooGuardianUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3WooGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3WooGuardianUpdated)
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
		it.Event = new(WooPPV3WooGuardianUpdated)
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
func (it *WooPPV3WooGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3WooGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3WooGuardianUpdated represents a WooGuardianUpdated event raised by the WooPPV3 contract.
type WooPPV3WooGuardianUpdated struct {
	NewWooGuardian common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWooGuardianUpdated is a free log retrieval operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_WooPPV3 *WooPPV3Filterer) FilterWooGuardianUpdated(opts *bind.FilterOpts, newWooGuardian []common.Address) (*WooPPV3WooGuardianUpdatedIterator, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3WooGuardianUpdatedIterator{contract: _WooPPV3.contract, event: "WooGuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchWooGuardianUpdated is a free log subscription operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_WooPPV3 *WooPPV3Filterer) WatchWooGuardianUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3WooGuardianUpdated, newWooGuardian []common.Address) (event.Subscription, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3WooGuardianUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
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

// ParseWooGuardianUpdated is a log parse operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_WooPPV3 *WooPPV3Filterer) ParseWooGuardianUpdated(log types.Log) (*WooPPV3WooGuardianUpdated, error) {
	event := new(WooPPV3WooGuardianUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3WooSwapIterator is returned from FilterWooSwap and is used to iterate over the raw logs and unpacked data for WooSwap events raised by the WooPPV3 contract.
type WooPPV3WooSwapIterator struct {
	Event *WooPPV3WooSwap // Event containing the contract specifics and raw log

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
func (it *WooPPV3WooSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3WooSwap)
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
		it.Event = new(WooPPV3WooSwap)
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
func (it *WooPPV3WooSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3WooSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3WooSwap represents a WooSwap event raised by the WooPPV3 contract.
type WooPPV3WooSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	From       common.Address
	To         common.Address
	RebateTo   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWooSwap is a free log retrieval operation binding the contract event 0x74ef34e2ea7c5d9f7b7ed44e97ad44b4303416c3a660c3fb5b3bdb95a1d6abd3.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_WooPPV3 *WooPPV3Filterer) FilterWooSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*WooPPV3WooSwapIterator, error) {

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

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3WooSwapIterator{contract: _WooPPV3.contract, event: "WooSwap", logs: logs, sub: sub}, nil
}

// WatchWooSwap is a free log subscription operation binding the contract event 0x74ef34e2ea7c5d9f7b7ed44e97ad44b4303416c3a660c3fb5b3bdb95a1d6abd3.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_WooPPV3 *WooPPV3Filterer) WatchWooSwap(opts *bind.WatchOpts, sink chan<- *WooPPV3WooSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3WooSwap)
				if err := _WooPPV3.contract.UnpackLog(event, "WooSwap", log); err != nil {
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

// ParseWooSwap is a log parse operation binding the contract event 0x74ef34e2ea7c5d9f7b7ed44e97ad44b4303416c3a660c3fb5b3bdb95a1d6abd3.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_WooPPV3 *WooPPV3Filterer) ParseWooSwap(log types.Log) (*WooPPV3WooSwap, error) {
	event := new(WooPPV3WooSwap)
	if err := _WooPPV3.contract.UnpackLog(event, "WooSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WooPPV3WooracleUpdatedIterator is returned from FilterWooracleUpdated and is used to iterate over the raw logs and unpacked data for WooracleUpdated events raised by the WooPPV3 contract.
type WooPPV3WooracleUpdatedIterator struct {
	Event *WooPPV3WooracleUpdated // Event containing the contract specifics and raw log

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
func (it *WooPPV3WooracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WooPPV3WooracleUpdated)
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
		it.Event = new(WooPPV3WooracleUpdated)
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
func (it *WooPPV3WooracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WooPPV3WooracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WooPPV3WooracleUpdated represents a WooracleUpdated event raised by the WooPPV3 contract.
type WooPPV3WooracleUpdated struct {
	NewWooracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWooracleUpdated is a free log retrieval operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV3 *WooPPV3Filterer) FilterWooracleUpdated(opts *bind.FilterOpts, newWooracle []common.Address) (*WooPPV3WooracleUpdatedIterator, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _WooPPV3.contract.FilterLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return &WooPPV3WooracleUpdatedIterator{contract: _WooPPV3.contract, event: "WooracleUpdated", logs: logs, sub: sub}, nil
}

// WatchWooracleUpdated is a free log subscription operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV3 *WooPPV3Filterer) WatchWooracleUpdated(opts *bind.WatchOpts, sink chan<- *WooPPV3WooracleUpdated, newWooracle []common.Address) (event.Subscription, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _WooPPV3.contract.WatchLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WooPPV3WooracleUpdated)
				if err := _WooPPV3.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
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

// ParseWooracleUpdated is a log parse operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_WooPPV3 *WooPPV3Filterer) ParseWooracleUpdated(log types.Log) (*WooPPV3WooracleUpdated, error) {
	event := new(WooPPV3WooracleUpdated)
	if err := _WooPPV3.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
