package starkex

import (
	"blep.ai/data/common"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.starkex.0_0_0.starkex_position": &Position{},
}

func (p *Position) UnmarshalPosition(vault Vault, assetKey string, asset Asset) {
	*p = Position{
		Ts:                 common.UnixToTimestampPb(1000 * int64(vault.Ts)),
		BlockNumber:        uint64(vault.BlockNumber),
		Asset:              assetKey,
		AssetType:          asset.Type,
		Amount:             asset.Amount,
		CachedFundingIndex: asset.AdditionalAttributes.CachedFundingIndex,
		VaultKey:           vault.Key,
	}
}
