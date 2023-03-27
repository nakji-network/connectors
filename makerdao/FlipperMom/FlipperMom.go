
package FlipperMom

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
		e := new(FLIPPERMOMDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Flip:  e.Flip.Bytes(),
				Cat:  e.Cat.Bytes(),
		}
	case "Rely":
		e := new(FLIPPERMOMRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Flip:  e.Flip.Bytes(),
				Usr:  e.Usr.Bytes(),
		}
	case "SetAuthority":
		e := new(FLIPPERMOMSetAuthority)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetAuthority{
				Ts:   timestamp,
				OldAuthority:  e.OldAuthority.Bytes(),
				NewAuthority:  e.NewAuthority.Bytes(),
		}
	case "SetOwner":
		e := new(FLIPPERMOMSetOwner)
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
