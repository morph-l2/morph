use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::metrics::METRICS;
use crate::util;
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
use tokio::time::sleep;

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
    pub blob_kzg: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

const MAX_RETRY_TIMES: u8 = 2;

pub async fn handle_challenge() -> Result<(), Box<dyn Error>> {
    // Prepare parameter.
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
    let wallet_address: Address = l1_signer.address();

    let l1_rollup: RollupType = Rollup::new(Address::from_str(l1_rollup_address.as_str())?, l1_signer);
    handle_with_prover(wallet_address, l1_provider, l1_rollup).await;

    Ok(())
}

async fn handle_with_prover(wallet_address: Address, l1_provider: Provider<Http>, l1_rollup: RollupType) {
    let l2_rpc = var("HANDLER_L2_RPC").expect("Cannot detect L2_RPC env var");

    loop {
        sleep(Duration::from_secs(12)).await;

        // Step1. fetch latest blocknum.
        let latest = match l1_provider.get_block_number().await {
            Ok(bn) => bn,
            Err(e) => {
                log::error!("L1 provider.get_block_number error: {:#?}", e);
                continue;
            }
        };
        log::info!("Current L1 block number: {:#?}", latest);

        // Record wallet balance.
        let balance = match l1_provider.get_balance(wallet_address, None).await {
            Ok(b) => b,
            Err(e) => {
                log::error!("handler_wallet.get_balance error: {:#?}", e);
                return;
            }
        };
        METRICS.wallet_balance.set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));

        // Step2. detecte challenge event.
        let batch_index = match detecte_challenge_event(latest, &l1_rollup, &l1_provider).await {
            Some(value) => value,
            None => continue,
        };
        log::warn!("Challenge event detected, batch index is: {:#?}", batch_index);
        METRICS.detected_batch_index.set(batch_index as i64);
        match query_proof(batch_index).await {
            Some(prove_result) => {
                log::info!("query proof and prove state: {:#?}", batch_index);
                if !prove_result.proof_data.is_empty() {
                    prove_state(batch_index, &l1_rollup).await;
                    continue;
                }
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
        log::info!(
            "batch inspect of: {:?}, chunks.len = {:?}, chunks = {:#?}",
            batch_index,
            batch_info.len(),
            batch_info
        );
        METRICS.chunks_len.set(batch_info.len() as i64);

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
            Some(info) => match info.as_str() {
                task_status::STARTED => log::info!("successfully submitted prove task, waiting for proof to be generated"),
                task_status::PROVING => log::info!("waiting for prev proof to be generated"),
                task_status::PROVED => {
                    log::info!("proof already generated");
                    prove_state(batch_index, &l1_rollup).await;
                    continue;
                }
                _ => {
                    log::error!("submit prove task failed: {:#?}", info);
                    continue;
                }
            },
            None => {
                log::error!("submit prove task failed");
                continue;
            }
        }

        // Step5. query proof and prove onchain state.
        let mut max_waiting_time: usize = 4800 * batch_info.len() + 1800; //chunk_prove_time =1h 20min，batch_prove_time = 24min
        while max_waiting_time > 300 {
            sleep(Duration::from_secs(300)).await;
            max_waiting_time -= 300;
            match query_proof(batch_index).await {
                Some(prove_result) => {
                    log::debug!("query proof and prove state: {:#?}", batch_index);
                    if !prove_result.proof_data.is_empty() {
                        prove_state(batch_index, &l1_rollup).await;
                        break;
                    }
                }
                None => {
                    log::error!("prover status unknown, resubmit task");
                    break;
                }
            }
        }
    }
}

async fn prove_state(batch_index: u64, l1_rollup: &RollupType) -> bool {
    for _ in 0..MAX_RETRY_TIMES {
        sleep(Duration::from_secs(12)).await;
        let prove_result = match query_proof(batch_index).await {
            Some(pr) => pr,
            None => continue,
        };

        if prove_result.proof_data.is_empty() {
            log::warn!("query proof of {:#?}, proof_data is empty", batch_index);
            continue;
        }

        log::info!("starting prove state onchain, batch index = {:#?}", batch_index);
        let aggr_proof = Bytes::from(prove_result.proof_data);
        let kzg_data = Bytes::from(prove_result.blob_kzg);

        let call = l1_rollup.prove_state(batch_index, aggr_proof, kzg_data, 10u32.pow(6));
        let rt = call.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => {
                log::info!("tx of prove_state has been sent: {:#?}", pending_tx.tx_hash());
                pending_tx
            }
            Err(err) => {
                log::error!("send tx of prove_state error: {:#?}", err);
                METRICS.verify_result.set(2);
                match err {
                    ContractError::Revert(data) => {
                        let msg = String::decode_with_selector(&data).unwrap_or(String::from("unknown, decode contract revert error"));
                        log::error!("send tx of prove_state error, msg: {:#?}", msg);
                    }
                    _ => (),
                }
                continue;
            }
        };
        match pending_tx.await {
            Ok(receipt) => {
                match receipt {
                    Some(receipt) => {
                        // Check the status of the tx receipt
                        if receipt.status == Some(1.into()) {
                            log::error!("tx of prove_state success, tx hash: {:?}", receipt.transaction_hash);
                        } else {
                            log::error!("tx of prove_state failed, tx hash: {:?}", receipt.transaction_hash);
                        }
                    }
                    None => {
                        log::error!("No tx receipt found, may still be in pending status or has been dropped");
                    }
                }
            }
            Err(error) => log::error!("provider error: {:?}", error),
        }
    }
    return false;
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

    return Some(prove_result);
}

async fn query_challenged_batch(latest: U64, l1_rollup: &RollupType, batch_index: u64, l1_provider: &Provider<Http>) -> Option<TxHash> {
    let start = if latest > U64::from(7200 * 3) {
        // Depends on challenge period
        latest - U64::from(7200 * 3)
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
        latest - U64::from(7200 * 3)
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
    log::info!(
        "{:#?} batches have already been challenged, and been found in recent 7200x3 L1 blocks.",
        logs.len()
    );

    if logs.is_empty() {
        log::debug!("no challenge state logs, start blocknum = {:#?}, latest blocknum = {:#?}", start, latest);
        return None;
    }
    logs.sort_by(|a, b| a.block_number.unwrap().cmp(&b.block_number.unwrap()));

    for log in logs {
        let batch_index: u64 = log.topics[1].to_low_u64_be();
        let batch_in_challenge: bool = match l1_rollup.batch_in_challenge(U256::from(batch_index)).await {
            Ok(x) => x,
            Err(e) => {
                log::info!("query l1_rollup.batch_in_challenge error, batch index = {:#?}, {:#?}", batch_index, e);
                return None;
            }
        };
        let is_batch_finalized: bool = match l1_rollup.is_batch_finalized(U256::from(batch_index)).await {
            Ok(x) => x,
            Err(e) => {
                log::info!("query l1_rollup.is_batch_finalized error, batch index = {:#?}, {:#?}", batch_index, e);
                return None;
            }
        };
        if batch_in_challenge && !is_batch_finalized {
            return Some(batch_index);
        }
        log::debug!("batch status not in challenge or already finalized, batch index = {:#?}", batch_index);
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
    let chunks: Vec<Bytes> = param.batch_data.chunks;
    return decode_chunks(chunks);
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
    METRICS.txn_len.set(txn_in_batch.into());
    log::debug!("decode_chunks_blocknum: {:#?}", chunk_with_blocks);
    log::info!("max_l2txn_in_chunk: {:#?}", max_txn_in_chunk);
    log::info!("total_l2txn_in_batch: {:#?}", txn_in_batch);
    return Some(chunk_with_blocks);
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
    let chunks: Vec<Bytes> = param.batch_data.chunks;
    let rt = decode_chunks(chunks).unwrap();
    assert!(rt.len() == 11);
    assert!(rt.get(3).unwrap().len() == 2);
}
