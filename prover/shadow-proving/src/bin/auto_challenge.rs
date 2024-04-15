use challenge_handler::{challenge, shadow_challenge, util};

use dotenv::dotenv;
use env_logger::Env;

#[tokio::main]
async fn main() {
    // Prepare environment.
    dotenv().ok();
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!("Starting challenge handler...");

    // Start challenger.
    let result = if util::read_env_var("SHADOW_PROVING", true) {
        shadow_challenge::run().await
    } else {
        challenge::run().await
    };

    // Handle result.
    match result {
        Ok(()) => (),
        Err(e) => {
            log::error!("challenge handler exec error: {:#?}", e);
        }
    }
}
