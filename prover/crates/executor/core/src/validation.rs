use alloy_consensus::BlockHeader as _;
use alloy_eips::eip2718::Encodable2718;
use anyhow::{Result, anyhow, bail, ensure};
use morph_chainspec::{MORPH_MAX_TX_PAYLOAD_BYTES_PER_BLOCK, MorphChainSpec, MorphHardforks};
use morph_primitives::{
    Block, MorphTxEnvelope,
    transaction::morph_transaction::{MORPH_TX_VERSION_1, MORPH_TX_VERSION_2},
};

/// Checks Morph transaction rules and the L2 payload size before execution.
pub fn validate_block_pre_execution(block: &Block, chain_spec: &MorphChainSpec) -> Result<()> {
    let timestamp = block.header.timestamp();
    let transactions = &block.body.transactions;
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
