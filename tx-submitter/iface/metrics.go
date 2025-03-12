package iface

type IMetrics interface {
	// ... existing code ...
	
	// Reorg metrics
	IncReorgs()
	SetReorgDepth(depth float64)
} 