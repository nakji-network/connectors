package flow_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/nakji-network/connectors/flow"

	"github.com/onflow/flow-go-sdk/access/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestSubscription(t *testing.T) {
	config := &flow.Config{
		Host:              grpc.MainnetHost,
		FromBlock:         0,
		NumBlocks:         0,
		MaxGrpcMsgSize:    64 * 1024 * 1024,
		MaxApiUsage:       20,
		MaxWorkerPoolSize: 5,
		CacheSize:         2000,
		ChannelSize:       1000,
		Timeout:           3 * time.Minute,
	}
	events := []string{
		"A.1654653399040a61.FlowToken.TokensInitialized",
		"A.1654653399040a61.FlowToken.TokensWithdrawn",
		"A.1654653399040a61.FlowToken.TokensDeposited",
		"A.1654653399040a61.FlowToken.TokensMinted",
		"A.1654653399040a61.FlowToken.TokensBurned",
		"A.1654653399040a61.FlowToken.MinterCreated",
		"A.1654653399040a61.FlowToken.BurnerCreated",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	sub, err := flow.NewSubscription(ctx, config, events)
	if err != nil {
		t.Fatal(err)
	}
	for {
		select {
		case <-sub.Done():
			return
		case block := <-sub.Blocks():
			t.Logf("Block: %s", protojson.Format(block))
		case tx := <-sub.Transactions():
			t.Logf("Transaction: %s", protojson.Format(tx))
		case log := <-sub.Logs():
			b, _ := json.MarshalIndent(log, "", "  ")
			t.Logf("Log: %s", string(b))
		case err := <-sub.Err():
			t.Error(err)
		}
	}
}
