// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {ISequencer} from "./ISequencer.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Staking} from "../../libraries/staking/Staking.sol";

contract L2Staking is IL2Staking, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    // unstaking info
    struct UnStakingInfo {
        uint256 amount;
        uint256 unlock;
    }

    /**
     * @notice Struct representing a staker information.
     *
     * @custom:field addr   Address of the sequencer.
     * @custom:field tmKey  Tendermint key(ED25519) of the seuqencer.
     * @custom:field blsKey BLS key of the seuqencer.
     * @custom:field active Active Status.
     */
    struct StakerInfo {
        address addr;
        bytes32 tmKey;
        bytes blsKey;
        bool active;
    }

    // Sync from l1 staking
    EnumerableSetUpgradeable.AddressSet internal allStakers;

    mapping(address => StakerInfo) public override stakerInfo;

    // user staking info
    mapping(address => mapping(address => uint256)) public override stakingInfo;

    // user unstaking info
    mapping(address => mapping(address => UnStakingInfo)) public override unstakingInfo;

    // staker's morph amount
    mapping(address => uint256) public override stakersAmount;

    // staker's all delegators
    mapping(address => EnumerableSetUpgradeable.AddressSet) internal delegators;

    // sequencer contract
    ISequencer public iSequencerSet;

    // morph contract
    IERC20 public iMorphToken;

    // Disdribute
    // IDisdribute

    // staking limit
    uint256 public override limit;

    // total number of sequencers
    uint256 public override sequencersSize;

    // epoch number
    uint256 public override epoch;

    /*********************** modifiers **************************/
    modifier isStaker(address _staker) {
        require(allStakers.contains(_staker), "staker not exist");
        _;
    }

    modifier checkStaker(address _staker) {
        require(stakerInfo[_staker].active, "staker not active");
        _;
    }

    /*********************** events **************************/
    /**
     * @notice stake info
     */
    event Delegated(address indexed staker, address indexed who, uint256 amount);
    /**
     * @notice unstake info
     */
    event UnDelegated(address indexed staker, address indexed who, uint256 amount);
    /**
     * @notice claim info
     */
    event Claimed(address indexed staker, address indexed who, uint256 amount);
    /**
     * @notice params updated
     */
    event ParamsUpdated(uint256 sequencersSize, uint256 limit, uint256 epoch);

    /*********************** Constructor **************************/
    constructor() {
        // _disableInitializers();
    }

    /*********************** Init **************************/
    /* * @notice initializer
     * @param _admin params admin
     * @param _sequencerContract sequencer contract address
     * @param _sequencersSize size of sequencer set
     * @param _limit smallest staking value
     * @param _lock withdraw lock time
     */
    function initialize(
        address _admin,
        address _sequencerContract,
        address _morphContract,
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _epoch
    ) public initializer {
        require(_sequencerContract != address(0), "invalid sequencer contract");
        require(_morphContract != address(0), "invalid morph contract");
        require(_sequencersSize > 0, "sequencersSize must greater than 0");
        require(_limit > 0, "staking limit must greater than 0");
        require(_epoch > 0, "epoch must greater than 0");

        iSequencerSet = ISequencer(_sequencerContract);
        iMorphToken = IERC20(_morphContract);

        // init params
        sequencersSize = _sequencersSize;
        limit = _limit;
        epoch = _epoch;

        // transfer owner to admin
        _transferOwnership(_admin);

        super.__ReentrancyGuard_init();
    }

    /*********************** External Functions **************************/
    /**
     * @notice user stake morph to staker
     * @param staker stake to
     * @param amount stake amount
     */
    function delegateStake(
        address staker,
        uint256 amount
    ) external isStaker(staker) checkStaker(staker) nonReentrant {
        uint256 userStakingAmount = stakingInfo[staker][msg.sender];
        if (allStakers.contains(msg.sender)) {
            if (userStakingAmount < limit) {
                // staker self, need to meet a minimum number
                require(amount >= limit, "staking amount is not enough");
            }
        }

        // Re-staking to the same staker is not allowed during the lock-up period
        require(
            block.number >= unstakingInfo[staker][msg.sender].unlock,
            "re-staking cannot be made during the lock-up period"
        );

        uint256 balanceBefore = iMorphToken.balanceOf(address(this));
        iMorphToken.transferFrom(msg.sender, address(this), amount);
        uint256 balanceAfter = iMorphToken.balanceOf(address(this));
        require(
            balanceAfter > balanceBefore && balanceAfter - balanceBefore == amount,
            "morph token transfer fail"
        );

        stakingInfo[staker][msg.sender] = userStakingAmount + amount;

        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount + amount;

        delegators[staker].add(msg.sender);

        emit Delegated(staker, msg.sender, amount);

        // update sequencer set
        _updateSequencers();

        // @todo
        // push record to distribute
    }

    /**
     * @notice user unstake morph
     * @param staker stake to
     */
    function unDelegateStake(address staker) external isStaker(staker) nonReentrant {
        UnStakingInfo memory info = unstakingInfo[staker][msg.sender];

        require(info.amount == 0, "need claim");
        require(_isStakingTo(staker), "staking amount is zero");

        uint256 userStakingAmount = stakingInfo[staker][msg.sender];

        // @todo
        // staker self, consider staking amount whether to keep the limit

        // record undeledate
        uint256 unlock = (block.number / epoch + 1) * epoch;
        info.amount = userStakingAmount;
        info.unlock = unlock;

        // update unstaking info
        unstakingInfo[staker][msg.sender] = info;

        // update staking info
        stakingInfo[staker][msg.sender] = 0;
        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount - userStakingAmount;

        emit UnDelegated(staker, msg.sender, userStakingAmount);

        // update sequencer set
        _updateSequencers();

        // @todo
        // push record to distribute
    }

    /**
     * @notice user claim morph
     * @param staker stake to
     */
    function claim(address staker) external nonReentrant {
        require(unstakingInfo[staker][msg.sender].amount > 0, "no information on unstaking");
        require(
            block.number >= unstakingInfo[staker][msg.sender].unlock,
            "claim cannot be made during the lock-up period"
        );

        uint256 unstakingAmount = unstakingInfo[staker][msg.sender].amount;

        uint256 balanceBefore = iMorphToken.balanceOf(address(this));
        iMorphToken.transfer(msg.sender, unstakingAmount);
        uint256 balanceAfter = iMorphToken.balanceOf(address(this));
        require(
            balanceBefore > balanceAfter && balanceBefore - balanceAfter == unstakingAmount,
            "morph token transfer fail"
        );

        emit Claimed(staker, msg.sender, unstakingAmount);

        delete unstakingInfo[staker][msg.sender];

        // @todo
        // distribute claim
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
    function getDelegators(address staker) external view returns (address[] memory) {
        return delegators[staker].values();
    }

    /**
     * @notice Get all stakers
     */
    function getStakers() external view returns (address[] memory) {
        return _getStakers();
    }

    /**
     * @notice update params
     * @param _sequencersSize sequencers size
     * @param _limit smallest staking value
     * @param _epoch epoch number
     */
    function updateParams(
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _epoch
    ) external onlyOwner {
        require(
            _sequencersSize > 0 &&
                _sequencersSize != sequencersSize &&
                _sequencersSize >= allStakers.length(),
            "invalid new sequencers size"
        );
        if (_limit > 0) {
            limit = _limit;
        }
        if (_epoch > 0) {
            epoch = _epoch;
        }
        sequencersSize = _sequencersSize;
        emit ParamsUpdated(sequencersSize, limit, epoch);

        // @todo check if the size less than current sequencer set size
        // if (sequencersSize < iSequencerSet.getCount()) {
        //     // update sequencer set
        //     _updateSequencers();
        // }
    }

    /**
     * @notice update stakers
     * @param stakers, sync from l1, {addr, tmKey, blsKey}
     * @param active, active & inActive
     */
    function updateStakers(Types.StakerInfo[] memory stakers, bool active) external {
        for (uint256 i = 0; i < stakers.length; i++) {
            if (allStakers.contains(stakers[i].addr)) {
                stakerInfo[stakers[i].addr].active = active;
            } else {
                allStakers.add(stakers[i].addr);
                stakerInfo[stakers[i].addr] = StakerInfo(
                    stakers[i].addr,
                    stakers[i].tmKey,
                    stakers[i].blsKey,
                    active
                );
            }
        }
    }

    /*********************** Internal Functions **************************/
    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function _isStakingTo(address staker) internal view returns (bool) {
        return stakingInfo[staker][msg.sender] > 0;
    }

    /**
     * @notice Select the size of staker with the largest staking amount
     */
    function _updateSequencers() internal {
        address[] memory mStakers = _getSortedStakers();

        uint256 size = sequencersSize;
        if (mStakers.length < sequencersSize) {
            size = mStakers.length;
        }

        address[] memory _update = new address[](size);
        for (uint256 i = 0; i < size; i++) {
            // need staker morph amount >= limit
            // staker is active
            uint256 amount = stakersAmount[mStakers[i]];
            if (amount >= limit && stakerInfo[mStakers[i]].active) {
                _update[i] = mStakers[i];
            }
        }

        //todo update sequencer set
        // iSequencerSet.updateSeqencerSet(_update)
    }

    /**
     * @notice get all stakers
     */
    function _getStakers() internal view returns (address[] memory) {
        return allStakers.values();
    }

    /**
     * @notice sort stakers by amount
     */
    function _getSortedStakers() internal view returns (address[] memory) {
        address[] memory mStakers = _getStakers();

        for (uint256 i = 0; i < mStakers.length; i++) {
            uint256 amount0 = stakersAmount[mStakers[i]];
            uint256 maxIndex = i;

            for (uint256 j = i + 1; j < mStakers.length; j++) {
                uint256 amount1 = stakersAmount[mStakers[j]];
                if (amount1 > amount0) {
                    maxIndex = j;
                }

                address temp = mStakers[i];
                mStakers[i] = mStakers[maxIndex];
                mStakers[maxIndex] = temp;
            }
        }

        return mStakers;
    }
}
