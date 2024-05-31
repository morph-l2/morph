use eyre::anyhow;
use serde_json::{json, Value};

use super::error::*;

pub struct ExecutionNode {
    pub rpc_url: String,
}

impl ExecutionNode {
    pub async fn query_blob_tx(&self, hash: &str) -> Result<Value, OverHeadError> {
        let params: serde_json::Value = json!([hash]);
        let rpc_url = self.rpc_url.clone();

        tokio::task::spawn_blocking(move || {
            Self::query_execution_node(
                rpc_url,
                &json!({
                    "jsonrpc": "2.0",
                    "method": "eth_getTransactionByHash",
                    "params": params,
                    "id": 1,
                }),
            )
        })
        .await
        .expect("Tokio spawn blocking issue with query_blob_tx")
    }

    pub async fn query_blob_tx_receipt(&self, hash: &str) -> Result<Value, OverHeadError> {
        let params: serde_json::Value = json!([hash]);
        let rpc_url = self.rpc_url.clone();

        tokio::task::spawn_blocking(move || {
            Self::query_execution_node(
                rpc_url,
                &json!({
                    "jsonrpc": "2.0",
                    "method": "eth_getTransactionReceipt",
                    "params": params,
                    "id": 1,
                }),
            )
        })
        .await
        .expect("Tokio spawn blocking issue with query_blob_tx_receipt")
    }

    pub async fn query_block(&self, hash: &str) -> Result<Value, OverHeadError> {
        let params: serde_json::Value = json!([hash, true]);
        let rpc_url = self.rpc_url.clone();

        tokio::task::spawn_blocking(move || {
            Self::query_execution_node(
                rpc_url,
                &json!({
                    "jsonrpc": "2.0",
                    "method": "eth_getBlockByHash",
                    "params": params,
                    "id": 1,
                }),
            )
        })
        .await
        .expect("Tokio spawn blocking issue with query_block")
    }

    pub async fn query_block_by_num(&self, num: u64) -> Result<Value, OverHeadError> {
        let params: serde_json::Value = if num != 0 {
            json!([format!("0x{:X}", num), false])
        } else {
            json!(["latest", false])
        };
        let rpc_url = self.rpc_url.clone();
        tokio::task::spawn_blocking(move || {
            Self::query_execution_node(
                rpc_url,
                &json!({
                    "jsonrpc": "2.0",
                    "method": "eth_getBlockByNumber",
                    "params": params,
                    "id": 1,
                }),
            )
        })
        .await
        .expect("Tokio spawn blocking issue with query_block_by_num")
    }

    fn query_execution_node(
        rpc_url: String,
        param: &serde_json::Value,
    ) -> Result<Value, OverHeadError> {
        let client = reqwest::blocking::Client::new();
        let response = client
            .post(rpc_url)
            .header(
                reqwest::header::CONTENT_TYPE,
                reqwest::header::HeaderValue::from_static("application/json"),
            )
            .json(param)
            .send();

        match response {
            Ok(rs) => match rs.text() {
                Ok(r) => serde_json::from_str::<Value>(&r).map_err(|e| {
                    OverHeadError::ExecutionNodeError(anyhow!(
                        "deserialize response failed, params= {:#?}, err: {:#?}",
                        param,
                        e
                    ))
                }),
                Err(e) => Err(OverHeadError::ExecutionNodeError(anyhow!(format!(
                    "fetch l1 execution node res_txt error, param= {:#?}, error = {:#?}",
                    param, e
                )))),
            },
            Err(e) => Err(OverHeadError::ExecutionNodeError(anyhow!(format!(
                "call l1 execution node error, param= {:#?}, error = {:#?}",
                param, e
            )))),
        }
    }
}

pub struct BeaconNode {
    pub rpc_url: String,
}
impl BeaconNode {
    pub async fn query_sidecars(
        &self,
        slot: String,
        indexes: Vec<u64>,
    ) -> Result<Value, OverHeadError> {
        let rpc_url = self.rpc_url.clone();

        tokio::task::spawn_blocking(move || {
            Self::query_beacon_node(rpc_url, slot.as_str(), indexes)
        })
        .await
        .expect("Tokio spawn blocking issue with query_sidecars")
    }

    fn query_beacon_node(
        l1_beacon_rpc: String,
        slot: &str,
        indexes: Vec<u64>,
    ) -> Result<Value, OverHeadError> {
        let client = reqwest::blocking::Client::new();
        let mut url = l1_beacon_rpc.to_owned() +
            "/eth/v1/beacon/blob_sidecars/" +
            slot.to_string().as_str() +
            "?";
        for index in indexes {
            url = url + "indices=" + index.to_string().as_str() + "&";
        }

        let response = client
            .get(url.clone())
            .header(
                reqwest::header::CONTENT_TYPE,
                reqwest::header::HeaderValue::from_static("application/json"),
            )
            .send();

        match response {
            Ok(rs) => match rs.text() {
                Ok(r) => serde_json::from_str::<Value>(&r).map_err(|e| {
                    OverHeadError::BeaconNodeError(anyhow!(
                        "deserialize response failed, slot= {:#?}, error = {:#?}",
                        slot,
                        e
                    ))
                }),
                Err(e) => Err(OverHeadError::BeaconNodeError(anyhow!(format!(
                    "fetch beacon node res_txt error, slot= {:#?}, error = {:#?}",
                    slot, e
                )))),
            },
            Err(e) => Err(OverHeadError::BeaconNodeError(anyhow!(format!(
                "query beacon node error, slot= {:#?}, error = {:#?}",
                slot, e
            )))),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::env::var;

    #[tokio::test]
    #[ignore]
    async fn test_query_execution_node() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let params: serde_json::Value =
            json!(["0x0037beafe424df970b35eb7eb5fadb5f34c16159f6ec58818947444b10e43cdd"]);

        let rpc_url: String =
            var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env var not found");

        let rt = tokio::task::spawn_blocking(move || {
            ExecutionNode::query_execution_node(
                rpc_url,
                &json!({
                    "jsonrpc": "2.0",
                    "method": "eth_getTransactionByHash",
                    "params": params,
                    "id": 1,
                }),
            )
        })
        .await
        .unwrap();

        match rt {
            Ok(transaction) => {
                log::info!(
                    "blobVersionedHashes: {:#?}",
                    transaction["result"]["blobVersionedHashes"]
                );
            }
            Err(e) => {
                log::error!("{:?}", e);
            }
        }
    }

    #[tokio::test]
    #[ignore]
    async fn test_query_block() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let execution_node = ExecutionNode {
            rpc_url: var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env var not found"),
        };

        let rt = execution_node
            .query_block("0x1f4b9b40ff7eba7c4cbf48c1a03551c25bb23ac69963016a72c9069db5ca00db")
            .await;

        match rt {
            Ok(info) => {
                log::info!(
                    "transactions: {:#?}",
                    info["result"]["transactions"].as_array().unwrap()
                );
            }
            Err(e) => {
                log::error!("{:?}", e);
            }
        }
    }

    #[tokio::test]
    #[ignore]
    async fn test_query_beacon_node() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let rpc_url: String =
            var("GAS_ORACLE_L1_BEACON_RPC").expect("GAS_ORACLE_L1_BEACON_RPC env var not found");

        let rt = tokio::task::spawn_blocking(move || {
            BeaconNode::query_beacon_node(rpc_url, "1053668", vec![0])
        })
        .await
        .unwrap();

        match rt {
            Ok(transaction) => {
                log::info!("blobVersionedHashes: {:#?}", transaction["data"][0]["kzg_commitment"]);
            }
            Err(e) => {
                log::error!("{:?}", e);
            }
        }
    }
}
