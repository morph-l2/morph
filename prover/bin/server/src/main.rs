use prover_server::{queue, server};

#[tokio::main]
async fn main() {
    println!("Server starting...");
    // server::start().await;
    queue::test_prove();
}
