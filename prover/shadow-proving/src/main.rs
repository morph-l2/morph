use std::time::Duration;

use axum::{routing::get, Router};
use dotenv::dotenv;
use env_logger::Env;
use prometheus::{Encoder, TextEncoder};
use shadow_proving::metrics::{METRICS, REGISTRY};
use shadow_proving::shadow_prove::ShadowProver;
use shadow_proving::shadow_rollup::BatchSyncer;
use shadow_proving::util::read_env_var;
use tokio::time::sleep;
use tower_http::trace::TraceLayer;

#[tokio::main]
async fn main() {
    // Prepare environment.
    dotenv().ok();
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!("Starting shadow proving...");

    // Start metric management.
    metric_mng().await;

    let batch_syncer = BatchSyncer::prepare().await;
    let shadow_prover = ShadowProver::prepare().await;

    loop {
        sleep(Duration::from_secs(12)).await;
        // Sync & Prove
        let result = match batch_syncer.sync().await {
            Ok(Some(batch)) => shadow_prover.prove(batch).await,
            Ok(None) => Ok(()),
            Err(e) => Err(e),
        };

        // Handle result.
        match result {
            Ok(()) => (),
            Err(e) => {
                log::error!("shadow proving exec error: {:#?}", e);
            }
        }
    }
}

// Metric management
async fn metric_mng() {
    register_metrics();
    let metric_address = read_env_var("SHADOW_PROVING_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
    tokio::spawn(async move {
        let metrics = Router::new().route("/metrics", get(handle_metrics)).layer(TraceLayer::new_for_http());
        axum::Server::bind(&metric_address.parse().unwrap())
            .serve(metrics.into_make_service())
            .await
            .unwrap();
    });
}

fn register_metrics() {
    // detected batch index.
    REGISTRY.register(Box::new(METRICS.detected_batch_index.clone())).unwrap();
    // chunks len.
    REGISTRY.register(Box::new(METRICS.chunks_len.clone())).unwrap();
    // txn len.
    REGISTRY.register(Box::new(METRICS.txn_len.clone())).unwrap();
    // prover status.
    REGISTRY.register(Box::new(METRICS.verify_result.clone())).unwrap();
    // wallet balance.
    REGISTRY.register(Box::new(METRICS.wallet_balance.clone())).unwrap();
}

async fn handle_metrics() -> String {
    let mut buffer = Vec::new();
    let encoder = TextEncoder::new();

    // Gather the metrics.
    let mut metric_families = REGISTRY.gather();
    metric_families.extend(prometheus::gather());

    // Encode metrics to send.
    match encoder.encode(&metric_families, &mut buffer) {
        Ok(()) => {
            let output = String::from_utf8(buffer.clone()).unwrap();
            return output;
        }
        Err(e) => {
            log::error!("encode metrics error: {:#?}", e);
            return String::from("");
        }
    }
}
