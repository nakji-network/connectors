package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi"

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

	c.Config.SetDefault("woofi.author", "nakji")
	c.Config.SetDefault("woofi.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	//	Woo initial block is #18675185
	pflag.StringP("network", "n", "bsc", "network to connect to e.g. bsc")
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	networkName := c.Config.GetString("network")

	// Register topic and protobuf type mappings
	protos := make([]protoreflect.ProtoMessage, len(woofi.TopicTypes[networkName]))
	i := 0
	for _, topicProto := range woofi.TopicTypes[networkName] {
		protos[i] = topicProto
		i++
	}

	c.RegisterProtos(protos...)

	conf := &woofi.Config{
		ConnectorName: "woofi-" + networkName,
		NetworkName:   networkName,
		FromBlock:     c.Config.GetUint64("from-block"),
		NumBlocks:     c.Config.GetUint64("num-blocks"),
	}

	m := woofi.New(c, conf)
	m.Start()
}

func validateFlags(conf config.IConfig) error {
	networkName := conf.GetString("network")
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	availableNetworks := map[string]bool{"bsc": true, "polygon": true}

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
