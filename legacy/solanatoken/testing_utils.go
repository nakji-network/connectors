package solanatoken

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"blep.ai/data/chain/solana/connector"
	. "blep.ai/data/chain/solana/testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spf13/viper"
)

type testInstruction interface {
	solana.Instruction
	InnerInstructions() []solana.Instruction
}

type testTxInstruction struct {
	accounts          []*solana.AccountMeta
	data              []byte
	programID         solana.PublicKey
	innerInstructions []solana.Instruction
}

func (t *testTxInstruction) Accounts() []*solana.AccountMeta {
	return t.accounts
}

func (t *testTxInstruction) ProgramID() solana.PublicKey {
	return t.programID
}

func (t *testTxInstruction) Data() ([]byte, error) {
	return t.data, nil
}

func (t *testTxInstruction) InnerInstructions() []solana.Instruction {
	return t.innerInstructions
}

var (
	testSignature = solana.MustSignatureFromBase58("5rPNRaUSNJeEoQdNDC7Y611pA3MpHtFvXwgrpBnF8Mv5hb8BLXGAiZcnrZd36R5BrjBmN66afZVfvWiB3o3Fc9X7")
	testTs        = solana.UnixTimeSeconds(1000)
)

func NewTxResult(instructions []testInstruction) (*rpc.GetTransactionResult, *solana.Transaction) {
	solInsts := make([]solana.Instruction, len(instructions))
	for i, inst := range instructions {
		solInsts[i] = inst.(solana.Instruction)
	}
	tx, err := solana.NewTransaction(solInsts, solana.MustHashFromBase58("A9QnpgfhCkmiBSjgBuWk76Wo3HxzxvDopUq9x6UUMmjn"))
	if err != nil {
		panic(err)
	}

	innerInstructions := make([]rpc.InnerInstruction, 0)
	for i, instruction := range instructions {
		if len(instruction.InnerInstructions()) == 0 {
			continue
		}

		compiledInsts := make([]solana.CompiledInstruction, len(instruction.InnerInstructions()))

		for i, instruction := range instruction.InnerInstructions() {
			data, err := instruction.Data()
			if err != nil {
				panic(err)
			}

			accountKeyIndex := map[string]uint16{}
			for i, accountKey := range tx.Message.AccountKeys {
				accountKeyIndex[accountKey.String()] = uint16(i)
			}

			accounts := instruction.Accounts()
			accountIndex := make([]uint16, len(accounts))
			for i, account := range accounts {
				accountIndex[i] = accountKeyIndex[account.PublicKey.String()]
			}

			compiledInsts[i] = solana.CompiledInstruction{
				ProgramIDIndex: accountKeyIndex[instruction.ProgramID().String()],
				AccountCount:   bin.Varuint16(len(instruction.Accounts())),
				DataLength:     bin.Varuint16(uint16(len(data))),
				Data:           data,
				Accounts:       accountIndex,
			}
		}

		innerInstructions = append(innerInstructions, rpc.InnerInstruction{
			Index:        uint16(i),
			Instructions: compiledInsts,
		})
	}

	meta := &rpc.TransactionMeta{
		InnerInstructions: innerInstructions,
	}
	txResult := &rpc.GetTransactionResult{
		BlockTime: &testTs,
		Meta:      meta,
	}

	tx.Signatures = []solana.Signature{testSignature}
	return txResult, tx
}

func testConnector(client *rpc.Client) *SolanaTokenConnector {
	conf := viper.New()
	conf.SetDefault("solana.maxRetries", 10)
	conf.SetDefault("solana.lowPrioRetryDelay", 100)
	conf.SetDefault("solana.lowPrioMaxRetries", 1)
	conf.SetDefault("solana.retryDelay", 100)
	connector := &connector.SolanaConnector{Client: client, Conf: conf}

	return &SolanaTokenConnector{
		SolanaConnector: connector,
	}
}

func wrapAccountResponse(accountData interface{}, encoder func(io.Writer) *bin.Encoder) json.RawMessage {
	buf := new(bytes.Buffer)
	encoder(buf).Encode(accountData)
	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	response := `{"context":{"slot":83986105},"value":{"data":["%s","base64"],"executable":true,"lamports":999999,"owner":"11111111111111111111111111111111","rentEpoch":207}}`
	response = fmt.Sprintf(response, str)
	return json.RawMessage(WrapIntoRPC(response))
}
