use std::str::FromStr;

pub mod provider;

/// Read an environment variable and parse it to the desired type, or return the default value.
pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}

/// Profile the given code block cycle count.
#[macro_export]
macro_rules! profile {
    ($name:expr, $block:block) => {{
        #[cfg(target_os = "zkvm")]
        {
            println!("cycle-tracker-start: {}", $name);
            let result = (|| $block)();
            println!("cycle-tracker-end: {}", $name);
            result
        }

        #[cfg(not(target_os = "zkvm"))]
        {
            $block
        }
    }};
}

/// Profile the given code block and add the cycle count to the execution report.
#[macro_export]
macro_rules! profile_report {
    ($name:expr, $block:block) => {{
        #[cfg(target_os = "zkvm")]
        {
            println!("cycle-tracker-report-start: {}", $name);
            let result = (|| $block)();
            println!("cycle-tracker-report-end: {}", $name);
            result
        }

        #[cfg(not(target_os = "zkvm"))]
        {
            $block
        }
    }};
}
