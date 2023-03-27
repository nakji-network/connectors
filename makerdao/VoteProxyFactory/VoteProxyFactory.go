
package VoteProxyFactory

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
	case "LinkConfirmed":
		e := new(VOTEPROXYFACTORYLinkConfirmed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LinkConfirmed{
				Ts:   timestamp,
				Cold:  e.Cold.Bytes(),
				Hot:  e.Hot.Bytes(),
				VoteProxy:  e.VoteProxy.Bytes(),
		}
	case "LinkRequested":
		e := new(VOTEPROXYFACTORYLinkRequested)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LinkRequested{
				Ts:   timestamp,
				Cold:  e.Cold.Bytes(),
				Hot:  e.Hot.Bytes(),
		}
	}
	return nil
}
