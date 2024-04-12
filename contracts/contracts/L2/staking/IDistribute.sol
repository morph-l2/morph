// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    // event of claimAll
    event ClaimAll(address from, address to, uint256 amount);

    // event of claim
    event Claim(address from, address to, uint256 amount);

    /**
     * @dev Initialization parameter, which can only be called once.
     * @param morphToken_ address
     * @param record_ address
     * @param stake_ address
     */
    function initialize(
        address morphToken_,
        address record_,
        address stake_
    ) external;

    function notify(uint256 blockTime, uint256 blockNumber) external;

    function notifyUnDelegate(
        address sequencer,
        address account,
        uint256 deadlineClaimEpochIndex
    ) external;

    function notifyDelegate(
        uint256 epochIndex,
        address sequencer,
        address account,
        uint256 amount,
        uint256 blockNumber
    ) external;

    function mint() external;

    function claimAll(address account) external;

    function claim(address sequencer, address account) external;
}
