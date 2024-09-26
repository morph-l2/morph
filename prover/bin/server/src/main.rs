use prover_server::server;

#[tokio::main]
async fn main() {
    println!("Server starting...");
    server::start().await;
}
