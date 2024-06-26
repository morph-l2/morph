use super::error::*;
use eyre::anyhow;
use serde_json::Value;

pub struct BeaconNode {
    pub rpc_url: String,
}
impl BeaconNode {
    pub async fn query_sidecars(
        &self,
        slot: String,
        indexes: Vec<u64>,
    ) -> Result<Value, ScalarError> {
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
    ) -> Result<Value, ScalarError> {
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
                    ScalarError::BeaconNodeError(anyhow!(
                        "deserialize response failed, slot= {:#?}, error = {:#?}",
                        slot,
                        e
                    ))
                }),
                Err(e) => Err(ScalarError::BeaconNodeError(anyhow!(format!(
                    "fetch beacon node res_txt error, slot= {:#?}, error = {:#?}",
                    slot, e
                )))),
            },
            Err(e) => Err(ScalarError::BeaconNodeError(anyhow!(format!(
                "query beacon node error, slot= {:#?}, error = {:#?}",
                slot, e
            )))),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use ethers::{prelude::*, utils::hex};
    use std::env::var;

    #[tokio::test]
    #[ignore]
    async fn test_query_beacon_node_sidecars() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let rpc_url = var("GAS_ORACLE_L1_BEACON_RPC").expect("GAS_ORACLE_L1_BEACON_RPC env emmpty");
        let blk_number = U64::from(1644796);

        let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env empty");
        let l1_provider = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();

        let blk_info =
            l1_provider.get_block(BlockNumber::Number(blk_number + 1)).await.unwrap().unwrap();
        let pre_beacon_root = blk_info.parent_beacon_block_root.unwrap();

        let beacon_node = BeaconNode { rpc_url };

        let sidecars =
            beacon_node.query_sidecars(hex::encode_prefixed(pre_beacon_root), vec![0]).await;

        assert!(sidecars.is_ok());

        let sidecars = sidecars.unwrap();
        assert!(!sidecars["data"][0]["kzg_commitment"].is_null());
    }

    #[tokio::test]
    #[ignore]
    async fn test_query_beacon_node() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let rpc_url = var("GAS_ORACLE_L1_BEACON_RPC").expect("GAS_ORACLE_L1_BEACON_RPC env emmpty");

        let rt = tokio::task::spawn_blocking(move || {
            BeaconNode::query_beacon_node(rpc_url, "1053668", vec![0])
        })
        .await
        .expect("Tokio spawn blocking issue with query_beacon_node");

        match rt {
            Ok(transaction) => {
                log::info!("kzg_commitment: {:#?}", transaction["data"][0]["kzg_commitment"]);
            }
            Err(e) => {
                log::error!("{:?}", e);
            }
        }
    }

    #[tokio::test]
    #[ignore]
    async fn test_get_blob_tx() {
        env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
        dotenv::dotenv().ok();

        let l1_rpc = var("GAS_ORACLE_L1_RPC").expect("GAS_ORACLE_L1_RPC env empty");
        let l1_provider = Provider::<Http>::try_from(l1_rpc.clone()).unwrap();

        // blob tx
        let tx_hash = "0x0037beafe424df970b35eb7eb5fadb5f34c16159f6ec58818947444b10e43cdd"
            .parse::<H256>()
            .unwrap();
        let transaction = l1_provider.get_transaction(tx_hash).await.unwrap().unwrap();
        let tx_blob_versioned_hashes = transaction
            .other
            .get_with("blobVersionedHashes", serde_json::from_value::<Vec<H256>>)
            .unwrap_or(Ok(Vec::<H256>::new()))
            .unwrap_or_default();
        assert!(!tx_blob_versioned_hashes.is_empty());

        let receipt = l1_provider.get_transaction_receipt(tx_hash).await.unwrap().unwrap();
        let blob_gas_price = receipt
            .other
            .get_with("blobGasPrice", serde_json::from_value::<U256>)
            .unwrap_or(Ok(U256::zero()))
            .unwrap_or_default();
        assert!(!blob_gas_price.is_zero());

        // legacy tx
        let tx_hash = "0xd29de24c7447fb6bf4664586b5ee9c146d72e6bad74b2e9003ed5f7da80ccf51"
            .parse::<H256>()
            .unwrap();
        let transaction = l1_provider.get_transaction(tx_hash).await.unwrap().unwrap();
        let tx_blob_versioned_hashes = transaction
            .other
            .get_with("blobVersionedHashes", serde_json::from_value::<Vec<H256>>)
            .unwrap_or(Ok(Vec::<H256>::new()))
            .unwrap_or_default();
        assert!(tx_blob_versioned_hashes.is_empty());

        let receipt = l1_provider.get_transaction_receipt(tx_hash).await.unwrap().unwrap();
        let blob_gas_price = receipt
            .other
            .get_with("blobGasPrice", serde_json::from_value::<U256>)
            .unwrap_or(Ok(U256::zero()))
            .unwrap_or_default();
        assert!(blob_gas_price.is_zero());
    }
}
