// Code generated by connectorgen - Edit as necessary.
package axie

import (
	"context"
	"fmt"
	"math/big"
	"runtime"
	"strings"
	"sync"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/common"
	"blep.ai/data/connectors/source/axie/axienft"
	"blep.ai/data/connectors/source/axie/axs"
	"blep.ai/data/connectors/source/axie/bridge"
	"blep.ai/data/connectors/source/axie/slp"
	"github.com/nakji-network/connector/kafkautils"

	"golang.org/x/sync/semaphore"
)

const (
	Namespace      = "axie"
	TokenNamespace = "ethereum"
)

type AxieConnector struct {
	KP         kafkautils.ProducerInterface
	Topics     map[string]kafkautils.Topic
	ClientPool ethclient.ETHClientPool
	addresses  map[string][]ethcommon.Address
	blockCache map[uint64]uint64
}

func NewConnector(
	kp kafkautils.ProducerInterface,
	addresses map[string][]ethcommon.Address,
	topics map[string]kafkautils.Topic,
	ethClientPool ethclient.ETHClientPool,

) *AxieConnector {
	var blockCache map[uint64]uint64

	return &AxieConnector{
		KP:         kp,
		Topics:     topics,
		ClientPool: ethClientPool,
		addresses:  addresses,
		blockCache: blockCache,
	}
}

func (c *AxieConnector) Start(ctx context.Context, backfillNumBlocks uint64) {
	// Mainly serves as Keepalive for websocket connection to RPC endpoint by subscribing to new heads
	unsubscribe := make(chan interface{})
	headers := c.ClientPool.ConsumeHeaders(unsubscribe)
	axienftLogs, sub0 := c.startListener(ctx, "axienft")
	axsLogs, sub1 := c.startListener(ctx, "axs")
	bridgeLogs, sub2 := c.startListener(ctx, "bridge")
	slpLogs, sub3 := c.startListener(ctx, "slp")

	sink := make(chan *kafkautils.Message, 10000)
	errorSubs := []<-chan error{sub0.Err(), sub1.Err(), sub2.Err(), sub3.Err()}
	out := common.MergeErrChans(errorSubs...)

	err := c.KP.EnableTransactions()
	if err != nil {
		log.Fatal().Err(err).Msg("Transaction was not enabled")
	}
	go c.KP.WriteAndCommitSink(sink)

	var once1 sync.Once
	var once2 sync.Once
	var once3 sync.Once
	var once4 sync.Once
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("worker cancelled and shutting down")
			return
		case header := <-headers:
			log.Debug().
				Str("block", header.Number.String()).
				Uint64("ts", header.Time).
				Msg("header received")

			ethclient.CacheBlockTimestamp(header.Hash(), header.Time)
		case err = <-out:
			log.Fatal().Err(err).Msg("Event listener failed")
			return
		case evLog := <-axienftLogs:
			if evLog.Removed {
				continue
			}

			go once1.Do(func() {
				c.backfill(sink, evLog.BlockNumber, backfillNumBlocks, "axienft", c.AxienftLogToMsg)
			})

			msg := c.AxienftLogToMsg(evLog)

			if msg != nil {
				sink <- msg
			}
		case evLog := <-axsLogs:
			if evLog.Removed {
				continue
			}

			go once2.Do(func() {
				c.backfill(sink, evLog.BlockNumber, backfillNumBlocks, "axs", c.AxsLogToMsg)
			})

			msg := c.AxsLogToMsg(evLog)

			if msg != nil {
				sink <- msg
			}
		case evLog := <-bridgeLogs:
			if evLog.Removed {
				continue
			}

			go once3.Do(func() {
				c.backfill(sink, evLog.BlockNumber, backfillNumBlocks, "bridge", c.BridgeLogToMsg)
			})

			msg := c.BridgeLogToMsg(evLog)

			if msg != nil {
				sink <- msg
			}
		case evLog := <-slpLogs:
			if evLog.Removed {
				continue
			}

			go once4.Do(func() {
				c.backfill(sink, evLog.BlockNumber, backfillNumBlocks, "slp", c.SlpLogToMsg)
			})

			msg := c.SlpLogToMsg(evLog)

			if msg != nil {
				sink <- msg
			}
		}
	}
}

func (c *AxieConnector) startListener(ctx context.Context, contractName string) (chan ethtypes.Log, event.Subscription) {
	query := geth.FilterQuery{
		Addresses: c.addresses[contractName],
	}
	eventLogs := make(chan ethtypes.Log)
	sub, err := c.ClientPool.SubscribeFilterLogs(ctx, query, eventLogs)
	if err != nil {
		msg := fmt.Sprintf("%s contract listener failed", contractName)
		log.Fatal().Err(err).
			Interface("query", query).
			Msg(msg)
	}
	msg := fmt.Sprintf("%s contract listener live", contractName)
	log.Info().Interface("query", query).Msg(msg)

	return eventLogs, sub
}

func (c *AxieConnector) AxienftLogToMsg(evLog ethtypes.Log) *kafkautils.Message {
	ts, err := c.ClientPool.GetLogTimestamp(evLog, c.blockCache)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", evLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	axienftAbi, err := abi.JSON(strings.NewReader(axienft.AxienftABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read Axienft ABI")
	}

	ev, err := axienftAbi.EventByID(evLog.Topics[0])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to find event")
	}

	if ev == nil {
		return nil
	}

	switch ev.Name {
	case "Approval":
		event := new(axienft.AxienftApproval)
		if err := ethclient.UnpackLog(axienftAbi, event, "Approval", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_approval"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.Approval{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				Owner:    event.Owner.Bytes(),
				Approved: event.Approved.Bytes(),
				TokenId:  event.TokenId.Bytes(),
			},
		}
	case "ApprovalForAll":
		event := new(axienft.AxienftApprovalForAll)
		if err := ethclient.UnpackLog(axienftAbi, event, "ApprovalForAll", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_approvalforall"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.ApprovalForAll{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				Owner:    event.Owner.Bytes(),
				Operator: event.Operator.Bytes(),
				Approved: event.Approved,
			},
		}
	case "AxieEvolved":
		event := new(axienft.AxienftAxieEvolved)
		if err := ethclient.UnpackLog(axienftAbi, event, "AxieEvolved", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_axieevolved"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.AxieEvolved{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				AxieId:   event.AxieId.Bytes(),
				OldGenes: event.OldGenes.Bytes(),
				NewGenes: event.NewGenes.Bytes(),
			},
		}
	case "AxieRebirthed":
		event := new(axienft.AxienftAxieRebirthed)
		if err := ethclient.UnpackLog(axienftAbi, event, "AxieRebirthed", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_axierebirthed"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.AxieRebirthed{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				AxieId: event.AxieId.Bytes(),
				Genes:  event.Genes.Bytes(),
			},
		}
	case "AxieRetired":
		event := new(axienft.AxienftAxieRetired)
		if err := ethclient.UnpackLog(axienftAbi, event, "AxieRetired", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_axieretired"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.AxieRetired{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				AxieId: event.AxieId.Bytes(),
			},
		}
	case "AxieSpawned":
		event := new(axienft.AxienftAxieSpawned)
		if err := ethclient.UnpackLog(axienftAbi, event, "AxieSpawned", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_axiespawned"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.AxieSpawned{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				AxieId: event.AxieId.Bytes(),
				Owner:  event.Owner.Bytes(),
				Genes:  event.Genes.Bytes(),
			},
		}
	case "Transfer":
		event := new(axienft.AxienftTransfer)
		if err := ethclient.UnpackLog(axienftAbi, event, "Transfer", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axienft_transfer"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axienft.Transfer{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				From:    event.From.Bytes(),
				To:      event.To.Bytes(),
				TokenId: event.TokenId.Bytes(),
			},
		}
	}

	return nil
}
func (c *AxieConnector) AxsLogToMsg(evLog ethtypes.Log) *kafkautils.Message {
	ts, err := c.ClientPool.GetLogTimestamp(evLog, c.blockCache)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", evLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	axsAbi, err := abi.JSON(strings.NewReader(axs.AxsABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read Axs ABI")
	}

	ev, err := axsAbi.EventByID(evLog.Topics[0])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to find event")
	}

	if ev == nil {
		return nil
	}

	switch ev.Name {
	case "Approval":
		event := new(axs.AxsApproval)
		if err := ethclient.UnpackLog(axsAbi, event, "Approval", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axs_approval"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axs.Approval{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Owner:   event.Owner.Bytes(),
				Spender: event.Spender.Bytes(),
				Value:   event.Value.Bytes(),
			},
		}
	case "Transfer":
		event := new(axs.AxsTransfer)
		if err := ethclient.UnpackLog(axsAbi, event, "Transfer", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["axs_transfer"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &axs.Transfer{
				Ts:    common.UnixToTimestampPb(int64(ts * 1000)),
				From:  event.From.Bytes(),
				To:    event.To.Bytes(),
				Value: event.Value.Bytes(),
			},
		}
	}

	return nil
}
func (c *AxieConnector) BridgeLogToMsg(evLog ethtypes.Log) *kafkautils.Message {
	ts, err := c.ClientPool.GetLogTimestamp(evLog, c.blockCache)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", evLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	bridgeAbi, err := abi.JSON(strings.NewReader(bridge.BridgeABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read Bridge ABI")
	}

	ev, err := bridgeAbi.EventByID(evLog.Topics[0])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to find event")
	}

	if ev == nil {
		return nil
	}

	switch ev.Name {
	case "AdminChanged":
		event := new(bridge.BridgeAdminChanged)
		if err := ethclient.UnpackLog(bridgeAbi, event, "AdminChanged", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_adminchanged"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.AdminChanged{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				OldAdmin: event.OldAdmin.Bytes(),
				NewAdmin: event.NewAdmin.Bytes(),
			},
		}
	case "AdminRemoved":
		event := new(bridge.BridgeAdminRemoved)
		if err := ethclient.UnpackLog(bridgeAbi, event, "AdminRemoved", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_adminremoved"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.AdminRemoved{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				OldAdmin: event.OldAdmin.Bytes(),
			},
		}
	case "Paused":
		event := new(bridge.BridgePaused)
		if err := ethclient.UnpackLog(bridgeAbi, event, "Paused", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_paused"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.Paused{
				Ts: common.UnixToTimestampPb(int64(ts * 1000)),
			},
		}
	case "ProxyUpdated":
		event := new(bridge.BridgeProxyUpdated)
		if err := ethclient.UnpackLog(bridgeAbi, event, "ProxyUpdated", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_proxyupdated"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.ProxyUpdated{
				Ts:  common.UnixToTimestampPb(int64(ts * 1000)),
				New: event.New.Bytes(),
				Old: event.Old.Bytes(),
			},
		}
	case "TokenDeposited":
		event := new(bridge.BridgeTokenDeposited)
		if err := ethclient.UnpackLog(bridgeAbi, event, "TokenDeposited", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_tokendeposited"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.TokenDeposited{
				Ts:               common.UnixToTimestampPb(int64(ts * 1000)),
				DepositId:        event.DepositId.Bytes(),
				Owner:            event.Owner.Bytes(),
				TokenAddress:     event.TokenAddress.Bytes(),
				SidechainAddress: event.SidechainAddress.Bytes(),
				Standard:         event.Standard,
				TokenNumber:      event.TokenNumber.Bytes(),
			},
		}
	case "TokenWithdrew":
		event := new(bridge.BridgeTokenWithdrew)
		if err := ethclient.UnpackLog(bridgeAbi, event, "TokenWithdrew", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_tokenwithdrew"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.TokenWithdrew{
				Ts:           common.UnixToTimestampPb(int64(ts * 1000)),
				WithdrawId:   event.WithdrawId.Bytes(),
				Owner:        event.Owner.Bytes(),
				TokenAddress: event.TokenAddress.Bytes(),
				TokenNumber:  event.TokenNumber.Bytes(),
			},
		}
	case "Unpaused":
		event := new(bridge.BridgeUnpaused)
		if err := ethclient.UnpackLog(bridgeAbi, event, "Unpaused", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["bridge_unpaused"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &bridge.Unpaused{
				Ts: common.UnixToTimestampPb(int64(ts * 1000)),
			},
		}
	}

	return nil
}
func (c *AxieConnector) SlpLogToMsg(evLog ethtypes.Log) *kafkautils.Message {
	ts, err := c.ClientPool.GetLogTimestamp(evLog, c.blockCache)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", evLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	slpAbi, err := abi.JSON(strings.NewReader(slp.SlpABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read Slp ABI")
	}

	ev, err := slpAbi.EventByID(evLog.Topics[0])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to find event")
	}

	if ev == nil {
		return nil
	}

	switch ev.Name {
	case "AdminChanged":
		event := new(slp.SlpAdminChanged)
		if err := ethclient.UnpackLog(slpAbi, event, "AdminChanged", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_adminchanged"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.AdminChanged{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				OldAdmin: event.OldAdmin.Bytes(),
				NewAdmin: event.NewAdmin.Bytes(),
			},
		}
	case "AdminRemoved":
		event := new(slp.SlpAdminRemoved)
		if err := ethclient.UnpackLog(slpAbi, event, "AdminRemoved", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_adminremoved"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.AdminRemoved{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				OldAdmin: event.OldAdmin.Bytes(),
			},
		}
	case "Approval":
		event := new(slp.SlpApproval)
		if err := ethclient.UnpackLog(slpAbi, event, "Approval", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_approval"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.Approval{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Owner:   event.Owner.Bytes(),
				Spender: event.Spender.Bytes(),
				Value:   event.Value.Bytes(),
			},
		}
	case "MinterAdded":
		event := new(slp.SlpMinterAdded)
		if err := ethclient.UnpackLog(slpAbi, event, "MinterAdded", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_minteradded"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.MinterAdded{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				Minter: event.Minter.Bytes(),
			},
		}
	case "MinterRemoved":
		event := new(slp.SlpMinterRemoved)
		if err := ethclient.UnpackLog(slpAbi, event, "MinterRemoved", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_minterremoved"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.MinterRemoved{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				Minter: event.Minter.Bytes(),
			},
		}
	case "Transfer":
		event := new(slp.SlpTransfer)
		if err := ethclient.UnpackLog(slpAbi, event, "Transfer", evLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return nil
		}

		return &kafkautils.Message{
			Topic: c.Topics["slp_transfer"],
			Key:   kafkautils.NewKey(Namespace, event.Raw.Address.Hex()),
			ProtoMsg: &slp.Transfer{
				Ts:    common.UnixToTimestampPb(int64(ts * 1000)),
				From:  event.From.Bytes(),
				To:    event.To.Bytes(),
				Value: event.Value.Bytes(),
			},
		}
	}

	return nil
}

// Backfill last 100 blocks
func (c *AxieConnector) backfill(out chan<- *kafkautils.Message, latestBlockNumber, backfillNumBlocks uint64, contract string, logToMsg func(ethtypes.Log) *kafkautils.Message) {
	filterQuery := geth.FilterQuery{
		FromBlock: big.NewInt(int64(latestBlockNumber - backfillNumBlocks)),
		ToBlock:   big.NewInt(int64(latestBlockNumber)),
		Addresses: c.addresses[contract],
	}

	logchan := make(chan ethtypes.Log)
	errchan := make(chan error)

	go c.ClientPool.ChunkedFilterLogs(context.Background(), filterQuery, 100, 1, logchan, errchan)

	maxWorkers := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(maxWorkers))

	for {
		select {
		case err := <-errchan:
			log.Error().Err(err).Msg("Failed to complete backfill")
		case evLog := <-logchan:
			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Error().Err(err).Msg("Failed to acquire semaaphor")
			}

			go func(evLog ethtypes.Log) {
				defer sem.Release(1)

				// Writes to out chan
				msg := logToMsg(evLog)
				if msg != nil {
					out <- msg
				}
			}(evLog)
		}
	}
}
