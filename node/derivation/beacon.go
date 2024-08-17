package derivation

import (
	"context"
	"crypto/sha256"
	"encoding/json"
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
	if len(hashes) != len(resp.Data) {
		return nil, fmt.Errorf("expected %v sidecars but got %v", len(hashes), len(resp.Data))
	}

	return resp.Data, nil
}

// GetBlobSidecar fetches blob sidecars that were confirmed in the specified L1 block with the
// given indexed hashes. Order of the returned sidecars is not guaranteed, and blob data is not
// checked for validity.
func (cl *L1BeaconClient) GetBlobSidecar(ctx context.Context, ref L1BlockRef, hashes []IndexedBlobHash) (types.BlobTxSidecar, error) {
	blobSidecars, err := cl.GetBlobSidecars(ctx, ref, hashes)
	if err != nil {
		return types.BlobTxSidecar{}, fmt.Errorf("%w: failed to get blob sidecars for L1BlockRef %v", err, ref)
	}
	return sidecarFromSidecars(blobSidecars, hashes)
}

func indexFunc(s []*BlobSidecar, f func(blobSidecars *BlobSidecar) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func sidecarFromSidecars(blobSidecars []*BlobSidecar, hashes []IndexedBlobHash) (types.BlobTxSidecar, error) {
	var blobTxSidecar types.BlobTxSidecar
	for i, ih := range hashes {
		// The beacon node api makes no guarantees on order of the returned blob sidecars, so
		// search for the sidecar that matches the current indexed hash to ensure blobs are
		// returned in the same order.
		scIndex := indexFunc(
			blobSidecars,
			func(sc *BlobSidecar) bool { return uint64(sc.Index) == ih.Index })
		if scIndex == -1 {
			return types.BlobTxSidecar{}, fmt.Errorf("no blob in response matches desired index: %v", ih.Index)
		}
		sidecar := blobSidecars[scIndex]

		// make sure the blob's kzg commitment hashes to the expected value
		hash := KZGToVersionedHash(kzg4844.Commitment(sidecar.KZGCommitment))
		if hash != ih.Hash {
			return types.BlobTxSidecar{}, fmt.Errorf("expected hash %s for blob at index %d but got %s", ih.Hash, ih.Index, hash)
		}

		// confirm blob data is valid by verifying its proof against the commitment
		var blob Blob
		b, err := hexutil.Decode(sidecar.Blob)
		if err != nil {
			return types.BlobTxSidecar{}, fmt.Errorf("hexutil.Decode(sidecar.Blob) error:%v", err)
		}
		copy(blob[:], b)
		if err := VerifyBlobProof(&blob, kzg4844.Commitment(sidecar.KZGCommitment), kzg4844.Proof(sidecar.KZGProof)); err != nil {
			return types.BlobTxSidecar{}, fmt.Errorf("%w: blob at index %d failed verification", err, i)
		}
		blobTxSidecar.Blobs = append(blobTxSidecar.Blobs, *blob.KZGBlob())
		blobTxSidecar.Commitments = append(blobTxSidecar.Commitments, kzg4844.Commitment(sidecar.KZGCommitment))
		blobTxSidecar.Proofs = append(blobTxSidecar.Proofs, kzg4844.Proof(sidecar.KZGProof))
	}
	return blobTxSidecar, nil
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

func VerifyBlobProof(blob *Blob, commitment kzg4844.Commitment, proof kzg4844.Proof) error {
	return kzg4844.VerifyBlobProof(blob.KZGBlob(), commitment, proof)
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
