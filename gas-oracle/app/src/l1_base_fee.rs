use std::time::Duration;

use crate::abi::gas_price_oracle_abi::GasPriceOracle;
use crate::metrics::ORACLE_SERVICE_METRICS;
use ethers::prelude::*;

static DEFAULT_SCALAR: f64 = 1000000000.0;

/// Update baseFee and scalar.
/// Set the gas data of L1 network to the GasPriceOrale contract on L2.
pub async fn update(
    l1_provider: Provider<Http>,
    l2_provider: Provider<Http>,
    l2_wallet: Address,
    l2_oracle: GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
    gas_threshold: u128,
) {
    // Step1. get l1 data.
    let (l1_base_fee, l1_gas_price) = match query_l1_base_fee(&l1_provider).await {
        (Some(base), Some(price)) => (base, price),
        _ => return,
    };
    if l1_base_fee.is_zero() || l1_gas_price.is_zero() {
        log::error!("current ethereum baseFee or gas_price is zero");
        return;
    }
    #[rustfmt::skip]
    log::info!("current ethereum baseFee is: {:#?}, gas_price is: {:#?}", l1_base_fee, l1_gas_price);
    ORACLE_SERVICE_METRICS.l1_base_fee.set(
        ethers::utils::format_units(l1_base_fee, "gwei")
            .unwrap_or(String::from("0"))
            .parse()
            .unwrap_or(0.0),
    );

    // Step2. get l2 data.
    let base_fee_on_l2: U256 = match l2_oracle.l_1_base_fee().await {
        Ok(fee) => fee,
        Err(e) => {
            log::error!("Failed to query query_l1_base_fee on l2: {:#?}", e);
            return;
        }
    };
    let scalar: U256 = match l2_oracle.scalar().await {
        Ok(s) => s,
        Err(e) => {
            log::error!("Failed to query scalar on l2: {:#?}", e);
            return;
        }
    };
    #[rustfmt::skip]
    log::info!("current l1BaseFee on l2 is: {:#?}, scalar is: {:#?}", base_fee_on_l2, scalar);
    ORACLE_SERVICE_METRICS.l1_base_fee_on_l2.set(
        ethers::utils::format_units(base_fee_on_l2, "gwei")
            .unwrap_or(String::from("0"))
            .parse()
            .unwrap_or(0.0),
    );

    // Step3. Update contract.
    update_base_fee(l1_base_fee, base_fee_on_l2, gas_threshold, &l2_oracle).await;

    // Step4. Record wallet balance.
    let balance = match l2_provider.get_balance(l2_wallet, None).await {
        Ok(b) => b,
        Err(e) => {
            log::error!("l2_wallet.get_balance error: {:#?}", e);
            return;
        }
    };
    ORACLE_SERVICE_METRICS
        .gas_oracle_owner_balance
        .set(ethers::utils::format_ether(balance).parse().unwrap_or(0.0));
}

async fn update_base_fee(
    l1_base_fee: U256,
    base_fee_on_l2: U256,
    gas_threshold: u128,
    l2_oracle: &GasPriceOracle<SignerMiddleware<Provider<Http>, LocalWallet>>,
) {
    let actual_change = l1_base_fee.as_u128().abs_diff(base_fee_on_l2.as_u128());
    let expected_change = base_fee_on_l2.as_u128() * gas_threshold / 100;
    log::info!(
        "l1BaseFee actual_change: {:#?}, expected_change: {:#?}",
        actual_change,
        expected_change
    );

    //min gas price
    if l1_base_fee > U256::from(0) && actual_change > expected_change {
        // Set l1_base_fee for l2.
        let tx = l2_oracle.set_l1_base_fee(l1_base_fee).legacy();
        let rt = tx.send().await;
        match rt {
            Ok(info) => {
                log::info!("tx of set_l1_base_fee has been sent: {:#?}", info.tx_hash());
            }
            Err(e) => log::error!("set_l1_base_fee error: {:#?}", e),
        }
    }
}

async fn query_l1_base_fee(l1_provider: &Provider<Http>) -> (Option<U256>, Option<U256>) {
    let l1_base_fee = match l1_provider.get_block(BlockNumber::Latest).await {
        Ok(ob) => {
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(1);
            if let Some(b) = ob {
                b.base_fee_per_gas
            } else {
                log::error!("Block missing base fee per gas");
                None
            }
        }
        Err(e) => {
            log::error!("Failed to query l1_base_fee_per_gas: {:#?}", e);
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(2);
            None
        }
    };
    let gas_price = match l1_provider.get_gas_price().await {
        Ok(gp) => Some(gp),
        Err(e) => {
            log::error!("Failed to query l1_gas_price: {:#?}", e);
            ORACLE_SERVICE_METRICS.l1_rpc_status.set(2);
            None
        }
    };

    (l1_base_fee, gas_price)
}
