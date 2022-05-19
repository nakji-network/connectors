package DIRECT_MOM

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
	case "Disable":
		e := new(DIRECTMOMDisable)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Disable{
			Ts:  timestamp,
			Who: e.Who.Bytes(),
		}
	case "SetAuthority":
		e := new(DIRECTMOMSetAuthority)
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
		e := new(DIRECTMOMSetOwner)
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
