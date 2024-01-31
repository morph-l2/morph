// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {L1MessageBaseTest} from "./L1MessageBase.t.sol";
import {Staking} from "../../L1/staking/Staking.sol";
import {L1Sequencer} from "../../L1/staking/L1Sequencer.sol";
import {Proxy} from "../../libraries/proxy/Proxy.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L1StakingBaseTest is L1MessageBaseTest {
    uint256 public beginSeq = 10;
    uint256 public version = 0;
    bytes[] public sequencerBLSKeys;

    // staking config
    Staking staking;

    uint256 public constant SEQUENCER_SIZE = 3;
    uint256 public LOCK = 3;

    event Registered(
        address addr,
        bytes32 tmKey,
        bytes blsKey,
        uint256 balance
    );
    event SequencerUpdated(bytes[] sequencers, uint256 version);

    // L1Sequencer config
    L1Sequencer l1Sequencer;

    address l2Sequencer = address(Predeploys.L2_SEQUENCER);

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);

        // deploy proxys
        Proxy stakingProxy = new Proxy(multisig);
        Proxy l1SequencerProxy = new Proxy(multisig);

        // deploy impls
        Staking stakingImpl = new Staking();
        L1Sequencer l1SequencerImpl = new L1Sequencer(
            payable(l1CrossDomainMessenger)
        );

        // upgrade proxys
        stakingProxy.upgradeToAndCall(
            address(stakingImpl),
            abi.encodeWithSelector(
                Staking.initialize.selector,
                address(alice),
                address(l1SequencerProxy),
                SEQUENCER_SIZE,
                MIN_DEPOSIT,
                LOCK
            )
        );

        l1SequencerProxy.upgradeToAndCall(
            address(l1SequencerImpl),
            abi.encodeWithSelector(
                L1Sequencer.initialize.selector,
                address(stakingProxy),
                address(rollup)
            )
        );
        // contracts address set
        staking = Staking(address(stakingProxy));
        l1Sequencer = L1Sequencer(payable(address(l1SequencerProxy)));
        hevm.stopPrank();
    }
}
