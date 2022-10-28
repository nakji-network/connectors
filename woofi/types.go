package woofi

import (
	"github.com/nakji-network/connectors/woofi/WOOPP"
	"github.com/nakji-network/connectors/woofi/bscWooPP"
	"github.com/nakji-network/connectors/woofi/bscWooPPV1"
	"github.com/nakji-network/connectors/woofi/bscWooPPV2"
	"github.com/nakji-network/connectors/woofi/bscWooRouterV1"
	"github.com/nakji-network/connectors/woofi/bscWooRouterV2"
	"github.com/nakji-network/connectors/woofi/polygonWooPP"

	"google.golang.org/protobuf/proto"
)

const (
	BscNetwork                 = "bsc"
	BscWooPPContractAddr       = "0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"
	BscWooPPV1ContractAddr     = "0x10C24658815585851a8888f059Cb4338790023F1"
	BscWooPPV2ContractAddr     = "0x8489d142Da126F4Ea01750e80ccAa12FD1642988"
	BscWooRouterV1ContractAddr = "0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83"
	BscWooRouterV2ContractAddr = "0xcEf5BE73ae943B77f9Bc08859367D923C030a269"

	PolygonNetwork           = "polygon"
	PolygonWooPPContractAddr = "0x7400B665C8f4f3a951a99f1ee9872efb8778723d"
)

var TopicTypes = []proto.Message{
	&WOOPP.FeeManagerUpdated{},
	&WOOPP.OwnershipTransferPrepared{},
	&WOOPP.OwnershipTransferred{},
	&WOOPP.ParametersUpdated{},
	&WOOPP.Paused{},
	&WOOPP.RewardManagerUpdated{},
	&WOOPP.StrategistUpdated{},
	&WOOPP.Unpaused{},
	&WOOPP.Withdraw{},
	&WOOPP.WooGuardianUpdated{},
	&WOOPP.WooracleUpdated{},
	&WOOPP.WooSwap{},

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

	&bscWooPPV1.OwnershipTransferPrepared{},
	&bscWooPPV1.OwnershipTransferred{},
	&bscWooPPV1.ParametersUpdated{},
	&bscWooPPV1.Paused{},
	&bscWooPPV1.RewardManagerUpdated{},
	&bscWooPPV1.StrategistUpdated{},
	&bscWooPPV1.Unpaused{},
	&bscWooPPV1.Withdraw{},
	&bscWooPPV1.WooGuardianUpdated{},
	&bscWooPPV1.WooracleUpdated{},
	&bscWooPPV1.WooSwap{},

	&bscWooPPV2.OwnershipTransferPrepared{},
	&bscWooPPV2.OwnershipTransferred{},
	&bscWooPPV2.ParametersUpdated{},
	&bscWooPPV2.Paused{},
	&bscWooPPV2.RewardManagerUpdated{},
	&bscWooPPV2.StrategistUpdated{},
	&bscWooPPV2.Unpaused{},
	&bscWooPPV2.Withdraw{},
	&bscWooPPV2.WooGuardianUpdated{},
	&bscWooPPV2.WooracleUpdated{},
	&bscWooPPV2.WooSwap{},

	&bscWooRouterV1.OwnershipTransferred{},
	&bscWooRouterV1.WooPoolChanged{},
	&bscWooRouterV1.WooRouterSwap{},

	&bscWooRouterV2.OwnershipTransferred{},
	&bscWooRouterV2.WooPoolChanged{},
	&bscWooRouterV2.WooRouterSwap{},

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
