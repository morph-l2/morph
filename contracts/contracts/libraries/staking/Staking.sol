// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ICrossDomainMessenger} from "../ICrossDomainMessenger.sol";

abstract contract Staking {
    /**
     * @notice Messenger contract on this domain.
     */
    ICrossDomainMessenger public immutable MESSENGER;

    /**
     * @notice Corresponding staking on the other domain.
     */
    Staking public immutable OTHER_STAKING;

    /**
     * @notice Ensures that the caller is a cross-chain message from the other staking.
     */
    modifier onlyOtherStaking() {
        require(
            msg.sender == address(MESSENGER) && MESSENGER.xDomainMessageSender() == address(OTHER_STAKING),
            "staking: only other staking contract allowed"
        );
        _;
    }

    /**
     * @param _messenger    Address of CrossDomainMessenger on this network.
     * @param _otherStaking Address of the other Staking contract.
     */
    constructor(address payable _messenger, address payable _otherStaking) {
        MESSENGER = ICrossDomainMessenger(_messenger);
        OTHER_STAKING = Staking(_otherStaking);
    }

    /**
     * @notice Getter for messenger contract.
     *
     * @return Messenger contract on this domain.
     */
    function messenger() external view returns (address) {
        return address(MESSENGER);
    }
}
