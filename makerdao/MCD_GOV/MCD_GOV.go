
package MCD_GOV

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
		e := new(MCDGOVApproval)
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
	case "Burn":
		e := new(MCDGOVBurn)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Burn{
				Ts:   timestamp,
				Guy:  e.Guy.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "LogSetAuthority":
		e := new(MCDGOVLogSetAuthority)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetAuthority{
				Ts:   timestamp,
				Authority:  e.Authority.Bytes(),
		}
	case "LogSetOwner":
		e := new(MCDGOVLogSetOwner)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetOwner{
				Ts:   timestamp,
				Owner:  e.Owner.Bytes(),
		}
	case "Mint":
		e := new(MCDGOVMint)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Mint{
				Ts:   timestamp,
				Guy:  e.Guy.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Transfer":
		e := new(MCDGOVTransfer)
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
