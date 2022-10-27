package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi"
	"github.com/nakji-network/connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/connectors/woofi/WooPP"
	"github.com/nakji-network/connectors/woofi/WooPPV1"
	"github.com/nakji-network/connectors/woofi/WooPPV2"
	"github.com/nakji-network/connectors/woofi/WooRouterV1"
	"github.com/nakji-network/connectors/woofi/WooRouterV2"
	"github.com/nakji-network/connectors/woofi/WooRouterV3"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// TODO: remove backfill from main binary
	if err := processBackfillFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	c.RegisterProtos(
		&WooPP.FeeManagerUpdated{},
		&WooPP.OwnershipTransferPrepared{},
		&WooPP.OwnershipTransferred{},
		&WooPP.ParametersUpdated{},
		&WooPP.Paused{},
		&WooPP.RewardManagerUpdated{},
		&WooPP.StrategistUpdated{},
		&WooPP.Unpaused{},
		&WooPP.Withdraw{},
		&WooPP.WooGuardianUpdated{},
		&WooPP.WooracleUpdated{},
		&WooPP.WooSwap{},

		&WooPP.FeeManagerUpdated{},
		&WooPP.OwnershipTransferPrepared{},
		&WooPP.OwnershipTransferred{},
		&WooPP.ParametersUpdated{},
		&WooPP.Paused{},
		&WooPP.RewardManagerUpdated{},
		&WooPP.StrategistUpdated{},
		&WooPP.Unpaused{},
		&WooPP.Withdraw{},
		&WooPP.WooGuardianUpdated{},
		&WooPP.WooracleUpdated{},
		&WooPP.WooSwap{},

		&WooPPV1.OwnershipTransferPrepared{},
		&WooPPV1.OwnershipTransferred{},
		&WooPPV1.ParametersUpdated{},
		&WooPPV1.Paused{},
		&WooPPV1.RewardManagerUpdated{},
		&WooPPV1.StrategistUpdated{},
		&WooPPV1.Unpaused{},
		&WooPPV1.Withdraw{},
		&WooPPV1.WooGuardianUpdated{},
		&WooPPV1.WooracleUpdated{},
		&WooPPV1.WooSwap{},

		&WooPPV2.OwnershipTransferPrepared{},
		&WooPPV2.OwnershipTransferred{},
		&WooPPV2.ParametersUpdated{},
		&WooPPV2.Paused{},
		&WooPPV2.RewardManagerUpdated{},
		&WooPPV2.StrategistUpdated{},
		&WooPPV2.Unpaused{},
		&WooPPV2.Withdraw{},
		&WooPPV2.WooGuardianUpdated{},
		&WooPPV2.WooracleUpdated{},
		&WooPPV2.WooSwap{},

		&WooRouterV1.OwnershipTransferred{},
		&WooRouterV1.WooPoolChanged{},
		&WooRouterV1.WooRouterSwap{},

		&WooRouterV2.OwnershipTransferred{},
		&WooRouterV2.WooPoolChanged{},
		&WooRouterV2.WooRouterSwap{},

		&WooRouterV3.OwnershipTransferred{},
		&WooRouterV3.WooPoolChanged{},
		&WooRouterV3.WooRouterSwap{},

		&WooCrossChainRouterV1.WooCrossSwapOnSrcChain{},
		&WooCrossChainRouterV1.WooCrossSwapOnDstChain{},
		&WooCrossChainRouterV1.OwnershipTransferred{},
	)

	conf := &woofi.Config{
		NetworkName: "avalanche",
		ContractAddresses: map[string]string{
			"WooPPV1":               "0xF8cE0D043891b62c55380fB1EFBfB4F186153D96",
			"WooPPV2":               "0x1df3009c57a8B143c6246149F00B090Bce3b8f88",
			"WooRouterV1":           "0x160020B09DeD3d862f7f851B5c50632BcF2062FF",
			"WooRouterV2":           "0x3e0Da0A9e4139B32b37710784b8DCa643c152001",
			"WooRouterV3":           "0x5AA6a4E96A9129562e2fc06660D07FEdDAAf7854",
			"WooCrossChainRouterV1": "0xdF37F7A85D4563f39A78494568824b4dF8669B7a",
		},
		FromBlock: c.Config.GetUint64("from-block"),
		NumBlocks: c.Config.GetUint64("num-blocks"),
	}

	m := woofi.New(c, conf)
	m.Start()
}

// Deprecated: backfill should be moved outside the main process
func processBackfillFlags(conf *viper.Viper) error {
	conf.SetDefault("blockTime", 15*time.Second)
	conf.SetDefault("waitBlocks", 4)

	//	Woo initial block is #18675185
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)

	return validateFlags(conf)
}

// Deprecated: backfill should be moved outside the main process
func validateFlags(conf config.IConfig) error {
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	if fromBlock < 0 {
		return fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return nil
}
