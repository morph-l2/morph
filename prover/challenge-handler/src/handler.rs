use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::util;
use dotenv::dotenv;
use ethers::providers::{Http, Provider};
use ethers::signers::Wallet;
use ethers::types::Address;
use ethers::types::Bytes;
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use std::env::var;
use std::error::Error;
use std::ops::Mul;
use std::str::FromStr;
use std::sync::Arc;
use std::time::Duration;

#[derive(Serialize)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
    pub rpc: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

pub async fn handle_challenge() -> Result<(), Box<dyn Error>> {
    // Prepare parameter.
    dotenv().ok();
    let l1_rpc = var("HANDLER_L1_RPC").expect("Cannot detect L1_RPC env var");
    let l1_rollup_address = var("HANDLER_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let _ = var("HANDLER_PROVER_RPC").expect("Cannot detect PROVER_RPC env var");
    let private_key = var("CHALLENGE_HANDLER_PRIVATE_KEY").expect("Cannot detect L1_ROLLUP_PRIVATE_KEY env var");

    let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc)?;
    let l1_signer = Arc::new(SignerMiddleware::new(
        l1_provider.clone(),
        Wallet::from_str(private_key.as_str())
            .unwrap()
            .with_chain_id(l1_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let l1_rollup: RollupType = Rollup::new(Address::from_str(l1_rollup_address.as_str())?, l1_signer);

    handle_with_prover(l1_provider, l1_rollup).await;

    Ok(())
}

async fn handle_with_prover(l1_provider: Provider<Http>, l1_rollup: RollupType) {
    let l2_rpc = var("HANDLER_L2_RPC").expect("Cannot detect L2_RPC env var");

    loop {
        std::thread::sleep(Duration::from_secs(12));

        // Step1. fetch latest blocknum.
        let latest = match l1_provider.get_block_number().await {
            Ok(bn) => bn,
            Err(e) => {
                log::error!("L1 provider.get_block_number error: {:#?}", e);
                continue;
            }
        };
        // Step2. detecte challenge event.
        let batch_index = match detecte_challenge_event(latest, &l1_rollup, &l1_provider).await {
            Some(value) => value,
            None => continue,
        };
        log::warn!("Challenge event detected, batch index is: {:#?}", batch_index);
        match query_proof(batch_index).await {
            Some(_) => {
                log::info!("query proof and prove state: {:#?}", batch_index);
                prove_state(batch_index, &l1_rollup, &l1_provider).await;
                continue;
            }
            None => (),
        }

        // Step3. query challenged batch for the past 3 days(7200blocks*3 = 3 day).
        let hash = match query_challenged_batch(latest, &l1_rollup, batch_index, &l1_provider).await {
            Some(value) => value,
            None => continue,
        };
        let batch_info = match batch_inspect(&l1_provider, hash).await {
            Some(batch) => batch,
            None => continue,
        };

        // Step4. Make a call to the Prove server.
        let request = ProveRequest {
            batch_index: batch_index,
            chunks: batch_info.clone(),
            rpc: l2_rpc.to_owned(),
        };
        let rt = tokio::task::spawn_blocking(move || util::call_prover(serde_json::to_string(&request).unwrap(), "/prove_batch"))
            .await
            .unwrap();

        match rt {
            Some(info) => {
                match info.as_str() {
                    task_status::STARTED => log::info!("successfully submitted prove task, waiting for proof to be generated"),
                    task_status::PROVING => {
                        log::info!("waiting for proof to be generated");
                        std::thread::sleep(Duration::from_secs(600)); //chunk_prove_time =1h 20min，batch_prove_time = 24min
                        prove_state(batch_index, &l1_rollup, &l1_provider).await;
                        continue;
                    }
                    task_status::PROVED => {
                        log::info!("proof already generated");
                        prove_state(batch_index, &l1_rollup, &l1_provider).await;
                        continue;
                    }
                    _ => {
                        log::error!("submit prove task failed: {:#?}", info);
                        continue;
                    }
                }
            }
            None => {
                log::error!("submit prove task failed");
                continue;
            }
        }

        // Step5. query proof and prove onchain state.
        std::thread::sleep(Duration::from_secs((4800 * batch_info.len() + 1800) as u64)); //chunk_prove_time =1h 20min，batch_prove_time = 24min
        log::info!("sleep finish, query proof and prove state: {:#?}", batch_index);
        prove_state(batch_index, &l1_rollup, &l1_provider).await;
    }
}

async fn prove_state(batch_index: u64, l1_rollup: &RollupType, l1_provider: &Provider<Http>) -> bool {
    loop {
        std::thread::sleep(Duration::from_secs(30));
        let prove_result = match query_proof(batch_index).await {
            Some(pr) => pr,
            None => continue,
        };

        let aggr_proof = Bytes::from(prove_result.proof_data);
        log::info!(
            "starting prove state onchain, batch index = {:#?}, aggr_proof = {:#?}",
            batch_index,
            aggr_proof
        );

        let tx = l1_rollup.prove_state(batch_index, aggr_proof);
        let rt = tx.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => {
                log::info!("tx of prove_state has been sent: {:#?}", pending_tx.tx_hash());
                pending_tx
            }
            Err(e) => {
                log::error!("send tx of prove_state error: {:#?}", e);
                match e {
                    ContractError::Revert(data) => {
                        let msg = String::decode_with_selector(&data).unwrap_or(String::from("decode contract revert error"));
                        log::error!("send tx of prove_state error, msg: {:#?}", msg);
                    }
                    _ => {}
                }
                continue;
            }
        };

        let check_receipt = || async {
            let receipt = l1_provider.get_transaction_receipt(pending_tx.tx_hash()).await;
            if let Err(e) = receipt {
                log::error!("get receipt of prove_state error, msg: {:#?}", e);
                return false;
            }
            match receipt.unwrap() {
                Some(tr) => {
                    // Either 1 (success) or 0 (failure).
                    match tr.status.unwrap_or_default().as_u64() {
                        1 => log::info!("prove_state receipt success: {:#?}", tr.transaction_hash),
                        _ => log::error!("prove_state receipt fail: {:#?}", tr),
                    };
                    return true;
                }
                // Maybe still pending
                None => {
                    log::info!("prove_state tx is pending: {:#?}", pending_tx.tx_hash());
                    return false;
                }
            }
        };
        for _ in 1..10 {
            std::thread::sleep(Duration::from_secs(12));
            if check_receipt().await {
                return true;
            };
        }
    }
}

/**
 * Query the zkevm proof for the specified batch index.
 * Only return result when proof data exists, otherwise return None.
 */
async fn query_proof(batch_index: u64) -> Option<ProveResult> {
    // Make a call to the Prove server.
    let rt = tokio::task::spawn_blocking(move || util::call_prover(batch_index.to_string(), "/query_proof"))
        .await
        .unwrap();
    let rt_text = match rt {
        Some(info) => info,
        None => {
            log::error!("query proof failed");
            return None;
        }
    };

    let prove_result: ProveResult = match serde_json::from_str(rt_text.as_str()) {
        Ok(pr) => pr,
        Err(_) => {
            log::error!("deserialize prove_result failed, batch index = {:#?}", batch_index);
            return None;
        }
    };

    if prove_result.pi_data.is_empty() || prove_result.proof_data.is_empty() {
        log::warn!("query proof of {:#?}, pi_data or  proof_data is empty", batch_index);
        return None;
    }

    return Some(prove_result);
}

async fn query_challenged_batch(latest: U64, l1_rollup: &RollupType, batch_index: u64, l1_provider: &Provider<Http>) -> Option<TxHash> {
    let start = if latest > U64::from(7200 * 3) {
        // Depends on challenge period
        // latest - U64::from(7200 * 3)
        U64::from(1)

    } else {
        U64::from(1)
    };
    let filter = l1_rollup
        .commit_batch_filter()
        .filter
        .from_block(start)
        .topic1(U256::from(batch_index))
        .address(l1_rollup.address());
    let logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("l1_rollup.commit_batch.get_logs error: {:#?}", e);
            return None;
        }
    };

    if logs.is_empty() {
        log::error!("no commit_batch log of {:?}, commit_batch logs is empty", batch_index);
        return None;
    }

    for log in logs {
        if log.topics[1].to_low_u64_be() != batch_index {
            continue;
        }
        let tx_hash = log.transaction_hash.unwrap();
        let receipt = l1_provider.get_transaction_receipt(tx_hash).await.unwrap();
        match receipt {
            Some(tr) => {
                match tr.status.unwrap_or_default().as_u64() {
                    1 => return Some(tx_hash),
                    _ => {
                        log::warn!("commit_batch receipt is fail: {:#?}", tr);
                        continue;
                    }
                };
            }
            None => {
                log::warn!("no commit_batch receipt, batch index = {:?}, tx_hash = {:?}", batch_index, tx_hash);
            }
        }
    }
    log::error!("unable to find valid commit_batch log, batch index = {:?}", batch_index);
    None
}

async fn detecte_challenge_event(latest: U64, l1_rollup: &RollupType, l1_provider: &Provider<Http>) -> Option<u64> {
    let start = if latest > U64::from(7200 * 3) {
        // Depends on challenge period
        // latest - U64::from(7200 * 3)
        U64::from(1)
    } else {
        U64::from(1)
    };
    let filter = l1_rollup.challenge_state_filter().filter.from_block(start).address(l1_rollup.address());
    let mut logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("l1_rollup.challenge_state.get_logs error: {:#?}", e);
            return None;
        }
    };
    log::debug!("l1_rollup.challenge_state.get_logs.len = {:#?}", logs.len());
    if logs.is_empty() {
        log::debug!("no challenge state logs, start blocknum = {:#?}, latest blocknum = {:#?}", start, latest);
        return None;
    }

    //TODO sort by batch index
    //asc
    logs.sort_by(|a, b| a.block_number.unwrap().cmp(&b.block_number.unwrap()));

    //desc
    // logs.sort_by(|a, b| b.block_number.unwrap().cmp(&a.block_number.unwrap()));

    for log in logs {
        let batch_index: u64 = log.topics[1].to_low_u64_be();
        let batch_in_challenge: bool = match l1_rollup.batch_in_challenge(U256::from(batch_index)).await {
            Ok(x) => x,
            Err(e) => {
                log::info!("query l1_rollup.batch_in_challenge error, batch index = {:#?}, {:#?}", batch_index, e);
                return None;
            }
        };

        if batch_in_challenge {
            return Some(batch_index);
        }
        log::info!("batch status not in challenge, batch index = {:#?}", batch_index);
    }
    log::info!("all batch's status not in challenge now");
    None
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
    // log::debug!("batch inspect: tx.input =  {:#?}", data);

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
    let min_gas_limit = param.min_gas_limit;
    log::info!("batch inspect: min_gas_limit =  {:#?}", min_gas_limit);

    let chunks: Vec<Bytes> = param.batch_data.chunks;
    return decode_chunks(chunks);
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

    log::debug!("decode_chunks_blocknum: {:#?}", chunk_with_blocks);
    return Some(chunk_with_blocks);
}

#[tokio::test]
async fn test_decode_chunks() {
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    use std::fs::File;
    use std::io::Read;
    let mut file = File::open("./src/staking.data").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    let input = Bytes::from_str(contents.as_str()).unwrap();

    let param = CommitBatchCall::decode(&input).unwrap();
    let chunks: Vec<Bytes> = param.batch_data.chunks;
    let rt = decode_chunks(chunks).unwrap();
    assert!(rt.len() == 11);
    assert!(rt.get(3).unwrap().len() == 2);
}

#[tokio::test]
async fn test_commit_batch_filter() {
    let l1_provider: Provider<Http> = Provider::<Http>::try_from("https://eth-mainnet.g.alchemy.com/v2/oGHoxvEj_29sdJlOjKMPMHvmqS8gYmLX").unwrap();
    let l1_rollup: Rollup<Provider<Http>> = Rollup::new(
        Address::from_str("0xa13BAF47339d63B743e7Da8741db5456DAc1E556").unwrap(),
        Arc::new(l1_provider.clone()),
    );

    let latest = l1_provider.get_block_number().await.unwrap();
    let filter = l1_rollup.commit_batch_filter().filter.from_block(18548027).topic1(U256::from(20856));
    let logs: Vec<Log> = l1_provider.get_logs(&filter).await.unwrap();

    println!("{:?}", logs.len());
    println!("{:?}", logs.first().unwrap().topics[0]);
}
