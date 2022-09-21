package subscription

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	DefaultLogChanSize = 1000
	DefaultErrChanSize = 1000
)

type GrpcClient[Log any] interface {
	GetLatestBlock(ctx context.Context) (height uint64, err error)
	GetLogsForHeightRange(ctx context.Context, topics []string, startHeight uint64, endHeight uint64) (logs <-chan Log, errs <-chan error)
}

type Subscription[Log any] struct {
	Grpc              GrpcClient[Log]
	BlockTime         time.Duration
	Topics            []string
	FromBlock         uint64
	NumBlocks         uint64
	LogChanSize       int
	ErrChanSize       int
	BackfillBatchSize uint64

	initOnce sync.Once

	curBlock uint64

	done    chan struct{}
	logChan chan Log
	errChan chan error
}

func (sub *Subscription[Log]) Subscribe(ctx context.Context) error {
	sub.initOnce.Do(sub.init)

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
	return nil
}

func (sub *Subscription[Log]) Unsubscribe() {
	log.Info().Msg("shutting down subscription")
	close(sub.done)
}

func (sub *Subscription[Log]) Done() <-chan struct{} {
	return sub.done
}

func (sub *Subscription[Log]) Errs() <-chan error {
	return sub.errChan
}

func (sub *Subscription[Log]) Logs() <-chan Log {
	return sub.logChan
}

func (sub *Subscription[Log]) init() {
	if sub.LogChanSize == 0 {
		sub.LogChanSize = DefaultLogChanSize
	}
	if sub.ErrChanSize == 0 {
		sub.ErrChanSize = DefaultErrChanSize
	}
	sub.done = make(chan struct{})
	sub.logChan = make(chan Log, sub.LogChanSize)
	sub.errChan = make(chan error, sub.ErrChanSize)
}

func (sub *Subscription[Log]) subscribe() {
	backfilled := false
	t := time.NewTicker(sub.BlockTime)
	for {
		select {
		case <-sub.done:
			t.Stop()
			return
		case <-t.C:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			latestBlock, err := sub.Grpc.GetLatestBlock(ctx)
			cancel()
			if err != nil {
				sub.errChan <- err
				continue
			}
			if sub.curBlock == latestBlock {
				continue
			}
			// Backfill if needed
			if !backfilled && (sub.FromBlock > 0 || sub.NumBlocks > 0) {
				sub.backfill(latestBlock)
				backfilled = true
			} else if sub.curBlock == 0 {
				go sub.getEvents(latestBlock, latestBlock)
			} else {
				go sub.getEvents(sub.curBlock+1, latestBlock)
			}
			sub.curBlock = latestBlock
		}
	}
}

func (sub *Subscription[Log]) backfill(latestBlock uint64) {
	log.Info().Msg("start backfilling")
	defer log.Info().Msg("stop backfilling")
	if sub.FromBlock > 0 {
		sub.getEvents(sub.FromBlock, latestBlock)
	} else if sub.NumBlocks > 0 {
		sub.getEvents(latestBlock-sub.NumBlocks, latestBlock)
	}
}

func (sub *Subscription[Log]) getEvents(startHeight uint64, endHeight uint64) {
	if sub.BackfillBatchSize > 0 && endHeight-startHeight > sub.BackfillBatchSize {
		for start := startHeight; start <= endHeight; start += sub.BackfillBatchSize {
			select {
			case <-sub.done:
				return
			default:
				end := start + sub.BackfillBatchSize - 1
				if end > endHeight {
					end = endHeight
				}
				sub.getEvents(start, end)
			}
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		logs, errs := sub.Grpc.GetLogsForHeightRange(ctx, sub.Topics, startHeight, endHeight)
		for {
			select {
			case log, ok := <-logs:
				if ok {
					sub.logChan <- log
				} else {
					logs = nil
				}
			case err, ok := <-errs:
				if ok {
					sub.errChan <- err
				} else {
					errs = nil
				}
			}
			if logs == nil && errs == nil {
				break
			}
		}
	}
}
