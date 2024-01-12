// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Staking} from "../L1/staking/Staking.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {Sequencer} from "../universal/Sequencer.sol";
import {L1Sequencer} from "../L1/staking/L1Sequencer.sol";
import {L2Sequencer} from "../L2/L2Sequencer.sol";
import {Types} from "../libraries/Types.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {MockCall} from "../mock/MockCall.sol";
import {MockCrossMessager} from "../mock/MockCrossMessager.sol";
import "./CommonTest.t.sol";
import "forge-std/console.sol";

contract L2Sequencer_Test is Staking_Initializer {
    mapping(uint256 => address[]) public sequencersAddr;
    mapping(uint256 => mapping(address => Types.SequencerInfo))
        public sequencers;

    function test_updateSequencers() external {
        vm.mockCall(
            address(l2Sequencer.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(l2Sequencer.OTHER_SEQUENCER()))
        );

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerBLSKeys.push(sequencerInfo.blsKey);
            sequencerInfos[i] = sequencerInfo;
        }
        version++;
        vm.prank(address(L2Messenger));
        // updateSequencers
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(l2Sequencer.sequencerAddresses(i), sequencerInfos[i].addr);

            (address user, bytes32 tmKey, bytes memory blsKey) = l2Sequencer
                .sequencerInfos(i);
            assertEq(user, sequencerInfos[i].addr);
            assertEq(tmKey, sequencerInfos[i].tmKey);
            assertEq(blsKey, sequencerInfos[i].blsKey);
        }
    }
}

contract L2Sequencer_CrossChain_Test is Staking_Initializer {
    mapping(uint256 => Types.SequencerInfo) public sequencers;

    function setUp() public virtual override {
        super.setUp();

        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            address user = address(uint160(beginSeq + i));
            vm.deal(user, 3 * MIN_DEPOSIT);
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = user;
            vm.prank(alice);
            staking.updateWhitelist(add, remove);
            vm.expectEmit(true, true, true, true);
            emit Registered(
                user,
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                2 * MIN_DEPOSIT
            );
            vm.prank(user);
            staking.register{value: 2 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defultGasLimit
            );
            (
                address addrCheck,
                bytes32 tmKeyCheck,
                bytes memory blsKeyCheck,
                uint256 balanceCheck
            ) = staking.stakings(user);
            assertEq(addrCheck, user);
            assertEq(tmKeyCheck, sequencerInfo.tmKey);
            assertEq(blsKeyCheck, sequencerInfo.blsKey);
            assertEq(balanceCheck, 2 * MIN_DEPOSIT);
            sequencers[i] = sequencerInfo;
        }
    }

    function test_updateSequencers_cross() external {
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
            alice
        );
        sequencers[SEQUENCER_SIZE - 1] = sequencerInfo;
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            sequencerBLSKeys.push(sequencers[i].blsKey);
        }
        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            sequencerInfos[i] = sequencers[i];
        }

        address[] memory add = new address[](1);
        address[] memory remove;
        add[0] = alice;
        staking.updateWhitelist(add, remove);

        version++;
        bytes memory data = abi.encodeWithSelector(
            L2Sequencer.updateSequencers.selector,
            // Because this call will be executed on the remote chain, we reverse the order of
            // the remote and local token addresses relative to their order in the
            // updateSequencers function.
            version,
            sequencerInfos
        );
        uint256 _nonce = L1Messenger.messageNonce();
        address _sender = AddressAliasHelper.applyL1ToL2Alias(
            address(L1Messenger)
        );

        vm.expectEmit(true, true, true, true);
        emit Registered(
            alice,
            sequencerInfo.tmKey,
            sequencerInfo.blsKey,
            2 * MIN_DEPOSIT
        );

        vm.expectEmit(true, true, true, true);
        emit SentMessage(
            address(l2Sequencer),
            address(l1Sequencer),
            data,
            _nonce,
            defultGasLimit
        );

        vm.expectEmit(true, true, true, true);
        emit SentMessageExtension1(address(l1Sequencer), 0);

        vm.expectEmit(true, true, true, true);
        emit SequencerUpdated(sequencerBLSKeys, version);

        staking.register{value: 2 * MIN_DEPOSIT}(
            sequencerInfo.tmKey,
            sequencerInfo.blsKey,
            defultGasLimit
        );
        vm.stopPrank();

        bytes32 versionedHash = Hashing.hashCrossDomainMessageV1(
            _nonce,
            address(l1Sequencer),
            address(l2Sequencer),
            0,
            defultGasLimit,
            data
        );
        vm.expectEmit(true, true, true, true);
        emit RelayedMessage(versionedHash);
        vm.prank(_sender);
        L2Messenger.relayMessage(
            _nonce,
            address(l1Sequencer),
            address(l2Sequencer),
            0,
            defultGasLimit,
            data
        );
        assertEq(L2Messenger.failedMessages(versionedHash), false);
        assertEq(L2Messenger.successfulMessages(versionedHash), true);
        assertEq(l2Sequencer.currentVersion(), 1);
    }
}
