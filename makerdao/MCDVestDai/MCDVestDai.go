
package MCDVestDai

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
	case "Deny":
		e := new(MCDVESTDAIDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "File":
		e := new(MCDVESTDAIFile)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "Init":
		e := new(MCDVESTDAIInit)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Init{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
				Usr:  e.Usr.Bytes(),
		}
	case "Move":
		e := new(MCDVESTDAIMove)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Move{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
				Dst:  e.Dst.Bytes(),
		}
	case "Rely":
		e := new(MCDVESTDAIRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Restrict":
		e := new(MCDVESTDAIRestrict)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Restrict{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
		}
	case "Unrestrict":
		e := new(MCDVESTDAIUnrestrict)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unrestrict{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
		}
	case "Vest":
		e := new(MCDVESTDAIVest)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Vest{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
				Amt:  e.Amt.Bytes(),
		}
	case "Yank":
		e := new(MCDVESTDAIYank)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Yank{
				Ts:   timestamp,
				Id:  e.Id.Bytes(),
				End:  e.End.Bytes(),
		}
	}
	return nil
}
