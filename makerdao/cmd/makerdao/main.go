package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/makerdao"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("contracts.url", "https://changelog.makerdao.com/releases/mainnet/active/contracts.json")
	c.Config.SetDefault("makerdao.author", "nakji")
	c.Config.SetDefault("makerdao.factoryAddress", "0x5a464C28D19848f44199D003BeF5ecc87d090F87")
	c.Config.SetDefault("makerdao.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	//	MakerDAO initial block is #8928141
	pflag.StringP("network", "n", "ethereum", "network to connect to e.g. ethereum")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	networkName := c.Config.GetString("network")

	// Register topic and protobuf type mappings
	protos := make([]protoreflect.ProtoMessage, len(makerdao.TopicTypes[networkName]))
	i := 0
	for _, topicProto := range makerdao.TopicTypes[networkName] {
		protos[i] = topicProto
		i++
	}

	c.RegisterProtos(protos...)

	conf := &makerdao.Config{
		ContractsUrl:   c.Config.GetString("contracts.url"),
		FactoryAddress: c.Config.GetString("makerdao.factoryAddress"),
		ConnectorName:  "makerdao-" + networkName,
		NetworkName:    networkName,
		FromBlock:      c.Config.GetUint64("from-block"),
		NumBlocks:      c.Config.GetUint64("num-blocks"),
	}

	m := makerdao.New(c, conf)
	m.Start()
}

func validateFlags(conf config.IConfig) error {
	networkName := conf.GetString("network")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableNetworks := map[string]bool{"ethereum": true, "arbitrum": true, "optimism": true}

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
