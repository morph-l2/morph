use crate::{metrics::METRICS, BatchInfo};
use alloy_consensus::Transaction;
use alloy_network::{Network, ReceiptResponse};
use alloy_primitives::{hex, Address, Bytes, Keccak256, TxHash, B256, U256};
use alloy_provider::{DynProvider, Provider};
use alloy_rpc_types::{BlockId, BlockNumberOrTag};

use anyhow::{anyhow, Context, Result};
use futures::future::join_all;

use crate::{
    rollup_compat::{
        decode_batch_data_input, resolve_canonical_commit, BatchDataInput, CanonicalCommit,
    },
    Rollup::{self, RollupInstance},
    ShadowRollup::{self, ShadowRollupInstance},
};

#[derive(Clone, Debug)]
pub struct BatchSyncer<P, N> {
    l1_provider: DynProvider,
    l2_provider: DynProvider,
    l1_rollup: RollupInstance<DynProvider>,
    l1_shadow_rollup: ShadowRollupInstance<P, N>,
    rollup_deployed_block: u64,
}

impl<P, N> BatchSyncer<P, N>
where
    P: Provider<N> + Clone,
    N: Network,
{
    async fn finalized_snapshot(&self) -> Result<(u64, B256)> {
        let block = self
            .l1_provider
            .get_block_by_number(BlockNumberOrTag::Finalized)
            .await?
            .ok_or_else(|| anyhow!("finalized L1 block is unavailable"))?;
        Ok((block.header.number, block.header.hash))
    }

    async fn assert_snapshot(&self, number: u64, expected_hash: B256) -> Result<()> {
        let block = self
            .l1_provider
            .get_block_by_number(BlockNumberOrTag::Number(number))
            .await?
            .ok_or_else(|| anyhow!("snapshot L1 block {number} is unavailable"))?;
        if block.header.hash != expected_hash {
            return Err(anyhow!(
                "finalized L1 snapshot changed: number={number}, expected={expected_hash:?}, actual={:?}",
                block.header.hash
            ));
        }
        Ok(())
    }

    pub fn new(
        rollup_address: Address,
        shadow_rollup_address: Address,
        l1_provider: DynProvider,
        l2_provider: DynProvider,
        wallet: P,
        rollup_deployed_block: u64,
    ) -> Self {
        let l1_rollup = Rollup::RollupInstance::new(rollup_address, l1_provider.clone());
        let l1_shadow_rollup = ShadowRollup::new(shadow_rollup_address, wallet);

        Self { l1_provider, l2_provider, l1_rollup, l1_shadow_rollup, rollup_deployed_block }
    }

    // Fetch the latest committed batch from l1-rollup if the batch_index has increased.
    pub async fn get_latest_batch(&self, batch_index: u64) -> Result<Option<(BatchInfo, Bytes)>> {
        let (snapshot, snapshot_hash) = self.finalized_snapshot().await?;
        let last_committed_batch_index = self
            .l1_rollup
            .lastCommittedBatchIndex()
            .block(BlockId::Number(BlockNumberOrTag::Number(snapshot)))
            .call()
            .await?
            .to::<u64>();
        self.assert_snapshot(snapshot, snapshot_hash).await?;
        if last_committed_batch_index <= batch_index.saturating_add(1) {
            log::info!(
                "The current batch_index has not increased, latest_batch_index: {last_committed_batch_index:?}"
            );
            return Ok(None);
        }
        if last_committed_batch_index <= 2 {
            return Ok(None);
        }
        self.get_specified_batch_at(last_committed_batch_index - 1, snapshot, snapshot_hash).await
    }

    // Fetch the latest committed batch from l1-rollup.
    pub async fn get_committed_batch(&self) -> Result<Option<(BatchInfo, Bytes)>> {
        let (snapshot, snapshot_hash) = self.finalized_snapshot().await?;
        let last_committed = self
            .l1_rollup
            .lastCommittedBatchIndex()
            .block(BlockId::Number(BlockNumberOrTag::Number(snapshot)))
            .call()
            .await?;
        self.assert_snapshot(snapshot, snapshot_hash).await?;
        let last_committed = last_committed.to::<u64>();
        if last_committed <= 2 {
            return Ok(None);
        }

        // The next canonical commit carries the target batch header as its
        // parentBatchHeader, so prove the newest batch that has a canonical
        // successor rather than relying on three adjacent logs.
        self.get_specified_batch_at(last_committed - 1, snapshot, snapshot_hash).await
    }

    /// Fetch a specified batch from l1-rollup by batch_num.
    pub async fn get_specified_batch(&self, batch_num: u64) -> Result<Option<(BatchInfo, Bytes)>> {
        let (snapshot, snapshot_hash) = self.finalized_snapshot().await?;
        self.get_specified_batch_at(batch_num, snapshot, snapshot_hash).await
    }

    async fn get_specified_batch_at(
        &self,
        batch_num: u64,
        snapshot: u64,
        snapshot_hash: B256,
    ) -> Result<Option<(BatchInfo, Bytes)>> {
        if batch_num <= 1 {
            // We need prev(batch_num-1) to infer start_block, and next(batch_num+1) to retrieve the
            // specified batch header (via parentBatchHeader).
            return Err(anyhow!("batch_num must be greater than 1"));
        }

        self.assert_snapshot(snapshot, snapshot_hash).await?;
        let prev = resolve_canonical_commit(
            &self.l1_provider,
            &self.l1_rollup,
            batch_num - 1,
            self.rollup_deployed_block,
            snapshot,
        )
        .await?;
        self.assert_snapshot(snapshot, snapshot_hash).await?;
        let current = resolve_canonical_commit(
            &self.l1_provider,
            &self.l1_rollup,
            batch_num,
            self.rollup_deployed_block,
            snapshot,
        )
        .await?;
        self.assert_snapshot(snapshot, snapshot_hash).await?;
        let next = resolve_canonical_commit(
            &self.l1_provider,
            &self.l1_rollup,
            batch_num + 1,
            self.rollup_deployed_block,
            snapshot,
        )
        .await?;
        self.assert_snapshot(snapshot, snapshot_hash).await?;
        let (Some(prev), Some(current), Some(next)) = (prev, current, next) else {
            return Ok(None);
        };

        let ((blocks, total_txn_count), current_input) =
            self.batch_blocks_inspect(&prev, &current).await?;

        let batch_info: BatchInfo = BatchInfo {
            batch_index: batch_num,
            start_block: blocks.0,
            end_block: blocks.1,
            total_txn: total_txn_count,
        };

        // next(batch_num+1) commitBatch calldata contains curr(batch_num) parentBatchHeader.
        let next_input =
            canonical_batch_input_inspect(&self.l1_provider, &next, *self.l1_rollup.address())
                .await?;
        let batch_header = next_input.parent_batch_header;
        validate_canonical_header(
            batch_header.as_ref(),
            batch_num,
            prev.batch_hash,
            &current_input,
        )?;
        let mut hasher = Keccak256::new();
        hasher.update(batch_header.as_ref());
        let header_hash = hasher.finalize();
        if header_hash != current.batch_hash {
            return Err(anyhow!(
                "canonical batch header hash mismatch: event={:?}, calldata={:?}",
                current.batch_hash,
                header_hash
            ));
        }
        self.assert_snapshot(snapshot, snapshot_hash).await?;

        log::info!(
            "Found canonical specified batch: batch_index={}, commit_identity={:?}",
            batch_num,
            current.identity
        );
        Ok(Some((batch_info, batch_header)))
    }

    /**
     * Sync a latest batch to l1-shadow-rollup.
     */
    pub async fn sync_batch(
        &self,
        batch_info: BatchInfo,
        batch_header: Bytes,
    ) -> Result<Option<BatchInfo>, anyhow::Error> {
        log::info!("start sync_batch...");
        // Batch should not have been verified yet.
        if self.is_prove_success(batch_info.batch_index).await? {
            log::info!("batch of {:?} already prove state successful", batch_info.batch_index);
            return Ok(None);
        };

        // Assembling a batche of the same commitment.
        #[rustfmt::skip]
        //   Below is the encoding for `BatchHeader`, reference: morph-repo/contracts/contracts/libraries/codec/BatchHeaderCodecV1.sol
        //    
        //   * Field                   Bytes       Type        Index   Comments
        //   * version                 1           uint8       0       The batch version
        //   * batchIndex              8           uint64      1       The index of the batch
        //   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch
        //   * totalL1MessagePopped    8           uint64      17      Number of total L1 messages popped after the batch
        //   * dataHash                32          bytes32     25      The data hash of the batch
        //   * blobVersionedHash       32          bytes32     57      The versioned hash of the blob with this batch’s data
        //   * prevStateHash           32          bytes32     89      Preview state root
        //   * postStateHash           32          bytes32     121     Post state root
        //   * withdrawRootHash        32          bytes32     153     L2 withdrawal tree root hash
        //   * sequencerSetVerifyHash  32          bytes32     185     L2 sequencers set verify hash
        //   * parentBatchHash         32          bytes32     217     The parent batch hash
        //   * skippedL1MessageBitmap  dynamic     uint256[]   249     A bitmap to indicate which L1 messages are skipped in the batch
        //   @dev Below is the feilds for `BatchHeader` V1
        //   * lastBlockNumber         8           uint64      249     The last block number in this batch
        // ```
        let batch_store = ShadowRollup::BatchStore {
            prevStateRoot: batch_header
                .get(89..121)
                .unwrap_or_default()
                .try_into()
                .unwrap_or_default(),
            postStateRoot: batch_header
                .get(121..153)
                .unwrap_or_default()
                .try_into()
                .unwrap_or_default(),
            withdrawalRoot: batch_header
                .get(153..185)
                .unwrap_or_default()
                .try_into()
                .unwrap_or_default(),
            dataHash: batch_header.get(25..57).unwrap_or_default().try_into().unwrap_or_default(),
            blobVersionedHash: batch_header
                .get(57..89)
                .unwrap_or_default()
                .try_into()
                .unwrap_or_default(),
            sequencerSetVerifyHash: batch_header
                .get(185..217)
                .unwrap_or_default()
                .try_into()
                .unwrap_or_default(),
        };

        // Commit the shadow batch.
        let shadow_tx = self.l1_shadow_rollup.commitBatch(batch_info.batch_index, batch_store);
        let rt = shadow_tx.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => pending_tx,
            Err(e) => {
                log::error!("send tx of shadow_rollup.commit_batch error: {:#?}", e);
                return Ok(None);
            }
        };
        let receipt =
            pending_tx.get_receipt().await.map_err(|e| anyhow!("get receipt error: {e}"))?;
        if !receipt.status() {
            log::error!("shadow_rollup.commit_batch check_receipt fail");
            return Ok(None);
        }

        log::info!(">Sync shadow batch complete: {:#?}", batch_info.batch_index);
        Ok(Some(batch_info))
    }

    // Inspect blocks range and total txn count in a batch.
    async fn batch_blocks_inspect(
        &self,
        prev: &CanonicalCommit,
        current: &CanonicalCommit,
    ) -> Result<(((u64, u64), u64), BatchDataInput)> {
        let rollup_address = *self.l1_rollup.address();
        let prev_batch_input =
            canonical_batch_input_inspect(&self.l1_provider, prev, rollup_address).await?;
        let current_batch_input =
            canonical_batch_input_inspect(&self.l1_provider, current, rollup_address).await?;
        let start_block = prev_batch_input
            .last_block_number
            .checked_add(1)
            .ok_or_else(|| anyhow!("previous batch last block overflows"))?;
        let end_block = current_batch_input.last_block_number;

        if start_block == 0 {
            return Err(anyhow!(
                "batch_blocks_inspect: start_block = 0, tx_hash = {:?}",
                prev.transaction_hash
            ));
        }
        if start_block > end_block {
            return Err(anyhow!(
                "batch_blocks_inspect: invalid block range {start_block}..={end_block}"
            ));
        }

        let mut total_tx_count: u64 = 0;
        let block_numbers: Vec<u64> = (start_block..=end_block).collect();
        for chunk in block_numbers.chunks(10) {
            let mut tasks = Vec::with_capacity(chunk.len());
            for &i in chunk {
                let provider = self.l2_provider.clone();
                tasks.push(async move {
                    provider.get_block_transaction_count_by_number(i.into()).await
                });
            }
            let results = join_all(tasks).await;
            for res in results {
                total_tx_count += res
                    .context("query L2 block transaction count")?
                    .ok_or_else(|| anyhow!("L2 block transaction count is unavailable"))?;
            }
        }

        log::info!(
            "decode_blocks, blocks_len: {:#?}, start_block: {:#?}, txn_in_batch: {:?}",
            end_block - start_block + 1,
            start_block,
            total_tx_count
        );

        METRICS.shadow_txn_len.set(total_tx_count as i64);

        Ok((((start_block, end_block), total_tx_count), current_batch_input))
    }

    // Check whether a batch has been proved successfully.
    async fn is_prove_success(&self, batch_index: u64) -> Result<bool, anyhow::Error> {
        self.l1_shadow_rollup
            .isProveSuccess(U256::from(batch_index))
            .call()
            .await
            .context("l1_shadow_rollup.is_prove_succes")
    }

    // Calculate the batch public input hash.
    pub fn calc_batch_pi(
        &self,
        chain_id: u64,
        batch_header: &Bytes,
    ) -> Result<B256, anyhow::Error> {
        // `version` is read only for logging here. Unlike the V2/non-V2 split in `lib.rs`
        // (where the prover dispatches different blob-decoding paths), the public-input
        // formula is uniform across versions: offset 57 holds the blob input field
        // (versioned hash for V0/V1, aggregated hash for V2), and the on-chain verifier
        // uses the same layout. Versioned routing happens earlier via the `batch_version`
        // request parameter; nothing in this function needs to branch on it.
        let version = batch_header.first().copied().unwrap_or(0);

        let prev_state_root: &[u8] = batch_header.get(89..121).unwrap_or_default();
        let post_state_root: &[u8] = batch_header.get(121..153).unwrap_or_default();
        let withdrawal_root: &[u8] = batch_header.get(153..185).unwrap_or_default();
        let data_hash: &[u8] = batch_header.get(25..57).unwrap_or_default();
        let sequencer_set_verify_hash: &[u8] = batch_header.get(185..217).unwrap_or_default();

        // All versions: blob input at offset 57 (aggregated hash for V2, versioned hash for V0/V1)
        let blob_input: &[u8] = batch_header.get(57..89).unwrap_or_default();

        log::info!(
            "calc_batch_pi, version = {}, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
            dataHash = {:?}, blobInput = {:?}, sequencerSetVerifyHash = {:?}",
            version,
            hex::encode_prefixed(prev_state_root),
            hex::encode_prefixed(post_state_root),
            hex::encode_prefixed(withdrawal_root),
            hex::encode_prefixed(data_hash),
            hex::encode_prefixed(blob_input),
            hex::encode_prefixed(sequencer_set_verify_hash),
        );
        let mut hasher = Keccak256::new();
        hasher.update(chain_id.to_be_bytes());
        hasher.update(prev_state_root);
        hasher.update(post_state_root);
        hasher.update(withdrawal_root);
        hasher.update(sequencer_set_verify_hash);
        hasher.update(data_hash);
        hasher.update(blob_input);
        Ok(hasher.finalize())
    }
}

pub async fn batch_input_inspect(l1_provider: &DynProvider, hash: TxHash) -> Option<(Bytes, u64)> {
    //Step1.  Get transaction
    let result = l1_provider.get_transaction_by_hash(hash).await;
    let tx = match result {
        Ok(Some(tx)) => tx,
        Ok(None) => {
            log::error!("l1_provider.get_transaction is none");
            return None;
        }
        Err(e) => {
            log::error!("l1_provider.get_transaction err: {:#?}", e);
            return None;
        }
    };

    //Step2. Parse transaction data
    let data = tx.input();

    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = match decode_batch_data_input(data) {
        Ok(param) => param,
        Err(err) => {
            log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
            log::debug!("batch inspect decode error: {err:#}");
            return None;
        }
    };
    let parent_batch_header = param.parent_batch_header;
    let last_block_number = param.last_block_number;
    Some((parent_batch_header, last_block_number))
}

async fn canonical_batch_input_inspect(
    l1_provider: &DynProvider,
    commit: &CanonicalCommit,
    rollup_address: Address,
) -> Result<BatchDataInput> {
    let tx = l1_provider
        .get_transaction_by_hash(commit.transaction_hash)
        .await?
        .ok_or_else(|| anyhow!("canonical commit transaction is unavailable"))?;
    if tx.to() != Some(rollup_address) ||
        tx.block_number != Some(commit.identity.block_number) ||
        tx.transaction_index != Some(commit.identity.transaction_index) ||
        tx.block_hash != Some(commit.block_hash)
    {
        return Err(anyhow!(
            "canonical commit transaction identity mismatch: expected={:?}, tx_block={:?}, tx_index={:?}, tx_block_hash={:?}",
            commit.identity,
            tx.block_number,
            tx.transaction_index,
            tx.block_hash
        ));
    }
    decode_batch_data_input(tx.input())
}

fn validate_canonical_header(
    header: &[u8],
    expected_batch_index: u64,
    expected_parent_hash: B256,
    input: &BatchDataInput,
) -> Result<()> {
    let minimum_len = if input.version == 0 { 249 } else { 257 };
    if header.len() < minimum_len {
        return Err(anyhow!(
            "canonical batch header is too short: version={}, len={}",
            input.version,
            header.len()
        ));
    }
    if header[0] != input.version {
        return Err(anyhow!("canonical batch header version mismatch"));
    }
    if u64::from_be_bytes(header[1..9].try_into().unwrap()) != expected_batch_index {
        return Err(anyhow!("canonical batch header index mismatch"));
    }
    if header[89..121] != input.prev_state_root ||
        header[121..153] != input.post_state_root ||
        header[153..185] != input.withdrawal_root
    {
        return Err(anyhow!("canonical batch header root mismatch"));
    }
    if header[217..249] != expected_parent_hash[..] {
        return Err(anyhow!("canonical batch header parent hash mismatch"));
    }
    if input.version > 0 &&
        u64::from_be_bytes(header[249..257].try_into().unwrap()) != input.last_block_number
    {
        return Err(anyhow!("canonical batch header last block mismatch"));
    }
    Ok(())
}

#[cfg(test)]
mod canonical_header_tests {
    use super::*;
    use crate::rollup_compat::SubmissionKind;

    fn fixture() -> (Vec<u8>, BatchDataInput, B256) {
        let parent_hash = B256::repeat_byte(0x44);
        let input = BatchDataInput {
            version: 2,
            parent_batch_header: Bytes::new(),
            last_block_number: 123,
            num_l1_messages: 0,
            prev_state_root: [0x11; 32],
            post_state_root: [0x22; 32],
            withdrawal_root: [0x33; 32],
            kind: SubmissionKind::CommitBatch,
            pre_submitter: false,
        };
        let mut header = vec![0u8; 257];
        header[0] = input.version;
        header[1..9].copy_from_slice(&7u64.to_be_bytes());
        header[89..121].copy_from_slice(&input.prev_state_root);
        header[121..153].copy_from_slice(&input.post_state_root);
        header[153..185].copy_from_slice(&input.withdrawal_root);
        header[217..249].copy_from_slice(parent_hash.as_slice());
        header[249..257].copy_from_slice(&input.last_block_number.to_be_bytes());
        (header, input, parent_hash)
    }

    #[test]
    fn canonical_header_matches_calldata_roots_parent_and_range() {
        let (header, input, parent_hash) = fixture();
        validate_canonical_header(&header, 7, parent_hash, &input).unwrap();
    }

    #[test]
    fn canonical_header_mismatch_fails_closed() {
        let (mut header, input, parent_hash) = fixture();
        header[121] ^= 1;
        assert!(validate_canonical_header(&header, 7, parent_hash, &input).is_err());

        let (mut header, input, parent_hash) = fixture();
        header[256] ^= 1;
        assert!(validate_canonical_header(&header, 7, parent_hash, &input).is_err());
    }
}
#[tokio::test]
#[ignore = "requires configured L1/L2 RPCs and a funded signer"]
async fn test_sync_batch() {
    use alloy_network::EthereumWallet;
    use alloy_primitives::Address;
    use alloy_provider::ProviderBuilder;
    use alloy_signer_local::PrivateKeySigner;
    use std::{env::var, str::FromStr};

    let l1_rpc: String = var("SHADOW_PROVING_VERIFY_L1_RPC").unwrap_or(
        var("SHADOW_PROVING_L1_RPC").expect("Shadow prove cannot detect L1_RPC env var"),
    );
    let l2_rpc: String = var("SHADOW_PROVING_VERIFY_L2_RPC").unwrap_or(
        var("SHADOW_PROVING_L2_RPC").expect("Shadow prove cannot detect L2_RPC env var"),
    );
    let private_key = var("SHADOW_PROVING_PRIVATE_KEY")
        .expect("Cannot detect SHADOW_PROVING_PRIVATE_KEY env var");

    let signer: PrivateKeySigner = private_key.parse().unwrap();
    let wallet: EthereumWallet = EthereumWallet::from(signer.clone());
    let l1_provider: DynProvider =
        ProviderBuilder::new().connect_http(l1_rpc.parse().unwrap()).erased();
    let l2_provider: DynProvider =
        ProviderBuilder::new().connect_http(l2_rpc.parse().unwrap()).erased();

    let rollup = var("SHADOW_PROVING_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let shadow_rollup =
        var("SHADOW_PROVING_L1_SHADOW_ROLLUP").expect("Cannot detect L1_SHADOW_ROLLUP env var");
    let rollup_deployed_block = var("SHADOW_PROVING_L1_ROLLUP_DEPLOY_BLOCK")
        .unwrap_or_else(|_| "1".to_string())
        .parse()
        .unwrap();

    let l1_signer = ProviderBuilder::new().wallet(wallet).connect_http(l1_rpc.parse().unwrap());

    let bs = BatchSyncer::new(
        Address::from_str(&rollup).unwrap(),
        Address::from_str(&shadow_rollup).unwrap(),
        l1_provider,
        l2_provider,
        l1_signer,
        rollup_deployed_block,
    );
    let batch = bs.get_committed_batch().await.unwrap().unwrap();
    bs.sync_batch(batch.0, batch.1).await.unwrap();
}

#[tokio::test]
#[ignore = "requires an authenticated historical L1 RPC fixture"]
async fn test_inspect_batch_header() {
    use alloy_primitives::B256;
    use alloy_provider::ProviderBuilder;
    use std::str::FromStr;

    let provider: DynProvider = ProviderBuilder::new()
        .connect_http(
            "https://eth-holesky.g.alchemy.com/v2/xxxxxxx".parse().expect("parse l1_rpc to Url"),
        )
        .erased();
    let next_tx_hash =
        B256::from_str("0x2bdfb2bd0b8c9210bfb593cc5734e3f092fcdd54fe74c46a938448b0422089f7")
            .unwrap();
    let batch_header = batch_input_inspect(&provider, next_tx_hash)
        .await
        .ok_or_else(|| "Failed to inspect batch header".to_string())
        .unwrap()
        .0;

    let batch_store = ShadowRollup::BatchStore {
        prevStateRoot: batch_header.get(89..121).unwrap_or_default().try_into().unwrap_or_default(),
        postStateRoot: batch_header
            .get(121..153)
            .unwrap_or_default()
            .try_into()
            .unwrap_or_default(),
        withdrawalRoot: batch_header
            .get(153..185)
            .unwrap_or_default()
            .try_into()
            .unwrap_or_default(),
        dataHash: batch_header.get(25..57).unwrap_or_default().try_into().unwrap_or_default(),
        blobVersionedHash: batch_header
            .get(57..89)
            .unwrap_or_default()
            .try_into()
            .unwrap_or_default(),
        sequencerSetVerifyHash: batch_header
            .get(185..217)
            .unwrap_or_default()
            .try_into()
            .unwrap_or_default(),
    };

    println!(
        "sync batch of {:?}, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
            dataHash = {:?}, blobVersionedHash = {:?}, sequencerSetVerifyHash = {:?}",
        "batch_info.batch_index",
        hex::encode(batch_store.prevStateRoot.as_slice()),
        hex::encode(batch_store.postStateRoot.as_slice()),
        hex::encode(batch_store.withdrawalRoot.as_slice()),
        hex::encode(batch_store.dataHash.as_slice()),
        hex::encode(batch_store.blobVersionedHash.as_slice()),
        hex::encode(batch_store.sequencerSetVerifyHash.as_slice()),
    );
    // prevStateRoot =
    // "13a862a764f09e1300ad485fadcc130741d400e8d5be3dbb968901e6590e25ca", postStateRoot =
    // "20a6aa14638839f76d2b233499439e45cd315434f9628902793c421ec71fcb0c", withdrawalRoot =
    // "eda0cccc67b86712eea4536d186be3d412b86c4c56741d641d1bbfdd26b5d40b",         dataHash =
    // "89a1c4692d97c7a4a516b35bc46963da3425af5273cb5a7b8ee2cdcf41c6fa65", blobVersionedHash =
    // "013f8fabf23fba03c52572d3403d175d952937cdd78bb8e9e06eb6ffa751fd2a", sequencerSetVerifyHash =
    // "60f10881edf25485d6d9db1c3a634c002bf4da64cce0f9a0f528e00f1ead3dec"
}
