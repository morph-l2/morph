// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/console2.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Types} from "../../libraries/common/Types.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Sequencer} from "../../L2/staking/Sequencer.sol";
import {L2Staking} from "../../L2/staking/L2Staking.sol";
import {Gov} from "../../L2/staking/Gov.sol";
import {L2MessageBaseTest} from "./L2MessageBase.t.sol";

contract L2StakingBaseTest is L2MessageBaseTest {
    uint256 public beginSeq = 10;
    uint256 public version = 0;
    bytes[] public sequencerBLSKeys;
    address[] public sequencerAddrs;

    // Sequencer config
    Sequencer public sequencer;

    uint256 public constant SEQUENCER_SIZE = 3;

    uint256 public NEXT_EPOCH_START = 1700000000;

    // Gov config
    Gov public l2Gov;

    // L2Staking instance
    L2Staking public l2Staking;

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
        TransparentUpgradeableProxy sequencerProxy = TransparentUpgradeableProxy(
                payable(Predeploys.SEQUENCER)
            );
        TransparentUpgradeableProxy l2GovProxy = TransparentUpgradeableProxy(
            payable(Predeploys.GOV)
        );
        TransparentUpgradeableProxy l2StakingProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_STAKING)
            );
        hevm.store(
            address(sequencerProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2GovProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2StakingProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );

        hevm.startPrank(multisig);
        // deploy impl contracts
        Sequencer sequencerImpl = new Sequencer();

        Gov govImpl = new Gov();

        L2Staking l2StakingImpl = new L2Staking(payable(NON_ZERO_ADDRESS));

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
        ITransparentUpgradeableProxy(address(l2GovProxy)).upgradeToAndCall(
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
                SEQUENCER_SIZE,
                ROLLUP_EPOCH
            )
        );

        // set address
        sequencer = Sequencer(payable(address(sequencerProxy)));
        l2Gov = Gov(payable(address(l2GovProxy)));
        l2Staking = L2Staking(payable(address(l2StakingProxy)));

        _changeAdmin(address(sequencer));
        _changeAdmin(address(l2Gov));
        _changeAdmin(address(l2Staking));

        hevm.stopPrank();
    }
}
