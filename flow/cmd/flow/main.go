package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/proto"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("rpc.url", grpc.MainnetHost)
	c.Config.SetDefault("maxRetry", 7)
	c.Config.SetDefault("maxGrpcMsgSize", 64*1024*1024) // 64 MB
	c.Config.SetDefault("maxApiUsage", 20)
	c.Config.SetDefault("maxWorkerPoolSize", 5)
	c.Config.SetDefault("cacheSize", 2000)
	c.Config.SetDefault("channelSize", 1000)
	c.Config.SetDefault("timeout", 3*time.Minute)
	c.Config.SetDefault("reportInterval", 1*time.Minute)

	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	conf := &flow.Config{
		Host:              c.Config.GetString("rpc.url"),
		FromBlock:         c.Config.GetUint64("from-block"),
		NumBlocks:         c.Config.GetUint64("num-blocks"),
		MaxGrpcMsgSize:    c.Config.GetInt("maxGrpcMsgSize"),
		MaxApiUsage:       c.Config.GetInt("maxApiUsage"),
		MaxWorkerPoolSize: c.Config.GetInt("maxWorkerPoolSize"),
		CacheSize:         c.Config.GetInt("cacheSize"),
		ChannelSize:       c.Config.GetInt("channelSize"),
		Timeout:           c.Config.GetDuration("timeout"),
		ReportInterval:    c.Config.GetDuration("reportInterval"),
	}

	var topicTypes []proto.Message
	topicTypes = append(topicTypes, &flow.Block{}, &flow.Transaction{})

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, topicTypes...)
	if conf.FromBlock > 0 || conf.NumBlocks > 0 {
		c.RegisterProtos(kafkautils.MsgTypeBf, topicTypes...)
	}

	m := flow.New(c, conf)
	m.Start()
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
