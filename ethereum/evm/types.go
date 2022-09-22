package evm

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
)

func UnmarshalEthHeader(in *types.Header) *Block {
	return &Block{
		Ts:         &timestamp.Timestamp{Seconds: int64(in.Time)},
		Hash:       in.Hash().Bytes(),
		Num:        in.Number.Uint64(),
		Difficulty: in.Difficulty.Uint64(),
		GasLimit:   in.GasLimit,
		GasUsed:    in.GasUsed,
		Nonce:      in.Nonce.Uint64(),
	}
}

// go-ethereum Block to nakji struct
//func (b *Block) UnmarshalEthBlock(in *types.Block) {
//*b = Block{
//Timestamp:  &timestamp.Timestamp{Seconds: int64(in.Time())},
//Hash:       in.Hash().Hex(),
//Difficulty: in.Difficulty().Uint64(),
//Number:     in.Number().Uint64(),
//GasLimit:   in.GasLimit(),
//GasUsed:    in.GasUsed(),
//Nonce:      in.Nonce(),
//}
//}

// go-ethereum Transaction to nakji struct
func UnmarshalEthTransaction(in *types.Transaction) *Transaction {
	V, R, S := in.RawSignatureValues()

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

	return &Transaction{
		AccountNonce: in.Nonce(),
		Price:        in.GasPrice().Uint64(),
		GasLimit:     in.Gas(),
		Recipient:    Recipient, // nil means contract creation
		Amount:       in.Value().Uint64(),
		Payload:      in.Data(),

		From: from.Bytes(),
		Hash: in.Hash().Hex(),
		Size: float64(in.Size()),

		V: V.Uint64(),
		R: R.Uint64(),
		S: S.Uint64(),
	}
}
