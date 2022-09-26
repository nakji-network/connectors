package flow

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxApiUsage    = 20 // Limit GetEventsForHeightRange API usage rate.
	defaultTimeout = 3 * time.Minute
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
	cache     *lru.Cache
	curBlock  uint64
	apiUsage  uint64 // GetEventsForHeightRange API usage rate.
	done      chan struct{}
	logChan   chan Log
	errChan   chan error
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
	cache, err := lru.New(100000)
	if err != nil {
		return nil, err
	}
	sub := &Subscription{
		host:      host,
		grpc:      cli,
		events:    events,
		fromBlock: fromBlock,
		numBlocks: numBlocks,
		cache:     cache,
		done:      make(chan struct{}),
		logChan:   make(chan Log, 1000),
		errChan:   make(chan error, 1000),
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
	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-sub.done:
				t.Stop()
				return
			case <-t.C:
				// Reset API usage to zero every second.
				atomic.StoreUint64(&sub.apiUsage, 0)
			}
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

func (s *Subscription) Logs() <-chan Log {
	return s.logChan
}

func (s *Subscription) Err() <-chan error {
	return s.errChan
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
			latest, err := s.getLatestBlockHeight()
			if err != nil {
				s.errChan <- err
				continue
			}
			if s.curBlock == latest {
				continue
			}
			// Backfill if needed
			if !backfilled && (s.fromBlock > 0 || s.numBlocks > 0) {
				s.backfill(latest)
				backfilled = true
			} else if s.curBlock == 0 {
				go s.getEvents(latest, latest)
			} else {
				go s.getEvents(s.curBlock+1, latest)
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
	// Get events for maximun of 250 blocks at once
	for start := startHeight; start <= endHeight; start += 250 {
		select {
		case <-s.done:
			return
		default:
			end := start + 249
			if end > endHeight {
				end = endHeight
			}
			for _, event := range s.events {
				blockEvents, err := s.getEventsForHeightRange(event, start, end)
				if err != nil {
					s.errChan <- err
					continue
				}
				for _, blockEvent := range blockEvents {
					for _, ev := range blockEvent.Events {
						t := strings.Split(ev.Type, ".")
						s.logChan <- Log{
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

func (s *Subscription) getLatestBlockHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	header, err := s.grpc.GetLatestBlockHeader(ctx, true)
	if err != nil {
		return 0, err
	}
	s.cache.Add(header.Height, header.ID)
	return header.Height, nil
}

func (s *Subscription) getBlockIDByHeight(height uint64) (flow.Identifier, error) {
	if id, ok := s.cache.Get(height); ok {
		return id.(flow.Identifier), nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	header, err := s.grpc.GetBlockHeaderByHeight(ctx, height)
	if err != nil {
		return flow.Identifier{}, err
	}
	s.cache.Add(header.Height, header.ID)
	return header.ID, nil
}

func (s *Subscription) getEventsForHeightRange(eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	if endHeight > startHeight && atomic.AddUint64(&s.apiUsage, 1) <= maxApiUsage {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()
		return s.grpc.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
	} else { // Exceed maximun API usage. Try using different APIs.
		blockIDs := make([]flow.Identifier, 0, endHeight-startHeight+1)
		for height := startHeight; height <= endHeight; height++ {
			id, err := s.getBlockIDByHeight(height)
			if err != nil {
				s.errChan <- err
				continue
			}
			blockIDs = append(blockIDs, id)
		}
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()
		return s.grpc.GetEventsForBlockIDs(ctx, eventType, blockIDs)
	}
}
