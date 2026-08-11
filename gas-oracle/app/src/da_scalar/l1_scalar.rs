use std::{
    collections::{BTreeSet, HashMap, HashSet},
    sync::Arc,
};

use super::{
    blob_client::BeaconNode,
    calculate::{data_and_hashes_from_txs, extract_tx_payload, extract_txn_count},
    error::ScalarError,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};
use crate::{
    abi::{gas_price_oracle_abi::GasPriceOracle, rollup_abi::Rollup},
    metrics::ORACLE_SERVICE_METRICS,
    rollup_compat::{
        decode_batch_data_input, resolve_canonical_commit, CanonicalCommit, SubmissionKind,
        LOG_QUERY_RANGE,
    },
    signer::send_transaction,
};
use ethers::{
    prelude::*,
    utils::{hex, keccak256},
};
use eyre::anyhow;
use remote_signer_client::SignerClient;
use serde_json::Value;

const PRECISION: u64 = 10u64.pow(9);
const MAX_COMMIT_SCALAR: u64 = 10u64.pow(9 + 6);
const MAX_BLOB_SCALAR: u64 = 10u64.pow(9 + 2);
const SAMPLE_LOOKBACK_BLOCKS: u64 = 100;
const ZERO_VERSIONED_HASH: H256 = H256([
    0x01, 0x06, 0x57, 0xf3, 0x75, 0x54, 0xc7, 0x81, 0x40, 0x2a, 0x22, 0x91, 0x7d, 0xee, 0x2f, 0x75,
    0xde, 0xf7, 0xab, 0x96, 0x6d, 0x7b, 0x77, 0x09, 0x05, 0x39, 0x8e, 0xba, 0x3c, 0x44, 0x40, 0x14,
]);

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct BlobSource {
    transaction_hash: TxHash,
    block_number: U64,
    block_hash: H256,
    transaction_index: U64,
    blob_hashes: Vec<H256>,
}

type HistoricalCommitIndex = HashMap<u64, Vec<Log>>;

fn index_commit_logs(logs: Vec<Log>) -> HistoricalCommitIndex {
    let mut index = HistoricalCommitIndex::new();
    for log in logs {
        if log.removed.unwrap_or_default() {
            continue;
        }
        let Some(topic) = log.topics.get(1) else {
            continue;
        };
        let batch_index = U256::from_big_endian(topic.as_bytes());
        if batch_index > U256::from(u64::MAX) {
            continue;
        }
        index.entry(batch_index.as_u64()).or_default().push(log);
    }
    index
}

fn batch_indices_from_commit_logs(logs: &[Log], max_batch_index: u64) -> BTreeSet<u64> {
    logs.iter()
        .filter(|log| !log.removed.unwrap_or_default())
        .filter_map(|log| log.topics.get(1))
        .filter_map(|topic| {
            let value = U256::from_big_endian(topic.as_bytes());
            (value <= U256::from(max_batch_index)).then(|| value.as_u64())
        })
        .filter(|index| *index > 0)
        .collect()
}

fn ensure_snapshot_hash(expected: H256, actual: Option<H256>) -> Result<(), ScalarError> {
    if actual != Some(expected) {
        return Err(ScalarError::Error(anyhow!(
            "finalized snapshot hash changed; discard the complete DA sample round"
        )));
    }
    Ok(())
}

fn transaction_blob_hashes(transaction: &Transaction) -> Vec<H256> {
    transaction
        .other
        .get_with("blobVersionedHashes", serde_json::from_value::<Vec<H256>>)
        .unwrap_or(Ok(Vec::new()))
        .unwrap_or_default()
}

fn receipt_matches_source(receipt: &TransactionReceipt, source: &BlobSource) -> bool {
    receipt.transaction_hash == source.transaction_hash
        && receipt.block_number == Some(source.block_number)
        && receipt.block_hash == Some(source.block_hash)
        && receipt.transaction_index == source.transaction_index
}

fn blob_commitment(version: u8, hashes: &[H256]) -> Result<H256, ScalarError> {
    match version {
        0 | 1 if hashes.len() == 1 => Ok(hashes[0]),
        0 | 1 => Err(ScalarError::Error(anyhow!(format!(
            "legacy batch must have exactly one blob, got {}",
            hashes.len()
        )))),
        2 if !hashes.is_empty() => {
            let bytes: Vec<u8> =
                hashes.iter().flat_map(|hash| hash.as_bytes().iter().copied()).collect();
            Ok(H256::from(keccak256(bytes)))
        }
        2 => Err(ScalarError::Error(anyhow!("V2 batch has no blob"))),
        _ => Err(ScalarError::Error(anyhow!(format!("unsupported batch version {version}")))),
    }
}

fn validate_fresh_blob_commitment(
    stored_commitment: H256,
    transaction_commitment: H256,
) -> Result<(), ScalarError> {
    // Rollup garbage-collects the mapping after successor finalization. A
    // direct canonical blob transaction remains a stable source even when the
    // snapshot getter has already returned zero.
    if !stored_commitment.is_zero() && stored_commitment != transaction_commitment {
        return Err(ScalarError::Error(anyhow!(format!(
            "canonical blob commitment mismatch: stored={stored_commitment:?}, tx={transaction_commitment:?}"
        ))));
    }
    Ok(())
}

fn recover_gc_commitment(snapshot: H256, canonical_block: H256) -> H256 {
    if snapshot.is_zero() {
        canonical_block
    } else {
        snapshot
    }
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
enum BloblessDisposition {
    ExcludeMissing,
    ExcludeZero,
    ResolveStored(H256),
}

fn classify_blobless_submission(
    version: u8,
    kind: SubmissionKind,
    pre_submitter: bool,
    stored_commitment: H256,
) -> Result<BloblessDisposition, ScalarError> {
    if kind == SubmissionKind::CommitBatch && !pre_submitter {
        return Err(ScalarError::Error(anyhow!(
            "current canonical commitBatch has no blob; refusing zero-data scalar sample"
        )));
    }

    if version <= 1 && stored_commitment == ZERO_VERSIONED_HASH {
        // Before the Submitter cutover, commitBatch legitimately accepted a
        // V0/V1 transaction without a blob and stored ZERO_VERSIONED_HASH.
        // Permissionless proof/state submissions can also reuse that value.
        return Ok(BloblessDisposition::ExcludeZero);
    }
    if stored_commitment.is_zero() {
        return Ok(BloblessDisposition::ExcludeMissing);
    }
    if kind == SubmissionKind::CommitBatch {
        return Err(ScalarError::Error(anyhow!(
            "legacy canonical commitBatch is blobless with a non-zero stored commitment"
        )));
    }
    Ok(BloblessDisposition::ResolveStored(stored_commitment))
}

// Main struct to manage overhead information
pub struct ScalarUpdater {
    l1_provider: Provider<Http>, // L1 provider for HTTP connections
    l2_provider: Provider<Http>,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>, // L2 gasPrice Oracle
    ext_signer: Option<SignerClient>,
    l1_rollup: Rollup<Provider<Http>>, // Rollup object for L1
    beacon_node: BeaconNode,           // Beacon node for blockchain
    l1_rollup_deploy_block: u64,
    gas_threshold: u64,
    commit_scalar_buffer: u64,
    blob_scalar_buffer: u64,
    finalize_batch_gas_used: u64,
    txn_per_batch: u64,
}

impl ScalarUpdater {
    // Constructor to initialize an OverHead object
    #[allow(clippy::too_many_arguments)]
    pub fn new(
        l1_provider: Provider<Http>,
        l2_provider: Provider<Http>,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        ext_signer: Option<SignerClient>,
        l1_rollup: Rollup<Provider<Http>>,
        l1_beacon_rpc: String,
        l1_rollup_deploy_block: u64,
        gas_threshold: u64,
        commit_scalar_buffer: u64,
        blob_scalar_buffer: u64,
        finalize_batch_gas_used: u64,
        txn_per_batch: u64,
    ) -> Self {
        // Create beacon nodes with provided RPC URLs
        let beacon_node = BeaconNode { rpc_url: l1_beacon_rpc };

        // Return a new OverHead instance with initialized values
        Self {
            l1_provider,
            l2_provider,
            l2_oracle,
            ext_signer,
            l1_rollup,
            beacon_node,
            l1_rollup_deploy_block,
            gas_threshold,
            commit_scalar_buffer,
            blob_scalar_buffer,
            finalize_batch_gas_used,
            txn_per_batch,
        }
    }

    /// Update commitScalar and blobScalar.
    /// Calculate the user's average cost of the latest rollup and set it to the GasPriceOrale
    /// contract on the L2 network.
    pub async fn update(&mut self) -> Result<(), ScalarError> {
        // Step1. Pin all reads to one finalized L1 block. A latest-head sample
        // can be reorged after the oracle has already written its L2 scalar.
        let snapshot = self
            .l1_provider
            .get_block(BlockNumber::Finalized)
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "overhead.l1_provider.get finalized block error: {:#?}",
                    e
                )))
            })?
            .ok_or_else(|| ScalarError::Error(anyhow!("finalized L1 block is unavailable")))?;
        let snapshot_number = snapshot
            .number
            .ok_or_else(|| ScalarError::Error(anyhow!("finalized L1 block has no number")))?;
        let snapshot_hash = snapshot
            .hash
            .ok_or_else(|| ScalarError::Error(anyhow!("finalized L1 block has no hash")))?;
        if self.l1_rollup_deploy_block == 0
            || self.l1_rollup_deploy_block > snapshot_number.as_u64()
        {
            return Err(ScalarError::Error(anyhow!(format!(
                "configured Rollup deployment block {} is outside finalized snapshot 1..={}",
                self.l1_rollup_deploy_block, snapshot_number
            ))));
        }
        let start = if snapshot_number > U64::from(SAMPLE_LOOKBACK_BLOCKS) {
            snapshot_number - U64::from(SAMPLE_LOOKBACK_BLOCKS)
        } else {
            U64::from(1)
        };

        let (mut commit_scalar, mut blob_scalar) =
            match self.calculate_scalar(start.as_u64(), snapshot_number, snapshot_hash).await {
                Ok(Some(scalar)) => scalar,
                Ok(None) => {
                    return Ok(());
                }
                Err(e) => return Err(e),
            };

        // Step2. fetch current scalar on l2.
        let current_commit_scalar: U256 = self.l2_oracle.commit_scalar().await.map_err(|e| {
            ScalarError::Error(anyhow!(format!("query l2_oracle.commit_scalar error: {:#?}", e)))
        })?;

        let current_blob_scalar: U256 = self.l2_oracle.blob_scalar().await.map_err(|e| {
            ScalarError::Error(anyhow!(format!("query l2_oracle.blob_scalar error: {:#?}", e)))
        })?;

        log::info!("set_commit_or_blob_scalar, latest commit_scalar: {:?}, latest blob_scalar: {:?}, current_commit_scalar on l2 is: {:?}, current_blob_scalar on l2 is: {:?}", 
        commit_scalar, blob_scalar, current_commit_scalar.as_u64(), current_blob_scalar.as_u64());

        // Fine tune the actual value
        commit_scalar += self.commit_scalar_buffer;
        blob_scalar += self.blob_scalar_buffer;

        commit_scalar = commit_scalar.min(MAX_COMMIT_SCALAR);
        blob_scalar = blob_scalar.min(MAX_BLOB_SCALAR);

        ORACLE_SERVICE_METRICS.commit_scalar.set((commit_scalar / PRECISION) as i64);
        ORACLE_SERVICE_METRICS
            .blob_scalar
            .set((1000.0 * blob_scalar as f64 / PRECISION as f64).round() / 1000.0);

        // Step3. set on L2chain
        let client: Arc<SignerMiddleware<Provider<Http>, LocalWallet>> = self.l2_oracle.client();
        if self.check_threshold_reached(
            commit_scalar,
            current_commit_scalar.as_u64(),
            "commit_scalar",
        ) {
            // Update commit_scalar
            let call = self.l2_oracle.set_commit_scalar(U256::from(commit_scalar));
            let tx_hash = send_transaction(call, &client, &self.ext_signer, &self.l2_provider)
                .await
                .map_err(|e| {
                    ScalarError::Error(anyhow!(format!("set_commit_scalar error: {:#?}", e)))
                })?;
            log::info!("set_commit_scalar success, tx_hash: {:#?}", tx_hash);
        }

        if self.check_threshold_reached(blob_scalar, current_blob_scalar.as_u64(), "blob_scalar") {
            // Update blob_scalar
            let call = self.l2_oracle.set_blob_scalar(U256::from(blob_scalar));
            let tx_hash = send_transaction(call, &client, &self.ext_signer, &self.l2_provider)
                .await
                .map_err(|e| {
                    ScalarError::Error(anyhow!(format!("set_blob_scalar error: {:#?}", e)))
                })?;
            log::info!("set_blob_scalar success, tx_hash: {:#?}", tx_hash);
        }

        Ok(())
    }

    fn check_threshold_reached(&mut self, latest: u64, current: u64, state_var_name: &str) -> bool {
        let actual_change = latest.abs_diff(current);
        let expected_change = current * self.gas_threshold / 100;
        let need_update = actual_change > expected_change;
        log::info!(
            "update {}, actual_change: {:?}, expected_change: {:?}, need_update: {:?}",
            state_var_name,
            actual_change,
            expected_change,
            need_update
        );
        need_update
    }

    async fn calculate_scalar(
        &mut self,
        start: u64,
        snapshot_number: U64,
        snapshot_hash: H256,
    ) -> Result<Option<(u64, u64)>, ScalarError> {
        let block_id = BlockId::Number(BlockNumber::Number(snapshot_number));
        let last_batch_index = self
            .l1_rollup
            .last_committed_batch_index()
            .block(block_id)
            .call()
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("query last committed batch: {e:?}"))))?
            .as_u64();
        if last_batch_index == 0 {
            log::warn!("no non-genesis batch exists at finalized L1 snapshot, skip update");
            return Ok(None);
        }

        let commit_topic = H256::from(keccak256("CommitBatch(uint256,bytes32)"));
        let candidate_logs = self
            .l1_provider
            .get_logs(
                &Filter::new()
                    .address(self.l1_rollup.address())
                    .from_block(start)
                    .to_block(snapshot_number)
                    .topic0(commit_topic),
            )
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!("query finalized Commit window: {e:?}")))
            })?;
        let batch_indices = batch_indices_from_commit_logs(&candidate_logs, last_batch_index);
        if batch_indices.is_empty() {
            log::warn!("no Commit event exists in the finalized DA sample window");
            return Ok(None);
        }

        let mut seen_sources = HashSet::new();
        let mut commit_scalar_sum = 0u128;
        let mut blob_scalar_sum = 0u128;
        let mut sample_count = 0u128;
        let mut historical_commits = None;

        for batch_index in batch_indices {
            let canonical = resolve_canonical_commit(
                &self.l1_rollup,
                &self.l1_provider,
                batch_index,
                start,
                snapshot_number,
            )
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "resolve canonical commit for batch {batch_index}: {e:#}"
                )))
            })?;
            let Some(canonical) = canonical else {
                // The window may contain a Commit that was canonically reverted
                // before the pinned snapshot. It contributes no DA sample.
                continue;
            };

            let canonical_tx = self.load_commit_transaction(canonical).await?;
            let input = decode_batch_data_input(canonical_tx.input.as_ref()).map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "decode canonical commit {:?}: {e:#}",
                    canonical.transaction_hash
                )))
            })?;
            let blob_source = self
                .resolve_blob_source(
                    batch_index,
                    canonical,
                    &canonical_tx,
                    input.version,
                    input.kind,
                    input.pre_submitter,
                    snapshot_number,
                    snapshot_hash,
                    &mut historical_commits,
                )
                .await?;
            let Some(blob_source) = blob_source else {
                ORACLE_SERVICE_METRICS.da_sample_exclusions.inc();
                log::warn!(
                    "canonical batch {} has no verifiable real DA blob source; excluded",
                    batch_index
                );
                continue;
            };
            if !seen_sources.insert(blob_source.clone()) {
                ORACLE_SERVICE_METRICS.da_sample_exclusions.inc();
                log::warn!(
                    "duplicate DA blob source {:?} in finalized window; counted once",
                    blob_source.transaction_hash
                );
                continue;
            }

            let (commit_scalar, blob_scalar) = self
                .calculate_from_rollup(blob_source, input.last_block_number)
                .await
                .inspect_err(|_| {
                    log::info!(
                        "scalar is unavailable, canonical_tx_hash ={:#?}",
                        canonical.transaction_hash
                    );
                })?;
            commit_scalar_sum += u128::from(commit_scalar);
            blob_scalar_sum += u128::from(blob_scalar);
            sample_count += 1;
        }

        let after = self
            .l1_provider
            .get_block(snapshot_number)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("recheck finalized snapshot: {e:?}"))))?
            .ok_or_else(|| ScalarError::Error(anyhow!("finalized snapshot disappeared")))?;
        ensure_snapshot_hash(snapshot_hash, after.hash)?;
        if sample_count == 0 {
            return Ok(None);
        }

        let commit_scalar = u64::try_from(commit_scalar_sum / sample_count)
            .map_err(|_| ScalarError::Error(anyhow!("averaged commit scalar overflow")))?;
        let blob_scalar = u64::try_from(blob_scalar_sum / sample_count)
            .map_err(|_| ScalarError::Error(anyhow!("averaged blob scalar overflow")))?;
        Ok(Some((commit_scalar, blob_scalar)))
    }

    async fn load_commit_transaction(
        &self,
        commit: CanonicalCommit,
    ) -> Result<Transaction, ScalarError> {
        let transaction = self
            .l1_provider
            .get_transaction(commit.transaction_hash)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("get commit transaction: {e:?}"))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "commit transaction {:?} is unavailable",
                    commit.transaction_hash
                )))
            })?;
        if transaction.to != Some(self.l1_rollup.address())
            || transaction.block_number != Some(commit.identity.block_number.into())
            || transaction.block_hash != Some(commit.block_hash)
            || transaction.transaction_index != Some(commit.identity.transaction_index.into())
        {
            return Err(ScalarError::Error(anyhow!(format!(
                "commit transaction {:?} moved to a different block",
                commit.transaction_hash
            ))));
        }
        Ok(transaction)
    }

    async fn load_canonical_block_commitment(
        &self,
        batch_index: u64,
        canonical: CanonicalCommit,
    ) -> Result<H256, ScalarError> {
        if canonical.identity.block_number < self.l1_rollup_deploy_block {
            return Err(ScalarError::Error(anyhow!(format!(
                "canonical commit block {} predates configured Rollup deployment block {}",
                canonical.identity.block_number, self.l1_rollup_deploy_block
            ))));
        }

        let canonical_number = U64::from(canonical.identity.block_number);
        let before = self
            .l1_provider
            .get_block(BlockNumber::Number(canonical_number))
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "load canonical commit block before historical getter: {e:?}"
                )))
            })?
            .ok_or_else(|| ScalarError::Error(anyhow!("canonical commit block disappeared")))?;
        if before.hash != Some(canonical.block_hash) {
            return Err(ScalarError::Error(anyhow!(
                "canonical commit block hash changed before historical getter"
            )));
        }

        // Use the block hash, rather than an unpinned number, for the historical
        // eth_call. This recovers the mapping value before later finalization GC.
        let commitment: H256 = self
            .l1_rollup
            .batch_blob_versioned_hashes(U256::from(batch_index))
            .block(BlockId::Hash(canonical.block_hash))
            .call()
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "query stored blob commitment at canonical commit block: {e:?}"
                )))
            })?
            .into();

        let after = self
            .l1_provider
            .get_block(BlockNumber::Number(canonical_number))
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!(
                    "recheck canonical commit block after historical getter: {e:?}"
                )))
            })?
            .ok_or_else(|| ScalarError::Error(anyhow!("canonical commit block disappeared")))?;
        if after.hash != Some(canonical.block_hash) {
            return Err(ScalarError::Error(anyhow!(
                "canonical commit block hash changed during historical getter"
            )));
        }
        Ok(commitment)
    }

    async fn load_historical_commit_index(
        &self,
        snapshot_number: U64,
    ) -> Result<HistoricalCommitIndex, ScalarError> {
        if self.l1_rollup_deploy_block > snapshot_number.as_u64() {
            return Err(ScalarError::Error(anyhow!(
                "Rollup deployment block is newer than the finalized snapshot"
            )));
        }

        // Replay CommitBatch once per scalar round and share the resulting
        // index between every stored-source batch. Replaying the entire chain
        // separately for each batch multiplies RPC load by the sample count.
        let commit_topic = H256::from(keccak256("CommitBatch(uint256,bytes32)"));
        let mut logs = Vec::new();
        let mut from = self.l1_rollup_deploy_block;
        let snapshot = snapshot_number.as_u64();
        loop {
            let to = from.saturating_add(LOG_QUERY_RANGE - 1).min(snapshot);
            logs.extend(
                self.l1_provider
                    .get_logs(
                        &Filter::new()
                            .address(self.l1_rollup.address())
                            .from_block(from)
                            .to_block(to)
                            .topic0(commit_topic),
                    )
                    .await
                    .map_err(|e| {
                        ScalarError::Error(anyhow!(format!(
                            "query historical blob-source commits in {from}..={to}: {e:?}"
                        )))
                    })?,
            );
            if to == snapshot {
                break;
            }
            from = to
                .checked_add(1)
                .ok_or_else(|| ScalarError::Error(anyhow!("commit replay block overflow")))?;
        }
        Ok(index_commit_logs(logs))
    }

    #[allow(clippy::too_many_arguments)]
    async fn resolve_blob_source(
        &self,
        batch_index: u64,
        canonical: CanonicalCommit,
        canonical_tx: &Transaction,
        version: u8,
        kind: SubmissionKind,
        pre_submitter: bool,
        snapshot_number: U64,
        snapshot_hash: H256,
        historical_commits: &mut Option<HistoricalCommitIndex>,
    ) -> Result<Option<BlobSource>, ScalarError> {
        let block_id = BlockId::Number(BlockNumber::Number(snapshot_number));
        let snapshot_commitment: H256 = self
            .l1_rollup
            .batch_blob_versioned_hashes(U256::from(batch_index))
            .block(block_id)
            .call()
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!("query stored blob commitment: {e:?}")))
            })?
            .into();

        let current_hashes = transaction_blob_hashes(canonical_tx);
        if !current_hashes.is_empty() {
            let actual_commitment = blob_commitment(version, &current_hashes)?;
            validate_fresh_blob_commitment(snapshot_commitment, actual_commitment)?;
            return Ok(Some(BlobSource {
                transaction_hash: canonical.transaction_hash,
                block_number: canonical.identity.block_number.into(),
                block_hash: canonical.block_hash,
                transaction_index: canonical.identity.transaction_index.into(),
                blob_hashes: current_hashes,
            }));
        }

        // A post-cutover commitBatch is required by the Rollup contract to
        // carry a blob. Keep that invariant fail-closed even if storage was GC'd.
        if kind == SubmissionKind::CommitBatch && !pre_submitter {
            return classify_blobless_submission(version, kind, pre_submitter, snapshot_commitment)
                .map(|_| None);
        }

        let canonical_block_commitment = if snapshot_commitment.is_zero() {
            self.load_canonical_block_commitment(batch_index, canonical).await?
        } else {
            snapshot_commitment
        };
        let expected_commitment =
            recover_gc_commitment(snapshot_commitment, canonical_block_commitment);
        match classify_blobless_submission(version, kind, pre_submitter, expected_commitment)? {
            BloblessDisposition::ExcludeZero => {
                log::warn!(
                    "canonical batch {} is a legal V0/V1 blobless ZERO_VERSIONED_HASH submission; excluded from DA scalar sampling",
                    batch_index
                );
                return Ok(None);
            }
            BloblessDisposition::ExcludeMissing => {
                log::warn!(
                    "batch {} stored blob commitment is zero at both finalized snapshot and canonical commit block; excluded and alert",
                    batch_index
                );
                return Ok(None);
            }
            BloblessDisposition::ResolveStored(commitment) => {
                debug_assert_eq!(commitment, expected_commitment);
            }
        }

        // Stored-source recommits intentionally look *before* the last Revert.
        // Canonical filtering here would discard the original blob transaction.
        if historical_commits.is_none() {
            *historical_commits = Some(self.load_historical_commit_index(snapshot_number).await?);
        }
        let mut candidates = historical_commits
            .as_ref()
            .and_then(|index| index.get(&batch_index))
            .cloned()
            .unwrap_or_default();
        candidates.retain(|log| {
            log.block_number
                .is_some_and(|number| number.as_u64() <= canonical.identity.block_number)
        });
        candidates.sort_by_key(|log| {
            (
                log.block_number.unwrap_or_default(),
                log.transaction_index.unwrap_or_default(),
                log.log_index.unwrap_or_default(),
            )
        });

        for candidate in candidates.into_iter().rev() {
            let Some(transaction_hash) = candidate.transaction_hash else {
                continue;
            };
            if transaction_hash == canonical.transaction_hash {
                continue;
            }
            let Some(block_number) = candidate.block_number else {
                continue;
            };
            let Some(block_hash) = candidate.block_hash else {
                continue;
            };
            let Some(transaction_index) = candidate.transaction_index else {
                continue;
            };
            let Some(transaction) =
                self.l1_provider.get_transaction(transaction_hash).await.map_err(|e| {
                    ScalarError::Error(anyhow!(format!("get blob-source transaction: {e:?}")))
                })?
            else {
                continue;
            };
            if transaction.to != Some(self.l1_rollup.address())
                || transaction.block_number != Some(block_number)
                || transaction.block_hash != Some(block_hash)
                || transaction.transaction_index != Some(transaction_index)
            {
                return Err(ScalarError::Error(anyhow!(format!(
                    "blob-source transaction {transaction_hash:?} changed block identity"
                ))));
            }
            let Ok(candidate_input) = decode_batch_data_input(transaction.input.as_ref()) else {
                continue;
            };
            if candidate_input.version != version {
                continue;
            }
            let hashes = transaction_blob_hashes(&transaction);
            if hashes.is_empty() {
                continue;
            }
            if blob_commitment(candidate_input.version, &hashes)? != expected_commitment {
                continue;
            }

            let after = self
                .l1_provider
                .get_block(snapshot_number)
                .await
                .map_err(|e| {
                    ScalarError::Error(anyhow!(format!("recheck finalized snapshot: {e:?}")))
                })?
                .ok_or_else(|| ScalarError::Error(anyhow!("finalized snapshot disappeared")))?;
            ensure_snapshot_hash(snapshot_hash, after.hash)?;
            log::info!(
                "resolved stored blob source: batch={}, canonical_tx={:?}, source_tx={:?}, source_block={}, blobs={}",
                batch_index,
                canonical.transaction_hash,
                transaction_hash,
                block_number,
                hashes.len()
            );
            return Ok(Some(BlobSource {
                transaction_hash,
                block_number,
                block_hash,
                transaction_index,
                blob_hashes: hashes,
            }));
        }

        log::warn!(
            "no historical blob source matches batch {} commitment {:?}; exclude and alert",
            batch_index,
            expected_commitment
        );
        Ok(None)
    }

    async fn calculate_from_rollup(
        &mut self,
        blob_source: BlobSource,
        last_block_number: u64,
    ) -> Result<(u64, u64), ScalarError> {
        //Step1. get_data_from_blob
        let (l2_data_len, num_blobs, l2_txn) =
            self.get_data_from_blob(&blob_source, last_block_number).await.map_err(|e| {
                log::error!("get_data_from_blob error: {:#?}", e);
                e
            })?;

        let source_tx_receipt = self
            .l1_provider
            .get_transaction_receipt(blob_source.transaction_hash)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get source transaction receipt return none, tx_hash= {:#?}",
                    blob_source.transaction_hash
                )))
            })?;
        if !receipt_matches_source(&source_tx_receipt, &blob_source) {
            return Err(ScalarError::Error(anyhow!(
                "original blob source receipt identity mismatch"
            )));
        }

        // rollup_gas_used
        let rollup_gas_used = source_tx_receipt.gas_used.unwrap_or_default();
        if rollup_gas_used.is_zero() {
            return Err(ScalarError::Error(anyhow!(format!(
                "blob tx calldata gas_used is none or 0, tx_hash = {:#?}",
                blob_source.transaction_hash
            ))));
        }

        //Step2. Calculate scalar
        let commit_scalar = (rollup_gas_used.as_u64() + self.finalize_batch_gas_used) * PRECISION
            / l2_txn.max(self.txn_per_batch);
        let blob_scalar = if l2_data_len > 0 {
            num_blobs.max(1) * MAX_BLOB_TX_PAYLOAD_SIZE as u64 * PRECISION / l2_data_len
        } else {
            MAX_BLOB_SCALAR
        };

        log::info!(
            "rollup_gas_used: {:?}, l2_txn: {:?}, l2_data_len:{:?}, commit_scalar: {:?}, blob_scalar: {:.4}",
            rollup_gas_used,
            l2_txn,
            l2_data_len,
            commit_scalar/PRECISION,
            blob_scalar as f64/PRECISION as f64,
        );

        // Set metric
        ORACLE_SERVICE_METRICS.txn_per_batch.set(l2_txn as f64);
        Ok((commit_scalar, blob_scalar))
    }

    async fn get_data_from_blob(
        &self,
        source: &BlobSource,
        last_block_num: u64,
    ) -> Result<(u64, u64, u64), ScalarError> {
        let blob_tx = self
            .l1_provider
            .get_transaction(source.transaction_hash)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get transaction return none, tx_hash: {:#?}",
                    source.transaction_hash
                )))
            })?;
        if blob_tx.to != Some(self.l1_rollup.address())
            || blob_tx.block_number != Some(source.block_number)
            || blob_tx.block_hash != Some(source.block_hash)
            || blob_tx.transaction_index != Some(source.transaction_index)
        {
            return Err(ScalarError::Error(anyhow!(
                "blob source transaction changed block identity"
            )));
        }

        let blob_block = self
            .l1_provider
            .get_block_with_txs(BlockNumber::Number(source.block_number))
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get block info return none, block_num: {:#?}",
                    source.block_number
                )))
            })?;
        if blob_block.hash != Some(source.block_hash) {
            return Err(ScalarError::Error(anyhow!(
                "blob source block hash changed before sidecar lookup"
            )));
        }

        let indexed_hashes = data_and_hashes_from_txs(&blob_block.transactions, &blob_tx);
        if indexed_hashes.is_empty() {
            return Err(ScalarError::Error(anyhow!(format!(
                "resolved blob source {:?} has no blob positions",
                source.transaction_hash
            ))));
        }
        let actual_hashes: Vec<H256> = indexed_hashes.iter().map(|item| item.hash).collect();
        if actual_hashes != source.blob_hashes {
            return Err(ScalarError::Error(anyhow!(
                "blob positions changed between source resolution and sidecar lookup"
            )));
        }

        // The sample is finalized, so the following block (whose parent beacon
        // root addresses this execution block's sidecars) must already exist.
        let next_block_num =
            U64::from(
                source.block_number.as_u64().checked_add(1).ok_or_else(|| {
                    ScalarError::Error(anyhow!("blob source block number overflow"))
                })?,
            );
        let next_block = self
            .l1_provider
            .get_block(BlockNumber::Number(next_block_num))
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("get next L1 block: {e:?}"))))?
            .ok_or_else(|| ScalarError::Error(anyhow!("next finalized L1 block is unavailable")))?;
        if next_block.parent_hash != source.block_hash {
            return Err(ScalarError::Error(anyhow!(
                "next L1 block is not the canonical child of the blob source block"
            )));
        }
        let prev_beacon_root = next_block.parent_beacon_block_root.ok_or_else(|| {
            ScalarError::Error(anyhow!(format!(
                "next block has no parent beacon root, block number: {next_block_num:?}"
            )))
        })?;

        let indexes: Vec<u64> = indexed_hashes.iter().map(|item| item.index).collect();
        let sidecars_rt = self
            .beacon_node
            .query_sidecars(hex::encode_prefixed(prev_beacon_root), indexes)
            .await?;

        let sidecars: &Vec<Value> = sidecars_rt["data"].as_array().ok_or_else(|| {
            ScalarError::Error(anyhow!(format!(
                "blob_sidecars is none, blk_num: {:?}, blk_root: {:?}",
                source.block_number, prev_beacon_root
            )))
        })?;

        if sidecars.is_empty() {
            return Err(ScalarError::Error(anyhow!(format!(
                "blob_sidecars is empty, blk_num: {:?}, blk_root: {:?}",
                source.block_number, prev_beacon_root
            ))));
        }

        // All blobs belong to the same batch: the batch is compressed as a whole,
        // split into multiple segments across blobs, then reconstructed by concatenating
        // the segments, trimming the valid compressed payload, and decompressing once.
        // This also remains compatible with the single-blob case.
        let num_blobs = indexed_hashes.len() as u64;
        let origin_batch = extract_tx_payload(indexed_hashes, sidecars)?;

        let batch_size = origin_batch.len() as u64;
        let txn_count = extract_txn_count(&origin_batch, last_block_num).ok_or_else(|| {
            ScalarError::CalculateError(anyhow!(
                "blob payload block range does not match canonical lastBlockNumber"
            ))
        })?;

        Ok((batch_size, num_blobs, txn_count))
    }
}

#[cfg(test)]
mod tests {

    use super::*;
    use std::{env::var, str::FromStr, sync::Arc};

    #[test]
    fn v2_blob_commitment_preserves_order() {
        let a = H256::from_low_u64_be(1);
        let b = H256::from_low_u64_be(2);
        assert_ne!(blob_commitment(2, &[a, b]).unwrap(), blob_commitment(2, &[b, a]).unwrap());
        assert!(blob_commitment(2, &[]).is_err());
        assert!(blob_commitment(1, &[a, b]).is_err());
    }

    #[test]
    fn fresh_blob_survives_mapping_gc_but_not_a_nonzero_mismatch() {
        let actual = H256::from_low_u64_be(7);
        assert!(validate_fresh_blob_commitment(H256::zero(), actual).is_ok());
        assert!(validate_fresh_blob_commitment(actual, actual).is_ok());
        assert!(validate_fresh_blob_commitment(H256::from_low_u64_be(8), actual).is_err());
    }

    #[test]
    fn mapping_gc_recovery_prefers_snapshot_then_canonical_block() {
        let snapshot = H256::from_low_u64_be(1);
        let historical = H256::from_low_u64_be(2);
        assert_eq!(recover_gc_commitment(snapshot, historical), snapshot);
        assert_eq!(recover_gc_commitment(H256::zero(), historical), historical);
        assert_eq!(recover_gc_commitment(H256::zero(), H256::zero()), H256::zero());
    }

    #[test]
    fn zero_blobless_classification_preserves_legacy_but_rejects_current_commit_batch() {
        assert_eq!(
            classify_blobless_submission(
                0,
                SubmissionKind::CommitBatch,
                true,
                ZERO_VERSIONED_HASH,
            )
            .unwrap(),
            BloblessDisposition::ExcludeZero
        );
        assert_eq!(
            classify_blobless_submission(
                1,
                SubmissionKind::CommitBatchWithProof,
                false,
                ZERO_VERSIONED_HASH,
            )
            .unwrap(),
            BloblessDisposition::ExcludeZero
        );
        assert_eq!(
            classify_blobless_submission(1, SubmissionKind::CommitState, false, H256::zero(),)
                .unwrap(),
            BloblessDisposition::ExcludeMissing
        );
        assert!(classify_blobless_submission(
            1,
            SubmissionKind::CommitBatch,
            false,
            ZERO_VERSIONED_HASH,
        )
        .is_err());
    }

    #[test]
    fn blob_source_dedup_uses_full_stable_identity() {
        let source = BlobSource {
            transaction_hash: H256::from_low_u64_be(1),
            block_number: 10.into(),
            block_hash: H256::from_low_u64_be(10),
            transaction_index: 0.into(),
            blob_hashes: vec![H256::from_low_u64_be(11)],
        };
        let mut sources = HashSet::new();
        assert!(sources.insert(source.clone()));
        assert!(!sources.insert(source.clone()));

        let mut orphan_replacement = source;
        orphan_replacement.block_hash = H256::from_low_u64_be(12);
        assert!(sources.insert(orphan_replacement));
    }

    #[test]
    fn source_receipt_must_match_full_stable_identity() {
        let source = BlobSource {
            transaction_hash: H256::from_low_u64_be(1),
            block_number: 10.into(),
            block_hash: H256::from_low_u64_be(10),
            transaction_index: 2.into(),
            blob_hashes: vec![H256::from_low_u64_be(11)],
        };
        let mut receipt = TransactionReceipt {
            transaction_hash: source.transaction_hash,
            block_number: Some(source.block_number),
            block_hash: Some(source.block_hash),
            transaction_index: source.transaction_index,
            ..Default::default()
        };
        assert!(receipt_matches_source(&receipt, &source));
        receipt.block_hash = Some(H256::from_low_u64_be(12));
        assert!(!receipt_matches_source(&receipt, &source));
    }

    #[test]
    fn commit_window_indices_are_unique_and_bounded() {
        let topic0 = H256::from_low_u64_be(1);
        let make_log = |index: u64| Log {
            topics: vec![topic0, H256::from_low_u64_be(index)],
            ..Default::default()
        };
        let removed = Log { removed: Some(true), ..make_log(2) };
        let indices = batch_indices_from_commit_logs(
            &[make_log(1), make_log(1), removed, make_log(3), make_log(9)],
            3,
        );
        assert_eq!(indices.into_iter().collect::<Vec<_>>(), vec![1, 3]);
    }

    #[test]
    fn historical_commit_replay_is_indexed_once_by_batch() {
        let make_log = |index: u64| Log {
            topics: vec![H256::zero(), H256::from_low_u64_be(index)],
            ..Default::default()
        };
        let removed = Log { removed: Some(true), ..make_log(2) };
        let index = index_commit_logs(vec![make_log(1), make_log(1), make_log(2), removed]);
        assert_eq!(index.get(&1).unwrap().len(), 2);
        assert_eq!(index.get(&2).unwrap().len(), 1);
    }

    #[test]
    fn changed_snapshot_hash_fails_closed() {
        let expected = H256::from_low_u64_be(1);
        assert!(ensure_snapshot_hash(expected, Some(expected)).is_ok());
        assert!(ensure_snapshot_hash(expected, Some(H256::from_low_u64_be(2))).is_err());
        assert!(ensure_snapshot_hash(expected, None).is_err());
    }

    #[tokio::test]
    #[ignore]
    async fn test_calculate_from_current_rollup() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let rollup_tx_hash = "0x87b09de64fd9c433226a0c683a3b3c1d1e8ab3fa24f3213fa63e2931f205f8d8"
            .parse::<H256>()
            .unwrap();
        let rollup_tx_block_num = U64::from(1489357);

        let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env empty");
        let l2_rpc = var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env empty");
        let gas_threshold = var("GAS_THRESHOLD").expect("GAS_THRESHOLD env empty").parse().unwrap();
        let l1_rollup_address =
            Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env empty")).unwrap();
        let l2_oracle_address =
            Address::from_str(&var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env empty"))
                .unwrap();
        let private_key =
            var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env empty");

        let l1_provider = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();
        let l1_rollup_contract = Rollup::new(l1_rollup_address, Arc::new(l1_provider.clone()));

        let l2_provider = Provider::<Http>::try_from(l2_rpc).unwrap();
        let l2_signer = Arc::new(SignerMiddleware::new(
            l2_provider.clone(),
            Wallet::from_str(private_key.as_str())
                .unwrap()
                .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
        ));

        let l2_oracle_contract = GasPriceOracle::new(l2_oracle_address, l2_signer);

        let ext_signer =
            SignerClient::new("appid", "privkey_pem", "address", "chain", "url").unwrap();
        let mut overhead: ScalarUpdater = ScalarUpdater::new(
            l1_provider,
            l2_provider,
            l2_oracle_contract,
            Some(ext_signer),
            l1_rollup_contract,
            var("GAS_ORACLE_L1_BEACON_RPC")
                .expect("Cannot detect GAS_ORACLE_L1_BEACON_RPC env empty")
                .parse()
                .expect("Cannot parse GAS_ORACLE_L1_BEACON_RPC env var empty"),
            1u64,
            gas_threshold,
            0u64,
            0u64,
            0u64,
            50u64,
        );

        let source_tx =
            overhead.l1_provider.get_transaction(rollup_tx_hash).await.unwrap().unwrap();
        let source = BlobSource {
            transaction_hash: rollup_tx_hash,
            block_number: rollup_tx_block_num,
            block_hash: source_tx.block_hash.unwrap(),
            transaction_index: source_tx.transaction_index.unwrap(),
            blob_hashes: transaction_blob_hashes(&source_tx),
        };
        let latest_overhead = overhead.calculate_from_rollup(source, 0).await;

        log::info!("latest_overhead: {:#?}", latest_overhead);
    }
}
