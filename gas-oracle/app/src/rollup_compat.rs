use ethers::{
    abi::AbiDecode,
    contract::EthCall,
    providers::Middleware,
    types::{Address, BlockId, BlockNumber, Filter, Log, TxHash, H256, U256, U64},
    utils::keccak256,
};
use eyre::{eyre, Result};

use crate::abi::{
    pre_submitter_rollup_abi::{
        CommitBatchCall as PreSubmitterCommitBatchCall,
        CommitBatchWithProofCall as PreSubmitterCommitBatchWithProofCall,
        CommitStateCall as PreSubmitterCommitStateCall,
    },
    rollup_abi::{CommitBatchCall, CommitBatchWithProofCall, CommitStateCall, Rollup},
};

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
#[allow(clippy::enum_variant_names)]
pub enum SubmissionKind {
    CommitBatch,
    CommitState,
    CommitBatchWithProof,
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct BatchDataInput {
    pub version: u8,
    pub last_block_number: u64,
    pub kind: SubmissionKind,
    pub pre_submitter: bool,
}

macro_rules! decoded {
    ($input:expr, $kind:expr, $pre:expr) => {{
        let input = $input;
        BatchDataInput {
            version: input.version,
            last_block_number: input.last_block_number,
            kind: $kind,
            pre_submitter: $pre,
        }
    }};
}

/// Decode only the batch fields consumed by the oracle. The selector is
/// matched before ABI decoding so malformed or unknown calldata fails closed.
pub fn decode_batch_data_input(data: &[u8]) -> Result<BatchDataInput> {
    let selector =
        data.get(..4).ok_or_else(|| eyre!("rollup calldata is shorter than a selector"))?;

    if selector == CommitBatchCall::selector() {
        let call = CommitBatchCall::decode(data)?;
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatch, false));
    }
    if selector == CommitStateCall::selector() {
        let call = CommitStateCall::decode(data)?;
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitState, false));
    }
    if selector == CommitBatchWithProofCall::selector() {
        let call = CommitBatchWithProofCall::decode(data)?;
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatchWithProof, false));
    }
    if selector == PreSubmitterCommitBatchCall::selector() {
        let call = PreSubmitterCommitBatchCall::decode(data)?;
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitBatch, true));
    }
    if selector == PreSubmitterCommitStateCall::selector() {
        let call = PreSubmitterCommitStateCall::decode(data)?;
        return Ok(decoded!(call.batch_data_input, SubmissionKind::CommitState, true));
    }
    if selector == PreSubmitterCommitBatchWithProofCall::selector() {
        let call = PreSubmitterCommitBatchWithProofCall::decode(data)?;
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

/// Select the last same-batch Commit after the last Revert. Full log identity
/// deliberately distinguishes a same-hash recommit from the reverted event.
pub fn select_canonical_commit(
    commits: &[Log],
    reverts: &[Log],
    canonical_hash: H256,
) -> Option<CanonicalCommit> {
    if canonical_hash.is_zero() {
        return None;
    }
    let last_revert =
        reverts.iter().filter(|log| !log.removed.unwrap_or_default()).filter_map(identity).max();

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

pub const LOG_QUERY_RANGE: u64 = 10_000;

pub async fn query_batch_logs<M: Middleware>(
    provider: &M,
    rollup: Address,
    batch_index: u64,
    event_topic: H256,
    from_block: u64,
    to_block: u64,
) -> Result<Vec<Log>> {
    if from_block > to_block {
        return Ok(Vec::new());
    }
    let mut logs = Vec::new();
    let mut from = from_block;
    loop {
        let to = from.saturating_add(LOG_QUERY_RANGE - 1).min(to_block);
        let filter = Filter::new()
            .address(rollup)
            .from_block(from)
            .to_block(to)
            .topic0(event_topic)
            .topic1(U256::from(batch_index));
        logs.extend(provider.get_logs(&filter).await.map_err(|err| eyre!(err.to_string()))?);
        if to == to_block {
            break;
        }
        from = to + 1;
    }
    Ok(logs)
}

pub async fn resolve_canonical_commit<M: Middleware + 'static>(
    rollup: &Rollup<M>,
    provider: &M,
    batch_index: u64,
    from_block: u64,
    snapshot_number: U64,
) -> Result<Option<CanonicalCommit>> {
    let before = provider
        .get_block(snapshot_number)
        .await
        .map_err(|err| eyre!(err.to_string()))?
        .ok_or_else(|| eyre!("snapshot block {} not found", snapshot_number))?;
    let before_hash = before.hash.ok_or_else(|| eyre!("snapshot block has no hash"))?;
    let block_id = BlockId::Number(BlockNumber::Number(snapshot_number));
    let canonical_hash: H256 = rollup
        .committed_batches(U256::from(batch_index))
        .block(block_id)
        .call()
        .await
        .map_err(|err| eyre!(err.to_string()))?
        .into();
    if canonical_hash.is_zero() {
        return Ok(None);
    }

    let commit_topic = H256::from(keccak256("CommitBatch(uint256,bytes32)"));
    let revert_topic = H256::from(keccak256("RevertBatch(uint256,bytes32)"));
    let commits = query_batch_logs(
        provider,
        rollup.address(),
        batch_index,
        commit_topic,
        from_block,
        snapshot_number.as_u64(),
    )
    .await?;
    let reverts = query_batch_logs(
        provider,
        rollup.address(),
        batch_index,
        revert_topic,
        from_block,
        snapshot_number.as_u64(),
    )
    .await?;

    let after = provider
        .get_block(snapshot_number)
        .await
        .map_err(|err| eyre!(err.to_string()))?
        .ok_or_else(|| eyre!("snapshot block {} disappeared", snapshot_number))?;
    if after.hash != Some(before_hash) {
        return Err(eyre!("snapshot block hash changed during canonical resolution"));
    }

    Ok(select_canonical_commit(&commits, &reverts, canonical_hash))
}

#[cfg(test)]
mod tests {
    use super::*;

    fn log(block: u64, marker: u64, hash: H256) -> Log {
        Log {
            block_number: Some(block.into()),
            transaction_index: Some(0.into()),
            log_index: Some(0.into()),
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
        let corpus: serde_json::Value = serde_json::from_str(include_str!(
            "../../../node/derivation/testdata/rollup_calldata_fixtures.json"
        ))
        .unwrap();
        let mut decoded_count = 0;
        for fixture in corpus["fixtures"].as_array().unwrap() {
            let epoch = fixture["epoch"].as_str().unwrap();
            if epoch != "pre-submitter" && epoch != "current" {
                continue;
            }
            let data = ethers::utils::hex::decode(
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
    fn selector_sets_are_frozen() {
        assert_eq!(CommitBatchCall::selector(), [0x41, 0xf7, 0x56, 0xda]);
        assert_eq!(CommitStateCall::selector(), [0x67, 0xca, 0xa3, 0x7a]);
        assert_eq!(CommitBatchWithProofCall::selector(), [0x15, 0x44, 0xba, 0x3a]);
        assert_eq!(PreSubmitterCommitBatchCall::selector(), [0x42, 0x88, 0x68, 0xb5]);
        assert_eq!(PreSubmitterCommitStateCall::selector(), [0x1e, 0x88, 0x25, 0xbe]);
        assert_eq!(PreSubmitterCommitBatchWithProofCall::selector(), [0x4e, 0x8f, 0x1d, 0x67]);
    }

    #[test]
    fn same_hash_recommit_after_revert_wins() {
        let hash = H256::from_low_u64_be(99);
        let resolved = select_canonical_commit(
            &[log(10, 1, hash), log(12, 3, hash)],
            &[log(11, 2, hash)],
            hash,
        )
        .unwrap();
        assert_eq!(resolved.transaction_hash, H256::from_low_u64_be(3));
    }
}
