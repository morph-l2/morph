package types

import (
	"crypto/rand"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	gokzg4844 "github.com/crate-crypto/go-eth-kzg"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// randFieldElement generates a random valid field element for BLS12-381
func randFieldElement() [32]byte {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to get random field element")
	}
	var r fr.Element
	r.SetBytes(bytes)
	return gokzg4844.SerializeScalar(r)
}

// randBlob creates a random valid blob
func randBlob() *kzg4844.Blob {
	var blob kzg4844.Blob
	for i := 0; i < len(blob); i += gokzg4844.SerializedScalarSize {
		fieldElementBytes := randFieldElement()
		copy(blob[i:i+gokzg4844.SerializedScalarSize], fieldElementBytes[:])
	}
	return &blob
}

// emptyBlob creates an empty (zero) blob
func emptyBlob() *kzg4844.Blob {
	return new(kzg4844.Blob)
}

func TestMakeBlobProof_SingleBlob(t *testing.T) {
	// Create a single blob
	blob := randBlob()

	// Generate commitment from blob
	commitment, err := kzg4844.BlobToCommitment(blob)
	require.NoError(t, err, "should create commitment from blob")

	// Generate proofs using MakeBlobProof
	blobs := []kzg4844.Blob{*blob}
	commitments := []kzg4844.Commitment{commitment}

	proofs, err := MakeBlobProof(blobs, commitments)
	require.NoError(t, err, "should generate proofs")
	require.Len(t, proofs, 1, "should generate one proof for one blob")

	// Verify the proof
	err = kzg4844.VerifyBlobProof(blob, commitment, proofs[0])
	assert.NoError(t, err, "proof should verify successfully")
}

func TestMakeBlobProof_MultipleBlobs(t *testing.T) {
	// Create multiple blobs
	numBlobs := 3
	blobs := make([]kzg4844.Blob, numBlobs)
	commitments := make([]kzg4844.Commitment, numBlobs)

	// Generate commitments for each blob
	for i := 0; i < numBlobs; i++ {
		blob := randBlob()
		blobs[i] = *blob

		commitment, err := kzg4844.BlobToCommitment(blob)
		require.NoError(t, err, "should create commitment from blob")
		commitments[i] = commitment
	}

	// Generate proofs using MakeBlobProof
	proofs, err := MakeBlobProof(blobs, commitments)
	require.NoError(t, err, "should generate proofs")
	require.Len(t, proofs, numBlobs, "should generate one proof per blob")

	// Verify each proof
	for i := 0; i < numBlobs; i++ {
		err = kzg4844.VerifyBlobProof(&blobs[i], commitments[i], proofs[i])
		assert.NoError(t, err, "proof %d should verify successfully", i)
	}
}

func TestMakeBlobProof_EmptyBlob(t *testing.T) {
	// Test with empty (zero) blob
	blob := emptyBlob()

	commitment, err := kzg4844.BlobToCommitment(blob)
	require.NoError(t, err, "should create commitment from empty blob")

	blobs := []kzg4844.Blob{*blob}
	commitments := []kzg4844.Commitment{commitment}

	proofs, err := MakeBlobProof(blobs, commitments)
	require.NoError(t, err, "should generate proof for empty blob")
	require.Len(t, proofs, 1, "should generate one proof")

	// Verify the proof
	err = kzg4844.VerifyBlobProof(blob, commitment, proofs[0])
	assert.NoError(t, err, "proof for empty blob should verify successfully")
}

func TestMakeBlobProof_InvalidCommitment(t *testing.T) {
	// Test with blob and mismatched commitment
	blob := randBlob()
	// Create a different blob and its commitment
	otherBlob := randBlob()
	wrongCommitment, err := kzg4844.BlobToCommitment(otherBlob)
	require.NoError(t, err)

	blobs := []kzg4844.Blob{*blob}
	commitments := []kzg4844.Commitment{wrongCommitment}

	// MakeBlobProof should succeed (it doesn't validate commitment matches blob)
	proofs, err := MakeBlobProof(blobs, commitments)
	require.NoError(t, err, "MakeBlobProof should succeed even with wrong commitment")
	require.Len(t, proofs, 1)

	// But verification should fail
	err = kzg4844.VerifyBlobProof(blob, wrongCommitment, proofs[0])
	assert.Error(t, err, "verification should fail with mismatched commitment")
}

func TestMakeBlobProof_EmptySlice(t *testing.T) {
	// Test with empty slices
	blobs := []kzg4844.Blob{}
	commitments := []kzg4844.Commitment{}

	proofs, err := MakeBlobProof(blobs, commitments)
	require.NoError(t, err, "should handle empty slices")
	require.Len(t, proofs, 0, "should return empty proofs slice")
}

func TestMakeCellProof_SingleBlob(t *testing.T) {
	// Create a single blob
	blob := randBlob()

	// Generate commitment from blob
	commitment, err := kzg4844.BlobToCommitment(blob)
	require.NoError(t, err, "should create commitment from blob")

	// Generate cell proofs using MakeCellProof
	blobs := []kzg4844.Blob{*blob}
	proofs, err := MakeCellProof(blobs)
	require.NoError(t, err, "should generate cell proofs")
	require.Len(t, proofs, kzg4844.CellProofsPerBlob, "should generate %d proofs for one blob", kzg4844.CellProofsPerBlob)

	// Verify the cell proofs
	err = kzg4844.VerifyCellProofs(blobs, []kzg4844.Commitment{commitment}, proofs)
	assert.NoError(t, err, "cell proofs should verify successfully")
}

func TestMakeCellProof_MultipleBlobs(t *testing.T) {
	// Create multiple blobs
	numBlobs := 3
	blobs := make([]kzg4844.Blob, numBlobs)
	commitments := make([]kzg4844.Commitment, numBlobs)

	// Generate commitments for each blob
	for i := 0; i < numBlobs; i++ {
		blob := randBlob()
		blobs[i] = *blob

		commitment, err := kzg4844.BlobToCommitment(blob)
		require.NoError(t, err, "should create commitment from blob")
		commitments[i] = commitment
	}

	// Generate cell proofs using MakeCellProof
	proofs, err := MakeCellProof(blobs)
	require.NoError(t, err, "should generate cell proofs")
	expectedProofCount := numBlobs * kzg4844.CellProofsPerBlob
	require.Len(t, proofs, expectedProofCount, "should generate %d proofs for %d blobs", expectedProofCount, numBlobs)

	// Verify all cell proofs
	err = kzg4844.VerifyCellProofs(blobs, commitments, proofs)
	assert.NoError(t, err, "cell proofs should verify successfully")
}

func TestMakeCellProof_EmptyBlob(t *testing.T) {
	// Test with empty (zero) blob
	blob := emptyBlob()

	commitment, err := kzg4844.BlobToCommitment(blob)
	require.NoError(t, err, "should create commitment from empty blob")

	blobs := []kzg4844.Blob{*blob}
	proofs, err := MakeCellProof(blobs)
	require.NoError(t, err, "should generate cell proofs for empty blob")
	require.Len(t, proofs, kzg4844.CellProofsPerBlob, "should generate %d proofs", kzg4844.CellProofsPerBlob)

	// Verify the cell proofs
	err = kzg4844.VerifyCellProofs(blobs, []kzg4844.Commitment{commitment}, proofs)
	assert.NoError(t, err, "cell proofs for empty blob should verify successfully")
}

func TestMakeCellProof_EmptySlice(t *testing.T) {
	// Test with empty slice
	blobs := []kzg4844.Blob{}

	proofs, err := MakeCellProof(blobs)
	require.NoError(t, err, "should handle empty slice")
	require.Len(t, proofs, 0, "should return empty proofs slice")
}

func TestMakeCellProof_ProofCount(t *testing.T) {
	// Test that each blob generates exactly CellProofsPerBlob proofs
	testCases := []struct {
		name     string
		numBlobs int
	}{
		{"single blob", 1},
		{"two blobs", 2},
		{"five blobs", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			blobs := make([]kzg4844.Blob, tc.numBlobs)
			for i := 0; i < tc.numBlobs; i++ {
				blob := randBlob()
				blobs[i] = *blob
			}

			proofs, err := MakeCellProof(blobs)
			require.NoError(t, err)
			expectedCount := tc.numBlobs * kzg4844.CellProofsPerBlob
			assert.Equal(t, expectedCount, len(proofs), "should generate %d proofs for %d blobs", expectedCount, tc.numBlobs)
		})
	}
}

func TestMakeCellProof_ProofGrouping(t *testing.T) {
	// Test that proofs are correctly grouped by blob
	// Each blob's proofs should be consecutive in the proofs slice
	numBlobs := 2
	blobs := make([]kzg4844.Blob, numBlobs)
	commitments := make([]kzg4844.Commitment, numBlobs)

	for i := 0; i < numBlobs; i++ {
		blob := randBlob()
		blobs[i] = *blob

		commitment, err := kzg4844.BlobToCommitment(blob)
		require.NoError(t, err)
		commitments[i] = commitment
	}

	// Generate all proofs
	allProofs, err := MakeCellProof(blobs)
	require.NoError(t, err)

	// Verify each blob's proofs individually
	for i := 0; i < numBlobs; i++ {
		startIdx := i * kzg4844.CellProofsPerBlob
		endIdx := startIdx + kzg4844.CellProofsPerBlob
		blobProofs := allProofs[startIdx:endIdx]

		// Verify this blob's proofs
		err = kzg4844.VerifyCellProofs(
			[]kzg4844.Blob{blobs[i]},
			[]kzg4844.Commitment{commitments[i]},
			blobProofs,
		)
		assert.NoError(t, err, "blob %d proofs should verify successfully", i)
	}
}
