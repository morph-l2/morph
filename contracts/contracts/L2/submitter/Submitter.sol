// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IL2Sequencer} from "../staking/IL2Sequencer.sol";
import {IGov} from "../staking/IGov.sol";
import {ISubmitter} from "./ISubmitter.sol";

contract Submitter is ISubmitter, OwnableUpgradeable {
    // l2SequencerContract address
    address public immutable L2_SEQUENCER_CONTRACT;
    // GovContract address
    address public immutable L2_GOV_CONTRACT;

    // uint256 next batch index;
    uint256 public override nextBatchIndex;
    // next batch start block
    uint256 public override nextBatchStartBlock;
    // batchIndex => batchInfo
    mapping(uint256 => Types.BatchInfo) public confirmedBatches;

    Types.EpochHistory[] public epochHistory;

    /**
     * @notice constructor
     */
    constructor() {
        L2_SEQUENCER_CONTRACT = Predeploys.L2_SEQUENCER;
        L2_GOV_CONTRACT = Predeploys.L2_GOV;
    }

    function initialize() public initializer {
        __Ownable_init();
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

        confirmedBatches[batchIndex] = Types.BatchInfo(
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
        epochHistory.push(Types.EpochHistory(epoch, block.timestamp));
    }

    // ============================================================================

    /**
     * @notice get the current sequencer's turn
     */
    function getTurn(
        address submitter
    ) external view returns (uint256, uint256) {
        uint256 currentVersion = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .currentVersion();

        uint256 start = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerHistory(currentVersion)
            .timestamp;

        if (
            epochHistory.length > 0 &&
            epochHistory[epochHistory.length - 1].timestamp > start
        ) {
            start = epochHistory[epochHistory.length - 1].timestamp;
        }

        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerHistory(currentVersion)
            .sequencerAddresses;
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
        uint256 currentVersion = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .currentVersion();

        uint256 start = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerHistory(currentVersion)
            .timestamp;

        if (
            epochHistory.length > 0 &&
            epochHistory[epochHistory.length - 1].timestamp > start
        ) {
            start = epochHistory[epochHistory.length - 1].timestamp;
        }

        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerHistory(currentVersion)
            .sequencerAddresses;

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
        return confirmedBatches[batchIndex];
    }
}
