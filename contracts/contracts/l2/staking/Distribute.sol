// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {IRecord} from "./IRecord.sol";
import {IDistribute} from "./IDistribute.sol";
import {IMorphToken} from "../system/IMorphToken.sol";

contract Distribute is IDistribute, OwnableUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    /*************
     * Constants *
     *************/

    /// @notice reward epoch, seconds of one day (3600 * 24)
    uint256 private constant REWARD_EPOCH = 86400;

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
    mapping(address delegatee => mapping(uint256 epochIndex => Distribution))
        private distributions;

    /// @notice delegatee's unclaimed commission
    mapping(address delegatee => uint256 epochIndex)
        public
        override unclaimedCommission;

    /// @notice delegator's unclaimed reward
    mapping(address delegator => Unclaimed) private unclaimed;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Ensures that the caller message from l2 staking contract.
    modifier onlyL2StakingContract() {
        require(
            _msgSender() == L2_STAKING_CONTRACT,
            "only l2 staking contract allowed"
        );
        _;
    }

    /// @notice Ensures that the caller message from record contract.
    modifier onlyRecordContract() {
        require(
            _msgSender() == RECORD_CONTRACT,
            "only record contract allowed"
        );

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
    }

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    function initialize() public initializer {
        __Ownable_init();
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
    /// @param remainsNumber     delegator number
    /// @param newDelegation     first delegate or additional delegate
    function notifyDelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 amount,
        uint256 totalAmount,
        uint256 remainsNumber,
        bool newDelegation
    ) public onlyL2StakingContract {
        // update distribution info
        distributions[delegatee][effectiveEpoch].delegationAmount = totalAmount;
        distributions[delegatee][effectiveEpoch].remainsNumber = remainsNumber;
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
    /// @param remainsNumber     delegator number
    function notifyUndelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 totalAmount,
        uint256 remainsNumber
    ) public onlyL2StakingContract {
        // update distribution info
        distributions[delegatee][effectiveEpoch].delegationAmount = totalAmount;
        distributions[delegatee][effectiveEpoch].remainsNumber = remainsNumber;

        // not start reward yet, or delegate and undelegation within the same epoch, remove unclaim info
        if (
            effectiveEpoch == 0 ||
            unclaimed[delegator].unclaimedStart[delegatee] == effectiveEpoch
        ) {
            // update distribution info
            distributions[delegatee][effectiveEpoch].delegators.remove(
                delegator
            );
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
    /// @param epochIndex        epoch index
    /// @param sequencers        sequencers
    /// @param delegatorRewards  sequencer's delegatorRewardAmount
    /// @param commissions       sequencers commission
    function updateEpochReward(
        uint256 epochIndex,
        address[] calldata sequencers,
        uint256[] calldata delegatorRewards,
        uint256[] calldata commissions
    ) external onlyRecordContract {
        mintedEpochCount++;
        require(mintedEpochCount - 1 == epochIndex, "invalid epoch index");

        require(
            delegatorRewards.length == sequencers.length &&
                commissions.length == sequencers.length,
            "invalid data length"
        );

        for (uint256 i = 0; i < sequencers.length; i++) {
            distributions[sequencers[i]][epochIndex]
                .delegatorRewardAmount = delegatorRewards[i];
            distributions[sequencers[i]][epochIndex]
                .commissionAmount = commissions[i];
        }
    }

    /// @dev claim delegation reward of a delegatee.
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    ///
    ///  If targetEpochIndex is zero, claim up to latest mint epoch,
    ///  otherwise it must be greater than the last claimed epoch index.
    function claim(
        address delegatee,
        address delegator,
        uint256 targetEpochIndex
    ) external onlyL2StakingContract {
        require(mintedEpochCount != 0, "not mint yet");
        uint256 endEpochIndex = targetEpochIndex;
        if (targetEpochIndex == 0 || targetEpochIndex > mintedEpochCount) {
            endEpochIndex = mintedEpochCount - 1;
        }
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
    function claimAll(
        address delegator,
        uint256 targetEpochIndex
    ) external onlyL2StakingContract {
        require(mintedEpochCount != 0, "not mint yet");
        uint256 endEpochIndex = targetEpochIndex;
        if (targetEpochIndex == 0 || targetEpochIndex > mintedEpochCount) {
            endEpochIndex = mintedEpochCount - 1;
        }
        uint256 reward;
        for (uint256 i = 0; i < unclaimed[delegator].delegatees.length(); i++) {
            address delegatee = unclaimed[delegator].delegatees.at(i);
            if (
                unclaimed[delegator].delegatees.contains(delegatee) &&
                unclaimed[delegator].unclaimedStart[delegatee] <= endEpochIndex
            ) {
                reward += _claim(delegatee, delegator, targetEpochIndex);
            }
        }
        if (reward > 0) {
            _transfer(delegator, reward);
        }
    }

    /// @dev claim commission reward
    /// @param delegatee         delegatee address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    function claimCommission(
        address delegatee,
        uint256 targetEpochIndex
    ) external onlyL2StakingContract {
        require(mintedEpochCount != 0, "not mint yet");
        uint256 end = targetEpochIndex;
        if (targetEpochIndex == 0 || targetEpochIndex > mintedEpochCount) {
            end = mintedEpochCount - 1;
        }
        require(
            unclaimedCommission[delegatee] <= end,
            "all commission claimed"
        );
        uint256 commission;
        for (uint256 i = unclaimedCommission[delegatee]; i <= end; i++) {
            commission += distributions[delegatee][i].commissionAmount;
        }
        if (commission > 0) {
            _transfer(delegatee, commission);
        }
        unclaimedCommission[delegatee] = end + 1;

        emit CommissionClaimed(delegatee, end, commission);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice query unclaimed morph reward on a delegatee
    /// @param delegatee     delegatee address
    /// @param delegator     delegatee address
    function queryUnclaimed(
        address delegatee,
        address delegator
    ) external view returns (uint256 reward) {
        uint256 totalAmount;
        uint256 delegatorAmount;
        uint start = unclaimed[delegator].unclaimedStart[delegatee];
        for (uint256 i = start; i < mintedEpochCount; i++) {
            if (distributions[delegatee][i].amounts[delegator] > 0) {
                delegatorAmount = distributions[delegatee][i].amounts[
                    delegator
                ];
            }
            if (distributions[delegatee][i].delegationAmount > 0) {
                totalAmount = distributions[delegatee][i].delegationAmount;
            }
            reward +=
                (distributions[delegatee][i].delegatorRewardAmount *
                    delegatorAmount) /
                totalAmount;
            if (
                unclaimed[delegator].undelegated[delegatee] &&
                unclaimed[delegator].unclaimedEnd[delegatee] == i
            ) {
                break;
            }
        }
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice transfer morph token
    function _transfer(address _to, uint256 _amount) internal {
        uint256 balanceBefore = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(
            _to
        );
        IMorphToken(MORPH_TOKEN_CONTRACT).transfer(_to, _amount);
        uint256 balanceAfter = IMorphToken(MORPH_TOKEN_CONTRACT).balanceOf(_to);
        require(
            _amount > 0 && balanceAfter - balanceBefore == _amount,
            "morph token transfer failed"
        );
    }

    /// @notice claim delegator morph reward
    function _claim(
        address delegatee,
        address delegator,
        uint256 endEpochIndex
    ) internal returns (uint256 reward) {
        require(
            unclaimed[delegator].delegatees.contains(delegatee),
            "no remaining reward"
        );
        require(
            unclaimed[delegator].unclaimedStart[delegatee] <= endEpochIndex,
            "all reward claimed"
        );

        for (
            uint256 i = unclaimed[delegator].unclaimedStart[delegatee];
            i <= endEpochIndex;
            i++
        ) {
            // compute reward
            reward +=
                (distributions[delegatee][i].delegatorRewardAmount *
                    distributions[delegatee][i].amounts[delegator]) /
                distributions[delegatee][i].delegationAmount;

            // if distribution is empty, update next distribution info
            if (
                !distributions[delegatee][i + 1].delegators.contains(delegator)
            ) {
                distributions[delegatee][i + 1].delegators.add(delegator);
                distributions[delegatee][i + 1].amounts[
                    delegator
                ] = distributions[delegatee][i].amounts[delegator];

                // if delegator info exist in next epoch, distribution must be updated
                if (distributions[delegatee][i + 1].delegationAmount == 0) {
                    distributions[delegatee][i + 1]
                        .delegationAmount = distributions[delegatee][i]
                        .delegationAmount;
                    distributions[delegatee][i + 1]
                        .remainsNumber = distributions[delegatee][i]
                        .remainsNumber;
                }
            }

            // update distribution info, delete if all claimed
            distributions[delegatee][i].remainsNumber--;
            if (distributions[delegatee][i].remainsNumber == 0) {
                delete distributions[delegatee][i];
            }

            // if undelegated, remove delegator unclaimed info after claimed all
            if (
                unclaimed[delegator].undelegated[delegatee] &&
                unclaimed[delegator].unclaimedEnd[delegatee] == i
            ) {
                unclaimed[delegator].delegatees.remove(delegatee);
                delete unclaimed[delegator].undelegated[delegatee];
                delete unclaimed[delegator].unclaimedStart[delegatee];
                delete unclaimed[delegator].unclaimedEnd[delegatee];
                break;
            }
            unclaimed[delegator].unclaimedStart[delegatee]++;
        }

        emit RewardClaimed(delegator, delegatee, endEpochIndex, reward);
    }
}
