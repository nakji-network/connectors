// Multilevel uniswap pair cache utilizing lru.ARCCache
//
// Levels:
// 1: cache
// 2: database
// 3: ethereum network
//
// Maps ethcommon.Address to common.Pair (aka proto.Message)
package uniswapv2

import (
	reflect "reflect"
	"time"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/tokencache"

	"blep.ai/data/common"
	"blep.ai/data/connectors/source/uniswapv2/pair"
	"blep.ai/data/database"
	ethcommon "github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const pairTable = "pair"
const (
	// TODO: do not hardcode
	contractTimeout = 15000 * time.Millisecond
	maxConcurrency  = 150
	EMPTY_ADDRESS   = "0x0000000000000000000000000000000000000000"
)

type PairCache struct {
	cache      *lru.ARCCache
	db         *database.Database
	clientPool *ethclient.ClientPool
	namespace  string // pair.Ns
	app        string // pair.App
	tokenCache *tokencache.TokenCache
}

func NewPairCache(blockchain, app string, db *database.Database, clientPool *ethclient.ClientPool, size int, tokenCache *tokencache.TokenCache) (*PairCache, error) {
	cache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}
	pairCache := PairCache{cache, db, clientPool, blockchain, app, tokenCache}

	return &pairCache, err
}

func (c *PairCache) Add(a ethcommon.Address, pair *common.Pair) {
	c.cache.Add(a, pair)
	pair.Insert(c.db)
}

// Get pair from cache
// Same usage as lru.ARCCache.Get
func (c *PairCache) Get(a ethcommon.Address) (res *common.Pair, ok bool) {
	// Level 1. Cache
	if t, ok := c.cache.Get(a); ok {
		return t.(*common.Pair), true
	}

	// Level 2: Database
	if res, ok = c.getFromDatabase(a); ok {
		c.cache.Add(a, res)
		return
	}

	// Level 3: Ethereum network
	log.Error().Msg("Should not have needed to get from network")
	if res, ok = c.getFromNetwork(a); ok {
		res.Ns = c.namespace
		res.App = c.app
		c.Add(a, res)
		// Write to db
		err := c.db.InsertProto2(pairTable, res)
		if err != nil {
			log.Error().Err(err).Msg("Unable to write result to db")
		}
		return
	}

	return
}

func (c *PairCache) getFromDatabase(a ethcommon.Address) (res *common.Pair, ok bool) {
	ts_, err := c.db.SelectByID(pairTable, c.namespace, a.Bytes(), reflect.TypeOf(common.Pair{}))
	if err != nil {
		log.Fatal().Err(err).Msg("db error")
		return nil, false
	}

	ts := ts_.([]common.Pair)
	if len(ts) == 1 {
		log.Debug().Msg("found in db")
		return &ts[0], true
	}

	log.Debug().Msg("not found in db")
	return nil, false
}

// 2 calls to Eth RPC
func (c *PairCache) getFromNetwork(a ethcommon.Address) (res *common.Pair, ok bool) {
	if a.Hex() == EMPTY_ADDRESS {
		log.Warn().Msg("empty address received")
		return nil, false
	}

	contract, err := pair.NewIUniswapV2PairCaller(a, c.clientPool)
	if err != nil {
		log.Error().Err(err).Msg("contract error")
		return nil, false
	}

	t0, err := contract.Token0(nil)
	if err != nil || t0.Hex() == EMPTY_ADDRESS {
		if t0.Hex() == EMPTY_ADDRESS {
			log.Warn().Msg("empty token0 address")
		} else {
			log.Error().Err(err).Msg("token0 error")
		}
		return nil, false
	}

	t1, err := contract.Token1(nil)
	if err != nil || t1.Hex() == EMPTY_ADDRESS {
		if t1.Hex() == EMPTY_ADDRESS {
			log.Warn().Msg("empty token1 address")
		} else {
			log.Error().Err(err).Msg("token1 error")
		}
		return nil, false
	}
	return &common.Pair{
		Ns:     c.namespace,
		Id:     a.Bytes(),
		Added:  timestamppb.Now(),
		App:    c.app,
		I:      -1,
		Token0: t0.Bytes(),
		Token1: t1.Bytes(),
	}, true
}
