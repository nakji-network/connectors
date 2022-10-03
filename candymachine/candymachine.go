package candymachine

import (
	"context"
	"os"
	"os/signal"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain"
	"github.com/nakji-network/connector/kafkautils"

	candymachinev2 "github.com/gagliardetto/metaplex-go/clients/nft-candy-machine-v2"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	Namespace = "candymachine"
)

type Connector struct {
	*connector.Connector
	CandyMachinePublicKey solana.PublicKey
	client                *chain.SolanaClient
	config                *viper.Viper
}

func NewConnector(conf *viper.Viper) *Connector {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	programID := conf.GetString("candymachine.programid")

	client := c.ChainClients.Solana()

	return &Connector{
		CandyMachinePublicKey: solana.MustPublicKeyFromBase58(programID),
		Connector:             c,
		config:                conf,
		client:                client,
	}
}

func (c *Connector) Start() {
	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register topic and protobuf type mappings
	protos := make([]protoreflect.ProtoMessage, len(TopicTypes))
	i := 0
	for _, topicProto := range TopicTypes {
		protos[i] = topicProto
		i++
	}

	c.RegisterProtos(kafkautils.MsgTypeFct, protos...)

	//backfillLimit := c.config.GetInt("candymachine.backfillLimit")

	go c.client.ProgramListener(ctx, c.CandyMachinePublicKey, Namespace, c.processListenerTransaction)

	// TODO: Need further refactoring because the solana wsclient is not being used.
	//go c.client.ProgramBackfiller(ctx, c.CandyMachinePublicKey, Namespace, c.processBackfillerTransaction, backfillLimit)

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			c.client.Close()
			return
		}
	}
}

func (c *Connector) processListenerTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) {
	instructions := c.getInstructions(txResult, tx)

	for _, instruction := range instructions {
		msg := c.ConvertInstructionToMessage(instruction, txResult, tx, signature)

		if msg != nil {
			log.Info().Interface("msg", msg).Msg("candy machine message published")
			c.EventSink <- &kafkautils.Message{
				MsgType:  kafkautils.MsgTypeFct,
				ProtoMsg: msg,
			}
		}
	}
}

func (c *Connector) processBackfillerTransaction(txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) {
	instructions := c.getInstructions(txResult, tx)

	for _, instruction := range instructions {
		msg := c.ConvertInstructionToMessage(instruction, txResult, tx, signature)

		if msg != nil {
			log.Info().Interface("msg", msg).Msg("candy machine message published via backfill")
			c.EventSink <- &kafkautils.Message{
				MsgType:  kafkautils.MsgTypeFct,
				ProtoMsg: msg,
			}
		}
	}
}

func (c *Connector) ConvertInstructionToMessage(instruction *candymachinev2.Instruction, txResult *rpc.GetTransactionResult, _ *solana.Transaction, signature solana.Signature) proto.Message {
	var protoMsg proto.Message

	switch impl := instruction.Impl.(type) {
	case *candymachinev2.MintNft:
		protoMsg = DecodeMintNFT(impl, txResult, signature)
	case *candymachinev2.UpdateCandyMachine:
		protoMsg = DecodeUpdateCandyMachine(impl, txResult, signature)
	case *candymachinev2.AddConfigLines:
		protoMsg = DecodeAddConfigLines(impl, txResult, signature)
	case *candymachinev2.InitializeCandyMachine:
		protoMsg = DecodeInitializeCandyMachine(impl, txResult, signature)
	case *candymachinev2.UpdateAuthority:
		protoMsg = DecodeUpdateAuthority(impl, txResult, signature)
	case *candymachinev2.WithdrawFunds:
		protoMsg = DecodeWithdrawFunds(impl, txResult, signature)
	default:
		return nil
	}

	return protoMsg
}

// Fetches and parses all CandyMachine V2 instructions and inner instructions from all transaction data.
// The returned slice preserves the ordering of these Token Program instructions in the transaction.
func (c *Connector) getInstructions(txResult *rpc.GetTransactionResult, tx *solana.Transaction) []*candymachinev2.Instruction {
	accountKeys := tx.Message.AccountKeys
	var (
		programIdIdx uint16
		instructions []*candymachinev2.Instruction
	)

	for i, publicKey := range accountKeys {
		if publicKey.String() == c.CandyMachinePublicKey.String() {
			programIdIdx = uint16(i)
			break
		}
	}

	// This map preserves the ordering of instructions and inner instructions within a transaction.
	// This is important since Solana uses that as basis for the order of their execution. The map's
	// key is the parent instruction's index within the transaction.
	innerInstructionsIndexMap := make(map[uint16]rpc.InnerInstruction)
	for _, innerInstruction := range txResult.Meta.InnerInstructions {
		innerInstructionsIndexMap[innerInstruction.Index] = innerInstruction
	}

	for i, instruction := range tx.Message.Instructions {
		// Check if the parent instruction is an SPL-Token Program instruction.
		// Add that to `instructions` slice, if it is.
		if instruction.ProgramIDIndex == programIdIdx {
			ins, err := decodeInstruction(instruction, tx)
			if err == nil {
				instructions = append(instructions, ins)
			}
		}

		// If the inner instruction's index matches the current index, we check if any of these
		// inner instructions are CandyMachine V2 instructions and add them to `instructions`.
		if innerInstruction, ok := innerInstructionsIndexMap[uint16(i)]; ok {
			for _, instruction := range innerInstruction.Instructions {
				if instruction.ProgramIDIndex == programIdIdx {
					ins, err := decodeInstruction(instruction, tx)
					if err == nil {
						instructions = append(instructions, ins)
					}
				}
			}
		}
	}

	return instructions
}

// Decodes a `solana.CompiledInstruction` into a `*candymachinev2.Instruction` which represents an Instruction
// from the CandyMachine V2 Program in Solana.
func decodeInstruction(instruction solana.CompiledInstruction, tx *solana.Transaction) (*candymachinev2.Instruction, error) {
	accountMetas := instruction.ResolveInstructionAccounts(&tx.Message)
	ins, err := candymachinev2.DecodeInstruction(accountMetas, instruction.Data)
	if err != nil {
		log.Error().Err(err).Msg("decoding instruction failed")
		return nil, err
	}
	return ins, nil
}
