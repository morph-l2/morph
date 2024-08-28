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

    /// @notice stakers removed
    /// @param stakers     stakers removed
    event StakersRemoved(address[] stakers);

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

    /// @notice slash remaining claimed
    /// @param receiver  receiver address
    /// @param amount    claimed amount
    event SlashRemainingClaimed(address receiver, uint256 amount);

    /// @notice whitelist updated
    /// @param add     addresses added
    /// @param remove  addresses removed
    event WhitelistUpdated(address[] add, address[] remove);

    /// @notice gas limit add staker updated
    /// @param oldGasLimit    old gas limit
    /// @param newGasLimit    new gas limit
    event GasLimitAddStakerUpdated(uint256 oldGasLimit, uint256 newGasLimit);

    /// @notice gas limit remove stakers updated
    /// @param oldGasLimit    old gas limit
    /// @param newGasLimit    new gas limit
    event GasLimitRemoveStakersUpdated(uint256 oldGasLimit, uint256 newGasLimit);

    /// @notice reward percentage updated
    /// @param oldPercentage    old percentage
    /// @param newPercentage    new percentage
    event RewardPercentageUpdated(uint256 oldPercentage, uint256 newPercentage);

    /// @notice challenge deposit value updated
    /// @param oldChallengeDeposit    old challengeDeposit
    /// @param newChallengeDeposit    new challengeDeposit
    event ChallengeDepositUpdated(uint256 oldChallengeDeposit, uint256 newChallengeDeposit);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return all stakers
    function getStakers() external view returns (address[255] memory);

    /// @notice return active stakers
    function getActiveStakers() external view returns (address[] memory);

    /// @notice return staking value
    function stakingValue() external view returns (uint256);

    /// @notice return challenge deposit value
    function challengeDeposit() external view returns (uint256);

    /// @notice whether address is staker
    /// @param addr  the address to check
    function isStaker(address addr) external view returns (bool);

    /// @notice whether address is active staker
    /// @param addr  the address to check
    function isActiveStaker(address addr) external view returns (bool);

    /// @notice get staker bitmap
    /// @param staker  the staker address
    function getStakerBitmap(address staker) external view returns (uint256);

    /// @notice get stakers bitmap
    /// @param stakers  the staker address array
    function getStakersBitmap(address[] calldata stakers) external view returns (uint256);

    /// @notice verify BLS signature
    /// @param signedSequencersBitmap bitmap of signed sequencers
    /// @param sequencerSet           sequencer set
    /// @param msgHash                bls message hash
    /// @param signature              batch signature
    function verifySignature(
        uint256 signedSequencersBitmap,
        address[] calldata sequencerSet,
        bytes32 msgHash,
        bytes calldata signature
    ) external view returns (bool);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice challenger win, slash sequencers
    /// @param sequencersBitmap  the sequencers to slash
    function slash(uint256 sequencersBitmap) external returns (uint256);
}
