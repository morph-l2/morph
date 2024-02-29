// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import {IL1MessageQueue} from "./IL1MessageQueue.sol";

interface IL1MessageQueueWithGasPriceOracle {
    /**********
     * Events *
     **********/

    /// @notice Emitted when current l2 base fee is updated.
    /// @param oldL2BaseFee The original l2 base fee before update.
    /// @param newL2BaseFee The current l2 base fee updated.
    event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee);


    /*************************
     * Public View Functions *
     *************************/

    /// @notice Return the latest known l2 base fee.
    function l2BaseFee() external view returns (uint256);
}
