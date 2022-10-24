package woofi

import (
	"context"
	"fmt"
	"sync"

	"github.com/nakji-network/connector"
	"github.com/rs/zerolog/log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Log struct {
	types.Log
	Network string
}

type SubscriptionGroup struct {
	logChan chan Log
	errChan chan error
	done    chan struct{}
	subs    map[string]connector.ISubscription
}

func NewSubscriptionGroup() *SubscriptionGroup {
	return &SubscriptionGroup{
		logChan: make(chan Log),
		errChan: make(chan error),
		done:    make(chan struct{}),
		subs:    make(map[string]connector.ISubscription),
	}
}

func (g *SubscriptionGroup) AddSubscription(ctx context.Context, c *connector.Connector, network string, addresses []ethcommon.Address, fromBlock uint64, numBlocks uint64) error {
	sub, err := connector.NewSubscription(ctx, c, network, addresses, fromBlock, numBlocks)
	if err != nil {
		return err
	}
	g.subs[network] = sub
	return nil
}

func (g *SubscriptionGroup) Init() {
	var wg sync.WaitGroup
	wg.Add(len(g.subs))
	for network, sub := range g.subs {
		go func(network string, sub connector.ISubscription) {
			defer wg.Done()
			for {
				select {
				case vLog := <-sub.Logs():
					g.logChan <- Log{
						Log:     vLog,
						Network: network,
					}
				case err := <-sub.Err():
					g.errChan <- err
				case <-sub.Done():
					log.Info().Str("network", network).Msg("subscription closed")
					return
				}
			}
		}(network, sub)
	}
	go func() {
		wg.Wait()
		close(g.done)
	}()
}

func (g *SubscriptionGroup) Logs() <-chan Log {
	return g.logChan
}

func (g *SubscriptionGroup) Err() <-chan error {
	return g.errChan
}

func (g *SubscriptionGroup) Done() <-chan struct{} {
	return g.done
}

func (g *SubscriptionGroup) GetBlockTime(ctx context.Context, vLog Log) (uint64, error) {
	sub, ok := g.subs[vLog.Network]
	if !ok {
		return 0, fmt.Errorf("unknown network: %s", vLog.Network)
	}
	return sub.GetBlockTime(ctx, vLog.Log)
}
