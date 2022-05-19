package abracadabra

import (
	"context"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/rs/zerolog/log"

	"blep.ai/data/chain/ethereum/connector"
	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/common"
	"blep.ai/data/connectors/source/abracadabra/magiccrv"
)

const (
	Namespace       = "abracadabra"
	MagicCRVAddress = "0x98BF3E7da6f5A81630730d538715E35d8c0D3eDE"
)

type AbracadabraConnector struct {
	*connector.EthereumConnector
}

func NewConnector() *AbracadabraConnector {
	c := connector.NewConnector()
	return &AbracadabraConnector{
		EthereumConnector: c,
	}
}

func (c *AbracadabraConnector) Start() {
	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	crvAddress := ethcommon.HexToAddress(MagicCRVAddress)
	go c.EventsListener(ctx, crvAddress, c.processEvent, magiccrv.MagiccrvABI)

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			c.ClientPool.Close()
			c.Producer.Close()
			return
		}
	}
}

func (c *AbracadabraConnector) processEvent(event *abi.Event, eventLog ethtypes.Log, ts uint64, magiccrvABI abi.ABI) {
	switch event.Name {
	case "Approval":
		event := new(magiccrv.MagiccrvApproval)
		if err := ethclient.UnpackLog(magiccrvABI, event, "Approval", eventLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return
		}

		msg := &kafkautils.Message{
			Topic: c.Topics["magiccrvapproval"],
			Key:   kafkautils.NewKey(Namespace, "approval"),
			ProtoMsg: &magiccrv.Approval{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Owner:   event.Owner.Bytes(),
				Spender: event.Spender.Bytes(),
				Amount:  event.Amount.Bytes(),
			},
		}

		c.QueueSink <- msg

	case "Transfer":
		event := new(magiccrv.MagiccrvTransfer)
		if err := ethclient.UnpackLog(magiccrvABI, event, "Transfer", eventLog); err != nil {
			log.Error().Err(err).Msg("Unpack event error")
			return
		}

		msg := &kafkautils.Message{
			Topic: c.Topics["magiccrvtransfer"],
			Key:   kafkautils.NewKey(Namespace, "transfer"),
			ProtoMsg: &magiccrv.Transfer{
				Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
				From:   event.From.Bytes(),
				To:     event.To.Bytes(),
				Amount: event.Amount.Bytes(),
			},
		}

		c.QueueSink <- msg
	}
}
