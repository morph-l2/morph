use crate::{
    abi::{gas_price_oracle_abi::GasPriceOracle, rollup_abi::Rollup},
    da_scalar::l1_scalar::ScalarUpdater,
    external_sign::ExternalSign,
    l1_base_fee::BaseFeeUpdater,
    read_parse_env,
};
use ethers::{
    prelude::*,
    providers::{Http, Provider},
    signers::Wallet,
    types::Address,
};
use eyre::anyhow;
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
    l1_base_fee_buffer: u64,
    l1_blob_base_fee_buffer: u64,
    commit_scalar_buffer: u64,
    blob_scalar_buffer: u64,
    finalize_batch_gas_used: u64,
    txn_per_batch: u64,
}

impl Config {
    fn new() -> Result<Self, Box<dyn Error>> {
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
            private_key: read_env_var(
                "L2_GAS_ORACLE_PRIVATE_KEY",
                "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80".to_string(),
            ),
            l1_beacon_rpc: read_parse_env("GAS_ORACLE_L1_BEACON_RPC"),
            l1_base_fee_buffer: read_env_var("GAS_ORACLE_L1_BASE_FEE_BUFFER", 0u64),
            l1_blob_base_fee_buffer: read_env_var("GAS_ORACLE_L1_BLOB_BASE_FEE_BUFFER", 0u64),
            commit_scalar_buffer: read_env_var("GAS_ORACLE_COMMIT_SCALAR_BUFFER", 0u64),
            blob_scalar_buffer: read_env_var("GAS_ORACLE_BLOB_SCALAR_BUFFER", 0u64),
            finalize_batch_gas_used: read_env_var("GAS_ORACLE_FINALIZE_BATCH_GAS_USED", 113100u64),
            txn_per_batch: read_env_var("TXN_PER_BATCH", 50),
        })
    }
}

/// Update data of gasPriceOrale contract on L2 network.
pub async fn update() -> Result<(), Box<dyn Error>> {
    let config = Config::new()?;
    check_config(&config)?;

    let (base_fee_updater, scalar_updater) = prepare_updater(&config).await?;

    // Start updater.
    let updater = start_updater(config, base_fee_updater, scalar_updater);

    // Start metric management.
    let metric = metric_mng();

    tokio::join!(updater, metric);
    Ok(())
}

fn check_config(config: &Config) -> Result<(), Box<dyn Error>> {
    log::info!("Check env config, l1_base_fee_buffer: {:?}, l1_blob_base_fee_buffer: {:?}, commit_scalar_buffer: {:?}, blob_scalar_buffer: {:?}",
    config.l1_base_fee_buffer,config.l1_blob_base_fee_buffer, config.commit_scalar_buffer, config.blob_scalar_buffer);

    if config.l1_base_fee_buffer > 100 * 10u64.pow(9) {
        // 1 means 1wei
        return Err(anyhow!(
            "Check env config error, GAS_ORACLE_L1_BASE_FEE_BUFFER should be less than 100 Gwei)"
        )
        .into());
    }
    if config.l1_blob_base_fee_buffer > 100 * 10u64.pow(9) {
        // 1 means 1wei
        return Err(anyhow!(
            "Check env config error, GAS_ORACLE_L1_BLOB_BASE_FEE_BUFFER should be less than 100 Gwei"
        )
        .into());
    }
    if config.commit_scalar_buffer > 10000 * 10u64.pow(9) {
        // 1 means 1Gas
        return Err(anyhow!(
            "Check env config error, GAS_ORACLE_COMMIT_SCALAR_BUFFER should be less than 10000 gas"
        )
        .into());
    }
    if config.blob_scalar_buffer > 10 * 10u64.pow(9) {
        //1*10^9 means an increase of 1 times
        return Err(anyhow!(
            "Check env config error, GAS_ORACLE_BLOB_SCALAR_BUFFER should be less than 10(means 10 times)"
        )
        .into());
    }
    if config.txn_per_batch < 10u64 {
        return Err(anyhow!("Check env config error, TXN_PER_BATCH should be more than 10").into());
    }

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

    let l2_oracle = GasPriceOracle::new(config.l2_oracle_address, l2_signer);
    let l1_rollup = Rollup::new(config.l1_rollup_address, Arc::new(l1_provider.clone()));

    //Signer
    let use_ext_sign: bool = read_env_var("GAS_ORACLE_EXTERNAL_SIGN", false);

    let ext_signer = if use_ext_sign {
        log::info!("Gas Oracle will use external signer");
        let gas_oracle_appid: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_APPID");
        let privkey_pem: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_RSA_PRIV");
        let sign_address: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_ADDRESS");
        let sign_chain: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_CHAIN");
        let sign_url: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_URL");
        let signer: ExternalSign = ExternalSign::new(
            &gas_oracle_appid,
            &privkey_pem,
            &sign_address,
            &sign_chain,
            &sign_url,
        )
        .map_err(|e| anyhow!(format!("Prepare ExternalSign err: {:?}", e)))
        .unwrap();
        Some(signer)
    } else {
        log::info!("Gas Oracle will use local signer");
        None
    };

    let base_fee_updater = BaseFeeUpdater::new(
        l1_provider.clone(),
        l2_provider.clone(),
        ext_signer.clone(),
        l2_oracle.clone(),
        config.gas_threshold,
        config.l1_base_fee_buffer,
        config.l1_blob_base_fee_buffer,
    );

    let scalar_updater = ScalarUpdater::new(
        l1_provider.clone(),
        l2_provider.clone(),
        l2_oracle.clone(),
        ext_signer,
        l1_rollup,
        config.l1_beacon_rpc.clone(),
        config.gas_threshold,
        config.commit_scalar_buffer,
        config.blob_scalar_buffer,
        config.finalize_batch_gas_used,
        config.txn_per_batch,
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
