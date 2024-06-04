use std::time::Duration;

use axum::{routing::get, Router};
use dotenv::dotenv;
use flexi_logger::{Cleanup, Criterion, Duplicate, FileSpec, Logger, Naming, WriteMode};
use log::Record;
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
    // Initialize logger.
    setup_logging();

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
    REGISTRY.register(Box::new(METRICS.shadow_batch_index.clone())).unwrap();
    // chunks len.
    REGISTRY.register(Box::new(METRICS.shadow_chunks_len.clone())).unwrap();
    // txn len.
    REGISTRY.register(Box::new(METRICS.shadow_txn_len.clone())).unwrap();
    // prover status.
    REGISTRY.register(Box::new(METRICS.shadow_verify_result.clone())).unwrap();
    // wallet balance.
    REGISTRY.register(Box::new(METRICS.shadow_wallet_balance.clone())).unwrap();
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

// Constants for configuration
const LOG_LEVEL: &str = "info";
const LOG_FILE_BASENAME: &str = "shadow_proving";
const LOG_FILE_SIZE_LIMIT: u64 = 200 * 10u64.pow(6); // 200MB
                                                     // const LOG_FILE_SIZE_LIMIT: u64 = 10u64.pow(3); // 1kB
const LOG_FILES_TO_KEEP: usize = 3;

fn setup_logging() {
    //configure the logger
    Logger::try_with_env_or_str(LOG_LEVEL)
        .unwrap()
        .log_to_file(
            FileSpec::default()
                .directory(read_env_var("LOG_DIR", String::from("/data/logs/morph-shadow-proving")))
                .basename(LOG_FILE_BASENAME),
        )
        .format(log_format)
        .duplicate_to_stdout(Duplicate::All)
        .rotate(
            Criterion::Size(LOG_FILE_SIZE_LIMIT),     // Scroll when file size reaches 10MB
            Naming::Timestamps,                       // Using timestamps as part of scrolling files
            Cleanup::KeepLogFiles(LOG_FILES_TO_KEEP), // Keep the latest 3 scrolling files
        )
        .write_mode(WriteMode::BufferAndFlush)
        .start()
        .unwrap();
}

fn log_format(w: &mut dyn std::io::Write, now: &mut flexi_logger::DeferredNow, record: &Record) -> Result<(), std::io::Error> {
    write!(
        w,
        "{} [{}] {} - {}",
        now.now().format("%Y-%m-%d %H:%M:%S"), // Custom time format
        record.level(),
        record.target(),
        record.args()
    )
}
