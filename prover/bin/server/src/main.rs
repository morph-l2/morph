use dotenv::dotenv;
use flexi_logger::{
    filter::{LogLineFilter, LogLineWriter},
    Cleanup, Criterion, DeferredNow, Duplicate, FileSpec, Logger, Naming, WriteMode,
};
use log::Record;
use prover_server::{read_env_var, server};

#[tokio::main]
async fn main() {
    dotenv().ok();
    setup_logging();
    log::info!("Starting server...");

    server::start().await;
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
                .directory(read_env_var("PROVER_LOG_DIR", String::from("/data/logs/morph-prover")))
                .basename(LOG_FILE_BASENAME),
        )
        .format(log_format)
        .filter(Box::new(ProveFilter))
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

pub struct ProveFilter;

impl LogLineFilter for ProveFilter {
    fn write(
        &self,
        now: &mut DeferredNow,
        record: &log::Record,
        log_line_writer: &dyn LogLineWriter,
    ) -> std::io::Result<()> {
        let module_path = record.module_path().unwrap_or("start");
        let args = record.args().to_string();
        if !args.contains("tracing::span") && !module_path.contains("p3_") {
            log_line_writer.write(now, record)?;
        }
        Ok(())
    }
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
