use serde_json::{json, Value};

pub struct ExecutionNode {
    pub rpc_url: String,
}

impl ExecutionNode {
    pub async fn query_blob_tx(&self, hash: &str) -> Option<Value> {
        let params: serde_json::Value = json!([hash]);
        let rpc_url = self.rpc_url.clone();

        let rt = tokio::task::spawn_blocking(move || {
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
        .map_or_else(
            |e| {
                log::error!("query_blob_tx error: {:?}", e);
                return None;
            },
            |v| v,
        )?;

        match serde_json::from_str::<Value>(&rt) {
            Ok(parsed) => return Some(parsed),
            Err(_) => {
                log::error!("deserialize query_blob_tx failed, hash= {:?}", hash);
                return None;
            }
        };
    }

    pub async fn query_blob_tx_receipt(&self, hash: &str) -> Option<Value> {
        let params: serde_json::Value = json!([hash]);
        let rpc_url = self.rpc_url.clone();

        let rt = tokio::task::spawn_blocking(move || {
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
        .await;

        match rt {
            Ok(Some(info)) => {
                match serde_json::from_str::<Value>(&info) {
                    Ok(parsed) => return Some(parsed),
                    Err(_) => {
                        log::error!("deserialize query_blob_tx_receipt failed, hash= {:?}", hash);
                        return None;
                    }
                };
            }
            _ => {
                log::error!("query_blob_tx_receipt failed");
                return None;
            }
        }
    }

    pub async fn query_block(&self, hash: &str) -> Option<Value> {
        let params: serde_json::Value = json!([hash, true]);
        let rpc_url = self.rpc_url.clone();

        let rt = tokio::task::spawn_blocking(move || {
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
        .await;

        match rt {
            Ok(Some(info)) => {
                match serde_json::from_str::<Value>(&info) {
                    Ok(parsed) => return Some(parsed),
                    Err(_) => {
                        log::error!("deserialize block failed, hash= {:?}", hash);
                        return None;
                    }
                };
            }
            _ => {
                log::error!("query block failed");
                return None;
            }
        }
    }

    pub async fn query_block_by_num(&self, num: u64) -> Option<Value> {
        let params: serde_json::Value = if num == 0 {
            json!([format!("0x{:X}", num), false])
        } else {
            json!(["latest", false])
        };
        let rpc_url = self.rpc_url.clone();
        let rt = tokio::task::spawn_blocking(move || {
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
        .unwrap();

        match rt {
            Some(info) => {
                match serde_json::from_str::<Value>(&info) {
                    Ok(parsed) => return Some(parsed),
                    Err(_) => {
                        log::error!("deserialize block failed, blockNum= {:?}", num);
                        return None;
                    }
                };
            }
            None => {
                log::error!("query block failed");
                return None;
            }
        }
    }

    fn query_execution_node(rpc_url: String, param: &serde_json::Value) -> Option<String> {
        let client = reqwest::blocking::Client::new();
        let response = client
            .post(rpc_url)
            .header(
                reqwest::header::CONTENT_TYPE,
                reqwest::header::HeaderValue::from_static("application/json"),
            )
            .json(param)
            .send();
        let rt: Result<String, reqwest::Error> = match response {
            Ok(x) => x.text(),
            Err(e) => {
                log::error!(
                    "call l1 execution node error, param =  {:#?}, error = {:#?}",
                    param,
                    e
                );
                return None;
            }
        };

        let rt_text = match rt {
            Ok(x) => x,
            Err(e) => {
                log::error!(
                    "fetch l1 execution node res_txt error, param =  {:#?}, error = {:#?}",
                    param,
                    e
                );
                return None;
            }
        };

        Some(rt_text)
    }
}

pub struct BeaconNode {
    pub rpc_url: String,
}
impl BeaconNode {
    pub async fn query_sidecars(&self, slot: String, indexes: Vec<u64>) -> Option<Value> {
        let rpc_url = self.rpc_url.clone();

        let rt = tokio::task::spawn_blocking(move || {
            Self::query_beacon_node(rpc_url, slot.as_str(), indexes)
        })
        .await;

        match rt {
            Ok(Some(info)) => {
                match serde_json::from_str::<Value>(&info) {
                    Ok(parsed) => return Some(parsed),
                    Err(_) => {
                        log::error!("deserialize blobSidecars from beacon failed",);
                        return None;
                    }
                };
            }
            _ => {
                log::error!("query blobSidecars from beacon node failed");
                return None;
            }
        }
    }

    fn query_beacon_node(l1_beacon_rpc: String, slot: &str, indexes: Vec<u64>) -> Option<String> {
        let client = reqwest::blocking::Client::new();
        let mut url = l1_beacon_rpc.to_owned()
            + "/eth/v1/beacon/blob_sidecars/"
            + slot.to_string().as_str()
            + "?";
        for index in indexes {
            url = url + "indices=" + index.to_string().as_str() + "&";
        }
        let response = client
            .get(url)
            .header(
                reqwest::header::CONTENT_TYPE,
                reqwest::header::HeaderValue::from_static("application/json"),
            )
            .send();
        let rt: Result<String, reqwest::Error> = match response {
            Ok(x) => x.text(),
            Err(e) => {
                log::error!(
                    "query beacon node error, slot =  {:#?}, error = {:#?}",
                    slot,
                    e
                );
                return None;
            }
        };

        let rt_text = match rt {
            Ok(x) => x,
            Err(e) => {
                log::error!(
                    "fetch beacon node res_txt error, slot =  {:#?}, error = {:#?}",
                    slot,
                    e
                );
                return None;
            }
        };

        Some(rt_text)
    }
}

#[tokio::test]
async fn test_query_execution_node() {
    use std::env::var;
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let params: serde_json::Value =
        json!(["0x541cee01d959a9c8ea9f6607763a1e048327dcaf312f1d435fddfbc4a1e78dc7"]);

    let rpc_url: String = var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env var not found");

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
        Some(info) => {
            log::info!("query result: {:#?}", info);
            let transaction = match serde_json::from_str::<Value>(&info) {
                Ok(parsed) => parsed,
                Err(_) => {
                    log::error!("deserialize rollup_batch failed, batch index");
                    return;
                }
            };
            log::info!(
                "blobVersionedHashes: {:#?}",
                transaction["result"]["blobVersionedHashes"]
            );
        }
        None => {
            log::error!("submitt prove task failed");
        }
    }
}

#[tokio::test]
async fn test_query_block() {
    use std::env::var;

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let execution_node = ExecutionNode {
        rpc_url: var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env var not found"),
    };

    let rt: Option<Value> = execution_node
        .query_block("0xa37b1b946129bc8c4d50daf31978c8f2954d0b3c6e2ceffd486e33ed94cbeec2")
        .await;

    match rt {
        Some(info) => {
            log::info!(
                "transactions: {:#?}",
                info["result"]["transactions"].as_array().unwrap()
            );
        }
        None => {
            log::error!("query_block failed");
        }
    }
}

#[tokio::test]
async fn test_query_beacon_node() {
    use std::env::var;
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let rpc_url: String =
        var("GAS_ORACLE_L1_BEACON_RPC").expect("GAS_ORACLE_L1_BEACON_RPC env var not found");

    let rt = tokio::task::spawn_blocking(move || {
        BeaconNode::query_beacon_node(rpc_url, "4481517", vec![0])
    })
    .await
    .unwrap();

    match rt {
        Some(info) => {
            // log::info!("query result: {:#?}", info);
            let transaction = match serde_json::from_str::<Value>(&info) {
                Ok(parsed) => parsed,
                Err(_) => {
                    log::error!("deserialize rollup_batch failed, batch index");
                    return;
                }
            };
            log::info!(
                "blobVersionedHashes: {:#?}",
                transaction["data"][0]["kzg_commitment"]
            );
        }
        None => {
            log::error!("submitt prove task failed");
        }
    }
}
