package ArbitrumDaiBridge

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Closed":
		e := new(ARBITRUMDAIBRIDGEClosed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Closed{
			Ts: timestamp,
		}
	case "Deny":
		e := new(ARBITRUMDAIBRIDGEDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "DepositInitiated":
		e := new(ARBITRUMDAIBRIDGEDepositInitiated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DepositInitiated{
			Ts:             timestamp,
			L1Token:        e.L1Token.Bytes(),
			From:           e.From.Bytes(),
			To:             e.To.Bytes(),
			SequenceNumber: e.SequenceNumber.Bytes(),
			Amount:         e.Amount.Bytes(),
		}
	case "Rely":
		e := new(ARBITRUMDAIBRIDGERely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "TxToL2":
		e := new(ARBITRUMDAIBRIDGETxToL2)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &TxToL2{
			Ts:     timestamp,
			From:   e.From.Bytes(),
			To:     e.To.Bytes(),
			SeqNum: e.SeqNum.Bytes(),
			Data:   e.Data[:],
		}
	case "WithdrawalFinalized":
		e := new(ARBITRUMDAIBRIDGEWithdrawalFinalized)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WithdrawalFinalized{
			Ts:      timestamp,
			L1Token: e.L1Token.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
			ExitNum: e.ExitNum.Bytes(),
			Amount:  e.Amount.Bytes(),
		}
	}
	return nil
}
