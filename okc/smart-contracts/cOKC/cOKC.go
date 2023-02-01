package cOKC

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
		event := new(COKCApproval)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Approval event error")
			return nil
		}
		return &Approval{
			Ts:      timestamp,
			Block:   evLog.BlockNumber,
			Idx:     uint64(evLog.Index),
			Tx:      evLog.TxHash.Bytes(),
			Owner:   event.Owner.Bytes(),
			Spender: event.Spender.Bytes(),
			Value:   event.Value.Bytes(),
		}
	case "MintClaimed":
		event := new(COKCMintClaimed)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack MintClaimed event error")
			return nil
		}
		return &MintClaimed{
			Ts:           timestamp,
			Block:        evLog.BlockNumber,
			Idx:          uint64(evLog.Index),
			Tx:           evLog.TxHash.Bytes(),
			User:         event.User.Bytes(),
			RewardAmount: event.RewardAmount.Bytes(),
		}
	case "RankClaimed":
		event := new(COKCRankClaimed)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack RankClaimed event error")
			return nil
		}
		return &RankClaimed{
			Ts:    timestamp,
			Block: evLog.BlockNumber,
			Idx:   uint64(evLog.Index),
			Tx:    evLog.TxHash.Bytes(),
			User:  event.User.Bytes(),
			Term:  event.Term.Bytes(),
			Rank:  event.Rank.Bytes(),
		}
	case "Staked":
		event := new(COKCStaked)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Staked event error")
			return nil
		}
		return &Staked{
			Ts:     timestamp,
			Block:  evLog.BlockNumber,
			Idx:    uint64(evLog.Index),
			Tx:     evLog.TxHash.Bytes(),
			User:   event.User.Bytes(),
			Amount: event.Amount.Bytes(),
			Term:   event.Term.Bytes(),
		}
	case "Transfer":
		event := new(COKCTransfer)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Transfer event error")
			return nil
		}
		return &Transfer{
			Ts:    timestamp,
			Block: evLog.BlockNumber,
			Idx:   uint64(evLog.Index),
			Tx:    evLog.TxHash.Bytes(),
			From:  event.From.Bytes(),
			To:    event.To.Bytes(),
			Value: event.Value.Bytes(),
		}
	case "Withdrawn":
		event := new(COKCWithdrawn)
		if err := common.UnpackLog(*contractAbi, event, eventName, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Withdrawn event error")
			return nil
		}
		return &Withdrawn{
			Ts:     timestamp,
			Block:  evLog.BlockNumber,
			Idx:    uint64(evLog.Index),
			Tx:     evLog.TxHash.Bytes(),
			User:   event.User.Bytes(),
			Amount: event.Amount.Bytes(),
			Reward: event.Reward.Bytes(),
		}
	}
	return nil
}
