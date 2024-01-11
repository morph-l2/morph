use challenge_handler::handler;
use env_logger::Env;

#[tokio::main]
async fn main() {
    // Initialize logger.
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!("Starting challenge handler...");

    // Start challenge handler.
    let result = handler::handle_challenge().await;

    // Handle result.
    match result {
        Ok(()) => (),
        Err(e) => {
            log::error!("challenge handler exec error: {:#?}", e);
        }
    }
}
