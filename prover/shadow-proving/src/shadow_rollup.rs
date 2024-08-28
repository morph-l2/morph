use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::abi::shadow_rollup_abi::{BatchStore, ShadowRollup};
use crate::metrics::METRICS;
use crate::{BatchInfo, RollupType, ShadowRollupType};
use ethers::signers::Wallet;
use ethers::{abi::AbiDecode, prelude::*};
use std::env::var;
use std::ops::Mul;
use std::str::FromStr;
use std::sync::Arc;
use std::time::Duration;
use tokio::time::sleep;

#[derive(Clone, Debug)]
pub struct BatchSyncer {
    l1_provider: Provider<Http>,
    l1_rollup: RollupType,
    l1_shadow_rollup: ShadowRollupType,
}

impl BatchSyncer {
    pub async fn prepare() -> Self {
        let l1_rpc = var("SHADOW_PROVING_L1_RPC").expect("Cannot detect L1_RPC env var");
        let l1_rollup_address = var("SHADOW_PROVING_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
        let l1_shadow_rollup_address = var("SHADOW_PROVING_L1_SHADOW_ROLLUP").expect("Cannot detect L1_SHADOW_ROLLUP env var");
        let private_key = var("SHADOW_PROVING_PRIVATE_KEY").expect("Cannot detect SHADOW_PROVING_PRIVATE_KEY env var");

        // Provider & Signer
        let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc).unwrap();
        let l1_signer = Arc::new(SignerMiddleware::new(
            l1_provider.clone(),
            Wallet::from_str(private_key.as_str())
                .unwrap()
                .with_chain_id(l1_provider.get_chainid().await.unwrap().as_u64()),
        ));
        let l1_rollup: RollupType = Rollup::new(Address::from_str(l1_rollup_address.as_str()).unwrap(), l1_signer.clone());
        let l1_shadow_rollup: ShadowRollupType = ShadowRollup::new(Address::from_str(l1_shadow_rollup_address.as_str()).unwrap(), l1_signer.clone());

        Self {
            l1_provider,
            l1_rollup,
            l1_shadow_rollup,
        }
    }
    /**
     * Sync a latest batch to l1-shadow-rollup.
     */
    pub async fn sync(&self) -> Result<Option<BatchInfo>, anyhow::Error> {
        log::info!("start sync_batch...");
        self.sync_batch().await
    }

    async fn sync_batch(&self) -> Result<Option<BatchInfo>, anyhow::Error> {
        let latest = self.l1_provider.get_block_number().await?;

        let (batch_info, batch_header) = match get_committed_batch(latest, &self.l1_rollup, &self.l1_provider).await {
            Ok(Some(committed_batch)) => committed_batch,
            Ok(None) => return Ok(None),
            Err(msg) => {
                log::error!("get_committed_batch error: {:?}", msg);
                return Ok(None);
            }
        };

        if is_prove_success(batch_info.batch_index, &self.l1_shadow_rollup).await.unwrap_or(true) {
            log::debug!("batch of {:?} already prove successful", batch_info.batch_index);
            return Ok(None);
        };

        // Prepare shadow batch
        let shadow_tx = self.l1_shadow_rollup.commit_batch(
            batch_info.batch_index,
            BatchStore {
                prev_state_root: batch_header.get(89..121).unwrap_or_default().try_into().unwrap_or_default(),
                post_state_root: batch_header.get(121..153).unwrap_or_default().try_into().unwrap_or_default(),
                withdrawal_root: batch_header.get(153..185).unwrap_or_default().try_into().unwrap_or_default(),
                data_hash: batch_header.get(25..57).unwrap_or_default().try_into().unwrap_or_default(),
                blob_versioned_hash: batch_header.get(57..89).unwrap_or_default().try_into().unwrap_or_default(),
                sequencer_set_verify_hash: batch_header.get(185..217).unwrap_or_default().try_into().unwrap_or_default(),
            },
        );
        let rt = shadow_tx.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => pending_tx,
            Err(e) => {
                log::error!("send tx of shadow_rollup.commit_batch error: {:#?}", e);
                if let ContractError::Revert(data) = e {
                    let msg = String::decode_with_selector(&data).unwrap_or(String::from("unknown, decode contract revert error"));
                    log::error!("send tx of shadow_rollup.commit_batch error, exec msg: {:#?}", msg);
                }
                return Ok(None);
            }
        };
        if !check_receipt("shadow_commit_batch", &self.l1_provider, pending_tx).await {
            return Ok(None);
        }
        log::info!(">Sync shadow batch complete: {:#?}", batch_info.batch_index);
        Ok(Some(batch_info))
    }
}

async fn get_committed_batch(latest: U64, l1_rollup: &RollupType, l1_provider: &Provider<Http>) -> Result<Option<(BatchInfo, Bytes)>, String> {
    log::info!("latest blocknum = {:#?}", latest);
    let start = if latest > U64::from(600) {
        latest - U64::from(600)
    } else {
        U64::from(1)
    };
    let filter = l1_rollup.commit_batch_filter().filter.from_block(start).address(l1_rollup.address());
    let mut logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("l1_rollup.commit_batch.get_logs error: {:#?}", e);
            return Err("l1_rollup.commit_batch.get_logs provider error".to_string());
        }
    };
    if logs.is_empty() {
        log::warn!("There have been no commit_batch logs for the last 600 blocks");
        return Ok(None);
    }
    if logs.len() < 2 {
        log::warn!("No enough commit_batch logs for the last 600 blocks");
        return Ok(None);
    }
    logs.sort_by(|a, b| a.block_number.unwrap().cmp(&b.block_number.unwrap()));

    let batch_header = {
        let tx_hash = match logs.last() {
            Some(log) => log.transaction_hash.unwrap_or_default(),

            None => {
                return Err("find commit_batch log error".to_string());
            }
        };
        batch_header_inspect(l1_provider, tx_hash)
            .await
            .ok_or_else(|| "Failed to inspect batch header".to_string())?
    };

    let (batch_index, tx_hash) = match logs.get(logs.len() - 2) {
        Some(log) => {
            let _index = log.topics[1].to_low_u64_be();
            let _tx_hash = log.transaction_hash.unwrap_or_default();
            (_index, _tx_hash)
        }
        None => {
            return Err("find commit_batch log error".to_string());
        }
    };

    let chunks = match batch_inspect(l1_provider, tx_hash).await {
        Some(chunks) => chunks,
        None => vec![],
    };

    if batch_index == 0 || chunks.is_empty() {
        return Err(String::from("batch_index == 0 or chunks.is_empty()"));
    }

    let batch_info: BatchInfo = BatchInfo { batch_index, chunks };

    log::info!("latest batch index = {:#?}", batch_index);
    Ok(Some((batch_info, batch_header)))
}

/// Below is the encoding for `BatchHeader`.
/// ```text
///   * Field                   Bytes       Type        Index   Comments
///   * version                 1           uint8       0       The batch version
///   * batchIndex              8           uint64      1       The index of the batch
///   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch
///   * totalL1MessagePopped    8           uint64      17      Number of total L1 messages popped after the batch
///   * dataHash                32          bytes32     25      The data hash of the batch
///   * blobVersionedHash       32          bytes32     57      The versioned hash of the blob with this batch’s data
///   * parentBatchHash         32          bytes32     89      The parent batch hash
///   * skippedL1MessageBitmap  dynamic     uint256[]   121     A bitmap to indicate which L1 messages are skipped in the batch
/// ``
async fn batch_header_inspect(l1_provider: &Provider<Http>, hash: TxHash) -> Option<Bytes> {
    //Step1.  Get transaction
    let result = l1_provider.get_transaction(hash).await;
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
    let data = tx.input;

    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = if let Ok(_param) = CommitBatchCall::decode(&data) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
        return None;
    };
    let parent_batch_header: Bytes = param.batch_data_input.parent_batch_header;
    Some(parent_batch_header)
}

async fn batch_inspect(l1_provider: &Provider<Http>, hash: TxHash) -> Option<Vec<Vec<u64>>> {
    //Step1.  Get transaction
    let result = l1_provider.get_transaction(hash).await;
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
    let data = tx.input;

    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = if let Ok(_param) = CommitBatchCall::decode(&data) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
        return None;
    };
    let chunks: Vec<Bytes> = param.batch_data_input.chunks;
    decode_chunks(chunks)
}

fn decode_chunks(chunks: Vec<Bytes>) -> Option<Vec<Vec<u64>>> {
    if chunks.is_empty() {
        return None;
    }

    let mut chunk_with_blocks: Vec<Vec<u64>> = vec![];
    let mut txn_in_batch = 0;
    let mut max_txn_in_chunk = 0;
    for chunk in chunks.iter() {
        let mut chunk_bn: Vec<u64> = vec![];
        let bs: &[u8] = chunk;

        // decode blocks from chunk
        // |   1 byte   | 60 bytes | ... | 60 bytes |
        // | num blocks |  block 1 | ... |  block n |
        let num_blocks = U256::from_big_endian(bs.get(..1)?);
        for i in 0..num_blocks.as_usize() {
            let block_num = U256::from_big_endian(bs.get((60.mul(i) + 1)..(60.mul(i) + 1 + 8))?);
            let txs_num = U256::from_big_endian(bs.get((60.mul(i) + 1 + 56)..(60.mul(i) + 1 + 58))?);
            max_txn_in_chunk = max_txn_in_chunk.max(txs_num.as_u32());
            txn_in_batch += txs_num.as_u32();
            chunk_bn.push(block_num.as_u64());
        }

        chunk_with_blocks.push(chunk_bn);
    }
    METRICS.shadow_txn_len.set(txn_in_batch.into());
    log::debug!("decode_chunks_blocknum: {:#?}", chunk_with_blocks);
    log::debug!("max_l2txn_in_chunk: {:#?}", max_txn_in_chunk);
    log::debug!("total_l2txn_in_batch: {:#?}", txn_in_batch);
    Some(chunk_with_blocks)
}

async fn check_receipt(method: &str, l1_provider: &Provider<Http>, pending_tx: PendingTransaction<'_, Http>) -> bool {
    let check = || async {
        let receipt = l1_provider.get_transaction_receipt(pending_tx.tx_hash()).await.unwrap();
        match receipt {
            Some(tr) => {
                // Either 1 (success) or 0 (failure).
                match tr.status.unwrap_or_default().as_u64() {
                    1 => log::info!("{:?} receipt success: {:#?}", method, pending_tx.tx_hash()),
                    _ => log::error!("{:?} receipt fail: {:#?}", method, tr),
                };
                true
            }
            // Maybe still pending.
            None => {
                log::info!("{:?} receipt pending", method);
                false
            }
        }
    };

    for _ in 1..5 {
        sleep(Duration::from_secs(12)).await;
        if check().await {
            return true;
        };
    }
    false
}

async fn is_prove_success(batch_index: u64, l1_rollup: &ShadowRollupType) -> Option<bool> {
    let is_prove_success: bool = match l1_rollup.is_prove_success(U256::from(batch_index)).await {
        Ok(x) => x,
        Err(e) => {
            log::info!(
                "query l1_shadow_rollup.is_prove_success error, batch index = {:#?}, {:#?}",
                batch_index,
                e
            );
            return None;
        }
    };
    Some(is_prove_success)
}

#[tokio::test]
async fn test_decode_chunks() {
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    use std::fs::File;
    use std::io::Read;
    let mut file = File::open("./src/input.data").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    let input = Bytes::from_str(contents.as_str()).unwrap();

    let param = CommitBatchCall::decode(&input).unwrap();
    let chunks: Vec<Bytes> = param.batch_data_input.chunks;
    let rt = decode_chunks(chunks).unwrap();
    assert!(rt.len() == 11);
    assert!(rt.get(3).unwrap().len() == 2);
}
