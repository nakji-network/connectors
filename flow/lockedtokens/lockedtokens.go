package lockedtokens

import (
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&SharedAccountRegistered{},
	&UnlockedAccountRegistered{},
	&UnlockLimitIncreased{},
	&LockedAccountRegisteredAsNode{},
	&LockedAccountRegisteredAsDelegator{},
	&LockedTokensDeposited{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "8d0e87b65159ae63"
}

func (c SmartContract) Name() string {
	return "LockedTokens"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.8d0e87b65159ae63.LockedTokens.SharedAccountRegistered",
		"A.8d0e87b65159ae63.LockedTokens.UnlockedAccountRegistered",
		"A.8d0e87b65159ae63.LockedTokens.UnlockLimitIncreased",
		"A.8d0e87b65159ae63.LockedTokens.LockedAccountRegisteredAsNode",
		"A.8d0e87b65159ae63.LockedTokens.LockedAccountRegisteredAsDelegator",
		"A.8d0e87b65159ae63.LockedTokens.LockedTokensDeposited",
	}
}

func (c SmartContract) Message(vLog *flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "SharedAccountRegistered":
		address := vLog.Value.Fields[0].(cadence.Address)
		msg := &SharedAccountRegistered{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
		}
		return msg
	case "UnlockedAccountRegistered":
		address := vLog.Value.Fields[0].(cadence.Address)
		msg := &UnlockedAccountRegistered{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
		}
		return msg
	case "UnlockLimitIncreased":
		address := vLog.Value.Fields[0].(cadence.Address)
		increaseAmount := vLog.Value.Fields[1].(cadence.UFix64)
		newLimit := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &UnlockLimitIncreased{
			Ts:             timestamppb.New(vLog.Timestamp),
			BlockNumber:    vLog.Height,
			TxID:           vLog.TransactionID[:],
			LogIndex:       uint64(vLog.TransactionIndex),
			Address:        address.Bytes(),
			IncreaseAmount: flow.UFix64ToFloat64(increaseAmount),
			NewLimit:       flow.UFix64ToFloat64(newLimit),
		}
		return msg
	case "LockedAccountRegisteredAsNode":
		address := vLog.Value.Fields[0].(cadence.Address)
		nodeID := vLog.Value.Fields[1].(cadence.String)
		msg := &LockedAccountRegisteredAsNode{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
			NodeID:      string(nodeID),
		}
		return msg
	case "LockedAccountRegisteredAsDelegator":
		address := vLog.Value.Fields[0].(cadence.Address)
		nodeID := vLog.Value.Fields[1].(cadence.String)
		msg := &LockedAccountRegisteredAsDelegator{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
			NodeID:      string(nodeID),
		}
		return msg
	case "LockedTokensDeposited":
		address := vLog.Value.Fields[0].(cadence.Address)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &LockedTokensDeposited{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
