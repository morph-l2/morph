use prover_server::server;

#[tokio::main]
async fn main() {
    server::start().await;
}
