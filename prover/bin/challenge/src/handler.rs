use crate::abi::rollup_abi::Rollup;
use crate::external_sign::ExternalSign;
use crate::metrics::METRICS;
use crate::rollup_compat::{
    authority_is_active, decode_batch_data_input, decode_batch_data_store_block_number, resolve_canonical_commit, CanonicalCommit,
};
use crate::util::read_env_var;
use crate::util::{self, read_parse_env};
use ethers::prelude::*;
use ethers::providers::{Http, Provider};
use ethers::signers::Wallet;
use ethers::types::Address;
use ethers::types::Bytes;
use eyre::anyhow;
use serde::{Deserialize, Serialize};
use std::env::var;
use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use std::time::Duration;
use tokio::time::sleep;
use transaction::eip2718::TypedTransaction;

#[derive(Serialize)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
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

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

const MAX_RETRY_TIMES: u8 = 2;

#[derive(Clone)]
pub struct ChallengeHandler {
    l1_rollup: RollupType,
    l1_provider: Provider<Http>,
    l2_rpc: String,
    ext_signer: Option<ExternalSign>,
    rollup_deployed_block: u64,
}

impl ChallengeHandler {
    pub async fn prepare() -> Self {
        // Prepare parameter.
        let l1_rpc = var("HANDLER_L1_RPC").expect("Cannot detect L1_RPC env var");
        let l2_rpc = var("HANDLER_L2_RPC").expect("Cannot detect L2_RPC env var");
        let l1_rollup_address = var("HANDLER_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
        let _ = var("HANDLER_PROVER_RPC").expect("Cannot detect PROVER_RPC env var");
        let rollup_deployed_block: u64 = read_parse_env("HANDLER_L1_ROLLUP_DEPLOY_BLOCK");

        let private_key = read_env_var(
            "CHALLENGE_HANDLER_PRIVATE_KEY",
            "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80".to_string(),
        );

        let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc).unwrap();
        let l1_signer = Arc::new(SignerMiddleware::new(
            l1_provider.clone(),
            Wallet::from_str(private_key.as_str())
                .unwrap()
                .with_chain_id(l1_provider.get_chainid().await.unwrap().as_u64()),
        ));
        let l1_rollup: RollupType = Rollup::new(Address::from_str(l1_rollup_address.as_str()).unwrap(), l1_signer);

        let use_ext_sign: bool = read_env_var("HANDLER_EXTERNAL_SIGN", false);

        let ext_signer = if use_ext_sign {
            log::info!("Challenge handler will use external signer");
            let handler_appid: String = read_parse_env("HANDLER_EXTERNAL_SIGN_APPID");
            let privkey_pem: String = read_parse_env("HANDLER_EXTERNAL_SIGN_RSA_PRIV");
            let sign_address: String = read_parse_env("HANDLER_EXTERNAL_SIGN_ADDRESS");
            let sign_chain: String = read_parse_env("HANDLER_EXTERNAL_SIGN_CHAIN");
            let sign_url: String = read_parse_env("HANDLER_EXTERNAL_SIGN_URL");
            let signer: ExternalSign = ExternalSign::new(&handler_appid, &privkey_pem, &sign_address, &sign_chain, &sign_url)
                .map_err(|e| anyhow!(format!("Prepare ExternalSign err: {:?}", e)))
                .unwrap();
            Some(signer)
        } else {
            log::info!("Challenge handler will use local signer");
            None
        };

        Self {
            l1_rollup,
            l1_provider,
            l2_rpc,
            ext_signer,
            rollup_deployed_block,
        }
    }

    pub async fn handle_challenge(&self) -> Result<(), Box<dyn Error>> {
        self.handle_with_prover(self.l2_rpc.clone(), &self.l1_provider, &self.l1_rollup).await;
        Ok(())
    }
    async fn handle_with_prover(&self, l2_rpc: String, l1_provider: &Provider<Http>, l1_rollup: &RollupType) {
        loop {
            sleep(Duration::from_secs(12)).await;

            // A challenge win/slash does not terminate the batch. Keep the
            // service in an explicit pending state until the owner executes
            // the canonical RevertBatch and clears revertReqIndex.
            match l1_rollup.revert_req_index().await {
                Ok(index) if !index.is_zero() => {
                    log::warn!("challenge result is awaiting owner RevertBatch, revert_req_index={}", index);
                    continue;
                }
                Ok(_) => {}
                Err(err) => {
                    log::error!("query revertReqIndex failed: {err:#}");
                    continue;
                }
            }

            // Step1. fetch latest blocknum.
            let (latest, snapshot_hash) = match l1_provider.get_block(BlockNumber::Finalized).await {
                Ok(Some(block)) => match (block.number, block.hash) {
                    (Some(number), Some(hash)) => (number, hash),
                    _ => {
                        log::error!("finalized L1 block has no stable identity");
                        continue;
                    }
                },
                Ok(None) => {
                    log::error!("finalized L1 block is unavailable");
                    continue;
                }
                Err(e) => {
                    log::error!("L1 provider.get finalized block error: {:#?}", e);
                    continue;
                }
            };
            log::info!("Current L1 block number: {:#?}", latest);

            let wallet = if let Some(signer) = &self.ext_signer {
                Address::from_str(&signer.address).unwrap_or_default()
            } else {
                self.l1_rollup.client().address()
            };
            // Record wallet balance.
            let balance = match l1_provider.get_balance(wallet, None).await {
                Ok(b) => b,
                Err(e) => {
                    log::error!("handler_wallet.get_balance error: {:#?}", e);
                    return;
                }
            };
            METRICS.wallet_balance.set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));

            // Step2. detect challenge events from the past 3 days.
            let batch_index = match detecte_challenge_event(latest, l1_rollup, l1_provider).await {
                Some(value) => value,
                None => {
                    METRICS.detected_batch_index.set(0i64);
                    continue;
                }
            };
            log::warn!("Challenge event detected, batch index is: {:#?}", batch_index);
            METRICS.detected_batch_index.set(batch_index as i64);

            // Step3. query challenged batch info.
            let (challenged_commit, batch_hash) = match query_batch_tx(latest, l1_rollup, batch_index, self.rollup_deployed_block, l1_provider).await
            {
                Some(value) => value,
                None => continue,
            };

            let mut batch_info = match batch_inspect(l1_rollup, l1_provider, batch_index, &challenged_commit, latest).await {
                Some(mut b) => {
                    b.batch_index = batch_index;
                    b.parent_batch_hash = batch_hash.as_bytes().try_into().unwrap_or_default();
                    b
                }
                None => continue,
            };
            if !snapshot_unchanged(l1_provider, latest, snapshot_hash).await {
                log::error!("finalized L1 snapshot changed while building challenge proof input");
                continue;
            }

            let blocks_len = batch_info.end_block - batch_info.start_block + 1;
            log::info!(
                "batch info: batch index = {:#?}, blocks len = {:#?}, start_block = {:#?}, end_block = {:#?}",
                batch_info.batch_index,
                batch_info.end_block - batch_info.start_block + 1,
                batch_info.start_block,
                batch_info.end_block,
            );
            METRICS.blocks_len.set((batch_info.end_block - batch_info.start_block + 1) as i64);

            if let Some(batch_proof) = query_proof(batch_index).await {
                if !batch_proof.proof_data.is_empty() {
                    log::info!("query proof and prove state: {:#?}", batch_index);
                    let batch_header = batch_info.fill_ext(batch_proof.batch_header.clone()).encode();
                    sleep(Duration::from_secs(600)).await;
                    self.prove_state(batch_index, batch_header, batch_proof, l1_rollup).await;
                    continue;
                }
            }

            // Step4. Make a call to the Prove server.
            let request = ProveRequest {
                batch_index,
                start_block: batch_info.start_block,
                end_block: batch_info.end_block,
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
                        if let Some(batch_proof) = query_proof(batch_index).await {
                            if !batch_proof.proof_data.is_empty() {
                                log::info!("query proof and prove state: {:#?}", batch_index);
                                let batch_header = batch_info.fill_ext(batch_proof.batch_header.clone()).encode();
                                self.prove_state(batch_index, batch_header, batch_proof, l1_rollup).await;
                            }
                        }
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
            let mut max_waiting_time: usize = 1600 * blocks_len as usize; //block_prove_time =30min
            while max_waiting_time > 300 {
                sleep(Duration::from_secs(300)).await;
                max_waiting_time -= 300;
                match query_proof(batch_index).await {
                    Some(batch_proof) => {
                        log::debug!("query proof and prove state: {:#?}", batch_index);
                        if !batch_proof.proof_data.is_empty() {
                            let batch_header = batch_info.fill_ext(batch_proof.batch_header.clone()).encode();
                            self.prove_state(batch_index, batch_header, batch_proof, l1_rollup).await;
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

    async fn prove_state(&self, batch_index: u64, batch_header: Bytes, batch_proof: ProveResult, l1_rollup: &RollupType) -> bool {
        let signer = if let Some(external) = &self.ext_signer {
            Address::from_str(&external.address).unwrap_or_default()
        } else {
            l1_rollup.client().address()
        };
        for _ in 0..MAX_RETRY_TIMES {
            sleep(Duration::from_secs(12)).await;
            if !self.proof_signer_is_active(l1_rollup, signer).await {
                // Re-run the gate before every retry: a threshold increase,
                // exit, or slash between attempts must stop transmission.
                return false;
            }
            log::info!("starting prove state onchain, batch index = {:#?}", batch_index);
            let proof = Bytes::from(batch_proof.proof_data.clone());

            let client: Arc<SignerMiddleware<Provider<Http>, LocalWallet>> = self.l1_rollup.client();
            let calldata = l1_rollup.prove_state(batch_header.clone(), proof).calldata();
            let result = send_transaction(self.l1_rollup.address(), calldata, &client, &self.ext_signer, &self.l1_provider).await;
            if let Ok(tx_hash) = result {
                METRICS.verify_result.set(1);
                log::info!("prove_state success, batch_index: {:?}, tx_hash: {:#?}", batch_index, tx_hash);
                return true;
            }

            if let Err(e) = result {
                METRICS.verify_result.set(2);
                log::error!("send tx of prove_state error, batch_index: {:?}, err_msg: {:#?}", batch_index, e);
                continue;
            }
        }
        false
    }

    async fn proof_signer_is_active(&self, l1_rollup: &RollupType, signer: Address) -> bool {
        let (number, hash) = match self.l1_provider.get_block(BlockNumber::Finalized).await {
            Ok(Some(block)) => match (block.number, block.hash) {
                (Some(number), Some(hash)) => (number, hash),
                _ => {
                    log::error!("finalized L1 block has no stable identity");
                    return false;
                }
            },
            Ok(None) => {
                log::error!("finalized L1 block is unavailable");
                return false;
            }
            Err(err) => {
                log::error!("query finalized L1 snapshot failed: {err:#}");
                return false;
            }
        };
        let snapshot = BlockId::Number(BlockNumber::Number(number));
        let active = match authority_is_active(l1_rollup, signer, snapshot).await {
            Ok(active) => active,
            Err(err) => {
                log::error!("Rollup authority active preflight failed: {err:#}");
                return false;
            }
        };
        if !snapshot_unchanged(&self.l1_provider, number, hash).await {
            log::error!("finalized snapshot changed during authority active preflight");
            return false;
        }
        if !active {
            log::error!("refusing proveState from inactive submitter signer {signer:?}");
        }
        active
    }
}

/**
 * Query the plonk proof for the specified batch index.
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

async fn query_batch_tx(
    latest: U64,
    l1_rollup: &RollupType,
    batch_index: u64,
    rollup_deployed_block: u64,
    l1_provider: &Provider<Http>,
) -> Option<(CanonicalCommit, H256)> {
    let commit = match resolve_canonical_commit(l1_rollup, l1_provider, batch_index, rollup_deployed_block, latest).await {
        Ok(Some(commit)) => commit,
        Ok(None) => {
            log::warn!("batch {batch_index} has no canonical Commit at snapshot {latest}");
            return None;
        }
        Err(err) => {
            log::error!("canonical Commit resolution failed: {err:#}");
            return None;
        }
    };

    let block_id = BlockId::Number(BlockNumber::Number(latest));
    let parent_index = batch_index.checked_sub(1)?;
    let parent_hash: [u8; 32] = match l1_rollup.committed_batches(U256::from(parent_index)).block(block_id).call().await {
        Ok(hash) => hash,
        Err(err) => {
            log::error!("query canonical parent batch hash failed: {err:#}");
            return None;
        }
    };
    Some((commit, H256::from_slice(&parent_hash)))
}

async fn detecte_challenge_event(latest: U64, l1_rollup: &RollupType, l1_provider: &Provider<Http>) -> Option<u64> {
    let start = if latest > U64::from(7200 * 3) {
        // Depends on challenge period
        latest - U64::from(7200 * 3)
    } else {
        U64::from(1)
    };
    let block_id = BlockId::Number(BlockNumber::Number(latest));
    let filter = l1_rollup
        .challenge_state_filter()
        .filter
        .from_block(start)
        .to_block(latest)
        .address(l1_rollup.address());
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
        let batch_in_challenge: bool = match l1_rollup.batch_in_challenge(U256::from(batch_index)).block(block_id).call().await {
            Ok(x) => x,
            Err(e) => {
                log::info!("query l1_rollup.batch_in_challenge error, batch index = {:#?}, {:#?}", batch_index, e);
                return None;
            }
        };
        let is_batch_finalized: bool = match l1_rollup.is_batch_finalized(U256::from(batch_index)).block(block_id).call().await {
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

#[derive(Default, Clone)]
struct BatchInfo {
    version: u8,
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    l1_message_popped: u64,
    total_l1_message_popped: u64,
    data_hash: [u8; 32],
    blob_hashes: Vec<[u8; 32]>,
    prev_state_root: [u8; 32],
    post_state_root: [u8; 32],
    withdrawal_root: [u8; 32],
    sequencer_set_verify_hash: [u8; 32],
    parent_batch_hash: [u8; 32],
}

async fn batch_inspect(
    l1_rollup: &RollupType,
    l1_provider: &Provider<Http>,
    batch_index: u64,
    commit: &CanonicalCommit,
    snapshot: U64,
) -> Option<BatchInfo> {
    let block_id = BlockId::Number(BlockNumber::Number(snapshot));
    let parent_index = batch_index.checked_sub(1)?;
    let calldata = l1_rollup.batch_data_store(U256::from(parent_index)).calldata()?;
    let request: TypedTransaction = TransactionRequest::new().to(l1_rollup.address()).data(calldata).into();
    let prev_batch_last_bn: U256 = match l1_provider.call(&request, Some(block_id)).await {
        Ok(data) => match decode_batch_data_store_block_number(data.as_ref()) {
            Ok(block_number) => block_number,
            Err(error) => {
                log::error!("decode compatible batchDataStore output failed: {error:#}");
                return None;
            }
        },
        Err(e) => {
            log::error!("l1_rollup compatible batchDataStore call failed: {:#?}", e);
            return None;
        }
    };
    //Step1.  Get transaction
    let result = l1_provider.get_transaction(commit.transaction_hash).await;
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
    if tx.to != Some(l1_rollup.address())
        || tx.block_number != Some(commit.identity.block_number.into())
        || tx.block_hash != Some(commit.block_hash)
        || tx.transaction_index != Some(commit.identity.transaction_index.into())
    {
        log::error!("canonical Commit transaction identity changed");
        return None;
    }

    //Step2. Parse transaction data
    let data = tx.input;
    if data.is_empty() {
        log::warn!("batch inspect: tx.input is empty, tx_hash =  {:#?}", commit.transaction_hash);
        return None;
    }
    let param = match decode_batch_data_input(&data) {
        Ok(param) => param,
        Err(err) => {
            log::error!("batch inspect: decode tx.input error, tx_hash = {:?}: {err:#}", commit.transaction_hash);
            return None;
        }
    };

    let version = param.version;
    let prev_state_root = param.prev_state_root;
    let post_state_root = param.post_state_root;
    let withdrawal_root = param.withdrawal_root;
    let last_block_number = param.last_block_number;
    let num_l1_messages = param.num_l1_messages;
    log::info!("======> batch inspect: decode tx.input, version =  {:#?}", version);
    log::info!("======> batch inspect: decode tx.input, param =  {:#?}", param);

    let mut batch_info = BatchInfo {
        version,
        prev_state_root,
        post_state_root,
        withdrawal_root,
        start_block: prev_batch_last_bn.as_u64() + 1,
        end_block: last_block_number,
        l1_message_popped: num_l1_messages as u64,
        ..Default::default()
    };

    // prev_batch_header
    let prev_batch_header = param.parent_batch_header;
    let prev_total_l1_message = prev_batch_header.get(17..25).unwrap_or_default();
    let post_total_l1_message = u64::from_be_bytes(prev_total_l1_message.try_into().unwrap()) + batch_info.l1_message_popped;
    batch_info.total_l1_message_popped = post_total_l1_message;

    Some(batch_info)
}

async fn snapshot_unchanged(provider: &Provider<Http>, number: U64, expected_hash: H256) -> bool {
    matches!(
        provider.get_block(number).await,
        Ok(Some(block)) if block.hash == Some(expected_hash)
    )
}

impl BatchInfo {
    fn fill_ext(&mut self, batch_header_ex: Vec<u8>) -> &Self {
        log::debug!("batch_header_ex len: {:#?}", batch_header_ex.len());

        self.data_hash = batch_header_ex.get(0..32).unwrap_or_default().try_into().unwrap_or_default();
        self.sequencer_set_verify_hash = batch_header_ex.get(64..96).unwrap_or_default().try_into().unwrap_or_default();

        // All versions: single hash at [32..64] (aggregated for V2, versioned for V0/V1)
        let hash: [u8; 32] = batch_header_ex.get(32..64).unwrap_or_default().try_into().unwrap_or_default();
        self.blob_hashes = vec![hash];
        self
    }

    fn encode(&self) -> Bytes {
        let mut batch_header: Vec<u8> = vec![];
        batch_header.extend_from_slice(&self.version.to_be_bytes());
        batch_header.extend_from_slice(&self.batch_index.to_be_bytes());
        batch_header.extend_from_slice(&self.l1_message_popped.to_be_bytes());
        batch_header.extend_from_slice(&self.total_l1_message_popped.to_be_bytes());
        batch_header.extend_from_slice(&self.data_hash);
        // blob hash at offset 57 (same position for all versions; aggregated for V2)
        batch_header.extend_from_slice(self.blob_hashes.first().unwrap_or(&[0u8; 32]));
        batch_header.extend_from_slice(&self.prev_state_root);
        batch_header.extend_from_slice(&self.post_state_root);
        batch_header.extend_from_slice(&self.withdrawal_root);
        batch_header.extend_from_slice(&self.sequencer_set_verify_hash);
        batch_header.extend_from_slice(&self.parent_batch_hash);
        batch_header.extend_from_slice(&self.end_block.to_be_bytes());
        Bytes::from(batch_header)
    }
}

async fn send_transaction(
    contract: Address,
    calldata: Option<Bytes>,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
    l2_provider: &Provider<Http>,
) -> Result<H256, Box<dyn Error>> {
    let req = Eip1559TransactionRequest::new().data(calldata.unwrap_or_default());
    let mut tx = TypedTransaction::Eip1559(req);
    tx.set_to(contract);
    if let Some(signer) = ext_signer {
        tx.set_from(Address::from_str(&signer.address).unwrap_or_default());
    } else {
        tx.set_from(local_signer.address());
    }
    local_signer.fill_transaction(&mut tx, None).await.map_err(|e| {
        let msg = contract_error(ContractError::<SignerMiddleware<Provider<Http>, LocalWallet>>::from_middleware_error(e));
        anyhow!("prove_state fill_transaction error: {:#?}", msg)
    })?;

    let signed_tx = sign_tx(tx, local_signer, ext_signer)
        .await
        .map_err(|e| anyhow!("prove_state sign_tx error: {}", e))?;

    let pending_tx = l2_provider.send_raw_transaction(signed_tx).await.map_err(|e| {
        let msg = contract_error(ContractError::<Provider<Http>>::from(e));
        anyhow!("prove_state call contract error: {}", msg)
    })?;

    let tx_hash = pending_tx.tx_hash();

    let receipt = pending_tx
        .await
        .map_err(|e| anyhow!(format!("prove_state check_receipt of {:#?} is error: {:#?}", tx_hash, e)))?
        .ok_or(anyhow!(format!("prove_state check_receipt is none, tx_hash: {:#?}", tx_hash)))?;

    if receipt.status == Some(1.into()) {
        Ok(tx_hash)
    } else {
        Err(anyhow!(format!("tx of prove_state failed, transaction_hash: {:#?}", receipt.transaction_hash)).into())
    }
}

async fn sign_tx(
    tx: TypedTransaction,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
) -> Result<Bytes, Box<dyn Error>> {
    if let Some(signer) = ext_signer {
        Ok(signer.request_sign(&tx).await?)
    } else {
        let signature = local_signer.signer().sign_transaction(&tx).await?;
        Ok(tx.rlp_signed(&signature))
    }
}

pub fn contract_error<M: Middleware>(e: ContractError<M>) -> String {
    let error_msg = if let Some(contract_err) = e.as_revert() {
        format!("contract error: {:?}", contract_err)
    } else {
        format!("error: {:?}", e)
    };
    error_msg
}
