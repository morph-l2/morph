// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CrossDomainMessenger} from "./CrossDomainMessenger.sol";

abstract contract RollupMessage {
    /**
     * @notice Messenger contract on this domain.
     */
    CrossDomainMessenger public immutable MESSENGER;

    /**
     * @notice Corresponding counterpart on the other domain.
     */
    address public immutable COUNTERPART;

    /**
     * @notice Ensures that the caller is a cross-chain message from the L1 rollup.
     */
    modifier onlyCounterpart() {
        require(
            msg.sender == address(MESSENGER) &&
                MESSENGER.xDomainMessageSender() == COUNTERPART,
            "Rollup: function can only be called from the L1 rollup"
        );
        _;
    }

    /**
     * @param _messenger    Address of CrossDomainMessenger on this network.
     * @param _counterpart  Address of the counterpart contract.
     */
    constructor(address payable _messenger, address payable _counterpart) {
        MESSENGER = CrossDomainMessenger(_messenger);
        COUNTERPART = _counterpart;
    }

    /**
     * @notice Getter for messenger contract.
     *
     * @return Messenger contract on this domain.
     */
    function messenger() external view returns (CrossDomainMessenger) {
        return MESSENGER;
    }
}
