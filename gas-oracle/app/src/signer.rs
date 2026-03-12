use ethers::{
    abi::FunctionExt,
    contract::builders::ContractCall,
    middleware::SignerMiddleware,
    prelude::*,
    providers::{Http, Provider},
    signers::LocalWallet,
    types::transaction::eip2718::TypedTransaction,
};
use eyre::anyhow;
use remote_signer_client::SignerClient;
use std::{error::Error, str::FromStr, sync::Arc};

use crate::{contract_error, read_env_var};

/// Send a contract call as a transaction.
///
/// `call` carries both the ABI method info (for extracting the method signature)
/// and the encoded calldata. The method signature is forwarded to the remote
/// signer so it can apply policy checks.
pub async fn send_transaction<M, D>(
    call: ContractCall<M, D>,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<SignerClient>,
    l2_provider: &Provider<Http>,
) -> Result<H256, Box<dyn Error>>
where
    M: ethers::providers::Middleware,
    D: ethers::abi::Detokenize,
{
    let contract = match call.tx.to() {
        Some(NameOrAddress::Address(addr)) => *addr,
        _ => return Err(anyhow!("send_transaction: contract address not set in call").into()),
    };
    let method_sig = call.function.abi_signature();
    let calldata = call.calldata();

    // Estimate eip1559_fees
    let gas_data = local_signer
        .estimate_eip1559_fees(Some(eip1559_estimator))
        .await
        .map_err(|e| anyhow!(format!("estimate_eip1559_fees error: {:#?}", e)))?;
    let req = Eip1559TransactionRequest::new()
        .data(calldata.unwrap_or_default())
        .max_fee_per_gas(gas_data.0)
        .max_priority_fee_per_gas(gas_data.1);
    let mut tx = TypedTransaction::Eip1559(req);
    tx.set_to(contract);
    if let Some(signer) = ext_signer {
        tx.set_from(Address::from_str(&signer.address).unwrap_or_default());
    } else {
        tx.set_from(local_signer.address());
    }

    // Fill nonce, gas
    local_signer.fill_transaction(&mut tx, None).await.map_err(|e| {
        let msg = contract_error(
            ContractError::<SignerMiddleware<Provider<Http>, LocalWallet>>::from_middleware_error(
                e,
            ),
        );
        anyhow!("fill_transaction error: {:#?}", msg)
    })?;

    // Sign and send
    let signed_tx = sign_tx(&mut tx, local_signer, ext_signer, &method_sig)
        .await
        .map_err(|e| anyhow!("sign_tx error: {}", e))?;

    let pending_tx = l2_provider.send_raw_transaction(signed_tx).await.map_err(|e| {
        let msg = contract_error(ContractError::<Provider<Http>>::from(e));
        anyhow!("call contract error: {}", msg)
    })?;
    let tx_hash = pending_tx.tx_hash();

    let receipt = pending_tx
        .await
        .map_err(|e| anyhow!(format!("check_receipt of {:#?} is error: {:#?}", tx_hash, e)))?
        .ok_or(anyhow!(format!("check_receipt is none, tx_hash: {:#?}", tx_hash)))?;

    if receipt.status == Some(1.into()) {
        Ok(tx_hash)
    } else {
        Err(anyhow!(format!("tx exec failed, transaction_hash: {:#?}", receipt.transaction_hash))
            .into())
    }
}

async fn sign_tx(
    tx: &mut TypedTransaction,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<SignerClient>,
    method_sig: &str,
) -> Result<Bytes, Box<dyn Error>> {
    if let Some(signer) = ext_signer {
        log::info!("request remote sign, method_sig: {}", method_sig);
        Ok(signer.sign(tx, method_sig).await?)
    } else {
        log::info!("request local sign");
        let signature = local_signer.signer().sign_transaction(tx).await?;
        Ok(tx.rlp_signed(&signature))
    }
}

//Fee estimator
lazy_static::lazy_static! {
    static ref EIP1559_FEE_ESTIMATION_MAX_FEE: u64= read_env_var("EIP1559_FEE_ESTIMATION_MAX_FEE", 200_000_000_000);
}

const PRIORITY_FEE_PER_GAS_WEI: u64 = 1_000_000; // 0.001 Gwei
fn eip1559_estimator(base_fee_per_gas: U256, _rewards: Vec<Vec<U256>>) -> (U256, U256) {
    let max_priority_fee_per_gas = U256::from(PRIORITY_FEE_PER_GAS_WEI);
    let max_fee_per_gas = std::cmp::min(
        U256::from(*EIP1559_FEE_ESTIMATION_MAX_FEE),
        base_fee_per_gas + max_priority_fee_per_gas,
    );

    (max_fee_per_gas, max_priority_fee_per_gas)
}



#[tokio::test]
#[ignore]
async fn test_estimate_eip1559_fees() {
    dotenv::dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    let l2_provider = Provider::<Http>::try_from("https://rpc-quicknode.morphl2.io").unwrap();
    let l2_signer = Arc::new(SignerMiddleware::new(
        l2_provider.clone(),
        Wallet::from_str("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
            .unwrap()
            .with_chain_id(l2_provider.get_chainid().await.unwrap().as_u64()),
    ));
    let data_defuault = l2_signer.estimate_eip1559_fees(None).await.unwrap();
    let data = l2_signer.estimate_eip1559_fees(Some(eip1559_estimator)).await.unwrap();
    log::info!("====gas data_defuault: {:?}, {:?}", data_defuault.0, data_defuault.1);
    log::info!("====gas data: {:?}, {:?}", data.0, data.1);
    let req = Eip1559TransactionRequest::new()
        .to(Address::from_str("0x099f9e4ecc7fb2b4fd759ce0c2c2c3072b77e9bc").unwrap())
        .from(Address::from_str("0x523bff68043C818e9b449dd3Bee8ecCfa85D7E50").unwrap())
        .max_fee_per_gas(data.0)
        .max_priority_fee_per_gas(data.1);
    let mut tx = TypedTransaction::Eip1559(req);
    l2_signer.fill_transaction(&mut tx, None).await.unwrap();

    let tx_typed = tx.as_eip1559_mut().unwrap();
    log::info!(
        "====tx_typed gas data: {:?}, {:?}",
        tx_typed.max_fee_per_gas, tx_typed.max_priority_fee_per_gas
    );
}
