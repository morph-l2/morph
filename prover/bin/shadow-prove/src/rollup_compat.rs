use std::collections::HashMap;

use alloy_primitives::{Bytes, TxHash, B256, U256};
use alloy_provider::{DynProvider, Provider};
use alloy_rpc_types::{BlockId, BlockNumberOrTag, Log};
use alloy_sol_types::SolCall;
use anyhow::{anyhow, Result};

use crate::abi::{
    pre_submitter::IRollup as PreSubmitterIRollup, IRollup, PreSubmitterRollup, Rollup,
    Rollup::RollupInstance,
};

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub enum SubmissionKind {
    CommitBatch,
    CommitState,
    CommitBatchWithProof,
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct BatchDataInput {
    pub version: u8,
    pub parent_batch_header: Bytes,
    pub last_block_number: u64,
    pub num_l1_messages: u16,
    pub prev_state_root: [u8; 32],
    pub post_state_root: [u8; 32],
    pub withdrawal_root: [u8; 32],
    pub kind: SubmissionKind,
    pub pre_submitter: bool,
}

#[derive(Clone, Copy, Debug, Eq, Ord, PartialEq, PartialOrd)]
pub struct LogIdentity {
    pub block_number: u64,
    pub transaction_index: u64,
    pub log_index: u64,
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub struct CanonicalCommit {
    pub identity: LogIdentity,
    pub transaction_hash: TxHash,
    pub block_hash: B256,
    pub batch_hash: B256,
}

fn log_identity(log: &Log) -> Option<LogIdentity> {
    Some(LogIdentity {
        block_number: log.block_number?,
        transaction_index: log.transaction_index?,
        log_index: log.log_index?,
    })
}

/// Select the last canonical Commit after the last Revert at one fixed L1
/// snapshot. Full log identity (rather than batch hash alone) deliberately
/// distinguishes a same-hash recommit.
pub fn select_canonical_commit(
    commits: &[Log],
    reverts: &[Log],
    canonical_hash: B256,
) -> Option<CanonicalCommit> {
    if canonical_hash == B256::ZERO {
        return None;
    }

    let last_revert = reverts.iter().filter(|log| !log.removed).filter_map(log_identity).max();

    commits
        .iter()
        .filter(|log| !log.removed)
        .filter_map(|log| {
            let identity = log_identity(log)?;
            if last_revert.is_some_and(|revert| identity <= revert) {
                return None;
            }
            let topics = log.topics();
            if topics.len() < 3 || topics[2] != canonical_hash {
                return None;
            }
            Some(CanonicalCommit {
                identity,
                transaction_hash: log.transaction_hash?,
                block_hash: log.block_hash?,
                batch_hash: topics[2],
            })
        })
        .max_by_key(|commit| commit.identity)
}

const LOG_QUERY_BLOCK_RANGE: u64 = 10_000;

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
struct IndexedSnapshot {
    number: u64,
    hash: B256,
}

/// A reusable Commit/Revert index anchored to one canonical finalized block.
///
/// The first refresh replays from the configured Rollup deployment block. A
/// later refresh first verifies the previous anchor and then scans only the
/// newly finalized suffix. If the anchor changed, the cache is discarded and
/// the current operation fails closed; a later poll may rebuild from genesis.
#[derive(Debug)]
pub struct CanonicalLogIndex {
    rollup_deployed_block: u64,
    indexed_snapshot: Option<IndexedSnapshot>,
    commits: HashMap<u64, Vec<Log>>,
    reverts: HashMap<u64, Vec<Log>>,
}

impl CanonicalLogIndex {
    pub fn new(rollup_deployed_block: u64) -> Result<Self> {
        if rollup_deployed_block == 0 {
            return Err(anyhow!("Rollup deployment block must be non-zero"));
        }
        Ok(Self {
            rollup_deployed_block,
            indexed_snapshot: None,
            commits: HashMap::new(),
            reverts: HashMap::new(),
        })
    }

    fn clear(&mut self) {
        self.indexed_snapshot = None;
        self.commits.clear();
        self.reverts.clear();
    }

    fn query_start(&self, snapshot_number: u64) -> Result<Option<u64>> {
        if snapshot_number < self.rollup_deployed_block {
            return Err(anyhow!(
                "Rollup deployment block {} is newer than snapshot {snapshot_number}",
                self.rollup_deployed_block
            ));
        }
        match self.indexed_snapshot {
            None => Ok(Some(self.rollup_deployed_block)),
            Some(indexed) if snapshot_number < indexed.number => Err(anyhow!(
                "finalized snapshot moved backwards: indexed={}, requested={snapshot_number}",
                indexed.number
            )),
            Some(indexed) if snapshot_number == indexed.number => Ok(None),
            Some(indexed) => Ok(Some(
                indexed
                    .number
                    .checked_add(1)
                    .ok_or_else(|| anyhow!("canonical log index block overflow"))?,
            )),
        }
    }

    fn verify_cached_anchor(&mut self, observed_hash: B256) -> Result<()> {
        let Some(indexed) = self.indexed_snapshot else {
            return Ok(());
        };
        if indexed.hash != observed_hash {
            self.clear();
            return Err(anyhow!(
                "canonical log index anchor changed at block {}; cache discarded",
                indexed.number
            ));
        }
        Ok(())
    }

    fn batch_index(log: &Log) -> Result<u64> {
        if log.removed {
            return Err(anyhow!("removed log returned for finalized canonical range"));
        }
        // Require every component used by canonical identity before admitting a
        // log to the cache. Resolution must never silently skip malformed RPC
        // results and fall back to an older Commit.
        log_identity(log).ok_or_else(|| anyhow!("canonical Rollup log has incomplete identity"))?;
        if log.transaction_hash.is_none() || log.block_hash.is_none() {
            return Err(anyhow!("canonical Rollup log has incomplete block/transaction identity"));
        }
        let topics = log.topics();
        if topics.len() < 3 {
            return Err(anyhow!("canonical Rollup log is missing indexed topics"));
        }
        let batch_index = U256::from_be_slice(topics[1].as_slice());
        if batch_index > U256::from(u64::MAX) {
            return Err(anyhow!("Rollup batch index exceeds u64"));
        }
        Ok(batch_index.to::<u64>())
    }

    fn append_logs(
        &mut self,
        snapshot_number: u64,
        snapshot_hash: B256,
        commits: Vec<Log>,
        reverts: Vec<Log>,
    ) -> Result<()> {
        let mut indexed_commits = Vec::with_capacity(commits.len());
        for log in commits {
            indexed_commits.push((Self::batch_index(&log)?, log));
        }
        let mut indexed_reverts = Vec::with_capacity(reverts.len());
        for log in reverts {
            indexed_reverts.push((Self::batch_index(&log)?, log));
        }
        for (batch_index, log) in indexed_commits {
            self.commits.entry(batch_index).or_default().push(log);
        }
        for (batch_index, log) in indexed_reverts {
            self.reverts.entry(batch_index).or_default().push(log);
        }
        self.indexed_snapshot =
            Some(IndexedSnapshot { number: snapshot_number, hash: snapshot_hash });
        Ok(())
    }

    fn resolve_cached(&self, batch_index: u64, canonical_hash: B256) -> Option<CanonicalCommit> {
        select_canonical_commit(
            self.commits.get(&batch_index).map(Vec::as_slice).unwrap_or_default(),
            self.reverts.get(&batch_index).map(Vec::as_slice).unwrap_or_default(),
            canonical_hash,
        )
    }

    pub async fn refresh(
        &mut self,
        provider: &DynProvider,
        rollup: &RollupInstance<DynProvider>,
        snapshot_number: u64,
        snapshot_hash: B256,
    ) -> Result<()> {
        let snapshot_tag = BlockNumberOrTag::Number(snapshot_number);
        let before = provider
            .get_block_by_number(snapshot_tag)
            .await?
            .ok_or_else(|| anyhow!("snapshot block {snapshot_number} not found"))?;
        if before.header.hash != snapshot_hash {
            return Err(anyhow!("finalized snapshot hash changed before canonical log refresh"));
        }

        if let Some(indexed) = self.indexed_snapshot {
            let anchor = provider
                .get_block_by_number(BlockNumberOrTag::Number(indexed.number))
                .await?
                .ok_or_else(|| anyhow!("canonical log index anchor block disappeared"))?;
            self.verify_cached_anchor(anchor.header.hash)?;
        }

        let query_start = match self.query_start(snapshot_number) {
            Ok(query_start) => query_start,
            Err(error) => {
                self.clear();
                return Err(error);
            }
        };
        let Some(mut from) = query_start else {
            // Equal-height reuse still verifies the caller supplied the exact
            // cached block hash, rather than merely the block number.
            self.verify_cached_anchor(snapshot_hash)?;
            return Ok(());
        };

        let mut commits = Vec::new();
        let mut reverts = Vec::new();
        while from <= snapshot_number {
            let to = from.saturating_add(LOG_QUERY_BLOCK_RANGE - 1).min(snapshot_number);
            let commit_filter = rollup
                .CommitBatch_filter()
                .filter
                .from_block(from)
                .to_block(to)
                .address(*rollup.address());
            commits.extend(provider.get_logs(&commit_filter).await?);

            let revert_filter = rollup
                .RevertBatch_filter()
                .filter
                .from_block(from)
                .to_block(to)
                .address(*rollup.address());
            reverts.extend(provider.get_logs(&revert_filter).await?);

            if to == snapshot_number {
                break;
            }
            from =
                to.checked_add(1).ok_or_else(|| anyhow!("canonical log query block overflow"))?;
        }

        let after = provider
            .get_block_by_number(snapshot_tag)
            .await?
            .ok_or_else(|| anyhow!("snapshot block {snapshot_number} disappeared"))?;
        if after.header.hash != snapshot_hash {
            return Err(anyhow!("snapshot block hash changed during canonical log refresh"));
        }

        self.append_logs(snapshot_number, snapshot_hash, commits, reverts)
    }

    /// Resolve from the already-refreshed fixed snapshot. Multiple batch
    /// lookups (shadow prev/current/next) reuse the same log replay.
    pub async fn resolve(
        &self,
        rollup: &RollupInstance<DynProvider>,
        batch_index: u64,
        snapshot_number: u64,
        snapshot_hash: B256,
    ) -> Result<Option<CanonicalCommit>> {
        if self.indexed_snapshot !=
            Some(IndexedSnapshot { number: snapshot_number, hash: snapshot_hash })
        {
            return Err(anyhow!("canonical log index is not anchored to the requested snapshot"));
        }
        let canonical_hash = rollup
            .committedBatches(U256::from(batch_index))
            .block(BlockId::Number(BlockNumberOrTag::Number(snapshot_number)))
            .call()
            .await?;
        if canonical_hash == B256::ZERO {
            return Ok(None);
        }
        Ok(self.resolve_cached(batch_index, canonical_hash))
    }
}

fn current_input(input: IRollup::BatchDataInput, kind: SubmissionKind) -> BatchDataInput {
    BatchDataInput {
        version: input.version,
        parent_batch_header: input.parentBatchHeader,
        last_block_number: input.lastBlockNumber,
        num_l1_messages: input.numL1Messages,
        prev_state_root: *input.prevStateRoot,
        post_state_root: *input.postStateRoot,
        withdrawal_root: *input.withdrawalRoot,
        kind,
        pre_submitter: false,
    }
}

fn historical_input(
    input: PreSubmitterIRollup::BatchDataInput,
    kind: SubmissionKind,
) -> BatchDataInput {
    BatchDataInput {
        version: input.version,
        parent_batch_header: input.parentBatchHeader,
        last_block_number: input.lastBlockNumber,
        num_l1_messages: input.numL1Messages,
        prev_state_root: *input.prevStateRoot,
        post_state_root: *input.postStateRoot,
        withdrawal_root: *input.withdrawalRoot,
        kind,
        pre_submitter: true,
    }
}

/// Decode only the first `BatchDataInput` argument from every submission ABI
/// supported across the Submitter cutover. Unknown and short selectors fail
/// closed; historical signature data is intentionally not exposed.
pub fn decode_batch_data_input(data: &[u8]) -> Result<BatchDataInput> {
    if data.len() < 4 {
        return Err(anyhow!("rollup calldata is shorter than a selector"));
    }

    if let Ok(call) = Rollup::commitBatchCall::abi_decode(data) {
        return Ok(current_input(call.batchDataInput, SubmissionKind::CommitBatch));
    }
    if let Ok(call) = Rollup::commitStateCall::abi_decode(data) {
        return Ok(current_input(call.batchDataInput, SubmissionKind::CommitState));
    }
    if let Ok(call) = Rollup::commitBatchWithProofCall::abi_decode(data) {
        return Ok(current_input(call.batchDataInput, SubmissionKind::CommitBatchWithProof));
    }

    if let Ok(call) = PreSubmitterRollup::commitBatchCall::abi_decode(data) {
        return Ok(historical_input(call.batchDataInput, SubmissionKind::CommitBatch));
    }
    if let Ok(call) = PreSubmitterRollup::commitStateCall::abi_decode(data) {
        return Ok(historical_input(call.batchDataInput, SubmissionKind::CommitState));
    }
    if let Ok(call) = PreSubmitterRollup::commitBatchWithProofCall::abi_decode(data) {
        return Ok(historical_input(call.batchDataInput, SubmissionKind::CommitBatchWithProof));
    }

    Err(anyhow!("unsupported rollup submission selector"))
}

#[cfg(test)]
mod tests {
    use super::*;
    use alloy_primitives::{Address, Log as PrimitiveLog};

    fn rpc_log(
        batch_index: u64,
        batch_hash: B256,
        block_number: u64,
        transaction_index: u64,
        log_index: u64,
        tx_marker: u8,
    ) -> Log {
        Log {
            inner: PrimitiveLog::new_unchecked(
                Address::ZERO,
                vec![
                    B256::ZERO,
                    B256::from(U256::from(batch_index).to_be_bytes::<32>()),
                    batch_hash,
                ],
                Bytes::new(),
            ),
            block_number: Some(block_number),
            transaction_index: Some(transaction_index),
            log_index: Some(log_index),
            transaction_hash: Some(B256::repeat_byte(tx_marker)),
            block_hash: Some(B256::repeat_byte(block_number as u8)),
            ..Default::default()
        }
    }

    #[test]
    fn rejects_short_and_unknown_calldata() {
        assert!(decode_batch_data_input(&[]).is_err());
        assert!(decode_batch_data_input(&[0xde, 0xad, 0xbe, 0xef]).is_err());
    }

    #[test]
    fn decodes_shared_pre_and_post_submitter_fixtures() {
        let corpus: serde_json::Value = serde_json::from_str(include_str!(
            "../../../../node/derivation/testdata/rollup_calldata_fixtures.json"
        ))
        .unwrap();
        let mut decoded_count = 0;
        for fixture in corpus["fixtures"].as_array().unwrap() {
            let epoch = fixture["epoch"].as_str().unwrap();
            if epoch != "pre-submitter" && epoch != "current" {
                continue;
            }
            let data = alloy_primitives::hex::decode(
                fixture["data"].as_str().unwrap().trim_start_matches("0x"),
            )
            .unwrap();
            let decoded = decode_batch_data_input(&data)
                .unwrap_or_else(|err| panic!("{}: {err:#}", fixture["name"]));
            assert_eq!(decoded.version as u64, fixture["expected"]["version"].as_u64().unwrap());
            assert_eq!(
                decoded.last_block_number,
                fixture["expected"]["lastBlockNumber"].as_u64().unwrap()
            );
            decoded_count += 1;
        }
        assert_eq!(decoded_count, 6);
    }

    #[test]
    fn selector_sets_are_distinct() {
        assert_eq!(Rollup::commitBatchCall::SELECTOR, [0x41, 0xf7, 0x56, 0xda]);
        assert_eq!(Rollup::commitStateCall::SELECTOR, [0x67, 0xca, 0xa3, 0x7a]);
        assert_eq!(Rollup::commitBatchWithProofCall::SELECTOR, [0x15, 0x44, 0xba, 0x3a]);

        assert_eq!(PreSubmitterRollup::commitBatchCall::SELECTOR, [0x42, 0x88, 0x68, 0xb5]);
        assert_eq!(PreSubmitterRollup::commitStateCall::SELECTOR, [0x1e, 0x88, 0x25, 0xbe]);
        assert_eq!(
            PreSubmitterRollup::commitBatchWithProofCall::SELECTOR,
            [0x4e, 0x8f, 0x1d, 0x67]
        );
    }

    #[test]
    fn recommit_after_revert_wins_even_when_hash_is_identical() {
        let hash = B256::repeat_byte(0x44);
        let first = rpc_log(7, hash, 10, 0, 0, 1);
        let revert = rpc_log(7, hash, 11, 0, 0, 2);
        let second = rpc_log(7, hash, 12, 0, 0, 3);

        let resolved = select_canonical_commit(&[first, second], &[revert], hash).unwrap();
        assert_eq!(resolved.transaction_hash, B256::repeat_byte(3));
        assert_eq!(resolved.identity.block_number, 12);
    }

    #[test]
    fn zero_hash_and_same_block_revert_order_fail_closed() {
        let hash = B256::repeat_byte(0x44);
        assert!(
            select_canonical_commit(&[rpc_log(7, hash, 10, 0, 0, 1)], &[], B256::ZERO).is_none()
        );

        let before_revert = rpc_log(7, hash, 10, 0, 1, 1);
        let revert = rpc_log(7, hash, 10, 0, 2, 2);
        assert!(select_canonical_commit(&[before_revert], std::slice::from_ref(&revert), hash)
            .is_none());

        let after_revert = rpc_log(7, hash, 10, 0, 3, 3);
        let resolved = select_canonical_commit(&[after_revert], &[revert], hash).unwrap();
        assert_eq!(resolved.transaction_hash, B256::repeat_byte(3));
    }

    #[test]
    fn canonical_log_index_reuses_one_snapshot_for_adjacent_batches_and_then_increments() {
        let snapshot_hash = B256::repeat_byte(0xaa);
        let mut index = CanonicalLogIndex::new(5).unwrap();
        assert_eq!(index.query_start(10).unwrap(), Some(5));

        let hash6 = B256::repeat_byte(6);
        let hash7 = B256::repeat_byte(7);
        let hash8 = B256::repeat_byte(8);
        index
            .append_logs(
                10,
                snapshot_hash,
                vec![
                    rpc_log(6, hash6, 6, 0, 0, 6),
                    rpc_log(7, hash7, 7, 0, 0, 7),
                    rpc_log(8, hash8, 8, 0, 0, 8),
                ],
                vec![],
            )
            .unwrap();

        // prev/current/next are all served by the same indexed replay.
        assert_eq!(index.resolve_cached(6, hash6).unwrap().identity.block_number, 6);
        assert_eq!(index.resolve_cached(7, hash7).unwrap().identity.block_number, 7);
        assert_eq!(index.resolve_cached(8, hash8).unwrap().identity.block_number, 8);
        assert_eq!(index.query_start(10).unwrap(), None);
        assert_eq!(index.query_start(12).unwrap(), Some(11));

        index
            .append_logs(
                12,
                B256::repeat_byte(0xbb),
                vec![rpc_log(9, B256::repeat_byte(9), 11, 0, 0, 9)],
                vec![],
            )
            .unwrap();
        assert!(index.resolve_cached(7, hash7).is_some());
        assert!(index.resolve_cached(9, B256::repeat_byte(9)).is_some());
    }

    #[test]
    fn canonical_log_index_discards_cache_and_fails_closed_on_anchor_reorg() {
        let old_hash = B256::repeat_byte(0xaa);
        let mut index = CanonicalLogIndex::new(5).unwrap();
        index
            .append_logs(10, old_hash, vec![rpc_log(7, B256::repeat_byte(7), 7, 0, 0, 7)], vec![])
            .unwrap();

        let error = index.verify_cached_anchor(B256::repeat_byte(0xbb)).unwrap_err();
        assert!(error.to_string().contains("cache discarded"));
        assert!(index.resolve_cached(7, B256::repeat_byte(7)).is_none());
        assert_eq!(index.query_start(12).unwrap(), Some(5));
    }
}
