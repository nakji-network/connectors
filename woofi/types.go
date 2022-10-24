package woofi

import (
	"github.com/nakji-network/connectors/woofi/bscWooPP"
	"github.com/nakji-network/connectors/woofi/bscWooRouter"
	"github.com/nakji-network/connectors/woofi/polygonWooPP"

	"google.golang.org/protobuf/proto"
)

const (
	BscNetwork               = "bsc"
	BscWooPPContractAddr     = "0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"
	BscWooRouterContractAddr = "0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83"

	PolygonNetwork           = "polygon"
	PolygonWooPPContractAddr = "0x7400B665C8f4f3a951a99f1ee9872efb8778723d"
)

var TopicTypes = []proto.Message{
	&bscWooPP.FeeManagerUpdated{},
	&bscWooPP.OwnershipTransferPrepared{},
	&bscWooPP.OwnershipTransferred{},
	&bscWooPP.ParametersUpdated{},
	&bscWooPP.Paused{},
	&bscWooPP.RewardManagerUpdated{},
	&bscWooPP.StrategistUpdated{},
	&bscWooPP.Unpaused{},
	&bscWooPP.Withdraw{},
	&bscWooPP.WooGuardianUpdated{},
	&bscWooPP.WooracleUpdated{},
	&bscWooPP.WooSwap{},

	&bscWooRouter.OwnershipTransferred{},
	&bscWooRouter.WooPoolChanged{},
	&bscWooRouter.WooRouterSwap{},

	&polygonWooPP.FeeManagerUpdated{},
	&polygonWooPP.OwnershipTransferPrepared{},
	&polygonWooPP.OwnershipTransferred{},
	&polygonWooPP.ParametersUpdated{},
	&polygonWooPP.Paused{},
	&polygonWooPP.RewardManagerUpdated{},
	&polygonWooPP.StrategistUpdated{},
	&polygonWooPP.Unpaused{},
	&polygonWooPP.Withdraw{},
	&polygonWooPP.WooGuardianUpdated{},
	&polygonWooPP.WooracleUpdated{},
	&polygonWooPP.WooSwap{},
}
