package flowfees

import (
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&TokensDeposited{},
	&TokensWithdrawn{},
	&FeesDeducted{},
	&FeeParametersChanged{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "f919ee77447b7497"
}

func (c SmartContract) Name() string {
	return "FlowFees"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.f919ee77447b7497.FlowFees.TokensDeposited",
		"A.f919ee77447b7497.FlowFees.TokensWithdrawn",
		"A.f919ee77447b7497.FlowFees.FeesDeducted",
		"A.f919ee77447b7497.FlowFees.FeeParametersChanged",
	}
}

func (c SmartContract) Message(vLog flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "TokensDeposited":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TokensDeposited{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "TokensWithdrawn":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TokensWithdrawn{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "FeesDeducted":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		inclusionEffort := vLog.Value.Fields[1].(cadence.UFix64)
		exclusionEffort := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &FeesDeducted{
			Ts:              timestamppb.New(vLog.Timestamp),
			BlockNumber:     vLog.Height,
			TxID:            vLog.TransactionID[:],
			LogIndex:        uint64(vLog.TransactionIndex),
			Amount:          flow.UFix64ToFloat64(amount),
			InclusionEffort: flow.UFix64ToFloat64(inclusionEffort),
			ExecutionEffort: flow.UFix64ToFloat64(exclusionEffort),
		}
		return msg
	case "FeeParametersChanged":
		surgeFactor := vLog.Value.Fields[0].(cadence.UFix64)
		inclusionEffortCost := vLog.Value.Fields[1].(cadence.UFix64)
		exclusionEffortCost := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &FeeParametersChanged{
			Ts:                  timestamppb.New(vLog.Timestamp),
			BlockNumber:         vLog.Height,
			TxID:                vLog.TransactionID[:],
			LogIndex:            uint64(vLog.TransactionIndex),
			SurgeFactor:         flow.UFix64ToFloat64(surgeFactor),
			InclusionEffortCost: flow.UFix64ToFloat64(inclusionEffortCost),
			ExecutionEffortCost: flow.UFix64ToFloat64(exclusionEffortCost),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
