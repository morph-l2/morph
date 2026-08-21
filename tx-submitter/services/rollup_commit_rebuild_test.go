package services

import (
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
)

const submitterRollupCommitABI = `[
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"},
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"}],"name":"commitState","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

func TestPendingCommitRebuildSwitchesOnStoredBlobHash(t *testing.T) {
	r, _, _, rollupContract := setupTestRollup(t)
	targetABI, err := abi.JSON(strings.NewReader(submitterRollupCommitABI))
	require.NoError(t, err)
	r.abi = &targetABI

	parentHeader := make([]byte, 9) // parent batch index is zero, so the target is batch 1
	batch := &eth.RPCRollupBatch{
		Version:           1,
		ParentBatchHeader: parentHeader,
		LastBlockNumber:   10,
	}
	r.batchCacheLegacy.Set(1, batch)
	batchInput := bindings.IRollupBatchDataInput{
		Version:           uint8(batch.Version),
		ParentBatchHeader: batch.ParentBatchHeader,
		LastBlockNumber:   batch.LastBlockNumber,
	}
	tip := big.NewInt(1)
	feeCap := big.NewInt(2)
	blobFeeCap := big.NewInt(3)
	head := &ethtypes.Header{}

	commitBatchData, err := targetABI.Pack("commitBatch", batchInput)
	require.NoError(t, err)
	commitBatchTx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{ChainID: r.chainId, Nonce: 4, Gas: 100_000, Data: commitBatchData})
	rollupContract.SetStoredBlobHash(common.HexToHash("0x1234"))

	rebuilt, handled, err := r.tryRebuildRollupCommitTx(commitBatchTx, tip, feeCap, blobFeeCap, head)
	require.NoError(t, err)
	require.True(t, handled)
	require.Equal(t, targetABI.Methods["commitState"].ID, rebuilt.Data()[:4])
	require.Equal(t, uint8(ethtypes.DynamicFeeTxType), rebuilt.Type())
	require.Empty(t, rebuilt.BlobHashes())
	args, err := targetABI.Methods["commitState"].Inputs.Unpack(rebuilt.Data()[4:])
	require.NoError(t, err)
	require.Len(t, args, 1, "new commitState calldata has no signature tuple")

	commitStateData, err := targetABI.Pack("commitState", batchInput)
	require.NoError(t, err)
	commitStateTx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{ChainID: r.chainId, Nonce: 4, Gas: 100_000, Data: commitStateData})
	rollupContract.SetStoredBlobHash([32]byte{})

	rebuilt, handled, err = r.tryRebuildRollupCommitTx(commitStateTx, tip, feeCap, blobFeeCap, head)
	require.NoError(t, err)
	require.True(t, handled)
	require.Equal(t, targetABI.Methods["commitBatch"].ID, rebuilt.Data()[:4])
	args, err = targetABI.Methods["commitBatch"].Inputs.Unpack(rebuilt.Data()[4:])
	require.NoError(t, err)
	require.Len(t, args, 1, "new commitBatch calldata has no signature tuple")
}
