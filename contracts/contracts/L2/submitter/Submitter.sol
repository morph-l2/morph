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

    // uint256 next batch index;
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
     * @notice constructor
     */
    constructor() {
        L2_SEQUENCER_CONTRACT = Predeploys.L2_SEQUENCER;
        L2_GOV_CONTRACT = Predeploys.L2_GOV;
    }

    function initialize(address[] memory sequencers) public initializer {
        __Ownable_init();
        sequencerHistory.push(SequencerHistory(sequencers, block.timestamp));
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

        uint256 epochStart = start + submitterIndex * epoch;
        uint256 turnPeriod = epoch * sequencersLen;

        if (block.timestamp > epochStart) {
            uint256 turns = (block.timestamp - epochStart) / turnPeriod + 1;
            epochStart += turns * turnPeriod;
        }

        return (epochStart, epochStart + epoch);
    }

    /**
     * @notice get current submitter
     */
    function getCurrentSubmitter()
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
        uint256 currentSubmitterIndex = turns % sequencersLen;

        uint256 currentEpochStart = block.timestamp -
            ((block.timestamp - start) % epoch);

        return (
            sequencers[currentSubmitterIndex],
            currentEpochStart,
            currentEpochStart + epoch
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
