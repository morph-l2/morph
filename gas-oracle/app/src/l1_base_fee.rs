use eyre::anyhow;

use crate::{
    abi::gas_price_oracle_abi::GasPriceOracle, metrics::ORACLE_SERVICE_METRICS, OracleError,
};
use ethers::prelude::*;

static DEFAULT_SCALAR: f64 = 1000000000.0;
const MAX_BASE_FEE: u128 = 1000 * 10i32.pow(9) as u128; // 1000Gwei
const MAX_SCALAR_RATIO: f64 = 100.0;

pub struct BaseFeeUpdater {
    l1_provider: Provider<Http>,
    l2_provider: Provider<Http>,
    l2_wallet: Address,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
    gas_threshold: u128,
}

impl BaseFeeUpdater {
    pub fn new(
        l1_provider: Provider<Http>,
        l2_provider: Provider<Http>,
        l2_wallet: Address,
        l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
        gas_threshold: u128,
    ) -> Self {
        BaseFeeUpdater { l1_provider, l2_provider, l2_wallet, l2_oracle, gas_threshold }
    }

    /// Update baseFee and scalar.
    /// Set the gas data of L1 network to the GasPriceOrale contract on L2.
    pub async fn update(&self) -> Result<(), OracleError> {
        // Step1. get l1 data.
        let (l1_base_fee, l1_gas_price) = query_l1_base_fee(&self.l1_provider).await?;
        let l1_base_fee = l1_base_fee.unwrap_or_default();
        let l1_gas_price = l1_gas_price.unwrap_or_default();

        if l1_base_fee.is_zero() || l1_gas_price.is_zero() {
            return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                "current ethereum baseFee or gas_price is zero"
            ))));
        }

        #[rustfmt::skip]
        log::debug!("current ethereum baseFee is: {:#?}, gas_price is: {:#?}", l1_base_fee, l1_gas_price);
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

        let scalar: U256 = self.l2_oracle.scalar().await.map_err(|e| {
            OracleError::L1BaseFeeError(anyhow!(format!("Failed to query scalar on l2: {:#?}", e)))
        })?;

        #[rustfmt::skip]
        log::info!("current l1BaseFee on l2 is: {:#?}, scalar is: {:#?}", base_fee_on_l2, scalar);
        ORACLE_SERVICE_METRICS.l1_base_fee_on_l2.set(
            ethers::utils::format_units(base_fee_on_l2, "gwei")
                .unwrap_or(String::from("0"))
                .parse()
                .unwrap_or(0.0),
        );

        self.update_base_fee(l1_base_fee, base_fee_on_l2).await?;
        self.update_scalar(l1_gas_price, l1_base_fee, scalar).await?;

        // Step4. Record wallet balance.
        let balance = self.l2_provider.get_balance(self.l2_wallet, None).await.map_err(|e| {
            OracleError::L1BaseFeeError(anyhow!(format!("l2_wallet.get_balance error: {:#?}", e)))
        })?;

        ORACLE_SERVICE_METRICS
            .gas_oracle_owner_balance
            .set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));
        Ok(())
    }

    async fn update_base_fee(
        &self,
        mut l1_base_fee: U256,
        base_fee_on_l2: U256,
    ) -> Result<(), OracleError> {
        l1_base_fee = l1_base_fee.min(U256::from(MAX_BASE_FEE));

        let actual_change = l1_base_fee.as_u128().abs_diff(base_fee_on_l2.as_u128());
        let expected_change = base_fee_on_l2.as_u128() * self.gas_threshold / 100;
        log::info!(
            "l1BaseFee actual_change: {:#?}, expected_change: {:#?}",
            actual_change,
            expected_change
        );

        //min gas price
        if l1_base_fee > U256::from(0) && actual_change > expected_change {
            // Set l1_base_fee for l2.
            let tx = self.l2_oracle.set_l1_base_fee(l1_base_fee).legacy();
            let rt = tx.send().await;
            let pending_tx = match rt {
                Ok(pending) => {
                    log::info!("tx of set_l1_base_fee has been sent: {:#?}", pending.tx_hash());
                    pending
                }
                Err(e) => {
                    return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                        "set_l1_base_fee error: {:#?}",
                        e
                    ))));
                }
            };
            pending_tx.await.map_err(|e| {
                OracleError::L1BaseFeeError(anyhow!(format!(
                    "set_l1_base_fee check_receipt error: {:#?}",
                    e
                )))
            })?;
        }

        Ok(())
    }

    async fn update_scalar(
        &self,
        l1_gas_price: U256,
        l1_base_fee: U256,
        current_scalar: U256,
    ) -> Result<(), OracleError> {
        if l1_gas_price < l1_base_fee {
            return Err(OracleError::L1BaseFeeError(anyhow!("l1_gas_price < l1_base_fee")));
        }
        let mut scalar_ratio_from_l1 = l1_gas_price.as_u128() as f64 / l1_base_fee.as_u128() as f64;
        scalar_ratio_from_l1 = scalar_ratio_from_l1.min(MAX_SCALAR_RATIO);
        #[rustfmt::skip]
        ORACLE_SERVICE_METRICS.scalar_ratio.set(format!("{:.2}", scalar_ratio_from_l1).parse().unwrap_or(0.00));

        let scalar_ratio_from_l2 = current_scalar.as_u128() as f64 / DEFAULT_SCALAR;
        let scalar_diff = scalar_ratio_from_l1 - scalar_ratio_from_l2;

        log::debug!(
            "scalar_ratio_from_l1 is: {:#?}, actual_change is: {:#?}%, expected_change is: {:#?}%",
            scalar_ratio_from_l1,
            scalar_diff,
            self.gas_threshold
        );

        if scalar_diff.abs() * 100.0 > self.gas_threshold as f64 {
            // Set scalar for l2.
            let scalar_expect = (DEFAULT_SCALAR * scalar_ratio_from_l1).ceil() as u128;
            let tx = self.l2_oracle.set_scalar(U256::from(scalar_expect)).legacy();
            let rt = tx.send().await;
            let pending_tx = match rt {
                Ok(pending) => {
                    log::info!("tx of set_scalar has been sent: {:#?}", pending.tx_hash());
                    pending
                }
                Err(e) => {
                    return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                        "set scalar error: {:#?}",
                        e
                    ))));
                }
            };
            pending_tx.await.map_err(|e| {
                OracleError::L1BaseFeeError(anyhow!(format!(
                    "set_scalar check_receipt error: {:#?}",
                    e
                )))
            })?;
        }
        Ok(())
    }
}

async fn query_l1_base_fee(
    l1_provider: &Provider<Http>,
) -> Result<(Option<U256>, Option<U256>), OracleError> {
    let l1_base_fee = match l1_provider.get_block(BlockNumber::Latest).await {
        Ok(ob) => {
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(1);
            if let Some(b) = ob {
                b.base_fee_per_gas
            } else {
                return Err(OracleError::L1BaseFeeError(anyhow!("Block missing base fee per gas")))
            }
        }
        Err(e) => {
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(2);
            return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                "Failed to query l1_base_fee_per_gas: {:#?}",
                e
            ))))
        }
    };
    let gas_price = match l1_provider.get_gas_price().await {
        Ok(gp) => Some(gp),
        Err(e) => {
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(2);
            return Err(OracleError::L1BaseFeeError(anyhow!(format!(
                "Failed to query l1_gas_price: {:#?}",
                e
            ))))
        }
    };

    Ok((l1_base_fee, gas_price))
}
