package flowidtablestaking

import (
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&NewEpoch{},
	&EpochTotalRewardsPaid{},
	&NewNodeCreated{},
	&TokensCommitted{},
	&TokensUnstaking{},
	&TokensUnstaked{},
	&NodeRemovedAndRefunded{},
	&RewardsPaid{},
	&UnstakedTokensWithdrawn{},
	&RewardTokensWithdrawn{},
	&NetworkingAddressUpdated{},
	&NewDelegatorCreated{},
	&DelegatorTokensCommitted{},
	&DelegatorTokensStaked{},
	&DelegatorTokensUnstaking{},
	&DelegatorTokensUnstaked{},
	&DelegatorRewardsPaid{},
	&DelegatorUnstakedTokensWithdrawn{},
	&DelegatorRewardTokensWithdrawn{},
	&NewDelegatorCutPercentage{},
	&NewWeeklyPayout{},
	&NewStakingMinimums{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "8624b52f9ddcd04a"
}

func (c SmartContract) Name() string {
	return "FlowIDTableStaking"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewEpoch",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.EpochTotalRewardsPaid",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewNodeCreated",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.TokensCommitted",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.TokensUnstaking",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.TokensUnstaked",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NodeRemovedAndRefunded",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.RewardsPaid",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.UnstakedTokensWithdrawn",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.RewardTokensWithdrawn",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NetworkingAddressUpdated",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewDelegatorCreated",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorTokensCommitted",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorTokensStaked",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorTokensUnstaking",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorTokensUnstaked",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorRewardsPaid",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorUnstakedTokensWithdrawn",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.DelegatorRewardTokensWithdrawn",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewDelegatorCutPercentage",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewWeeklyPayout",
		"A.8624b52f9ddcd04a.FlowIDTableStaking.NewStakingMinimums",
	}
}

func (c SmartContract) Message(vLog flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "NewEpoch":
		totalStaked := vLog.Value.Fields[0].(cadence.UFix64)
		totalRewardPayout := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &NewEpoch{
			Ts:                timestamppb.New(vLog.Timestamp),
			BlockNumber:       vLog.Height,
			TxID:              vLog.TransactionID[:],
			LogIndex:          uint64(vLog.TransactionIndex),
			TotalStaked:       flow.UFix64ToFloat64(totalStaked),
			TotalRewardPayout: flow.UFix64ToFloat64(totalRewardPayout),
		}
		return msg
	case "EpochTotalRewardsPaid":
		total := vLog.Value.Fields[0].(cadence.UFix64)
		fromFees := vLog.Value.Fields[1].(cadence.UFix64)
		minted := vLog.Value.Fields[2].(cadence.UFix64)
		feesBurned := vLog.Value.Fields[3].(cadence.UFix64)
		msg := &EpochTotalRewardsPaid{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Total:       flow.UFix64ToFloat64(total),
			FromFees:    flow.UFix64ToFloat64(fromFees),
			Minted:      flow.UFix64ToFloat64(minted),
			FeesBurned:  flow.UFix64ToFloat64(feesBurned),
		}
		return msg
	case "NewNodeCreated":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		role := vLog.Value.Fields[1].(cadence.UInt32)
		amountCommitted := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &NewNodeCreated{
			Ts:              timestamppb.New(vLog.Timestamp),
			BlockNumber:     vLog.Height,
			TxID:            vLog.TransactionID[:],
			LogIndex:        uint64(vLog.TransactionIndex),
			NodeID:          string(nodeID),
			Role:            uint32(role),
			AmountCommitted: flow.UFix64ToFloat64(amountCommitted),
		}
		return msg
	case "TokensCommitted":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &TokensCommitted{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "TokensUnstaking":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &TokensUnstaking{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "TokensUnstaked":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &TokensUnstaked{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "NodeRemovedAndRefunded":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &NodeRemovedAndRefunded{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "RewardsPaid":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &RewardsPaid{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "UnstakedTokensWithdrawn":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &UnstakedTokensWithdrawn{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "RewardTokensWithdrawn":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		amount := vLog.Value.Fields[1].(cadence.UFix64)
		msg := &RewardTokensWithdrawn{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "NetworkingAddressUpdated":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		newAddress := vLog.Value.Fields[1].(cadence.String)
		msg := &NetworkingAddressUpdated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			NewAddress:  string(newAddress),
		}
		return msg
	case "NewDelegatorCreated":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		msg := &NewDelegatorCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
		}
		return msg
	case "DelegatorTokensCommitted":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorTokensCommitted{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "DelegatorTokensStaked":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorTokensStaked{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "DelegatorTokensUnstaking":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorTokensUnstaking{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "DelegatorTokensUnstaked":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorTokensUnstaked{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "DelegatorRewardsPaid":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorRewardsPaid{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "DelegatorUnstakedTokensWithdrawn":
		nodeID := vLog.Value.Fields[0].(cadence.String)
		delegatorID := vLog.Value.Fields[1].(cadence.UInt32)
		amount := vLog.Value.Fields[2].(cadence.UFix64)
		msg := &DelegatorUnstakedTokensWithdrawn{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NodeID:      string(nodeID),
			DelegatorID: uint32(delegatorID),
			Amount:      flow.UFix64ToFloat64(amount),
		}
		return msg
	case "NewDelegatorCutPercentage":
		newCutPercentage := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &NewDelegatorCutPercentage{
			Ts:               timestamppb.New(vLog.Timestamp),
			BlockNumber:      vLog.Height,
			TxID:             vLog.TransactionID[:],
			LogIndex:         uint64(vLog.TransactionIndex),
			NewCutPercentage: flow.UFix64ToFloat64(newCutPercentage),
		}
		return msg
	case "NewWeeklyPayout":
		newPayout := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &NewWeeklyPayout{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NewPayout:   flow.UFix64ToFloat64(newPayout),
		}
		return msg
	case "NewStakingMinimums":
		newMininums := vLog.Value.Fields[0].(cadence.Dictionary)
		msg := &NewStakingMinimums{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NewMinimums: make(map[uint32]float64),
		}
		for _, p := range newMininums.Pairs {
			k := p.Key.(cadence.UInt32)
			v := p.Value.(cadence.UFix64)
			msg.NewMinimums[uint32(k)] = flow.UFix64ToFloat64(v)
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
