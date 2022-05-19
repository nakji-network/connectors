package FLIPPER_MOM

import (
	blepEthclient "blep.ai/data/chain/ethereum/ethclient"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Deny":
		e := new(FLIPPERMOMDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:   timestamp,
			Flip: e.Flip.Bytes(),
			Cat:  e.Cat.Bytes(),
		}
	case "Rely":
		e := new(FLIPPERMOMRely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:   timestamp,
			Flip: e.Flip.Bytes(),
			Usr:  e.Usr.Bytes(),
		}
	case "SetAuthority":
		e := new(FLIPPERMOMSetAuthority)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetAuthority{
			Ts:           timestamp,
			OldAuthority: e.OldAuthority.Bytes(),
			NewAuthority: e.NewAuthority.Bytes(),
		}
	case "SetOwner":
		e := new(FLIPPERMOMSetOwner)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SetOwner{
			Ts:       timestamp,
			OldOwner: e.OldOwner.Bytes(),
			NewOwner: e.NewOwner.Bytes(),
		}
	}
	return nil
}
