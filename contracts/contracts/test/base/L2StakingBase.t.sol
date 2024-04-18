// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/console2.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Sequencer} from "../../L2/staking/Sequencer.sol";
import {L2Staking} from "../../L2/staking/L2Staking.sol";
import {Distribute} from "../../L2/staking/Distribute.sol";
import {Gov} from "../../L2/staking/Gov.sol";
import {L2MessageBaseTest} from "./L2MessageBase.t.sol";
import {MorphToken} from "../../L2/system/MorphToken.sol";
import {Record} from "../../L2/staking/Record.sol";
import {L2MessageBaseTest} from "./L2MessageBase.t.sol";

contract L2StakingBaseTest is L2MessageBaseTest {
    uint256 public beginSeq = 10;
    uint256 public version = 0;

    bytes[] public sequencerBLSKeys;
    address[] public sequencerAddrs;

    uint256 public constant SEQUENCER_SIZE = 3;

    uint256 public NEXT_EPOCH_START = 1700000000;

    uint256 public REWARD_START_TIME = 86400;

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

    uint256 public PROPOSAL_INTERVAL = 1000;
    uint256 public ROLLUP_EPOCH = 1000;
    uint256 public MAX_CHUNKS = 1000000000;

    function setUp() public virtual override {
        super.setUp();
        // Set the proxy at the correct address
        hevm.etch(
            Predeploys.SEQUENCER,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.GOV,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_STAKING,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.MORPH_TOKEN,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.DISTRIBUTE,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.RECORD,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        TransparentUpgradeableProxy sequencerProxy = TransparentUpgradeableProxy(
                payable(Predeploys.SEQUENCER)
            );
        TransparentUpgradeableProxy govProxy = TransparentUpgradeableProxy(
            payable(Predeploys.GOV)
        );
        TransparentUpgradeableProxy l2StakingProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_STAKING)
            );
        TransparentUpgradeableProxy morphTokenProxy = TransparentUpgradeableProxy(
                payable(Predeploys.MORPH_TOKEN)
            );
        TransparentUpgradeableProxy distributeProxy = TransparentUpgradeableProxy(
                payable(Predeploys.DISTRIBUTE)
            );
        TransparentUpgradeableProxy recordProxy = TransparentUpgradeableProxy(
            payable(Predeploys.RECORD)
        );
        hevm.store(
            address(sequencerProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(govProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2StakingProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(morphTokenProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(distributeProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(recordProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );

        hevm.startPrank(multisig);
        // deploy impl contracts
        Sequencer sequencerImpl = new Sequencer();

        Gov govImpl = new Gov();

        L2Staking l2StakingImpl = new L2Staking(payable(NON_ZERO_ADDRESS));

        MorphToken morphTokenImpl = new MorphToken();

        Distribute distributeImpl = new Distribute();

        Record recordImpl = new Record();

        // upgrade proxy
        Types.StakerInfo[] memory stakerInfos = new Types.StakerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakingInfo(user);
            stakerInfos[i] = stakerInfo;
            sequencerAddrs.push(stakerInfo.addr);
        }
        ITransparentUpgradeableProxy(address(sequencerProxy)).upgradeToAndCall(
            address(sequencerImpl),
            abi.encodeWithSelector(
                Sequencer.initialize.selector,
                sequencerAddrs
            )
        );
        ITransparentUpgradeableProxy(address(govProxy)).upgradeToAndCall(
            address(govImpl),
            abi.encodeWithSelector(
                Gov.initialize.selector,
                PROPOSAL_INTERVAL, // _proposalInterval
                0, // _batchBlockInterval
                0, // _batchMaxBytes
                FINALIZATION_PERIOD_SECONDS, // _batchTimeout
                ROLLUP_EPOCH, // rollupEpoch
                MAX_CHUNKS // maxChunks
            )
        );

        ITransparentUpgradeableProxy(address(l2StakingProxy)).upgradeToAndCall(
            address(l2StakingImpl),
            abi.encodeWithSelector(
                L2Staking.initialize.selector,
                multisig,
                SEQUENCER_SIZE * 2,
                ROLLUP_EPOCH,
                REWARD_START_TIME,
                stakerInfos
            )
        );

        ITransparentUpgradeableProxy(address(morphTokenProxy)).upgradeToAndCall(
                address(morphTokenImpl),
                abi.encodeWithSelector(
                    MorphToken.initialize.selector,
                    "Morph",
                    "MPH",
                    address(distributeProxy),
                    1000000000 ether,
                    1596535874529,
                    REWARD_START_TIME
                )
            );
        ITransparentUpgradeableProxy(address(distributeProxy)).upgradeToAndCall(
                address(distributeImpl),
                abi.encodeWithSelector(
                    Distribute.initialize.selector,
                    address(morphTokenProxy),
                    address(recordProxy),
                    address(l2StakingProxy)
                )
            );
        ITransparentUpgradeableProxy(address(recordProxy)).upgradeToAndCall(
            address(recordImpl),
            abi.encodeWithSelector(
                Record.initialize.selector,
                multisig,
                address(1)
            )
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
