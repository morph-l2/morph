use crate::{Block, TxTrace};
use alloy_primitives::{Address, Bytes, B256, U256, U64};
use morph_primitives::MorphTxEnvelope;
use serde::{Deserialize, Serialize};
use serde_with::{serde_as, Map};

mod authorization_list;
/// Blob related types
pub mod blob;
/// L2 Block
pub mod block;
mod tx;
/// Alternative fee transaction types
pub use authorization_list::{ArchivedSignedAuthorization, AuthorizationList};
pub use tx::TransactionTrace;

/// Block header
#[derive(Serialize, Deserialize, Default, Debug, Clone, Hash, PartialEq, Eq)]
pub struct BlockHeader {
    /// block number
    pub number: U256,
    /// block hash
    pub hash: B256,
    /// the start index of L1Message needs to be processed
    #[serde(default, rename = "nextL1MsgIndex")]
    pub next_l1_msg_index: U64,
    /// state root
    #[serde(rename = "stateRoot")]
    pub state_root: B256,
    /// timestamp
    pub timestamp: U256,
    /// gas limit
    #[serde(rename = "gasLimit")]
    pub gas_limit: U256,
    /// base fee per gas
    #[serde(rename = "baseFeePerGas")]
    pub base_fee_per_gas: Option<U256>,
    /// difficulty
    pub difficulty: U256,
    /// mix hash
    #[serde(rename = "mixHash")]
    pub mix_hash: Option<B256>,
}

/// Coinbase
#[derive(Serialize, Deserialize, Default, Debug, Clone, Hash, PartialEq, Eq)]
struct Coinbase {
    /// address of coinbase
    address: Address,
}

/// Bytecode trace
#[derive(Serialize, Deserialize, Default, Debug, Clone, Hash)]
struct BytecodeTrace {
    /// bytecode
    code: Bytes,
}

/// storage trace
#[serde_as]
#[derive(Serialize, Deserialize, Default, Debug, Clone, Eq, PartialEq, Hash)]
pub struct StorageTrace {
    /// root before
    #[serde(rename = "rootBefore")]
    root_before: B256,
    /// root after
    #[serde(rename = "rootAfter")]
    root_after: B256,
    /// account proofs
    #[serde_as(as = "Option<Map<_, _>>")]
    pub proofs: Option<Vec<(Address, Vec<Bytes>)>>,
    #[serde(rename = "storageProofs", default)]
    #[serde_as(as = "Map<_,Map<_, _>>")]
    /// storage proofs for each account
    #[allow(clippy::type_complexity)]
    pub storage_proofs: Vec<(Address, Vec<(B256, Vec<Bytes>)>)>,
    /// proofs
    #[serde(rename = "flattenProofs")]
    #[serde_as(as = "Option<Map<_, _>>")]
    pub flatten_proofs: Option<Vec<(B256, Bytes)>>,
}

/// Block trace format
#[derive(Serialize, Deserialize, Default, Debug, Clone)]

pub struct BlockTrace {
    /// chain id
    #[serde(rename = "chainID", default)]
    chain_id: u64,
    /// coinbase
    coinbase: Coinbase,
    /// block
    pub header: BlockHeader,
    /// txs
    pub transactions: Vec<TransactionTrace>,
    /// bytecodes
    codes: Vec<BytecodeTrace>,
    /// storage trace BEFORE execution
    #[serde(rename = "storageTrace")]
    pub storage_trace: StorageTrace,
    //d tx_storage_trace
    /// l1 tx queue
    #[serde(rename = "startL1QueueIndex", default)]
    start_l1_queue_index: u64,
}

impl BlockTrace {
    /// Returns the typed transactions in the block.
    pub fn typed_transactions(&self) -> Vec<MorphTxEnvelope> {
        self.transactions.iter().map(|tx_trace| tx_trace.try_build_tx_envelope().unwrap()).collect()
    }
}

impl Block for BlockTrace {
    type Tx = TransactionTrace;

    fn number(&self) -> u64 {
        self.header.number.to()
    }
    fn block_hash(&self) -> B256 {
        self.header.hash
    }
    fn chain_id(&self) -> u64 {
        self.chain_id
    }
    fn coinbase(&self) -> Address {
        self.coinbase.address
    }
    fn timestamp(&self) -> U256 {
        self.header.timestamp
    }

    fn gas_limit(&self) -> U256 {
        self.header.gas_limit
    }

    fn base_fee_per_gas(&self) -> Option<U256> {
        self.header.base_fee_per_gas
    }

    fn difficulty(&self) -> U256 {
        self.header.difficulty
    }

    fn prevrandao(&self) -> Option<B256> {
        self.header.mix_hash
    }

    fn transactions(&self) -> impl Iterator<Item = &Self::Tx> {
        self.transactions.iter()
    }

    fn root_before(&self) -> B256 {
        self.storage_trace.root_before
    }

    fn root_after(&self) -> B256 {
        self.storage_trace.root_after
    }

    fn codes(&self) -> impl ExactSizeIterator<Item = &[u8]> {
        self.codes.iter().map(|code| code.code.as_ref())
    }

    fn start_l1_queue_index(&self) -> u64 {
        self.start_l1_queue_index
    }

    fn flatten_proofs(&self) -> impl Iterator<Item = (&B256, &[u8])> {
        self.storage_trace
            .flatten_proofs
            .as_ref()
            .into_iter()
            .flat_map(|proof| proof.iter().map(|(k, v)| (k, v.as_ref())))
    }
}
