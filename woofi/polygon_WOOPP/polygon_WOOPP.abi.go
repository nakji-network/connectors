// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygon_WOOPP

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

// PolygonWOOPPMetaData contains all meta data concerning the PolygonWOOPP contract.
var PolygonWOOPPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newQuoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"ParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRewardManager\",\"type\":\"address\"}],\"name\":\"RewardManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"StrategistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"WooGuardianUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"WooracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractIWooFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStrategist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairsInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"removeBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPairsInfo\",\"type\":\"string\"}],\"name\":\"setPairsInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooGuardian\",\"type\":\"address\"}],\"name\":\"setWooGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWooracle\",\"type\":\"address\"}],\"name\":\"setWooracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"threshold\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastResetTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"R\",\"type\":\"uint64\"},{\"internalType\":\"uint112\",\"name\":\"target\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newR\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAllToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooGuardian\",\"outputs\":[{\"internalType\":\"contractIWooGuardian\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonWOOPPABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonWOOPPMetaData.ABI instead.
var PolygonWOOPPABI = PolygonWOOPPMetaData.ABI

// PolygonWOOPP is an auto generated Go binding around an Ethereum contract.
type PolygonWOOPP struct {
	PolygonWOOPPCaller     // Read-only binding to the contract
	PolygonWOOPPTransactor // Write-only binding to the contract
	PolygonWOOPPFilterer   // Log filterer for contract events
}

// PolygonWOOPPCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonWOOPPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonWOOPPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonWOOPPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonWOOPPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonWOOPPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonWOOPPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonWOOPPSession struct {
	Contract     *PolygonWOOPP     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonWOOPPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonWOOPPCallerSession struct {
	Contract *PolygonWOOPPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PolygonWOOPPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonWOOPPTransactorSession struct {
	Contract     *PolygonWOOPPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PolygonWOOPPRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonWOOPPRaw struct {
	Contract *PolygonWOOPP // Generic contract binding to access the raw methods on
}

// PolygonWOOPPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonWOOPPCallerRaw struct {
	Contract *PolygonWOOPPCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonWOOPPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonWOOPPTransactorRaw struct {
	Contract *PolygonWOOPPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonWOOPP creates a new instance of PolygonWOOPP, bound to a specific deployed contract.
func NewPolygonWOOPP(address common.Address, backend bind.ContractBackend) (*PolygonWOOPP, error) {
	contract, err := bindPolygonWOOPP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPP{PolygonWOOPPCaller: PolygonWOOPPCaller{contract: contract}, PolygonWOOPPTransactor: PolygonWOOPPTransactor{contract: contract}, PolygonWOOPPFilterer: PolygonWOOPPFilterer{contract: contract}}, nil
}

// NewPolygonWOOPPCaller creates a new read-only instance of PolygonWOOPP, bound to a specific deployed contract.
func NewPolygonWOOPPCaller(address common.Address, caller bind.ContractCaller) (*PolygonWOOPPCaller, error) {
	contract, err := bindPolygonWOOPP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPCaller{contract: contract}, nil
}

// NewPolygonWOOPPTransactor creates a new write-only instance of PolygonWOOPP, bound to a specific deployed contract.
func NewPolygonWOOPPTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonWOOPPTransactor, error) {
	contract, err := bindPolygonWOOPP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPTransactor{contract: contract}, nil
}

// NewPolygonWOOPPFilterer creates a new log filterer instance of PolygonWOOPP, bound to a specific deployed contract.
func NewPolygonWOOPPFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonWOOPPFilterer, error) {
	contract, err := bindPolygonWOOPP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPFilterer{contract: contract}, nil
}

// bindPolygonWOOPP binds a generic wrapper to an already deployed contract.
func bindPolygonWOOPP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolygonWOOPPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonWOOPP *PolygonWOOPPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonWOOPP.Contract.PolygonWOOPPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonWOOPP *PolygonWOOPPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.PolygonWOOPPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonWOOPP *PolygonWOOPPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.PolygonWOOPPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonWOOPP *PolygonWOOPPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonWOOPP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonWOOPP *PolygonWOOPPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonWOOPP *PolygonWOOPPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.contract.Transact(opts, method, params...)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) NEWOWNER() (common.Address, error) {
	return _PolygonWOOPP.Contract.NEWOWNER(&_PolygonWOOPP.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) NEWOWNER() (common.Address, error) {
	return _PolygonWOOPP.Contract.NEWOWNER(&_PolygonWOOPP.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) OWNER() (common.Address, error) {
	return _PolygonWOOPP.Contract.OWNER(&_PolygonWOOPP.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) OWNER() (common.Address, error) {
	return _PolygonWOOPP.Contract.OWNER(&_PolygonWOOPP.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) FeeManager() (common.Address, error) {
	return _PolygonWOOPP.Contract.FeeManager(&_PolygonWOOPP.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) FeeManager() (common.Address, error) {
	return _PolygonWOOPP.Contract.FeeManager(&_PolygonWOOPP.CallOpts)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPCaller) IsStrategist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "isStrategist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _PolygonWOOPP.Contract.IsStrategist(&_PolygonWOOPP.CallOpts, arg0)
}

// IsStrategist is a free data retrieval call binding the contract method 0x6734faee.
//
// Solidity: function isStrategist(address ) view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) IsStrategist(arg0 common.Address) (bool, error) {
	return _PolygonWOOPP.Contract.IsStrategist(&_PolygonWOOPP.CallOpts, arg0)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_PolygonWOOPP *PolygonWOOPPCaller) PairsInfo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "pairsInfo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_PolygonWOOPP *PolygonWOOPPSession) PairsInfo() (string, error) {
	return _PolygonWOOPP.Contract.PairsInfo(&_PolygonWOOPP.CallOpts)
}

// PairsInfo is a free data retrieval call binding the contract method 0x3ef31236.
//
// Solidity: function pairsInfo() view returns(string)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) PairsInfo() (string, error) {
	return _PolygonWOOPP.Contract.PairsInfo(&_PolygonWOOPP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPSession) Paused() (bool, error) {
	return _PolygonWOOPP.Contract.Paused(&_PolygonWOOPP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) Paused() (bool, error) {
	return _PolygonWOOPP.Contract.Paused(&_PolygonWOOPP.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_PolygonWOOPP *PolygonWOOPPCaller) PoolSize(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "poolSize", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_PolygonWOOPP *PolygonWOOPPSession) PoolSize(token common.Address) (*big.Int, error) {
	return _PolygonWOOPP.Contract.PoolSize(&_PolygonWOOPP.CallOpts, token)
}

// PoolSize is a free data retrieval call binding the contract method 0xfa75d160.
//
// Solidity: function poolSize(address token) view returns(uint256)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) PoolSize(token common.Address) (*big.Int, error) {
	return _PolygonWOOPP.Contract.PoolSize(&_PolygonWOOPP.CallOpts, token)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPCaller) QuerySellBase(opts *bind.CallOpts, baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "querySellBase", baseToken, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _PolygonWOOPP.Contract.QuerySellBase(&_PolygonWOOPP.CallOpts, baseToken, baseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address baseToken, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) QuerySellBase(baseToken common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _PolygonWOOPP.Contract.QuerySellBase(&_PolygonWOOPP.CallOpts, baseToken, baseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPCaller) QuerySellQuote(opts *bind.CallOpts, baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "querySellQuote", baseToken, quoteAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _PolygonWOOPP.Contract.QuerySellQuote(&_PolygonWOOPP.CallOpts, baseToken, quoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address baseToken, uint256 quoteAmount) view returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) QuerySellQuote(baseToken common.Address, quoteAmount *big.Int) (*big.Int, error) {
	return _PolygonWOOPP.Contract.QuerySellQuote(&_PolygonWOOPP.CallOpts, baseToken, quoteAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) QuoteToken() (common.Address, error) {
	return _PolygonWOOPP.Contract.QuoteToken(&_PolygonWOOPP.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) QuoteToken() (common.Address, error) {
	return _PolygonWOOPP.Contract.QuoteToken(&_PolygonWOOPP.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_PolygonWOOPP *PolygonWOOPPCaller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "tokenInfo", arg0)

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
func (_PolygonWOOPP *PolygonWOOPPSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _PolygonWOOPP.Contract.TokenInfo(&_PolygonWOOPP.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(uint112 reserve, uint112 threshold, uint32 lastResetTimestamp, uint64 R, uint112 target, bool isValid)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) TokenInfo(arg0 common.Address) (struct {
	Reserve            *big.Int
	Threshold          *big.Int
	LastResetTimestamp uint32
	R                  uint64
	Target             *big.Int
	IsValid            bool
}, error) {
	return _PolygonWOOPP.Contract.TokenInfo(&_PolygonWOOPP.CallOpts, arg0)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) WooGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "wooGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) WooGuardian() (common.Address, error) {
	return _PolygonWOOPP.Contract.WooGuardian(&_PolygonWOOPP.CallOpts)
}

// WooGuardian is a free data retrieval call binding the contract method 0x3313429d.
//
// Solidity: function wooGuardian() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) WooGuardian() (common.Address, error) {
	return _PolygonWOOPP.Contract.WooGuardian(&_PolygonWOOPP.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCaller) Wooracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonWOOPP.contract.Call(opts, &out, "wooracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPSession) Wooracle() (common.Address, error) {
	return _PolygonWOOPP.Contract.Wooracle(&_PolygonWOOPP.CallOpts)
}

// Wooracle is a free data retrieval call binding the contract method 0xbc8530f9.
//
// Solidity: function wooracle() view returns(address)
func (_PolygonWOOPP *PolygonWOOPPCallerSession) Wooracle() (common.Address, error) {
	return _PolygonWOOPP.Contract.Wooracle(&_PolygonWOOPP.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) AddBaseToken(opts *bind.TransactOpts, baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "addBaseToken", baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.AddBaseToken(&_PolygonWOOPP.TransactOpts, baseToken, threshold, R)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0xdb77c0a1.
//
// Solidity: function addBaseToken(address baseToken, uint256 threshold, uint256 R) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) AddBaseToken(baseToken common.Address, threshold *big.Int, R *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.AddBaseToken(&_PolygonWOOPP.TransactOpts, baseToken, threshold, R)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PolygonWOOPP *PolygonWOOPPSession) ClaimOwnership() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.ClaimOwnership(&_PolygonWOOPP.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.ClaimOwnership(&_PolygonWOOPP.TransactOpts)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.InitOwner(&_PolygonWOOPP.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.InitOwner(&_PolygonWOOPP.TransactOpts, newOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolygonWOOPP *PolygonWOOPPSession) Pause() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Pause(&_PolygonWOOPP.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) Pause() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Pause(&_PolygonWOOPP.TransactOpts)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) RemoveBaseToken(opts *bind.TransactOpts, baseToken common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "removeBaseToken", baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.RemoveBaseToken(&_PolygonWOOPP.TransactOpts, baseToken)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address baseToken) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) RemoveBaseToken(baseToken common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.RemoveBaseToken(&_PolygonWOOPP.TransactOpts, baseToken)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPTransactor) SellBase(opts *bind.TransactOpts, baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "sellBase", baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SellBase(&_PolygonWOOPP.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellBase is a paid mutator transaction binding the contract method 0x6846fb50.
//
// Solidity: function sellBase(address baseToken, uint256 baseAmount, uint256 minQuoteAmount, address to, address rebateTo) returns(uint256 quoteAmount)
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SellBase(baseToken common.Address, baseAmount *big.Int, minQuoteAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SellBase(&_PolygonWOOPP.TransactOpts, baseToken, baseAmount, minQuoteAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPTransactor) SellQuote(opts *bind.TransactOpts, baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "sellQuote", baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SellQuote(&_PolygonWOOPP.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SellQuote is a paid mutator transaction binding the contract method 0xf3287c2f.
//
// Solidity: function sellQuote(address baseToken, uint256 quoteAmount, uint256 minBaseAmount, address to, address rebateTo) returns(uint256 baseAmount)
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SellQuote(baseToken common.Address, quoteAmount *big.Int, minBaseAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SellQuote(&_PolygonWOOPP.TransactOpts, baseToken, quoteAmount, minBaseAmount, to, rebateTo)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) SetFeeManager(opts *bind.TransactOpts, newFeeManager common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "setFeeManager", newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetFeeManager(&_PolygonWOOPP.TransactOpts, newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetFeeManager(&_PolygonWOOPP.TransactOpts, newFeeManager)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) SetPairsInfo(opts *bind.TransactOpts, newPairsInfo string) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "setPairsInfo", newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetPairsInfo(&_PolygonWOOPP.TransactOpts, newPairsInfo)
}

// SetPairsInfo is a paid mutator transaction binding the contract method 0xbec6c1c1.
//
// Solidity: function setPairsInfo(string newPairsInfo) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SetPairsInfo(newPairsInfo string) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetPairsInfo(&_PolygonWOOPP.TransactOpts, newPairsInfo)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) SetStrategist(opts *bind.TransactOpts, strategist common.Address, flag bool) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "setStrategist", strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetStrategist(&_PolygonWOOPP.TransactOpts, strategist, flag)
}

// SetStrategist is a paid mutator transaction binding the contract method 0x4c341e13.
//
// Solidity: function setStrategist(address strategist, bool flag) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SetStrategist(strategist common.Address, flag bool) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetStrategist(&_PolygonWOOPP.TransactOpts, strategist, flag)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) SetWooGuardian(opts *bind.TransactOpts, newWooGuardian common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "setWooGuardian", newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetWooGuardian(&_PolygonWOOPP.TransactOpts, newWooGuardian)
}

// SetWooGuardian is a paid mutator transaction binding the contract method 0x301ed02d.
//
// Solidity: function setWooGuardian(address newWooGuardian) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SetWooGuardian(newWooGuardian common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetWooGuardian(&_PolygonWOOPP.TransactOpts, newWooGuardian)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) SetWooracle(opts *bind.TransactOpts, newWooracle common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "setWooracle", newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetWooracle(&_PolygonWOOPP.TransactOpts, newWooracle)
}

// SetWooracle is a paid mutator transaction binding the contract method 0xa1ae8490.
//
// Solidity: function setWooracle(address newWooracle) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) SetWooracle(newWooracle common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.SetWooracle(&_PolygonWOOPP.TransactOpts, newWooracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.TransferOwnership(&_PolygonWOOPP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.TransferOwnership(&_PolygonWOOPP.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) TuneParameters(opts *bind.TransactOpts, token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "tuneParameters", token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.TuneParameters(&_PolygonWOOPP.TransactOpts, token, newThreshold, newR)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x567b5d6d.
//
// Solidity: function tuneParameters(address token, uint256 newThreshold, uint256 newR) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) TuneParameters(token common.Address, newThreshold *big.Int, newR *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.TuneParameters(&_PolygonWOOPP.TransactOpts, token, newThreshold, newR)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolygonWOOPP *PolygonWOOPPSession) Unpause() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Unpause(&_PolygonWOOPP.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) Unpause() (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Unpause(&_PolygonWOOPP.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Withdraw(&_PolygonWOOPP.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.Withdraw(&_PolygonWOOPP.TransactOpts, token, to, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) WithdrawAll(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "withdrawAll", token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.WithdrawAll(&_PolygonWOOPP.TransactOpts, token, to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address token, address to) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) WithdrawAll(token common.Address, to common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.WithdrawAll(&_PolygonWOOPP.TransactOpts, token, to)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactor) WithdrawAllToOwner(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.contract.Transact(opts, "withdrawAllToOwner", token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_PolygonWOOPP *PolygonWOOPPSession) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.WithdrawAllToOwner(&_PolygonWOOPP.TransactOpts, token)
}

// WithdrawAllToOwner is a paid mutator transaction binding the contract method 0x48d9751e.
//
// Solidity: function withdrawAllToOwner(address token) returns()
func (_PolygonWOOPP *PolygonWOOPPTransactorSession) WithdrawAllToOwner(token common.Address) (*types.Transaction, error) {
	return _PolygonWOOPP.Contract.WithdrawAllToOwner(&_PolygonWOOPP.TransactOpts, token)
}

// PolygonWOOPPFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPFeeManagerUpdatedIterator struct {
	Event *PolygonWOOPPFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPFeeManagerUpdated)
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
		it.Event = new(PolygonWOOPPFeeManagerUpdated)
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
func (it *PolygonWOOPPFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPFeeManagerUpdated represents a FeeManagerUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPFeeManagerUpdated struct {
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts, newFeeManager []common.Address) (*PolygonWOOPPFeeManagerUpdatedIterator, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPFeeManagerUpdatedIterator{contract: _PolygonWOOPP.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address indexed newFeeManager)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPFeeManagerUpdated, newFeeManager []common.Address) (event.Subscription, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "FeeManagerUpdated", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPFeeManagerUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseFeeManagerUpdated(log types.Log) (*PolygonWOOPPFeeManagerUpdated, error) {
	event := new(PolygonWOOPPFeeManagerUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the PolygonWOOPP contract.
type PolygonWOOPPOwnershipTransferPreparedIterator struct {
	Event *PolygonWOOPPOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPOwnershipTransferPrepared)
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
		it.Event = new(PolygonWOOPPOwnershipTransferPrepared)
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
func (it *PolygonWOOPPOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the PolygonWOOPP contract.
type PolygonWOOPPOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonWOOPPOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPOwnershipTransferPreparedIterator{contract: _PolygonWOOPP.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPOwnershipTransferPrepared)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseOwnershipTransferPrepared(log types.Log) (*PolygonWOOPPOwnershipTransferPrepared, error) {
	event := new(PolygonWOOPPOwnershipTransferPrepared)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolygonWOOPP contract.
type PolygonWOOPPOwnershipTransferredIterator struct {
	Event *PolygonWOOPPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPOwnershipTransferred)
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
		it.Event = new(PolygonWOOPPOwnershipTransferred)
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
func (it *PolygonWOOPPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPOwnershipTransferred represents a OwnershipTransferred event raised by the PolygonWOOPP contract.
type PolygonWOOPPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonWOOPPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPOwnershipTransferredIterator{contract: _PolygonWOOPP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPOwnershipTransferred)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonWOOPPOwnershipTransferred, error) {
	event := new(PolygonWOOPPOwnershipTransferred)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPParametersUpdatedIterator is returned from FilterParametersUpdated and is used to iterate over the raw logs and unpacked data for ParametersUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPParametersUpdatedIterator struct {
	Event *PolygonWOOPPParametersUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPParametersUpdated)
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
		it.Event = new(PolygonWOOPPParametersUpdated)
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
func (it *PolygonWOOPPParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPParametersUpdated represents a ParametersUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPParametersUpdated struct {
	BaseToken    common.Address
	NewThreshold *big.Int
	NewR         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParametersUpdated is a free log retrieval operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterParametersUpdated(opts *bind.FilterOpts, baseToken []common.Address) (*PolygonWOOPPParametersUpdatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPParametersUpdatedIterator{contract: _PolygonWOOPP.contract, event: "ParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchParametersUpdated is a free log subscription operation binding the contract event 0x5a1f36141a0cb942bd02bfae3796688b4c89d39ca58a36b286c71492a9b18012.
//
// Solidity: event ParametersUpdated(address indexed baseToken, uint256 newThreshold, uint256 newR)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchParametersUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPParametersUpdated, baseToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "ParametersUpdated", baseTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPParametersUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseParametersUpdated(log types.Log) (*PolygonWOOPPParametersUpdated, error) {
	event := new(PolygonWOOPPParametersUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "ParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolygonWOOPP contract.
type PolygonWOOPPPausedIterator struct {
	Event *PolygonWOOPPPaused // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPPaused)
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
		it.Event = new(PolygonWOOPPPaused)
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
func (it *PolygonWOOPPPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPPaused represents a Paused event raised by the PolygonWOOPP contract.
type PolygonWOOPPPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterPaused(opts *bind.FilterOpts) (*PolygonWOOPPPausedIterator, error) {

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPPausedIterator{contract: _PolygonWOOPP.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPPaused) (event.Subscription, error) {

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPPaused)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParsePaused(log types.Log) (*PolygonWOOPPPaused, error) {
	event := new(PolygonWOOPPPaused)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPRewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPRewardManagerUpdatedIterator struct {
	Event *PolygonWOOPPRewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPRewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPRewardManagerUpdated)
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
		it.Event = new(PolygonWOOPPRewardManagerUpdated)
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
func (it *PolygonWOOPPRewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPRewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPRewardManagerUpdated represents a RewardManagerUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPRewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*PolygonWOOPPRewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPRewardManagerUpdatedIterator{contract: _PolygonWOOPP.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPRewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPRewardManagerUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseRewardManagerUpdated(log types.Log) (*PolygonWOOPPRewardManagerUpdated, error) {
	event := new(PolygonWOOPPRewardManagerUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPStrategistUpdatedIterator is returned from FilterStrategistUpdated and is used to iterate over the raw logs and unpacked data for StrategistUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPStrategistUpdatedIterator struct {
	Event *PolygonWOOPPStrategistUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPStrategistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPStrategistUpdated)
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
		it.Event = new(PolygonWOOPPStrategistUpdated)
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
func (it *PolygonWOOPPStrategistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPStrategistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPStrategistUpdated represents a StrategistUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPStrategistUpdated struct {
	Strategist common.Address
	Flag       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategistUpdated is a free log retrieval operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterStrategistUpdated(opts *bind.FilterOpts, strategist []common.Address) (*PolygonWOOPPStrategistUpdatedIterator, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPStrategistUpdatedIterator{contract: _PolygonWOOPP.contract, event: "StrategistUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategistUpdated is a free log subscription operation binding the contract event 0xaee0e1c61a3f9668cbb6c91cbe9352a6cbb5334687e9dcf76655d72a23b942ce.
//
// Solidity: event StrategistUpdated(address indexed strategist, bool flag)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchStrategistUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPStrategistUpdated, strategist []common.Address) (event.Subscription, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "StrategistUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPStrategistUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseStrategistUpdated(log types.Log) (*PolygonWOOPPStrategistUpdated, error) {
	event := new(PolygonWOOPPStrategistUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "StrategistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolygonWOOPP contract.
type PolygonWOOPPUnpausedIterator struct {
	Event *PolygonWOOPPUnpaused // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPUnpaused)
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
		it.Event = new(PolygonWOOPPUnpaused)
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
func (it *PolygonWOOPPUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPUnpaused represents a Unpaused event raised by the PolygonWOOPP contract.
type PolygonWOOPPUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolygonWOOPPUnpausedIterator, error) {

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPUnpausedIterator{contract: _PolygonWOOPP.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPUnpaused) (event.Subscription, error) {

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPUnpaused)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseUnpaused(log types.Log) (*PolygonWOOPPUnpaused, error) {
	event := new(PolygonWOOPPUnpaused)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PolygonWOOPP contract.
type PolygonWOOPPWithdrawIterator struct {
	Event *PolygonWOOPPWithdraw // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPWithdraw)
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
		it.Event = new(PolygonWOOPPWithdraw)
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
func (it *PolygonWOOPPWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPWithdraw represents a Withdraw event raised by the PolygonWOOPP contract.
type PolygonWOOPPWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*PolygonWOOPPWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPWithdrawIterator{contract: _PolygonWOOPP.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPWithdraw)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseWithdraw(log types.Log) (*PolygonWOOPPWithdraw, error) {
	event := new(PolygonWOOPPWithdraw)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPWooGuardianUpdatedIterator is returned from FilterWooGuardianUpdated and is used to iterate over the raw logs and unpacked data for WooGuardianUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPWooGuardianUpdatedIterator struct {
	Event *PolygonWOOPPWooGuardianUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPWooGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPWooGuardianUpdated)
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
		it.Event = new(PolygonWOOPPWooGuardianUpdated)
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
func (it *PolygonWOOPPWooGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPWooGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPWooGuardianUpdated represents a WooGuardianUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPWooGuardianUpdated struct {
	NewWooGuardian common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWooGuardianUpdated is a free log retrieval operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterWooGuardianUpdated(opts *bind.FilterOpts, newWooGuardian []common.Address) (*PolygonWOOPPWooGuardianUpdatedIterator, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPWooGuardianUpdatedIterator{contract: _PolygonWOOPP.contract, event: "WooGuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchWooGuardianUpdated is a free log subscription operation binding the contract event 0x95543a0f1d7d53325b1264bff29b0cb8704c48f0da6e578463e10a951caf2ebb.
//
// Solidity: event WooGuardianUpdated(address indexed newWooGuardian)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchWooGuardianUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPWooGuardianUpdated, newWooGuardian []common.Address) (event.Subscription, error) {

	var newWooGuardianRule []interface{}
	for _, newWooGuardianItem := range newWooGuardian {
		newWooGuardianRule = append(newWooGuardianRule, newWooGuardianItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "WooGuardianUpdated", newWooGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPWooGuardianUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseWooGuardianUpdated(log types.Log) (*PolygonWOOPPWooGuardianUpdated, error) {
	event := new(PolygonWOOPPWooGuardianUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "WooGuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPWooSwapIterator is returned from FilterWooSwap and is used to iterate over the raw logs and unpacked data for WooSwap events raised by the PolygonWOOPP contract.
type PolygonWOOPPWooSwapIterator struct {
	Event *PolygonWOOPPWooSwap // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPWooSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPWooSwap)
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
		it.Event = new(PolygonWOOPPWooSwap)
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
func (it *PolygonWOOPPWooSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPWooSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPWooSwap represents a WooSwap event raised by the PolygonWOOPP contract.
type PolygonWOOPPWooSwap struct {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterWooSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*PolygonWOOPPWooSwapIterator, error) {

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

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPWooSwapIterator{contract: _PolygonWOOPP.contract, event: "WooSwap", logs: logs, sub: sub}, nil
}

// WatchWooSwap is a free log subscription operation binding the contract event 0x74ef34e2ea7c5d9f7b7ed44e97ad44b4303416c3a660c3fb5b3bdb95a1d6abd3.
//
// Solidity: event WooSwap(address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchWooSwap(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPWooSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "WooSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPWooSwap)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "WooSwap", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseWooSwap(log types.Log) (*PolygonWOOPPWooSwap, error) {
	event := new(PolygonWOOPPWooSwap)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "WooSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonWOOPPWooracleUpdatedIterator is returned from FilterWooracleUpdated and is used to iterate over the raw logs and unpacked data for WooracleUpdated events raised by the PolygonWOOPP contract.
type PolygonWOOPPWooracleUpdatedIterator struct {
	Event *PolygonWOOPPWooracleUpdated // Event containing the contract specifics and raw log

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
func (it *PolygonWOOPPWooracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonWOOPPWooracleUpdated)
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
		it.Event = new(PolygonWOOPPWooracleUpdated)
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
func (it *PolygonWOOPPWooracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonWOOPPWooracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonWOOPPWooracleUpdated represents a WooracleUpdated event raised by the PolygonWOOPP contract.
type PolygonWOOPPWooracleUpdated struct {
	NewWooracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWooracleUpdated is a free log retrieval operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_PolygonWOOPP *PolygonWOOPPFilterer) FilterWooracleUpdated(opts *bind.FilterOpts, newWooracle []common.Address) (*PolygonWOOPPWooracleUpdatedIterator, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.FilterLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonWOOPPWooracleUpdatedIterator{contract: _PolygonWOOPP.contract, event: "WooracleUpdated", logs: logs, sub: sub}, nil
}

// WatchWooracleUpdated is a free log subscription operation binding the contract event 0x59ddfdd1bf7e3ea08a07e8aaa1fe2ce9c840fab69fe5ede6ea727a45eb42fc66.
//
// Solidity: event WooracleUpdated(address indexed newWooracle)
func (_PolygonWOOPP *PolygonWOOPPFilterer) WatchWooracleUpdated(opts *bind.WatchOpts, sink chan<- *PolygonWOOPPWooracleUpdated, newWooracle []common.Address) (event.Subscription, error) {

	var newWooracleRule []interface{}
	for _, newWooracleItem := range newWooracle {
		newWooracleRule = append(newWooracleRule, newWooracleItem)
	}

	logs, sub, err := _PolygonWOOPP.contract.WatchLogs(opts, "WooracleUpdated", newWooracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonWOOPPWooracleUpdated)
				if err := _PolygonWOOPP.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
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
func (_PolygonWOOPP *PolygonWOOPPFilterer) ParseWooracleUpdated(log types.Log) (*PolygonWOOPPWooracleUpdated, error) {
	event := new(PolygonWOOPPWooracleUpdated)
	if err := _PolygonWOOPP.contract.UnpackLog(event, "WooracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
