package derivation

import (
	"context"
	"errors"
)

// SPEC-005 §3.3 path B (degraded batch-content verification).
//
// When path A (online beacon blob) is unavailable, this path rebuilds the
// versioned blob hash from local L2 blocks and compares it against the
// blob hash recorded in L1 commitBatch calldata. State / withdrawal root
// verification (verify.go::verifyBatchRoots) runs independently and is
// never gated on either path; see SPEC-005 §3.4.
//
// Trigger conditions (must all hold per SPEC-005 §3.3):
//   1. Path A returned an empty / unavailable result for this batch.
//   2. The batch's last L2 block is at or below safe_head — i.e. the batch
//      is in the historical tail, the only segment where blob retention
//      can legitimately have lapsed.
//   3. The local node still holds every L2 block in the batch range.
//
// Default-on/off behaviour and whether to retry path A on success are the
// SPEC-005 §8 #3 open question.

// errPathBUnavailable indicates the caller must fall back to the standard
// path-A failure handling (rollback / re-derive) — i.e. path B was either
// not eligible to run or failed to reproduce the blob hash.
var errPathBUnavailable = errors.New("path B unavailable")

// verifyBatchContentPathB attempts the degraded path B verification for the
// given batch. Returns nil on success.
//
// Eligibility check (returns errPathBUnavailable when not eligible) is
// kept inside this function so callers can blindly invoke it as a
// fallback after path A has failed — there is no separate "isEligible"
// query to keep two-stage races out of the main loop.
func (d *Derivation) verifyBatchContentPathB(ctx context.Context, batchInfo *BatchInfo) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !d.pathBEnabled() {
		return errPathBUnavailable
	}
	if !d.pathBEligible(batchInfo) {
		return errPathBUnavailable
	}

	// TODO(spec-005-path-b): rebuild versioned blob hash from local L2 blocks.
	//
	// Implementation sketch:
	//   1. For each L2 block in [batchInfo.firstBlockNumber, batchInfo.lastBlockNumber]:
	//        - fetch local block (already on disk; geth eth_getBlockByNumber).
	//        - encode tx list using node/types.MaxBlobBytesSize / RetrieveBlobBytes
	//          inverse: see node/types/blob.go for the path-A decode helpers.
	//   2. Compress with node/zstd, slice to blob-sized chunks (see SPEC-002 batching).
	//   3. For each chunk, compute kzg4844 commitment + versioned hash.
	//   4. Compare ordered versioned hashes against batchInfo.blobHashes.
	//
	// This is gated on confirming there's no double-implementation cost vs the
	// existing tx-submitter blob construction path (open question per
	// tech-design §8 / per-module §5 #3); production-grade code should reuse
	// existing helpers rather than reimplementing the encoder.

	d.metrics.IncPathBTriggered()
	d.logger.Info("path B verification triggered (skeleton — not yet implemented)",
		"batchIndex", batchInfo.batchIndex)
	return errPathBUnavailable
}

// pathBEnabled reports whether the operator has opted into the degraded path.
//
// TODO(spec-005-path-b): wire this to a flag once SPEC-005 §8 #3 is decided
// (default-on vs default-off). Until then, path B is permanently disabled.
func (d *Derivation) pathBEnabled() bool {
	return false
}

// pathBEligible reports whether path B can run for the given batch.
// Per SPEC-005 §3.3: batch must be historical (lastBlock <= safe_head) AND
// every L2 block in the range must exist locally.
func (d *Derivation) pathBEligible(batchInfo *BatchInfo) bool {
	safe := d.readSafeHead()
	if safe == nil {
		return false
	}
	if batchInfo.lastBlockNumber > safe.L2Number {
		// Live segment, not eligible — Path A failure here is a real anomaly.
		return false
	}
	// TODO(spec-005-path-b): walk [first, last] confirming local presence.
	// Skipped for skeleton — pathBEnabled() is false anyway.
	return true
}
