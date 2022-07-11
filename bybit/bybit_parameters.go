package bybit

import (
	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitmex"
)

// Inspired by gocryptotrader bitmex_parameters.

type GenericRequestParams struct {
	Symbol string
}

// VerifyData verifies outgoing data sets
func (p *GenericRequestParams) VerifyData() error {
	return nil
}

// ToURLVals converts struct values to url.values and encodes it on the supplied
// path
func (p *GenericRequestParams) ToURLVals(path string) (string, error) {
	values, err := bitmex.StructValsToURLVals(p)
	if err != nil {
		return "", err
	}
	return common.EncodeURLValues(path, values), nil
}

// IsNil checks to see if any values has been set for the paramater
func (p *GenericRequestParams) IsNil() bool {
	return *p == (GenericRequestParams{})
}
