use eyre::anyhow;
use std::{str::FromStr, time::Duration};
use tokio::time::sleep;

use super::{
    calculate::{
        calc_blob_gasprice, calc_tx_overhead, data_and_hashes_from_txs, data_gas_cost,
        extract_tx_payload, extract_txn_num,
    },
    error::OverHeadError,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};
use crate::{
    abi::{
        gas_price_oracle_abi::GasPriceOracle,
        rollup_abi::{CommitBatchCall, Rollup},
    },
    metrics::ORACLE_SERVICE_METRICS,
};
use ethers::{abi::AbiDecode, prelude::*, utils::hex};
use serde_json::Value;

use super::blob_client::{BeaconNode, ExecutionNode};

// Information about the previous rollup
struct PrevRollupInfo {
    gas_used: u128,        // Exec layer gas used
    l2_data_gas_used: u64, // Gas used by L2txn data
    l2_txn: u64,           // Number of L2 transactions
}

// Main struct to manage overhead information
pub struct OverHeadUpdater {
    l1_provider: Provider<Http>, // L1 provider for HTTP connections
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>, /* Oracle for L2
                                  * gas prices */
    l1_rollup: Rollup<Provider<Http>>, // Rollup object for L1
    overhead_threshold: u128,          // Threshold for overhead
    execution_node: ExecutionNode,     // Node for executing transactions
    beacon_node: BeaconNode,           // Beacon node for blockchain
    overhead_switch: bool,             // Flag to enable/disable overhead calculations
    max_overhead: u128,                // Maximum allowable overhead
    prev_rollup_info: Option<PrevRollupInfo>, // Optional info about the previous rollup
}

impl OverHeadUpdater {
    // Constructor to initialize an OverHead object
    #[allow(clippy::too_many_arguments)]
    pub fn new(
        l1_provider: Provider<Http>,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        l1_rollup: Rollup<Provider<Http>>,
        overhead_threshold: u128,
        l1_rpc: String,
        l1_beacon_rpc: String,
        overhead_switch: bool,
        max_overhead: u128,
    ) -> Self {
        // Create execution and beacon nodes with provided RPC URLs
        let execution_node = ExecutionNode { rpc_url: l1_rpc };
        let beacon_node = BeaconNode { rpc_url: l1_beacon_rpc };

        // Return a new OverHead instance with initialized values
        OverHeadUpdater {
            l1_provider,
            l2_oracle,
            l1_rollup,
            overhead_threshold,
            execution_node,
            beacon_node,
            overhead_switch,
            max_overhead,
            prev_rollup_info: None, // Initialize with no previous rollup info
        }
    }

    /// Update overhead
    /// Calculate the user's average cost of the latest rollup and set it to the GasPriceOrale
    /// contract on the L2 network.
    pub async fn update(&mut self) -> Result<(), OverHeadError> {
        // Step1. fetch latest batches and calculate overhead.
        let latest = self.l1_provider.get_block_number().await.map_err(|e| {
            OverHeadError::Error(anyhow!(format!(
                "overhead.l1_provider.get_block_number error: {:#?}",
                e
            )))
        })?;
        let start = if latest > U64::from(100) {
            latest - U64::from(100) //100
        } else {
            U64::from(1)
        };

        let mut latest_overhead = match self.calculate_overhead(start.as_u64()).await {
            Ok(Some(overhead)) => overhead,
            Ok(None) => {
                return Ok(());
            }
            Err(e) => return Err(e),
        };

        // Step2. fetch current overhead on l2.
        let current_overhead = self.l2_oracle.overhead().await.map_err(|e| {
            OverHeadError::Error(anyhow!(format!("query l2_oracle.overhead error: {:#?}", e)))
        })?;
        log::info!("current overhead on l2 is: {:#?}", current_overhead.as_u128());
        ORACLE_SERVICE_METRICS.overhead.set(latest_overhead as i64);

        latest_overhead = latest_overhead.min(self.max_overhead as usize);
        let abs_diff = U256::from(latest_overhead).abs_diff(current_overhead);
        log::info!(
            "overhead actual_change = {:#?}, expected_change =  {:#?}",
            abs_diff,
            self.overhead_threshold
        );

        if abs_diff < U256::from(self.overhead_threshold) {
            log::info!(
                "overhead update abort, diff: {:?} < overhead_threshold: {:?}",
                abs_diff,
                self.overhead_threshold
            );
            return Ok(())
        }

        // Step3. update overhead
        if self.overhead_switch {
            let tx = self.l2_oracle.set_overhead(U256::from(latest_overhead)).legacy();
            let rt = tx.send().await;
            match rt {
                Ok(info) => log::info!("tx of set_overhead has been sent: {:?}", info.tx_hash()),
                Err(e) => log::error!("update overhead error: {:#?}", e),
            }
        }

        Ok(())
    }

    async fn calculate_overhead(&mut self, start: u64) -> Result<Option<usize>, OverHeadError> {
        let filter = self
            .l1_rollup
            .commit_batch_filter()
            .filter
            .from_block(start)
            .address(self.l1_rollup.address());

        let mut logs = self.l1_provider.get_logs(&filter).await.map_err(|e| {
            OverHeadError::Error(anyhow!(format!("overhead.l1_provider.get_logs error: {:#?}", e)))
        })?;

        log::debug!("overhead.l1_provider.submit_batches.get_logs.len ={:#?}", logs.len());

        logs.retain(|x| x.transaction_hash.is_some() && x.block_number.is_some());
        if logs.is_empty() {
            log::warn!("rollup logs for the last 100 blocks of l1 is empty");
            if self.prev_rollup_info.is_some() {
                log::info!("calculate overhead from prev rollup");
                return self.calculate_from_prev_rollup().await;
            }
            log::warn!(
                "rollup logs for the last 100 blocks of l1 is empty and prev_rollup_info is none, skip update"
            );
            return Ok(None);
        }

        let log = logs.iter().max_by_key(|log| log.block_number.unwrap()).ok_or_else(|| {
            OverHeadError::Error(anyhow!(format!(
                "no submit batches logs, start blocknum ={:#?}",
                start
            )))
        })?;

        let latest_overhead = self
            .calculate_from_current_rollup(log.transaction_hash.unwrap(), log.block_number.unwrap())
            .await
            .map_err(|e| {
                log::info!(
                    "last_overhead is none, skip update, tx_hash ={:#?}",
                    log.transaction_hash.unwrap()
                );
                e
            })?;

        Ok(Some(latest_overhead))
    }

    async fn calculate_from_prev_rollup(&mut self) -> Result<Option<usize>, OverHeadError> {
        let latest_block = self.execution_node.query_block_by_num(0).await?;

        let excess_blob_gas =
            U256::from_str(latest_block["result"]["excessBlobGas"].as_str().unwrap_or("0x0"))
                .unwrap_or(U256::from(0));

        let blob_fee = calc_blob_gasprice(excess_blob_gas.as_u64());

        let gas_price = self.l1_provider.get_gas_price().await.unwrap_or_default();

        let prev_rollup = self.prev_rollup_info.as_ref().unwrap();

        //Step4. Calculate overhead
        let x: f64 = blob_fee as f64 / gas_price.as_u128() as f64;
        let overhead = calc_tx_overhead(
            prev_rollup.gas_used,
            MAX_BLOB_TX_PAYLOAD_SIZE as f64,
            x,
            prev_rollup.l2_data_gas_used,
            prev_rollup.l2_txn,
        );

        Ok(Some(overhead as usize))
    }

    async fn calculate_from_current_rollup(
        &mut self,
        tx_hash: TxHash,
        block_num: U64,
    ) -> Result<usize, OverHeadError> {
        //Step1.  Get transaction
        let tx = self
            .l1_provider
            .get_transaction(tx_hash)
            .await
            .map_err(|e| {
                OverHeadError::Error(anyhow!(format!("l1_provider.get_transaction err: {:#?}", e)))
            })?
            .ok_or_else(|| {
                OverHeadError::Error(anyhow!(format!(
                    "ll1_provider.get_transaction is none, tx_hash= {:#?}",
                    tx_hash
                )))
            })?;

        log::info!("rollup tx hash: {:#?}, blocknum: {:#?}", tx_hash, block_num);

        //Step2. Parse transaction data
        let data = tx.input;
        if data.is_empty() {
            return Err(OverHeadError::Error(anyhow!(format!(
                "overhead_inspect tx.input is empty, tx_hash= {:#?}",
                tx_hash
            ))));
        }
        let param = CommitBatchCall::decode(&data).map_err(|e| {
            OverHeadError::Error(anyhow!(format!(
                "overhead_inspect decode tx.input error, tx_hash= {:#?}, err= {:#?}",
                tx_hash, e
            )))
        })?;

        let chunks: Vec<Bytes> = param.batch_data_input.chunks;
        let l2_txn = extract_txn_num(chunks).unwrap_or(0);

        //Step3. Calculate l2 data gas
        let l2_data_gas = self
            .calculate_l2_data_gas_from_blob(tx_hash, tx.block_hash.unwrap(), block_num, l2_txn)
            .await
            .map_err(|e| {
                log::error!("calculate_l2_data_gas_from_blob error: {:#?}", e);
                e
            })?;

        let blob_tx_receipt = self
            .execution_node
            .query_blob_tx_receipt(hex::encode_prefixed(tx_hash).as_str())
            .await?;

        // rollup_gas_used
        let rollup_gas_used =
            U256::from_str(blob_tx_receipt["result"]["gasUsed"].as_str().unwrap_or("0x0"))
                .unwrap_or(U256::from(0));
        log::info!("rollup_calldata_gas_used: {:?}", rollup_gas_used);

        if rollup_gas_used.is_zero() {
            return Err(OverHeadError::Error(anyhow!(format!(
                "blob tx calldata gas_used is None or 0, tx_hash = {:#?}",
                tx_hash
            ))));
        }

        // blob_gas_price
        let blob_gas_price =
            U256::from_str(blob_tx_receipt["result"]["blobGasPrice"].as_str().unwrap_or("0x0"))
                .unwrap_or(U256::from(0));

        // effective_gas_price
        let effective_gas_price = U256::from_str(
            blob_tx_receipt["result"]["effectiveGasPrice"].as_str().unwrap_or("0x0"),
        )
        .unwrap_or(U256::from(0));
        log::info!(
            "blob_gas_price: {:?}, calldata_gas_price: {:?}",
            blob_gas_price,
            effective_gas_price
        );
        if effective_gas_price.is_zero() {
            return Err(OverHeadError::Error(anyhow!(format!(
                "blob tx calldata effective_gas_price is None or 0, tx_hash = {:#?}",
                tx_hash
            ))));
        }

        //Step4. Calculate overhead
        let x: f64 = blob_gas_price.as_u128() as f64 / effective_gas_price.as_u128() as f64;

        let overhead = calc_tx_overhead(
            rollup_gas_used.as_u128(),
            MAX_BLOB_TX_PAYLOAD_SIZE as f64,
            x,
            l2_data_gas,
            l2_txn,
        );

        let prev_rollup = PrevRollupInfo {
            gas_used: rollup_gas_used.as_u128(),
            l2_data_gas_used: l2_data_gas,
            l2_txn,
        };

        self.prev_rollup_info = Some(prev_rollup);

        let blob_fee_ratio = (MAX_BLOB_TX_PAYLOAD_SIZE as f64 * blob_gas_price.as_u128() as f64)
            .ceil() /
            ((rollup_gas_used * effective_gas_price).as_usize() as f64);

        log::info!(
            "lastest_overhead: {:?},  x:{:?}, l2_txn:{:?}, blob gasFee ratio: {:?}",
            overhead,
            x,
            l2_txn,
            format!("{:.5}", blob_fee_ratio)
        );

        // Set metric
        ORACLE_SERVICE_METRICS.txn_per_batch.set(l2_txn as f64);
        Ok(overhead as usize)
    }

    async fn calculate_l2_data_gas_from_blob(
        &self,
        tx_hash: TxHash,
        block_hash: TxHash,
        block_num: U64,
        l2_txn: u64,
    ) -> Result<u64, OverHeadError> {
        if l2_txn == 0 {
            return Ok(0);
        }
        let blob_tx =
            self.execution_node.query_blob_tx(hex::encode_prefixed(tx_hash).as_str()).await?;

        let blob_block =
            self.execution_node.query_block(hex::encode_prefixed(block_hash).as_str()).await?;

        let indexed_hashes = data_and_hashes_from_txs(
            blob_block["result"]["transactions"].as_array().unwrap_or(&Vec::<Value>::new()),
            &blob_tx["result"],
        );

        if indexed_hashes.is_empty() {
            log::info!("No blob in this batch, batchTxHash ={:#?}", tx_hash);
            return Ok(0);
        }

        //Waiting for the next L1 block to be produced.
        sleep(Duration::from_secs(12)).await;
        let next_block = self.execution_node.query_block_by_num((block_num + 1).as_u64()).await?;

        let prev_beacon_root = next_block["result"]["parentBeaconBlockRoot"]
            .as_str()
            .ok_or_else(|| OverHeadError::Error(anyhow!("Next block not produce")))?;

        let indexes: Vec<u64> = indexed_hashes.iter().map(|item| item.index).collect();
        let sidecars_rt =
            self.beacon_node.query_sidecars(prev_beacon_root.to_string(), indexes).await?;

        let sidecars: &Vec<Value> = sidecars_rt["data"]
            .as_array()
            .ok_or_else(|| OverHeadError::Error(anyhow!("query blob_sidecars empty")))?;

        if sidecars.is_empty() {
            return Err(OverHeadError::Error(anyhow!(format!(
                "query sidecars.is_empty: {:?}",
                sidecars_rt
            ))));
        }

        let tx_payload = extract_tx_payload(indexed_hashes, sidecars)?;

        let tx_payload_gas = data_gas_cost(&tx_payload);
        log::info!("sum(tx_data_gas) in blob: {}", tx_payload_gas);

        Ok(tx_payload_gas)
    }
}

#[tokio::test]
#[ignore]
async fn test_calculate_from_current_rollup() {
    use std::{env::var, sync::Arc};

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let rollup_tx_hash = "0x87b09de64fd9c433226a0c683a3b3c1d1e8ab3fa24f3213fa63e2931f205f8d8";
    let rollup_tx_block_num = 1489357;
    log::info!("rollup_tx_block_num: {:#?}", rollup_tx_block_num);

    let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("Cannot detect GAS_ORACLE_L1_RPC env var");
    let l2_rpc = var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env");
    let overhead_threshold =
        var("OVERHEAD_THRESHOLD").expect("OVERHEAD_THRESHOLD env").parse().unwrap();
    let l1_rollup_address = Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env")).unwrap();
    let l2_oracle_address =
        Address::from_str(&var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env")).unwrap();
    let private_key = var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env");

    let l1_provider: Provider<Http> = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();
    let l1_rollup: Rollup<Provider<Http>> =
        Rollup::new(l1_rollup_address, Arc::new(l1_provider.clone()));

    let l2_provider: Provider<Http> = Provider::<Http>::try_from(l2_rpc).unwrap();
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str(private_key.as_str())
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, _>> =
        GasPriceOracle::new(l2_oracle_address, l2_signer);

    let mut overhead: OverHeadUpdater = OverHeadUpdater::new(
        l1_provider,
        l2_oracle,
        l1_rollup,
        overhead_threshold,
        l1_rpc,
        var("GAS_ORACLE_L1_BEACON_RPC")
            .expect("Cannot detect GAS_ORACLE_L1_BEACON_RPC env var")
            .parse()
            .expect("Cannot parse GAS_ORACLE_L1_BEACON_RPC env var"),
        false,
        200000u128,
    );

    let latest_overhead = overhead
        .calculate_from_current_rollup(
            TxHash::from_str(rollup_tx_hash).unwrap(),
            U64::from(rollup_tx_block_num),
        )
        .await;
    println!("latest_overhead: {:?}", latest_overhead);
    return;
}
