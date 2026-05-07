package batch

import (
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"morph-l2/common/blob"
)

// TestBatchHeaderV2 exercises the V2 header variant: it reuses the V1 wire
// layout (257 bytes) but the 32-byte field at offset 57 carries an aggregated
// blob hash (keccak256(blobhash(0)||...||blobhash(N-1))) rather than a single
// versioned hash. Parsing helpers must route callers accordingly.
func TestBatchHeaderV2(t *testing.T) {
	aggregated := common.BigToHash(big.NewInt(0xABCDEF))

	// Start from a V1 encoding (identical byte layout), then flip the version
	// byte to V2. This matches the on-chain behavior where a V2 header is
	// produced by tx-submitter with the aggregated hash stored at offset 57.
	raw := BatchHeaderV1{
		BatchHeaderV0: BatchHeaderV0{
			BatchIndex:             42,
			L1MessagePopped:        1,
			TotalL1MessagePopped:   3,
			DataHash:               common.BigToHash(big.NewInt(0x11)),
			BlobVersionedHash:      aggregated,
			PrevStateRoot:          common.BigToHash(big.NewInt(0x22)),
			PostStateRoot:          common.BigToHash(big.NewInt(0x33)),
			WithdrawalRoot:         common.BigToHash(big.NewInt(0x44)),
			SequencerSetVerifyHash: common.BigToHash(big.NewInt(0x55)),
			ParentBatchHash:        common.BigToHash(big.NewInt(0x66)),
		},
		LastBlockNumber: 9_876,
	}.Bytes()
	raw[0] = BatchHeaderVersion2

	version, err := raw.Version()
	require.NoError(t, err)
	require.EqualValues(t, BatchHeaderVersion2, version)

	batchIndex, err := raw.BatchIndex()
	require.NoError(t, err)
	require.EqualValues(t, 42, batchIndex)

	lastBlockNumber, err := raw.LastBlockNumber()
	require.NoError(t, err)
	require.EqualValues(t, 9_876, lastBlockNumber)

	// V2 headers must route callers to BlobHashesHash; the legacy accessor
	// intentionally errors to prevent silent mis-use.
	_, err = raw.BlobVersionedHash()
	require.Error(t, err)

	aggHash, err := raw.BlobHashesHash()
	require.NoError(t, err)
	require.EqualValues(t, aggregated, aggHash)

	// Length check: a V2 header with the wrong length must fail validate().
	short := make(BatchHeaderBytes, expectedLengthV2-1)
	short[0] = BatchHeaderVersion2
	_, err = short.BatchIndex()
	require.ErrorIs(t, err, ErrInvalidBatchHeaderLength)
}

// TestBlobHashesHashUnavailableForLegacy guards the inverse direction: V0 and
// V1 headers must reject BlobHashesHash so that callers reach for the correct
// accessor.
func TestBlobHashesHashUnavailableForLegacy(t *testing.T) {
	v0 := BatchHeaderV0{
		BatchIndex:        1,
		BlobVersionedHash: blob.EmptyVersionedHash,
	}.Bytes()
	_, err := v0.BlobHashesHash()
	require.Error(t, err)

	v1 := BatchHeaderV1{
		BatchHeaderV0: BatchHeaderV0{
			BatchIndex:        2,
			BlobVersionedHash: blob.EmptyVersionedHash,
		},
		LastBlockNumber: 10,
	}.Bytes()
	_, err = v1.BlobHashesHash()
	require.Error(t, err)
}
