package derivation

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	tmlog "github.com/tendermint/tendermint/libs/log"

	commonbatch "morph-l2/common/batch"
	commonblob "morph-l2/common/blob"
)

// SPEC-005 section 4.3 + 5.1 Path B core encoding tests. These cover the
// verify_path_b.go logic via the extracted verifyPathBContent free function
// and a fake pathBBlockReader, avoiding the full L1 / beacon / authclient
// stack required by Derivation construction.
//
// The round-trip tests use makeEmptyL2Block (zero L2 txs / zero L1 messages)
// so ParsingTxs returns an empty payload. The codec is still exercised end
// to end -- BatchData.Append, V1/V2 payload selection, CompressBatchBytes,
// MakeBlobTxSidecar, BlobHashes() -- and the resulting hashes round-trip
// against the same code path. Tx-bearing blocks add no Path B coverage that
// the existing common/batch unit tests don't already provide.

type fakePathBBlockReader struct {
	blocks map[uint64]*eth.Block
	errs   map[uint64]error
}

func (f *fakePathBBlockReader) BlockByNumber(_ context.Context, n *big.Int) (*eth.Block, error) {
	num := n.Uint64()
	if e, ok := f.errs[num]; ok && e != nil {
		return nil, e
	}
	return f.blocks[num], nil // nil block when not registered -- exercises the "missing" branch
}

// makeEmptyL2Block builds a header-only block. ParsingTxs / BuildBlockContext
// only read fields verifyPathBContent already owns; no signer / state / receipts
// machinery is required.
func makeEmptyL2Block(num uint64) *eth.Block {
	h := &eth.Header{
		Number:   new(big.Int).SetUint64(num),
		Time:     1700000000 + num,
		GasLimit: 30_000_000,
		BaseFee:  big.NewInt(0),
	}
	return eth.NewBlockWithHeader(h)
}

// rebuildExpectedBlobHashes runs the same encoding pipeline as
// verifyPathBContent against the supplied blocks and returns the versioned
// hashes a real L1 commitBatch tx would have recorded for that batch. The
// round-trip tests use this as the L1-side oracle.
func rebuildExpectedBlobHashes(t *testing.T, blocks []*eth.Block, version, parentTotalL1Popped uint64, blobCount int) []common.Hash {
	t.Helper()

	bd := commonbatch.NewBatchData()
	total := parentTotalL1Popped
	for _, b := range blocks {
		txsPayload, l1Hashes, newTotal, l2TxNum, err := commonbatch.ParsingTxs(b.Transactions(), total)
		if err != nil {
			t.Fatalf("ParsingTxs(block %d): %v", b.NumberU64(), err)
		}
		l1MsgNum := int(newTotal - total)
		bd.Append(commonbatch.BuildBlockContext(b.Header(), l2TxNum+l1MsgNum, l1MsgNum), txsPayload, l1Hashes)
		total = newTotal
	}

	var payload []byte
	if version >= 2 {
		payload = bd.TxsPayloadV2()
	} else {
		payload = bd.TxsPayload()
	}

	compressed, err := commonblob.CompressBatchBytes(payload)
	if err != nil {
		t.Fatalf("CompressBatchBytes: %v", err)
	}
	sidecar, err := commonblob.MakeBlobTxSidecar(compressed, blobCount)
	if err != nil {
		t.Fatalf("MakeBlobTxSidecar: %v", err)
	}
	return sidecar.BlobHashes()
}

func TestPathB_RoundTripOK_V1(t *testing.T) {
	blocks := []*eth.Block{makeEmptyL2Block(10), makeEmptyL2Block(11), makeEmptyL2Block(12)}
	hashes := rebuildExpectedBlobHashes(t, blocks, 1, 0, 1)

	reader := &fakePathBBlockReader{blocks: map[uint64]*eth.Block{
		10: blocks[0], 11: blocks[1], 12: blocks[2],
	}}
	bi := &BatchInfo{
		batchIndex:                 7,
		version:                    1,
		firstBlockNumber:           10,
		lastBlockNumber:            12,
		parentTotalL1MessagePopped: 0,
		blobHashes:                 hashes,
	}

	if err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), bi); err != nil {
		t.Fatalf("V1 round-trip failed: %v", err)
	}
}

func TestPathB_RoundTripOK_V2(t *testing.T) {
	blocks := []*eth.Block{makeEmptyL2Block(20), makeEmptyL2Block(21)}
	hashes := rebuildExpectedBlobHashes(t, blocks, 2, 5, 1)

	reader := &fakePathBBlockReader{blocks: map[uint64]*eth.Block{
		20: blocks[0], 21: blocks[1],
	}}
	bi := &BatchInfo{
		batchIndex:                 8,
		version:                    2,
		firstBlockNumber:           20,
		lastBlockNumber:            21,
		parentTotalL1MessagePopped: 5,
		blobHashes:                 hashes,
	}

	if err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), bi); err != nil {
		t.Fatalf("V2 round-trip failed: %v", err)
	}
}

func TestPathB_VersionedHashMismatch(t *testing.T) {
	blocks := []*eth.Block{makeEmptyL2Block(10)}
	hashes := rebuildExpectedBlobHashes(t, blocks, 1, 0, 1)
	// Flip a single byte so the rebuilt hash cannot possibly match.
	tampered := make([]common.Hash, len(hashes))
	copy(tampered, hashes)
	tampered[0][0] ^= 0xff

	reader := &fakePathBBlockReader{blocks: map[uint64]*eth.Block{10: blocks[0]}}
	bi := &BatchInfo{
		batchIndex:                 9,
		version:                    1,
		firstBlockNumber:           10,
		lastBlockNumber:            10,
		parentTotalL1MessagePopped: 0,
		blobHashes:                 tampered,
	}

	err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), bi)
	if err == nil {
		t.Fatal("expected versioned hash mismatch error, got nil")
	}
	if !strings.Contains(err.Error(), "versioned hash mismatch") {
		t.Fatalf("error should mention 'versioned hash mismatch'; got: %v", err)
	}
}

func TestPathB_LocalBlockMissing(t *testing.T) {
	// Pre-build hashes that match a 2-block batch, then deliberately omit
	// block 11 from the reader so verifyPathBContent observes it as nil.
	blocks := []*eth.Block{makeEmptyL2Block(10), makeEmptyL2Block(11)}
	hashes := rebuildExpectedBlobHashes(t, blocks, 1, 0, 1)

	reader := &fakePathBBlockReader{blocks: map[uint64]*eth.Block{10: blocks[0]}}
	bi := &BatchInfo{
		batchIndex:                 11,
		version:                    1,
		firstBlockNumber:           10,
		lastBlockNumber:            11,
		parentTotalL1MessagePopped: 0,
		blobHashes:                 hashes,
	}

	err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), bi)
	if err == nil {
		t.Fatal("expected local block missing error, got nil")
	}
	if !strings.Contains(err.Error(), "missing") {
		t.Fatalf("error should mention 'missing'; got: %v", err)
	}
}

func TestPathB_LocalBlockReadError(t *testing.T) {
	blocks := []*eth.Block{makeEmptyL2Block(10)}
	hashes := rebuildExpectedBlobHashes(t, blocks, 1, 0, 1)

	reader := &fakePathBBlockReader{
		blocks: map[uint64]*eth.Block{10: blocks[0]},
		errs:   map[uint64]error{10: errors.New("rpc down")},
	}
	bi := &BatchInfo{
		batchIndex:       12,
		version:          1,
		firstBlockNumber: 10,
		lastBlockNumber:  10,
		blobHashes:       hashes,
	}

	err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), bi)
	if err == nil {
		t.Fatal("expected wrapped read error, got nil")
	}
	if !strings.Contains(err.Error(), "read local block") {
		t.Fatalf("error should mention 'read local block'; got: %v", err)
	}
}

func TestPathB_RejectsInvalidInputs(t *testing.T) {
	cases := []struct {
		name    string
		bi      *BatchInfo
		wantSub string
	}{
		{
			name:    "firstBlockNumber zero",
			bi:      &BatchInfo{firstBlockNumber: 0, lastBlockNumber: 5, blobHashes: []common.Hash{{}}},
			wantSub: "invalid block range",
		},
		{
			name:    "last < first",
			bi:      &BatchInfo{firstBlockNumber: 10, lastBlockNumber: 9, blobHashes: []common.Hash{{}}},
			wantSub: "invalid block range",
		},
		{
			name:    "empty blobHashes",
			bi:      &BatchInfo{firstBlockNumber: 5, lastBlockNumber: 5, blobHashes: nil},
			wantSub: "no blob hashes",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := &fakePathBBlockReader{}
			err := verifyPathBContent(context.Background(), reader, newDiscardMetrics(), tmlog.NewNopLogger(), tc.bi)
			if err == nil {
				t.Fatal("expected validation error, got nil")
			}
			if !strings.Contains(err.Error(), tc.wantSub) {
				t.Fatalf("error should mention %q; got: %v", tc.wantSub, err)
			}
		})
	}
}
