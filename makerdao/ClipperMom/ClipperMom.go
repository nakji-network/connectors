
package ClipperMom

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
	case "SetAuthority":
		e := new(CLIPPERMOMSetAuthority)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetAuthority{
				Ts:   timestamp,
				OldAuthority:  e.OldAuthority.Bytes(),
				NewAuthority:  e.NewAuthority.Bytes(),
		}
	case "SetBreaker":
		e := new(CLIPPERMOMSetBreaker)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetBreaker{
				Ts:   timestamp,
				Clip:  e.Clip.Bytes(),
				Level:  e.Level.Bytes(),
		}
	case "SetOwner":
		e := new(CLIPPERMOMSetOwner)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetOwner{
				Ts:   timestamp,
				OldOwner:  e.OldOwner.Bytes(),
				NewOwner:  e.NewOwner.Bytes(),
		}
	}
	return nil
}
