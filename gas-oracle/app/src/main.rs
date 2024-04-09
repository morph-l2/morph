use app::gas_price_oracle;
use flexi_logger::{Cleanup, Criterion, Duplicate, FileSpec, Logger, Naming, WriteMode};
use log::Record;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Initialize logger.
    setup_logging();
    log::info!("Start updating gas oracle...");

    // Update GasPriceOrale contract on L2 network.
    let result = gas_price_oracle::update().await;

    // Handle result.
    match result {
        Ok(()) => Ok(()),
        Err(e) => {
            log::error!("gas oracle exec error: {:#?}", e.to_string());
            Err(e)
        }
    }
}

// Constants for configuration
const LOG_LEVEL: &str = "info";
const LOG_FILE_BASENAME: &str = "gasoracle";
const LOG_FILE_SIZE_LIMIT: u64 = 100 * 10u64.pow(6); // 100MB
const LOG_FILES_TO_KEEP: usize = 3;

fn setup_logging() {
    //configure the logger
    Logger::try_with_env_or_str(LOG_LEVEL)
        .unwrap()
        .log_to_file(FileSpec::default().basename(LOG_FILE_BASENAME))
        .format(custom_format)
        .duplicate_to_stdout(Duplicate::All)
        .rotate(
            Criterion::Size(LOG_FILE_SIZE_LIMIT), // Scroll when file size reaches 10MB
            Naming::TimestampsDirect,             // Using timestamps as part of scrolling files
            Cleanup::KeepLogFiles(LOG_FILES_TO_KEEP), // Keep the latest 3 scrolling files
        )
        .write_mode(WriteMode::BufferAndFlush)
        .start()
        .unwrap();
}

fn custom_format(
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
