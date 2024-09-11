use sbv_core::VerificationError;
use sbv_primitives::{types::BlockTrace, B256};

pub mod blob;
pub mod evm;

// struct Verifier = Box<dyn Fn(&BlockTrace) -> Result<B256, VerificationError>>;

pub trait Verifier {
    fn verify(request: &BlockTrace) -> Result<B256, VerificationError>;
}
