package compound

import (
	. "blep.ai/data/connectors/source/compound/ctoken"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.compound.0_0_0.compound_mint":            &Mint{},
	"nakji.compound.0_0_0.compound_redeem":          &Redeem{},
	"nakji.compound.0_0_0.compound_borrow":          &Borrow{},
	"nakji.compound.0_0_0.compound_repayborrow":     &RepayBorrow{},
	"nakji.compound.0_0_0.compound_liquidateborrow": &LiquidateBorrow{},
}
