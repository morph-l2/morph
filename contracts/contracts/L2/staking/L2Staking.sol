// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {IDistribute} from "./IDistribute.sol";
import {IGov} from "./IGov.sol";

contract L2Staking is
    IL2Staking,
    Staking,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    // SequencerContract address
    address public immutable SEQUENCER_CONTRACT;
    // MorphTokenContract address
    address public immutable MORPH_TOKEN_CONTRACT;
    // DistributeContract address
    address public immutable DISTRIBUTE_CONTRACT;
    // GovContract address
    address public immutable GOV_CONTRACT;

    // undelegate lock blocks
    uint256 public lockBlocks;
    // contract stage. {0: init, 1: issue_token, 2: reward}
    uint256 public stage;

    // Sync from l1 staking
    EnumerableSetUpgradeable.AddressSet internal stakerList;

    // staker's all delegators
    mapping(address => EnumerableSetUpgradeable.AddressSet) internal delegators;

    // staker info
    mapping(address => Types.StakerInfo) public override stakers;

    // staker status
    mapping(address => bool) public override stakerStatus;

    // user staking info
    mapping(address => mapping(address => uint256)) public override stakings;

    // user unstaking info
    mapping(address => mapping(address => Unstaking))
        public
        override unstakings;

    // staker's morph amount
    mapping(address => uint256) public override stakersAmount;

    // total number of sequencers
    uint256 public override sequencersSize;

    /*********************** modifiers **************************/

    modifier isStaker(address _staker) {
        require(stakerList.contains(_staker), "staker not exist");
        _;
    }

    modifier checkStaker(address _staker) {
        require(stakerStatus[_staker], "staker not active");
        _;
    }

    /*********************** events **************************/
    /**
     * @notice stake info
     */
    event Delegated(
        address indexed staker,
        address indexed who,
        uint256 amount
    );
    /**
     * @notice unstake info
     */
    event UnDelegated(
        address indexed staker,
        address indexed who,
        uint256 amount
    );
    /**
     * @notice claim info
     */
    event Claimed(address indexed staker, address indexed who, uint256 amount);
    /**
     * @notice withdrawal info
     */
    event Withdrawn(
        address indexed staker,
        address indexed who,
        uint256 amount
    );
    /**
     * @notice params updated
     */
    event ParamsUpdated(uint256 sequencersSize);

    /*********************** Constructor **************************/
    /**
     * @notice constructor
     * @param _otherStaking Address of the staking contract on the other network.
     */
    constructor(
        address payable _otherStaking
    ) Staking(payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER), _otherStaking) {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
        GOV_CONTRACT = Predeploys.GOV;
    }

    /*********************** Init **************************/
    /* * @notice initializer
     * @param _admin            params admin
     * @param _sequencersSize   size of sequencer set
     */
    function initialize(
        address _admin,
        uint256 _sequencersSize
    ) public initializer {
        require(_sequencersSize > 0, "sequencersSize must greater than 0");
        sequencersSize = _sequencersSize;

        // transfer owner to admin
        _transferOwnership(_admin);

        super.__ReentrancyGuard_init();
    }

    /*********************** External Functions **************************/
    /**
     * @notice delegator stake morph to staker
     * @param staker    stake to whom
     * @param amount    stake amount
     */
    function delegateStake(
        address staker,
        uint256 amount
    ) external isStaker(staker) checkStaker(staker) nonReentrant {
        // Re-staking to the same staker is not allowed during the lock-up period
        require(
            block.number >= unstakings[staker][msg.sender].unlock,
            "re-staking cannot be made during the lock-up period"
        );

        uint256 balanceBefore = IERC20(MORPH_TOKEN_CONTRACT).balanceOf(
            address(this)
        );
        IERC20(MORPH_TOKEN_CONTRACT).transferFrom(
            msg.sender,
            address(this),
            amount
        );
        uint256 balanceAfter = IERC20(MORPH_TOKEN_CONTRACT).balanceOf(
            address(this)
        );
        require(
            balanceAfter > balanceBefore &&
                balanceAfter - balanceBefore == amount,
            "morph token transfer fail"
        );

        uint256 userStakingAmount = stakings[staker][msg.sender];
        stakings[staker][msg.sender] = userStakingAmount + amount;

        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount + amount;

        delegators[staker].add(msg.sender);

        emit Delegated(staker, msg.sender, amount);

        // update sequencer set
        _updateSequencerSet();

        // push record to distribute
        // IDistribute(DISTRIBUTE_CONTRACT).notifyDelegate(
        //     block.number / epoch,
        //     staker,
        //     msg.sender,
        //     amount,
        //     block.number
        // );
    }

    /**
     * @notice delegator unstake morph
     * @param staker stake to whom
     */
    function undelegateStake(
        address staker
    ) external isStaker(staker) nonReentrant {
        Unstaking memory info = unstakings[staker][msg.sender];

        require(info.amount == 0, "needs to be withdrawn");
        require(_isStakingTo(staker), "staking amount is zero");

        uint256 delegatorStakingAmount = stakings[staker][msg.sender];

        // record undeledate
        uint256 rewardEpoch = IGov(GOV_CONTRACT).rewardEpoch();
        uint256 unlock = (block.number / rewardEpoch + 1) * rewardEpoch;
        info.amount = delegatorStakingAmount;
        info.unlock = unlock;

        // update unstaking info
        unstakings[staker][msg.sender] = info;

        // update staking info
        stakings[staker][msg.sender] = 0;
        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount - delegatorStakingAmount;

        emit UnDelegated(staker, msg.sender, delegatorStakingAmount);

        // update sequencer set
        _updateSequencerSet();

        // push record to distribute
        // IDistribute(DISTRIBUTE_CONTRACT).notifyUnDelegate(
        //     staker,
        //     msg.sender,
        //     block.number / epoch
        // );
    }

    /**
     * @notice redelegate stake
     */
    function redelegateStake() external nonReentrant {
        // TODO
    }

    /**
     * @notice delegator withdrawal
     * @param staker stake to whom
     */
    function withdrawal(address staker) external isStaker(staker) nonReentrant {
        require(
            unstakings[staker][msg.sender].amount > 0,
            "no information on unstaking"
        );
        require(
            block.number >= unstakings[staker][msg.sender].unlock,
            "withdrawal cannot be made during the lock-up period"
        );

        uint256 unstakingAmount = unstakings[staker][msg.sender].amount;

        uint256 balanceBefore = IERC20(MORPH_TOKEN_CONTRACT).balanceOf(
            address(this)
        );
        IERC20(MORPH_TOKEN_CONTRACT).transfer(msg.sender, unstakingAmount);
        uint256 balanceAfter = IERC20(MORPH_TOKEN_CONTRACT).balanceOf(
            address(this)
        );
        require(
            balanceBefore > balanceAfter &&
                balanceBefore - balanceAfter == unstakingAmount,
            "morph token transfer fail"
        );

        emit Withdrawn(staker, msg.sender, unstakingAmount);

        delete unstakings[staker][msg.sender];
    }

    /**
     * @notice delegator claim reward
     * @param staker stake to whom
     */
    function claim(address staker) external isStaker(staker) nonReentrant {
        // claim reward
        // reward from distribution
        IDistribute(DISTRIBUTE_CONTRACT).claim(staker, msg.sender);

        // emit Claimed(staker, msg.sender, unstakingAmount);
    }

    /**
     * @notice delegator claim all reward
     */
    function claimAll() external nonReentrant {
        // claim all reward
        // reward from distribution，if staking to multiple staker.
        IDistribute(DISTRIBUTE_CONTRACT).claimAll(msg.sender);
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
     * @notice Get all stakers
     */
    function getStakers() external view returns (address[] memory) {
        return _getStakers();
    }

    /**
     * @notice get stakers info
     */
    function getStakesInfo(
        address[] memory stakerAddrs
    ) external view returns (Types.StakerInfo[] memory) {
        Types.StakerInfo[] memory _stakers = new Types.StakerInfo[](
            stakerAddrs.length
        );
        for (uint256 i = 0; i < stakerAddrs.length; i++) {
            _stakers[i] = Types.StakerInfo(
                stakers[stakerAddrs[i]].addr,
                stakers[stakerAddrs[i]].tmKey,
                stakers[stakerAddrs[i]].blsKey
            );
        }
        return _stakers;
    }

    /**
     * @notice update params
     * @param _sequencersSize   sequencers size
     */
    function updateParams(uint256 _sequencersSize) external onlyOwner {
        require(
            _sequencersSize > 0 &&
                _sequencersSize != sequencersSize &&
                _sequencersSize < stakerList.length(),
            "invalid new sequencers size"
        );
        sequencersSize = _sequencersSize;
        emit ParamsUpdated(sequencersSize);
        // @todo check if the size less than current sequencer set size
        if (
            // TODO: update accurate judgment conditions
            sequencersSize <
            ISequencer(SEQUENCER_CONTRACT).getCurrentSeqeuncerSetSize()
        ) {
            // update sequencer set
            _updateSequencerSet();
        }
    }

    /**
     * @notice add staker, sync from L1
     * @param add   staker to add. {addr, tmKey, blsKey}
     */
    function addStaker(Types.StakerInfo memory add) external onlyOtherStaking {
        if (!stakerList.contains(add.addr)) {
            stakerList.add(add.addr);
            stakers[add.addr] = add;
        }
        // TODO: update sequencer set
        if (stage < 2 && true) {
            _updateSequencerSet();
        }
    }

    /**
     * @notice remove stakers, sync from L1
     * @param remove    staker to remove
     */
    function removeStakers(address[] memory remove) external onlyOtherStaking {
        for (uint256 i = 0; i < remove.length; i++) {
            if (!stakerList.contains(remove[i])) {
                stakerList.remove(remove[i]);
            }
        }
        // TODO: update sequencer set
        if (stage < 2 && true) {
            _updateSequencerSet();
        }
    }

    function stageAdvance() external onlyOwner {
        if (stage == 0) {
            // TODO: check & do sth
            stage = 1;
            return;
        }
        if (stage == 1) {
            // TODO: check & do sth
            stage = 2;
            return;
        }
    }

    /*********************** Internal Functions **************************/
    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function _isStakingTo(address staker) internal view returns (bool) {
        return stakings[staker][msg.sender] > 0;
    }

    /**
     * @notice select the size of staker with the largest staking amount, the max size is ${sequencersSize}
     */
    function _updateSequencerSet() internal {
        address[] memory mStakers = _getSortedStakers();

        uint256 size = sequencersSize;
        if (mStakers.length < sequencersSize) {
            size = mStakers.length;
        }

        // determination of total update size
        uint256 updateSize = 0;
        for (uint256 i = 0; i < size; i++) {
            // staker is active
            // @todo checkou amount > 0
            // uint256 amount = stakersAmount[mStakers[i]];
            if (stakerStatus[mStakers[i]]) {
                updateSize = updateSize + 1;
            }
        }

        if (updateSize == 0) {
            revert("the number of updates required is 0");
        }

        uint256 index = 0;
        address[] memory newSequencerSet = new address[](updateSize);
        for (uint256 i = 0; i < size; i++) {
            Types.StakerInfo memory info = stakers[mStakers[i]];
            // staker is active
            // @todo checkou amount > 0
            // uint256 amount = stakersAmount[mStakers[i]];
            if (stakerStatus[mStakers[i]]) {
                newSequencerSet[index] = info.addr;
                index = index + 1;
            }
        }

        // update sequencer set
        ISequencer(SEQUENCER_CONTRACT).updateSequencerSet(newSequencerSet);
    }

    /**
     * @notice get all stakers
     */
    function _getStakers() internal view returns (address[] memory) {
        return stakerList.values();
    }

    /**
     * @notice sort stakers by amount
     */
    function _getSortedStakers() internal view returns (address[] memory) {
        address[] memory mStakers = _getStakers();

        for (uint256 i = 0; i < mStakers.length; i++) {
            uint256 maxIndex = i;
            for (uint256 j = i + 1; j < mStakers.length; j++) {
                uint256 amount0 = stakersAmount[mStakers[maxIndex]];
                uint256 amount1 = stakersAmount[mStakers[j]];
                if (amount1 > amount0) {
                    maxIndex = j;
                }
            }
            if (i != maxIndex) {
                address temp = mStakers[i];
                mStakers[i] = mStakers[maxIndex];
                mStakers[maxIndex] = temp;
            }
        }

        return mStakers;
    }
}
