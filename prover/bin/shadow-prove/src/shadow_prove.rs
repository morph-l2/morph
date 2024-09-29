use crate::{metrics::METRICS, util, BatchInfo, ShadowRollup::ShadowRollupInstance};
use alloy::{
    network::{Network, ReceiptResponse},
    primitives::{Address, Bytes},
    providers::{Provider, RootProvider},
    transports::{
        http::{Client, Http},
        Transport,
    },
};
use serde::{Deserialize, Serialize};
use std::{env::var, time::Duration};
use tokio::time::sleep;

const MAX_RETRY_TIMES: u8 = 2;

#[derive(Serialize)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
    pub shadow: bool,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
    pub batch_header: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

#[derive(Clone, Debug)]
pub struct ShadowProver<T, P, N> {
    l1_provider: RootProvider<Http<Client>>,
    l1_shadow_rollup: ShadowRollupInstance<T, P, N>,
    wallet_address: Address,
}

impl<T, P, N> ShadowProver<T, P, N>
where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
    pub fn new(
        wallet_address: Address,
        shadow_rollup_address: Address,
        provider: RootProvider<Http<Client>>,
        wallet: P,
    ) -> Self {
        let l1_shadow_rollup = ShadowRollupInstance::new(shadow_rollup_address, wallet);

        Self { l1_provider: provider, l1_shadow_rollup, wallet_address }
    }

    pub async fn prove(&self, batch_info: BatchInfo) -> Result<(), anyhow::Error> {
        log::info!(">Start shadow prove for batch: {:#?}", batch_info.batch_index);

        // Record wallet balance.
        let balance = match self.l1_provider.get_balance(self.wallet_address).await {
            Ok(b) => b,
            Err(e) => {
                log::error!("shadow_proving_wallet.get_balance error: {:#?}", e);
                return Ok(());
            }
        };
        METRICS
            .shadow_wallet_balance
            .set(alloy::primitives::utils::format_ether(balance).parse().unwrap_or(0.0));

        handle_with_prover(&batch_info, &self.l1_shadow_rollup).await;

        Ok(())
    }
}

async fn handle_with_prover<T, P, N>(
    batch_info: &BatchInfo,
    l1_shadow_rollup: &ShadowRollupInstance<T, P, N>,
) where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
    let l2_rpc = var("SHADOW_PROVING_L2_RPC").expect("Cannot detect L2_RPC env var");
    let batch_index = batch_info.batch_index;
    let blocks = &batch_info.blocks;
    let blocks_len = blocks.len();

    METRICS.shadow_blocks_len.set(blocks_len as i64);
    METRICS.shadow_batch_index.set(batch_index as i64);

    for _ in 0..MAX_RETRY_TIMES {
        sleep(Duration::from_secs(12)).await;

        log::info!(
            "Start prove batch of: {:?}, blocks.len = {:?}, block_start = {:#?}",
            batch_index,
            blocks_len,
            blocks[0]
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
            start_block: *blocks.first().unwrap_or(&0u64),
            end_block: *blocks.last().unwrap_or(&0u64),
            rpc: l2_rpc.to_owned(),
            shadow: true,
        };
        let rt = tokio::task::spawn_blocking(move || {
            util::call_prover(serde_json::to_string(&request).unwrap(), "/prove_batch")
        })
        .await
        .unwrap();

        match rt {
            Some(info) => match info.as_str() {
                task_status::STARTED => log::info!(
                    "successfully submitted prove task, waiting for proof to be generated"
                ),
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
        let mut max_waiting_time: usize = 1600 * blocks_len; //block_prove_time =30min
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

async fn prove_state<T, P, N>(
    batch_index: u64,
    shadow_rollup: &ShadowRollupInstance<T, P, N>,
) -> bool
where
    P: Provider<T, N> + Clone,
    T: Transport + Clone,
    N: Network,
{
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
        log::info!("proof_data {:#?}", prove_result.proof_data);

        log::info!(">Starting prove state onchain, batch index = {:#?}", batch_index);
        let aggr_proof = Bytes::from(prove_result.proof_data);
        let shadow_tx = shadow_rollup.proveState(batch_index, aggr_proof);
        let send = shadow_tx.send().await;

        let pending_tx = match send {
            Ok(pending_tx) => pending_tx,
            Err(e) => {
                log::error!("send tx of prove_state error: {:#?}", e);
                METRICS.shadow_verify_result.set(2);
                continue;
            }
        };
        let receipt = pending_tx.get_receipt().await.unwrap();
        if receipt.status() {
            log::info!("tx of prove_state success, tx hash: {:?}", receipt.transaction_hash());
            return true;
        }
        log::error!("tx of prove_state failed, tx hash: {:?}", receipt.transaction_hash());
    }
    false
}

/**
 * Query the zkevm proof for the specified batch index.
 * Only return result when proof data exists, otherwise return None.
 */
async fn query_proof(batch_index: u64) -> Option<ProveResult> {
    // Make a call to the Prove server.
    let rt = tokio::task::spawn_blocking(move || {
        util::call_prover(batch_index.to_string(), "/query_proof")
    })
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

#[test]
fn test() {
    use crate::abi::SP1Verifier;
    use alloy::rpc::json_rpc::ErrorPayload;
    // Sample JSON error payload from an Ethereum JSON RPC response.
    let json = r#"{"code":3,"message":"execution reverted: ","data":"0x810f00230000000000000000000000000000000000000000000000000000000000000001"}"#;

    // Parse the JSON into an `ErrorPayload` struct.
    let payload: ErrorPayload = serde_json::from_str(json).unwrap();

    let err = payload.as_decoded_error::<SP1Verifier::SP1VerifierErrors>(false).unwrap();
    match err {
        SP1Verifier::SP1VerifierErrors::WrongVerifierSelector(s) => {
            println!(
                "WrongVerifierSelector -  expected: {:?}, received: {:?}",
                s.expected, s.received
            );
        }
        SP1Verifier::SP1VerifierErrors::InvalidProof(_) => println!("WrongVerifierSelector"),
    }
}
