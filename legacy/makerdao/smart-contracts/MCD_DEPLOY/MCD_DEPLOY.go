package MCD_DEPLOY

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
	case "LogSetAuthority":
		e := new(MCDDEPLOYLogSetAuthority)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetAuthority{
			Ts:        timestamp,
			Authority: e.Authority.Bytes(),
		}
	case "LogSetOwner":
		e := new(MCDDEPLOYLogSetOwner)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetOwner{
			Ts:    timestamp,
			Owner: e.Owner.Bytes(),
		}
	}
	return nil
}
