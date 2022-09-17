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
	DefaultLogChanSize       = 1000
	DefaultErrChanSize       = 1000
	DefaultBackfillBatchSize = 100
)

type GrpcClient[Log any] interface {
	GetLatestBlock(ctx context.Context) (height uint64, err error)
	CheckIfBlockExists(ctx context.Context, height uint64) (exists bool, err error)
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
	errchan chan error
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
	return sub.errchan
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
	if sub.BackfillBatchSize == 0 {
		sub.BackfillBatchSize = DefaultBackfillBatchSize
	}
	sub.done = make(chan struct{})
	sub.logChan = make(chan Log, sub.LogChanSize)
	sub.errchan = make(chan error, sub.ErrChanSize)
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
			// Try getting next block until success
			if ok, err := sub.nextBlock(); err != nil {
				sub.errchan <- err
				continue
			} else if !ok {
				continue
			}
			// Backfill if needed
			if !backfilled && (sub.FromBlock > 0 || sub.NumBlocks > 0) {
				sub.backfill()
				backfilled = true
			}
			// Get events for the lastest block
			go sub.getEvents(sub.curBlock, sub.curBlock)
		}
	}
}

func (sub *Subscription[Log]) nextBlock() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if sub.curBlock == 0 {
		lastestBlock, err := sub.Grpc.GetLatestBlock(ctx)
		if err != nil {
			return false, err
		}
		sub.curBlock = lastestBlock
		return true, nil
	} else {
		ok, err := sub.Grpc.CheckIfBlockExists(ctx, sub.curBlock+1)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
		sub.curBlock++
		return true, nil
	}
}

func (sub *Subscription[Log]) backfill() {
	log.Info().Msg("start backfilling")
	defer log.Info().Msg("stop backfilling")
	var startHeight uint64
	if sub.FromBlock > 0 {
		startHeight = sub.FromBlock
	} else if sub.NumBlocks > 0 {
		startHeight = sub.curBlock - sub.NumBlocks
	}
	// Backfill {{ BackfillBatchSize }} blocks at once
	for ; startHeight <= sub.curBlock; startHeight += sub.BackfillBatchSize {
		select {
		case <-sub.done:
			return
		default:
			endHeight := startHeight + sub.BackfillBatchSize - 1
			if startHeight > sub.curBlock {
				startHeight = sub.curBlock
			}
			if endHeight > sub.curBlock {
				endHeight = sub.curBlock
			}
			sub.getEvents(startHeight, endHeight)
		}
	}
}

func (sub *Subscription[Log]) getEvents(startHeight uint64, endHeight uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	logs, errs := sub.Grpc.GetLogsForHeightRange(ctx, sub.Topics, startHeight, endHeight)
	for log := range logs {
		sub.logChan <- log
	}
	for err := range errs {
		sub.errchan <- err
	}
}
