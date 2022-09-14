package flow

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
)

type EventType struct {
	ContractAddr string
	ContractName string
	EventName    string
}

func (t EventType) String() string {
	return fmt.Sprintf("A.%s.%s.%s", t.ContractAddr, t.ContractName, t.EventName)
}

type Log struct {
	BlockID   flow.Identifier
	Height    uint64
	Timestamp time.Time
	// Type is the qualified event type.
	Type EventType
	// TransactionID is the ID of the transaction this event was emitted from.
	TransactionID flow.Identifier
	// TransactionIndex is the index of the transaction this event was emitted from, within its containing block.
	TransactionIndex int
	// EventIndex is the index of the event within the transaction it was emitted from.
	EventIndex int
	// Value contains the event data.
	Value cadence.Event
	// Bytes representing event data.
	Payload []byte
}

type Subscription struct {
	host      string
	grpc      *grpc.Client
	fromBlock uint64
	numBlocks uint64
	events    []string
	curBlock  *flow.BlockHeader
	logs      chan Log
	done      chan bool
	errchan   chan error
}

func NewSubscription(ctx context.Context, host string, events []string, fromBlock uint64, numBlocks uint64) (*Subscription, error) {
	cli, err := grpc.NewClient(host)
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx); err != nil {
		return nil, err
	}
	sub := &Subscription{
		host:      host,
		grpc:      cli,
		events:    events,
		fromBlock: fromBlock,
		numBlocks: numBlocks,
		logs:      make(chan Log, 1000),
		done:      make(chan bool),
		errchan:   make(chan error, 1000),
	}
	go func() {
		interupt := make(chan os.Signal, 1)
		signal.Notify(interupt, os.Interrupt)
		select {
		case <-interupt:
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
	s.done <- true
	close(s.logs)
	close(s.done)
	close(s.errchan)
}

func (s *Subscription) Done() <-chan bool {
	return s.done
}

func (s *Subscription) Err() <-chan error {
	return s.errchan
}

func (s *Subscription) Logs() <-chan Log {
	return s.logs
}

func (s *Subscription) subscribe() {
	backfilled := false
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-s.done:
			t.Stop()
			return
		case <-t.C:
			// Try getting next block until success
			if err := s.nextBlock(); err != nil {
				s.errchan <- err
				continue
			}
			// backfill if needed
			if !backfilled && (s.fromBlock > 0 || s.numBlocks > 0) {
				s.backfill()
				backfilled = true
			}
			// get events for lastest block
			s.getEvents(s.curBlock.Height, s.curBlock.Height)
		}
	}
}

func (s *Subscription) nextBlock() error {
	var (
		nextBlock *flow.BlockHeader
		err       error
	)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if s.curBlock == nil {
		nextBlock, err = s.grpc.GetLatestBlockHeader(ctx, true)
	} else {
		nextBlock, err = s.grpc.GetBlockHeaderByHeight(ctx, s.curBlock.Height+1)
	}
	if err != nil {
		return err
	}
	s.curBlock = nextBlock
	return nil
}

func (s *Subscription) backfill() {
	log.Info().Msg("start backfilling")
	defer log.Info().Msg("stop backfilling")
	var startHeight uint64
	if s.fromBlock > 0 {
		startHeight = s.fromBlock
	} else if s.numBlocks > 0 {
		startHeight = s.curBlock.Height - s.numBlocks
	}
	// backfill 250 blocks at once
	for ; startHeight <= s.curBlock.Height; startHeight += 250 {
		select {
		case <-s.done:
			return
		default:
			endHeight := startHeight + 249
			if startHeight > s.curBlock.Height {
				startHeight = s.curBlock.Height
			}
			if endHeight > s.curBlock.Height {
				endHeight = s.curBlock.Height
			}
			s.getEvents(startHeight, endHeight)
		}
	}
}

func (s *Subscription) getEvents(startHeight uint64, endHeight uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	for _, event := range s.events {
		blockEvents, err := s.grpc.GetEventsForHeightRange(ctx, event, startHeight, endHeight)
		if err != nil {
			s.errchan <- err
			continue
		}
		for _, blockEvent := range blockEvents {
			for _, ev := range blockEvent.Events {
				t := strings.Split(ev.Type, ".")
				s.logs <- Log{
					BlockID:   blockEvent.BlockID,
					Height:    blockEvent.Height,
					Timestamp: blockEvent.BlockTimestamp,
					Type: EventType{
						ContractAddr: t[1],
						ContractName: t[2],
						EventName:    t[3],
					},
					TransactionID:    ev.TransactionID,
					TransactionIndex: ev.TransactionIndex,
					EventIndex:       ev.EventIndex,
					Value:            ev.Value,
					Payload:          ev.Payload,
				}
			}
		}
	}
}
