package l1sequencer

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mdlayher/vsock"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// Wire protocol constants — must match crates/protocol/src/lib.rs in
// the morph-enclave-signer repo. Big-endian throughout.
//
// Init (op=0x01) is the key-creation flow handled by ops-cli on a
// throwaway EC2; the morph node never sends it.
const (
	opSign      byte = 0x02
	opGetPubkey byte = 0x03

	statusOk byte = 0x00

	hashLen    = 32
	sigLen     = 65
	pubkeyLen  = 33
	addressLen = 20
	maxErrMsg  = 1024

	// dialTimeout bounds vsock connect. mdlayher/vsock has no
	// context/timeout-aware Dial, so we wrap it in a goroutine and
	// race against a timer; without this the constructor and any
	// reconnect on the Sign hot path could block forever if the
	// enclave accepted the FD but never sends.
	dialTimeout = 3 * time.Second
	// requestTimeout bounds one round-trip on an established conn.
	// Applied via SetDeadline at the start of probe/signOnce, cleared
	// on exit so the persistent conn can be reused for the next call.
	requestTimeout = 1 * time.Second
)

// EnclaveSigner implements Signer by talking to a Nitro Enclave signer
// over vsock. One persistent connection is reused across all Sign
// calls; on any wire error the next call transparently reconnects.
//
// Address format: "CID:port" (e.g. "16:5000"). The enclave's EVM
// address is fetched at startup via GetPubkey and cached.
type EnclaveSigner struct {
	cid     uint32
	port    uint32
	address common.Address
	logger  tmlog.Logger

	mu   sync.Mutex
	conn net.Conn
}

// NewEnclaveSigner dials the enclave, fetches the loaded EVM address
// via GetPubkey, then runs a signature self-test (sign a zero hash,
// recover the address from the resulting signature, assert it matches).
// The self-test catches misconfig where the enclave reports one
// address but signs with a key for a different one — wrong SECRET_ID
// baked into the .eif, vsock-proxy MITM swap, etc.
//
// Failure here is fatal: the caller (node startup in cmd/node/main.go)
// propagates the error and the process exits before tendermint comes
// up. The self-test mismatch case panics with the recovered vs.
// reported addresses; other init errors return through the caller.
func NewEnclaveSigner(addr string, logger tmlog.Logger) (*EnclaveSigner, error) {
	cid, port, err := parseAddr(addr)
	if err != nil {
		return nil, err
	}
	s := &EnclaveSigner{
		cid:    cid,
		port:   port,
		logger: logger.With("module", "signer", "kind", "enclave", "addr", addr),
	}
	if err := s.probe(); err != nil {
		s.logger.Error("enclave signer init failed; node cannot start", "err", err)
		return nil, fmt.Errorf("enclave signer probe %s: %w", addr, err)
	}
	if err := s.verifySigningIdentity(); err != nil {
		s.logger.Error("enclave signer self-test failed; node cannot start", "err", err)
		_ = s.Close()
		return nil, fmt.Errorf("enclave signer self-test: %w", err)
	}
	s.logger.Info("enclave signer ready", "address", s.address.Hex())
	return s, nil
}

// Sign signs a 32-byte prehash. Returns the 65-byte [r||s||v] signature.
//
// Up to 3 attempts: each tries the cached conn (or dials fresh if
// none), and on any wire error drops the conn and retries. After 3
// failures returns an error wrapping the last cause; the caller is
// expected to treat that as fatal (block production can't proceed
// without a signer, so tendermint will halt consensus).
func (s *EnclaveSigner) Sign(data []byte) ([]byte, error) {
	if len(data) != hashLen {
		return nil, fmt.Errorf("enclave signer: hash must be %d bytes, got %d", hashLen, len(data))
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	const maxAttempts = 3
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if s.conn == nil {
			conn, err := s.dial()
			if err != nil {
				lastErr = err
				s.logger.Info("sign attempt failed at dial",
					"attempt", attempt, "max", maxAttempts, "err", err)
				continue
			}
			s.conn = conn
		}
		sig, err := s.signOnce(s.conn, data)
		if err == nil {
			return sig, nil
		}
		lastErr = err
		s.logger.Info("sign attempt failed",
			"attempt", attempt, "max", maxAttempts, "err", err)
		_ = s.conn.Close()
		s.conn = nil
	}
	s.logger.Error("sign exhausted all retries", "attempts", maxAttempts, "err", lastErr)
	return nil, fmt.Errorf("enclave sign failed after %d attempts: %w", maxAttempts, lastErr)
}

// Address returns the enclave's EVM address (fetched at construction).
func (s *EnclaveSigner) Address() common.Address { return s.address }

// Close releases the persistent connection. Safe to call multiple times.
func (s *EnclaveSigner) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.conn == nil {
		return nil
	}
	err := s.conn.Close()
	s.conn = nil
	return err
}

func (s *EnclaveSigner) probe() error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	success := false
	defer func() {
		if !success {
			_ = conn.Close()
		}
	}()
	if err := conn.SetDeadline(time.Now().Add(requestTimeout)); err != nil {
		return fmt.Errorf("set deadline: %w", err)
	}
	defer conn.SetDeadline(time.Time{}) // clear on exit so signOnce can manage its own deadline

	if _, err := conn.Write([]byte{opGetPubkey}); err != nil {
		return fmt.Errorf("write GetPubkey: %w", err)
	}
	if err := readOkStatus(conn); err != nil {
		return err
	}
	body := make([]byte, pubkeyLen+addressLen)
	if _, err := io.ReadFull(conn, body); err != nil {
		return fmt.Errorf("read pubkey+address: %w", err)
	}
	copy(s.address[:], body[pubkeyLen:])
	s.conn = conn
	success = true
	return nil
}

// verifySigningIdentity asks the enclave to sign a 32-byte zero hash,
// recovers the public key from the signature using secp256k1 ECDSA
// recovery, derives the EVM address, and returns an error if it
// doesn't match the address GetPubkey reported. This catches:
//   - SECRET_ID baked into the .eif points at a different bundle than
//     the operator believes
//   - keystore loaded a key whose claimed address disagrees with what
//     the signature actually recovers to
//   - any layer between the SDK and the wire (vsock-proxy, in-enclave
//     stub) tampering with payloads
//
// Called only from NewEnclaveSigner. A non-nil return aborts node
// startup via main.go's error path; tendermint never sees the bad
// signer, so we don't need to panic for runtime safety.
func (s *EnclaveSigner) verifySigningIdentity() error {
	var zero [hashLen]byte
	sig, err := s.signOnce(s.conn, zero[:])
	if err != nil {
		return fmt.Errorf("self-test sign: %w", err)
	}
	pub, err := crypto.SigToPub(zero[:], sig)
	if err != nil {
		return fmt.Errorf("self-test recover pubkey: %w", err)
	}
	derived := crypto.PubkeyToAddress(*pub)
	if derived != s.address {
		return fmt.Errorf(
			"address mismatch: GetPubkey reported %s but Sign recovers to %s",
			s.address.Hex(), derived.Hex())
	}
	return nil
}

func (s *EnclaveSigner) signOnce(conn net.Conn, data []byte) ([]byte, error) {
	if err := conn.SetDeadline(time.Now().Add(requestTimeout)); err != nil {
		return nil, err
	}
	defer conn.SetDeadline(time.Time{})

	req := make([]byte, 1+hashLen)
	req[0] = opSign
	copy(req[1:], data)
	if _, err := conn.Write(req); err != nil {
		return nil, fmt.Errorf("write Sign: %w", err)
	}
	if err := readOkStatus(conn); err != nil {
		return nil, err
	}
	sig := make([]byte, sigLen)
	if _, err := io.ReadFull(conn, sig); err != nil {
		return nil, fmt.Errorf("read signature: %w", err)
	}
	return sig, nil
}

// dial wraps vsock.Dial with dialTimeout. mdlayher/vsock has no
// context-aware Dial, so we run it in a goroutine and race against a
// timer; if Dial returns after we've given up, the late conn is
// closed in a background goroutine to avoid an FD leak.
func (s *EnclaveSigner) dial() (net.Conn, error) {
	type result struct {
		conn net.Conn
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		c, e := vsock.Dial(s.cid, s.port, nil)
		ch <- result{c, e}
	}()
	select {
	case r := <-ch:
		if r.err != nil {
			return nil, fmt.Errorf("vsock dial %d:%d: %w", s.cid, s.port, r.err)
		}
		return r.conn, nil
	case <-time.After(dialTimeout):
		go func() {
			r := <-ch
			if r.conn != nil {
				_ = r.conn.Close()
			}
		}()
		return nil, fmt.Errorf("vsock dial %d:%d timed out after %v", s.cid, s.port, dialTimeout)
	}
}

func parseAddr(addr string) (uint32, uint32, error) {
	parts := strings.SplitN(addr, ":", 2)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("addr must be CID:port, got %q", addr)
	}
	cid, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid CID %q: %w", parts[0], err)
	}
	port, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid port %q: %w", parts[1], err)
	}
	return uint32(cid), uint32(port), nil
}

func readOkStatus(conn net.Conn) error {
	var statusBuf [1]byte
	if _, err := io.ReadFull(conn, statusBuf[:]); err != nil {
		return fmt.Errorf("read status: %w", err)
	}
	if statusBuf[0] == statusOk {
		return nil
	}
	var lenBuf [2]byte
	if _, err := io.ReadFull(conn, lenBuf[:]); err != nil {
		return fmt.Errorf("enclave signer status=0x%02X (read msg len: %w)", statusBuf[0], err)
	}
	msgLen := binary.BigEndian.Uint16(lenBuf[:])
	if msgLen > maxErrMsg {
		return fmt.Errorf("enclave signer status=0x%02X (msg len %d exceeds max %d)",
			statusBuf[0], msgLen, maxErrMsg)
	}
	msg := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msg); err != nil {
		return fmt.Errorf("enclave signer status=0x%02X (read msg: %w)", statusBuf[0], err)
	}
	return fmt.Errorf("enclave signer error: status=0x%02X msg=%q", statusBuf[0], string(msg))
}
