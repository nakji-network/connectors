package WooCrossChainRouterV1

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, ts *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "WooCrossSwapOnSrcChain":
		e := new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooCrossSwapOnSrcChain{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			RefId:           e.RefId.Bytes(),
			Sender:          e.Sender.Bytes(),
			To:              e.To.Bytes(),
			FromToken:       e.FromToken.Bytes(),
			FromAmount:      e.FromAmount.Bytes(),
			MinQuoteAmount:  e.MinQuoteAmount.Bytes(),
			RealQuoteAmount: e.RealQuoteAmount.Bytes(),
		}
	case "WooCrossSwapOnDstChain":
		e := new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooCrossSwapOnDstChain{
			Ts:            ts,
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			RefId:         e.RefId.Bytes(),
			Sender:        e.Sender.Bytes(),
			To:            e.To.Bytes(),
			BridgedToken:  e.BridgedToken.Bytes(),
			BridgedAmount: e.BridgedAmount.Bytes(),
			ToToken:       e.ToToken.Bytes(),
			RealToToken:   e.RealToToken.Bytes(),
			MinToAmount:   e.MinToAmount.Bytes(),
			RealToAmount:  e.RealToAmount.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(WooCrossChainRouterV1OwnershipTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferred{
			Ts:            ts,
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	default:
		log.Error().Msgf("invalid event: %s", eventName)
		return nil
	}
}
