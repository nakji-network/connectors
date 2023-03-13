package erc721

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, evLog types.Log, timestamp *timestamppb.Timestamp) proto.Message {
	switch eventName {
	case "Approval":
		event := new(IERC721Approval)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &Approval{
			Ts:          timestamp,
			Owner:       event.Owner.Bytes(),
			Approved:    event.Approved.Bytes(),
			TokenId:     event.TokenId.Bytes(),
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	case "ApprovalForAll":
		event := new(IERC721ApprovalForAll)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &ApprovalForAll{
			Ts:          timestamp,
			Owner:       event.Owner.Bytes(),
			Operator:    event.Operator.Bytes(),
			Approved:    event.Approved,
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	case "Transfer":
		event := new(IERC721Transfer)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &Transfer{
			Ts:          timestamp,
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			TokenId:     event.TokenId.Bytes(),
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	}
	return nil
}
