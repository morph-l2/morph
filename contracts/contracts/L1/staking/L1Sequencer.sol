// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {IL1Sequencer} from "./IL1Sequencer.sol";

contract L1Sequencer is
    Initializable,
    IL1Sequencer,
    Sequencer,
    PausableUpgradeable
{
    // staking contract
    address public stakingContract;
    // rollup Contract
    address public rollupContract;

    // current sequencers version
    uint256 public override currentVersion = 0;
    // newest sequencers version
    uint256 public override newestVersion = 0;
    // map(version => sequencerAddresses)
    mapping(uint256 => address[]) public sequencerAddresses;
    // map(version => sequencerBLSKeys)
    mapping(uint256 => bytes[]) public sequencerBLSKeys;

    /**
     * @notice only staking contract
     */
    modifier onlyStakingContract() {
        require(msg.sender == stakingContract, "only staking contract");
        _;
    }

    /**
     * @notice only rollup contract
     */
    modifier onlyRollupContract() {
        require(msg.sender == rollupContract, "only rollup contract");
        _;
    }

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     */
    constructor(
        address payable _messenger
    ) Sequencer(_messenger, payable(Predeploys.L2_SEQUENCER)) {
        _disableInitializers();
    }

    /**
     * @notice initializer
     * @param _stakingContract staking contract address
     * @param _rollupContract rollup contract address
     */
    function initialize(
        address _stakingContract,
        address _rollupContract
    ) public initializer {
        require(_stakingContract != address(0), "invalid staking contract");
        require(_rollupContract != address(0), "invalid rollup contract");
        stakingContract = _stakingContract;
        rollupContract = _rollupContract;
        _pause();
    }

    function pause() external override onlyStakingContract {
        _pause();
    }

    function unpause() external override onlyStakingContract {
        _unpause();
    }

    function updateSequencersVersion(
        address[] memory _sequencerAddresses,
        bytes[] memory _sequencerBLSKeys
    ) internal {
        newestVersion++;
        sequencerAddresses[newestVersion] = _sequencerAddresses;
        sequencerBLSKeys[newestVersion] = _sequencerBLSKeys;
    }

    /**
     * @notice verify BLS signature
     * @param version sequencer set version
     * @param sequencers sequencers signed
     * @param signature batch signature
     * @param batchHash batch hash
     */
    function verifySignature(
        uint256 version,
        address[] memory sequencers,
        bytes memory signature,
        bytes32 batchHash
    ) external onlyRollupContract whenNotPaused returns (bool) {
        _confirmVersion(version);
        // TODO: verify BLS signature
        sequencers = sequencers;
        signature = signature;
        batchHash = batchHash;
        return true;
    }

    function updateAndSendSequencerSet(
        bytes memory _sequencerBytes,
        address[] memory _sequencerAddresses,
        bytes[] memory _sequencerBLSKeys,
        uint32 _gasLimit
    ) external override onlyStakingContract {
        if (newestVersion == 0 && sequencerAddresses[0].length == 0) {
            // init sequencers
            sequencerAddresses[0] = _sequencerAddresses;
            sequencerBLSKeys[0] = _sequencerBLSKeys;
            _unpause();
        } else {
            require(!paused(), "send message when unpaused");
            updateSequencersVersion(_sequencerAddresses, _sequencerBLSKeys);

            MESSENGER.sendMessage(
                address(OTHER_SEQUENCER),
                0,
                _sequencerBytes,
                _gasLimit
            );
        }

        emit SequencerUpdated(
            newestVersion,
            _sequencerAddresses,
            _sequencerBLSKeys
        );
    }

    /**
     * @notice confirm sequencer ser version
     * @param version sequencer set version
     */
    function _confirmVersion(uint256 version) internal {
        require(
            version >= currentVersion && version <= newestVersion,
            "invalid sequencer version"
        );
        for (uint256 i = currentVersion; i < version; i++) {
            delete sequencerAddresses[i];
            delete sequencerBLSKeys[i];
        }
        currentVersion = version;
    }

    function sequencerNum(uint256 version) external view returns (uint256) {
        return sequencerBLSKeys[version].length;
    }

    /**
     * @notice whether is current sequencer
     * @param addr address
     * @param version sequencer version
     */
    function isSequencer(
        address addr,
        uint256 version
    ) external view returns (bool) {
        require(
            version >= currentVersion && version <= newestVersion,
            "invalid version"
        );
        for (uint256 i = 0; i < sequencerAddresses[version].length; i++) {
            if (sequencerAddresses[currentVersion][i] == addr) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice sequencer addresses
     * @param version version
     */
    function getSequencerAddresses(
        uint256 version
    ) external view returns (address[] memory) {
        return sequencerAddresses[version];
    }

    /**
     * @notice sequencer BLS keys
     * @param version version
     */
    function getSequencerBLSKeys(
        uint256 version
    ) external view returns (bytes[] memory) {
        return sequencerBLSKeys[version];
    }
}
