package batch

import (
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBatchConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  BatchConfig
		wantErr bool
	}{
		{name: "interval only", config: BatchConfig{BlockInterval: 10}},
		{name: "timeout only", config: BatchConfig{Timeout: 60}},
		{name: "both", config: BatchConfig{BlockInterval: 10, Timeout: 60}},
		{name: "both zero", config: BatchConfig{}, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.config.Validate()
			require.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestBatchConfigHashIsStableAndFieldSensitive(t *testing.T) {
	config := BatchConfig{BlockInterval: 10, Timeout: 60}
	require.Equal(t, config.Hash(), config.Hash())
	require.NotEqual(t, common.Hash{}, config.Hash())
	require.NotEqual(t, config.Hash(), (BatchConfig{BlockInterval: 11, Timeout: 60}).Hash())
	require.NotEqual(t, config.Hash(), (BatchConfig{BlockInterval: 10, Timeout: 61}).Hash())
}
