use prover_server::server;

#[tokio::main]
async fn main() {
    println!("server starting...");
    server::start().await;
}
