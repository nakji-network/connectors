package klaytn

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/rs/zerolog/log"
)

func UnmarshalEthHeader(in *types.Header) *Block0 {
	return &Block0{
		Ts:         &timestamp.Timestamp{Seconds: in.Time.Int64()},
		Num:        in.Number.Uint64(),
		Hash:       in.Hash().Bytes(),
		BlockScore: in.BlockScore.Uint64(),
		GasUsed:    in.GasUsed,
	}
}

// go-klaytn Block to nakji struct
func (x *Block) UnmarshalKlaytnBlock(in *types.Block) {
	*x = Block{
		Ts:         &timestamp.Timestamp{Seconds: in.Time().Int64()},
		Hash:       in.Hash().Hex(),
		BlockScore: in.BlockScore().Uint64(),
		Number:     in.Number().Uint64(),
		GasUsed:    in.GasUsed(),
	}
}

// go-klaytn Transaction to nakji struct
func (tx *Transaction) UnmarshalKlaytnTransaction(in *types.Transaction) {

	// handle nil recipients for contract creations
	Recipient := []byte{}
	if in.To() != nil {
		Recipient = in.To().Bytes()
	}

	// Get Sender (.From()) address
	from, err := types.Sender(types.LatestSignerForChainID(in.ChainId()), in)
	if err != nil {
		log.Fatal().Err(err).
			Interface("tx", in).
			Msg("UnmarshallEthTransaction .AsMessage error")
	}

	var sigs []*Signature
	for _, txSig := range in.RawSignatureValues() {
		sigs = append(sigs, &Signature{
			V: txSig.V.Uint64(),
			R: txSig.R.Uint64(),
			S: txSig.S.Uint64(),
		})
	}

	*tx = Transaction{
		AccountNonce: in.Nonce(),
		Price:        in.GasPrice().Uint64(),
		GasLimit:     in.Gas(),
		Recipient:    Recipient, // nil means contract creation
		Amount:       in.Value().Uint64(),
		Payload:      in.Data(),

		From:       from.Bytes(),
		Hash:       in.Hash().Hex(),
		Size:       float64(in.Size()),
		Signatures: sigs,
	}
}
