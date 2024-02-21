// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL1MessageQueue} from "./IL1MessageQueue.sol";
import {IL1MessageQueueWithGasPriceOracle} from "./IL1MessageQueueWithGasPriceOracle.sol";

import {L1MessageQueue} from "./L1MessageQueue.sol";

contract L1MessageQueueWithGasPriceOracle is
    L1MessageQueue,
    IL1MessageQueueWithGasPriceOracle
{
    /*************
     * Constants *
     *************/

    /// @notice The intrinsic gas for transaction.
    uint256 private constant INTRINSIC_GAS_TX = 21000;

    /// @notice The appropriate intrinsic gas for each byte.
    uint256 private constant APPROPRIATE_INTRINSIC_GAS_PER_BYTE = 16;

    /*************
     * Variables *
     *************/

    /// @inheritdoc IL1MessageQueueWithGasPriceOracle
    uint256 public override l2BaseFee;

    /***************
     * Constructor *
     ***************/

    /// @notice Constructor for `L1MessageQueueWithGasPriceOracle` implementation contract.
    ///
    /// @param _messenger The address of `L1CrossDomainMessenger` contract.
    /// @param _rollup The address of `Rollup` contract.
    /// @param _enforcedTxGateway The address of `EnforcedTxGateway` contract.
    constructor(
        address _messenger,
        address _rollup,
        address _enforcedTxGateway
    ) L1MessageQueue(_messenger, _rollup, _enforcedTxGateway) {}

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IL1MessageQueue
    function estimateCrossDomainMessageFee(
        uint256 _gasLimit
    )
        external
        view
        override(IL1MessageQueue, L1MessageQueue)
        returns (uint256)
    {
        return _gasLimit * l2BaseFee;
    }

    /// @inheritdoc IL1MessageQueue
    function calculateIntrinsicGasFee(
        bytes calldata _calldata
    )
        public
        pure
        virtual
        override(IL1MessageQueue, L1MessageQueue)
        returns (uint256)
    {
        // no way this can overflow `uint256`
        unchecked {
            return
                INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE;
        }
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Allows whitelistCheckered caller to modify the l2 base fee.
    /// @param _newL2BaseFee The new l2 base fee.
    function setL2BaseFee(uint256 _newL2BaseFee) public onlyOwner {
        uint256 _oldL2BaseFee = l2BaseFee;
        l2BaseFee = _newL2BaseFee;

        emit UpdateL2BaseFee(_oldL2BaseFee, _newL2BaseFee);
    }
}
