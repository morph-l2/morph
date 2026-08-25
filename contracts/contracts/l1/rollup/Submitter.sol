// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {IRollupSubmitterView, ISubmitter} from "./ISubmitter.sol";

/// @title Submitter
/// @notice Registration, stake, exit, and slashing for batch submitters.
contract Submitter is ISubmitter, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    address public rollupContract;
    uint256 public minimumStake;
    uint256 public override challengeDeposit;
    uint256 public rewardPercentage;
    uint256 public slashRemaining;

    mapping(address submitter => bool isRegistered) public registered;
    mapping(address submitter => uint256 amount) public stakeOf;
    mapping(address submitter => bool isWithdrawing) public withdrawing;
    mapping(address submitter => uint256 batchIndex) public withdrawalBatchIndex;

    event SubmitterAdded(address indexed submitter);
    event SubmitterRemoved(address indexed submitter);
    event MinimumStakeUpdated(uint256 oldValue, uint256 newValue);
    event ChallengeDepositUpdated(uint256 oldValue, uint256 newValue);
    event RewardPercentageUpdated(uint256 oldValue, uint256 newValue);
    event Staked(address indexed submitter, uint256 amount, uint256 total);
    event WithdrawalRequested(address indexed submitter, uint256 amount);
    event WithdrawalClaimed(address indexed submitter, address indexed receiver, uint256 amount);
    event Slashed(address indexed submitter, uint256 amount, uint256 reward);
    event SlashRemainingClaimed(address indexed receiver, uint256 amount);

    modifier onlyRollup() {
        require(_msgSender() == rollupContract, "only rollup contract");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner_,
        address rollupContract_,
        uint256 minimumStake_,
        uint256 challengeDeposit_,
        uint256 rewardPercentage_
    ) external initializer {
        require(owner_ != address(0), "invalid owner");
        require(rollupContract_ != address(0) && rollupContract_.code.length > 0, "invalid rollup contract");
        require(minimumStake_ > 0, "invalid minimum stake");
        require(challengeDeposit_ > 0, "invalid challenge deposit");
        require(rewardPercentage_ > 0 && rewardPercentage_ <= 100, "invalid reward percentage");

        __Ownable_init();
        __ReentrancyGuard_init();
        _transferOwnership(owner_);

        rollupContract = rollupContract_;
        minimumStake = minimumStake_;
        challengeDeposit = challengeDeposit_;
        rewardPercentage = rewardPercentage_;

        emit MinimumStakeUpdated(0, minimumStake_);
        emit ChallengeDepositUpdated(0, challengeDeposit_);
        emit RewardPercentageUpdated(0, rewardPercentage_);
    }

    function isActive(address submitter) public view override returns (bool) {
        return registered[submitter] && !withdrawing[submitter] && stakeOf[submitter] >= minimumStake;
    }

    function addSubmitter(address submitter) external onlyOwner {
        require(
            submitter != address(0) && !registered[submitter] && !withdrawing[submitter] && stakeOf[submitter] == 0,
            "invalid submitter"
        );
        registered[submitter] = true;
        emit SubmitterAdded(submitter);
    }

    function removeSubmitter(address submitter) external onlyOwner {
        require(registered[submitter] && !withdrawing[submitter], "invalid submitter");
        registered[submitter] = false;
        uint256 amount = stakeOf[submitter];
        if (amount > 0) {
            _startWithdrawal(submitter, amount);
        }
        emit SubmitterRemoved(submitter);
    }

    function setMinimumStake(uint256 value) external onlyOwner {
        require(value > 0 && value != minimumStake, "invalid minimum stake");
        emit MinimumStakeUpdated(minimumStake, value);
        minimumStake = value;
    }

    function updateChallengeDeposit(uint256 value) external onlyOwner {
        require(value > 0 && value != challengeDeposit, "invalid challenge deposit");
        emit ChallengeDepositUpdated(challengeDeposit, value);
        challengeDeposit = value;
    }

    function updateRewardPercentage(uint256 value) external onlyOwner {
        require(value > 0 && value <= 100 && value != rewardPercentage, "invalid reward percentage");
        emit RewardPercentageUpdated(rewardPercentage, value);
        rewardPercentage = value;
    }

    /// @notice Credits the sent ETH to a registered submitter. The payer does not need to be the submitter.
    function stake(address submitter) external payable nonReentrant {
        require(registered[submitter] && !withdrawing[submitter], "not stakeable");
        uint256 total = stakeOf[submitter] + msg.value;
        require(msg.value > 0 && total >= minimumStake, "below minimum stake");
        stakeOf[submitter] = total;
        emit Staked(submitter, msg.value, total);
    }

    function withdraw() external {
        address submitter = _msgSender();
        require(registered[submitter] && !withdrawing[submitter] && stakeOf[submitter] > 0, "not withdrawable");

        // Exiting always revokes the owner's previous authorization. Re-entry requires addSubmitter.
        registered[submitter] = false;
        _startWithdrawal(submitter, stakeOf[submitter]);
    }

    function claimWithdrawal(address receiver) external nonReentrant {
        address submitter = _msgSender();
        require(receiver != address(0), "invalid receiver");
        require(withdrawing[submitter], "not withdrawing");
        require(
            IRollupSubmitterView(rollupContract).lastFinalizedBatchIndex() >= withdrawalBatchIndex[submitter],
            "withdrawal batch not finalized"
        );

        uint256 amount = stakeOf[submitter];
        stakeOf[submitter] = 0;
        withdrawing[submitter] = false;
        delete withdrawalBatchIndex[submitter];
        _transfer(receiver, amount);
        emit WithdrawalClaimed(submitter, receiver, amount);
    }

    function slash(address submitter) external override onlyRollup nonReentrant returns (uint256 reward) {
        if (submitter == address(0)) return 0;

        uint256 amount = stakeOf[submitter];
        registered[submitter] = false;
        stakeOf[submitter] = 0;
        withdrawing[submitter] = false;
        delete withdrawalBatchIndex[submitter];

        reward = (amount * rewardPercentage) / 100;
        slashRemaining += amount - reward;
        _transfer(rollupContract, reward);
        emit Slashed(submitter, amount, reward);
    }

    function claimSlashRemaining(address receiver) external onlyOwner nonReentrant {
        require(receiver != address(0), "invalid receiver");
        uint256 amount = slashRemaining;
        require(amount > 0, "no slash remaining");
        slashRemaining = 0;
        _transfer(receiver, amount);
        emit SlashRemainingClaimed(receiver, amount);
    }

    function _startWithdrawal(address submitter, uint256 amount) internal {
        withdrawing[submitter] = true;
        withdrawalBatchIndex[submitter] = IRollupSubmitterView(rollupContract).lastCommittedBatchIndex();
        emit WithdrawalRequested(submitter, amount);
    }

    function _transfer(address receiver, uint256 amount) internal {
        if (amount == 0) return;
        (bool success,) = receiver.call{value: amount}("");
        require(success, "ETH transfer failed");
    }

    uint256[49] private __gap;
}
