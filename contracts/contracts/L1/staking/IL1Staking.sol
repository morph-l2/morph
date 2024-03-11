// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IL1Staking {
    /**
     * @notice staking value
     */
    function STAKING_VALUE() external returns (uint256);

    /**
     * @notice whether address is staker
     */
    function isStaker(address addr) external returns (bool);

    /**
     * @notice verify BLS signature
     * @param signedSequencers  signed sequencers
     * @param sequencerSet      sequencer set
     * @param signature         batch signature
     */
    function verifySignature(
        address[] memory signedSequencers,
        address[] memory sequencerSet,
        bytes memory signature
    ) external returns (bool);

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(address[] memory sequencers, address challenger) external;
}
