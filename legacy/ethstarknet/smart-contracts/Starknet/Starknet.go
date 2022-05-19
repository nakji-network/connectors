package Starknet

import (
	"bytes"
	"encoding/gob"

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
	case "ConsumedMessageToL1":
		e := new(StarknetConsumedMessageToL1)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(e.Payload)

		return &ConsumedMessageToL1{
			Ts:          timestamp,
			Fromaddress: e.FromAddress.Bytes(),
			Toaddress:   e.ToAddress.Bytes(),
			Payload:     buf.Bytes(),
		}
	case "ConsumedMessageToL2":
		e := new(StarknetConsumedMessageToL2)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(e.Payload)

		return &ConsumedMessageToL2{
			Ts:          timestamp,
			Fromaddress: e.FromAddress.Bytes(),
			Toaddress:   e.ToAddress.Bytes(),
			Selector:    e.Selector.Bytes(),
			Payload:     buf.Bytes(),
			Nonce:       e.Nonce.Bytes(),
		}
	case "LogMessageToL1":
		e := new(StarknetLogMessageToL1)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(e.Payload)

		return &LogMessageToL1{
			Ts:          timestamp,
			Fromaddress: e.FromAddress.Bytes(),
			Toaddress:   e.ToAddress.Bytes(),
			Payload:     buf.Bytes(),
		}
	case "LogMessageToL2":
		e := new(StarknetLogMessageToL2)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(e.Payload)

		return &LogMessageToL2{
			Ts:          timestamp,
			Fromaddress: e.FromAddress.Bytes(),
			Toaddress:   e.ToAddress.Bytes(),
			Selector:    e.Selector.Bytes(),
			Payload:     buf.Bytes(),
			Nonce:       e.Nonce.Bytes(),
		}
	case "LogNewGovernorAccepted":
		e := new(StarknetLogNewGovernorAccepted)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogNewGovernorAccepted{
			Ts:               timestamp,
			AcceptedGovernor: e.AcceptedGovernor.Bytes(),
		}
	case "LogNominatedGovernor":
		e := new(StarknetLogNominatedGovernor)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogNominatedGovernor{
			Ts:                timestamp,
			NominatedGovernor: e.NominatedGovernor.Bytes(),
		}
	case "LogNominationCancelled":
		e := new(StarknetLogNominationCancelled)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogNominationCancelled{
			Ts: timestamp,
		}
	case "LogOperatorAdded":
		e := new(StarknetLogOperatorAdded)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogOperatorAdded{
			Ts:       timestamp,
			Operator: e.Operator.Bytes(),
		}
	case "LogOperatorRemoved":
		e := new(StarknetLogOperatorRemoved)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogOperatorRemoved{
			Ts:       timestamp,
			Operator: e.Operator.Bytes(),
		}
	case "LogRemovedGovernor":
		e := new(StarknetLogRemovedGovernor)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogRemovedGovernor{
			Ts:              timestamp,
			RemovedGovernor: e.RemovedGovernor.Bytes(),
		}
	case "LogStateTransitionFact":
		e := new(StarknetLogStateTransitionFact)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogStateTransitionFact{
			Ts:                  timestamp,
			StateTransitionFact: e.StateTransitionFact[:],
		}
	case "LogStateUpdate":
		e := new(StarknetLogStateUpdate)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogStateUpdate{
			Ts:          timestamp,
			GlobalRoot:  e.GlobalRoot.Bytes(),
			BlockNumber: e.BlockNumber.Bytes(),
		}
	}
	return nil
}
