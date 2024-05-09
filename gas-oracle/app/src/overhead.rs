use crate::abi::gas_price_oracle_abi::GasPriceOracle;
use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::blob::{kzg_to_versioned_hash, Blob, PREV_ROLLUP_L1_BLOCK};
use crate::blob_client::{BeaconNode, ExecutionNode};
use crate::metrics::ORACLE_SERVICE_METRICS;
use crate::typed_tx::TypedTransaction;
use ethers::utils::{hex, rlp};
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use serde_json::Value;
use std::env::var;
use std::ops::Mul;
use std::str::FromStr;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb
const MAX_OVERHEAD: usize = 10000; // 937,500

pub struct OverHead {
    l1_provider: Provider<Http>,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
    l1_rollup: Rollup<Provider<Http>>,
    overhead_threshold: u128,
    execution_node: ExecutionNode,
    beacon_node: BeaconNode,
    overhead_switch: bool,
}

impl OverHead {
    pub fn new(
        l1_provider: Provider<Http>,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        l1_rollup: Rollup<Provider<Http>>,
        overhead_threshold: u128,
        l1_rpc: String,
        l1_beacon_rpc: String,
        overhead_switch: bool,
    ) -> Self {
        let execution_node = ExecutionNode { rpc_url: l1_rpc };
        let beacon_node = BeaconNode {
            rpc_url: l1_beacon_rpc,
        };

        OverHead {
            l1_provider,
            l2_oracle,
            l1_rollup,
            overhead_threshold,
            execution_node,
            beacon_node,
            overhead_switch,
        }
    }

    /// Update overhead
    /// Calculate the average cost of the latest roll up and set it to the GasPriceOrale contract on the L2 network.
    pub async fn update(&self) {
        // Step1. fetch latest batches and calculate overhead.
        let latest = match self.l1_provider.get_block_number().await {
            Ok(bn) => bn,
            Err(e) => {
                log::error!("overhead.l1_provider.get_block_number error: {:#?}", e);
                return;
            }
        };
        let start = if latest > U64::from(100) {
            latest - U64::from(100) //100
        } else {
            latest
        };
        let filter = self
            .l1_rollup
            .commit_batch_filter()
            .filter
            .from_block(start)
            .address(self.l1_rollup.address());

        let mut logs: Vec<Log> = match self.l1_provider.get_logs(&filter).await {
            Ok(logs) => logs,
            Err(e) => {
                log::error!("overhead.l1_provider.get_logs error: {:#?}", e);
                return;
            }
        };
        log::debug!(
            "overhead.l1_provider.submit_batches.get_logs.len ={:#?}",
            logs.len()
        );

        logs.retain(|x| x.transaction_hash != None && x.block_number != None);
        if logs.is_empty() {
            log::warn!("rollup logs for the last 100 blocks of l1 is empty");
            return;
        }
        logs.sort_by(|a, b| b.block_number.unwrap().cmp(&a.block_number.unwrap()));
        let log = match logs.first() {
            Some(log) => log,
            None => {
                log::info!("no submit batches logs, latest blocknum ={:#?}", latest);
                return;
            }
        };

        let current_rollup_l1_block = log.block_number.unwrap_or(U64::from(0)).as_u64();
        if current_rollup_l1_block <= *PREV_ROLLUP_L1_BLOCK.lock().await {
            log::info!(
                "No new batch has been committed, latest blocknum ={:#?}",
                latest
            );
            return;
        }

        let mut latest_overhead = match self
            .overhead_inspect(log.transaction_hash.unwrap(), log.block_number.unwrap())
            .await
        {
            Some(overhead) => overhead,
            None => {
                log::info!(
                    "last_overhead is none, skip update, tx_hash ={:#?}",
                    log.transaction_hash.unwrap()
                );
                return;
            }
        };

        // Step2. fetch current overhead on l2.
        let current_overhead = match self.l2_oracle.overhead().await {
            Ok(ov) => ov,
            Err(e) => {
                log::error!("query l2_oracle.overhead error: {:#?}", e);
                return;
            }
        };
        log::info!(
            "current overhead on l2 is: {:#?}",
            current_overhead.as_u128()
        );
        ORACLE_SERVICE_METRICS.overhead.set(latest_overhead as i64);

        latest_overhead = latest_overhead.min(MAX_OVERHEAD);
        let abs_diff = U256::from(latest_overhead).abs_diff(current_overhead);
        log::info!(
            "overhead actual_change = {:#?}, expected_change =  {:#?}",
            abs_diff,
            self.overhead_threshold
        );

        if abs_diff < U256::from(self.overhead_threshold) {
            return;
        }

        // Step3. update overhead
        if self.overhead_switch {
            let tx = self
                .l2_oracle
                .set_overhead(U256::from(latest_overhead))
                .legacy();
            let rt = tx.send().await;
            match rt {
                Ok(info) => log::info!("tx of set_overhead has been sent: {:?}", info.tx_hash()),
                Err(e) => log::error!("update overhead error: {:#?}", e),
            }
        }
        *PREV_ROLLUP_L1_BLOCK.lock().await = current_rollup_l1_block;
    }

    async fn overhead_inspect(&self, tx_hash: TxHash, block_num: U64) -> Option<usize> {
        let txn_per_batch_expect: u128 = var("TXN_PER_BATCH")
            .expect("Cannot detect TXN_PER_BATCH env var")
            .parse()
            .expect("Cannot parse TXN_PER_BATCH env var");

        //Step1.  Get transaction
        let result = self.l1_provider.get_transaction(tx_hash).await;
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

        log::info!(
            "rollup l1_tx hash: {:#?}, rollup l1_blocknum: {:#?}",
            tx_hash,
            block_num
        );

        //Step2. Parse transaction data
        let data = tx.input;
        if data.is_empty() {
            log::warn!(
                "overhead_inspect tx.input is empty, tx_hash =  {:#?}",
                tx_hash
            );
            return None;
        }
        let param = match CommitBatchCall::decode(&data) {
            Ok(_param) => _param,
            Err(e) => {
                log::error!(
                    "overhead_inspect decode tx.input error, tx_hash =  {:#?}, err= {:#?}",
                    tx_hash,
                    e
                );
                return None;
            }
        };
        let chunks: Vec<Bytes> = param.batch_data_input.chunks;
        let l2_txn = extract_txn_num(chunks).unwrap_or(0);

        //Step3. Calculate l2 data gas
        let l2_data_gas = match self
            .calculate_l2_data_gas_from_blob(tx_hash, tx.block_hash.unwrap(), block_num, l2_txn)
            .await
        {
            Ok(Some(value)) => value,
            Ok(None) => return None, // Waiting for the next L1 block.
            Err(e) => {
                log::error!("calculate_l2_data_gas_from_blob error: {:#?}", e);
                return None;
            }
        };

        let blob_tx_receipt = match self
            .execution_node
            .query_blob_tx_receipt(hex::encode_prefixed(tx_hash).as_str())
            .await
        {
            Some(r) => r,
            None => {
                log::error!("query_blob_tx_receipt err");
                return None;
            }
        };

        // rollup_gas_used
        let rollup_gas_used = U256::from_str(
            &blob_tx_receipt["result"]["gasUsed"]
                .as_str()
                .unwrap_or("0x0"),
        )
        .unwrap_or(U256::from(0));
        log::info!("rollup_calldata_gas_used: {:?}", rollup_gas_used);
        if rollup_gas_used.is_zero() {
            log::error!(
                "blob tx calldata gas_used is None or 0, tx_hash = {:#?}",
                tx_hash
            );
            return None;
        }

        // blob_gas_price
        let blob_gas_price = U256::from_str(
            &blob_tx_receipt["result"]["blobGasPrice"]
                .as_str()
                .unwrap_or("0x0"),
        )
        .unwrap_or(U256::from(0));

        // effective_gas_price
        let effective_gas_price = U256::from_str(
            &blob_tx_receipt["result"]["effectiveGasPrice"]
                .as_str()
                .unwrap_or("0x0"),
        )
        .unwrap_or(U256::from(0));
        log::info!(
            "blob_gas_price: {:?}, calldata_gas_price: {:?}",
            blob_gas_price,
            effective_gas_price
        );
        if effective_gas_price.is_zero() {
            log::error!(
                "blob tx calldata effective_gas_price is None or 0, tx_hash = {:#?}",
                tx_hash
            );
            return None;
        }

        //Step4. Calculate overhead
        let x: f64 = blob_gas_price.as_u128() as f64 / effective_gas_price.as_u128() as f64;
        let blob_gas_used = if l2_txn > 0 {
            MAX_BLOB_TX_PAYLOAD_SIZE as f64
        } else {
            0.0
        };

        let mut sys_gas: u128 =
            rollup_gas_used.as_u128() + 156400 + (blob_gas_used * x).ceil() as u128;
        sys_gas = if sys_gas < l2_data_gas as u128 {
            0u128
        } else {
            sys_gas - l2_data_gas as u128
        };
        let overhead = if l2_txn as u128 > txn_per_batch_expect {
            sys_gas / l2_txn as u128
        } else {
            sys_gas / txn_per_batch_expect
        };

        let blob_fee_ratio = (blob_gas_used * blob_gas_price.as_u128() as f64).ceil()
            / ((rollup_gas_used * effective_gas_price).as_usize() as f64);

        log::info!(
            "lastest_overhead: {:?},  x:{:?}, l2_txn:{:?}, blob gasFee ratio: {:?}",
            overhead,
            x,
            l2_txn,
            format!("{:.5}", blob_fee_ratio)
        );

        // Set metric
        ORACLE_SERVICE_METRICS.txn_per_batch.set(l2_txn as f64);
        Some(overhead as usize)
    }

    async fn calculate_l2_data_gas_from_blob(
        &self,
        tx_hash: TxHash,
        block_hash: TxHash,
        block_num: U64,
        l2_txn: u64,
    ) -> Result<Option<u64>, String> {
        if l2_txn == 0 {
            return Ok(Some(0));
        }
        let blob_tx = self
            .execution_node
            .query_blob_tx(hex::encode_prefixed(tx_hash).as_str())
            .await
            .ok_or_else(|| "Failed to query blob tx".to_string())?;

        let blob_block = self
            .execution_node
            .query_block(hex::encode_prefixed(block_hash).as_str())
            .await
            .ok_or_else(|| "Failed to query blob block".to_string())?;

        let indexed_hashes = data_and_hashes_from_txs(
            &blob_block["result"]["transactions"]
                .as_array()
                .unwrap_or(&Vec::<Value>::new()),
            &blob_tx["result"],
        );

        if indexed_hashes.is_empty() {
            log::info!("No blob in this batch, batchTxHash ={:#?}", tx_hash);
            return Ok(Some(0));
        }

        let next_block_num = block_num + 1;
        let next_block = self
            .execution_node
            .query_block_by_num(next_block_num.as_u64())
            .await
            .ok_or_else(|| format!("Failed to query block {}", next_block_num))?;

        let prev_beacon_root = match next_block["result"]["parentBeaconBlockRoot"].as_str() {
            Some(r) => r,
            None => {
                log::error!("Next block not produce");
                return Ok(None);
            } // Waiting for the next L1 block.
        };

        let indexes: Vec<u64> = indexed_hashes.iter().map(|item| item.index).collect();
        let sidecars_rt = self
            .beacon_node
            .query_sidecars(prev_beacon_root.to_string(), indexes)
            .await
            .ok_or_else(|| "Failed to query side car".to_string())?;

        let sidecars: &Vec<Value> = sidecars_rt["data"]
            .as_array()
            .ok_or_else(|| "query blob_sidecars empty".to_string())?;

        let tx_payload = extract_tx_payload(indexed_hashes, sidecars)?;

        let tx_payload_gas = data_gas_cost(&tx_payload);
        log::info!("tx_payload_in_blob gas: {}", tx_payload_gas);

        Ok(Some(tx_payload_gas))
    }
}

fn extract_tx_payload(
    indexed_hashes: Vec<IndexedBlobHash>,
    sidecars: &Vec<Value>,
) -> Result<Vec<u8>, String> {
    let mut tx_payload = Vec::<u8>::new();
    for i_h in indexed_hashes {
        if let Some(sidecar) = sidecars.iter().find(|sidecar| {
            sidecar["index"]
                .as_str()
                .unwrap_or("1000")
                .parse::<u64>()
                .unwrap_or(1000)
                == i_h.index
        }) {
            let kzg_commitment = sidecar["kzg_commitment"]
                .as_str()
                .ok_or_else(|| "Failed to fetch kzg commitment from blob".to_string())?;
            let decoded_commitment: Vec<u8> =
                hex::decode(kzg_commitment).map_err(|e| e.to_string())?;
            let actual_versioned_hash = kzg_to_versioned_hash(&decoded_commitment);

            if i_h.hash != actual_versioned_hash {
                log::error!(
                    "expected hash {:?} for blob at index {:?} but got {:?}",
                    i_h.hash,
                    i_h.index,
                    actual_versioned_hash
                );
                return Err("Invalid versionedHash for Blob".to_string());
            }

            let encoded_blob = sidecar["blob"]
                .as_str()
                .ok_or_else(|| format!("Missing blob value in blob_hash: {:?}", i_h.hash))?;
            let decoded_blob = hex::decode(encoded_blob).map_err(|e| {
                format!(
                    "Failed to decode blob, blob_hash: {:?}, err: {}",
                    i_h.hash, e
                )
            })?;

            if decoded_blob.len() != MAX_BLOB_TX_PAYLOAD_SIZE {
                return Err("Invalid length for Blob".to_string());
            }

            let blob_array: [u8; MAX_BLOB_TX_PAYLOAD_SIZE] = decoded_blob.try_into().unwrap();
            let blob_struct = Blob(blob_array);
            let mut decoded_payload = blob_struct
                .decode_raw_tx_payload()
                .map_err(|e| format!("Failed to decode blob tx payload: {}", e))?;
            tx_payload.append(&mut decoded_payload);
        } else {
            return Err(format!(
                "no blob in response matches desired index: {}",
                i_h.index
            ));
        }
    }
    Ok(tx_payload)
}

fn extract_txn_num(chunks: Vec<Bytes>) -> Option<u64> {
    if chunks.is_empty() {
        return None;
    }

    let mut txn_in_batch = 0;
    let mut l1_txn_in_batch = 0;
    for chunk in chunks.iter() {
        let mut chunk_bn: Vec<u64> = vec![];
        let bs: &[u8] = chunk;
        // decode blockcontext from chunk
        // |   1 byte   | 60 bytes | ... | 60 bytes |
        // | num blocks |  block 1 | ... |  block n |
        let num_blocks = U256::from_big_endian(bs.get(..1)?);
        for i in 0..num_blocks.as_usize() {
            let block_num = U256::from_big_endian(bs.get((60.mul(i) + 1)..(60.mul(i) + 1 + 8))?);
            let txs_num =
                U256::from_big_endian(bs.get((60.mul(i) + 1 + 56)..(60.mul(i) + 1 + 58))?);
            let l1_txs_num =
                U256::from_big_endian(bs.get((60.mul(i) + 1 + 58)..(60.mul(i) + 1 + 60))?);
            txn_in_batch += txs_num.as_u32();
            l1_txn_in_batch += l1_txs_num.as_u32();
            chunk_bn.push(block_num.as_u64());
        }
    }
    log::debug!(
        "total_txn_in_batch: {:#?}, l1_txn_in_batch: {:#?}",
        txn_in_batch,
        l1_txn_in_batch
    );
    if txn_in_batch < l1_txn_in_batch {
        log::error!("total_txn_in_batch < l1_txn_in_batch");
        return None;
    }
    return Some((txn_in_batch - l1_txn_in_batch) as u64);
}

fn data_gas_cost(data: &[u8]) -> u64 {
    if data.len() == 0 {
        return 0;
    }
    let (zeroes, ones) = zeroes_and_ones(data);
    // 4 Paid for every zero byte of data or code for a transaction.
    // 16 Paid for every non-zero byte of data or code for a transaction.
    let zeroes_gas = zeroes * 4;
    let ones_gas = ones * 16;
    zeroes_gas + ones_gas
}

fn zeroes_and_ones(data: &[u8]) -> (u64, u64) {
    let mut zeroes = 0;
    let mut ones = 0;

    for &byt in data {
        if byt == 0 {
            zeroes += 1;
        } else {
            ones += 1;
        }
    }
    (zeroes, ones)
}

#[derive(Debug, Serialize, Deserialize)]
struct BlockInfo {
    block_number: U256,
    timestamp: u64,
    base_fee: U256,
    gas_limit: u64,
    num_txs: u64,
}

fn decode_transactions_from_blob(bs: &[u8]) -> Vec<TypedTransaction> {
    let mut txs_decoded: Vec<TypedTransaction> = Vec::new();

    let mut offset: usize = 0;
    while offset < bs.len() {
        let tx_len = *bs.get(offset + 1).unwrap() as usize;
        if tx_len == 0 {
            break;
        }
        let tx_bytes = bs[offset..offset + 2 + tx_len].to_vec();
        let tx_decoded: TypedTransaction = match rlp::decode(&tx_bytes) {
            Ok(tx) => tx,
            Err(e) => {
                log::error!("decode_transactions_from_blob error: {:?}", e);
                return vec![];
            }
        };

        txs_decoded.push(tx_decoded);
        offset += tx_len + 2
    }

    txs_decoded
}

#[derive(Debug, Clone)]
struct IndexedBlobHash {
    index: u64,
    hash: H256,
}

fn data_and_hashes_from_txs(txs: &[Value], target_tx: &Value) -> Vec<IndexedBlobHash> {
    let mut hashes = Vec::new();
    let mut blob_index = 0u64; // index of each blob in the block's blob sidecar

    for tx in txs {
        // skip any non-batcher transactions
        if tx["hash"] != target_tx["hash"] {
            if let Some(blob_hashes) = tx["blobVersionedHashes"].as_array() {
                blob_index += blob_hashes.len() as u64;
            }
            continue;
        }
        if let Some(blob_hashes) = tx["blobVersionedHashes"].as_array() {
            for h in blob_hashes {
                let idh = IndexedBlobHash {
                    index: blob_index,
                    hash: H256::from_str(h.as_str().unwrap()).unwrap(),
                };
                hashes.push(idh);
                blob_index += 1;
            }
        }
    }
    hashes
}

#[tokio::test]
async fn test_decode_transactions_from_blob() {
    use ethers::prelude::*;
    use ethers::types::transaction::eip2718::TypedTransaction;
    use ethers::utils::to_checksum;

    let wallet: LocalWallet = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
        .parse()
        .unwrap();

    let addresses = vec![
        "0x4e6bA705D14b2237374cF3a308ec466cAb24cA6a",
        "0x0425266311AA5858625cD399EADBBfab183494f7",
        "0x1f68c776FBe7285eBe02111F0A982D1640b0a483",
    ];

    let txs: Vec<TypedTransaction> = addresses
        .iter()
        .map(|to| {
            let req = TransactionRequest::new()
                .to(*to)
                .value(1000000000000000000u64)
                .gas(21000)
                .chain_id(1u64);
            TypedTransaction::Legacy(req)
        })
        .collect();

    let mut txs_bytes: Vec<u8> = Vec::new();
    for tx in txs {
        let sig = wallet.sign_transaction(&tx).await.unwrap();
        txs_bytes.extend_from_slice(&tx.rlp_signed(&sig));
    }

    let txs_decoded: Vec<crate::typed_tx::TypedTransaction> =
        decode_transactions_from_blob(txs_bytes.as_slice());

    for (tx, address_str) in txs_decoded.iter().zip(addresses) {
        if let crate::typed_tx::TypedTransaction::Legacy(tr) = tx.clone() {
            let address_to = tr.to.unwrap();
            let to_tx = to_checksum(address_to.as_address().unwrap(), None);
            assert_eq!(to_tx.as_str(), address_str);
        };
    }
}

#[tokio::test]
async fn test_overhead_inspect() {
    use std::sync::Arc;

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let rollup_tx_hash = "0x87b09de64fd9c433226a0c683a3b3c1d1e8ab3fa24f3213fa63e2931f205f8d8";
    let rollup_tx_block_num = 1489357;
    log::info!("rollup_tx_block_num: {:#?}", rollup_tx_block_num);

    let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("Cannot detect GAS_ORACLE_L1_RPC env var");
    let l2_rpc = var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env");
    let overhead_threshold = var("OVERHEAD_THRESHOLD")
        .expect("OVERHEAD_THRESHOLD env")
        .parse()
        .unwrap();
    let l1_rollup_address = Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env")).unwrap();
    let l2_oracle_address =
        Address::from_str(&var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env")).unwrap();
    let private_key = var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env");

    let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();
    let l1_rollup: Rollup<Provider<Http>> =
        Rollup::new(l1_rollup_address, Arc::new(l1_provider.clone()));

    let l2_provider: Provider<Http> = Provider::<Http>::try_from(l2_rpc).unwrap();
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str(private_key.as_str())
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, _>> =
        GasPriceOracle::new(l2_oracle_address, l2_signer);

    let overhead: OverHead = OverHead::new(
        l1_provider,
        l2_oracle,
        l1_rollup,
        overhead_threshold,
        l1_rpc,
        var("GAS_ORACLE_L1_BEACON_RPC")
            .expect("Cannot detect GAS_ORACLE_L1_BEACON_RPC env var")
            .parse()
            .expect("Cannot parse GAS_ORACLE_L1_BEACON_RPC env var"),
        false,
    );

    let latest_overhead = overhead
        .overhead_inspect(
            TxHash::from_str(rollup_tx_hash).unwrap(),
            U64::from(rollup_tx_block_num),
        )
        .await;
    println!("latest_overhead: {:?}", latest_overhead);
    return;
}
