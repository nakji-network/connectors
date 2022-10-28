package bscWooRouterV1

import (
	"fmt"
	"strings"

	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	network string
	address string
	abi     abi.ABI
}

func NewContract(network string, address string) (*SmartContract, error) {
	contractAbi, err := abi.JSON(strings.NewReader(BscWooRouterV1ABI))
	if err != nil {
		return nil, fmt.Errorf("error reading BscWooRouterV1ABI: %s", err)
	}
	return &SmartContract{network: network, address: address, abi: contractAbi}, nil
}

func (sc *SmartContract) Network() string {
	return sc.network
}

func (sc *SmartContract) Address() string {
	return sc.address
}

func (sc *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "OwnershipTransferred":
		e := new(BscWooRouterV1OwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
		e := new(BscWooRouterV1WooPoolChanged)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
		e := new(BscWooRouterV1WooRouterSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}