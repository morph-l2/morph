package services

import (
	"crypto/sha256"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/assert"
)

// TestCalcBlobHashV1VsKZGToVersionedHash verifies that CalcBlobHashV1 and kZGToVersionedHash
// produce the same results for the same commitment.
func TestCalcBlobHashV1VsKZGToVersionedHash(t *testing.T) {
	tests := []struct {
		name       string
		commitment kzg4844.Commitment
	}{
		{
			name:       "zero commitment",
			commitment: kzg4844.Commitment{},
		},
		{
			name: "all ones commitment",
			commitment: func() kzg4844.Commitment {
				var c kzg4844.Commitment
				for i := range c {
					c[i] = 0xFF
				}
				return c
			}(),
		},
		{
			name: "pattern commitment",
			commitment: func() kzg4844.Commitment {
				var c kzg4844.Commitment
				for i := range c {
					c[i] = byte(i % 256)
				}
				return c
			}(),
		},
		{
			name: "reversed pattern commitment",
			commitment: func() kzg4844.Commitment {
				var c kzg4844.Commitment
				for i := range c {
					c[i] = byte(255 - (i % 256))
				}
				return c
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Calculate using kZGToVersionedHash (v0 method)
			hashV0 := kZGToVersionedHash(tt.commitment)

			// Calculate using CalcBlobHashV1 (v1 method)
			hasher := sha256.New()
			hashV1 := kzg4844.CalcBlobHashV1(hasher, &tt.commitment)

			// Convert hashV1 to common.Hash for comparison
			hashV1CommonHash := common.Hash(hashV1)

			// They should be equal
			assert.Equal(t, hashV0, hashV1CommonHash,
				"CalcBlobHashV1 and kZGToVersionedHash should produce the same result")

			// Verify the first byte is 0x01 (version byte)
			assert.Equal(t, uint8(0x01), hashV0[0], "First byte should be version 0x01")
			assert.Equal(t, uint8(0x01), hashV1[0], "First byte should be version 0x01")
		})
	}
}

// TestCalcBlobHashV1ReuseHasher verifies that CalcBlobHashV1 can correctly reuse a hasher
// for multiple commitments.
func TestCalcBlobHashV1ReuseHasher(t *testing.T) {
	// Create multiple commitments
	commitments := []kzg4844.Commitment{
		{0x01, 0x02, 0x03}, // Will be padded with zeros
		{0xFF, 0xFE, 0xFD}, // Will be padded with zeros
		{},                 // Zero commitment
	}

	// Calculate hashes using reused hasher
	hasher := sha256.New()
	hashesV1 := make([]common.Hash, len(commitments))
	for i, commit := range commitments {
		hashV1 := kzg4844.CalcBlobHashV1(hasher, &commit)
		hashesV1[i] = common.Hash(hashV1)
	}

	// Calculate hashes using kZGToVersionedHash (should produce same results)
	hashesV0 := make([]common.Hash, len(commitments))
	for i, commit := range commitments {
		hashesV0[i] = kZGToVersionedHash(commit)
	}

	// Compare results
	for i := range commitments {
		assert.Equal(t, hashesV0[i], hashesV1[i],
			"Hash %d should be equal regardless of hasher reuse", i)
	}
}

// TestCalcBlobHashV1MultipleCalls verifies that CalcBlobHashV1 produces consistent
// results when called multiple times with the same commitment.
func TestCalcBlobHashV1MultipleCalls(t *testing.T) {
	commitment := kzg4844.Commitment{0x42}

	hasher1 := sha256.New()
	hash1 := kzg4844.CalcBlobHashV1(hasher1, &commitment)

	hasher2 := sha256.New()
	hash2 := kzg4844.CalcBlobHashV1(hasher2, &commitment)

	// Same commitment should produce same hash
	assert.Equal(t, hash1, hash2, "Same commitment should produce same hash")

	// Should also match kZGToVersionedHash
	hashV0 := kZGToVersionedHash(commitment)
	assert.Equal(t, common.Hash(hash1), hashV0, "Should match kZGToVersionedHash result")
}
