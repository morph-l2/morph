use crate::{metrics::METRICS, BatchInfo};
use alloy_consensus::Transaction;
use alloy_network::{Network, ReceiptResponse};
use alloy_primitives::{hex, Address, Bytes, TxHash, U256, U64};
use alloy_provider::{DynProvider, Provider};
use alloy_rpc_types::Log;
use alloy_sol_types::SolCall;

use anyhow::{anyhow, Context, Result};
use futures::future::join_all;

use crate::{
    Rollup::{self, RollupInstance},
    ShadowRollup::{self, ShadowRollupInstance},
};

#[derive(Clone, Debug)]
pub struct BatchSyncer<P, N> {
    l1_provider: DynProvider,
    l2_provider: DynProvider,
    l1_rollup: RollupInstance<DynProvider>,
    l1_shadow_rollup: ShadowRollupInstance<P, N>,
}

impl<P, N> BatchSyncer<P, N>
where
    P: Provider<N> + Clone,
    N: Network,
{
    pub fn new(
        rollup_address: Address,
        shadow_rollup_address: Address,
        l1_provider: DynProvider,
        l2_provider: DynProvider,
        wallet: P,
    ) -> Self {
        let l1_rollup = Rollup::RollupInstance::new(rollup_address, l1_provider.clone());
        let l1_shadow_rollup = ShadowRollup::new(shadow_rollup_address, wallet);

        Self { l1_provider, l2_provider, l1_rollup, l1_shadow_rollup }
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

        log::info!(
            "sync batch of {:?}, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
            dataHash = {:?}, blobVersionedHash = {:?}, sequencerSetVerifyHash = {:?}",
            batch_info.batch_index,
            hex::encode_prefixed(batch_store.prevStateRoot),
            hex::encode_prefixed(batch_store.postStateRoot),
            hex::encode_prefixed(batch_store.withdrawalRoot),
            hex::encode_prefixed(batch_store.dataHash),
            hex::encode_prefixed(batch_store.blobVersionedHash),
            hex::encode_prefixed(batch_store.sequencerSetVerifyHash),
        );

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

    pub async fn get_committed_batch(&self) -> Result<Option<(BatchInfo, Bytes)>> {
        let latest = match self.l1_provider.get_block_number().await {
            Ok(v) => U64::from(v),
            Err(e) => {
                log::error!("l1_provider.get_block_number error: {:?}", e);
                return Err(anyhow!("l1_provider.get_block_number error"));
            }
        };

        log::info!("latest l1 blocknum = {:#?}", latest);
        let start = if latest > U64::from(600) { latest - U64::from(600) } else { U64::from(1) };
        let filter = self
            .l1_rollup
            .CommitBatch_filter()
            .filter
            .from_block(start)
            .address(*self.l1_rollup.address());
        let mut logs: Vec<Log> = match self.l1_provider.get_logs(&filter).await {
            Ok(logs) => logs,
            Err(e) => {
                log::error!("l1_rollup.commit_batch.get_logs error: {:#?}", e);
                return Err(anyhow!("l1_rollup.commit_batch.get_logs provider error"));
            }
        };
        if logs.is_empty() {
            log::warn!("There have been no commit_batch logs for the last 600 blocks");
            return Ok(None);
        }
        if logs.len() < 3 {
            log::warn!("No enough commit_batch logs for the last 600 blocks");
            return Ok(None);
        }
        logs.sort_by(|a, b| a.block_number.unwrap().cmp(&b.block_number.unwrap()));

        let batch_index_hash = match logs.get(logs.len() - 2) {
            Some(log) => {
                let _index = U256::from_be_slice(log.topics()[1].as_slice());
                (_index.to::<u64>(), log.transaction_hash.unwrap_or_default())
            }
            None => {
                return Err(anyhow!("find commit_batch log error"));
            }
        };

        if batch_index_hash.0 == 0 {
            return Err(anyhow!("batch_index is 0"));
        }

        let prev_tx_hash = match logs.get(logs.len() - 3) {
            Some(log) => log.transaction_hash.unwrap_or_default(),
            None => {
                return Err(anyhow!("find commit_batch log error"));
            }
        };

        let (blocks, total_txn_count) =
            match self.batch_blocks_inspect(prev_tx_hash, batch_index_hash.1).await {
                Some(block_txn) => block_txn,
                None => return Err(anyhow!("batch_blocks_inspect none")),
            };

        if blocks.0 >= blocks.1 {
            return Err(anyhow!("blocks is empty"));
        }

        let batch_info: BatchInfo = BatchInfo {
            batch_index: batch_index_hash.0,
            start_block: blocks.0,
            end_block: blocks.1,
            total_txn: total_txn_count,
        };

        // A rollup commit_batch_input contains prev batch_header.
        let next_tx_hash = match logs.last() {
            Some(log) => log.transaction_hash.unwrap_or_default(),

            None => {
                return Err(anyhow!("find commit_batch log error"));
            }
        };
        let batch_input = batch_input_inspect(&self.l1_provider, next_tx_hash)
            .await
            .ok_or_else(|| anyhow!("Failed to inspect batch header"))?;

        log::info!("Found the committed batch, batch index = {:#?}", batch_index_hash.0);
        Ok(Some((batch_info, batch_input.0)))
    }

    async fn batch_blocks_inspect(
        &self,
        prev_batch_hash: TxHash,
        current_batch_hash: TxHash,
    ) -> Option<((u64, u64), u64)> {
        let prev_batch_input = batch_input_inspect(&self.l1_provider, prev_batch_hash).await?;
        let current_batch_input =
            batch_input_inspect(&self.l1_provider, current_batch_hash).await?;
        let start_block = prev_batch_input.1 + 1;
        let end_block = current_batch_input.1;

        if start_block == 0 {
            log::error!("batch_blocks_inspect: start_block = 0, tx_hash =  {:#?}", prev_batch_hash);
            return None;
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
                total_tx_count += res.unwrap_or_default().unwrap_or_default();
            }
        }

        log::info!(
            "decode_blocks, blocks_len: {:#?}, start_block: {:#?}, txn_in_batch: {:?}",
            end_block - start_block + 1,
            start_block,
            total_tx_count
        );

        METRICS.shadow_txn_len.set(total_tx_count as i64);

        Some(((start_block, end_block), total_tx_count))
    }

    // Check whether a batch has been proved successfully.
    async fn is_prove_success(&self, batch_index: u64) -> Result<bool, anyhow::Error> {
        self.l1_shadow_rollup
            .isProveSuccess(U256::from(batch_index))
            .call()
            .await
            .context("l1_shadow_rollup.is_prove_succes")
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
    let param = if let Ok(_param) = Rollup::commitBatchCall::abi_decode(data) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
        return None;
    };
    let parent_batch_header: Bytes = param.batchDataInput.parentBatchHeader;
    let last_block_number: u64 = param.batchDataInput.lastBlockNumber;
    Some((parent_batch_header, last_block_number))
}
#[tokio::test]
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

    let l1_signer = ProviderBuilder::new().wallet(wallet).connect_http(l1_rpc.parse().unwrap());

    let bs = BatchSyncer::new(
        Address::from_str(&rollup).unwrap(),
        Address::from_str(&shadow_rollup).unwrap(),
        l1_provider,
        l2_provider,
        l1_signer,
    );
    let batch = bs.get_committed_batch().await.unwrap().unwrap();
    bs.sync_batch(batch.0, batch.1).await.unwrap();
}

#[tokio::test]
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
