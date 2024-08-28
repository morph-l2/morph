use ethers::{
    middleware::SignerMiddleware,
    prelude::*,
    providers::{Http, Provider},
    signers::LocalWallet,
    types::{transaction::eip2718::TypedTransaction, Eip1559TransactionRequest},
};
use eyre::anyhow;
use std::{error::Error, str::FromStr, sync::Arc};

use crate::{contract_error, external_sign::ExternalSign};

pub async fn send_transaction(
    contract: Address,
    calldata: Option<Bytes>,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
    l2_provider: &Provider<Http>,
) -> Result<H256, Box<dyn Error>> {
    let req = Eip1559TransactionRequest::new().data(calldata.unwrap_or_default());
    let mut tx = TypedTransaction::Eip1559(req);
    tx.set_to(contract);
    if let Some(signer) = ext_signer {
        tx.set_from(Address::from_str(&signer.address).unwrap_or_default());
    } else {
        tx.set_from(local_signer.address());
    }
    local_signer.fill_transaction(&mut tx, None).await.map_err(|e| {
        let msg = contract_error(
            ContractError::<SignerMiddleware<Provider<Http>, LocalWallet>>::from_middleware_error(
                e,
            ),
        );
        anyhow!("fill_transaction error: {:#?}", msg)
    })?;

    let signed_tx = sign_tx(&mut tx, local_signer, ext_signer)
        .await
        .map_err(|e| anyhow!("sign_tx error: {}", e))?;

    let pending_tx = l2_provider.send_raw_transaction(signed_tx).await.map_err(|e| {
        let msg = contract_error(ContractError::<Provider<Http>>::from(e));
        anyhow!("call contract error: {}", msg)
    })?;
    let tx_hash = pending_tx.tx_hash();

    let receipt = pending_tx
        .await
        .map_err(|e| anyhow!(format!("check_receipt of {:#?} is error: {:#?}", tx_hash, e)))?
        .ok_or(anyhow!(format!("check_receipt is none, tx_hash: {:#?}", tx_hash)))?;

    if receipt.status == Some(1.into()) {
        Ok(tx_hash)
    } else {
        Err(anyhow!(format!("tx exec failed, transaction_hash: {:#?}", receipt.transaction_hash))
            .into())
    }
}

async fn sign_tx(
    tx: &mut TypedTransaction,
    local_signer: &Arc<SignerMiddleware<Provider<Http>, LocalWallet>>,
    ext_signer: &Option<ExternalSign>,
) -> Result<Bytes, Box<dyn Error>> {
    if let Some(signer) = ext_signer {
        log::info!("request ext sign");
        Ok(signer.request_sign(tx).await?)
    } else {
        log::info!("request local sign");
        let signature = local_signer.signer().sign_transaction(tx).await?;
        Ok(tx.rlp_signed(&signature))
    }
}
