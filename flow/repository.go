package flow

import (
	"context"
	"fmt"
	"time"

	"github.com/onflow/flow-go-sdk"
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxRetryCount  = 5
	retryInterval  = 1 * time.Second
	maxGrpcMsgSize = 64 * 1024 * 1024 // 64 MB
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
	grpc *flowgrpc.Client
}

func NewRepository(host string) (Repository, error) {
	cli, err := flowgrpc.NewClient(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxGrpcMsgSize)),
	)
	if err != nil {
		return nil, err
	}
	return &repositoryImpl{cli}, nil
}

func (r *repositoryImpl) GetLatestBlock(ctx context.Context, isSealed bool) (*flow.Block, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetLatestBlock: %w", ctx.Err())
		default:
			block, err := r.grpc.GetLatestBlock(ctx, isSealed)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
					continue
				}
				return nil, fmt.Errorf("GetLatestBlock: %w", err)
			}
			return block, nil
		}
	}
}

func (r *repositoryImpl) GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetBlockByHeight: %w", ctx.Err())
		default:
			block, err := r.grpc.GetBlockByHeight(ctx, height)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
					time.Sleep(retryInterval)
					continue
				}
				return nil, fmt.Errorf("GetBlockByHeight: %w", err)
			}
			return block, nil
		}
	}
}

func (r *repositoryImpl) GetCollection(ctx context.Context, colID flow.Identifier) (*flow.Collection, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetCollection: %w", ctx.Err())
		default:
			col, err := r.grpc.GetCollection(ctx, colID)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable, codes.NotFound) {
					time.Sleep(retryInterval)
					continue
				}
				return nil, fmt.Errorf("GetCollection: %w", err)
			}
			return col, nil
		}
	}
}

func (r *repositoryImpl) GetTransaction(ctx context.Context, txID flow.Identifier) (*flow.Transaction, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetTransaction: %w", ctx.Err())
		default:
			tx, err := r.grpc.GetTransaction(ctx, txID)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
					time.Sleep(retryInterval)
					continue
				}
				return nil, fmt.Errorf("GetTransaction: %w", err)
			}
			return tx, nil
		}
	}
}

func (r *repositoryImpl) GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetEventsForBlockIDs: %w", ctx.Err())
		default:
			events, err := r.grpc.GetEventsForBlockIDs(ctx, eventType, blockIDs)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
					time.Sleep(retryInterval)
					continue
				}
				return nil, fmt.Errorf("GetEventsForBlockIDs: %w", err)
			}
			return events, nil
		}
	}
}

func (r *repositoryImpl) GetEventsForHeightRange(ctx context.Context, eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	for retryCount := 0; ; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("GetEventsForHeightRange: %w", ctx.Err())
		default:
			events, err := r.grpc.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
			if err != nil {
				if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
					time.Sleep(retryInterval)
					continue
				}
				return nil, fmt.Errorf("GetEventsForHeightRange: %w", err)
			}
			return events, nil
		}
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
