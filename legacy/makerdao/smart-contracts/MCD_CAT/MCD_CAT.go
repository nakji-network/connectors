package MCD_CAT

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
	case "Bite":
		e := new(MCDCATBite)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
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
			Flip: e.Flip.Bytes(),
			Id:   e.Id.Bytes(),
		}
	}
	return nil
}
