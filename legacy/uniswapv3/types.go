// Code generated by connectorgen - Edit as necessary.
package uniswapv3

import (
	"blep.ai/data/connectors/source/uniswapv3/factory"
	"blep.ai/data/connectors/source/uniswapv3/pool"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.uniswapv3.0_0_0.factory_feeamountenabled":                &factory.FeeAmountEnabled{},
	"nakji.uniswapv3.0_0_0.factory_ownerchanged":                    &factory.OwnerChanged{},
	"nakji.uniswapv3.0_0_0.factory_poolcreated":                     &factory.PoolCreated{},
	"nakji.uniswapv3.0_0_0.pool_burn":                               &pool.Burn{},
	"nakji.uniswapv3.0_0_0.pool_collect":                            &pool.Collect{},
	"nakji.uniswapv3.0_0_0.pool_collectprotocol":                    &pool.CollectProtocol{},
	"nakji.uniswapv3.0_0_0.pool_flash":                              &pool.Flash{},
	"nakji.uniswapv3.0_0_0.pool_increaseobservationcardinalitynext": &pool.IncreaseObservationCardinalityNext{},
	"nakji.uniswapv3.0_0_0.pool_initialize":                         &pool.Initialize{},
	"nakji.uniswapv3.0_0_0.pool_mint":                               &pool.Mint{},
	"nakji.uniswapv3.0_0_0.pool_setfeeprotocol":                     &pool.SetFeeProtocol{},
	"nakji.uniswapv3.0_0_0.pool_swap":                               &pool.Swap{},
}
