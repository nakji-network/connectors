package flow

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/rs/zerolog/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	*Config
	grpc                      Repository
	events                    []string
	cache                     *lru.Cache
	curBlock                  uint64
	apiUsage                  uint64 // GetEventsForHeightRange API usage rate.
	done                      chan struct{}
	blockChan                 chan *Block
	transactionChan           chan *Transaction
	logChan                   chan *Log
	historicalBlockChan       chan *Block
	historicalTransactionChan chan *Transaction
	historicalLogChan         chan *Log
	errChan                   chan error
}

func NewSubscription(ctx context.Context, config *Config, events []string) (*Subscription, error) {
	repo, err := NewRepository(ctx, config)
	if err != nil {
		return nil, err
	}
	cache, err := lru.New(config.CacheSize)
	if err != nil {
		return nil, err
	}
	sub := &Subscription{
		Config:                    config,
		grpc:                      repo,
		events:                    events,
		cache:                     cache,
		done:                      make(chan struct{}),
		blockChan:                 make(chan *Block, config.ChannelSize),
		transactionChan:           make(chan *Transaction, config.ChannelSize),
		logChan:                   make(chan *Log, config.ChannelSize),
		historicalBlockChan:       make(chan *Block, config.ChannelSize),
		historicalTransactionChan: make(chan *Transaction, config.ChannelSize),
		historicalLogChan:         make(chan *Log, config.ChannelSize),
		errChan:                   make(chan error, config.ChannelSize),
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
	log.Info().Str("host", s.Host).Msg("shutting down subscription")
	close(s.done)
}

func (s *Subscription) Done() <-chan struct{} {
	return s.done
}

func (s *Subscription) Blocks() <-chan *Block {
	return s.blockChan
}

func (s *Subscription) HistoricalBlocks() <-chan *Block {
	return s.historicalBlockChan
}

func (s *Subscription) Transactions() <-chan *Transaction {
	return s.transactionChan
}

func (s *Subscription) HistoricalTransactions() <-chan *Transaction {
	return s.historicalTransactionChan
}

func (s *Subscription) Logs() <-chan *Log {
	return s.logChan
}

func (s *Subscription) HistoricalLogs() <-chan *Log {
	return s.historicalLogChan
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
			if !backfilled && (s.FromBlock > 0 || s.NumBlocks > 0) {
				go s.backfill(latest - 1)
				backfilled = true
			}
			if s.curBlock == 0 {
				s.subscribeEvents(latest, latest)
			} else {
				s.subscribeEvents(s.curBlock+1, latest)
			}
			s.curBlock = latest
		}
	}
}

func (s *Subscription) backfill(latestBlock uint64) {
	log.Info().Msg("start backfilling")
	defer log.Info().Msg("stop backfilling")
	if s.FromBlock > 0 {
		s.subscribeHistoricalEvents(s.FromBlock, latestBlock)
	} else if s.NumBlocks > 0 {
		s.subscribeHistoricalEvents(latestBlock-s.NumBlocks+1, latestBlock)
	}
}

func (s *Subscription) subscribeEvents(startHeight uint64, endHeight uint64) {
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
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				s.subscribeBlocks(false, start, end)
				wg.Done()
			}()
			go func() {
				s.subscribeLogs(false, start, end)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func (s *Subscription) subscribeHistoricalEvents(startHeight uint64, endHeight uint64) {
	// Get events for maximun of 250 blocks at once in reverse order
	for end := endHeight; end >= startHeight; end -= 250 {
		select {
		case <-s.done:
			return
		default:
			start := end - 249
			if start < startHeight {
				start = startHeight
			}
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				s.subscribeBlocks(true, start, end)
				wg.Done()
			}()
			go func() {
				s.subscribeLogs(true, start, end)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func (s *Subscription) subscribeBlocks(isHistorical bool, startHeight uint64, endHeight uint64) {
	var (
		txMtx        sync.Mutex
		blockMtx     sync.Mutex
		transactions []*Transaction
		blocks       = make([]*Block, 0, endHeight-startHeight+1)
	)
	wpSize1 := endHeight - startHeight + 1
	if wpSize1 > uint64(s.MaxWorkerPoolSize) {
		wpSize1 = uint64(s.MaxWorkerPoolSize) // Limit WorkerPool size
	}
	wp1 := NewWorkerPool(int(wpSize1))
	for height := startHeight; height <= endHeight; height++ {
		blockHeight := height
		wp1.Run(func() {
			// Get Block
			block, err := s.getBlockByHeight(blockHeight)
			if err != nil {
				s.errChan <- err
				return
			}
			b := &Block{
				Ts:                   timestamppb.New(block.Timestamp),
				Id:                   block.ID[:],
				ParentID:             block.ParentID[:],
				Height:               block.Height,
				CollectionGuarantees: make([]*CollectionGuarantee, 0, len(block.CollectionGuarantees)),
				Seals:                make([]*BlockSeal, 0, len(block.Seals)),
			}
			for _, collectionGuarantee := range block.CollectionGuarantees {
				b.CollectionGuarantees = append(b.CollectionGuarantees, &CollectionGuarantee{
					CollectionID: collectionGuarantee.CollectionID[:],
				})
			}
			for _, seal := range block.Seals {
				b.Seals = append(b.Seals, &BlockSeal{
					BlockID:            seal.BlockID[:],
					ExecutionReceiptID: seal.ExecutionReceiptID[:],
				})
			}
			blockMtx.Lock()
			blocks = append(blocks, b)
			blockMtx.Unlock()
			// Get Transactions for Block
			wpSize2 := len(block.CollectionGuarantees)
			if wpSize2 > s.MaxWorkerPoolSize {
				wpSize2 = s.MaxWorkerPoolSize // Limit WorkerPool size
			}
			wp2 := NewWorkerPool(wpSize2)
			for i := range block.CollectionGuarantees {
				collID := block.CollectionGuarantees[i].CollectionID
				wp2.Run(func() {
					txs := s.getTransactions(block, collID)
					txMtx.Lock()
					transactions = append(transactions, txs...)
					txMtx.Unlock()
				})
			}
			wp2.Wait()
		})
	}
	wp1.Wait()
	// Sort blocks by height
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Height < blocks[j].Height
	})
	// Sort transactions by timestamp and collection ID
	sort.Slice(transactions, func(i, j int) bool {
		tx1 := transactions[i]
		tx2 := transactions[j]
		return tx1.Ts.AsTime().Before(tx2.Ts.AsTime()) || (tx1.Ts.AsTime().Equal(tx2.Ts.AsTime()) && string(tx1.CollectionID) < string(tx2.CollectionID))
	})
	if isHistorical {
		for _, block := range blocks {
			s.historicalBlockChan <- block
		}
		for _, transaction := range transactions {
			s.historicalTransactionChan <- transaction
		}
	} else {
		for _, block := range blocks {
			s.blockChan <- block
		}
		for _, transaction := range transactions {
			s.transactionChan <- transaction
		}
	}
}

func (s *Subscription) getTransactions(block *flow.Block, collID flow.Identifier) []*Transaction {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	coll, err := s.grpc.GetCollection(ctx, collID)
	if err != nil {
		s.errChan <- err
		return nil
	}
	var mtx sync.Mutex
	transactions := make([]*Transaction, 0, len(coll.TransactionIDs))
	wpSize := len(coll.TransactionIDs)
	if wpSize > s.MaxWorkerPoolSize {
		wpSize = s.MaxWorkerPoolSize // Limit WorkerPool size
	}
	wp := NewWorkerPool(wpSize)
	for i := range coll.TransactionIDs {
		txID := coll.TransactionIDs[i]
		wp.Run(func() {
			transaction, err := s.getTransaction(block, collID, txID)
			if err != nil {
				s.errChan <- err
				return
			}
			mtx.Lock()
			transactions = append(transactions, transaction)
			mtx.Unlock()
		})
	}
	wp.Wait()
	return transactions
}

func (s *Subscription) getTransaction(block *flow.Block, collID flow.Identifier, txID flow.Identifier) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	tx, err := s.grpc.GetTransaction(ctx, txID)
	if err != nil {
		return nil, err
	}
	transaction := Transaction{
		Ts:               timestamppb.New(block.Timestamp),
		Id:               txID[:],
		BlockNumber:      block.Height,
		BlockID:          block.ID[:],
		CollectionID:     collID[:],
		Script:           tx.Script,
		Arguments:        tx.Arguments,
		ReferenceBlockID: tx.ReferenceBlockID[:],
		GasLimit:         tx.GasLimit,
		ProposalKey: &ProposalKey{
			Address:        tx.ProposalKey.Address[:],
			KeyIndex:       int64(tx.ProposalKey.KeyIndex),
			SequenceNumber: tx.ProposalKey.SequenceNumber,
		},
		Payer:              tx.Payer[:],
		Authorizers:        make([][]byte, 0, len(tx.Authorizers)),
		PayloadSignatures:  make([]*TransactionSignature, 0, len(tx.PayloadSignatures)),
		EnvelopeSignatures: make([]*TransactionSignature, 0, len(tx.EnvelopeSignatures)),
	}
	for _, authorizers := range tx.Authorizers {
		transaction.Authorizers = append(transaction.Authorizers, authorizers[:])
	}
	for _, sig := range tx.PayloadSignatures {
		transaction.PayloadSignatures = append(transaction.PayloadSignatures, &TransactionSignature{
			Address:     sig.Address[:],
			SignerIndex: int64(sig.SignerIndex),
			KeyIndex:    int64(sig.KeyIndex),
			Signature:   sig.Signature,
		})
	}
	for _, sig := range tx.EnvelopeSignatures {
		transaction.EnvelopeSignatures = append(transaction.EnvelopeSignatures, &TransactionSignature{
			Address:     sig.Address[:],
			SignerIndex: int64(sig.SignerIndex),
			KeyIndex:    int64(sig.KeyIndex),
			Signature:   sig.Signature,
		})
	}
	return &transaction, nil
}

func (s *Subscription) subscribeLogs(isHistorical bool, startHeight uint64, endHeight uint64) {
	var logs []*Log
	for _, event := range s.events {
		blockEvents, err := s.getEventsForHeightRange(event, startHeight, endHeight)
		if err != nil {
			s.errChan <- err
			continue
		}
		for _, blockEvent := range blockEvents {
			for _, ev := range blockEvent.Events {
				t := strings.Split(ev.Type, ".")
				log := Log{
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
				logs = append(logs, &log)
			}
		}
	}
	// Sort logs by block height, transaction index and event index
	sort.Slice(logs, func(i, j int) bool {
		log1 := logs[i]
		log2 := logs[j]
		return (log1.Height < log2.Height) || (log1.Height == log2.Height && log1.TransactionIndex < log2.TransactionIndex) || (log1.Height == log2.Height && log1.TransactionIndex == log2.TransactionIndex && log1.EventIndex < log2.EventIndex)
	})
	if isHistorical {
		for _, log := range logs {
			s.historicalLogChan <- log
		}
	} else {
		for _, log := range logs {
			s.logChan <- log
		}
	}
}

func (s *Subscription) getLatestBlockHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	block, err := s.grpc.GetLatestBlock(ctx, true)
	if err != nil {
		return 0, err
	}
	s.cache.Add(block.Height, block)
	return block.Height, nil
}

func (s *Subscription) getBlockByHeight(height uint64) (*flow.Block, error) {
	if block, ok := s.cache.Get(height); ok {
		return block.(*flow.Block), nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	block, err := s.grpc.GetBlockByHeight(ctx, height)
	if err != nil {
		return nil, err
	}
	s.cache.Add(height, block)
	return block, nil
}

func (s *Subscription) getEventsForHeightRange(eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	if endHeight > startHeight && atomic.AddUint64(&s.apiUsage, 1) <= uint64(s.MaxApiUsage) {
		ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
		defer cancel()
		return s.grpc.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
	} else { // Exceed maximun API usage. Try using different APIs.
		blockIDs := make([]flow.Identifier, 0, endHeight-startHeight+1)
		for height := startHeight; height <= endHeight; height++ {
			block, err := s.getBlockByHeight(height)
			if err != nil {
				s.errChan <- err
				continue
			}
			blockIDs = append(blockIDs, block.ID)
		}
		ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
		defer cancel()
		return s.grpc.GetEventsForBlockIDs(ctx, eventType, blockIDs)
	}
}
