use crate::abi::shadow_rollup_abi::ShadowRollup;
use crate::metrics::METRICS;
use crate::{util, BatchInfo, ShadowRollupType};
use ethers::prelude::*;
use ethers::providers::{Http, Provider};
use ethers::signers::Wallet;
use ethers::types::Address;
use ethers::types::Bytes;
use serde::{Deserialize, Serialize};
use std::env::var;
use std::str::FromStr;
use std::sync::Arc;
use std::time::Duration;
use tokio::time::sleep;

const MAX_RETRY_TIMES: u8 = 2;

#[derive(Serialize)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
    pub rpc: String,
    pub shadow: bool,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
    pub blob_kzg: Vec<u8>,
    pub batch_header: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

#[derive(Clone, Debug)]
pub struct ShadowProver {
    l1_shadow_rollup: ShadowRollupType,
    l1_provider: Provider<Http>,
    wallet_address: Address,
}

impl ShadowProver {
    pub async fn prepare() -> Self {
        let l1_rpc = var("SHADOW_PROVING_L1_RPC").expect("Cannot detect L1_RPC env var");
        let l1_shadow_rollup_address = var("SHADOW_PROVING_L1_SHADOW_ROLLUP").expect("Cannot detect L1_SHADOW_ROLLUP env var");
        let _ = var("SHADOW_PROVING_PROVER_RPC").expect("Cannot detect PROVER_RPC env var");
        let private_key = var("SHADOW_PROVING_PRIVATE_KEY").expect("Cannot detect SHADOW_PROVING_PRIVATE_KEY env var");

        let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc).unwrap();
        let l1_signer = Arc::new(SignerMiddleware::new(
            l1_provider.clone(),
            Wallet::from_str(private_key.as_str())
                .unwrap()
                .with_chain_id(l1_provider.get_chainid().await.unwrap().as_u64()),
        ));
        let wallet_address: Address = l1_signer.address();
        let l1_shadow_rollup = ShadowRollup::new(Address::from_str(l1_shadow_rollup_address.as_str()).unwrap(), l1_signer.clone());
        Self {
            l1_shadow_rollup,
            l1_provider,
            wallet_address,
        }
    }

    pub async fn prove(&self, batch_info: BatchInfo) -> Result<(), anyhow::Error> {
        log::info!(">Start shadow prove for batch: {:#?}", batch_info.batch_index);

        // Record wallet balance.
        let balance = match self.l1_provider.get_balance(self.wallet_address, None).await {
            Ok(b) => b,
            Err(e) => {
                log::error!("shadow_proving_wallet.get_balance error: {:#?}", e);
                return Ok(());
            }
        };
        METRICS
            .shadow_wallet_balance
            .set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));

        handle_with_prover(&batch_info, &self.l1_shadow_rollup).await;

        Ok(())
    }
}

async fn handle_with_prover(batch_info: &BatchInfo, l1_shadow_rollup: &ShadowRollupType) {
    let l2_rpc = var("SHADOW_PROVING_L2_RPC").expect("Cannot detect L2_RPC env var");
    let batch_index = batch_info.batch_index;
    let chunks = &batch_info.chunks;
    let chunks_len = chunks.len();

    METRICS.shadow_chunks_len.set(chunks_len as i64);
    METRICS.shadow_batch_index.set(batch_index as i64);

    for _ in 0..MAX_RETRY_TIMES {
        sleep(Duration::from_secs(12)).await;

        log::debug!(
            "Start prove batch of: {:?}, chunks.len = {:?}, chunks = {:#?}",
            batch_index,
            chunks_len,
            chunks
        );

        // Query existing proof
        if let Some(prove_result) = query_proof(batch_index).await {
            if !prove_result.proof_data.is_empty() {
                log::info!("query proof and prove state: {:?}", batch_index);
                prove_state(batch_index, l1_shadow_rollup).await;
                break;
            }
        }

        // Request the proverServer to prove.
        let request = ProveRequest {
            batch_index,
            chunks: chunks.clone(),
            rpc: l2_rpc.to_owned(),
            shadow: true,
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
                    prove_state(batch_index, l1_shadow_rollup).await;
                    break;
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
        let mut max_waiting_time: usize = 4800 * chunks_len + 2400; //chunk_prove_time =1h 20min，batch_prove_time = 24min
        while max_waiting_time > 300 {
            sleep(Duration::from_secs(300)).await;
            max_waiting_time -= 300; // Query results every 5 minutes.
            match query_proof(batch_index).await {
                Some(prove_result) => {
                    log::debug!("query proof and prove state: {:#?}", batch_index);
                    if !prove_result.proof_data.is_empty() {
                        prove_state(batch_index, l1_shadow_rollup).await;
                        return;
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

async fn prove_state(batch_index: u64, l1_rollup: &ShadowRollupType) -> bool {
    for _ in 0..MAX_RETRY_TIMES {
        sleep(Duration::from_secs(12)).await;
        let prove_result = match query_proof(batch_index).await {
            Some(pr) => pr,
            None => continue,
        };

        l1_rollup.address();

        if prove_result.proof_data.is_empty() {
            log::warn!("query proof of {:#?}, proof_data is empty", batch_index);
            continue;
        }

        log::info!(">Starting prove state onchain, batch index = {:#?}", batch_index);
        let aggr_proof = Bytes::from(prove_result.proof_data);
        let kzg_data = Bytes::from(prove_result.blob_kzg);

        let call = l1_rollup.prove_state(batch_index, aggr_proof, kzg_data);
        let rt = call.send().await;
        let pending_tx = match rt {
            Ok(pending_tx) => {
                log::info!("tx of prove_state has been sent: {:#?}", pending_tx.tx_hash());
                pending_tx
            }
            Err(err) => {
                log::error!("send tx of prove_state error: {:#?}", err);
                METRICS.shadow_verify_result.set(2);
                if let ContractError::Revert(data) = err {
                    let msg = String::decode_with_selector(&data).unwrap_or(String::from("unknown, decode contract revert error"));
                    log::error!("send tx of prove_state error, msg: {:#?}", msg);
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
                            log::info!("tx of prove_state success, tx hash: {:?}", receipt.transaction_hash);
                            return true;
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
    false
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

    Some(prove_result)
}
