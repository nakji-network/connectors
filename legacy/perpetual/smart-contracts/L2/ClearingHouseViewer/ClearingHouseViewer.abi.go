// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ClearingHouseViewer

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

// ClearingHousePosition is an auto generated low-level Go binding around an user-defined struct.
type ClearingHousePosition struct {
	Size                                 SignedDecimalsignedDecimal
	Margin                               Decimaldecimal
	OpenNotional                         Decimaldecimal
	LastUpdatedCumulativePremiumFraction SignedDecimalsignedDecimal
	LiquidityHistoryIndex                *big.Int
	BlockNumber                          *big.Int
}

// Decimaldecimal is an auto generated low-level Go binding around an user-defined struct.
type Decimaldecimal struct {
	D *big.Int
}

// SignedDecimalsignedDecimal is an auto generated low-level Go binding around an user-defined struct.
type SignedDecimalsignedDecimal struct {
	D *big.Int
}

// ClearingHouseViewerABI is the input ABI used to generate the binding from.
const ClearingHouseViewerABI = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_clearingHouse\",\"internalType\":\"contractClearingHouse\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractClearingHouse\"}],\"name\":\"clearingHouse\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getMarginRatio\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"margin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getPersonalBalanceWithFundingPayment\",\"inputs\":[{\"type\":\"address\",\"name\":\"_quoteToken\",\"internalType\":\"contractIERC20\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"position\",\"internalType\":\"structClearingHouse.Position\",\"components\":[{\"type\":\"tuple\",\"name\":\"size\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"margin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"openNotional\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"lastUpdatedCumulativePremiumFraction\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"uint256\",\"name\":\"liquidityHistoryIndex\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\"}]}],\"name\":\"getPersonalPositionWithFundingPayment\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getUnrealizedPnl\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"},{\"type\":\"uint8\",\"name\":\"_pnlCalcOption\",\"internalType\":\"enumClearingHouse.PnlCalcOption\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isPositionNeedToBeMigrated\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]}]"

// ClearingHouseViewer is an auto generated Go binding around an Ethereum contract.
type ClearingHouseViewer struct {
	ClearingHouseViewerCaller     // Read-only binding to the contract
	ClearingHouseViewerTransactor // Write-only binding to the contract
	ClearingHouseViewerFilterer   // Log filterer for contract events
}

// ClearingHouseViewerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClearingHouseViewerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseViewerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClearingHouseViewerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseViewerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClearingHouseViewerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseViewerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClearingHouseViewerSession struct {
	Contract     *ClearingHouseViewer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClearingHouseViewerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClearingHouseViewerCallerSession struct {
	Contract *ClearingHouseViewerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ClearingHouseViewerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClearingHouseViewerTransactorSession struct {
	Contract     *ClearingHouseViewerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ClearingHouseViewerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClearingHouseViewerRaw struct {
	Contract *ClearingHouseViewer // Generic contract binding to access the raw methods on
}

// ClearingHouseViewerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClearingHouseViewerCallerRaw struct {
	Contract *ClearingHouseViewerCaller // Generic read-only contract binding to access the raw methods on
}

// ClearingHouseViewerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClearingHouseViewerTransactorRaw struct {
	Contract *ClearingHouseViewerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClearingHouseViewer creates a new instance of ClearingHouseViewer, bound to a specific deployed contract.
func NewClearingHouseViewer(address common.Address, backend bind.ContractBackend) (*ClearingHouseViewer, error) {
	contract, err := bindClearingHouseViewer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseViewer{ClearingHouseViewerCaller: ClearingHouseViewerCaller{contract: contract}, ClearingHouseViewerTransactor: ClearingHouseViewerTransactor{contract: contract}, ClearingHouseViewerFilterer: ClearingHouseViewerFilterer{contract: contract}}, nil
}

// NewClearingHouseViewerCaller creates a new read-only instance of ClearingHouseViewer, bound to a specific deployed contract.
func NewClearingHouseViewerCaller(address common.Address, caller bind.ContractCaller) (*ClearingHouseViewerCaller, error) {
	contract, err := bindClearingHouseViewer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseViewerCaller{contract: contract}, nil
}

// NewClearingHouseViewerTransactor creates a new write-only instance of ClearingHouseViewer, bound to a specific deployed contract.
func NewClearingHouseViewerTransactor(address common.Address, transactor bind.ContractTransactor) (*ClearingHouseViewerTransactor, error) {
	contract, err := bindClearingHouseViewer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseViewerTransactor{contract: contract}, nil
}

// NewClearingHouseViewerFilterer creates a new log filterer instance of ClearingHouseViewer, bound to a specific deployed contract.
func NewClearingHouseViewerFilterer(address common.Address, filterer bind.ContractFilterer) (*ClearingHouseViewerFilterer, error) {
	contract, err := bindClearingHouseViewer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseViewerFilterer{contract: contract}, nil
}

// bindClearingHouseViewer binds a generic wrapper to an already deployed contract.
func bindClearingHouseViewer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClearingHouseViewerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClearingHouseViewer *ClearingHouseViewerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClearingHouseViewer.Contract.ClearingHouseViewerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClearingHouseViewer *ClearingHouseViewerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouseViewer.Contract.ClearingHouseViewerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClearingHouseViewer *ClearingHouseViewerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClearingHouseViewer.Contract.ClearingHouseViewerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClearingHouseViewer *ClearingHouseViewerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClearingHouseViewer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClearingHouseViewer *ClearingHouseViewerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouseViewer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClearingHouseViewer *ClearingHouseViewerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClearingHouseViewer.Contract.contract.Transact(opts, method, params...)
}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_ClearingHouseViewer *ClearingHouseViewerCaller) ClearingHouse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "clearingHouse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_ClearingHouseViewer *ClearingHouseViewerSession) ClearingHouse() (common.Address, error) {
	return _ClearingHouseViewer.Contract.ClearingHouse(&_ClearingHouseViewer.CallOpts)
}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) ClearingHouse() (common.Address, error) {
	return _ClearingHouseViewer.Contract.ClearingHouse(&_ClearingHouseViewer.CallOpts)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerCaller) GetMarginRatio(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "getMarginRatio", _amm, _trader)

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerSession) GetMarginRatio(_amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouseViewer.Contract.GetMarginRatio(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) GetMarginRatio(_amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouseViewer.Contract.GetMarginRatio(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}

// GetPersonalBalanceWithFundingPayment is a free data retrieval call binding the contract method 0xbafe871c.
//
// Solidity: function getPersonalBalanceWithFundingPayment(address _quoteToken, address _trader) view returns((uint256) margin)
func (_ClearingHouseViewer *ClearingHouseViewerCaller) GetPersonalBalanceWithFundingPayment(opts *bind.CallOpts, _quoteToken common.Address, _trader common.Address) (Decimaldecimal, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "getPersonalBalanceWithFundingPayment", _quoteToken, _trader)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetPersonalBalanceWithFundingPayment is a free data retrieval call binding the contract method 0xbafe871c.
//
// Solidity: function getPersonalBalanceWithFundingPayment(address _quoteToken, address _trader) view returns((uint256) margin)
func (_ClearingHouseViewer *ClearingHouseViewerSession) GetPersonalBalanceWithFundingPayment(_quoteToken common.Address, _trader common.Address) (Decimaldecimal, error) {
	return _ClearingHouseViewer.Contract.GetPersonalBalanceWithFundingPayment(&_ClearingHouseViewer.CallOpts, _quoteToken, _trader)
}

// GetPersonalBalanceWithFundingPayment is a free data retrieval call binding the contract method 0xbafe871c.
//
// Solidity: function getPersonalBalanceWithFundingPayment(address _quoteToken, address _trader) view returns((uint256) margin)
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) GetPersonalBalanceWithFundingPayment(_quoteToken common.Address, _trader common.Address) (Decimaldecimal, error) {
	return _ClearingHouseViewer.Contract.GetPersonalBalanceWithFundingPayment(&_ClearingHouseViewer.CallOpts, _quoteToken, _trader)
}

// GetPersonalPositionWithFundingPayment is a free data retrieval call binding the contract method 0x16f6c558.
//
// Solidity: function getPersonalPositionWithFundingPayment(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouseViewer *ClearingHouseViewerCaller) GetPersonalPositionWithFundingPayment(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "getPersonalPositionWithFundingPayment", _amm, _trader)

	if err != nil {
		return *new(ClearingHousePosition), err
	}

	out0 := *abi.ConvertType(out[0], new(ClearingHousePosition)).(*ClearingHousePosition)

	return out0, err

}

// GetPersonalPositionWithFundingPayment is a free data retrieval call binding the contract method 0x16f6c558.
//
// Solidity: function getPersonalPositionWithFundingPayment(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouseViewer *ClearingHouseViewerSession) GetPersonalPositionWithFundingPayment(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouseViewer.Contract.GetPersonalPositionWithFundingPayment(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}

// GetPersonalPositionWithFundingPayment is a free data retrieval call binding the contract method 0x16f6c558.
//
// Solidity: function getPersonalPositionWithFundingPayment(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) GetPersonalPositionWithFundingPayment(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouseViewer.Contract.GetPersonalPositionWithFundingPayment(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}

// GetUnrealizedPnl is a free data retrieval call binding the contract method 0xa2420238.
//
// Solidity: function getUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerCaller) GetUnrealizedPnl(opts *bind.CallOpts, _amm common.Address, _trader common.Address, _pnlCalcOption uint8) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "getUnrealizedPnl", _amm, _trader, _pnlCalcOption)

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetUnrealizedPnl is a free data retrieval call binding the contract method 0xa2420238.
//
// Solidity: function getUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerSession) GetUnrealizedPnl(_amm common.Address, _trader common.Address, _pnlCalcOption uint8) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouseViewer.Contract.GetUnrealizedPnl(&_ClearingHouseViewer.CallOpts, _amm, _trader, _pnlCalcOption)
}

// GetUnrealizedPnl is a free data retrieval call binding the contract method 0xa2420238.
//
// Solidity: function getUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((int256))
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) GetUnrealizedPnl(_amm common.Address, _trader common.Address, _pnlCalcOption uint8) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouseViewer.Contract.GetUnrealizedPnl(&_ClearingHouseViewer.CallOpts, _amm, _trader, _pnlCalcOption)
}

// IsPositionNeedToBeMigrated is a free data retrieval call binding the contract method 0x7db3b72d.
//
// Solidity: function isPositionNeedToBeMigrated(address _amm, address _trader) view returns(bool)
func (_ClearingHouseViewer *ClearingHouseViewerCaller) IsPositionNeedToBeMigrated(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (bool, error) {
	var out []interface{}
	err := _ClearingHouseViewer.contract.Call(opts, &out, "isPositionNeedToBeMigrated", _amm, _trader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPositionNeedToBeMigrated is a free data retrieval call binding the contract method 0x7db3b72d.
//
// Solidity: function isPositionNeedToBeMigrated(address _amm, address _trader) view returns(bool)
func (_ClearingHouseViewer *ClearingHouseViewerSession) IsPositionNeedToBeMigrated(_amm common.Address, _trader common.Address) (bool, error) {
	return _ClearingHouseViewer.Contract.IsPositionNeedToBeMigrated(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}

// IsPositionNeedToBeMigrated is a free data retrieval call binding the contract method 0x7db3b72d.
//
// Solidity: function isPositionNeedToBeMigrated(address _amm, address _trader) view returns(bool)
func (_ClearingHouseViewer *ClearingHouseViewerCallerSession) IsPositionNeedToBeMigrated(_amm common.Address, _trader common.Address) (bool, error) {
	return _ClearingHouseViewer.Contract.IsPositionNeedToBeMigrated(&_ClearingHouseViewer.CallOpts, _amm, _trader)
}
