pub mod error;
pub mod executor;
pub mod input;

// 保持对外 API 不变：仍然从 crate 根导出 EthEvm / EthEvmContext
pub use executor::{EthEvm, EthEvmContext};
