// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {Rollup} from "../L1/rollup/Rollup.sol";

contract MockRollup is Rollup {
    constructor() Rollup(0, payable(address(1)), address(1), address(1)) {}

    function setLastFinalizedBatchIndex(
        uint256 _lastFinalizedBatchIndex
    ) external {
        lastFinalizedBatchIndex = _lastFinalizedBatchIndex;
    }
}
