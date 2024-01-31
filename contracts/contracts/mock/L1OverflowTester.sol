// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {ICrossDomainMessenger} from "./../libraries/ICrossDomainMessenger.sol";
import {L2OverflowTester} from "./L2OverflowTester.sol";

contract L1OverflowTester {
    /**
     * @notice Messenger contract on this domain.
     */
    ICrossDomainMessenger public immutable MESSENGER;

    address public immutable OTHERTESTER;

    uint32 internal gas_limit = 2_000_000;

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     * @param _otherTester Address of L2OverflowTester
     */
    constructor(
        address payable _messenger,
        address _otherTester,
        uint32 _gasLimit
    ) {
        MESSENGER = ICrossDomainMessenger(_messenger);
        OTHERTESTER = _otherTester;
        gas_limit = _gasLimit;
    }

    function updateGasLimit(uint32 _gasLimit) public {
        gas_limit = _gasLimit;
    }

    function crossHash(string calldata _message, uint count) public {
        MESSENGER.sendMessage(
            OTHERTESTER,
            0,
            abi.encodeWithSelector(
                L2OverflowTester.hash.selector,
                _message,
                count
            ),
            gas_limit
        );
    }
}
