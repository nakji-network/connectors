package woofi

import (
	"github.com/nakji-network/connectors/woofi/bsc_WOOPP"
	"github.com/nakji-network/connectors/woofi/polygon_WOOPP"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]map[string]proto.Message{
	"bsc": {
		"nakji.woofi.0_0_0.bsc_woopp_feemanagerupdated":         &bsc_WOOPP.FeeManagerUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_ownershiptransferprepared": &bsc_WOOPP.OwnershipTransferPrepared{},
		"nakji.woofi.0_0_0.bsc_woopp_ownershiptransferred":      &bsc_WOOPP.OwnershipTransferred{},
		"nakji.woofi.0_0_0.bsc_woopp_parametersupdated":         &bsc_WOOPP.ParametersUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_paused":                    &bsc_WOOPP.Paused{},
		"nakji.woofi.0_0_0.bsc_woopp_rewardmanagerupdated":      &bsc_WOOPP.RewardManagerUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_strategistupdated":         &bsc_WOOPP.StrategistUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_unpaused":                  &bsc_WOOPP.Unpaused{},
		"nakji.woofi.0_0_0.bsc_woopp_withdraw":                  &bsc_WOOPP.Withdraw{},
		"nakji.woofi.0_0_0.bsc_woopp_wooguardianupdated":        &bsc_WOOPP.WooGuardianUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_wooracleupdated":           &bsc_WOOPP.WooracleUpdated{},
		"nakji.woofi.0_0_0.bsc_woopp_wooswap":                   &bsc_WOOPP.WooSwap{},
	},
	"polygon": {
		"nakji.woofi.0_0_0.polygon_woopp_feemanagerupdated":         &polygon_WOOPP.FeeManagerUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_ownershiptransferprepared": &polygon_WOOPP.OwnershipTransferPrepared{},
		"nakji.woofi.0_0_0.polygon_woopp_ownershiptransferred":      &polygon_WOOPP.OwnershipTransferred{},
		"nakji.woofi.0_0_0.polygon_woopp_parametersupdated":         &polygon_WOOPP.ParametersUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_paused":                    &polygon_WOOPP.Paused{},
		"nakji.woofi.0_0_0.polygon_woopp_rewardmanagerupdated":      &polygon_WOOPP.RewardManagerUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_strategistupdated":         &polygon_WOOPP.StrategistUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_unpaused":                  &polygon_WOOPP.Unpaused{},
		"nakji.woofi.0_0_0.polygon_woopp_withdraw":                  &polygon_WOOPP.Withdraw{},
		"nakji.woofi.0_0_0.polygon_woopp_wooguardianupdated":        &polygon_WOOPP.WooGuardianUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_wooracleupdated":           &polygon_WOOPP.WooracleUpdated{},
		"nakji.woofi.0_0_0.polygon_woopp_wooswap":                   &polygon_WOOPP.WooSwap{},
	},
}

var ABIs = map[string]string{
	"bsc_WOOPP":     bsc_WOOPP.BscWOOPPABI,
	"polygon_WOOPP": polygon_WOOPP.PolygonWOOPPABI,
}
