package main

import (
	"fmt"
	"time"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/perpetual"
	"blep.ai/data/monitor"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	conf := config.GetConfig()

	conf.SetDefault("contracts.url", "https://metadata.perp.exchange/production.json")
	conf.SetDefault("perpetual.factoryAddress", "0x8C29F6F7fc1999aB84b476952E986F974Acb3824")
	conf.SetDefault("perpetual.kafka.txID", "perpetual")
	conf.SetDefault("network", "xdai")
	conf.SetDefault("perpetual.author", "nakji")
	conf.SetDefault("perpetual.version", "0_0_0")
	conf.SetDefault("blockTime", 15*time.Second)
	conf.SetDefault("waitBlocks", 4)

	//	Perpetual Protocol Layer 1 initial block: 12463451
	//	Perpetual Protocol Layer 2 initial block: 13507265
	pflag.StringP("network", "n", "xdai", "network to connect to e.g. ethereum")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)

	if err := validateFlags(conf); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	networkName := conf.GetString("network")
	connectorName := "perpetual-" + networkName

	//	Register for Prometheus alerts
	monitor.StartMonitor(connectorName)

	config := &perpetual.Config{
		DbConfig:       conf.GetString("timescaledb.connection"),
		Env:            conf.GetString("kafka.env"),
		KafkaUrl:       conf.GetString("kafka.url"),
		ContractsUrl:   conf.GetString("contracts.url"),
		FactoryAddress: conf.GetString("perpetual.factoryAddress"),
		Author:         conf.GetString("perpetual.author"),
		ProtocolName:   "perpetual",
		ConnectorName:  connectorName,
		NetworkName:    networkName,
		Version:        conf.GetString("perpetual.version"),
		RpcUrls:        conf.GetStringSlice(fmt.Sprintf("%s.rpc", networkName)),
		FromBlock:      conf.GetUint64("from-block"),
		NumBlocks:      conf.GetUint64("num-blocks"),
	}

	connector := perpetual.New(config)
	connector.Start()
}

func validateFlags(conf config.IConfig) error {
	networkName := conf.GetString("network")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableNetworks := map[string]bool{"ethereum": true, "xdai": true}

	if isExists := availableNetworks[networkName]; !isExists {
		errorMsg := "network is not supported, please try again with one of these: "
		for k, _ := range availableNetworks {
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
