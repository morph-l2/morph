// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {IL2Sequencer} from "./IL2Sequencer.sol";
import {ISubmitter} from "../submitter/ISubmitter.sol";

contract L2Sequencer is Initializable, IL2Sequencer, Sequencer {
    // submitter contract address
    address public immutable L2_SUBMITTER_CONTRACT;

    // current sequencers version
    uint256 public override currentVersion = 0;
    uint256 public override currentVersionHeight = 0;
    uint256 public override preVersionHeight = 0;

    mapping(uint256 => Types.SequencerHistory) public sequencerHistory;

    // sequencer infos array
    Types.SequencerInfo[] public sequencerInfos;
    Types.SequencerInfo[] public preSequencerInfos;

    // event of sequencer update
    event SequencerUpdated(address[] sequencers, uint256 version);

    /**
     *
     * @param _otherSequencer Address of the sequencer on the other network.
     */
    constructor(
        address payable _otherSequencer
    )
        Sequencer(
            payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER),
            _otherSequencer
        )
    {
        L2_SUBMITTER_CONTRACT = Predeploys.L2_SUBMITTER;
    }

    function initialize(
        Types.SequencerInfo[] memory _sequencers
    ) public initializer {
        address[] memory sequencerAddresses = new address[](_sequencers.length);
        for (uint256 i = 0; i < _sequencers.length; i++) {
            sequencerAddresses[i] = _sequencers[i].addr;
            sequencerInfos.push(_sequencers[i]);
        }
        sequencerHistory[0] = Types.SequencerHistory(
            sequencerAddresses,
            block.timestamp
        );
    }

    function updateSequencers(
        uint256 version,
        Types.SequencerInfo[] memory _sequencers
    ) public onlyOtherSequencer {
        require(version == currentVersion + 1, "invalid version");
        currentVersion = version;

        delete preSequencerInfos;
        preSequencerInfos = sequencerInfos;
        delete sequencerInfos;

        preVersionHeight = currentVersionHeight;
        currentVersionHeight = block.number;

        address[] memory sequencerAddresses = new address[](_sequencers.length);
        for (uint256 i = 0; i < _sequencers.length; i++) {
            sequencerAddresses[i] = _sequencers[i].addr;
            sequencerInfos.push(_sequencers[i]);
        }

        sequencerHistory[version] = Types.SequencerHistory(
            sequencerAddresses,
            block.timestamp
        );

        emit SequencerUpdated(sequencerAddresses, version);
    }

    /**
     * @notice get sequencer history
     */
    function getSequencerHistory(
        uint256 version
    ) external view returns (Types.SequencerHistory memory) {
        return sequencerHistory[version];
    }

    /**
     * @notice get sequencers addresses
     */
    function getSequencerAddresses(
        bool previous
    ) external view returns (address[] memory) {
        if (previous && currentVersion > 0) {
            return sequencerHistory[currentVersion - 1].sequencerAddresses;
        }
        return sequencerHistory[currentVersion].sequencerAddresses;
    }

    /**
     * @notice get sequencers infos
     */
    function getSequencerInfos(
        bool previous
    ) external view returns (Types.SequencerInfo[] memory) {
        if (previous) {
            return preSequencerInfos;
        }
        return sequencerInfos;
    }

    /**
     * @notice get address is in sequencers set
     */
    function inSequencersSet(
        bool previous,
        address checkAddr
    ) external view returns (bool, uint256) {
        if (previous && currentVersion > 0) {
            for (
                uint256 i = 0;
                i <
                sequencerHistory[currentVersion - 1].sequencerAddresses.length;
                i++
            ) {
                if (
                    checkAddr ==
                    sequencerHistory[currentVersion - 1].sequencerAddresses[i]
                ) {
                    return (true, currentVersion - 1);
                }
            }
            return (false, currentVersion - 1);
        }

        for (
            uint256 i = 0;
            i < sequencerHistory[currentVersion].sequencerAddresses.length;
            i++
        ) {
            if (
                checkAddr ==
                sequencerHistory[currentVersion].sequencerAddresses[i]
            ) {
                return (true, currentVersion);
            }
        }
        return (false, currentVersion);
    }

    /**
     * @notice get the index of address in sequencers set
     */
    function sequencerIndex(
        bool previous,
        address checkAddr
    ) external view returns (uint256, uint256) {
        if (previous && currentVersion > 0) {
            for (
                uint256 i = 0;
                i <
                sequencerHistory[currentVersion - 1].sequencerAddresses.length;
                i++
            ) {
                if (
                    checkAddr ==
                    sequencerHistory[currentVersion - 1].sequencerAddresses[i]
                ) {
                    return (i, currentVersion - 1);
                }
            }
            revert("sequencer not exist");
        }
        for (
            uint256 i = 0;
            i < sequencerHistory[currentVersion].sequencerAddresses.length;
            i++
        ) {
            if (
                checkAddr ==
                sequencerHistory[currentVersion].sequencerAddresses[i]
            ) {
                return (i, currentVersion);
            }
        }
        revert("sequencer not exist");
    }

    /**
     * @notice get the length of sequencerAddresses
     */
    function sequencersLen(
        bool previous
    ) external view returns (uint256, uint256) {
        if (previous && currentVersion > 0) {
            return (
                sequencerHistory[currentVersion - 1].sequencerAddresses.length,
                currentVersion - 1
            );
        }
        return (
            sequencerHistory[currentVersion].sequencerAddresses.length,
            currentVersion
        );
    }
}
