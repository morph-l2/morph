// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ICrossDomainMessenger} from "../ICrossDomainMessenger.sol";

abstract contract Sequencer {
    /**
     * @notice Messenger contract on this domain.
     */
    ICrossDomainMessenger public immutable MESSENGER;

    /**
     * @notice Corresponding sequencer on the other domain.
     */
    Sequencer public immutable OTHER_SEQUENCER;

    /**
     * @notice Ensures that the caller is a cross-chain message from the other sequencer.
     */
    modifier onlyOtherSequencer() {
        require(
            msg.sender == address(MESSENGER) &&
                MESSENGER.xDomainMessageSender() == address(OTHER_SEQUENCER),
            "Sequencer: function can only be called from the other sequencer"
        );
        _;
    }

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     * @param _otherSequencer Address of the other Sequencer contract.
     */
    constructor(address payable _messenger, address payable _otherSequencer) {
        MESSENGER = ICrossDomainMessenger(_messenger);
        OTHER_SEQUENCER = Sequencer(_otherSequencer);
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
