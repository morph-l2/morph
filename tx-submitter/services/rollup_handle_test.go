package services

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/event"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/metrics"
	"morph-l2/tx-submitter/mock"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"
)

var (
	finalizeCalldata = "13361101000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000f90000000000000036a700000000000000000000000000001d630ed957794a607e274cdd09fd72217c7c2d9c0be5c11f993d66bf58f1ba098d950170e28b2b52dabe6efe96dbcfe85f691872adb80af292efbb7ded8e53ada1a000d75f5b0490922155046b8428750e9eede0df103c264d67434112e0b4320b8702e53bafa072ec3bbdafbbed9601a82f193d6d64db604ed931a0617779ccebb63d83b92dd5879bfb5e44ba433bafbb65afc822c7d000210ec0d296c59446d2cfa4d7461c3b8d74650efdca322d5fe15e699d534ee47ec9321bcb82bc7a9ec6e9260ff211c1f2db0cf21655a25035100797bc47c339502a0f1a791c0739d4e96000000000000000"
)

// setupTestRollup creates a test Rollup instance with mocked dependencies
func setupTestRollup(t *testing.T) (*Rollup, *mock.L1ClientWrapper, *mock.L2ClientWrapper, *mock.MockRollup) {
	// Create mock clients
	l1Mock := mock.NewL1ClientWrapper()
	l2Mock := mock.NewL2ClientWrapper()

	// Set mock values for gas estimation
	l1Mock.BaseFee = big.NewInt(1e9) // 1 gwei
	l1Mock.TipCap = big.NewInt(1e9)  // 1 gwei

	// Create a private key for testing
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	// Create mock metrics
	metrics := metrics.NewMetrics()
	t.Cleanup(func() {
		metrics.UnregisterMetrics()
	})

	// Create mock event storage
	eventStorage := mock.NewMockEventInfoStorage()
	err = eventStorage.Load()
	require.NoError(t, err)

	// Initialize event storage with test data
	eventStorage.SetBlockProcessed(1000)
	eventStorage.SetBlockTime(uint64(time.Now().Unix()))
	err = eventStorage.Store()
	require.NoError(t, err)

	// Create mock event indexer
	indexer := event.NewEventIndexer(
		nil, // We don't need a real ethclient.Client for testing
		big.NewInt(0),
		ethereum.FilterQuery{},
		100,
		eventStorage,
	)

	// Create mock rotator
	rotator := NewRotator(common.Address{}, common.Address{}, indexer)

	// Create mock L1Staking
	l1Staking := mock.NewMockL1Staking()
	// Set some test stakers
	testStakers := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}
	l1Staking.SetActiveStakers(testStakers)

	// Create rollup config
	cfg := utils.Config{
		MaxTip:         10e9,
		MaxBaseFee:     100e9,
		MinTip:         1e9,
		TipFeeBump:     100,
		TxTimeout:      10 * time.Second,
		PriorityRollup: true,
	}

	// Create mock journal
	mockJournal := mock.NewMockJournal()

	// Create mock rollup
	mockRollup := mock.NewMockRollup()

	// Get Rollup ABI
	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	require.NotNil(t, rollupAbi)

	// Create rollup instance
	rollup := NewRollup(
		context.Background(),
		metrics,
		nil,                      // l1RpcClient
		l1Mock,                   // l1Client
		[]iface.L2Client{l2Mock}, // l2Clients
		mockRollup,               // rollup
		l1Staking,                // staking
		big.NewInt(1),            // chainId
		privateKey,               // privKey
		common.Address{},         // rollupAddr
		rollupAbi,                // abi
		cfg,                      // cfg
		nil,                      // rsaPriv
		rotator,                  // rotator
		nil,                      // ldb
		nil,                      // bm
		eventStorage,             // eventInfoStorage
	)

	// Initialize pending transactions
	rollup.pendingTxs = NewPendingTxs([]byte{}, []byte{}, mockJournal)

	// Initialize reorg detector
	rollup.reorgDetector = &ReorgDetector{
		l1Client: l1Mock,
		metrics:  metrics,
	}

	return rollup, l1Mock, l2Mock, mockRollup
}

// TestHandleDiscardedTx tests the handling of discarded transactions
func TestHandleDiscardedTx(t *testing.T) {
	r, l1Mock, _, _ := setupTestRollup(t)

	// Create a test transaction
	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		GasTipCap: big.NewInt(1e9),
		GasFeeCap: big.NewInt(2e9),
		Gas:       21000,
		To:        &common.Address{},
	})

	txRecord := &types.TxRecord{
		Tx:         tx,
		SendTime:   uint64(time.Now().Unix()),
		QueryTimes: 5, // Set high enough to trigger discard handling
	}

	// Test case 1: Transaction was confirmed in a reorg
	l1Mock.SendTxErr = core.ErrNonceTooLow
	err := r.handleDiscardedTx(txRecord, tx, "commitBatch")
	require.NoError(t, err)
	require.Equal(t, 0, len(r.pendingTxs.GetAll()), "Transaction should be removed from pending pool")

	// Test case 2: Successful resubmission
	l1Mock.SendTxErr = nil
	tx = ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     2,
		GasTipCap: big.NewInt(1e9),
		GasFeeCap: big.NewInt(2e9),
		Gas:       21000,
		To:        &common.Address{},
	})
	txRecord.Tx = tx

	err = r.handleDiscardedTx(txRecord, tx, "commitBatch")
	require.NoError(t, err)
	require.Equal(t, 1, len(r.pendingTxs.GetAll()), "New transaction should be added to pending pool")
}

// TestHandleReorg tests the handling of chain reorganizations
func TestHandleReorg(t *testing.T) {
	r, _, _, _ := setupTestRollup(t)

	// Test reorg handling
	depth := uint64(2)
	err := r.handleReorg(depth)
	require.NoError(t, err)

	// Verify metrics
	require.Equal(t, float64(depth), r.metrics.GetReorgDepth())
	require.Equal(t, float64(1), r.metrics.GetReorgCount())

}

// TestHandleMissingTx tests the handling of missing transactions
func TestHandleMissingTx(t *testing.T) {
	r, l1Mock, _, _ := setupTestRollup(t)

	// Create a test transaction
	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		GasTipCap: big.NewInt(1e9),
		GasFeeCap: big.NewInt(2e9),
		Gas:       21000,
		To:        &common.Address{},
	})

	// Test case 1: Transaction with low query times
	txRecord := &types.TxRecord{
		Tx:         tx,
		SendTime:   uint64(time.Now().Unix()),
		QueryTimes: 2,
	}

	// add record to localpool
	r.pendingTxs.txinfos[tx.Hash()] = txRecord

	err := r.handleMissingTx(txRecord, tx, "commitBatch")
	require.NoError(t, err)
	record := r.pendingTxs.GetTxRecord(tx.Hash())
	require.NotNil(t, record, "Transaction record should exist")
	require.Equal(t, uint64(3), record.QueryTimes)

	// Test case 2: Transaction with high query times
	oldHash := tx.Hash()
	txRecord.QueryTimes = 5
	// Set up mock for successful resubmission
	l1Mock.SendTxErr = nil
	err = r.handleMissingTx(txRecord, tx, "commitBatch")
	require.NoError(t, err)

	// The old transaction should be removed
	record = r.pendingTxs.GetTxRecord(oldHash)
	require.Nil(t, record, "Old transaction should be removed")

	// Find the new transaction by checking all pending transactions
	found := false
	for _, txRecord := range r.pendingTxs.GetAll() {
		if txRecord.Tx.Nonce() == tx.Nonce() {
			found = true
			require.Equal(t, uint64(0), txRecord.QueryTimes, "New transaction should have reset query times")
			break
		}
	}
	require.True(t, found, "New transaction should exist in pending pool")
}
