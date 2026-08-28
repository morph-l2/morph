package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsExecutionRevertErr(t *testing.T) {
	require.False(t, IsExecutionRevertErr(nil))
	require.False(t, IsExecutionRevertErr(errors.New("connection refused")))
	require.False(t, IsExecutionRevertErr(errors.New("timeout")))
	require.True(t, IsExecutionRevertErr(errors.New("execution reverted: only active submitter allowed")))
	require.True(t, IsExecutionRevertErr(errors.New("Execution Reverted")))
	require.True(t, IsExecutionRevertErr(errors.New("always failing transaction")))
}
