package solanatoken

import (
	"context"
	"os"
	"os/signal"

	"blep.ai/data/chain/solana/connector"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	Namespace = "solanatoken"
)

type SolanaTokenConnector struct {
	*connector.SolanaConnector
}

func NewConnector(
	conf *viper.Viper,
	topics map[string]kafkautils.Topic,
) *SolanaTokenConnector {
	connector := connector.NewConnector(Namespace, conf, topics)

	return &SolanaTokenConnector{
		SolanaConnector: connector,
	}
}

var NFTPubKeys = map[string]solana.PublicKey{
	"magiceden":   solana.MustPublicKeyFromBase58(MagicEdenProgramId),
	"magicedenv2": solana.MustPublicKeyFromBase58(MagicEdenV2ProgramId),
	"solanart":    solana.MustPublicKeyFromBase58(SolanartProgramId),
	"solsea":      solana.MustPublicKeyFromBase58(SolseaProgramId),
	"exchangeart": solana.MustPublicKeyFromBase58(ExchangeArtProgramId),
	"digitaleyes": solana.MustPublicKeyFromBase58(DigitalEyesProgramId),
}

func (c *SolanaTokenConnector) Start() {
	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	backfillLimit := c.Conf.GetInt("solanatoken.backfillLimit")
	defer cancel()

	for namespace, pubKey := range NFTPubKeys {
		go c.ProgramListener(ctx, pubKey, namespace, c.processTransaction)
		go c.ProgramBackfiller(ctx, pubKey, namespace, c.processBackfillTransaction, backfillLimit)
	}

	go c.ProgramListener(ctx, solana.TokenMetadataProgramID, "solananft", c.processNFTTransaction)
	go c.ProgramBackfiller(ctx, solana.TokenMetadataProgramID, "solananft", c.processNFTBackfillTransaction, backfillLimit)

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			c.WsClient.Close()
			c.Producer.Close()
			return
		}
	}
}

func (c *SolanaTokenConnector) processTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) {
	if connector.HasError(txResult.Meta.Err) {
		return
	}

	trade := c.parseTrade(txResult, tx)
	if trade != nil {
		msg := &kafkautils.Message{
			Topic:    c.Topics["tokentrade"],
			Key:      kafkautils.NewKey(Namespace+"trade", string(trade.Transaction)),
			ProtoMsg: trade,
		}
		c.QueueSink <- msg
	}
}

func (c *SolanaTokenConnector) processBackfillTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, _ solana.Signature) {
	if connector.HasError(txResult.Meta.Err) {
		return
	}

	trade := c.parseTrade(txResult, tx)

	if trade != nil {
		msg := &kafkautils.Message{
			Topic:    c.Topics["tokentrade"],
			Key:      kafkautils.NewKey(Namespace+"trade", string(trade.Transaction)),
			ProtoMsg: trade,
		}

		db := c.Db.DbInterface.(*database.Database)
		if err := db.InsertKafkaMessage(*msg); err != nil {
			log.Fatal().Err(err).Msg("Database write error fail fast")
		}
	}
}

func (c *SolanaTokenConnector) processNFTTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) {
	if connector.HasError(txResult.Meta.Err) {
		log.Debug().Msgf("tx %s has error: %v", signature.String(), txResult.Meta.Err)
		return
	}

	nft := c.parseNFT(txResult, tx)

	if nft != nil {
		msg := &kafkautils.Message{
			Topic:    c.Topics["tokennft"],
			Key:      kafkautils.NewKey(Namespace+"nft", string(nft.Address)),
			ProtoMsg: nft,
		}
		c.QueueSink <- msg
	}
}

func (c *SolanaTokenConnector) processNFTBackfillTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) {
	if connector.HasError(txResult.Meta.Err) {
		log.Debug().Msgf("tx %s has error: %v", signature.String(), txResult.Meta.Err)
		return
	}

	nft := c.parseNFT(txResult, tx)

	if nft != nil {
		msg := &kafkautils.Message{
			Topic:    c.Topics["tokennft"],
			Key:      kafkautils.NewKey(Namespace+"nft", string(nft.Address)),
			ProtoMsg: nft,
		}

		db := c.Db.DbInterface.(*database.Database)
		if err := db.InsertKafkaMessage(*msg); err != nil {
			log.Fatal().Err(err).Msg("Database write error fail fast")
		}
	}
}

// parseNFT parses the transaction to see whether a new NFT was minted/created and gets all the
// relevant NFT data.
func (c *SolanaTokenConnector) parseNFT(txResult *rpc.GetTransactionResult, tx *solana.Transaction) *NFT {
	var mintPubKey solana.PublicKey

	// Find only the Token MintTo or MintToChecked instruction
	for _, instruction := range tx.Message.Instructions {
		tokenInstruction, err := connector.DecodeTokenInstruction(instruction, tx)
		if err != nil {
			continue
		}

		switch event := tokenInstruction.Impl.(type) {
		case *token.MintTo: // New NFTs are minted during this MintTo instruction
			mintPubKey = event.GetMintAccount().PublicKey
			if *event.Amount > 1 {
				return nil
			}
			break
		case *token.MintToChecked: // New NFTs are minted during this MintToChecked instruction
			mintPubKey = event.GetMintAccount().PublicKey
			if *event.Amount > 1 {
				return nil
			}
			break
		}
	}

	if mintPubKey.Equals(solana.PublicKey{}) {
		return nil
	}

	ctx := context.Background()
	mint := c.GetMint(ctx, mintPubKey)

	if !isNFT(mint) {
		return nil
	}

	meta, err := c.GetMetadata(ctx, mintPubKey)
	if err != nil {
		log.Warn().Err(err).Msgf("failed to get metadata for mint %s", mintPubKey.String())
		return nil
	}

	return DecodeNFT(mint, meta, txResult)
}

// An NFT in Solana is a Token with only 1 supply and 0 decimals.
func isNFT(mint *token.Mint) bool {
	return mint.Decimals == 0 && mint.Supply == 1
}

var authorityTypeToString = map[bin.BorshEnum]string{
	0: "MintTokens",
	1: "FreezeAccount",
	2: "AccountOwner",
	3: "CloseAccount",
}

// All events are found at https://pkg.go.dev/github.com/gagliardetto/solana-go/programs/token#InstructionImplDef.
func (c *SolanaTokenConnector) InstructionToEventMessage(
	tokenInstruction *token.Instruction,
	txResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
	signature solana.Signature,
) *kafkautils.Message {
	switch event := tokenInstruction.Impl.(type) {
	case *token.Transfer:
		protoMsg := DecodeTransfer(event, txResult, tx, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokentransfer"],
			Key:      kafkautils.NewKey(Namespace+"transfer", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeMint:
		protoMsg := DecodeInitializeMint(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializemint"],
			Key:      kafkautils.NewKey(Namespace+"initializemint", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeAccount:
		protoMsg := DecodeInitializeAccount(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializeaccount"],
			Key:      kafkautils.NewKey(Namespace+"initializeaccount", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.MintTo:
		protoMsg := DecodeMintTo(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenmintto"],
			Key:      kafkautils.NewKey(Namespace, string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.SetAuthority:
		protoMsg := DecodeSetAuthority(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokensetauthority"],
			Key:      kafkautils.NewKey(Namespace+"setauthority", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.CloseAccount:
		protoMsg := DecodeCloseAccount(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokencloseaccount"],
			Key:      kafkautils.NewKey(Namespace+"closeaccount", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.Burn:
		protoMsg := DecodeBurn(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenburn"],
			Key:      kafkautils.NewKey(Namespace+"burn", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.Approve:
		protoMsg := DecodeApprove(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenapprove"],
			Key:      kafkautils.NewKey(Namespace+"approve", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.Revoke:
		protoMsg := DecodeRevoke(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenrevoke"],
			Key:      kafkautils.NewKey(Namespace+"revoke", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.FreezeAccount:
		protoMsg := DecodeFreezeAccount(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenfreezeaccount"],
			Key:      kafkautils.NewKey(Namespace+"freezeaccount", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeMultisig:
		protoMsg := DecodeInitializeMultisig(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializemultisig"],
			Key:      kafkautils.NewKey(Namespace+"initializemultisig", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.SyncNative:
		protoMsg := DecodeSyncNative(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokensyncnative"],
			Key:      kafkautils.NewKey(Namespace+"syncnative", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.ThawAccount:
		protoMsg := DecodeThawAccount(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenthawaccount"],
			Key:      kafkautils.NewKey(Namespace+"thawaccount", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.TransferChecked:
		protoMsg := DecodeTransferChecked(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokentransferchecked"],
			Key:      kafkautils.NewKey(Namespace+"transferchecked", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.ApproveChecked:
		protoMsg := DecodeApproveChecked(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenapprovechecked"],
			Key:      kafkautils.NewKey(Namespace+"approvechecked", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.MintToChecked:
		protoMsg := DecodeMintToChecked(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenminttochecked"],
			Key:      kafkautils.NewKey(Namespace+"minttochecked", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.BurnChecked:
		protoMsg := DecodeBurnChecked(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokenburnchecked"],
			Key:      kafkautils.NewKey(Namespace+"burnchecked", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeAccount2:
		protoMsg := DecodeInitializeAccount2(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializeaccount2"],
			Key:      kafkautils.NewKey(Namespace+"initializeaccount2", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeAccount3:
		protoMsg := DecodeInitializeAccount3(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializeaccount3"],
			Key:      kafkautils.NewKey(Namespace+"initializeaccount3", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeMultisig2:
		protoMsg := DecodeInitializeMultisig2(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializemultisig2"],
			Key:      kafkautils.NewKey(Namespace+"initializemultisig2", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	case *token.InitializeMint2:
		protoMsg := DecodeInitializeMint2(event, txResult, signature)

		return &kafkautils.Message{
			Topic:    c.Topics["tokeninitializemint2"],
			Key:      kafkautils.NewKey(Namespace+"initializemint2", string(protoMsg.Id)),
			ProtoMsg: protoMsg,
		}

	default:
		return nil
	}
}
