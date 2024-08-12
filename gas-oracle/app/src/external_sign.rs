// main.rs (or lib.rs)

use ethers::{abi::AbiEncode, types::*};
use reqwest::Client;
use rsa::{
    pkcs1::{DecodeRsaPrivateKey, EncodeRsaPublicKey},
    Pkcs1v15Sign, RsaPrivateKey,
};
use serde::{Deserialize, Serialize};
use sha2::Sha256;
// use sha2::{Digest, Sha256};
use sha2::Digest;
use std::{error::Error, time::SystemTime};
use transaction::eip2718::TypedTransaction;
use uuid::Uuid;

#[derive(Clone)]
pub struct ExternalSign {
    pub appid: String,
    pub address: String,
    pub privkey: RsaPrivateKey,
    pub chain: String,
    pub client: Client,
    pub url: String,
}

#[derive(Serialize)]
struct BusinessData {
    appid: String,
    data: String,
    noncestr: String,
    timestamp: String,
}

#[derive(Serialize)]
struct ReqData {
    business_data: BusinessData,
    biz_signature: String,
    public_key: String,
}

#[derive(Serialize)]
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
    sign_datas: Vec<SignData>,
}

#[derive(Deserialize)]
struct SignData {
    sign: String,
}

fn parse_rsa_private_key(
    private_key_str: &str,
) -> Result<RsaPrivateKey, Box<dyn std::error::Error>> {
    // let pem = parse(private_key_str)?;
    let private_key = RsaPrivateKey::from_pkcs1_pem(private_key_str)?;

    Ok(private_key)
}

impl ExternalSign {
    pub fn new(
        appid: &str,
        privkey_pem: &str,
        address: &str,
        chain: &str,
        url: &str,
    ) -> Result<Self, Box<dyn Error>> {
        let privkey = RsaPrivateKey::from_pkcs1_pem(privkey_pem)?;
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
        let tx_hash = tx.sighash();
        let data = self.new_data(&hex::encode(tx_hash.encode()))?;

        let req_data = self.craft_req_data(data)?;

        let resp = self.do_request(&self.url, &req_data).await?;

        let response: Response = serde_json::from_str(&resp)?;

        // This assumes the response includes the signature in hex form
        if response.result.sign_datas.is_empty() {
            return Err("response sha3 empty".into());
        }

        let sig = hex::decode(&response.result.sign_datas[0].sign)?;
        // Apply the signature to the transaction (example, adjust as needed)
        let signed_tx: Bytes = tx.rlp_signed(&Signature::try_from(sig.as_slice()).unwrap());
        Ok(signed_tx)
    }

    fn new_data(&self, hash: &str) -> Result<Data, Box<dyn Error>> {
        Ok(Data {
            address: self.address.clone(),
            chain: self.chain.clone(),
            sha3: hash.to_string(),
        })
    }

    fn craft_req_data(&self, data: Data) -> Result<ReqData, Box<dyn Error>> {
        let nonce_str = Uuid::new_v4().to_string();
        let data_bs = serde_json::to_string(&data)?;

        let business_data = BusinessData {
            appid: self.appid.clone(),
            data: data_bs,
            noncestr: nonce_str,
            timestamp: SystemTime::now()
                .duration_since(SystemTime::UNIX_EPOCH)
                .unwrap()
                .as_secs()
                .to_string(),
        };

        let business_data_bs = serde_json::to_string(&business_data)?;
        let hashed = sha2::Sha256::digest(business_data_bs.as_bytes());

        // let signing_key = SigningKey::<Sha256>::new(self.privkey.clone());

        let signature = self.privkey.sign(Pkcs1v15Sign::new::<Sha256>(), &hashed)?;
        let hex_sig = hex::encode(signature);

        let pubkey = self.privkey.to_public_key().to_pkcs1_der().unwrap();
        let pubkey_str = pem::encode(&pem::Pem {
            tag: "RSA PUBLIC KEY".into(),
            contents: pubkey.as_ref().to_vec(),
        });

        Ok(ReqData { business_data, biz_signature: hex_sig, public_key: pubkey_str })
    }

    async fn do_request(&self, url: &str, payload: &ReqData) -> Result<String, Box<dyn Error>> {
        let response = self.client.post(url).json(payload).send().await?;

        if !response.status().is_success() {
            return Err(format!("response status not ok: {:?}", response.status()).into());
        }

        Ok(response.text().await?)
    }

    fn hash_transaction(&self, tx: &TransactionRequest) -> String {
        // Implement transaction hashing suitable to your needs
        // This is a stub placeholder
        "your_hash_hex".to_string()
    }
}

#[tokio::test]
async fn test_sign() -> Result<(), Box<dyn Error>> {
    // Initialize your ExternalSign instance and perform signing as needed
    // You would populate the actual Ethereum transaction details and other logistics here

    Ok(())
}
