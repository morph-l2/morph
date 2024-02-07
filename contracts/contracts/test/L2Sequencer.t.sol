// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract L2SequencerTest is L2StakingBaseTest {
    mapping(uint256 => address[]) public sequencersAddr;
    mapping(uint256 => mapping(address => Types.SequencerInfo))
        public sequencers;

    function testUpdateSequencers() external {
        hevm.mockCall(
            address(l2Sequencer.messenger()),
            abi.encodeWithSelector(
                ICrossDomainMessenger.xDomainMessageSender.selector
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
        hevm.prank(address(l2CrossDomainMessenger));
        // updateSequencers
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(l2Sequencer.sequencerAddresses(i), sequencerInfos[i].addr);

            (address user, bytes32 tmKey, bytes memory blsKey) = l2Sequencer
                .sequencerInfos(i);
            assertEq(user, sequencerInfos[i].addr);
            assertEq(tmKey, sequencerInfos[i].tmKey);
            assertBytesEq(blsKey, sequencerInfos[i].blsKey);
        }
    }
}
