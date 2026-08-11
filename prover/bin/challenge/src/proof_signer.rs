use crate::abi::rollup_abi::Rollup;
use crate::external_sign::ExternalSign;
use crate::rollup_compat::authority_is_active;
use ethers::prelude::{BlockId, BlockNumber, LocalWallet, Middleware, Signer};
use ethers::types::{transaction::eip2718::TypedTransaction, Address, Bytes, H256, U64};
use eyre::{eyre, Result};
use std::future::Future;

#[derive(Clone)]
pub enum ApprovedProofSigner {
    Local(LocalWallet),
    External { address: Address, signer: Box<ExternalSign> },
}

impl ApprovedProofSigner {
    pub fn local(wallet: LocalWallet) -> Self {
        Self::Local(wallet)
    }

    pub fn external(address: Address, signer: ExternalSign) -> Self {
        Self::External {
            address,
            signer: Box::new(signer),
        }
    }

    pub fn address(&self) -> Address {
        match self {
            Self::Local(wallet) => wallet.address(),
            Self::External { address, .. } => *address,
        }
    }

    pub async fn sign_transaction(&self, tx: &TypedTransaction) -> Result<Bytes> {
        match self {
            Self::Local(wallet) => {
                let signature = wallet.sign_transaction(tx).await?;
                Ok(tx.rlp_signed(&signature))
            }
            Self::External { signer, .. } => signer.request_sign(tx).await.map_err(|err| eyre!(err.to_string())),
        }
    }
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub struct FixedSnapshot {
    pub number: U64,
    pub hash: H256,
}

#[derive(Clone)]
pub struct SelectedProofSigner {
    pub signer: ApprovedProofSigner,
    pub approved_index: usize,
    pub snapshot: FixedSnapshot,
}

async fn first_active_index<F, Fut, E>(approved: &[Address], mut is_active: F) -> std::result::Result<Option<usize>, E>
where
    F: FnMut(Address) -> Fut,
    Fut: Future<Output = std::result::Result<bool, E>>,
{
    for (index, address) in approved.iter().copied().enumerate() {
        if is_active(address).await? {
            return Ok(Some(index));
        }
    }
    Ok(None)
}

async fn fixed_finalized_snapshot<M: Middleware>(client: &M) -> Result<FixedSnapshot> {
    let block = client
        .get_block(BlockNumber::Finalized)
        .await
        .map_err(|err| eyre!("query finalized L1 snapshot failed: {err}"))?
        .ok_or_else(|| eyre!("finalized L1 block is unavailable"))?;
    let number = block.number.ok_or_else(|| eyre!("finalized L1 block has no number"))?;
    let hash = block.hash.ok_or_else(|| eyre!("finalized L1 block has no hash"))?;
    Ok(FixedSnapshot { number, hash })
}

async fn ensure_snapshot_unchanged<M: Middleware>(client: &M, snapshot: FixedSnapshot) -> Result<()> {
    let current = client
        .get_block(snapshot.number)
        .await
        .map_err(|err| eyre!("recheck finalized L1 snapshot failed: {err}"))?
        .ok_or_else(|| eyre!("finalized L1 snapshot disappeared during signer selection"))?;
    if current.hash != Some(snapshot.hash) {
        return Err(eyre!("finalized L1 snapshot changed during signer selection"));
    }
    Ok(())
}

/// Selects the first approved signer that is active at one immutable finalized
/// block hash. The caller must treat `None` and every error as a hard no-send.
pub async fn select_active_proof_signer<M: Middleware + 'static>(
    rollup: &Rollup<M>,
    approved: &[ApprovedProofSigner],
) -> Result<Option<SelectedProofSigner>> {
    if approved.is_empty() {
        return Err(eyre!("no approved proof signers are configured"));
    }

    let snapshot = fixed_finalized_snapshot(rollup.client().as_ref()).await?;
    let block_id = BlockId::Hash(snapshot.hash);
    let addresses = approved.iter().map(ApprovedProofSigner::address).collect::<Vec<_>>();
    let approved_index = first_active_index(&addresses, |address| async move {
        authority_is_active(rollup, address, block_id)
            .await
            .map_err(|err| eyre!("active check failed for approved proof signer {address:?}: {err}"))
    })
    .await?;
    ensure_snapshot_unchanged(rollup.client().as_ref(), snapshot).await?;
    Ok(approved_index.map(|approved_index| SelectedProofSigner {
        signer: approved[approved_index].clone(),
        approved_index,
        snapshot,
    }))
}

#[cfg(test)]
mod tests {
    use super::first_active_index;
    use ethers::types::Address;
    use std::convert::Infallible;
    use std::future::ready;

    #[tokio::test]
    async fn primary_inactivation_selects_first_active_backup() {
        let primary = Address::from_low_u64_be(1);
        let backup = Address::from_low_u64_be(2);
        let approved = [primary, backup];

        let initial = first_active_index(&approved, |_| ready(Ok::<_, Infallible>(true))).await.unwrap();
        let after_primary_inactivation = first_active_index(&approved, |address| ready(Ok::<_, Infallible>(address == backup)))
            .await
            .unwrap();

        assert_eq!(initial, Some(0));
        assert_eq!(after_primary_inactivation, Some(1));
    }

    #[tokio::test]
    async fn all_inactive_signers_fail_closed_without_dispatch() {
        let approved = [Address::from_low_u64_be(1), Address::from_low_u64_be(2)];
        let selected = first_active_index(&approved, |_| ready(Ok::<_, Infallible>(false))).await.unwrap();
        let mut sent_transactions = 0;
        if selected.is_some() {
            sent_transactions += 1;
        }

        assert_eq!(selected, None);
        assert_eq!(sent_transactions, 0);
    }
}
