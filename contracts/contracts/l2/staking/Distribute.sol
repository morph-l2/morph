// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IDistribute} from "./IDistribute.sol";
import {IMorphToken} from "../system/IMorphToken.sol";

contract Distribute is IDistribute, OwnableUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    /*************
     * Constants *
     *************/

    /// @notice MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;

    /// @notice l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// @notice record contract address
    address public immutable RECORD_CONTRACT;

    /*************
     * Variables *
     *************/

    /// @notice total minted epoch
    uint256 private mintedEpochCount;

    /// @notice distribution info, delete after all claimed
    mapping(address delegatee => mapping(uint256 epochIndex => Distribution)) private distributions;

    /// oldest distribution
    mapping(address delegatee => uint256 epochIndex) private oldestDistribution;

    /// @notice delegatee's unclaimed commission
    mapping(address delegatee => uint256 amount) private commissions;

    /// @notice delegator's unclaimed reward
    mapping(address delegator => Unclaimed) private unclaimed;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Ensures that the caller message from l2 staking contract.
    modifier onlyL2StakingContract() {
        require(_msgSender() == L2_STAKING_CONTRACT, "only l2 staking contract allowed");
        _;
    }

    /// @notice Ensures that the caller message from record contract.
    modifier onlyRecordContract() {
        require(_msgSender() == RECORD_CONTRACT, "only record contract allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        RECORD_CONTRACT = Predeploys.RECORD;
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    /// @param _owner owner
    function initialize(address _owner) public initializer {
        require(_owner != address(0), "invalid owner address");
        _transferOwnership(_owner);
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @dev notify delegation
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param effectiveEpoch    delegation effective epoch
    /// @param amount            delegator total amount, not increment
    /// @param totalAmount       delegatee total amount
    /// @param newDelegation     first delegate or additional delegate
    function notifyDelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 amount,
        uint256 totalAmount,
        bool newDelegation
    ) public onlyL2StakingContract {
        // update distribution info
        distributions[delegatee][effectiveEpoch].delegationAmount = totalAmount;
        distributions[delegatee][effectiveEpoch].delegators.add(delegator);
        distributions[delegatee][effectiveEpoch].amounts[delegator] = amount;

        // update unclaimed info
        if (newDelegation) {
            unclaimed[delegator].delegatees.add(delegatee);
            unclaimed[delegator].unclaimedStart[delegatee] = effectiveEpoch;
        }
    }

    /// @dev notify unDelegation
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param effectiveEpoch    delegation effective epoch
    /// @param totalAmount       delegatee total amount
    function notifyUndelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 totalAmount
    ) public onlyL2StakingContract {
        // update distribution info
        distributions[delegatee][effectiveEpoch].delegationAmount = totalAmount;

        // not start reward yet, or delegate and undelegation within the same epoch, remove unclaim info
        if (effectiveEpoch == 0 || unclaimed[delegator].unclaimedStart[delegatee] == effectiveEpoch) {
            // update distribution info
            distributions[delegatee][effectiveEpoch].delegators.remove(delegator);
            delete distributions[delegatee][effectiveEpoch].amounts[delegator];

            // update unclaimed info
            unclaimed[delegator].delegatees.remove(delegatee);
            delete unclaimed[delegator].undelegated[delegatee];
            delete unclaimed[delegator].unclaimedStart[delegatee];
            delete unclaimed[delegator].unclaimedEnd[delegatee];
            return;
        }

        // update unclaimed info
        unclaimed[delegator].undelegated[delegatee] = true;
        unclaimed[delegator].unclaimedEnd[delegatee] = effectiveEpoch - 1;
    }

    /// @dev update epoch reward
    /// @param epochIndex         epoch index
    /// @param sequencers         sequencers
    /// @param delegatorRewards   sequencer's delegator reward amount
    /// @param commissionsAmount  sequencers commission amount
    function updateEpochReward(
        uint256 epochIndex,
        address[] calldata sequencers,
        uint256[] calldata delegatorRewards,
        uint256[] calldata commissionsAmount
    ) external onlyRecordContract {
        mintedEpochCount++;
        require(mintedEpochCount - 1 == epochIndex, "invalid epoch index");
        require(
            delegatorRewards.length == sequencers.length && commissionsAmount.length == sequencers.length,
            "invalid data length"
        );

        for (uint256 i = 0; i < sequencers.length; i++) {
            distributions[sequencers[i]][epochIndex].delegatorRewardAmount = delegatorRewards[i];
            if (distributions[sequencers[i]][epochIndex].delegationAmount == 0 && epochIndex > 0) {
                distributions[sequencers[i]][epochIndex].delegationAmount = distributions[sequencers[i]][epochIndex - 1]
                    .delegationAmount;
            }
            commissions[sequencers[i]] += commissionsAmount[i];
        }
    }

    /// @dev clean distributions
    /// @param delegatee         delegatee address
    /// @param targetEpochIndex  the epoch index to delete up to
    ///
    /// check off-chain that all rewards have been claimed
    function cleanDistributions(address delegatee, uint256 targetEpochIndex) external onlyOwner {
        uint256 i = oldestDistribution[delegatee];
        for (; i <= targetEpochIndex; i++) {
            delete distributions[delegatee][i];
        }
        oldestDistribution[delegatee] = i;
    }

    /// @dev claim delegation reward of a delegatee.
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    ///
    ///  If targetEpochIndex is zero, claim up to latest mint epoch,
    ///  otherwise it must be greater than the last claimed epoch index.
    function claim(address delegatee, address delegator, uint256 targetEpochIndex) external onlyL2StakingContract {
        require(mintedEpochCount != 0, "not minted yet");
        uint256 endEpochIndex = (targetEpochIndex == 0 || targetEpochIndex > mintedEpochCount - 1)
            ? mintedEpochCount - 1
            : targetEpochIndex;
        uint256 reward = _claim(delegatee, delegator, endEpochIndex);
        if (reward > 0) {
            _transfer(delegator, reward);
        }
    }

    /// @dev claim delegation reward of all sequencers.
    /// @param delegator         delegator address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    ///
    ///  If targetEpochIndex is zero, claim up to latest mint epoch,
    ///  otherwise it must be greater than the last claimed epoch index.
    function claimAll(address delegator, uint256 targetEpochIndex) external onlyL2StakingContract {
        require(mintedEpochCount != 0, "not minted yet");
        uint256 endEpochIndex = (targetEpochIndex == 0 || targetEpochIndex > mintedEpochCount - 1)
            ? mintedEpochCount - 1
            : targetEpochIndex;
        uint256 reward;
        for (uint256 i = 0; i < unclaimed[delegator].delegatees.length(); i++) {
            address delegatee = unclaimed[delegator].delegatees.at(i);
            if (
                unclaimed[delegator].delegatees.contains(delegatee) &&
                unclaimed[delegator].unclaimedStart[delegatee] <= endEpochIndex
            ) {
                reward += _claim(delegatee, delegator, endEpochIndex);
            }
        }
        if (reward > 0) {
            _transfer(delegator, reward);
        }
    }

    /// @dev claim commission reward
    /// @param delegatee         delegatee address
    function claimCommission(address delegatee) external onlyL2StakingContract {
        require(commissions[delegatee] > 0, "no commission to claim");

        uint256 amount = commissions[delegatee];
        delete commissions[delegatee];
        _transfer(delegatee, amount);

        emit CommissionClaimed(delegatee, amount);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice query oldest distribution
    /// @param delegatee     delegatee address
    function queryOldestDistribution(address delegatee) external view returns (uint256 epochIndex) {
        return oldestDistribution[delegatee];
    }

    /// @notice query whether all rewards have been claimed for a delegatee
    /// @param delegator     delegatee address
    /// @param delegatee     delegatee address
    function isRewardClaimed(address delegator, address delegatee) external view returns (bool claimed) {
        return !unclaimed[delegator].delegatees.contains(delegatee);
    }

    /// @notice query all unclaimed commission of a delegatee
    /// @param delegatee     delegatee address
    function queryUnclaimedCommission(address delegatee) external view returns (uint256 amount) {
        return commissions[delegatee];
    }

    /// @notice query unclaimed morph reward on a delegatee
    /// @param delegatee     delegatee address
    /// @param delegator     delegatee address
    function queryUnclaimed(address delegatee, address delegator) external view returns (uint256 reward) {
        require(unclaimed[delegator].delegatees.length() != 0, "invalid delegator or no remaining reward");
        require(unclaimed[delegator].delegatees.contains(delegatee), "no remaining reward of the delegatee");
        uint256 delegateeAmount;
        uint256 delegatorAmount;
        uint256 start = unclaimed[delegator].unclaimedStart[delegatee];
        for (uint256 i = start; i < mintedEpochCount; i++) {
            if (distributions[delegatee][i].amounts[delegator] > 0) {
                delegatorAmount = distributions[delegatee][i].amounts[delegator];
            }
            if (distributions[delegatee][i].delegationAmount > 0) {
                delegateeAmount = distributions[delegatee][i].delegationAmount;
            }
            reward += (distributions[delegatee][i].delegatorRewardAmount * delegatorAmount) / delegateeAmount;
            if (unclaimed[delegator].undelegated[delegatee] && unclaimed[delegator].unclaimedEnd[delegatee] == i) {
                break;
            }
        }
    }

    /// @notice query all unclaimed morph reward
    /// @param delegator     delegatee address
    function queryAllUnclaimed(
        address delegator
    ) external view returns (address[] memory delegatees, uint256[] memory rewards) {
        uint256 length = unclaimed[delegator].delegatees.length();
        require(length != 0, "invalid delegator or no remaining reward");
        delegatees = new address[](length);
        rewards = new uint256[](length);
        for (uint256 j = 0; j < unclaimed[delegator].delegatees.length(); j++) {
            address delegatee = unclaimed[delegator].delegatees.at(j);
            uint256 reward;
            uint256 delegateeAmount;
            uint256 delegatorAmount;
            uint256 start = unclaimed[delegator].unclaimedStart[delegatee];
            for (uint256 i = start; i < mintedEpochCount; i++) {
                if (distributions[delegatee][i].amounts[delegator] > 0) {
                    delegatorAmount = distributions[delegatee][i].amounts[delegator];
                }
                if (distributions[delegatee][i].delegationAmount > 0) {
                    delegateeAmount = distributions[delegatee][i].delegationAmount;
                }
                reward += (distributions[delegatee][i].delegatorRewardAmount * delegatorAmount) / delegateeAmount;
                if (unclaimed[delegator].undelegated[delegatee] && unclaimed[delegator].unclaimedEnd[delegatee] == i) {
                    break;
                }
            }
            delegatees[j] = delegatee;
            rewards[j] = reward;
        }
    }

    /// @notice query all unclaimed morph reward epochs info
    /// @param delegator     delegatee address
    function queryAllUnclaimedEpochs(
        address delegator
    ) external view returns (address[] memory, bool[] memory, uint256[] memory, uint256[] memory) {
        uint256 length = unclaimed[delegator].delegatees.length();
        address[] memory delegatees = new address[](length);
        bool[] memory undelegated = new bool[](length);
        uint256[] memory unclaimedStart = new uint256[](length);
        uint256[] memory unclaimedEnd = new uint256[](length);
        for (uint256 i = 0; i < length; i++) {
            delegatees[i] = unclaimed[delegator].delegatees.at(i);
            undelegated[i] = unclaimed[delegator].undelegated[delegatees[i]];
            unclaimedStart[i] = unclaimed[delegator].unclaimedStart[delegatees[i]];
            unclaimedEnd[i] = unclaimed[delegator].unclaimedEnd[delegatees[i]];
        }
        return (delegatees, undelegated, unclaimedStart, unclaimedEnd);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice transfer morph token
    function _transfer(address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(address(this));
        IMorphToken(MORPH_TOKEN_CONTRACT).transfer(_to, _amount);
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(address(this));
        require(_amount > 0 && balanceBefore - balanceAfter == _amount, "morph token transfer failed");
    }

    /// @notice claim delegator morph reward
    function _claim(address delegatee, address delegator, uint256 endEpochIndex) internal returns (uint256 reward) {
        require(unclaimed[delegator].delegatees.contains(delegatee), "no remaining reward");
        require(unclaimed[delegator].unclaimedStart[delegatee] <= endEpochIndex, "all reward claimed");

        uint256 delegateeAmount;
        uint256 delegatorAmount;
        for (uint256 i = unclaimed[delegator].unclaimedStart[delegatee]; i <= endEpochIndex; i++) {
            if (distributions[delegatee][i].amounts[delegator] > 0) {
                delegatorAmount = distributions[delegatee][i].amounts[delegator];
            }
            if (distributions[delegatee][i].delegationAmount > 0) {
                delegateeAmount = distributions[delegatee][i].delegationAmount;
            }
            reward += (distributions[delegatee][i].delegatorRewardAmount * delegatorAmount) / delegateeAmount;

            // if undelegated, remove delegator unclaimed info after claimed all
            if (unclaimed[delegator].undelegated[delegatee] && unclaimed[delegator].unclaimedEnd[delegatee] == i) {
                unclaimed[delegator].delegatees.remove(delegatee);
                delete unclaimed[delegator].undelegated[delegatee];
                delete unclaimed[delegator].unclaimedStart[delegatee];
                delete unclaimed[delegator].unclaimedEnd[delegatee];
                break;
            }
        }
        unclaimed[delegator].unclaimedStart[delegatee] = endEpochIndex + 1;
        if (distributions[delegatee][endEpochIndex + 1].amounts[delegator] == 0) {
            distributions[delegatee][endEpochIndex + 1].amounts[delegator] = delegatorAmount;
        }

        emit RewardClaimed(delegator, delegatee, endEpochIndex, reward);
    }
}
