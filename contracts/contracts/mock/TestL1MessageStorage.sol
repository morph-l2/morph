// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {BitMapsUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/BitMapsUpgradeable.sol";
import {L1MessageStorage} from "../L1/L1MessageStorage.sol";

contract TestL1MessageStorage is L1MessageStorage {
    function queueTransaction(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes calldata _data
    ) external {
        _queueTransaction(_sender, _target, _value, _gasLimit, _data);
    }

    function popCrossDomainMessage(
        uint256 _startIndex,
        uint256 _count,
        uint256 _skippedBitmap
    ) external {
        _popCrossDomainMessage(_startIndex, _count, _skippedBitmap);
    }
}
