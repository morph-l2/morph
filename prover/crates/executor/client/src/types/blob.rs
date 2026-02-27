#[cfg(not(target_os = "zkvm"))]
use {
    alloy_eips::eip2718::Decodable2718,
    alloy_rlp::{Decodable, Header},
    prover_primitives::MorphTxEnvelope,
};
use anyhow::{anyhow, Context};
use ruzstd::StreamingDecoder;
use std::io::Read;

/// This magic number is included at the start of a single Zstandard frame
pub const MAGIC_NUM: u32 = 0xFD2F_B528;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in
/// evaluationform.
pub const BLOB_WIDTH: usize = 4096;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Clone, Debug)]
pub struct BlobData {}

pub fn get_origin_batch(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    // Decode blob, recovering BLS12-381 scalars.
    let mut batch_data = vec![0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
    for i in 0..4096 {
        if blob_data[i * 32] != 0 {
            return Err(anyhow!(format!(
                "Invalid blob, found non-zero high order byte {:x} of field element {}",
                blob_data[i * 32],
                i
            )));
        }
        batch_data[i * 31..i * 31 + 31].copy_from_slice(&blob_data[i * 32 + 1..i * 32 + 32]);
    }
    decompress_batch(&batch_data)
}

pub fn decompress_batch(compressed_batch: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    if compressed_batch.iter().all(|&x| x == 0) {
        // empty batch
        return Ok(Vec::new());
    }

    let mut content = MAGIC_NUM.to_le_bytes().to_vec();
    content.append(&mut compressed_batch.to_vec());
    let mut x = content.as_slice();

    let mut decoder = StreamingDecoder::new(&mut x)?;
    let mut result = Vec::new();
    decoder.read_to_end(&mut result).context("Failed to decompress batch")?;
    #[cfg(not(target_os = "zkvm"))]
    log::info!("decompressed_batch: {:?}", result.len());
    Ok(result)
}

#[cfg(not(target_os = "zkvm"))]
pub fn decode_transactions(bs: &[u8]) -> Vec<MorphTxEnvelope> {
    let mut txs_decoded: Vec<MorphTxEnvelope> = Vec::new();

    let mut offset: usize = 0;
    while offset < bs.len() {
        let first_byte = *bs.get(offset).unwrap();
        if first_byte == 0 {
            // zero byte is found after valid tx bytes, break the loop
            log::info!("zero byte");
            break;
        }

        let tx_len_size = if first_byte > 0xf7 {
            (first_byte - 0xf7) as usize
        } else {
            // Support transaction types: 0x01, 0x02, 0x04, 0x7f
            if first_byte != 0x01 && first_byte != 0x02 && first_byte != 0x04 && first_byte != 0x7f
            {
                log::info!("not supported tx type: 0x{first_byte:02x}");
                break;
            }
            (*bs.get(offset + 1).unwrap() - 0xf7) as usize
        };

        let mut tx_len_bytes = [0u8; 4];
        if first_byte > 0xf7 {
            tx_len_bytes[4 - tx_len_size..]
                .copy_from_slice(bs.get(offset + 1..offset + tx_len_size + 1).unwrap_or_default());
        } else {
            tx_len_bytes[4 - tx_len_size..]
                .copy_from_slice(bs.get(offset + 2..offset + tx_len_size + 2).unwrap_or_default());
        }

        let rlp_tx_len = if first_byte > 0xf7 {
            1 + tx_len_size + u32::from_be_bytes(tx_len_bytes) as usize
        } else {
            2 + tx_len_size + u32::from_be_bytes(tx_len_bytes) as usize
        };

        let tx_bytes = bs[offset..offset + rlp_tx_len].to_vec();
        let tx_decoded = if first_byte == 0x7f {
            // Morph transaction: bypass MorphTxEnvelope::decode_2718 because morph-reth's
            // TxMorph::rlp_decode_fields incorrectly expects a version-byte prefix that is
            // NOT present in the signed (batch) encoding.
            decode_morph_tx_envelope(&tx_bytes)
                .inspect_err(|e| {
                    log::error!("decode morph tx error: {e:?}");
                })
                .unwrap()
        } else {
            MorphTxEnvelope::decode_2718(&mut tx_bytes.as_slice())
                .inspect_err(|e| {
                    log::error!("decode_transaction error: {e:?}");
                })
                .unwrap()
        };

        txs_decoded.push(tx_decoded);
        offset += rlp_tx_len;
    }

    log::info!("Successfully decoded {} transactions", txs_decoded.len());
    txs_decoded
}

/// Decode a morph transaction (type 0x7F) from EIP-2718 encoded bytes.
///
/// This manually decodes the fields + signature instead of going through morph-reth's
/// `TxMorph::rlp_decode_fields` which has a bug: it calls `decode_fields` which tries
/// to interpret the first RLP byte as a version-byte prefix, but the signed encoding
/// does NOT include the version byte. V0 vs V1 is distinguished by the number of fields.
#[cfg(not(target_os = "zkvm"))]
fn decode_morph_tx_envelope(buf: &[u8]) -> Result<MorphTxEnvelope, alloy_rlp::Error> {
    use alloy_primitives::{Bytes, Signature, B256, U256};
    use alloy_eips::eip2930::AccessList;
    use alloy_primitives::TxKind;
    use prover_primitives::TxMorph;
    use prover_primitives::alloy_consensus::SignableTransaction;

    let mut reader = &buf[1..]; // skip type byte (0x7F)

    // Decode the outer RLP list header
    let header = Header::decode(&mut reader)?;
    if !header.list {
        return Err(alloy_rlp::Error::UnexpectedString);
    }
    let payload_end = reader.len() - header.payload_length;

    // Decode common fields (V0 and V1 share these 11 fields)
    let chain_id: u64 = Decodable::decode(&mut reader)?;
    let nonce: u64 = Decodable::decode(&mut reader)?;
    let max_priority_fee_per_gas: u128 = Decodable::decode(&mut reader)?;
    let max_fee_per_gas: u128 = Decodable::decode(&mut reader)?;
    let gas_limit: u128 = Decodable::decode(&mut reader)?;
    let to: TxKind = Decodable::decode(&mut reader)?;
    let value: U256 = Decodable::decode(&mut reader)?;
    let input: Bytes = Decodable::decode(&mut reader)?;
    let access_list: AccessList = Decodable::decode(&mut reader)?;
    let fee_token_id: u16 = Decodable::decode(&mut reader)?;
    let fee_limit: U256 = Decodable::decode(&mut reader)?;

    // Signature is 3 fields: v (bool), r (U256), s (U256).
    // After the common 11 fields, check remaining bytes to determine V0 vs V1:
    // - V0: next 3 fields are the signature (v, r, s)
    // - V1: next 2 fields are reference + memo, then signature (v, r, s)
    //
    // We detect V1 by saving position and trying to read a bool for v.
    // If remaining > 3 fields, it's V1 (reference + memo + sig).
    let saved = reader;
    let _probe: bool = Decodable::decode(&mut { reader })?;
    // For V0, the next field is v (bool: 0x00 or 0x01).
    // For V1, the next field is reference (32 bytes or empty bytes).
    // We can distinguish by checking: if the probe succeeded AND the following
    // two fields (r, s) would consume exactly the remaining bytes → V0.
    // Otherwise → V1.
    //
    // Simpler heuristic: check the first byte of the next field.
    // - v (bool) is encoded as 0x00 or 0x01 (single byte)
    // - reference (B256) is encoded as 0xa0 + 32 bytes, or 0x80 (empty)
    let next_byte = saved[0];
    let (version, reference, memo) = if next_byte == 0x00 || next_byte == 0x01 {
        // This looks like a bool (v field) → V0 format
        (0u8, None, None)
    } else {
        // V1: decode reference and memo
        reader = saved;
        let reference_bytes: Bytes = Decodable::decode(&mut reader)?;
        let reference = if reference_bytes.is_empty() {
            None
        } else if reference_bytes.len() == 32 {
            Some(B256::from_slice(&reference_bytes))
        } else {
            return Err(alloy_rlp::Error::Custom("invalid reference length"));
        };
        let memo_bytes: Bytes = Decodable::decode(&mut reader)?;
        let memo = if memo_bytes.is_empty() { None } else { Some(memo_bytes) };
        (1u8, reference, memo)
    };

    // Decode signature (v, r, s)
    if version != 0 {
        // reader is already advanced past reference + memo
    } else {
        reader = saved;
    }
    let sig_v: bool = Decodable::decode(&mut reader)?;
    let sig_r: U256 = Decodable::decode(&mut reader)?;
    let sig_s: U256 = Decodable::decode(&mut reader)?;

    if reader.len() != payload_end {
        return Err(alloy_rlp::Error::ListLengthMismatch {
            expected: header.payload_length,
            got: header.payload_length - (reader.len() - payload_end),
        });
    }

    let tx = TxMorph {
        chain_id,
        nonce,
        gas_limit,
        max_fee_per_gas,
        max_priority_fee_per_gas,
        to,
        value,
        access_list,
        input,
        fee_token_id,
        fee_limit,
        version,
        reference,
        memo,
    };

    let signature = Signature::new(sig_r, sig_s, sig_v);
    Ok(MorphTxEnvelope::Morph(tx.into_signed(signature)))
}
