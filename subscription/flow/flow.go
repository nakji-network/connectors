package flow

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	flowgrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

type GrpcClient struct {
	gprc *flowgrpc.Client
}

func NewGrpcClient(ctx context.Context, host string) (*GrpcClient, error) {
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
	return &GrpcClient{cli}, nil
}

func (cli *GrpcClient) GetLatestBlock(ctx context.Context) (uint64, error) {
	header, err := cli.gprc.GetLatestBlockHeader(ctx, true)
	if err != nil {
		return 0, err
	}
	return header.Height, nil
}

func (cli *GrpcClient) GetLogsForHeightRange(ctx context.Context, topics []string, startHeight uint64, endHeight uint64) (<-chan Log, <-chan error) {
	logChan := make(chan Log, 100)
	errChan := make(chan error, 100)

	go func() {
		defer close(logChan)
		defer close(errChan)
		for _, topic := range topics {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				blockEvents, err := cli.gprc.GetEventsForHeightRange(ctx, topic, startHeight, endHeight)
				if err != nil {
					if grpcErr, ok := err.(flowgrpc.RPCError); ok && grpcErr.GRPCStatus().Code() == codes.OutOfRange {
						return
					}
					errChan <- err
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
						logChan <- log
					}
				}
			}
		}
	}()

	return logChan, errChan
}
