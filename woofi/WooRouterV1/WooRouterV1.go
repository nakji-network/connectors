package WooRouterV1

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
	case "OwnershipTransferred":
		e := new(WooRouterV1OwnershipTransferred)
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
	case "WooPoolChanged":
		e := new(WooRouterV1WooPoolChanged)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooPoolChanged{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			NewPool:     e.NewPool.Bytes(),
		}
	case "WooRouterSwap":
		e := new(WooRouterV1WooRouterSwap)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooRouterSwap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			SwapType:    uint32(e.SwapType),
			FromToken:   e.FromToken.Bytes(),
			ToToken:     e.ToToken.Bytes(),
			FromAmount:  e.FromToken.Bytes(),
			ToAmount:    e.ToAmount.Bytes(),
			From:        e.FromToken.Bytes(),
			To:          e.To.Bytes(),
		}
	default:
		log.Error().Msgf("invalid event: %s", eventName)
		return nil
	}
}
