// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {ICrossDomainMessenger} from "./../libraries/ICrossDomainMessenger.sol";
import {L2OverflowTester} from "./L2OverflowTester.sol";

contract L1OverflowTester {
    /**
     * @notice Messenger contract on this domain.
     */
    ICrossDomainMessenger public immutable MESSENGER;

    address public immutable OTHERTESTER;

    uint32 internal gasLimit = 2_000_000;

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     * @param _otherTester Address of L2OverflowTester
     */
    constructor(address payable _messenger, address _otherTester, uint32 _gasLimit) {
        MESSENGER = ICrossDomainMessenger(_messenger);
        OTHERTESTER = _otherTester;
        gasLimit = _gasLimit;
    }

    function updateGasLimit(uint32 _gasLimit) public {
        gasLimit = _gasLimit;
    }

    function crossHash(string calldata _message, uint256 count) public {
        MESSENGER.sendMessage(OTHERTESTER, 0, abi.encodeCall(L2OverflowTester.hash, (_message, count)), gasLimit);
    }
}
