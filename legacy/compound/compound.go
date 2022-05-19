package compound

import (
	"context"
	"os"
	"os/signal"
	"strings"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/common"
	. "blep.ai/data/connectors/source/compound/ctoken"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type Connector struct {
	KP                *kafkautils.Producer
	Topics            map[string]kafkautils.Topic
	ClientPool        *ethclient.ClientPool
	ContractAddresses []ethcommon.Address
	*Status
}

const (
	Namespace = "compound"
)

// Status for kubernetes health check
type Status struct {
	Live  bool
	Ready bool
	Data  map[string]interface{}
}

func NewConnector(kp *kafkautils.Producer, topics map[string]kafkautils.Topic, clientPool *ethclient.ClientPool, ContractAddresses []ethcommon.Address) *Connector {
	return &Connector{
		KP:                kp,
		Topics:            topics,
		ClientPool:        clientPool,
		ContractAddresses: ContractAddresses,
		Status:            new(Status),
	}
}

func (c *Connector) Start() {
	c.Status.Live = true

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer c.ClientPool.Close()

	// Keepalive by subscribing to new heads
	unsubscribe := make(chan interface{})

	headers := c.ClientPool.ConsumeHeaders(unsubscribe)

	logs := make(chan ethtypes.Log)

	defer close(logs)

	c.KP.EnableTransactions()
	c.Status.Ready = true

	// Initialize CEther ABI for reading logs
	contractAbi1, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read CEther abi")
	}
	contractAbi2, err := abi.JSON(strings.NewReader(CtokenMetaData.ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read CEther abi")
	}

	contractAbis := []abi.ABI{contractAbi1, contractAbi2}

	sink := make(chan *kafkautils.Message) // 0 length for real time streams

	// out chan written to kafka
	go c.KP.WriteAndCommitSink(sink)

	sub := c.CEtherLogsListener(logs)

	for {
		select {
		case header := <-headers:
			log.Debug().
				Str("block", header.Number.String()).
				Uint64("ts", header.Time).
				Msg("header received")

			ethclient.CacheBlockTimestamp(header.Hash(), header.Time)
		case err := <-sub.Err():
			c.Status.Live = false
			c.Status.Ready = false
			log.Error().Err(err).Msg("sub error")
			return
		case evLog := <-logs:
			c.ProcessLogEvent(c.ClientPool, contractAbis, evLog, sink)
		case <-interrupt:
			log.Debug().Msg("interrupt")
			return
		}
	}
}

// CEtherLogsListener listens to all contracts that emit CEther related events.
func (c *Connector) CEtherLogsListener(logs chan ethtypes.Log) ethereum.Subscription {
	query := ethereum.FilterQuery{
		Addresses: c.ContractAddresses,
	}

	sub, err := c.ClientPool.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal().Err(err).Msg("subscribing CEther filter logs failed")
	}

	return sub
}

func (c *Connector) ProcessLogEvent(clientPool *ethclient.ClientPool, contractAbis []abi.ABI, evLog ethtypes.Log, sink chan *kafkautils.Message) {
	// Add timestamp from block since logs don't include timestamp
	ts, err := clientPool.GetLogTimestamp(evLog, nil)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", evLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	// Writes to out chan
	err = c.WriteEventToChan(contractAbis, evLog, ts, sink)
	if err != nil {
		log.Error().Err(err).
			Interface("evLog", evLog).
			Msg("ProcessLogEvent error")
	}
}

// WriteEventToChan parses the event and writes it to the channel for kafka.
func (c *Connector) WriteEventToChan(contractAbis []abi.ABI, evLog ethtypes.Log, ts uint64, out chan<- *kafkautils.Message) error {
	var ev *abi.Event
	var err error
	var idx int

	for i, cAbi := range contractAbis {
		ev, err = cAbi.EventByID(evLog.Topics[0])
		if err == nil {
			idx = i
			break
		}
	}

	if ev == nil {
		log.Warn().Msg("ignore if event id isn't defined in a partial ABI")
		return nil
	}

	contractAbi := contractAbis[idx]

	switch ev.Name {
	case "Mint":
		event := new(CompoundMint)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Mint event error")
			return nil
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["mint"],
			Key:   kafkautils.NewKey(Namespace, evLog.Address.Hex()),
			ProtoMsg: &Mint{
				Ts:         common.UnixToTimestampPb(int64(ts * 1000)),
				Block:      evLog.BlockNumber,
				Idx:        uint64(evLog.Index),
				Tx:         evLog.TxHash.Bytes(),
				Minter:     event.Minter.Bytes(),
				MintAmount: event.MintAmount.Bytes(),
				MintTokens: event.MintTokens.Bytes(),
			},
		}
	case "Redeem":
		event := new(CompoundRedeem)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Redeem event error")
			return nil
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["redeem"],
			Key:   kafkautils.NewKey(Namespace, evLog.Address.Hex()),
			ProtoMsg: &Redeem{
				Ts:           common.UnixToTimestampPb(int64(ts * 1000)),
				Block:        evLog.BlockNumber,
				Idx:          uint64(evLog.Index),
				Tx:           evLog.TxHash.Bytes(),
				Redeemer:     event.Redeemer.Bytes(),
				RedeemAmount: event.RedeemAmount.Bytes(),
				RedeemTokens: event.RedeemTokens.Bytes(),
			},
		}
	case "Borrow":
		event := new(CompoundBorrow)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack Borrow event error")
			return nil
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["borrow"],
			Key:   kafkautils.NewKey(Namespace, evLog.Address.Hex()),
			ProtoMsg: &Borrow{
				Ts:             common.UnixToTimestampPb(int64(ts * 1000)),
				Block:          evLog.BlockNumber,
				Idx:            uint64(evLog.Index),
				Tx:             evLog.TxHash.Bytes(),
				Borrower:       event.Borrower.Bytes(),
				BorrowAmount:   event.BorrowAmount.Bytes(),
				AccountBorrows: event.AccountBorrows.Bytes(),
				TotalBorrows:   event.TotalBorrows.Bytes(),
			},
		}
	case "RepayBorrow":
		event := new(CompoundRepayBorrow)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack RepayBorrow event error")
			return nil
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["repayborrow"],
			Key:   kafkautils.NewKey(Namespace, evLog.Address.Hex()),
			ProtoMsg: &RepayBorrow{
				Ts:             common.UnixToTimestampPb(int64(ts * 1000)),
				Block:          evLog.BlockNumber,
				Idx:            uint64(evLog.Index),
				Tx:             evLog.TxHash.Bytes(),
				Payer:          event.Payer.Bytes(),
				Borrower:       event.Borrower.Bytes(),
				RepayAmount:    event.RepayAmount.Bytes(),
				AccountBorrows: event.AccountBorrows.Bytes(),
				TotalBorrows:   event.TotalBorrows.Bytes(),
			},
		}
	case "LiquidateBorrow":
		event := new(CompoundLiquidateBorrow)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, evLog); err != nil {
			log.Error().Err(err).Msg("Unpack LiquidateBorrow event error")
			return nil
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["liquidateborrow"],
			Key:   kafkautils.NewKey(Namespace, evLog.Address.Hex()),
			ProtoMsg: &LiquidateBorrow{
				Ts:               common.UnixToTimestampPb(int64(ts * 1000)),
				Block:            evLog.BlockNumber,
				Idx:              uint64(evLog.Index),
				Tx:               evLog.TxHash.Bytes(),
				Liquidator:       event.Liquidator.Bytes(),
				Borrower:         event.Borrower.Bytes(),
				RepayAmount:      event.RepayAmount.Bytes(),
				CTokenCollateral: event.CTokenCollateral.Bytes(),
				SeizeTokens:      event.SeizeTokens.Bytes(),
			},
		}
	}

	return nil
}

func ConvertRawAddress(rawAddresses ...string) []ethcommon.Address {
	var addresses []ethcommon.Address

	for _, addr := range rawAddresses {
		addresses = append(addresses, ethcommon.HexToAddress(addr))
	}

	return addresses
}
