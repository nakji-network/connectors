package mai

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) proto.Message {
	switch eventName {
	case "Approval":
		e := new(MaiApproval)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			Owner:   e.Owner.Bytes(),
			Spender: e.Spender.Bytes(),
			Value:   e.Value.Bytes(),
		}
	case "BorrowToken":
		e := new(MaiBorrowToken)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &BorrowToken{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
			Amount:  e.Amount.Bytes(),
		}
	case "BuyRiskyVault":
		e := new(MaiBuyRiskyVault)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &BuyRiskyVault{
			Ts:         timestamp,
			Block:      vLog.BlockNumber,
			Idx:        uint64(vLog.Index),
			Tx:         vLog.TxHash.Bytes(),
			VaultID:    e.VaultID.Bytes(),
			Owner:      e.Owner.Bytes(),
			Buyer:      e.Buyer.Bytes(),
			AmountPaid: e.AmountPaid.Bytes(),
		}
	case "CreateVault":
		e := new(MaiCreateVault)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &CreateVault{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
			Creator: e.Creator.Bytes(),
		}
	case "DepositCollateral":
		e := new(MaiDepositCollateral)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DepositCollateral{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
			Amount:  e.Amount.Bytes(),
		}
	case "DestroyVault":
		e := new(MaiDestroyVault)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DestroyVault{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(MaiOwnershipTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
			Ts:            timestamp,
			Block:         vLog.BlockNumber,
			Idx:           uint64(vLog.Index),
			Tx:            vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "PayBackToken":
		e := new(MaiPayBackToken)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &PayBackToken{
			Ts:         timestamp,
			Block:      vLog.BlockNumber,
			Idx:        uint64(vLog.Index),
			Tx:         vLog.TxHash.Bytes(),
			VaultID:    e.VaultID.Bytes(),
			Amount:     e.Amount.Bytes(),
			ClosingFee: e.ClosingFee.Bytes(),
		}
	case "Transfer":
		e := new(MaiTransfer)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:    timestamp,
			Block: vLog.BlockNumber,
			Idx:   uint64(vLog.Index),
			Tx:    vLog.TxHash.Bytes(),
			From:  e.From.Bytes(),
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	case "TransferVault":
		e := new(MaiTransferVault)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &TransferVault{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
		}
	case "WithdrawCollateral":
		e := new(MaiWithdrawCollateral)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WithdrawCollateral{
			Ts:      timestamp,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			VaultID: e.VaultID.Bytes(),
			Amount:  e.Amount.Bytes(),
		}
	}
	return nil
}
