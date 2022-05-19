package googletrends

import (
	reflect "reflect"
	"time"

	"blep.ai/data/common"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/groovili/gogtrends"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var TopicTypes = map[string]proto.Message{
	"nakji.google.0_0_0.trend_score": &TrendScore{},
}

type GoogleTrends struct {
	*kafkautils.Producer
	BaseURL    string
	Keywords   []string
	NS         string
	TrendTopic kafkautils.Topic
	DB         *database.Database
	Table      string
	Sleep      time.Duration
	Interval   time.Duration // interval at which the loop runs
	Backoff    time.Duration
}

type Config struct {
	BaseURL    string
	Keywords   []string
	NS         string
	KP         *kafkautils.Producer
	TrendTopic kafkautils.Topic
	DB         *database.Database
	Table      string
	Sleep      time.Duration
	Interval   time.Duration // interval at which the loop runs
	Backoff    time.Duration
}

// Unmarshal gogtrends to TrendScore
func UnmarshalGogtrendsTimeline(data []*gogtrends.Timeline, tsFetched *timestamppb.Timestamp) []*TrendScore {
	res := make([]*TrendScore, 0)
	for _, d := range data {
		if d.HasData[0] {
			unixts := rowTimestampAsUnix(d)

			res = append(res, &TrendScore{
				Ts:      common.UnixToTimestampPb(unixts * 1000),
				Fetched: tsFetched,
				Value:   float64(d.Value[0]),
				// ValueNormalized is set elsewhere
			})
		}
	}

	return res
}

// TrendScore Viewer interface
func (s *TrendScore) Viewer() map[reflect.Type]database.View {
	return map[reflect.Type]database.View{
		reflect.TypeOf(TrendScore{}): {
			Src: "timescaledb",
			Cmd: `
				time_bucket('%s', "ts")        AS "ts",
				LAST ("Fetched", "ts")         AS "Fetched",
				AVG ("Value")                  AS "Value",
				AVG ("ValueNormalized")        AS "ValueNormalized"
			`,
		},
		reflect.TypeOf(common.OHLC{}): {
			Src: "timescaledb",
			Cmd: `
				time_bucket('%s', "ts")               AS "ts",
				FIRST("ValueNormalized", "ts")        AS "open",
				MAX  ("ValueNormalized")              AS "high",
				MIN  ("ValueNormalized")              AS "low",
				LAST ("ValueNormalized", "ts")        AS "close"
			`,
		},
	}
}
