package ethereum

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lru "github.com/hashicorp/golang-lru"
)

type Log struct {
	Timestamp time.Time
	// Consensus fields:
	// address of the contract that generated the event
	Address common.Address
	// list of topics provided by the contract.
	Topics []common.Hash
	// supplied by the contract, usually ABI-encoded
	Data []byte

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber uint64
	// hash of the transaction
	TxHash common.Hash
	// index of the transaction in the block
	TxIndex uint
	// hash of the block in which the transaction was included
	BlockHash common.Hash
	// index of the log in the block
	Index uint

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool
}

type GrpcClient struct {
	grpc  *ethclient.Client
	cache *lru.Cache
}

func NewGrpcClient(ctx context.Context, url string) (*GrpcClient, error) {
	cli, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	cache, err := lru.New(100000)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{grpc: cli, cache: cache}, nil
}

func (cli *GrpcClient) GetLatestBlock(ctx context.Context) (uint64, error) {
	return cli.grpc.BlockNumber(ctx)
}

func (cli *GrpcClient) GetLogsForHeightRange(ctx context.Context, topics []string, startHeight uint64, endHeight uint64) (<-chan Log, <-chan error) {
	logChan := make(chan Log, 100)
	errChan := make(chan error, 1)

	go func() {
		defer close(logChan)
		defer close(errChan)
		addresses := make([]common.Address, 0, len(topics))
		for _, topic := range topics {
			addresses = append(addresses, common.HexToAddress(topic))
		}
		logs, err := cli.grpc.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startHeight)),
			ToBlock:   big.NewInt(int64(endHeight)),
			Addresses: addresses,
		})
		if err != nil {
			errChan <- err
		} else {
			for _, log := range logs {
				select {
				case <-ctx.Done():
					return
				default:
					ts, err := cli.getTimestamp(ctx, log.BlockNumber)
					if err != nil {
						errChan <- err
					}
					logChan <- Log{
						Timestamp:   ts,
						Address:     log.Address,
						Topics:      log.Topics,
						Data:        log.Data,
						BlockNumber: log.BlockNumber,
						TxHash:      log.TxHash,
						TxIndex:     log.TxIndex,
						BlockHash:   log.BlockHash,
						Index:       log.Index,
						Removed:     log.Removed,
					}
				}
			}
		}
	}()

	return logChan, errChan
}

func (cli *GrpcClient) getTimestamp(ctx context.Context, height uint64) (time.Time, error) {
	if t, ok := cli.cache.Get(height); ok {
		return t.(time.Time), nil
	}
	header, err := cli.grpc.HeaderByNumber(ctx, big.NewInt(int64(height)))
	if err != nil {
		return time.Time{}, err
	}
	t := blockTimeToTime(header.Time)
	cli.cache.Add(height, t)
	return t, nil
}

func blockTimeToTime(blockTime uint64) time.Time {
	unixtime := int64(blockTime * 1000)
	return time.Unix(unixtime/1000, int64(unixtime%1000*1e6))
}
