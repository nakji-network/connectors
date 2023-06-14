package main

import (
	"fmt"

	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/mai"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	conf := &mai.Config{
		Blockchain: cfg.GetString("blockchain"),
		FromBlock:  cfg.GetUint64("from-block"),
		NumBlocks:  cfg.GetUint64("num-blocks"),
	}

	connector := mai.New(conf)
	connector.Start()
}

func parseConfig() (config.IConfig, error) {
	pflag.StringP("blockchain", "c", "polygon", "blockchain to connect to e.g. ethereum")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")
	pflag.StringP("author", "a", "nakji", "author of the connector (optional)")
	pflag.Parse()

	conf := config.GetConfig()
	conf.BindPFlags(pflag.CommandLine)

	blockchain := conf.GetString("blockchain")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableBlockchains := map[string]bool{
		"polygon": true,
	}

	if _, isExists := availableBlockchains[blockchain]; !isExists {
		errorMsg := "blockchain is not supported, please try again with one of these: "
		for k := range availableBlockchains {
			errorMsg += k + " "
		}
		return nil, fmt.Errorf(errorMsg)
	}

	if fromBlock < 0 {
		return nil, fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return nil, fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return conf, nil
}
