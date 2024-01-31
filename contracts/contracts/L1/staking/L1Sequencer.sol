// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Sequencer} from "../../libraries/sequencer/Sequencer.sol";
import {Semver} from "../../libraries/common/Semver.sol";
import {IL1Sequencer} from "./IL1Sequencer.sol";

contract L1Sequencer is
    Initializable,
    IL1Sequencer,
    Sequencer,
    Semver,
    Pausable
{
    // staking contract
    address public stakingContract;
    // rollup Contract
    address public rollupContract;

    // current sequencers version
    uint256 public override currentVersion = 0;
    // newest sequencers version
    uint256 public override newestVersion = 0;
    // map(version => sequencerBLSkeys)
    mapping(uint256 => bytes[]) public sequencerBLSKeys;

    /**
     * @notice xxx
     */
    event SequencerConfirmed(address[] sequencers, uint256 version);

    /**
     * @notice xxx
     */
    modifier onlyStakingContract() {
        require(msg.sender == stakingContract, "only staking contract");
        _;
    }

    /**
     * @notice xxx
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
    )
        Semver(1, 0, 0)
        Pausable()
        Sequencer(_messenger, payable(Predeploys.L2_SEQUENCER))
    {
        _pause();
    }

    /**
     * @notice do not receive ETH
     */
    receive() external payable {
        require(false);
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

    function pause() external override onlyStakingContract whenNotPaused {
        _pause();
    }

    function updateSequencersVersion(
        bytes[] memory _sequencerBLSKeys
    ) internal {
        if (newestVersion == 0) {
            _unpause();
        }
        require(!paused(), "send message when unpaused");
        newestVersion++;
        sequencerBLSKeys[newestVersion] = _sequencerBLSKeys;
    }

    /**
     * @notice verify BLS signature
     * @param version sequencer set version
     * @param indexs sequencer index
     * @param signature batch signature
     */
    function verifySignature(
        uint256 version,
        uint256[] memory indexs,
        bytes memory signature
    ) external onlyRollupContract whenNotPaused {
        confirmVersion(version);
        // TODO verify BLS signature
    }

    /**
     * @notice confirm sequencer ser version
     * @param version sequencer set version
     */
    function confirmVersion(uint256 version) internal {
        require(
            version >= currentVersion && version <= newestVersion,
            "invalid version"
        );
        for (uint256 i = 1; i <= version; i++) {
            delete sequencerBLSKeys[i];
        }
        currentVersion = version;
    }

    function updateAndSendSequencerSet(
        bytes memory _sequencerBytes,
        bytes[] memory _sequencerBLSKeys,
        uint32 _gasLimit,
        address _refundAddress
    ) external payable override onlyStakingContract {
        updateSequencersVersion(_sequencerBLSKeys);
        require(!paused(), "send message when unpaused");
        MESSENGER.sendMessage(
            address(OTHER_SEQUENCER),
            0,
            _sequencerBytes,
            _gasLimit,
            _refundAddress
        );
    }

    function getSequencerBLSKeys(
        uint256 version,
        uint256 index
    ) external view returns (bytes memory) {
        uint256 blsKeyNum = sequencerBLSKeys[version].length;
        if (blsKeyNum > 0 && index + 1 <= blsKeyNum) {
            return sequencerBLSKeys[version][index];
        }
        return bytes("");
    }

    function getSequencerBLSKeysLength(
        uint256 version
    ) external view returns (uint256) {
        uint256 blsKeyNum = sequencerBLSKeys[version].length;
        return blsKeyNum;
    }
}
