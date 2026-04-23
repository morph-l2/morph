package l1sequencer

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// Signer manages sequencer identity and signing capabilities.
// It abstracts the private key management, allowing for local key storage
// or remote signing services (e.g., HSM, KMS) in the future.
type Signer interface {
	// Sign signs data with the sequencer's private key
	Sign(data []byte) ([]byte, error)

	// Address returns the sequencer's address
	Address() common.Address
}

// LocalSigner implements Signer with a local private key
type LocalSigner struct {
	privKey *ecdsa.PrivateKey
	address common.Address
	logger  tmlog.Logger
}

// NewLocalSigner creates a new LocalSigner with a local private key
func NewLocalSigner(privKey *ecdsa.PrivateKey, logger tmlog.Logger) (*LocalSigner, error) {
	if privKey == nil {
		return nil, fmt.Errorf("private key is required")
	}

	return &LocalSigner{
		privKey: privKey,
		address: crypto.PubkeyToAddress(privKey.PublicKey),
		logger:  logger.With("module", "signer"),
	}, nil
}

// Sign signs data with the sequencer's private key
func (s *LocalSigner) Sign(data []byte) ([]byte, error) {
	signature, err := crypto.Sign(data, s.privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}
	return signature, nil
}

// Address returns the sequencer's address
func (s *LocalSigner) Address() common.Address {
	return s.address
}

