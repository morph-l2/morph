use std::{str::FromStr, sync::Arc};

use crate::{
    abi::gas_price_oracle_abi::GasPriceOracle, calc_blob_basefee, external_sign::ExternalSign, metrics::ORACLE_SERVICE_METRICS, signer::send_transaction, OracleError
};
use ethers::prelude::*;
use eyre::anyhow;

const MAX_BASE_FEE: u128 = 1000 * 10i32.pow(9) as u128; // 1000Gwei
const MAX_BLOB_BASE_FEE: u128 = 100 * 10i32.pow(9) as u128; // 100Gwei

pub struct BaseFeeUpdater {
    l1_provider: Provider<Http>,
    l2_provider: Provider<Http>,
    ext_signer: Option<ExternalSign>,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
    gas_threshold: u64,
    l1_base_fee_buffer: u64,
    l1_blob_base_fee_buffer: u64,
}

impl BaseFeeUpdater {
    pub fn new(
        l1_provider: Provider<Http>,
        l2_provider: Provider<Http>,
        ext_signer: Option<ExternalSign>,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        gas_threshold: u64,
        l1_base_fee_buffer: u64,
        l1_blob_base_fee_buffer: u64,
    ) -> Self {
        BaseFeeUpdater {
            l1_provider,
            l2_provider,
            ext_signer,
            l2_oracle,
            gas_threshold,
            l1_base_fee_buffer,
            l1_blob_base_fee_buffer,
        }
    }

    /// Update baseFee and scalar.
    /// Set the gas data of L1 network to the GasPriceOrale contract on L2.
    pub async fn update(&self) -> Result<(), OracleError> {
        // Step0. Check & Record wallet balance.
        let wallet = if let Some(signer) = &self.ext_signer {
            Address::from_str(&signer.address).unwrap_or_default()
        } else {
            self.l2_oracle.client().address()
        };
        let balance = self.l2_provider.get_balance(wallet, None).await.map_err(|e| {
            OracleError::L1BaseFeeError(anyhow!(format!("l2_wallet.get_balance error: {:#?}", e)))
        })?;

        ORACLE_SERVICE_METRICS
            .gas_oracle_owner_balance
            .set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));

        // Step1. get l1 data.
        let (mut l1_base_fee, mut l1_blob_base_fee, l1_gas_price) =
            query_l1_base_fee(&self.l1_provider).await?;

        if l1_base_fee.is_zero() || l1_blob_base_fee.is_zero() || l1_gas_price.is_zero() {
            return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                "current l1 baseFee or l1_blob_base_fee or gas_price is zero/none"
            ))));
        }

        // Fine tune the actual value
        l1_base_fee += U256::from(self.l1_base_fee_buffer);
        l1_blob_base_fee += U256::from(self.l1_blob_base_fee_buffer);

        #[rustfmt::skip]
        log::debug!("current ethereum baseFee is: {:#?}, l1_blob_base_fee is {:#?}, gas_price is: {:#?}", l1_base_fee,l1_blob_base_fee,l1_gas_price);
        ORACLE_SERVICE_METRICS.l1_base_fee.set(
            ethers::utils::format_units(l1_base_fee, "gwei")
                .unwrap_or(String::from("0"))
                .parse()
                .unwrap_or(0.0),
        );

        // Step2. get l2 data.
        let base_fee_on_l2: U256 = self.l2_oracle.l_1_base_fee().await.map_err(|e| {
            OracleError::L1BaseFeeError(anyhow!(format!(
                "Failed to query query_l1_base_fee on l2: {:#?}",
                e
            )))
        })?;

        let blob_fee_on_l2: U256 = self.l2_oracle.l_1_blob_base_fee().await.map_err(|e| {
            OracleError::L1BaseFeeError(anyhow!(format!(
                "Failed to query l_1_blob_base_fee on l2: {:#?}",
                e
            )))
        })?;

        log::debug!(
            "current l1BaseFee on l2 is: {:#?}, blob_fee_on_l2 on l2 is: {:#?}",
            base_fee_on_l2,
            blob_fee_on_l2
        );
        ORACLE_SERVICE_METRICS.l1_base_fee_on_l2.set(
            ethers::utils::format_units(base_fee_on_l2, "gwei")
                .unwrap_or(String::from("0"))
                .parse()
                .unwrap_or(0.0),
        );
        ORACLE_SERVICE_METRICS.l1_blob_base_fee_on_l2.set(
            ethers::utils::format_units(blob_fee_on_l2, "gwei")
                .unwrap_or(String::from("0"))
                .parse()
                .unwrap_or(0.0),
        );

        self.update_base_fee(l1_base_fee, l1_blob_base_fee, base_fee_on_l2, blob_fee_on_l2).await?;

        Ok(())
    }

    async fn update_base_fee(
        &self,
        mut l1_base_fee: U256,
        mut l1_blob_base_fee: U256,
        base_fee_on_l2: U256,
        blob_fee_on_l2: U256,
    ) -> Result<(), OracleError> {
        l1_base_fee = l1_base_fee.min(MAX_BASE_FEE.into());
        l1_blob_base_fee = l1_blob_base_fee.min(MAX_BLOB_BASE_FEE.into());

        // update_base_fee
        let actual_change = l1_blob_base_fee.abs_diff(blob_fee_on_l2);
        let expected_change = blob_fee_on_l2 * self.gas_threshold / 100;
        let need_update = !l1_blob_base_fee.is_zero() && actual_change > expected_change;

        log::info!(
            "set_l1_base_fee_and_blob_base_fee, blob_fee actual_change: {:?}, expected_change: {:?}, need_update: {:?}", 
            actual_change,
            expected_change,
            need_update
        );

        let client: Arc<SignerMiddleware<Provider<Http>, LocalWallet>> = self.l2_oracle.client();
        if need_update {
            // Update calldata basefee and blob baseFee
            let calldata = self
                .l2_oracle
                .set_l1_base_fee_and_blob_base_fee(l1_base_fee, l1_blob_base_fee)
                .calldata();
            let tx_hash = send_transaction(
                self.l2_oracle.address(),
                calldata,
                &client,
                &self.ext_signer,
                &self.l2_provider,
            )
            .await
            .map_err(|e| {
                OracleError::L1BaseFeeError(anyhow!(format!(
                    "set_l1_base_fee_and_blob_base_fee error: {:#?}",
                    e
                )))
            })?;
            log::info!("set_l1_base_fee_and_blob_base_fee success, tx_hash: {:#?}", tx_hash);

            return Ok(());
        }

        // Only update calldata basefee.
        let actual_change = l1_base_fee.as_u64().abs_diff(base_fee_on_l2.as_u64());
        let expected_change = base_fee_on_l2.as_u64() * self.gas_threshold / 100;
        let need_update = !l1_blob_base_fee.is_zero() && actual_change > expected_change;

        log::info!(
            "set_l1_base_fee, l1BaseFee actual_change: {:?}, expected_change: {:?}, need_update: {:?}",
            actual_change,
            expected_change,
            need_update
        );
        if need_update {
            // Set l1_base_fee for l2.
            let calldata = self.l2_oracle.set_l1_base_fee(l1_base_fee).calldata();
            let tx_hash = send_transaction(
                self.l2_oracle.address(),
                calldata,
                &client,
                &self.ext_signer,
                &self.l2_provider,
            )
            .await
            .map_err(|e| {
                OracleError::L1BaseFeeError(anyhow!(format!("set_l1_base_fee error: {:#?}", e)))
            })?;
            log::info!("set_l1_base_fee success, tx_hash: {:#?}", tx_hash);
        }

        Ok(())
    }
}

async fn query_l1_base_fee(
    l1_provider: &Provider<Http>,
) -> Result<(U256, U256, U256), OracleError> {
    let latest_block = l1_provider
        .get_block(BlockNumber::Latest)
        .await
        .map_err(|e| OracleError::L1BaseFeeError(anyhow!(format!("query_l1_base_fee: {:#?}", e))))?
        .ok_or_else(|| {
            OracleError::L1BaseFeeError(anyhow!(format!("l1 latest block info is none.")))
        })?;

    let l1_base_fee = latest_block.base_fee_per_gas.unwrap_or_default();
    let excess_blob_gas = latest_block.excess_blob_gas.unwrap_or_default();
    let latest_blob_fee = calc_blob_basefee(excess_blob_gas.as_u64());

    let gas_price = match l1_provider.get_gas_price().await {
        Ok(gp) => gp,
        Err(e) => {
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(2);
            return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                "Failed to query l1_gas_price: {:#?}",
                e
            ))))
        }
    };

    Ok((l1_base_fee, U256::from(latest_blob_fee), gas_price))
}
