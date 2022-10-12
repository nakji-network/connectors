package near

import (
	"context"
	"embed"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	reflect "reflect"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	channelSize  = 1000
	port         = "3030"
	backfillPort = "3031"
)

//go:embed target/release/nakji_near_client
var f embed.FS

var endpointMap = map[string]proto.Message{
	"/block":   &Block{},
	"/tx":      &Transaction{},
	"/receipt": &Receipt{},
	"/outcome": &ExecutionOutcome{},
}

type Subscription struct {
	host                 string
	port                 string
	backfillPort         string
	fromBlock            uint64
	numBlocks            uint64
	events               []string
	msgTypes             []proto.Message
	done                 chan struct{}
	ctx                  context.Context
	blockChan            chan proto.Message
	transactionChan      chan proto.Message
	receiptChan          chan proto.Message
	executionOutcomeChan chan proto.Message
	errChan              chan error
}

func NewSubscription(ctx context.Context, host string, msgTypes []proto.Message, events []string, fromBlock uint64, numBlocks uint64) (*Subscription, error) {
	bin, err := f.ReadFile("target/release/nakji_near_client")
	if err != nil {
		log.Fatal().Err(err).Msg("failed read embedded Nakji Near Client binary")
	}

	err = os.WriteFile("nakji_near_client", bin, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to write Nakji Near Client binary to local fs")
	}

	sub := &Subscription{
		host:                 host,
		port:                 port,
		backfillPort:         backfillPort,
		events:               events,
		msgTypes:             msgTypes,
		fromBlock:            fromBlock,
		numBlocks:            numBlocks,
		done:                 make(chan struct{}),
		ctx:                  ctx,
		blockChan:            make(chan proto.Message, channelSize),
		transactionChan:      make(chan proto.Message, channelSize),
		receiptChan:          make(chan proto.Message, channelSize),
		executionOutcomeChan: make(chan proto.Message, channelSize),
		errChan:              make(chan error, channelSize),
	}
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		select {
		case <-interrupt:
			sub.Unsubscribe()
		case <-ctx.Done():
			sub.Unsubscribe()
		}
	}()

	go sub.subscribe()
	return sub, nil
}

func (s *Subscription) Unsubscribe() {
	log.Info().Str("host", s.host).Msg("shutting down subscription")
	close(s.done)
}

func (s *Subscription) Done() <-chan struct{} {
	return s.done
}

func (s *Subscription) Blocks() <-chan proto.Message {
	return s.blockChan
}

func (s *Subscription) Transactions() <-chan proto.Message {
	return s.transactionChan
}

func (s *Subscription) Receipts() <-chan proto.Message {
	return s.receiptChan
}

func (s *Subscription) ExecutionOutcomes() <-chan proto.Message {
	return s.executionOutcomeChan
}

func (s *Subscription) Err() <-chan error {
	return s.errChan
}

func (s *Subscription) subscribe() {
	// Start Lake Stream from current block
	go s.initLakeStream(s.port, 0, 0)
	// Backfill if needed
	if s.fromBlock > 0 || s.numBlocks > 0 {
		go s.backfill()
	}
}

func (s *Subscription) backfill() {
	log.Info().Msg("start backfilling")
	if s.fromBlock > 0 {
		//s.getEvents(s.fromBlock, latestBlock)
		s.initLakeStream(s.backfillPort, s.fromBlock, 0)
	} else if s.numBlocks > 0 {
		//s.getEvents(latestBlock-s.numBlocks, latestBlock)
		s.initLakeStream(s.backfillPort, 0, s.numBlocks)
	}
}

func (s *Subscription) initLakeStream(port string, fromBlock uint64, numBlocks uint64) {
	isBackfill := false

	// Setup command for Lake stream
	cmd := []string{
		"./nakji_near_client",
	}

	cmd = append(cmd, fmt.Sprintf("--port=%s", port))

	// Arguments for binary are detrmined by backfill config
	if fromBlock > 0 {
		cmd = append(cmd, "from-block")
		cmd = append(cmd, fmt.Sprintf("%d", fromBlock))
		isBackfill = true
	} else if numBlocks > 0 {
		cmd = append(cmd, "num-blocks")
		cmd = append(cmd, fmt.Sprintf("%d", numBlocks))
		isBackfill = true
	} else {
		cmd = append(cmd, "from-latest")
	}

	execCmd := exec.CommandContext(s.ctx, cmd[0], cmd[1:]...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Run binary in goroutine
	go func() {
		if err := execCmd.Run(); err != nil {
			s.Unsubscribe()
			log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error running NEAR Rust binary: ", err))
		}
	}()

	// Allow time for ws server to start
	time.Sleep(5 * time.Second)

	// Connect to ws server to receive data from Lake stream
	for endpoint, msgType := range endpointMap {
		// Check if msgType is in sub's msgTypes
		for _, subType := range s.msgTypes {
			if reflect.TypeOf(msgType) == reflect.TypeOf(subType) {
				go s.wsStream(port, endpoint, msgType, isBackfill)
			}
		}
	}

}

func (s *Subscription) wsStream(port string, endpoint string, msgType proto.Message, isBackfill bool) {
	host := fmt.Sprintf("%s%s%s", s.host, ":", port)
	u := url.URL{Scheme: "ws", Host: host, Path: endpoint}
	log.Info().Str("host", host).Str("endpoint", endpoint).Msg("connecting to ws server")

	// Output chan determined by msgType
	var ch chan protoreflect.ProtoMessage
	switch reflect.TypeOf(msgType) {
	case reflect.TypeOf(&Block{}):
		ch = s.blockChan
	case reflect.TypeOf(&Transaction{}):
		ch = s.transactionChan
	case reflect.TypeOf(&Receipt{}):
		ch = s.receiptChan
	case reflect.TypeOf(&ExecutionOutcome{}):
		ch = s.executionOutcomeChan
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error connecting to websocket", err))
	}
	defer conn.Close()

	for {
		select {
		case <-s.done:
			return
		default:
			_, wsMessage, err := conn.ReadMessage()
			if err != nil && isBackfill {
				log.Debug().Msg(fmt.Sprintf("%s Backfill ws connection closed", err))
				return
			} else if err != nil {
				log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error reading ws message", err))
			}

			msg := proto.Clone(msgType)

			if err := proto.Unmarshal(wsMessage, msg); err != nil {
				log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error unmarshalling message: ", err))
			}

			ch <- msg
		}
	}
}
