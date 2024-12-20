use base64::{engine::general_purpose, Engine};
use ethers::{abi::AbiEncode, types::*, utils::hex};
use pem::{encode_config, EncodeConfig, Pem};
use reqwest::Client;
use rsa::{
    pkcs8::{DecodePrivateKey, EncodePublicKey},
    Pkcs1v15Sign, RsaPrivateKey,
};
use serde::{Deserialize, Serialize};
use sha2::{Digest, Sha256};
use std::{error::Error, time::UNIX_EPOCH};
use transaction::eip2718::TypedTransaction;
use uuid::Uuid;

#[derive(Clone)]
pub struct ExternalSign {
    appid: String,
    pub address: String,
    privkey: RsaPrivateKey,
    chain: String,
    client: Client,
    url: String,
}

#[derive(Serialize, Debug)]
struct BusinessData {
    appid: String,
    data: String,
    noncestr: String,
    timestamp: String,
}

#[derive(Serialize, Debug)]
struct ReqData {
    #[serde(flatten)]
    business_data: BusinessData,
    #[serde(rename = "bizSignature")]
    biz_signature: String,
    #[serde(rename = "publicKey")]
    public_key: String,
    #[serde(rename = "txData")]
    tx_data: String, // hex string of tx
}

#[derive(Serialize, Debug)]
struct Data {
    address: String,
    chain: String,
    sha3: String,
}

#[derive(Deserialize)]
struct Response {
    result: ResultData,
}

#[derive(Deserialize)]
struct ResultData {
    #[serde(rename = "signDatas")]
    sign_datas: Vec<SignData>,
}

#[derive(Deserialize)]
struct SignData {
    sign: String,
}

impl ExternalSign {
    pub fn new(appid: &str, privkey_pem: &str, address: &str, chain: &str, url: &str) -> Result<Self, Box<dyn Error>> {
        let privkey = RsaPrivateKey::from_pkcs8_pem(&reformat_pem(privkey_pem))?;
        let client = Client::new();
        Ok(ExternalSign {
            appid: appid.to_string(),
            address: address.to_string(),
            privkey,
            chain: chain.to_string(),
            client,
            url: url.to_string(),
        })
    }

    pub async fn request_sign(&self, tx: &TypedTransaction) -> Result<Bytes, Box<dyn Error>> {
        let data = self.new_data(&hex::encode_prefixed(tx.sighash().encode()))?;
        let tx_info = hex::encode(tx.rlp_signed(&Signature::try_from(vec![0u8; 65].as_slice()).unwrap()));

        let req_data = self.craft_req_data(data, tx_info)?;

        let rt = self.do_request(&self.url, &req_data).await?;
        log::debug!("ext_sign response: {:?}", rt);

        let response: Response = serde_json::from_str(&rt)?;
        if response.result.sign_datas.is_empty() {
            return Err("ext_sign response sign_datas empty".into());
        }
        if response.result.sign_datas[0].sign.len() < 132 {
            //hex prefix + 65bytes sig
            return Err("ext_sign response sign data invalid".into());
        }

        let sig = hex::decode(&response.result.sign_datas[0].sign[2..])?;
        let signed_tx: Bytes = tx.rlp_signed(&Signature::try_from(sig.as_slice())?);
        Ok(signed_tx)
    }

    fn new_data(&self, hash: &str) -> Result<Data, Box<dyn Error>> {
        Ok(Data {
            address: self.address.clone(),
            chain: self.chain.clone(),
            sha3: hash.to_string(),
        })
    }

    fn craft_req_data(&self, data: Data, tx_info: String) -> Result<ReqData, Box<dyn Error>> {
        let nonce_str = Uuid::new_v4().to_string();
        let data_bs = serde_json::to_string(&data)?;

        let business_data = BusinessData {
            appid: self.appid.clone(),
            data: data_bs,
            noncestr: nonce_str,
            timestamp: UNIX_EPOCH.elapsed()?.as_secs().to_string(),
        };

        let business_data_bs = serde_json::to_string(&business_data)?;
        let hashed = sha2::Sha256::digest(business_data_bs.as_bytes());
        let signature = self.privkey.sign(Pkcs1v15Sign::new::<Sha256>(), &hashed)?;
        let hex_sig = hex::encode(signature);

        let pubkey = self.privkey.to_public_key().to_public_key_der()?;
        let pubkey_base64 = base64::engine::general_purpose::STANDARD.encode(pubkey.as_ref());
        Ok(ReqData {
            business_data,
            biz_signature: hex_sig,
            public_key: pubkey_base64,
            tx_data: tx_info,
        })
    }

    async fn do_request(&self, url: &str, payload: &ReqData) -> Result<String, Box<dyn Error>> {
        log::debug!("===payload: {:?}", serde_json::to_string(payload).unwrap());
        let response: reqwest::Response = self.client.post(url).json(&payload).send().await?;
        if !response.status().is_success() {
            return Err(format!("ext_sign response status not ok: {:?}", response.status()).into());
        }
        Ok(response.text().await?)
    }
}

fn reformat_pem(pem_string: &str) -> String {
    // Decode the base64 encoded string
    let key_bytes = general_purpose::STANDARD.decode(pem_string).expect("Failed to decode base64");

    // Create a PEM object with the required format
    let pem = Pem {
        tag: String::from("PRIVATE KEY"),
        contents: key_bytes,
    };

    // Encode the PEM object to a string with standard line width of 64 characters
    let config = EncodeConfig {
        line_ending: pem::LineEnding::LF,
    };
    encode_config(&pem, config)
}
