package flowstoragefees

import (
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&StorageMegaBytesPerReservedFLOWChanged{},
	&MinimumStorageReservationChanged{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "e467b9dd11fa00df"
}

func (c SmartContract) Name() string {
	return "FlowStorageFees"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.e467b9dd11fa00df.FlowStorageFees.StorageMegaBytesPerReservedFLOWChanged",
		"A.e467b9dd11fa00df.FlowStorageFees.MinimumStorageReservationChanged",
	}
}

func (c SmartContract) Message(vLog *flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "StorageMegaBytesPerReservedFLOWChanged":
		storageMegaBytesPerReservedFLOW := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &StorageMegaBytesPerReservedFLOWChanged{
			Ts:                              timestamppb.New(vLog.Timestamp),
			BlockNumber:                     vLog.Height,
			TxID:                            vLog.TransactionID[:],
			LogIndex:                        uint64(vLog.TransactionIndex),
			StorageMegaBytesPerReservedFLOW: flow.UFix64ToFloat64(storageMegaBytesPerReservedFLOW),
		}
		return msg
	case "MinimumStorageReservationChanged":
		minimumStorageReservation := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &MinimumStorageReservationChanged{
			Ts:                        timestamppb.New(vLog.Timestamp),
			BlockNumber:               vLog.Height,
			TxID:                      vLog.TransactionID[:],
			LogIndex:                  uint64(vLog.TransactionIndex),
			MinimumStorageReservation: flow.UFix64ToFloat64(minimumStorageReservation),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
