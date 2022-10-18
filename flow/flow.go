package flow

import (
	"context"
	"time"

	"github.com/nakji-network/connector"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type ContractID struct {
	Addr string
	Name string
}

type Config struct {
	Host              string
	FromBlock         uint64
	NumBlocks         uint64
	MaxRetry          int
	MaxGrpcMsgSize    int
	MaxApiUsage       int // Limit GetEventsForHeightRange API usage rate.
	MaxWorkerPoolSize int
	CacheSize         int
	ChannelSize       int
	Timeout           time.Duration
	ReportInterval    time.Duration
}

type Connector struct {
	*connector.Connector
	*Config
	contracts    map[ContractID]ISmartContract
	errCounter   uint64
	blockCounter uint64
	txCounter    uint64
	eventCounter map[string]uint64
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector:    c,
		Config:       config,
		contracts:    make(map[ContractID]ISmartContract),
		eventCounter: make(map[string]uint64),
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[ContractID{sc.Address(), sc.Name()}] = sc
}

func (c *Connector) GetContract(addr string, name string) ISmartContract {
	return c.contracts[ContractID{addr, name}]
}

func (c *Connector) Start() {
	var events []string
	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
	}

	ctx := context.Background()

	sub, err := NewSubscription(ctx, c.Config, events)
	if err != nil {
		log.Fatal().Err(err).Str("host", c.Host).Msg("connection error")
	}

	t := time.NewTicker(c.ReportInterval)
	defer t.Stop()

	for {
		select {
		case <-sub.Done():
			c.reportCounters()
			log.Info().Msg("connector shutdown")
			return

		// Listen to error channel
		case err := <-sub.Err():
			c.errCounter++
			log.Error().Err(err).Str("host", c.Host).Msg("subscription failed")

		// Listen for new blocks
		case block := <-sub.Blocks():
			c.blockCounter++
			c.EventSink <- block

		// Listen for new transactions
		case tx := <-sub.Transactions():
			c.txCounter++
			c.EventSink <- tx

		// Listen to event logs
		case vLog := <-sub.Logs():
			if msg := c.parse(vLog); msg != nil {
				c.eventCounter[vLog.Type.String()]++
				c.EventSink <- msg
			}

		// Report counters
		case <-t.C:
			c.reportCounters()
		}
	}
}

func (c *Connector) parse(vLog *Log) proto.Message {
	contract := c.GetContract(vLog.Type.ContractAddr, vLog.Type.ContractName)
	if contract == nil {
		log.Info().Str("address", vLog.Type.ContractAddr).Str("name", vLog.Type.ContractName).Msg("event from unsupported address")
		return nil
	}
	return contract.Message(vLog)
}

func (c *Connector) reportCounters() {
	log.Info().Msgf("# Error: %d", c.errCounter)
	log.Info().Msgf("# Block: %d", c.blockCounter)
	log.Info().Msgf("# Transaction: %d", c.txCounter)
	for event, counter := range c.eventCounter {
		log.Info().Msgf("# %s: %d", event, counter)
	}
}
