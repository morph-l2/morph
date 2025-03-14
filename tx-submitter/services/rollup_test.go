package services

import (
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/utils"
)

func TestSendTx(t *testing.T) {
	// Create a new dynamic fee transaction
	_tx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     0,
		GasTipCap: big.NewInt(1000000000),
		GasFeeCap: big.NewInt(1000000000),
		Gas:       1000000,
		To:        &common.Address{},
		Value:     big.NewInt(0),
		Data:      []byte{},
	}
	// new tx
	tx := types.NewTx(_tx)

	// required err
	err := sendTx(nil, 1, tx)
	require.ErrorContains(t, err, utils.ErrExceedFeeLimit.Error())

	// blob tx
	blobTx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(big.NewInt(1)),
		Nonce:      0,
		GasTipCap:  uint256.MustFromBig(big.NewInt(1000000000)),
		GasFeeCap:  uint256.MustFromBig(big.NewInt(1000000000)),
		Gas:        1000000,
		To:         common.Address{},
		Data:       []byte{},
		BlobFeeCap: uint256.MustFromBig(big.NewInt(1000000000)),
		BlobHashes: []common.Hash{
			{},
		},
	},
	)
	err = sendTx(nil, 1, blobTx)
	require.ErrorContains(t, err, utils.ErrExceedFeeLimit.Error())
}

func TestGetGasTipAndCap(t *testing.T) {
	initTip := big.NewInt(1e9)
	baseFee := big.NewInt(1e9)
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee: baseFee,
		},
	)

	r, l1Mock, _, _ := setupTestRollup(t)
	l1Mock.TipCap = initTip
	l1Mock.Block = block

	tip, feecap, blobfee, err := r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, initTip, tip)

	// Test with different TipFeeBump
	r, l1Mock, _, _ = setupTestRollup(t)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	r.cfg.TipFeeBump = 200

	tip, feecap, blobfee, err = r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, tip, initTip.Mul(initTip, big.NewInt(2)))

	// Test with base fee too high
	r, l1Mock, _, _ = setupTestRollup(t)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	r.cfg.MaxBaseFee = baseFee.Uint64() - 1

	_, _, _, err = r.GetGasTipAndCap()
	require.ErrorContains(t, err, "base fee is too high")

	// Test with tip too high
	r, l1Mock, _, _ = setupTestRollup(t)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	r.cfg.MaxTip = initTip.Uint64() - 1

	_, _, _, err = r.GetGasTipAndCap()
	require.ErrorContains(t, err, "tip is too high")
}

func TestReSubmitTx(t *testing.T) {
	marketTip := big.NewInt(3e9) // 3 Gwei market tip
	baseFee := big.NewInt(2e9)   // 2 Gwei base fee
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee: baseFee,
		},
	)

	r, l1Mock, _, _ := setupTestRollup(t)
	l1Mock.TipCap = marketTip
	l1Mock.Block = block
	r.cfg.MaxTip = 10e12
	r.cfg.MaxBaseFee = 100e9
	r.cfg.MinTip = 1e9
	r.cfg.TipFeeBump = 0 // no bump for replace mode

	// Test nil tx
	_, err := r.ReSubmitTx(false, nil)
	require.ErrorContains(t, err, "nil tx")

	t.Run("DynamicFeeTx", func(t *testing.T) {
		oldDynamicTx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   big.NewInt(1),
			Nonce:     1,
			GasTipCap: big.NewInt(2e9),  // 2 Gwei
			GasFeeCap: big.NewInt(10e9), // 10 Gwei
			Gas:       100000,
			To:        &common.Address{},
			Value:     big.NewInt(0),
			Data:      []byte{1, 2, 3, 4},
		})

		// Test Replace Mode
		t.Run("Replace", func(t *testing.T) {
			newTx, err := r.ReSubmitTx(false, oldDynamicTx)
			require.NoError(t, err)
			require.NotNil(t, newTx)

			// Verify fields preserved
			require.Equal(t, oldDynamicTx.Nonce(), newTx.Nonce())
			require.Equal(t, oldDynamicTx.Gas(), newTx.Gas())
			require.Equal(t, oldDynamicTx.Data(), newTx.Data())
			require.Equal(t, oldDynamicTx.Value(), newTx.Value())

			// Verify fees are at least 1.1x of original
			originalTip := oldDynamicTx.GasTipCap()
			newTip := newTx.GasTipCap()
			expectedTip := new(big.Int).Mul(originalTip, big.NewInt(110))
			expectedTip = expectedTip.Div(expectedTip, big.NewInt(100))
			require.True(t, newTip.Cmp(expectedTip) >= 0, "new tip should be at least 1.1x of original")

			originalFeeCap := oldDynamicTx.GasFeeCap()
			newFeeCap := newTx.GasFeeCap()
			expectedFeeCap := new(big.Int).Mul(originalFeeCap, big.NewInt(110))
			expectedFeeCap = expectedFeeCap.Div(expectedFeeCap, big.NewInt(100))
			require.True(t, newFeeCap.Cmp(expectedFeeCap) >= 0, "new fee cap should be at least 1.1x of original")
		})

		// Test Resubmit Mode
		t.Run("Resubmit", func(t *testing.T) {
			newTx, err := r.ReSubmitTx(true, oldDynamicTx)
			require.NoError(t, err)
			require.NotNil(t, newTx)

			// Verify fields preserved
			require.Equal(t, oldDynamicTx.Nonce(), newTx.Nonce())
			require.Equal(t, oldDynamicTx.Gas(), newTx.Gas())
			require.Equal(t, oldDynamicTx.Data(), newTx.Data())
			require.Equal(t, oldDynamicTx.Value(), newTx.Value())
			require.Equal(t, oldDynamicTx.ChainId(), newTx.ChainId())

			// Verify fees are market prices
			require.Equal(t, marketTip.Uint64(), newTx.GasTipCap().Uint64(), "new tip should be market price")
			require.True(t, newTx.GasFeeCap().Cmp(baseFee) > 0, "new fee cap should be higher than base fee")
		})
	})

	t.Run("BlobTx", func(t *testing.T) {
		oldBlobTx := types.NewTx(&types.BlobTx{
			ChainID:    uint256.MustFromBig(big.NewInt(1)),
			Nonce:      2,
			GasTipCap:  uint256.MustFromBig(big.NewInt(2e9)),
			GasFeeCap:  uint256.MustFromBig(big.NewInt(10e9)),
			Gas:        200000,
			To:         common.Address{},
			Value:      uint256.NewInt(0),
			Data:       []byte{1, 2, 3, 4},
			BlobFeeCap: uint256.MustFromBig(big.NewInt(5e9)),
			BlobHashes: []common.Hash{{1}},
			Sidecar: &types.BlobTxSidecar{
				Blobs:       []kzg4844.Blob{{1}},
				Commitments: []kzg4844.Commitment{{1}},
				Proofs:      []kzg4844.Proof{{1}},
			},
		})

		// Test Replace Mode
		t.Run("Replace", func(t *testing.T) {
			newTx, err := r.ReSubmitTx(false, oldBlobTx)
			require.NoError(t, err)
			require.NotNil(t, newTx)

			// Verify fields preserved
			require.Equal(t, oldBlobTx.Nonce(), newTx.Nonce())
			require.Equal(t, oldBlobTx.Gas(), newTx.Gas())
			require.Equal(t, oldBlobTx.Data(), newTx.Data())
			require.Equal(t, oldBlobTx.Value(), newTx.Value())
			require.Equal(t, len(oldBlobTx.BlobHashes()), len(newTx.BlobHashes()))

			// Verify fees are at least 2x of original
			originalTip := oldBlobTx.GasTipCap()
			newTip := newTx.GasTipCap()
			expectedTip := new(big.Int).Mul(originalTip, big.NewInt(200))
			expectedTip = expectedTip.Div(expectedTip, big.NewInt(100))
			require.True(t, newTip.Cmp(expectedTip) >= 0, "new tip should be at least 2x of original")

			originalFeeCap := oldBlobTx.GasFeeCap()
			newFeeCap := newTx.GasFeeCap()
			expectedFeeCap := new(big.Int).Mul(originalFeeCap, big.NewInt(200))
			expectedFeeCap = expectedFeeCap.Div(expectedFeeCap, big.NewInt(100))
			require.True(t, newFeeCap.Cmp(expectedFeeCap) >= 0, "new fee cap should be at least 2x of original")

			originalBlobFeeCap := oldBlobTx.BlobGasFeeCap()
			newBlobFeeCap := newTx.BlobGasFeeCap()
			expectedBlobFeeCap := new(big.Int).Mul(originalBlobFeeCap, big.NewInt(200))
			expectedBlobFeeCap = expectedBlobFeeCap.Div(expectedBlobFeeCap, big.NewInt(100))
			require.True(t, newBlobFeeCap.Cmp(expectedBlobFeeCap) >= 0, "new blob fee cap should be at least 2x of original")
		})

		// Test Resubmit Mode
		t.Run("Resubmit", func(t *testing.T) {
			newTx, err := r.ReSubmitTx(true, oldBlobTx)
			require.NoError(t, err)
			require.NotNil(t, newTx)

			// Verify fields preserved
			require.Equal(t, oldBlobTx.Nonce(), newTx.Nonce())
			require.Equal(t, oldBlobTx.Gas(), newTx.Gas())
			require.Equal(t, oldBlobTx.Data(), newTx.Data())
			require.Equal(t, oldBlobTx.Value(), newTx.Value())
			require.Equal(t, len(oldBlobTx.BlobHashes()), len(newTx.BlobHashes()))

			// Verify fees are market prices
			require.Equal(t, marketTip.Uint64(), newTx.GasTipCap().Uint64(), "new tip should be market price")
			require.True(t, newTx.GasFeeCap().Cmp(baseFee) > 0, "new fee cap should be higher than base fee")
			require.NotNil(t, newTx.BlobGasFeeCap(), "new blob tx should have blob fee cap")
		})
	})
}

func TestCancelTx(t *testing.T) {
	// Setup mock L1 client
	initTip := big.NewInt(1e9)
	baseFee := big.NewInt(1e9)
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee: baseFee,
		},
	)

	r, l1Mock, _, _ := setupTestRollup(t)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	r.cfg.MaxTip = 10e12
	r.cfg.MaxBaseFee = 100e9
	r.cfg.MinTip = 1e9
	r.cfg.TipFeeBump = 120 // 20% bump

	// Test 1: Cancel nil transaction
	_, err := r.CancelTx(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil tx")

	// Test 2: Cancel DynamicFeeTx
	originalDynamicTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		GasTipCap: big.NewInt(2e9),  // 2 Gwei
		GasFeeCap: big.NewInt(10e9), // 10 Gwei
		Gas:       100000,
		To:        &common.Address{},
		Value:     big.NewInt(0),
		Data:      []byte{1, 2, 3, 4}, // Some non-empty data
	})

	cancelTx, err := r.CancelTx(originalDynamicTx)
	require.NoError(t, err)
	require.NotNil(t, cancelTx)

	// Verify the cancel transaction
	require.Equal(t, originalDynamicTx.Nonce(), cancelTx.Nonce())
	require.Equal(t, originalDynamicTx.Gas(), cancelTx.Gas())
	require.Equal(t, 0, len(cancelTx.Data()))

	// Verify fee multipliers for DynamicFeeTx (1.2x)
	originalTip := originalDynamicTx.GasTipCap()
	cancelTip := cancelTx.GasTipCap()
	expectedTip := new(big.Int).Mul(originalTip, big.NewInt(110))
	expectedTip = expectedTip.Div(expectedTip, big.NewInt(100))
	require.True(t, cancelTip.Cmp(expectedTip) >= 0, "cancel tx tip should be at least 1.1x of original")

	originalFeeCap := originalDynamicTx.GasFeeCap()
	cancelFeeCap := cancelTx.GasFeeCap()
	expectedFeeCap := new(big.Int).Mul(originalFeeCap, big.NewInt(110))
	expectedFeeCap = expectedFeeCap.Div(expectedFeeCap, big.NewInt(100))
	require.True(t, cancelFeeCap.Cmp(expectedFeeCap) >= 0, "cancel tx fee cap should be at least 1.1x of original")

	// Test 3: Cancel BlobTx
	blobTx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(big.NewInt(1)),
		Nonce:      2,
		GasTipCap:  uint256.MustFromBig(big.NewInt(2e9)),
		GasFeeCap:  uint256.MustFromBig(big.NewInt(10e9)),
		Gas:        200000,
		To:         common.Address{},
		Value:      uint256.NewInt(0),
		Data:       []byte{1, 2, 3, 4},
		BlobFeeCap: uint256.MustFromBig(big.NewInt(5e9)),
		BlobHashes: []common.Hash{{1}},
		Sidecar: &types.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{{1}},
			Commitments: []kzg4844.Commitment{{1}},
			Proofs:      []kzg4844.Proof{{1}},
		},
	})

	cancelBlobTx, err := r.CancelTx(blobTx)
	require.NoError(t, err)
	require.NotNil(t, cancelBlobTx)

	// Verify the cancel blob transaction
	require.Equal(t, blobTx.Nonce(), cancelBlobTx.Nonce())
	require.Equal(t, blobTx.Gas(), cancelBlobTx.Gas())
	require.Equal(t, 0, len(cancelBlobTx.Data()))

	// Verify fee multipliers for BlobTx (2x)
	originalTip = blobTx.GasTipCap()
	cancelTip = cancelBlobTx.GasTipCap()
	require.True(t, cancelTip.Cmp(new(big.Int).Mul(originalTip, big.NewInt(2))) >= 0, "cancel blob tx tip should be at least 2x of original")

	originalFeeCap = blobTx.GasFeeCap()
	cancelFeeCap = cancelBlobTx.GasFeeCap()
	require.True(t, cancelFeeCap.Cmp(new(big.Int).Mul(originalFeeCap, big.NewInt(2))) >= 0, "cancel blob tx fee cap should be at least 2x of original")

	originalBlobFeeCap := blobTx.BlobGasFeeCap()
	cancelBlobFeeCap := cancelBlobTx.BlobGasFeeCap()
	require.True(t, cancelBlobFeeCap.Cmp(new(big.Int).Mul(originalBlobFeeCap, big.NewInt(2))) >= 0, "cancel blob tx blob fee cap should be at least 2x of original")
	require.Equal(t, 1, len(cancelBlobTx.BlobHashes()))
	require.Equal(t, 1, len(cancelBlobTx.BlobTxSidecar().Blobs))
}

func TestTxStateTransition(t *testing.T) {
	// Create test transactions
	tx1 := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     0,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(100),
		Gas:       21000,
		To:        &common.Address{},
		Value:     big.NewInt(1),
	})
	receipt := &types.Receipt{
		TxHash:       tx1.Hash(),
		BlockNumber:  big.NewInt(1000),
		Status:       1,
		BlobGasUsed:  0,
		BlobGasPrice: big.NewInt(0),
	}

	// Create rollup instance
	rollup, l1Mock, _, _ := setupTestRollup(t)

	// Test transaction state transitions
	t.Run("Transaction State Flow", func(t *testing.T) {
		// Step 1: Transaction exists only locally (not in mempool or block)
		err := rollup.pendingTxs.Add(tx1)
		require.NoError(t, err)
		status, err := rollup.getTxStatus(tx1)
		require.NoError(t, err)
		require.Equal(t, txStatusMissing, status.state)

		// Step 2: Transaction detected in mempool
		l1Mock.AddTx(tx1)
		status, err = rollup.getTxStatus(tx1)
		require.NoError(t, err)
		require.Equal(t, txStatusPending, status.state)

		// Step 3: Transaction included in block
		l1Mock.AddReceipt(receipt)
		status, err = rollup.getTxStatus(tx1)
		require.NoError(t, err)
		require.Equal(t, txStatusConfirmed, status.state)

		// Step 4: Transaction finalized (after 6 blocks)
		l1Mock.Block = types.NewBlockWithHeader(
			&types.Header{
				Number: big.NewInt(1006),
			},
		)
		status, err = rollup.getTxStatus(tx1)
		require.NoError(t, err)
		require.Equal(t, txStatusConfirmed, status.state)

		// Step 5: Process transaction and verify cleanup
		rollup.ProcessTx()
		// Verify transaction is removed from pendingTxs after finalization
		txRecord := rollup.pendingTxs.GetTxRecord(tx1.Hash())
		require.Nil(t, txRecord)
	})
}
