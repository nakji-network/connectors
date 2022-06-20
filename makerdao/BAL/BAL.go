
package BAL

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
		e := new(BALApproval)
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
	case "RoleGranted":
		e := new(BALRoleGranted)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RoleGranted{
				Ts:   timestamp,
				Role:  e.Role[:],
				Account:  e.Account.Bytes(),
				Sender:  e.Sender.Bytes(),
		}
	case "RoleRevoked":
		e := new(BALRoleRevoked)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RoleRevoked{
				Ts:   timestamp,
				Role:  e.Role[:],
				Account:  e.Account.Bytes(),
				Sender:  e.Sender.Bytes(),
		}
	case "Snapshot":
		e := new(BALSnapshot)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Snapshot{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
		}
	case "Transfer":
		e := new(BALTransfer)
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
	}
	return nil
}
