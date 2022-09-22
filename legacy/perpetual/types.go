package perpetual

import (
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L1/PerpRewardVesting"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/Amm"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/AmmReader"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/ChainlinkPriceFeed"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/ClearingHouse"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/InsuranceFund"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/L2PriceFeed"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.perpetual.0_0_0.amm_capchanged":                 &Amm.CapChanged{},
	"nakji.perpetual.0_0_0.amm_fundingrateupdated":         &Amm.FundingRateUpdated{},
	"nakji.perpetual.0_0_0.amm_liquiditychanged":           &Amm.LiquidityChanged{},
	"nakji.perpetual.0_0_0.amm_ownershiptransferred":       &Amm.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.amm_pricefeedupdated":           &Amm.PriceFeedUpdated{},
	"nakji.perpetual.0_0_0.amm_reservesnapshotted":         &Amm.ReserveSnapshotted{},
	"nakji.perpetual.0_0_0.amm_shutdown":                   &Amm.Shutdown{},
	"nakji.perpetual.0_0_0.amm_swapinput":                  &Amm.SwapInput{},
	"nakji.perpetual.0_0_0.amm_swapoutput":                 &Amm.SwapOutput{},
	"nakji.perpetual.0_0_0.ammreader_ownershiptransferred": &AmmReader.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.ammreader_rewardtransferred":    &AmmReader.RewardTransferred{},
	"nakji.perpetual.0_0_0.ammreader_rewardwithdrawn":      &AmmReader.RewardWithdrawn{},
	// "nakji.perpetual.0_0_0.chainlinkl1_adminchanged":                 &ChainlinkL1.AdminChanged{},
	// "nakji.perpetual.0_0_0.chainlinkl1_upgraded":                     &ChainlinkL1.Upgraded{},
	"nakji.perpetual.0_0_0.chainlinkpricefeed_ownershiptransferred":  &ChainlinkPriceFeed.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.clearinghouse_liquidationfeeratiochanged": &ClearingHouse.LiquidationFeeRatioChanged{},
	"nakji.perpetual.0_0_0.clearinghouse_marginchanged":              &ClearingHouse.MarginChanged{},
	"nakji.perpetual.0_0_0.clearinghouse_marginratiochanged":         &ClearingHouse.MarginRatioChanged{},
	"nakji.perpetual.0_0_0.clearinghouse_ownershiptransferred":       &ClearingHouse.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.clearinghouse_paused":                     &ClearingHouse.Paused{},
	"nakji.perpetual.0_0_0.clearinghouse_positionadjusted":           &ClearingHouse.PositionAdjusted{},
	"nakji.perpetual.0_0_0.clearinghouse_positionchanged":            &ClearingHouse.PositionChanged{},
	"nakji.perpetual.0_0_0.clearinghouse_positionliquidated":         &ClearingHouse.PositionLiquidated{},
	"nakji.perpetual.0_0_0.clearinghouse_positionsettled":            &ClearingHouse.PositionSettled{},
	"nakji.perpetual.0_0_0.clearinghouse_referredpositionchanged":    &ClearingHouse.ReferredPositionChanged{},
	"nakji.perpetual.0_0_0.clearinghouse_restrictionmodeentered":     &ClearingHouse.RestrictionModeEntered{},
	"nakji.perpetual.0_0_0.clearinghouse_unpaused":                   &ClearingHouse.Unpaused{},
	// "nakji.perpetual.0_0_0.clientbridge_adminchanged":                &ClientBridge.AdminChanged{},
	// "nakji.perpetual.0_0_0.clientbridge_upgraded":                    &ClientBridge.Upgraded{},
	// "nakji.perpetual.0_0_0.feerewardpooll1_adminchanged":             &FeeRewardPoolL1.AdminChanged{},
	// "nakji.perpetual.0_0_0.feerewardpooll1_upgraded":                 &FeeRewardPoolL1.Upgraded{},
	// "nakji.perpetual.0_0_0.feetokenpooldispatcherl1_adminchanged":    &FeeTokenPoolDispatcherL1.AdminChanged{},
	// "nakji.perpetual.0_0_0.feetokenpooldispatcherl1_upgraded":        &FeeTokenPoolDispatcherL1.Upgraded{},
	"nakji.perpetual.0_0_0.insurancefund_ownershiptransferred": &InsuranceFund.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.insurancefund_shutdownallamms":      &InsuranceFund.ShutdownAllAmms{},
	"nakji.perpetual.0_0_0.insurancefund_tokenadded":           &InsuranceFund.TokenAdded{},
	"nakji.perpetual.0_0_0.insurancefund_tokenremoved":         &InsuranceFund.TokenRemoved{},
	"nakji.perpetual.0_0_0.insurancefund_withdrawn":            &InsuranceFund.Withdrawn{},
	"nakji.perpetual.0_0_0.l2pricefeed_ownershiptransferred":   &L2PriceFeed.OwnershipTransferred{},
	"nakji.perpetual.0_0_0.l2pricefeed_pricefeeddataset":       &L2PriceFeed.PriceFeedDataSet{},
	// "nakji.perpetual.0_0_0.metatxgateway_adminchanged":               &MetaTxGateway.AdminChanged{},
	// "nakji.perpetual.0_0_0.metatxgateway_upgraded":                   &MetaTxGateway.Upgraded{},
	"nakji.perpetual.0_0_0.perprewardvesting_claimed":              &PerpRewardVesting.Claimed{},
	"nakji.perpetual.0_0_0.perprewardvesting_ownershiptransferred": &PerpRewardVesting.OwnershipTransferred{},
	// "nakji.perpetual.0_0_0.rootbridge_adminchanged":                  &RootBridge.AdminChanged{},
	// "nakji.perpetual.0_0_0.rootbridge_upgraded":                      &RootBridge.Upgraded{},
	// "nakji.perpetual.0_0_0.rootbridgev2_adminchanged":                &RootBridgeV2.AdminChanged{},
	// "nakji.perpetual.0_0_0.rootbridgev2_upgraded":                    &RootBridgeV2.Upgraded{},
	// "nakji.perpetual.0_0_0.stakedperptoken_adminchanged":             &StakedPerpToken.AdminChanged{},
	// "nakji.perpetual.0_0_0.stakedperptoken_upgraded":                 &StakedPerpToken.Upgraded{},
}

var ABIs = map[string]string{
	"Amm":       Amm.AmmABI,
	"AmmReader": AmmReader.AmmReaderABI,
	// "ChainlinkL1":              ChainlinkL1.ChainlinkL1ABI,
	"ChainlinkPriceFeed": ChainlinkPriceFeed.ChainlinkPriceFeedABI,
	"ClearingHouse":      ClearingHouse.ClearingHouseABI,
	// "ClientBridge":             ClientBridge.ClientBridgeABI,
	// "FeeRewardPoolL1":          FeeRewardPoolL1.FeeRewardPoolL1ABI,
	// "FeeTokenPoolDispatcherL1": FeeTokenPoolDispatcherL1.FeeTokenPoolDispatcherL1ABI,
	"InsuranceFund": InsuranceFund.InsuranceFundABI,
	"L2PriceFeed":   L2PriceFeed.L2PriceFeedABI,
	// "MetaTxGateway":            MetaTxGateway.MetaTxGatewayABI,
	"PerpRewardVesting": PerpRewardVesting.PerpRewardVestingABI,
	// "RootBridge":               RootBridge.RootBridgeABI,
	// "RootBridgeV2":             RootBridgeV2.RootBridgeV2ABI,
	// "StakedPerpToken":          StakedPerpToken.StakedPerpTokenABI,
}
