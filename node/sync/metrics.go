package sync

import "github.com/go-kit/kit/metrics"

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "syncer"
)

//go:generate go run ../ops-morph/metricsgen -struct=Metrics

type Metrics struct {
	SyncedL1Height       metrics.Gauge   `metrics_name:"l1height"`
	SyncedL1MessageNonce metrics.Gauge   `metrics_name:"message_nonce"`
	SyncedL1MessageCount metrics.Counter `metrics_name:"message_count"`
}
