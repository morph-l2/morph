package services

import (
	"errors"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/mock"
)

// A submitter can be removed, slashed, priced out by a raised minimum stake, or
// start withdrawing at any time after startup. Every submitting method carries
// onlyActiveSubmitter on L1, so the loops must re-check instead of trusting the
// startup PreCheck: submitting while inactive mines a reverting tx, and a
// reverting blob tx is still charged its full blob fee.
func TestRollupSkipsWhenSubmitterBecameInactive(t *testing.T) {
	r, _, _, _ := setupTestRollup(t)
	r.Submitter.(*mock.MockSubmitter).SetActive(r.WalletAddr(), false)

	require.ErrorContains(t, r.rollup(), "not an active submitter")
	require.Zero(t, r.pendingTxs.Len())
}

func TestFinalizeSkipsWhenSubmitterBecameInactive(t *testing.T) {
	r, _, _, rollupContract := setupTestRollup(t)
	rollupContract.SetLastFinalizedBatchIndex(big.NewInt(10))
	rollupContract.SetLastCommittedBatchIndex(big.NewInt(12))
	r.Submitter.(*mock.MockSubmitter).SetActive(r.WalletAddr(), false)

	require.ErrorContains(t, r.finalize(), "not an active submitter")
	require.Zero(t, r.pendingTxs.Len())
}

// An RPC failure on the activity probe must also stop submission: assuming the
// wallet is still eligible is what burns funds.
func TestRollupSkipsWhenActivityProbeFails(t *testing.T) {
	r, _, _, _ := setupTestRollup(t)
	r.Submitter.(*mock.MockSubmitter).SetError(errors.New("rpc unavailable"))

	require.ErrorContains(t, r.rollup(), "check Submitter activity")
	require.Zero(t, r.pendingTxs.Len())
}

// RoughEstimateGas exists to survive a flaky eth_estimateGas, not to paper over a
// contract revert. Guessing a gas limit past a revert submits a tx that can only
// fail, so the revert must abort submission even with the flag on.
func TestFinalizeDoesNotRoughEstimatePastRevert(t *testing.T) {
	r, l1Mock, _, rollupContract := setupTestRollup(t)
	r.cfg.RoughEstimateGas = true
	rollupContract.SetLastFinalizedBatchIndex(big.NewInt(10))
	rollupContract.SetLastCommittedBatchIndex(big.NewInt(12))
	rollupContract.SetBatchExists(true)
	rollupContract.SetBatchInsideChallengeWindow(false)
	r.batchCacheLegacy.Set(12, &eth.RPCRollupBatch{
		ParentBatchHeader: hexutil.Bytes{0x01, 0x02},
	})

	l1Mock.EstimateGasErr = errors.New("execution reverted: only active submitter allowed")
	require.ErrorContains(t, r.finalize(), "estimate finalize gas error")
	require.Zero(t, r.pendingTxs.Len())

	// A transport-level failure still gets the rough fallback, so the flag keeps
	// doing what it was added for.
	l1Mock.EstimateGasErr = errors.New("connection refused")
	require.NoError(t, r.finalize())
	require.Equal(t, 1, r.pendingTxs.Len())
}
