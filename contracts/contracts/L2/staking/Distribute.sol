// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../IMorphToken.sol";

interface IRecords {
    // return epoch index start and end
    function epochInfo(uint256 index) external returns (uint256, uint256);

    function sequencerEpochRatio(uint256 epochIndex, address sequencer) external returns (uint256);
}

contract Distribute is Initializable, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;


    struct Set {
        EnumerableSet.AddressSet index;
        mapping(address => uint256) value;
    }

    struct Uint256Set {
        EnumerableSet.UintSet index;
        mapping(uint256 => uint256) value;
    }

    struct Distribution {
        uint256 totalAmount;
        uint256 totalShare;
        uint256 remainNumber;
        // mapping(delegator => share)
        Set shares;
        // mapping(delegator => amount)
        Set amounts;
        bool valid;
    }

    uint256 private _usedEpochIndex;
    address private _morphToken;
    address private _record;
    address private _gov;
    // delegator => [sequencer]
    mapping(address => EnumerableSet.AddressSet) private _vestIn;
    //mapping(sequencer => mapping(epochIndex => Distribution)) private a;
    mapping(address => mapping(uint256 => Distribution)) private _collect;

    // todo mapping(sequencer => mapping(delegator => deadlineEpochIndex))
    mapping(address => mapping(address => uint256))  private _deadlineRecord;

    // mapping(sequencer => mapping(delegator => claimedEpochIndex)) private b;
    mapping(address => mapping(address => uint256)) private _claimedEpoch;
    // epoch index => reward
    mapping(uint256 => uint256) private _rewards;
    // The start time of each day and the corresponding block number
    // time => number
    Uint256Set  private _timeToBlockNumber;


    /**
     * @notice Ensures that the caller message from record contract.
     */
    modifier onlyRecord() {
        require(msg.sender == _record, "only record contract can call");
        _;
    }

    /**
     * @notice Ensures that the caller message from gov contract.
     */
    modifier onlyGov() {
        require(msg.sender == _gov, "only gov contract can call");
        _;
    }


    function initialize(address morphToken_, address record_, address gov_) initializer public {
        require(morphToken_ != address(0), "invalid morph token contract address");
        require(record_ != address(0), "invalid record contract address");
        require(gov_ != address(0), "invalid gov contract address");
        _morphToken = morphToken_;
        _record = record_;
        _gov = gov_;
    }

    function notify(uint256 blockTime, uint256 blockNumber) public onlyGov {
        require(blockTime <= block.timestamp, "blockTime must be smaller than or equal to the current block time");
        require(blockNumber <= block.number, "blockNumber must be smaller than or equal to the current block number");
        _timeToBlockNumber.index.add(blockTime);
        _timeToBlockNumber.value[blockTime] = blockNumber;
    }

    function mint() public onlyRecord {
        (uint256 mintBegin, uint256 mintEnd) = IMorphToken(_morphToken).mint();

        uint256 internalDays = (mintEnd - mintBegin) / 86400;

        for (uint256 i = 0; i < internalDays; i++) {
            //? if not data
            for (uint256 j = _timeToBlockNumber.index.length() - 1; j >= 0; j--) {
                uint256 tm = mintBegin + (i * 86400);
                if (_timeToBlockNumber.index.at(j) >= tm && _timeToBlockNumber.index.at(j) < tm + 86400) {

                    uint256 rewardOfOneDay = IMorphToken(_morphToken).reward(tm);

                    uint256 beginTimeOfOneDay = _timeToBlockNumber.index.at(j);
                    // ?
                    uint256 nextBeginTimeOfOneDay = _timeToBlockNumber.index.at(j + 1);
                    uint256 beginBlockNumberOfOneDay = _timeToBlockNumber.value[beginTimeOfOneDay];
                    uint256 endBlockNumberOfOneDay = _timeToBlockNumber.value[nextBeginTimeOfOneDay] - 1;

                    uint256 totalBlockNumberOfOneDay = endBlockNumberOfOneDay - beginBlockNumberOfOneDay + 1;
                    for (uint256 k = _usedEpochIndex;;k++) {
                        (uint256 epochIndexBeginNumber, uint256 epochIndexEndNumber) = IRecords(_record).epochInfo(k);

                        if (beginBlockNumberOfOneDay <= epochIndexBeginNumber && epochIndexEndNumber <= endBlockNumberOfOneDay) {
                            _rewards[k] = rewardOfOneDay * (epochIndexEndNumber - epochIndexBeginNumber + 1) / totalBlockNumberOfOneDay;
                            _usedEpochIndex = k;
                            continue;
                        }else if (beginBlockNumberOfOneDay > epochIndexBeginNumber && epochIndexEndNumber < endBlockNumberOfOneDay) {
                            _rewards[k] += rewardOfOneDay * (epochIndexEndNumber - beginBlockNumberOfOneDay + 1) / totalBlockNumberOfOneDay;
                            _usedEpochIndex = k;
                            continue;
                        }else if (beginBlockNumberOfOneDay < epochIndexBeginNumber && epochIndexEndNumber > endBlockNumberOfOneDay) {
                            _rewards[k] += rewardOfOneDay * (endBlockNumberOfOneDay - epochIndexBeginNumber + 1) / totalBlockNumberOfOneDay;
                            _usedEpochIndex = k;
                            continue;
                        }
                        break;
                    }
                }
            }
        }
    }

    function notify(uint256 epochIndex, uint256 startNumber, uint256 endNumber, address sequencer, address account, uint256 amount, uint256 blockNumber) public {

        _vestIn[account].add(sequencer);

        Distribution storage dt = _collect[sequencer][epochIndex];

        if (!_collect[sequencer][epochIndex].valid) {
            // Iterate over epoch index to find the nearest valid value
            for (uint i = epochIndex - 1; i > 0; i--) {
                if (_collect[sequencer][i].valid) {

                    (uint256 totalShare, uint256 newAccountShare) = _additiveDilution(startNumber, endNumber, blockNumber);

                    dt.totalAmount = _collect[sequencer][i].totalAmount + amount;
                    dt.totalShare = totalShare;

                    for (uint256 j = 0; j < _collect[sequencer][i].amounts.index.length(); j++) {
                        address delegator = _collect[sequencer][i].amounts.index.at(j);
                        uint256 delegateAmount = _collect[sequencer][i].amounts.value[delegator];

                        dt.shares.index.add(delegator);
                        dt.shares.value[delegator] = delegateAmount;

                        dt.amounts.index.add(delegator);
                        dt.amounts.value[delegator] = delegateAmount;
                    }

                    if (!_collect[sequencer][i].shares.index.contains(account)) {
                        // when it doesn't exist
                        dt.remainNumber = _collect[sequencer][i].amounts.index.length() + 1;

                        dt.shares.index.add(account);
                        dt.shares.value[account] = newAccountShare;

                        dt.amounts.index.add(account);
                        dt.amounts.value[account] = amount;

                    }else {
                        // when it exist
                        dt.remainNumber = _collect[sequencer][i].amounts.index.length();

                        dt.shares.value[account] = newAccountShare;

                        dt.amounts.value[account] += amount;
                    }
                    dt.valid = true;
                    return;
                }
            }

            // When none existed
            dt.totalAmount = amount;
            dt.totalShare = amount;
            dt.remainNumber = 1;
            dt.shares.index.add(account);
            dt.shares.value[account] = amount;
            dt.amounts.index.add(account);
            dt.amounts.value[account] = amount;
            dt.valid = true;
        } else {

            // todo
            (uint256 totalShare, uint256 newAccountShare) = _additiveDilution(startNumber, endNumber, blockNumber);

            dt.totalAmount += amount;
            dt.totalShare = totalShare;

            if (!dt.shares.index.contains(account)) {
                // when it doesn't exist
                dt.remainNumber += 1;

                dt.shares.index.add(account);
                // todo
                dt.shares.value[account] = newAccountShare;

                dt.amounts.index.add(account);
                dt.amounts.value[account] = amount;
            }else {
                // when it exist
                // todo
                dt.shares.value[account] = newAccountShare;

                dt.amounts.value[account] += amount;
            }
        }
    }

    // equity increase : additive dilution
    function _additiveDilution(uint256 startNumber, uint256 endNumber, uint256 blockNumber) internal returns (uint256, uint256) {
        // uint256 totalMolecule = _collect[sequencer][i].totalAmount * (endNumber - startNumber) * (_collect[sequencer][i].totalAmount + amount);
        //                    uint256 molecule = _collect[sequencer][i].totalAmount * amount * (endNumber - blockNumber);
        //                    uint256 denominator = ((blockNumber - startNumber) * amount + (endNumber - startNumber) * _collect[sequencer][i].totalAmount);
        // totalMolecule / denominator;
        return (0, 0);
    }

    function claimAll(address account) public {
        uint256 accountTotalReward = 0;
        for (uint256 i = 0; i < _vestIn[account].length(); i++) {
            accountTotalReward += _claim(_vestIn[account].at(i), account);
        }
        IMorphToken(_morphToken).transfer(account, accountTotalReward);
    }

    function claim(address sequencer, address account) public {
        uint256 accountReward = _claim(sequencer, account);
        IMorphToken(_morphToken).transfer(account, accountReward);
    }

    function _claim(address sequencer, address account) internal returns (uint256) {
        // todo
        uint256 endEpochIndex = 0;

        //?
        if (_deadlineRecord[sequencer][account] != 0) {
            // endEpochIndex > _deadlineRecord[sequencer][account] ????
            endEpochIndex = _deadlineRecord[sequencer][account];
        }

        uint256 claimedEpochIndex = _claimedEpoch[sequencer][account];
        uint256 accountReward = 0;
        uint256 validEpochIndex = 0;
        for (uint256 i = claimedEpochIndex + 1; i <= endEpochIndex; i++) {
            uint256 ratio = IRecords(_record).sequencerEpochRatio(i, sequencer);
            uint256 epochTotalReward = _rewards[i];
            //uint256 sequencerReward = epochTotalReward * ratio / 100;
            if (_collect[sequencer][i].valid) {
                accountReward = (epochTotalReward * ratio * _collect[sequencer][i].shares.value[account]) / (_collect[sequencer][i].totalShare * 100);
                validEpochIndex = i;
            }else {
                for (uint j = i - 1; j > 0; j--) {
                    if (_collect[sequencer][j].valid) {
                        accountReward = (epochTotalReward * ratio * _collect[sequencer][j].amounts.value[account]) / (_collect[sequencer][j].totalAmount * 100);
                        validEpochIndex = j;
                    }
                }
            }
        }

        if (endEpochIndex != validEpochIndex) {
            Distribution storage dt = _collect[sequencer][endEpochIndex];

            dt.totalAmount = _collect[sequencer][validEpochIndex].totalAmount;
            dt.totalShare = _collect[sequencer][validEpochIndex].totalAmount;
            dt.remainNumber = _collect[sequencer][validEpochIndex].amounts.index.length();

            for (uint256 j = 0; j < _collect[sequencer][validEpochIndex].amounts.index.length(); j++) {
                address delegator = _collect[sequencer][validEpochIndex].amounts.index.at(j);
                uint256 delegateAmount = _collect[sequencer][validEpochIndex].amounts.value[delegator];

                dt.shares.index.add(delegator);
                dt.shares.value[delegator] = delegateAmount;

                dt.amounts.index.add(delegator);
                dt.amounts.value[delegator] = delegateAmount;
            }

            dt.valid = true;
        }
        _collect[sequencer][endEpochIndex].remainNumber -= 1;

        _claimedEpoch[sequencer][account] = endEpochIndex;
        //？
        if (endEpochIndex == _deadlineRecord[sequencer][account]) {
            delete  _deadlineRecord[sequencer][account];
        }

        return accountReward;
    }
}

