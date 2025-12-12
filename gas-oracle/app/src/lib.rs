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