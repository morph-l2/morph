
/// Predeployed L2ToL1Message
pub mod l2_to_l1_message {
    use alloy_primitives::{address, uint, Address, U256};

    /// Withdraw root address
    pub const WITHDRAW_ROOT_ADDRESS: Address =
        address!("0x5300000000000000000000000000000000000001");
    /// Withdraw root slot
    pub const WITHDRAW_ROOT_SLOT: U256 = uint!(33_U256);
}
