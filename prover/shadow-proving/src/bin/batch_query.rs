use challenge_handler::abi::rollup_abi::{CommitBatchCall, Rollup};
use dotenv::dotenv;
use env_logger::Env;
use ethers::prelude::*;
use ethers::signers::Wallet;
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use std::env::var;
use std::error::Error;
use std::ops::Mul;
use std::str::FromStr;
use std::sync::Arc;

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

/**
 * Search for the latest batch to challenge
 */
#[tokio::main]
pub async fn main() -> Result<(), Box<dyn Error>> {
    // Prepare env.
    env_logger::Builder::from_env(Env::default().default_filter_or("debug")).init();
    dotenv().ok();
    let l1_rpc = var("L1_RPC").expect("Cannot detect L1_RPC env var");
    let l1_rollup_address = var("L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let private_key = var("CHALLENGER_PRIVATEKEY").expect("Cannot detect CHALLENGER_PRIVATEKEY env var");

    let target_batch: u64 = var("TARGET_BATCH")
        .expect("Cannot detect TARGET_BATCH env var")
        .parse()
        .expect("Cannot parse TARGET_BATCH env var");


    let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc)?;
    let l1_signer = Arc::new(SignerMiddleware::new(
        l1_provider.clone(),
        Wallet::from_str(private_key.as_str())
            .unwrap()
            .with_chain_id(l1_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let challenger_address = l1_signer.address();
    let l1_rollup: RollupType = Rollup::new(Address::from_str(l1_rollup_address.as_str())?, l1_signer);

    // Check rollup state.
    let is_challenger: bool = match l1_rollup.is_challenger(challenger_address).await {
        Ok(x) => x,
        Err(e) => {
            log::info!("query l1_rollup.is_challenger error: {:#?}", e);
            return Ok(());
        }
    };
    log::info!("address({:#?})  is_challenger: {:#?}", challenger_address, is_challenger);

    let challenger_balance = l1_provider.get_balance(challenger_address, None).await.unwrap();
    log::info!("challenger_eth_balance: {:#?}", challenger_balance);

    let finalization_period = l1_rollup.finalization_period_seconds().await?;
    let proof_window = l1_rollup.proof_window().await?;
    log::info!("finalization_period: ({:#?})  proof_window: {:#?}", finalization_period, proof_window);

    // Search for the latest batch
    let latest = match l1_provider.get_block_number().await {
        Ok(bn) => bn,
        Err(e) => {
            log::error!("L1 provider.get_block_number error: {:#?}", e);
            return Ok(());
        }
    };

    log::info!("latest blocknum = {:#?}", latest);
    let start = if latest > U64::from(1000) {
        // latest - U64::from(1000)
        U64::from(1)
    } else {
        U64::from(1)
    };

    let filter = l1_rollup
        .commit_batch_filter()
        .filter
        .from_block(start)
        .topic1(U256::from(target_batch))
        .address(l1_rollup.address());
    let logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("l1_rollup.commit_batch.get_logs error: {:#?}", e);
            vec![]
        }
    };

    if logs.is_empty() {
        log::error!("no commit_batch log of {:?}, commit_batch logs is empty", target_batch);
        return Ok(());
    }

    assert!(logs.len() == 1, "logs.len() should equals 1");

    let log = logs.first().unwrap();
    let tx_hash = log.transaction_hash.unwrap();
    let batch_index = log.topics[1].to_low_u64_be();
    let result = l1_provider.get_transaction(tx_hash).await.unwrap().unwrap();
    let data = result.input;

    // log::info!("batch inspect: tx.input =  {:#?}", data);
    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", tx_hash);
        return Ok(());
    }
    let param = if let Ok(_param) = CommitBatchCall::decode(&data) {
        _param
    } else {
        log::error!("batch inspect: decode tx.input error, tx_hash =  {:#?}", tx_hash);
        return Ok(());
    };
    // let min_gas_limit = param.min_gas_limit;
    // log::info!("batch inspect: min_gas_limit =  {:#?}", min_gas_limit);

    let chunks: Vec<Bytes> = param.batch_data.chunks;

    let chunk_with_blocks = decode_chunks(chunks).unwrap_or_default();
    log::info!("batch_index: {:#?},  decode_chunks_blocknum: {:#?}", batch_index, chunk_with_blocks);

    Ok(())
}

fn decode_chunks(chunks: Vec<Bytes>) -> Option<Vec<Vec<u64>>> {
    if chunks.is_empty() {
        return None;
    }

    let mut chunk_with_blocks: Vec<Vec<u64>> = vec![];
    for chunk in chunks.iter() {
        let mut chunk_bn: Vec<u64> = vec![];
        let bs: &[u8] = chunk;

        // decode blocks from chunk
        // |   1 byte   | 60 bytes | ... | 60 bytes |
        // | num blocks |  block 1 | ... |  block n |
        let num_blocks = U256::from_big_endian(bs.get(..1)?);
        for i in 0..num_blocks.as_usize() {
            let block_num = U256::from_big_endian(bs.get((60.mul(i) + 1)..(60.mul(i) + 1 + 8))?);
            chunk_bn.push(block_num.as_u64());
        }

        chunk_with_blocks.push(chunk_bn);
    }

    // log::debug!("decode_chunks_blocknum: {:#?}", chunk_with_blocks);
    return Some(chunk_with_blocks);
}
