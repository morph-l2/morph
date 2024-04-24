// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {IDistribute} from "./IDistribute.sol";

contract L2Staking is
    IL2Staking,
    Staking,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    // sequencer contract address
    address public immutable SEQUENCER_CONTRACT;
    // MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;
    // distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;

    // reward epoch, seconds of one day (3600 * 24)
    uint256 public immutable REWARD_EPOCH = 86400;
    // is reward start
    bool public REWARD_STARTED;
    // reward start time
    uint256 public REWARD_START_TIME;
    // max number of sequencers
    uint256 public SEQUENCER_MAX_SIZE;
    // undelegate lock epochs
    uint256 public UNDELEGATE_LOCK_EPOCHS;

    // latest sequencer set size
    uint256 public latestSequencerSetSize;
    // sequencer candidate number
    uint256 public candidateNumber;
    // sync from l1 staking
    address[] public stakerAddresses;
    // mapping(staker => staker_ranking)
    mapping(address => uint256) public stakerRankings;
    // mapping(staker => staker_info)
    mapping(address => Types.StakerInfo) public stakers;
    // mapping(staker => commission_percentage). default commission is zero if not set
    mapping(address => uint256) public commissions;

    // mapping(staker => total_amount)
    mapping(address => uint256) public stakerDelegations;
    // mapping(staker => delegators)
    mapping(address => EnumerableSetUpgradeable.AddressSet) internal delegators;
    // mapping(staker => mapping(delegator => amount))
    mapping(address => mapping(address => uint256)) public delegations;
    // mapping(delegator => Undelegation[])
    mapping(address => Undelegation[]) public undelegations;

    /*********************** modifiers **************************/

    /// @notice must be staker
    modifier isStaker(address addr) {
        require(stakerRankings[addr] > 0, "not staker");
        _;
    }

    /// @notice only staker allowed
    modifier onlyStaker() {
        require(stakerRankings[msg.sender] > 0, "only staker allowed");
        _;
    }

    /*********************** Constructor **************************/

    /**
     * @notice constructor
     * @param _otherStaking Address of the staking contract on the other network.
     */
    constructor(
        address payable _otherStaking
    ) Staking(payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER), _otherStaking) {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
    }

    /*********************** Init **************************/

    /**
     * @notice initializer
     * @param _admin                params admin
     * @param _sequencersMaxSize    max size of sequencer set
     * @param _undelegateLockEpochs undelegate lock epochs
     * @param _rewardStartTime      reward start time
     * @param _stakers              initial stakers, must be same as initial sequencer set in sequencer contract
     **/
    function initialize(
        address _admin,
        uint256 _sequencersMaxSize,
        uint256 _undelegateLockEpochs,
        uint256 _rewardStartTime,
        Types.StakerInfo[] calldata _stakers
    ) public initializer {
        require(_sequencersMaxSize > 0, "sequencersSize must greater than 0");
        SEQUENCER_MAX_SIZE = _sequencersMaxSize;
        require(_undelegateLockEpochs > 0, "invalid undelegateLockEpochs");
        UNDELEGATE_LOCK_EPOCHS = _undelegateLockEpochs;
        require(
            _rewardStartTime > block.timestamp &&
                _rewardStartTime % REWARD_EPOCH == 0,
            "invalid reward start time"
        );

        REWARD_START_TIME = _rewardStartTime;

        latestSequencerSetSize = _stakers.length;
        require(latestSequencerSetSize > 0, "invalid initial stakers");
        for (uint256 i = 0; i < latestSequencerSetSize; i++) {
            stakers[_stakers[i].addr] = _stakers[i];
            stakerAddresses.push(_stakers[i].addr);
            stakerRankings[_stakers[i].addr] = i + 1;
        }

        // transfer owner to admin
        _transferOwnership(_admin);

        super.__ReentrancyGuard_init();
    }

    /*********************** External Functions **************************/

    /**
     * @notice add staker, sync from L1
     * @param add   staker to add. {addr, tmKey, blsKey}
     */
    function addStaker(Types.StakerInfo memory add) external onlyOtherStaking {
        if (stakerRankings[add.addr] == 0) {
            stakerAddresses.push(add.addr);
            stakerRankings[add.addr] = stakerAddresses.length;
        }
        stakers[add.addr] = add;
        emit StakerAdded(add.addr, add.tmKey, add.blsKey);

        if (!REWARD_STARTED && stakerAddresses.length <= SEQUENCER_MAX_SIZE) {
            _updateSequencerSet();
        }
    }

    /**
     * @notice remove stakers, sync from L1. If new sequencer set is nil, layer2 will stop producing blocks
     * @param remove    staker to remove
     */
    function removeStakers(address[] memory remove) external onlyOtherStaking {
        bool updateSequencerSet = false;
        for (uint256 i = 0; i < remove.length; i++) {
            updateSequencerSet = REWARD_STARTED
                ? stakerRankings[remove[i]] <= latestSequencerSetSize
                : stakerRankings[remove[i]] <= SEQUENCER_MAX_SIZE;

            if (stakerRankings[remove[i]] > 0) {
                // update stakerRankings
                for (
                    uint256 j = stakerRankings[remove[i]] - 1;
                    j < stakerAddresses.length - 1;
                    j++
                ) {
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
        }
        emit StakerRemoved(remove);

        if (updateSequencerSet) {
            _updateSequencerSet();
        }
    }

    /**
     * @notice setCommissionRate set delegate commission percentage
     * @param commission    commission percentage
     */
    function setCommissionRate(uint256 commission) external onlyStaker {
        require(commission <= 20, "invalid commission");
        commissions[msg.sender] = commission;
        uint256 epochEffective = 0;
        if (REWARD_STARTED) {
            epochEffective = currentEpoch() + 1;
        }
        emit CommissionUpdated(msg.sender, commission, epochEffective);
    }

    /**
     * @notice delegator stake morph to staker
     * @param staker    stake to whom
     * @param amount    stake amount
     */
    function delegateStake(
        address staker,
        uint256 amount
    ) external isStaker(staker) nonReentrant {
        require(amount > 0, "invalid stake amount");
        // Re-staking to the same staker is not allowed before claiming undelegation
        require(!_unclaimed(msg.sender, staker), "undelegation unclaimed");

        stakerDelegations[staker] += amount;
        delegations[staker][msg.sender] += amount;
        delegators[staker].add(msg.sender); // will not be added repeatedly

        if (stakerDelegations[staker] == amount) {
            candidateNumber += 1;
        }

        uint256 beforeRanking = stakerRankings[staker];
        if (REWARD_STARTED && beforeRanking > 1) {
            // update stakers and rankings
            for (uint256 i = beforeRanking - 1; i > 0; i--) {
                if (
                    stakerDelegations[stakerAddresses[i]] >
                    stakerDelegations[stakerAddresses[i - 1]]
                ) {
                    address tmp = stakerAddresses[i - 1];
                    stakerAddresses[i - 1] = stakerAddresses[i];
                    stakerAddresses[i] = tmp;

                    stakerRankings[stakerAddresses[i - 1]] = i;
                    stakerRankings[stakerAddresses[i]] = i + 1;
                }
            }
        }
        uint256 effectiveEpoch = REWARD_STARTED ? currentEpoch() + 1 : 0;

        emit Delegated(
            staker,
            msg.sender,
            delegations[staker][msg.sender], // new amount, not incremental
            effectiveEpoch
        );

        // notify delegation to distribute contract
        IDistribute(DISTRIBUTE_CONTRACT).notifyDelegation(
            staker,
            msg.sender,
            effectiveEpoch,
            delegations[staker][msg.sender],
            stakerDelegations[staker],
            delegators[staker].length(),
            delegations[staker][msg.sender] == amount
        );

        // transfer morph token from delegator to this
        _transferFrom(msg.sender, address(this), amount);

        if (
            REWARD_STARTED &&
            beforeRanking > latestSequencerSetSize &&
            stakerRankings[staker] <= latestSequencerSetSize
        ) {
            _updateSequencerSet();
        }
    }

    /**
     * @notice delegator unstake morph
     * @param delegatee delegatee address
     */
    function undelegateStake(address delegatee) external nonReentrant {
        // must claim before you can delegate stake again
        require(!_unclaimed(msg.sender, delegatee), "undelegation unclaimed");
        require(_isStakingTo(delegatee), "staking amount is zero");

        // staker has been removed, unlock next epoch
        bool removed = stakerRankings[delegatee] == 0;

        uint256 effectiveEpoch;
        uint256 unlockEpoch;

        if (REWARD_STARTED) {
            effectiveEpoch = currentEpoch() + 1;
            unlockEpoch = removed
                ? effectiveEpoch
                : effectiveEpoch + UNDELEGATE_LOCK_EPOCHS;
        }

        Undelegation memory undelegation = Undelegation(
            delegatee,
            delegations[delegatee][msg.sender],
            unlockEpoch
        );

        undelegations[msg.sender].push(undelegation);
        delete delegations[delegatee][msg.sender];
        stakerDelegations[delegatee] -= undelegation.amount;
        delegators[delegatee].remove(msg.sender);

        // update candidateNumber
        if (!removed && stakerDelegations[delegatee] == 0) {
            candidateNumber -= 1;
        }

        uint256 beforeRanking = stakerRankings[delegatee];
        if (
            !removed &&
            REWARD_STARTED &&
            stakerRankings[delegatee] < candidateNumber
        ) {
            // update stakers and rankings
            for (
                uint256 i = stakerRankings[delegatee] - 1;
                i < candidateNumber - 1;
                i++
            ) {
                if (
                    stakerDelegations[stakerAddresses[i + 1]] >
                    stakerDelegations[stakerAddresses[i]]
                ) {
                    address tmp = stakerAddresses[i];
                    stakerAddresses[i] = stakerAddresses[i + 1];
                    stakerAddresses[i + 1] = tmp;

                    stakerRankings[stakerAddresses[i]] = i + 1;
                    stakerRankings[stakerAddresses[i + 1]] = i + 2;
                }
            }
        }

        // notify undelegation to distribute contract
        IDistribute(DISTRIBUTE_CONTRACT).notifyUndelegation(
            delegatee,
            msg.sender,
            effectiveEpoch,
            undelegation.amount,
            delegators[delegatee].length()
        );

        emit Undelegated(
            delegatee,
            msg.sender,
            undelegation.amount,
            effectiveEpoch,
            unlockEpoch
        );

        if (
            !removed &&
            REWARD_STARTED &&
            beforeRanking <= latestSequencerSetSize &&
            stakerRankings[delegatee] > latestSequencerSetSize
        ) {
            _updateSequencerSet();
        }
    }

    /**
     * @notice delegator cliam delegate staking value
     */
    function claimUndelegation() external nonReentrant {
        uint256 totalAmount;
        for (uint256 i = 0; i < undelegations[msg.sender].length; i++) {
            if (undelegations[msg.sender][i].unlockEpoch <= currentEpoch()) {
                totalAmount += undelegations[msg.sender][i].amount;
                if (undelegations[msg.sender].length > 1) {
                    undelegations[msg.sender][i] = undelegations[msg.sender][
                        undelegations[msg.sender].length - 1
                    ];
                }
                undelegations[msg.sender].pop();
            }
        }
        require(totalAmount > 0, "no Morph token to claim");
        _transfer(msg.sender, totalAmount);

        emit UndelegationClaimed(msg.sender, totalAmount);
    }

    /**
     * @notice delegator claim reward
     * @param delegatee         delegatee address, claim all if empty
     * @param targetEpochIndex  up to the epoch index that the delegator wants to claim
     */
    function claimReward(
        address delegatee,
        uint256 targetEpochIndex
    ) external nonReentrant {
        if (delegatee == address(0)) {
            IDistribute(DISTRIBUTE_CONTRACT).claimAll(
                msg.sender,
                targetEpochIndex
            );
        } else {
            IDistribute(DISTRIBUTE_CONTRACT).claim(
                delegatee,
                msg.sender,
                targetEpochIndex
            );
        }
    }

    /**
     * @notice claimCommission claim commission reward
     * @param targetEpochIndex   up to the epoch index that the staker wants to claim
     */
    function claimCommission(
        uint256 targetEpochIndex
    ) external onlyStaker nonReentrant {
        IDistribute(DISTRIBUTE_CONTRACT).claimCommission(
            msg.sender,
            targetEpochIndex
        );
    }

    /**
     * @notice update params
     * @param _sequencersMaxSize   max size of sequencer set
     */
    function updateParams(uint256 _sequencersMaxSize) external onlyOwner {
        require(
            _sequencersMaxSize > 0 && _sequencersMaxSize != SEQUENCER_MAX_SIZE,
            "invalid new sequencers size"
        );
        SEQUENCER_MAX_SIZE = _sequencersMaxSize;
        emit ParamsUpdated(SEQUENCER_MAX_SIZE);

        if (SEQUENCER_MAX_SIZE < latestSequencerSetSize) {
            // update sequencer set
            _updateSequencerSet();
        }
    }

    /**
     * @notice advance layer2 stage
     * @param _rewardStartTime   reward start time
     */
    function updateRewardStartTime(
        uint256 _rewardStartTime
    ) external onlyOwner {
        require(!REWARD_STARTED, "reward already started");
        require(
            _rewardStartTime > block.timestamp &&
                _rewardStartTime % REWARD_EPOCH == 0,
            "invalid reward start time"
        );
        REWARD_START_TIME = _rewardStartTime;
        emit RewardStartTimeUpdated(REWARD_START_TIME);
    }

    /**
     * @notice start reward
     */
    function startReward() external onlyOwner {
        require(
            block.timestamp >= REWARD_START_TIME,
            "can't start before reward start time"
        );
        require(candidateNumber > 0, "none candidate");

        REWARD_STARTED = true;

        // sort stakers by insertion sort
        for (uint256 i = 1; i < stakerAddresses.length; i++) {
            for (uint256 j = 0; j < i; j++) {
                if (
                    stakerDelegations[stakerAddresses[i]] >
                    stakerDelegations[stakerAddresses[j]]
                ) {
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

    /*********************** External View Functions **************************/

    /**
     * @notice return current reward epoch index
     */
    function currentEpoch() public view returns (uint256) {
        return
            block.timestamp > REWARD_START_TIME
                ? (block.timestamp - REWARD_START_TIME) / REWARD_EPOCH
                : 0;
    }

    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function isStakingTo(address staker) external view returns (bool) {
        return _isStakingTo(staker);
    }

    /**
     * @notice Get all the delegators which staked to staker
     * @param staker sequencers size
     */
    function getDelegators(
        address staker
    ) external view returns (address[] memory) {
        return delegators[staker].values();
    }

    /**
     * @notice get stakers info
     */
    function getStakesInfo(
        address[] memory _stakerAddresses
    ) external view returns (Types.StakerInfo[] memory) {
        Types.StakerInfo[] memory _stakers = new Types.StakerInfo[](
            _stakerAddresses.length
        );
        for (uint256 i = 0; i < _stakerAddresses.length; i++) {
            _stakers[i] = Types.StakerInfo(
                stakers[_stakerAddresses[i]].addr,
                stakers[_stakerAddresses[i]].tmKey,
                stakers[_stakerAddresses[i]].blsKey
            );
        }
        return _stakers;
    }

    /*********************** Internal Functions **************************/

    /**
     * @notice transfer morph token
     */
    function _transfer(address _to, uint256 _amount) internal {
        uint256 balanceBefore = IERC20Upgradeable(MORPH_TOKEN_CONTRACT)
            .balanceOf(_to);
        IERC20Upgradeable(MORPH_TOKEN_CONTRACT).transfer(_to, _amount);
        uint256 balanceAfter = IERC20Upgradeable(MORPH_TOKEN_CONTRACT)
            .balanceOf(_to);
        require(
            _amount > 0 && balanceAfter - balanceBefore == _amount,
            "morph token transfer failed"
        );
    }

    /**
     * @notice transfer morph token from
     */
    function _transferFrom(
        address _from,
        address _to,
        uint256 _amount
    ) internal {
        uint256 balanceBefore = IERC20Upgradeable(MORPH_TOKEN_CONTRACT)
            .balanceOf(_to);
        IERC20Upgradeable(MORPH_TOKEN_CONTRACT).transferFrom(
            _from,
            _to,
            _amount
        );
        uint256 balanceAfter = IERC20Upgradeable(MORPH_TOKEN_CONTRACT)
            .balanceOf(_to);
        require(
            _amount > 0 && balanceAfter - balanceBefore == _amount,
            "morph token transfer failed"
        );
    }

    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function _isStakingTo(address staker) internal view returns (bool) {
        return delegations[staker][msg.sender] > 0;
    }

    /**
     * @notice select the size of staker with the largest staking amount, the max size is ${SEQUENCER_MAX_SIZE}
     */
    function _updateSequencerSet() internal {
        uint256 sequencerSize = SEQUENCER_MAX_SIZE;
        if (REWARD_STARTED) {
            if (candidateNumber < SEQUENCER_MAX_SIZE) {
                sequencerSize = candidateNumber;
            }
        } else if (stakerAddresses.length < SEQUENCER_MAX_SIZE) {
            sequencerSize = stakerAddresses.length;
        }
        address[] memory sequencerSet = new address[](sequencerSize);
        for (uint256 i = 0; i < sequencerSize; i++) {
            sequencerSet[i] = stakerAddresses[i];
        }
        ISequencer(SEQUENCER_CONTRACT).updateSequencerSet(sequencerSet);
        latestSequencerSetSize = sequencerSet.length;
    }

    /**
     * @notice whether there is a undedeletion unclaimed
     */
    function _unclaimed(
        address delegator,
        address delegatee
    ) internal view returns (bool) {
        for (uint256 i = 0; i < undelegations[delegator].length; i++) {
            if (undelegations[delegator][i].delegatee == delegatee) {
                return true;
            }
        }
        return false;
    }
}
