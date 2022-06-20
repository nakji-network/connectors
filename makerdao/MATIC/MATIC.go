
package MATIC

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
		e := new(MATICApproval)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
				Ts:   timestamp,
				Owner:  e.Owner.Bytes(),
				Spender:  e.Spender.Bytes(),
				Value:  e.Value.Bytes(),
		}
	case "Paused":
		e := new(MATICPaused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Paused{
				Ts:   timestamp,
				Account:  e.Account.Bytes(),
		}
	case "PauserAdded":
		e := new(MATICPauserAdded)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PauserAdded{
				Ts:   timestamp,
				Account:  e.Account.Bytes(),
		}
	case "PauserRemoved":
		e := new(MATICPauserRemoved)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PauserRemoved{
				Ts:   timestamp,
				Account:  e.Account.Bytes(),
		}
	case "Transfer":
		e := new(MATICTransfer)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
				Ts:   timestamp,
				From:  e.From.Bytes(),
				To:  e.To.Bytes(),
				Value:  e.Value.Bytes(),
		}
	case "Unpaused":
		e := new(MATICUnpaused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unpaused{
				Ts:   timestamp,
				Account:  e.Account.Bytes(),
		}
	}
	return nil
}
