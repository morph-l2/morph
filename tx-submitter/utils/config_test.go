package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateBatchSealingConfig(t *testing.T) {
	require.Error(t, validateBatchSealingConfig(0, 0))
	require.NoError(t, validateBatchSealingConfig(1, 0))
	require.NoError(t, validateBatchSealingConfig(0, 1))
	require.NoError(t, validateBatchSealingConfig(1, 1))
}
