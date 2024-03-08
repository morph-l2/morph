use serde_json::{json, Value};
use std::env::var;

pub async fn query_blob_tx(hash: &str) -> Option<Value> {
    let params: serde_json::Value = json!([hash]);

    let rt = tokio::task::spawn_blocking(move || {
        query_execution_node(&json!({
            "jsonrpc": "2.0",
            "method": "eth_getTransactionByHash",
            "params": params,
            "id": 1,
        }))
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

pub async fn query_blob_tx_receipt(hash: &str) -> Option<Value> {
    let params: serde_json::Value = json!([hash]);

    let rt = tokio::task::spawn_blocking(move || {
        query_execution_node(&json!({
            "jsonrpc": "2.0",
            "method": "eth_getTransactionReceipt",
            "params": params,
            "id": 1,
        }))
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

pub async fn query_block(hash: &str) -> Option<Value> {
    let params: serde_json::Value = json!([hash, true]);

    let rt = tokio::task::spawn_blocking(move || {
        query_execution_node(&json!({
            "jsonrpc": "2.0",
            "method": "eth_getBlockByHash",
            "params": params,
            "id": 1,
        }))
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

pub async fn query_block_by_num(num: u64) -> Option<Value> {
    let params: serde_json::Value = json!([format!("0x{:X}", num), true]);

    let rt = tokio::task::spawn_blocking(move || {
        query_execution_node(&json!({
            "jsonrpc": "2.0",
            "method": "eth_getBlockByNumber",
            "params": params,
            "id": 1,
        }))
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

pub async fn query_sidecars(slot: String, indexes: Vec<u64>) -> Option<Value> {
    let rt = tokio::task::spawn_blocking(move || query_beacon_node(slot.as_str(), indexes)).await;

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

fn query_execution_node(param: &serde_json::Value) -> Option<String> {
    let l1_rpc = if let Ok(v) = var("GAS_ORACLE_L1_RPC") {
        v
    } else {
        log::info!("query_execution_node: Cannot detect GAS_ORACLE_L1_RPC env var");
        return None;
    };

    let client = reqwest::blocking::Client::new();
    let url = l1_rpc.to_owned();
    let response = client
        .post(url)
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

fn query_beacon_node(slot: &str, indexes: Vec<u64>) -> Option<String> {
    let l1_beacon_rpc = if let Ok(v) = var("GAS_ORACLE_L1_BEACON_RPC") {
        v
    } else {
        log::info!("query_execution_node: Cannot detect GAS_ORACLE_L1_BEACON_RPC env var");
        return None;
    };

    let client = reqwest::blocking::Client::new();
    let mut url = l1_beacon_rpc.to_owned() + slot.to_string().as_str() + "?";
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

#[tokio::test]
async fn test_query_execution_node() {
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let params: serde_json::Value =
        json!(["0x541cee01d959a9c8ea9f6607763a1e048327dcaf312f1d435fddfbc4a1e78dc7"]);

    let rt = tokio::task::spawn_blocking(move || {
        query_execution_node(&json!({
            "jsonrpc": "2.0",
            "method": "eth_getTransactionByHash",
            "params": params,
            "id": 1,
        }))
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
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let rt: Option<Value> =
        query_block("0xa37b1b946129bc8c4d50daf31978c8f2954d0b3c6e2ceffd486e33ed94cbeec2").await;

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
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    let rt = tokio::task::spawn_blocking(move || query_beacon_node("4481517", vec![0]))
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
