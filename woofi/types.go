package woofi

import (
	"github.com/nakji-network/connectors/woofi/bscWOOPP"
	"github.com/nakji-network/connectors/woofi/polygonWOOPP"

	"google.golang.org/protobuf/proto"
)

const (
	BscNetwork           = "bsc"
	BscWOOPPContractAddr = "0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"

	PolygonNetwork           = "polygon"
	PolygonWOOPPContractAddr = "0x7400B665C8f4f3a951a99f1ee9872efb8778723d"
)

var TopicTypes = []proto.Message{
	&bscWOOPP.FeeManagerUpdated{},
	&bscWOOPP.OwnershipTransferPrepared{},
	&bscWOOPP.OwnershipTransferred{},
	&bscWOOPP.ParametersUpdated{},
	&bscWOOPP.Paused{},
	&bscWOOPP.RewardManagerUpdated{},
	&bscWOOPP.StrategistUpdated{},
	&bscWOOPP.Unpaused{},
	&bscWOOPP.Withdraw{},
	&bscWOOPP.WooGuardianUpdated{},
	&bscWOOPP.WooracleUpdated{},
	&bscWOOPP.WooSwap{},

	&polygonWOOPP.FeeManagerUpdated{},
	&polygonWOOPP.OwnershipTransferPrepared{},
	&polygonWOOPP.OwnershipTransferred{},
	&polygonWOOPP.ParametersUpdated{},
	&polygonWOOPP.Paused{},
	&polygonWOOPP.RewardManagerUpdated{},
	&polygonWOOPP.StrategistUpdated{},
	&polygonWOOPP.Unpaused{},
	&polygonWOOPP.Withdraw{},
	&polygonWOOPP.WooGuardianUpdated{},
	&polygonWOOPP.WooracleUpdated{},
	&polygonWOOPP.WooSwap{},
}
