package main

import (
	"fmt"
	"time"

	"blep.ai/data/config"
	starknet "blep.ai/data/connectors/source/ethstarknet"
	"blep.ai/data/monitor"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	conf := config.GetConfig()

	conf.SetDefault("contracts.url", "https://alpha-mainnet.starknet.io/feeder_gateway/get_contract_addresses")
	conf.SetDefault("starknet.kafka.txID", "starknet")
	conf.SetDefault("starknet.author", "nakji")
	conf.SetDefault("starknet.version", "0_0_0")
	conf.SetDefault("blockTime", 15*time.Second)
	conf.SetDefault("waitBlocks", 4)

	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)

	if err := validateFlags(conf); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	//	Register for Prometheus alerts
	monitor.StartMonitor("starknet")

	config := &starknet.Config{
		DbConfig:     conf.GetString("timescaledb.connection"),
		Env:          conf.GetString("kafka.env"),
		KafkaUrl:     conf.GetString("kafka.url"),
		ContractsUrl: conf.GetString("contracts.url"),
		Author:       conf.GetString("starknet.author"),
		ProtocolName: "starknet",
		NetworkName:  "ethereum",
		Version:      conf.GetString("starknet.version"),
		RpcUrls:      conf.GetStringSlice("ethereum.rpc"),
		FromBlock:    conf.GetUint64("from-block"),
		NumBlocks:    conf.GetUint64("num-blocks"),
	}

	connector := starknet.New(config)
	connector.Start()
}

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
