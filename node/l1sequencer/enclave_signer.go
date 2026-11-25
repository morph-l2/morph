package l1sequencer

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"slices"
	"sort"
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
// The enclave image carries no key configuration: the node names every
// key (KMS key id + SM secret id) in LoadKeys (op=0x04), together with
// the AWS credentials, and signs with Sign(address, hash) (op=0x02).
// Init (op=0x01) is the key-creation flow run by ops-cli; ListKeys
// (op=0x03) is for operators. The node never sends either.
const (
	opSign     byte = 0x02 // [address:20][hash:32]
	opLoadKeys byte = 0x04 // [creds][n:u8] + n × [kms_lp][sid_lp]

	statusOk byte = 0x00

	hashLen    = 32
	sigLen     = 65
	pubkeyLen  = 33
	addressLen = 20
	maxErrMsg  = 1024
	// maxIDLen caps a KMS key id / SM secret id (crates/protocol
	// MAX_SECRET_ID_LEN).
	maxIDLen = 512

	// dialTimeout bounds vsock connect. mdlayher/vsock has no
	// context/timeout-aware Dial, so we wrap it in a goroutine and
	// race against a timer; without this the constructor and any
	// reconnect on the sign hot path could block forever if the
	// enclave accepted the FD but never sends.
	dialTimeout = 3 * time.Second
	// requestTimeout bounds one Sign round-trip on an established conn.
	// Applied via SetDeadline at the start of each RPC, cleared on
	// exit so the persistent conn can be reused for the next call.
	requestTimeout = 1 * time.Second
	// perKeyLoadTimeout budgets one key's SM Get + KMS Decrypt inside
	// LoadKeys; the whole request gets this times the key count. Keys the
	// enclave already holds answer from memory and cost nothing.
	perKeyLoadTimeout = 5 * time.Second
)

// Env var names for the enclave key config: one KMS key id / SM secret id
// pair per key, numbered by <n>, e.g.
//
//	MORPH_NODE_SEQUENCER_ENCLAVE_KMS_KEY_ID_0=arn:aws:kms:...
//	MORPH_NODE_SEQUENCER_ENCLAVE_SECRET_ID_0=arn:aws:secretsmanager:...
const (
	EnclaveKMSKeyIDEnvPrefix   = "MORPH_NODE_SEQUENCER_ENCLAVE_KMS_KEY_ID_"
	EnclaveSecretIDEnvPrefix   = "MORPH_NODE_SEQUENCER_ENCLAVE_SECRET_ID_"
	enclaveKeyEnvPrefixSummary = "MORPH_NODE_SEQUENCER_ENCLAVE_{KMS_KEY_ID,SECRET_ID}_<n>"
)

// EnclaveKey names one key the enclave should load.
type EnclaveKey struct {
	KMSKeyID string
	SecretID string
}

// ParseEnclaveKeys collects the enclave key config from environment
// entries ("K=V", as from os.Environ). Every <n> that appears is one key;
// keys are ordered by <n>, which need not be contiguous. <n> must be a
// plain decimal number ("0", "1", "10"; not "01"). An <n> with only one
// half of the pair (or an empty value) is an error, as is having no key
// at all.
func ParseEnclaveKeys(environ []string) ([]EnclaveKey, error) {
	type pair struct{ kms, sid string }
	pairs := map[uint64]*pair{}
	for _, kv := range environ {
		k, v, _ := strings.Cut(kv, "=")
		var suffix string
		var isKMS bool
		switch {
		case strings.HasPrefix(k, EnclaveKMSKeyIDEnvPrefix):
			suffix, isKMS = strings.TrimPrefix(k, EnclaveKMSKeyIDEnvPrefix), true
		case strings.HasPrefix(k, EnclaveSecretIDEnvPrefix):
			suffix = strings.TrimPrefix(k, EnclaveSecretIDEnvPrefix)
		default:
			continue
		}
		// Canonical decimal only: "01" and "1" would otherwise both parse
		// to 1 and one pair would silently overwrite the other.
		n, err := strconv.ParseUint(suffix, 10, 32)
		if err != nil || strconv.FormatUint(n, 10) != suffix {
			return nil, fmt.Errorf("%s: suffix must be a decimal number without leading zeros", k)
		}
		p := pairs[n]
		if p == nil {
			p = &pair{}
			pairs[n] = p
		}
		if isKMS {
			p.kms = v
		} else {
			p.sid = v
		}
	}
	if len(pairs) == 0 {
		return nil, fmt.Errorf("no enclave key configured: set %s", enclaveKeyEnvPrefixSummary)
	}
	ns := make([]uint64, 0, len(pairs))
	for n := range pairs {
		ns = append(ns, n)
	}
	sort.Slice(ns, func(i, j int) bool { return ns[i] < ns[j] })
	keys := make([]EnclaveKey, 0, len(ns))
	for _, n := range ns {
		p := pairs[n]
		if p.kms == "" || p.sid == "" {
			return nil, fmt.Errorf("%s%d and %s%d must both be set", EnclaveKMSKeyIDEnvPrefix, n, EnclaveSecretIDEnvPrefix, n)
		}
		keys = append(keys, EnclaveKey{KMSKeyID: p.kms, SecretID: p.sid})
	}
	return keys, nil
}

// EnclaveSigner implements Signer by talking to a Nitro Enclave signer
// over vsock. One persistent connection is reused across all sign calls;
// on any wire error the next call transparently reconnects.
//
// The node owns the key config (which KMS key / SM secret backs each
// key) and pushes it on every new connection with LoadKeys, so a
// restarted enclave — which has lost its in-memory keys — is reloaded
// without a node restart.
//
// The signer's identity — the address of every configured key — is fixed
// by the first connection and never changes afterwards: every later
// connection must load exactly the same addresses or it is refused.
// Changing keys means changing the config and restarting the node.
type EnclaveSigner struct {
	cid    uint32
	port   uint32
	logger tmlog.Logger
	keys   []EnclaveKey
	// dialFn opens a fresh connection to the enclave (vsock in production);
	// credsFn resolves the AWS credentials sent with LoadKeys (IRSA, nil
	// when absent). Both replaceable in tests.
	dialFn  func() (net.Conn, error)
	credsFn func() (*awsCreds, error)

	// addrs[i] is the address of keys[i]. Set once by the first connection
	// and read-only from then on, so it needs no lock.
	addrs []common.Address

	// mu guards conn and is held for a whole sign (including a reconnect).
	mu   sync.Mutex
	conn net.Conn
}

// NewEnclaveSigner connects to the enclave and has it load every key in
// keys (see connect). Any failure is fatal: the caller (node startup in
// cmd/node/main.go) exits before tendermint comes up. Starting with only
// some keys could make this node leader with no valid key at a rotation
// height.
func NewEnclaveSigner(addr string, keys []EnclaveKey, logger tmlog.Logger) (*EnclaveSigner, error) {
	cid, port, err := parseAddr(addr)
	if err != nil {
		return nil, err
	}
	s := &EnclaveSigner{
		cid:     cid,
		port:    port,
		keys:    keys,
		logger:  logger.With("module", "signer", "kind", "enclave", "addr", addr),
		credsFn: resolveCreds,
	}
	s.dialFn = s.dialVsock
	if err := s.connect(); err != nil {
		return nil, fmt.Errorf("enclave signer %s: %w", addr, err)
	}
	return s, nil
}

// connect opens a connection on which the enclave holds every configured
// key: dial, LoadKeys (idempotent enclave-side: a key still in memory is
// answered from there, a restarted enclave loads it again), then a
// signature self-test per key.
//
// It is all-or-nothing: a key that fails to load, two keys with the same
// address or a failed self-test fail the whole connection, on startup and
// on every reconnect alike. The first connection fixes the signer's
// identity (s.addrs); a later one that loads different addresses is
// refused, so a reconnect can only ever restore the startup state.
//
// Caller holds s.mu, or is the constructor.
func (s *EnclaveSigner) connect() (err error) {
	creds, err := s.credsFn()
	if err != nil {
		return err
	}
	conn, err := s.dialFn()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	results, err := loadKeysOnce(conn, creds, s.keys)
	if err != nil {
		return fmt.Errorf("LoadKeys: %w", err)
	}
	if len(results) != len(s.keys) {
		return fmt.Errorf("LoadKeys: %d results for %d keys", len(results), len(s.keys))
	}
	addrs := make([]common.Address, len(results))
	for i, r := range results {
		if !r.ok() {
			return fmt.Errorf("key %d (%s) failed to load: status=0x%02X %s", i, s.keys[i].SecretID, r.status, r.msg)
		}
		if j := slices.Index(addrs[:i], r.address); j >= 0 {
			return fmt.Errorf("keys %d and %d resolve to the same address %s", j, i, r.address.Hex())
		}
		addrs[i] = r.address
	}
	if s.addrs != nil && !slices.Equal(addrs, s.addrs) {
		return fmt.Errorf("enclave keys changed since startup: loaded %v, expected %v", AddrHexes(addrs), AddrHexes(s.addrs))
	}
	if err := verifySigningIdentity(conn, addrs); err != nil {
		return fmt.Errorf("self-test: %w", err)
	}

	s.conn = conn
	if s.addrs == nil {
		s.addrs = addrs
	}
	return nil
}

// verifySigningIdentity asks the enclave to sign a 32-byte zero hash with
// every key in addrs, recovers the public key from each signature, derives
// the EVM address, and returns an error if any doesn't match. This catches
// a keystore whose claimed address disagrees with what its signature
// actually recovers to, and any layer between the SDK and the wire
// (vsock-proxy, in-enclave stub) tampering with payloads.
func verifySigningIdentity(conn net.Conn, addrs []common.Address) error {
	var zero [hashLen]byte
	for _, addr := range addrs {
		sig, err := signOnce(conn, addr, zero[:])
		if err != nil {
			return fmt.Errorf("self-test sign with %s: %w", addr.Hex(), err)
		}
		pub, err := crypto.SigToPub(zero[:], sig)
		if err != nil {
			return fmt.Errorf("self-test recover pubkey for %s: %w", addr.Hex(), err)
		}
		if derived := crypto.PubkeyToAddress(*pub); derived != addr {
			return fmt.Errorf("address mismatch: enclave reported %s but its signature recovers to %s",
				addr.Hex(), derived.Hex())
		}
	}
	return nil
}

// Addresses returns the address of every configured key, in config order.
func (s *EnclaveSigner) Addresses() []common.Address {
	return slices.Clone(s.addrs)
}

// HasAddress reports whether addr is one of the signer's keys.
func (s *EnclaveSigner) HasAddress(addr common.Address) bool {
	return slices.Contains(s.addrs, addr)
}

// SignWith signs a 32-byte prehash with the key of addr — the address the
// L1 contract designates for the block being produced. Returns the
// 65-byte [r||s||v] signature.
//
// Up to 3 attempts: each uses the cached conn or reconnects (see connect),
// and drops the conn on any error so the next attempt starts fresh. After
// 3 failures returns an error wrapping the last cause; the caller abandons
// this block round, and the next sign starts over with a reconnect.
func (s *EnclaveSigner) SignWith(addr common.Address, data []byte) ([]byte, error) {
	if len(data) != hashLen {
		return nil, fmt.Errorf("enclave signer: hash must be %d bytes, got %d", hashLen, len(data))
	}
	if !s.HasAddress(addr) {
		return nil, fmt.Errorf("enclave signer has no key for %s", addr.Hex())
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	const maxAttempts = 3
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if s.conn == nil {
			if lastErr = s.connect(); lastErr != nil {
				s.logger.Error("enclave reconnect failed", "attempt", attempt, "max", maxAttempts, "err", lastErr)
				continue
			}
			s.logger.Info("enclave signer reconnected", "addresses", AddrHexes(s.addrs))
		}
		sig, err := signOnce(s.conn, addr, data)
		if err == nil {
			return sig, nil
		}
		lastErr = err
		s.logger.Info("sign attempt failed",
			"attempt", attempt, "max", maxAttempts, "signer", addr.Hex(), "err", err)
		s.dropConn()
	}
	s.logger.Error("sign exhausted all retries", "attempts", maxAttempts, "signer", addr.Hex(), "err", lastErr)
	return nil, fmt.Errorf("enclave sign with %s failed after %d attempts: %w", addr.Hex(), maxAttempts, lastErr)
}

func (s *EnclaveSigner) dropConn() {
	if s.conn != nil {
		_ = s.conn.Close()
		s.conn = nil
	}
}

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

// keyResult is one LoadKeys outcome, in request order.
type keyResult struct {
	status  byte
	address common.Address
	msg     string
}

func (r keyResult) ok() bool { return r.status == statusOk }

// loadKeysOnce sends LoadKeys(creds, keys) on conn and decodes the
// per-key results.
func loadKeysOnce(conn net.Conn, creds *awsCreds, keys []EnclaveKey) ([]keyResult, error) {
	if len(keys) > 255 {
		return nil, fmt.Errorf("too many enclave keys: %d", len(keys))
	}
	if err := conn.SetDeadline(time.Now().Add(time.Duration(len(keys)) * perKeyLoadTimeout)); err != nil {
		return nil, fmt.Errorf("set deadline: %w", err)
	}
	defer func() { _ = conn.SetDeadline(time.Time{}) }()

	req, err := appendCreds([]byte{opLoadKeys}, creds)
	if err != nil {
		return nil, err
	}
	req = append(req, byte(len(keys)))
	for _, k := range keys {
		if len(k.KMSKeyID) > maxIDLen || len(k.SecretID) > maxIDLen {
			return nil, fmt.Errorf("kms key id / secret id longer than %d bytes", maxIDLen)
		}
		req = appendLPString(req, k.KMSKeyID)
		req = appendLPString(req, k.SecretID)
	}
	if _, err := conn.Write(req); err != nil {
		return nil, fmt.Errorf("write LoadKeys: %w", err)
	}
	if err := readOkStatus(conn); err != nil {
		return nil, err
	}
	var n [1]byte
	if _, err := io.ReadFull(conn, n[:]); err != nil {
		return nil, fmt.Errorf("read result count: %w", err)
	}
	results := make([]keyResult, n[0])
	body := make([]byte, 1+pubkeyLen+addressLen)
	for i := range results {
		if _, err := io.ReadFull(conn, body); err != nil {
			return nil, fmt.Errorf("read result %d: %w", i, err)
		}
		msg, err := readLPString(conn, maxErrMsg)
		if err != nil {
			return nil, fmt.Errorf("read result %d msg: %w", i, err)
		}
		results[i] = keyResult{
			status:  body[0],
			address: common.BytesToAddress(body[1+pubkeyLen:]),
			msg:     msg,
		}
	}
	return results, nil
}

// signOnce sends one Sign(addr, data) round-trip on conn.
func signOnce(conn net.Conn, addr common.Address, data []byte) ([]byte, error) {
	if err := conn.SetDeadline(time.Now().Add(requestTimeout)); err != nil {
		return nil, err
	}
	defer func() { _ = conn.SetDeadline(time.Time{}) }()

	req := make([]byte, 0, 1+addressLen+hashLen)
	req = append(req, opSign)
	req = append(req, addr.Bytes()...)
	req = append(req, data...)
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

// dialVsock wraps vsock.Dial with dialTimeout. mdlayher/vsock has no
// context-aware Dial, so we run it in a goroutine and race against a
// timer; if Dial returns after we've given up, the late conn is
// closed in a background goroutine to avoid an FD leak.
func (s *EnclaveSigner) dialVsock() (net.Conn, error) {
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
	// Accept both "vsock:CID:port" (matches ops-cli's --addr convention)
	// and bare "CID:port" (legacy — kept so existing systemd Environment
	// / k8s ConfigMap entries keep working through the format unification).
	s := strings.TrimPrefix(addr, "vsock:")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("addr must be CID:port or vsock:CID:port, got %q", addr)
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

// readLPString reads a [len:u16][utf8] field capped at max.
func readLPString(conn net.Conn, max int) (string, error) {
	var l [2]byte
	if _, err := io.ReadFull(conn, l[:]); err != nil {
		return "", err
	}
	n := int(binary.BigEndian.Uint16(l[:]))
	if n > max {
		return "", fmt.Errorf("field len %d exceeds max %d", n, max)
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(conn, b); err != nil {
		return "", err
	}
	return string(b), nil
}

func readOkStatus(conn net.Conn) error {
	var statusBuf [1]byte
	if _, err := io.ReadFull(conn, statusBuf[:]); err != nil {
		return fmt.Errorf("read status: %w", err)
	}
	if statusBuf[0] == statusOk {
		return nil
	}
	msg, err := readLPString(conn, maxErrMsg)
	if err != nil {
		return fmt.Errorf("enclave signer status=0x%02X (read msg: %w)", statusBuf[0], err)
	}
	return &enclaveError{status: statusBuf[0], msg: msg}
}

// enclaveError is a well-formed non-OK status frame from the enclave
// (status byte + message).
type enclaveError struct {
	status byte
	msg    string
}

func (e *enclaveError) Error() string {
	return fmt.Sprintf("enclave signer error: status=0x%02X msg=%q", e.status, e.msg)
}
