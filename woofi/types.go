package woofi

import (
	"github.com/nakji-network/connectors/woofi/WooPP"
	"github.com/nakji-network/connectors/woofi/WooPPV1"
	"github.com/nakji-network/connectors/woofi/WooPPV2"
	"github.com/nakji-network/connectors/woofi/WooRouterV1"
	"github.com/nakji-network/connectors/woofi/WooRouterV2"
)

var ABIs = map[string]string{
	"WooPP":       WooPP.WooPPABI,
	"WooPPV1":     WooPPV1.WooPPV1ABI,
	"WooPPV2":     WooPPV2.WooPPV2ABI,
	"WooRouterV1": WooRouterV1.WooRouterV1ABI,
	"WooRouterV2": WooRouterV2.WooRouterV2ABI,
}
