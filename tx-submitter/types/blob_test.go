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
