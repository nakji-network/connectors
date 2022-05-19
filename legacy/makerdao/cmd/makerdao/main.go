package main

import (
	"fmt"
	"time"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/makerdao"
	"blep.ai/data/monitor"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	conf := config.GetConfig()

	conf.SetDefault("contracts.url", "https://changelog.makerdao.com/releases/mainnet/active/contracts.json")
	conf.SetDefault("makerdao.author", "nakji")
	conf.SetDefault("makerdao.factoryAddress", "0x5a464C28D19848f44199D003BeF5ecc87d090F87")
	conf.SetDefault("makerdao.version", "0_0_0")
	conf.SetDefault("blockTime", 15*time.Second)
	conf.SetDefault("waitBlocks", 4)

	//	MakerDAO initial block is #8928141
	pflag.StringP("network", "n", "ethereum", "network to connect to e.g. ethereum")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)

	if err := validateFlags(conf); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	networkName := conf.GetString("network")
	connectorName := "makerdao-" + networkName

	//	Register for Prometheus alerts
	monitor.StartMonitor(connectorName)

	config := &makerdao.Config{
		DbConfig:       conf.GetString("timescaledb.connection"),
		Env:            conf.GetString("kafka.env"),
		KafkaUrl:       conf.GetString("kafka.url"),
		ContractsUrl:   conf.GetString("contracts.url"),
		FactoryAddress: conf.GetString("makerdao.factoryAddress"),
		Author:         conf.GetString("makerdao.author"),
		ProtocolName:   "makerdao",
		ConnectorName:  connectorName,
		NetworkName:    networkName,
		Version:        conf.GetString("makerdao.version"),
		RpcUrls:        conf.GetStringSlice(fmt.Sprintf("%s.rpc", networkName)),
		FromBlock:      conf.GetUint64("from-block"),
		NumBlocks:      conf.GetUint64("num-blocks"),
	}

	connector := makerdao.New(config)
	connector.Start()
}

func validateFlags(conf config.IConfig) error {
	networkName := conf.GetString("network")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableNetworks := map[string]bool{"ethereum": true, "arbitrum": true, "optimism": true}

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
