package L2PriceFeed

import (
	blepethclient "blep.ai/data/chain/ethereum/ethclient"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "OwnershipTransferred":
		e := new(L2PriceFeedOwnershipTransferred)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "PriceFeedDataSet":
		e := new(L2PriceFeedPriceFeedDataSet)
		if err := blepethclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PriceFeedDataSet{
			Ts:        timestamp,
			Key:       e.Key[:],
			Price:     e.Price.Bytes(),
			Timestamp: e.Timestamp.Bytes(),
			RoundId:   e.RoundId.Bytes(),
		}
	}
	return nil
}
