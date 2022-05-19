// Logic for manually parsing trades for NFT's and publishing them
package solanatoken

import (
	"context"
	"fmt"
	"time"

	"blep.ai/data/chain/solana/connector"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	MagicEdenProgramId     = "MEisE1HzehtrDpAAT8PnLHjpSSkRYakotTuJRPjTpo8"
	MagicEdenEscrowAccount = "GUfCR9mK6azb9vcpsxgXyj7XRPAKJd4KMHTTVvtncGgp"
	MagicEdenV2ProgramId   = "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K"
	SolanartProgramId      = "CJsLwbP1iu5DuUikHEJnLfANgKy6stB2uFgvBBHoyxwz"
	SolanartFeeAddress     = "39fEpihLATXPJCQuSiXLUSiCbGchGYjeL39eyXh3KbyT"
	SolseaProgramId        = "617jbWo616ggkDxvW1Le8pV38XLbVSyWY8ae6QUmGBAU"
	SolseaFeeAddress       = "6T4f5bdrd9ffTtehqAj9BGyxahysRGcaUZeDzA1XN52N"
	ExchangeArtProgramId   = "AmK5g2XcyptVLCFESBCJqoSfwV3znGoVYQnqEnaAZKWn"
	DigitalEyesProgramId   = "A7p8451ktDCHq5yYaHczeLMYsjRsAkzc3hCXcSrwYHU7"
	DigitalEyesFeeAddress  = "3iYf9hHQPciwgJ1TCjpRUp1A3QW4AfaK7J6vCmETRMuu"
)

func (c *SolanaTokenConnector) parseTrade(txResult *rpc.GetTransactionResult, tx *solana.Transaction) *Trade {
	idx := tx.Message.Instructions[0].ProgramIDIndex
	programId := tx.Message.AccountKeys[idx]

	switch programId.String() {
	case MagicEdenProgramId:
		return c.parseMagicEdenTrade(txResult, tx)
	case MagicEdenV2ProgramId:
		return c.parseMagicEdenV2Trade(txResult, tx)
	}

	if len(tx.Message.Instructions) < 2 {
		return nil
	}

	idx = tx.Message.Instructions[1].ProgramIDIndex
	programId = tx.Message.AccountKeys[idx]

	switch programId.String() {
	case SolanartProgramId:
		return c.parseTokenTrade("solanart", txResult, tx)
	case ExchangeArtProgramId:
		return c.parseTokenTrade("exchangeart", txResult, tx)
	case DigitalEyesProgramId:
		return c.parseTokenTrade("digitaleyes", txResult, tx)
	}

	if len(tx.Message.Instructions) < 3 {
		return nil
	}

	idx = tx.Message.Instructions[2].ProgramIDIndex
	programId = tx.Message.AccountKeys[idx]

	if programId.String() == SolseaProgramId {
		return c.parseSolseaTrade(txResult, tx)
	}

	return nil
}

// parseTokenTrade parses trades of Solanart, ExchangeArt and DigitalEyes. They all share the same transactions
// for trading NFTs.
// A sample ExchangeArt trade looks like:
// https://solscan.io/tx/3UgMV5id6wd4LK7HSiHPo2G8qWTpg32yTPaDTJsKCtErr62FCS4izZ815Jpa1XfhccwiJ3rywZqnVejgbq7hFLXS
// A sample Solanart trade looks like:
// https://solscan.io/tx/Y3oaunErxoBdqwzDbwGNHNhjJ9A8WoEJvXb3XM4r3F9NTCvnnkWhRzfo71WCJtM2zbkXrjWrkqe2SiaMTso3ZtJ
// A sample DigitalEyes trade looks like:
// https://solscan.io/tx/tPyKk6ucj8Mca7r9mdUfzcz3D9rLcc4FYivS4nuRfEpkn7Q4PywWgsHkECYizNzYFX46hNwbFQgJRynE26K4HL3
func (c *SolanaTokenConnector) parseTokenTrade(
	exchange string,
	txResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
) *Trade {
	if len(txResult.Meta.InnerInstructions) < 2 {
		return nil
	}
	createAssocAccountInstructions := txResult.Meta.InnerInstructions[0].Instructions
	initializeAccountInstruction := createAssocAccountInstructions[len(createAssocAccountInstructions)-1]
	buyInstructions := txResult.Meta.InnerInstructions[1].Instructions

	if len(buyInstructions) < 4 {
		return nil
	}

	// Get NFT's TokenAddress through SPL-Token's InitializeAccount instruction
	initializeAccount, ok := parseTokenInitializeAccount(initializeAccountInstruction, tx)
	if !ok {
		return nil
	}

	// Verify existence of a Token.Transfer instruction
	secondLastInstruction := buyInstructions[len(buyInstructions)-2]
	_, ok = parseTokenTransfer(secondLastInstruction, tx)
	if !ok {
		return nil
	}

	// Get buyer and seller info from the last SOL Transfer instruction
	transfer, err := getLastSolTransfer(buyInstructions, tx)
	if err != nil {
		log.Fatal().Err(err).Msg("no sol transfers")
	}

	timestamp := timestamppb.New(txResult.BlockTime.Time())
	trade := &Trade{
		Ts:          timestamp,
		Transaction: tx.Signatures[0].String(),
		SoldOn:      exchange,
		NftAddress:  initializeAccount.GetMintAccount().PublicKey.String(),
		Buyer:       transfer.GetFundingAccount().PublicKey.String(),
		Seller:      transfer.GetRecipientAccount().PublicKey.String(),
		Price:       getPrice(buyInstructions, tx),
	}

	log.Debug().Interface(exchange+"trade", trade).Msgf("parsed %s trade", exchange)

	return trade
}

// Manually parse a Solsea trade. A sample Solsea trade looks like:
// https://solscan.io/tx/484ATxiDpTvzB9XLekCnYGDtAcKW5kb6Zy4civ3zYmUMffqaSnuBas6ubwumnp8rsJ2Gb4QXDBu3S6u1Tc6v6s5Y
func (c *SolanaTokenConnector) parseSolseaTrade(
	txResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
) *Trade {
	initializeAccountInstruction := tx.Message.Instructions[1]
	buyInstructions := txResult.Meta.InnerInstructions[0].Instructions

	if len(buyInstructions) < 4 {
		return nil
	}

	// Get NFT's TokenAddress through SPL-Token's InitializeAccount instruction
	initializeAccount, ok := parseTokenInitializeAccount(initializeAccountInstruction, tx)
	if !ok {
		return nil
	}

	// Verify existence of a Token.Transfer instruction
	lastInstruction := buyInstructions[len(buyInstructions)-1]
	_, ok = parseTokenTransfer(lastInstruction, tx)
	if !ok {
		return nil
	}

	// Verify that the 1st buy instruction is a SOL transfer to Solsea's Fee Address
	transferToFeeAddress, err := connector.DecodeSystemInstruction(buyInstructions[0], tx)
	if err != nil {
		return nil
	}
	transfer, ok := transferToFeeAddress.Impl.(*system.Transfer)
	if !ok {
		return nil
	}
	if transfer.GetRecipientAccount().PublicKey.String() != SolseaFeeAddress {
		return nil
	}

	// Get buyer and seller info from the last SOL Transfer instruction
	transfer, err = getLastSolTransfer(buyInstructions, tx)
	if err != nil {
		log.Fatal().Err(err).Msg("no sol transfers")
	}

	timestamp := timestamppb.New(txResult.BlockTime.Time())
	trade := &Trade{
		Ts:          timestamp,
		Transaction: tx.Signatures[0].String(),
		SoldOn:      "solsea",
		NftAddress:  initializeAccount.GetMintAccount().PublicKey.String(),
		Buyer:       transfer.GetFundingAccount().PublicKey.String(),
		Seller:      transfer.GetRecipientAccount().PublicKey.String(),
		Price:       getPrice(buyInstructions, tx),
	}

	log.Debug().Interface("solseaTrade", trade).Msg("parsed Solsea trade")

	return trade
}

func (c *SolanaTokenConnector) parseMagicEdenV2Trade(
	txResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
) *Trade {
	instructions := tx.Message.Instructions

	// A Buy transaction in MagicEdenV2 has 3 instructions or 2 instructions
	// in some rare cases. Like this: https://solscan.io/tx/47mrYfEq7sZRt23AsgKHk32soKREk9UneDScmZbBN6QZBLNHW6LHieHAE26X4q15JfBYjNDUnMiqCUhp3Z8fQjDQ
	if len(instructions) != 2 && len(instructions) != 3 {
		return nil
	}

	// Second to last InnerInstruction must be a TokenTransfer for it to be a Buy instruction
	var lastInnerInstructions []solana.CompiledInstruction
	for _, innerInstruction := range txResult.Meta.InnerInstructions {
		if innerInstruction.Index == uint16(len(instructions)-1) {
			lastInnerInstructions = innerInstruction.Instructions
		}
	}

	// Buy instructions must be at least 8
	if len(lastInnerInstructions) < 8 {
		return nil
	}

	secondToLastInst := lastInnerInstructions[len(lastInnerInstructions)-2]
	_, ok := parseTokenTransfer(secondToLastInst, tx)
	if !ok {
		return nil
	}

	lastInstruction := instructions[len(instructions)-1]
	buyer := tx.Message.AccountKeys[lastInstruction.Accounts[0]]
	seller := tx.Message.AccountKeys[lastInstruction.Accounts[1]]
	tokenAddress := tx.Message.AccountKeys[lastInstruction.Accounts[4]]
	timestamp := timestamppb.New(txResult.BlockTime.Time())

	var price uint64
	for _, instruction := range lastInnerInstructions {
		transferInstruction, err := connector.DecodeSystemInstruction(instruction, tx)
		if err != nil {
			break
		}

		transfer, ok := transferInstruction.Impl.(*system.Transfer)
		if !ok {
			break
		}

		price += *transfer.Lamports
	}

	trade := &Trade{
		Ts:          timestamp,
		Transaction: tx.Signatures[0].String(),
		SoldOn:      "magicedenv2",
		Buyer:       buyer.String(),
		Seller:      seller.String(),
		NftAddress:  tokenAddress.String(),
		Price:       price,
	}

	log.Debug().Interface("magicedenv2", trade).Msg("parsed MagicEdenV2 trade")

	return trade
}

// Manually parse a MagicEden trade. A sample MagicEden trade looks like:
// https://solscan.io/tx/A6ka4Yu7XxfGuo3CpJRZSGYxRJEY5Yh7iJXvS5RPW1rWeVTkXvZ4i899FJYfsDCoWNoHLgWY5evFrgWUsDtjCAm
func (c *SolanaTokenConnector) parseMagicEdenTrade(
	txResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
) *Trade {
	instructions := txResult.Meta.InnerInstructions[0].Instructions

	// A Buy instruction in MagicEden has 4 or more inner instructions
	if len(instructions) < 4 {
		return nil
	}

	// Last instruction must be a token.SetAuthority instruction
	lastInstruction := instructions[len(instructions)-1]
	tokenInstruction, err := connector.DecodeTokenInstruction(lastInstruction, tx)
	if err != nil {
		return nil
	}

	setAuthority, ok := tokenInstruction.Impl.(*token.SetAuthority)
	if !ok {
		return nil
	}

	// SetAuthority instruction must change Account Owner from Magic Eden's Escrow account to
	// a new account
	if *setAuthority.AuthorityType != token.AuthorityAccountOwner {
		return nil
	}
	authority := setAuthority.GetAuthorityAccount()
	if authority.PublicKey.String() != MagicEdenEscrowAccount {
		return nil
	}

	// Get buyer and seller info from the last SOL Transfer instruction
	transfer, err := getLastSolTransfer(instructions, tx)
	if err != nil {
		log.Fatal().Err(err).Msg("no sol transfers")
	}

	tokenAccountPubKey := setAuthority.GetSubjectAccount().PublicKey
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	trade := &Trade{
		Ts:          timestamp,
		Transaction: tx.Signatures[0].String(),
		SoldOn:      "magiceden",
		Buyer:       transfer.GetFundingAccount().PublicKey.String(),
		Seller:      transfer.GetRecipientAccount().PublicKey.String(),
		Price:       getPrice(instructions, tx),
		NftAddress:  c.getNFTAddress(tokenAccountPubKey),
	}

	log.Debug().Interface("magicEdenTrade", trade).Msg("parsed MagicEden trade")

	return trade
}

// getLastSolTransfer returns the last system.Transfer in a slice of instruction which we use to get buyer and seller info.
func getLastSolTransfer(buyInstructions []solana.CompiledInstruction, tx *solana.Transaction) (*system.Transfer, error) {
	for i := len(buyInstructions) - 1; i >= 0; i-- {
		buyInstruction := buyInstructions[i]
		transferInstruction, err := connector.DecodeSystemInstruction(buyInstruction, tx)
		if err != nil {
			continue
		}
		transfer, ok := transferInstruction.Impl.(*system.Transfer)
		if !ok {
			continue
		}
		return transfer, nil
	}

	return nil, fmt.Errorf("no sol transfer")
}

// getPrice returns uint64 price from all Sol Transfers in the Trade/Buy instruction
func getPrice(buyInstructions []solana.CompiledInstruction, tx *solana.Transaction) uint64 {
	var price uint64
	for _, instruction := range buyInstructions {
		transferInstruction, err := connector.DecodeSystemInstruction(instruction, tx)
		if err != nil {
			continue
		}
		transfer, ok := transferInstruction.Impl.(*system.Transfer)
		if !ok {
			continue
		}

		price += *transfer.Lamports
	}
	return price
}

// parseTokenInitializeAccount decodes the compiled instruction into a Token.InitializeAccount
func parseTokenInitializeAccount(instruction solana.CompiledInstruction, tx *solana.Transaction) (*token.InitializeAccount, bool) {
	tokenInstruction, err := connector.DecodeTokenInstruction(instruction, tx)
	if err != nil {
		return nil, false
	}

	initializeAccount, ok := tokenInstruction.Impl.(*token.InitializeAccount)
	if !ok {
		return nil, false
	}

	return initializeAccount, true
}

// parseTokenInitializeAccount decodes the compiled instruction into a Token.Transfer
func parseTokenTransfer(instruction solana.CompiledInstruction, tx *solana.Transaction) (*token.Transfer, bool) {
	tokenTransferInstruction, err := connector.DecodeTokenInstruction(instruction, tx)
	if err != nil {
		return nil, false
	}

	tokenTransfer, ok := tokenTransferInstruction.Impl.(*token.Transfer)
	if !ok {
		return nil, false
	}

	return tokenTransfer, true
}

// We get the NFT's address via the TokenAccount. If the new authority is no longer a token account,
// we get the last transaction of that Account and get the NFT's address in the InitializeAccount
// instruction. If an Account is no longer a Token Account, it is because token.CloseAccount instruction
// was run on it. This means its last transaction includes CloseAccount and InitializeAccount instructions.
func (c *SolanaTokenConnector) getNFTAddress(tokenAccountPubKey solana.PublicKey) string {
	ctx := context.Background()
	tokenAccount, err := c.GetTokenAccount(ctx, tokenAccountPubKey)

	if err != nil {
		// Get last transaction which is the transaction that closes the token account and
		// reverts it to a regular account.
		var txSignatures []*rpc.TransactionSignature
		limit := 1

		//`GetSignaturesForAddressWithOpts` sometimes returns 0 even when transactions exist. An
		// account/address will always have at least 1 transaction. We retry until we get 1
		// signature.
		for len(txSignatures) == 0 {
			opts := &rpc.GetSignaturesForAddressOpts{Limit: &limit}
			txSignatures, err = c.Client.GetSignaturesForAddressWithOpts(ctx, tokenAccountPubKey, opts)
			if err != nil {
				log.Fatal().Err(err).Msg("failed getting signature for account")
			}
			time.Sleep(time.Duration(500) * time.Millisecond)
		}

		txResult, tx := c.GetTransaction(ctx, txSignatures[0].Signature)
		tokenInstructions := c.GetTokenInstructions(txResult, tx)

		// We find the InitializeAccount instruction since that instruction has the
		// NFT's address data.
		for _, tokenInstruction := range tokenInstructions {
			initializeAccount, ok := tokenInstruction.Impl.(*token.InitializeAccount)
			if !ok {
				continue
			}
			mint := initializeAccount.GetMintAccount()
			return mint.PublicKey.String()
		}
	}
	return tokenAccount.Mint.String()
}
