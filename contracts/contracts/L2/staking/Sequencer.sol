// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";

contract L2Sequencer is Initializable, ISequencer {
    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// @notice The hash of sequencer set list.
    bytes32 public SEQUENCER_SET_VERIFY_HASH;

    // The latest three sequencerSet changes and effective height
    // multiple changes within a block only record the final state
    uint256 public blockHeight0;
    address[] public sequencerSet0;
    uint256 public blockHeight1;
    address[] public sequencerSet1;
    uint256 public blockHeight2;
    address[] public sequencerSet2;

    // event of sequencer update
    event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight);

    /**
     * @notice only L2Staking contract
     */
    modifier onlyL2StakingContract() {
        require(msg.sender == L2_STAKING_CONTRACT, "only L2Staking contract");
        _;
    }

    /**
     * @notice constructor
     */
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
    }

    function initialize(address[] memory sequencerSet) public initializer {
        require(sequencerSet.length > 0, "invalid sequencer set");
        sequencerSet0 = sequencerSet;
        sequencerSet1 = sequencerSet;
        sequencerSet2 = sequencerSet;
    }

    /**
     * @notice update sequencer set
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

    /**
     * @notice get current sequencer set
     */
    function getCurrentSeqeuncerSet() external view returns (address[] memory) {
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
    function getCurrentSeqeuncerSetSize() external view returns (uint256) {
        if (block.number >= blockHeight2) {
            return sequencerSet2.length;
        }
        if (block.number >= blockHeight1) {
            return sequencerSet1.length;
        }
        return sequencerSet0.length;
    }

    /**
     * @notice whether the address is a sequencer
     */
    function isSequencer(address addr) external view returns (bool) {
        if (block.number >= blockHeight2) {
            return inAddressList(addr, sequencerSet2);
        }
        if (block.number >= blockHeight1) {
            return inAddressList(addr, sequencerSet1);
        }
        return inAddressList(addr, sequencerSet0);
    }

    /**
     * @notice whether the address is the address list
     */
    function inAddressList(
        address addr,
        address[] memory addressList
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < addressList.length; i++) {
            if (addr == addressList[i]) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice get latest sequencer set
     */
    function getLatestSeqeuncerSet() external view returns (address[] memory) {
        return sequencerSet2;
    }
}
