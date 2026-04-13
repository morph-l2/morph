use std::str::FromStr;

pub mod provider;

/// Read an environment variable and parse it to the desired type, or return the default value.
pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}
