// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {IDistribute} from "./IDistribute.sol";
import {IMorphToken} from "../system/IMorphToken.sol";

contract L2Staking is IL2Staking, Staking, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    /*************
     * Constants *
     *************/

    /// @notice reward epoch, seconds of one day (3600 * 24)
    uint256 private constant REWARD_EPOCH = 86400;

    /// @notice MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;

    /// @notice sequencer contract address
    address public immutable SEQUENCER_CONTRACT;

    /// @notice distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;

    /*************
     * Variables *
     *************/

    /// @notice is reward started
    bool public rewardStarted;

    /// @notice reward start time
    uint256 public rewardStartTime;

    /// @notice max number of sequencer set
    uint256 public sequencerSetMaxSize;

    /// @notice undelegate lock epochs
    uint256 public undelegateLockEpochs;

    /// @notice latest sequencer set size
    uint256 public latestSequencerSetSize;

    /// @notice sequencer candidate number
    uint256 public candidateNumber;

    /// @notice sync from l1 staking
    address[] public stakerAddresses;

    /// @notice staker rankings
    mapping(address staker => uint256 ranking) public stakerRankings;

    /// @notice stakers info
    mapping(address staker => Types.StakerInfo) public stakers;

    /// @notice staker commissions, default commission is zero if not set
    mapping(address staker => uint256 amount) public commissions;

    /// @notice staker's total delegation amount
    mapping(address staker => uint256 totalDelegationAmount) public stakerDelegations;

    /// @notice delegators of staker
    mapping(address staker => EnumerableSetUpgradeable.AddressSet) internal delegators;

    /// @notice delegations of a staker
    mapping(address staker => mapping(address delegator => uint256 amount)) public delegations;

    /// @notice delegator's undelegations
    mapping(address delegator => Undelegation[]) public undelegations;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice must be staker
    modifier isStaker(address addr) {
        require(stakerRankings[addr] > 0, "not staker");
        _;
    }

    /// @notice only staker allowed
    modifier onlyStaker() {
        require(stakerRankings[_msgSender()] > 0, "only staker allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    /// @param _otherStaking Address of the staking contract on the other network.
    constructor(address payable _otherStaking) Staking(payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER), _otherStaking) {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
    }

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    /// @param _owner                owner
    /// @param _sequencersMaxSize    max size of sequencer set
    /// @param _undelegateLockEpochs undelegate lock epochs
    /// @param _rewardStartTime      reward start time
    /// @param _stakers              initial stakers, must be same as initial sequencer set in sequencer contract
    function initialize(
        address _owner,
        uint256 _sequencersMaxSize,
        uint256 _undelegateLockEpochs,
        uint256 _rewardStartTime,
        Types.StakerInfo[] calldata _stakers
    ) public initializer {
        require(_owner != address(0), "invalid owner address");
        require(_sequencersMaxSize > 0, "sequencersSize must greater than 0");
        require(_undelegateLockEpochs > 0, "invalid undelegateLockEpochs");
        require(
            _rewardStartTime > block.timestamp && _rewardStartTime % REWARD_EPOCH == 0,
            "invalid reward start time"
        );
        require(_stakers.length > 0, "invalid initial stakers");

        _transferOwnership(_owner);
        __ReentrancyGuard_init();

        sequencerSetMaxSize = _sequencersMaxSize;
        undelegateLockEpochs = _undelegateLockEpochs;
        rewardStartTime = _rewardStartTime;
        latestSequencerSetSize = _stakers.length;
        for (uint256 i = 0; i < latestSequencerSetSize; i++) {
            stakers[_stakers[i].addr] = _stakers[i];
            stakerAddresses.push(_stakers[i].addr);
            stakerRankings[_stakers[i].addr] = i + 1;
        }

        emit SequencerSetMaxSizeUpdated(0, _sequencersMaxSize);
        emit RewardStartTimeUpdated(0, _rewardStartTime);
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice add staker, sync from L1
    /// @param add   staker to add. {addr, tmKey, blsKey}
    function addStaker(Types.StakerInfo calldata add) external onlyOtherStaking {
        if (stakerRankings[add.addr] == 0) {
            stakerAddresses.push(add.addr);
            stakerRankings[add.addr] = stakerAddresses.length;
        }
        stakers[add.addr] = add;
        emit StakerAdded(add.addr, add.tmKey, add.blsKey);

        if (!rewardStarted && stakerAddresses.length <= sequencerSetMaxSize) {
            _updateSequencerSet();
        }
    }

    /// @notice remove stakers, sync from L1. If new sequencer set is nil, layer2 will stop producing blocks
    /// @param remove    staker to remove
    function removeStakers(address[] calldata remove) external onlyOtherStaking {
        bool updateSequencerSet = false;
        for (uint256 i = 0; i < remove.length; i++) {
            if (stakerRankings[remove[i]] <= latestSequencerSetSize) {
                updateSequencerSet = true;
            }

            if (stakerRankings[remove[i]] > 0) {
                // update stakerRankings
                for (uint256 j = stakerRankings[remove[i]] - 1; j < stakerAddresses.length - 1; j++) {
                    stakerAddresses[j] = stakerAddresses[j + 1];
                    stakerRankings[stakerAddresses[j]] -= 1;
                }
                stakerAddresses.pop();
                delete stakerRankings[remove[i]];

                // update candidateNumber
                if (stakerDelegations[remove[i]] > 0) {
                    candidateNumber -= 1;
                }
            }

            delete stakers[remove[i]];
        }
        emit StakerRemoved(remove);

        if (updateSequencerSet) {
            _updateSequencerSet();
        }
    }

    /// @notice setCommissionRate set delegate commission percentage
    /// @param commission    commission percentage
    function setCommissionRate(uint256 commission) external onlyStaker {
        require(commission <= 20, "invalid commission");
        commissions[_msgSender()] = commission;
        uint256 epochEffective = rewardStarted ? currentEpoch() + 1 : 0;
        emit CommissionUpdated(_msgSender(), commission, epochEffective);
    }

    /// @notice claimCommission claim commission reward
    /// @param targetEpochIndex   up to the epoch index that the staker wants to claim
    function claimCommission(uint256 targetEpochIndex) external onlyStaker nonReentrant {
        IDistribute(DISTRIBUTE_CONTRACT).claimCommission(_msgSender(), targetEpochIndex);
    }

    /// @notice update params
    /// @param _sequencerSetMaxSize   max size of sequencer set
    function updateSequencerSetMaxSize(uint256 _sequencerSetMaxSize) external onlyOwner {
        require(
            _sequencerSetMaxSize > 0 && _sequencerSetMaxSize != sequencerSetMaxSize,
            "invalid new sequencer set max size"
        );
        uint256 _oldSequencerSetMaxSize = sequencerSetMaxSize;
        sequencerSetMaxSize = _sequencerSetMaxSize;
        emit SequencerSetMaxSizeUpdated(_oldSequencerSetMaxSize, _sequencerSetMaxSize);

        if (sequencerSetMaxSize < latestSequencerSetSize) {
            // update sequencer set
            _updateSequencerSet();
        }
    }

    /// @notice advance layer2 stage
    /// @param _rewardStartTime   reward start time
    function updateRewardStartTime(uint256 _rewardStartTime) external onlyOwner {
        require(!rewardStarted && rewardStartTime > block.timestamp, "reward already started");
        require(
            _rewardStartTime > block.timestamp &&
                _rewardStartTime % REWARD_EPOCH == 0 &&
                _rewardStartTime != rewardStartTime,
            "invalid reward start time"
        );
        uint256 _oldTime = rewardStartTime;
        rewardStartTime = _rewardStartTime;
        emit RewardStartTimeUpdated(_oldTime, _rewardStartTime);
    }

    /// @notice start reward
    function startReward() external onlyOwner {
        require(block.timestamp >= rewardStartTime, "can't start before reward start time");
        require(candidateNumber > 0, "none candidate");

        rewardStarted = true;

        // sort stakers by insertion sort
        for (uint256 i = 1; i < stakerAddresses.length; i++) {
            for (uint256 j = 0; j < i; j++) {
                if (stakerDelegations[stakerAddresses[i]] > stakerDelegations[stakerAddresses[j]]) {
                    address tmp = stakerAddresses[j];
                    stakerAddresses[j] = stakerAddresses[i];
                    stakerAddresses[i] = tmp;
                }
            }
        }
        // update rankings
        for (uint256 i = 0; i < stakerAddresses.length; i++) {
            stakerRankings[stakerAddresses[i]] = i + 1;
        }

        // update sequencer set
        _updateSequencerSet();
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice delegator stake morph to delegatee
    /// @param delegatee    stake to whom
    /// @param amount       stake amount
    function delegateStake(address delegatee, uint256 amount) external isStaker(delegatee) nonReentrant {
        require(amount > 0, "invalid stake amount");
        // Re-staking to the same delegatee is not allowed before claiming undelegation
        require(!_unclaimed(_msgSender(), delegatee), "undelegation unclaimed");

        stakerDelegations[delegatee] += amount;
        delegations[delegatee][_msgSender()] += amount;
        delegators[delegatee].add(_msgSender()); // will not be added repeatedly

        if (stakerDelegations[delegatee] == amount) {
            candidateNumber += 1;
        }

        uint256 beforeRanking = stakerRankings[delegatee];
        if (rewardStarted && beforeRanking > 1) {
            // update stakers and rankings
            for (uint256 i = beforeRanking - 1; i > 0; i--) {
                if (stakerDelegations[stakerAddresses[i]] > stakerDelegations[stakerAddresses[i - 1]]) {
                    address tmp = stakerAddresses[i - 1];
                    stakerAddresses[i - 1] = stakerAddresses[i];
                    stakerAddresses[i] = tmp;

                    stakerRankings[stakerAddresses[i - 1]] = i;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                }
            }
        }
        uint256 effectiveEpoch = rewardStarted ? currentEpoch() + 1 : 0;

        emit Delegated(
            delegatee,
            _msgSender(),
            delegations[delegatee][_msgSender()], // new amount, not incremental
            amount,
            effectiveEpoch
        );

        // notify delegation to distribute contract
        IDistribute(DISTRIBUTE_CONTRACT).notifyDelegation(
            delegatee,
            _msgSender(),
            effectiveEpoch,
            delegations[delegatee][_msgSender()],
            stakerDelegations[delegatee],
            delegators[delegatee].length(),
            delegations[delegatee][_msgSender()] == amount
        );

        // transfer morph token from delegator to this
        _transferFrom(_msgSender(), address(this), amount);

        if (
            rewardStarted && beforeRanking > latestSequencerSetSize && stakerRankings[delegatee] <= sequencerSetMaxSize
        ) {
            _updateSequencerSet();
        }
    }

    /// @notice delegator unstake morph
    /// @param delegatee delegatee address
    function undelegateStake(address delegatee) external nonReentrant {
        // must claim any undelegation first
        require(!_unclaimed(_msgSender(), delegatee), "undelegation unclaimed");
        require(_isStakingTo(delegatee), "staking amount is zero");

        // staker has been removed, unlock next epoch
        bool removed = stakerRankings[delegatee] == 0;

        uint256 effectiveEpoch = rewardStarted ? currentEpoch() + 1 : 0;
        uint256 unlockEpoch = (rewardStarted && !removed)
            ? effectiveEpoch + undelegateLockEpochs // reward started and staker is active
            : effectiveEpoch; // equal to 0 if reward not started. equal to effectiveEpoch is staker is removed

        Undelegation memory undelegation = Undelegation(delegatee, delegations[delegatee][_msgSender()], unlockEpoch);

        undelegations[_msgSender()].push(undelegation);
        delete delegations[delegatee][_msgSender()];
        stakerDelegations[delegatee] -= undelegation.amount;
        delegators[delegatee].remove(_msgSender());

        uint256 beforeRanking = stakerRankings[delegatee];
        if (!removed && rewardStarted && beforeRanking < candidateNumber) {
            // update stakers and rankings
            for (uint256 i = stakerRankings[delegatee] - 1; i < candidateNumber - 1; i++) {
                if (stakerDelegations[stakerAddresses[i + 1]] > stakerDelegations[stakerAddresses[i]]) {
                    address tmp = stakerAddresses[i];
                    stakerAddresses[i] = stakerAddresses[i + 1];
                    stakerAddresses[i + 1] = tmp;

                    stakerRankings[stakerAddresses[i]] = i + 1;
                    stakerRankings[stakerAddresses[i + 1]] = i + 2;
                }
            }
        }

        // update candidateNumber
        if (!removed && stakerDelegations[delegatee] == 0) {
            candidateNumber -= 1;
        }

        // notify undelegation to distribute contract
        IDistribute(DISTRIBUTE_CONTRACT).notifyUndelegation(
            delegatee,
            _msgSender(),
            effectiveEpoch,
            stakerDelegations[delegatee],
            delegators[delegatee].length()
        );

        emit Undelegated(delegatee, _msgSender(), undelegation.amount, effectiveEpoch, unlockEpoch);

        if (
            !removed &&
            rewardStarted &&
            beforeRanking <= latestSequencerSetSize &&
            (stakerRankings[delegatee] > latestSequencerSetSize || stakerRankings[delegatee] > candidateNumber)
        ) {
            _updateSequencerSet();
        }
    }

    /// @notice delegator cliam delegate staking value
    function claimUndelegation() external nonReentrant {
        uint256 totalAmount;
        uint256 length = undelegations[_msgSender()].length;

        for (uint256 i = 0; i < length; ) {
            // if the reward is not started yet, claiming directly is allowed
            if (!rewardStarted || undelegations[_msgSender()][i].unlockEpoch <= currentEpoch()) {
                totalAmount += undelegations[_msgSender()][i].amount;

                // event params
                address delegatee = undelegations[_msgSender()][i].delegatee;
                uint256 unlockEpoch = undelegations[_msgSender()][i].unlockEpoch;
                uint256 amount = undelegations[_msgSender()][i].amount;

                if (i < length - 1) {
                    undelegations[_msgSender()][i] = undelegations[_msgSender()][length - 1];
                }
                undelegations[_msgSender()].pop();
                length = length - 1;

                emit UndelegationClaimed(delegatee, _msgSender(), unlockEpoch, amount);
            } else {
                i = i + 1;
            }
        }

        require(totalAmount > 0, "no Morph token to claim");
        _transfer(_msgSender(), totalAmount);
    }

    /// @notice delegator claim reward
    /// @param delegatee         delegatee address, claim all if empty
    /// @param targetEpochIndex  up to the epoch index that the delegator wants to claim
    function claimReward(address delegatee, uint256 targetEpochIndex) external nonReentrant {
        if (delegatee == address(0)) {
            IDistribute(DISTRIBUTE_CONTRACT).claimAll(_msgSender(), targetEpochIndex);
        } else {
            IDistribute(DISTRIBUTE_CONTRACT).claim(delegatee, _msgSender(), targetEpochIndex);
        }
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return current reward epoch index
    function currentEpoch() public view returns (uint256) {
        require(block.timestamp >= rewardStartTime, "reward is not started yet");
        return (block.timestamp - rewardStartTime) / REWARD_EPOCH;
    }

    /// @notice check if the user has staked to staker
    /// @param staker sequencers size
    function isStakingTo(address staker) external view returns (bool) {
        return _isStakingTo(staker);
    }

    /// @notice Get all the delegators which staked to staker
    /// @param staker staker address
    function getAllDelegators(address staker) external view returns (address[] memory) {
        return delegators[staker].values();
    }

    /// @notice Get the delegators length which staked to staker
    /// @param staker staker address
    function getDelegatorsLength(address staker) external view returns (uint256) {
        return delegators[staker].length();
    }

    /// @notice Get the delegators which staked to staker in pagination
    /// @param staker       staker address
    /// @param pageSize     page size
    /// @param pageIndex    page index
    function getAllDelegatorsInPagination(
        address staker,
        uint256 pageSize,
        uint256 pageIndex
    ) external view returns (uint256 delegatorsTotalNumber, address[] memory delegatorsInPage) {
        require(pageSize > 0, "invalid page size");

        delegatorsTotalNumber = delegators[staker].length();
        delegatorsInPage = new address[](pageSize);

        uint256 start = pageSize * pageIndex;
        uint256 end = pageSize * (pageIndex + 1) - 1;
        if (end > (delegatorsTotalNumber - 1)) {
            end = delegatorsTotalNumber - 1;
        }
        for (uint256 i = start; i <= end; i++) {
            delegatorsInPage[i] = delegators[staker].at(i);
        }
        return (delegatorsTotalNumber, delegatorsInPage);
    }

    /// @notice get stakers info
    function getStakesInfo(address[] calldata _stakerAddresses) external view returns (Types.StakerInfo[] memory) {
        Types.StakerInfo[] memory _stakers = new Types.StakerInfo[](_stakerAddresses.length);
        for (uint256 i = 0; i < _stakerAddresses.length; i++) {
            _stakers[i] = Types.StakerInfo(
                stakers[_stakerAddresses[i]].addr,
                stakers[_stakerAddresses[i]].tmKey,
                stakers[_stakerAddresses[i]].blsKey
            );
        }
        return _stakers;
    }

    /// @notice get stakers
    function getStakers() external view returns (Types.StakerInfo[] memory) {
        Types.StakerInfo[] memory _stakers = new Types.StakerInfo[](stakerAddresses.length);
        for (uint256 i = 0; i < stakerAddresses.length; i++) {
            _stakers[i] = Types.StakerInfo(
                stakers[stakerAddresses[i]].addr,
                stakers[stakerAddresses[i]].tmKey,
                stakers[stakerAddresses[i]].blsKey
            );
        }
        return _stakers;
    }

    /// @notice get staker addresses length
    function getStakerAddressesLength() external view returns (uint256) {
        return stakerAddresses.length;
    }

    /// @notice get undelegations of a delegator
    /// @param delegator delegator
    function getUndelegations(address delegator) external view returns (Undelegation[] memory) {
        return undelegations[delegator];
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice transfer morph token
    function _transfer(address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        IMorphToken(MORPH_TOKEN_CONTRACT).transfer(_to, _amount);
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        require(_amount > 0 && balanceAfter - balanceBefore == _amount, "morph token transfer failed");
    }

    /// @notice transfer morph token from
    function _transferFrom(address _from, address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        IMorphToken(MORPH_TOKEN_CONTRACT).transferFrom(_from, _to, _amount);
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        require(_amount > 0 && balanceAfter - balanceBefore == _amount, "morph token transfer failed");
    }

    /// @notice check if the user has staked to staker
    /// @param staker sequencers size
    function _isStakingTo(address staker) internal view returns (bool) {
        return delegations[staker][_msgSender()] > 0;
    }

    /// @notice select the size of staker with the largest staking amount, the max size is ${sequencerSetMaxSize}
    function _updateSequencerSet() internal {
        uint256 sequencerSize = sequencerSetMaxSize;
        if (rewardStarted) {
            if (candidateNumber < sequencerSetMaxSize) {
                sequencerSize = candidateNumber;
            }
        } else if (stakerAddresses.length < sequencerSetMaxSize) {
            sequencerSize = stakerAddresses.length;
        }
        address[] memory sequencerSet = new address[](sequencerSize);
        for (uint256 i = 0; i < sequencerSize; i++) {
            sequencerSet[i] = stakerAddresses[i];
        }
        ISequencer(SEQUENCER_CONTRACT).updateSequencerSet(sequencerSet);
        latestSequencerSetSize = sequencerSet.length;
    }

    /// @notice whether there is a undedeletion unclaimed
    function _unclaimed(address delegator, address delegatee) internal view returns (bool) {
        for (uint256 i = 0; i < undelegations[delegator].length; i++) {
            if (undelegations[delegator][i].delegatee == delegatee) {
                return true;
            }
        }
        return false;
    }
}
