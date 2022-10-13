package flow

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/onflow/flow-go-sdk"
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Repository interface {
	GetLatestBlock(ctx context.Context, isSealed bool) (*flow.Block, error)
	GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error)
	GetCollection(ctx context.Context, colID flow.Identifier) (*flow.Collection, error)
	GetTransaction(ctx context.Context, txID flow.Identifier) (*flow.Transaction, error)
	GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error)
	GetEventsForHeightRange(ctx context.Context, eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error)
}

type repositoryImpl struct {
	*Config
	grpc *flowgrpc.Client
}

func NewRepository(ctx context.Context, config *Config) (Repository, error) {
	cli, err := flowgrpc.NewClient(
		config.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxGrpcMsgSize)),
	)
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx); err != nil {
		return nil, err
	}
	return &repositoryImpl{config, cli}, nil
}

func (r *repositoryImpl) GetLatestBlock(ctx context.Context, isSealed bool) (*flow.Block, error) {
	for retry := 0; ; retry++ {
		block, err := r.grpc.GetLatestBlock(ctx, isSealed)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetLatestBlock: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetLatestBlock: %w", err)
		}
		return block, nil
	}
}

func (r *repositoryImpl) GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	for retry := 0; ; retry++ {
		block, err := r.grpc.GetBlockByHeight(ctx, height)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetBlockByHeight: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetBlockByHeight: %w", err)
		}
		return block, nil
	}
}

func (r *repositoryImpl) GetCollection(ctx context.Context, colID flow.Identifier) (*flow.Collection, error) {
	for retry := 0; ; retry++ {
		col, err := r.grpc.GetCollection(ctx, colID)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable, codes.NotFound) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetCollection: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetCollection: %w", err)
		}
		return col, nil
	}
}

func (r *repositoryImpl) GetTransaction(ctx context.Context, txID flow.Identifier) (*flow.Transaction, error) {
	for retry := 0; ; retry++ {
		tx, err := r.grpc.GetTransaction(ctx, txID)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetTransaction: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetTransaction: %w", err)
		}
		return tx, nil
	}
}

func (r *repositoryImpl) GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	for retry := 0; ; retry++ {
		events, err := r.grpc.GetEventsForBlockIDs(ctx, eventType, blockIDs)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetEventsForBlockIDs: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetEventsForBlockIDs: %w", err)
		}
		return events, nil
	}
}

func (r *repositoryImpl) GetEventsForHeightRange(ctx context.Context, eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	for retry := 0; ; retry++ {
		events, err := r.grpc.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
		if err != nil {
			if retry < r.MaxRetry && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetEventsForHeightRange: %w", ctx.Err())
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			return nil, fmt.Errorf("GetEventsForHeightRange: %w", err)
		}
		return events, nil
	}
}

func checkCode(err error, codes ...codes.Code) bool {
	if rpcErr, ok := err.(flowgrpc.RPCError); ok {
		code := rpcErr.GRPCStatus().Code()
		for i := range codes {
			if code == codes[i] {
				return true
			}
		}
	}
	return false
}
