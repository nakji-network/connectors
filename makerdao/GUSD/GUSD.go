
package GUSD

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
		e := new(GUSDApproval)
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
	case "CustodianChangeConfirmed":
		e := new(GUSDCustodianChangeConfirmed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &CustodianChangeConfirmed{
				Ts:   timestamp,
				LockId:  e.LockId[:],
				NewCustodian:  e.NewCustodian.Bytes(),
		}
	case "CustodianChangeRequested":
		e := new(GUSDCustodianChangeRequested)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &CustodianChangeRequested{
				Ts:   timestamp,
				LockId:  e.LockId[:],
				MsgSender:  e.MsgSender.Bytes(),
				ProposedCustodian:  e.ProposedCustodian.Bytes(),
		}
	case "ImplChangeConfirmed":
		e := new(GUSDImplChangeConfirmed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ImplChangeConfirmed{
				Ts:   timestamp,
				LockId:  e.LockId[:],
				NewImpl:  e.NewImpl.Bytes(),
		}
	case "ImplChangeRequested":
		e := new(GUSDImplChangeRequested)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ImplChangeRequested{
				Ts:   timestamp,
				LockId:  e.LockId[:],
				MsgSender:  e.MsgSender.Bytes(),
				ProposedImpl:  e.ProposedImpl.Bytes(),
		}
	case "Transfer":
		e := new(GUSDTransfer)
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
