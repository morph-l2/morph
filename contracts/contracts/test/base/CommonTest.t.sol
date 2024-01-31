// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

/* Testing utilities */
import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import {Test} from "forge-std/Test.sol";

import {MockTree} from "../../mock/MockTree.sol";
import {Types} from "../../libraries/common/Types.sol";

contract CommonTest is DSTestPlus, MockTree {
    address immutable NON_ZERO_ADDRESS = address(1);

    address alice = address(128);
    address bob = address(256);
    address multisig = address(512);

    FFIInterface ffi;

    bytes32 PROXY_OWNER_KEY =
        0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
    uint256 public FINALIZATION_PERIOD_SECONDS = 2;

    function setUp() public virtual {
        // Give alice and bob some ETH
        hevm.deal(alice, 1 << 16);
        hevm.deal(bob, 1 << 16);
        hevm.deal(multisig, 1 << 16);

        hevm.label(alice, "alice");
        hevm.label(bob, "bob");
        hevm.label(multisig, "multisig");

        // Make sure we have a non-zero base fee
        hevm.fee(1000000000);

        ffi = new FFIInterface();
    }

    function _encodeXDomainCalldata(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _messageNonce,
        bytes memory _message
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "relayMessage(address,address,uint256,uint256,bytes)",
                _sender,
                _target,
                _value,
                _messageNonce,
                _message
            );
    }
}

contract FFIInterface is Test {
    function getTest() external returns (uint64) {
        string[] memory cmds = new string[](2);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "getTest";

        bytes memory result = vm.ffi(cmds);
        uint64 num = abi.decode(result, (uint64));
        return num;
    }

    function getProveWithdrawalTransactionInputs(
        bytes32 withdrawalHash
    ) external returns (bytes32, bytes32[32] memory, bytes32) {
        string[] memory cmds = new string[](3);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "getProveWithdrawalTransactionInputs";
        cmds[2] = vm.toString(withdrawalHash);

        bytes memory result = vm.ffi(cmds);
        (
            bytes32 withdrawalHashRes,
            bytes32[32] memory withdrawalProof,
            bytes32 withdrawalRoot
        ) = abi.decode(result, (bytes32, bytes32[32], bytes32));

        return (withdrawalHashRes, withdrawalProof, withdrawalRoot);
    }

    function generateStakingInfo(
        address _staker
    ) external returns (Types.SequencerInfo memory) {
        string[] memory cmds = new string[](3);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "generateStakingInfo";
        cmds[2] = vm.toString(_staker);

        bytes memory result = vm.ffi(cmds);
        Types.SequencerInfo memory sequencerInfo = abi.decode(
            result,
            (Types.SequencerInfo)
        );
        return sequencerInfo;
    }
}
