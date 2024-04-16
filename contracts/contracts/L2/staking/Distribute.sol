// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../IMorphToken.sol";
import {IDistribute} from "./IDistribute.sol";

interface IRecords {
    // return epoch index start and end
    function epochInfo(uint256 index) external returns (uint256, uint256);

    function sequencerEpochRatio(
        uint256 epochIndex,
        address sequencer
    ) external returns (uint256);
}

contract Distribute is IDistribute, Initializable, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    // the maximum value of the epoch index recorded after mint execution.
    uint256 private _latestMintedEpochIndex;
    address private _morphToken;
    address private _record;
    address private _stake;
    // delegator => [sequencer]
    mapping(address => EnumerableSet.AddressSet) private _vestIn;
    // mapping(sequencer => mapping(epochIndex => Distribution));
    mapping(address => mapping(uint256 => Distribution)) private _collect;

    // mapping(sequencer => mapping(delegator => DelegatorEpochRecord));
    mapping(address => mapping(address => DelegatorEpochRecord))
    private _epochRecord;

    // epoch index => reward
    mapping(uint256 => uint256) private _rewards;

    // The start time of each day and the corresponding block number
    // block time => block number
    Uint256Set private _blockInfo;

    /**
     * @notice Ensures that the caller message from record contract.
     */
    modifier onlyRecord() {
        require(msg.sender == _record, "only record contract can call");
        _;
    }

    modifier onlyStake() {
        require(msg.sender == _stake, "only stake contract can call");
        _;
    }

    /**
     * @dev See {IDistribute-initialize}.
     */
    function initialize(
        address morphToken_,
        address record_,
        address stake_
    ) public initializer {
        require(
            morphToken_ != address(0),
            "invalid morph token contract address"
        );
        require(record_ != address(0), "invalid record contract address");
        require(stake_ != address(0), "invalid stake contract address");
        _morphToken = morphToken_;
        _record = record_;
        _stake = stake_;
    }

    /**
     * @dev See {IDistribute-initialize}.
     */
    function notify(uint256 blockTime, uint256 blockNumber) public onlyRecord {
        // todo blockTime
        require(
            blockTime <= block.timestamp,
            "blockTime must be smaller than or equal to the current block time"
        );
        require(
            blockNumber <= block.number,
            "blockNumber must be smaller than or equal to the current block number"
        );
        _blockInfo.index.add(blockTime);
        _blockInfo.value[blockTime] = blockNumber;
    }

    /**
     * @dev See {IDistribute-notifyUnDelegate}.
     */
    function notifyUnDelegate(
        address sequencer,
        address account,
        uint256 deadlineClaimEpochIndex
    ) public onlyStake {
        require(sequencer != address(0), "invalid sequencer address");
        require(account != address(0), "invalid account address");

        if (
            _epochRecord[sequencer][account].claimed >= deadlineClaimEpochIndex
        ) {
            revert(
                "deadline claim epoch index must be granter than claimed epoch index"
            );
        }

        _epochRecord[sequencer][account].deadline = deadlineClaimEpochIndex;

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
    ) public onlyStake {
        require(sequencer != address(0), "invalid sequencer address");
        require(account != address(0), "invalid account address");

        (uint256 startNumber, uint256 endNumber) = IRecords(_record).epochInfo(
            epochIndex
        );

        if (blockNumber < startNumber || blockNumber > endNumber) {
            revert("invalid args");
        }

        // the epoch index that actually took effect.
        epochIndex += 1;
        _epochRecord[sequencer][account].begin = epochIndex;
        _vestIn[account].add(sequencer);

        Distribution storage dt = _collect[sequencer][epochIndex];

        if (!_collect[sequencer][epochIndex].valid) {
            // Iterate over epoch index to find the nearest valid value
            for (uint i = epochIndex - 1; i > 0; i--) {
                if (_collect[sequencer][i].valid) {
                    dt.totalAmount =
                        _collect[sequencer][i].totalAmount +
                        amount;
                    for (
                        uint256 j = 0;
                        j < _collect[sequencer][i].amounts.index.length();
                        j++
                    ) {
                        address delegator = _collect[sequencer][i]
                            .amounts
                            .index
                            .at(j);
                        uint256 delegateAmount = _collect[sequencer][i]
                            .amounts
                            .value[delegator];

                        dt.amounts.index.add(delegator);
                        dt.amounts.value[delegator] = delegateAmount;
                    }

                    if (
                        !_collect[sequencer][i].amounts.index.contains(account)
                    ) {
                        // when it doesn't exist
                        dt.remainNumber =
                            _collect[sequencer][i].amounts.index.length() +
                            1;

                        dt.amounts.index.add(account);
                        dt.amounts.value[account] = amount;
                    } else {
                        // when it exist
                        dt.remainNumber = _collect[sequencer][i]
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

        emit NotifyDelegate(sequencer, epochIndex, account, amount, blockNumber);
    }

    /**
     * @dev See {IDistribute-mint}.
     */
    function mint() public onlyRecord {
        (uint256 mintBegin, uint256 mintEnd) = IMorphToken(_morphToken).mint();

        uint256 internalDays = (mintEnd - mintBegin) / 86400;

        for (uint256 i = 0; i < internalDays; i++) {
            if (_blockInfo.index.length() <= internalDays) {
                revert("mapping block time to block number data not enable");
            }

            uint256 tm = mintBegin + (i * 86400);

            for (uint256 j = 0; j < _blockInfo.index.length(); j++) {
                uint256 beginTimeOfOneDay = _blockInfo.index.at(j);

                if (beginTimeOfOneDay >= tm && beginTimeOfOneDay < tm + 86400) {
                    uint256 rewardOfOneDay = IMorphToken(_morphToken).reward(
                        tm
                    );
                    uint256 nextBeginTimeOfOneDay = _blockInfo.index.at(j + 1);
                    uint256 beginBlockNumberOfOneDay = _blockInfo.value[
                                beginTimeOfOneDay
                        ];
                    uint256 endBlockNumberOfOneDay = _blockInfo.value[
                                nextBeginTimeOfOneDay
                        ] - 1;

                    uint256 totalBlockNumberOfOneDay = endBlockNumberOfOneDay -
                                beginBlockNumberOfOneDay +
                                1;
                    for (uint256 k = _latestMintedEpochIndex; ; k++) {
                        (
                            uint256 epochIndexBeginNumber,
                            uint256 epochIndexEndNumber
                        ) = IRecords(_record).epochInfo(k);

                        if (epochIndexEndNumber <= beginBlockNumberOfOneDay) {
                            continue;
                        }

                        if (
                            beginBlockNumberOfOneDay <= epochIndexBeginNumber &&
                            epochIndexEndNumber <= endBlockNumberOfOneDay
                        ) {
                            _rewards[k] =
                                (rewardOfOneDay *
                                    (epochIndexEndNumber -
                                    epochIndexBeginNumber +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            _latestMintedEpochIndex = k;
                            continue;
                        } else if (
                            beginBlockNumberOfOneDay > epochIndexBeginNumber &&
                            beginBlockNumberOfOneDay < epochIndexEndNumber
                        ) {
                            _rewards[k] +=
                                (rewardOfOneDay *
                                    (epochIndexEndNumber -
                                    beginBlockNumberOfOneDay +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            _latestMintedEpochIndex = k;
                            continue;
                        } else if (
                            epochIndexBeginNumber < endBlockNumberOfOneDay &&
                            epochIndexEndNumber > endBlockNumberOfOneDay
                        ) {
                            _rewards[k] +=
                                (rewardOfOneDay *
                                    (endBlockNumberOfOneDay -
                                    epochIndexBeginNumber +
                                        1)) /
                                totalBlockNumberOfOneDay;
                            _latestMintedEpochIndex = k;
                            continue;
                        }
                        break;
                    }

                    for (uint256 m = j; m >= 0; m--) {
                        uint256 timeIndex = _blockInfo.index.at(m);
                        _blockInfo.index.remove(timeIndex);
                        delete _blockInfo.value[timeIndex];
                    }
                }
            }
        }
    }

    /**
     * @dev See {IDistribute-latestMintedEpochIndex}.
     */
    function latestMintedEpochIndex() public view returns (uint256) {
        return _latestMintedEpochIndex;
    }

    /**
     * @dev See {IDistribute-claimedEpochIndex}.
     */
    function claimedEpochIndex(address sequencer, address account) public view returns (uint256) {
        return _epochRecord[sequencer][account].claimed;
    }

    /**
     * @dev See {IDistribute-claimAll}.
     */
    function claimAll(address account, uint256 targetEpochIndex) public {
        uint256 accountTotalReward = 0;
        for (uint256 i = 0; i < _vestIn[account].length(); i++) {
            accountTotalReward += _claim(_vestIn[account].at(i), account, targetEpochIndex);
        }
        IMorphToken(_morphToken).transfer(account, accountTotalReward);

        emit ClaimAll(address(this), account, accountTotalReward);
    }

    /**
     * @dev See {IDistribute-claim}.
     */
    function claim(address sequencer, address account, uint256 targetEpochIndex) public {
        if (!_vestIn[account].contains(sequencer)) {
            revert("not delegate to the sequencer");
        }

        uint256 accountReward = _claim(sequencer, account, targetEpochIndex);
        IMorphToken(_morphToken).transfer(account, accountReward);

        emit Claim(address(this), account, accountReward);
    }

    function _claim(
        address sequencer,
        address account,
        uint256 targetEpochIndex
    ) internal returns (uint256) {
        // determine the epoch index of the end claim
        uint256 endClaimEpochIndex = _latestMintedEpochIndex;

        (uint256 startNumber, uint256 endNumber) = IRecords(_record).epochInfo(
            _latestMintedEpochIndex
        );
        // determine whether the epoch index starts in one day and ends in another
        if (startNumber / 86400 != endNumber / 86400) {
            endClaimEpochIndex = _latestMintedEpochIndex - 1;
        }

        uint256 accountDeadlineEpochIndex = _epochRecord[sequencer][account]
            .deadline;
        if (
            accountDeadlineEpochIndex != 0 &&
            endClaimEpochIndex > accountDeadlineEpochIndex
        ) {
            endClaimEpochIndex = accountDeadlineEpochIndex;
        }

        if (targetEpochIndex != 0) {
            if (targetEpochIndex < _epochRecord[sequencer][account].claimed + 1) {
                revert("claim epoch index must be greater than claimed epoch index");
            }

            if (targetEpochIndex < endClaimEpochIndex) {
                endClaimEpochIndex = targetEpochIndex;
            }
        }

        // determine the epoch index of the begin claim
        uint256 beginClaimEpochIndex = 0;
        uint256 claimedEpochIndex = _epochRecord[sequencer][account].claimed;
        if (claimedEpochIndex == 0) {
            beginClaimEpochIndex = _epochRecord[sequencer][account].begin;
        } else {
            beginClaimEpochIndex = claimedEpochIndex + 1;
        }

        uint256 accountReward = 0;
        uint256 validEpochIndex = 0;
        for (uint256 i = beginClaimEpochIndex; i <= endClaimEpochIndex; i++) {
            uint256 ratio = IRecords(_record).sequencerEpochRatio(i, sequencer);
            uint256 epochTotalReward = _rewards[i];
            //uint256 sequencerReward = epochTotalReward * ratio / 100;
            if (_collect[sequencer][i].valid) {
                accountReward +=
                    (epochTotalReward *
                    ratio *
                        _collect[sequencer][i].amounts.value[account]) /
                    (_collect[sequencer][i].totalAmount * 100);
                validEpochIndex = i;
            } else {
                for (uint j = i - 1; j > 0; j--) {
                    if (_collect[sequencer][j].valid) {
                        accountReward +=
                            (epochTotalReward *
                            ratio *
                                _collect[sequencer][j].amounts.value[account]) /
                            (_collect[sequencer][j].totalAmount * 100);
                        validEpochIndex = j;
                    }
                }
            }
        }

        if (endClaimEpochIndex != validEpochIndex) {
            // first copy
            Distribution storage dt = _collect[sequencer][endClaimEpochIndex];

            dt.totalAmount = _collect[sequencer][validEpochIndex].totalAmount;
            dt.remainNumber = _collect[sequencer][validEpochIndex]
                .amounts
                .index
                .length();

            for (
                uint256 j = 0;
                j < _collect[sequencer][validEpochIndex].amounts.index.length();
                j++
            ) {
                address delegator = _collect[sequencer][validEpochIndex]
                    .amounts
                    .index
                    .at(j);
                uint256 delegateAmount = _collect[sequencer][validEpochIndex]
                    .amounts
                    .value[delegator];

                dt.amounts.index.add(delegator);
                dt.amounts.value[delegator] = delegateAmount;
            }

            dt.valid = true;
        }

        _collect[sequencer][endClaimEpochIndex].remainNumber -= 1;

        _epochRecord[sequencer][account].claimed = endClaimEpochIndex;

        if (endClaimEpochIndex == _epochRecord[sequencer][account].deadline) {
            _vestIn[account].remove(sequencer);

            delete _epochRecord[sequencer][account];

            _collect[sequencer][endClaimEpochIndex].totalAmount -= _collect[
                        sequencer
                ][endClaimEpochIndex].amounts.value[account];


            _collect[sequencer][endClaimEpochIndex].amounts.index.remove(
                account
            );
            delete _collect[sequencer][endClaimEpochIndex].amounts.value[
            account
            ];
        }

        return accountReward;
    }
}
