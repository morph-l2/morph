// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {ISubmitter} from "./ISubmitter.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L2Sequencer is Initializable, ISequencer {
    // submitter contract address
    address public immutable L2_SUBMITTER_CONTRACT;
    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    // current sequencers version
    uint256 public override currentVersion = 0;
    uint256 public override currentVersionHeight = 0;
    uint256 public override preVersion = 0;
    uint256 public override preVersionHeight = 0;

    // addresses of sequencerSet
    address[] public override sequencerAddresses;
    address[] public override preSequencerAddresses;

    // sequencer infos array
    Types.StakerInfo[] public sequencerInfos;
    Types.StakerInfo[] public preSequencerInfos;

    // event of sequencer update
    event SequencerUpdated(address[] sequencers, uint256 version);

    /**
     * @notice only L2Staking contract
     */
    modifier onlyL2StakingContract() {
        require(msg.sender == L2_STAKING_CONTRACT, "only L2Staking contract");
        _;
    }

    constructor() {
        L2_SUBMITTER_CONTRACT = Predeploys.SUBMITTER;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
    }

    function initialize(
        Types.StakerInfo[] memory _sequencers
    ) public initializer {
        for (uint256 i = 0; i < _sequencers.length; i++) {
            sequencerAddresses.push(_sequencers[i].addr);
            sequencerInfos.push(_sequencers[i]);
        }
    }

    function updateSequencers(
        uint256 version,
        Types.StakerInfo[] memory _sequencers
    ) public onlyL2StakingContract {
        ISubmitter(L2_SUBMITTER_CONTRACT).sequencersUpdated(sequencerAddresses);

        preVersion = currentVersion;
        delete preSequencerInfos;
        delete preSequencerAddresses;
        preSequencerInfos = sequencerInfos;
        preSequencerAddresses = sequencerAddresses;
        preVersionHeight = currentVersionHeight;

        currentVersion = version;
        delete sequencerInfos;
        delete sequencerAddresses;
        currentVersionHeight = block.number;

        for (uint256 i = 0; i < _sequencers.length; i++) {
            sequencerAddresses.push(_sequencers[i].addr);
            sequencerInfos.push(_sequencers[i]);
        }
        emit SequencerUpdated(sequencerAddresses, version);
    }

    /**
     * @notice get sequencers addresses
     */
    function getSequencerAddresses(
        bool previous
    ) external view returns (address[] memory) {
        if (previous) {
            return preSequencerAddresses;
        }
        return sequencerAddresses;
    }

    /**
     * @notice get sequencers infos
     */
    function getSequencerInfos(
        bool previous
    ) external view returns (Types.StakerInfo[] memory) {
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
        if (previous) {
            for (uint256 i = 0; i < preSequencerAddresses.length; i++) {
                if (checkAddr == preSequencerAddresses[i]) {
                    return (true, preVersion);
                }
            }
            return (false, preVersion);
        }

        for (uint256 i = 0; i < sequencerAddresses.length; i++) {
            if (checkAddr == sequencerAddresses[i]) {
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
        if (previous) {
            for (uint256 i = 0; i < preSequencerAddresses.length; i++) {
                if (checkAddr == preSequencerAddresses[i]) {
                    return (i, preVersion);
                }
            }
            revert("sequencer not exist");
        }
        for (uint256 i = 0; i < sequencerAddresses.length; i++) {
            if (checkAddr == sequencerAddresses[i]) {
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
        if (previous) {
            return (preSequencerAddresses.length, preVersion);
        }
        return (sequencerAddresses.length, currentVersion);
    }
}
