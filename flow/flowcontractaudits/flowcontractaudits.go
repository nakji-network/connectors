package flowcontractaudits

import (
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&AuditorCreated{},
	&VoucherCreated{},
	&VoucherUsed{},
	&VoucherRemoved{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "e467b9dd11fa00df"
}

func (c SmartContract) Name() string {
	return "FlowContractAudits"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.e467b9dd11fa00df.FlowContractAudits.AuditorCreated",
		"A.e467b9dd11fa00df.FlowContractAudits.VoucherCreated",
		"A.e467b9dd11fa00df.FlowContractAudits.VoucherUsed",
		"A.e467b9dd11fa00df.FlowContractAudits.VoucherRemoved",
	}
}

func (c SmartContract) Message(vLog *flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "AuditorCreated":
		msg := &AuditorCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		return msg
	case "VoucherCreated":
		address := vLog.Value.Fields[0].(cadence.Optional)
		recurrent := vLog.Value.Fields[1].(cadence.Bool)
		expiryBlockHeight := vLog.Value.Fields[2].(cadence.Optional)
		codeHash := vLog.Value.Fields[3].(cadence.String)
		msg := &VoucherCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Recurrent:   bool(recurrent),
			CodeHash:    string(codeHash),
		}
		if address.Value != nil {
			msg.Address = address.Value.(cadence.Address).Bytes()
		}
		if expiryBlockHeight.Value != nil {
			msg.ExpiryBlockHeight = uint64(expiryBlockHeight.Value.(cadence.UInt64))
		}
		return msg
	case "VoucherUsed":
		address := vLog.Value.Fields[0].(cadence.Address)
		key := vLog.Value.Fields[1].(cadence.String)
		recurrent := vLog.Value.Fields[2].(cadence.Bool)
		expiryBlockHeight := vLog.Value.Fields[3].(cadence.Optional)
		msg := &VoucherUsed{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Address:     address.Bytes(),
			Key:         string(key),
			Recurrent:   bool(recurrent),
		}
		if expiryBlockHeight.Value != nil {
			msg.ExpiryBlockHeight = uint64(expiryBlockHeight.Value.(cadence.UInt64))
		}
		return msg
	case "VoucherRemoved":
		key := vLog.Value.Fields[0].(cadence.String)
		recurrent := vLog.Value.Fields[1].(cadence.Bool)
		expiryBlockHeight := vLog.Value.Fields[2].(cadence.Optional)
		msg := &VoucherRemoved{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
			Key:         string(key),
			Recurrent:   bool(recurrent),
		}
		if expiryBlockHeight.Value != nil {
			msg.ExpiryBlockHeight = uint64(expiryBlockHeight.Value.(cadence.UInt64))
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
