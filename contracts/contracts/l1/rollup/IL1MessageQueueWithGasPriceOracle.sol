// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface IL1MessageQueueWithGasPriceOracle {
    /**********
     * Events *
     **********/

    /// @notice Emitted when current l2 base fee is updated.
    /// @param oldL2BaseFee The original l2 base fee before update.
    /// @param newL2BaseFee The current l2 base fee updated.
    event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee);

    /// @notice Emitted when owner updates whitelist checker contract.
    /// @param _oldWhitelistChecker The address of old whitelist checker contract.
    /// @param _newWhitelistChecker The address of new whitelist checker contract.
    event UpdateWhitelistChecker(address indexed _oldWhitelistChecker, address indexed _newWhitelistChecker);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice Return the latest known l2 base fee.
    function l2BaseFee() external view returns (uint256);

    /// @notice Return the address of whitelist checker contract.
    function whitelistChecker() external view returns (address);
}
