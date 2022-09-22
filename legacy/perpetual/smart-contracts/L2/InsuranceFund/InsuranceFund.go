package InsuranceFund

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
		e := new(InsuranceFundOwnershipTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "ShutdownAllAmms":
		e := new(InsuranceFundShutdownAllAmms)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ShutdownAllAmms{
			Ts:          timestamp,
			BlockNumber: e.BlockNumber.Bytes(),
		}
	case "TokenAdded":
		e := new(InsuranceFundTokenAdded)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &TokenAdded{
			Ts:           timestamp,
			TokenAddress: e.TokenAddress.Bytes(),
		}
	case "TokenRemoved":
		e := new(InsuranceFundTokenRemoved)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &TokenRemoved{
			Ts:           timestamp,
			TokenAddress: e.TokenAddress.Bytes(),
		}
	case "Withdrawn":
		e := new(InsuranceFundWithdrawn)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Withdrawn{
			Ts:         timestamp,
			Withdrawer: e.Withdrawer.Bytes(),
			Amount:     e.Amount.Bytes(),
		}
	}
	return nil
}
