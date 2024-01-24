// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Semver} from "../../libraries/common/Semver.sol";
import {RollupMessage} from "../../libraries/submitter/RollupMessage.sol";
import {IL2Sequencer} from "../staking/IL2Sequencer.sol";
import {IGov} from "../staking/IGov.sol";
import {ISubmitter} from "./ISubmitter.sol";

contract Submitter is Initializable, Semver, ISubmitter, RollupMessage {
    // l2SequencerContract address
    address public immutable L2_SEQUENCER_CONTRACT;
    // GovContract address
    address public immutable L2_GOV_CONTRACT;

    // uint256 public override nextSubmitterIndex;
    uint256 public override nextBatchIndex;
    // next batch start block
    uint256 public override nextBatchStartBlock;
    // bathcIndex => batchInfo
    mapping(uint256 => Types.BatchInfo) public confirmedBatchs;

    // next epoch start time
    uint256 public nextEpochStart;
    // next submitter index
    uint256 public nextSubmitterIndex;
    // calculated epoch index
    uint256 public calculatedEpochIndex;
    // epoch info
    mapping(uint256 => Types.EpochInfo) public epochs;

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
    constructor(
        address payable _rollup
    )
        Semver(1, 0, 0)
        RollupMessage(payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER), _rollup)
    {
        L2_SEQUENCER_CONTRACT = Predeploys.L2_SEQUENCER;
        L2_GOV_CONTRACT = Predeploys.L2_GOV;
    }

    /**
     * @notice Initializer.
     * @param _nextEpochStart next epoch start time
     */
    function initialize(uint256 _nextEpochStart) public initializer {
        require(_nextEpochStart > 0, "invalid firstEpochStart");
        nextEpochStart = _nextEpochStart;
    }

    /**
     * @notice set rollup acknowledge, only call by bridge
     */
    function ackRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    ) public onlyCounterpart {
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
        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerAddresses(false);
        updateEpoch(epoch, sequencers);
    }

    /**
     * @notice sequencers updated
     */
    function sequencersUpdated(address[] memory sequencers) public {
        require(
            msg.sender == L2_SEQUENCER_CONTRACT,
            "only l2 sequencer contract"
        );
        uint256 epoch = IGov(L2_GOV_CONTRACT).rollupEpoch();
        nextSubmitterIndex = 0;
        updateEpoch(epoch, sequencers);
    }

    /**
     * @notice statistics timeout behavior
     */
    function updateEpoch(uint256 epoch, address[] memory sequencers) internal {
        uint256 sequencersLen = sequencers.length;
        while (nextEpochStart + epoch <= block.timestamp) {
            calculatedEpochIndex++;
            epochs[calculatedEpochIndex] = Types.EpochInfo(
                sequencers[nextSubmitterIndex],
                nextEpochStart,
                nextEpochStart + epoch
            );
            nextSubmitterIndex++;
            if (nextSubmitterIndex == sequencersLen) {
                nextSubmitterIndex = 0;
            }
            nextEpochStart += epoch;
        }
        emit EpochUpdated(epoch, sequencersLen);
    }

    /**
     * @notice update epoch external
     */
    function updateEpochExternal() public {
        // update epoch
        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerAddresses(false);
        uint256 epoch = IGov(L2_GOV_CONTRACT).rollupEpoch();
        updateEpoch(epoch, sequencers);
    }

    // ============================================================================

    /**
     * @notice get the current sequencer's turn
     */
    function getTurn(
        address submitter
    ) external view returns (uint256, uint256) {
        // update epoch
        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerAddresses(false);
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

        uint256 _nextEpochStart = nextEpochStart;
        uint256 _nextSubmitterIndex = nextSubmitterIndex;

        while (_nextEpochStart + epoch <= block.timestamp) {
            _nextSubmitterIndex++;
            if (_nextSubmitterIndex == sequencersLen) {
                _nextSubmitterIndex = 0;
            }
            _nextEpochStart += epoch;
        }

        if (submitterIndex > _nextSubmitterIndex) {
            uint256 startTime = (submitterIndex - _nextSubmitterIndex) * epoch;
            return (startTime, startTime + epoch);
        } else if (submitterIndex < _nextSubmitterIndex) {
            uint256 startTime = (sequencersLen -
                _nextSubmitterIndex +
                submitterIndex) * epoch;
            return (startTime, startTime + epoch);
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
        // update epoch
        address[] memory sequencers = IL2Sequencer(L2_SEQUENCER_CONTRACT)
            .getSequencerAddresses(false);
        uint256 epoch = IGov(L2_GOV_CONTRACT).rollupEpoch();
        uint256 sequencersLen = sequencers.length;

        uint256 _nextEpochStart = nextEpochStart;
        uint256 _nextSubmitterIndex = nextSubmitterIndex;

        while (_nextEpochStart + epoch <= block.timestamp) {
            _nextSubmitterIndex++;
            if (_nextSubmitterIndex == sequencersLen) {
                _nextSubmitterIndex = 0;
            }
            _nextEpochStart += epoch;
        }

        return (
            sequencers[nextSubmitterIndex],
            _nextEpochStart,
            _nextEpochStart + epoch
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
