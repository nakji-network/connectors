package WooCrossChainRouterV1

import (
	"strings"

	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	addr string
	abi  abi.ABI
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(WooCrossChainRouterV1ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooCrossChainRouterV1ABI")
	}
	return &SmartContract{addr: addr, abi: contractAbi}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&WooCrossSwapOnSrcChain{},
		&WooCrossSwapOnDstChain{},
		&OwnershipTransferred{},
	}
}

func (sc *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "WooCrossSwapOnSrcChain":
		e := new(WooCrossChainRouterV1WooCrossSwapOnSrcChain)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
			ContractAddress: sc.Address(),
		}
	case "WooCrossSwapOnDstChain":
		e := new(WooCrossChainRouterV1WooCrossSwapOnDstChain)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooCrossSwapOnDstChain{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			RefId:           e.RefId.Bytes(),
			Sender:          e.Sender.Bytes(),
			To:              e.To.Bytes(),
			BridgedToken:    e.BridgedToken.Bytes(),
			BridgedAmount:   e.BridgedAmount.Bytes(),
			ToToken:         e.ToToken.Bytes(),
			RealToToken:     e.RealToToken.Bytes(),
			MinToAmount:     e.MinToAmount.Bytes(),
			RealToAmount:    e.RealToAmount.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "OwnershipTransferred":
		e := new(WooCrossChainRouterV1OwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferred{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.Address(),
		}
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
