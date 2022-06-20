
package MCD_CAT

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
	case "Bite":
		e := new(MCDCATBite)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Bite{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Urn:  e.Urn.Bytes(),
				Ink:  e.Ink.Bytes(),
				Art:  e.Art.Bytes(),
				Tab:  e.Tab.Bytes(),
				Flip:  e.Flip.Bytes(),
				Id:  e.Id.Bytes(),
		}
	}
	return nil
}
