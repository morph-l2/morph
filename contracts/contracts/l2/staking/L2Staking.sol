// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {CountersUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import {DoubleEndedQueueUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/DoubleEndedQueueUpgradeable.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {IMorphToken} from "../system/IMorphToken.sol";

contract L2Staking is IL2Staking, Staking, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;
    using CountersUpgradeable for CountersUpgradeable.Counter;
    using DoubleEndedQueueUpgradeable for DoubleEndedQueueUpgradeable.Bytes32Deque;

    /*************
     * Constants *
     *************/

    /// @notice reward epoch, seconds of one day (3600 * 24)
    uint256 private constant REWARD_EPOCH = 86400;

    /// @notice commission rate base
    uint256 private constant COMMISSION_RATE_BASE = 100; // 100%

    /// @notice MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;

    /// @notice sequencer contract address
    address public immutable SEQUENCER_CONTRACT;

    /// @notice system address
    address public immutable SYSTEM_ADDRESS;

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

    /// @notice nonce of staking L1 => L2 msg
    uint256 public nonce;

    /// @notice sync from l1 staking
    address[] public stakerAddresses;

    /// @notice staker rankings
    mapping(address staker => uint256 ranking) public stakerRankings;

    /// @notice stakers info
    mapping(address staker => Types.StakerInfo) public stakers;

    /// @notice staker commissions info, default commission percentage is zero if not set
    mapping(address staker => Commission) public commissions;

    /// @notice delegators of a delegatee
    mapping(address staker => EnumerableSetUpgradeable.AddressSet) internal delegators;

    /// @notice delegation of a delegatee
    mapping(address staker => DelegateeDelegation) public delegateeDelegations;

    /// @notice the delegation of a delegator
    mapping(address staker => mapping(address delegator => uint256 share)) public delegatorDelegations;

    // hash of the undelegate request => undelegate request
    mapping(bytes32 => UndelegateRequest) private _undelegateRequests;

    // delegator address => undelegate request queue(hash of the request)
    mapping(address => DoubleEndedQueueUpgradeable.Bytes32Deque) private _undelegateRequestsQueue;

    // delegator address => personal undelegate sequence
    mapping(address => CountersUpgradeable.Counter) private _undelegateSequence;

    // total blocks of an epoch
    EnumerableSetUpgradeable.AddressSet private epochSequencers;

    // sequencers that has produced blocks
    uint256 public epochTotalBlocks;

    // blocks produced by sequencers
    mapping(address seequencer => uint256) public epochSequencerBlocks;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice must be staker
    modifier onlyStaker(address addr) {
        if (stakerRankings[addr] == 0) revert ErrNotStaker();
        _;
    }

    /// @notice check nonce
    modifier checkNonce(uint256 _nonce) {
        if (_nonce != nonce) revert ErrInvalidNonce();
        _;
    }

    /// @notice Ensures that the caller message from system
    modifier onlSystem() {
        if (_msgSender() != SYSTEM_ADDRESS) revert ErrOnlySystem();
        _;
    }

    /// @notice Ensures that the caller message from system
    modifier onlyMorphTokenContract() {
        if (_msgSender() != MORPH_TOKEN_CONTRACT) revert ErrOnlyMorphTokenContract();
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
        SYSTEM_ADDRESS = Predeploys.SYSTEM;
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
        if (_owner == address(0)) revert ErrInvalidOwner();
        if (_sequencersMaxSize == 0) revert ErrZeroSequencerSize();
        if (_undelegateLockEpochs == 0) revert ErrZeroLockEpochs();
        if (_rewardStartTime <= block.timestamp || _rewardStartTime % REWARD_EPOCH != 0) revert ErrInvalidStartTime();
        if (_stakers.length == 0) revert ErrNoStakers();

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
    /// @param _nonce   msg nonce
    /// @param add      staker to add. {addr, tmKey, blsKey}
    function addStaker(uint256 _nonce, Types.StakerInfo calldata add) external onlyOtherStaking checkNonce(_nonce) {
        nonce = _nonce + 1;
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
    /// @param _nonce   msg nonce
    /// @param remove   staker to remove
    function removeStakers(uint256 _nonce, address[] calldata remove) external onlyOtherStaking checkNonce(_nonce) {
        nonce = _nonce + 1;
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
                if (delegateeDelegations[remove[i]].amount > 0) {
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

    /// @notice add staker. Only can be called when a serious bug causes L1 and L2 data to be out of sync
    /// @param _nonce   msg nonce
    /// @param add      staker to add. {addr, tmKey, blsKey}
    function emergencyAddStaker(uint256 _nonce, Types.StakerInfo calldata add) external onlyOwner checkNonce(_nonce) {
        nonce = _nonce + 1;
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

    /// @notice remove stakers. Only can be called when a serious bug causes L1 and L2 data to be out of sync
    /// @param _nonce   msg nonce
    /// @param remove   staker to remove
    function emergencyRemoveStakers(uint256 _nonce, address[] calldata remove) external onlyOwner checkNonce(_nonce) {
        nonce = _nonce + 1;
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
                if (delegateeDelegations[remove[i]].amount > 0) {
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

    /// @notice setCommissionPercentage set delegate commission percentage
    /// @param rate    commission percentage
    function setCommissionRate(uint256 rate) external onlyStaker(_msgSender()) {
        if (rate > 20) revert ErrInvalidCommissionRate();
        uint256 oldRate = commissions[_msgSender()].rate;
        commissions[_msgSender()] = Commission({rate: rate, amount: commissions[_msgSender()].amount});
        emit CommissionUpdated(_msgSender(), rate, oldRate);
    }

    /// @notice claimCommission claim unclaimed commission reward of a staker
    function claimCommission() external nonReentrant {
        if (commissions[_msgSender()].amount == 0) revert ErrNoCommission();
        uint256 amount = commissions[_msgSender()].amount;
        commissions[_msgSender()].amount = 0;
        _transfer(_msgSender(), amount);
        emit CommissionClaimed(_msgSender(), amount);
    }

    /// @notice update params
    /// @param _sequencerSetMaxSize   max size of sequencer set
    function updateSequencerSetMaxSize(uint256 _sequencerSetMaxSize) external onlyOwner {
        if (_sequencerSetMaxSize == 0 || _sequencerSetMaxSize == sequencerSetMaxSize) revert ErrInvalidSequencerSize();
        uint256 _oldSequencerSetMaxSize = sequencerSetMaxSize;
        sequencerSetMaxSize = _sequencerSetMaxSize;
        emit SequencerSetMaxSizeUpdated(_oldSequencerSetMaxSize, _sequencerSetMaxSize);

        uint256 candidate = rewardStarted ? candidateNumber : stakerAddresses.length;
        uint256 newSequencerSetSize = candidate < sequencerSetMaxSize ? candidate : sequencerSetMaxSize;
        // latest_sequencer_set_size = Min(candidate, old_sequencer_set_max_size)
        // new_sequencer_set_size = Min(candidate, new_sequencer_set_max_size)
        // if new_sequencer_set_size != latest_sequencer_set_size, update sequencer set
        if (newSequencerSetSize != latestSequencerSetSize) {
            _updateSequencerSet();
        }
    }

    /// @notice advance layer2 stage
    /// @param _rewardStartTime   reward start time
    function updateRewardStartTime(uint256 _rewardStartTime) external onlyOwner {
        if (rewardStarted) revert ErrRewardStarted();
        if (
            _rewardStartTime <= block.timestamp ||
            _rewardStartTime % REWARD_EPOCH != 0 ||
            _rewardStartTime == rewardStartTime
        ) revert ErrInvalidStartTime();

        uint256 _oldTime = rewardStartTime;
        rewardStartTime = _rewardStartTime;
        emit RewardStartTimeUpdated(_oldTime, _rewardStartTime);
    }

    /// @notice start reward
    function startReward() external onlyOwner {
        if (block.timestamp < rewardStartTime) revert ErrStartTimeNotReached();
        if (candidateNumber == 0) revert ErrNoCandidate();

        rewardStarted = true;

        // sort stakers by insertion sort
        for (uint256 i = 1; i < stakerAddresses.length; i++) {
            for (uint256 j = 0; j < i; j++) {
                if (delegateeDelegations[stakerAddresses[i]].amount > delegateeDelegations[stakerAddresses[j]].amount) {
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
    function delegate(address delegatee, uint256 amount) external onlyStaker(delegatee) nonReentrant {
        if (amount == 0) revert ErrZeroAmount();

        delegators[delegatee].add(_msgSender()); // will not be added repeatedly

        // ***********************************************************************************************

        if (!rewardStarted) {
            delegatorDelegations[delegatee][_msgSender()] += amount;
            delegateeDelegations[delegatee].amount += amount;
            delegateeDelegations[delegatee].share = delegateeDelegations[delegatee].amount; // {share == amount} before reward start
        } else {
            uint256 _tshare = delegateeDelegations[delegatee].share;
            uint256 _tAmount = delegateeDelegations[delegatee].amount;
            uint256 _uShare = delegatorDelegations[delegatee][_msgSender()];

            if (_tAmount == 0) {
                delegatorDelegations[delegatee][_msgSender()] = amount;
                delegateeDelegations[delegatee].share = amount;
                delegateeDelegations[delegatee].amount = amount;
            } else {
                delegatorDelegations[delegatee][_msgSender()] = _uShare + (amount * _tshare) / _tAmount;
                delegateeDelegations[delegatee].amount += amount;
                delegateeDelegations[delegatee].share = _tshare + (amount * _tshare) / _tAmount;
            }
        }

        // ***********************************************************************************************

        if (delegateeDelegations[delegatee].amount == amount) {
            candidateNumber += 1;
        }

        uint256 beforeRanking = stakerRankings[delegatee];
        if (rewardStarted && beforeRanking > 1) {
            // update stakers and rankings
            for (uint256 i = beforeRanking - 1; i > 0; i--) {
                if (
                    delegateeDelegations[stakerAddresses[i]].amount >
                    delegateeDelegations[stakerAddresses[i - 1]].amount
                ) {
                    address tmp = stakerAddresses[i - 1];
                    stakerAddresses[i - 1] = stakerAddresses[i];
                    stakerAddresses[i] = tmp;

                    stakerRankings[stakerAddresses[i - 1]] = i;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                }
            }
        }

        emit Delegated(delegatee, _msgSender(), amount, delegateeDelegations[delegatee].amount);

        // transfer morph token from delegator to this
        _transferFrom(_msgSender(), address(this), amount);

        if (
            rewardStarted && beforeRanking > latestSequencerSetSize && stakerRankings[delegatee] <= sequencerSetMaxSize
        ) {
            _updateSequencerSet();
        }
    }

    /// @notice delegator undelegate stake morph token
    /// @param delegatee    delegatee address
    /// @param amount       undelegate stake amount, undelegate all if set 0
    function undelegate(address delegatee, uint256 amount) external nonReentrant {
        if (amount == 0) revert ErrZeroAmount();
        if (_getDelegationAmount(delegatee, _msgSender()) == 0) revert ErrZeroShares();
        if (_getDelegationAmount(delegatee, _msgSender()) < amount) revert ErrInsufficientBalance();

        // weather staker has been removed
        bool removed = stakerRankings[delegatee] == 0;

        uint256 unlockEpoch = rewardStarted ? undelegateLockEpochs + 1 : 0;

        UndelegateRequest memory request = UndelegateRequest({amount: amount, unlockEpoch: unlockEpoch});
        bytes32 hash = keccak256(abi.encodePacked(_msgSender(), _useSequence(_msgSender())));
        // the hash should not exist in the queue
        // this will not happen in normal cases
        if (_undelegateRequests[hash].amount != 0) revert ErrRequestExisted();
        _undelegateRequests[hash] = request;
        _undelegateRequestsQueue[_msgSender()].pushBack(hash);

        // update delegatorDelegations & delegateeDelegations
        if (!rewardStarted) {
            delegatorDelegations[delegatee][_msgSender()] -= amount; // {share = amount} before reward start
            delegateeDelegations[delegatee].amount -= amount;
            delegateeDelegations[delegatee].share = delegateeDelegations[delegatee].amount; // {share = amount} before reward start
        } else {
            uint256 _tshare = delegateeDelegations[delegatee].share;
            uint256 _tAmount = delegateeDelegations[delegatee].amount;
            uint256 _uShare = delegatorDelegations[delegatee][_msgSender()];

            delegatorDelegations[delegatee][_msgSender()] = _uShare - (amount * _tshare) / _tAmount;
            delegateeDelegations[delegatee].amount -= amount;
            delegateeDelegations[delegatee].share = _tshare - (amount * _tshare) / _tAmount;
        }

        uint256 beforeRanking = stakerRankings[delegatee];
        if (!removed && rewardStarted && beforeRanking < candidateNumber) {
            // update stakers and rankings
            for (uint256 i = stakerRankings[delegatee] - 1; i < candidateNumber - 1; i++) {
                if (
                    delegateeDelegations[stakerAddresses[i + 1]].amount >
                    delegateeDelegations[stakerAddresses[i]].amount
                ) {
                    address tmp = stakerAddresses[i];
                    stakerAddresses[i] = stakerAddresses[i + 1];
                    stakerAddresses[i + 1] = tmp;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                    stakerRankings[stakerAddresses[i + 1]] = i + 2;
                }
            }
        }

        // update candidateNumber
        if (!removed && delegateeDelegations[delegatee].amount == 0) {
            candidateNumber -= 1;
        }

        uint256 delegateeAmount = delegateeDelegations[delegatee].amount;
        emit Undelegated(delegatee, _msgSender(), amount, delegateeAmount, unlockEpoch);

        if (
            !removed &&
            rewardStarted &&
            beforeRanking <= latestSequencerSetSize &&
            (stakerRankings[delegatee] > latestSequencerSetSize || stakerRankings[delegatee] > candidateNumber)
        ) {
            _updateSequencerSet();
        }
    }

    /// @notice delegator redelegate stake morph token to new delegatee
    /// @param delegateeFrom    old delegatee
    /// @param delegateeTo      new delegatee
    /// @param amount           amount
    function redelegate(
        address delegateeFrom,
        address delegateeTo,
        uint256 amount
    ) external onlyStaker(delegateeFrom) onlyStaker(delegateeTo) nonReentrant {
        if (amount == 0) revert ErrZeroAmount();
        if (_getDelegationAmount(delegateeFrom, _msgSender()) == 0) revert ErrZeroShares();
        if (_getDelegationAmount(delegateeFrom, _msgSender()) < amount) revert ErrInsufficientBalance();

        bool updateSequencerSet;

        // ***************************** undelegate from old delegatee ***************************** //
        // weather staker has been removed
        bool removed = stakerRankings[delegateeFrom] == 0;

        // update delegatorDelegations & delegateeDelegations
        if (!rewardStarted) {
            delegatorDelegations[delegateeFrom][_msgSender()] -= amount; // {share = amount} before reward start
            delegateeDelegations[delegateeFrom].amount -= amount;
            delegateeDelegations[delegateeFrom].share = delegateeDelegations[delegateeFrom].amount; // {share = amount} before reward start
        } else {
            uint256 _tshare = delegateeDelegations[delegateeFrom].share;
            uint256 _tAmount = delegateeDelegations[delegateeFrom].amount;
            uint256 _uShare = delegatorDelegations[delegateeFrom][_msgSender()];

            delegatorDelegations[delegateeFrom][_msgSender()] = _uShare - (amount * _tshare) / _tAmount;
            delegateeDelegations[delegateeFrom].amount -= amount;
            delegateeDelegations[delegateeFrom].share = _tshare - (amount * _tshare) / _tAmount;
        }

        uint256 beforeRanking = stakerRankings[delegateeFrom];
        if (!removed && rewardStarted && beforeRanking < candidateNumber) {
            // update stakers and rankings
            for (uint256 i = stakerRankings[delegateeFrom] - 1; i < candidateNumber - 1; i++) {
                if (
                    delegateeDelegations[stakerAddresses[i + 1]].amount >
                    delegateeDelegations[stakerAddresses[i]].amount
                ) {
                    address tmp = stakerAddresses[i];
                    stakerAddresses[i] = stakerAddresses[i + 1];
                    stakerAddresses[i + 1] = tmp;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                    stakerRankings[stakerAddresses[i + 1]] = i + 2;
                }
            }
        }

        // update candidateNumber
        if (!removed && delegateeDelegations[delegateeFrom].amount == 0) {
            candidateNumber -= 1;
        }

        if (
            !removed &&
            rewardStarted &&
            beforeRanking <= latestSequencerSetSize &&
            (stakerRankings[delegateeFrom] > latestSequencerSetSize || stakerRankings[delegateeFrom] > candidateNumber)
        ) {
            updateSequencerSet = true;
        }

        // ***************************** bond to new delegatee ***************************** //

        delegators[delegateeTo].add(_msgSender()); // will not be added repeatedly

        // update delegatorDelegations & delegateeDelegations
        if (!rewardStarted) {
            delegatorDelegations[delegateeTo][_msgSender()] += amount; // {share = amount} before reward start
            delegateeDelegations[delegateeTo].amount += amount;
            delegateeDelegations[delegateeTo].share = delegateeDelegations[delegateeTo].amount; // {share = amount} before reward start
        } else {
            uint256 _tshare = delegateeDelegations[delegateeTo].share;
            uint256 _tAmount = delegateeDelegations[delegateeTo].amount;
            uint256 _uShare = delegatorDelegations[delegateeTo][_msgSender()];

            delegatorDelegations[delegateeTo][_msgSender()] = _uShare + (amount * _tshare) / _tAmount;
            delegateeDelegations[delegateeTo].amount += amount;
            delegateeDelegations[delegateeTo].share = _tshare + (amount * _tshare) / _tAmount;
        }

        if (delegateeDelegations[delegateeTo].amount == amount) {
            candidateNumber += 1;
        }

        beforeRanking = stakerRankings[delegateeTo];
        if (rewardStarted && beforeRanking > 1) {
            // update stakers and rankings
            for (uint256 i = beforeRanking - 1; i > 0; i--) {
                if (
                    delegateeDelegations[stakerAddresses[i]].amount >
                    delegateeDelegations[stakerAddresses[i - 1]].amount
                ) {
                    address tmp = stakerAddresses[i - 1];
                    stakerAddresses[i - 1] = stakerAddresses[i];
                    stakerAddresses[i] = tmp;

                    stakerRankings[stakerAddresses[i - 1]] = i;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                }
            }
        }

        if (
            rewardStarted &&
            beforeRanking > latestSequencerSetSize &&
            stakerRankings[delegateeTo] <= sequencerSetMaxSize
        ) {
            updateSequencerSet = true;
        }

        // ********************************************************************************* //

        if (updateSequencerSet) {
            _updateSequencerSet();
        }

        uint256 delegateeFromAmount = delegateeDelegations[delegateeFrom].amount;
        uint256 delegateeToAmount = delegateeDelegations[delegateeTo].amount;
        emit Redelegated(delegateeFrom, delegateeTo, _msgSender(), amount, delegateeFromAmount, delegateeToAmount);
    }

    /// @notice delegator cliam delegate staking value
    /// @param number   the number of undelegate requests to be claimed. 0 means claim all
    /// @return amount  the total amount of MPH claimed
    function claimUndelegation(uint256 number) external nonReentrant returns (uint256) {
        // number == 0 means claim all
        // number should not exceed the length of the queue
        if (_undelegateRequestsQueue[_msgSender()].length() == 0) revert ErrNoUndelegateRequest();

        number = (number == 0 || number > _undelegateRequestsQueue[_msgSender()].length())
            ? _undelegateRequestsQueue[_msgSender()].length()
            : number;

        uint256 totalAmount;
        while (number != 0) {
            bytes32 hash = _undelegateRequestsQueue[_msgSender()].front();
            UndelegateRequest memory request = _undelegateRequests[hash];
            if (currentEpoch() < request.unlockEpoch) {
                break;
            }

            // remove from the queue
            _undelegateRequestsQueue[_msgSender()].popFront();

            totalAmount += request.amount;
            --number;
        }
        if (totalAmount == 0) revert ErrNoClaimableUndelegateRequest();

        _transfer(_msgSender(), totalAmount);

        emit UndelegationClaimed(_msgSender(), totalAmount);

        return totalAmount;
    }

    /// @dev distribute inflation by MorphTokenContract on epoch end
    /// @param amount        total reward amount
    function distribute(uint256 amount) external onlyMorphTokenContract {
        if (epochTotalBlocks != 0) {
            for (uint256 i = 0; i < epochSequencers.length(); i++) {
                uint256 commissionRate = commissions[epochSequencers.at(i)].rate;
                uint256 rewardAmount = (amount * epochSequencerBlocks[epochSequencers.at(i)]) / epochTotalBlocks;
                uint256 commissionAmount = (rewardAmount * commissionRate) / COMMISSION_RATE_BASE;
                uint256 delegatorRewardAmount = rewardAmount - commissionAmount;

                commissions[epochSequencers.at(i)].amount += commissionAmount;
                delegateeDelegations[epochSequencers.at(i)].amount += delegatorRewardAmount;

                emit Distributed(epochSequencers.at(i), delegatorRewardAmount, commissionAmount);
            }
        }

        // clean block record
        uint256 sequencerSetSize = epochSequencers.length();
        for (uint256 i = 0; i < sequencerSetSize; i++) {
            delete epochSequencerBlocks[epochSequencers.at(0)];
            // remove the first element, then the last one will be moved to the first
            epochSequencers.remove(epochSequencers.at(0));
        }
        epochTotalBlocks = 0;
    }

    /// @dev record block producer
    /// @param sequencerAddr    producer address
    function recordBlocks(address sequencerAddr) external onlSystem {
        epochSequencers.add(sequencerAddr);
        epochTotalBlocks += 1;
        epochSequencerBlocks[sequencerAddr] += 1;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return the total length of delegator's pending undelegate queue.
    /// @param delegator    delegator
    function pendingUndelegateRequest(address delegator) public view returns (uint256) {
        return _undelegateRequestsQueue[delegator].length();
    }

    /// @notice return the total number of delegator's claimable undelegate requests.
    /// @param delegator   delegator
    function claimableUndelegateRequest(address delegator) public view returns (uint256) {
        uint256 length = _undelegateRequestsQueue[delegator].length();
        uint256 count;
        for (uint256 i; i < length; ++i) {
            bytes32 hash = _undelegateRequestsQueue[delegator].at(i);
            UndelegateRequest memory request = _undelegateRequests[hash];
            if (currentEpoch() >= request.unlockEpoch) {
                ++count;
            } else {
                break;
            }
        }
        return count;
    }

    /// @notice return the sum of first `number` requests' MPH locked in delegator's undelegate queue.
    /// @param delegator    delegator
    /// @param number       number
    function lockedAmount(address delegator, uint256 number) public view returns (uint256) {
        // number == 0 means all
        // number should not exceed the length of the queue
        if (_undelegateRequestsQueue[delegator].length() == 0) {
            return 0;
        }
        number = (number == 0 || number > _undelegateRequestsQueue[delegator].length())
            ? _undelegateRequestsQueue[delegator].length()
            : number;

        uint256 _totalAmount;
        for (uint256 i; i < number; ++i) {
            bytes32 hash = _undelegateRequestsQueue[delegator].at(i);
            UndelegateRequest memory request = _undelegateRequests[hash];
            _totalAmount += request.amount;
        }
        return _totalAmount;
    }

    /// @notice return the undelegate request at _index.
    /// @param delegator    delegator
    /// @param _index       index
    function undelegateRequest(address delegator, uint256 _index) public view returns (UndelegateRequest memory) {
        bytes32 hash = _undelegateRequestsQueue[delegator].at(_index);
        return _undelegateRequests[hash];
    }

    /// @notice return the personal undelegate sequence of the delegator.
    /// @param delegator    delegator
    function undelegateSequence(address delegator) public view returns (uint256) {
        return _undelegateSequence[delegator].current();
    }

    /// @notice return current reward epoch index
    function currentEpoch() public view returns (uint256) {
        if (block.timestamp < rewardStartTime) revert ErrRewardNotStarted();
        return (block.timestamp - rewardStartTime) / REWARD_EPOCH;
    }

    /// @notice check if the user has staked to staker
    /// @param staker sequencers size
    function isStakingTo(address staker) external view returns (bool) {
        return _isStakingTo(staker);
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
        if (pageSize == 0) revert ErrInvalidPageSize();

        delegatorsTotalNumber = delegators[staker].length();
        delegatorsInPage = new address[](pageSize);

        uint256 start = pageSize * pageIndex;
        uint256 end = pageSize * (pageIndex + 1) - 1;
        if (end > (delegatorsTotalNumber - 1)) {
            end = delegatorsTotalNumber - 1;
        }
        uint256 i = start;
        uint256 j = 0;
        while (i <= end) {
            delegatorsInPage[j++] = delegators[staker].at(i++);
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

    /// @notice query all unclaimed commission of a delegatee
    /// @param delegatee     delegatee address
    function queryUnclaimedCommission(address delegatee) external view returns (uint256 amount) {
        return commissions[delegatee].amount;
    }

    /// @notice query delegation amount of a delegator
    /// @param delegatee     delegatee address
    /// @param delegator     delegator address
    function queryDelegationAmount(address delegatee, address delegator) external view returns (uint256 amount) {
        return _getDelegationAmount(delegatee, delegator);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice use sequence
    function _useSequence(address delegator) internal returns (uint256 current) {
        CountersUpgradeable.Counter storage sequence = _undelegateSequence[delegator];
        current = sequence.current();
        sequence.increment();
    }

    /// @notice transfer morph token
    function _transfer(address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        if (!IMorphToken(MORPH_TOKEN_CONTRACT).transfer(_to, _amount)) revert ErrTransferFailed();
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        if (_amount == 0 || balanceAfter - balanceBefore != _amount) revert ErrTransferFailed();
    }

    /// @notice transfer morph token from
    function _transferFrom(address _from, address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        if (!IMorphToken(MORPH_TOKEN_CONTRACT).transferFrom(_from, _to, _amount)) revert ErrTransferFailed();
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        if (_amount == 0 || balanceAfter - balanceBefore != _amount) revert ErrTransferFailed();
    }

    /// @notice check if the user has staked to delegatee
    /// @param delegatee    delegatee
    function _isStakingTo(address delegatee) internal view returns (bool) {
        return delegatorDelegations[delegatee][_msgSender()] > 0;
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

    /// @notice query delegation amount of a delegator
    /// @param delegatee     delegatee address
    /// @param delegator     delegator address
    function _getDelegationAmount(address delegatee, address delegator) internal view returns (uint256 amount) {
        return
            delegatorDelegations[delegatee][delegator] == 0
                ? 0
                : (delegateeDelegations[delegatee].amount * delegatorDelegations[delegatee][delegator]) /
                    delegateeDelegations[delegatee].share;
    }
}
