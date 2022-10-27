package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi"
	"github.com/nakji-network/connectors/woofi/WOOPP"

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
		&WOOPP.WooSwap{})

	conf := &woofi.Config{
		NetworkName: "polygon",
		ContractAddresses: map[string]string{
			"WOOPP": "0x7400B665C8f4f3a951a99f1ee9872efb8778723d",
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
