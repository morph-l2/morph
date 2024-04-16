use crate::abi::rollup_abi::{CommitBatchCall, Rollup};
use crate::abi::shadow_rollup_abi::ShadowRollup;
use crate::metrics::METRICS;
use crate::{util, BatchInfo};
use ethers::providers::{Http, Provider};
use ethers::signers::Wallet;
use ethers::types::Address;
use ethers::types::Bytes;
use ethers::{abi::AbiDecode, prelude::*};
use serde::{Deserialize, Serialize};
use serde_json::Value;
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
    pub blob_kzg: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

type ShadowRollupType = ShadowRollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

pub async fn prove(batch_info: BatchInfo) -> Result<(), Box<dyn Error>> {
    // Prepare parameter.
    let l1_rpc = var("HANDLER_L1_RPC").expect("Cannot detect L1_RPC env var");
    let l1_shadow_rollup_address = var("HANDLER_L1_SHADOW_ROLLUP").expect("Cannot detect L1_SHADOW_ROLLUP env var");
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

    let l1_shadow_rollup: ShadowRollupType = ShadowRollup::new(Address::from_str(l1_shadow_rollup_address.as_str())?, l1_signer.clone());

    // Record wallet balance.
    let balance = match l1_provider.get_balance(wallet_address, None).await {
        Ok(b) => b,
        Err(e) => {
            log::error!("handler_wallet.get_balance error: {:#?}", e);
            return Ok(());
        }
    };
    METRICS.wallet_balance.set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));

    handle_with_prover(&batch_info, l1_shadow_rollup).await;

    Ok(())
}

async fn handle_with_prover(batch_info: &BatchInfo, l1_shadow_rollup: ShadowRollupType) {
    let l2_rpc = var("HANDLER_L2_RPC").expect("Cannot detect L2_RPC env var");
    let batch_index = batch_info.batch_index;
    let chunks = &batch_info.chunks;
    let chunks_len = chunks.len();

    loop {
        std::thread::sleep(Duration::from_secs(12));

        // Step2. detecte challenge event.
        log::warn!("Challenge event detected, batch index is: {:#?}", batch_index);
        METRICS.detected_batch_index.set(batch_index as i64);
        match query_proof(batch_index).await {
            Some(prove_result) => {
                log::info!("query proof and prove state: {:#?}", batch_index);
                if !prove_result.proof_data.is_empty() {
                    prove_state(batch_index, &l1_shadow_rollup).await;
                    continue;
                }
            }
            None => (),
        }

        // Step3. query challenged batch for the past 3 days(7200blocks*3 = 3 day).

        log::info!("batch inspect: chunks.len =  {:#?}", chunks_len);
        METRICS.chunks_len.set(chunks_len as i64);

        // Step4. Make a call to the Prove server.
        let request = ProveRequest {
            batch_index: batch_index,
            chunks: chunks.clone(),
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
                    prove_state(batch_index, &l1_shadow_rollup).await;
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
        let mut max_waiting_time: usize = 4800 * chunks_len + 1800; //chunk_prove_time =1h 20min，batch_prove_time = 24min
        while max_waiting_time > 300 {
            std::thread::sleep(Duration::from_secs(300));
            max_waiting_time -= 300;
            match query_proof(batch_index).await {
                Some(prove_result) => {
                    log::debug!("query proof and prove state: {:#?}", batch_index);
                    if !prove_result.proof_data.is_empty() {
                        prove_state(batch_index, &l1_shadow_rollup).await;
                        break;
                    }
                }
                None => {
                    log::error!("prover status unknown, resubmit task");
                    break; // resubmit task
                }
            }
        }
    }
}

async fn prove_state(batch_index: u64, l1_rollup: &ShadowRollupType) -> bool {
    for _ in 0..2 {
        std::thread::sleep(Duration::from_secs(30));
        let prove_result = match query_proof(batch_index).await {
            Some(pr) => pr,
            None => continue,
        };

        l1_rollup.address();

        if prove_result.proof_data.is_empty() {
            log::warn!("query proof of {:#?}, proof_data is empty", batch_index);
            continue;
        }

        log::info!("starting prove state onchain, batch index = {:#?}", batch_index);
        let aggr_proof = Bytes::from(prove_result.proof_data);
        let kzg_data = Bytes::from(prove_result.blob_kzg);

        let call = l1_rollup.prove_state(batch_index, aggr_proof, kzg_data);
        let rt = call.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => {
                log::info!("tx of prove_state has been sent: {:#?}", pending_tx.tx_hash());
                pending_tx
            }
            Err(e) => {
                log::error!("send tx of prove_state error: {:#?}", e);
                METRICS.verify_result.set(2);
                match e {
                    ContractError::Revert(data) => {
                        let msg = String::decode_with_selector(&data).unwrap_or(String::from("decode contract revert error"));
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
                        // Check the status of the transaction receipt
                        if receipt.status == Some(1.into()) {
                            println!(
                                "Transaction successfully included in the blockchain, transaction hash: {:?}",
                                receipt.transaction_hash
                            );
                        } else {
                            println!("Transaction failed, transaction hash: {:?}", receipt.transaction_hash);
                        }
                    }
                    None => {
                        println!("No transaction receipt found, the transaction might still be in pending status or has been dropped by the network");
                    }
                }
            }
            Err(error) => match error {
                ProviderError::JsonRpcClientError(tx_error) => {
                    println!("Transaction rpc error: {:?}", tx_error);
                }
                _ => {
                    println!("Encountered an unknown error: {:?}", error);
                }
            },
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

/**
 * Query the kzg commitment of batch.
 */
async fn query_kzg_commitment(batch_index: u64) -> Option<Vec<u8>> {
    // Make a call to the sequencer.
    let rt = tokio::task::spawn_blocking(move || util::call_sequencer(batch_index.to_string(), "/morph_getRollupBatchByIndex"))
        .await
        .unwrap();
    let rt_text = match rt {
        Some(info) => info,
        None => {
            log::error!("query kzg commitment failed");
            return None;
        }
    };

    let rollup_batch = match serde_json::from_str::<Value>(&rt_text) {
        Ok(parsed) => parsed,
        Err(_) => {
            log::error!("deserialize rollup_batch failed, batch index = {:#?}", batch_index);
            return None;
        }
    };

    let commitments = match rollup_batch["sidecar"]["commitments"].as_array() {
        Some(c) => c,
        None => {
            log::error!("deserialize rollup_batch_commitments failed, batch index = {:#?}", batch_index);
            return None;
        }
    };
    if commitments.is_empty() {
        log::error!("rollup batch has empty kzg commitments, batch index = {:#?}", batch_index);
        return None;
    }

    // ComputeProof computes the KZG proof at the given point for the polynomial

    let commitment_bytes: Vec<u8> = match hex::decode(commitments.first().unwrap().as_str().unwrap_or("0xf")) {
        Ok(cb) => cb,
        Err(_) => {
            log::error!("deserialize commitment failed, batch index = {:#?}", batch_index);
            return None;
        }
    };
    return Some(commitment_bytes);
}
