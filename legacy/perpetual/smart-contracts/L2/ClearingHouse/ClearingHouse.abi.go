// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ClearingHouse

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

// ClearingHouseABI is the input ABI used to generate the binding from.
const ClearingHouseABI = "[{\"type\":\"event\",\"name\":\"LiquidationFeeRatioChanged\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"liquidationFeeRatio\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarginChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"int256\",\"name\":\"amount\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"fundingPayment\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarginRatioChanged\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"marginRatio\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionAdjusted\",\"inputs\":[{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"trader\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"int256\",\"name\":\"newPositionSize\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"oldLiquidityIndex\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"newLiquidityIndex\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"trader\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"margin\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"positionNotional\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"exchangedPositionSize\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"positionSizeAfter\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"realizedPnl\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"unrealizedPnlAfter\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"badDebt\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"liquidationPenalty\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"spotPrice\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"fundingPayment\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionLiquidated\",\"inputs\":[{\"type\":\"address\",\"name\":\"trader\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"positionNotional\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"positionSize\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"liquidationFee\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"liquidator\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"badDebt\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionSettled\",\"inputs\":[{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"trader\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"valueTransferred\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferredPositionChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"referralCode\",\"internalType\":\"bytes32\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestrictionModeEntered\",\"inputs\":[{\"type\":\"address\",\"name\":\"amm\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addMargin\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"tuple\",\"name\":\"_addedMargin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"adjustPosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"candidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"closePosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"closePositionWithReferral\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"bytes32\",\"name\":\"_referralCode\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIMultiTokenRewardRecipient\"}],\"name\":\"feePool\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getLatestCumulativePremiumFraction\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getMarginRatio\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structClearingHouse.Position\",\"components\":[{\"type\":\"tuple\",\"name\":\"size\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"margin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"openNotional\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"lastUpdatedCumulativePremiumFraction\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"uint256\",\"name\":\"liquidityHistoryIndex\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\"}]}],\"name\":\"getPosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"positionNotional\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"unrealizedPnl\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getPositionNotionalAndUnrealizedPnl\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"},{\"type\":\"uint8\",\"name\":\"_pnlCalcOption\",\"internalType\":\"enumClearingHouse.PnlCalcOption\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"position\",\"internalType\":\"structClearingHouse.Position\",\"components\":[{\"type\":\"tuple\",\"name\":\"size\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"margin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"openNotional\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"lastUpdatedCumulativePremiumFraction\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"uint256\",\"name\":\"liquidityHistoryIndex\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\"}]}],\"name\":\"getUnadjustedPosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"initMarginRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_initMarginRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maintenanceMarginRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_liquidationFeeRatio\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_insuranceFund\",\"internalType\":\"contractIInsuranceFund\"},{\"type\":\"address\",\"name\":\"_trustedForwarder\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIInsuranceFund\"}],\"name\":\"insuranceFund\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isTrustedForwarder\",\"inputs\":[{\"type\":\"address\",\"name\":\"forwarder\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"liquidate\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"address\",\"name\":\"_trader\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"liquidationFeeRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"maintenanceMarginRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"openInterestNotionalMap\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"openPosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"uint8\",\"name\":\"_side\",\"internalType\":\"enumClearingHouse.Side\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_leverage\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_baseAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"openPositionWithReferral\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"uint8\",\"name\":\"_side\",\"internalType\":\"enumClearingHouse.Side\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_leverage\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_baseAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"bytes32\",\"name\":\"_referralCode\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"partialLiquidationRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"pause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"paused\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"payFunding\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeMargin\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"},{\"type\":\"tuple\",\"name\":\"_removedMargin\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLiquidationFeeRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_liquidationFeeRatio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setMaintenanceMarginRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_maintenanceMarginRatio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setPartialLiquidationRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_ratio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTollPool\",\"inputs\":[{\"type\":\"address\",\"name\":\"_feePool\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setWhitelist\",\"inputs\":[{\"type\":\"address\",\"name\":\"_whitelist\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"settlePosition\",\"inputs\":[{\"type\":\"address\",\"name\":\"_amm\",\"internalType\":\"contractIAmm\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"trustedForwarder\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unpause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateOwner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"versionRecipient\",\"inputs\":[]}]"

// ClearingHouse is an auto generated Go binding around an Ethereum contract.
type ClearingHouse struct {
	ClearingHouseCaller     // Read-only binding to the contract
	ClearingHouseTransactor // Write-only binding to the contract
	ClearingHouseFilterer   // Log filterer for contract events
}

// ClearingHouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClearingHouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClearingHouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClearingHouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearingHouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClearingHouseSession struct {
	Contract     *ClearingHouse    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClearingHouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClearingHouseCallerSession struct {
	Contract *ClearingHouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClearingHouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClearingHouseTransactorSession struct {
	Contract     *ClearingHouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClearingHouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClearingHouseRaw struct {
	Contract *ClearingHouse // Generic contract binding to access the raw methods on
}

// ClearingHouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClearingHouseCallerRaw struct {
	Contract *ClearingHouseCaller // Generic read-only contract binding to access the raw methods on
}

// ClearingHouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClearingHouseTransactorRaw struct {
	Contract *ClearingHouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClearingHouse creates a new instance of ClearingHouse, bound to a specific deployed contract.
func NewClearingHouse(address common.Address, backend bind.ContractBackend) (*ClearingHouse, error) {
	contract, err := bindClearingHouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClearingHouse{ClearingHouseCaller: ClearingHouseCaller{contract: contract}, ClearingHouseTransactor: ClearingHouseTransactor{contract: contract}, ClearingHouseFilterer: ClearingHouseFilterer{contract: contract}}, nil
}

// NewClearingHouseCaller creates a new read-only instance of ClearingHouse, bound to a specific deployed contract.
func NewClearingHouseCaller(address common.Address, caller bind.ContractCaller) (*ClearingHouseCaller, error) {
	contract, err := bindClearingHouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseCaller{contract: contract}, nil
}

// NewClearingHouseTransactor creates a new write-only instance of ClearingHouse, bound to a specific deployed contract.
func NewClearingHouseTransactor(address common.Address, transactor bind.ContractTransactor) (*ClearingHouseTransactor, error) {
	contract, err := bindClearingHouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseTransactor{contract: contract}, nil
}

// NewClearingHouseFilterer creates a new log filterer instance of ClearingHouse, bound to a specific deployed contract.
func NewClearingHouseFilterer(address common.Address, filterer bind.ContractFilterer) (*ClearingHouseFilterer, error) {
	contract, err := bindClearingHouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseFilterer{contract: contract}, nil
}

// bindClearingHouse binds a generic wrapper to an already deployed contract.
func bindClearingHouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClearingHouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClearingHouse *ClearingHouseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClearingHouse.Contract.ClearingHouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClearingHouse *ClearingHouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClearingHouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClearingHouse *ClearingHouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClearingHouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClearingHouse *ClearingHouseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClearingHouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClearingHouse *ClearingHouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClearingHouse *ClearingHouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClearingHouse.Contract.contract.Transact(opts, method, params...)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ClearingHouse *ClearingHouseCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ClearingHouse *ClearingHouseSession) Candidate() (common.Address, error) {
	return _ClearingHouse.Contract.Candidate(&_ClearingHouse.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ClearingHouse *ClearingHouseCallerSession) Candidate() (common.Address, error) {
	return _ClearingHouse.Contract.Candidate(&_ClearingHouse.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(address)
func (_ClearingHouse *ClearingHouseCaller) FeePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "feePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(address)
func (_ClearingHouse *ClearingHouseSession) FeePool() (common.Address, error) {
	return _ClearingHouse.Contract.FeePool(&_ClearingHouse.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(address)
func (_ClearingHouse *ClearingHouseCallerSession) FeePool() (common.Address, error) {
	return _ClearingHouse.Contract.FeePool(&_ClearingHouse.CallOpts)
}

// GetLatestCumulativePremiumFraction is a free data retrieval call binding the contract method 0x6891397b.
//
// Solidity: function getLatestCumulativePremiumFraction(address _amm) view returns((int256))
func (_ClearingHouse *ClearingHouseCaller) GetLatestCumulativePremiumFraction(opts *bind.CallOpts, _amm common.Address) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "getLatestCumulativePremiumFraction", _amm)

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetLatestCumulativePremiumFraction is a free data retrieval call binding the contract method 0x6891397b.
//
// Solidity: function getLatestCumulativePremiumFraction(address _amm) view returns((int256))
func (_ClearingHouse *ClearingHouseSession) GetLatestCumulativePremiumFraction(_amm common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouse.Contract.GetLatestCumulativePremiumFraction(&_ClearingHouse.CallOpts, _amm)
}

// GetLatestCumulativePremiumFraction is a free data retrieval call binding the contract method 0x6891397b.
//
// Solidity: function getLatestCumulativePremiumFraction(address _amm) view returns((int256))
func (_ClearingHouse *ClearingHouseCallerSession) GetLatestCumulativePremiumFraction(_amm common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouse.Contract.GetLatestCumulativePremiumFraction(&_ClearingHouse.CallOpts, _amm)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouse *ClearingHouseCaller) GetMarginRatio(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "getMarginRatio", _amm, _trader)

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouse *ClearingHouseSession) GetMarginRatio(_amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouse.Contract.GetMarginRatio(&_ClearingHouse.CallOpts, _amm, _trader)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0xea0a45f9.
//
// Solidity: function getMarginRatio(address _amm, address _trader) view returns((int256))
func (_ClearingHouse *ClearingHouseCallerSession) GetMarginRatio(_amm common.Address, _trader common.Address) (SignedDecimalsignedDecimal, error) {
	return _ClearingHouse.Contract.GetMarginRatio(&_ClearingHouse.CallOpts, _amm, _trader)
}

// GetPosition is a free data retrieval call binding the contract method 0xb33dc190.
//
// Solidity: function getPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256))
func (_ClearingHouse *ClearingHouseCaller) GetPosition(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "getPosition", _amm, _trader)

	if err != nil {
		return *new(ClearingHousePosition), err
	}

	out0 := *abi.ConvertType(out[0], new(ClearingHousePosition)).(*ClearingHousePosition)

	return out0, err

}

// GetPosition is a free data retrieval call binding the contract method 0xb33dc190.
//
// Solidity: function getPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256))
func (_ClearingHouse *ClearingHouseSession) GetPosition(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouse.Contract.GetPosition(&_ClearingHouse.CallOpts, _amm, _trader)
}

// GetPosition is a free data retrieval call binding the contract method 0xb33dc190.
//
// Solidity: function getPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256))
func (_ClearingHouse *ClearingHouseCallerSession) GetPosition(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouse.Contract.GetPosition(&_ClearingHouse.CallOpts, _amm, _trader)
}

// GetPositionNotionalAndUnrealizedPnl is a free data retrieval call binding the contract method 0x8bedf3bb.
//
// Solidity: function getPositionNotionalAndUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((uint256) positionNotional, (int256) unrealizedPnl)
func (_ClearingHouse *ClearingHouseCaller) GetPositionNotionalAndUnrealizedPnl(opts *bind.CallOpts, _amm common.Address, _trader common.Address, _pnlCalcOption uint8) (struct {
	PositionNotional Decimaldecimal
	UnrealizedPnl    SignedDecimalsignedDecimal
}, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "getPositionNotionalAndUnrealizedPnl", _amm, _trader, _pnlCalcOption)

	outstruct := new(struct {
		PositionNotional Decimaldecimal
		UnrealizedPnl    SignedDecimalsignedDecimal
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PositionNotional = *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)
	outstruct.UnrealizedPnl = *abi.ConvertType(out[1], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return *outstruct, err

}

// GetPositionNotionalAndUnrealizedPnl is a free data retrieval call binding the contract method 0x8bedf3bb.
//
// Solidity: function getPositionNotionalAndUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((uint256) positionNotional, (int256) unrealizedPnl)
func (_ClearingHouse *ClearingHouseSession) GetPositionNotionalAndUnrealizedPnl(_amm common.Address, _trader common.Address, _pnlCalcOption uint8) (struct {
	PositionNotional Decimaldecimal
	UnrealizedPnl    SignedDecimalsignedDecimal
}, error) {
	return _ClearingHouse.Contract.GetPositionNotionalAndUnrealizedPnl(&_ClearingHouse.CallOpts, _amm, _trader, _pnlCalcOption)
}

// GetPositionNotionalAndUnrealizedPnl is a free data retrieval call binding the contract method 0x8bedf3bb.
//
// Solidity: function getPositionNotionalAndUnrealizedPnl(address _amm, address _trader, uint8 _pnlCalcOption) view returns((uint256) positionNotional, (int256) unrealizedPnl)
func (_ClearingHouse *ClearingHouseCallerSession) GetPositionNotionalAndUnrealizedPnl(_amm common.Address, _trader common.Address, _pnlCalcOption uint8) (struct {
	PositionNotional Decimaldecimal
	UnrealizedPnl    SignedDecimalsignedDecimal
}, error) {
	return _ClearingHouse.Contract.GetPositionNotionalAndUnrealizedPnl(&_ClearingHouse.CallOpts, _amm, _trader, _pnlCalcOption)
}

// GetUnadjustedPosition is a free data retrieval call binding the contract method 0x0475db5a.
//
// Solidity: function getUnadjustedPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouse *ClearingHouseCaller) GetUnadjustedPosition(opts *bind.CallOpts, _amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "getUnadjustedPosition", _amm, _trader)

	if err != nil {
		return *new(ClearingHousePosition), err
	}

	out0 := *abi.ConvertType(out[0], new(ClearingHousePosition)).(*ClearingHousePosition)

	return out0, err

}

// GetUnadjustedPosition is a free data retrieval call binding the contract method 0x0475db5a.
//
// Solidity: function getUnadjustedPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouse *ClearingHouseSession) GetUnadjustedPosition(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouse.Contract.GetUnadjustedPosition(&_ClearingHouse.CallOpts, _amm, _trader)
}

// GetUnadjustedPosition is a free data retrieval call binding the contract method 0x0475db5a.
//
// Solidity: function getUnadjustedPosition(address _amm, address _trader) view returns(((int256),(uint256),(uint256),(int256),uint256,uint256) position)
func (_ClearingHouse *ClearingHouseCallerSession) GetUnadjustedPosition(_amm common.Address, _trader common.Address) (ClearingHousePosition, error) {
	return _ClearingHouse.Contract.GetUnadjustedPosition(&_ClearingHouse.CallOpts, _amm, _trader)
}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCaller) InitMarginRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "initMarginRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseSession) InitMarginRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.InitMarginRatio(&_ClearingHouse.CallOpts)
}

// InitMarginRatio is a free data retrieval call binding the contract method 0xe57d5636.
//
// Solidity: function initMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCallerSession) InitMarginRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.InitMarginRatio(&_ClearingHouse.CallOpts)
}

// InsuranceFund is a free data retrieval call binding the contract method 0xb7902303.
//
// Solidity: function insuranceFund() view returns(address)
func (_ClearingHouse *ClearingHouseCaller) InsuranceFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "insuranceFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InsuranceFund is a free data retrieval call binding the contract method 0xb7902303.
//
// Solidity: function insuranceFund() view returns(address)
func (_ClearingHouse *ClearingHouseSession) InsuranceFund() (common.Address, error) {
	return _ClearingHouse.Contract.InsuranceFund(&_ClearingHouse.CallOpts)
}

// InsuranceFund is a free data retrieval call binding the contract method 0xb7902303.
//
// Solidity: function insuranceFund() view returns(address)
func (_ClearingHouse *ClearingHouseCallerSession) InsuranceFund() (common.Address, error) {
	return _ClearingHouse.Contract.InsuranceFund(&_ClearingHouse.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ClearingHouse *ClearingHouseCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ClearingHouse *ClearingHouseSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ClearingHouse.Contract.IsTrustedForwarder(&_ClearingHouse.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ClearingHouse *ClearingHouseCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ClearingHouse.Contract.IsTrustedForwarder(&_ClearingHouse.CallOpts, forwarder)
}

// LiquidationFeeRatio is a free data retrieval call binding the contract method 0xcfe71103.
//
// Solidity: function liquidationFeeRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCaller) LiquidationFeeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "liquidationFeeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeRatio is a free data retrieval call binding the contract method 0xcfe71103.
//
// Solidity: function liquidationFeeRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseSession) LiquidationFeeRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.LiquidationFeeRatio(&_ClearingHouse.CallOpts)
}

// LiquidationFeeRatio is a free data retrieval call binding the contract method 0xcfe71103.
//
// Solidity: function liquidationFeeRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCallerSession) LiquidationFeeRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.LiquidationFeeRatio(&_ClearingHouse.CallOpts)
}

// MaintenanceMarginRatio is a free data retrieval call binding the contract method 0x83acb48a.
//
// Solidity: function maintenanceMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCaller) MaintenanceMarginRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "maintenanceMarginRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaintenanceMarginRatio is a free data retrieval call binding the contract method 0x83acb48a.
//
// Solidity: function maintenanceMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseSession) MaintenanceMarginRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.MaintenanceMarginRatio(&_ClearingHouse.CallOpts)
}

// MaintenanceMarginRatio is a free data retrieval call binding the contract method 0x83acb48a.
//
// Solidity: function maintenanceMarginRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCallerSession) MaintenanceMarginRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.MaintenanceMarginRatio(&_ClearingHouse.CallOpts)
}

// OpenInterestNotionalMap is a free data retrieval call binding the contract method 0xb0ac5f2f.
//
// Solidity: function openInterestNotionalMap(address ) view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCaller) OpenInterestNotionalMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "openInterestNotionalMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenInterestNotionalMap is a free data retrieval call binding the contract method 0xb0ac5f2f.
//
// Solidity: function openInterestNotionalMap(address ) view returns(uint256 d)
func (_ClearingHouse *ClearingHouseSession) OpenInterestNotionalMap(arg0 common.Address) (*big.Int, error) {
	return _ClearingHouse.Contract.OpenInterestNotionalMap(&_ClearingHouse.CallOpts, arg0)
}

// OpenInterestNotionalMap is a free data retrieval call binding the contract method 0xb0ac5f2f.
//
// Solidity: function openInterestNotionalMap(address ) view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCallerSession) OpenInterestNotionalMap(arg0 common.Address) (*big.Int, error) {
	return _ClearingHouse.Contract.OpenInterestNotionalMap(&_ClearingHouse.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClearingHouse *ClearingHouseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClearingHouse *ClearingHouseSession) Owner() (common.Address, error) {
	return _ClearingHouse.Contract.Owner(&_ClearingHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClearingHouse *ClearingHouseCallerSession) Owner() (common.Address, error) {
	return _ClearingHouse.Contract.Owner(&_ClearingHouse.CallOpts)
}

// PartialLiquidationRatio is a free data retrieval call binding the contract method 0x090f05c2.
//
// Solidity: function partialLiquidationRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCaller) PartialLiquidationRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "partialLiquidationRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PartialLiquidationRatio is a free data retrieval call binding the contract method 0x090f05c2.
//
// Solidity: function partialLiquidationRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseSession) PartialLiquidationRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.PartialLiquidationRatio(&_ClearingHouse.CallOpts)
}

// PartialLiquidationRatio is a free data retrieval call binding the contract method 0x090f05c2.
//
// Solidity: function partialLiquidationRatio() view returns(uint256 d)
func (_ClearingHouse *ClearingHouseCallerSession) PartialLiquidationRatio() (*big.Int, error) {
	return _ClearingHouse.Contract.PartialLiquidationRatio(&_ClearingHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClearingHouse *ClearingHouseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClearingHouse *ClearingHouseSession) Paused() (bool, error) {
	return _ClearingHouse.Contract.Paused(&_ClearingHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClearingHouse *ClearingHouseCallerSession) Paused() (bool, error) {
	return _ClearingHouse.Contract.Paused(&_ClearingHouse.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_ClearingHouse *ClearingHouseCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_ClearingHouse *ClearingHouseSession) TrustedForwarder() (common.Address, error) {
	return _ClearingHouse.Contract.TrustedForwarder(&_ClearingHouse.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_ClearingHouse *ClearingHouseCallerSession) TrustedForwarder() (common.Address, error) {
	return _ClearingHouse.Contract.TrustedForwarder(&_ClearingHouse.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_ClearingHouse *ClearingHouseCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClearingHouse.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_ClearingHouse *ClearingHouseSession) VersionRecipient() (string, error) {
	return _ClearingHouse.Contract.VersionRecipient(&_ClearingHouse.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_ClearingHouse *ClearingHouseCallerSession) VersionRecipient() (string, error) {
	return _ClearingHouse.Contract.VersionRecipient(&_ClearingHouse.CallOpts)
}

// AddMargin is a paid mutator transaction binding the contract method 0xac06a96e.
//
// Solidity: function addMargin(address _amm, (uint256) _addedMargin) returns()
func (_ClearingHouse *ClearingHouseTransactor) AddMargin(opts *bind.TransactOpts, _amm common.Address, _addedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "addMargin", _amm, _addedMargin)
}

// AddMargin is a paid mutator transaction binding the contract method 0xac06a96e.
//
// Solidity: function addMargin(address _amm, (uint256) _addedMargin) returns()
func (_ClearingHouse *ClearingHouseSession) AddMargin(_amm common.Address, _addedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.AddMargin(&_ClearingHouse.TransactOpts, _amm, _addedMargin)
}

// AddMargin is a paid mutator transaction binding the contract method 0xac06a96e.
//
// Solidity: function addMargin(address _amm, (uint256) _addedMargin) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) AddMargin(_amm common.Address, _addedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.AddMargin(&_ClearingHouse.TransactOpts, _amm, _addedMargin)
}

// AdjustPosition is a paid mutator transaction binding the contract method 0x2db33aa1.
//
// Solidity: function adjustPosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactor) AdjustPosition(opts *bind.TransactOpts, _amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "adjustPosition", _amm)
}

// AdjustPosition is a paid mutator transaction binding the contract method 0x2db33aa1.
//
// Solidity: function adjustPosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseSession) AdjustPosition(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.AdjustPosition(&_ClearingHouse.TransactOpts, _amm)
}

// AdjustPosition is a paid mutator transaction binding the contract method 0x2db33aa1.
//
// Solidity: function adjustPosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) AdjustPosition(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.AdjustPosition(&_ClearingHouse.TransactOpts, _amm)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xa8c1b0bc.
//
// Solidity: function closePosition(address _amm, (uint256) _quoteAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseTransactor) ClosePosition(opts *bind.TransactOpts, _amm common.Address, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "closePosition", _amm, _quoteAssetAmountLimit)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xa8c1b0bc.
//
// Solidity: function closePosition(address _amm, (uint256) _quoteAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseSession) ClosePosition(_amm common.Address, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClosePosition(&_ClearingHouse.TransactOpts, _amm, _quoteAssetAmountLimit)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xa8c1b0bc.
//
// Solidity: function closePosition(address _amm, (uint256) _quoteAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) ClosePosition(_amm common.Address, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClosePosition(&_ClearingHouse.TransactOpts, _amm, _quoteAssetAmountLimit)
}

// ClosePositionWithReferral is a paid mutator transaction binding the contract method 0x3e810714.
//
// Solidity: function closePositionWithReferral(address _amm, (uint256) _quoteAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseTransactor) ClosePositionWithReferral(opts *bind.TransactOpts, _amm common.Address, _quoteAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "closePositionWithReferral", _amm, _quoteAssetAmountLimit, _referralCode)
}

// ClosePositionWithReferral is a paid mutator transaction binding the contract method 0x3e810714.
//
// Solidity: function closePositionWithReferral(address _amm, (uint256) _quoteAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseSession) ClosePositionWithReferral(_amm common.Address, _quoteAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClosePositionWithReferral(&_ClearingHouse.TransactOpts, _amm, _quoteAssetAmountLimit, _referralCode)
}

// ClosePositionWithReferral is a paid mutator transaction binding the contract method 0x3e810714.
//
// Solidity: function closePositionWithReferral(address _amm, (uint256) _quoteAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) ClosePositionWithReferral(_amm common.Address, _quoteAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.Contract.ClosePositionWithReferral(&_ClearingHouse.TransactOpts, _amm, _quoteAssetAmountLimit, _referralCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xaab954d0.
//
// Solidity: function initialize(uint256 _initMarginRatio, uint256 _maintenanceMarginRatio, uint256 _liquidationFeeRatio, address _insuranceFund, address _trustedForwarder) returns()
func (_ClearingHouse *ClearingHouseTransactor) Initialize(opts *bind.TransactOpts, _initMarginRatio *big.Int, _maintenanceMarginRatio *big.Int, _liquidationFeeRatio *big.Int, _insuranceFund common.Address, _trustedForwarder common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "initialize", _initMarginRatio, _maintenanceMarginRatio, _liquidationFeeRatio, _insuranceFund, _trustedForwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xaab954d0.
//
// Solidity: function initialize(uint256 _initMarginRatio, uint256 _maintenanceMarginRatio, uint256 _liquidationFeeRatio, address _insuranceFund, address _trustedForwarder) returns()
func (_ClearingHouse *ClearingHouseSession) Initialize(_initMarginRatio *big.Int, _maintenanceMarginRatio *big.Int, _liquidationFeeRatio *big.Int, _insuranceFund common.Address, _trustedForwarder common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.Initialize(&_ClearingHouse.TransactOpts, _initMarginRatio, _maintenanceMarginRatio, _liquidationFeeRatio, _insuranceFund, _trustedForwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xaab954d0.
//
// Solidity: function initialize(uint256 _initMarginRatio, uint256 _maintenanceMarginRatio, uint256 _liquidationFeeRatio, address _insuranceFund, address _trustedForwarder) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) Initialize(_initMarginRatio *big.Int, _maintenanceMarginRatio *big.Int, _liquidationFeeRatio *big.Int, _insuranceFund common.Address, _trustedForwarder common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.Initialize(&_ClearingHouse.TransactOpts, _initMarginRatio, _maintenanceMarginRatio, _liquidationFeeRatio, _insuranceFund, _trustedForwarder)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address _amm, address _trader) returns()
func (_ClearingHouse *ClearingHouseTransactor) Liquidate(opts *bind.TransactOpts, _amm common.Address, _trader common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "liquidate", _amm, _trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address _amm, address _trader) returns()
func (_ClearingHouse *ClearingHouseSession) Liquidate(_amm common.Address, _trader common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.Liquidate(&_ClearingHouse.TransactOpts, _amm, _trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address _amm, address _trader) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) Liquidate(_amm common.Address, _trader common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.Liquidate(&_ClearingHouse.TransactOpts, _amm, _trader)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x893d242d.
//
// Solidity: function openPosition(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseTransactor) OpenPosition(opts *bind.TransactOpts, _amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "openPosition", _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x893d242d.
//
// Solidity: function openPosition(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseSession) OpenPosition(_amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.OpenPosition(&_ClearingHouse.TransactOpts, _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x893d242d.
//
// Solidity: function openPosition(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) OpenPosition(_amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.OpenPosition(&_ClearingHouse.TransactOpts, _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit)
}

// OpenPositionWithReferral is a paid mutator transaction binding the contract method 0x5559c775.
//
// Solidity: function openPositionWithReferral(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseTransactor) OpenPositionWithReferral(opts *bind.TransactOpts, _amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "openPositionWithReferral", _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit, _referralCode)
}

// OpenPositionWithReferral is a paid mutator transaction binding the contract method 0x5559c775.
//
// Solidity: function openPositionWithReferral(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseSession) OpenPositionWithReferral(_amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.Contract.OpenPositionWithReferral(&_ClearingHouse.TransactOpts, _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit, _referralCode)
}

// OpenPositionWithReferral is a paid mutator transaction binding the contract method 0x5559c775.
//
// Solidity: function openPositionWithReferral(address _amm, uint8 _side, (uint256) _quoteAssetAmount, (uint256) _leverage, (uint256) _baseAssetAmountLimit, bytes32 _referralCode) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) OpenPositionWithReferral(_amm common.Address, _side uint8, _quoteAssetAmount Decimaldecimal, _leverage Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _referralCode [32]byte) (*types.Transaction, error) {
	return _ClearingHouse.Contract.OpenPositionWithReferral(&_ClearingHouse.TransactOpts, _amm, _side, _quoteAssetAmount, _leverage, _baseAssetAmountLimit, _referralCode)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClearingHouse *ClearingHouseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClearingHouse *ClearingHouseSession) Pause() (*types.Transaction, error) {
	return _ClearingHouse.Contract.Pause(&_ClearingHouse.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClearingHouse *ClearingHouseTransactorSession) Pause() (*types.Transaction, error) {
	return _ClearingHouse.Contract.Pause(&_ClearingHouse.TransactOpts)
}

// PayFunding is a paid mutator transaction binding the contract method 0x3e09fa10.
//
// Solidity: function payFunding(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactor) PayFunding(opts *bind.TransactOpts, _amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "payFunding", _amm)
}

// PayFunding is a paid mutator transaction binding the contract method 0x3e09fa10.
//
// Solidity: function payFunding(address _amm) returns()
func (_ClearingHouse *ClearingHouseSession) PayFunding(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.PayFunding(&_ClearingHouse.TransactOpts, _amm)
}

// PayFunding is a paid mutator transaction binding the contract method 0x3e09fa10.
//
// Solidity: function payFunding(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) PayFunding(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.PayFunding(&_ClearingHouse.TransactOpts, _amm)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0x30e2ae65.
//
// Solidity: function removeMargin(address _amm, (uint256) _removedMargin) returns()
func (_ClearingHouse *ClearingHouseTransactor) RemoveMargin(opts *bind.TransactOpts, _amm common.Address, _removedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "removeMargin", _amm, _removedMargin)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0x30e2ae65.
//
// Solidity: function removeMargin(address _amm, (uint256) _removedMargin) returns()
func (_ClearingHouse *ClearingHouseSession) RemoveMargin(_amm common.Address, _removedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.RemoveMargin(&_ClearingHouse.TransactOpts, _amm, _removedMargin)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0x30e2ae65.
//
// Solidity: function removeMargin(address _amm, (uint256) _removedMargin) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) RemoveMargin(_amm common.Address, _removedMargin Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.RemoveMargin(&_ClearingHouse.TransactOpts, _amm, _removedMargin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClearingHouse *ClearingHouseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClearingHouse *ClearingHouseSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClearingHouse.Contract.RenounceOwnership(&_ClearingHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClearingHouse *ClearingHouseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClearingHouse.Contract.RenounceOwnership(&_ClearingHouse.TransactOpts)
}

// SetLiquidationFeeRatio is a paid mutator transaction binding the contract method 0x203b3788.
//
// Solidity: function setLiquidationFeeRatio((uint256) _liquidationFeeRatio) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetLiquidationFeeRatio(opts *bind.TransactOpts, _liquidationFeeRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setLiquidationFeeRatio", _liquidationFeeRatio)
}

// SetLiquidationFeeRatio is a paid mutator transaction binding the contract method 0x203b3788.
//
// Solidity: function setLiquidationFeeRatio((uint256) _liquidationFeeRatio) returns()
func (_ClearingHouse *ClearingHouseSession) SetLiquidationFeeRatio(_liquidationFeeRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetLiquidationFeeRatio(&_ClearingHouse.TransactOpts, _liquidationFeeRatio)
}

// SetLiquidationFeeRatio is a paid mutator transaction binding the contract method 0x203b3788.
//
// Solidity: function setLiquidationFeeRatio((uint256) _liquidationFeeRatio) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetLiquidationFeeRatio(_liquidationFeeRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetLiquidationFeeRatio(&_ClearingHouse.TransactOpts, _liquidationFeeRatio)
}

// SetMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x6c127446.
//
// Solidity: function setMaintenanceMarginRatio((uint256) _maintenanceMarginRatio) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetMaintenanceMarginRatio(opts *bind.TransactOpts, _maintenanceMarginRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setMaintenanceMarginRatio", _maintenanceMarginRatio)
}

// SetMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x6c127446.
//
// Solidity: function setMaintenanceMarginRatio((uint256) _maintenanceMarginRatio) returns()
func (_ClearingHouse *ClearingHouseSession) SetMaintenanceMarginRatio(_maintenanceMarginRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetMaintenanceMarginRatio(&_ClearingHouse.TransactOpts, _maintenanceMarginRatio)
}

// SetMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x6c127446.
//
// Solidity: function setMaintenanceMarginRatio((uint256) _maintenanceMarginRatio) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetMaintenanceMarginRatio(_maintenanceMarginRatio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetMaintenanceMarginRatio(&_ClearingHouse.TransactOpts, _maintenanceMarginRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ClearingHouse *ClearingHouseSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetOwner(&_ClearingHouse.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetOwner(&_ClearingHouse.TransactOpts, newOwner)
}

// SetPartialLiquidationRatio is a paid mutator transaction binding the contract method 0x862c04e3.
//
// Solidity: function setPartialLiquidationRatio((uint256) _ratio) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetPartialLiquidationRatio(opts *bind.TransactOpts, _ratio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setPartialLiquidationRatio", _ratio)
}

// SetPartialLiquidationRatio is a paid mutator transaction binding the contract method 0x862c04e3.
//
// Solidity: function setPartialLiquidationRatio((uint256) _ratio) returns()
func (_ClearingHouse *ClearingHouseSession) SetPartialLiquidationRatio(_ratio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetPartialLiquidationRatio(&_ClearingHouse.TransactOpts, _ratio)
}

// SetPartialLiquidationRatio is a paid mutator transaction binding the contract method 0x862c04e3.
//
// Solidity: function setPartialLiquidationRatio((uint256) _ratio) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetPartialLiquidationRatio(_ratio Decimaldecimal) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetPartialLiquidationRatio(&_ClearingHouse.TransactOpts, _ratio)
}

// SetTollPool is a paid mutator transaction binding the contract method 0x7b25e922.
//
// Solidity: function setTollPool(address _feePool) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetTollPool(opts *bind.TransactOpts, _feePool common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setTollPool", _feePool)
}

// SetTollPool is a paid mutator transaction binding the contract method 0x7b25e922.
//
// Solidity: function setTollPool(address _feePool) returns()
func (_ClearingHouse *ClearingHouseSession) SetTollPool(_feePool common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetTollPool(&_ClearingHouse.TransactOpts, _feePool)
}

// SetTollPool is a paid mutator transaction binding the contract method 0x7b25e922.
//
// Solidity: function setTollPool(address _feePool) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetTollPool(_feePool common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetTollPool(&_ClearingHouse.TransactOpts, _feePool)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_ClearingHouse *ClearingHouseTransactor) SetWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "setWhitelist", _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_ClearingHouse *ClearingHouseSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetWhitelist(&_ClearingHouse.TransactOpts, _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SetWhitelist(&_ClearingHouse.TransactOpts, _whitelist)
}

// SettlePosition is a paid mutator transaction binding the contract method 0x36405257.
//
// Solidity: function settlePosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactor) SettlePosition(opts *bind.TransactOpts, _amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "settlePosition", _amm)
}

// SettlePosition is a paid mutator transaction binding the contract method 0x36405257.
//
// Solidity: function settlePosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseSession) SettlePosition(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SettlePosition(&_ClearingHouse.TransactOpts, _amm)
}

// SettlePosition is a paid mutator transaction binding the contract method 0x36405257.
//
// Solidity: function settlePosition(address _amm) returns()
func (_ClearingHouse *ClearingHouseTransactorSession) SettlePosition(_amm common.Address) (*types.Transaction, error) {
	return _ClearingHouse.Contract.SettlePosition(&_ClearingHouse.TransactOpts, _amm)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClearingHouse *ClearingHouseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClearingHouse *ClearingHouseSession) Unpause() (*types.Transaction, error) {
	return _ClearingHouse.Contract.Unpause(&_ClearingHouse.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClearingHouse *ClearingHouseTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClearingHouse.Contract.Unpause(&_ClearingHouse.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ClearingHouse *ClearingHouseTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClearingHouse.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ClearingHouse *ClearingHouseSession) UpdateOwner() (*types.Transaction, error) {
	return _ClearingHouse.Contract.UpdateOwner(&_ClearingHouse.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ClearingHouse *ClearingHouseTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _ClearingHouse.Contract.UpdateOwner(&_ClearingHouse.TransactOpts)
}

// ClearingHouseLiquidationFeeRatioChangedIterator is returned from FilterLiquidationFeeRatioChanged and is used to iterate over the raw logs and unpacked data for LiquidationFeeRatioChanged events raised by the ClearingHouse contract.
type ClearingHouseLiquidationFeeRatioChangedIterator struct {
	Event *ClearingHouseLiquidationFeeRatioChanged // Event containing the contract specifics and raw log

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
func (it *ClearingHouseLiquidationFeeRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseLiquidationFeeRatioChanged)
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
		it.Event = new(ClearingHouseLiquidationFeeRatioChanged)
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
func (it *ClearingHouseLiquidationFeeRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseLiquidationFeeRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseLiquidationFeeRatioChanged represents a LiquidationFeeRatioChanged event raised by the ClearingHouse contract.
type ClearingHouseLiquidationFeeRatioChanged struct {
	LiquidationFeeRatio *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLiquidationFeeRatioChanged is a free log retrieval operation binding the contract event 0x6f825f18b787836f7cf974799b160538ab237aa0d14af3611fa69683a40c6d22.
//
// Solidity: event LiquidationFeeRatioChanged(uint256 liquidationFeeRatio)
func (_ClearingHouse *ClearingHouseFilterer) FilterLiquidationFeeRatioChanged(opts *bind.FilterOpts) (*ClearingHouseLiquidationFeeRatioChangedIterator, error) {

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "LiquidationFeeRatioChanged")
	if err != nil {
		return nil, err
	}
	return &ClearingHouseLiquidationFeeRatioChangedIterator{contract: _ClearingHouse.contract, event: "LiquidationFeeRatioChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidationFeeRatioChanged is a free log subscription operation binding the contract event 0x6f825f18b787836f7cf974799b160538ab237aa0d14af3611fa69683a40c6d22.
//
// Solidity: event LiquidationFeeRatioChanged(uint256 liquidationFeeRatio)
func (_ClearingHouse *ClearingHouseFilterer) WatchLiquidationFeeRatioChanged(opts *bind.WatchOpts, sink chan<- *ClearingHouseLiquidationFeeRatioChanged) (event.Subscription, error) {

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "LiquidationFeeRatioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseLiquidationFeeRatioChanged)
				if err := _ClearingHouse.contract.UnpackLog(event, "LiquidationFeeRatioChanged", log); err != nil {
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

// ParseLiquidationFeeRatioChanged is a log parse operation binding the contract event 0x6f825f18b787836f7cf974799b160538ab237aa0d14af3611fa69683a40c6d22.
//
// Solidity: event LiquidationFeeRatioChanged(uint256 liquidationFeeRatio)
func (_ClearingHouse *ClearingHouseFilterer) ParseLiquidationFeeRatioChanged(log types.Log) (*ClearingHouseLiquidationFeeRatioChanged, error) {
	event := new(ClearingHouseLiquidationFeeRatioChanged)
	if err := _ClearingHouse.contract.UnpackLog(event, "LiquidationFeeRatioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseMarginChangedIterator is returned from FilterMarginChanged and is used to iterate over the raw logs and unpacked data for MarginChanged events raised by the ClearingHouse contract.
type ClearingHouseMarginChangedIterator struct {
	Event *ClearingHouseMarginChanged // Event containing the contract specifics and raw log

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
func (it *ClearingHouseMarginChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseMarginChanged)
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
		it.Event = new(ClearingHouseMarginChanged)
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
func (it *ClearingHouseMarginChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseMarginChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseMarginChanged represents a MarginChanged event raised by the ClearingHouse contract.
type ClearingHouseMarginChanged struct {
	Sender         common.Address
	Amm            common.Address
	Amount         *big.Int
	FundingPayment *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarginChanged is a free log retrieval operation binding the contract event 0x1c822fe92555a7e529201522fbd166b03a55712d945737f4cc5c1e18199caa7b.
//
// Solidity: event MarginChanged(address indexed sender, address indexed amm, int256 amount, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) FilterMarginChanged(opts *bind.FilterOpts, sender []common.Address, amm []common.Address) (*ClearingHouseMarginChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "MarginChanged", senderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseMarginChangedIterator{contract: _ClearingHouse.contract, event: "MarginChanged", logs: logs, sub: sub}, nil
}

// WatchMarginChanged is a free log subscription operation binding the contract event 0x1c822fe92555a7e529201522fbd166b03a55712d945737f4cc5c1e18199caa7b.
//
// Solidity: event MarginChanged(address indexed sender, address indexed amm, int256 amount, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) WatchMarginChanged(opts *bind.WatchOpts, sink chan<- *ClearingHouseMarginChanged, sender []common.Address, amm []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "MarginChanged", senderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseMarginChanged)
				if err := _ClearingHouse.contract.UnpackLog(event, "MarginChanged", log); err != nil {
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

// ParseMarginChanged is a log parse operation binding the contract event 0x1c822fe92555a7e529201522fbd166b03a55712d945737f4cc5c1e18199caa7b.
//
// Solidity: event MarginChanged(address indexed sender, address indexed amm, int256 amount, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) ParseMarginChanged(log types.Log) (*ClearingHouseMarginChanged, error) {
	event := new(ClearingHouseMarginChanged)
	if err := _ClearingHouse.contract.UnpackLog(event, "MarginChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseMarginRatioChangedIterator is returned from FilterMarginRatioChanged and is used to iterate over the raw logs and unpacked data for MarginRatioChanged events raised by the ClearingHouse contract.
type ClearingHouseMarginRatioChangedIterator struct {
	Event *ClearingHouseMarginRatioChanged // Event containing the contract specifics and raw log

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
func (it *ClearingHouseMarginRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseMarginRatioChanged)
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
		it.Event = new(ClearingHouseMarginRatioChanged)
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
func (it *ClearingHouseMarginRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseMarginRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseMarginRatioChanged represents a MarginRatioChanged event raised by the ClearingHouse contract.
type ClearingHouseMarginRatioChanged struct {
	MarginRatio *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarginRatioChanged is a free log retrieval operation binding the contract event 0xce738a6382894b012492e42016c4321ac69da80c4827f80a2e0907585818ac31.
//
// Solidity: event MarginRatioChanged(uint256 marginRatio)
func (_ClearingHouse *ClearingHouseFilterer) FilterMarginRatioChanged(opts *bind.FilterOpts) (*ClearingHouseMarginRatioChangedIterator, error) {

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "MarginRatioChanged")
	if err != nil {
		return nil, err
	}
	return &ClearingHouseMarginRatioChangedIterator{contract: _ClearingHouse.contract, event: "MarginRatioChanged", logs: logs, sub: sub}, nil
}

// WatchMarginRatioChanged is a free log subscription operation binding the contract event 0xce738a6382894b012492e42016c4321ac69da80c4827f80a2e0907585818ac31.
//
// Solidity: event MarginRatioChanged(uint256 marginRatio)
func (_ClearingHouse *ClearingHouseFilterer) WatchMarginRatioChanged(opts *bind.WatchOpts, sink chan<- *ClearingHouseMarginRatioChanged) (event.Subscription, error) {

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "MarginRatioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseMarginRatioChanged)
				if err := _ClearingHouse.contract.UnpackLog(event, "MarginRatioChanged", log); err != nil {
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

// ParseMarginRatioChanged is a log parse operation binding the contract event 0xce738a6382894b012492e42016c4321ac69da80c4827f80a2e0907585818ac31.
//
// Solidity: event MarginRatioChanged(uint256 marginRatio)
func (_ClearingHouse *ClearingHouseFilterer) ParseMarginRatioChanged(log types.Log) (*ClearingHouseMarginRatioChanged, error) {
	event := new(ClearingHouseMarginRatioChanged)
	if err := _ClearingHouse.contract.UnpackLog(event, "MarginRatioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ClearingHouse contract.
type ClearingHouseOwnershipTransferredIterator struct {
	Event *ClearingHouseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClearingHouseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseOwnershipTransferred)
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
		it.Event = new(ClearingHouseOwnershipTransferred)
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
func (it *ClearingHouseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseOwnershipTransferred represents a OwnershipTransferred event raised by the ClearingHouse contract.
type ClearingHouseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClearingHouse *ClearingHouseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClearingHouseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseOwnershipTransferredIterator{contract: _ClearingHouse.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClearingHouse *ClearingHouseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClearingHouseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseOwnershipTransferred)
				if err := _ClearingHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ClearingHouse *ClearingHouseFilterer) ParseOwnershipTransferred(log types.Log) (*ClearingHouseOwnershipTransferred, error) {
	event := new(ClearingHouseOwnershipTransferred)
	if err := _ClearingHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHousePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ClearingHouse contract.
type ClearingHousePausedIterator struct {
	Event *ClearingHousePaused // Event containing the contract specifics and raw log

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
func (it *ClearingHousePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHousePaused)
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
		it.Event = new(ClearingHousePaused)
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
func (it *ClearingHousePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHousePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHousePaused represents a Paused event raised by the ClearingHouse contract.
type ClearingHousePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ClearingHouse *ClearingHouseFilterer) FilterPaused(opts *bind.FilterOpts) (*ClearingHousePausedIterator, error) {

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ClearingHousePausedIterator{contract: _ClearingHouse.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ClearingHouse *ClearingHouseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ClearingHousePaused) (event.Subscription, error) {

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHousePaused)
				if err := _ClearingHouse.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ClearingHouse *ClearingHouseFilterer) ParsePaused(log types.Log) (*ClearingHousePaused, error) {
	event := new(ClearingHousePaused)
	if err := _ClearingHouse.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHousePositionAdjustedIterator is returned from FilterPositionAdjusted and is used to iterate over the raw logs and unpacked data for PositionAdjusted events raised by the ClearingHouse contract.
type ClearingHousePositionAdjustedIterator struct {
	Event *ClearingHousePositionAdjusted // Event containing the contract specifics and raw log

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
func (it *ClearingHousePositionAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHousePositionAdjusted)
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
		it.Event = new(ClearingHousePositionAdjusted)
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
func (it *ClearingHousePositionAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHousePositionAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHousePositionAdjusted represents a PositionAdjusted event raised by the ClearingHouse contract.
type ClearingHousePositionAdjusted struct {
	Amm               common.Address
	Trader            common.Address
	NewPositionSize   *big.Int
	OldLiquidityIndex *big.Int
	NewLiquidityIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPositionAdjusted is a free log retrieval operation binding the contract event 0xdbb898e847eca1f5626ba536593caa10241fb631e7c8a1cb06dfdce8d3934f49.
//
// Solidity: event PositionAdjusted(address indexed amm, address indexed trader, int256 newPositionSize, uint256 oldLiquidityIndex, uint256 newLiquidityIndex)
func (_ClearingHouse *ClearingHouseFilterer) FilterPositionAdjusted(opts *bind.FilterOpts, amm []common.Address, trader []common.Address) (*ClearingHousePositionAdjustedIterator, error) {

	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "PositionAdjusted", ammRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHousePositionAdjustedIterator{contract: _ClearingHouse.contract, event: "PositionAdjusted", logs: logs, sub: sub}, nil
}

// WatchPositionAdjusted is a free log subscription operation binding the contract event 0xdbb898e847eca1f5626ba536593caa10241fb631e7c8a1cb06dfdce8d3934f49.
//
// Solidity: event PositionAdjusted(address indexed amm, address indexed trader, int256 newPositionSize, uint256 oldLiquidityIndex, uint256 newLiquidityIndex)
func (_ClearingHouse *ClearingHouseFilterer) WatchPositionAdjusted(opts *bind.WatchOpts, sink chan<- *ClearingHousePositionAdjusted, amm []common.Address, trader []common.Address) (event.Subscription, error) {

	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "PositionAdjusted", ammRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHousePositionAdjusted)
				if err := _ClearingHouse.contract.UnpackLog(event, "PositionAdjusted", log); err != nil {
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

// ParsePositionAdjusted is a log parse operation binding the contract event 0xdbb898e847eca1f5626ba536593caa10241fb631e7c8a1cb06dfdce8d3934f49.
//
// Solidity: event PositionAdjusted(address indexed amm, address indexed trader, int256 newPositionSize, uint256 oldLiquidityIndex, uint256 newLiquidityIndex)
func (_ClearingHouse *ClearingHouseFilterer) ParsePositionAdjusted(log types.Log) (*ClearingHousePositionAdjusted, error) {
	event := new(ClearingHousePositionAdjusted)
	if err := _ClearingHouse.contract.UnpackLog(event, "PositionAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHousePositionChangedIterator is returned from FilterPositionChanged and is used to iterate over the raw logs and unpacked data for PositionChanged events raised by the ClearingHouse contract.
type ClearingHousePositionChangedIterator struct {
	Event *ClearingHousePositionChanged // Event containing the contract specifics and raw log

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
func (it *ClearingHousePositionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHousePositionChanged)
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
		it.Event = new(ClearingHousePositionChanged)
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
func (it *ClearingHousePositionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHousePositionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHousePositionChanged represents a PositionChanged event raised by the ClearingHouse contract.
type ClearingHousePositionChanged struct {
	Trader                common.Address
	Amm                   common.Address
	Margin                *big.Int
	PositionNotional      *big.Int
	ExchangedPositionSize *big.Int
	Fee                   *big.Int
	PositionSizeAfter     *big.Int
	RealizedPnl           *big.Int
	UnrealizedPnlAfter    *big.Int
	BadDebt               *big.Int
	LiquidationPenalty    *big.Int
	SpotPrice             *big.Int
	FundingPayment        *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPositionChanged is a free log retrieval operation binding the contract event 0x4c7b764f428c13bbea8cc8da90ebe6eef4dafeb27a4e3d9041d64208c47ca7c2.
//
// Solidity: event PositionChanged(address indexed trader, address indexed amm, uint256 margin, uint256 positionNotional, int256 exchangedPositionSize, uint256 fee, int256 positionSizeAfter, int256 realizedPnl, int256 unrealizedPnlAfter, uint256 badDebt, uint256 liquidationPenalty, uint256 spotPrice, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) FilterPositionChanged(opts *bind.FilterOpts, trader []common.Address, amm []common.Address) (*ClearingHousePositionChangedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "PositionChanged", traderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHousePositionChangedIterator{contract: _ClearingHouse.contract, event: "PositionChanged", logs: logs, sub: sub}, nil
}

// WatchPositionChanged is a free log subscription operation binding the contract event 0x4c7b764f428c13bbea8cc8da90ebe6eef4dafeb27a4e3d9041d64208c47ca7c2.
//
// Solidity: event PositionChanged(address indexed trader, address indexed amm, uint256 margin, uint256 positionNotional, int256 exchangedPositionSize, uint256 fee, int256 positionSizeAfter, int256 realizedPnl, int256 unrealizedPnlAfter, uint256 badDebt, uint256 liquidationPenalty, uint256 spotPrice, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) WatchPositionChanged(opts *bind.WatchOpts, sink chan<- *ClearingHousePositionChanged, trader []common.Address, amm []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "PositionChanged", traderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHousePositionChanged)
				if err := _ClearingHouse.contract.UnpackLog(event, "PositionChanged", log); err != nil {
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

// ParsePositionChanged is a log parse operation binding the contract event 0x4c7b764f428c13bbea8cc8da90ebe6eef4dafeb27a4e3d9041d64208c47ca7c2.
//
// Solidity: event PositionChanged(address indexed trader, address indexed amm, uint256 margin, uint256 positionNotional, int256 exchangedPositionSize, uint256 fee, int256 positionSizeAfter, int256 realizedPnl, int256 unrealizedPnlAfter, uint256 badDebt, uint256 liquidationPenalty, uint256 spotPrice, int256 fundingPayment)
func (_ClearingHouse *ClearingHouseFilterer) ParsePositionChanged(log types.Log) (*ClearingHousePositionChanged, error) {
	event := new(ClearingHousePositionChanged)
	if err := _ClearingHouse.contract.UnpackLog(event, "PositionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHousePositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the ClearingHouse contract.
type ClearingHousePositionLiquidatedIterator struct {
	Event *ClearingHousePositionLiquidated // Event containing the contract specifics and raw log

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
func (it *ClearingHousePositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHousePositionLiquidated)
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
		it.Event = new(ClearingHousePositionLiquidated)
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
func (it *ClearingHousePositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHousePositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHousePositionLiquidated represents a PositionLiquidated event raised by the ClearingHouse contract.
type ClearingHousePositionLiquidated struct {
	Trader           common.Address
	Amm              common.Address
	PositionNotional *big.Int
	PositionSize     *big.Int
	LiquidationFee   *big.Int
	Liquidator       common.Address
	BadDebt          *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0x5225638a979dd133f201045ab4169ec2189874d864f2b5a10be72ac6e4b421b4.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed amm, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator, uint256 badDebt)
func (_ClearingHouse *ClearingHouseFilterer) FilterPositionLiquidated(opts *bind.FilterOpts, trader []common.Address, amm []common.Address) (*ClearingHousePositionLiquidatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "PositionLiquidated", traderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHousePositionLiquidatedIterator{contract: _ClearingHouse.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0x5225638a979dd133f201045ab4169ec2189874d864f2b5a10be72ac6e4b421b4.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed amm, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator, uint256 badDebt)
func (_ClearingHouse *ClearingHouseFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *ClearingHousePositionLiquidated, trader []common.Address, amm []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "PositionLiquidated", traderRule, ammRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHousePositionLiquidated)
				if err := _ClearingHouse.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
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

// ParsePositionLiquidated is a log parse operation binding the contract event 0x5225638a979dd133f201045ab4169ec2189874d864f2b5a10be72ac6e4b421b4.
//
// Solidity: event PositionLiquidated(address indexed trader, address indexed amm, uint256 positionNotional, uint256 positionSize, uint256 liquidationFee, address liquidator, uint256 badDebt)
func (_ClearingHouse *ClearingHouseFilterer) ParsePositionLiquidated(log types.Log) (*ClearingHousePositionLiquidated, error) {
	event := new(ClearingHousePositionLiquidated)
	if err := _ClearingHouse.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHousePositionSettledIterator is returned from FilterPositionSettled and is used to iterate over the raw logs and unpacked data for PositionSettled events raised by the ClearingHouse contract.
type ClearingHousePositionSettledIterator struct {
	Event *ClearingHousePositionSettled // Event containing the contract specifics and raw log

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
func (it *ClearingHousePositionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHousePositionSettled)
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
		it.Event = new(ClearingHousePositionSettled)
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
func (it *ClearingHousePositionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHousePositionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHousePositionSettled represents a PositionSettled event raised by the ClearingHouse contract.
type ClearingHousePositionSettled struct {
	Amm              common.Address
	Trader           common.Address
	ValueTransferred *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPositionSettled is a free log retrieval operation binding the contract event 0xba70e8457921bdd49042182753cb6259c62f923e2b64216b68372cd85ec3af6c.
//
// Solidity: event PositionSettled(address indexed amm, address indexed trader, uint256 valueTransferred)
func (_ClearingHouse *ClearingHouseFilterer) FilterPositionSettled(opts *bind.FilterOpts, amm []common.Address, trader []common.Address) (*ClearingHousePositionSettledIterator, error) {

	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "PositionSettled", ammRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHousePositionSettledIterator{contract: _ClearingHouse.contract, event: "PositionSettled", logs: logs, sub: sub}, nil
}

// WatchPositionSettled is a free log subscription operation binding the contract event 0xba70e8457921bdd49042182753cb6259c62f923e2b64216b68372cd85ec3af6c.
//
// Solidity: event PositionSettled(address indexed amm, address indexed trader, uint256 valueTransferred)
func (_ClearingHouse *ClearingHouseFilterer) WatchPositionSettled(opts *bind.WatchOpts, sink chan<- *ClearingHousePositionSettled, amm []common.Address, trader []common.Address) (event.Subscription, error) {

	var ammRule []interface{}
	for _, ammItem := range amm {
		ammRule = append(ammRule, ammItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "PositionSettled", ammRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHousePositionSettled)
				if err := _ClearingHouse.contract.UnpackLog(event, "PositionSettled", log); err != nil {
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

// ParsePositionSettled is a log parse operation binding the contract event 0xba70e8457921bdd49042182753cb6259c62f923e2b64216b68372cd85ec3af6c.
//
// Solidity: event PositionSettled(address indexed amm, address indexed trader, uint256 valueTransferred)
func (_ClearingHouse *ClearingHouseFilterer) ParsePositionSettled(log types.Log) (*ClearingHousePositionSettled, error) {
	event := new(ClearingHousePositionSettled)
	if err := _ClearingHouse.contract.UnpackLog(event, "PositionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseReferredPositionChangedIterator is returned from FilterReferredPositionChanged and is used to iterate over the raw logs and unpacked data for ReferredPositionChanged events raised by the ClearingHouse contract.
type ClearingHouseReferredPositionChangedIterator struct {
	Event *ClearingHouseReferredPositionChanged // Event containing the contract specifics and raw log

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
func (it *ClearingHouseReferredPositionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseReferredPositionChanged)
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
		it.Event = new(ClearingHouseReferredPositionChanged)
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
func (it *ClearingHouseReferredPositionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseReferredPositionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseReferredPositionChanged represents a ReferredPositionChanged event raised by the ClearingHouse contract.
type ClearingHouseReferredPositionChanged struct {
	ReferralCode [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReferredPositionChanged is a free log retrieval operation binding the contract event 0x09a07769667a46c2b89124a4d731c76fb5203061073aec16319ba9c6a79a3248.
//
// Solidity: event ReferredPositionChanged(bytes32 indexed referralCode)
func (_ClearingHouse *ClearingHouseFilterer) FilterReferredPositionChanged(opts *bind.FilterOpts, referralCode [][32]byte) (*ClearingHouseReferredPositionChangedIterator, error) {

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "ReferredPositionChanged", referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &ClearingHouseReferredPositionChangedIterator{contract: _ClearingHouse.contract, event: "ReferredPositionChanged", logs: logs, sub: sub}, nil
}

// WatchReferredPositionChanged is a free log subscription operation binding the contract event 0x09a07769667a46c2b89124a4d731c76fb5203061073aec16319ba9c6a79a3248.
//
// Solidity: event ReferredPositionChanged(bytes32 indexed referralCode)
func (_ClearingHouse *ClearingHouseFilterer) WatchReferredPositionChanged(opts *bind.WatchOpts, sink chan<- *ClearingHouseReferredPositionChanged, referralCode [][32]byte) (event.Subscription, error) {

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "ReferredPositionChanged", referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseReferredPositionChanged)
				if err := _ClearingHouse.contract.UnpackLog(event, "ReferredPositionChanged", log); err != nil {
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

// ParseReferredPositionChanged is a log parse operation binding the contract event 0x09a07769667a46c2b89124a4d731c76fb5203061073aec16319ba9c6a79a3248.
//
// Solidity: event ReferredPositionChanged(bytes32 indexed referralCode)
func (_ClearingHouse *ClearingHouseFilterer) ParseReferredPositionChanged(log types.Log) (*ClearingHouseReferredPositionChanged, error) {
	event := new(ClearingHouseReferredPositionChanged)
	if err := _ClearingHouse.contract.UnpackLog(event, "ReferredPositionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseRestrictionModeEnteredIterator is returned from FilterRestrictionModeEntered and is used to iterate over the raw logs and unpacked data for RestrictionModeEntered events raised by the ClearingHouse contract.
type ClearingHouseRestrictionModeEnteredIterator struct {
	Event *ClearingHouseRestrictionModeEntered // Event containing the contract specifics and raw log

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
func (it *ClearingHouseRestrictionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseRestrictionModeEntered)
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
		it.Event = new(ClearingHouseRestrictionModeEntered)
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
func (it *ClearingHouseRestrictionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseRestrictionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseRestrictionModeEntered represents a RestrictionModeEntered event raised by the ClearingHouse contract.
type ClearingHouseRestrictionModeEntered struct {
	Amm         common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRestrictionModeEntered is a free log retrieval operation binding the contract event 0x97f920fec5d17ec336df7ad438ca8a76dcfd206ddc1d19b263b6f3e6ba3c4f31.
//
// Solidity: event RestrictionModeEntered(address amm, uint256 blockNumber)
func (_ClearingHouse *ClearingHouseFilterer) FilterRestrictionModeEntered(opts *bind.FilterOpts) (*ClearingHouseRestrictionModeEnteredIterator, error) {

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "RestrictionModeEntered")
	if err != nil {
		return nil, err
	}
	return &ClearingHouseRestrictionModeEnteredIterator{contract: _ClearingHouse.contract, event: "RestrictionModeEntered", logs: logs, sub: sub}, nil
}

// WatchRestrictionModeEntered is a free log subscription operation binding the contract event 0x97f920fec5d17ec336df7ad438ca8a76dcfd206ddc1d19b263b6f3e6ba3c4f31.
//
// Solidity: event RestrictionModeEntered(address amm, uint256 blockNumber)
func (_ClearingHouse *ClearingHouseFilterer) WatchRestrictionModeEntered(opts *bind.WatchOpts, sink chan<- *ClearingHouseRestrictionModeEntered) (event.Subscription, error) {

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "RestrictionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseRestrictionModeEntered)
				if err := _ClearingHouse.contract.UnpackLog(event, "RestrictionModeEntered", log); err != nil {
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

// ParseRestrictionModeEntered is a log parse operation binding the contract event 0x97f920fec5d17ec336df7ad438ca8a76dcfd206ddc1d19b263b6f3e6ba3c4f31.
//
// Solidity: event RestrictionModeEntered(address amm, uint256 blockNumber)
func (_ClearingHouse *ClearingHouseFilterer) ParseRestrictionModeEntered(log types.Log) (*ClearingHouseRestrictionModeEntered, error) {
	event := new(ClearingHouseRestrictionModeEntered)
	if err := _ClearingHouse.contract.UnpackLog(event, "RestrictionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearingHouseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ClearingHouse contract.
type ClearingHouseUnpausedIterator struct {
	Event *ClearingHouseUnpaused // Event containing the contract specifics and raw log

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
func (it *ClearingHouseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearingHouseUnpaused)
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
		it.Event = new(ClearingHouseUnpaused)
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
func (it *ClearingHouseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearingHouseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearingHouseUnpaused represents a Unpaused event raised by the ClearingHouse contract.
type ClearingHouseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ClearingHouse *ClearingHouseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ClearingHouseUnpausedIterator, error) {

	logs, sub, err := _ClearingHouse.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ClearingHouseUnpausedIterator{contract: _ClearingHouse.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ClearingHouse *ClearingHouseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ClearingHouseUnpaused) (event.Subscription, error) {

	logs, sub, err := _ClearingHouse.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearingHouseUnpaused)
				if err := _ClearingHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ClearingHouse *ClearingHouseFilterer) ParseUnpaused(log types.Log) (*ClearingHouseUnpaused, error) {
	event := new(ClearingHouseUnpaused)
	if err := _ClearingHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
