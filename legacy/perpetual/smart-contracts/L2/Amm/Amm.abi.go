// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Amm

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

// Decimaldecimal is an auto generated low-level Go binding around an user-defined struct.
type Decimaldecimal struct {
	D *big.Int
}

// IAmmLiquidityChangedSnapshot is an auto generated low-level Go binding around an user-defined struct.
type IAmmLiquidityChangedSnapshot struct {
	CumulativeNotional SignedDecimalsignedDecimal
	QuoteAssetReserve  Decimaldecimal
	BaseAssetReserve   Decimaldecimal
	TotalPositionSize  SignedDecimalsignedDecimal
}

// SignedDecimalsignedDecimal is an auto generated low-level Go binding around an user-defined struct.
type SignedDecimalsignedDecimal struct {
	D *big.Int
}

// AmmABI is the input ABI used to generate the binding from.
const AmmABI = "[{\"type\":\"event\",\"name\":\"CapChanged\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxHoldingBaseAsset\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"openInterestNotionalCap\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundingRateUpdated\",\"inputs\":[{\"type\":\"int256\",\"name\":\"rate\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"underlyingPrice\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityChanged\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"quoteReserve\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"baseReserve\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"cumulativeNotional\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"priceFeed\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReserveSnapshotted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"quoteAssetReserve\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"baseAssetReserve\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Shutdown\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"settlementPrice\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwapInput\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"dir\",\"internalType\":\"enumIAmm.Dir\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"quoteAssetAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"baseAssetAmount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwapOutput\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"dir\",\"internalType\":\"enumIAmm.Dir\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"quoteAssetAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"baseAssetAmount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"MAX_ORACLE_SPREAD_RATIO\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"baseAssetReserve\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"calcBaseAssetAfterLiquidityMigration\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"_fromQuoteReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_fromBaseReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"calcFee\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"candidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"fluctuationLimitRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"fundingBufferPeriod\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"fundingPeriod\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}],\"name\":\"fundingRate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getBaseAssetDelta\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getBaseAssetDeltaThisFundingPeriod\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"getCumulativeNotional\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getInputPrice\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfQuote\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getInputPriceWithReserves\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfQuote\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_quoteAssetPoolAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_baseAssetPoolAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getInputTwap\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfQuote\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIAmm.LiquidityChangedSnapshot\",\"components\":[{\"type\":\"tuple\",\"name\":\"cumulativeNotional\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"quoteAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"baseAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"totalPositionSize\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}]}],\"name\":\"getLatestLiquidityChangedSnapshots\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIAmm.LiquidityChangedSnapshot\",\"components\":[{\"type\":\"tuple\",\"name\":\"cumulativeNotional\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]},{\"type\":\"tuple\",\"name\":\"quoteAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"baseAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"totalPositionSize\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}]}],\"name\":\"getLiquidityChangedSnapshots\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"i\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLiquidityHistoryLength\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getMaxHoldingBaseAsset\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getOpenInterestNotionalCap\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getOutputPrice\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfBase\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getOutputPriceWithReserves\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfBase\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_quoteAssetPoolAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_baseAssetPoolAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getOutputTwap\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfBase\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getReserve\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getSettlementPrice\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getSnapshotLen\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getSpotPrice\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getTwapPrice\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_intervalInSeconds\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getUnderlyingPrice\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"getUnderlyingTwapPrice\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_intervalInSeconds\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"globalShutdown\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_quoteAssetReserve\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_baseAssetReserve\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_tradeLimitRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_fundingPeriod\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_priceFeed\",\"internalType\":\"contractIPriceFeed\"},{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"_quoteAsset\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_fluctuationLimitRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_tollRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_spreadRatio\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isOverFluctuationLimit\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfBase\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isOverSpreadLimit\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"nextFundingTime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"open\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIPriceFeed\"}],\"name\":\"priceFeed\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"priceFeedKey\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIERC20\"}],\"name\":\"quoteAsset\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"quoteAssetReserve\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"quoteAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"baseAssetReserve\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\"}],\"name\":\"reserveSnapshots\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setCap\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_maxHoldingBaseAsset\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_openInterestNotionalCap\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setCounterParty\",\"inputs\":[{\"type\":\"address\",\"name\":\"_counterParty\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFluctuationLimitRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_fluctuationLimitRatio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setGlobalShutdown\",\"inputs\":[{\"type\":\"address\",\"name\":\"_globalShutdown\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOpen\",\"inputs\":[{\"type\":\"bool\",\"name\":\"_open\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setPriceFeed\",\"inputs\":[{\"type\":\"address\",\"name\":\"_priceFeed\",\"internalType\":\"contractIPriceFeed\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSpotPriceTwapInterval\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_interval\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSpreadRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_spreadRatio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTollRatio\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_tollRatio\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structSignedDecimal.signedDecimal\",\"components\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}]}],\"name\":\"settleFunding\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"shutdown\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"spotPriceTwapInterval\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"spreadRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"swapInput\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfQuote\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_baseAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"bool\",\"name\":\"_canOverFluctuationLimit\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}],\"name\":\"swapOutput\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_dirOfBase\",\"internalType\":\"enumIAmm.Dir\"},{\"type\":\"tuple\",\"name\":\"_baseAssetAmount\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple\",\"name\":\"_quoteAssetAmountLimit\",\"internalType\":\"structDecimal.decimal\",\"components\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"tollAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"tollRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"int256\",\"name\":\"d\",\"internalType\":\"int256\"}],\"name\":\"totalPositionSize\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"d\",\"internalType\":\"uint256\"}],\"name\":\"tradeLimitRatio\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateOwner\",\"inputs\":[]}]"

// Amm is an auto generated Go binding around an Ethereum contract.
type Amm struct {
	AmmCaller     // Read-only binding to the contract
	AmmTransactor // Write-only binding to the contract
	AmmFilterer   // Log filterer for contract events
}

// AmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmmSession struct {
	Contract     *Amm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmmCallerSession struct {
	Contract *AmmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmmTransactorSession struct {
	Contract     *AmmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmmRaw struct {
	Contract *Amm // Generic contract binding to access the raw methods on
}

// AmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmmCallerRaw struct {
	Contract *AmmCaller // Generic read-only contract binding to access the raw methods on
}

// AmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmmTransactorRaw struct {
	Contract *AmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmm creates a new instance of Amm, bound to a specific deployed contract.
func NewAmm(address common.Address, backend bind.ContractBackend) (*Amm, error) {
	contract, err := bindAmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Amm{AmmCaller: AmmCaller{contract: contract}, AmmTransactor: AmmTransactor{contract: contract}, AmmFilterer: AmmFilterer{contract: contract}}, nil
}

// NewAmmCaller creates a new read-only instance of Amm, bound to a specific deployed contract.
func NewAmmCaller(address common.Address, caller bind.ContractCaller) (*AmmCaller, error) {
	contract, err := bindAmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmmCaller{contract: contract}, nil
}

// NewAmmTransactor creates a new write-only instance of Amm, bound to a specific deployed contract.
func NewAmmTransactor(address common.Address, transactor bind.ContractTransactor) (*AmmTransactor, error) {
	contract, err := bindAmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmmTransactor{contract: contract}, nil
}

// NewAmmFilterer creates a new log filterer instance of Amm, bound to a specific deployed contract.
func NewAmmFilterer(address common.Address, filterer bind.ContractFilterer) (*AmmFilterer, error) {
	contract, err := bindAmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmmFilterer{contract: contract}, nil
}

// bindAmm binds a generic wrapper to an already deployed contract.
func bindAmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.AmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transact(opts, method, params...)
}

// MAXORACLESPREADRATIO is a free data retrieval call binding the contract method 0x05172a25.
//
// Solidity: function MAX_ORACLE_SPREAD_RATIO() view returns(uint256)
func (_Amm *AmmCaller) MAXORACLESPREADRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "MAX_ORACLE_SPREAD_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXORACLESPREADRATIO is a free data retrieval call binding the contract method 0x05172a25.
//
// Solidity: function MAX_ORACLE_SPREAD_RATIO() view returns(uint256)
func (_Amm *AmmSession) MAXORACLESPREADRATIO() (*big.Int, error) {
	return _Amm.Contract.MAXORACLESPREADRATIO(&_Amm.CallOpts)
}

// MAXORACLESPREADRATIO is a free data retrieval call binding the contract method 0x05172a25.
//
// Solidity: function MAX_ORACLE_SPREAD_RATIO() view returns(uint256)
func (_Amm *AmmCallerSession) MAXORACLESPREADRATIO() (*big.Int, error) {
	return _Amm.Contract.MAXORACLESPREADRATIO(&_Amm.CallOpts)
}

// BaseAssetReserve is a free data retrieval call binding the contract method 0xc9566fcc.
//
// Solidity: function baseAssetReserve() view returns(uint256 d)
func (_Amm *AmmCaller) BaseAssetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "baseAssetReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseAssetReserve is a free data retrieval call binding the contract method 0xc9566fcc.
//
// Solidity: function baseAssetReserve() view returns(uint256 d)
func (_Amm *AmmSession) BaseAssetReserve() (*big.Int, error) {
	return _Amm.Contract.BaseAssetReserve(&_Amm.CallOpts)
}

// BaseAssetReserve is a free data retrieval call binding the contract method 0xc9566fcc.
//
// Solidity: function baseAssetReserve() view returns(uint256 d)
func (_Amm *AmmCallerSession) BaseAssetReserve() (*big.Int, error) {
	return _Amm.Contract.BaseAssetReserve(&_Amm.CallOpts)
}

// CalcBaseAssetAfterLiquidityMigration is a free data retrieval call binding the contract method 0x40d71cd9.
//
// Solidity: function calcBaseAssetAfterLiquidityMigration((int256) _baseAssetAmount, (uint256) _fromQuoteReserve, (uint256) _fromBaseReserve) view returns((int256))
func (_Amm *AmmCaller) CalcBaseAssetAfterLiquidityMigration(opts *bind.CallOpts, _baseAssetAmount SignedDecimalsignedDecimal, _fromQuoteReserve Decimaldecimal, _fromBaseReserve Decimaldecimal) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "calcBaseAssetAfterLiquidityMigration", _baseAssetAmount, _fromQuoteReserve, _fromBaseReserve)

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// CalcBaseAssetAfterLiquidityMigration is a free data retrieval call binding the contract method 0x40d71cd9.
//
// Solidity: function calcBaseAssetAfterLiquidityMigration((int256) _baseAssetAmount, (uint256) _fromQuoteReserve, (uint256) _fromBaseReserve) view returns((int256))
func (_Amm *AmmSession) CalcBaseAssetAfterLiquidityMigration(_baseAssetAmount SignedDecimalsignedDecimal, _fromQuoteReserve Decimaldecimal, _fromBaseReserve Decimaldecimal) (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.CalcBaseAssetAfterLiquidityMigration(&_Amm.CallOpts, _baseAssetAmount, _fromQuoteReserve, _fromBaseReserve)
}

// CalcBaseAssetAfterLiquidityMigration is a free data retrieval call binding the contract method 0x40d71cd9.
//
// Solidity: function calcBaseAssetAfterLiquidityMigration((int256) _baseAssetAmount, (uint256) _fromQuoteReserve, (uint256) _fromBaseReserve) view returns((int256))
func (_Amm *AmmCallerSession) CalcBaseAssetAfterLiquidityMigration(_baseAssetAmount SignedDecimalsignedDecimal, _fromQuoteReserve Decimaldecimal, _fromBaseReserve Decimaldecimal) (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.CalcBaseAssetAfterLiquidityMigration(&_Amm.CallOpts, _baseAssetAmount, _fromQuoteReserve, _fromBaseReserve)
}

// CalcFee is a free data retrieval call binding the contract method 0x62267955.
//
// Solidity: function calcFee((uint256) _quoteAssetAmount) view returns((uint256), (uint256))
func (_Amm *AmmCaller) CalcFee(opts *bind.CallOpts, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "calcFee", _quoteAssetAmount)

	if err != nil {
		return *new(Decimaldecimal), *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)
	out1 := *abi.ConvertType(out[1], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, out1, err

}

// CalcFee is a free data retrieval call binding the contract method 0x62267955.
//
// Solidity: function calcFee((uint256) _quoteAssetAmount) view returns((uint256), (uint256))
func (_Amm *AmmSession) CalcFee(_quoteAssetAmount Decimaldecimal) (Decimaldecimal, Decimaldecimal, error) {
	return _Amm.Contract.CalcFee(&_Amm.CallOpts, _quoteAssetAmount)
}

// CalcFee is a free data retrieval call binding the contract method 0x62267955.
//
// Solidity: function calcFee((uint256) _quoteAssetAmount) view returns((uint256), (uint256))
func (_Amm *AmmCallerSession) CalcFee(_quoteAssetAmount Decimaldecimal) (Decimaldecimal, Decimaldecimal, error) {
	return _Amm.Contract.CalcFee(&_Amm.CallOpts, _quoteAssetAmount)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_Amm *AmmCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_Amm *AmmSession) Candidate() (common.Address, error) {
	return _Amm.Contract.Candidate(&_Amm.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_Amm *AmmCallerSession) Candidate() (common.Address, error) {
	return _Amm.Contract.Candidate(&_Amm.CallOpts)
}

// FluctuationLimitRatio is a free data retrieval call binding the contract method 0x1d3acb44.
//
// Solidity: function fluctuationLimitRatio() view returns(uint256 d)
func (_Amm *AmmCaller) FluctuationLimitRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "fluctuationLimitRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FluctuationLimitRatio is a free data retrieval call binding the contract method 0x1d3acb44.
//
// Solidity: function fluctuationLimitRatio() view returns(uint256 d)
func (_Amm *AmmSession) FluctuationLimitRatio() (*big.Int, error) {
	return _Amm.Contract.FluctuationLimitRatio(&_Amm.CallOpts)
}

// FluctuationLimitRatio is a free data retrieval call binding the contract method 0x1d3acb44.
//
// Solidity: function fluctuationLimitRatio() view returns(uint256 d)
func (_Amm *AmmCallerSession) FluctuationLimitRatio() (*big.Int, error) {
	return _Amm.Contract.FluctuationLimitRatio(&_Amm.CallOpts)
}

// FundingBufferPeriod is a free data retrieval call binding the contract method 0xc2de442f.
//
// Solidity: function fundingBufferPeriod() view returns(uint256)
func (_Amm *AmmCaller) FundingBufferPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "fundingBufferPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingBufferPeriod is a free data retrieval call binding the contract method 0xc2de442f.
//
// Solidity: function fundingBufferPeriod() view returns(uint256)
func (_Amm *AmmSession) FundingBufferPeriod() (*big.Int, error) {
	return _Amm.Contract.FundingBufferPeriod(&_Amm.CallOpts)
}

// FundingBufferPeriod is a free data retrieval call binding the contract method 0xc2de442f.
//
// Solidity: function fundingBufferPeriod() view returns(uint256)
func (_Amm *AmmCallerSession) FundingBufferPeriod() (*big.Int, error) {
	return _Amm.Contract.FundingBufferPeriod(&_Amm.CallOpts)
}

// FundingPeriod is a free data retrieval call binding the contract method 0x74d7c62b.
//
// Solidity: function fundingPeriod() view returns(uint256)
func (_Amm *AmmCaller) FundingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "fundingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPeriod is a free data retrieval call binding the contract method 0x74d7c62b.
//
// Solidity: function fundingPeriod() view returns(uint256)
func (_Amm *AmmSession) FundingPeriod() (*big.Int, error) {
	return _Amm.Contract.FundingPeriod(&_Amm.CallOpts)
}

// FundingPeriod is a free data retrieval call binding the contract method 0x74d7c62b.
//
// Solidity: function fundingPeriod() view returns(uint256)
func (_Amm *AmmCallerSession) FundingPeriod() (*big.Int, error) {
	return _Amm.Contract.FundingPeriod(&_Amm.CallOpts)
}

// FundingRate is a free data retrieval call binding the contract method 0x41d3c84c.
//
// Solidity: function fundingRate() view returns(int256 d)
func (_Amm *AmmCaller) FundingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "fundingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRate is a free data retrieval call binding the contract method 0x41d3c84c.
//
// Solidity: function fundingRate() view returns(int256 d)
func (_Amm *AmmSession) FundingRate() (*big.Int, error) {
	return _Amm.Contract.FundingRate(&_Amm.CallOpts)
}

// FundingRate is a free data retrieval call binding the contract method 0x41d3c84c.
//
// Solidity: function fundingRate() view returns(int256 d)
func (_Amm *AmmCallerSession) FundingRate() (*big.Int, error) {
	return _Amm.Contract.FundingRate(&_Amm.CallOpts)
}

// GetBaseAssetDelta is a free data retrieval call binding the contract method 0x2f848859.
//
// Solidity: function getBaseAssetDelta() view returns((int256))
func (_Amm *AmmCaller) GetBaseAssetDelta(opts *bind.CallOpts) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getBaseAssetDelta")

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetBaseAssetDelta is a free data retrieval call binding the contract method 0x2f848859.
//
// Solidity: function getBaseAssetDelta() view returns((int256))
func (_Amm *AmmSession) GetBaseAssetDelta() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetBaseAssetDelta(&_Amm.CallOpts)
}

// GetBaseAssetDelta is a free data retrieval call binding the contract method 0x2f848859.
//
// Solidity: function getBaseAssetDelta() view returns((int256))
func (_Amm *AmmCallerSession) GetBaseAssetDelta() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetBaseAssetDelta(&_Amm.CallOpts)
}

// GetBaseAssetDeltaThisFundingPeriod is a free data retrieval call binding the contract method 0xe805d6fc.
//
// Solidity: function getBaseAssetDeltaThisFundingPeriod() view returns((int256))
func (_Amm *AmmCaller) GetBaseAssetDeltaThisFundingPeriod(opts *bind.CallOpts) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getBaseAssetDeltaThisFundingPeriod")

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetBaseAssetDeltaThisFundingPeriod is a free data retrieval call binding the contract method 0xe805d6fc.
//
// Solidity: function getBaseAssetDeltaThisFundingPeriod() view returns((int256))
func (_Amm *AmmSession) GetBaseAssetDeltaThisFundingPeriod() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetBaseAssetDeltaThisFundingPeriod(&_Amm.CallOpts)
}

// GetBaseAssetDeltaThisFundingPeriod is a free data retrieval call binding the contract method 0xe805d6fc.
//
// Solidity: function getBaseAssetDeltaThisFundingPeriod() view returns((int256))
func (_Amm *AmmCallerSession) GetBaseAssetDeltaThisFundingPeriod() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetBaseAssetDeltaThisFundingPeriod(&_Amm.CallOpts)
}

// GetCumulativeNotional is a free data retrieval call binding the contract method 0x29f9b17b.
//
// Solidity: function getCumulativeNotional() view returns((int256))
func (_Amm *AmmCaller) GetCumulativeNotional(opts *bind.CallOpts) (SignedDecimalsignedDecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getCumulativeNotional")

	if err != nil {
		return *new(SignedDecimalsignedDecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedDecimalsignedDecimal)).(*SignedDecimalsignedDecimal)

	return out0, err

}

// GetCumulativeNotional is a free data retrieval call binding the contract method 0x29f9b17b.
//
// Solidity: function getCumulativeNotional() view returns((int256))
func (_Amm *AmmSession) GetCumulativeNotional() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetCumulativeNotional(&_Amm.CallOpts)
}

// GetCumulativeNotional is a free data retrieval call binding the contract method 0x29f9b17b.
//
// Solidity: function getCumulativeNotional() view returns((int256))
func (_Amm *AmmCallerSession) GetCumulativeNotional() (SignedDecimalsignedDecimal, error) {
	return _Amm.Contract.GetCumulativeNotional(&_Amm.CallOpts)
}

// GetInputPrice is a free data retrieval call binding the contract method 0xe1f1027f.
//
// Solidity: function getInputPrice(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmCaller) GetInputPrice(opts *bind.CallOpts, _dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getInputPrice", _dirOfQuote, _quoteAssetAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetInputPrice is a free data retrieval call binding the contract method 0xe1f1027f.
//
// Solidity: function getInputPrice(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmSession) GetInputPrice(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputPrice(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount)
}

// GetInputPrice is a free data retrieval call binding the contract method 0xe1f1027f.
//
// Solidity: function getInputPrice(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmCallerSession) GetInputPrice(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputPrice(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount)
}

// GetInputPriceWithReserves is a free data retrieval call binding the contract method 0x50799c81.
//
// Solidity: function getInputPriceWithReserves(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmCaller) GetInputPriceWithReserves(opts *bind.CallOpts, _dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getInputPriceWithReserves", _dirOfQuote, _quoteAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetInputPriceWithReserves is a free data retrieval call binding the contract method 0x50799c81.
//
// Solidity: function getInputPriceWithReserves(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmSession) GetInputPriceWithReserves(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputPriceWithReserves(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)
}

// GetInputPriceWithReserves is a free data retrieval call binding the contract method 0x50799c81.
//
// Solidity: function getInputPriceWithReserves(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmCallerSession) GetInputPriceWithReserves(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputPriceWithReserves(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)
}

// GetInputTwap is a free data retrieval call binding the contract method 0x4894d183.
//
// Solidity: function getInputTwap(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmCaller) GetInputTwap(opts *bind.CallOpts, _dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getInputTwap", _dirOfQuote, _quoteAssetAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetInputTwap is a free data retrieval call binding the contract method 0x4894d183.
//
// Solidity: function getInputTwap(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmSession) GetInputTwap(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputTwap(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount)
}

// GetInputTwap is a free data retrieval call binding the contract method 0x4894d183.
//
// Solidity: function getInputTwap(uint8 _dirOfQuote, (uint256) _quoteAssetAmount) view returns((uint256))
func (_Amm *AmmCallerSession) GetInputTwap(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetInputTwap(&_Amm.CallOpts, _dirOfQuote, _quoteAssetAmount)
}

// GetLatestLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x1b584d6c.
//
// Solidity: function getLatestLiquidityChangedSnapshots() view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmCaller) GetLatestLiquidityChangedSnapshots(opts *bind.CallOpts) (IAmmLiquidityChangedSnapshot, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getLatestLiquidityChangedSnapshots")

	if err != nil {
		return *new(IAmmLiquidityChangedSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(IAmmLiquidityChangedSnapshot)).(*IAmmLiquidityChangedSnapshot)

	return out0, err

}

// GetLatestLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x1b584d6c.
//
// Solidity: function getLatestLiquidityChangedSnapshots() view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmSession) GetLatestLiquidityChangedSnapshots() (IAmmLiquidityChangedSnapshot, error) {
	return _Amm.Contract.GetLatestLiquidityChangedSnapshots(&_Amm.CallOpts)
}

// GetLatestLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x1b584d6c.
//
// Solidity: function getLatestLiquidityChangedSnapshots() view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmCallerSession) GetLatestLiquidityChangedSnapshots() (IAmmLiquidityChangedSnapshot, error) {
	return _Amm.Contract.GetLatestLiquidityChangedSnapshots(&_Amm.CallOpts)
}

// GetLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x237f17ee.
//
// Solidity: function getLiquidityChangedSnapshots(uint256 i) view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmCaller) GetLiquidityChangedSnapshots(opts *bind.CallOpts, i *big.Int) (IAmmLiquidityChangedSnapshot, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getLiquidityChangedSnapshots", i)

	if err != nil {
		return *new(IAmmLiquidityChangedSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(IAmmLiquidityChangedSnapshot)).(*IAmmLiquidityChangedSnapshot)

	return out0, err

}

// GetLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x237f17ee.
//
// Solidity: function getLiquidityChangedSnapshots(uint256 i) view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmSession) GetLiquidityChangedSnapshots(i *big.Int) (IAmmLiquidityChangedSnapshot, error) {
	return _Amm.Contract.GetLiquidityChangedSnapshots(&_Amm.CallOpts, i)
}

// GetLiquidityChangedSnapshots is a free data retrieval call binding the contract method 0x237f17ee.
//
// Solidity: function getLiquidityChangedSnapshots(uint256 i) view returns(((int256),(uint256),(uint256),(int256)))
func (_Amm *AmmCallerSession) GetLiquidityChangedSnapshots(i *big.Int) (IAmmLiquidityChangedSnapshot, error) {
	return _Amm.Contract.GetLiquidityChangedSnapshots(&_Amm.CallOpts, i)
}

// GetLiquidityHistoryLength is a free data retrieval call binding the contract method 0x42b3198b.
//
// Solidity: function getLiquidityHistoryLength() view returns(uint256)
func (_Amm *AmmCaller) GetLiquidityHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getLiquidityHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidityHistoryLength is a free data retrieval call binding the contract method 0x42b3198b.
//
// Solidity: function getLiquidityHistoryLength() view returns(uint256)
func (_Amm *AmmSession) GetLiquidityHistoryLength() (*big.Int, error) {
	return _Amm.Contract.GetLiquidityHistoryLength(&_Amm.CallOpts)
}

// GetLiquidityHistoryLength is a free data retrieval call binding the contract method 0x42b3198b.
//
// Solidity: function getLiquidityHistoryLength() view returns(uint256)
func (_Amm *AmmCallerSession) GetLiquidityHistoryLength() (*big.Int, error) {
	return _Amm.Contract.GetLiquidityHistoryLength(&_Amm.CallOpts)
}

// GetMaxHoldingBaseAsset is a free data retrieval call binding the contract method 0x11377394.
//
// Solidity: function getMaxHoldingBaseAsset() view returns((uint256))
func (_Amm *AmmCaller) GetMaxHoldingBaseAsset(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getMaxHoldingBaseAsset")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetMaxHoldingBaseAsset is a free data retrieval call binding the contract method 0x11377394.
//
// Solidity: function getMaxHoldingBaseAsset() view returns((uint256))
func (_Amm *AmmSession) GetMaxHoldingBaseAsset() (Decimaldecimal, error) {
	return _Amm.Contract.GetMaxHoldingBaseAsset(&_Amm.CallOpts)
}

// GetMaxHoldingBaseAsset is a free data retrieval call binding the contract method 0x11377394.
//
// Solidity: function getMaxHoldingBaseAsset() view returns((uint256))
func (_Amm *AmmCallerSession) GetMaxHoldingBaseAsset() (Decimaldecimal, error) {
	return _Amm.Contract.GetMaxHoldingBaseAsset(&_Amm.CallOpts)
}

// GetOpenInterestNotionalCap is a free data retrieval call binding the contract method 0x52545410.
//
// Solidity: function getOpenInterestNotionalCap() view returns((uint256))
func (_Amm *AmmCaller) GetOpenInterestNotionalCap(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getOpenInterestNotionalCap")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetOpenInterestNotionalCap is a free data retrieval call binding the contract method 0x52545410.
//
// Solidity: function getOpenInterestNotionalCap() view returns((uint256))
func (_Amm *AmmSession) GetOpenInterestNotionalCap() (Decimaldecimal, error) {
	return _Amm.Contract.GetOpenInterestNotionalCap(&_Amm.CallOpts)
}

// GetOpenInterestNotionalCap is a free data retrieval call binding the contract method 0x52545410.
//
// Solidity: function getOpenInterestNotionalCap() view returns((uint256))
func (_Amm *AmmCallerSession) GetOpenInterestNotionalCap() (Decimaldecimal, error) {
	return _Amm.Contract.GetOpenInterestNotionalCap(&_Amm.CallOpts)
}

// GetOutputPrice is a free data retrieval call binding the contract method 0xda0c5927.
//
// Solidity: function getOutputPrice(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmCaller) GetOutputPrice(opts *bind.CallOpts, _dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getOutputPrice", _dirOfBase, _baseAssetAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetOutputPrice is a free data retrieval call binding the contract method 0xda0c5927.
//
// Solidity: function getOutputPrice(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmSession) GetOutputPrice(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputPrice(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// GetOutputPrice is a free data retrieval call binding the contract method 0xda0c5927.
//
// Solidity: function getOutputPrice(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmCallerSession) GetOutputPrice(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputPrice(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// GetOutputPriceWithReserves is a free data retrieval call binding the contract method 0x9bf5d1d4.
//
// Solidity: function getOutputPriceWithReserves(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmCaller) GetOutputPriceWithReserves(opts *bind.CallOpts, _dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getOutputPriceWithReserves", _dirOfBase, _baseAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetOutputPriceWithReserves is a free data retrieval call binding the contract method 0x9bf5d1d4.
//
// Solidity: function getOutputPriceWithReserves(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmSession) GetOutputPriceWithReserves(_dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputPriceWithReserves(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)
}

// GetOutputPriceWithReserves is a free data retrieval call binding the contract method 0x9bf5d1d4.
//
// Solidity: function getOutputPriceWithReserves(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetPoolAmount, (uint256) _baseAssetPoolAmount) pure returns((uint256))
func (_Amm *AmmCallerSession) GetOutputPriceWithReserves(_dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetPoolAmount Decimaldecimal, _baseAssetPoolAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputPriceWithReserves(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount, _quoteAssetPoolAmount, _baseAssetPoolAmount)
}

// GetOutputTwap is a free data retrieval call binding the contract method 0x6fa42ede.
//
// Solidity: function getOutputTwap(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmCaller) GetOutputTwap(opts *bind.CallOpts, _dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getOutputTwap", _dirOfBase, _baseAssetAmount)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetOutputTwap is a free data retrieval call binding the contract method 0x6fa42ede.
//
// Solidity: function getOutputTwap(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmSession) GetOutputTwap(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputTwap(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// GetOutputTwap is a free data retrieval call binding the contract method 0x6fa42ede.
//
// Solidity: function getOutputTwap(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns((uint256))
func (_Amm *AmmCallerSession) GetOutputTwap(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (Decimaldecimal, error) {
	return _Amm.Contract.GetOutputTwap(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns((uint256), (uint256))
func (_Amm *AmmCaller) GetReserve(opts *bind.CallOpts) (Decimaldecimal, Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getReserve")

	if err != nil {
		return *new(Decimaldecimal), *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)
	out1 := *abi.ConvertType(out[1], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, out1, err

}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns((uint256), (uint256))
func (_Amm *AmmSession) GetReserve() (Decimaldecimal, Decimaldecimal, error) {
	return _Amm.Contract.GetReserve(&_Amm.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns((uint256), (uint256))
func (_Amm *AmmCallerSession) GetReserve() (Decimaldecimal, Decimaldecimal, error) {
	return _Amm.Contract.GetReserve(&_Amm.CallOpts)
}

// GetSettlementPrice is a free data retrieval call binding the contract method 0xec2c0e63.
//
// Solidity: function getSettlementPrice() view returns((uint256))
func (_Amm *AmmCaller) GetSettlementPrice(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getSettlementPrice")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetSettlementPrice is a free data retrieval call binding the contract method 0xec2c0e63.
//
// Solidity: function getSettlementPrice() view returns((uint256))
func (_Amm *AmmSession) GetSettlementPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetSettlementPrice(&_Amm.CallOpts)
}

// GetSettlementPrice is a free data retrieval call binding the contract method 0xec2c0e63.
//
// Solidity: function getSettlementPrice() view returns((uint256))
func (_Amm *AmmCallerSession) GetSettlementPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetSettlementPrice(&_Amm.CallOpts)
}

// GetSnapshotLen is a free data retrieval call binding the contract method 0x0d451c8f.
//
// Solidity: function getSnapshotLen() view returns(uint256)
func (_Amm *AmmCaller) GetSnapshotLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getSnapshotLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSnapshotLen is a free data retrieval call binding the contract method 0x0d451c8f.
//
// Solidity: function getSnapshotLen() view returns(uint256)
func (_Amm *AmmSession) GetSnapshotLen() (*big.Int, error) {
	return _Amm.Contract.GetSnapshotLen(&_Amm.CallOpts)
}

// GetSnapshotLen is a free data retrieval call binding the contract method 0x0d451c8f.
//
// Solidity: function getSnapshotLen() view returns(uint256)
func (_Amm *AmmCallerSession) GetSnapshotLen() (*big.Int, error) {
	return _Amm.Contract.GetSnapshotLen(&_Amm.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0xdc76fabc.
//
// Solidity: function getSpotPrice() view returns((uint256))
func (_Amm *AmmCaller) GetSpotPrice(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getSpotPrice")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0xdc76fabc.
//
// Solidity: function getSpotPrice() view returns((uint256))
func (_Amm *AmmSession) GetSpotPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetSpotPrice(&_Amm.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0xdc76fabc.
//
// Solidity: function getSpotPrice() view returns((uint256))
func (_Amm *AmmCallerSession) GetSpotPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetSpotPrice(&_Amm.CallOpts)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xf1b5df86.
//
// Solidity: function getTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmCaller) GetTwapPrice(opts *bind.CallOpts, _intervalInSeconds *big.Int) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getTwapPrice", _intervalInSeconds)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetTwapPrice is a free data retrieval call binding the contract method 0xf1b5df86.
//
// Solidity: function getTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmSession) GetTwapPrice(_intervalInSeconds *big.Int) (Decimaldecimal, error) {
	return _Amm.Contract.GetTwapPrice(&_Amm.CallOpts, _intervalInSeconds)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xf1b5df86.
//
// Solidity: function getTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmCallerSession) GetTwapPrice(_intervalInSeconds *big.Int) (Decimaldecimal, error) {
	return _Amm.Contract.GetTwapPrice(&_Amm.CallOpts, _intervalInSeconds)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0x468f02d2.
//
// Solidity: function getUnderlyingPrice() view returns((uint256))
func (_Amm *AmmCaller) GetUnderlyingPrice(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getUnderlyingPrice")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0x468f02d2.
//
// Solidity: function getUnderlyingPrice() view returns((uint256))
func (_Amm *AmmSession) GetUnderlyingPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetUnderlyingPrice(&_Amm.CallOpts)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0x468f02d2.
//
// Solidity: function getUnderlyingPrice() view returns((uint256))
func (_Amm *AmmCallerSession) GetUnderlyingPrice() (Decimaldecimal, error) {
	return _Amm.Contract.GetUnderlyingPrice(&_Amm.CallOpts)
}

// GetUnderlyingTwapPrice is a free data retrieval call binding the contract method 0x34953249.
//
// Solidity: function getUnderlyingTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmCaller) GetUnderlyingTwapPrice(opts *bind.CallOpts, _intervalInSeconds *big.Int) (Decimaldecimal, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getUnderlyingTwapPrice", _intervalInSeconds)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// GetUnderlyingTwapPrice is a free data retrieval call binding the contract method 0x34953249.
//
// Solidity: function getUnderlyingTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmSession) GetUnderlyingTwapPrice(_intervalInSeconds *big.Int) (Decimaldecimal, error) {
	return _Amm.Contract.GetUnderlyingTwapPrice(&_Amm.CallOpts, _intervalInSeconds)
}

// GetUnderlyingTwapPrice is a free data retrieval call binding the contract method 0x34953249.
//
// Solidity: function getUnderlyingTwapPrice(uint256 _intervalInSeconds) view returns((uint256))
func (_Amm *AmmCallerSession) GetUnderlyingTwapPrice(_intervalInSeconds *big.Int) (Decimaldecimal, error) {
	return _Amm.Contract.GetUnderlyingTwapPrice(&_Amm.CallOpts, _intervalInSeconds)
}

// GlobalShutdown is a free data retrieval call binding the contract method 0x0244accf.
//
// Solidity: function globalShutdown() view returns(address)
func (_Amm *AmmCaller) GlobalShutdown(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "globalShutdown")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalShutdown is a free data retrieval call binding the contract method 0x0244accf.
//
// Solidity: function globalShutdown() view returns(address)
func (_Amm *AmmSession) GlobalShutdown() (common.Address, error) {
	return _Amm.Contract.GlobalShutdown(&_Amm.CallOpts)
}

// GlobalShutdown is a free data retrieval call binding the contract method 0x0244accf.
//
// Solidity: function globalShutdown() view returns(address)
func (_Amm *AmmCallerSession) GlobalShutdown() (common.Address, error) {
	return _Amm.Contract.GlobalShutdown(&_Amm.CallOpts)
}

// IsOverFluctuationLimit is a free data retrieval call binding the contract method 0x130234cf.
//
// Solidity: function isOverFluctuationLimit(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns(bool)
func (_Amm *AmmCaller) IsOverFluctuationLimit(opts *bind.CallOpts, _dirOfBase uint8, _baseAssetAmount Decimaldecimal) (bool, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "isOverFluctuationLimit", _dirOfBase, _baseAssetAmount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOverFluctuationLimit is a free data retrieval call binding the contract method 0x130234cf.
//
// Solidity: function isOverFluctuationLimit(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns(bool)
func (_Amm *AmmSession) IsOverFluctuationLimit(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (bool, error) {
	return _Amm.Contract.IsOverFluctuationLimit(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// IsOverFluctuationLimit is a free data retrieval call binding the contract method 0x130234cf.
//
// Solidity: function isOverFluctuationLimit(uint8 _dirOfBase, (uint256) _baseAssetAmount) view returns(bool)
func (_Amm *AmmCallerSession) IsOverFluctuationLimit(_dirOfBase uint8, _baseAssetAmount Decimaldecimal) (bool, error) {
	return _Amm.Contract.IsOverFluctuationLimit(&_Amm.CallOpts, _dirOfBase, _baseAssetAmount)
}

// IsOverSpreadLimit is a free data retrieval call binding the contract method 0x9e010362.
//
// Solidity: function isOverSpreadLimit() view returns(bool)
func (_Amm *AmmCaller) IsOverSpreadLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "isOverSpreadLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOverSpreadLimit is a free data retrieval call binding the contract method 0x9e010362.
//
// Solidity: function isOverSpreadLimit() view returns(bool)
func (_Amm *AmmSession) IsOverSpreadLimit() (bool, error) {
	return _Amm.Contract.IsOverSpreadLimit(&_Amm.CallOpts)
}

// IsOverSpreadLimit is a free data retrieval call binding the contract method 0x9e010362.
//
// Solidity: function isOverSpreadLimit() view returns(bool)
func (_Amm *AmmCallerSession) IsOverSpreadLimit() (bool, error) {
	return _Amm.Contract.IsOverSpreadLimit(&_Amm.CallOpts)
}

// NextFundingTime is a free data retrieval call binding the contract method 0xe0037a6c.
//
// Solidity: function nextFundingTime() view returns(uint256)
func (_Amm *AmmCaller) NextFundingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "nextFundingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextFundingTime is a free data retrieval call binding the contract method 0xe0037a6c.
//
// Solidity: function nextFundingTime() view returns(uint256)
func (_Amm *AmmSession) NextFundingTime() (*big.Int, error) {
	return _Amm.Contract.NextFundingTime(&_Amm.CallOpts)
}

// NextFundingTime is a free data retrieval call binding the contract method 0xe0037a6c.
//
// Solidity: function nextFundingTime() view returns(uint256)
func (_Amm *AmmCallerSession) NextFundingTime() (*big.Int, error) {
	return _Amm.Contract.NextFundingTime(&_Amm.CallOpts)
}

// Open is a free data retrieval call binding the contract method 0xfcfff16f.
//
// Solidity: function open() view returns(bool)
func (_Amm *AmmCaller) Open(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "open")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Open is a free data retrieval call binding the contract method 0xfcfff16f.
//
// Solidity: function open() view returns(bool)
func (_Amm *AmmSession) Open() (bool, error) {
	return _Amm.Contract.Open(&_Amm.CallOpts)
}

// Open is a free data retrieval call binding the contract method 0xfcfff16f.
//
// Solidity: function open() view returns(bool)
func (_Amm *AmmCallerSession) Open() (bool, error) {
	return _Amm.Contract.Open(&_Amm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Amm *AmmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Amm *AmmSession) Owner() (common.Address, error) {
	return _Amm.Contract.Owner(&_Amm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Amm *AmmCallerSession) Owner() (common.Address, error) {
	return _Amm.Contract.Owner(&_Amm.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Amm *AmmCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Amm *AmmSession) PriceFeed() (common.Address, error) {
	return _Amm.Contract.PriceFeed(&_Amm.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Amm *AmmCallerSession) PriceFeed() (common.Address, error) {
	return _Amm.Contract.PriceFeed(&_Amm.CallOpts)
}

// PriceFeedKey is a free data retrieval call binding the contract method 0x58a4c3dc.
//
// Solidity: function priceFeedKey() view returns(bytes32)
func (_Amm *AmmCaller) PriceFeedKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "priceFeedKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceFeedKey is a free data retrieval call binding the contract method 0x58a4c3dc.
//
// Solidity: function priceFeedKey() view returns(bytes32)
func (_Amm *AmmSession) PriceFeedKey() ([32]byte, error) {
	return _Amm.Contract.PriceFeedKey(&_Amm.CallOpts)
}

// PriceFeedKey is a free data retrieval call binding the contract method 0x58a4c3dc.
//
// Solidity: function priceFeedKey() view returns(bytes32)
func (_Amm *AmmCallerSession) PriceFeedKey() ([32]byte, error) {
	return _Amm.Contract.PriceFeedKey(&_Amm.CallOpts)
}

// QuoteAsset is a free data retrieval call binding the contract method 0xfdf262b7.
//
// Solidity: function quoteAsset() view returns(address)
func (_Amm *AmmCaller) QuoteAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "quoteAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteAsset is a free data retrieval call binding the contract method 0xfdf262b7.
//
// Solidity: function quoteAsset() view returns(address)
func (_Amm *AmmSession) QuoteAsset() (common.Address, error) {
	return _Amm.Contract.QuoteAsset(&_Amm.CallOpts)
}

// QuoteAsset is a free data retrieval call binding the contract method 0xfdf262b7.
//
// Solidity: function quoteAsset() view returns(address)
func (_Amm *AmmCallerSession) QuoteAsset() (common.Address, error) {
	return _Amm.Contract.QuoteAsset(&_Amm.CallOpts)
}

// QuoteAssetReserve is a free data retrieval call binding the contract method 0xe4bc2eb9.
//
// Solidity: function quoteAssetReserve() view returns(uint256 d)
func (_Amm *AmmCaller) QuoteAssetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "quoteAssetReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteAssetReserve is a free data retrieval call binding the contract method 0xe4bc2eb9.
//
// Solidity: function quoteAssetReserve() view returns(uint256 d)
func (_Amm *AmmSession) QuoteAssetReserve() (*big.Int, error) {
	return _Amm.Contract.QuoteAssetReserve(&_Amm.CallOpts)
}

// QuoteAssetReserve is a free data retrieval call binding the contract method 0xe4bc2eb9.
//
// Solidity: function quoteAssetReserve() view returns(uint256 d)
func (_Amm *AmmCallerSession) QuoteAssetReserve() (*big.Int, error) {
	return _Amm.Contract.QuoteAssetReserve(&_Amm.CallOpts)
}

// ReserveSnapshots is a free data retrieval call binding the contract method 0x21e00985.
//
// Solidity: function reserveSnapshots(uint256 ) view returns((uint256) quoteAssetReserve, (uint256) baseAssetReserve, uint256 timestamp, uint256 blockNumber)
func (_Amm *AmmCaller) ReserveSnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	QuoteAssetReserve Decimaldecimal
	BaseAssetReserve  Decimaldecimal
	Timestamp         *big.Int
	BlockNumber       *big.Int
}, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "reserveSnapshots", arg0)

	outstruct := new(struct {
		QuoteAssetReserve Decimaldecimal
		BaseAssetReserve  Decimaldecimal
		Timestamp         *big.Int
		BlockNumber       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.QuoteAssetReserve = *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)
	outstruct.BaseAssetReserve = *abi.ConvertType(out[1], new(Decimaldecimal)).(*Decimaldecimal)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReserveSnapshots is a free data retrieval call binding the contract method 0x21e00985.
//
// Solidity: function reserveSnapshots(uint256 ) view returns((uint256) quoteAssetReserve, (uint256) baseAssetReserve, uint256 timestamp, uint256 blockNumber)
func (_Amm *AmmSession) ReserveSnapshots(arg0 *big.Int) (struct {
	QuoteAssetReserve Decimaldecimal
	BaseAssetReserve  Decimaldecimal
	Timestamp         *big.Int
	BlockNumber       *big.Int
}, error) {
	return _Amm.Contract.ReserveSnapshots(&_Amm.CallOpts, arg0)
}

// ReserveSnapshots is a free data retrieval call binding the contract method 0x21e00985.
//
// Solidity: function reserveSnapshots(uint256 ) view returns((uint256) quoteAssetReserve, (uint256) baseAssetReserve, uint256 timestamp, uint256 blockNumber)
func (_Amm *AmmCallerSession) ReserveSnapshots(arg0 *big.Int) (struct {
	QuoteAssetReserve Decimaldecimal
	BaseAssetReserve  Decimaldecimal
	Timestamp         *big.Int
	BlockNumber       *big.Int
}, error) {
	return _Amm.Contract.ReserveSnapshots(&_Amm.CallOpts, arg0)
}

// SpotPriceTwapInterval is a free data retrieval call binding the contract method 0x5f1ba1fd.
//
// Solidity: function spotPriceTwapInterval() view returns(uint256)
func (_Amm *AmmCaller) SpotPriceTwapInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "spotPriceTwapInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpotPriceTwapInterval is a free data retrieval call binding the contract method 0x5f1ba1fd.
//
// Solidity: function spotPriceTwapInterval() view returns(uint256)
func (_Amm *AmmSession) SpotPriceTwapInterval() (*big.Int, error) {
	return _Amm.Contract.SpotPriceTwapInterval(&_Amm.CallOpts)
}

// SpotPriceTwapInterval is a free data retrieval call binding the contract method 0x5f1ba1fd.
//
// Solidity: function spotPriceTwapInterval() view returns(uint256)
func (_Amm *AmmCallerSession) SpotPriceTwapInterval() (*big.Int, error) {
	return _Amm.Contract.SpotPriceTwapInterval(&_Amm.CallOpts)
}

// SpreadRatio is a free data retrieval call binding the contract method 0x6febdd50.
//
// Solidity: function spreadRatio() view returns(uint256 d)
func (_Amm *AmmCaller) SpreadRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "spreadRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadRatio is a free data retrieval call binding the contract method 0x6febdd50.
//
// Solidity: function spreadRatio() view returns(uint256 d)
func (_Amm *AmmSession) SpreadRatio() (*big.Int, error) {
	return _Amm.Contract.SpreadRatio(&_Amm.CallOpts)
}

// SpreadRatio is a free data retrieval call binding the contract method 0x6febdd50.
//
// Solidity: function spreadRatio() view returns(uint256 d)
func (_Amm *AmmCallerSession) SpreadRatio() (*big.Int, error) {
	return _Amm.Contract.SpreadRatio(&_Amm.CallOpts)
}

// TollAmount is a free data retrieval call binding the contract method 0x2d666e38.
//
// Solidity: function tollAmount() view returns(uint256 d)
func (_Amm *AmmCaller) TollAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "tollAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TollAmount is a free data retrieval call binding the contract method 0x2d666e38.
//
// Solidity: function tollAmount() view returns(uint256 d)
func (_Amm *AmmSession) TollAmount() (*big.Int, error) {
	return _Amm.Contract.TollAmount(&_Amm.CallOpts)
}

// TollAmount is a free data retrieval call binding the contract method 0x2d666e38.
//
// Solidity: function tollAmount() view returns(uint256 d)
func (_Amm *AmmCallerSession) TollAmount() (*big.Int, error) {
	return _Amm.Contract.TollAmount(&_Amm.CallOpts)
}

// TollRatio is a free data retrieval call binding the contract method 0x6baccaba.
//
// Solidity: function tollRatio() view returns(uint256 d)
func (_Amm *AmmCaller) TollRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "tollRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TollRatio is a free data retrieval call binding the contract method 0x6baccaba.
//
// Solidity: function tollRatio() view returns(uint256 d)
func (_Amm *AmmSession) TollRatio() (*big.Int, error) {
	return _Amm.Contract.TollRatio(&_Amm.CallOpts)
}

// TollRatio is a free data retrieval call binding the contract method 0x6baccaba.
//
// Solidity: function tollRatio() view returns(uint256 d)
func (_Amm *AmmCallerSession) TollRatio() (*big.Int, error) {
	return _Amm.Contract.TollRatio(&_Amm.CallOpts)
}

// TotalPositionSize is a free data retrieval call binding the contract method 0x4cb876f2.
//
// Solidity: function totalPositionSize() view returns(int256 d)
func (_Amm *AmmCaller) TotalPositionSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "totalPositionSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPositionSize is a free data retrieval call binding the contract method 0x4cb876f2.
//
// Solidity: function totalPositionSize() view returns(int256 d)
func (_Amm *AmmSession) TotalPositionSize() (*big.Int, error) {
	return _Amm.Contract.TotalPositionSize(&_Amm.CallOpts)
}

// TotalPositionSize is a free data retrieval call binding the contract method 0x4cb876f2.
//
// Solidity: function totalPositionSize() view returns(int256 d)
func (_Amm *AmmCallerSession) TotalPositionSize() (*big.Int, error) {
	return _Amm.Contract.TotalPositionSize(&_Amm.CallOpts)
}

// TradeLimitRatio is a free data retrieval call binding the contract method 0x8f40d932.
//
// Solidity: function tradeLimitRatio() view returns(uint256 d)
func (_Amm *AmmCaller) TradeLimitRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "tradeLimitRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradeLimitRatio is a free data retrieval call binding the contract method 0x8f40d932.
//
// Solidity: function tradeLimitRatio() view returns(uint256 d)
func (_Amm *AmmSession) TradeLimitRatio() (*big.Int, error) {
	return _Amm.Contract.TradeLimitRatio(&_Amm.CallOpts)
}

// TradeLimitRatio is a free data retrieval call binding the contract method 0x8f40d932.
//
// Solidity: function tradeLimitRatio() view returns(uint256 d)
func (_Amm *AmmCallerSession) TradeLimitRatio() (*big.Int, error) {
	return _Amm.Contract.TradeLimitRatio(&_Amm.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x620be067.
//
// Solidity: function initialize(uint256 _quoteAssetReserve, uint256 _baseAssetReserve, uint256 _tradeLimitRatio, uint256 _fundingPeriod, address _priceFeed, bytes32 _priceFeedKey, address _quoteAsset, uint256 _fluctuationLimitRatio, uint256 _tollRatio, uint256 _spreadRatio) returns()
func (_Amm *AmmTransactor) Initialize(opts *bind.TransactOpts, _quoteAssetReserve *big.Int, _baseAssetReserve *big.Int, _tradeLimitRatio *big.Int, _fundingPeriod *big.Int, _priceFeed common.Address, _priceFeedKey [32]byte, _quoteAsset common.Address, _fluctuationLimitRatio *big.Int, _tollRatio *big.Int, _spreadRatio *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "initialize", _quoteAssetReserve, _baseAssetReserve, _tradeLimitRatio, _fundingPeriod, _priceFeed, _priceFeedKey, _quoteAsset, _fluctuationLimitRatio, _tollRatio, _spreadRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0x620be067.
//
// Solidity: function initialize(uint256 _quoteAssetReserve, uint256 _baseAssetReserve, uint256 _tradeLimitRatio, uint256 _fundingPeriod, address _priceFeed, bytes32 _priceFeedKey, address _quoteAsset, uint256 _fluctuationLimitRatio, uint256 _tollRatio, uint256 _spreadRatio) returns()
func (_Amm *AmmSession) Initialize(_quoteAssetReserve *big.Int, _baseAssetReserve *big.Int, _tradeLimitRatio *big.Int, _fundingPeriod *big.Int, _priceFeed common.Address, _priceFeedKey [32]byte, _quoteAsset common.Address, _fluctuationLimitRatio *big.Int, _tollRatio *big.Int, _spreadRatio *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Initialize(&_Amm.TransactOpts, _quoteAssetReserve, _baseAssetReserve, _tradeLimitRatio, _fundingPeriod, _priceFeed, _priceFeedKey, _quoteAsset, _fluctuationLimitRatio, _tollRatio, _spreadRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0x620be067.
//
// Solidity: function initialize(uint256 _quoteAssetReserve, uint256 _baseAssetReserve, uint256 _tradeLimitRatio, uint256 _fundingPeriod, address _priceFeed, bytes32 _priceFeedKey, address _quoteAsset, uint256 _fluctuationLimitRatio, uint256 _tollRatio, uint256 _spreadRatio) returns()
func (_Amm *AmmTransactorSession) Initialize(_quoteAssetReserve *big.Int, _baseAssetReserve *big.Int, _tradeLimitRatio *big.Int, _fundingPeriod *big.Int, _priceFeed common.Address, _priceFeedKey [32]byte, _quoteAsset common.Address, _fluctuationLimitRatio *big.Int, _tollRatio *big.Int, _spreadRatio *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.Initialize(&_Amm.TransactOpts, _quoteAssetReserve, _baseAssetReserve, _tradeLimitRatio, _fundingPeriod, _priceFeed, _priceFeedKey, _quoteAsset, _fluctuationLimitRatio, _tollRatio, _spreadRatio)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Amm *AmmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Amm *AmmSession) RenounceOwnership() (*types.Transaction, error) {
	return _Amm.Contract.RenounceOwnership(&_Amm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Amm *AmmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Amm.Contract.RenounceOwnership(&_Amm.TransactOpts)
}

// SetCap is a paid mutator transaction binding the contract method 0x0dd68c70.
//
// Solidity: function setCap((uint256) _maxHoldingBaseAsset, (uint256) _openInterestNotionalCap) returns()
func (_Amm *AmmTransactor) SetCap(opts *bind.TransactOpts, _maxHoldingBaseAsset Decimaldecimal, _openInterestNotionalCap Decimaldecimal) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setCap", _maxHoldingBaseAsset, _openInterestNotionalCap)
}

// SetCap is a paid mutator transaction binding the contract method 0x0dd68c70.
//
// Solidity: function setCap((uint256) _maxHoldingBaseAsset, (uint256) _openInterestNotionalCap) returns()
func (_Amm *AmmSession) SetCap(_maxHoldingBaseAsset Decimaldecimal, _openInterestNotionalCap Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetCap(&_Amm.TransactOpts, _maxHoldingBaseAsset, _openInterestNotionalCap)
}

// SetCap is a paid mutator transaction binding the contract method 0x0dd68c70.
//
// Solidity: function setCap((uint256) _maxHoldingBaseAsset, (uint256) _openInterestNotionalCap) returns()
func (_Amm *AmmTransactorSession) SetCap(_maxHoldingBaseAsset Decimaldecimal, _openInterestNotionalCap Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetCap(&_Amm.TransactOpts, _maxHoldingBaseAsset, _openInterestNotionalCap)
}

// SetCounterParty is a paid mutator transaction binding the contract method 0x9ece77c8.
//
// Solidity: function setCounterParty(address _counterParty) returns()
func (_Amm *AmmTransactor) SetCounterParty(opts *bind.TransactOpts, _counterParty common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setCounterParty", _counterParty)
}

// SetCounterParty is a paid mutator transaction binding the contract method 0x9ece77c8.
//
// Solidity: function setCounterParty(address _counterParty) returns()
func (_Amm *AmmSession) SetCounterParty(_counterParty common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetCounterParty(&_Amm.TransactOpts, _counterParty)
}

// SetCounterParty is a paid mutator transaction binding the contract method 0x9ece77c8.
//
// Solidity: function setCounterParty(address _counterParty) returns()
func (_Amm *AmmTransactorSession) SetCounterParty(_counterParty common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetCounterParty(&_Amm.TransactOpts, _counterParty)
}

// SetFluctuationLimitRatio is a paid mutator transaction binding the contract method 0xa8f8be4e.
//
// Solidity: function setFluctuationLimitRatio((uint256) _fluctuationLimitRatio) returns()
func (_Amm *AmmTransactor) SetFluctuationLimitRatio(opts *bind.TransactOpts, _fluctuationLimitRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setFluctuationLimitRatio", _fluctuationLimitRatio)
}

// SetFluctuationLimitRatio is a paid mutator transaction binding the contract method 0xa8f8be4e.
//
// Solidity: function setFluctuationLimitRatio((uint256) _fluctuationLimitRatio) returns()
func (_Amm *AmmSession) SetFluctuationLimitRatio(_fluctuationLimitRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetFluctuationLimitRatio(&_Amm.TransactOpts, _fluctuationLimitRatio)
}

// SetFluctuationLimitRatio is a paid mutator transaction binding the contract method 0xa8f8be4e.
//
// Solidity: function setFluctuationLimitRatio((uint256) _fluctuationLimitRatio) returns()
func (_Amm *AmmTransactorSession) SetFluctuationLimitRatio(_fluctuationLimitRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetFluctuationLimitRatio(&_Amm.TransactOpts, _fluctuationLimitRatio)
}

// SetGlobalShutdown is a paid mutator transaction binding the contract method 0xb2ed32c7.
//
// Solidity: function setGlobalShutdown(address _globalShutdown) returns()
func (_Amm *AmmTransactor) SetGlobalShutdown(opts *bind.TransactOpts, _globalShutdown common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setGlobalShutdown", _globalShutdown)
}

// SetGlobalShutdown is a paid mutator transaction binding the contract method 0xb2ed32c7.
//
// Solidity: function setGlobalShutdown(address _globalShutdown) returns()
func (_Amm *AmmSession) SetGlobalShutdown(_globalShutdown common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetGlobalShutdown(&_Amm.TransactOpts, _globalShutdown)
}

// SetGlobalShutdown is a paid mutator transaction binding the contract method 0xb2ed32c7.
//
// Solidity: function setGlobalShutdown(address _globalShutdown) returns()
func (_Amm *AmmTransactorSession) SetGlobalShutdown(_globalShutdown common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetGlobalShutdown(&_Amm.TransactOpts, _globalShutdown)
}

// SetOpen is a paid mutator transaction binding the contract method 0x6fdca5e0.
//
// Solidity: function setOpen(bool _open) returns()
func (_Amm *AmmTransactor) SetOpen(opts *bind.TransactOpts, _open bool) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setOpen", _open)
}

// SetOpen is a paid mutator transaction binding the contract method 0x6fdca5e0.
//
// Solidity: function setOpen(bool _open) returns()
func (_Amm *AmmSession) SetOpen(_open bool) (*types.Transaction, error) {
	return _Amm.Contract.SetOpen(&_Amm.TransactOpts, _open)
}

// SetOpen is a paid mutator transaction binding the contract method 0x6fdca5e0.
//
// Solidity: function setOpen(bool _open) returns()
func (_Amm *AmmTransactorSession) SetOpen(_open bool) (*types.Transaction, error) {
	return _Amm.Contract.SetOpen(&_Amm.TransactOpts, _open)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Amm *AmmTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Amm *AmmSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetOwner(&_Amm.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Amm *AmmTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetOwner(&_Amm.TransactOpts, newOwner)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Amm *AmmTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Amm *AmmSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetPriceFeed(&_Amm.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Amm *AmmTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Amm.Contract.SetPriceFeed(&_Amm.TransactOpts, _priceFeed)
}

// SetSpotPriceTwapInterval is a paid mutator transaction binding the contract method 0x62e7a176.
//
// Solidity: function setSpotPriceTwapInterval(uint256 _interval) returns()
func (_Amm *AmmTransactor) SetSpotPriceTwapInterval(opts *bind.TransactOpts, _interval *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setSpotPriceTwapInterval", _interval)
}

// SetSpotPriceTwapInterval is a paid mutator transaction binding the contract method 0x62e7a176.
//
// Solidity: function setSpotPriceTwapInterval(uint256 _interval) returns()
func (_Amm *AmmSession) SetSpotPriceTwapInterval(_interval *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.SetSpotPriceTwapInterval(&_Amm.TransactOpts, _interval)
}

// SetSpotPriceTwapInterval is a paid mutator transaction binding the contract method 0x62e7a176.
//
// Solidity: function setSpotPriceTwapInterval(uint256 _interval) returns()
func (_Amm *AmmTransactorSession) SetSpotPriceTwapInterval(_interval *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.SetSpotPriceTwapInterval(&_Amm.TransactOpts, _interval)
}

// SetSpreadRatio is a paid mutator transaction binding the contract method 0x524f15a1.
//
// Solidity: function setSpreadRatio((uint256) _spreadRatio) returns()
func (_Amm *AmmTransactor) SetSpreadRatio(opts *bind.TransactOpts, _spreadRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setSpreadRatio", _spreadRatio)
}

// SetSpreadRatio is a paid mutator transaction binding the contract method 0x524f15a1.
//
// Solidity: function setSpreadRatio((uint256) _spreadRatio) returns()
func (_Amm *AmmSession) SetSpreadRatio(_spreadRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetSpreadRatio(&_Amm.TransactOpts, _spreadRatio)
}

// SetSpreadRatio is a paid mutator transaction binding the contract method 0x524f15a1.
//
// Solidity: function setSpreadRatio((uint256) _spreadRatio) returns()
func (_Amm *AmmTransactorSession) SetSpreadRatio(_spreadRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetSpreadRatio(&_Amm.TransactOpts, _spreadRatio)
}

// SetTollRatio is a paid mutator transaction binding the contract method 0x7ec3246d.
//
// Solidity: function setTollRatio((uint256) _tollRatio) returns()
func (_Amm *AmmTransactor) SetTollRatio(opts *bind.TransactOpts, _tollRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "setTollRatio", _tollRatio)
}

// SetTollRatio is a paid mutator transaction binding the contract method 0x7ec3246d.
//
// Solidity: function setTollRatio((uint256) _tollRatio) returns()
func (_Amm *AmmSession) SetTollRatio(_tollRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetTollRatio(&_Amm.TransactOpts, _tollRatio)
}

// SetTollRatio is a paid mutator transaction binding the contract method 0x7ec3246d.
//
// Solidity: function setTollRatio((uint256) _tollRatio) returns()
func (_Amm *AmmTransactorSession) SetTollRatio(_tollRatio Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SetTollRatio(&_Amm.TransactOpts, _tollRatio)
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns((int256))
func (_Amm *AmmTransactor) SettleFunding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "settleFunding")
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns((int256))
func (_Amm *AmmSession) SettleFunding() (*types.Transaction, error) {
	return _Amm.Contract.SettleFunding(&_Amm.TransactOpts)
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns((int256))
func (_Amm *AmmTransactorSession) SettleFunding() (*types.Transaction, error) {
	return _Amm.Contract.SettleFunding(&_Amm.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Amm *AmmTransactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Amm *AmmSession) Shutdown() (*types.Transaction, error) {
	return _Amm.Contract.Shutdown(&_Amm.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Amm *AmmTransactorSession) Shutdown() (*types.Transaction, error) {
	return _Amm.Contract.Shutdown(&_Amm.TransactOpts)
}

// SwapInput is a paid mutator transaction binding the contract method 0x75df6389.
//
// Solidity: function swapInput(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _baseAssetAmountLimit, bool _canOverFluctuationLimit) returns((uint256))
func (_Amm *AmmTransactor) SwapInput(opts *bind.TransactOpts, _dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _canOverFluctuationLimit bool) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "swapInput", _dirOfQuote, _quoteAssetAmount, _baseAssetAmountLimit, _canOverFluctuationLimit)
}

// SwapInput is a paid mutator transaction binding the contract method 0x75df6389.
//
// Solidity: function swapInput(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _baseAssetAmountLimit, bool _canOverFluctuationLimit) returns((uint256))
func (_Amm *AmmSession) SwapInput(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _canOverFluctuationLimit bool) (*types.Transaction, error) {
	return _Amm.Contract.SwapInput(&_Amm.TransactOpts, _dirOfQuote, _quoteAssetAmount, _baseAssetAmountLimit, _canOverFluctuationLimit)
}

// SwapInput is a paid mutator transaction binding the contract method 0x75df6389.
//
// Solidity: function swapInput(uint8 _dirOfQuote, (uint256) _quoteAssetAmount, (uint256) _baseAssetAmountLimit, bool _canOverFluctuationLimit) returns((uint256))
func (_Amm *AmmTransactorSession) SwapInput(_dirOfQuote uint8, _quoteAssetAmount Decimaldecimal, _baseAssetAmountLimit Decimaldecimal, _canOverFluctuationLimit bool) (*types.Transaction, error) {
	return _Amm.Contract.SwapInput(&_Amm.TransactOpts, _dirOfQuote, _quoteAssetAmount, _baseAssetAmountLimit, _canOverFluctuationLimit)
}

// SwapOutput is a paid mutator transaction binding the contract method 0xd71ec2ad.
//
// Solidity: function swapOutput(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetAmountLimit) returns((uint256))
func (_Amm *AmmTransactor) SwapOutput(opts *bind.TransactOpts, _dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "swapOutput", _dirOfBase, _baseAssetAmount, _quoteAssetAmountLimit)
}

// SwapOutput is a paid mutator transaction binding the contract method 0xd71ec2ad.
//
// Solidity: function swapOutput(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetAmountLimit) returns((uint256))
func (_Amm *AmmSession) SwapOutput(_dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SwapOutput(&_Amm.TransactOpts, _dirOfBase, _baseAssetAmount, _quoteAssetAmountLimit)
}

// SwapOutput is a paid mutator transaction binding the contract method 0xd71ec2ad.
//
// Solidity: function swapOutput(uint8 _dirOfBase, (uint256) _baseAssetAmount, (uint256) _quoteAssetAmountLimit) returns((uint256))
func (_Amm *AmmTransactorSession) SwapOutput(_dirOfBase uint8, _baseAssetAmount Decimaldecimal, _quoteAssetAmountLimit Decimaldecimal) (*types.Transaction, error) {
	return _Amm.Contract.SwapOutput(&_Amm.TransactOpts, _dirOfBase, _baseAssetAmount, _quoteAssetAmountLimit)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Amm *AmmTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Amm *AmmSession) UpdateOwner() (*types.Transaction, error) {
	return _Amm.Contract.UpdateOwner(&_Amm.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Amm *AmmTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _Amm.Contract.UpdateOwner(&_Amm.TransactOpts)
}

// AmmCapChangedIterator is returned from FilterCapChanged and is used to iterate over the raw logs and unpacked data for CapChanged events raised by the Amm contract.
type AmmCapChangedIterator struct {
	Event *AmmCapChanged // Event containing the contract specifics and raw log

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
func (it *AmmCapChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmCapChanged)
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
		it.Event = new(AmmCapChanged)
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
func (it *AmmCapChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmCapChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmCapChanged represents a CapChanged event raised by the Amm contract.
type AmmCapChanged struct {
	MaxHoldingBaseAsset     *big.Int
	OpenInterestNotionalCap *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCapChanged is a free log retrieval operation binding the contract event 0x7338f3784ceb8f9456bac0c4a69f1c6354dc325fa6455e3e3f6a8a9bf9249a7c.
//
// Solidity: event CapChanged(uint256 maxHoldingBaseAsset, uint256 openInterestNotionalCap)
func (_Amm *AmmFilterer) FilterCapChanged(opts *bind.FilterOpts) (*AmmCapChangedIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "CapChanged")
	if err != nil {
		return nil, err
	}
	return &AmmCapChangedIterator{contract: _Amm.contract, event: "CapChanged", logs: logs, sub: sub}, nil
}

// WatchCapChanged is a free log subscription operation binding the contract event 0x7338f3784ceb8f9456bac0c4a69f1c6354dc325fa6455e3e3f6a8a9bf9249a7c.
//
// Solidity: event CapChanged(uint256 maxHoldingBaseAsset, uint256 openInterestNotionalCap)
func (_Amm *AmmFilterer) WatchCapChanged(opts *bind.WatchOpts, sink chan<- *AmmCapChanged) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "CapChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmCapChanged)
				if err := _Amm.contract.UnpackLog(event, "CapChanged", log); err != nil {
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

// ParseCapChanged is a log parse operation binding the contract event 0x7338f3784ceb8f9456bac0c4a69f1c6354dc325fa6455e3e3f6a8a9bf9249a7c.
//
// Solidity: event CapChanged(uint256 maxHoldingBaseAsset, uint256 openInterestNotionalCap)
func (_Amm *AmmFilterer) ParseCapChanged(log types.Log) (*AmmCapChanged, error) {
	event := new(AmmCapChanged)
	if err := _Amm.contract.UnpackLog(event, "CapChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmFundingRateUpdatedIterator is returned from FilterFundingRateUpdated and is used to iterate over the raw logs and unpacked data for FundingRateUpdated events raised by the Amm contract.
type AmmFundingRateUpdatedIterator struct {
	Event *AmmFundingRateUpdated // Event containing the contract specifics and raw log

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
func (it *AmmFundingRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmFundingRateUpdated)
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
		it.Event = new(AmmFundingRateUpdated)
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
func (it *AmmFundingRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmFundingRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmFundingRateUpdated represents a FundingRateUpdated event raised by the Amm contract.
type AmmFundingRateUpdated struct {
	Rate            *big.Int
	UnderlyingPrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundingRateUpdated is a free log retrieval operation binding the contract event 0xd2805fe76d30598332a67c1061cee82e2e102b0f59f5457b1729bce028a054a0.
//
// Solidity: event FundingRateUpdated(int256 rate, uint256 underlyingPrice)
func (_Amm *AmmFilterer) FilterFundingRateUpdated(opts *bind.FilterOpts) (*AmmFundingRateUpdatedIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "FundingRateUpdated")
	if err != nil {
		return nil, err
	}
	return &AmmFundingRateUpdatedIterator{contract: _Amm.contract, event: "FundingRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFundingRateUpdated is a free log subscription operation binding the contract event 0xd2805fe76d30598332a67c1061cee82e2e102b0f59f5457b1729bce028a054a0.
//
// Solidity: event FundingRateUpdated(int256 rate, uint256 underlyingPrice)
func (_Amm *AmmFilterer) WatchFundingRateUpdated(opts *bind.WatchOpts, sink chan<- *AmmFundingRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "FundingRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmFundingRateUpdated)
				if err := _Amm.contract.UnpackLog(event, "FundingRateUpdated", log); err != nil {
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

// ParseFundingRateUpdated is a log parse operation binding the contract event 0xd2805fe76d30598332a67c1061cee82e2e102b0f59f5457b1729bce028a054a0.
//
// Solidity: event FundingRateUpdated(int256 rate, uint256 underlyingPrice)
func (_Amm *AmmFilterer) ParseFundingRateUpdated(log types.Log) (*AmmFundingRateUpdated, error) {
	event := new(AmmFundingRateUpdated)
	if err := _Amm.contract.UnpackLog(event, "FundingRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the Amm contract.
type AmmLiquidityChangedIterator struct {
	Event *AmmLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *AmmLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmLiquidityChanged)
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
		it.Event = new(AmmLiquidityChanged)
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
func (it *AmmLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmLiquidityChanged represents a LiquidityChanged event raised by the Amm contract.
type AmmLiquidityChanged struct {
	QuoteReserve       *big.Int
	BaseReserve        *big.Int
	CumulativeNotional *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0xeba648fc83514e32ba75e27696ee5fa3610ebc7a77460a3ac3e8724ffa1be4cc.
//
// Solidity: event LiquidityChanged(uint256 quoteReserve, uint256 baseReserve, int256 cumulativeNotional)
func (_Amm *AmmFilterer) FilterLiquidityChanged(opts *bind.FilterOpts) (*AmmLiquidityChangedIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "LiquidityChanged")
	if err != nil {
		return nil, err
	}
	return &AmmLiquidityChangedIterator{contract: _Amm.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0xeba648fc83514e32ba75e27696ee5fa3610ebc7a77460a3ac3e8724ffa1be4cc.
//
// Solidity: event LiquidityChanged(uint256 quoteReserve, uint256 baseReserve, int256 cumulativeNotional)
func (_Amm *AmmFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *AmmLiquidityChanged) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "LiquidityChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmLiquidityChanged)
				if err := _Amm.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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

// ParseLiquidityChanged is a log parse operation binding the contract event 0xeba648fc83514e32ba75e27696ee5fa3610ebc7a77460a3ac3e8724ffa1be4cc.
//
// Solidity: event LiquidityChanged(uint256 quoteReserve, uint256 baseReserve, int256 cumulativeNotional)
func (_Amm *AmmFilterer) ParseLiquidityChanged(log types.Log) (*AmmLiquidityChanged, error) {
	event := new(AmmLiquidityChanged)
	if err := _Amm.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Amm contract.
type AmmOwnershipTransferredIterator struct {
	Event *AmmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AmmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmOwnershipTransferred)
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
		it.Event = new(AmmOwnershipTransferred)
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
func (it *AmmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmOwnershipTransferred represents a OwnershipTransferred event raised by the Amm contract.
type AmmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Amm *AmmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AmmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Amm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AmmOwnershipTransferredIterator{contract: _Amm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Amm *AmmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AmmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Amm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmOwnershipTransferred)
				if err := _Amm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Amm *AmmFilterer) ParseOwnershipTransferred(log types.Log) (*AmmOwnershipTransferred, error) {
	event := new(AmmOwnershipTransferred)
	if err := _Amm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmPriceFeedUpdatedIterator is returned from FilterPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedUpdated events raised by the Amm contract.
type AmmPriceFeedUpdatedIterator struct {
	Event *AmmPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *AmmPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmPriceFeedUpdated)
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
		it.Event = new(AmmPriceFeedUpdated)
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
func (it *AmmPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmPriceFeedUpdated represents a PriceFeedUpdated event raised by the Amm contract.
type AmmPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdated is a free log retrieval operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address priceFeed)
func (_Amm *AmmFilterer) FilterPriceFeedUpdated(opts *bind.FilterOpts) (*AmmPriceFeedUpdatedIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return &AmmPriceFeedUpdatedIterator{contract: _Amm.contract, event: "PriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdated is a free log subscription operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address priceFeed)
func (_Amm *AmmFilterer) WatchPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *AmmPriceFeedUpdated) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmPriceFeedUpdated)
				if err := _Amm.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
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

// ParsePriceFeedUpdated is a log parse operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address priceFeed)
func (_Amm *AmmFilterer) ParsePriceFeedUpdated(log types.Log) (*AmmPriceFeedUpdated, error) {
	event := new(AmmPriceFeedUpdated)
	if err := _Amm.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmReserveSnapshottedIterator is returned from FilterReserveSnapshotted and is used to iterate over the raw logs and unpacked data for ReserveSnapshotted events raised by the Amm contract.
type AmmReserveSnapshottedIterator struct {
	Event *AmmReserveSnapshotted // Event containing the contract specifics and raw log

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
func (it *AmmReserveSnapshottedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmReserveSnapshotted)
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
		it.Event = new(AmmReserveSnapshotted)
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
func (it *AmmReserveSnapshottedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmReserveSnapshottedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmReserveSnapshotted represents a ReserveSnapshotted event raised by the Amm contract.
type AmmReserveSnapshotted struct {
	QuoteAssetReserve *big.Int
	BaseAssetReserve  *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterReserveSnapshotted is a free log retrieval operation binding the contract event 0x3a3348362552c3897fd1f06a3233519ebd8bd76ad6e99a418a9741155fe90515.
//
// Solidity: event ReserveSnapshotted(uint256 quoteAssetReserve, uint256 baseAssetReserve, uint256 timestamp)
func (_Amm *AmmFilterer) FilterReserveSnapshotted(opts *bind.FilterOpts) (*AmmReserveSnapshottedIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "ReserveSnapshotted")
	if err != nil {
		return nil, err
	}
	return &AmmReserveSnapshottedIterator{contract: _Amm.contract, event: "ReserveSnapshotted", logs: logs, sub: sub}, nil
}

// WatchReserveSnapshotted is a free log subscription operation binding the contract event 0x3a3348362552c3897fd1f06a3233519ebd8bd76ad6e99a418a9741155fe90515.
//
// Solidity: event ReserveSnapshotted(uint256 quoteAssetReserve, uint256 baseAssetReserve, uint256 timestamp)
func (_Amm *AmmFilterer) WatchReserveSnapshotted(opts *bind.WatchOpts, sink chan<- *AmmReserveSnapshotted) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "ReserveSnapshotted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmReserveSnapshotted)
				if err := _Amm.contract.UnpackLog(event, "ReserveSnapshotted", log); err != nil {
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

// ParseReserveSnapshotted is a log parse operation binding the contract event 0x3a3348362552c3897fd1f06a3233519ebd8bd76ad6e99a418a9741155fe90515.
//
// Solidity: event ReserveSnapshotted(uint256 quoteAssetReserve, uint256 baseAssetReserve, uint256 timestamp)
func (_Amm *AmmFilterer) ParseReserveSnapshotted(log types.Log) (*AmmReserveSnapshotted, error) {
	event := new(AmmReserveSnapshotted)
	if err := _Amm.contract.UnpackLog(event, "ReserveSnapshotted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Amm contract.
type AmmShutdownIterator struct {
	Event *AmmShutdown // Event containing the contract specifics and raw log

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
func (it *AmmShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmShutdown)
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
		it.Event = new(AmmShutdown)
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
func (it *AmmShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmShutdown represents a Shutdown event raised by the Amm contract.
type AmmShutdown struct {
	SettlementPrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 settlementPrice)
func (_Amm *AmmFilterer) FilterShutdown(opts *bind.FilterOpts) (*AmmShutdownIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &AmmShutdownIterator{contract: _Amm.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 settlementPrice)
func (_Amm *AmmFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *AmmShutdown) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmShutdown)
				if err := _Amm.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 settlementPrice)
func (_Amm *AmmFilterer) ParseShutdown(log types.Log) (*AmmShutdown, error) {
	event := new(AmmShutdown)
	if err := _Amm.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmSwapInputIterator is returned from FilterSwapInput and is used to iterate over the raw logs and unpacked data for SwapInput events raised by the Amm contract.
type AmmSwapInputIterator struct {
	Event *AmmSwapInput // Event containing the contract specifics and raw log

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
func (it *AmmSwapInputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmSwapInput)
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
		it.Event = new(AmmSwapInput)
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
func (it *AmmSwapInputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmSwapInputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmSwapInput represents a SwapInput event raised by the Amm contract.
type AmmSwapInput struct {
	Dir              uint8
	QuoteAssetAmount *big.Int
	BaseAssetAmount  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapInput is a free log retrieval operation binding the contract event 0xae6a2b946841d9afc0e1e19a94ae4af26f01125b87b5095bbfb177a9741a2ede.
//
// Solidity: event SwapInput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) FilterSwapInput(opts *bind.FilterOpts) (*AmmSwapInputIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "SwapInput")
	if err != nil {
		return nil, err
	}
	return &AmmSwapInputIterator{contract: _Amm.contract, event: "SwapInput", logs: logs, sub: sub}, nil
}

// WatchSwapInput is a free log subscription operation binding the contract event 0xae6a2b946841d9afc0e1e19a94ae4af26f01125b87b5095bbfb177a9741a2ede.
//
// Solidity: event SwapInput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) WatchSwapInput(opts *bind.WatchOpts, sink chan<- *AmmSwapInput) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "SwapInput")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmSwapInput)
				if err := _Amm.contract.UnpackLog(event, "SwapInput", log); err != nil {
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

// ParseSwapInput is a log parse operation binding the contract event 0xae6a2b946841d9afc0e1e19a94ae4af26f01125b87b5095bbfb177a9741a2ede.
//
// Solidity: event SwapInput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) ParseSwapInput(log types.Log) (*AmmSwapInput, error) {
	event := new(AmmSwapInput)
	if err := _Amm.contract.UnpackLog(event, "SwapInput", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmSwapOutputIterator is returned from FilterSwapOutput and is used to iterate over the raw logs and unpacked data for SwapOutput events raised by the Amm contract.
type AmmSwapOutputIterator struct {
	Event *AmmSwapOutput // Event containing the contract specifics and raw log

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
func (it *AmmSwapOutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmSwapOutput)
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
		it.Event = new(AmmSwapOutput)
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
func (it *AmmSwapOutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmSwapOutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmSwapOutput represents a SwapOutput event raised by the Amm contract.
type AmmSwapOutput struct {
	Dir              uint8
	QuoteAssetAmount *big.Int
	BaseAssetAmount  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapOutput is a free log retrieval operation binding the contract event 0x0dd4066b1a6ce97fb670c3e4201e908c644193f38cbdaffd0229d7e26da3e533.
//
// Solidity: event SwapOutput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) FilterSwapOutput(opts *bind.FilterOpts) (*AmmSwapOutputIterator, error) {

	logs, sub, err := _Amm.contract.FilterLogs(opts, "SwapOutput")
	if err != nil {
		return nil, err
	}
	return &AmmSwapOutputIterator{contract: _Amm.contract, event: "SwapOutput", logs: logs, sub: sub}, nil
}

// WatchSwapOutput is a free log subscription operation binding the contract event 0x0dd4066b1a6ce97fb670c3e4201e908c644193f38cbdaffd0229d7e26da3e533.
//
// Solidity: event SwapOutput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) WatchSwapOutput(opts *bind.WatchOpts, sink chan<- *AmmSwapOutput) (event.Subscription, error) {

	logs, sub, err := _Amm.contract.WatchLogs(opts, "SwapOutput")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmSwapOutput)
				if err := _Amm.contract.UnpackLog(event, "SwapOutput", log); err != nil {
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

// ParseSwapOutput is a log parse operation binding the contract event 0x0dd4066b1a6ce97fb670c3e4201e908c644193f38cbdaffd0229d7e26da3e533.
//
// Solidity: event SwapOutput(uint8 dir, uint256 quoteAssetAmount, uint256 baseAssetAmount)
func (_Amm *AmmFilterer) ParseSwapOutput(log types.Log) (*AmmSwapOutput, error) {
	event := new(AmmSwapOutput)
	if err := _Amm.contract.UnpackLog(event, "SwapOutput", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
