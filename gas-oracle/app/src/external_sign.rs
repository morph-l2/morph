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
        let data = self.new_data(&hex::encode(tx.sighash().encode()))?;
        let req_data = self.craft_req_data(data)?;

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

    fn craft_req_data(&self, data: Data) -> Result<ReqData, Box<dyn Error>> {
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
        Ok(ReqData { business_data, biz_signature: hex_sig, public_key: pubkey_base64 })
    }

    async fn do_request(&self, url: &str, payload: &ReqData) -> Result<String, Box<dyn Error>> {
        let response = self.client.post(url).json(&payload).send().await?;
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
    let pem = Pem { tag: String::from("PRIVATE KEY"), contents: key_bytes };

    // Encode the PEM object to a string with standard line width of 64 characters
    let config = EncodeConfig { line_ending: pem::LineEnding::LF };
    encode_config(&pem, config)
}

#[tokio::test]
#[ignore]
async fn test_sign() -> Result<(), Box<dyn Error>> {
    use ethers::{
        middleware::SignerMiddleware,
        providers::{Http, Middleware, Provider},
        signers::{Signer, Wallet},
    };
    use std::{str::FromStr, sync::Arc};
    use crate::read_parse_env;

    // let path = "./../ext_signer_private.key";
    // let file = File::open(path)?;
    // let reader = BufReader::new(file);

    // let mut privkey_base64 = String::new();
    // for line in reader.lines() {
    //     let line_str = line?;
    //     privkey_base64.push_str(&line_str);
    //     privkey_base64.push('\n');
    // }
    dotenv::dotenv().ok();
    let privkey_pem: String = read_parse_env("GAS_ORACLE_EXTERNAL_SIGN_RSA_PRIV");
    // appid := "morph-tx-submitter-399A1722-3F2C-4E39-ABD2-1B65D02C66BA"
    // rsaPrivStr := ""
    // url := "http://localhost:8080/v1/sign/tx_sign"
    // addr := "0x33d5b507868b7e8ac930cd3bde9eadd60c638479"
    // chain := "QANET-L1"
    // chainid := big.NewInt(900)
    // signer := types.LatestSignerForChainID(chainid)
    let ext_signer: ExternalSign = ExternalSign::new(
        "morph-tx-submitter-399A1722-3F2C-4E39-ABD2-1B65D02C66BA",
        &privkey_pem,
        "0x33d5b507868b7e8ac930cd3bde9eadd60c638479",
        "QANET-L1",
        "http://localhost:8080/v1/sign/tx_sign",
    )?;

    let l2_provider = Provider::<Http>::try_from("http://localhost:8545")?;
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")?
            .with_chain_id(l2_provider.get_chainid().await?.as_u64()),
    ));
    let req = Eip1559TransactionRequest::new()
        .to(Address::from_str("0xA081b6A7Fcb129f1381C1128a9A42c486464c61c")?)
        .data(Bytes::new())
        .value(100)
        .from(Address::from_str("0x33d5b507868b7e8ac930cd3bde9eadd60c638479")?);
    let mut tx = TypedTransaction::Eip1559(req);
    l2_signer.fill_transaction(&mut tx, None).await?;

    let raw_tx = ext_signer.request_sign(&tx).await?;
    let pending_tx = l2_provider.send_raw_transaction(raw_tx).await?;
    pending_tx.await.expect("send_raw_transaction");
    Ok(())
}
