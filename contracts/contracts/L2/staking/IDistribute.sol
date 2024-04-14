// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    // event of claimAll
    event ClaimAll(address indexed from, address indexed to, uint256 amount);

    // event of claim
    event Claim(address indexed from, address indexed to, uint256 amount);

    event NotifyDelegate(
        address indexed sequencer,
        uint256 indexed epochIndex,
        address indexed account,
        uint256 amount,
        uint256 blockNumber);

    event NotifyUnDelegate(
        address indexed sequencer,
        address indexed account,
        uint256 deadlineClaimEpochIndex
    );

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
        address sequencer,
        uint256 epochIndex,
        address account,
        uint256 amount,
        uint256 blockNumber
    ) external;

    function mint() external;

    function claimAll(address account) external;

    function claim(address sequencer, address account) external;
}
