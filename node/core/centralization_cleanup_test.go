package node

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/upgrade"

	nodesync "morph-l2/node/sync"
)

func TestParseStaticValidatorTmKeys(t *testing.T) {
	keyA := strings.Repeat("11", tmKeySize)
	keyB := strings.Repeat("22", tmKeySize)

	keysAB, hashAB, err := parseStaticValidatorTmKeys([]string{keyB, "0x" + keyA})
	require.NoError(t, err)
	require.Len(t, keysAB, 2)
	require.True(t, bytes.Compare(keysAB[0], keysAB[1]) < 0, "keys must be canonicalized lexicographically")

	keysBA, hashBA, err := parseStaticValidatorTmKeys([]string{keyA + "," + keyB})
	require.NoError(t, err)
	require.Equal(t, keysAB, keysBA)
	require.Equal(t, hashAB, hashBA, "input order must not change the network key-set hash")

	for _, tc := range []struct {
		name   string
		values []string
	}{
		{name: "empty"},
		{name: "short", values: []string{"01"}},
		{name: "invalid hex", values: []string{strings.Repeat("zz", tmKeySize)}},
		{name: "duplicate", values: []string{keyA, "0x" + keyA}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := parseStaticValidatorTmKeys(tc.values)
			require.Error(t, err)
		})
	}
}

func TestValidatorSetAtHeightUpgradeBoundary(t *testing.T) {
	originalHeight := upgrade.UpgradeBlockHeight()
	upgrade.SetUpgradeBlockHeight(100)
	t.Cleanup(func() { upgrade.SetUpgradeBlockHeight(originalHeight) })

	legacyKeys := [][]byte{bytes.Repeat([]byte{0x11}, tmKeySize)}
	staticKeys := [][]byte{bytes.Repeat([]byte{0x22}, tmKeySize)}
	legacyCalls := 0
	executor := &Executor{
		staticValidatorTmKeys: staticKeys,
		legacyValidatorSetAtHeightFunc: func(int64) ([][]byte, error) {
			legacyCalls++
			return cloneValidatorKeys(legacyKeys), nil
		},
	}

	for _, height := range []uint64{99, 100} {
		got, err := executor.validatorSetAtHeight(height)
		require.NoError(t, err)
		require.Equal(t, legacyKeys, got)
	}
	got, err := executor.validatorSetAtHeight(101)
	require.NoError(t, err)
	require.Equal(t, staticKeys, got)
	require.Equal(t, 2, legacyCalls, "H+1 must not call either historical contract")

	got[0][0] ^= 0xff
	require.Equal(t, byte(0x22), executor.staticValidatorTmKeys[0][0], "callers must not mutate the configured set")
}

func TestEnsureSyncerStartedIndependentOfValidatorMembership(t *testing.T) {
	created := 0
	fakeSyncer := nodesync.NewFakeSyncer(nil)
	executor := &Executor{
		newSyncerFunc: func() (*nodesync.Syncer, error) {
			created++
			return fakeSyncer, nil
		},
	}

	require.NoError(t, executor.ensureSyncerStarted())
	require.NoError(t, executor.ensureSyncerStarted())
	require.Equal(t, 1, created)
	require.Same(t, fakeSyncer, executor.l1MsgReader)
}

func TestEnsureSyncerStartedKeepsFactoryError(t *testing.T) {
	created := 0
	wantErr := errors.New("factory failed")
	executor := &Executor{
		newSyncerFunc: func() (*nodesync.Syncer, error) {
			created++
			return nil, wantErr
		},
	}

	require.ErrorIs(t, executor.ensureSyncerStarted(), wantErr)
	require.ErrorIs(t, executor.ensureSyncerStarted(), wantErr)
	require.Equal(t, 1, created)
	require.Nil(t, executor.l1MsgReader)
}

func TestAllValidatorEntryPointsUseUnifiedResolver(t *testing.T) {
	source, err := os.ReadFile("executor.go")
	require.NoError(t, err)
	require.Equal(t, 3, strings.Count(string(source), ".validatorSetAtHeight("),
		"NewExecutor and both DeliverBlock branches must use the H/H+1 resolver")
}
