// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Staking} from "../L1/staking/Staking.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {Sequencer} from "../universal/Sequencer.sol";
import {L2Sequencer} from "../L2/L2Sequencer.sol";
import {Types} from "../libraries/Types.sol";
import "forge-std/console.sol";
import "./CommonTest.t.sol";

contract L1Sequencer_Test is Staking_Initializer {
    function test_updateAndSendSequencerSet() external {
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
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(l2Sequencer),
                abi.encodeWithSelector(
                    L2Sequencer.updateSequencers.selector,
                    sequencerInfos
                ),
                defultGasLimit
            )
        );
        vm.prank(address(staking));
        l1Sequencer.updateAndSendSequencerSet(
            abi.encodeWithSelector(
                L2Sequencer.updateSequencers.selector,
                sequencerInfos
            ),
            sequencerBLSKeys,
            defultGasLimit
        );
        version++;
        checkSequencerBLSKeys(version);
    }

    function checkSequencerBLSKeys(uint256 version) internal {
        for (uint256 i = 0; i < sequencerBLSKeys.length; i++) {
            assertEq(
                sequencerBLSKeys[i],
                l1Sequencer.sequencerBLSKeys(version, i)
            );
        }
    }
}

contract L1Sequencer_Verify_Test is Staking_Initializer {
    mapping(uint256 => mapping(uint256 => Types.SequencerInfo)) sequencersInfosStorage;

    function setUp() public virtual override {
        super.setUp();

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            Types.SequencerInfo[]
                memory sequencerInfos = new Types.SequencerInfo[](
                    SEQUENCER_SIZE
                );
            version++;
            for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
                address user = address(uint160(beginSeq + i));
                Types.SequencerInfo memory sequencerInfo = ffi
                    .generateStakingInfo(user);
                sequencerBLSKeys.push(sequencerInfo.blsKey);
                sequencerInfos[j] = sequencerInfo;
                sequencersInfosStorage[version][j] = sequencerInfo;
            }
            bytes memory data = abi.encodeWithSelector(
                L2Sequencer.updateSequencers.selector,
                // Because this call will be executed on the remote chain, we reverse the order of
                // the remote and local token addresses relative to their order in the
                // updateSequencers function.
                sequencerInfos
            );
            vm.prank(address(staking));
            l1Sequencer.updateAndSendSequencerSet(
                data,
                sequencerBLSKeys,
                defultGasLimit
            );
            checkSequencerBLSKeys(version);
            delete sequencerBLSKeys;
        }
    }

    function test_getSequencerBLSKeys() external {
        uint256 newnestVersion = l1Sequencer.newestVersion();
        for (uint256 i = 1; i <= newnestVersion; i++) {
            for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
                assertEq(
                    l1Sequencer.getSequencerBLSKeys(i, j),
                    sequencersInfosStorage[i][j].blsKey
                );
            }
        }
    }

    function test_verifySignature_onebyone() external {
        uint256[] memory indexs = new uint256[](0);
        bytes memory signature = bytes("");
        vm.startPrank(address(rollup));
        for (uint256 version = 1; version <= SEQUENCER_SIZE; version++) {
            l1Sequencer.verifySignature(version, indexs, signature);
            assertEq(l1Sequencer.getSequencerBLSKeys(version, 0), bytes(""));
        }
        vm.stopPrank();
    }

    function test_verifySignature_newnest() external {
        uint256[] memory indexs = new uint256[](0);
        bytes memory signature = bytes("");
        vm.startPrank(address(rollup));
        uint256 newnestVersion = l1Sequencer.newestVersion();
        l1Sequencer.verifySignature(newnestVersion, indexs, signature);
        for (uint256 i = 1; i <= newnestVersion; i++) {
            assertEq(l1Sequencer.getSequencerBLSKeys(version, 0), bytes(""));
        }
        vm.stopPrank();
    }

    function checkSequencerBLSKeys(uint256 version) internal {
        for (uint256 i = 0; i < sequencerBLSKeys.length; i++) {
            assertEq(
                sequencerBLSKeys[i],
                l1Sequencer.sequencerBLSKeys(version, i)
            );
        }
    }
}
