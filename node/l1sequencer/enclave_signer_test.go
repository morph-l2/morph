package l1sequencer

import (
	"crypto/ecdsa"
	"encoding/binary"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// The production transport is vsock, which CI usually lacks, so the
// wire-protocol tests below run over TCP against fakeEnclave
// (EnclaveSigner's dialFn makes the transport swappable).

func TestParseAddr(t *testing.T) {
	cases := []struct {
		in       string
		wantCID  uint32
		wantPort uint32
		wantErr  bool
	}{
		// Legacy bare CID:port — accepted for backward compat.
		{"16:5000", 16, 5000, false},
		{"3:8000", 3, 8000, false},
		{"0:1", 0, 1, false},
		// vsock:CID:port — matches ops-cli convention.
		{"vsock:16:5000", 16, 5000, false},
		{"vsock:3:8000", 3, 8000, false},
		{"vsock:0:1", 0, 1, false},
		// Errors.
		{"abc:5000", 0, 0, true},
		{"16:bad", 0, 0, true},
		{"16", 0, 0, true},
		{"vsock:16", 0, 0, true},
		{"vsock:abc:5000", 0, 0, true},
		{"", 0, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			cid, port, err := parseAddr(tc.in)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.wantCID, cid)
			require.Equal(t, tc.wantPort, port)
		})
	}
}

// fakeEnclave speaks the signer wire protocol over TCP so EnclaveSigner can
// be exercised without vsock. "Secrets" are a map secret_id → key; LoadKeys
// loads named secrets into memory (additive), Sign signs with a loaded key.
type fakeEnclave struct {
	t  *testing.T
	ln net.Listener

	mu      sync.Mutex
	secrets map[string]*ecdsa.PrivateKey // what SM would return
	loaded  map[string]*ecdsa.PrivateKey // enclave memory
	loads   int                          // LoadKeys requests served
	// lie, if set, makes that address sign with a different key.
	lie common.Address
}

func newFakeEnclave(t *testing.T, secrets map[string]*ecdsa.PrivateKey) *fakeEnclave {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	f := &fakeEnclave{t: t, ln: ln, secrets: secrets, loaded: map[string]*ecdsa.PrivateKey{}}
	go f.serve()
	t.Cleanup(func() { _ = ln.Close() })
	return f
}

// restart models an enclave restart: in-memory keys are lost.
func (f *fakeEnclave) restart() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.loaded = map[string]*ecdsa.PrivateKey{}
}

func (f *fakeEnclave) setSecret(id string, k *ecdsa.PrivateKey) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if k == nil {
		delete(f.secrets, id)
	} else {
		f.secrets[id] = k
	}
}

func (f *fakeEnclave) dial() (net.Conn, error) { return net.Dial("tcp", f.ln.Addr().String()) }

func (f *fakeEnclave) serve() {
	for {
		c, err := f.ln.Accept()
		if err != nil {
			return
		}
		go f.handle(c)
	}
}

func readLP(c net.Conn) string {
	var l [2]byte
	if _, err := io.ReadFull(c, l[:]); err != nil {
		return ""
	}
	b := make([]byte, binary.BigEndian.Uint16(l[:]))
	_, _ = io.ReadFull(c, b)
	return string(b)
}

func writeErr(c net.Conn, status byte, msg string) {
	_, _ = c.Write(appendLPString([]byte{status}, msg))
}

func (f *fakeEnclave) handle(c net.Conn) {
	defer c.Close()
	for {
		var op [1]byte
		if _, err := io.ReadFull(c, op[:]); err != nil {
			return
		}
		switch op[0] {
		case opLoadKeys:
			readLP(c) // akid
			readLP(c) // secret key
			readLP(c) // token
			var exp [8]byte
			var n [1]byte
			if _, err := io.ReadFull(c, exp[:]); err != nil {
				return
			}
			if _, err := io.ReadFull(c, n[:]); err != nil {
				return
			}
			resp := []byte{statusOk, n[0]}
			f.mu.Lock()
			f.loads++
			for i := 0; i < int(n[0]); i++ {
				readLP(c) // kms key id
				sid := readLP(c)
				k := f.loaded[sid]
				if k == nil {
					k = f.secrets[sid]
					if k != nil {
						f.loaded[sid] = k
					}
				}
				if k == nil {
					resp = append(resp, 0x02)
					resp = append(resp, make([]byte, pubkeyLen+addressLen)...)
					resp = appendLPString(resp, "secret not found: "+sid)
					continue
				}
				resp = append(resp, statusOk)
				resp = append(resp, crypto.CompressPubkey(&k.PublicKey)...)
				resp = append(resp, crypto.PubkeyToAddress(k.PublicKey).Bytes()...)
				resp = appendLPString(resp, "")
			}
			f.mu.Unlock()
			_, _ = c.Write(resp)
		case opSign:
			buf := make([]byte, addressLen+hashLen)
			if _, err := io.ReadFull(c, buf); err != nil {
				return
			}
			addr := common.BytesToAddress(buf[:addressLen])
			f.mu.Lock()
			var key *ecdsa.PrivateKey
			for _, k := range f.loaded {
				if crypto.PubkeyToAddress(k.PublicKey) == addr {
					key = k
				}
			}
			empty, lie := len(f.loaded) == 0, f.lie
			f.mu.Unlock()
			switch {
			case empty:
				writeErr(c, 0x02, "no key loaded")
				continue
			case key == nil:
				writeErr(c, 0x06, "no such key")
				continue
			case addr == lie:
				key, _ = crypto.GenerateKey()
			}
			sig, err := crypto.Sign(buf[addressLen:], key)
			require.NoError(f.t, err)
			_, _ = c.Write(append([]byte{statusOk}, sig...))
		default:
			return
		}
	}
}

func newTestEnclaveSigner(t *testing.T, f *fakeEnclave, secretIDs ...string) (*EnclaveSigner, error) {
	keys := make([]EnclaveKey, len(secretIDs))
	for i, id := range secretIDs {
		keys[i] = EnclaveKey{KMSKeyID: "alias/k", SecretID: id}
	}
	s := &EnclaveSigner{
		logger:  tmlog.NewNopLogger(),
		keys:    keys,
		dialFn:  f.dial,
		credsFn: func() (*awsCreds, error) { return &awsCreds{accessKeyID: "AKIA", secretAccessKey: "s"}, nil },
	}
	if err := s.init(); err != nil {
		return nil, err
	}
	t.Cleanup(func() { _ = s.Close() })
	return s, nil
}

func mustKey(t *testing.T) *ecdsa.PrivateKey {
	k, err := crypto.GenerateKey()
	require.NoError(t, err)
	return k
}

func addrOf(k *ecdsa.PrivateKey) common.Address { return crypto.PubkeyToAddress(k.PublicKey) }

func assertSignsAs(t *testing.T, s Signer, addr common.Address) {
	t.Helper()
	hash := crypto.Keccak256([]byte("block"))
	sig, err := s.SignWith(addr, hash)
	require.NoError(t, err)
	pub, err := crypto.SigToPub(hash, sig)
	require.NoError(t, err)
	require.Equal(t, addr, crypto.PubkeyToAddress(*pub))
}

func TestParseEnclaveKeys(t *testing.T) {
	kms, sid := EnclaveKMSKeyIDEnvPrefix, EnclaveSecretIDEnvPrefix
	keys, err := ParseEnclaveKeys([]string{
		sid + "1=s1", "PATH=/bin", kms + "0=k0", kms + "1=k1", sid + "0=s0",
		kms + "10=k10", sid + "10=arn:aws:secretsmanager:r:1:secret:a=b",
	})
	require.NoError(t, err)
	require.Equal(t, []EnclaveKey{
		{KMSKeyID: "k0", SecretID: "s0"},
		{KMSKeyID: "k1", SecretID: "s1"},
		{KMSKeyID: "k10", SecretID: "arn:aws:secretsmanager:r:1:secret:a=b"},
	}, keys)

	for _, bad := range [][]string{
		{},
		{"PATH=/bin"},
		{kms + "0=k0"},
		{kms + "0=k0", sid + "0="},
		{kms + "x=k", sid + "x=s"},
		{kms + "=k", sid + "=s"},
		// Non-canonical suffixes: "01" would collide with "1".
		{kms + "01=k", sid + "01=s"},
		{kms + "1=k1", sid + "1=s1", kms + "01=k01", sid + "01=s01"},
		{kms + "00=k", sid + "00=s"},
		{kms + "+1=k", sid + "+1=s"},
	} {
		_, err := ParseEnclaveKeys(bad)
		require.Error(t, err, "%v", bad)
	}
}

func TestEnclaveSigner_MultiKey(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	s, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.NoError(t, err)
	require.Equal(t, []common.Address{addrOf(a), addrOf(b)}, s.Addresses())

	assertSignsAs(t, s, addrOf(a))
	assertSignsAs(t, s, addrOf(b))

	_, err = s.SignWith(addrOf(mustKey(t)), make([]byte, hashLen))
	require.ErrorContains(t, err, "no loaded key")
	_, err = s.SignWith(addrOf(a), []byte{1, 2, 3})
	require.ErrorContains(t, err, "hash must be")
}

func TestEnclaveSigner_AnyKeyFailingFailsStartup(t *testing.T) {
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": mustKey(t)})
	_, err := newTestEnclaveSigner(t, f, "sa", "missing")
	require.ErrorContains(t, err, "1 of 2 enclave key(s) failed to load")
	require.ErrorContains(t, err, "secret not found: missing")
}

func TestEnclaveSigner_DuplicateAddressFailsStartup(t *testing.T) {
	a := mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"s1": a, "s2": a})
	_, err := newTestEnclaveSigner(t, f, "s1", "s2")
	require.ErrorContains(t, err, "same address")
}

func TestEnclaveSigner_SelfTestCatchesWrongKey(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	f.lie = addrOf(b)
	_, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.ErrorContains(t, err, "address mismatch")
}

func TestEnclaveSigner_ReconnectReloadsKeys(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	s, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.NoError(t, err)

	// Enclave restarted (memory lost): the next sign reconnects, LoadKeys
	// reloads both keys, signing continues — no node restart needed.
	require.NoError(t, s.Close())
	f.restart()
	assertSignsAs(t, s, addrOf(b))
	require.Equal(t, 2, f.loads)
}

func TestEnclaveSigner_ReconnectKeyMissingOnlyDisablesThatKey(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	s, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.NoError(t, err)

	// After a restart, "sa" can no longer be loaded (e.g. its IAM access
	// was revoked): only that key is unusable; "sb" still signs.
	require.NoError(t, s.Close())
	f.restart()
	f.setSecret("sa", nil)
	assertSignsAs(t, s, addrOf(b))
	require.Equal(t, []common.Address{addrOf(b)}, s.Addresses())
	require.False(t, s.HasAddress(addrOf(a)))
	_, err = s.SignWith(addrOf(a), make([]byte, hashLen))
	require.ErrorContains(t, err, "no loaded key")

	// Access restored: the key comes back on the next reconnect.
	f.setSecret("sa", a)
	assertSignsAs(t, s, addrOf(a))
}

func TestEnclaveSigner_ReconnectKeySetNotComparedWithStartup(t *testing.T) {
	a, b, c := mustKey(t), mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	s, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.NoError(t, err)

	// After a restart "sa" holds a different key. The new set is used as
	// is — only the requested (L1-designated) address has to be in it.
	require.NoError(t, s.Close())
	f.restart()
	f.setSecret("sa", c)
	assertSignsAs(t, s, addrOf(b))
	require.Equal(t, []common.Address{addrOf(c), addrOf(b)}, s.Addresses())
	require.True(t, s.HasAddress(addrOf(c)))
	require.False(t, s.HasAddress(addrOf(a)))
	assertSignsAs(t, s, addrOf(c))
}

func TestEnclaveSigner_ReconnectSelfTestStillEnforced(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	f := newFakeEnclave(t, map[string]*ecdsa.PrivateKey{"sa": a, "sb": b})
	s, err := newTestEnclaveSigner(t, f, "sa", "sb")
	require.NoError(t, err)

	// The enclave now signs "sb" with a key other than the one it reports.
	require.NoError(t, s.Close())
	f.restart()
	f.mu.Lock()
	f.lie = addrOf(b)
	f.mu.Unlock()
	_, err = s.SignWith(addrOf(a), make([]byte, hashLen))
	require.ErrorContains(t, err, "self-test after reconnect")
}

func TestLocalSigner_MultiKey(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	s, err := NewLocalSigner(tmlog.NewNopLogger(), a, b)
	require.NoError(t, err)
	require.Equal(t, []common.Address{addrOf(a), addrOf(b)}, s.Addresses())
	require.True(t, s.HasAddress(addrOf(a)))
	require.True(t, s.HasAddress(addrOf(b)))
	assertSignsAs(t, s, addrOf(a))
	assertSignsAs(t, s, addrOf(b))
	other := addrOf(mustKey(t))
	require.False(t, s.HasAddress(other))
	_, err = s.SignWith(other, make([]byte, hashLen))
	require.Error(t, err)

	_, err = NewLocalSigner(tmlog.NewNopLogger(), a, a)
	require.ErrorContains(t, err, "duplicates")
	_, err = NewLocalSigner(tmlog.NewNopLogger())
	require.Error(t, err)
}

func TestParsePrivateKeys(t *testing.T) {
	a, b := mustKey(t), mustKey(t)
	ha := common.Bytes2Hex(crypto.FromECDSA(a))
	hb := common.Bytes2Hex(crypto.FromECDSA(b))

	keys, err := ParsePrivateKeys(ha)
	require.NoError(t, err)
	require.Len(t, keys, 1)

	keys, err = ParsePrivateKeys("0x" + ha + " , " + hb)
	require.NoError(t, err)
	require.Len(t, keys, 2)
	require.Equal(t, addrOf(b), addrOf(keys[1]))

	_, err = ParsePrivateKeys(ha + ",")
	require.ErrorContains(t, err, "#1 is empty")
	_, err = ParsePrivateKeys("zz")
	require.Error(t, err)
	require.NotContains(t, err.Error(), "zz")
}

// TestEnclaveSigner_RealEnclave runs against a real morph-signer built
// without the vsock feature (TCP, file backend). Skipped unless
// ENCLAVE_TCP_ADDR (host:port) and ENCLAVE_TEST_SECRETS (comma-separated
// bundle paths, as written by morph-key-init) are set.
func TestEnclaveSigner_RealEnclave(t *testing.T) {
	addr, secrets := os.Getenv("ENCLAVE_TCP_ADDR"), os.Getenv("ENCLAVE_TEST_SECRETS")
	if addr == "" || secrets == "" {
		t.Skip("ENCLAVE_TCP_ADDR / ENCLAVE_TEST_SECRETS not set")
	}
	var keys []EnclaveKey
	for _, sid := range strings.Split(secrets, ",") {
		keys = append(keys, EnclaveKey{KMSKeyID: "local", SecretID: sid})
	}
	s := &EnclaveSigner{
		logger:  tmlog.NewNopLogger(),
		keys:    keys,
		dialFn:  func() (net.Conn, error) { return net.Dial("tcp", addr) },
		credsFn: func() (*awsCreds, error) { return nil, nil },
	}
	require.NoError(t, s.init())
	defer s.Close()
	require.Len(t, s.Addresses(), len(keys))
	for _, a := range s.Addresses() {
		assertSignsAs(t, s, a)
	}
	// Reconnect path: drop the conn, sign again (LoadKeys answered from memory).
	require.NoError(t, s.Close())
	assertSignsAs(t, s, s.Addresses()[0])
}
