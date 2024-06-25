// update errors
#[derive(Debug, thiserror::Error)]
pub enum ScalarError {
    #[error("{0}")]
    Error(eyre::Error),
    #[error("{0}")]
    CalculateError(eyre::Error),
    #[error("{0}")]
    BeaconNodeError(eyre::Error),
}

