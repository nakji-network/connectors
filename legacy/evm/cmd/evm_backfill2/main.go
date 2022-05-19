// go run ./connectors/source/evm/cmd/evm_backfill2/main.go --fromBlock=10000000 --toBlock=10200000
// 3 minutes for 2M blocks on Chainstack dedicated node
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/evm"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/avast/retry-go"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	// Load config in here to support flags
	pflag.Int64("fromBlock", 10000835, "backfill from this block (inclusive)")
	pflag.Int64("toBlock", 11945500, "backfill to this block (exclusive)")
	pflag.Int64("chunk", 100000, "batch size per graphql query")
	pflag.String("ns", "eth", "namespace")
	pflag.String("s", "eth", "subject")

	conf := config.GetConfig()

	endpoint := url.URL{
		Scheme: conf.GetString("ethereum.archival.scheme"),
		User:   url.UserPassword(conf.GetString("ethereum.archival.username"), conf.GetString("ethereum.archival.password")),
		Host:   conf.GetString("ethereum.archival.host"),
		Path:   "/graphql",
	}

	// Init historical db
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	start := time.Now()

	toBlock := conf.GetInt64("toBlock")
	fromBlock := conf.GetInt64("fromBlock")
	ns, s := conf.GetString("ns"), conf.GetString("s")
	chunkSize := conf.GetInt64("chunk")

	type queryRange struct {
		from int64
		to   int64 // not inclusive
	}
	chunks := []queryRange{}
	for i := fromBlock; i < toBlock; i += chunkSize {
		log.Info().Ints64("chunk", []int64{i, min(i+chunkSize, toBlock)}).Msg("chunk")
		chunks = append(chunks, queryRange{i, min(i+chunkSize, toBlock)})
	}

	for _, chunk := range chunks {
		// Do the request to get []Block0's
		var blocks []evm.Block0
		err := retry.Do(
			func() error {
				var err error
				blocks, err = getBlocks(endpoint, getQuery(chunk.from, chunk.to-1), 20*time.Second)
				return err
			},
			retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
				log.Warn().Err(err).Uint("n", n).Msg("retrying...")
				return retry.BackOffDelay(n, err, config)
			}),
		)
		if err != nil {
			log.Fatal().Msg("getBlocks from graphql api failed")
		}

		msgs := make([]*kafkautils.Message, len(blocks))
		for i := range blocks {
			msgs[i] = &kafkautils.Message{
				Key:      kafkautils.NewKey(ns, s),
				ProtoMsg: &blocks[i],
			}
		}

		err = db.CopyFromKafkaMessages(context.Background(), "block", msgs)
		if err != nil {
			log.Error().Err(err).Msg("CopyFromKafkaMessages error")
		}
	}

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

func getQuery(from, to int64) string {
	return fmt.Sprintf(`
{"query":"{\n  blocks(from: %d, to: %d) {\n    timestamp\n    number\n    hash\n    difficulty\n    gasLimit\n    gasUsed\n    nonce\n  }\n}","variables":null}
	`, from, to)
}

func getBlocks(url url.URL, query string, timeout time.Duration) ([]evm.Block0, error) {
	api := http.Client{
		Timeout: timeout,
	}

	reqBody := bytes.NewBuffer([]byte(query))
	req, err := http.NewRequest(http.MethodPost, url.String(), reqBody)
	if err != nil {
		return []evm.Block0{}, err
	}

	res, getErr := api.Do(req)
	if getErr != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	jsonRes := GQLResponse{}
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		log.Fatal().Err(jsonErr).Msg("failed to parse")
	}

	block0s := make([]evm.Block0, len(jsonRes.Data.Blocks))
	for i, block := range jsonRes.Data.Blocks {
		block0s[i] = evm.Block0{
			Ts:         &timestamp.Timestamp{Seconds: block.Ts.Int64()},
			Hash:       block.Hash.Bytes(),
			Num:        block.Number,
			Difficulty: block.Difficulty.Uint64(),
			GasLimit:   block.GasLimit,
			GasUsed:    block.GasUsed,
			Nonce:      block.Nonce.Uint64(),
		}
	}

	return block0s, nil
}

type GQLResponse struct {
	Data struct {
		Blocks []GQLBlock `json:"blocks"`
	} `json:"data"`
}

type GQLBlock struct {
	Ts         BigIntHex
	Number     uint64
	Hash       BigIntHex
	Difficulty BigIntHex
	GasLimit   uint64
	GasUsed    uint64
	Nonce      BigIntHex
}

// ByteString is a byte array that serializes to hex
type BigIntHex struct{ big.Int }

// MarshalJSON serializes ByteArray to hex
//func (s BigIntHex) MarshalJSON() ([]byte, error) {
//bytes, err := json.Marshal(fmt.Sprintf("%x", string(s)))
//return bytes, err
//}

// UnmarshalJSON deserializes ByteArray to hex
func (s *BigIntHex) UnmarshalJSON(data []byte) error {
	var x big.Int
	unquoted, _ := strconv.Unquote(string(data))
	_, ok := x.SetString(unquoted, 0)
	if !ok {
		return fmt.Errorf("%s cannot be unmarshaled to big.Int", data)
	}
	s.Int = x
	return nil
}
