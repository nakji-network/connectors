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
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	grpc      *flowgrpc.Client
	fromBlock uint64
	numBlocks uint64
	events    []string
	curBlock  *flow.BlockHeader
	logs      chan Log
	done      chan bool
	errchan   chan error
}

func NewSubscription(ctx context.Context, host string, events []string, fromBlock uint64, numBlocks uint64) (*Subscription, error) {
	cli, err := flowgrpc.NewClient(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1<<25)), // 32 MB
	)
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
	t := time.NewTicker(2 * time.Second) // Block initialization time on main net is 2 seconds
	for {
		select {
		case <-s.done:
			t.Stop()
			return
		case <-t.C:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			latest, err := s.grpc.GetLatestBlockHeader(ctx, true)
			cancel()
			if err != nil {
				s.errchan <- err
				continue
			}
			if s.curBlock != nil && s.curBlock.Height == latest.Height {
				continue
			}
			// Backfill if needed
			if !backfilled && (s.fromBlock > 0 || s.numBlocks > 0) {
				s.backfill(latest.Height)
				backfilled = true
			} else if s.curBlock == nil {
				go s.getEvents(latest.Height, latest.Height)
			} else {
				go s.getEvents(s.curBlock.Height+1, latest.Height)
			}
			s.curBlock = latest
		}
	}
}

func (s *Subscription) backfill(latestBlock uint64) {
	log.Info().Msg("start backfilling")
	defer log.Info().Msg("stop backfilling")
	if s.fromBlock > 0 {
		s.getEvents(s.fromBlock, latestBlock)
	} else if s.numBlocks > 0 {
		s.getEvents(latestBlock-s.numBlocks, latestBlock)
	}
}

func (s *Subscription) getEvents(startHeight uint64, endHeight uint64) {
	for start := startHeight; start <= endHeight; start += 250 {
		select {
		case <-s.done:
			return
		default:
			end := start + 249
			if start > endHeight {
				start = endHeight
			}
			if end > endHeight {
				end = endHeight
			}
			for _, event := range s.events {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
				blockEvents, err := s.grpc.GetEventsForHeightRange(ctx, event, start, end)
				cancel()
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
	}
}
