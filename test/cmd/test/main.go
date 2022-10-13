package main

import (
	"os"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/test"

	"github.com/hugjobk/go-benchmark"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	WorkerCount       int
	TaskCount         int
	BenchmarkDuration time.Duration
	Tps               int
	LatencyStart      time.Duration
	LatencyStep       int
	ShowProcess       bool
	MsgSize           int
)

type Duration time.Duration

func (d *Duration) String() string {
	return time.Duration(*d).String()
}

func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}

func (d *Duration) Type() string {
	return "Duration"
}

func init() {
	var (
		duration     = Duration(30 * time.Second)
		latencyStart = Duration(1 * time.Microsecond)
	)

	workerCount := pflag.UintP("worker-count", "w", 1, "Number of worker goroutines")
	taskCount := pflag.UintP("task-count", "n", 0, "Number of messages sent to Kafka")
	pflag.VarP(&duration, "duration", "d", "Benchmark duration")
	tps := pflag.UintP("tps", "t", 10000, "Limit number of mesages per second sent to Kafka")
	pflag.VarP(&latencyStart, "latency-start", "l", "Latency report starting interval")
	latencyStep := pflag.UintP("latency-step", "e", 6, "Number of latency report intervals")
	showProcess := pflag.BoolP("process", "p", true, "Whether to show benchmark process")
	msgSize := pflag.UintP("msg-size", "s", 100, "Message size in byte")

	pflag.Parse()

	WorkerCount = int(*workerCount)
	TaskCount = int(*taskCount)
	BenchmarkDuration = time.Duration(duration)
	Tps = int(*tps)
	LatencyStart = time.Duration(latencyStart)
	LatencyStep = int(*latencyStep)
	ShowProcess = *showProcess
	MsgSize = int(*msgSize)
}

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.RegisterProtos(&test.Message{})

	b := benchmark.Benchmark{
		WorkerCount:  WorkerCount,
		TaskCount:    TaskCount,
		Duration:     BenchmarkDuration,
		Tps:          Tps,
		LatencyStart: LatencyStart,
		LatencyStep:  LatencyStep,
		ShowProcess:  ShowProcess,
	}
	b.Run("Benckmark Connector", func(i int) error {
		c.EventSink <- &test.Message{
			Ts:   timestamppb.Now(),
			Id:   uint64(i),
			Data: make([]byte, MsgSize),
		}
		return nil
	}).Report(os.Stdout)

	// Prevent the benchmark from being shutdown/restarted on k8s
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		for {
			time.Sleep(1 * time.Second)
		}
	}
}
