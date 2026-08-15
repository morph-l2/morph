/// Predeployed Gas Price Oracle
pub mod l1_gas_price_oracle {
    use alloy_primitives::{address, Address, U256};

    /// L1GasPriceOracle predeployed address
    pub const ADDRESS: Address = address!("5300000000000000000000000000000000000002");
    /// L1 base fee slot in L1GasPriceOracle
    pub const BASE_FEE_SLOT: U256 = U256::from_limbs([1, 0, 0, 0]);

    /// The following 2 slots will be depreciated after curie fork
    /// L1 overhead slot in L1GasPriceOracle
    pub const OVERHEAD_SLOT: U256 = U256::from_limbs([2, 0, 0, 0]);
    /// L1 scalar slot in L1GasPriceOracle
    pub const SCALAR_SLOT: U256 = U256::from_limbs([3, 0, 0, 0]);

    /// THe following 3 slots plus `BASE_FEE_SLOT` will be used for l1 fee after curie fork
    /// L1 BlobBaseFee slot in L1GasPriceOracle after Curie fork
    pub const L1_BLOB_BASEFEE_SLOT: U256 = U256::from_limbs([5, 0, 0, 0]);
    /// L1 commitScalar slot in L1GasPriceOracle after Curie fork
    pub const COMMIT_SCALAR_SLOT: U256 = U256::from_limbs([6, 0, 0, 0]);
    /// L1 blob_scalar slot in L1GasPriceOracle after Curie fork
    pub const BLOB_SCALAR_SLOT: U256 = U256::from_limbs([7, 0, 0, 0]);
    /// L1 isCurie slot in L1GasPriceOracle after Curie fork
    pub const IS_CURIE_SLOT: U256 = U256::from_limbs([8, 0, 0, 0]);
    /// Initial commit scalar after curie fork
    pub const INITIAL_COMMIT_SCALAR: U256 = U256::from_limbs([230759955285, 0, 0, 0]);
    /// Initial blob scalar after curie fork
    pub const INITIAL_BLOB_SCALAR: U256 = U256::from_limbs([417565260, 0, 0, 0]);
}

/// Predeployed L2ToL1Message
pub mod l2_to_l1_message {
    use alloy_primitives::{address, uint, Address, U256};

    /// Withdraw root address
    pub const WITHDRAW_ROOT_ADDRESS: Address =
        address!("0x5300000000000000000000000000000000000001");
    /// Withdraw root slot
    pub const WITHDRAW_ROOT_SLOT: U256 = uint!(33_U256);
}
