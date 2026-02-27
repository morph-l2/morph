use std::{str::FromStr, time::Duration};

use alloy_network::{Ethereum, EthereumWallet};
use alloy_primitives::Address;
use alloy_provider::{DynProvider, Provider, ProviderBuilder};
use alloy_signer_local::PrivateKeySigner;
use axum::{routing::get, Router};
use dotenv::dotenv;
use flexi_logger::{Cleanup, Criterion, Duplicate, FileSpec, Logger, Naming, WriteMode};
use log::Record;
use prometheus::{Encoder, TextEncoder};
use shadow_proving::{
    execute::try_execute_batch,
    metrics::{METRICS, REGISTRY},
    shadow_prove::{BatchProveInfo, ShadowProver},
    shadow_rollup::BatchSyncer,
    util::{read_env_var, read_parse_env},
    SHADOW_EXECUTE, SHADOW_PROVING_MAX_BLOCK, SHADOW_PROVING_MAX_TXN, SHADOW_PROVING_PROVER_RPC,
};

use tokio::time::interval;
use tokio::{sync::broadcast, time::MissedTickBehavior};

#[tokio::main]
async fn main() {
    // Prepare environment.
    dotenv().ok();
    setup_logging();
    log::info!("Starting shadow proving...");
    log::info!("Loading with env SHADOW_PROVING_MAX_BLOCK: {}", *SHADOW_PROVING_MAX_BLOCK);
    log::info!("Loading with env SHADOW_PROVING_MAX_TXN: {}", *SHADOW_PROVING_MAX_TXN);
    log::info!("Loading with env SHADOW_PROVING_PROVER_RPC: {}", *SHADOW_PROVING_PROVER_RPC);

    // Start metric management.
    metric_mng().await;

    let (batch_syncer, shadow_prover, l2_provider) = init_shadow_proving();
    let chain_id = l2_provider.get_chain_id().await.unwrap_or_default();

    let (tx, mut rx) = broadcast::channel(4);
    let batch_syncer_exec = batch_syncer.clone();

    // Spawn batch fetch and execute task.
    tokio::spawn(async move {
        // Track the latest processed batch index
        let mut latest_processed_batch: u64 = 0;
        let mut ticker = interval(Duration::from_secs(60));
        ticker.set_missed_tick_behavior(MissedTickBehavior::Skip);
        loop {
            ticker.tick().await;

            // Get committed batch
            let (batch_info, batch_header) =
                match batch_syncer_exec.get_latest_batch(latest_processed_batch).await {
                    Ok(Some(committed_batch)) => committed_batch,
                    Ok(None) => continue,
                    Err(e) => {
                        log::error!("get_committed_batch error: {:?}", e);
                        continue;
                    }
                };

            // Check if batch has already been processed
            if batch_info.batch_index <= latest_processed_batch {
                log::info!("Batch {} has already been processed, skipping", batch_info.batch_index);
                continue;
            }
            latest_processed_batch = batch_info.batch_index;

            // Shadow execute checks
            if *SHADOW_EXECUTE {
                log::info!(">Start shadow execute batch: {:#?}", batch_info.batch_index);
                // Execute batch
                let offchain_batch_pi = match try_execute_batch(&batch_info, &l2_provider).await {
                    Ok(pi) => {
                        // Update the latest processed batch index
                        pi
                    }
                    Err(e) => {
                        log::error!("execute_batch error: {:?}", e);
                        continue;
                    }
                };
                let onchain_batch_pi =
                    batch_syncer_exec.calc_batch_pi(chain_id, &batch_header).unwrap_or_default();
                if offchain_batch_pi != onchain_batch_pi {
                    log::error!(
                        "Shadow execute batch pi mismatch! offchain: {:?}, onchain: {:?}",
                        offchain_batch_pi,
                        onchain_batch_pi
                    );
                    continue;
                }
            }

            // Sync & Prove checks
            if batch_info.end_block - batch_info.start_block + 1 > *SHADOW_PROVING_MAX_BLOCK {
                log::warn!("Too many blocks in the latest batch to shadow prove");
                continue;
            }

            if batch_info.total_txn > *SHADOW_PROVING_MAX_TXN {
                log::warn!("Too many txn in the latest batch to shadow prove");
                continue;
            }

            if let Err(e) = tx.send((batch_info, batch_header)) {
                log::error!("Failed to send batch to prove queue: {:?}", e);
            }
        }
    });

    // Start prove worker loop.
    loop {
        let (batch_info, batch_header) = match rx.recv().await {
            Ok(v) => v,
            Err(broadcast::error::RecvError::Lagged(skipped)) => {
                log::warn!("Prove thread lagged, skipped {} batches", skipped);
                continue;
            }
            Err(broadcast::error::RecvError::Closed) => break,
        };

        let result = match batch_syncer.sync_batch(batch_info, batch_header.clone()).await {
            Ok(Some(batch)) => {
                let prove_info = BatchProveInfo { batch_info: batch, batch_header };
                shadow_prover.prove(prove_info).await
            }
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

fn init_shadow_proving(
) -> (BatchSyncer<DynProvider, Ethereum>, ShadowProver<DynProvider, Ethereum>, DynProvider) {
    let l1_verify_rpc: String = read_parse_env("SHADOW_PROVING_VERIFY_L1_RPC");
    let l1_rpc: String = read_parse_env("SHADOW_PROVING_L1_RPC");
    let l2_rpc: String = read_parse_env("SHADOW_PROVING_L2_RPC");

    let private_key: String = read_parse_env("SHADOW_PROVING_PRIVATE_KEY");
    let rollup: String = read_parse_env("SHADOW_PROVING_L1_ROLLUP");
    let shadow_rollup: String = read_parse_env("SHADOW_PROVING_L1_SHADOW_ROLLUP");

    let signer: PrivateKeySigner = private_key.parse().expect("parse PrivateKeySigner");
    let wallet: EthereumWallet = EthereumWallet::from(signer.clone());
    let l1_provider =
        ProviderBuilder::new().connect_http(l1_rpc.parse().expect("parse l1_rpc to Url")).erased();

    let l2_provider =
        ProviderBuilder::new().connect_http(l2_rpc.parse().expect("parse l2_rpc to Url")).erased();

    let verify_provider = ProviderBuilder::new()
        .connect_http(l1_verify_rpc.parse().expect("parse l1_rpc to Url"))
        .erased();

    let l1_signer =
        ProviderBuilder::new().wallet(wallet).connect_provider(verify_provider.clone()).erased();

    let batch_syncer = BatchSyncer::new(
        Address::from_str(&rollup).unwrap(),
        Address::from_str(&shadow_rollup).unwrap(),
        l1_provider.clone(),
        l2_provider.clone(),
        l1_signer.clone(),
    );

    let shadow_prover = ShadowProver::new(
        signer.address(),
        Address::from_str(&rollup).unwrap(),
        Address::from_str(&shadow_rollup).unwrap(),
        l1_provider,
        verify_provider,
        l1_signer,
    );

    (batch_syncer, shadow_prover, l2_provider)
}

// Metric management
async fn metric_mng() {
    register_metrics();
    let metric_address = read_env_var("SHADOW_PROVING_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
    tokio::spawn(async move {
        let metrics = Router::new().route("/metrics", get(handle_metrics));
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
    REGISTRY.register(Box::new(METRICS.shadow_blocks_len.clone())).unwrap();
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
const LOG_FILE_BASENAME: &str = "app_info";
const LOG_FILE_SIZE_LIMIT: u64 = 200 * 10u64.pow(6); // 200MB
                                                     // const LOG_FILE_SIZE_LIMIT: u64 = 10u64.pow(3); // 1kB
const LOG_FILES_TO_KEEP: usize = 3;

fn setup_logging() {
    //configure the logger
    Logger::try_with_env_or_str(LOG_LEVEL)
        .unwrap()
        .log_to_file(
            FileSpec::default()
                .directory(read_env_var(
                    "SHADOW_PROVING_LOG_DIR",
                    String::from("/data/logs/morph-shadow-proving"),
                ))
                .basename(LOG_FILE_BASENAME),
        )
        .format(log_format)
        .duplicate_to_stdout(Duplicate::All)
        .rotate(
            Criterion::Size(LOG_FILE_SIZE_LIMIT), // Scroll when file size reaches 200MB
            Naming::TimestampsCustomFormat {
                current_infix: Some(""),
                format: "r%Y-%m-%d_%H-%M-%S",
            }, // Using timestamps as part of scrolling files
            Cleanup::KeepLogFiles(LOG_FILES_TO_KEEP), // Keep the latest 3 scrolling files
        )
        .write_mode(WriteMode::BufferAndFlush)
        .start()
        .unwrap();
}

fn log_format(
    w: &mut dyn std::io::Write,
    now: &mut flexi_logger::DeferredNow,
    record: &Record,
) -> Result<(), std::io::Error> {
    write!(
        w,
        "{} [{}] {} - {}",
        now.now().format("%Y-%m-%d %H:%M:%S"), // Custom time format
        record.level(),
        record.target(),
        record.args()
    )
}

#[tokio::test]
async fn test_shadow() {
    dotenv().ok();
    env_logger::Builder::new().filter_level(log::LevelFilter::Info).format_target(false).init();
    log::info!("Starting shadow proving...");

    // cargo test -p shadow-proving --bin shadow-proving -- test_shadow --exact --nocapture -- --batch-num 100 --no-prove
    let (batch_num, prove) = test_args::read_shadow_test_args_from_argv();

    let (batch_syncer, shadow_prover, l2_provider) = init_shadow_proving();
    let chain_id = l2_provider.get_chain_id().await.unwrap_or_default();

    let (batch_info, batch_header) =
        batch_syncer.get_specified_batch(batch_num).await.unwrap().unwrap();

    let offchain_batch_pi = try_execute_batch(&batch_info, &l2_provider).await.unwrap();

    let onchain_batch_pi = batch_syncer.calc_batch_pi(chain_id, &batch_header).unwrap_or_default();
    if offchain_batch_pi != onchain_batch_pi {
        log::error!(
            "Shadow execute batch pi mismatch! offchain: {:?}, onchain: {:?}",
            offchain_batch_pi,
            onchain_batch_pi
        );
        return;
    }
    if prove {
        let batch =
            batch_syncer.sync_batch(batch_info, batch_header.clone()).await.unwrap().unwrap();
        let prove_info = BatchProveInfo { batch_info: batch, batch_header };
        shadow_prover.prove(prove_info).await.unwrap();
    }
}

#[cfg(test)]
mod test_args {
    use clap::Parser;

    const DEFAULT_BATCH_NUM: u64 = 100;
    const DEFAULT_PROVE: bool = true;

    /// Shadow prove test parameters.
    #[derive(Parser, Debug)]
    #[command(author, version, about, long_about = None, disable_help_flag = true)]
    struct ShadowTestArgs {
        /// Batch index to test.
        #[arg(
            long = "batch-num",
            alias = "batch",
            default_value_t = DEFAULT_BATCH_NUM,
            value_parser = parse_u64_auto_radix
        )]
        batch_num: u64,

        /// Disable proving step (only do shadow execute/pi check).
        #[arg(long = "no-prove", default_value_t = DEFAULT_PROVE, action = clap::ArgAction::SetFalse)]
        prove: bool,
    }

    pub(super) fn read_shadow_test_args_from_argv() -> (u64, bool) {
        let filtered = filter_argv(&["--batch-num", "--batch", "--no-prove"]);
        let args = ShadowTestArgs::parse_from(filtered);
        (args.batch_num, args.prove)
    }

    fn filter_argv(allowed_flags: &[&str]) -> Vec<String> {
        let argv: Vec<String> = std::env::args().skip(1).collect();
        let mut filtered: Vec<String> = Vec::with_capacity(argv.len() + 1);
        // clap expects argv[0] to be the binary name, so we use a placeholder.
        filtered.push("shadow_test".to_string());

        let mut it = argv.into_iter();
        while let Some(arg) = it.next() {
            if allowed_flags.iter().any(|f| *f == arg) {
                filtered.push(arg);
                // Only flags that take a value need to consume the next argv.
                if filtered.last().map(|s| s.as_str()) != Some("--no-prove") {
                    if let Some(v) = it.next() {
                        filtered.push(v);
                    }
                }
            } else {
                // ignore unknown args
            }
        }

        filtered
    }

    fn parse_u64_auto_radix(s: &str) -> Result<u64, String> {
        let s = s.trim();
        if let Some(hex) = s.strip_prefix("0x").or_else(|| s.strip_prefix("0X")) {
            u64::from_str_radix(hex, 16).map_err(|e| e.to_string())
        } else {
            s.parse::<u64>().map_err(|e| e.to_string())
        }
    }
}
