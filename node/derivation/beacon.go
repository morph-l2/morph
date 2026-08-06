package derivation

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/params"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

const (
	genesisMethod        = "eth/v1/beacon/genesis"
	specMethod           = "eth/v1/config/spec"
	sidecarsMethodPrefix = "eth/v1/beacon/blob_sidecars/"
)

type L1BeaconClient struct {
	cl HTTP

	initLock     sync.Mutex
	timeToSlotFn TimeToSlotFn
}

// NewL1BeaconClient returns a client for making requests to an L1 consensus layer node.
func NewL1BeaconClient(cl HTTP) *L1BeaconClient {
	return &L1BeaconClient{cl: cl}
}

func (cl *L1BeaconClient) apiReq(ctx context.Context, dest any, method string) error {
	headers := http.Header{}
	headers.Add("Accept", "application/json")
	resp, err := cl.cl.Get(ctx, method, headers)
	if err != nil {
		return fmt.Errorf("%w: http Get failed", err)
	}
	if resp.StatusCode != http.StatusOK {
		errMsg, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return fmt.Errorf("failed request with status %d: %s", resp.StatusCode, string(errMsg))
	}
	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		_ = resp.Body.Close()
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("%w: failed to close response body", err)
	}
	return nil
}

type TimeToSlotFn func(timestamp uint64) (uint64, error)

// GetTimeToSlotFn returns a function that converts a timestamp to a slot number.
func (cl *L1BeaconClient) GetTimeToSlotFn(ctx context.Context) (TimeToSlotFn, error) {
	cl.initLock.Lock()
	defer cl.initLock.Unlock()
	if cl.timeToSlotFn != nil {
		return cl.timeToSlotFn, nil
	}

	var genesisResp APIGenesisResponse
	if err := cl.apiReq(ctx, &genesisResp, genesisMethod); err != nil {
		return nil, err
	}

	var configResp APIConfigResponse
	if err := cl.apiReq(ctx, &configResp, specMethod); err != nil {
		return nil, err
	}

	genesisTime := uint64(genesisResp.Data.GenesisTime)
	secondsPerSlot := uint64(configResp.Data.SecondsPerSlot)
	if secondsPerSlot == 0 {
		return nil, fmt.Errorf("got bad value for seconds per slot: %v", configResp.Data.SecondsPerSlot)
	}
	cl.timeToSlotFn = func(timestamp uint64) (uint64, error) {
		if timestamp < genesisTime {
			return 0, fmt.Errorf("provided timestamp (%v) precedes genesis time (%v)", timestamp, genesisTime)
		}
		return (timestamp - genesisTime) / secondsPerSlot, nil
	}
	return cl.timeToSlotFn, nil
}

type L1BlockRef struct {
	Hash       common.Hash `json:"hash"`
	Number     uint64      `json:"number"`
	ParentHash common.Hash `json:"parentHash"`
	Time       uint64      `json:"timestamp"`
}

// GetBlobSidecars fetches blob sidecars that were confirmed in the specified L1 block with the
// given indexed hashes. Order of the returned sidecars is not guaranteed, and blob data is not
// checked for validity.
func (cl *L1BeaconClient) GetBlobSidecars(ctx context.Context, ref L1BlockRef, hashes []IndexedBlobHash) ([]*BlobSidecar, error) {
	if len(hashes) == 0 {
		return []*BlobSidecar{}, nil
	}
	slotFn, err := cl.GetTimeToSlotFn(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get time to slot function", err)
	}
	slot, err := slotFn(ref.Time)
	if err != nil {
		return nil, fmt.Errorf("%w: error in converting ref.Time to slot", err)
	}

	builder := strings.Builder{}
	builder.WriteString(sidecarsMethodPrefix)
	builder.WriteString(strconv.FormatUint(slot, 10))
	builder.WriteRune('?')
	v := url.Values{}

	for i := range hashes {
		v.Add("indices", strconv.FormatUint(hashes[i].Index, 10))
	}
	builder.WriteString(v.Encode())

	var resp APIGetBlobSidecarsResponse
	if err := cl.apiReq(ctx, &resp, builder.String()); err != nil {
		return nil, fmt.Errorf("%w: failed to fetch blob sidecars for slot %v block %v", err, slot, ref)
	}
	// Some Beacon nodes may ignore the indices parameter and return all sidecars for the slot.
	// We only need to ensure we have at least the number of sidecars we requested.
	// Callers are responsible for filtering the correct sidecars by index if needed.
	if len(resp.Data) < len(hashes) {
		return nil, fmt.Errorf("expected at least %v sidecars but got %v", len(hashes), len(resp.Data))
	}

	return resp.Data, nil
}

// IndexedBlobHash represents a blob hash that commits to a single blob confirmed in a block.  The
// index helps us avoid unnecessary blob to blob hash conversions to find the right content in a
// sidecar.
type IndexedBlobHash struct {
	Index uint64      // absolute index in the block, a.k.a. position in sidecar blobs array
	Hash  common.Hash // hash of the blob, used for consistency checks
}

func KZGToVersionedHash(commitment kzg4844.Commitment) (out common.Hash) {
	// EIP-4844 spec:
	//	def kzg_to_versioned_hash(commitment: KZGCommitment) -> VersionedHash:
	//		return VERSIONED_HASH_VERSION_KZG + sha256(commitment)[1:]
	h := sha256.New()
	h.Write(commitment[:])
	_ = h.Sum(out[:0])
	out[0] = params.BlobTxHashVersion
	return out
}

// verifyBlob authenticates a blob against the L1-signed versioned blob hash
// by recomputing the KZG commitment locally and checking
//
//	KZGToVersionedHash(BlobToCommitment(blob)) == expectedHash
//
// We deliberately do NOT verify a beacon-supplied kzg_proof. After
// EIP-7594 (PeerDAS / Osaka) the beacon /eth/v1/beacon/blob_sidecars
// endpoint's kzg_proof field is no longer guaranteed to be a legacy
// single-blob proof across forks/clients, and the new
// /eth/v1/beacon/blobs endpoint does not return proofs at all. The
// commitment round-trip gives us the same security property
// (blob bytes -> commitment -> versioned hash matches the L1-signed
// hash) without depending on those fields.
func verifyBlob(blob *Blob, expectedHash common.Hash) error {
	commitment, err := kzg4844.BlobToCommitment(blob.KZGBlob())
	if err != nil {
		return fmt.Errorf("cannot compute KZG commitment for blob: %w", err)
	}
	got := KZGToVersionedHash(commitment)
	if got != expectedHash {
		return fmt.Errorf("recomputed blob hash %s does not match expected %s", got.Hex(), expectedHash.Hex())
	}
	return nil
}

// dataAndHashesFromTxs extracts calldata and datahashes from the input transactions and returns them. It
// creates a placeholder blobOrCalldata element for each returned blob hash that must be populated
// by fillBlobPointers after blob bodies are retrieved.
func dataAndHashesFromTxs(txs types.Transactions, targetTx *types.Transaction) []IndexedBlobHash {
	var hashes []IndexedBlobHash
	blobIndex := 0 // index of each blob in the block's blob sidecar
	for _, tx := range txs {
		// skip any non-batcher transactions
		if tx.Hash() != targetTx.Hash() {
			blobIndex += len(tx.BlobHashes())
			continue
		}
		for _, h := range tx.BlobHashes() {
			idh := IndexedBlobHash{
				Index: uint64(blobIndex),
				Hash:  h,
			}
			hashes = append(hashes, idh)
			blobIndex++
		}
	}
	return hashes
}

// FallbackBeaconClient queries several beacon nodes in order and rotates to
// the next one when an endpoint cannot serve verified blobs (error, missing
// data, or content failing hash verification). Failing endpoints are recorded
// in the beacon_request_failure_total metric.
//
// Scope: fallback only covers per-endpoint data faults (corruption, pruned or
// unsynced data, client bugs). It is not a consistency mechanism and should
// not grow into one — safety against bad data comes from the KZG hash
// verification itself, and EL/CL fork mismatches near the chain head are
// eliminated by running derivation with confirmations=finalized, not by
// trying more beacons.
type FallbackBeaconClient struct {
	clients   []*L1BeaconClient
	endpoints []string // parallel to clients, used only for logs/metrics
	log       tmlog.Logger
	metrics   *Metrics
}

func NewFallbackBeaconClient(endpoints []string, log tmlog.Logger, metrics *Metrics) *FallbackBeaconClient {
	clients := make([]*L1BeaconClient, 0, len(endpoints))
	for _, endpoint := range endpoints {
		clients = append(clients, NewL1BeaconClient(NewBasicHTTPClient(endpoint, log)))
	}
	return &FallbackBeaconClient{
		clients:   clients,
		endpoints: endpoints,
		log:       log,
		metrics:   metrics,
	}
}

// GetVerifiedBlobSidecar fetches and content-verifies the blobs identified by
// wantHashes — the L1 tx's versioned blob hashes, in tx order — trying each
// configured beacon in turn. It returns the assembled BlobTxSidecar (blobs +
// commitments, in wantHashes order). A beacon that errors or serves an
// incomplete/invalid set is recorded as a failure and skipped; if every
// beacon fails, the last error is returned.
//
// indexHints is an optional optimization, NOT a correctness input: when
// non-empty, the first fetch attempt asks the beacon for only those sidecar
// indices (?indices=) instead of the whole slot. Verification is purely by
// hash — verifyBlob re-derives each commitment from the blob bytes and checks
// it against wantHashes — so a caller that cannot build indices (e.g. the L1
// block body is unavailable) may pass nil and every sidecar at the slot is
// fetched and matched by hash instead. Keeping index (a data-fetch detail)
// and hash (the security check) apart is deliberate: conflating them is what
// previously made a block-body fetch failure look "unverifiable" and led to
// dropping the fetch-all self-heal.
func (c *FallbackBeaconClient) GetVerifiedBlobSidecar(ctx context.Context, ref L1BlockRef, wantHashes []common.Hash, indexHints []IndexedBlobHash) (types.BlobTxSidecar, error) {
	if len(wantHashes) == 0 {
		return types.BlobTxSidecar{}, nil
	}
	// Config validation rejects an empty beacon list at startup; this guards
	// direct construction, where falling through would silently return an
	// empty sidecar as success.
	if len(c.clients) == 0 {
		return types.BlobTxSidecar{}, errors.New("no beacon endpoints configured")
	}
	var lastErr error
	for i, cl := range c.clients {
		sidecars, err := cl.GetBlobSidecarsEnhanced(ctx, ref, indexHints)
		// An empty list (slot pruned / not yet indexed) is an availability
		// failure of this endpoint; report it as such rather than as a
		// hash-match failure inside blobsFromSidecars.
		if err == nil && len(sidecars) == 0 {
			err = errors.New("beacon returned no sidecars for slot")
		}
		if err == nil {
			var verified types.BlobTxSidecar
			verified, err = blobsFromSidecars(sidecars, wantHashes)
			if err == nil {
				return verified, nil
			}
		}
		if ctxErr := ctx.Err(); ctxErr != nil {
			return types.BlobTxSidecar{}, ctxErr
		}
		if c.metrics != nil {
			c.metrics.IncBeaconRequestFailure(c.endpoints[i])
		}
		if c.log != nil {
			c.log.Error("beacon failed to serve valid blob sidecars, trying next endpoint",
				"endpoint", c.endpoints[i], "err", err)
		}
		lastErr = err
	}
	return types.BlobTxSidecar{}, lastErr
}

// blobsFromSidecars matches each wanted versioned hash to a sidecar via its
// commitment-derived versioned hash, authenticates the blob bytes with
// verifyBlob, and assembles the result in wantHashes (i.e. tx) order — batches
// are decoded by concatenating blob bodies, so order matters. Extra sidecars
// are ignored. Proofs are intentionally left empty: no consumer needs them
// and computing them costs an extra KZG op per blob.
func blobsFromSidecars(sidecars []*BlobSidecar, wantHashes []common.Hash) (types.BlobTxSidecar, error) {
	byHash := make(map[common.Hash]*BlobSidecar, len(sidecars))
	for _, sidecar := range sidecars {
		// JSON null entries decode to nil; skipping them surfaces as a
		// "not found" error below instead of a panic.
		if sidecar == nil {
			continue
		}
		var commitment kzg4844.Commitment
		copy(commitment[:], sidecar.KZGCommitment[:])
		byHash[KZGToVersionedHash(commitment)] = sidecar
	}
	out := types.BlobTxSidecar{
		Blobs:       make([]kzg4844.Blob, 0, len(wantHashes)),
		Commitments: make([]kzg4844.Commitment, 0, len(wantHashes)),
	}
	for _, expected := range wantHashes {
		sidecar, ok := byHash[expected]
		if !ok {
			return types.BlobTxSidecar{}, fmt.Errorf("blob (hash=%s) not found in beacon response", expected.Hex())
		}
		b, err := hexutil.Decode(sidecar.Blob)
		if err != nil {
			return types.BlobTxSidecar{}, fmt.Errorf("failed to decode blob (hash=%s): %w", expected.Hex(), err)
		}
		if len(b) != BlobSize {
			return types.BlobTxSidecar{}, fmt.Errorf("blob (hash=%s): unexpected length %d (want %d)", expected.Hex(), len(b), BlobSize)
		}
		var blob Blob
		copy(blob[:], b)
		if err := verifyBlob(&blob, expected); err != nil {
			return types.BlobTxSidecar{}, err
		}
		var commitment kzg4844.Commitment
		copy(commitment[:], sidecar.KZGCommitment[:])
		out.Blobs = append(out.Blobs, *blob.KZGBlob())
		out.Commitments = append(out.Commitments, commitment)
	}
	return out, nil
}

// Note: ForceGetAllBlobs is defined in derivation.go in the same package

// GetBlobSidecarsEnhanced is an enhanced version of GetBlobSidecars method, combining two approaches to fetch blob data
// If the first method fails or returns no blobs, it will try the second method
func (cl *L1BeaconClient) GetBlobSidecarsEnhanced(ctx context.Context, ref L1BlockRef, hashes []IndexedBlobHash) ([]*BlobSidecar, error) {
	// First try using the original GetBlobSidecars method
	blobSidecars, err := cl.GetBlobSidecars(ctx, ref, hashes)
	if err != nil || len(blobSidecars) == 0 {
		// If failed or no blobs retrieved, try the second method
		slotFn, err := cl.GetTimeToSlotFn(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get timeToSlotFn: %w", err)
		}

		slot, err := slotFn(ref.Time)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate slot: %w", err)
		}

		// Build request URL and use apiReq method directly
		method := fmt.Sprintf("%s%d", sidecarsMethodPrefix, slot)
		var blobResp APIGetBlobSidecarsResponse
		if err := cl.apiReq(ctx, &blobResp, method); err != nil {
			return nil, fmt.Errorf("failed to request blob sidecars: %w", err)
		}

		return blobResp.Data, nil
	}

	return blobSidecars, nil
}
