package WooRouterV1

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
	contractAbi, err := abi.JSON(strings.NewReader(WooRouterV1ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooRouterV1ABI")
	}
	return &SmartContract{addr: addr, abi: contractAbi}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&OwnershipTransferred{},
		&WooPoolChanged{},
		&WooRouterSwap{},
	}
}

func (sc *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "OwnershipTransferred":
		e := new(WooRouterV1OwnershipTransferred)
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
	case "WooPoolChanged":
		e := new(WooRouterV1WooPoolChanged)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooPoolChanged{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewPool:         e.NewPool.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "WooRouterSwap":
		e := new(WooRouterV1WooRouterSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooRouterSwap{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			SwapType:        uint32(e.SwapType),
			FromToken:       e.FromToken.Bytes(),
			ToToken:         e.ToToken.Bytes(),
			FromAmount:      e.FromToken.Bytes(),
			ToAmount:        e.ToAmount.Bytes(),
			From:            e.FromToken.Bytes(),
			To:              e.To.Bytes(),
			ContractAddress: sc.Address(),
		}
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
