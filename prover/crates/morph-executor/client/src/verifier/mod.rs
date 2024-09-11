use sbv_core::VerificationError;
use sbv_primitives::{types::BlockTrace, B256};

pub mod blob_verifier;
pub mod evm_verifier;

// struct Verifier = Box<dyn Fn(&BlockTrace) -> Result<B256, VerificationError>>;

// pub trait Verifier {
//     fn verify(request: &BlockTrace) -> Result<B256, VerificationError>;
// }
