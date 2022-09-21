package flow

import (
	"strconv"

	"github.com/onflow/cadence"
)

func UFix64ToFloat64(u cadence.UFix64) float64 {
	f, _ := strconv.ParseFloat(u.String(), 64)
	return f
}
