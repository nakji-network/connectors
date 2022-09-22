// go run ./connectors/source/daofi/cmd/daofi_backfill/main.go --fromBlock 11945500 --toBlock 11945502
package main

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/aggregator"
	"blep.ai/data/connectors/aggregator/ohlcv"
	"blep.ai/data/connectors/source/daofi"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
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
	// Load config in here to support flags
	pflag.Int64("fromBlock", 15066279, "backfill from this block")
	pflag.Int64("toBlock", 15069503, "backfill to this block")
	pflag.Int64("blockChunk", 20, "get 100 blocks at a time to prevent timeouts")
	pflag.Int64P("maxConnections", "c", 200, "max simultaneous connections")
	pflag.Bool("upsert", true, "uses batch instead of COPY FROM when existing entries might exist")

	var conf = config.GetConfig()
	conf.SetDefault("daofi.kafka.topic.liquiditypool_reserve", ".fct.nakji.common.0_0_0.liquiditypool_reserve")
	conf.SetDefault("daofi.kafka.topic.liquiditypool_change", ".fct.nakji.common.0_0_0.liquiditypool_change")
	conf.SetDefault("daofi.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("ohlcv.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("daofi.factoryAddress", "0x408b10d4a4EA307017B647732D7BcE95A3fD76B2")
	conf.SetDefault("xdai.archival-rpc", []string{
		"https://xdai-archive.blockscout.com",
	})

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
	aggregators := map[string]aggregator.Aggregator{
		"trade": {
			Source:   make(chan interface{}),
			Reducer:  ohlcv.ReducerOHLCVV,
			Topic:    kafkautils.MustParseTopic(conf.GetString("ohlcv.kafka.topic.market_trade"), conf.GetString("kafka.env")),
			Duration: 1 * time.Minute,
			KafkaUrl: "", // don't need to write directly to db
		},
	}

	// Start aggregator processors
	for _, a := range aggregators {
		go a.AggregateToDb(db)
	}

	connector := daofi.New(daofi.Config{
		//EnableBlockCache: true,
		Namespace:      "daofi",
		TokenNamespace: "xdai",
		KP:             nil,
		RPCURLs:        conf.GetStringSlice("xdai.archival-rpc"),
		LPTopic:        kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:   kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:     kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Db:             db,
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("daofi.factoryAddress")),
	})

	messageCh := connector.PairHistoryPipeline(ethereum.FilterQuery{
		FromBlock: big.NewInt(conf.GetInt64("fromBlock")),
		ToBlock:   big.NewInt(conf.GetInt64("toBlock")),
	}, conf.GetInt64("blockChunk"), conf.GetInt64("maxConnections"))

	start := time.Now()

	// database write buffer
	buffer := make([]*kafkautils.Message, 0)

	for msg := range messageCh {
		status.EventsProcessed++

		if msg.Topic.Schema() != "nakji.common.0_0_0.liquiditypool_change" {
			continue
		}

		buffer = append(buffer, msg)

		// Flush to db
		if len(buffer) == 2000 {
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
