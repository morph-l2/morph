use base64::Engine;
use ethers::{abi::AbiEncode, types::*};
use reqwest::Client;
use rsa::{
    pkcs8::{DecodePrivateKey, EncodePublicKey},
    Pkcs1v15Sign, RsaPrivateKey,
};
use serde::{Deserialize, Serialize};
use sha2::{Digest, Sha256};
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
    pub fn new(
        appid: &str,
        privkey_pem: &str,
        address: &str,
        chain: &str,
        url: &str,
    ) -> Result<Self, Box<dyn Error>> {
        let privkey = RsaPrivateKey::from_pkcs8_pem(privkey_pem)?;
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

        if response.result.sign_datas.is_empty() {
            return Err("response sha3 empty".into());
        }

        let sig = hex::decode(&response.result.sign_datas[0].sign[2..])?;
        // println!("==resp: {:#?}", resp);

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
        let signature = self.privkey.sign(Pkcs1v15Sign::new::<Sha256>(), &hashed)?;
        let hex_sig = hex::encode(signature);

        let pubkey = self.privkey.to_public_key().to_public_key_der().unwrap();
        let pubkey_base64 =
            base64::engine::general_purpose::STANDARD.encode(pubkey.as_ref());
        Ok(ReqData { business_data, biz_signature: hex_sig, public_key: pubkey_base64 })
    }

    async fn do_request(&self, url: &str, payload: &ReqData) -> Result<String, Box<dyn Error>> {
        let response = self.client.post(url).json(&payload).send().await?;

        println!("response: {:#?}", response);
        if !response.status().is_success() {
            return Err(format!("response status not ok: {:?}", response.status()).into());
        }

        Ok(response.text().await?)
    }
}

#[tokio::test]
async fn test_sign() -> Result<(), Box<dyn Error>> {
    use ethers::{
        middleware::SignerMiddleware,
        providers::{Http, Middleware, Provider},
        signers::{Signer, Wallet},
    };
    use std::{
        fs::File,
        io::{BufRead, BufReader},
        str::FromStr,
        sync::Arc,
    };

    let path = "../../../../applib/ext_signer_private.key";
    let file = File::open(path).unwrap();
    let reader = BufReader::new(file);

    let mut privkey_base64 = String::new();
    for line in reader.lines() {
        let line_str = line?;
        privkey_base64.push_str(&line_str);
        privkey_base64.push('\n');
    }

    // appid := "morph-tx-submitter-399A1722-3F2C-4E39-ABD2-1B65D02C66BA"
    // rsaPrivStr := ""
    // url := "http://localhost:8080/v1/sign/tx_sign"
    // addr := "0x33d5b507868b7e8ac930cd3bde9eadd60c638479"
    // chain := "QANET-L1"
    // chainid := big.NewInt(900)
    // signer := types.LatestSignerForChainID(chainid)

    let l2_provider = Provider::<Http>::try_from("http://localhost:8545")?;
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let req = Eip1559TransactionRequest::new()
        .to(Address::from_str("0xA081b6A7Fcb129f1381C1128a9A42c486464c61c").unwrap())
        .data(Bytes::new())
        .value(100)
        .from(Address::from_str("0x33d5b507868b7e8ac930cd3bde9eadd60c638479").unwrap());
    //set_from(Address::from_str("0x33d5b507868b7e8ac930cd3bde9eadd60c638479").unwrap())
    let mut tx = TypedTransaction::Eip1559(req);
    l2_signer.fill_transaction(&mut tx, None).await.unwrap();

    let ext_signer: ExternalSign = ExternalSign::new(
        "morph-tx-submitter-399A1722-3F2C-4E39-ABD2-1B65D02C66BA",
        &privkey_base64,
        "0x33d5b507868b7e8ac930cd3bde9eadd60c638479",
        "QANET-L1",
        "http://localhost:8080/v1/sign/tx_sign",
    )
    .unwrap();
    let raw_tx = ext_signer.request_sign(&tx).await.unwrap();
    let pending_tx = l2_provider.send_raw_transaction(raw_tx).await.unwrap();
    pending_tx.await.expect("send_raw_transaction");
    Ok(())
}
