package node

import "github.com/go-kit/kit/metrics"

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "executor"
)

//go:generate go run ../ops-morph/metricsgen -struct=Metrics

type Metrics struct {
	Height                  metrics.Gauge
	BatchPointHeight        metrics.Gauge
	BatchIndex              metrics.Gauge
	NextL1MessageQueueIndex metrics.Gauge
}
