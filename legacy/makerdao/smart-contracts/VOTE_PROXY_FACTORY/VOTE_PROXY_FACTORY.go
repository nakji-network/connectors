package VOTE_PROXY_FACTORY

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
	case "LinkConfirmed":
		e := new(VOTEPROXYFACTORYLinkConfirmed)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LinkConfirmed{
			Ts:        timestamp,
			Cold:      e.Cold.Bytes(),
			Hot:       e.Hot.Bytes(),
			VoteProxy: e.VoteProxy.Bytes(),
		}
	case "LinkRequested":
		e := new(VOTEPROXYFACTORYLinkRequested)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LinkRequested{
			Ts:   timestamp,
			Cold: e.Cold.Bytes(),
			Hot:  e.Hot.Bytes(),
		}
	}
	return nil
}
