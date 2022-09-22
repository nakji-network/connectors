package binance

import (
	"encoding/json"
)

type MultiStreamResponse struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}
