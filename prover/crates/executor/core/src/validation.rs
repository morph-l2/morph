use alloy_consensus::BlockHeader as _;
use alloy_eips::eip2718::Encodable2718;
use anyhow::{Result, anyhow, bail, ensure};
use morph_chainspec::{MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK, MorphChainSpec, MorphHardforks};
use morph_primitives::{
    MorphHeader, MorphTxEnvelope,
    transaction::morph_transaction::{MORPH_TX_VERSION_1, MORPH_TX_VERSION_2},
};

use crate::executor::chain_spec_by_chain_id;

/// Validates the L1 queue cursor against the block's leading L1 messages.
/// A missing parent header leaves only the initial cursor unauthenticated.
pub fn validate_next_l1_msg_index(
    chain_id: u64,
    parent_header: Option<&MorphHeader>,
    header: &MorphHeader,
    transactions: &[MorphTxEnvelope],
) -> Result<()> {
    let chain_spec = chain_spec_by_chain_id(chain_id);
    validate_next_l1_msg_index_impl(
        parent_header.map(|parent| parent.next_l1_msg_index),
        header.next_l1_msg_index,
        transactions,
        chain_spec.is_jade_active_at_timestamp(header.timestamp()),
    )
}

fn validate_next_l1_msg_index_impl(
    parent_next: Option<u64>,
    next: u64,
    transactions: &[MorphTxEnvelope],
    is_jade: bool,
) -> Result<()> {
    if let Some(parent_next) = parent_next {
        ensure!(
            next >= parent_next,
            "next L1 message index decreased: expected at least {parent_next}, got {next}"
        );
    }

    let mut expected = parent_next;
    let mut expected_queue_index = None;
    let mut has_l1_message = false;
    for tx in transactions.iter().take_while(|tx| tx.is_l1_msg()) {
        has_l1_message = true;
        let queue_index = tx
            .queue_index()
            .ok_or_else(|| anyhow!("L1 message transaction is missing queue index"))?;
        if is_jade && let Some(expected) = expected_queue_index {
            ensure!(
                queue_index == expected,
                "invalid L1 message queue index: expected {expected}, got {queue_index}"
            );
        }
        let next_queue_index =
            queue_index.checked_add(1).ok_or_else(|| anyhow!("L1 message queue index overflow"))?;
        expected_queue_index = Some(next_queue_index);
        expected = Some(next_queue_index);
    }

    let Some(expected) = expected else {
        return Ok(());
    };
    if is_jade {
        ensure!(next == expected, "invalid next L1 message index: expected {expected}, got {next}");
    } else if has_l1_message {
        ensure!(
            next >= expected,
            "invalid pre-Jade next L1 message index: expected at least {expected}, got {next}"
        );
    }
    Ok(())
}

pub(crate) fn validate_block(
    header: &MorphHeader,
    transactions: &[MorphTxEnvelope],
    chain_spec: &MorphChainSpec,
) -> Result<()> {
    let timestamp = header.timestamp();
    validate_transactions(
        transactions,
        chain_spec.is_emerald_active_at_timestamp(timestamp),
        chain_spec.is_jade_active_at_timestamp(timestamp),
        chain_spec.is_celadon_active_at_timestamp(timestamp),
    )?;
    validate_l2_tx_payload_size(transactions)
}

fn validate_transactions(
    transactions: &[MorphTxEnvelope],
    is_emerald: bool,
    is_jade: bool,
    is_celadon: bool,
) -> Result<()> {
    for tx in transactions {
        let morph_tx = match tx {
            MorphTxEnvelope::Morph(signed) => signed.tx(),
            _ => continue,
        };

        if !is_emerald {
            bail!("MorphTx type is not yet active (emerald fork not reached)");
        }
        if !is_jade && morph_tx.version == MORPH_TX_VERSION_1 {
            bail!("MorphTx version 1 is not yet active (jade fork not reached)");
        }
        if !is_celadon && morph_tx.version == MORPH_TX_VERSION_2 {
            bail!("MorphTx version 2 is not yet active (celadon fork not reached)");
        }

        morph_tx.validate().map_err(|reason| anyhow!(reason.to_owned()))?;
    }
    Ok(())
}

fn validate_l2_tx_payload_size(transactions: &[MorphTxEnvelope]) -> Result<()> {
    let size = transactions
        .iter()
        .filter(|tx| !tx.is_l1_msg())
        .map(|tx| tx.encode_2718_len() as u64)
        .fold(0, u64::saturating_add);
    ensure!(
        size <= MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK,
        "invalid block payload size: {size} exceeds limit {MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK}"
    );
    Ok(())
}

#[cfg(test)]
mod tests {
    use alloy_consensus::{Sealed, Signed};
    use alloy_primitives::{B256, Bytes, Signature, U256};
    use morph_primitives::{
        TxL1Msg, TxMorph,
        transaction::morph_transaction::{
            MORPH_TX_VERSION_0, MORPH_TX_VERSION_1, MORPH_TX_VERSION_2,
        },
    };

    use super::*;

    fn signed(tx: TxMorph) -> MorphTxEnvelope {
        MorphTxEnvelope::Morph(Signed::new_unchecked(tx, Signature::test_signature(), B256::ZERO))
    }

    fn morph_tx(version: u8) -> MorphTxEnvelope {
        signed(TxMorph { version, fee_token_id: 1, ..Default::default() })
    }

    fn l1(queue_index: u64) -> MorphTxEnvelope {
        MorphTxEnvelope::L1Msg(Sealed::new(TxL1Msg { queue_index, ..Default::default() }))
    }

    #[test]
    fn validates_next_l1_message_index_within_a_batch() {
        validate_next_l1_msg_index_impl(Some(5), 7, &[l1(5), l1(6)], true).unwrap();
        validate_next_l1_msg_index_impl(Some(7), 7, &[], true).unwrap();
        validate_next_l1_msg_index_impl(Some(7), 8, &[l1(7)], true).unwrap();

        // The initial parent cursor is explicitly out of scope, but the header must still agree
        // with any L1 messages in the first block.
        validate_next_l1_msg_index_impl(None, 7, &[l1(5), l1(6)], true).unwrap();
        validate_next_l1_msg_index_impl(None, 123, &[], true).unwrap();
    }

    #[test]
    fn allows_jade_l1_message_forward_skip_from_parent() {
        validate_next_l1_msg_index_impl(Some(5), 7, &[l1(6)], true).unwrap();
        validate_next_l1_msg_index_impl(Some(2), 7, &[l1(5), l1(6)], true).unwrap();
    }

    #[test]
    fn rejects_inconsistent_next_l1_message_index() {
        validate_next_l1_msg_index_impl(None, 8, &[l1(5), l1(6)], true).unwrap_err();
        validate_next_l1_msg_index_impl(Some(2), 8, &[l1(5), l1(6)], true).unwrap_err();
        validate_next_l1_msg_index_impl(Some(2), 8, &[l1(5), l1(7)], true).unwrap_err();
        validate_next_l1_msg_index_impl(Some(7), 8, &[], true).unwrap_err();
        validate_next_l1_msg_index_impl(Some(7), 6, &[], true).unwrap_err();
    }

    #[test]
    fn preserves_pre_jade_l1_message_index_leniency() {
        validate_next_l1_msg_index_impl(Some(5), 10, &[l1(5)], false).unwrap();
        validate_next_l1_msg_index_impl(Some(5), 10, &[], false).unwrap();
        validate_next_l1_msg_index_impl(Some(5), 5, &[l1(5)], false).unwrap_err();
    }

    #[test]
    fn accepts_supported_versions_after_activation() {
        for version in [MORPH_TX_VERSION_0, MORPH_TX_VERSION_1, MORPH_TX_VERSION_2] {
            validate_transactions(&[morph_tx(version)], true, true, true).unwrap();
        }
    }

    #[test]
    fn rejects_transactions_before_their_activation() {
        for (version, forks) in [
            (MORPH_TX_VERSION_0, (false, false, false)),
            (MORPH_TX_VERSION_1, (true, false, false)),
            (MORPH_TX_VERSION_2, (true, true, false)),
        ] {
            let (is_emerald, is_jade, is_celadon) = forks;
            validate_transactions(&[morph_tx(version)], is_emerald, is_jade, is_celadon)
                .unwrap_err();
        }
    }

    #[test]
    fn rejects_invalid_version_fields() {
        let v0_without_fee_token =
            signed(TxMorph { version: MORPH_TX_VERSION_0, ..Default::default() });
        validate_transactions(&[v0_without_fee_token], true, true, true).unwrap_err();

        let v1_with_eth_fee_limit = signed(TxMorph {
            version: MORPH_TX_VERSION_1,
            fee_limit: U256::from(1),
            ..Default::default()
        });
        validate_transactions(&[v1_with_eth_fee_limit], true, true, true).unwrap_err();
    }

    #[test]
    fn enforces_payload_limit_only_for_l2_transactions() {
        let oversized = signed(TxMorph {
            version: MORPH_TX_VERSION_1,
            fee_token_id: 1,
            input: Bytes::from(vec![0; MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK as usize + 1]),
            ..Default::default()
        });
        validate_l2_tx_payload_size(&[oversized]).unwrap_err();

        let l1_message = MorphTxEnvelope::L1Msg(Sealed::new(TxL1Msg {
            input: Bytes::from(vec![0; MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK as usize + 1]),
            ..Default::default()
        }));
        validate_l2_tx_payload_size(&[l1_message]).unwrap();
    }
}
