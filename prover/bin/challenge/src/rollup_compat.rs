use ethers::{
    abi::AbiDecode,
    providers::Middleware,
    types::{transaction::eip2718::TypedTransaction, Address, BlockId, BlockNumber, Bytes, Filter, Log, TransactionRequest, TxHash, H256, U256, U64},
    utils::{id, keccak256},
};
use eyre::{eyre, Result};

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

pub async fn resolve_canonical_commit<M, P>(
    rollup: &Rollup<M>,
    provider: &P,
    batch_index: u64,
    rollup_deployed_block: u64,
    snapshot_number: U64,
) -> Result<Option<CanonicalCommit>>
where
    M: Middleware + 'static,
    P: Middleware + 'static,
{
    if rollup_deployed_block == 0 || rollup_deployed_block > snapshot_number.as_u64() {
        return Err(eyre!(
            "invalid Rollup deployment block {rollup_deployed_block} for snapshot {snapshot_number}"
        ));
    }
    let before = provider
        .get_block(snapshot_number)
        .await?
        .ok_or_else(|| eyre!("snapshot block {} not found", snapshot_number))?;
    let before_hash = before.hash.ok_or_else(|| eyre!("snapshot block has no hash"))?;
    let block_id = BlockId::Number(BlockNumber::Number(snapshot_number));
    let canonical_hash: H256 = rollup.committed_batches(U256::from(batch_index)).block(block_id).call().await?.into();
    if canonical_hash.is_zero() {
        return Ok(None);
    }

    let commit_topic = H256::from(keccak256("CommitBatch(uint256,bytes32)"));
    let revert_topic = H256::from(keccak256("RevertBatch(uint256,bytes32)"));
    let mut commits = Vec::new();
    let mut reverts = Vec::new();
    let mut from = rollup_deployed_block;
    while from <= snapshot_number.as_u64() {
        let to = from.saturating_add(QUERY_RANGE - 1).min(snapshot_number.as_u64());
        let base = Filter::new()
            .address(rollup.address())
            .from_block(from)
            .to_block(to)
            .topic1(U256::from(batch_index));
        commits.extend(provider.get_logs(&base.clone().topic0(commit_topic)).await?);
        reverts.extend(provider.get_logs(&base.topic0(revert_topic)).await?);
        if to == snapshot_number.as_u64() {
            break;
        }
        from = to + 1;
    }

    let after = provider
        .get_block(snapshot_number)
        .await?
        .ok_or_else(|| eyre!("snapshot block {} disappeared", snapshot_number))?;
    if after.hash != Some(before_hash) {
        return Err(eyre!("snapshot block hash changed during canonical resolution"));
    }

    Ok(select_canonical_commit(&commits, &reverts, canonical_hash))
}

#[cfg(test)]
mod tests {
    use super::*;

    fn log(block: u64, tx: u64, index: u64, hash: H256, marker: u64) -> Log {
        Log {
            block_number: Some(block.into()),
            transaction_index: Some(tx.into()),
            log_index: Some(index.into()),
            transaction_hash: Some(H256::from_low_u64_be(marker)),
            block_hash: Some(H256::from_low_u64_be(block)),
            topics: vec![H256::zero(), H256::from_low_u64_be(7), hash],
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
}
