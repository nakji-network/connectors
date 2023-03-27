
package ProxyFactory

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
	if eventName == "Created"{
		e := new(PROXYFACTORYCreated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Created{
				Ts:   timestamp,
				Sender:  e.Sender.Bytes(),
				Owner:  e.Owner.Bytes(),
				Proxy:  e.Proxy.Bytes(),
				Cache:  e.Cache.Bytes(),
		}
	}
	return nil
}
