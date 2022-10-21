// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bscWOOPP

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

// BscWOOPPMetaData contains all meta data concerning the BscWOOPP contract.
var BscWOOPPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newQuoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"ParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRewardManager\",\"type\":\"address\"}],\"name\":\"RewardManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"StrategistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"WooGuardianUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"WooracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractIWooFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStrategist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairsInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"removeBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPairsInfo\",\"type\":\"string\"}],\"name\":\"setPairsInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"setWooGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"setWooracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"threshold\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastResetTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"R\",\"type\":\"uint64\"},{\"internalType\":\"uint112\",\"name\":\"target\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAllToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooGuardian\",\"outputs\":[{\"internalType\":\"contractIWooGuardian\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BscWOOPPABI is the input ABI used to generate the binding from.
// Deprecated: Use BscWOOPPMetaData.ABI instead.
var BscWOOPPABI = BscWOOPPMetaData.ABI

// BscWOOPP is an auto generated Go binding around an Ethereum contract.
type BscWOOPP struct {
	BscWOOPPCaller     // Read-only binding to the contract
	BscWOOPPTransactor // Write-only binding to the contract
	BscWOOPPFilterer   // Log filterer for contract events
}

// BscWOOPPCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscWOOPPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWOOPPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscWOOPPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWOOPPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscWOOPPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscWOOPPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscWOOPPSession struct {
	Contract     *BscWOOPP         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscWOOPPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscWOOPPCallerSession struct {
	Contract *BscWOOPPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BscWOOPPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscWOOPPTransactorSession struct {
	Contract     *BscWOOPPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BscWOOPPRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscWOOPPRaw struct {
	Contract *BscWOOPP // Generic contract binding to access the raw methods on
}

// BscWOOPPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscWOOPPCallerRaw struct {
	Contract *BscWOOPPCaller // Generic read-only contract binding to access the raw methods on
}

// BscWOOPPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscWOOPPTransactorRaw struct {
	Contract *BscWOOPPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscWOOPP creates a new instance of BscWOOPP, bound to a specific deployed contract.
func NewBscWOOPP(address common.Address, backend bind.ContractBackend) (*BscWOOPP, error) {
	contract, err := bindBscWOOPP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscWOOPP{BscWOOPPCaller: BscWOOPPCaller{contract: contract}, BscWOOPPTransactor: BscWOOPPTransactor{contract: contract}, BscWOOPPFilterer: BscWOOPPFilterer{contract: contract}}, nil
}

// NewBscWOOPPCaller creates a new read-only instance of BscWOOPP, bound to a specific deployed contract.
func NewBscWOOPPCaller(address common.Address, caller bind.ContractCaller) (*BscWOOPPCaller, error) {
	contract, err := bindBscWOOPP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPCaller{contract: contract}, nil
}

// NewBscWOOPPTransactor creates a new write-only instance of BscWOOPP, bound to a specific deployed contract.
func NewBscWOOPPTransactor(address common.Address, transactor bind.ContractTransactor) (*BscWOOPPTransactor, error) {
	contract, err := bindBscWOOPP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPTransactor{contract: contract}, nil
}

// NewBscWOOPPFilterer creates a new log filterer instance of BscWOOPP, bound to a specific deployed contract.
func NewBscWOOPPFilterer(address common.Address, filterer bind.ContractFilterer) (*BscWOOPPFilterer, error) {
	contract, err := bindBscWOOPP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPFilterer{contract: contract}, nil
}

// bindBscWOOPP binds a generic wrapper to an already deployed contract.
func bindBscWOOPP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscWOOPPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWOOPP *BscWOOPPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWOOPP.Contract.BscWOOPPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWOOPP *BscWOOPPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWOOPP.Contract.BscWOOPPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWOOPP *BscWOOPPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWOOPP.Contract.BscWOOPPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscWOOPP *BscWOOPPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscWOOPP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscWOOPP *BscWOOPPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWOOPP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscWOOPP *BscWOOPPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscWOOPP.Contract.contract.Transact(opts, method, params...)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPSession) NEWOWNER() (common.Address, error) {
	return _BscWOOPP.Contract.NEWOWNER(&_BscWOOPP.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) NEWOWNER() (common.Address, error) {
	return _BscWOOPP.Contract.NEWOWNER(&_BscWOOPP.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPSession) OWNER() (common.Address, error) {
	return _BscWOOPP.Contract.OWNER(&_BscWOOPP.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) OWNER() (common.Address, error) {
	return _BscWOOPP.Contract.OWNER(&_BscWOOPP.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_BscWOOPP *BscWOOPPSession) FeeManager() (common.Address, error) {
	return _BscWOOPP.Contract.FeeManager(&_BscWOOPP.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) FeeManager() (common.Address, error) {
	return _BscWOOPP.Contract.FeeManager(&_BscWOOPP.CallOpts)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWOOPP *BscWOOPPCaller) IsStrategist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "isStrategist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWOOPP *BscWOOPPSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _BscWOOPP.Contract.IsStrategist(&_BscWOOPP.CallOpts, arg0)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_BscWOOPP *BscWOOPPCallerSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _BscWOOPP.Contract.IsStrategist(&_BscWOOPP.CallOpts, arg0)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWOOPP *BscWOOPPCaller) PairsInfo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "pairsInfo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWOOPP *BscWOOPPSession) PairsInfo() (string, error) {
	return _BscWOOPP.Contract.PairsInfo(&_BscWOOPP.CallOpts)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_BscWOOPP *BscWOOPPCallerSession) PairsInfo() (string, error) {
	return _BscWOOPP.Contract.PairsInfo(&_BscWOOPP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWOOPP *BscWOOPPCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWOOPP *BscWOOPPSession) Paused() (bool, error) {
	return _BscWOOPP.Contract.Paused(&_BscWOOPP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BscWOOPP *BscWOOPPCallerSession) Paused() (bool, error) {
	return _BscWOOPP.Contract.Paused(&_BscWOOPP.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWOOPP *BscWOOPPCaller) PoolSize(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "poolSize", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWOOPP *BscWOOPPSession) PoolSize(token common.Address) (*big.Int, error) {
	return _BscWOOPP.Contract.PoolSize(&_BscWOOPP.CallOpts, token)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_BscWOOPP *BscWOOPPCallerSession) PoolSize(token common.Address) (*big.Int, error) {
	return _BscWOOPP.Contract.PoolSize(&_BscWOOPP.CallOpts, token)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPCaller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWOOPP.Contract.QuerySellBase(&_BscWOOPP.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPCallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _BscWOOPP.Contract.QuerySellBase(&_BscWOOPP.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPCaller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWOOPP.Contract.QuerySellQuote(&_BscWOOPP.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPCallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _BscWOOPP.Contract.QuerySellQuote(&_BscWOOPP.CallOpts, baseToken, quoteAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWOOPP *BscWOOPPSession) QuoteToken() (common.Address, error) {
	return _BscWOOPP.Contract.QuoteToken(&_BscWOOPP.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) QuoteToken() (common.Address, error) {
	return _BscWOOPP.Contract.QuoteToken(&_BscWOOPP.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_BscWOOPP *BscWOOPPCaller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "tokenInfo", arg0)

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
func (_BscWOOPP *BscWOOPPSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _BscWOOPP.Contract.TokenInfo(&_BscWOOPP.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_BscWOOPP *BscWOOPPCallerSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _BscWOOPP.Contract.TokenInfo(&_BscWOOPP.CallOpts, arg0)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) WooGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "wooGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWOOPP *BscWOOPPSession) WooGuardian() (common.Address, error) {
	return _BscWOOPP.Contract.WooGuardian(&_BscWOOPP.CallOpts)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) WooGuardian() (common.Address, error) {
	return _BscWOOPP.Contract.WooGuardian(&_BscWOOPP.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWOOPP *BscWOOPPCaller) Wooracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscWOOPP.contract.Call(opts, &out, "wooracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWOOPP *BscWOOPPSession) Wooracle() (common.Address, error) {
	return _BscWOOPP.Contract.Wooracle(&_BscWOOPP.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_BscWOOPP *BscWOOPPCallerSession) Wooracle() (common.Address, error) {
	return _BscWOOPP.Contract.Wooracle(&_BscWOOPP.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_BscWOOPP *BscWOOPPTransactor) AddBaseToken(opts *bind.TransactOpts, baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "addBaseToken", baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_BscWOOPP *BscWOOPPSession) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.AddBaseToken(&_BscWOOPP.TransactOpts, baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.AddBaseToken(&_BscWOOPP.TransactOpts, baseToken, threshold, R)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWOOPP *BscWOOPPTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWOOPP *BscWOOPPSession) ClaimOwnership() (*types.Transaction, error) {
	return _BscWOOPP.Contract.ClaimOwnership(&_BscWOOPP.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BscWOOPP *BscWOOPPTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _BscWOOPP.Contract.ClaimOwnership(&_BscWOOPP.TransactOpts)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWOOPP *BscWOOPPTransactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWOOPP *BscWOOPPSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.InitOwner(&_BscWOOPP.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.InitOwner(&_BscWOOPP.TransactOpts, newOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWOOPP *BscWOOPPTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWOOPP *BscWOOPPSession) Pause() (*types.Transaction, error) {
	return _BscWOOPP.Contract.Pause(&_BscWOOPP.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BscWOOPP *BscWOOPPTransactorSession) Pause() (*types.Transaction, error) {
	return _BscWOOPP.Contract.Pause(&_BscWOOPP.TransactOpts)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWOOPP *BscWOOPPTransactor) RemoveBaseToken(opts *bind.TransactOpts, baseToken common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "removeBaseToken", baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWOOPP *BscWOOPPSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.RemoveBaseToken(&_BscWOOPP.TransactOpts, baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.RemoveBaseToken(&_BscWOOPP.TransactOpts, baseToken)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPTransactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SellBase(&_BscWOOPP.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_BscWOOPP *BscWOOPPTransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SellBase(&_BscWOOPP.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPTransactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SellQuote(&_BscWOOPP.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_BscWOOPP *BscWOOPPTransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SellQuote(&_BscWOOPP.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_BscWOOPP *BscWOOPPTransactor) SetFeeManager(opts *bind.TransactOpts, newFeeManager common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "setFeeManager", newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_BscWOOPP *BscWOOPPSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetFeeManager(&_BscWOOPP.TransactOpts, newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetFeeManager(&_BscWOOPP.TransactOpts, newFeeManager)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWOOPP *BscWOOPPTransactor) SetPairsInfo(opts *bind.TransactOpts, newPairsInfo string) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "setPairsInfo", newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWOOPP *BscWOOPPSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetPairsInfo(&_BscWOOPP.TransactOpts, newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetPairsInfo(&_BscWOOPP.TransactOpts, newPairsInfo)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWOOPP *BscWOOPPTransactor) SetStrategist(opts *bind.TransactOpts, strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "setStrategist", strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWOOPP *BscWOOPPSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetStrategist(&_BscWOOPP.TransactOpts, strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetStrategist(&_BscWOOPP.TransactOpts, strategist, flag)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWOOPP *BscWOOPPTransactor) SetWooGuardian(opts *bind.TransactOpts, newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "setWooGuardian", newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWOOPP *BscWOOPPSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetWooGuardian(&_BscWOOPP.TransactOpts, newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetWooGuardian(&_BscWOOPP.TransactOpts, newWooGuardian)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWOOPP *BscWOOPPTransactor) SetWooracle(opts *bind.TransactOpts, newWooracle common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "setWooracle", newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWOOPP *BscWOOPPSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetWooracle(&_BscWOOPP.TransactOpts, newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.SetWooracle(&_BscWOOPP.TransactOpts, newWooracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWOOPP *BscWOOPPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWOOPP *BscWOOPPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.TransferOwnership(&_BscWOOPP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.TransferOwnership(&_BscWOOPP.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_BscWOOPP *BscWOOPPTransactor) TuneParameters(opts *bind.TransactOpts, token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "tuneParameters", token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_BscWOOPP *BscWOOPPSession) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.TuneParameters(&_BscWOOPP.TransactOpts, token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.TuneParameters(&_BscWOOPP.TransactOpts, token, newThreshold, newR)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWOOPP *BscWOOPPTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWOOPP *BscWOOPPSession) Unpause() (*types.Transaction, error) {
	return _BscWOOPP.Contract.Unpause(&_BscWOOPP.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BscWOOPP *BscWOOPPTransactorSession) Unpause() (*types.Transaction, error) {
	return _BscWOOPP.Contract.Unpause(&_BscWOOPP.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWOOPP *BscWOOPPTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWOOPP *BscWOOPPSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.Withdraw(&_BscWOOPP.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BscWOOPP.Contract.Withdraw(&_BscWOOPP.TransactOpts, token, to, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_BscWOOPP *BscWOOPPTransactor) WithdrawAll(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "withdrawAll", token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_BscWOOPP *BscWOOPPSession) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.WithdrawAll(&_BscWOOPP.TransactOpts, token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.WithdrawAll(&_BscWOOPP.TransactOpts, token, to)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_BscWOOPP *BscWOOPPTransactor) WithdrawAllToOwner(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BscWOOPP.contract.Transact(opts, "withdrawAllToOwner", token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_BscWOOPP *BscWOOPPSession) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.WithdrawAllToOwner(&_BscWOOPP.TransactOpts, token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_BscWOOPP *BscWOOPPTransactorSession) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _BscWOOPP.Contract.WithdrawAllToOwner(&_BscWOOPP.TransactOpts, token)
}

// BscWOOPPFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the BscWOOPP contract.
type BscWOOPPFeeManagerUpdatedIterator struct {
	Event *BscWOOPPFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPFeeManagerUpdated)
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
		it.Event = new(BscWOOPPFeeManagerUpdated)
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
func (it *BscWOOPPFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPFeeManagerUpdated represents a FeeManagerUpdated event raised by the BscWOOPP contract.
type BscWOOPPFeeManagerUpdated struct {
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_BscWOOPP *BscWOOPPFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts, newFeeManager []common.Address) (*BscWOOPPFeeManagerUpdatedIterator, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPFeeManagerUpdatedIterator{contract: _BscWOOPP.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_BscWOOPP *BscWOOPPFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPFeeManagerUpdated, newFeeManager []common.Address) (event.Subscription, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPFeeManagerUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseFeeManagerUpdated(log types.Log) (*BscWOOPPFeeManagerUpdated, error) {
	event := new(BscWOOPPFeeManagerUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the BscWOOPP contract.
type BscWOOPPOwnershipTransferPreparedIterator struct {
	Event *BscWOOPPOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *BscWOOPPOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPOwnershipTransferPrepared)
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
		it.Event = new(BscWOOPPOwnershipTransferPrepared)
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
func (it *BscWOOPPOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the BscWOOPP contract.
type BscWOOPPOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_BscWOOPP *BscWOOPPFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscWOOPPOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPOwnershipTransferPreparedIterator{contract: _BscWOOPP.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_BscWOOPP *BscWOOPPFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *BscWOOPPOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPOwnershipTransferPrepared)
				if err := _BscWOOPP.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseOwnershipTransferPrepared(log types.Log) (*BscWOOPPOwnershipTransferPrepared, error) {
	event := new(BscWOOPPOwnershipTransferPrepared)
	if err := _BscWOOPP.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscWOOPP contract.
type BscWOOPPOwnershipTransferredIterator struct {
	Event *BscWOOPPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscWOOPPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPOwnershipTransferred)
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
		it.Event = new(BscWOOPPOwnershipTransferred)
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
func (it *BscWOOPPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPOwnershipTransferred represents a OwnershipTransferred event raised by the BscWOOPP contract.
type BscWOOPPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWOOPP *BscWOOPPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscWOOPPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPOwnershipTransferredIterator{contract: _BscWOOPP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscWOOPP *BscWOOPPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscWOOPPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPOwnershipTransferred)
				if err := _BscWOOPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseOwnershipTransferred(log types.Log) (*BscWOOPPOwnershipTransferred, error) {
	event := new(BscWOOPPOwnershipTransferred)
	if err := _BscWOOPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPParametersUpdatedIterator is returned from FilterParametersUpdated and is used to iterate over the raw logs and unpacked data for ParametersUpdated events raised by the BscWOOPP contract.
type BscWOOPPParametersUpdatedIterator struct {
	Event *BscWOOPPParametersUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPParametersUpdated)
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
		it.Event = new(BscWOOPPParametersUpdated)
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
func (it *BscWOOPPParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPParametersUpdated represents a ParametersUpdated event raised by the BscWOOPP contract.
type BscWOOPPParametersUpdated struct {
	BaseToken    common.Address
	NewThreshold *big.Int
	NewR         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParametersUpdated is a free log retrieval operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_BscWOOPP *BscWOOPPFilterer) FilterParametersUpdated(opts *bind.FilterOpts, baseToken []common.Address) (*BscWOOPPParametersUpdatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPParametersUpdatedIterator{contract: _BscWOOPP.contract, event: "ParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchParametersUpdated is a free log subscription operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_BscWOOPP *BscWOOPPFilterer) WatchParametersUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPParametersUpdated, baseToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPParametersUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseParametersUpdated(log types.Log) (*BscWOOPPParametersUpdated, error) {
	event := new(BscWOOPPParametersUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BscWOOPP contract.
type BscWOOPPPausedIterator struct {
	Event *BscWOOPPPaused // Event containing the contract specifics and raw log

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
func (it *BscWOOPPPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPPaused)
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
		it.Event = new(BscWOOPPPaused)
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
func (it *BscWOOPPPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPPaused represents a Paused event raised by the BscWOOPP contract.
type BscWOOPPPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BscWOOPP *BscWOOPPFilterer) FilterPaused(opts *bind.FilterOpts) (*BscWOOPPPausedIterator, error) {

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BscWOOPPPausedIterator{contract: _BscWOOPP.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BscWOOPP *BscWOOPPFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BscWOOPPPaused) (event.Subscription, error) {

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPPaused)
				if err := _BscWOOPP.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParsePaused(log types.Log) (*BscWOOPPPaused, error) {
	event := new(BscWOOPPPaused)
	if err := _BscWOOPP.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPRewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the BscWOOPP contract.
type BscWOOPPRewardManagerUpdatedIterator struct {
	Event *BscWOOPPRewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPRewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPRewardManagerUpdated)
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
		it.Event = new(BscWOOPPRewardManagerUpdated)
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
func (it *BscWOOPPRewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPRewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPRewardManagerUpdated represents a RewardManagerUpdated event raised by the BscWOOPP contract.
type BscWOOPPRewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_BscWOOPP *BscWOOPPFilterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*BscWOOPPRewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPRewardManagerUpdatedIterator{contract: _BscWOOPP.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_BscWOOPP *BscWOOPPFilterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPRewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPRewardManagerUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseRewardManagerUpdated(log types.Log) (*BscWOOPPRewardManagerUpdated, error) {
	event := new(BscWOOPPRewardManagerUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPStrategistUpdatedIterator is returned from FilterStrategistUpdated and is used to iterate over the raw logs and unpacked data for StrategistUpdated events raised by the BscWOOPP contract.
type BscWOOPPStrategistUpdatedIterator struct {
	Event *BscWOOPPStrategistUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPStrategistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPStrategistUpdated)
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
		it.Event = new(BscWOOPPStrategistUpdated)
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
func (it *BscWOOPPStrategistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPStrategistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPStrategistUpdated represents a StrategistUpdated event raised by the BscWOOPP contract.
type BscWOOPPStrategistUpdated struct {
	Strategist common.Address
	Flag       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategistUpdated is a free log retrieval operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_BscWOOPP *BscWOOPPFilterer) FilterStrategistUpdated(opts *bind.FilterOpts, strategist []common.Address) (*BscWOOPPStrategistUpdatedIterator, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPStrategistUpdatedIterator{contract: _BscWOOPP.contract, event: "StrategistUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategistUpdated is a free log subscription operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_BscWOOPP *BscWOOPPFilterer) WatchStrategistUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPStrategistUpdated, strategist []common.Address) (event.Subscription, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPStrategistUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseStrategistUpdated(log types.Log) (*BscWOOPPStrategistUpdated, error) {
	event := new(BscWOOPPStrategistUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BscWOOPP contract.
type BscWOOPPUnpausedIterator struct {
	Event *BscWOOPPUnpaused // Event containing the contract specifics and raw log

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
func (it *BscWOOPPUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPUnpaused)
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
		it.Event = new(BscWOOPPUnpaused)
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
func (it *BscWOOPPUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPUnpaused represents a Unpaused event raised by the BscWOOPP contract.
type BscWOOPPUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BscWOOPP *BscWOOPPFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BscWOOPPUnpausedIterator, error) {

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BscWOOPPUnpausedIterator{contract: _BscWOOPP.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BscWOOPP *BscWOOPPFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BscWOOPPUnpaused) (event.Subscription, error) {

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPUnpaused)
				if err := _BscWOOPP.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseUnpaused(log types.Log) (*BscWOOPPUnpaused, error) {
	event := new(BscWOOPPUnpaused)
	if err := _BscWOOPP.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BscWOOPP contract.
type BscWOOPPWithdrawIterator struct {
	Event *BscWOOPPWithdraw // Event containing the contract specifics and raw log

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
func (it *BscWOOPPWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPWithdraw)
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
		it.Event = new(BscWOOPPWithdraw)
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
func (it *BscWOOPPWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPWithdraw represents a Withdraw event raised by the BscWOOPP contract.
type BscWOOPPWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_BscWOOPP *BscWOOPPFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BscWOOPPWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPWithdrawIterator{contract: _BscWOOPP.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_BscWOOPP *BscWOOPPFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BscWOOPPWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPWithdraw)
				if err := _BscWOOPP.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseWithdraw(log types.Log) (*BscWOOPPWithdraw, error) {
	event := new(BscWOOPPWithdraw)
	if err := _BscWOOPP.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPWooGuardianUpdatedIterator is returned from FilterWooGuardianUpdated and is used to iterate over the raw logs and unpacked data for WooGuardianUpdated events raised by the BscWOOPP contract.
type BscWOOPPWooGuardianUpdatedIterator struct {
	Event *BscWOOPPWooGuardianUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPWooGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPWooGuardianUpdated)
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
		it.Event = new(BscWOOPPWooGuardianUpdated)
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
func (it *BscWOOPPWooGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPWooGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPWooGuardianUpdated represents a WooGuardianUpdated event raised by the BscWOOPP contract.
type BscWOOPPWooGuardianUpdated struct {
	NewWooGuardian common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWooGuardianUpdated is a free log retrieval operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_BscWOOPP *BscWOOPPFilterer) FilterWooGuardianUpdated(opts *bind.FilterOpts, newWooGuardian []common.Address) (*BscWOOPPWooGuardianUpdatedIterator, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPWooGuardianUpdatedIterator{contract: _BscWOOPP.contract, event: "WooGuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchWooGuardianUpdated is a free log subscription operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_BscWOOPP *BscWOOPPFilterer) WatchWooGuardianUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPWooGuardianUpdated, newWooGuardian []common.Address) (event.Subscription, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPWooGuardianUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseWooGuardianUpdated(log types.Log) (*BscWOOPPWooGuardianUpdated, error) {
	event := new(BscWOOPPWooGuardianUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPWooSwapIterator is returned from FilterWooSwap and is used to iterate over the raw logs and unpacked data for WooSwap events raised by the BscWOOPP contract.
type BscWOOPPWooSwapIterator struct {
	Event *BscWOOPPWooSwap // Event containing the contract specifics and raw log

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
func (it *BscWOOPPWooSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPWooSwap)
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
		it.Event = new(BscWOOPPWooSwap)
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
func (it *BscWOOPPWooSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPWooSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPWooSwap represents a WooSwap event raised by the BscWOOPP contract.
type BscWOOPPWooSwap struct {
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
func (_BscWOOPP *BscWOOPPFilterer) FilterWooSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*BscWOOPPWooSwapIterator, error) {

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

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPWooSwapIterator{contract: _BscWOOPP.contract, event: "WooSwap", logs: logs, sub: sub}, nil
}

// WatchWooSwap is a free log subscription operation binding the contract event 0x74ef34e2ea7c5d9f7b7ed44e97ad44b4303416c3a660c3fb5b3bdb95a1d6abd3.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_BscWOOPP *BscWOOPPFilterer) WatchWooSwap(opts *bind.WatchOpts, sink chan<- *BscWOOPPWooSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPWooSwap)
				if err := _BscWOOPP.contract.UnpackLog(event, "WooSwap", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseWooSwap(log types.Log) (*BscWOOPPWooSwap, error) {
	event := new(BscWOOPPWooSwap)
	if err := _BscWOOPP.contract.UnpackLog(event, "WooSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscWOOPPWooracleUpdatedIterator is returned from FilterWooracleUpdated and is used to iterate over the raw logs and unpacked data for WooracleUpdated events raised by the BscWOOPP contract.
type BscWOOPPWooracleUpdatedIterator struct {
	Event *BscWOOPPWooracleUpdated // Event containing the contract specifics and raw log

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
func (it *BscWOOPPWooracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscWOOPPWooracleUpdated)
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
		it.Event = new(BscWOOPPWooracleUpdated)
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
func (it *BscWOOPPWooracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscWOOPPWooracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscWOOPPWooracleUpdated represents a WooracleUpdated event raised by the BscWOOPP contract.
type BscWOOPPWooracleUpdated struct {
	NewWooracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWooracleUpdated is a free log retrieval operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_BscWOOPP *BscWOOPPFilterer) FilterWooracleUpdated(opts *bind.FilterOpts, newWooracle []common.Address) (*BscWOOPPWooracleUpdatedIterator, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _BscWOOPP.contract.FilterLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return &BscWOOPPWooracleUpdatedIterator{contract: _BscWOOPP.contract, event: "WooracleUpdated", logs: logs, sub: sub}, nil
}

// WatchWooracleUpdated is a free log subscription operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_BscWOOPP *BscWOOPPFilterer) WatchWooracleUpdated(opts *bind.WatchOpts, sink chan<- *BscWOOPPWooracleUpdated, newWooracle []common.Address) (event.Subscription, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _BscWOOPP.contract.WatchLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscWOOPPWooracleUpdated)
				if err := _BscWOOPP.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
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
func (_BscWOOPP *BscWOOPPFilterer) ParseWooracleUpdated(log types.Log) (*BscWOOPPWooracleUpdated, error) {
	event := new(BscWOOPPWooracleUpdated)
	if err := _BscWOOPP.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
