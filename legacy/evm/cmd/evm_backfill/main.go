package main

import (
	"context"
	"math/big"
	"net/url"
	"sync"
	"time"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/evm"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/avast/retry-go"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	// Load config in here to support flags
	pflag.Int64("fromBlock", 10000835, "backfill from this block (inclusive)")
	pflag.Int64("toBlock", 11945500, "backfill to this block (inclusive)")
	pflag.Int64("batch", 100, "batch size per db write")
	pflag.String("ns", "eth", "namespace")
	pflag.String("s", "eth", "subject")
	pflag.Int64("n", 100, "concurrent connections")

	conf := config.GetConfig()

	endpoint := url.URL{
		Scheme: conf.GetString("ethereum.archival.scheme"),
		User:   url.UserPassword(conf.GetString("ethereum.archival.username"), conf.GetString("ethereum.archival.password")),
		Host:   conf.GetString("ethereum.archival.host"),
		Path:   conf.GetString("ethereum.archival.path"),
	}

	// Init historical db
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	log.Info().Msgf("connecting to %s", endpoint.String())
	client, err := ethclient.Dial(endpoint.String())
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer client.Close()

	start := time.Now()

	toBlock := conf.GetInt64("toBlock")
	fromBlock := conf.GetInt64("fromBlock")
	ns, s := conf.GetString("ns"), conf.GetString("s")
	batchi := conf.GetInt("batch") - 1
	concurrency := conf.GetInt64("n")

	chunkSize := (toBlock - fromBlock) / concurrency

	type queryRange struct {
		from int64
		to   int64 // not inclusive
	}
	chunks := []queryRange{}
	for i := fromBlock; i < toBlock; i += chunkSize {
		log.Info().Ints64("chunk", []int64{i, min(i+chunkSize, toBlock)}).Msg("chunk")
		chunks = append(chunks, queryRange{i, min(i+chunkSize, toBlock)})
	}

	var wg sync.WaitGroup
	wg.Add(len(chunks))
	for _, chunk := range chunks {
		go func(chunk queryRange) {
			defer wg.Done()

			// buffer
			var msgs []*kafkautils.Message // make([]*kafkautils.Message, conf.GetInt64("batch"))
			// i for block num, bi for batch num
			for i, bi := chunk.to-1, 0; i >= chunk.from; i, bi = i-1, bi+1 {
				var header *ethtypes.Header
				log.Info().Int64("header", i).Msg("getting header")

				// Retrieve block header from network
				err := retry.Do(
					func() error {
						var err error
						header, err = client.HeaderByNumber(context.Background(), big.NewInt(i))
						return err
					},
					retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
						log.Warn().Err(err).Uint("n", n).Msg("retrying...")
						return retry.BackOffDelay(n, err, config)
					}),
				)
				if err != nil {
					log.Fatal().Msg("HeaderByNumber failed")
				}

				msgs = append(msgs, &kafkautils.Message{
					Key:      kafkautils.NewKey(ns, s),
					ProtoMsg: evm.UnmarshalEthHeader(header),
				})

				// Batch to use CopyBy
				if bi == batchi {
					//log.Info().Interface("msgs", msgs).Msg("messages")
					err = db.CopyFromKafkaMessages(context.Background(), "nakji.evm.0_0_0.chain_block", msgs)
					if err != nil {
						log.Error().Err(err).Msg("CopyFromKafkaMessages error")
					}
					bi = -1
					msgs = nil
				}

			}

		}(chunk)
	}

	wg.Wait()

	log.Info().
		Dur("duration", time.Since(start)).
		Msg("Done")
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
