use ethers::{contract::ContractRevert, providers::Middleware};
use std::str::FromStr;
pub mod abi;
mod error;
pub mod gas_price_oracle;
mod l1_base_fee;
mod metrics;

mod da_scalar;
mod external_sign;
mod signer;
use abi::gas_price_oracle_abi::GasPriceOracleErrors;
pub use error::*;
use ethers::contract::ContractError;

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

pub fn contract_error<M: Middleware>(e: ContractError<M>) -> String {
    let error_msg = if let Some(contract_err) = e.as_revert() {
        if let Some(data) = GasPriceOracleErrors::decode_with_selector(contract_err.as_ref()) {
            format!("exec error: {:#?}", data)
        } else {
            format!("unknown contract error: {:#?}", contract_err)
        }
    } else {
        format!("error: {:#?}", e)
    };
    error_msg
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
