package uniswapv2

import (
	"context"
	"fmt"
	"net"

	"blep.ai/data/connectors/source"
	"blep.ai/data/connectors/source/evm"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	source.UnimplementedConnectorServiceServer
	connector *UniswapV2
}

// SayHello implements helloworld.GreeterServer
//func (s *server) GetPairs(ctx context.Context, in *GetPairsReq) (*GetPairsRes, error) {
//for _, id := range in.Addresses {
//log.Printf("Received: %v", id)
//}
//return &GetPairsRes{
//Pairs: []*Pair{{
//Id: []byte("asdf"),
//Token0: []byte("asdf"),
//Token1: []byte("asdf"),
//}},
//}, nil
//}

// Convert pair id to symbol-symbol (eg WETH/USDT)
func (s *server) GetHumanName(ctx context.Context, in *source.GetHumanNameReq) (*source.GetHumanNameRes, error) {
	u := s.connector
	pair, ok := u.pairCache.Get(ethcommon.BytesToAddress(in.Address))
	if !ok {
		return nil, fmt.Errorf("Not Found: %s", ethcommon.BytesToAddress(in.Address).Hex())
	}

	token0 := ethcommon.BytesToAddress(pair.Token0)
	t0, ok := u.tokenCache.Get(token0)
	if !ok {
		log.Error().Str("token", token0.Hex()).Msg("Get token failed")
		t0 = &evm.ERC20{Symbol: token0.Hex()}
	}

	token1 := ethcommon.BytesToAddress(pair.Token1)
	t1, ok := u.tokenCache.Get(token1)
	if !ok {
		log.Error().Str("token", token1.Hex()).Msg("Get token failed")
		t1 = &evm.ERC20{Symbol: token1.Hex()}
	}

	return &source.GetHumanNameRes{
		Name: t0.Symbol + "-" + t1.Symbol,
	}, nil
}

// Convert many pair ids to many symbol-symbol (eg WETH/USDT)
// Missing pairs return pair address
// Missing erc20 tokens show address instead of symbol
func (s *server) GetHumanNameBulk(ctx context.Context, in *source.GetHumanNameBulkReq) (*source.GetHumanNameBulkRes, error) {
	u := s.connector

	// prepare response
	res := source.GetHumanNameBulkRes{
		Names: make([]string, len(in.Addresses)),
	}

	for i, addr := range in.Addresses {
		a := ethcommon.BytesToAddress(addr)

		pair, ok := u.pairCache.Get(a)
		if !ok {
			log.Error().
				Str("erc20", a.Hex()).
				Msg("not found in pair map]")
			res.Names[i] = a.Hex()
		}

		// get symbols from tokens
		token0 := ethcommon.BytesToAddress(pair.Token0)
		t0, ok := u.tokenCache.Get(token0)
		if !ok {
			log.Error().Str("token", token0.Hex()).Msg("Get token failed")
			t0 = &evm.ERC20{Symbol: token0.Hex()}
		}

		token1 := ethcommon.BytesToAddress(pair.Token1)
		t1, ok := u.tokenCache.Get(token1)
		if !ok {
			log.Error().Str("token", token1.Hex()).Msg("Get token failed")
			t1 = &evm.ERC20{Symbol: token1.Hex()}
		}
		res.Names[i] = t0.Symbol + "-" + t1.Symbol
	}

	return &res, nil
}

func GrpcServer(port string, u *UniswapV2) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	s := grpc.NewServer()
	source.RegisterConnectorServiceServer(s, &server{
		connector: u,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
