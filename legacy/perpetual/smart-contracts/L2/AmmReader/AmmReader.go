package AmmReader

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
	case "OwnershipTransferred":
		e := new(AmmReaderOwnershipTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "RewardTransferred":
		e := new(AmmReaderRewardTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RewardTransferred{
			Ts:     timestamp,
			Amount: e.Amount.Bytes(),
		}
	case "RewardWithdrawn":
		e := new(AmmReaderRewardWithdrawn)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RewardWithdrawn{
			Ts:     timestamp,
			Staker: e.Staker.Bytes(),
			Amount: e.Amount.Bytes(),
		}
	}
	return nil
}
