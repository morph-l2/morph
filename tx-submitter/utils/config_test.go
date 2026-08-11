package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigBatchConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{name: "interval only", cfg: Config{BatchBlockInterval: 10}},
		{name: "timeout only", cfg: Config{BatchTimeout: 60}},
		{name: "both", cfg: Config{BatchBlockInterval: 10, BatchTimeout: 60}},
		{name: "both zero", cfg: Config{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantErr, tt.cfg.BatchConfig().Validate() != nil)
		})
	}
}

func TestBatchConfigReportIncludesSourceAndStableHash(t *testing.T) {
	cfg := Config{
		BatchBlockInterval:           10,
		BatchTimeout:                 60,
		BatchConfigSourceBlockNumber: 123,
		BatchConfigSourceBlockHash:   "0xabc",
	}
	report := cfg.BatchConfigReport()
	require.Equal(t, uint64(10), report.BlockInterval)
	require.Equal(t, uint64(60), report.Timeout)
	require.Equal(t, uint64(123), report.SourceBlockNumber)
	require.Equal(t, "0xabc", report.SourceBlockHash)
	require.Equal(t, cfg.BatchConfig().Hash().Hex(), report.Hash)
}
