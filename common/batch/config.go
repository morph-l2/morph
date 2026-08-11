package batch

import (
	"encoding/binary"
	"errors"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
)

const batchConfigHashDomain = "morph.batch-config.v1"

// BatchConfig contains the two static batching thresholds used by a submitter.
// BlockInterval is measured in L2 blocks and Timeout is measured in seconds.
// Either threshold may be disabled independently, but disabling both would make
// it impossible to guarantee that an in-progress batch is ever sealed.
type BatchConfig struct {
	BlockInterval uint64
	Timeout       uint64
}

// Validate checks the protocol-level zero-value rule for static batching
// configuration.
func (c BatchConfig) Validate() error {
	if c.BlockInterval == 0 && c.Timeout == 0 {
		return errors.New("batch block interval and timeout cannot both be zero")
	}
	return nil
}

// Hash returns a stable, domain-separated hash suitable for comparing the
// effective configuration across preflight, startup and backup instances.
func (c BatchConfig) Hash() common.Hash {
	encoded := make([]byte, len(batchConfigHashDomain)+16)
	copy(encoded, batchConfigHashDomain)
	binary.BigEndian.PutUint64(encoded[len(batchConfigHashDomain):], c.BlockInterval)
	binary.BigEndian.PutUint64(encoded[len(batchConfigHashDomain)+8:], c.Timeout)
	return crypto.Keccak256Hash(encoded)
}
