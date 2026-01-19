use alloy_primitives::{Address, Bytes, B256, U256, U64};
use alloy_rpc_types::EIP1186AccountProofResponse;
use alloy_rpc_types::EIP1186StorageProof;
use reth_primitives_traits::Account;
use reth_trie::{AccountProof, StorageProof, EMPTY_ROOT_HASH};

/// Converts an [EIP1186AccountProofResponse] to an [AccountProof].
pub fn eip1186_proof_to_account_proof(proof: EIP1186AccountProofResponse) -> AccountProof {
    let address = proof.address;
    let balance = proof.balance;
    let code_hash = proof.code_hash;
    let storage_root = proof.storage_hash;
    let account_proof = proof.account_proof;
    let storage_proofs = proof
        .storage_proof
        .into_iter()
        .map(|storage_proof| {
            let key = storage_proof.key;
            let value = storage_proof.value;
            let proof = storage_proof.proof;
            let mut sp = StorageProof::new(key.as_b256());
            sp.value = value;
            sp.proof = proof;
            sp
        })
        .collect();

    let (storage_root, info) =
        if proof.nonce == 0 && balance.is_zero() && storage_root.is_zero() && code_hash.is_zero() {
            // Account does not exist in state. Return `None` here to prevent proof verification.
            (EMPTY_ROOT_HASH, None)
        } else {
            (
                storage_root,
                Some(Account { nonce: proof.nonce, balance, bytecode_hash: code_hash.into() }),
            )
        };

    AccountProof { address, info, proof: account_proof, storage_root, storage_proofs }
}

#[derive(Clone, Debug, Default, PartialEq, Eq, serde::Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct EIP1186AccountProofResponseCompat {
    /// The account address.
    pub address: Address,
    /// The account balance.
    pub balance: U256,
    /// The hash of the code of the account.
    #[serde(alias = "keccakCodeHash")]
    pub code_hash: B256,
    /// The account nonce.
    pub nonce: U64,
    /// The hash of the storage of the account.
    pub storage_hash: B256,
    /// The account proof.
    pub account_proof: Vec<Bytes>,
    /// The storage proof.
    pub storage_proof: Vec<EIP1186StorageProof>,
}

impl From<EIP1186AccountProofResponseCompat> for EIP1186AccountProofResponse {
    fn from(v: EIP1186AccountProofResponseCompat) -> Self {
        Self {
            address: v.address,
            balance: v.balance,
            code_hash: v.code_hash,
            nonce: v.nonce.to::<u64>(),
            storage_hash: v.storage_hash,
            account_proof: v.account_proof,
            storage_proof: v.storage_proof,
        }
    }
}
