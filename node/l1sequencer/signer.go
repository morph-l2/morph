package l1sequencer

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// Signer manages the sequencer's signing identity. It may hold several keys
// (e.g. the current and the next sequencer key during a planned rotation);
// the block producer picks, per height, the one address the L1 contract
// designates and signs with exactly that key.
//
// Method set matches tendermint's sequencer.Signer.
type Signer interface {
	// Addresses returns every address this signer can currently sign as,
	// in configuration order. For logging / reporting; block production
	// uses HasAddress.
	Addresses() []common.Address

	// HasAddress reports whether the signer currently holds a key for addr.
	HasAddress(addr common.Address) bool

	// SignWith signs data with the private key of addr. Returns an error
	// if the signer holds no key for addr.
	SignWith(addr common.Address, data []byte) ([]byte, error)
}

// LocalSigner implements Signer with one or more in-memory private keys.
// Intended for devnet / testing; production uses EnclaveSigner.
type LocalSigner struct {
	addrs  []common.Address
	keys   map[common.Address]*ecdsa.PrivateKey
	logger tmlog.Logger
}

// NewLocalSigner creates a LocalSigner holding privKeys, in that order.
// At least one key is required; the same key twice is rejected.
func NewLocalSigner(logger tmlog.Logger, privKeys ...*ecdsa.PrivateKey) (*LocalSigner, error) {
	if len(privKeys) == 0 {
		return nil, fmt.Errorf("at least one private key is required")
	}
	s := &LocalSigner{
		keys:   make(map[common.Address]*ecdsa.PrivateKey, len(privKeys)),
		logger: logger.With("module", "signer"),
	}
	for i, k := range privKeys {
		if k == nil {
			return nil, fmt.Errorf("private key #%d is nil", i)
		}
		addr := crypto.PubkeyToAddress(k.PublicKey)
		if _, dup := s.keys[addr]; dup {
			return nil, fmt.Errorf("private key #%d duplicates address %s", i, addr.Hex())
		}
		s.keys[addr] = k
		s.addrs = append(s.addrs, addr)
	}
	return s, nil
}

// ParsePrivateKeys parses a comma-separated list of hex private keys
// (optional 0x prefix, surrounding whitespace ignored), e.g. the value of
// --sequencer.privateKey.
func ParsePrivateKeys(csv string) ([]*ecdsa.PrivateKey, error) {
	var keys []*ecdsa.PrivateKey
	for i, part := range strings.Split(csv, ",") {
		h := strings.TrimPrefix(strings.TrimSpace(part), "0x")
		if h == "" {
			return nil, fmt.Errorf("private key #%d is empty", i)
		}
		k, err := crypto.HexToECDSA(h)
		if err != nil {
			// Never echo the key material itself.
			return nil, fmt.Errorf("private key #%d is invalid: %w", i, err)
		}
		keys = append(keys, k)
	}
	return keys, nil
}

// AddrHexes renders addrs as hex strings, for logging.
func AddrHexes(addrs []common.Address) []string {
	hexes := make([]string, len(addrs))
	for i, a := range addrs {
		hexes[i] = a.Hex()
	}
	return hexes
}

// Addresses returns the signer's addresses in configuration order.
func (s *LocalSigner) Addresses() []common.Address {
	return append([]common.Address(nil), s.addrs...)
}

// HasAddress reports whether the signer holds a key for addr.
func (s *LocalSigner) HasAddress(addr common.Address) bool {
	_, ok := s.keys[addr]
	return ok
}

// SignWith signs data with the key of addr.
func (s *LocalSigner) SignWith(addr common.Address, data []byte) ([]byte, error) {
	k, ok := s.keys[addr]
	if !ok {
		return nil, fmt.Errorf("local signer has no key for %s", addr.Hex())
	}
	signature, err := crypto.Sign(data, k)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}
	return signature, nil
}
