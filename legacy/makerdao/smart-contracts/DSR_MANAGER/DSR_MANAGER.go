package DSR_MANAGER

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
	case "Exit":
		e := new(DSRMANAGERExit)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Exit{
			Ts:  timestamp,
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Join":
		e := new(DSRMANAGERJoin)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Join{
			Ts:  timestamp,
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	}
	return nil
}
