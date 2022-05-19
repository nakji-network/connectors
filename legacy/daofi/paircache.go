// Multilevel uniswap pair cache utilizing lru.ARCCache
//
// Levels:
// 1: cache
// 2: database
// 3: ethereum network
//
// Maps ethcommon.Address to Pair (aka proto.Message)
package daofi

import (
	"context"
	reflect "reflect"
	"time"

	"blep.ai/data/common"
	. "blep.ai/data/connectors/source/daofi/pair"
	"blep.ai/data/database"
	"blep.ai/data/tokencache"
	sq "github.com/Masterminds/squirrel"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const pairTable = "pair"
const (
	// TODO: do not hardcode
	contractTimeout = 5000 * time.Millisecond
	maxConcurrency  = 150
)

type PairCache struct {
	cache      *lru.ARCCache
	db         *database.Database
	ethconn    *ethclient.Client
	namespace  string // pair.Ns
	app        string // pair.App
	tokenCache *tokencache.TokenCache
}

func NewPairCache(blockchain, app string, db *database.Database, ethconn *ethclient.Client, size int, tokenCache *tokencache.TokenCache) (*PairCache, error) {
	cache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}

	pairCache := PairCache{cache, db, ethconn, blockchain, app, tokenCache}

	return &pairCache, err
}

// TODO: Get from database and convert to address
func (c *PairCache) AllPairAddresses() ([]ethcommon.Address, error) {
	var sql, args, _ = sq.
		Select("id").
		From(pairTable).
		Where(sq.And{
			sq.Eq{"ns": c.namespace},
			sq.Eq{"app": c.app},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := c.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]ethcommon.Address, 0)

	for rows.Next() {
		var addr ethcommon.Address
		err = rows.Scan(&addr)
		if err != nil {
			return nil, err
		}
		res = append(res, addr)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return res, nil
}

func (c *PairCache) Add(a ethcommon.Address, pair *common.Pair) {
	c.cache.Add(a, pair)
	err := c.db.InsertProto2(pairTable, pair)
	if err != nil {
		log.Error().Err(err).
			Hex("pair", pair.Id).
			Hex("token0", pair.Token0).
			Hex("token1", pair.Token1).
			Msg("Unable to write result to db")
	}
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
	ts_, err := c.db.SelectByID(pairTable, c.namespace, a.Bytes(), reflect.TypeOf(Pair{}))
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
	contract, err := NewIDAOfiV1PairCaller(a, c.ethconn)
	if err != nil {
		log.Error().Err(err).Msg("contract error")
		return nil, false
	}

	t0, err := contract.BaseToken(nil)
	if err != nil {
		log.Error().Err(err).Msg("contract error")
		return nil, false
	}
	t1, err := contract.QuoteToken(nil)
	if err != nil {
		log.Error().Err(err).Msg("contract error")
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
