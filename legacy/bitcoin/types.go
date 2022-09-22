package bitcoin

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.bitcoin.0_0_0.chain_block": &Block{},
	"nakji.bitcoin.0_0_0.chain_tx":    &Transaction{},
}

// Convert from btcd MsgBlock to Blep's Bitcoin Block
func (b *Block) UnmarshalBTCBlock(in *btcjson.GetBlockVerboseResult) {
	*b = Block{
		MerkleRoot:    in.MerkleRoot,
		Bits:          in.Bits,
		Hash:          in.Hash,
		Difficulty:    in.Difficulty,
		Nonce:         in.Nonce,
		PreviousHash:  in.PreviousHash,
		NextHash:      in.NextHash,
		Ts:            &timestamp.Timestamp{Seconds: in.Time},
		Version:       in.Version,
		VersionHex:    in.VersionHex,
		Height:        in.Height,
		Weight:        in.Weight,
		Size:          in.Size,
		StrippedSize:  in.StrippedSize,
		Confirmations: in.Confirmations,
	}
}

func (tx *Transaction) UnmarshalBTCTransaction(in *btcjson.TxRawResult) {
	vins := make([]*Vin, len(in.Vin))
	for _, vin := range in.Vin {
		var asm, hex string
		if vin.ScriptSig != nil {
			asm = vin.ScriptSig.Asm
			hex = vin.ScriptSig.Hex
		}

		newVin := Vin{
			Coinbase: vin.Coinbase,
			Txid:     vin.Txid,
			Vout:     vin.Vout,
			Asm:      asm,
			Hex:      hex,
			Sequence: vin.Sequence,
			Witness:  vin.Witness,
		}
		vins = append(vins, &newVin)
	}

	vouts := make([]*Vout, len(in.Vout))
	for _, vout := range in.Vout {
		newVout := Vout{
			Value:     vout.Value,
			N:         vout.N,
			Asm:       vout.ScriptPubKey.Asm,
			Hex:       vout.ScriptPubKey.Hex,
			ReqSigs:   vout.ScriptPubKey.ReqSigs,
			Type:      vout.ScriptPubKey.Type,
			Addresses: vout.ScriptPubKey.Addresses,
		}
		vouts = append(vouts, &newVout)
	}

	*tx = Transaction{
		Hex:           in.Hex,
		Txid:          in.Txid,
		Hash:          in.Hash,
		Size:          in.Size,
		Vsize:         in.Vsize,
		Weight:        in.Weight,
		Version:       in.Version,
		LockTime:      in.LockTime,
		BlockHash:     in.BlockHash,
		Confirmations: in.Confirmations,
		Ts:            &timestamp.Timestamp{Seconds: in.Time},
		Blocktime:     in.Blocktime,
		Vin:           vins,
		Vout:          vouts,
	}
}
