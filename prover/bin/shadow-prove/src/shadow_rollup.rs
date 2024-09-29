use crate::{metrics::METRICS, BatchInfo};
use alloy::{
    network::{Ethereum, Network, ReceiptResponse},
    primitives::{Address, Bytes, TxHash, U256, U64},
    providers::{Provider, ReqwestProvider, RootProvider},
    rpc::{client::RpcClient, types::Log},
    sol_types::SolCall,
    transports::{
        http::{Client, Http},
        Transport,
    },
};
use std::ops::Mul;

use crate::{
    Rollup::{self, RollupInstance},
    ShadowRollup::{self, ShadowRollupInstance},
};

#[derive(Clone, Debug)]
pub struct BatchSyncer<T, P, N> {
    l1_provider: RootProvider<Http<Client>>,
    l1_rollup: RollupInstance<T, P, N>,
    l1_shadow_rollup: ShadowRollupInstance<T, P, N>,
}

impl<T, P, N> BatchSyncer<T, P, N>
where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
    pub fn new(
        rollup_address: Address,
        shadow_rollup_address: Address,
        provider: RootProvider<Http<Client>>,
        wallet: P,
    ) -> Self {
        let l1_rollup = Rollup::RollupInstance::new(rollup_address, wallet.clone());
        let l1_shadow_rollup = ShadowRollup::new(shadow_rollup_address, wallet);

        Self { l1_provider: provider, l1_rollup, l1_shadow_rollup }
    }

    /**
     * Sync a latest batch to l1-shadow-rollup.
     */
    pub async fn sync_batch(&self) -> Result<Option<BatchInfo>, anyhow::Error> {
        log::info!("start sync_batch...");

        let latest = self.l1_provider.get_block_number().await?;

        let (batch_info, batch_header) = match get_committed_batch(
            U64::from(latest),
            &self.l1_rollup,
            &self.l1_provider,
        )
        .await
        {
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
            hex::encode(batch_store.prevStateRoot.as_slice()),
            hex::encode(batch_store.postStateRoot.as_slice()),
            hex::encode(batch_store.withdrawalRoot.as_slice()),
            hex::encode(batch_store.dataHash.as_slice()),
            hex::encode(batch_store.blobVersionedHash.as_slice()),
            hex::encode(batch_store.sequencerSetVerifyHash.as_slice()),
        );

        // Prepare shadow batch
        let shadow_tx = self.l1_shadow_rollup.commitBatch(batch_info.batch_index, batch_store);
        let rt = shadow_tx.send().await;

        let pending_tx = match rt {
            Ok(pending_tx) => pending_tx,
            Err(e) => {
                log::error!("send tx of shadow_rollup.commit_batch error: {:#?}", e);
                return Ok(None);
            }
        };
        let receipt = pending_tx.get_receipt().await.unwrap();
        if !receipt.status() {
            log::error!("shadow_rollup.commit_batch check_receipt fail");
            return Ok(None);
        }

        log::info!(">Sync shadow batch complete: {:#?}", batch_info.batch_index);
        Ok(Some(batch_info))
    }
}

/// Returns an HTTP provider for the given URL.
pub fn http_provider(url: &str) -> ReqwestProvider<Ethereum> {
    let url = url.parse().unwrap();
    let http = Http::<Client>::new(url);
    ReqwestProvider::new(RpcClient::new(http, true))
}

async fn get_committed_batch<T, P, N>(
    latest: U64,
    l1_rollup: &RollupInstance<T, P, N>,
    l1_provider: &RootProvider<Http<Client>>,
) -> Result<Option<(BatchInfo, Bytes)>, String>
where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
    log::info!("latest blocknum = {:#?}", latest);
    let start = if latest > U64::from(600) { latest - U64::from(600) } else { U64::from(1) };
    let filter =
        l1_rollup.CommitBatch_filter().filter.from_block(start).address(*l1_rollup.address());
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
            let _index = U256::from_be_slice(log.topics()[1].as_slice());
            let _tx_hash = log.transaction_hash.unwrap_or_default();
            (_index.to::<u64>(), _tx_hash)
        }
        None => {
            return Err("find commit_batch log error".to_string());
        }
    };

    if batch_index == 0 {
        return Err(String::from("batch_index == 0"));
    }
    let blocks = match batch_inspect(l1_provider, tx_hash).await {
        Some(blocks) => blocks,
        None => vec![],
    };

    if blocks.is_empty() {
        return Err(String::from("blocks.is_empty()"));
    }

    let batch_info: BatchInfo = BatchInfo { batch_index, blocks };

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
async fn batch_header_inspect(
    l1_provider: &RootProvider<Http<Client>>,
    hash: TxHash,
) -> Option<Bytes> {
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
    let data = tx.input;

    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = if let Ok(_param) = Rollup::commitBatchCall::abi_decode(&data, false) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
        return None;
    };
    let parent_batch_header: Bytes = param.batchDataInput.parentBatchHeader;
    Some(parent_batch_header)
}

async fn batch_inspect(l1_provider: &RootProvider<Http<Client>>, hash: TxHash) -> Option<Vec<u64>> {
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
    let data = tx.input;

    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = if let Ok(_param) = Rollup::commitBatchCall::abi_decode(&data, false) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", hash);
        return None;
    };
    let block_contexts: Bytes = param.batchDataInput.blockContexts;
    decode_blocks(block_contexts)
}

fn decode_blocks(block_contexts: Bytes) -> Option<Vec<u64>> {
    if block_contexts.is_empty() {
        return None;
    }

    let mut blocks: Vec<u64> = vec![];
    let mut txn_in_batch = 0u32;
    let bs: &[u8] = &block_contexts;

    // decode blocks from batch
    // |   1 byte   | 60 bytes | ... | 60 bytes |
    // | num blocks |  block 1 | ... |  block n |
    let num_blocks: u16 = ((bs[0] as u16) << 8) | (bs[1] as u16);

    for i in 0..num_blocks as usize {
        let block_num =
            u64::from_be_bytes(bs.get((60.mul(i) + 2)..(60.mul(i) + 2 + 8))?.try_into().unwrap());
        let txs_num = u16::from_be_bytes(
            bs.get((60.mul(i) + 2 + 56)..(60.mul(i) + 2 + 58))?.try_into().unwrap(),
        );
        txn_in_batch += txs_num as u32;
        blocks.push(block_num);
    }

    METRICS.shadow_txn_len.set(txn_in_batch.into());
    log::debug!("total_l2txn_in_batch: {:#?}", txn_in_batch);
    log::debug!("num_blocks: {:#?}, decode_blocks: {:#?}", num_blocks, blocks);
    Some(blocks)
}

async fn is_prove_success<T, P, N>(
    batch_index: u64,
    l1_rollup: &ShadowRollupInstance<T, P, N>,
) -> Option<bool>
where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
    let is_prove_success: bool =
        match l1_rollup.isProveSuccess(U256::from(batch_index)).call().await {
            Ok(x) => x._0,
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
async fn test_decode_blocks() {
    use std::{fs::File, io::Read, str::FromStr};

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    let mut file = File::open("./src/input.data").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    let input = Bytes::from_str(contents.as_str()).unwrap();

    let param = Rollup::commitBatchCall::abi_decode(&input, false).unwrap();
    let blocks: Bytes = param.batchDataInput.blockContexts;
    let rt = decode_blocks(blocks).unwrap();
    assert!(rt.len() == 11);
    assert!(rt[3] == 1112u64);
}

#[tokio::test]
async fn test_sync_batch() {
    use alloy::{
        network::EthereumWallet,
        primitives::Address,
        providers::{ProviderBuilder, RootProvider},
        signers::local::PrivateKeySigner,
        transports::http::{Client, Http},
    };
    use std::{env::var, str::FromStr};

    let l1_rpc: String = var("SHADOW_PROVING_VERIFY_L1_RPC").unwrap_or(
        var("SHADOW_PROVING_L1_RPC").expect("Shadow prove cannot detect L1_RPC env var"),
    );
    let private_key = var("SHADOW_PROVING_PRIVATE_KEY")
        .expect("Cannot detect SHADOW_PROVING_PRIVATE_KEY env var");

    let signer: PrivateKeySigner = private_key.parse().unwrap();
    let wallet: EthereumWallet = EthereumWallet::from(signer.clone());
    let provider: RootProvider<Http<Client>> =
        ProviderBuilder::new().on_http(l1_rpc.parse().unwrap());

    let rollup = var("SHADOW_PROVING_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let shadow_rollup =
        var("SHADOW_PROVING_L1_SHADOW_ROLLUP").expect("Cannot detect L1_SHADOW_ROLLUP env var");

    let l1_signer = ProviderBuilder::new()
        .with_recommended_fillers()
        .wallet(wallet)
        .on_http(l1_rpc.parse().unwrap());

    let bs = BatchSyncer::new(
        Address::from_str(&rollup).unwrap(),
        Address::from_str(&shadow_rollup).unwrap(),
        provider,
        l1_signer,
    );
    bs.sync_batch().await.unwrap();
}
