// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IMorphToken} from "../system/IMorphToken.sol";
import {IDistribute} from "./IDistribute.sol";
import {IRecord} from "./IRecord.sol";

contract Distribute is IDistribute, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;

    // MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;
    // record contract address
    address public immutable RECORD_CONTRACT;
    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    // the maximum value of the epoch index recorded after mint execution.
    uint256 public latestMintedEpochIndex;

    // delegator => []sequencer
    mapping(address => EnumerableSet.AddressSet) private vestIn;
    // mapping(sequencer => mapping(epochIndex => Distribution));
    mapping(address => mapping(uint256 => Distribution)) private collect;

    // mapping(sequencer => mapping(delegator => DelegatorEpochRecord));
    mapping(address => mapping(address => DelegatorEpochRecord))
        private epochRecord;

    // epoch index => reward
    mapping(uint256 => uint256) private rewards;

    // The start time of each day and the corresponding block number
    // block time => block number (incremental)
    TimeOrderedSet private blockInfo;

    /*********************** modifiers **************************/

    /**
     * @notice Ensures that the caller message from record contract.
     */
    modifier onlyRecordContract() {
        require(msg.sender == RECORD_CONTRACT, "only record contract allowed");
        _;
    }

    /**
     * @notice Ensures that the caller message from staking contract.
     */
    modifier onlyStakingContract() {
        require(
            msg.sender == L2_STAKING_CONTRACT,
            "only stake contract allowed"
        );
        _;
    }

    /*********************** Constructor **************************/

    /**
     * @notice constructor
     */
    constructor() {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        RECORD_CONTRACT = Predeploys.RECORD;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
    }

    /*********************** Init ****************************************/

    /**
     * @dev See {IDistribute-initialize}.
     */
    function initialize() public initializer {}

    /*********************** External Functions **************************/

    /**
     * @dev See {IDistribute-initialize}.
     */
    function notify(
        uint256 blockTime,
        uint256 blockNumber
    ) public onlyRecordContract {
        // todo blockTime
        require(
            blockTime <= block.timestamp,
            "blockTime must be smaller than or equal to the current block time"
        );
        require(
            blockNumber <= block.number,
            "blockNumber must be smaller than or equal to the current block number"
        );
        blockInfo.index.pushBack(bytes32(blockTime));
        blockInfo.value[blockTime] = blockNumber;
    }

    /**
     * @dev See {IDistribute-notifyUnDelegate}.
     */
    function notifyUnDelegate(
        address sequencer,
        address account,
        uint256 deadlineClaimEpochIndex
    ) public onlyStakingContract {
        require(sequencer != address(0), "invalid sequencer address");
        require(account != address(0), "invalid account address");

        if (
            epochRecord[sequencer][account].claimed >= deadlineClaimEpochIndex
        ) {
            revert(
                "deadline claim epoch index must be granter than claimed epoch index"
            );
        }

        epochRecord[sequencer][account].deadline = deadlineClaimEpochIndex;

        emit NotifyUnDelegate(sequencer, account, deadlineClaimEpochIndex);
    }

    /**
     * @dev See {IDistribute-notifyDelegate}.
     */
    function notifyDelegate(
        address sequencer,
        uint256 epochIndex,
        address account,
        uint256 amount,
        uint256 blockNumber
    ) public onlyStakingContract {
        require(sequencer != address(0), "invalid sequencer address");
        require(account != address(0), "invalid account address");

        (uint256 startNumber, uint256 endNumber) = IRecord(RECORD_CONTRACT)
            .epochInfo(epochIndex);

        if (blockNumber < startNumber || blockNumber > endNumber) {
            revert("invalid args");
        }

        // the epoch index that actually took effect.
        epochIndex += 1;
        epochRecord[sequencer][account].begin = epochIndex;
        vestIn[account].add(sequencer);

        Distribution storage dt = collect[sequencer][epochIndex];

        if (!collect[sequencer][epochIndex].valid) {
            // Iterate over epoch index to find the nearest valid value
            for (uint i = epochIndex - 1; i > 0; i--) {
                if (collect[sequencer][i].valid) {
                    dt.totalAmount = collect[sequencer][i].totalAmount + amount;
                    for (
                        uint256 j = 0;
                        j < collect[sequencer][i].amounts.index.length();
                        j++
                    ) {
                        address delegator = collect[sequencer][i]
                            .amounts
                            .index
                            .at(j);
                        uint256 delegateAmount = collect[sequencer][i]
                            .amounts
                            .value[delegator];

                        dt.amounts.index.add(delegator);
                        dt.amounts.value[delegator] = delegateAmount;
                    }

                    if (
                        !collect[sequencer][i].amounts.index.contains(account)
                    ) {
                        // when it doesn't exist
                        dt.remainNumber =
                            collect[sequencer][i].amounts.index.length() +
                            1;

                        dt.amounts.index.add(account);
                        dt.amounts.value[account] = amount;
                    } else {
                        // when it exist
                        dt.remainNumber = collect[sequencer][i]
                            .amounts
                            .index
                            .length();

                        dt.amounts.value[account] += amount;
                    }
                    dt.valid = true;
                }
            }

            if (!dt.valid) {
                // When none existed
                dt.totalAmount = amount;
                dt.remainNumber = 1;
                dt.amounts.index.add(account);
                dt.amounts.value[account] = amount;
                dt.valid = true;
            }
        } else {
            dt.totalAmount += amount;

            if (!dt.amounts.index.contains(account)) {
                // when it doesn't exist
                dt.remainNumber += 1;

                dt.amounts.index.add(account);
                dt.amounts.value[account] = amount;
            } else {
                // when it exist
                dt.amounts.value[account] += amount;
            }
        }

        emit NotifyDelegate(
            sequencer,
            epochIndex,
            account,
            amount,
            blockNumber
        );
    }

    /**
     * @dev See {IDistribute-mint}.
     */
    function mint() public onlyRecordContract {
        (uint256 mintBegin, uint256 mintEnd) = IMorphToken(MORPH_TOKEN_CONTRACT)
            .mint();

        uint256 internalDays = (mintEnd - mintBegin) / 86400;

        for (uint256 i = 0; i < internalDays; i++) {
            if (blockInfo.index.length() <= internalDays) {
                revert("mapping block time to block number data not enable");
            }

            uint256 tm = mintBegin + (i * 86400);

            uint256 usedIndex = 0;

            for (uint256 j = 0; j < blockInfo.index.length(); j++) {
                uint256 beginTimeOfOneDay = uint256(blockInfo.index.at(j));

                if (beginTimeOfOneDay >= tm && beginTimeOfOneDay < tm + 86400) {
                    usedIndex = j;

                    uint256 rewardOfOneDay = IMorphToken(MORPH_TOKEN_CONTRACT)
                        .reward(tm);
                    uint256 nextBeginTimeOfOneDay = uint256(
                        blockInfo.index.at(j + 1)
                    );
                    uint256 beginBlockNumberOfOneDay = blockInfo.value[
                        beginTimeOfOneDay
                    ];
                    uint256 endBlockNumberOfOneDay = blockInfo.value[
                        nextBeginTimeOfOneDay
                    ] - 1;

                    uint256 totalBlockNumberOfOneDay = endBlockNumberOfOneDay -
                        beginBlockNumberOfOneDay +
                        1;
                    for (uint256 k = latestMintedEpochIndex; ; k++) {
                        (
                            uint256 epochIndexBeginNumber,
                            uint256 epochIndexEndNumber
                        ) = IRecord(RECORD_CONTRACT).epochInfo(k);

                        if (epochIndexEndNumber <= beginBlockNumberOfOneDay) {
                            continue;
                        }

                        if (
                            beginBlockNumberOfOneDay <= epochIndexBeginNumber &&
                            epochIndexEndNumber <= endBlockNumberOfOneDay
                        ) {
                            rewards[k] =
                                (rewardOfOneDay *
                                    (epochIndexEndNumber -
                                        epochIndexBeginNumber +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            latestMintedEpochIndex = k;
                            continue;
                        } else if (
                            beginBlockNumberOfOneDay > epochIndexBeginNumber &&
                            beginBlockNumberOfOneDay < epochIndexEndNumber
                        ) {
                            rewards[k] +=
                                (rewardOfOneDay *
                                    (epochIndexEndNumber -
                                        beginBlockNumberOfOneDay +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            latestMintedEpochIndex = k;
                            continue;
                        } else if (
                            epochIndexBeginNumber < endBlockNumberOfOneDay &&
                            epochIndexEndNumber > endBlockNumberOfOneDay
                        ) {
                            rewards[k] +=
                                (rewardOfOneDay *
                                    (endBlockNumberOfOneDay -
                                        epochIndexBeginNumber +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            latestMintedEpochIndex = k;
                            continue;
                        }
                        break;
                    }
                }
            }
            for (uint256 m = 0; m <= usedIndex; m++) {
                bytes32 value = blockInfo.index.popFront();
                delete blockInfo.value[uint256(value)];
            }
        }
    }

    /**
     * @dev See {IDistribute-claimAll}.
     */
    function claimAll(address account, uint256 targetEpochIndex) public {
        uint256 accountTotalReward = 0;
        for (uint256 i = 0; i < vestIn[account].length(); i++) {
            accountTotalReward += _claim(
                vestIn[account].at(i),
                account,
                targetEpochIndex
            );
        }
        IMorphToken(MORPH_TOKEN_CONTRACT).transfer(account, accountTotalReward);

        emit ClaimAll(address(this), account, accountTotalReward);
    }

    /**
     * @dev See {IDistribute-claim}.
     */
    function claim(
        address sequencer,
        address account,
        uint256 targetEpochIndex
    ) public {
        if (!vestIn[account].contains(sequencer)) {
            revert("not delegate to the sequencer");
        }

        uint256 accountReward = _claim(sequencer, account, targetEpochIndex);
        IMorphToken(MORPH_TOKEN_CONTRACT).transfer(account, accountReward);

        emit Claim(address(this), account, accountReward);
    }

    /*********************** External View Functions **************************/

    /**
     * @dev See {IDistribute-claimedEpochIndex}.
     */
    function claimedEpochIndex(
        address sequencer,
        address account
    ) public view returns (uint256) {
        return epochRecord[sequencer][account].claimed;
    }

    /*********************** Internal Functions **************************/

    function _claim(
        address sequencer,
        address account,
        uint256 targetEpochIndex
    ) internal returns (uint256) {
        // determine the epoch index of the end claim
        uint256 endClaimEpochIndex = latestMintedEpochIndex;

        (uint256 startNumber, uint256 endNumber) = IRecord(RECORD_CONTRACT)
            .epochInfo(latestMintedEpochIndex);
        // determine whether the epoch index starts in one day and ends in another
        if (startNumber / 86400 != endNumber / 86400) {
            endClaimEpochIndex = latestMintedEpochIndex - 1;
        }

        uint256 accountDeadlineEpochIndex = epochRecord[sequencer][account]
            .deadline;
        if (
            accountDeadlineEpochIndex != 0 &&
            endClaimEpochIndex > accountDeadlineEpochIndex
        ) {
            endClaimEpochIndex = accountDeadlineEpochIndex;
        }

        if (targetEpochIndex != 0) {
            if (
                targetEpochIndex < epochRecord[sequencer][account].claimed + 1
            ) {
                revert(
                    "claim epoch index must be greater than claimed epoch index"
                );
            }

            if (targetEpochIndex < endClaimEpochIndex) {
                endClaimEpochIndex = targetEpochIndex;
            }
        }

        // determine the epoch index of the begin claim
        uint256 epochBegin = 0;
        uint256 epochClaimed = epochRecord[sequencer][account].claimed;
        if (epochClaimed == 0) {
            epochBegin = epochRecord[sequencer][account].begin;
        } else {
            epochBegin = epochClaimed + 1;
        }

        uint256 accountReward = 0;
        uint256 validEpochIndex = 0;
        for (uint256 i = epochBegin; i <= endClaimEpochIndex; i++) {
            uint256 ratio = IRecord(RECORD_CONTRACT).sequencerEpochRatio(
                i,
                sequencer
            );
            uint256 epochTotalReward = rewards[i];
            //uint256 sequencerReward = epochTotalReward * ratio / 100;
            if (collect[sequencer][i].valid) {
                accountReward +=
                    (epochTotalReward *
                        ratio *
                        collect[sequencer][i].amounts.value[account]) /
                    (collect[sequencer][i].totalAmount * 100);
                validEpochIndex = i;
            } else {
                for (uint j = i - 1; j > 0; j--) {
                    if (collect[sequencer][j].valid) {
                        accountReward +=
                            (epochTotalReward *
                                ratio *
                                collect[sequencer][j].amounts.value[account]) /
                            (collect[sequencer][j].totalAmount * 100);
                        validEpochIndex = j;
                    }
                }
            }
        }

        if (endClaimEpochIndex != validEpochIndex) {
            // first copy
            Distribution storage dt = collect[sequencer][endClaimEpochIndex];

            dt.totalAmount = collect[sequencer][validEpochIndex].totalAmount;
            dt.remainNumber = collect[sequencer][validEpochIndex]
                .amounts
                .index
                .length();

            for (
                uint256 j = 0;
                j < collect[sequencer][validEpochIndex].amounts.index.length();
                j++
            ) {
                address delegator = collect[sequencer][validEpochIndex]
                    .amounts
                    .index
                    .at(j);
                uint256 delegateAmount = collect[sequencer][validEpochIndex]
                    .amounts
                    .value[delegator];

                dt.amounts.index.add(delegator);
                dt.amounts.value[delegator] = delegateAmount;
            }

            dt.valid = true;
        }

        collect[sequencer][endClaimEpochIndex].remainNumber -= 1;
        epochRecord[sequencer][account].claimed = endClaimEpochIndex;

        if (endClaimEpochIndex == epochRecord[sequencer][account].deadline) {
            vestIn[account].remove(sequencer);

            delete epochRecord[sequencer][account];

            collect[sequencer][endClaimEpochIndex].totalAmount -= collect[
                sequencer
            ][endClaimEpochIndex].amounts.value[account];

            collect[sequencer][endClaimEpochIndex].amounts.index.remove(
                account
            );
            delete collect[sequencer][endClaimEpochIndex].amounts.value[
                account
            ];
        }

        return accountReward;
    }
}
