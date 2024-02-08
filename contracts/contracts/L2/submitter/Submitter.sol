// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IL2Sequencer} from "../staking/IL2Sequencer.sol";
import {IGov} from "../staking/IGov.sol";
import {ISubmitter} from "./ISubmitter.sol";

contract Submitter is ISubmitter, OwnableUpgradeable {
    struct SequencerHistory {
        address[] sequencerAddresses;
        uint256 timestamp;
    }

    struct EpochHistory {
        uint256 epoch;
        uint256 timestamp;
    }

    // l2SequencerContract address
    address public immutable L2_SEQUENCER_CONTRACT;
    // GovContract address
    address public immutable L2_GOV_CONTRACT;

    // uint256 bathc index;
    uint256 public override nextBatchIndex;
    // next batch start block
    uint256 public override nextBatchStartBlock;
    // bathcIndex => batchInfo
    mapping(uint256 => Types.BatchInfo) public confirmedBatchs;

    // epoch info
    mapping(uint256 => Types.EpochInfo) public epochs;

    SequencerHistory[] public sequencerHistory;
    EpochHistory[] public epochHistory;

    /**
     * @notice ack rollup
     */
    event ACKRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    );

    /**
     * @notice epoch updated
     */
    event EpochUpdated(uint256 interval, uint256 sequencersLen);

    /**
     * @notice constructor
     */
    constructor() {
        L2_SEQUENCER_CONTRACT = Predeploys.L2_SEQUENCER;
        L2_GOV_CONTRACT = Predeploys.L2_GOV;
    }

    /**
     * @notice set rollup acknowledge
     */
    function ackRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    ) public onlyOwner {
        require(batchIndex == nextBatchIndex, "invalid batchIndex");
        require(
            batchStartBlock == nextBatchStartBlock,
            "invalid batchStartBlock"
        );
        require(true);
        confirmedBatchs[batchIndex] = Types.BatchInfo(
            submitter,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );

        emit ACKRollup(
            batchIndex,
            submitter,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );

        nextBatchIndex++;
        nextBatchStartBlock = batchEndBlock + 1;
    }

    /**
     * @notice epoch updated
     */
    function epochUpdated(uint256 epoch) public {
        require(msg.sender == L2_GOV_CONTRACT, "only gov contract");
        epochHistory.push(EpochHistory(epoch, block.timestamp));
    }

    /**
     * @notice sequencers updated
     */
    function sequencersUpdated(address[] memory sequencers) public {
        require(
            msg.sender == L2_SEQUENCER_CONTRACT,
            "only l2 sequencer contract"
        );
        sequencerHistory.push(SequencerHistory(sequencers, block.timestamp));
    }

    // ============================================================================

    /**
     * @notice get the current sequencer's turn
     */
    function getTurn(
        address submitter
    ) external view returns (uint256, uint256) {
        uint256 start = sequencerHistory[sequencerHistory.length - 1].timestamp;

        if (
            epochHistory.length > 0 &&
            epochHistory[epochHistory.length - 1].timestamp > start
        ) {
            start = epochHistory[epochHistory.length - 1].timestamp;
        }

        address[] memory sequencers = sequencerHistory[
            sequencerHistory.length - 1
        ].sequencerAddresses;
        uint256 epoch = IGov(L2_GOV_CONTRACT).rollupEpoch();

        uint256 sequencersLen = sequencers.length;

        bool exist = false;
        uint256 submitterIndex = 0;
        while (submitterIndex < sequencersLen) {
            if (submitter == sequencers[submitterIndex]) {
                exist = true;
                break;
            }
            submitterIndex++;
        }
        require(exist, "invalid submitter");

        uint256 nextEpochStart = start + submitterIndex * epoch;

        uint256 turnPeriod = epoch * sequencersLen;

        if (block.timestamp > nextEpochStart) {
            uint256 turns = (block.timestamp - nextEpochStart) / turnPeriod + 1;
            nextEpochStart += turns * turnPeriod;
        }

        return (nextEpochStart, nextEpochStart + epoch);
    }

    /**
     * @notice get next submitter
     */
    function getNextSubmitter()
        external
        view
        returns (address, uint256, uint256)
    {
        require(sequencerHistory.length > 0, "invalid sequencer");
        uint256 start = sequencerHistory[sequencerHistory.length - 1].timestamp;

        if (
            epochHistory.length > 0 &&
            epochHistory[epochHistory.length - 1].timestamp > start
        ) {
            start = epochHistory[epochHistory.length - 1].timestamp;
        }

        address[] memory sequencers = sequencerHistory[
            sequencerHistory.length - 1
        ].sequencerAddresses;
        uint256 epoch = IGov(L2_GOV_CONTRACT).rollupEpoch();
        uint256 sequencersLen = sequencers.length;

        uint256 turns = (block.timestamp - start) / epoch;
        uint256 nextSubmitterIndex = turns % sequencersLen;

        uint256 nextEpochStart = block.timestamp -
            ((block.timestamp - start) % epoch);

        return (
            sequencers[nextSubmitterIndex],
            nextEpochStart,
            nextEpochStart + epoch
        );
    }

    /**
     * @notice get confirmed batch info
     */
    function getConfirmedBatch(
        uint256 batchIndex
    ) external view returns (Types.BatchInfo memory batchInfo) {
        return confirmedBatchs[batchIndex];
    }

    /**
     * @notice get epoch info
     */
    function getEpoch(
        uint256 epochIndex
    ) external view returns (Types.EpochInfo memory epochInfo) {
        return epochs[epochIndex];
    }
}
