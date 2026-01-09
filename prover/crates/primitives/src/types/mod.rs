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

#[cfg(test)]
mod tests {
    use super::*;
    use alloy_primitives::*;

    const TRACE: &str = include_str!("../../../../testdata/dev.json");

    #[test]
    fn test_deserialize() {
        // NOTE: testdata/dev.json is a nested list: [[ { ..block trace.. } ]]
        let trace = serde_json::from_str::<serde_json::Value>(TRACE).unwrap()[0][0].clone();

        let coinbase: Coinbase = serde_json::from_value(trace["coinbase"].clone()).unwrap();
        assert_eq!(coinbase.address, address!("5300000000000000000000000000000000000005"));

        let header: BlockHeader = serde_json::from_value(trace["header"].clone()).unwrap();
        assert_eq!(header.number, U256::from(8370400));
        assert_eq!(
            header.hash,
            b256!("3aec6d882b0548a6f073d5aed65cfa527809ea71528d5e387a9a37436f0c6f9c")
        );
        assert_eq!(header.timestamp, U256::from(0x66bcc0a0));
        assert_eq!(header.gas_limit, U256::from(0x989680));
        assert_eq!(header.base_fee_per_gas, Some(U256::from(0x3c9d282)));
        assert_eq!(header.difficulty, U256::from(0x2));
        assert_eq!(
            header.mix_hash,
            Some(b256!("0000000000000000000000000000000000000000000000000000000000000000"))
        );

        let transactions: Vec<TransactionTrace> =
            serde_json::from_value(trace["transactions"].clone()).unwrap();
        assert_eq!(transactions[0].ty, 2);
        assert_eq!(transactions[1].nonce, 52786);
        assert_eq!(
            transactions[2].tx_hash,
            b256!("66fc61ff7dd747503aed97a04ca47f00534d914a3b242ea4b2c75b66d720d4b8")
        );
        assert_eq!(transactions[3].gas, 53340);
        assert_eq!(transactions[4].gas_price, U256::from(0x48c1a5b));
        assert_eq!(transactions[5].gas_tip_cap, Some(U256::from(0x416e)));
        assert_eq!(transactions[6].gas_fee_cap, Some(U256::from(0x4f4f99a)));
        assert_eq!(transactions[7].from, address!("32a29339a23afff0febd75bef656b9bf32967085"));
        assert_eq!(transactions[8].to, Some(address!("a2a9fd768d482caf519d749d3123a133db278a66")),);
        assert_eq!(transactions[9].chain_id, U64::from(0x82750));
        assert_eq!(transactions[10].value, uint!(0x297ee5eafa233c_U256));
        assert_eq!(transactions[11].data, bytes!("21c69a19000000000000000000000000f610a9dfb7c89644979b4a0f27063e9e7d7cda320000000000000000000000000000000000000000000000000086e6edc73882e2000000000000000000000000956df8424b556f0076e8abf5481605f5a791cc7f000000000000000000000000956df8424b556f0076e8abf5481605f5a791cc7f00000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000244627dd56a000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001ea01a200020000000000000000000000000000000000000000000000000086e6edc73882e2006018df4b145f074d63efa24dbd61dd00da2cdb3697f610a9dfb7c89644979b4a0f27063e9e7d7cda3200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066bcce7c000000000000000000000000000000000000000000000000009b579427d194d20000000000000000000000000000000000000000000000000086e6edc73882e202007e07060300e6000000000000000000000000fffd8963efd1fc6a506488495d951d5263988d2500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000014010037060300920300e603010e03012e03014e03007e6e6cc7163fa93a693ee8491bb2b01656ba5ea1f3070a00000000000000000000000000000000000000000000000000000000000003019805000002007e02018405004053000000000000000000000000000000000000040201c705004002009202006a050040004802008000000106010e00000000e40040016e01840184070040002001b801be0000010060000001be01c70000050040000001db01e10000030060000001e101ea00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"));
        assert!(!transactions[12].is_create);
        assert_eq!(transactions[13].access_list.len(), 0);
        assert_eq!(transactions[14].v.to::<u64>(), 0x1);
        assert_eq!(
            transactions[15].r,
            uint!(0xb6c02ca3114e6eb88b49e122a24903579112bdf93a615a145949c60c0d5f0b4b_U256)
        );
        assert_eq!(
            transactions[15].s,
            uint!(0x6989f79234aa582bbe1efbd6b5d5dcb1ba0247bcc3b641bc2ef1573b330254cf_U256)
        );

        let _codes: Vec<BytecodeTrace> = serde_json::from_value(trace["codes"].clone()).unwrap();

        let storage_trace: StorageTrace =
            serde_json::from_value(trace["storageTrace"].clone()).unwrap();
        assert_eq!(
            storage_trace.root_before,
            b256!("1cee7c2b120a46f498630029b2463d67e579c27d49ebedcb82f31960d5c03a7c")
        );
        assert_eq!(
            storage_trace.root_after,
            b256!("2f6af5a76ddd2fcd78b2f72e39282a782bfa30625b7d8d8f7506603b7511ba38")
        );

        let _block: BlockTrace = serde_json::from_value(trace).unwrap();
    }
}
