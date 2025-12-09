use crate::{util, BatchInfo};
use anyhow::anyhow;
use serde::{Deserialize, Serialize};
#[derive(Serialize)]
pub struct ExecuteRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ExecuteResult {
    pub error_msg: String,
    pub error_code: String,
}

pub async fn execute_batch(batch: &BatchInfo) -> Result<(), anyhow::Error> {
    // Request the proverServer to prove.
    let request = ExecuteRequest {
        batch_index: batch.batch_index,
        start_block: batch.start_block,
        end_block: batch.end_block,
        rpc: "http://localhost:8545".to_owned(),
    };
    let rt = tokio::task::spawn_blocking(move || {
        util::call_prover(serde_json::to_string(&request).unwrap(), "/prove_batch")
    })
    .await;

    match rt {
        Ok(Some(_)) => Ok(()),
        Ok(None) => Err(anyhow!("call_prover result empty")),
        Err(e) => Err(anyhow::Error::from(e)),
    }
}
