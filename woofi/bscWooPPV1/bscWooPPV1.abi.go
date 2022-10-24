// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bscWooPPV1

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

// BscWooPPV1MetaData contains all meta data concerning the BscWooPPV1 contract.
var BscWooPPV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newQuoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"ParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRewardManager\",\"type\":\"address\"}],\"name\":\"RewardManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"StrategistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"WooGuardianUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WooSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"WooracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStrategist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairsInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"removeBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPairsInfo\",\"type\":\"string\"}],\"name\":\"setPairsInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRewardManager\",\"type\":\"address\"}],\"name\":\"setRewardManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"setWooGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"setWooracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"threshold\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastResetTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lpFeeRate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"R\",\"type\":\"uint64\"},{\"internalType\":\"uint112\",\"name\":\"target\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooGuardian\",\"outputs\":[{\"internalType\":\"contractIWooGuardian\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BscWooPPV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use BscWooPPV1MetaData.ABI instead.
var BscWooPPV1ABI = BscWooPPV1MetaData.ABI

// BscWooPPV1 is an auto generated Go binding around an Ethereum contract.
type BscWooPPV1 struct {
	BscWooPPV1Caller     // Read-only binding to the contract
	BscWooPPV1Transactor // Write-only binding to the contract
	BscWooPPV1Filterer   // Log filterer for contract events
}

// BscWooPPV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BscWooPPV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooPPV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BscWooPPV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooPPV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscWooPPV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWooPPV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscWooPPV1Session struct {
	Contract     *BscWooPPV1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscWooPPV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscWooPPV1CallerSession struct {
	Contract *BscWooPPV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BscWooPPV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscWooPPV1TransactorSession struct {
	Contract     *BscWooPPV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BscWooPPV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BscWooPPV1Raw struct {
	Contract *BscWooPPV1 // Generic contract binding to access the raw methods on
}

// BscWooPPV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscWooPPV1CallerRaw struct {
	Contract *BscWooPPV1Caller // Generic read-only contract binding to access the raw methods on
}

// BscWooPPV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscWooPPV1TransactorRaw struct {
	Contract *BscWooPPV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBscWooPPV1 creates a new instance of BscWooPPV1, bound to a specific deployed contract.
func NewBscWooPPV1(address common.Address, backend bind.ContractBackend) (*BscWooPPV1, error) {
	contract, err := bindBscWooPPV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1{BscWooPPV1Caller: BscWooPPV1Caller{contract: contract}, BscWooPPV1Transactor: BscWooPPV1Transactor{contract: contract}, BscWooPPV1Filterer: BscWooPPV1Filterer{contract: contract}}, nil
}

// NewBscWooPPV1Caller creates a new read-only instance of BscWooPPV1, bound to a specific deployed contract.
func NewBscWooPPV1Caller(address common.Address, caller bind.ContractCaller) (*BscWooPPV1Caller, error) {
	contract, err := bindBscWooPPV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1Caller{contract: contract}, nil
}

// NewBscWooPPV1Transactor creates a new write-only instance of BscWooPPV1, bound to a specific deployed contract.
func NewBscWooPPV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BscWooPPV1Transactor, error) {
	contract, err := bindBscWooPPV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1Transactor{contract: contract}, nil
}

// NewBscWooPPV1Filterer creates a new log filterer instance of BscWooPPV1, bound to a specific deployed contract.
func NewBscWooPPV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BscWooPPV1Filterer, error) {
	contract, err := bindBscWooPPV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1Filterer{contract: contract}, nil
}

// bindBscWooPPV1 binds a generic wrapper to an already deployed contract.
func bindBscWooPPV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscWooPPV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWooPPV1 *BscWooPPV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWooPPV1.Contract.BscWooPPV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWooPPV1 *BscWooPPV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.BscWooPPV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWooPPV1 *BscWooPPV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.BscWooPPV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWooPPV1 *BscWooPPV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWooPPV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWooPPV1 *BscWooPPV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWooPPV1 *BscWooPPV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.contract.Transact(opts, method, params...)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) NEWOWNER() (common.Address, error) {
	return _BscWooPPV1.Contract.NEWOWNER(&_BscWooPPV1.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) NEWOWNER() (common.Address, error) {
	return _BscWooPPV1.Contract.NEWOWNER(&_BscWooPPV1.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) OWNER() (common.Address, error) {
	return _BscWooPPV1.Contract.OWNER(&_BscWooPPV1.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) OWNER() (common.Address, error) {
	return _BscWooPPV1.Contract.OWNER(&_BscWooPPV1.CallOpts)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWooPPV1 *BscWooPPV1Caller) IsStrategist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "isStrategist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWooPPV1 *BscWooPPV1Session) IsStrategist(arg0 common.Address) (bool, error) {
	return _BscWooPPV1.Contract.IsStrategist(&_BscWooPPV1.CallOpts, arg0)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWooPPV1 *BscWooPPV1CallerSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _BscWooPPV1.Contract.IsStrategist(&_BscWooPPV1.CallOpts, arg0)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWooPPV1 *BscWooPPV1Caller) PairsInfo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "pairsInfo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWooPPV1 *BscWooPPV1Session) PairsInfo() (string, error) {
	return _BscWooPPV1.Contract.PairsInfo(&_BscWooPPV1.CallOpts)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWooPPV1 *BscWooPPV1CallerSession) PairsInfo() (string, error) {
	return _BscWooPPV1.Contract.PairsInfo(&_BscWooPPV1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWooPPV1 *BscWooPPV1Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWooPPV1 *BscWooPPV1Session) Paused() (bool, error) {
	return _BscWooPPV1.Contract.Paused(&_BscWooPPV1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWooPPV1 *BscWooPPV1CallerSession) Paused() (bool, error) {
	return _BscWooPPV1.Contract.Paused(&_BscWooPPV1.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWooPPV1 *BscWooPPV1Caller) PoolSize(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "poolSize", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWooPPV1 *BscWooPPV1Session) PoolSize(token common.Address) (*big.Int, error) {
	return _BscWooPPV1.Contract.PoolSize(&_BscWooPPV1.CallOpts, token)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWooPPV1 *BscWooPPV1CallerSession) PoolSize(token common.Address) (*big.Int, error) {
	return _BscWooPPV1.Contract.PoolSize(&_BscWooPPV1.CallOpts, token)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1Caller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1Session) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWooPPV1.Contract.QuerySellBase(&_BscWooPPV1.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1CallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWooPPV1.Contract.QuerySellBase(&_BscWooPPV1.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1Caller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1Session) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWooPPV1.Contract.QuerySellQuote(&_BscWooPPV1.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1CallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWooPPV1.Contract.QuerySellQuote(&_BscWooPPV1.CallOpts, baseToken, quoteAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) QuoteToken() (common.Address, error) {
	return _BscWooPPV1.Contract.QuoteToken(&_BscWooPPV1.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) QuoteToken() (common.Address, error) {
	return _BscWooPPV1.Contract.QuoteToken(&_BscWooPPV1.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) RewardManager() (common.Address, error) {
	return _BscWooPPV1.Contract.RewardManager(&_BscWooPPV1.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) RewardManager() (common.Address, error) {
	return _BscWooPPV1.Contract.RewardManager(&_BscWooPPV1.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 lpFeeRate, uint64 R, uint112 target, bool isValid)
func (_BscWooPPV1 *BscWooPPV1Caller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	LpFeeRate          uint64
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "tokenInfo", arg0)

	outstruct := new(struct {
		Reserve            *big.Int
		Threshold          *big.Int
		LastResetTimestamp uint32
		LpFeeRate          uint64
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
	outstruct.LpFeeRate = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.R = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Target = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 lpFeeRate, uint64 R, uint112 target, bool isValid)
func (_BscWooPPV1 *BscWooPPV1Session) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	LpFeeRate          uint64
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _BscWooPPV1.Contract.TokenInfo(&_BscWooPPV1.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 lpFeeRate, uint64 R, uint112 target, bool isValid)
func (_BscWooPPV1 *BscWooPPV1CallerSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	LpFeeRate          uint64
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _BscWooPPV1.Contract.TokenInfo(&_BscWooPPV1.CallOpts, arg0)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) WooGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "wooGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) WooGuardian() (common.Address, error) {
	return _BscWooPPV1.Contract.WooGuardian(&_BscWooPPV1.CallOpts)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) WooGuardian() (common.Address, error) {
	return _BscWooPPV1.Contract.WooGuardian(&_BscWooPPV1.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Caller) Wooracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWooPPV1.contract.Call(opts, &out, "wooracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWooPPV1 *BscWooPPV1Session) Wooracle() (common.Address, error) {
	return _BscWooPPV1.Contract.Wooracle(&_BscWooPPV1.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWooPPV1 *BscWooPPV1CallerSession) Wooracle() (common.Address, error) {
	return _BscWooPPV1.Contract.Wooracle(&_BscWooPPV1.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x4bb330b4.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 lpFeeRate, uint256 R) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) AddBaseToken(opts *bind.TransactOpts, baseToken common.Address, threshold *big.Int, lpFeeRate *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "addBaseToken", baseToken, threshold, lpFeeRate, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x4bb330b4.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 lpFeeRate, uint256 R) returns()
func (_BscWooPPV1 *BscWooPPV1Session) AddBaseToken(baseToken common.Address, threshold *big.Int, lpFeeRate *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.AddBaseToken(&_BscWooPPV1.TransactOpts, baseToken, threshold, lpFeeRate, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x4bb330b4.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 lpFeeRate, uint256 R) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) AddBaseToken(baseToken common.Address, threshold *big.Int, lpFeeRate *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.AddBaseToken(&_BscWooPPV1.TransactOpts, baseToken, threshold, lpFeeRate, R)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWooPPV1 *BscWooPPV1Session) ClaimOwnership() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.ClaimOwnership(&_BscWooPPV1.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.ClaimOwnership(&_BscWooPPV1.TransactOpts)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1Session) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.InitOwner(&_BscWooPPV1.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.InitOwner(&_BscWooPPV1.TransactOpts, newOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWooPPV1 *BscWooPPV1Session) Pause() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Pause(&_BscWooPPV1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) Pause() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Pause(&_BscWooPPV1.TransactOpts)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) RemoveBaseToken(opts *bind.TransactOpts, baseToken common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "removeBaseToken", baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWooPPV1 *BscWooPPV1Session) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.RemoveBaseToken(&_BscWooPPV1.TransactOpts, baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.RemoveBaseToken(&_BscWooPPV1.TransactOpts, baseToken)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1Transactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1Session) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SellBase(&_BscWooPPV1.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SellBase(&_BscWooPPV1.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1Transactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1Session) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SellQuote(&_BscWooPPV1.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SellQuote(&_BscWooPPV1.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) SetPairsInfo(opts *bind.TransactOpts, newPairsInfo string) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "setPairsInfo", newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWooPPV1 *BscWooPPV1Session) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetPairsInfo(&_BscWooPPV1.TransactOpts, newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetPairsInfo(&_BscWooPPV1.TransactOpts, newPairsInfo)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address newRewardManager) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) SetRewardManager(opts *bind.TransactOpts, newRewardManager common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "setRewardManager", newRewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address newRewardManager) returns()
func (_BscWooPPV1 *BscWooPPV1Session) SetRewardManager(newRewardManager common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetRewardManager(&_BscWooPPV1.TransactOpts, newRewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address newRewardManager) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SetRewardManager(newRewardManager common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetRewardManager(&_BscWooPPV1.TransactOpts, newRewardManager)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) SetStrategist(opts *bind.TransactOpts, strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "setStrategist", strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWooPPV1 *BscWooPPV1Session) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetStrategist(&_BscWooPPV1.TransactOpts, strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetStrategist(&_BscWooPPV1.TransactOpts, strategist, flag)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) SetWooGuardian(opts *bind.TransactOpts, newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "setWooGuardian", newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWooPPV1 *BscWooPPV1Session) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetWooGuardian(&_BscWooPPV1.TransactOpts, newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetWooGuardian(&_BscWooPPV1.TransactOpts, newWooGuardian)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) SetWooracle(opts *bind.TransactOpts, newWooracle common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "setWooracle", newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWooPPV1 *BscWooPPV1Session) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetWooracle(&_BscWooPPV1.TransactOpts, newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.SetWooracle(&_BscWooPPV1.TransactOpts, newWooracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.TransferOwnership(&_BscWooPPV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.TransferOwnership(&_BscWooPPV1.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x25a0141e.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) TuneParameters(opts *bind.TransactOpts, token common.Address, newThreshold *big.Int, newLpFeeRate *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "tuneParameters", token, newThreshold, newLpFeeRate, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x25a0141e.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR) returns()
func (_BscWooPPV1 *BscWooPPV1Session) TuneParameters(token common.Address, newThreshold *big.Int, newLpFeeRate *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.TuneParameters(&_BscWooPPV1.TransactOpts, token, newThreshold, newLpFeeRate, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x25a0141e.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) TuneParameters(token common.Address, newThreshold *big.Int, newLpFeeRate *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.TuneParameters(&_BscWooPPV1.TransactOpts, token, newThreshold, newLpFeeRate, newR)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWooPPV1 *BscWooPPV1Session) Unpause() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Unpause(&_BscWooPPV1.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) Unpause() (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Unpause(&_BscWooPPV1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1Session) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Withdraw(&_BscWooPPV1.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.Withdraw(&_BscWooPPV1.TransactOpts, token, to, amount)
}

// WithdrawToOwner is a paid mutator transaction binding the contract method 0x47c62a0d.
//
// Solidity: function withdrawToOwner(address token, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1Transactor) WithdrawToOwner(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.contract.Transact(opts, "withdrawToOwner", token, amount)
}

// WithdrawToOwner is a paid mutator transaction binding the contract method 0x47c62a0d.
//
// Solidity: function withdrawToOwner(address token, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1Session) WithdrawToOwner(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.WithdrawToOwner(&_BscWooPPV1.TransactOpts, token, amount)
}

// WithdrawToOwner is a paid mutator transaction binding the contract method 0x47c62a0d.
//
// Solidity: function withdrawToOwner(address token, uint256 amount) returns()
func (_BscWooPPV1 *BscWooPPV1TransactorSession) WithdrawToOwner(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWooPPV1.Contract.WithdrawToOwner(&_BscWooPPV1.TransactOpts, token, amount)
}

// BscWooPPV1OwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the BscWooPPV1 contract.
type BscWooPPV1OwnershipTransferPreparedIterator struct {
	Event *BscWooPPV1OwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1OwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1OwnershipTransferPrepared)
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
		it.Event = new(BscWooPPV1OwnershipTransferPrepared)
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
func (it *BscWooPPV1OwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1OwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1OwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the BscWooPPV1 contract.
type BscWooPPV1OwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscWooPPV1OwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1OwnershipTransferPreparedIterator{contract: _BscWooPPV1.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *BscWooPPV1OwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1OwnershipTransferPrepared)
				if err := _BscWooPPV1.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseOwnershipTransferPrepared(log types.Log) (*BscWooPPV1OwnershipTransferPrepared, error) {
	event := new(BscWooPPV1OwnershipTransferPrepared)
	if err := _BscWooPPV1.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscWooPPV1 contract.
type BscWooPPV1OwnershipTransferredIterator struct {
	Event *BscWooPPV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1OwnershipTransferred)
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
		it.Event = new(BscWooPPV1OwnershipTransferred)
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
func (it *BscWooPPV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1OwnershipTransferred represents a OwnershipTransferred event raised by the BscWooPPV1 contract.
type BscWooPPV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscWooPPV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1OwnershipTransferredIterator{contract: _BscWooPPV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscWooPPV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1OwnershipTransferred)
				if err := _BscWooPPV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseOwnershipTransferred(log types.Log) (*BscWooPPV1OwnershipTransferred, error) {
	event := new(BscWooPPV1OwnershipTransferred)
	if err := _BscWooPPV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1ParametersUpdatedIterator is returned from FilterParametersUpdated and is used to iterate over the raw logs and unpacked data for ParametersUpdated events raised by the BscWooPPV1 contract.
type BscWooPPV1ParametersUpdatedIterator struct {
	Event *BscWooPPV1ParametersUpdated // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1ParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1ParametersUpdated)
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
		it.Event = new(BscWooPPV1ParametersUpdated)
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
func (it *BscWooPPV1ParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1ParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1ParametersUpdated represents a ParametersUpdated event raised by the BscWooPPV1 contract.
type BscWooPPV1ParametersUpdated struct {
	BaseToken    common.Address
	NewThreshold *big.Int
	NewLpFeeRate *big.Int
	NewR         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParametersUpdated is a free log retrieval operation binding the contract event 0xd057107c38b4759e3eb195ba65ec10fba9f2018d5b01c5019522e52cb4558bc2.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterParametersUpdated(opts *bind.FilterOpts, baseToken []common.Address) (*BscWooPPV1ParametersUpdatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1ParametersUpdatedIterator{contract: _BscWooPPV1.contract, event: "ParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchParametersUpdated is a free log subscription operation binding the contract event 0xd057107c38b4759e3eb195ba65ec10fba9f2018d5b01c5019522e52cb4558bc2.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchParametersUpdated(opts *bind.WatchOpts, sink chan<- *BscWooPPV1ParametersUpdated, baseToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1ParametersUpdated)
				if err := _BscWooPPV1.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
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

// ParseParametersUpdated is a log parse operation binding the contract event 0xd057107c38b4759e3eb195ba65ec10fba9f2018d5b01c5019522e52cb4558bc2.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newLpFeeRate, uint256 newR)
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseParametersUpdated(log types.Log) (*BscWooPPV1ParametersUpdated, error) {
	event := new(BscWooPPV1ParametersUpdated)
	if err := _BscWooPPV1.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BscWooPPV1 contract.
type BscWooPPV1PausedIterator struct {
	Event *BscWooPPV1Paused // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1Paused)
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
		it.Event = new(BscWooPPV1Paused)
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
func (it *BscWooPPV1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1Paused represents a Paused event raised by the BscWooPPV1 contract.
type BscWooPPV1Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterPaused(opts *bind.FilterOpts) (*BscWooPPV1PausedIterator, error) {

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1PausedIterator{contract: _BscWooPPV1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BscWooPPV1Paused) (event.Subscription, error) {

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1Paused)
				if err := _BscWooPPV1.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParsePaused(log types.Log) (*BscWooPPV1Paused, error) {
	event := new(BscWooPPV1Paused)
	if err := _BscWooPPV1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1RewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the BscWooPPV1 contract.
type BscWooPPV1RewardManagerUpdatedIterator struct {
	Event *BscWooPPV1RewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1RewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1RewardManagerUpdated)
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
		it.Event = new(BscWooPPV1RewardManagerUpdated)
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
func (it *BscWooPPV1RewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1RewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1RewardManagerUpdated represents a RewardManagerUpdated event raised by the BscWooPPV1 contract.
type BscWooPPV1RewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*BscWooPPV1RewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1RewardManagerUpdatedIterator{contract: _BscWooPPV1.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *BscWooPPV1RewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1RewardManagerUpdated)
				if err := _BscWooPPV1.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseRewardManagerUpdated(log types.Log) (*BscWooPPV1RewardManagerUpdated, error) {
	event := new(BscWooPPV1RewardManagerUpdated)
	if err := _BscWooPPV1.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1StrategistUpdatedIterator is returned from FilterStrategistUpdated and is used to iterate over the raw logs and unpacked data for StrategistUpdated events raised by the BscWooPPV1 contract.
type BscWooPPV1StrategistUpdatedIterator struct {
	Event *BscWooPPV1StrategistUpdated // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1StrategistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1StrategistUpdated)
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
		it.Event = new(BscWooPPV1StrategistUpdated)
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
func (it *BscWooPPV1StrategistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1StrategistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1StrategistUpdated represents a StrategistUpdated event raised by the BscWooPPV1 contract.
type BscWooPPV1StrategistUpdated struct {
	Strategist common.Address
	Flag       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategistUpdated is a free log retrieval operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterStrategistUpdated(opts *bind.FilterOpts, strategist []common.Address) (*BscWooPPV1StrategistUpdatedIterator, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1StrategistUpdatedIterator{contract: _BscWooPPV1.contract, event: "StrategistUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategistUpdated is a free log subscription operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchStrategistUpdated(opts *bind.WatchOpts, sink chan<- *BscWooPPV1StrategistUpdated, strategist []common.Address) (event.Subscription, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1StrategistUpdated)
				if err := _BscWooPPV1.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseStrategistUpdated(log types.Log) (*BscWooPPV1StrategistUpdated, error) {
	event := new(BscWooPPV1StrategistUpdated)
	if err := _BscWooPPV1.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BscWooPPV1 contract.
type BscWooPPV1UnpausedIterator struct {
	Event *BscWooPPV1Unpaused // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1Unpaused)
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
		it.Event = new(BscWooPPV1Unpaused)
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
func (it *BscWooPPV1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1Unpaused represents a Unpaused event raised by the BscWooPPV1 contract.
type BscWooPPV1Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterUnpaused(opts *bind.FilterOpts) (*BscWooPPV1UnpausedIterator, error) {

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1UnpausedIterator{contract: _BscWooPPV1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BscWooPPV1Unpaused) (event.Subscription, error) {

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1Unpaused)
				if err := _BscWooPPV1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseUnpaused(log types.Log) (*BscWooPPV1Unpaused, error) {
	event := new(BscWooPPV1Unpaused)
	if err := _BscWooPPV1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BscWooPPV1 contract.
type BscWooPPV1WithdrawIterator struct {
	Event *BscWooPPV1Withdraw // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1Withdraw)
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
		it.Event = new(BscWooPPV1Withdraw)
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
func (it *BscWooPPV1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1Withdraw represents a Withdraw event raised by the BscWooPPV1 contract.
type BscWooPPV1Withdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BscWooPPV1WithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1WithdrawIterator{contract: _BscWooPPV1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BscWooPPV1Withdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1Withdraw)
				if err := _BscWooPPV1.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseWithdraw(log types.Log) (*BscWooPPV1Withdraw, error) {
	event := new(BscWooPPV1Withdraw)
	if err := _BscWooPPV1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1WooGuardianUpdatedIterator is returned from FilterWooGuardianUpdated and is used to iterate over the raw logs and unpacked data for WooGuardianUpdated events raised by the BscWooPPV1 contract.
type BscWooPPV1WooGuardianUpdatedIterator struct {
	Event *BscWooPPV1WooGuardianUpdated // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1WooGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1WooGuardianUpdated)
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
		it.Event = new(BscWooPPV1WooGuardianUpdated)
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
func (it *BscWooPPV1WooGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1WooGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1WooGuardianUpdated represents a WooGuardianUpdated event raised by the BscWooPPV1 contract.
type BscWooPPV1WooGuardianUpdated struct {
	NewWooGuardian common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWooGuardianUpdated is a free log retrieval operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterWooGuardianUpdated(opts *bind.FilterOpts, newWooGuardian []common.Address) (*BscWooPPV1WooGuardianUpdatedIterator, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1WooGuardianUpdatedIterator{contract: _BscWooPPV1.contract, event: "WooGuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchWooGuardianUpdated is a free log subscription operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchWooGuardianUpdated(opts *bind.WatchOpts, sink chan<- *BscWooPPV1WooGuardianUpdated, newWooGuardian []common.Address) (event.Subscription, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1WooGuardianUpdated)
				if err := _BscWooPPV1.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseWooGuardianUpdated(log types.Log) (*BscWooPPV1WooGuardianUpdated, error) {
	event := new(BscWooPPV1WooGuardianUpdated)
	if err := _BscWooPPV1.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1WooSwapIterator is returned from FilterWooSwap and is used to iterate over the raw logs and unpacked data for WooSwap events raised by the BscWooPPV1 contract.
type BscWooPPV1WooSwapIterator struct {
	Event *BscWooPPV1WooSwap // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1WooSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1WooSwap)
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
		it.Event = new(BscWooPPV1WooSwap)
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
func (it *BscWooPPV1WooSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1WooSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1WooSwap represents a WooSwap event raised by the BscWooPPV1 contract.
type BscWooPPV1WooSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	From       common.Address
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWooSwap is a free log retrieval operation binding the contract event 0x90d51bc25f44dddb7620733c08a051656d433bed5de7322f9988d78f123ab5d1.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterWooSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*BscWooPPV1WooSwapIterator, error) {

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

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1WooSwapIterator{contract: _BscWooPPV1.contract, event: "WooSwap", logs: logs, sub: sub}, nil
}

// WatchWooSwap is a free log subscription operation binding the contract event 0x90d51bc25f44dddb7620733c08a051656d433bed5de7322f9988d78f123ab5d1.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchWooSwap(opts *bind.WatchOpts, sink chan<- *BscWooPPV1WooSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1WooSwap)
				if err := _BscWooPPV1.contract.UnpackLog(event, "WooSwap", log); err != nil {
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

// ParseWooSwap is a log parse operation binding the contract event 0x90d51bc25f44dddb7620733c08a051656d433bed5de7322f9988d78f123ab5d1.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to)
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseWooSwap(log types.Log) (*BscWooPPV1WooSwap, error) {
	event := new(BscWooPPV1WooSwap)
	if err := _BscWooPPV1.contract.UnpackLog(event, "WooSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWooPPV1WooracleUpdatedIterator is returned from FilterWooracleUpdated and is used to iterate over the raw logs and unpacked data for WooracleUpdated events raised by the BscWooPPV1 contract.
type BscWooPPV1WooracleUpdatedIterator struct {
	Event *BscWooPPV1WooracleUpdated // Event containing the contract specifics and raw log

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
func (it *BscWooPPV1WooracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWooPPV1WooracleUpdated)
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
		it.Event = new(BscWooPPV1WooracleUpdated)
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
func (it *BscWooPPV1WooracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWooPPV1WooracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWooPPV1WooracleUpdated represents a WooracleUpdated event raised by the BscWooPPV1 contract.
type BscWooPPV1WooracleUpdated struct {
	NewWooracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWooracleUpdated is a free log retrieval operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_BscWooPPV1 *BscWooPPV1Filterer) FilterWooracleUpdated(opts *bind.FilterOpts, newWooracle []common.Address) (*BscWooPPV1WooracleUpdatedIterator, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _BscWooPPV1.contract.FilterLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return &BscWooPPV1WooracleUpdatedIterator{contract: _BscWooPPV1.contract, event: "WooracleUpdated", logs: logs, sub: sub}, nil
}

// WatchWooracleUpdated is a free log subscription operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_BscWooPPV1 *BscWooPPV1Filterer) WatchWooracleUpdated(opts *bind.WatchOpts, sink chan<- *BscWooPPV1WooracleUpdated, newWooracle []common.Address) (event.Subscription, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _BscWooPPV1.contract.WatchLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWooPPV1WooracleUpdated)
				if err := _BscWooPPV1.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
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
func (_BscWooPPV1 *BscWooPPV1Filterer) ParseWooracleUpdated(log types.Log) (*BscWooPPV1WooracleUpdated, error) {
	event := new(BscWooPPV1WooracleUpdated)
	if err := _BscWooPPV1.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
