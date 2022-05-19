// go run ./connectors/source/uniswapv2/cmd/uniswapv2_backfill/main.go --fromBlock 11945500 --toBlock 11945502
package main

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"

	_ "net/http/pprof"
)

type Status struct {
	Step              int
	EventsProcessed   int64
	Errors            []error
	FromBlock         int64
	ToBlock           int64
	SuccessfulQueries []ethereum.FilterQuery
	FailedQueries     []ethereum.FilterQuery
}

func main() {
	// Memory profiling
	go http.ListenAndServe(":6060", nil)

	// Load config in here to support flags
	pflag.Int64("fromBlock", 10000835, "backfill from this block")
	pflag.Int64("toBlock", 12000000, "backfill to this block")
	pflag.Int64("blockChunk", 3000, "get 100 blocks at a time to prevent timeouts")
	pflag.Int64P("maxConnections", "c", 10, "max simultaneous connections")
	pflag.Bool("upsert", true, "uses batch instead of COPY FROM when existing entries might exist")
	pflag.Int("dbBatch", 3000, "batch db writes by x items")
	pflag.Int64("processPairConcurrency", 4096, "max goroutines for pairProcessLog")

	var conf = config.GetConfig()
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_reserve", ".fct.nakji.common.0_0_0.liquiditypool_reserve")
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_change", ".fct.nakji.common.0_0_0.liquiditypool_change")
	conf.SetDefault("uniswapv2.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("ohlcv.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("uniswapv2.factoryAddress", "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")

	// get status of this long running status
	status := Status{
		FromBlock: conf.GetInt64("fromBlock"),
		ToBlock:   conf.GetInt64("toBlock"),
	}
	http.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		js, err := json.Marshal(status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	go http.ListenAndServe(":8080", nil)

	// Init historical db
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(uniswapv2.TopicTypes)
	kafkautils.TopicTypeRegistry.Load(common.TopicTypes)

	// Initialize aggregators chans and handlers
	//aggregators := map[string]aggregator.Aggregator{
	//"trade": {
	//Source: make(chan interface{}),
	//Reducer: ohlcv.ReducerOHLCVV,
	//Topic: kafkautils.MustParseTopic(conf.GetString("ohlcv.kafka.topic.market_trade"), conf.GetString("kafka.env")),
	//Duration: 1 * time.Minute,
	//KafkaUrl: "", // don't need to write directly to db
	//},
	//}

	// Start aggregator processors
	//for _, a := range aggregators {
	//go a.AggregateToDb(db)
	//}

	connector := uniswapv2.New(uniswapv2.Config{
		EnableBlockCache: true,
		Namespace:        "uniswapv2",
		TokenNamespace:   "ethereum",
		KP:               nil,
		RPCURLs: []string{(&url.URL{
			Scheme: conf.GetString("ethereum.archival.scheme"),
			User:   url.UserPassword(conf.GetString("ethereum.archival.username"), conf.GetString("ethereum.archival.password")),
			Host:   conf.GetString("ethereum.archival.host"),
			Path:   conf.GetString("ethereum.archival.path"),
		}).String(),
			(&url.URL{ // hacky to get past wss error
				Scheme: "wss",
				User:   url.UserPassword(conf.GetString("ethereum.archival.username"), conf.GetString("ethereum.archival.password")),
				Host:   conf.GetString("ethereum.archival.host"),
				Path:   conf.GetString("ethereum.archival.path"),
			}).String()},
		LPTopic:                kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:           kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:             kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Db:                     db,
		FactoryAddress:         ethcommon.HexToAddress(conf.GetString("uniswapv2.factoryAddress")),
		ProcessPairConcurrency: conf.GetInt64("processPairConcurrency"),
	})

	topics := []ethcommon.Hash{
		ethcommon.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"), // Mint
		ethcommon.HexToHash("0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496"), // Burn
		ethcommon.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"), // Swap
	}

	messageCh := connector.PairHistoryPipeline(ethereum.FilterQuery{
		FromBlock: big.NewInt(conf.GetInt64("fromBlock")),
		ToBlock:   big.NewInt(conf.GetInt64("toBlock")),
		Topics:    [][]ethcommon.Hash{topics},
	}, conf.GetInt64("blockChunk"), conf.GetInt64("maxConnections"))

	start := time.Now()

	// database write buffer
	buffer := make([]*kafkautils.Message, 0, conf.GetInt("dbBatch"))

	for msg := range messageCh {
		status.EventsProcessed++

		if msg.Topic.Schema() != "nakji.common.0_0_0.liquiditypool_change" {
			continue
		}

		buffer = append(buffer, msg)

		// Flush to db
		if len(buffer) == conf.GetInt("dbBatch") {
			if conf.GetBool("upsert") {
				errs := db.InsertBatchKafkaMessages(buffer)
				for _, err := range errs {
					log.Error().Err(err).
						Msg("InsertBatchKafkaMessages error")
					status.Errors = append(status.Errors, err)
				}
			} else {
				err = db.CopyFromKafkaMessages(context.Background(), "nakji.common.0_0_0.liquiditypool_change", buffer)
				if err != nil {
					log.Error().Err(err).
						Msg("CopyFromKafkaMessages error")
					status.Errors = append(status.Errors, err)
				}
			}

			log.Info().
				Dur("time", time.Since(start)).
				Interface("lastmsg", msg).
				//Interface("status", status).
				Msg("PairHistoryStreamer in progress")

			buffer = nil
		}
	}

	// Flush remaining buffer
	if conf.GetBool("upsert") {
		errs := db.InsertBatchKafkaMessages(buffer)
		for _, err := range errs {
			log.Error().Err(err).
				Msg("InsertBatchKafkaMessages error")
			status.Errors = append(status.Errors, err)
		}
	} else {
		err = db.CopyFromKafkaMessages(context.Background(), "nakji.common.0_0_0.liquiditypool_change", buffer)
		if err != nil {
			log.Error().Err(err).
				Msg("CopyFromKafkaMessages error")
			status.Errors = append(status.Errors, err)
		}
	}

	log.Info().
		Dur("totaltime", time.Since(start)).
		Interface("status", &status).
		Msg("PairHistoryStreamer completed")

	// loop through all events in history in order
	//for msg := range messageCh {
	//log.Info().
	//Interface("msg", msg).
	//Msg("")

	////err := db.InsertProto(msg.Topic.Schema(), msg.ProtoMsg, msg.Key.Tuple())
	////if err != nil {
	////log.Fatal().Msg("Database write error fail fast")
	////}
	////aggMsg := aggregator.Message{Key: msg.Message.Key, Message: msg.ProtoMsg, TopicPartition: &msg.TopicPartition}

	////if a, ok := aggregators[msg.Topic.Schema()]; ok {
	////a.Source <- aggMsg
	////} else {
	////log.Warn().Str("schema", msg.Topic.Schema()).Msg("Unhandled topic schema")
	////}
	//}
}
