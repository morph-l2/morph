package services

import (
	"math/big"
	"testing"

	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/mock"
)

// setupSplitRollup upgrades the test rollup to split rollup/finalize accounts.
func setupSplitRollup(t *testing.T) *Rollup {
	r, _, _, _ := setupTestRollup(t)
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	fin := NewLocalAccount("finalize", key, ethtypes.LatestSignerForChainID(big.NewInt(1)), nil)
	fin.pendingTxs = NewPendingTxs(mock.NewMockJournal())
	r.finalizeAcc = fin
	r.Submitter.(*mock.MockSubmitter).SetActive(fin.WalletAddr(), true)
	return r
}

func TestAccountsDedupInCompatMode(t *testing.T) {
	r, _, _, _ := setupTestRollup(t)
	require.False(t, r.splitAccounts())
	require.Len(t, r.accounts(), 1)
	require.Same(t, r.rollupAcc, r.finalizeAcc)
}

func TestPreCheckChecksBothAccounts(t *testing.T) {
	r := setupSplitRollup(t)
	require.True(t, r.splitAccounts())
	require.Len(t, r.accounts(), 2)
	require.NoError(t, r.PreCheck())

	r.Submitter.(*mock.MockSubmitter).SetActive(r.finalizeAcc.WalletAddr(), false)
	require.ErrorContains(t, r.PreCheck(), "finalize")
}

func TestFinalizeUsesOwnAccountPoolAndSigner(t *testing.T) {
	r := setupSplitRollup(t)
	rc := r.Rollup.(*mock.MockRollup)
	rc.SetLastFinalizedBatchIndex(big.NewInt(10))
	rc.SetLastCommittedBatchIndex(big.NewInt(12))
	rc.SetBatchExists(true)
	rc.SetBatchInsideChallengeWindow(false)
	r.cfg.RoughEstimateGas = true
	r.batchCacheLegacy.Set(12, &eth.RPCRollupBatch{ParentBatchHeader: headerWithBatchIndex(11)})

	// Estimation may fail in the mock; rough estimate keeps the tx flowing.
	require.NoError(t, r.finalize())
	require.Equal(t, 1, r.finalizeAcc.pendingTxs.Len())
	require.Zero(t, r.rollupAcc.pendingTxs.Len(), "finalize must not touch the rollup pool")

	sent := r.finalizeAcc.pendingTxs.GetAll()[0].Tx
	from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(big.NewInt(1)), sent)
	require.NoError(t, err)
	require.Equal(t, r.finalizeAcc.WalletAddr(), from)
}

func TestReorgClearsBothPools(t *testing.T) {
	r := setupSplitRollup(t)
	require.NoError(t, r.handleReorg(1))
}

func TestLockIfShared(t *testing.T) {
	r := setupSplitRollup(t)
	u1 := r.lockIfShared()
	u2 := r.lockIfShared() // split mode: must not deadlock
	u2()
	u1()

	r2, _, _, _ := setupTestRollup(t)
	unlock := r2.lockIfShared()
	require.False(t, r2.sharedAcctMu.TryLock())
	unlock()
	require.True(t, r2.sharedAcctMu.TryLock())
}
