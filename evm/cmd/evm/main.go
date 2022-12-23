package main

import (
	"fmt"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/evm"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	pflag.StringP("network", "n", "ethereum", "network to connect to e.g. ethereum")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")
	pflag.StringP("author", "a", "nakji", "author of the connector (optional)")
	pflag.Parse()

	cfg := config.GetConfig()
	cfg.BindPFlags(pflag.CommandLine)

	if err := validateFlags(cfg); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	networkName := cfg.GetString("network")
	authorName := cfg.GetString("author")
	m := connector.NewManifest(networkName, authorName, "0.0.0")

	c, err := connector.NewConnector(connector.WithManifest(m))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	conf := &evm.Config{
		NetworkName: cfg.GetString("network"),
		FromBlock:   cfg.GetUint64("from-block"),
		NumBlocks:   cfg.GetUint64("num-blocks"),
	}

	connector := evm.New(c, conf)
	connector.Start()
}

func validateFlags(conf config.IConfig) error {
	networkName := conf.GetString("network")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableNetworks := map[string]bool{
		"arbitrum":  true,
		"avalanche": true,
		"bsc":       true,
		"ethereum":  true,
		"fantom":    true,
		"harmony":   true,
		"heco":      true,
		"klaytn":    true,
		"metis":     true,
		"moonbeam":  true,
		"moonriver": true,
		"okc":       true,
		"optimism":  true,
		"polygon":   true,
		"gnosis":    true,
	}

	if _, isExists := availableNetworks[networkName]; !isExists {
		errorMsg := "network is not supported, please try again with one of these: "
		for k := range availableNetworks {
			errorMsg += k + " "
		}
		return fmt.Errorf(errorMsg)
	}

	if fromBlock < 0 {
		return fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return nil
}
