package batch

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"math/big"
	"testing"

	"morph-l2/bindings/bindings"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

func TestSelectCanonicalCommitUsesRecommitAfterLastRevert(t *testing.T) {
	hash := common.HexToHash("0x1234")
	commits := []canonicalCommitEvent{
		{batchIndex: 9, batchHash: hash, raw: ethtypes.Log{BlockNumber: 10, TxIndex: 1, Index: 2}},
		{batchIndex: 9, batchHash: hash, raw: ethtypes.Log{BlockNumber: 12, TxIndex: 0, Index: 0}},
	}
	reverts := []canonicalRevertEvent{
		{batchIndex: 9, batchHash: hash, raw: ethtypes.Log{BlockNumber: 11, TxIndex: 3, Index: 4}},
	}

	selected, err := selectCanonicalCommit(9, hash, commits, reverts)
	require.NoError(t, err)
	require.Equal(t, uint64(12), selected.raw.BlockNumber)

	// Ordering includes transaction and log index, not just block number.
	commits = append(commits, canonicalCommitEvent{
		batchIndex: 9,
		batchHash:  hash,
		raw:        ethtypes.Log{BlockNumber: 12, TxIndex: 0, Index: 3},
	})
	selected, err = selectCanonicalCommit(9, hash, commits, reverts)
	require.NoError(t, err)
	require.Equal(t, uint(3), selected.raw.Index)
}

func TestSelectCanonicalCommitRejectsEmptyCheckpoint(t *testing.T) {
	_, err := selectCanonicalCommit(7, common.Hash{}, nil, nil)
	require.ErrorContains(t, err, "no canonical committed hash")
}

type queueBackend struct {
	messages     map[uint64]common.Hash
	pending      uint64
	next         uint64
	callBlocks   []uint64
	latest       uint64
	headers      []*ethtypes.Header
	headerReads  int
	transactions map[common.Hash]*ethtypes.Transaction
}

func (b *queueBackend) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return []byte{1}, nil
}

func (b *queueBackend) CallContract(_ context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	if blockNumber == nil {
		return nil, errors.New("queue getter was not pinned to snapshot")
	}
	b.callBlocks = append(b.callBlocks, blockNumber.Uint64())
	if len(call.Data) < 4 {
		return nil, errors.New("short queue calldata")
	}
	selector := call.Data[:4]
	switch {
	case bytes.Equal(selector, crypto.Keccak256([]byte("getCrossDomainMessage(uint256)"))[:4]):
		if len(call.Data) != 36 {
			return nil, errors.New("invalid message getter calldata")
		}
		index := new(big.Int).SetBytes(call.Data[4:]).Uint64()
		hash, ok := b.messages[index]
		if !ok {
			return nil, errors.New("message index unavailable")
		}
		return hash.Bytes(), nil
	case bytes.Equal(selector, crypto.Keccak256([]byte("pendingQueueIndex()"))[:4]):
		return common.LeftPadBytes(new(big.Int).SetUint64(b.pending).Bytes(), 32), nil
	case bytes.Equal(selector, crypto.Keccak256([]byte("nextCrossDomainMessageIndex()"))[:4]):
		return common.LeftPadBytes(new(big.Int).SetUint64(b.next).Bytes(), 32), nil
	default:
		return nil, errors.New("unexpected queue getter")
	}
}

func (b *queueBackend) BlockNumber(context.Context) (uint64, error) { return b.latest, nil }

func (b *queueBackend) HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error) {
	if len(b.headers) == 0 {
		return nil, errors.New("header unavailable")
	}
	index := b.headerReads
	if index >= len(b.headers) {
		index = len(b.headers) - 1
	}
	b.headerReads++
	return b.headers[index], nil
}

func (b *queueBackend) TransactionByHash(_ context.Context, hash common.Hash) (*ethtypes.Transaction, bool, error) {
	tx, ok := b.transactions[hash]
	if !ok {
		return nil, false, errors.New("transaction unavailable")
	}
	return tx, false, nil
}

type canonicalRollupStub struct {
	blobHash       common.Hash
	queueAddress   common.Address
	callBlockReads []uint64
}

func (r *canonicalRollupStub) record(opts *bind.CallOpts) error {
	if opts == nil || opts.BlockNumber == nil {
		return errors.New("rollup getter was not pinned to snapshot")
	}
	r.callBlockReads = append(r.callBlockReads, opts.BlockNumber.Uint64())
	return nil
}

func (r *canonicalRollupStub) BatchBlobVersionedHashes(opts *bind.CallOpts, _ *big.Int) ([32]byte, error) {
	if err := r.record(opts); err != nil {
		return [32]byte{}, err
	}
	return [32]byte(r.blobHash), nil
}

func (r *canonicalRollupStub) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	if err := r.record(opts); err != nil {
		return common.Address{}, err
	}
	return r.queueAddress, nil
}

func (*canonicalRollupStub) CommittedBatches(*bind.CallOpts, *big.Int) ([32]byte, error) {
	return [32]byte{}, errors.New("not implemented")
}
func (*canonicalRollupStub) LastCommittedBatchIndex(*bind.CallOpts) (*big.Int, error) {
	return nil, errors.New("not implemented")
}
func (*canonicalRollupStub) LastFinalizedBatchIndex(*bind.CallOpts) (*big.Int, error) {
	return nil, errors.New("not implemented")
}
func (*canonicalRollupStub) BatchDataStore(*bind.CallOpts, *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	return struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}{}, errors.New("not implemented")
}
func (*canonicalRollupStub) FinalizedStateRoots(*bind.CallOpts, *big.Int) ([32]byte, error) {
	return [32]byte{}, errors.New("not implemented")
}
func (*canonicalRollupStub) FilterCommitBatch(*bind.FilterOpts, []*big.Int, [][32]byte) (*bindings.RollupCommitBatchIterator, error) {
	return nil, errors.New("not implemented")
}
func (*canonicalRollupStub) FilterRevertBatch(*bind.FilterOpts, []*big.Int, [][32]byte) (*bindings.RollupRevertBatchIterator, error) {
	return nil, errors.New("not implemented")
}

func TestReconstructCurrentCanonicalCommitAtPinnedSnapshot(t *testing.T) {
	parent := BatchHeaderV1{
		BatchHeaderV0: BatchHeaderV0{
			BatchIndex:           7,
			TotalL1MessagePopped: 3,
			PostStateRoot:        common.HexToHash("0x70"),
		},
		LastBlockNumber: 90,
	}.Bytes()
	input := bindings.IRollupBatchDataInput{
		Version:           BatchHeaderVersion1,
		ParentBatchHeader: parent,
		LastBlockNumber:   100,
		NumL1Messages:     2,
		PrevStateRoot:     common.HexToHash("0x70"),
		PostStateRoot:     common.HexToHash("0x80"),
		WithdrawalRoot:    common.HexToHash("0x90"),
	}
	abis, err := DefaultRollupABIs()
	require.NoError(t, err)
	calldata, err := abis.Current.Pack("commitBatch", input)
	require.NoError(t, err)

	messages := map[uint64]common.Hash{
		3: common.HexToHash("0x300"),
		4: common.HexToHash("0x400"),
	}
	data := make([]byte, 10)
	binary.BigEndian.PutUint64(data[:8], input.LastBlockNumber)
	binary.BigEndian.PutUint16(data[8:], input.NumL1Messages)
	data = append(data, messages[3].Bytes()...)
	data = append(data, messages[4].Bytes()...)
	parentHash, err := parent.Hash()
	require.NoError(t, err)
	blobHash := common.HexToHash("0xb10b")
	expected := BatchHeaderV1{
		BatchHeaderV0: BatchHeaderV0{
			BatchIndex:           8,
			L1MessagePopped:      2,
			TotalL1MessagePopped: 5,
			DataHash:             crypto.Keccak256Hash(data),
			BlobVersionedHash:    blobHash,
			PrevStateRoot:        common.Hash(input.PrevStateRoot),
			PostStateRoot:        common.Hash(input.PostStateRoot),
			WithdrawalRoot:       common.Hash(input.WithdrawalRoot),
			ParentBatchHash:      parentHash,
		},
		LastBlockNumber: input.LastBlockNumber,
	}.Bytes()
	expectedHash, err := expected.Hash()
	require.NoError(t, err)

	backend := &queueBackend{messages: messages, pending: 5, next: 6}
	rollup := &canonicalRollupStub{
		blobHash:     blobHash,
		queueAddress: common.HexToAddress("0x1234"),
	}
	cache := &BatchCache{
		ctx:                context.Background(),
		l1Client:           backend,
		rollupContract:     rollup,
		l1Snapshot:         &confirmedL1Snapshot{number: 99, hash: common.HexToHash("0x99")},
		sealedBatches:      make(map[uint64]*eth.RPCRollupBatch),
		sealedBatchHeaders: make(map[uint64]*BatchHeaderBytes),
	}
	tx := ethtypes.NewTx(&ethtypes.BlobTx{
		ChainID:    uint256.NewInt(1),
		GasTipCap:  uint256.NewInt(1),
		GasFeeCap:  uint256.NewInt(1),
		Value:      uint256.NewInt(0),
		Data:       calldata,
		BlobFeeCap: uint256.NewInt(1),
		BlobHashes: []common.Hash{blobHash},
	})

	got, err := cache.reconstructCommittedBatchHeader(8, expectedHash, tx)
	require.NoError(t, err)
	require.Equal(t, expected, got)
	require.NotEmpty(t, rollup.callBlockReads)
	for _, number := range rollup.callBlockReads {
		require.Equal(t, uint64(99), number)
	}
	for _, number := range backend.callBlocks {
		require.Equal(t, uint64(99), number)
	}

	backend.pending = 4
	_, err = cache.reconstructCommittedBatchHeader(8, expectedHash, tx)
	require.ErrorContains(t, err, "cursor mismatch")
}

func TestConfirmedL1SnapshotHashChangeFailsClosed(t *testing.T) {
	number := uint64(14)
	backend := &queueBackend{
		latest: 20,
		headers: []*ethtypes.Header{
			{Number: new(big.Int).SetUint64(number), Extra: []byte("before")},
			{Number: new(big.Int).SetUint64(number), Extra: []byte("after")},
		},
	}
	cache := &BatchCache{ctx: context.Background(), l1Client: backend}
	err := cache.withConfirmedL1Snapshot(func() error {
		require.Equal(t, number, cache.snapshotCallOpts().BlockNumber.Uint64())
		return nil
	})
	require.ErrorContains(t, err, "snapshot changed")
	require.Nil(t, cache.l1Snapshot)
}

type scriptedL2 struct {
	safe     *ethtypes.Block
	numbered map[uint64][]*ethtypes.Block
	reads    map[uint64]int
}

func (c *scriptedL2) BlockNumber(context.Context) (uint64, error) { return c.safe.NumberU64(), nil }
func (c *scriptedL2) Len() int                                    { return 1 }
func (c *scriptedL2) BlockByNumber(_ context.Context, number *big.Int) (*ethtypes.Block, error) {
	if number.Int64() == int64(rpc.SafeBlockNumber) {
		return c.safe, nil
	}
	n := number.Uint64()
	values := c.numbered[n]
	if len(values) == 0 {
		return nil, nil
	}
	read := c.reads[n]
	if read >= len(values) {
		read = len(values) - 1
	}
	c.reads[n]++
	return values[read], nil
}

func l2Block(number uint64, marker string) *ethtypes.Block {
	return ethtypes.NewBlockWithHeader(&ethtypes.Header{
		Number: new(big.Int).SetUint64(number),
		Extra:  []byte(marker),
	})
}

func initializedSafeHeadCache(last uint64, l2 L2MultiClient) *BatchCache {
	parent := BatchHeaderV1{
		BatchHeaderV0:   BatchHeaderV0{BatchIndex: 3},
		LastBlockNumber: last,
	}.Bytes()
	return &BatchCache{
		ctx:                   context.Background(),
		initDone:              true,
		parentBatchHeader:     &parent,
		lastPackedBlockHeight: last,
		currentBlockNumber:    last,
		batchData:             NewBatchData(),
		batchConfig:           BatchConfig{BlockInterval: 100},
		l2Clients:             l2,
	}
}

func TestAssembleCurrentBatchRejectsChangedSafeHeadHash(t *testing.T) {
	l2 := &scriptedL2{
		safe:     l2Block(10, "safe-before"),
		numbered: map[uint64][]*ethtypes.Block{10: {l2Block(10, "safe-after")}},
		reads:    make(map[uint64]int),
	}
	cache := initializedSafeHeadCache(10, l2)
	err := cache.AssembleCurrentBatchHeader()
	require.ErrorContains(t, err, "safe head changed")
}

func TestAssembleCurrentBatchRejectsL2Gap(t *testing.T) {
	l2 := &scriptedL2{
		safe: l2Block(12, "safe"),
		numbered: map[uint64][]*ethtypes.Block{
			10: {l2Block(10, "parent")},
			11: {l2Block(12, "wrong-height")},
		},
		reads: make(map[uint64]int),
	}
	cache := initializedSafeHeadCache(10, l2)
	err := cache.AssembleCurrentBatchHeader()
	require.ErrorContains(t, err, "L2 gap")
}

func TestResetCanonicalBootstrapLeavesNoCheckpoint(t *testing.T) {
	storage := NewBatchStorage(openTestKV(t))
	storeTestChain(t, storage, []uint64{1, 2})
	header := makeTestHeader(2)
	cache := &BatchCache{
		initDone:           true,
		batchStorage:       storage,
		sealedBatches:      map[uint64]*eth.RPCRollupBatch{2: {}},
		sealedBatchHeaders: map[uint64]*BatchHeaderBytes{2: &header},
		parentBatchHeader:  &header,
		batchData:          NewBatchData(),
		currentBlockHash:   common.HexToHash("0x1234"),
	}

	require.NoError(t, cache.ResetCanonicalBootstrap())
	_, err := storage.loadBatchIndices()
	require.True(t, isKVNotFound(err), "checkpoint must be absent after rollback: %v", err)
	require.False(t, cache.initDone)
	require.Empty(t, cache.sealedBatches)
	require.Empty(t, cache.sealedBatchHeaders)
	require.Nil(t, cache.parentBatchHeader)
	require.Nil(t, cache.batchData)
	require.Zero(t, cache.currentBlockHash)
}
