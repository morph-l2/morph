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
        bool exit;
    }

    // Sync from l1 staking
    EnumerableSetUpgradeable.AddressSet internal stakers;

    // user staking info
    mapping(address => mapping(address => uint256)) public stakingInfo;

    // user unstaking info
    mapping(address => mapping(address => UnStakingInfo)) public unstakingInfo;

    // staker's morph amount
    mapping(address => uint256) public stakersAmount;

    // staker's all delegators
    mapping(address => EnumerableSetUpgradeable.AddressSet) internal delegators;

    // sequencer contract
    ISequencer public iSequencerSet;

    // morph contract
    IERC20 public iMorphToken;

    // exit lock blocks
    uint256 public override lock;

    // staking limit
    uint256 public override limit;

    // total number of sequencers
    uint256 public override sequencersSize;

    /*********************** modifiers **************************/
    modifier isStaker(address _staker) {
        require(stakers.contains(_staker), "staker not exist");
        _;
    }

    /*********************** events **************************/
    /**
     * @notice stake info
     */
    event Staked(address indexed who, address indexed staker, uint256 amount);
    /**
     * @notice unstake info
     */
    event UnStaked(address indexed who, address indexed staker, uint256 amount);
    /**
     * @notice claim info
     */
    event Withdrawed(address indexed who, address indexed staker, uint256 amount);
    /**
     * @notice params updated
     */
    event ParamsUpdated(uint256 _sequencersSize, uint256 _limit, uint256 _lock);

    /*********************** Constructor **************************/
    constructor() {
        _disableInitializers();
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
        uint256 _lock
    ) public initializer {
        require(_sequencerContract != address(0), "invalid sequencer contract");
        require(_morphContract != address(0), "invalid morph contract");
        require(_sequencersSize > 0, "sequencersSize must greater than 0");
        require(_limit > 0, "staking limit must greater than 0");

        iSequencerSet = ISequencer(_sequencerContract);
        iMorphToken = IERC20(_morphContract);

        // init params
        sequencersSize = _sequencersSize;
        limit = _limit;
        lock = _lock;

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
    function stake(address staker, uint256 amount) external isStaker(staker) nonReentrant {
        uint256 userStakingAmount = stakingInfo[msg.sender][staker];
        if (stakers.contains(msg.sender)) {
            if (userStakingAmount < limit) {
                // staker self, need to meet a minimum number
                require(amount >= limit, "staking amount is not enough");
            }
        }

        // Re-staking to the same staker is not allowed during the lock-up period
        require(block.number > unstakingInfo[msg.sender][staker].unlock, "not allowed");

        uint256 balanceBefore = iMorphToken.balanceOf(address(this));
        iMorphToken.transferFrom(msg.sender, address(this), amount);
        uint256 balanceAfter = iMorphToken.balanceOf(address(this));
        require(
            balanceAfter > balanceBefore && balanceAfter - balanceBefore == amount,
            "morph token transfer fail"
        );

        stakingInfo[msg.sender][staker] = userStakingAmount + amount;

        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount + amount;

        delegators[staker].add(msg.sender);

        emit Staked(msg.sender, staker, amount);

        // update sequencer set
        _updateSequencers();
    }

    /**
     * @notice user unstake morph to staker
     * @param staker stake to
     */
    function unstake(address staker) external isStaker(staker) nonReentrant {
        UnStakingInfo memory info = unstakingInfo[msg.sender][staker];

        require(!info.exit, "no withdrawal");
        require(_isStakingTo(staker), "staking amount is zero");

        uint256 userStakingAmount = stakingInfo[msg.sender][staker];

        // @todo
        // staker self, need to consider limit amount
        stakingInfo[msg.sender][staker] = 0;

        uint256 stakerAmount = stakersAmount[staker];
        stakersAmount[staker] = stakerAmount - userStakingAmount;

        info.amount = userStakingAmount;
        info.unlock = block.number + lock;
        info.exit = true;

        // update unstaking info
        unstakingInfo[msg.sender][staker] = info;

        emit UnStaked(msg.sender, staker, userStakingAmount);

        // update sequencer set
        _updateSequencers();
    }

    /**
     * @notice user withdrawal
     * @param staker stake to
     */
    function withdrawal(address staker) external nonReentrant {
        require(
            unstakingInfo[msg.sender][staker].amount > 0 &&
                block.number > unstakingInfo[msg.sender][staker].unlock,
            "invalid withdrawal"
        );

        uint256 unstakingAmount = unstakingInfo[msg.sender][staker].amount;
        uint256 balanceBefore = iMorphToken.balanceOf(address(this));
        iMorphToken.transfer(msg.sender, unstakingAmount);
        uint256 balanceAfter = iMorphToken.balanceOf(address(this));
        require(
            balanceBefore > balanceAfter && balanceBefore - balanceAfter == unstakingAmount,
            "morph token transfer fail"
        );

        emit Withdrawed(msg.sender, staker, unstakingAmount);

        delete unstakingInfo[msg.sender][staker];
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
     * @param _limit smallest staking value
     * @param _lock withdraw lock time
     * @param _sequencersSize sequencers size
     */
    function updateParams(
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _lock
    ) external onlyOwner {
        require(
            _sequencersSize > 0 &&
                _sequencersSize != sequencersSize &&
                _sequencersSize >= stakers.length(),
            "invalid new sequencers size"
        );
        if (_limit > 0) {
            limit = _limit;
        }
        if (_lock > 0) {
            lock = _lock;
        }
        sequencersSize = _sequencersSize;
        emit ParamsUpdated(sequencersSize, limit, lock);

        // @todo check if the size less than current sequencer set size
        // if (sequencersSize < iSequencerSet.getCount()) {
        //     // update sequencer set
        //     _updateSequencers();
        // }
    }

    /**
     * @notice update stakers
     */
    function updateStakers(
        Types.StakerInfo[] memory add,
        Types.StakerInfo[] memory remove
    ) external {}

    /*********************** Internal Functions **************************/
    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function _isStakingTo(address staker) internal view returns (bool) {
        return stakingInfo[msg.sender][staker] > 0;
    }

    /**
     * @notice Select the size of staker with the largest staking amount
     */
    function _updateSequencers() internal {
        address[] memory _stakers = _getSortedStakers();

        uint256 size = sequencersSize;
        if (_stakers.length < sequencersSize) {
            size = _stakers.length;
        }

        address[] memory _update = new address[](size);
        for (uint256 i = 0; i < size; i++) {
            _update[i] = _stakers[i];
        }

        //todo update sequencer set
        // iSequencerSet.updateSeqencerSet(_update)
    }

    /**
     * @notice get all stakers
     */
    function _getStakers() internal view returns (address[] memory) {
        return stakers.values();
    }

    /**
     * @notice Sort stakers by amount
     */
    function _getSortedStakers() internal view returns (address[] memory) {
        address[] memory _stakers = _getStakers();

        for (uint256 i = 0; i < _stakers.length; i++) {
            uint256 amount0 = stakersAmount[_stakers[i]];
            uint256 maxIndex = i;

            for (uint256 j = i + 1; j < _stakers.length; j++) {
                uint256 amount1 = stakersAmount[_stakers[j]];
                if (amount1 > amount0) {
                    maxIndex = j;
                }

                address temp = _stakers[i];
                _stakers[i] = _stakers[maxIndex];
                _stakers[maxIndex] = temp;
            }
        }

        return _stakers;
    }
}
