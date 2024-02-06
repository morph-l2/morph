use app::gas_price_oracle;
use env_logger::Env;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Initialize logger.
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!("Start updating gas oracle...");

    // Update GasPriceOrale contract on L2 network.
    let result = gas_price_oracle::update().await;

    // Handle result.
    match result {
        Ok(()) => Ok(()),
        Err(e) => {
            log::error!("oracle exec error: {:#?}", e);
            Err(e)
        }
    }
}
