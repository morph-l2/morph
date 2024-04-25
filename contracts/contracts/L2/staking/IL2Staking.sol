// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /***********
     * Structs *
     ***********/

    /// @notice Undelegation representing a undelegation info.
    ///
    /// @custom:field delegatee  delegatee
    /// @custom:field amount     staking amount
    /// @custom:field unlock     unlock epoch index
    struct Undelegation {
        address delegatee;
        uint256 amount;
        uint256 unlockEpoch;
    }

    /**********
     * Events *
     **********/

    /// @notice Emitted delegated stake
    /// @param delegatee          delegatee
    /// @param delegator          unlock epoch index
    /// @param amount             new delegation amount, not increment
    /// @param effectiveEpoch     effective epoch
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch
    );

    /// @notice Emitted undelegated stake
    /// @param delegatee          delegatee
    /// @param delegator          unlock epoch index
    /// @param amount             delegation amount
    /// @param effectiveEpoch     effective epoch
    /// @param ublockEpoch        effective epoch
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch,
        uint256 ublockEpoch
    );

    /// @notice Emitted claim info
    /// @param delegator    delegator
    /// @param amount       amount
    event UndelegationClaimed(address indexed delegator, uint256 amount);

    /// @notice Emitted commission updated
    /// @param staker           staker address
    /// @param percentage       commission percentage
    /// @param epochEffective   epoch effective
    event CommissionUpdated(
        address indexed staker,
        uint256 percentage,
        uint256 epochEffective
    );

    /// @notice Emitted staker added
    /// @param addr     staker address
    /// @param tmKey    staker tendermint pubkey
    /// @param blsKey   staker BLS pubkey
    event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey);

    /// @notice Emitted stakers removed
    /// @param stakerAddresses  stakers removed
    event StakerRemoved(address[] stakerAddresses);

    /// @notice Emitted reward start time updated
    /// @param oldTime    The old reward start time
    /// @param newTime    The new reward start time
    event RewardStartTimeUpdated(uint256 oldTime, uint256 newTime);

    /// @notice Emitted sequencer set max size updated
    /// @param oldSize    The old sequencer set max size
    /// @param newSize    The new sequencer set max size
    event SequencerSetMaxSizeUpdated(uint256 oldSize, uint256 newSize);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice reward epoch
    function rewardStartTime() external view returns (uint256);

    /// @notice max size of sequencer set
    function sequencerSetMaxSize() external view returns (uint256);

    /// @notice undelegate lock epochs
    function undelegateLockEpochs() external view returns (uint256);

    /// @notice start reward
    function startReward() external;

    /// @notice return current reward epoch index
    function currentEpoch() external view returns (uint256);

    /// @notice check if the user has staked to staker
    /// @param staker sequencers size
    function isStakingTo(address staker) external view returns (bool);

    /// @notice Get all the delegators which staked to staker
    /// @param staker sequencers size
    function getDelegators(
        address staker
    ) external view returns (address[] memory);

    /// @notice get stakers info
    function getStakesInfo(
        address[] calldata _stakerAddresses
    ) external view returns (Types.StakerInfo[] memory);

    /// @notice get stakers
    function getStakers() external view returns (Types.StakerInfo[] memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice add staker, sync from L1
    /// @param add       staker to add. {addr, tmKey, blsKey}
    function addStaker(Types.StakerInfo calldata add) external;

    /// @notice remove stakers, sync from L1
    /// @param remove    staker to remove
    function removeStakers(address[] calldata remove) external;

    /// @notice setCommissionRate set delegate commission percentage
    /// @param commission    commission percentage
    function setCommissionRate(uint256 commission) external;

    /// @notice delegator stake morph to staker
    /// @param staker    stake to whom
    /// @param amount    stake amount
    function delegateStake(address staker, uint256 amount) external;

    /// @notice delegator unstake morph
    /// @param delegatee delegatee address
    function undelegateStake(address delegatee) external;

    /// @notice delegator cliam delegate staking value
    function claimUndelegation() external;

    /// @notice delegator claim reward
    /// @param delegatee         delegatee address, claim all if empty
    /// @param targetEpochIndex  up to the epoch index that the delegator wants to claim
    function claimReward(address delegatee, uint256 targetEpochIndex) external;

    /// @notice claimCommission claim commission reward
    /// @param targetEpochIndex   up to the epoch index that the staker wants to claim
    function claimCommission(uint256 targetEpochIndex) external;
}
