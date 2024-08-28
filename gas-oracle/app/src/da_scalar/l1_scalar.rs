use std::sync::Arc;

use eyre::anyhow;
use tokio::time::{sleep, Duration};

use super::{
    blob_client::BeaconNode,
    calculate::{data_and_hashes_from_txs, extract_tx_payload, extract_txn_num},
    error::ScalarError,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};
use crate::{
    abi::{
        gas_price_oracle_abi::GasPriceOracle,
        rollup_abi::{CommitBatchCall, Rollup},
    },
    external_sign::ExternalSign,
    metrics::ORACLE_SERVICE_METRICS,
    signer::send_transaction,
};
use ethers::{abi::AbiDecode, prelude::*, utils::hex};
use serde_json::Value;

const PRECISION: u64 = 10u64.pow(9);
const MAX_COMMIT_SCALAR: u64 = 10u64.pow(9 + 6);
const MAX_BLOB_SCALAR: u64 = 10u64.pow(9 + 2);

// Main struct to manage overhead information
pub struct ScalarUpdater {
    l1_provider: Provider<Http>, // L1 provider for HTTP connections
    l2_provider: Provider<Http>,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>, // L2 gasPrice Oracle
    ext_signer: Option<ExternalSign>,
    l1_rollup: Rollup<Provider<Http>>, // Rollup object for L1
    beacon_node: BeaconNode,           // Beacon node for blockchain
    gas_threshold: u64,
    commit_scalar_buffer: u64,
    blob_scalar_buffer: u64,
    finalize_batch_gas_used: u64,
    txn_per_batch: u64,
}

impl ScalarUpdater {
    // Constructor to initialize an OverHead object
    #[allow(clippy::too_many_arguments)]
    pub fn new(
        l1_provider: Provider<Http>,
        l2_provider: Provider<Http>,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        ext_signer: Option<ExternalSign>,
        l1_rollup: Rollup<Provider<Http>>,
        l1_beacon_rpc: String,
        gas_threshold: u64,
        commit_scalar_buffer: u64,
        blob_scalar_buffer: u64,
        finalize_batch_gas_used: u64,
        txn_per_batch: u64,
    ) -> Self {
        // Create beacon nodes with provided RPC URLs
        let beacon_node = BeaconNode { rpc_url: l1_beacon_rpc };

        // Return a new OverHead instance with initialized values
        Self {
            l1_provider,
            l2_provider,
            l2_oracle,
            ext_signer,
            l1_rollup,
            beacon_node,
            gas_threshold,
            commit_scalar_buffer,
            blob_scalar_buffer,
            finalize_batch_gas_used,
            txn_per_batch,
        }
    }

    /// Update commitScalar and blobScalar.
    /// Calculate the user's average cost of the latest rollup and set it to the GasPriceOrale
    /// contract on the L2 network.
    pub async fn update(&mut self) -> Result<(), ScalarError> {
        // Step1. fetch latest batches and calculate scalar.
        let latest = self.l1_provider.get_block_number().await.map_err(|e| {
            ScalarError::Error(anyhow!(format!(
                "overhead.l1_provider.get_block_number error: {:#?}",
                e
            )))
        })?;
        let start = if latest > U64::from(100) {
            latest - U64::from(100) //100
        } else {
            U64::from(1)
        };

        let (mut commit_scalar, mut blob_scalar) = match self.calculate_scalar(start.as_u64()).await
        {
            Ok(Some(scalar)) => scalar,
            Ok(None) => {
                return Ok(());
            }
            Err(e) => return Err(e),
        };

        // Step2. fetch current scalar on l2.
        let current_commit_scalar: U256 = self.l2_oracle.commit_scalar().await.map_err(|e| {
            ScalarError::Error(anyhow!(format!("query l2_oracle.commit_scalar error: {:#?}", e)))
        })?;

        let current_blob_scalar: U256 = self.l2_oracle.blob_scalar().await.map_err(|e| {
            ScalarError::Error(anyhow!(format!("query l2_oracle.blob_scalar error: {:#?}", e)))
        })?;

        log::info!("set_commit_or_blob_scalar, latest commit_scalar: {:?}, latest blob_scalar: {:?}, current_commit_scalar on l2 is: {:?}, current_blob_scalar on l2 is: {:?}", 
        commit_scalar, blob_scalar, current_commit_scalar.as_u64(), current_blob_scalar.as_u64());

        // Fine tune the actual value
        commit_scalar += self.commit_scalar_buffer;
        blob_scalar += self.blob_scalar_buffer;

        commit_scalar = commit_scalar.min(MAX_COMMIT_SCALAR);
        blob_scalar = blob_scalar.min(MAX_BLOB_SCALAR);

        ORACLE_SERVICE_METRICS.commit_scalar.set((commit_scalar / PRECISION) as i64);
        ORACLE_SERVICE_METRICS
            .blob_scalar
            .set((1000.0 * blob_scalar as f64 / PRECISION as f64).round() / 1000.0);

        // Step3. set on L2chain
        let client: Arc<SignerMiddleware<Provider<Http>, LocalWallet>> = self.l2_oracle.client();
        if self.check_threshold_reached(
            commit_scalar,
            current_commit_scalar.as_u64(),
            "commit_scalar",
        ) {
            // Update commit_scalar
            let calldata = self.l2_oracle.set_commit_scalar(U256::from(commit_scalar)).calldata();
            let tx_hash = send_transaction(
                self.l2_oracle.address(),
                calldata,
                &client,
                &self.ext_signer,
                &self.l2_provider,
            )
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!("set_commit_scalar error: {:#?}", e)))
            })?;
            log::info!("set_commit_scalar success, tx_hash: {:#?}", tx_hash);
        }

        if self.check_threshold_reached(blob_scalar, current_blob_scalar.as_u64(), "blob_scalar") {
            // Update blob_scalar
            let calldata = self.l2_oracle.set_blob_scalar(U256::from(blob_scalar)).calldata();
            let tx_hash = send_transaction(
                self.l2_oracle.address(),
                calldata,
                &client,
                &self.ext_signer,
                &self.l2_provider,
            )
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("set_blob_scalar error: {:#?}", e))))?;
            log::info!("set_blob_scalar success, tx_hash: {:#?}", tx_hash);
        }

        Ok(())
    }

    fn check_threshold_reached(&mut self, latest: u64, current: u64, state_var_name: &str) -> bool {
        let actual_change = latest.abs_diff(current);
        let expected_change = current * self.gas_threshold / 100;
        let need_update = actual_change > expected_change;
        log::info!(
            "update {}, actual_change: {:?}, expected_change: {:?}, need_update: {:?}",
            state_var_name,
            actual_change,
            expected_change,
            need_update
        );
        need_update
    }

    async fn calculate_scalar(&mut self, start: u64) -> Result<Option<(u64, u64)>, ScalarError> {
        let filter = self
            .l1_rollup
            .commit_batch_filter()
            .filter
            .from_block(start)
            .address(self.l1_rollup.address());

        let mut logs = self.l1_provider.get_logs(&filter).await.map_err(|e| {
            ScalarError::Error(anyhow!(format!("overhead.l1_provider.get_logs error: {:#?}", e)))
        })?;

        log::debug!("overhead.l1_provider.submit_batches.get_logs.len ={:#?}", logs.len());

        logs.retain(|x| x.transaction_hash.is_some() && x.block_number.is_some());
        if logs.is_empty() {
            log::warn!("rollup logs for the last 100 blocks of l1 is empty, skip update");
            return Ok(None);
        }

        let log = logs.iter().max_by_key(|log| log.block_number.unwrap()).ok_or_else(|| {
            ScalarError::Error(anyhow!(format!(
                "no submit batches logs, start blocknum ={:#?}",
                start
            )))
        })?;

        #[allow(clippy::manual_inspect)]
        let (commit_scalar, blob_scalar) = self
            .calculate_from_rollup(log.transaction_hash.unwrap(), log.block_number.unwrap())
            .await
            .map_err(|e| {
                log::info!(
                    "scalar is none, skip update, tx_hash ={:#?}",
                    log.transaction_hash.unwrap()
                );
                e
            })?;

        Ok(Some((commit_scalar, blob_scalar)))
    }

    async fn calculate_from_rollup(
        &mut self,
        tx_hash: TxHash,
        block_num: U64,
    ) -> Result<(u64, u64), ScalarError> {
        //Step1.  Get transaction
        let tx = self
            .l1_provider
            .get_transaction(tx_hash)
            .await
            .map_err(|e| {
                ScalarError::Error(anyhow!(format!("l1_provider.get_transaction err: {:#?}", e)))
            })?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "ll1_provider.get_transaction is none, tx_hash= {:#?}",
                    tx_hash
                )))
            })?;

        log::info!("hit self rollup tx hash: {:#?}, blocknum: {:#?}", tx_hash, block_num);

        //Step2. Parse transaction data
        let data = tx.input;
        if data.is_empty() {
            return Err(ScalarError::Error(anyhow!(format!(
                "overhead_inspect tx.input is empty, tx_hash= {:#?}",
                tx_hash
            ))));
        }
        let param = CommitBatchCall::decode(&data).map_err(|e| {
            ScalarError::Error(anyhow!(format!(
                "overhead_inspect decode tx.input error, tx_hash= {:#?}, err= {:#?}",
                tx_hash, e
            )))
        })?;

        let chunks: Vec<Bytes> = param.batch_data_input.chunks;
        let l2_txn = extract_txn_num(chunks).unwrap_or(0);

        //Step3. Calculate l2 data gas
        let l2_data_len = self
            .calculate_l2_data_len_from_blob(tx_hash, block_num, l2_txn)
            .await
            .map_err(|e| {
                log::error!("calculate_l2_data_len_from_blob error: {:#?}", e);
                e
            })?;

        let blob_tx_receipt = self
            .l1_provider
            .get_transaction_receipt(tx_hash)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get transaction receipt return none, tx_hash= {:#?}",
                    tx_hash
                )))
            })?;

        // rollup_gas_used
        let rollup_gas_used = blob_tx_receipt.gas_used.unwrap_or_default();
        if rollup_gas_used.is_zero() {
            return Err(ScalarError::Error(anyhow!(format!(
                "blob tx calldata gas_used is none or 0, tx_hash = {:#?}",
                tx_hash
            ))));
        }

        //Step4. Calculate scalar
        let commit_scalar = (rollup_gas_used.as_u64() + self.finalize_batch_gas_used) * PRECISION /
            l2_txn.max(self.txn_per_batch);
        let blob_scalar = if l2_data_len > 0 {
            MAX_BLOB_TX_PAYLOAD_SIZE as u64 * PRECISION / l2_data_len
        } else {
            MAX_BLOB_SCALAR
        };

        log::info!(
            "rollup_gas_used: {:?}, l2_txn: {:?}, l2_data_len:{:?}, commit_scalar: {:?}, blob_scalar: {:.4}",
            rollup_gas_used,
            l2_txn,
            l2_data_len,
            commit_scalar/PRECISION,
            blob_scalar as f64/PRECISION as f64,
        );

        // Set metric
        ORACLE_SERVICE_METRICS.txn_per_batch.set(l2_txn as f64);
        Ok((commit_scalar, blob_scalar))
    }

    async fn calculate_l2_data_len_from_blob(
        &self,
        tx_hash: TxHash,
        block_num: U64,
        l2_txn: u64,
    ) -> Result<u64, ScalarError> {
        if l2_txn == 0 {
            return Ok(0);
        }
        let blob_tx = self
            .l1_provider
            .get_transaction(tx_hash)
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get transaction return none, tx_hash: {:#?}",
                    tx_hash
                )))
            })?;

        let blob_block = self
            .l1_provider
            .get_block_with_txs(BlockNumber::Number(block_num))
            .await
            .map_err(|e| ScalarError::Error(anyhow!(format!("{:#?}", e))))?
            .ok_or_else(|| {
                ScalarError::Error(anyhow!(format!(
                    "l1 get block info return none, block_num: {:#?}",
                    block_num
                )))
            })?;

        let indexed_hashes = data_and_hashes_from_txs(&blob_block.transactions, &blob_tx);
        if indexed_hashes.is_empty() {
            log::info!("no blob in this batch, batch_tx_hash: {:#?}", tx_hash);
            return Ok(0);
        }

        // Waiting for the next L1 block to be produced.
        let next_block_num = block_num + 1;
        // Max delay 5 * 3 = 15 secs
        let mut retry_times = 5;
        let prev_beacon_root = loop {
            let blk_info = self.l1_provider.get_block(BlockNumber::Number(next_block_num)).await;
            if let Ok(Some(info)) = blk_info {
                if let Some(beacon_blk_root) = info.parent_beacon_block_root {
                    break beacon_blk_root;
                } else {
                    return Err(ScalarError::Error(anyhow!(format!(
                        "next block info's pre_beacon_root is none, block number: {:?}",
                        next_block_num
                    ))));
                }
            } else if retry_times > 0 {
                retry_times -= 1;
                sleep(Duration::from_secs(3)).await;

                log::info!(
                    "request next block info, retry times= {:?}, block number: {:?}",
                    retry_times,
                    next_block_num
                );
                continue;
            } else {
                return Err(ScalarError::Error(anyhow!(format!(
                    "maximum number of requests next block info reached: {:?}, block number:{:?}",
                    retry_times, next_block_num
                ))));
            }
        };

        let indexes: Vec<u64> = indexed_hashes.iter().map(|item| item.index).collect();
        let sidecars_rt = self
            .beacon_node
            .query_sidecars(hex::encode_prefixed(prev_beacon_root), indexes)
            .await?;

        let sidecars: &Vec<Value> = sidecars_rt["data"].as_array().ok_or_else(|| {
            ScalarError::Error(anyhow!(format!(
                "blob_sidecars is none, blk_num: {:?}, blk_root: {:?}",
                block_num, prev_beacon_root
            )))
        })?;

        if sidecars.is_empty() {
            return Err(ScalarError::Error(anyhow!(format!(
                "blob_sidecars is empty, blk_num: {:?}, blk_root: {:?}",
                block_num, prev_beacon_root
            ))));
        }

        let tx_payload = extract_tx_payload(indexed_hashes, sidecars)?;

        Ok(tx_payload.len() as u64)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::{env::var, str::FromStr, sync::Arc};

    #[tokio::test]
    #[ignore]
    async fn test_calculate_from_current_rollup() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let rollup_tx_hash = "0x87b09de64fd9c433226a0c683a3b3c1d1e8ab3fa24f3213fa63e2931f205f8d8"
            .parse::<H256>()
            .unwrap();
        let rollup_tx_block_num = U64::from(1489357);

        let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env empty");
        let l2_rpc = var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env empty");
        let gas_threshold = var("GAS_THRESHOLD").expect("GAS_THRESHOLD env empty").parse().unwrap();
        let l1_rollup_address =
            Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env empty")).unwrap();
        let l2_oracle_address =
            Address::from_str(&var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env empty"))
                .unwrap();
        let private_key =
            var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env empty");

        let l1_provider = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();
        let l1_rollup_contract = Rollup::new(l1_rollup_address, Arc::new(l1_provider.clone()));

        let l2_provider = Provider::<Http>::try_from(l2_rpc).unwrap();
        let l2_signer = Arc::new(SignerMiddleware::new(
            l2_provider.clone(),
            Wallet::from_str(private_key.as_str())
                .unwrap()
                .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
        ));

        let l2_oracle_contract = GasPriceOracle::new(l2_oracle_address, l2_signer);

        let ext_signer: ExternalSign =
            ExternalSign::new("appid", "privkey_pem", "address", "chain", "url").unwrap();
        let mut overhead: ScalarUpdater = ScalarUpdater::new(
            l1_provider,
            l2_provider,
            l2_oracle_contract,
            Some(ext_signer),
            l1_rollup_contract,
            var("GAS_ORACLE_L1_BEACON_RPC")
                .expect("Cannot detect GAS_ORACLE_L1_BEACON_RPC env empty")
                .parse()
                .expect("Cannot parse GAS_ORACLE_L1_BEACON_RPC env var empty"),
            gas_threshold,
            0u64,
            0u64,
            0u64,
            50u64,
        );

        let latest_overhead =
            overhead.calculate_from_rollup(rollup_tx_hash, rollup_tx_block_num).await;

        log::info!("latest_overhead: {:#?}", latest_overhead);
    }
}
