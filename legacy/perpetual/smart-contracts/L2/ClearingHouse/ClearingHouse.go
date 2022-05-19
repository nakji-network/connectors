package ClearingHouse

import (
	blepethclient "blep.ai/data/chain/ethereum/ethclient"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "LiquidationFeeRatioChanged":
		e := new(ClearingHouseLiquidationFeeRatioChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LiquidationFeeRatioChanged{
			Ts:                  timestamp,
			LiquidationFeeRatio: e.LiquidationFeeRatio.Bytes(),
		}
	case "MarginChanged":
		e := new(ClearingHouseMarginChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &MarginChanged{
			Ts:             timestamp,
			Sender:         e.Sender.Bytes(),
			Amm:            e.Amm.Bytes(),
			Amount:         e.Amount.Bytes(),
			FundingPayment: e.FundingPayment.Bytes(),
		}
	case "MarginRatioChanged":
		e := new(ClearingHouseMarginRatioChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &MarginRatioChanged{
			Ts:          timestamp,
			MarginRatio: e.MarginRatio.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(ClearingHouseOwnershipTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "Paused":
		e := new(ClearingHousePaused)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Paused{
			Ts:      timestamp,
			Account: e.Account.Bytes(),
		}
	case "PositionAdjusted":
		e := new(ClearingHousePositionAdjusted)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PositionAdjusted{
			Ts:                timestamp,
			Amm:               e.Amm.Bytes(),
			Trader:            e.Trader.Bytes(),
			NewPositionSize:   e.NewPositionSize.Bytes(),
			OldLiquidityIndex: e.OldLiquidityIndex.Bytes(),
			NewLiquidityIndex: e.NewLiquidityIndex.Bytes(),
		}
	case "PositionChanged":
		e := new(ClearingHousePositionChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PositionChanged{
			Ts:                    timestamp,
			Trader:                e.Trader.Bytes(),
			Amm:                   e.Amm.Bytes(),
			Margin:                e.Margin.Bytes(),
			PositionNotional:      e.PositionNotional.Bytes(),
			ExchangedPositionSize: e.ExchangedPositionSize.Bytes(),
			Fee:                   e.Fee.Bytes(),
			PositionSizeAfter:     e.PositionSizeAfter.Bytes(),
			RealizedPnl:           e.RealizedPnl.Bytes(),
			UnrealizedPnlAfter:    e.UnrealizedPnlAfter.Bytes(),
			BadDebt:               e.BadDebt.Bytes(),
			LiquidationPenalty:    e.LiquidationPenalty.Bytes(),
			SpotPrice:             e.SpotPrice.Bytes(),
			FundingPayment:        e.FundingPayment.Bytes(),
		}
	case "PositionLiquidated":
		e := new(ClearingHousePositionLiquidated)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PositionLiquidated{
			Ts:               timestamp,
			Trader:           e.Trader.Bytes(),
			Amm:              e.Amm.Bytes(),
			PositionNotional: e.PositionNotional.Bytes(),
			PositionSize:     e.PositionSize.Bytes(),
			LiquidationFee:   e.LiquidationFee.Bytes(),
			Liquidator:       e.Liquidator.Bytes(),
			BadDebt:          e.BadDebt.Bytes(),
		}
	case "PositionSettled":
		e := new(ClearingHousePositionSettled)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PositionSettled{
			Ts:               timestamp,
			Amm:              e.Amm.Bytes(),
			Trader:           e.Trader.Bytes(),
			ValueTransferred: e.ValueTransferred.Bytes(),
		}
	case "ReferredPositionChanged":
		e := new(ClearingHouseReferredPositionChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ReferredPositionChanged{
			Ts:           timestamp,
			ReferralCode: e.ReferralCode[:],
		}
	case "RestrictionModeEntered":
		e := new(ClearingHouseRestrictionModeEntered)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RestrictionModeEntered{
			Ts:          timestamp,
			Amm:         e.Amm.Bytes(),
			BlockNumber: e.BlockNumber.Bytes(),
		}
	case "Unpaused":
		e := new(ClearingHouseUnpaused)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unpaused{
			Ts:      timestamp,
			Account: e.Account.Bytes(),
		}
	}
	return nil
}
