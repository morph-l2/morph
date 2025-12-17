// use alloy::consensus::{
//     crypto::RecoveryError,
//     error::{UnsupportedTransactionType, ValueError},
//     transaction::Either,
//     EthereumTxEnvelope, SignableTransaction, Signed, Transaction, TxEip1559, TxEip2930, TxEip7702,
//     TxLegacy, TxType, TypedTransaction,
// };
// use alloy::primitives::{hex, Address, Bytes, Signature, TxKind, B256, U256};

// #[derive(Clone, Debug, alloy::consensus::TransactionEnvelope)]
// #[envelope(
//     tx_type_name = MorphTxType,
//     typed = MorphTypedTransaction,
//     arbitrary_cfg(any(test, feature = "arbitrary")),
//     serde_cfg(feature = "serde")
// )]
// #[doc(alias = "TransactionEnvelope")]
// pub enum MorphTxEnvelope {
//     /// An untagged [`TxLegacy`].
//     #[envelope(ty = 0)]
//     Legacy(Signed<TxLegacy>),
//     /// A [`TxEip2930`] tagged with type 1.
//     #[envelope(ty = 1)]
//     Eip2930(Signed<TxEip2930>),
//     /// A [`TxEip1559`] tagged with type 2.
//     #[envelope(ty = 2)]
//     Eip1559(Signed<TxEip1559>),
//     /// A TxEip4844 tagged with type 3.
//     /// An EIP-4844 transaction has two network representations:
//     /// 1 - The transaction itself, which is a regular RLP-encoded transaction and used to retrieve
//     /// historical transactions..
//     ///
//     /// 2 - The transaction with a sidecar, which is the form used to
//     /// send transactions to the network.
//     /// A [`TxEip7702`] tagged with type 4.
//     #[envelope(ty = 4)]
//     Eip7702(Signed<TxEip7702>),
// }

// impl TryFrom<TxType> for MorphTxType {
//     type Error = UnsupportedTransactionType<TxType>;

//     fn try_from(value: TxType) -> Result<Self, Self::Error> {
//         Ok(match value {
//             TxType::Legacy => Self::Legacy,
//             TxType::Eip2930 => Self::Eip2930,
//             TxType::Eip1559 => Self::Eip1559,
//             TxType::Eip4844 => return Err(UnsupportedTransactionType::new(TxType::Eip4844)),
//             TxType::Eip7702 => Self::Eip7702,
//         })
//     }
// }

// impl TryFrom<MorphTxType> for TxType {
//     type Error = UnsupportedTransactionType<MorphTxType>;

//     fn try_from(value: MorphTxType) -> Result<Self, Self::Error> {
//         Ok(match value {
//             MorphTxType::Legacy => Self::Legacy,
//             MorphTxType::Eip2930 => Self::Eip2930,
//             MorphTxType::Eip1559 => Self::Eip1559,
//             MorphTxType::Eip7702 => Self::Eip7702,
//         })
//     }
// }

// impl MorphTxEnvelope {
//     /// Return the [`TempoTxType`] of the inner txn.
//     pub const fn tx_type(&self) -> MorphTxType {
//         match self {
//             Self::Legacy(_) => MorphTxType::Legacy,
//             Self::Eip2930(_) => MorphTxType::Eip2930,
//             Self::Eip1559(_) => MorphTxType::Eip1559,
//             Self::Eip7702(_) => MorphTxType::Eip7702,
//         }
//     }
// }
