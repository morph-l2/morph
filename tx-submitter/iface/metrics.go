package iface

type IMetrics interface {
	// Reorg metrics
	IncReorgs()
	SetReorgDepth(depth float64)
}
