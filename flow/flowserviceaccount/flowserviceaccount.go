package flowserviceaccount

import (
	"github.com/nakji-network/connectors/flow"
	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&TransactionFeeUpdated{},
	&AccountCreationFeeUpdated{},
	&AccountCreatorAdded{},
	&AccountCreatorRemoved{},
	&IsAccountCreationRestrictedUpdated{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "e467b9dd11fa00df"
}

func (c SmartContract) Name() string {
	return "FlowServiceAccount"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.e467b9dd11fa00df.FlowServiceAccount.TransactionFeeUpdated",
		"A.e467b9dd11fa00df.FlowServiceAccount.AccountCreationFeeUpdated",
		"A.e467b9dd11fa00df.FlowServiceAccount.AccountCreatorAdded",
		"A.e467b9dd11fa00df.FlowServiceAccount.AccountCreatorRemoved",
		"A.e467b9dd11fa00df.FlowServiceAccount.IsAccountCreationRestrictedUpdated",
	}
}

func (c SmartContract) Message(vLog flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "TransactionFeeUpdated":
		newFee := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TransactionFeeUpdated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NewFee:      flow.UFix64ToFloat64(newFee),
		}
		return msg
	case "AccountCreationFeeUpdated":
		newFee := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &AccountCreationFeeUpdated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			NewFee:      flow.UFix64ToFloat64(newFee),
		}
		return msg
	case "AccountCreatorAdded":
		accountCreater := vLog.Value.Fields[0].(cadence.Address)
		msg := &AccountCreatorAdded{
			Ts:             timestamppb.New(vLog.Timestamp),
			BlockNumber:    vLog.Height,
			TxID:           vLog.TransactionID[:],
			LogIndex:       uint64(vLog.TransactionIndex),
			AccountCreator: accountCreater[:],
		}
		return msg
	case "AccountCreatorRemoved":
		accountCreater := vLog.Value.Fields[0].(cadence.Address)
		msg := &AccountCreatorRemoved{
			Ts:             timestamppb.New(vLog.Timestamp),
			BlockNumber:    vLog.Height,
			TxID:           vLog.TransactionID[:],
			LogIndex:       uint64(vLog.TransactionIndex),
			AccountCreator: accountCreater[:],
		}
		return msg
	case "IsAccountCreationRestrictedUpdated":
		isRestricted := vLog.Value.Fields[0].(cadence.Bool)
		msg := &IsAccountCreationRestrictedUpdated{
			Ts:           timestamppb.New(vLog.Timestamp),
			BlockNumber:  vLog.Height,
			TxID:         vLog.TransactionID[:],
			LogIndex:     uint64(vLog.TransactionIndex),
			IsRestricted: bool(isRestricted),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
