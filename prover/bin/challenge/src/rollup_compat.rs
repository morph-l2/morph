use ethers::{
    abi::AbiDecode,
    providers::Middleware,
    types::{transaction::eip2718::TypedTransaction, Address, BlockId, BlockNumber, Bytes, Filter, Log, TransactionRequest, TxHash, H256, U256, U64},
    utils::{id, keccak256},
};
use eyre::{eyre, Result};
use std::collections::HashMap;

use crate::abi::{
    legacy_l1_staking_abi::LegacyL1Staking,
    pre_submitter_rollup_abi::{
        CommitBatchCall as PreSubmitterCommitBatchCall, CommitBatchWithProofCall as PreSubmitterCommitBatchWithProofCall,
        CommitStateCall as PreSubmitterCommitStateCall,
    },
    rollup_abi::{CommitBatchCall, CommitBatchWithProofCall, CommitStateCall, Rollup},
    submitter_abi::Submitter,
};

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub enum AuthorityKind {
    LegacyL1Staking,
    Submitter,
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub struct RollupAuthority {
    pub address: Address,
    pub kind: AuthorityKind,
}

fn decode_address_word(data: &[u8]) -> Result<Address> {
    if data.len() != 32 || data[..12].iter().any(|byte| *byte != 0) {
        return Err(eyre!("authority getter returned malformed address data"));
    }
    let address = Address::from_slice(&data[12..]);
    if address.is_zero() {
        return Err(eyre!("authority getter returned zero address"));
    }
    Ok(address)
}

pub async fn resolve_rollup_authority<M: Middleware + 'static>(rollup: &Rollup<M>, snapshot: BlockId) -> Result<RollupAuthority> {
    match rollup.submitter_contract().block(snapshot).call().await {
        Ok(address) if !address.is_zero() => Ok(RollupAuthority {
            address,
            kind: AuthorityKind::Submitter,
        }),
        Ok(_) => Err(eyre!("Rollup.submitterContract returned zero address")),
        Err(current_error) => {
            // Before initialize4 the proxy still runs the frozen Rollup ABI and
            // rejects submitterContract(). Read the old slot through its
            // historical getter at the exact same numbered snapshot.
            let request: TypedTransaction = TransactionRequest::new()
                .to(rollup.address())
                .data(Bytes::from(id("l1StakingContract()")[..4].to_vec()))
                .into();
            let output = rollup.client().call(&request, Some(snapshot)).await.map_err(|legacy_error| {
                eyre!("neither current nor pre-submit authority getter succeeded: current={current_error}; legacy={legacy_error}")
            })?;
            Ok(RollupAuthority {
                address: decode_address_word(output.as_ref())?,
                kind: AuthorityKind::LegacyL1Staking,
            })
        }
    }
}

pub async fn authority_is_active<M: Middleware + 'static>(rollup: &Rollup<M>, signer: Address, snapshot: BlockId) -> Result<bool> {
    let authority = resolve_rollup_authority(rollup, snapshot).await?;
    match authority.kind {
        AuthorityKind::Submitter => Ok(Submitter::new(authority.address, rollup.client())
            .is_active(signer)
            .block(snapshot)
            .call()
            .await?),
        AuthorityKind::LegacyL1Staking => Ok(LegacyL1Staking::new(authority.address, rollup.client())
            .is_active_staker(signer)
            .block(snapshot)
            .call()
            .await?),
    }
}

pub async fn challenge_deposit_at<M: Middleware + 'static>(rollup: &Rollup<M>, snapshot: BlockId) -> Result<U256> {
    let authority = resolve_rollup_authority(rollup, snapshot).await?;
    match authority.kind {
        AuthorityKind::Submitter => Ok(Submitter::new(authority.address, rollup.client())
            .challenge_deposit()
            .block(snapshot)
            .call()
            .await?),
        AuthorityKind::LegacyL1Staking => Ok(LegacyL1Staking::new(authority.address, rollup.client())
            .challenge_deposit()
            .block(snapshot)
            .call()
            .await?),
    }
}

pub fn decode_batch_data_store_block_number(data: &[u8]) -> Result<U256> {
    if data.len() != 128 {
        return Err(eyre!("batchDataStore returned {}, expected exactly 128 bytes", data.len()));
    }
    Ok(U256::from_big_endian(&data[64..96]))
}

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

macro_rules! decoded {
    ($input:expr, $kind:expr, $pre:expr) => {{
        let input = $input;
        BatchDataInput {
            version: input.version,
            parent_batch_header: input.parent_batch_header,
            last_block_number: input.last_block_number,
            num_l1_messages: input.num_l1_messages,
            prev_state_root: input.prev_state_root,
            post_state_root: input.post_state_root,
            withdrawal_root: input.withdrawal_root,
            kind: $kind,
            pre_submitter: $pre,
        }
    }};
}

pub fn decode_batch_data_input(data: &[u8]) -> Result<BatchDataInput> {
    if data.len() < 4 {
        return Err(eyre!("rollup calldata is shorter than a selector"));
    }

    if let Ok(call) = CommitBatchCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatch, false));
    }
    if let Ok(call) = CommitStateCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitState, false));
    }
    if let Ok(call) = CommitBatchWithProofCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatchWithProof, false));
    }

    if let Ok(call) = PreSubmitterCommitBatchCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatch, true));
    }
    if let Ok(call) = PreSubmitterCommitStateCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitState, true));
    }
    if let Ok(call) = PreSubmitterCommitBatchWithProofCall::decode(data) {
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatchWithProof, true));
    }

    Err(eyre!("unsupported rollup submission selector"))
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
    pub block_hash: H256,
    pub batch_hash: H256,
}

fn identity(log: &Log) -> Option<LogIdentity> {
    Some(LogIdentity {
        block_number: log.block_number?.as_u64(),
        transaction_index: log.transaction_index?.as_u64(),
        log_index: log.log_index?.as_u64(),
    })
}

pub fn select_canonical_commit(commits: &[Log], reverts: &[Log], canonical_hash: H256) -> Option<CanonicalCommit> {
    if canonical_hash.is_zero() {
        return None;
    }
    let last_revert = reverts.iter().filter(|log| !log.removed.unwrap_or_default()).filter_map(identity).max();

    commits
        .iter()
        .filter(|log| !log.removed.unwrap_or_default())
        .filter_map(|log| {
            let identity = identity(log)?;
            if last_revert.is_some_and(|revert| identity <= revert) {
                return None;
            }
            if log.topics.len() < 3 || log.topics[2] != canonical_hash {
                return None;
            }
            Some(CanonicalCommit {
                identity,
                transaction_hash: log.transaction_hash?,
                block_hash: log.block_hash?,
                batch_hash: log.topics[2],
            })
        })
        .max_by_key(|commit| commit.identity)
}

const QUERY_RANGE: u64 = 10_000;

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
struct IndexedSnapshot {
    number: u64,
    hash: H256,
}

/// Reusable finalized Commit/Revert index. The first poll replays from the
/// configured Rollup deployment block; subsequent polls validate the previous
/// block-hash anchor and query only the newly finalized suffix.
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
            return Err(eyre!("Rollup deployment block must be non-zero"));
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
            return Err(eyre!(
                "Rollup deployment block {} is newer than snapshot {snapshot_number}",
                self.rollup_deployed_block
            ));
        }
        match self.indexed_snapshot {
            None => Ok(Some(self.rollup_deployed_block)),
            Some(indexed) if snapshot_number < indexed.number => Err(eyre!(
                "finalized snapshot moved backwards: indexed={}, requested={snapshot_number}",
                indexed.number
            )),
            Some(indexed) if snapshot_number == indexed.number => Ok(None),
            Some(indexed) => Ok(Some(
                indexed.number.checked_add(1).ok_or_else(|| eyre!("canonical log index block overflow"))?,
            )),
        }
    }

    fn verify_cached_anchor(&mut self, observed_hash: H256) -> Result<()> {
        let Some(indexed) = self.indexed_snapshot else {
            return Ok(());
        };
        if indexed.hash != observed_hash {
            self.clear();
            return Err(eyre!("canonical log index anchor changed at block {}; cache discarded", indexed.number));
        }
        Ok(())
    }

    fn batch_index(log: &Log) -> Result<u64> {
        if log.removed.unwrap_or_default() {
            return Err(eyre!("removed log returned for finalized canonical range"));
        }
        identity(log).ok_or_else(|| eyre!("canonical Rollup log has incomplete identity"))?;
        if log.transaction_hash.is_none() || log.block_hash.is_none() {
            return Err(eyre!("canonical Rollup log has incomplete block/transaction identity"));
        }
        if log.topics.len() < 3 {
            return Err(eyre!("canonical Rollup log is missing indexed topics"));
        }
        let batch_index = U256::from_big_endian(log.topics[1].as_bytes());
        if batch_index > U256::from(u64::MAX) {
            return Err(eyre!("Rollup batch index exceeds u64"));
        }
        Ok(batch_index.as_u64())
    }

    fn append_logs(&mut self, snapshot_number: u64, snapshot_hash: H256, commits: Vec<Log>, reverts: Vec<Log>) -> Result<()> {
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
        self.indexed_snapshot = Some(IndexedSnapshot {
            number: snapshot_number,
            hash: snapshot_hash,
        });
        Ok(())
    }

    fn resolve_cached(&self, batch_index: u64, canonical_hash: H256) -> Option<CanonicalCommit> {
        select_canonical_commit(
            self.commits.get(&batch_index).map(Vec::as_slice).unwrap_or_default(),
            self.reverts.get(&batch_index).map(Vec::as_slice).unwrap_or_default(),
            canonical_hash,
        )
    }

    pub async fn refresh<P: Middleware + 'static>(&mut self, provider: &P, rollup: Address, snapshot_number: U64, snapshot_hash: H256) -> Result<()> {
        let before = provider
            .get_block(snapshot_number)
            .await?
            .ok_or_else(|| eyre!("snapshot block {} not found", snapshot_number))?;
        if before.hash != Some(snapshot_hash) {
            return Err(eyre!("finalized snapshot hash changed before canonical log refresh"));
        }

        if let Some(indexed) = self.indexed_snapshot {
            let anchor = provider
                .get_block(U64::from(indexed.number))
                .await?
                .ok_or_else(|| eyre!("canonical log index anchor block disappeared"))?;
            let anchor_hash = anchor.hash.ok_or_else(|| eyre!("canonical log index anchor has no hash"))?;
            self.verify_cached_anchor(anchor_hash)?;
        }

        let query_start = match self.query_start(snapshot_number.as_u64()) {
            Ok(query_start) => query_start,
            Err(error) => {
                self.clear();
                return Err(error);
            }
        };
        let Some(mut from) = query_start else {
            self.verify_cached_anchor(snapshot_hash)?;
            return Ok(());
        };

        let commit_topic = H256::from(keccak256("CommitBatch(uint256,bytes32)"));
        let revert_topic = H256::from(keccak256("RevertBatch(uint256,bytes32)"));
        let mut commits = Vec::new();
        let mut reverts = Vec::new();
        while from <= snapshot_number.as_u64() {
            let to = from.saturating_add(QUERY_RANGE - 1).min(snapshot_number.as_u64());
            let base = Filter::new().address(rollup).from_block(from).to_block(to);
            commits.extend(provider.get_logs(&base.clone().topic0(commit_topic)).await?);
            reverts.extend(provider.get_logs(&base.topic0(revert_topic)).await?);
            if to == snapshot_number.as_u64() {
                break;
            }
            from = to.checked_add(1).ok_or_else(|| eyre!("canonical log query block overflow"))?;
        }

        let after = provider
            .get_block(snapshot_number)
            .await?
            .ok_or_else(|| eyre!("snapshot block {} disappeared", snapshot_number))?;
        if after.hash != Some(snapshot_hash) {
            return Err(eyre!("snapshot block hash changed during canonical log refresh"));
        }

        self.append_logs(snapshot_number.as_u64(), snapshot_hash, commits, reverts)
    }

    pub async fn resolve<M: Middleware + 'static>(
        &self,
        rollup: &Rollup<M>,
        batch_index: u64,
        snapshot_number: U64,
        snapshot_hash: H256,
    ) -> Result<Option<CanonicalCommit>> {
        if self.indexed_snapshot
            != Some(IndexedSnapshot {
                number: snapshot_number.as_u64(),
                hash: snapshot_hash,
            })
        {
            return Err(eyre!("canonical log index is not anchored to the requested snapshot"));
        }
        let block_id = BlockId::Number(BlockNumber::Number(snapshot_number));
        let canonical_hash: H256 = rollup.committed_batches(U256::from(batch_index)).block(block_id).call().await?.into();
        if canonical_hash.is_zero() {
            return Ok(None);
        }
        Ok(self.resolve_cached(batch_index, canonical_hash))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn batch_log(batch: u64, block: u64, tx: u64, index: u64, hash: H256, marker: u64) -> Log {
        Log {
            block_number: Some(block.into()),
            transaction_index: Some(tx.into()),
            log_index: Some(index.into()),
            transaction_hash: Some(H256::from_low_u64_be(marker)),
            block_hash: Some(H256::from_low_u64_be(block)),
            topics: vec![H256::zero(), H256::from_low_u64_be(batch), hash],
            ..Default::default()
        }
    }

    fn log(block: u64, tx: u64, index: u64, hash: H256, marker: u64) -> Log {
        batch_log(7, block, tx, index, hash, marker)
    }

    #[test]
    fn rejects_short_and_unknown_calldata() {
        assert!(decode_batch_data_input(&[]).is_err());
        assert!(decode_batch_data_input(&[0xde, 0xad, 0xbe, 0xef]).is_err());
    }

    #[test]
    fn decodes_shared_pre_and_post_submitter_fixtures() {
        let corpus: serde_json::Value =
            serde_json::from_str(include_str!("../../../../node/derivation/testdata/rollup_calldata_fixtures.json")).unwrap();
        let mut decoded_count = 0;
        for fixture in corpus["fixtures"].as_array().unwrap() {
            let epoch = fixture["epoch"].as_str().unwrap();
            if epoch != "pre-submitter" && epoch != "current" {
                continue;
            }
            let data = ethers::utils::hex::decode(fixture["data"].as_str().unwrap().trim_start_matches("0x")).unwrap();
            let decoded = decode_batch_data_input(&data).unwrap_or_else(|err| panic!("{}: {err:#}", fixture["name"]));
            assert_eq!(decoded.version as u64, fixture["expected"]["version"].as_u64().unwrap());
            assert_eq!(decoded.last_block_number, fixture["expected"]["lastBlockNumber"].as_u64().unwrap());
            decoded_count += 1;
        }
        assert_eq!(decoded_count, 6);
    }

    #[test]
    fn same_hash_recommit_after_revert_is_selected_by_identity() {
        let hash = H256::from_low_u64_be(99);
        let resolved = select_canonical_commit(&[log(10, 0, 0, hash, 1), log(12, 0, 0, hash, 3)], &[log(11, 0, 0, hash, 2)], hash).unwrap();
        assert_eq!(resolved.transaction_hash, H256::from_low_u64_be(3));
    }

    #[test]
    fn zero_hash_and_same_block_revert_order_fail_closed() {
        let hash = H256::from_low_u64_be(99);
        assert!(select_canonical_commit(&[log(10, 0, 0, hash, 1)], &[], H256::zero()).is_none());

        let before_revert = log(10, 0, 1, hash, 1);
        let revert = log(10, 0, 2, hash, 2);
        assert!(select_canonical_commit(&[before_revert], std::slice::from_ref(&revert), hash).is_none());

        let after_revert = log(10, 0, 3, hash, 3);
        let resolved = select_canonical_commit(&[after_revert], &[revert], hash).unwrap();
        assert_eq!(resolved.transaction_hash, H256::from_low_u64_be(3));
    }

    #[test]
    fn compatibility_words_reject_malformed_authority_and_ignore_legacy_bitmap_type() {
        let address = Address::from_low_u64_be(7);
        let mut address_word = [0u8; 32];
        address_word[12..].copy_from_slice(address.as_bytes());
        assert_eq!(decode_address_word(&address_word).unwrap(), address);
        assert!(decode_address_word(&address_word[..31]).is_err());
        address_word[0] = 1;
        assert!(decode_address_word(&address_word).is_err());

        let mut old_tuple = [0u8; 128];
        U256::from(123).to_big_endian(&mut old_tuple[64..96]);
        old_tuple[96..128].fill(0xff); // legacy uint256 bitmap, not an ABI address
        assert_eq!(decode_batch_data_store_block_number(&old_tuple).unwrap(), U256::from(123));
        assert!(decode_batch_data_store_block_number(&old_tuple[..127]).is_err());
    }

    #[test]
    fn canonical_log_index_reuses_snapshot_and_only_plans_new_suffix() {
        let snapshot_hash = H256::from_low_u64_be(100);
        let mut index = CanonicalLogIndex::new(5).unwrap();
        assert_eq!(index.query_start(10).unwrap(), Some(5));

        let hash7 = H256::from_low_u64_be(7);
        let hash8 = H256::from_low_u64_be(8);
        index
            .append_logs(
                10,
                snapshot_hash,
                vec![batch_log(7, 7, 0, 0, hash7, 7), batch_log(8, 8, 0, 0, hash8, 8)],
                vec![],
            )
            .unwrap();

        assert!(index.resolve_cached(7, hash7).is_some());
        assert!(index.resolve_cached(8, hash8).is_some());
        assert_eq!(index.query_start(10).unwrap(), None);
        assert_eq!(index.query_start(12).unwrap(), Some(11));

        let hash9 = H256::from_low_u64_be(9);
        index
            .append_logs(12, H256::from_low_u64_be(102), vec![batch_log(9, 11, 0, 0, hash9, 9)], vec![])
            .unwrap();
        assert!(index.resolve_cached(7, hash7).is_some());
        assert!(index.resolve_cached(9, hash9).is_some());
    }

    #[test]
    fn canonical_log_index_discards_cache_and_fails_closed_on_anchor_reorg() {
        let old_hash = H256::from_low_u64_be(100);
        let mut index = CanonicalLogIndex::new(5).unwrap();
        index
            .append_logs(10, old_hash, vec![batch_log(7, 7, 0, 0, H256::from_low_u64_be(7), 7)], vec![])
            .unwrap();

        let error = index.verify_cached_anchor(H256::from_low_u64_be(101)).unwrap_err();
        assert!(error.to_string().contains("cache discarded"));
        assert!(index.resolve_cached(7, H256::from_low_u64_be(7)).is_none());
        assert_eq!(index.query_start(12).unwrap(), Some(5));
    }
}
