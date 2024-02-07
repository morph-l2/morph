use crate::abi::gas_price_oracle_abi::GasPriceOracle;
use crate::abi::rollup_abi::{Rollup, SubmitBatchesCall};
use crate::metrics::ORACLE_SERVICE_METRICS;
use crate::typed_tx::TypedTransaction;
use ethers::utils::rlp;
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use std::env::var;

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
    let filter = l1_rollup.submit_batches_filter().filter.from_block(start);

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

    let overhead = match overhead_inspect(&l1_provider, log.transaction_hash.unwrap()).await {
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

async fn overhead_inspect(l1_provider: &Provider<Http>, hash: TxHash) -> Option<usize> {
    let txn_per_block_expect: f64 = var("TXN_PER_BLOCK")
        .expect("Cannot detect TXN_PER_BLOCK env var")
        .parse()
        .expect("Cannot parse TXN_PER_BLOCK env var");
    let txn_per_batch_expect: f64 = var("TXN_PER_BATCH")
        .expect("Cannot detect TXN_PER_BATCH env var")
        .parse()
        .expect("Cannot parse TXN_PER_BATCH env var");

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
        log::warn!("overhead_inspect tx.input is empty, tx_hash =  {:#?}", hash);
        return None;
    }
    let param = if let Ok(_param) = SubmitBatchesCall::decode(&data) {
        _param
    } else {
        log::error!(
            "overhead_inspect decode tx.input error, tx_hash =  {:#?}",
            hash
        );
        return None;
    };
    let batches = param.batches;
    if batches.is_empty() {
        return None;
    }

    //Step3. Calculate gas consumption
    let mut block_num = 0;
    let mut tx_num = 0;
    let mut flexible_cost = 0u64;
    let mut witness_cost = 0u64;
    let mut transactions_cost = 0u64;

    for batch in batches.iter() {
        let this_witness_c = data_gas_cost(&batch.block_witness);
        let this_tx_c = data_gas_cost(&batch.transactions);

        witness_cost += this_witness_c;
        transactions_cost += this_tx_c;
        flexible_cost += this_witness_c + this_tx_c;

        let block_contexts = decode_block_context(&batch.block_witness);
        block_num += block_contexts.unwrap_or(vec![]).len();

        let txs = decode_transactions(&batch.transactions);
        tx_num += txs.len();
    }

    let data_gas = data_gas_cost(&data);

    let receipt = match l1_provider.get_transaction_receipt(hash).await {
        Ok(Some(r)) => r,
        Ok(None) => {
            log::error!(
                "l1_provider.get_transaction_receipt is none, tx_hash = {:#?}",
                hash
            );
            return None;
        }
        Err(e) => {
            log::error!("l1_provider.get_transaction_receipt err: {:#?}", e);
            return None;
        }
    };
    let rollup_gas_used = receipt.gas_used.unwrap_or_default();
    if rollup_gas_used.is_zero() {
        log::error!(
            "l1_provider.get_transaction_receipt gas_used is None or 0, tx_hash = {:#?}",
            hash
        );
        return None;
    }

    //Step4. Calculate overhead
    let g = rollup_gas_used - flexible_cost;
    log::info!(
        "gasinspect => tx_hash: {:?}, batch_num: {:?}, block_num: {:?}, tx_num: {:?}, gas used total: {:?}, call_data: {:?}, witness_cost: {:?},
         transactions_cost: {:?}, otherFixedCost: {:?}, execution cost: {:?}, G: fixedCallData+execution cost: {:?}",
         hash, batches.len(), block_num, tx_num, rollup_gas_used, data_gas, witness_cost, transactions_cost, data_gas - flexible_cost, rollup_gas_used - data_gas, g
    );
    let mut overhead = 512.0;
    let txn_per_block = tx_num as f64 / block_num as f64;
    let txn_per_batch = tx_num as f64 / batches.len() as f64;
    // Set metric
    ORACLE_SERVICE_METRICS.txn_per_batch.set(txn_per_batch);

    overhead += if txn_per_block > txn_per_block_expect {
        580.0 / txn_per_block
    } else {
        580.0 / txn_per_block_expect
    };

    overhead += if txn_per_batch > txn_per_batch_expect {
        g.as_usize() as f64 / txn_per_batch
    } else {
        g.as_usize() as f64 / txn_per_batch_expect
    };

    log::info!("overhead inspection result is: {:#?}", overhead as usize);
    Some(overhead as usize)
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

fn decode_block_context(bs: &[u8]) -> Result<Vec<BlockInfo>, ethers::core::abi::Error> {
    let mut block_contexts = Vec::new();
    let mut reader = bs;

    while !reader.is_empty() {
        let mut block = BlockInfo {
            block_number: 0.into(),
            timestamp: 0,
            base_fee: 0.into(),
            gas_limit: 0,
            num_txs: 0,
        };

        //block_number
        let bs_block_number = reader.get(..32).unwrap();
        reader = &reader[32..];
        block.block_number = U256::from_big_endian(bs_block_number);

        //timestamp
        let timestamp: [u8; 8] = reader[..8].try_into().unwrap();
        block.timestamp = u64::from_be_bytes(timestamp);
        reader = &reader[8..];

        //base_fee
        let bs_base_fee = reader.get(..32).unwrap();
        block.base_fee = U256::from_big_endian(bs_base_fee);
        reader = &reader[32..];

        //gas_limit
        let gas_limit: [u8; 8] = reader[..8].try_into().unwrap();
        block.gas_limit = u64::from_be_bytes(gas_limit);
        reader = &reader[8..];

        //num_txs
        let num_txs: [u8; 8] = reader[..8].try_into().unwrap();
        block.num_txs = u64::from_be_bytes(num_txs);
        reader = &reader[8..];

        //drop txHash
        reader = &reader[block.num_txs as usize * 32..];

        block_contexts.push(block);
    }

    Ok(block_contexts)
}

fn decode_transactions(bs: &[u8]) -> Vec<TypedTransaction> {
    if bs.is_empty() {
        return vec![];
    }
    let rlp = rlp::Rlp::new(bs);
    let transactions: Vec<TypedTransaction> = rlp::decode_list(bs);
    if rlp.item_count().unwrap() != transactions.len() {
        log::error!(
            "txn rlp.item_count is wrong: {:#?}",
            rlp.item_count().unwrap()
        );
    }
    transactions
}
