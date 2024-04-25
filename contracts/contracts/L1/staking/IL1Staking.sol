// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IL1Staking {
    /**********
     * Events *
     **********/

    /// @notice staker registered
    /// @param addr     staker address
    /// @param tmKey    tendermint pubkey
    /// @param blsKey   BLS pubkey
    event Registered(address addr, bytes32 tmKey, bytes blsKey);

    /// @notice Withdrawn
    /// @param addr             staker address
    /// @param unlockHeight     unlock block height
    event Withdrawn(address indexed addr, uint256 unlockHeight);

    /// @notice staker claimed
    /// @param staker       staker claimed
    /// @param receiver     receiver address
    event Claimed(address indexed staker, address receiver);

    /// @notice stakers were slashed
    /// @param stakers  slashed stakers
    event Slashed(address[] stakers);

    /// @notice params updated
    /// @param gasLimit     new gas limit
    event ParamsUpdated(uint256 gasLimit);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice staking value
    function getStakers() external view returns (address[] memory);

    /// @notice staking value
    function stakingValue() external view returns (uint256);

    /// @notice whether address is staker
    /// @param addr  the address to check
    function isStaker(address addr) external view returns (bool);

    /// @notice verify BLS signature
    /// @param signedSequencers  signed sequencers
    /// @param sequencerSet      sequencer set
    /// @param msgHash           bls message hash
    /// @param signature         batch signature
    function verifySignature(
        address[] memory signedSequencers,
        address[] memory sequencerSet,
        bytes32 msgHash,
        bytes memory signature
    ) external view returns (bool);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice challenger win, slash sequencers
    /// @param sequencers  the sequencers to slash
    function slash(address[] memory sequencers) external returns (uint256);
}
