package erc1155

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
	case "ApprovalForAll":
		event := new(IERC1155ApprovalForAll)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &ApprovalForAll{
			Ts:          timestamp,
			Account:     event.Account.Bytes(),
			Operator:    event.Operator.Bytes(),
			Approved:    event.Approved,
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	case "TransferBatch":
		event := new(IERC1155TransferBatch)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}

		var ids [][]byte
		for _, id := range event.Ids {
			ids = append(ids, id.Bytes())
		}

		var values [][]byte
		for _, value := range event.Values {
			values = append(values, value.Bytes())
		}

		return &TransferBatch{
			Ts:          timestamp,
			Operator:    event.Operator.Bytes(),
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			Ids:         ids,
			Values:      values,
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	case "TransferSingle":
		event := new(IERC1155TransferSingle)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &TransferSingle{
			Ts:          timestamp,
			Operator:    event.Operator.Bytes(),
			From:        event.From.Bytes(),
			To:          event.To.Bytes(),
			Id:          event.Id.Bytes(),
			Value:       event.Value.Bytes(),
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	case "URI":
		event := new(IERC1155URI)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("failed to unpack event ")
			return nil
		}
		return &URI{
			Ts:          timestamp,
			Value:       event.Value,
			Id:          event.Id.Bytes(),
			BlockNumber: evLog.BlockNumber,
			LogIndex:    uint64(evLog.Index),
			TxHash:      evLog.TxHash.Bytes(),
		}
	}
	return nil
}
