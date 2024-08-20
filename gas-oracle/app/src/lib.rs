use ethers::{
    contract::ContractRevert,
    prelude::*,
    types::{transaction::eip2718::TypedTransaction, Eip1559TransactionRequest},
};
use external_sign::ExternalSign;
use eyre::anyhow;
use std::{error::Error, str::FromStr, sync::Arc};
pub mod abi;
mod da_scalar;
mod error;
mod external_sign;
pub mod gas_price_oracle;
mod l1_base_fee;
mod metrics;

use abi::gas_price_oracle_abi::GasPriceOracleErrors;
pub use error::*;
use ethers::{
    contract::ContractError,
    middleware::SignerMiddleware,
    providers::{Http, Provider},
    signers::LocalWallet,
};

pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}

pub fn read_parse_env<T: Clone + FromStr>(var_name: &'static str) -> T {
    let var_value =
        std::env::var(var_name).unwrap_or_else(|_| panic!("Can not read env of {}", var_name));
    match var_value.parse::<T>() {
        Ok(v) => v,
        Err(_) => panic!("Cannot parse {} env var", var_name),
    }
}

pub fn contract_error(e: ContractError<Provider<Http>>) -> String {
    let error_msg = if let Some(contract_err) = e.as_revert() {
        if let Some(data) = GasPriceOracleErrors::decode_with_selector(contract_err.as_ref()) {
            format!("exec error: {:?}", data)
        } else {
            format!("unknown contract error: {:?}", contract_err)
        }
    } else {
        format!("error: {:?}", e)
    };
    error_msg
}

async fn send_transaction(
    calldata: Option<Bytes>,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
    l2_provider: &Provider<Http>,
) -> Result<(), Box<dyn Error>> {
    let req = Eip1559TransactionRequest::new().data(calldata.unwrap_or_default());
    let mut tx = TypedTransaction::Eip1559(req);
    local_signer
        .fill_transaction(&mut tx, None)
        .await
        .map_err(|e| anyhow!("fill_transaction error: {:#?}", e))?;

    let signed_tx =
        sign_tx(tx, local_signer, ext_signer).await.map_err(|e| anyhow!("sign_tx error: {}", e))?;

    let pending_tx = l2_provider
        .send_raw_transaction(signed_tx)
        .await
        .map_err(|e| anyhow!("call contract error: {}", contract_error(e.into())))?;
    pending_tx.await.map_err(|e| anyhow!(format!("check_receipt error: {:#?}", e)))?;
    Ok(())
}

async fn sign_tx(
    tx: TypedTransaction,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
) -> Result<Bytes, Box<dyn Error>> {
    if let Some(signer) = ext_signer {
        Ok(signer.request_sign(&tx).await?)
    } else {
        let signature = local_signer.signer().sign_transaction(&tx).await?;
        Ok(tx.rlp_signed(&signature))
    }
}

/// Minimum gas price for data blobs.
pub const MIN_BLOB_GASPRICE: u64 = 1;

/// Controls the maximum rate of change for blob gas price.
pub const BLOB_GASPRICE_UPDATE_FRACTION: u64 = 3338477;

pub fn calc_blob_basefee(excess_blob_gas: u64) -> u128 {
    fake_exponential(MIN_BLOB_GASPRICE, excess_blob_gas, BLOB_GASPRICE_UPDATE_FRACTION)
}

fn fake_exponential(factor: u64, numerator: u64, denominator: u64) -> u128 {
    assert_ne!(denominator, 0, "attempt to divide by zero");
    let factor = factor as u128;
    let numerator = numerator as u128;
    let denominator = denominator as u128;

    let mut i = 1;
    let mut output = 0;
    let mut numerator_accum = factor * denominator;
    while numerator_accum > 0 {
        output += numerator_accum;

        // Denominator is asserted as not zero at the start of the function.
        numerator_accum = (numerator_accum * numerator) / (denominator * i);
        i += 1;
    }
    output / denominator
}
