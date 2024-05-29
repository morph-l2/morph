use std::str::FromStr;

pub mod abi;
pub mod gas_price_oracle;
mod l1_base_fee;
mod metrics;
mod overhead;

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
