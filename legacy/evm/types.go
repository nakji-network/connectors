package evm

import (
	reflect "reflect"

	"blep.ai/data/common"
	"blep.ai/data/database"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.evm.0_0_0.chain_block": &Block{},
	"nakji.evm.0_0_0.chain_tx":    &Transaction{},
}

func UnmarshalEthHeader(in *ethtypes.Header) *Block0 {
	return &Block0{
		Ts:         &timestamp.Timestamp{Seconds: int64(in.Time)},
		Hash:       in.Hash().Bytes(),
		Num:        in.Number.Uint64(),
		Difficulty: in.Difficulty.Uint64(),
		GasLimit:   in.GasLimit,
		GasUsed:    in.GasUsed,
		Nonce:      in.Nonce.Uint64(),
	}
}

// go-ethereum Block to blep struct
func (b *Block) UnmarshalEthBlock(in *ethtypes.Block) {
	*b = Block{
		Ts:         &timestamp.Timestamp{Seconds: int64(in.Time())},
		Hash:       in.Hash().Hex(),
		Difficulty: in.Difficulty().Uint64(),
		Number:     in.Number().Uint64(),
		GasLimit:   in.GasLimit(),
		GasUsed:    in.GasUsed(),
		Nonce:      in.Nonce(),
	}
}

// go-ethereum Transaction to blep struct
func (tx *Transaction) UnmarshalEthTransaction(in *ethtypes.Transaction) error {
	V, R, S := in.RawSignatureValues()

	// handle nil recipients for contract creations
	Recipient := []byte{}
	if in.To() != nil {
		Recipient = in.To().Bytes()
	}

	// Get Sender (.From()) address
	from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(in.ChainId()), in)
	if err != nil {
		log.Warn().Err(err).
			Interface("tx", in).
			Msg("UnmarshallEthTransaction .AsMessage error")
		return err
	}

	*tx = Transaction{
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

	return nil
}

func (_ *Transaction) Viewer() map[reflect.Type]database.View {
	return map[reflect.Type]database.View{
		reflect.TypeOf(common.Kline{}): {
			Src: "timescaledb",
			Cmd: `
				time_bucket('%s', "Ts")  		AS "ts",
				SUM  ("Amount")                 AS "open",
				SUM  ("Amount")                 AS "high",
				SUM  ("Amount")                 AS "low",
				SUM  ("Amount")                 AS "close",
				CAST(COUNT(DISTINCT "Recipient") AS FLOAT)    AS "volume",
				0.0                             AS "buy_volume",
				0.0                             AS "sell_volume"
			`,
		},
	}
}

func (_ *Block) Viewer() map[reflect.Type]database.View {
	return map[reflect.Type]database.View{
		reflect.TypeOf(common.Kline{}): {
			Src: "timescaledb",
			Cmd: `
				time_bucket('%s', "Ts")  		AS "ts",
				0.0                             AS "open",
				0.0                             AS "high",
				0.0                             AS "low",
				SUM  ("Difficulty")             AS "close",
				SUM ("GasUsed")                 AS "volume",
				0.0                             AS "buy_volume",
				0.0                             AS "sell_volume"
			`,
		},
	}
}
