package woofi

import (
	"github.com/nakji-network/connectors/woofi/WOOPP"
	"google.golang.org/protobuf/proto"
)

const (
	WOOPPContractAddr = "0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"
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
}
