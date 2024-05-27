use crate::{
    abi::{gas_price_oracle_abi::GasPriceOracle, rollup_abi::Rollup},
    l1_base_fee::BaseFeeUpdater,
    overhead::overhead::OverHeadUpdater,
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
use tokio::runtime::Runtime;

struct Config {
    l1_rpc: String,
    l2_rpc: String,
    gas_threshold: u128,
    overhead_threshold: u128,
    interval: u64,
    overhead_interval: u64,
    l1_rollup_address: Address,
    l2_oracle_address: Address,
    private_key: String,
    l1_beacon_rpc: String,
    overhead_switch: bool,
    max_overhead: u128,
}

impl Config {
    fn new() -> Result<Self, Box<dyn Error>> {
        let _: f64 = read_parse_env("TXN_PER_BLOCK");
        let _: u64 = read_parse_env("TXN_PER_BATCH");

        Ok(Self {
            l1_rpc: var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env"),
            l2_rpc: var("GAS_ORACLE_L2_RPC").expect("GAS_ORACLE_L2_RPC env"),
            gas_threshold: read_parse_env("GAS_THRESHOLD"),
            overhead_threshold: read_parse_env("OVERHEAD_THRESHOLD"),
            interval: read_parse_env("INTERVAL"),
            overhead_interval: read_parse_env("OVERHEAD_INTERVAL"),
            l1_rollup_address: Address::from_str(&var("L1_ROLLUP").expect("L1_ROLLUP env"))?,
            l2_oracle_address: Address::from_str(
                &var("L2_GAS_PRICE_ORACLE").expect("L2_GAS_PRICE_ORACLE env"),
            )?,
            private_key: var("L2_GAS_ORACLE_PRIVATE_KEY").expect("L2_GAS_ORACLE_PRIVATE_KEY env"),
            l1_beacon_rpc: read_parse_env("GAS_ORACLE_L1_BEACON_RPC"),
            overhead_switch: read_env_var("OVERHEAD_SWITCH", false),
            max_overhead: read_env_var("MAX_OVERHEAD", 200000),
        })
    }
}

/// Update data of gasPriceOrale contract on L2 network.
pub async fn update() -> Result<(), Box<dyn Error>> {
    let config = Config::new()?;

    // Prepare signer.
    let l1_provider: Provider<Http> = Provider::<Http>::try_from(config.l1_rpc.clone())?;
    let l2_provider: Provider<Http> = Provider::<Http>::try_from(config.l2_rpc.clone())?;
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str(config.private_key.as_str())
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let l2_wallet_address: Address = l2_signer.address();

    // Prepare contract.
    let l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, _>> =
        GasPriceOracle::new(config.l2_oracle_address, l2_signer);
    let l1_rollup: Rollup<Provider<Http>> =
        Rollup::new(config.l1_rollup_address, Arc::new(l1_provider.clone()));

    // Create baseFeeUpdater
    let base_fee_updater = BaseFeeUpdater::new(
        l1_provider.clone(),
        l2_provider.clone(),
        l2_wallet_address,
        l2_oracle.clone(),
        config.gas_threshold,
    );

    // Create overheadUpdater
    let mut overhead_updater = OverHeadUpdater::new(
        l1_provider.clone(),
        l2_oracle.clone(),
        l1_rollup,
        config.overhead_threshold,
        config.l1_rpc.clone(),
        config.l1_beacon_rpc,
        config.overhead_switch,
        config.max_overhead,
    );

    // Register metrics.
    register_metrics();

    // Update data of gasPriceOrale contract.
    tokio::spawn(async move {
        let mut update_times = 0;
        loop {
            sleep(Duration::from_millis(config.interval)).await;
            update_times += 1;
            base_fee_updater.update().await;

            if update_times < config.overhead_interval {
                // Waiting for the latest batch to be submitted.
                continue;
            }
            // Waiting for confirmation of the previous transaction.
            sleep(Duration::from_millis(8000)).await;
            overhead_updater.update().await;
            update_times = 0
        }
    });

    // Metrics serveice.
    let runtime = Runtime::new().unwrap();
    runtime
        .spawn(async {
            let metrics = Router::new()
                .route("/metrics", get(handle_metrics))
                .layer(TraceLayer::new_for_http());

            axum::Server::bind(&"0.0.0.0:6060".parse().unwrap())
                .serve(metrics.into_make_service())
                .await
                .unwrap();
        })
        .await
        .unwrap();

    Ok(())
}

fn register_metrics() {
    // l1 base fee.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_base_fee.clone())).unwrap();
    // l1 base fee on l2.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.l1_base_fee_on_l2.clone())).unwrap();
    // gas oracle owner balance.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.gas_oracle_owner_balance.clone())).unwrap();
    // gas oracle overhead.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.overhead.clone())).unwrap();
    // scalar ratio.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.scalar_ratio.clone())).unwrap();
    // txn per batch.
    REGISTRY.register(Box::new(ORACLE_SERVICE_METRICS.txn_per_batch.clone())).unwrap();
    // l1 prc.
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

#[tokio::test(flavor = "multi_thread", worker_threads = 4)]
async fn test() -> Result<(), Box<dyn Error>> {
    println!("update");
    let result = update().await;

    match result {
        Ok(()) => Ok(()),
        Err(e) => {
            println!("exec error:");
            Err(e)
        }
    }
}
