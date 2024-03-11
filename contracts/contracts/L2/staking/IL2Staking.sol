// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice update stakers
     */
    function updateStakers(
        Types.StakerInfo[] memory add,
        Types.StakerInfo[] memory remove
    ) external;
}
