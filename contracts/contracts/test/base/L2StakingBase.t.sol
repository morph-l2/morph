// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {MorphToken} from "../../l2/system/MorphToken.sol";
import {L2Staking} from "../../l2/staking/L2Staking.sol";
import {Sequencer} from "../../l2/staking/Sequencer.sol";
import {Distribute} from "../../l2/staking/Distribute.sol";
import {Record} from "../../l2/staking/Record.sol";
import {Gov} from "../../l2/staking/Gov.sol";
import {L2MessageBaseTest} from "./L2MessageBase.t.sol";

contract L2StakingBaseTest is L2MessageBaseTest {
    uint256 public constant SEQUENCER_RATIO_PRECISION = 1e8;
    uint256 public constant INFLATION_RATIO_PRECISION = 1e16;
    uint256 public constant REWARD_EPOCH = 86400;

    uint256 public beginSeq = 10;
    uint256 public version = 0;

    bytes[] public sequencerBLSKeys;
    address[] public sequencerAddresses;

    uint256 public constant SEQUENCER_SIZE = 3;
    uint256 public constant NEXT_EPOCH_START = 1700000000;
    uint256 public rewardStartTime = 86400;

    // Sequencer config
    Sequencer public sequencer;

    // Gov config
    Gov public gov;

    // L2Staking instance
    L2Staking public l2Staking;

    // Morph token
    MorphToken public morphToken;

    // Distribute
    Distribute public distribute;

    // Record
    Record public record;

    //Oracle address
    address public oracleAddress = address(1000);
    uint256 public nextBatchSubmissionIndex = 1;

    uint256 public constant VOTING_DURATION = 1000;
    uint256 public constant ROLLUP_EPOCH = 1000;
    uint256 public constant MAX_CHUNKS = 1000000000;

    function setUp() public virtual override {
        super.setUp();
        // Set the proxy at the correct address
        hevm.etch(
            Predeploys.SEQUENCER,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.GOV,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.L2_STAKING,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.MORPH_TOKEN,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.DISTRIBUTE,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.RECORD,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        TransparentUpgradeableProxy sequencerProxy = TransparentUpgradeableProxy(payable(Predeploys.SEQUENCER));
        TransparentUpgradeableProxy govProxy = TransparentUpgradeableProxy(payable(Predeploys.GOV));
        TransparentUpgradeableProxy l2StakingProxy = TransparentUpgradeableProxy(payable(Predeploys.L2_STAKING));
        TransparentUpgradeableProxy morphTokenProxy = TransparentUpgradeableProxy(payable(Predeploys.MORPH_TOKEN));
        TransparentUpgradeableProxy distributeProxy = TransparentUpgradeableProxy(payable(Predeploys.DISTRIBUTE));
        TransparentUpgradeableProxy recordProxy = TransparentUpgradeableProxy(payable(Predeploys.RECORD));
        hevm.store(address(sequencerProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(address(govProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(address(l2StakingProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(address(morphTokenProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(address(distributeProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(address(recordProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));

        hevm.startPrank(multisig);
        // deploy impl contracts
        MorphToken morphTokenImpl = new MorphToken();
        L2Staking l2StakingImpl = new L2Staking(payable(NON_ZERO_ADDRESS));
        Sequencer sequencerImpl = new Sequencer();
        Distribute distributeImpl = new Distribute();
        Record recordImpl = new Record();
        Gov govImpl = new Gov();

        // upgrade proxy
        Types.StakerInfo[] memory stakerInfos = new Types.StakerInfo[](SEQUENCER_SIZE);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos[i] = stakerInfo;
            sequencerAddresses.push(stakerInfo.addr);
        }
        ITransparentUpgradeableProxy(address(sequencerProxy)).upgradeToAndCall(
            address(sequencerImpl),
            abi.encodeCall(Sequencer.initialize, (multisig, sequencerAddresses))
        );
        ITransparentUpgradeableProxy(address(govProxy)).upgradeToAndCall(
            address(govImpl),
            abi.encodeCall(
                Gov.initialize,
                (
                    multisig,
                    VOTING_DURATION, // _votingDuration
                    0, // _batchBlockInterval
                    0, // _batchMaxBytes
                    finalizationPeriodSeconds, // _batchTimeout
                    MAX_CHUNKS, // maxChunks
                    ROLLUP_EPOCH // rollupEpoch
                )
            )
        );

        ITransparentUpgradeableProxy(address(l2StakingProxy)).upgradeToAndCall(
            address(l2StakingImpl),
            abi.encodeCall(
                L2Staking.initialize,
                (multisig, SEQUENCER_SIZE * 2, ROLLUP_EPOCH, rewardStartTime, stakerInfos)
            )
        );
        ITransparentUpgradeableProxy(address(morphTokenProxy)).upgradeToAndCall(
            address(morphTokenImpl),
            abi.encodeCall(MorphToken.initialize, ("Morph", "MPH", multisig, 1000000000 ether, 1596535874529))
        );
        ITransparentUpgradeableProxy(address(distributeProxy)).upgradeToAndCall(
            address(distributeImpl),
            abi.encodeCall(Distribute.initialize, (multisig))
        );
        ITransparentUpgradeableProxy(address(recordProxy)).upgradeToAndCall(
            address(recordImpl),
            abi.encodeCall(Record.initialize, (multisig, oracleAddress, nextBatchSubmissionIndex))
        );

        // set address
        sequencer = Sequencer(payable(address(sequencerProxy)));
        gov = Gov(payable(address(govProxy)));
        l2Staking = L2Staking(payable(address(l2StakingProxy)));
        distribute = Distribute(payable(address(distributeProxy)));
        record = Record(payable(address(recordProxy)));
        morphToken = MorphToken(payable(address(morphTokenProxy)));

        _changeAdmin(address(sequencer));
        _changeAdmin(address(gov));
        _changeAdmin(address(l2Staking));
        _changeAdmin(address(distribute));
        _changeAdmin(address(record));
        _changeAdmin(address(morphToken));

        hevm.stopPrank();
    }
}
