package subscription_test

import (
	"context"
	"testing"
	"time"

	"github.com/nakji-network/connectors/subscription"
	"github.com/nakji-network/connectors/subscription/ethereum"
	"github.com/nakji-network/connectors/subscription/flow"

	"github.com/onflow/flow-go-sdk/access/grpc"
)

func TestFlowSubscription(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	grpc, err := flow.NewGrpcClient(ctx, grpc.MainnetHost)
	if err != nil {
		t.Fatal(err)
	}
	sub := subscription.Subscription[flow.Log]{
		Grpc:      grpc,
		BlockTime: 2 * time.Second,
		Topics: []string{
			"A.1654653399040a61.FlowToken.TokensInitialized",
			"A.1654653399040a61.FlowToken.TokensWithdrawn",
			"A.1654653399040a61.FlowToken.TokensDeposited",
			"A.1654653399040a61.FlowToken.TokensMinted",
			"A.1654653399040a61.FlowToken.TokensBurned",
			"A.1654653399040a61.FlowToken.MinterCreated",
			"A.1654653399040a61.FlowToken.BurnerCreated",
		},
		FromBlock:         0,
		NumBlocks:         1000,
		LogChanSize:       1000,
		ErrChanSize:       1000,
		BackfillBatchSize: 250,
	}
	if err := sub.Subscribe(ctx); err != nil {
		t.Fatal(err)
	}
loop:
	for {
		select {
		case <-sub.Done():
			break loop
		case err := <-sub.Errs():
			t.Error(err)
		case log := <-sub.Logs():
			t.Log(log)
		}
	}
}

func TestEthereumSubscription(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	grpc, err := ethereum.NewGrpcClient(ctx, "https://hopeful-leakey:spud-chaste-trance-width-upward-harbor@nd-334-593-626.p2pify.com")
	if err != nil {
		t.Fatal(err)
	}
	sub := subscription.Subscription[ethereum.Log]{
		Grpc:      grpc,
		BlockTime: 10 * time.Second,
		Topics: []string{
			"0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9", // aave
		},
		FromBlock:         0,
		NumBlocks:         1000,
		LogChanSize:       1000,
		ErrChanSize:       1000,
		BackfillBatchSize: 250,
	}
	if err := sub.Subscribe(ctx); err != nil {
		t.Fatal(err)
	}
loop:
	for {
		select {
		case <-sub.Done():
			break loop
		case err := <-sub.Errs():
			t.Error(err)
		case log := <-sub.Logs():
			t.Log(log)
		}
	}
}
