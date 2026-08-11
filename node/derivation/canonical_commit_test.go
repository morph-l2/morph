package derivation

import (
	"context"
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

type resolverFakeClient struct {
	header          *eth.Header
	changedHeader   *eth.Header
	changeAfterCall int
	headerCalls     int
	callHash        common.Hash
	logs            []eth.Log
	filterCalls     int
	lastQuery       ethereum.FilterQuery
	transactions    map[common.Hash]*eth.Transaction
}

func (f *resolverFakeClient) HeaderByNumber(context.Context, *big.Int) (*eth.Header, error) {
	f.headerCalls++
	if f.changedHeader != nil && f.changeAfterCall > 0 && f.headerCalls >= f.changeAfterCall {
		return f.changedHeader, nil
	}
	return f.header, nil
}

func (f *resolverFakeClient) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	return canonicalResolverRollupABI.Methods["committedBatches"].Outputs.Pack(f.callHash)
}

func (f *resolverFakeClient) FilterLogs(_ context.Context, query ethereum.FilterQuery) ([]eth.Log, error) {
	f.filterCalls++
	f.lastQuery = query
	filtered := make([]eth.Log, 0, len(f.logs))
	for _, log := range f.logs {
		if len(query.Addresses) > 0 && log.Address != query.Addresses[0] {
			continue
		}
		matched := true
		for position, alternatives := range query.Topics {
			if len(alternatives) == 0 {
				continue
			}
			if len(log.Topics) <= position {
				matched = false
				break
			}
			positionMatched := false
			for _, topic := range alternatives {
				if log.Topics[position] == topic {
					positionMatched = true
					break
				}
			}
			if !positionMatched {
				matched = false
				break
			}
		}
		if matched {
			filtered = append(filtered, log)
		}
	}
	return filtered, nil
}

func (f *resolverFakeClient) TransactionByHash(_ context.Context, hash common.Hash) (*eth.Transaction, bool, error) {
	return f.transactions[hash], false, nil
}

func resolverTestHeader(number uint64, marker byte) *eth.Header {
	return &eth.Header{Number: new(big.Int).SetUint64(number), Time: number, Extra: []byte{marker}}
}

func resolverTestLog(address common.Address, event, hash common.Hash, index, block uint64, txIndex, logIndex uint) eth.Log {
	return eth.Log{
		Address:     address,
		Topics:      []common.Hash{event, common.BigToHash(uint64Big(index)), hash},
		BlockNumber: block,
		TxIndex:     txIndex,
		Index:       logIndex,
	}
}

func TestResolveCanonicalCommitAfterLastRevertWithSameHashRecommit(t *testing.T) {
	const snapshotNumber = uint64(100)
	const batchIndex = uint64(7)
	rollupAddress := common.HexToAddress("0x1234")
	batchHash := common.HexToHash("0xaaaa")
	wrongHash := common.HexToHash("0xbbbb")
	wantTx := common.HexToHash("0x3333")

	oldCommit := resolverTestLog(rollupAddress, canonicalCommitEventTopic, batchHash, batchIndex, 10, 0, 0)
	revert := resolverTestLog(rollupAddress, canonicalRevertEventTopic, batchHash, batchIndex, 11, 0, 0)
	wrongCommit := resolverTestLog(rollupAddress, canonicalCommitEventTopic, wrongHash, batchIndex, 12, 0, 0)
	recommit := resolverTestLog(rollupAddress, canonicalCommitEventTopic, batchHash, batchIndex, 12, 1, 3)
	recommit.TxHash = wantTx
	client := &resolverFakeClient{
		header:   resolverTestHeader(snapshotNumber, 1),
		callHash: batchHash,
		logs:     []eth.Log{recommit, oldCommit, wrongCommit, revert}, // deliberately unsorted
	}

	got, err := ResolveCanonicalCommit(context.Background(), client, rollupAddress, batchIndex, snapshotNumber)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, batchHash, got.BatchHash)
	require.Equal(t, wantTx, got.Log.TxHash, "same-hash recommit must be identified by its post-Revert log")
	require.Equal(t, snapshotNumber, client.lastQuery.ToBlock.Uint64())
	require.Equal(t, 3, client.headerCalls, "snapshot must be pinned and verified before/after reads")
}

func TestResolveCanonicalCommitReturnsNilForZeroGetter(t *testing.T) {
	client := &resolverFakeClient{header: resolverTestHeader(100, 1)}
	got, err := ResolveCanonicalCommit(context.Background(), client, common.HexToAddress("0x1234"), 1, 100)
	require.NoError(t, err)
	require.Nil(t, got)
	require.Zero(t, client.filterCalls, "zero committedBatches must not scan logs")
}

func TestResolveCanonicalCommitRejectsChangedSnapshot(t *testing.T) {
	client := &resolverFakeClient{
		header:          resolverTestHeader(100, 1),
		changedHeader:   resolverTestHeader(100, 2),
		changeAfterCall: 3,
		callHash:        common.HexToHash("0xaaaa"),
		logs: []eth.Log{
			resolverTestLog(common.HexToAddress("0x1234"), canonicalCommitEventTopic, common.HexToHash("0xaaaa"), 1, 10, 0, 0),
		},
	}
	_, err := ResolveCanonicalCommit(context.Background(), client, common.HexToAddress("0x1234"), 1, 100)
	require.ErrorIs(t, err, ErrSnapshotChanged)
}

func TestResolveStoredBlobSourceSearchesBeforeRevert(t *testing.T) {
	const snapshotNumber = uint64(100)
	const batchIndex = uint64(7)
	rollupAddress := common.HexToAddress("0x1234")
	corpus := loadCalldataFixtureCorpus(t)
	fixtureData := make(map[string][]byte, len(corpus.Fixtures))
	for _, fixture := range corpus.Fixtures {
		data, err := hexutil.Decode(fixture.Data)
		require.NoError(t, err)
		fixtureData[fixture.Name] = data
	}

	blobHashes := []common.Hash{common.HexToHash("0x1111"), common.HexToHash("0x2222")}
	storedHashBytes := make([]byte, 0, len(blobHashes)*common.HashLength)
	for _, hash := range blobHashes {
		storedHashBytes = append(storedHashBytes, hash[:]...)
	}
	storedHash := crypto.Keccak256Hash(storedHashBytes)
	originalTx := newResolverBlobTx(1, fixtureData["pre_submitter_commitBatch"], blobHashes)
	bloblessTx := eth.NewTx(&eth.LegacyTx{Nonce: 2, To: &rollupAddress, Data: fixtureData["current_commitState"]})

	originalCommit := resolverTestLog(rollupAddress, canonicalCommitEventTopic, common.HexToHash("0xaaaa"), batchIndex, 10, 0, 0)
	originalCommit.TxHash = originalTx.Hash()
	revert := resolverTestLog(rollupAddress, canonicalRevertEventTopic, common.HexToHash("0xaaaa"), batchIndex, 11, 0, 0)
	bloblessRecommit := resolverTestLog(rollupAddress, canonicalCommitEventTopic, common.HexToHash("0xbbbb"), batchIndex, 12, 0, 0)
	bloblessRecommit.TxHash = bloblessTx.Hash()
	client := &resolverFakeClient{
		header:   resolverTestHeader(snapshotNumber, 1),
		callHash: storedHash,
		logs:     []eth.Log{bloblessRecommit, revert, originalCommit},
		transactions: map[common.Hash]*eth.Transaction{
			originalTx.Hash(): originalTx,
			bloblessTx.Hash(): bloblessTx,
		},
	}
	d := newFixtureDerivation(t)

	source, err := ResolveStoredBlobSource(context.Background(), client, rollupAddress, batchIndex, snapshotNumber, d.UnPackData)
	require.NoError(t, err)
	require.Equal(t, originalTx.Hash(), source.CommitLog.TxHash)
	require.Equal(t, blobHashes, source.BlobHashes)
	require.Equal(t, []common.Hash{canonicalCommitEventTopic}, client.lastQuery.Topics[0],
		"stored-source lookup must scan every historical Commit and must not reuse the post-Revert canonical filter")
}

func TestResolveStoredBlobSourcePreservesZeroHashSpecialCase(t *testing.T) {
	client := &resolverFakeClient{
		header:   resolverTestHeader(100, 1),
		callHash: zeroBlobVersionedHash,
	}
	_, err := ResolveStoredBlobSource(context.Background(), client, common.HexToAddress("0x1234"), 1, 100, newFixtureDerivation(t).UnPackData)
	require.ErrorIs(t, err, ErrZeroBlobVersionedHash)
	require.Zero(t, client.filterCalls)
}

func TestStoredBlobHashMatchesLegacyAndV2(t *testing.T) {
	hashes := []common.Hash{common.HexToHash("0x1111"), common.HexToHash("0x2222")}
	require.True(t, storedBlobHashMatches(1, hashes[:1], hashes[0]))
	require.False(t, storedBlobHashMatches(1, hashes[:1], hashes[1]))
	encoded := append(append([]byte{}, hashes[0][:]...), hashes[1][:]...)
	require.True(t, storedBlobHashMatches(2, hashes, crypto.Keccak256Hash(encoded)))
	require.False(t, storedBlobHashMatches(2, hashes, hashes[0]))
}

func newResolverBlobTx(nonce uint64, data []byte, blobHashes []common.Hash) *eth.Transaction {
	return eth.NewTx(&eth.BlobTx{
		ChainID:    uint256.NewInt(1),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(1),
		GasFeeCap:  uint256.NewInt(2),
		Gas:        100_000,
		To:         common.HexToAddress("0x1234"),
		Value:      uint256.NewInt(0),
		Data:       data,
		BlobFeeCap: uint256.NewInt(1),
		BlobHashes: blobHashes,
		V:          uint256.NewInt(0),
		R:          uint256.NewInt(0),
		S:          uint256.NewInt(0),
	})
}
