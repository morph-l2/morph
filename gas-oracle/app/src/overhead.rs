use crate::abi::gas_price_oracle_abi::GasPriceOracle;
use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::blob::Blob;
use crate::blob_client;
use crate::metrics::ORACLE_SERVICE_METRICS;
use crate::typed_tx::TypedTransaction;
use ethers::utils::{hex, rlp};
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use serde_json::Value;
use std::env::var;
use std::io::{Cursor, Read};
use std::ops::Mul;
use std::str::FromStr;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

/// Update overhead
/// Calculate the average cost of the latest roll up and set it to the GasPriceOrale contract on the L2 network.
pub async fn update(
    l1_provider: Provider<Http>,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
    l1_rollup: Rollup<Provider<Http>>,
    overhead_threshold: u128,
) {
    // Step1. fetch latest batches and calculate overhead.
    let latest = match l1_provider.get_block_number().await {
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
    let filter = l1_rollup
        .commit_batch_filter()
        .filter
        .from_block(start)
        .address(l1_rollup.address());

    let mut logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
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

    let overhead =
        match overhead_inspect(&l1_provider, log.transaction_hash.unwrap(), U64::from(100)).await {
            Some(overhead) => overhead,
            None => {
                log::info!(
                    "overhead is none, skip update, tx_hash ={:#?}",
                    log.transaction_hash.unwrap()
                );
                return;
            }
        };

    // Step2. fetch current overhead on l2.
    let current_overhead = match l2_oracle.overhead().await {
        Ok(ov) => ov,
        Err(e) => {
            log::error!("query l2_oracle.overhead error: {:#?}", e);
            return;
        }
    };
    log::info!("current overhead is: {:#?}", current_overhead.as_u128());
    ORACLE_SERVICE_METRICS
        .overhead
        .set(i64::try_from(current_overhead).unwrap_or(-1));

    let abs_diff = U256::from(overhead).abs_diff(current_overhead);
    if abs_diff < U256::from(overhead_threshold) {
        log::info!(
            "overhead change value below threshold, change value = : {:#?}",
            abs_diff
        );
        return;
    }

    // Step3. update overhead
    let tx = l2_oracle.set_overhead(U256::from(overhead)).legacy();
    let rt = tx.send().await;
    match rt {
        Ok(info) => log::info!("tx of update_overhead has been sent: {:?}", info.tx_hash()),
        Err(e) => log::error!("update overhead error: {:#?}", e),
    }
}

async fn overhead_inspect(
    l1_provider: &Provider<Http>,
    tx_hash: TxHash,
    block_num: U64,
) -> Option<usize> {
    let txn_per_batch_expect: u128 = var("TXN_PER_BATCH")
        .expect("Cannot detect TXN_PER_BATCH env var")
        .parse()
        .expect("Cannot parse TXN_PER_BATCH env var");

    //Step1.  Get transaction
    let result = l1_provider.get_transaction(tx_hash).await;
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

    log::info!("rollup tx hash: {:#?}", tx_hash);
    log::info!("rollup blocknum = {:#?}", block_num);

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
    let chunks: Vec<Bytes> = param.batch_data.chunks;
    if chunks.is_empty() {
        return None;
    }
    let l2_txn = extract_txn_num(chunks).unwrap_or(0);

    let l2_data_gas =
        match calculate_l2_data_gas_from_blob(tx_hash, tx.block_hash.unwrap(), block_num, l2_txn)
            .await
        {
            Ok(Some(value)) => value,
            Ok(None) => return None, // Waiting for the next L1 block.
            Err(e) => {
                log::error!("calculate_l2_data_gas_from_blob error: {:?}", e);
                return None;
            }
        };

    let blob_tx_receipt =
        match blob_client::query_blob_tx_receipt(hex::encode_prefixed(tx_hash).as_str()).await {
            Some(r) => r,
            None => {
                log::error!("l1_provider.get_transaction_receipt err");
                return None;
            }
        };

    let rollup_gas_used = U256::from_str(
        &blob_tx_receipt["result"]["gasUsed"]
            .as_str()
            .unwrap_or("0x0"),
    )
    .unwrap_or(U256::from(0));
    log::info!("rollup_gas_used: {:?}", rollup_gas_used);

    if rollup_gas_used.is_zero() {
        log::error!(
            "l1_provider.get_transaction_receipt gas_used is None or 0, tx_hash = {:#?}",
            tx_hash
        );
        return None;
    }
    let blob_gas_price = U256::from_str(
        &blob_tx_receipt["result"]["blobGasPrice"]
            .as_str()
            .unwrap_or("0x0"),
    )
    .unwrap_or(U256::from(0));

    let effective_gas_price = U256::from_str(
        &blob_tx_receipt["result"]["blobGasUsed"]
            .as_str()
            .unwrap_or("0x0"),
    )
    .unwrap_or(U256::from(0));

    log::info!("blob_gas_price: {:?}", blob_gas_price);
    log::info!("effective_gas_price: {:?}", effective_gas_price);

    //Step4. Calculate overhead
    let x = blob_gas_price.as_u128() as f64 / effective_gas_price.as_u128() as f64;
    let blob_gas_used = if l2_txn > 0 {
        MAX_BLOB_TX_PAYLOAD_SIZE as f64
    } else {
        0.0
    };
    let sys_gas: u128 =
        rollup_gas_used.as_u128() + (blob_gas_used * x).ceil() as u128 - l2_data_gas as u128;
    let overhead = if l2_txn as u128 > txn_per_batch_expect {
        sys_gas / l2_txn as u128
    } else {
        sys_gas / txn_per_batch_expect
    };
    log::info!("overhead: {:?}", overhead);

    // Set metric
    ORACLE_SERVICE_METRICS.txn_per_batch.set(l2_txn as f64);
    Some(overhead as usize)
}

use std::convert::TryInto;

async fn calculate_l2_data_gas_from_blob(
    tx_hash: TxHash,
    block_hash: TxHash,
    block_num: U64,
    l2_txn: u64,
) -> Result<Option<u64>, String> {
    if l2_txn == 0 {
        return Ok(Some(0));
    }
    let blob_tx = blob_client::query_blob_tx(hex::encode_prefixed(tx_hash).as_str())
        .await
        .ok_or_else(|| "Failed to query blob tx".to_string())?;

    let blob_block = blob_client::query_block(hex::encode_prefixed(block_hash).as_str())
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
    let next_block = blob_client::query_block_by_num(next_block_num.as_u64())
        .await
        .ok_or_else(|| format!("Failed to query block {}", next_block_num))?;

    let prev_beacon_root = match next_block["result"]["parentBeaconBlockRoot"].as_str() {
        Some(r) => r,
        None => return Ok(None), // Waiting for the next L1 block.
    };

    let indexes: Vec<u64> = indexed_hashes.iter().map(|item| item.index).collect();
    let sidecars_rt = blob_client::query_sidecars(prev_beacon_root.to_string(), indexes)
        .await
        .ok_or_else(|| "Failed to query side car".to_string())?;

    let sidecars: &Vec<Value> = sidecars_rt["data"]
        .as_array()
        .ok_or_else(|| "query blob_sidecars empty".to_string())?;

    let mut tx_payload = Vec::<u8>::new();
    for i_h in indexed_hashes {
        if let Some(sidecar) = sidecars
            .iter()
            .find(|sidecar| sidecar["index"] == i_h.index)
        {
            let versioned_hash = i_h.hash;
            //TODO KZGToVersionedHash

            let blob_value = sidecar["blob"]
                .as_str()
                .ok_or_else(|| format!("Missing blob value in tx_hash: {:?}", tx_hash))?;

            let blob_bytes = hex::decode(blob_value).map_err(|e| {
                format!("Failed to decode blob, tx_hash: {:?}, err: {}", tx_hash, e)
            })?;

            if blob_bytes.len() != MAX_BLOB_TX_PAYLOAD_SIZE {
                return Err("Invalid length for Blob".to_string());
            }

            let array: [u8; MAX_BLOB_TX_PAYLOAD_SIZE] = blob_bytes.try_into().unwrap();
            let blob = Blob(array);
            let mut data = blob
                .decode_raw_tx_payload()
                .map_err(|e| format!("Failed to decode raw tx payload: {}", e))?;
            tx_payload.append(&mut data);
        } else {
            return Err(format!(
                "no blob in response matches desired index: {}",
                i_h.index
            ));
        }
    }

    let blob_data_gas = data_gas_cost(&tx_payload);
    log::info!("tx_payload_in_blob gas: {}", blob_data_gas);

    let (txs, position) = decode_transactions_from_blob(&tx_payload, l2_txn);

    let l2_data_gas = data_gas_cost(&tx_payload[0..position as usize]);
    log::info!(
        "decode transactions from blob, position: {:#?}, l2_data_gas: {:#?} , txs.len(): {:#?}",
        position,
        l2_data_gas,
        txs.len()
    );

    if let Some(tx) = txs.last() {
        log_chain_id(tx);
    }
    Ok(Some(l2_data_gas))
}

fn log_chain_id(tx: &TypedTransaction) {
    match tx {
        TypedTransaction::Legacy(tx_req) => {
            log::info!("Legacy.chain_id: {}", tx_req.chain_id.unwrap());
        }
        TypedTransaction::Eip2930(tx_req) => {
            log::info!("Eip2930.chain_id: {}", tx_req.tx.chain_id.unwrap());
        }
        TypedTransaction::Eip1559(tx_req) => {
            log::info!("Eip1559.chain_id: {}", tx_req.chain_id.unwrap());
        }
        TypedTransaction::L1MessageTx(tx_req) => {
            log::info!("L1MessageTx.chain_id: {}", tx_req.chain_id.unwrap());
        }
    }
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
    log::info!("total_txn_in_batch: {:#?}", txn_in_batch);
    log::info!("l1_txn_in_batch: {:#?}", l1_txn_in_batch);
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

fn decode_transactions_from_blob(bs: &[u8], tx_num: u64) -> (Vec<TypedTransaction>, u64) {
    let mut cursor = Cursor::new(bs);
    let mut txs = Vec::new();

    for _ in 0..tx_num {
        let mut tx_len_buf = [0u8; 4];
        if cursor.read_exact(&mut tx_len_buf).is_err() {
            break;
        }

        let tx_len = u32::from_be_bytes(tx_len_buf);
        let mut tx_bytes = vec![0u8; tx_len as usize];
        cursor.read_exact(&mut tx_bytes).unwrap();

        let tx: TypedTransaction = rlp::decode(&tx_bytes).unwrap();
        txs.push(tx);
    }

    (txs, cursor.position())
}

#[derive(Debug, Clone)]
struct IndexedBlobHash {
    index: u64,
    hash: H256,
}

fn data_and_hashes_from_txs(txs: &[Value], target_tx: &Value) -> Vec<IndexedBlobHash> {
    let mut hashes = Vec::new();
    let mut blob_index = 0u64; // index of each blob in the block's blob sidecar

    log::info!("txs.len ={:#?}", txs.len());

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
async fn test_overhead_inspect() {
    use std::sync::Arc;

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("Cannot detect GAS_ORACLE_L1_RPC env var");
    let rollup_tx_hash = var("ROLLUP_TX_HASH");
    let rollup_tx_block_num = var("ROLLUP_TX_BLOCK_NUM");
    log::info!("rollup_tx_block_num: {:#?}", rollup_tx_block_num);

    let l1_rollup_address: String = var("L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc).unwrap();
    let l1_rollup: Rollup<Provider<Http>> = Rollup::new(
        Address::from_str(l1_rollup_address.as_str()).unwrap(),
        Arc::new(l1_provider.clone()),
    );

    if rollup_tx_hash.is_ok() && rollup_tx_block_num.is_ok() {
        overhead_inspect(
            &l1_provider,
            TxHash::from_str(rollup_tx_hash.unwrap().as_str()).unwrap(),
            U64::from(rollup_tx_block_num.unwrap().parse::<u64>().unwrap()),
        )
        .await;
        return;
    }

    // Step1. fetch latest batches and calculate overhead.
    let latest = match l1_provider.get_block_number().await {
        Ok(bn) => bn,
        Err(e) => {
            log::error!("overhead.l1_provider.get_block_number error: {:#?}", e);
            return;
        }
    };
    let start = if latest > U64::from(200) {
        latest - U64::from(200) //200
    } else {
        U64::from(1)
    };
    let filter = l1_rollup
        .commit_batch_filter()
        .filter
        .from_block(start)
        .address(l1_rollup.address());

    let mut logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("overhead.l1_provider.get_logs error: {:#?}", e);
            return;
        }
    };
    log::info!(
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

    let test_hash =
        TxHash::from_str("0xff6546936f2e095e214c1e1b3bb8bfc7751524948e3af70c7811a9769ca026d9")
            .unwrap();

    let overhead = match overhead_inspect(
        &l1_provider,
        log.transaction_hash.unwrap(),
        log.block_number.unwrap(),
    )
    .await
    {
        Some(overhead) => overhead,
        None => {
            log::info!(
                "overhead is none, skip update, tx_hash ={:#?}",
                "log.transaction_hash.unwrap()"
            );
            return;
        }
    };
}

