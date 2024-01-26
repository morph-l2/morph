// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CommonTest} from "./CommonTest.t.sol";
import {Types} from "../../libraries/common/Types.sol";

contract FFITest is CommonTest {
    function testget() external {
        uint64 num = ffi.getTest();
        assertEq(num, 12345);
    }

    function testGetProveWithdrawalTransactionInputs() external {
        bytes32 wdHash = bytes32(uint256(2));
        (bytes32 wdHashRes, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(wdHash);
        assertEq(wdHash, wdHashRes);
        assertTrue(verifyMerkleProof(wdHashRes, wdProof, 0, wdRoot));
    }

    function testGenerateStakingInfo() external {
        for (uint256 i = 0; i < 10; i++) {
            address user = address(uint160(10 + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            assertEq(sequencerInfo.addr, user);
        }
    }
}
