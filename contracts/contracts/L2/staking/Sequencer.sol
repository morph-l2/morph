// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";

contract Sequencer is Initializable, ISequencer {
    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// The hash of latest three sequencer set.
    bytes32 public SEQUENCER_SET_VERIFY_HASH;

    // The latest three sequencerSet changes and effective height
    // multiple changes within a block only record the final state
    uint256 public blockHeight0;
    address[] public sequencerSet0;
    uint256 public blockHeight1;
    address[] public sequencerSet1;
    uint256 public blockHeight2;
    address[] public sequencerSet2;

    /*********************** modifiers **************************/

    /**
     * @notice only L2Staking contract
     */
    modifier onlyL2StakingContract() {
        require(msg.sender == L2_STAKING_CONTRACT, "only L2Staking contract");
        _;
    }

    /*********************** Constructor **************************/

    /**
     * @notice constructor
     */
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
    }

    /*********************** Init **************************/

    /**
     * @notice Initializer.
     * @param sequencerSet  sequencer set
     */
    function initialize(address[] memory sequencerSet) public initializer {
        require(sequencerSet.length > 0, "invalid sequencer set");
        sequencerSet0 = sequencerSet;
        sequencerSet1 = sequencerSet;
        sequencerSet2 = sequencerSet;
    }

    /*********************** External Functions **************************/

    /**
     * @notice update sequencer set. If new sequencer set is nil, layer2 will stop producing blocks
     */
    function updateSequencerSet(
        address[] memory newSequencerSet
    ) public onlyL2StakingContract {
        // sequencerSet changes will take effect after two blocks
        // The current block height +2 can only be greater than or equal to the last record
        if ((block.number + 2) > blockHeight2) {
            blockHeight0 = blockHeight1;
            blockHeight1 = blockHeight2;
            blockHeight2 = block.number + 2;

            sequencerSet0 = sequencerSet1;
            sequencerSet1 = sequencerSet2;
            sequencerSet2 = newSequencerSet;
        } else {
            sequencerSet2 = newSequencerSet;
        }

        // ************************************************
        // update SEQUENCER_VERIFY_HASH
        // ************************************************
        SEQUENCER_SET_VERIFY_HASH = keccak256(
            abi.encodePacked(
                blockHeight0,
                sequencerSet0,
                blockHeight1,
                sequencerSet1,
                blockHeight2,
                sequencerSet2
            )
        );

        emit SequencerSetUpdated(newSequencerSet, block.number + 2);
    }

    /*********************** External View Functions **************************/

    /**
     * @notice get current sequencer set
     */
    function getCurrentSequencerSet() external view returns (address[] memory) {
        if (block.number >= blockHeight2) {
            return sequencerSet2;
        }
        if (block.number >= blockHeight1) {
            return sequencerSet1;
        }
        return sequencerSet0;
    }

    /**
     * @notice get current sequencer set size
     */
    function getCurrentSequencerSetSize() external view returns (uint256) {
        if (block.number >= blockHeight2) {
            return sequencerSet2.length;
        }
        if (block.number >= blockHeight1) {
            return sequencerSet1.length;
        }
        return sequencerSet0.length;
    }

    /**
     * @notice get sequencer set 0
     */
    function getSequencerSet0() external view returns (address[] memory) {
        return sequencerSet0;
    }

    /**
     * @notice get size of sequencer set 0
     */
    function getSequencerSet0Size() external view returns (uint256) {
        return sequencerSet0.length;
    }

    /**
     * @notice get sequencer set 1
     */
    function getSequencerSet1() external view returns (address[] memory) {
        return sequencerSet1;
    }

    /**
     * @notice get size of sequencer set 1
     */
    function getSequencerSet1Size() external view returns (uint256) {
        return sequencerSet1.length;
    }

    /**
     * @notice get sequencer set 2
     */
    function getSequencerSet2() external view returns (address[] memory) {
        return sequencerSet2;
    }

    /**
     * @notice get size of sequencer set 2
     */
    function getSequencerSet2Size() external view returns (uint256) {
        return sequencerSet2.length;
    }

    /**
     * @notice whether the address is a sequencer
     */
    function isSequencer(address addr) external view returns (bool) {
        return _contains(sequencerSet2, addr);
    }

    /**
     * @notice whether the address is a current sequencer
     */
    function isCurrentSequencer(address addr) external view returns (bool) {
        if (block.number >= blockHeight2) {
            return _contains(sequencerSet2, addr);
        }
        if (block.number >= blockHeight1) {
            return _contains(sequencerSet1, addr);
        }
        return _contains(sequencerSet0, addr);
    }

    /*********************** Internal Functions **************************/

    /**
     * @notice whether the address is the address list
     */
    function _contains(
        address[] memory addressList,
        address addr
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < addressList.length; i++) {
            if (addr == addressList[i]) {
                return true;
            }
        }
        return false;
    }
}
