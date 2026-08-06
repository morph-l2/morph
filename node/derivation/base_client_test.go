package derivation

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/require"
)

// The fallback verifies blob content, so accept-path stubs must serve a real
// (blob, commitment) pair; the all-zero blob is the cheapest valid one.
var (
	hex32 = "0x" + strings.Repeat("00", 32)
	hex48 = "0x" + strings.Repeat("00", 48)

	zeroBlobHex        = "0x" + strings.Repeat("00", BlobSize)
	zeroBlobCommitment = mustZeroBlobCommitment()
	zeroBlobHash       = KZGToVersionedHash(zeroBlobCommitment)

	// Valid blob bytes served under the zero blob's commitment: count and
	// commitment lookup pass, only verifyBlob catches it.
	corruptBlobHex = "0x01" + strings.Repeat("00", BlobSize-1)
)

func mustZeroBlobCommitment() kzg4844.Commitment {
	var blob Blob
	commitment, err := kzg4844.BlobToCommitment(blob.KZGBlob())
	if err != nil {
		panic(err)
	}
	return commitment
}

func sidecarJSON(index int, blobHex, commitmentHex string) string {
	return fmt.Sprintf(`{"block_root":%q,"slot":"1","blob":%q,"index":"%d","kzg_commitment":%q,"kzg_proof":%q}`,
		hex32, blobHex, index, commitmentHex, hex48)
}

// beaconBehavior controls what a stub beacon returns for blob_sidecars.
type beaconBehavior int

const (
	beaconServesBlob             beaconBehavior = iota // 200 with one valid sidecar
	beaconServesEmpty                                  // 200 with an empty list (pruned / not indexed)
	beaconServerError                                  // 500
	beaconServesCorruptBlob                            // 200, right count and commitment, blob bytes do not match
	beaconServesWrongCommitment                        // 200, right count, commitment of some other blob
	beaconServesNullSidecar                            // 200, JSON null entry in the sidecar list
)

// newStubBeacon serves the genesis + spec endpoints (needed for slot math) and
// answers blob_sidecars according to behavior. It returns the base URL and a
// counter of blob_sidecars requests received.
func newStubBeacon(t *testing.T, behavior beaconBehavior) (string, *int32) {
	t.Helper()
	var blobHits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, genesisMethod):
			_, _ = w.Write([]byte(`{"data":{"genesis_time":"0"}}`))
		case strings.Contains(r.URL.Path, specMethod):
			_, _ = w.Write([]byte(`{"data":{"SECONDS_PER_SLOT":"12"}}`))
		case strings.Contains(r.URL.Path, sidecarsMethodPrefix):
			atomic.AddInt32(&blobHits, 1)
			switch behavior {
			case beaconServesBlob:
				_, _ = w.Write([]byte(`{"data":[` + sidecarJSON(0, zeroBlobHex, hexutil.Encode(zeroBlobCommitment[:])) + `]}`))
			case beaconServesEmpty:
				_, _ = w.Write([]byte(`{"data":[]}`))
			case beaconServerError:
				w.WriteHeader(http.StatusInternalServerError)
			case beaconServesCorruptBlob:
				_, _ = w.Write([]byte(`{"data":[` + sidecarJSON(0, corruptBlobHex, hexutil.Encode(zeroBlobCommitment[:])) + `]}`))
			case beaconServesWrongCommitment:
				_, _ = w.Write([]byte(`{"data":[` + sidecarJSON(0, zeroBlobHex, hex48) + `]}`))
			case beaconServesNullSidecar:
				_, _ = w.Write([]byte(`{"data":[null]}`))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(srv.Close)
	return srv.URL, &blobHits
}

func oneHash() []IndexedBlobHash {
	return []IndexedBlobHash{{Index: 0, Hash: zeroBlobHash}}
}

// wantHashes is the verification set the caller always has (tx.BlobHashes()),
// independent of whether index hints could be built.
func wantHashes() []common.Hash {
	return []common.Hash{zeroBlobHash}
}

type canceledHTTP struct {
	calls *int32
}

func (h canceledHTTP) Get(ctx context.Context, _ string, _ http.Header) (*http.Response, error) {
	atomic.AddInt32(h.calls, 1)
	return nil, ctx.Err()
}

func fetch(t *testing.T, c *FallbackBeaconClient) (types.BlobTxSidecar, error) {
	t.Helper()
	return c.GetVerifiedBlobSidecar(context.Background(), L1BlockRef{Time: 12}, wantHashes(), oneHash())
}

// The primary serves the blob and the fallback is never queried.
func TestFallbackBeacon_PrimaryServesBlob(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesBlob)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.EqualValues(t, 1, atomic.LoadInt32(primaryHits))
	require.EqualValues(t, 0, atomic.LoadInt32(fallbackHits), "fallback must not be queried while primary serves the blob")
}

// The key case: primary answers 200 but with NO sidecars (pruned / not indexed).
// This must fall back to a beacon that actually has the blob.
func TestFallbackBeacon_FallsBackOnEmptyResult(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesEmpty)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits))
}

// A 5xx on the primary also falls back.
func TestFallbackBeacon_FallsBackOnServerError(t *testing.T) {
	primary, _ := newStubBeacon(t, beaconServerError)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(fallbackHits))
}

// An unreachable primary (transport error) falls back.
func TestFallbackBeacon_FallsBackOnTransportError(t *testing.T) {
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{"http://127.0.0.1:0", fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(fallbackHits))
}

// 200 with the right count and commitment but corrupted blob bytes must
// trigger fallback instead of handing bad bytes downstream.
func TestFallbackBeacon_FallsBackOnCorruptBlobContent(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesCorruptBlob)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits), "corrupt blob content must trigger fallback")
}

// 200 with the right count but none of the sidecars carries the requested
// hash (e.g. another fork's sidecars at the same slot) must trigger fallback.
func TestFallbackBeacon_FallsBackOnMissingRequestedHash(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesWrongCommitment)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits), "missing requested hash must trigger fallback")
}

// A JSON null in the sidecar list decodes to a nil pointer; it must count as
// a verification failure and trigger fallback, not panic.
func TestFallbackBeacon_FallsBackOnNullSidecar(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesNullSidecar)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecars, err := fetch(t, c)
	require.NoError(t, err)
	require.Len(t, sidecars.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits), "null sidecar must trigger fallback")
}

// The #745 self-heal: when the caller cannot build index hints (the L1 block
// body was unavailable), it passes nil hints. The client must still fetch all
// sidecars at the slot and authenticate them by hash — a nil hint is NOT a
// failure. Guards against turning a BlockByNumber failure back into a hard
// error; see the DO NOT comment in fetchRollupDataByTxHash.
func TestFallbackBeacon_VerifiesWithoutIndexHints(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary}, nil, nil)
	sidecar, err := c.GetVerifiedBlobSidecar(context.Background(), L1BlockRef{Time: 12}, wantHashes(), nil)
	require.NoError(t, err)
	require.Len(t, sidecar.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
}

// Verification stays active on the no-index fetch-all path: a corrupt primary
// must still rotate to a healthy fallback even without index hints.
func TestFallbackBeacon_FallsBackWithoutIndexHintsOnCorruptContent(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesCorruptBlob)
	fallback, fallbackHits := newStubBeacon(t, beaconServesBlob)

	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, nil)
	sidecar, err := c.GetVerifiedBlobSidecar(context.Background(), L1BlockRef{Time: 12}, wantHashes(), nil)
	require.NoError(t, err)
	require.Len(t, sidecar.Blobs, 1)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits), "corrupt content must trigger fallback even without index hints")
}

// A client constructed with no endpoints must error instead of silently
// returning an empty sidecar as success.
func TestFallbackBeacon_NoEndpointsReturnsError(t *testing.T) {
	c := NewFallbackBeaconClient(nil, nil, nil)
	sidecar, err := c.GetVerifiedBlobSidecar(context.Background(), L1BlockRef{Time: 12}, wantHashes(), nil)
	require.Error(t, err)
	require.Empty(t, sidecar.Blobs)
}

func TestFallbackBeacon_StopsOnContextCancellation(t *testing.T) {
	var primaryCalls, fallbackCalls int32
	c := &FallbackBeaconClient{
		clients: []*L1BeaconClient{
			NewL1BeaconClient(canceledHTTP{calls: &primaryCalls}),
			NewL1BeaconClient(canceledHTTP{calls: &fallbackCalls}),
		},
		endpoints: []string{"primary", "fallback"},
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	sidecars, err := c.GetVerifiedBlobSidecar(ctx, L1BlockRef{Time: 12}, wantHashes(), oneHash())

	require.ErrorIs(t, err, context.Canceled)
	require.Empty(t, sidecars.Blobs)
	require.Positive(t, atomic.LoadInt32(&primaryCalls))
	require.Zero(t, atomic.LoadInt32(&fallbackCalls))
}

// When every beacon fails to serve a valid blob, an error is returned and
// every endpoint's failure is recorded in metrics.
func TestFallbackBeacon_AllFailReturnsError(t *testing.T) {
	primary, primaryHits := newStubBeacon(t, beaconServesEmpty)
	fallback, fallbackHits := newStubBeacon(t, beaconServesCorruptBlob)

	m := PrometheusMetrics("morphnode_test_" + t.Name())
	c := NewFallbackBeaconClient([]string{primary, fallback}, nil, m)
	sidecars, err := fetch(t, c)
	require.Error(t, err)
	require.Empty(t, sidecars.Blobs)
	require.Positive(t, atomic.LoadInt32(primaryHits))
	require.Positive(t, atomic.LoadInt32(fallbackHits))
}
