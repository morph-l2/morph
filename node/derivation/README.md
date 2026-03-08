# node/derivation — Batch Verification

## Overview

This package previously contained a standalone `Derivation` service that ran its own
block-production loop: it watched L1 for `CommitBatch` events and re-derived L2 blocks
from blob/calldata, writing them via `NewSafeL2Block`.

The derivation node has been removed. All nodes now run Tendermint-based block production
and receive L2 blocks through P2P sync. The verification logic that lived inside the old
service has been extracted into `BatchVerifier` — a **stateless, schedule-free component**
that other parts of the node can call on demand.

`BlockTagService` (`node/blocktag`) is the primary consumer: it calls `BatchVerifier` as
part of its polling loop to perform full batch verification before promoting a batch to
safe or finalized.

---

## Architecture

```
BlockTagService (scheduler)
│  polls L1 every N seconds
│  determines safe / finalized batch index
│  calls validateBatch(tagType, batchIndex, lastL2Block)
│
└─► BatchVerifier (stateless, no goroutines)
        FetchBatchRoots(txHash)   — parse L1 calldata → BatchRoots
        FetchBatchData(txHash)    — fetch blobs → BatchInfo (optional)
        VerifyBatch(roots, data)  — 5-step verification against local L2
```

`BatchVerifier` holds long-lived RPC connections (L1, L2, optional L1 Beacon) but owns
no goroutines and performs no scheduling. The caller decides when to invoke it.

---

## BatchVerifier

### Construction

```go
bv, err := derivation.NewBatchVerifier(ctx, cfg, validator, logger)
defer bv.Close()
```

`cfg` is `*derivation.Config`. Only a subset of fields is required when constructing
`BatchVerifier` directly (e.g. from `main.go` without a full derivation setup):

| Field | Required | Notes |
|---|---|---|
| `L1.Addr` | yes | L1 RPC endpoint |
| `L2.EthAddr` | yes | L2 eth RPC endpoint |
| `RollupContractAddress` | yes | Rollup contract on L1 |
| `BeaconRpc` | no | Enables blob fetching (Step 5); skipped if empty |
| `BaseHeight` | no | Snapshot base height; blocks ≤ this value are skipped |

`NewBatchVerifier` fetches geth upgrade config at startup (retries until geth is ready).

### Key types

**`BatchRoots`** — roots and block metadata parsed from L1 calldata only (no blobs):

```go
type BatchRoots struct {
    BatchIndex     uint64
    FirstBlockNum  uint64          // 0 for blob-based (v0) batches
    LastBlockNum   uint64
    PrevStateRoot  common.Hash     // zero for v0 batches
    PostStateRoot  common.Hash
    WithdrawalRoot common.Hash
    NumL1Messages  uint16
    BlockContexts  []BatchBlockContext // nil for v0 batches
}
```

**`BatchBlockContext`** — per-block fields decoded from calldata (v1+ batches):

```go
type BatchBlockContext struct {
    Number    uint64
    Timestamp uint64
    GasLimit  uint64
    BaseFee   *big.Int
    NumTxs    uint16   // total txs (L2 user + L1 message)
    NumL1Msgs uint16   // L1 message txs (type 0x7E)
}
```

**`BatchInfo`** — full blob-decoded batch data, returned by `FetchBatchData`. Contains
per-block user transaction lists used for Step 5 content verification.

### Methods

#### `FetchBatchRoots(ctx, txHash, batchIndex) (*BatchRoots, error)`

Fetches the L1 transaction and parses its calldata to extract state roots and block
metadata. No blob fetching is performed.

Calldata is unpacked across three ABI versions (legacy, beforeMoveBlockCtx, current).
For v0/legacy batches where `LastBlockNumber` is absent from calldata, it falls back to
querying `BatchDataStore` on-chain.

#### `FetchBatchData(ctx, txHash) (*BatchInfo, error)`

Fetches blob sidecars from the L1 Beacon API and decodes the full batch (user
transactions per block). Returns an error if `BeaconRpc` is not configured.

This is optional for `VerifyBatch` — pass `nil` to skip Step 5 (transaction content
check).

#### `VerifyBatch(ctx, l2Client, roots, batchData) error`

Runs up to five verification steps against the local L2 node:

| Step | Check | Condition |
|---|---|---|
| 1 | `PostStateRoot`: L2 block state root == L1-committed root | always |
| 2 | `WithdrawalRoot`: `L2ToL1MessagePasser.MessageRoot` == L1-committed root | always |
| 3 | `PrevStateRoot`: block before batch has correct state root | v1+, FirstBlockNum > baseHeight |
| 4 | Block contexts: Number, Timestamp, GasLimit, BaseFee, NumTxs, NumL1Msgs | v1+ (BlockContexts in calldata) |
| 5 | L2 user tx content: blob-decoded txs match actual L2 block txs byte-for-byte | batchData != nil |

Steps 1 and 2 trigger `validator.ChallengeState` on mismatch when a Validator is
configured and challenge is enabled. Steps 3–5 return errors only.

Verification is silently skipped during an upgrade transition window (ZK→MPT geth
switch) to avoid false positives while both geth versions coexist.

Blocks at or below `baseHeight` are skipped in all steps — this is required for nodes
that started from a snapshot rather than genesis.

#### `Close()`

Closes the L1 and L2 RPC connections. Call once when the verifier is no longer needed.

---

## Integration with BlockTagService

`BlockTagService` accepts an optional `*BatchVerifier` at construction:

```go
blockTagSvc, err := blocktag.NewBlockTagService(ctx, l2Client, blockTagConfig, bv, logger)
```

- **`bv != nil`**: full batch verification via `BatchVerifier` (Steps 1–5 above).
- **`bv == nil`**: lightweight fallback — only `CommittedStateRoots` on the Rollup
  contract is compared against the local L2 state root.

### CommitBatch log search

To call `FetchBatchRoots`, `BlockTagService` needs the L1 transaction hash of the
`CommitBatch` event for the target batch. It finds this by calling `FilterLogs` on the
L1 node, filtering by `RollupEventTopicHash` and the indexed `batchIndex`.

Safe and finalized tags maintain **independent search trackers** (`safeSearchTracker`,
`finalizedSearchTracker`) so that safe batch queries (which target more recent L1 blocks)
cannot advance the search start beyond the point where finalized batch events are found.

#### `l1SearchTracker`

A small helper that encapsulates the `FromBlock` management for `FilterLogs` calls.

Two modes, selected once at construction:

| Mode | Condition | Behaviour |
|---|---|---|
| **Fixed** | `L1StartBlock > 0` in config | `FromBlock` always returns the configured value; `Advance` is a no-op |
| **Auto** | `L1StartBlock == 0` (default) | `FromBlock` starts at 0, refined at startup from the last finalized batch, then progressively advanced after each successful log query |

At startup, `initSearchFromBlock` queries the `CommitBatch` event for
`LastFinalizedBatchIndex` and advances both trackers to that L1 block number. This
avoids a full genesis-to-tip scan on every restart. (If the query fails, trackers stay
at their initial value — auto mode starts from 0, a one-time cost.)

---

## Configuration

Relevant flags (all under `node/flags`):

| Flag | Config field | Default | Notes |
|---|---|---|---|
| `--l1.node.addr` | `L1.Addr` | — | L1 RPC URL |
| `--l1.beacon.addr` | `BeaconRpc` | — | L1 Beacon URL; blob fetching disabled if empty |
| `--rollup.contract.address` | `RollupContractAddress` | — | Rollup contract on L1 |
| `--derivation.baseHeight` | `BaseHeight` | `0` | Snapshot base height; 0 = started from genesis |
| `--blocktag.safeConfirmations` | `SafeConfirmations` | `10` | L1 blocks before a batch is safe |
| *(no flag)* | `L1StartBlock` | `0` | Fixed L1 scan start; set programmatically only. `0` = auto mode (tracker refined from last finalized batch at startup) |

---

## File map

| File | Purpose |
|---|---|
| `batch_verifier.go` | `BatchVerifier`, `BatchRoots`, `BatchBlockContext`, `VerifyBatch` and all sub-checks |
| `batch_info.go` | `BatchInfo`, blob-decoded batch representation |
| `batch_decode.go` | `unpackCalldataWithABIs`, ABI version selection |
| `beacon.go` | `L1BeaconClient`, blob sidecar fetching |
| `blob_type.go` | Blob encoding/decoding helpers |
| `blobs.go` | KZG commitment verification helpers |
| `config.go` | `Config`, `DefaultConfig`, `SetCliContext` |
| `base_client.go` | `BasicHTTPClient` for Beacon API |
| `database.go` | Retained for potential future use (not active) |
| `metrics.go` | Prometheus metrics stubs (retained, not wired to `BatchVerifier`) |

