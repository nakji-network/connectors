package briq

import (
	"context"
	_ "embed"
	"encoding/json"
	"os"
	"os/signal"
	"strings"
	"time"

	"blep.ai/data/chain/starknet"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	TokenNamespace = "starknet"
	Namespace      = "briq"
)

//go:embed smart-contracts/briq_abi.json
var briqABI []byte

type ABIPayload struct {
	Name     string `json:"Name"`
	Type     string `json:"type"`
	Selector string `json:"selector"`
}

type Connector struct {
	Topics       map[string]kafkautils.Topic
	Producer     *kafkautils.Producer
	QueueSink    chan *kafkautils.Message
	Db           *database.Db
	SelectorName map[string]ABIPayload
	Addresses    map[string]bool
	Network      string
}

func NewConnector(topics map[string]kafkautils.Topic, kp *kafkautils.Producer, qs chan *kafkautils.Message, db *database.Db, network string) *Connector {
	return &Connector{
		Topics:       topics,
		Producer:     kp,
		QueueSink:    qs,
		Db:           db,
		SelectorName: make(map[string]ABIPayload),
		Network:      network,
	}
}

func (c *Connector) Start() {
	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go c.SubscribeNewBlock(ctx)

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			c.Producer.Close()
			return
		}
	}
}

func (c *Connector) SubscribeNewBlock(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	var num int64 = -1000

	for {
		select {
		case <-ticker.C:
			block, err := starknet.GetBlock(c.Network, num+1)
			if err != nil {
				log.Error().Err(err).Int64("block number", num+1).Msg("failed to get new block")
				continue
			}
			if int64(block.Number) == num {
				continue
			}
			log.Info().Uint64("block number", block.Number).Msg("received new block")

			num = int64(block.Number)
			c.ParseBlock(block)
		case <-ctx.Done():
			log.Info().Msg("worker cancelled and shutting down")
			ticker.Stop()
			return
		}
	}
}

func (c *Connector) ParseBlock(block *starknet.Block) {
	txReceipts := block.TxReceipts
	bts := block.Ts
	ts := timestamppb.New(time.Unix(int64(bts), 0))
	for _, txr := range txReceipts {
		events := txr.Events
		for _, event := range events {
			if _, ok := c.Addresses[event.FromAddress]; ok {
				key := event.Keys[0]
				if abi, ok := c.SelectorName[key]; ok {
					ev := &BriqEvent{
						Ts:          ts,
						BlockHash:   block.Hash,
						TxHash:      txr.TxHash,
						FromAddress: event.FromAddress,
						Name:        abi.Name,
						Type:        abi.Type,
						Args:        event.Data,
					}

					msg := &kafkautils.Message{
						Topic:    c.Topics[strings.ToLower(abi.Name)],
						Key:      kafkautils.NewKey(Namespace, event.FromAddress),
						ProtoMsg: ev,
					}
					c.QueueSink <- msg

				} else {
					log.Error().Str("selector", key).Msg("failed to find the selector")
				}
			}
		}
	}
}

func (c *Connector) BuildSelectorName(data []ABIPayload) {
	for _, d := range data {
		c.SelectorName[d.Selector] = d
	}
}

func ParseABIPayload() []ABIPayload {
	var data []ABIPayload
	err := json.Unmarshal(briqABI, &data)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshall into ABIPayload")
	}

	return data
}
