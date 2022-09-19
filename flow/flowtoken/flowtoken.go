package flowtoken

import (
	"strconv"

	"github.com/nakji-network/connectors/flow"
	"github.com/onflow/cadence"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var Types = []proto.Message{
	&TokensWithdrawn{},
	&TokensDeposited{},
}

type SmartContract struct{}

func (c SmartContract) Address() string {
	return "1654653399040a61"
}

func (c SmartContract) Name() string {
	return "FlowToken"
}

func (c SmartContract) Events() []string {
	return []string{
		"A.1654653399040a61.FlowToken.TokensInitialized",
		"A.1654653399040a61.FlowToken.TokensWithdrawn",
		"A.1654653399040a61.FlowToken.TokensDeposited",
		"A.1654653399040a61.FlowToken.TokensMinted",
		"A.1654653399040a61.FlowToken.TokensBurned",
		"A.1654653399040a61.FlowToken.MinterCreated",
		"A.1654653399040a61.FlowToken.BurnerCreated",
	}
}

func (c SmartContract) Message(vLog flow.Log) proto.Message {
	switch vLog.Type.EventName {
	case "TokensInitialized":
		initialSupply := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TokensInitialized{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.InitialSupply, _ = strconv.ParseFloat(initialSupply.String(), 64)
		return msg
	case "TokensWithdrawn":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		from := vLog.Value.Fields[1].(cadence.Optional)
		msg := &TokensWithdrawn{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.Amount, _ = strconv.ParseFloat(amount.String(), 64)
		if from.Value != nil {
			msg.From = from.Value.(cadence.Address).Bytes()
		}
		return msg
	case "TokensDeposited":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		to := vLog.Value.Fields[1].(cadence.Optional)
		msg := &TokensDeposited{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.Amount, _ = strconv.ParseFloat(amount.String(), 64)
		if to.Value != nil {
			msg.To = to.Value.(cadence.Address).Bytes()
		}
		return msg
	case "TokensMinted":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TokensMinted{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.Amount, _ = strconv.ParseFloat(amount.String(), 64)
		return msg
	case "TokensBurned":
		amount := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &TokensBurned{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.Amount, _ = strconv.ParseFloat(amount.String(), 64)
		return msg
	case "MinterCreated":
		allowedAmount := vLog.Value.Fields[0].(cadence.UFix64)
		msg := &MinterCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		msg.AllowedAmount, _ = strconv.ParseFloat(allowedAmount.String(), 64)
		return msg
	case "BurnerCreated":
		msg := &BurnerCreated{
			Ts:          timestamppb.New(vLog.Timestamp),
			BlockNumber: vLog.Height,
			TxID:        vLog.TransactionID[:],
			LogIndex:    uint64(vLog.TransactionIndex),
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", vLog.Type.EventName)
		return nil
	}
}
