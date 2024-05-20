package oracle

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestSubmitRecord(t *testing.T) {
	o := testNewOracleClient(t)
	o.submitRecord()

}

func TestLastFinalizedBatchIndex(t *testing.T) {
	o := testNewOracleClient(t)
	lastFinalized, err := o.rollup.LastFinalizedBatchIndex(nil)
	require.NoError(t, err)
	fmt.Println("lastFinalized:", lastFinalized)
	nextBatchSubmissionIndex, err := o.record.NextBatchSubmissionIndex(nil)
	require.NoError(t, err)
	fmt.Println("nextBatchSubmissionIndex:", nextBatchSubmissionIndex)
	blockNumber, err := o.GetStartBlock(big.NewInt(38))
	require.NoError(t, err)
	fmt.Println("blockNumber:", blockNumber)
	bs, err := o.record.BatchSubmissions(nil, nextBatchSubmissionIndex.Sub(nextBatchSubmissionIndex, big.NewInt(1)))
	require.NoError(t, err)
	fmt.Println("bs=", bs)

	lastBatchIndex, err := o.rollup.LastCommittedBatchIndex(nil)
	require.NoError(t, err)
	fmt.Println("latestBatchIndex", lastBatchIndex)

}

func TestFetchRollupLog(t *testing.T) {
	o := testNewOracleClient(t)
	//rLogs, err := o.fetchRollupLog(o.ctx, 1, 200)
	//fmt.Println("lastFinalized:", len(rLogs))
	//require.NoError(t, err)
	//fmt.Println(o.cfg.RollupAddr)
	epoch, err := o.gov.RollupEpoch(nil)
	require.NoError(t, err)
	fmt.Println("epoch:", epoch)
}

func TestOracle_GetNextBatchSubmissionIndex(t *testing.T) {
	o := testNewOracleClient(t)
	err := o.submitRecord()
	require.NoError(t, err)
}
