package UniV2

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Approval":
		e := new(UNIV2Approval)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:      timestamp,
			Owner:   e.Owner.Bytes(),
			Spender: e.Spender.Bytes(),
			Value:   e.Value.Bytes(),
		}
	case "Burn":
		e := new(UNIV2Burn)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Burn{
			Ts:      timestamp,
			Sender:  e.Sender.Bytes(),
			Amount0: e.Amount0.Bytes(),
			Amount1: e.Amount1.Bytes(),
			To:      e.To.Bytes(),
		}
	case "Mint":
		e := new(UNIV2Mint)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Mint{
			Ts:      timestamp,
			Sender:  e.Sender.Bytes(),
			Amount0: e.Amount0.Bytes(),
			Amount1: e.Amount1.Bytes(),
		}
	case "Swap":
		e := new(UNIV2Swap)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Swap{
			Ts:         timestamp,
			Sender:     e.Sender.Bytes(),
			Amount0In:  e.Amount0In.Bytes(),
			Amount1In:  e.Amount1In.Bytes(),
			Amount0Out: e.Amount0Out.Bytes(),
			Amount1Out: e.Amount1Out.Bytes(),
			To:         e.To.Bytes(),
		}
	case "Sync":
		e := new(UNIV2Sync)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Sync{
			Ts:       timestamp,
			Reserve0: e.Reserve0.Bytes(),
			Reserve1: e.Reserve1.Bytes(),
		}
	case "Transfer":
		e := new(UNIV2Transfer)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:    timestamp,
			From:  e.From.Bytes(),
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	}
	return nil
}
