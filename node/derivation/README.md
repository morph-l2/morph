# node/derivation — Batch Verification

## Overview

This package previously contained a standalone `Derivation` service that ran its own
block-production loop: it watched L1 for `CommitBatch` events and re-derived L2 blocks
from blob/calldata, writing them via `NewSafeL2Block`.

The derivation node has been removed. All nodes now run Tendermint-based block production
and receive L2 blocks through P2P sync. The verification logic that lived inside the old
service has been extracted into `BatchVerifier` — a **stateless, schedule-free component**
that other parts of the node can call on demand.

`BatchProcessor` (`node/batchprocessor`) is the primary consumer: it sequentially walks
every committed batch, calls `BatchVerifier` for full verification, and updates
safe/finalized block tags on L2.

---

## Architecture

```
BatchProcessor (scheduler, node/batchprocessor)
│  polls L1 every N seconds
│  maintains two sequential cursors: lastSafeBatchIndex, lastFinalizedBatchIndex
│  walks batches from cursor+1 up to on-chain head
│
└─► BatchVerifier (stateless, no goroutines, node/derivation)
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
| `BaseHeight` | no | Snapshot base height; blocks <= this value are skipped |

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
configured and challenge is enabled. Steps 3-5 return errors only.

Verification is silently skipped during an upgrade transition window (ZK->MPT geth
switch) to avoid false positives while both geth versions coexist.

Blocks at or below `baseHeight` are skipped in all steps — this is required for nodes
that started from a snapshot rather than genesis.

#### `Close()`

Closes the L1 and L2 RPC connections. Call once when the verifier is no longer needed.

---

## BatchProcessor (node/batchprocessor)

`BatchProcessor` replaces the previous `BlockTagService` as the main scheduler for
tracking safe/finalized block tags. Key differences from the old approach:

- **Sequential processing**: walks every batch in order from cursor+1 to the on-chain
  head. No batch is skipped.
- **Single polling loop**: one goroutine, one ticker. No separate safe/finalized
  trackers or binary search.
- **Simple cursor model**: two integer cursors (`lastSafeBatchIndex`,
  `lastFinalizedBatchIndex`) track progress. `finalized <= safe` is always guaranteed.

### Flow per tick

```
processTick()
│
├─ getSafeL1Head() = latest L1 block - safeConfirmations
│  └─ getLastCommittedBatchAtBlock(safeL1Head)  → safe on-chain head
│     └─ advanceSafe: for idx in (cursor+1 .. safeHead):
│           processOneBatch(idx) → verify + get L2 block hash
│           update lastSafeBatchIndex, safeL2BlockHash
│
├─ getLastCommittedBatchAtBlock(finalized)  → finalized on-chain head
│  └─ advanceFinalized: for idx in (cursor+1 .. min(finalizedHead, safeCursor)):
│        processOneBatch(idx) → verify + get L2 block hash
│        update lastFinalizedBatchIndex, finalizedL2BlockHash
│
└─ notifyGeth() → SetBlockTags(safe, finalized) RPC to L2 geth
```

### processOneBatch(batchIndex)

1. `rollup.BatchDataStore(batchIndex)` — get the batch's `lastL2Block`
2. If `lastL2Block > l2Head` — node not synced yet, stop advancing
3. If `BatchVerifier` is available:
   - `fetchCommitBatchTxHash(batchIndex)` — find L1 tx via `FilterLogs`
   - `FetchBatchRoots` + `FetchBatchData` + `VerifyBatch`
   - On failure: log error (TODO: define failure behavior)
4. `HeaderByNumber(lastL2Block)` — get the L2 block hash
5. Return `(lastL2Block, blockHash)`

### Cursor initialization

At startup, both cursors are initialized to `LastFinalizedBatchIndex` from the L1
rollup contract, skipping already-finalized history. If the query fails, cursors
start from 0.

### Configuration (node/batchprocessor/config.go)

| Flag | Config field | Default | Notes |
|---|---|---|---|
| `--l1.node.addr` | `L1Addr` | — | L1 RPC URL |
| `--rollup.contract.address` | `RollupAddress` | — | Rollup contract on L1 |
| `--blocktag.safeConfirmations` | `SafeConfirmations` | `10` | L1 blocks before a batch is considered safe |
| *(inherited)* | `PollInterval` | `12s` | Polling interval |

---

## main.go integration

```go
// Build BatchVerifier (optional, non-critical)
bv, bvErr := derivation.NewBatchVerifier(rootCtx, bvCfg, nil, logger)
if bvErr != nil {
    logger.Error("BatchVerifier creation failed, verification disabled", "error", bvErr)
    bv = nil
}

// Create and start BatchProcessor
bp, err := batchprocessor.NewBatchProcessor(rootCtx, l2Client, bpCfg, bv, logger)
bp.Start()

// ...
<-rootCtx.Done()
bp.Stop()
```

`rootCtx` is created via `signal.NotifyContext` so that OS signals (SIGTERM, SIGINT)
propagate to startup retries (e.g. `FetchGethConfigWithRetry` inside `NewBatchVerifier`)
and the main blocking loop.

---

## Configuration (derivation)

Relevant flags (all under `node/flags`):

| Flag | Config field | Default | Notes |
|---|---|---|---|
| `--l1.node.addr` | `L1.Addr` | — | L1 RPC URL |
| `--l1.beacon.addr` | `BeaconRpc` | — | L1 Beacon URL; blob fetching disabled if empty |
| `--rollup.contract.address` | `RollupContractAddress` | — | Rollup contract on L1 |
| `--derivation.baseHeight` | `BaseHeight` | `0` | Snapshot base height; 0 = started from genesis |

---

## File map

| File | Package | Purpose |
|---|---|---|
| `derivation/batch_verifier.go` | derivation | `BatchVerifier`, `BatchRoots`, `BatchBlockContext`, `VerifyBatch` and all sub-checks |
| `derivation/batch_info.go` | derivation | `BatchInfo`, blob-decoded batch representation |
| `derivation/batch_decode.go` | derivation | `unpackCalldataWithABIs`, ABI version selection |
| `derivation/beacon.go` | derivation | `L1BeaconClient`, blob sidecar fetching |
| `derivation/blob_type.go` | derivation | Blob encoding/decoding helpers |
| `derivation/blobs.go` | derivation | KZG commitment verification helpers |
| `derivation/config.go` | derivation | `Config`, `DefaultConfig`, `SetCliContext` |
| `derivation/base_client.go` | derivation | `BasicHTTPClient` for Beacon API |
| `batchprocessor/processor.go` | batchprocessor | `BatchProcessor` — main scheduling loop, tag updates |
| `batchprocessor/config.go` | batchprocessor | `Config`, `DefaultConfig`, `SetCliContext` |
