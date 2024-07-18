use crate::{
    abi::{gas_price_oracle_abi::GasPriceOracle, rollup_abi::Rollup},
    da_scalar::l1_scalar::ScalarUpdater,
    l1_base_fee::BaseFeeUpdater,
    read_parse_env,
};
use ethers::{
    prelude::*,
    providers::{Http, Provider},
    signers::Wallet,
    types::Address,
};
use std::{env::var, error::Error, str::FromStr, sync::Arc, time::Duration};
use tokio::time::sleep;

use axum::{routing::get, Router};
use prometheus::{self, Encoder, TextEncoder};
use tower_http::trace::TraceLayer;

use crate::{metrics::*, read_env_var};

struct Config {
    l1_rpc: String,
    l2_rpc: String,
    gas_threshold: u64,
    interval: u64,
    overhead_interval: u64,
    l1_rollup_address: Address,
    l2_oracle_address: Address,
    private_key: String,
    l1_beacon_rpc: String,
}

impl Config {
    fn new() -> Result<Self, Box<dyn Error>> {
        let _: f64 = read_parse_env("TXN_PER_BLOCK");
        let _: u64 = read_parse_env("TXN_PER_BATCH");

        Ok(Self {
            l1_rpc: var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env"),
            l2_rpc: var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env"),
            gas_threshold: read_parse_env("GAS_THRESHOLD"),
            interval: read_parse_env("INTERVAL"),
            overhead_interval: read_parse_env("OVERHEAD_INTERVAL"),
            l1_rollup_address: Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env"))?,
            l2_oracle_address: Address::from_str(
                &var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env"),
            )?,
            private_key: var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env"),
            l1_beacon_rpc: read_parse_env("GAS_ORACLE_L1_BEACON_RPC"),
        })
    }
}

/// Update data of gasPriceOrale contract on L2 network.
pub async fn update() -> Result<(), Box<dyn Error>> {
    let config = Config::new()?;

    let (base_fee_updater, scalar_updater) = prepare_updater(&config).await?;

    // Start updater.
    let updater = start_updater(config, base_fee_updater, scalar_updater);

    // Start metric management.
    let metric = metric_mng();

    tokio::join!(updater, metric);
    Ok(())
}

async fn start_updater(
    config: Config,
    base_fee_updater: BaseFeeUpdater,
    mut scalar_updater: ScalarUpdater,
) {
    tokio::spawn(async move {
        let mut update_times = 0;
        loop {
            sleep(Duration::from_millis(config.interval)).await;
            update_times += 1;

            let _ = base_fee_updater
                .update()
                .await
                .map_err(|e| log::error!("base_fee_updater err: {:?}", e));

            if update_times < config.overhead_interval {
                // Waiting for the latest batch to be submitted.
                continue;
            }
            // Waiting for confirmation of the previous transaction.
            sleep(Duration::from_millis(6000)).await;
            let _ = scalar_updater
                .update()
                .await
                .map_err(|e| log::error!("overhead_updater err: {:?}", e));
            update_times = 0
        }
    });
}

async fn prepare_updater(
    config: &Config,
) -> Result<(BaseFeeUpdater, ScalarUpdater), Box<dyn Error>> {
    let l1_provider = Provider::<Http>::try_from(config.l1_rpc.clone())?;
    let l2_provider = Provider::<Http>::try_from(config.l2_rpc.clone())?;
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str(config.private_key.as_str())
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));

    let l2_wallet_address = l2_signer.address();
    let l2_oracle = GasPriceOracle::new(config.l2_oracle_address, l2_signer);
    let l1_rollup = Rollup::new(config.l1_rollup_address, Arc::new(l1_provider.clone()));

    let base_fee_updater = BaseFeeUpdater::new(
        l1_provider.clone(),
        l2_provider.clone(),
        l2_wallet_address,
        l2_oracle.clone(),
        config.gas_threshold,
    );

    let scalar_updater = ScalarUpdater::new(
        l1_provider.clone(),
        l2_oracle.clone(),
        l1_rollup,
        config.l1_beacon_rpc.clone(),
        config.gas_threshold,
    );
    Ok((base_fee_updater, scalar_updater))
}

async fn metric_mng() {
    // Register metrics.
    register_metrics();
    let metric_address = read_env_var("GAS_ORACLE_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
    tokio::spawn(async move {
        let metrics =
            Router::new().route("/metrics", get(handle_metrics)).layer(TraceLayer::new_for_http());
        axum::Server::bind(&metric_address.parse().unwrap())
            .serve(metrics.into_make_service())
            .await
            .unwrap();
    })
    .await
    .unwrap();
}

fn register_metrics() {
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_base_fee.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_base_fee_on_l2.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_blob_base_fee_on_l2.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.base_fee_scalar.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.commit_scalar.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.blob_scalar.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.txn_per_batch.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.gas_oracle_owner_balance.clone())).unwrap();
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_rpc_status.clone())).unwrap();
}

async fn handle_metrics() -> String {
    let mut buffer = Vec::new();
    let encoder = TextEncoder::new();

    // Gather the metrics.
    let mut metric_families = REGISTRY.gather();
    metric_families.extend(prometheus::gather());

    // Encode metrics to send.
    match encoder.encode(&metric_families, &mut buffer) {
        Ok(()) => String::from_utf8(buffer.clone()).unwrap(),
        Err(e) => {
            log::error!("encode metrics error: {:#?}", e);
            String::from("")
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test(flavor = "multi_thread", worker_threads = 4)]
    #[ignore]
    async fn test_start_update() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let result = update().await;

        assert!(result.is_ok());
    }
}
