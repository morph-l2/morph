package derivation

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	eth "github.com/morph-l2/go-ethereum/core/types"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"

	"morph-l2/node/types"
	"morph-l2/node/zstd"
)

// buildBlockContexts returns the concatenated 60-byte encoding of `count`
// sequential, tx-empty blocks starting at `startBlock`. The produced layout
// matches what tx-submitter places at the head of a V1/V2 batch payload.
func buildBlockContexts(startBlock uint64, count int) []byte {
	buf := make([]byte, 0, count*60)
	for i := 0; i < count; i++ {
		wb := &types.WrappedBlock{
			Number:    startBlock + uint64(i),
			Timestamp: 1_700_000_000 + uint64(i)*6,
			BaseFee:   big.NewInt(1_000_000_000),
			GasLimit:  30_000_000,
		}
		buf = append(buf, wb.BlockContextBytes(0, 0)...)
	}
	return buf
}

// buildV1ParentHeader encodes a minimal V1 parent header whose LastBlockNumber
// is one below `nextStartBlock`, so that ParseBatch can derive blockCount via
// the (batch.LastBlockNumber - parent.LastBlockNumber) path.
func buildV1ParentHeader(parentIndex, nextStartBlock uint64) []byte {
	return types.BatchHeaderV1{
		BatchHeaderV0: types.BatchHeaderV0{
			BatchIndex:        parentIndex,
			BlobVersionedHash: types.EmptyVersionedHash,
		},
		LastBlockNumber: nextStartBlock - 1,
	}.Bytes()
}

// splitCompressedIntoBlobs mirrors the tx-submitter strategy of compressing
// the entire payload as a single zstd stream and then slicing the compressed
// bytes into MaxBlobBytesSize chunks, each packed into a canonical blob.
func splitCompressedIntoBlobs(t *testing.T, compressed []byte) []kzg4844.Blob {
	t.Helper()
	var blobs []kzg4844.Blob
	for offset := 0; offset < len(compressed); offset += types.MaxBlobBytesSize {
		end := offset + types.MaxBlobBytesSize
		if end > len(compressed) {
			end = len(compressed)
		}
		blob, err := types.MakeBlobCanonical(compressed[offset:end])
		require.NoError(t, err)
		blobs = append(blobs, *blob)
	}
	if len(blobs) == 0 {
		// An empty payload still requires at least one (empty) blob so that
		// downstream consumers can iterate. Production submitters never emit
		// an empty batch, but the helper should remain total.
		blobs = append(blobs, kzg4844.Blob{})
	}
	return blobs
}

// TestParseBatchSingleBlob covers the backward-compatible path where a V1
// batch fits in a single blob. It guards against regressions in the recent
// "concatenate then decompress" refactor: the single-blob flow must still
// yield the same block contexts it did before multi-blob support landed.
func TestParseBatchSingleBlob(t *testing.T) {
	const (
		parentIndex = 99
		startBlock  = 1_000
		blockCount  = 5
	)

	blockCtx := buildBlockContexts(startBlock, blockCount)
	payload := append(blockCtx, 0x00) // empty tx stream terminator

	compressed, err := zstd.CompressBatchBytes(payload)
	require.NoError(t, err)
	require.LessOrEqual(t, len(compressed), types.MaxBlobBytesSize,
		"single-blob test expects compressed payload to fit in one blob")

	blobs := splitCompressedIntoBlobs(t, compressed)
	require.Len(t, blobs, 1)

	batch := geth.RPCRollupBatch{
		Version:           1,
		ParentBatchHeader: buildV1ParentHeader(parentIndex, startBlock),
		LastBlockNumber:   startBlock + blockCount - 1,
		Sidecar:           eth.BlobTxSidecar{Blobs: blobs},
	}

	var bi BatchInfo
	require.NoError(t, bi.ParseBatch(batch))

	require.EqualValues(t, parentIndex+1, bi.batchIndex)
	require.EqualValues(t, startBlock, bi.FirstBlockNumber())
	require.EqualValues(t, startBlock+blockCount-1, bi.LastBlockNumber())
	require.Len(t, bi.blockContexts, blockCount)
	for i, bc := range bi.blockContexts {
		require.EqualValues(t, uint64(startBlock+i), bc.Number,
			"block %d number mismatch", i)
	}
}

// TestParseBatchMultiBlob is the core multi-blob regression: it forces the
// compressed payload to exceed a single blob's capacity and verifies that
// ParseBatch reconstructs the decompressed stream by concatenating all blob
// bodies before running zstd.Decompress. A naive per-blob decompression loop
// would fail on blob[1] since it is mid-zstd-frame data, so a successful
// parse here proves the concatenation path is wired correctly.
//
// Compression-resistant random bytes are appended after the block-context
// header (past the tx terminator) purely to inflate the compressed size; the
// tx decoder stops at the first 0x00 byte and trailing random bytes are never
// interpreted as transactions.
func TestParseBatchMultiBlob(t *testing.T) {
	const (
		parentIndex = 123
		startBlock  = 2_000
		blockCount  = 8
	)

	blockCtx := buildBlockContexts(startBlock, blockCount)

	// 1 byte tx terminator + ~1.2x blob capacity of incompressible noise to
	// guarantee the zstd output straddles a blob boundary.
	padLen := types.MaxBlobBytesSize + types.MaxBlobBytesSize/5
	pad := make([]byte, padLen)
	_, err := rand.Read(pad)
	require.NoError(t, err)

	payload := make([]byte, 0, len(blockCtx)+1+padLen)
	payload = append(payload, blockCtx...)
	payload = append(payload, 0x00)
	payload = append(payload, pad...)

	compressed, err := zstd.CompressBatchBytes(payload)
	require.NoError(t, err)
	require.Greater(t, len(compressed), types.MaxBlobBytesSize,
		"multi-blob test requires compressed payload to overflow a single blob")

	blobs := splitCompressedIntoBlobs(t, compressed)
	require.GreaterOrEqual(t, len(blobs), 2, "expected at least 2 blobs for multi-blob path")

	batch := geth.RPCRollupBatch{
		Version:           2,
		ParentBatchHeader: buildV1ParentHeader(parentIndex, startBlock),
		LastBlockNumber:   startBlock + blockCount - 1,
		PrevStateRoot:     common.BigToHash(big.NewInt(1)),
		PostStateRoot:     common.BigToHash(big.NewInt(2)),
		WithdrawRoot:      common.BigToHash(big.NewInt(3)),
		Sidecar:           eth.BlobTxSidecar{Blobs: blobs},
	}

	var bi BatchInfo
	require.NoError(t, bi.ParseBatch(batch))

	require.EqualValues(t, parentIndex+1, bi.batchIndex)
	require.EqualValues(t, 2, bi.version)
	require.EqualValues(t, startBlock, bi.FirstBlockNumber())
	require.EqualValues(t, startBlock+blockCount-1, bi.LastBlockNumber())
	require.Len(t, bi.blockContexts, blockCount)
	for i, bc := range bi.blockContexts {
		require.EqualValues(t, uint64(startBlock+i), bc.Number,
			"block %d number mismatch", i)
		require.EqualValues(t, 1_700_000_000+uint64(i)*6, bc.Timestamp,
			"block %d timestamp mismatch", i)
	}
	require.EqualValues(t, batch.PostStateRoot, bi.root)
	require.EqualValues(t, batch.WithdrawRoot, bi.withdrawalRoot)
}

// TestParseBatchMultiBlobConcatDecompressInvariant directly exercises the
// low-level invariant that multi-blob ParseBatch relies on: the compressed
// stream can only be recovered by concatenating blob bodies in submission
// order and decompressing once. Per-blob decompression must fail on any
// non-initial blob, and reordering blobs must corrupt the decompressed
// output. Keeping this explicit protects the invariant even if ParseBatch is
// later refactored to hide the concatenation step.
func TestParseBatchMultiBlobConcatDecompressInvariant(t *testing.T) {
	pad := make([]byte, types.MaxBlobBytesSize+types.MaxBlobBytesSize/5)
	_, err := rand.Read(pad)
	require.NoError(t, err)

	compressed, err := zstd.CompressBatchBytes(pad)
	require.NoError(t, err)
	require.Greater(t, len(compressed), types.MaxBlobBytesSize)

	blobs := splitCompressedIntoBlobs(t, compressed)
	require.GreaterOrEqual(t, len(blobs), 2)

	// In-order concatenation round-trips.
	var concat []byte
	for i := range blobs {
		body, err := types.RetrieveBlobBytes(&blobs[i])
		require.NoError(t, err)
		concat = append(concat, body...)
	}
	decoded, err := zstd.DecompressBatchBytes(concat)
	require.NoError(t, err)
	require.Equal(t, pad, decoded)

	// Reversing blob order must corrupt the stream; decompression should
	// either error or yield a different payload.
	var reversed []byte
	for i := len(blobs) - 1; i >= 0; i-- {
		body, err := types.RetrieveBlobBytes(&blobs[i])
		require.NoError(t, err)
		reversed = append(reversed, body...)
	}
	if out, err := zstd.DecompressBatchBytes(reversed); err == nil {
		require.NotEqual(t, pad, out,
			"reversed-blob decompression unexpectedly matched payload")
	}
}
