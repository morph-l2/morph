// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IL1Staking {
    /**
     * @notice staking value
     */
    function STAKING_VALUE() external returns (uint256);

    /**
     * @notice whether address is staker
     * @param addr  the address to check
     */
    function isStaker(address addr) external returns (bool);

    /**
     * @notice verify BLS signature
     * @param signedSequencers  signed sequencers
     * @param sequencerSet      sequencer set
     * @param msgHash           bls message hash
     * @param signature         batch signature
     */
    function verifySignature(
        address[] memory signedSequencers,
        address[] memory sequencerSet,
        bytes32 msgHash,
        bytes memory signature
    ) external returns (bool);

    /**
     * @notice challenger win, slash sequencers
     * @param sequencers  the sequencers to slash
     */
    function slash(address[] memory sequencers) external returns (uint256);
}