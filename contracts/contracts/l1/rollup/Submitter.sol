// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {IRollupSubmitterView, ISubmitter} from "./ISubmitter.sol";

/// @title Submitter
/// @notice Registration, stake, exit, and slashing for centralized batch submitters.
contract Submitter is ISubmitter, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    uint256 internal constant LIABILITY_QUERY_GAS_LIMIT = 50_000;
    uint256 internal constant REFUND_GAS_LIMIT = 50_000;

    address public override rollupContract;
    uint256 public minimumStake;
    uint256 public override challengeDeposit;
    uint256 public rewardPercentage;
    uint256 public slashRemaining;

    mapping(address submitter => bool isRegistered) public registered;
    mapping(address submitter => uint256 amount) public stakeOf;
    mapping(address submitter => bool isExiting) public exiting;
    mapping(address submitter => uint256 amount) public claimable;

    event SubmitterAdded(address indexed submitter);
    event SubmitterRemoved(address indexed submitter);
    event MinimumStakeUpdated(uint256 oldValue, uint256 newValue);
    event ChallengeDepositUpdated(uint256 oldValue, uint256 newValue);
    event RewardPercentageUpdated(uint256 oldValue, uint256 newValue);
    event Staked(address indexed submitter, uint256 amount, uint256 total);
    event WithdrawalRequested(address indexed submitter, address indexed receiver, uint256 amount);
    event WithdrawalPaid(address indexed submitter, address indexed receiver, uint256 amount);
    event WithdrawalCredited(address indexed submitter, uint256 amount, uint256 totalCredit);
    event CreditClaimed(address indexed submitter, address indexed receiver, uint256 amount);
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
        return registered[submitter] && !exiting[submitter] && stakeOf[submitter] >= minimumStake;
    }

    function addSubmitter(address submitter) external onlyOwner {
        require(submitter != address(0) && !registered[submitter] && !exiting[submitter], "invalid submitter");
        registered[submitter] = true;
        emit SubmitterAdded(submitter);
    }

    function removeSubmitter(address submitter) external onlyOwner nonReentrant {
        require(registered[submitter] && !exiting[submitter], "invalid submitter");
        registered[submitter] = false;
        exiting[submitter] = true;
        emit SubmitterRemoved(submitter);
        emit WithdrawalRequested(submitter, submitter, stakeOf[submitter]);
        _settleIfReleased(submitter, submitter);
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

    function stake() external payable nonReentrant {
        address submitter = _msgSender();
        require(registered[submitter] && !exiting[submitter], "not stakeable");
        uint256 total = stakeOf[submitter] + msg.value;
        require(msg.value > 0 && total >= minimumStake, "below minimum stake");
        stakeOf[submitter] = total;
        emit Staked(submitter, msg.value, total);
    }

    function withdraw(address receiver) external nonReentrant {
        address submitter = _msgSender();
        require(receiver != address(0), "invalid receiver");
        require(registered[submitter] && !exiting[submitter] && stakeOf[submitter] > 0, "not withdrawable");

        // A completed exit always requires a fresh owner registration before staking again.
        registered[submitter] = false;
        exiting[submitter] = true;
        emit WithdrawalRequested(submitter, receiver, stakeOf[submitter]);
        _settleIfReleased(submitter, receiver);
    }

    function claimWithdrawal(address receiver) external nonReentrant {
        address submitter = _msgSender();
        require(receiver != address(0), "invalid receiver");
        require(exiting[submitter], "not exiting");
        require(_liabilityReleased(submitter), "pending batch liability");
        _settle(submitter, receiver);
    }

    function claimCredit(address receiver) external nonReentrant {
        address submitter = _msgSender();
        require(receiver != address(0), "invalid receiver");
        uint256 amount = claimable[submitter];
        require(amount > 0, "no credit");
        claimable[submitter] = 0;
        _transfer(receiver, amount);
        emit CreditClaimed(submitter, receiver, amount);
    }

    function slash(address submitter) external override onlyRollup nonReentrant returns (uint256 reward) {
        if (submitter == address(0)) return 0;

        uint256 amount = stakeOf[submitter];
        registered[submitter] = false;
        stakeOf[submitter] = 0;
        exiting[submitter] = true;

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

    function _settleIfReleased(address submitter, address receiver) internal {
        if (_liabilityReleased(submitter)) _settle(submitter, receiver);
    }

    function _settle(address submitter, address receiver) internal {
        uint256 amount = stakeOf[submitter];
        stakeOf[submitter] = 0;
        exiting[submitter] = false;

        if (amount == 0) {
            emit WithdrawalPaid(submitter, receiver, 0);
            return;
        }

        (bool success,) = receiver.call{value: amount, gas: REFUND_GAS_LIMIT}("");
        if (success) {
            emit WithdrawalPaid(submitter, receiver, amount);
        } else {
            claimable[submitter] += amount;
            emit WithdrawalCredited(submitter, amount, claimable[submitter]);
        }
    }

    function _liabilityReleased(address submitter) internal view returns (bool) {
        bytes memory input = abi.encodeCall(IRollupSubmitterView.pendingBatchCount, (submitter));
        address target = rollupContract;
        bool ok;
        uint256 size;
        uint256 count;
        assembly ("memory-safe") {
            ok := staticcall(LIABILITY_QUERY_GAS_LIMIT, target, add(input, 0x20), mload(input), 0, 0)
            size := returndatasize()
            if and(ok, eq(size, 0x20)) {
                returndatacopy(0, 0, 0x20)
                count := mload(0)
            }
        }
        return ok && size == 32 && count == 0;
    }

    function _transfer(address receiver, uint256 amount) internal {
        if (amount == 0) return;
        (bool success,) = receiver.call{value: amount}("");
        require(success, "ETH transfer failed");
    }

    uint256[50] private __gap;
}
