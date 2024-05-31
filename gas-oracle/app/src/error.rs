// update errors
#[derive(Debug, thiserror::Error)]
pub enum OracleError {
    #[error("{0}")]
    Error(eyre::Error),
    #[error("{0}")]
    L1BaseFeeError(eyre::Error),
}
