use challenge_handler::abi::rollup_abi::Rollup;
use dotenv::dotenv;
use env_logger::Env;
use ethers::prelude::*;
use ethers::signers::Wallet;
use std::env::var;
use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use std::time::Duration;

#[tokio::main]
async fn main() {
    // Prepare environment.
    dotenv().ok();
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!("Starting challenge handler...");

    // Start challenger.
    let result = challenge().await;

    // Handle result.
    match result {
        Ok(()) => (),
        Err(e) => {
            log::error!("challenge handler exec error: {:#?}", e);
        }
    }
}

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;

/**
 * Automatically challenge the latest batch.
 */
pub async fn challenge() -> Result<(), Box<dyn Error>> {
    // Prepare env.
    log::info!("Starting shadow-challenge...");
    let l1_rpc = var("CHALLENGER_L1_RPC").expect("Cannot detect L1_RPC env var");
    let l1_rollup_address = var("CHALLENGER_L1_ROLLUP").expect("Cannot detect L1_ROLLUP env var");
    let private_key = var("CHALLENGER_PRIVATEKEY").expect("Cannot detect CHALLENGER_PRIVATEKEY env var");
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
    log::info!("address({:#?})  is_challenger: {:#?}", challenger_address.to_string(), is_challenger);

    let challenger_balance = l1_provider.get_balance(challenger_address, None).await.unwrap();
    log::info!("challenger_eth_balance: {:#?}", ethers::utils::format_ether(challenger_balance));

    let finalization_period = l1_rollup.finalization_period_seconds().await?;
    let proof_window = l1_rollup.proof_window().await?;
    log::info!("finalization_period: {:#?}  proof_window: {:#?}", finalization_period, proof_window);

    let min_deposit: U256 = U256::from(4 * 10u128.pow(18));
    log::info!("min_deposit: {:#?}", min_deposit);

    loop {
        std::thread::sleep(Duration::from_secs(12));
        let _ = auto_challenge(&l1_provider, &l1_rollup, min_deposit).await;
    }
}

async fn auto_challenge(l1_provider: &Provider<Http>, l1_rollup: &RollupType, min_deposit: U256) -> Result<(), Box<dyn Error>> {
    // Search for the latest batch.
    let latest = match l1_provider.get_block_number().await {
        Ok(bn) => bn,
        Err(e) => {
            log::error!("L1 provider.get_block_number error: {:#?}", e);
            return Ok(());
        }
    };

    // Check layer2 state.
    verify_state_transition().await;

    // Check prev challenge.
    match detecte_challenge(latest, l1_rollup, l1_provider).await {
        Some(true) => {
            return Ok(());
        }
        Some(false) => (),
        None => {
            log::warn!("prev challenge unknown");
            return Ok(());
        }
    }

    log::info!("latest blocknum = {:#?}", latest);
    let start = if latest > U64::from(600) {
        latest - U64::from(600)
    } else {
        U64::from(1)
    };

    let filter = l1_rollup.commit_batch_filter().filter.from_block(start).address(l1_rollup.address());
    let mut logs: Vec<Log> = match l1_provider.get_logs(&filter).await {
        Ok(logs) => logs,
        Err(e) => {
            log::error!("l1_rollup.commit_batch.get_logs error: {:#?}", e);
            return Ok(());
        }
    };

    if logs.is_empty() {
        log::error!("There have been no commit_batch logs for the last 600 blocks.");
        return Ok(());
    }

    logs.sort_by(|a, b| a.block_number.unwrap().cmp(&b.block_number.unwrap()));
    let batch_index = match logs.get(logs.len() - 2) {
        Some(log) => log.topics[1].to_low_u64_be(),
        None => {
            log::error!("find commit_batch log error");
            return Ok(());
        }
    };
    log::info!("latest batch index = {:#?}", batch_index);

    // Challenge state.
    let is_batch_finalized = l1_rollup.is_batch_finalized(U256::from(batch_index)).await?;
    if is_batch_finalized {
        log::info!("is_batch_finalized = true, No need for challenge, batch index = {:#?}", batch_index);
        return Ok(());
    }

    let challenges = match l1_rollup.challenges(U256::from(batch_index)).await {
        Ok(x) => x,
        Err(e) => {
            log::info!("query l1_rollup.challenges error, batch index = {:#?}, {:#?}", batch_index, e);
            return Ok(());
        }
    };

    if challenges.1 != Address::default() {
        log::info!("already challenge, batch index = {:#?}", batch_index);
        return Ok(());
    }

    // l1_rollup.connect()
    let tx: FunctionCall<_, _, _> = l1_rollup.challenge_state(batch_index).value(min_deposit);
    let rt = tx.send().await;
    let pending_tx = match rt {
        Ok(pending_tx) => {
            log::info!("tx of challenge_state has been sent: {:#?}", pending_tx.tx_hash());
            pending_tx
        }
        Err(e) => {
            log::error!("send tx of challenge_state error hex: {:#?}", e);
            if let ContractError::Revert(data) = e {
                let msg = String::decode_with_selector(&data).unwrap_or(String::from("decode contract revert error"));
                log::error!("send tx of challenge_state error msg: {:#?}", msg);
            }
            return Ok(());
        }
    };

    match pending_tx.await {
        Ok(receipt) => {
            match receipt {
                Some(receipt) => {
                    // Check the status of the tx receipt
                    if receipt.status == Some(1.into()) {
                        log::info!(
                            "tx of challenge_state success, batch_index: {:?}, gasUsed: {:?}, txHash: {:?}",
                            batch_index,
                            receipt.gas_used,
                            receipt.transaction_hash
                        );
                    } else {
                        log::error!(
                            "tx of challenge_state failed, batch_index: {:?}, txHash: {:?}",
                            batch_index,
                            receipt.transaction_hash
                        );
                    }
                }
                None => {
                    log::error!("No challenge_state tx receipt found, may still be in pending status or has been dropped");
                }
            }
        }
        Err(error) => log::error!("provider error: {:?}", error),
    }

    Ok(())
}

async fn detecte_challenge(latest: U64, l1_rollup: &RollupType, l1_provider: &Provider<Http>) -> Option<bool> {
    let start = if latest > U64::from(7200) {
        // Depends on proof window
        latest - U64::from(7200)
        // U64::from(1)
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
        return Some(false);
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
        let _is_batch_finalized: bool = l1_rollup.is_batch_finalized(U256::from(batch_index)).await.unwrap();

        if batch_in_challenge {
            log::info!("prev challenge not finalized, batch index = {:#?}", batch_index);
            return Some(true);
        }
        log::debug!("batch status not in challenge, batch index = {:#?}", batch_index);
    }
    log::info!("all batch's status not in challenge now");
    Some(false)
}

// Check layer2 state.
async fn verify_state_transition() {
    // Do nothing
}
