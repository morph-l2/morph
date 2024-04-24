// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {IL2Sequencer} from "../L2/staking/IL2Sequencer.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract L1SequencerTest is L1MessageBaseTest {
    string sendMessage4 = "sendMessage(address,uint256,bytes,uint256)";
    address refundAddress = address(2048);

    function test_updateAndSendSequencerSet() external {
        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerAddresses.push(sequencerInfo.addr);
            sequencerBLSKeys.push(sequencerInfo.blsKey);
            sequencerInfos[i] = sequencerInfo;
        }

        // test sequencer set initialized
        hevm.startPrank(address(staking));
        l1Sequencer.updateAndSendSequencerSet(
            abi.encodeCall(
                IL2Sequencer.updateSequencers,
                (l1Sequencer.newestVersion() + 1, sequencerInfos)
            ),
            sequencerAddresses,
            sequencerBLSKeys,
            defaultGasLimit
        );
        checkSequencers(version);

        // test sequencer set updated
        hevm.expectCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSelector(
                bytes4(keccak256(bytes(sendMessage4))),
                address(l2Sequencer),
                0,
                abi.encodeCall(
                    IL2Sequencer.updateSequencers,
                    (l1Sequencer.newestVersion() + 1, sequencerInfos)
                ),
                defaultGasLimit
            )
        );

        l1Sequencer.updateAndSendSequencerSet(
            abi.encodeCall(
                IL2Sequencer.updateSequencers,
                (l1Sequencer.newestVersion() + 1, sequencerInfos)
            ),
            sequencerAddresses,
            sequencerBLSKeys,
            defaultGasLimit
        );
        version++;
        checkSequencers(version);
        hevm.stopPrank();
    }

    function checkSequencers(uint256 version) internal {
        for (uint256 i = 0; i < sequencerAddresses.length; i++) {
            assertEq(
                sequencerAddresses[i],
                l1Sequencer.sequencerAddresses(version, i)
            );
            assertBytesEq(
                sequencerBLSKeys[i],
                l1Sequencer.sequencerBLSKeys(version, i)
            );
        }
    }
}

contract L1SequencerVerifyTest is L1SequencerTest {
    mapping(uint256 => mapping(uint256 => Types.SequencerInfo)) sequencersInfosStorage;

    function setUp() public virtual override {
        super.setUp();

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            Types.SequencerInfo[]
                memory sequencerInfos = new Types.SequencerInfo[](
                    SEQUENCER_SIZE
                );
            for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
                address user = address(uint160(beginSeq + i));
                Types.SequencerInfo memory sequencerInfo = ffi
                    .generateStakingInfo(user);
                sequencerAddresses.push(sequencerInfo.addr);
                sequencerBLSKeys.push(sequencerInfo.blsKey);
                sequencerInfos[j] = sequencerInfo;
                sequencersInfosStorage[version][j] = sequencerInfo;
            }
            bytes memory data = abi.encodeCall(
                IL2Sequencer.updateSequencers,
                // Because this call will be executed on the remote chain, we reverse the order of
                // the remote and local token addresses relative to their order in the
                // updateSequencers function.
                (l1Sequencer.newestVersion() + 1, sequencerInfos)
            );
            hevm.prank(address(staking));
            l1Sequencer.updateAndSendSequencerSet(
                data,
                sequencerAddresses,
                sequencerBLSKeys,
                defaultGasLimit
            );
            checkSequencers(version);
            delete sequencerAddresses;
            delete sequencerBLSKeys;
            version++;
        }
    }

    function testGetSequencerAddresses() external {
        uint256 newestVersion = l1Sequencer.newestVersion();
        for (uint256 i = 1; i <= newestVersion; i++) {
            for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
                assertEq(
                    l1Sequencer.getSequencerAddresses(i)[j],
                    sequencersInfosStorage[i][j].addr
                );
            }
        }
    }

    function testGetSequencerBLSKeys() external {
        uint256 newestVersion = l1Sequencer.newestVersion();
        for (uint256 i = 1; i <= newestVersion; i++) {
            for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
                assertBytesEq(
                    l1Sequencer.getSequencerBLSKeys(i)[j],
                    sequencersInfosStorage[i][j].blsKey
                );
            }
        }
    }

    function testVerifySignatureNewest() external {
        address[] memory sequencers = new address[](0);
        bytes memory signature = bytes("");
        hevm.startPrank(address(rollup));
        uint256 currentVersion = l1Sequencer.currentVersion();
        uint256 newestVersion = l1Sequencer.newestVersion();
        for (uint256 i = currentVersion; i <= newestVersion; i++) {
            assertTrue(
                l1Sequencer.verifySignature(
                    newestVersion,
                    sequencers,
                    signature,
                    bytes32(0)
                )
            );
        }
        hevm.stopPrank();
    }
}
