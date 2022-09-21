package flowstakingcollection

import (
	"github.com/nakji-network/connectors/flow"
	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&NodeAddedToStakingCollection{},
	&DelegatorAddedToStakingCollection{},
	&NodeRemovedFromStakingCollection{},
	&DelegatorRemovedFromStakingCollection{},
	&MachineAccountCreated{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "8d0e87b65159ae63"
}

func (c SmartContract) Name() string {
	return "FlowStakingCollection"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.8d0e87b65159ae63.FlowStakingCollection.NodeAddedToStakingCollection",
		"A.8d0e87b65159ae63.FlowStakingCollection.DelegatorAddedToStakingCollection",
		"A.8d0e87b65159ae63.FlowStakingCollection.NodeRemovedFromStakingCollection",
		"A.8d0e87b65159ae63.FlowStakingCollection.DelegatorRemovedFromStakingCollection",
		"A.8d0e87b65159ae63.FlowStakingCollection.MachineAccountCreated",
	}
}

func (c SmartContract) Message(vLog flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "NodeAddedToStakingCollection":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		role := vLog.Value.Fields[1].(cadence.UInt32)
		amountCommitted := vLog.Value.Fields[2].(cadence.UFix64)
		address := vLog.Value.Fields[3].(cadence.Optional)
		msg := &NodeAddedToStakingCollection{
			Ts:              timestamppb.New(vLog.Timestamp),
			BlockNumber:     vLog.Height,
			TxID:            vLog.TransactionID[:],
			LogIndex:        uint64(vLog.TransactionIndex),
			NodeID:          string(nodeID),
			Role:            uint32(role),
			AmountCommitted: float64(amountCommitted),
		}
		if address.Value != nil {
			msg.Address = address.Value.(cadence.Address).Bytes()
		}
		return msg
	case "DelegatorAddedToStakingCollection":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amountCommitted := vLog.Value.Fields[2].(cadence.UFix64)
		address := vLog.Value.Fields[3].(cadence.Optional)
		msg := &DelegatorAddedToStakingCollection{
			Ts:              timestamppb.New(vLog.Timestamp),
			BlockNumber:     vLog.Height,
			TxID:            vLog.TransactionID[:],
			LogIndex:        uint64(vLog.TransactionIndex),
			NodeID:          string(nodeID),
			DelegatorID:     uint32(delegatorID),
			AmountCommitted: float64(amountCommitted),
		}
		if address.Value != nil {
			msg.Address = address.Value.(cadence.Address).Bytes()
		}
		return msg
	case "NodeRemovedFromStakingCollection":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		role := vLog.Value.Fields[1].(cadence.UInt32)
		address := vLog.Value.Fields[2].(cadence.Optional)
		msg := &NodeRemovedFromStakingCollection{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Role:        uint32(role),
		}
		if address.Value != nil {
			msg.Address = address.Value.(cadence.Address).Bytes()
		}
		return msg
	case "DelegatorRemovedFromStakingCollection":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		address := vLog.Value.Fields[2].(cadence.Optional)
		msg := &DelegatorRemovedFromStakingCollection{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
		}
		if address.Value != nil {
			msg.Address = address.Value.(cadence.Address).Bytes()
		}
		return msg
	case "MachineAccountCreated":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		role := vLog.Value.Fields[1].(cadence.UInt32)
		address := vLog.Value.Fields[2].(cadence.Address)
		msg := &MachineAccountCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Role:        uint32(role),
			Address:     address.Bytes(),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
