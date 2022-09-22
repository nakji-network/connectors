// Package chain contains types for the solana connector. These
// types represent the native entities on the Solana blockchain.

package chain

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.solana.0_0_0.solana_block":       &Block{},
	"nakji.solana.0_0_0.solana_account":     &Account{},
	"nakji.solana.0_0_0.solana_message":     &Message{},
	"nakji.solana.0_0_0.solana_transaction": &Transaction{},
	"nakji.solana.0_0_0.solana_reward":      &BlockReward{},
	"nakji.solana.0_0_0.solana_solana":      &Solana{},
}

func (b *Block) UnmarshalSolanaBlock(in *rpc.GetBlockResult) {
	signatures := make([]string, len(in.Signatures))
	for i, signature := range in.Signatures {
		signatures[i] = signature.String()
	}

	*b = Block{
		Hash:              in.Blockhash.String(),
		PreviousBlockhash: in.PreviousBlockhash.String(),
		ParentSlot:        uint64(in.ParentSlot),
		Signatures:        signatures,
		BlockTime:         int64(*in.BlockTime),
		BlockHeight:       uint64(*in.BlockHeight),
	}
}

func (a *Account) UnmarshalSolanaAccount(in *rpc.Account, pubKey solana.PublicKey) {
	*a = Account{
		PublicKey:    pubKey.String(),
		Lamports:     uint64(in.Lamports),
		Executable:   in.Executable,
		RentEpoch:    uint64(in.RentEpoch),
		DataEncoding: "jsonParsed",
	}

	if !in.Owner.IsZero() && in.Owner.String() != "" {
		aString := in.Owner.String()
		a.Owner = &aString
	}

	dataJson, err := in.Data.MarshalJSON()
	if err != nil {
		log.Error().Err(err).Str("account pubkey", pubKey.String()).Msg("failed to marshal json data for account")
	}

	a.DataJson = dataJson
}

func (r *BlockReward) UnmarshalSolanaBlockRewardForTx(in rpc.BlockReward, tx *Transaction) {
	r.UnmarshalSolanaBlockReward(in, tx.BlockHash)
	r.TransactionId = tx.Id
}

func (r *BlockReward) UnmarshalSolanaBlockReward(in rpc.BlockReward, blockHash string) {
	value := BlockReward_RewardType_value[string(in.RewardType)]
	rewardType := BlockReward_RewardType(value).Enum()

	*r = BlockReward{
		BlockHash:   blockHash,
		Pubkey:      in.Pubkey.String(),
		Lamports:    int64(in.Lamports),
		PostBalance: uint64(in.PostBalance),
		RewardType:  *rewardType,
	}
}

func (s *Solana) UnmarshalSolana(msg isSolana_Msg, table string) {
	*s = Solana{
		Table: table,
		Msg:   msg,
	}
}

func (t *Transaction) UnmarshalSolanaTransaction(in rpc.TransactionWithMeta, blockHash string) {
	preBalances := make([]uint64, len(in.Meta.PreBalances))
	for i, prebalance := range in.Meta.PreBalances {
		preBalances[i] = prebalance
	}

	postBalances := make([]uint64, len(in.Meta.PostBalances))
	for i, postbalance := range in.Meta.PostBalances {
		postBalances[i] = postbalance
	}

	signatures := make([]string, len(in.Transaction.Signatures))
	for i, signature := range in.Transaction.Signatures {
		signatures[i] = signature.String()
	}

	// When tested, preTokenBalances were always empty slices. This means that writing to PG
	// may lead to errors when these are not empty and may require extra handling of slices of
	// `protoreflect.Messaage`s to convert them to jsonb literals in Postgres.
	preTokenBalances := make([]*TokenBalance, len(in.Meta.PreTokenBalances))
	for i, preTokenBalance := range in.Meta.PreTokenBalances {
		preTokenBalances[i] = &TokenBalance{
			AccountIndex: uint32(preTokenBalance.AccountIndex),
			Mint:         preTokenBalance.Mint.String(),
			UiTokenAmount: &UiTokenAmount{
				Amount:   preTokenBalance.UiTokenAmount.Amount,
				Decimals: uint32(preTokenBalance.UiTokenAmount.Decimals),
				UiAmount: preTokenBalance.UiTokenAmount.UiAmountString,
			},
		}
	}

	// When tested, postTokenBalances were always empty slices. This means that writing to PG
	// may lead to errors when these are not empty and may require extra handling of slices of
	// `protoreflect.Messaage`s to convert them to jsonb literals in Postgres.
	postTokenBalances := make([]*TokenBalance, len(in.Meta.PostTokenBalances))
	for i, postTokenBalance := range in.Meta.PostTokenBalances {
		postTokenBalances[i] = &TokenBalance{
			AccountIndex: uint32(postTokenBalance.AccountIndex),
			Mint:         postTokenBalance.Mint.String(),
			UiTokenAmount: &UiTokenAmount{
				Amount:   postTokenBalance.UiTokenAmount.Amount,
				Decimals: uint32(postTokenBalance.UiTokenAmount.Decimals),
				UiAmount: postTokenBalance.UiTokenAmount.UiAmountString,
			},
		}
	}

	err := ""
	switch metaErr := in.Meta.Err.(type) {
	case map[string][]interface{}:
		for k, v := range metaErr {
			err = k
			for _, detail := range v {
				err = fmt.Sprintf("%s %v", err, detail)
			}
		}
	default:
		err = ""
	}

	*t = Transaction{
		Id:                in.Transaction.Signatures[0].String(),
		BlockHash:         blockHash,
		Err:               err,
		Fee:               in.Meta.Fee,
		PreBalances:       preBalances,
		PostBalances:      postBalances,
		PreTokenBalances:  preTokenBalances,
		PostTokenBalances: postTokenBalances,
		LogMessages:       in.Meta.LogMessages,
	}
}

func (m *Message) UnmarshalSolanaMsg(in solana.Message, tx *Transaction) {
	accountKeys := make([]string, len(in.AccountKeys))
	for i, accountKey := range in.AccountKeys {
		accountKeys[i] = accountKey.String()
	}

	instructions := make([]*CompiledInstruction, len(in.Instructions))
	for i, instruction := range in.Instructions {
		accounts := make([]uint32, len(instruction.Accounts))
		for i, account := range instruction.Accounts {
			accounts[i] = uint32(account)
		}

		instructions[i] = &CompiledInstruction{
			ProgramIdIndex: uint32(instruction.ProgramIDIndex),
			AccountCount:   uint32(instruction.AccountCount),
			DataLength:     uint32(instruction.DataLength),
			Accounts:       accounts,
			Data:           instruction.Data.String(),
		}
	}

	*m = Message{
		TransactionId:               tx.Id,
		NumRequiredSignatures:       uint32(in.Header.NumRequiredSignatures),
		NumReadOnlySignedAccounts:   uint32(in.Header.NumReadonlySignedAccounts),
		NumReadOnlyUnsignedAccounts: uint32(in.Header.NumReadonlyUnsignedAccounts),
		AccountKeys:                 accountKeys,
		RecentBlockHash:             in.RecentBlockhash.String(),
		Instructions:                instructions,
	}
}
