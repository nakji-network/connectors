package woofi_polygon

import (
	"context"
	"time"

	"github.com/nakji-network/connector/chain/ethereum"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	NetworkName string
	FromBlock   uint64
	NumBlocks   uint64
}

type Connector struct {
	*Config
	*ethereum.Connector
	contracts map[string]ISmartContract
}

func New(config *Config) *Connector {
	return &Connector{
		Config:    config,
		contracts: make(map[string]ISmartContract),
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[sc.Address()] = sc
}

func (c *Connector) GetContract(addr string) ISmartContract {
	return c.contracts[addr]
}

func (c *Connector) Start() {
	var (
		events    []proto.Message
		addresses []ethcommon.Address
	)

	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
		addresses = append(addresses, ethcommon.HexToAddress(contract.Address()))
	}

	c.Connector = ethereum.NewConnector(context.Background(), addresses, c.NetworkName)

	c.RegisterProtos(kafkautils.MsgTypeFct, events...)

	c.Sub.Subscribe(context.Background())

	if c.FromBlock > 0 || c.NumBlocks > 0 {
		c.RegisterProtos(kafkautils.MsgTypeBf, events...)
		go c.backfill(addresses)
	}

	for {
		select {
		case <-c.Sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.Sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.Sub.Logs():
			if msg := c.parse(vLog.Log); msg != nil {
				c.EventSink <- &kafkautils.Message{
					MsgType:  kafkautils.MsgTypeFct,
					ProtoMsg: msg,
				}
			}
		}
	}
}

func (c *Connector) backfill(addresses []ethcommon.Address) {
	logs := make(chan types.Log, 1000)
	go func() {
		for vLog := range logs {
			if msg := c.parse(vLog); msg != nil {
				c.EventSink <- &kafkautils.Message{
					MsgType:  kafkautils.MsgTypeBf,
					ProtoMsg: msg,
				}
			}
		}
		log.Info().Msg("backfill stopped")
	}()
	log.Info().Msg("backfill started")
	if err := ethereum.BackfillFrom(context.Background(), c.Client, addresses, logs, c.FromBlock, c.NumBlocks); err != nil {
		log.Error().Err(err).Msg("backfill failed")
	}
}

func (c *Connector) parse(vLog types.Log) protoreflect.ProtoMessage {
	contract := c.GetContract(vLog.Address.String())
	if contract == nil {
		log.Info().Str("address", vLog.Address.String()).Msg("event from unsupported address")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	ts, err := c.Sub.GetBlockTime(ctx, vLog)
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve block timestamp")
	}

	var sender *ethcommon.Address
	var receiver *ethcommon.Address

	tx, err := c.Sub.TransactionByHash(ctx, vLog.TxHash)
	if err != nil {
		log.Error().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction by hash")
	} else {
		receiver = tx.To()
		adrr, err := c.Sub.TransactionSender(ctx, tx, vLog.BlockHash, vLog.TxIndex)
		if err != nil {
			log.Error().Err(err).Str("txHash", vLog.TxHash.String()).Msg("failed to retrieve transaction sender")
		} else {
			sender = &adrr
		}
	}

	return contract.Message(Log{
		Log:             vLog,
		BlockTime:       time.Unix(int64(ts), 0),
		SenderAddress:   sender,
		ReceiverAddress: receiver,
	})
}
