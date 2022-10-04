package flow

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/onflow/flow-go-sdk"
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxRetryCount  = 6
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
		block, err := r.grpc.GetLatestBlock(ctx, isSealed)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetLatestBlock: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetLatestBlock: retry=%d after %s", retryCount+1, retryAfter)
					continue
				}
			}
			return nil, fmt.Errorf("GetLatestBlock: %w", err)
		}
		return block, nil
	}
}

func (r *repositoryImpl) GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	for retryCount := 0; ; retryCount++ {
		block, err := r.grpc.GetBlockByHeight(ctx, height)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetBlockByHeight: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetBlockByHeight: retry=%d after %s", retryCount+1, retryAfter)
					continue
				}
			}
			return nil, fmt.Errorf("GetBlockByHeight: %w", err)
		}
		return block, nil
	}
}

func (r *repositoryImpl) GetCollection(ctx context.Context, colID flow.Identifier) (*flow.Collection, error) {
	for retryCount := 0; ; retryCount++ {
		col, err := r.grpc.GetCollection(ctx, colID)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable, codes.NotFound) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetCollection: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetCollection: retry=%d after %s", retryCount+1, retryAfter)
					continue
				}
			}
			return nil, fmt.Errorf("GetCollection: %w", err)
		}
		return col, nil
	}
}

func (r *repositoryImpl) GetTransaction(ctx context.Context, txID flow.Identifier) (*flow.Transaction, error) {
	for retryCount := 0; ; retryCount++ {
		tx, err := r.grpc.GetTransaction(ctx, txID)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetTransaction: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetTransaction: retry=%d after %s", retryCount+1, retryAfter)
					continue
				}
			}
			return nil, fmt.Errorf("GetTransaction: %w", err)
		}
		return tx, nil
	}
}

func (r *repositoryImpl) GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	for retryCount := 0; ; retryCount++ {
		events, err := r.grpc.GetEventsForBlockIDs(ctx, eventType, blockIDs)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetEventsForBlockIDs: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetEventsForBlockIDs: retry=%d after %s", retryCount+1, retryAfter)
					continue
				}
			}
			return nil, fmt.Errorf("GetEventsForBlockIDs: %w", err)
		}
		return events, nil
	}
}

func (r *repositoryImpl) GetEventsForHeightRange(ctx context.Context, eventType string, startHeight uint64, endHeight uint64) ([]flow.BlockEvents, error) {
	for retryCount := 0; ; retryCount++ {
		events, err := r.grpc.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
		if err != nil {
			if retryCount < maxRetryCount && checkCode(err, codes.ResourceExhausted, codes.Unavailable) {
				backoff := float64(int(1) << retryCount)
				backoff += backoff * (0.1 * rand.Float64())
				retryAfter := time.Duration(backoff) * time.Second
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("GetEventsForHeightRange: %w", ctx.Err())
				case <-time.After(retryAfter):
					log.Info().Err(err).Msgf("GetEventsForHeightRange: retry=%d after %s", retryCount+1, retryAfter)
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
