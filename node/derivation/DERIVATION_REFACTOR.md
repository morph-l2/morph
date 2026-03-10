# Derivation Refactor: Batch Verification & L1 Reorg Detection

## Background

The derivation module is the core component that syncs L2 state from L1 batch data. Previously it only ran on validator nodes and used a challenge mechanism when state mismatches were detected. This refactor makes two fundamental changes:

1. **L1 batch data is the source of truth** — when local L2 blocks don't match L1 batch data, roll back and re-derive from L1 instead of issuing a challenge.
2. **Support `latest` mode** for fetching L1 batches (instead of only `finalized`), with L1 reorg detection to handle the reduced confirmation window.

## Design Principles

- **L2 rollback is only triggered by batch data mismatch**, never by L1 reorg alone.
  - L1 reorg → clean up DB → re-derive from reorg point → batch comparison decides if L2 needs rollback.
  - Most L1 reorgs just re-include the same batch tx in a different block — L2 stays valid.
- **Derivation can run as a verification thread** — when blocks already exist locally (e.g. produced by sequencer), derivation compares them against L1 batch data instead of skipping.

## What Changed

### Removed

| Item | Reason |
|------|--------|
| `validator` field in `Derivation` struct | Challenge mechanism removed |
| `validator.Validator` parameter in `NewDerivationClient()` | No longer needed |
| `ChallengeState` / `ChallengeEnable` logic in `derivationBlock()` | Replaced by rollback + re-derive |
| `validator` import in `node/cmd/node/main.go` | No longer referenced |

### Added — L1 Reorg Detection

When `confirmations` is not `finalized` (i.e. using `latest` or `safe`), each derivation loop checks recent L1 blocks for hash changes before processing new batches.

**New DB layer** (`node/db/`):

- `DerivationL1Block` struct — stores `{Number, Hash, BatchIndex, L2EndBlock}` per L1 block
- `WriteDerivationL1Block` / `ReadDerivationL1Block` / `ReadDerivationL1BlockRange` / `DeleteDerivationL1BlocksFrom`
- DB key prefix: `derivL1Block` + uint64 big-endian height

**New config** (`node/derivation/config.go`):

- `ReorgCheckDepth uint64` — how many recent L1 blocks to verify each loop (default: 64)
- CLI flag: `--derivation.reorgCheckDepth` / env `MORPH_DERIVATION_REORG_CHECK_DEPTH`

**New methods** (`node/derivation/derivation.go`):

| Method | Purpose |
|--------|---------|
| `detectReorg(ctx)` | Iterates recent L1 block hashes from DB, compares against current L1 chain. Returns the height where a mismatch is found, or nil. |
| `handleL1Reorg(height)` | Cleans DB records from the reorg point and resets `latestDerivationL1Height`. Does NOT rollback L2 — the next derivation loop re-fetches batches and the normal comparison logic decides. |
| `recordL1Blocks(ctx, from, to)` | After each derivation round, records L1 block hashes for the processed range. |

**Flow**:

```
derivationBlock() loop start
│
├─ [if not finalized] detectReorg()
│   ├─ no reorg → continue
│   └─ reorg at height X → handleL1Reorg(X)
│       ├─ DeleteDerivationL1BlocksFrom(X)
│       ├─ WriteLatestDerivationL1Height(X-1)
│       └─ return (next loop re-processes from X)
│
├─ fetch CommitBatch logs from L1
├─ process each batch → derive() + verifyBatchRoots()
├─ recordL1Blocks(start, end)
└─ WriteLatestDerivationL1Height(end)
```

### Added — Batch Data Verification

When `derive()` encounters an L2 block that already exists locally, it now **compares** the block against the L1 batch data instead of blindly skipping it.

**New methods**:

| Method | Purpose |
|--------|---------|
| `verifyBlockContext(localHeader, blockData)` | Compares timestamp, gasLimit, baseFee between local L2 block header and batch block context. |
| `verifyBatchRoots(batchInfo, lastHeader)` | Compares stateRoot and withdrawalRoot between L1 batch and last derived L2 block. Extracted from the old inline logic. |
| `rollbackLocalChain(targetBlockNumber)` | **TODO stub** — will call geth `SetHead` API to rewind L2 chain. |

**`derive()` new flow for each block in batch**:

```
block.Number <= latestBlockNumber?
├─ YES (block exists)
│   ├─ verifyBlockContext() passes → skip, continue
│   └─ verifyBlockContext() fails
│       ├─ IncBlockMismatchCount()
│       ├─ rollbackLocalChain(block.Number - 1)
│       └─ fall through to NewSafeL2Block (re-execute)
│
└─ NO (new block)
    └─ NewSafeL2Block (execute normally)
```

**`derivationBlock()` batch-level verification**:

```
After derive(batchInfo) completes:
│
├─ verifyBatchRoots() passes → normal
└─ verifyBatchRoots() fails
    ├─ IncRollbackCount()
    ├─ rollbackLocalChain(firstBlockNumber - 1)
    ├─ re-derive(batchInfo)
    ├─ verifyBatchRoots() again
    │   ├─ passes → recovered
    │   └─ fails → CRITICAL error, stop (manual intervention needed)
```

### Added — Metrics

| Metric | Type | Description |
|--------|------|-------------|
| `morphnode_derivation_l1_reorg_detected_total` | Counter | L1 reorg detection count |
| `morphnode_derivation_l2_rollback_total` | Counter | L2 rollbacks triggered by batch mismatch |
| `morphnode_derivation_block_mismatch_total` | Counter | Block-level context mismatches |

## Modified Files

| File | Changes |
|------|---------|
| `node/derivation/derivation.go` | Core refactor: removed validator/challenge, added reorg detection, batch verification, rollback flow |
| `node/derivation/database.go` | Extended `Reader`/`Writer` interfaces for L1 block hash tracking |
| `node/derivation/config.go` | Added `ReorgCheckDepth` config field |
| `node/derivation/metrics.go` | Added 3 new counter metrics |
| `node/db/keys.go` | Added `derivationL1BlockPrefix` and `DerivationL1BlockKey()` |
| `node/db/store.go` | Added `DerivationL1Block` struct and 4 CRUD methods |
| `node/flags/flags.go` | Added `DerivationReorgCheckDepth` CLI flag |
| `node/cmd/node/main.go` | Removed `validator` dependency from `NewDerivationClient` call |

## TODO (follow-up work)

### `rollbackLocalChain()` — geth SetHead integration

Currently a stub that logs and returns nil. Needs:

1. Expose `SetL2Head(number uint64)` in `go-ethereum/eth/catalyst/l2_api.go`
2. Add `SetHead` method to `go-ethereum/ethclient/authclient`
3. Add `SetHead` method to `node/types/retryable_client.go`
4. Call `d.l2Client.SetHead(d.ctx, targetBlockNumber)` in `rollbackLocalChain()`

Note: geth already has `BlockChain.SetHead(head uint64) error` — we just need to expose it through the engine API chain.

### Concurrency safety

When running as a verification thread alongside a sequencer, concurrent access between block production and rollback needs locking. This will be handled separately.

## How to Test

1. **Existing behavior preserved**: Set `--derivation.confirmations` to finalized (default) — reorg detection is skipped, batch verification still runs.
2. **Latest mode**: Set `--derivation.confirmations` to `-2` (latest) — reorg detection activates, L1 block hashes are tracked.
3. **Reorg detection**: Simulate by modifying a saved L1 block hash in DB — next loop should detect and clean up.
4. **Batch verification**: When an existing L2 block matches L1 batch data, it logs "block verified" and skips. When mismatched, it logs the mismatch and attempts rollback (currently a no-op stub).
