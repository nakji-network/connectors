package Amm

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
	case "CapChanged":
		e := new(AmmCapChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &CapChanged{
			Ts:                      timestamp,
			MaxHoldingBaseAsset:     e.MaxHoldingBaseAsset.Bytes(),
			OpenInterestNotionalCap: e.OpenInterestNotionalCap.Bytes(),
		}
	case "FundingRateUpdated":
		e := new(AmmFundingRateUpdated)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &FundingRateUpdated{
			Ts:              timestamp,
			Rate:            e.Rate.Bytes(),
			UnderlyingPrice: e.UnderlyingPrice.Bytes(),
		}
	case "LiquidityChanged":
		e := new(AmmLiquidityChanged)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LiquidityChanged{
			Ts:                 timestamp,
			QuoteReserve:       e.QuoteReserve.Bytes(),
			BaseReserve:        e.BaseReserve.Bytes(),
			CumulativeNotional: e.CumulativeNotional.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(AmmOwnershipTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "PriceFeedUpdated":
		e := new(AmmPriceFeedUpdated)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PriceFeedUpdated{
			Ts:        timestamp,
			PriceFeed: e.PriceFeed.Bytes(),
		}
	case "ReserveSnapshotted":
		e := new(AmmReserveSnapshotted)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ReserveSnapshotted{
			Ts:                timestamp,
			QuoteAssetReserve: e.QuoteAssetReserve.Bytes(),
			BaseAssetReserve:  e.BaseAssetReserve.Bytes(),
			Timestamp:         e.Timestamp.Bytes(),
		}
	case "Shutdown":
		e := new(AmmShutdown)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Shutdown{
			Ts:              timestamp,
			SettlementPrice: e.SettlementPrice.Bytes(),
		}
	case "SwapInput":
		e := new(AmmSwapInput)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SwapInput{
			Ts:               timestamp,
			QuoteAssetAmount: e.QuoteAssetAmount.Bytes(),
			BaseAssetAmount:  e.BaseAssetAmount.Bytes(),
		}
	case "SwapOutput":
		e := new(AmmSwapOutput)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SwapOutput{
			Ts:               timestamp,
			QuoteAssetAmount: e.QuoteAssetAmount.Bytes(),
			BaseAssetAmount:  e.BaseAssetAmount.Bytes(),
		}
	}
	return nil
}
