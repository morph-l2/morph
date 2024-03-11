// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

// import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
// import {Types} from "../libraries/common/Types.sol";
// import {IL2Sequencer} from "../L2/staking/IL2Sequencer.sol";
// import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

// contract L1SequencerTest is L1MessageBaseTest {
//     string sendMessage4 = "sendMessage(address,uint256,bytes,uint256,address)";
//     address refundAddress = address(2048);

//     function test_updateAndSendSequencerSet() external {
//         Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
//             SEQUENCER_SIZE
//         );

//         for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
//             address user = address(uint160(beginSeq + i));
//             Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
//                 user
//             );
//             sequencerAddrs.push(sequencerInfo.addr);
//             sequencerBLSKeys.push(sequencerInfo.blsKey);
//             sequencerInfos[i] = sequencerInfo;
//         }

//         // test sequencer set initialized
//         hevm.prank(address(staking));
//         l1Sequencer.updateAndSendSequencerSet(
//             abi.encodeWithSelector(
//                 IL2Sequencer.updateSequencers.selector,
//                 sequencerInfos
//             ),
//             sequencerAddrs,
//             sequencerBLSKeys,
//             defaultGasLimit,
//             refundAddress
//         );
//         checkSequencers(version);

//         // test sequencer set updated
//         hevm.expectCall(
//             address(l1CrossDomainMessenger),
//             abi.encodeWithSelector(
//                 bytes4(keccak256(bytes(sendMessage4))),
//                 address(l2Sequencer),
//                 0,
//                 abi.encodeWithSelector(
//                     IL2Sequencer.updateSequencers.selector,
//                     sequencerInfos
//                 ),
//                 defaultGasLimit,
//                 refundAddress
//             )
//         );
//         hevm.prank(address(staking));
//         l1Sequencer.updateAndSendSequencerSet(
//             abi.encodeWithSelector(
//                 IL2Sequencer.updateSequencers.selector,
//                 sequencerInfos
//             ),
//             sequencerAddrs,
//             sequencerBLSKeys,
//             defaultGasLimit,
//             refundAddress
//         );
//         version++;
//         checkSequencers(version);
//     }

//     function checkSequencers(uint256 version) internal {
//         for (uint256 i = 0; i < sequencerAddrs.length; i++) {
//             assertEq(sequencerAddrs[i], l1Sequencer.sequencerAddrs(version, i));
//             assertBytesEq(
//                 sequencerBLSKeys[i],
//                 l1Sequencer.sequencerBLSKeys(version, i)
//             );
//         }
//     }
// }

// contract L1SequencerVerifyTest is L1SequencerTest {
//     mapping(uint256 => mapping(uint256 => Types.SequencerInfo)) sequencersInfosStorage;

//     function setUp() public virtual override {
//         super.setUp();

//         for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
//             Types.SequencerInfo[]
//                 memory sequencerInfos = new Types.SequencerInfo[](
//                     SEQUENCER_SIZE
//                 );
//             for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
//                 address user = address(uint160(beginSeq + i));
//                 Types.SequencerInfo memory sequencerInfo = ffi
//                     .generateStakingInfo(user);
//                 sequencerAddrs.push(sequencerInfo.addr);
//                 sequencerBLSKeys.push(sequencerInfo.blsKey);
//                 sequencerInfos[j] = sequencerInfo;
//                 sequencersInfosStorage[version][j] = sequencerInfo;
//             }
//             bytes memory data = abi.encodeWithSelector(
//                 IL2Sequencer.updateSequencers.selector,
//                 // Because this call will be executed on the remote chain, we reverse the order of
//                 // the remote and local token addresses relative to their order in the
//                 // updateSequencers function.
//                 sequencerInfos
//             );
//             hevm.prank(address(staking));
//             l1Sequencer.updateAndSendSequencerSet(
//                 data,
//                 sequencerAddrs,
//                 sequencerBLSKeys,
//                 defaultGasLimit,
//                 refundAddress
//             );
//             checkSequencers(version);
//             delete sequencerAddrs;
//             delete sequencerBLSKeys;
//             version++;
//         }
//     }

//     function testGetSequencerAddrs() external {
//         uint256 newnestVersion = l1Sequencer.newestVersion();
//         for (uint256 i = 1; i <= newnestVersion; i++) {
//             for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
//                 assertEq(
//                     l1Sequencer.getSequencerAddrs(i)[j],
//                     sequencersInfosStorage[i][j].addr
//                 );
//             }
//         }
//     }

//     function testGetSequencerBLSKeys() external {
//         uint256 newnestVersion = l1Sequencer.newestVersion();
//         for (uint256 i = 1; i <= newnestVersion; i++) {
//             for (uint256 j = 0; j < SEQUENCER_SIZE; j++) {
//                 assertBytesEq(
//                     l1Sequencer.getSequencerBLSKeys(i)[j],
//                     sequencersInfosStorage[i][j].blsKey
//                 );
//             }
//         }
//     }

//     function testVerifySignatureNewnest() external {
//         uint256[] memory indexs = new uint256[](0);
//         bytes memory signature = bytes("");
//         hevm.startPrank(address(rollup));
//         uint256 currentVersion = l1Sequencer.currentVersion();
//         uint256 newnestVersion = l1Sequencer.newestVersion();
//         for (uint256 i = currentVersion; i <= newnestVersion; i++) {
//             assertTrue(
//                 l1Sequencer.verifySignature(newnestVersion, indexs, signature)
//             );
//         }
//         hevm.stopPrank();
//     }
// }
