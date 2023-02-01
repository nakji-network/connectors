package main

import (
	"fmt"

	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/okc"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	conf := &okc.Config{
		BlockchainName: "okc",
		FromBlock:      cfg.GetUint64("from-block"),
		NumBlocks:      cfg.GetUint64("num-blocks"),
	}

	connector := okc.New(conf)
	connector.Start()
}

func parseConfig() (config.IConfig, error) {
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")
	pflag.StringP("author", "a", "nakji", "author of the connector (optional)")
	pflag.Parse()

	conf := config.GetConfig()
	conf.BindPFlags(pflag.CommandLine)

	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	if fromBlock < 0 {
		return nil, fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return nil, fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return conf, nil
}
