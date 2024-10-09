// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {ZkEvmVerifierV1} from "../libraries/verifier/ZkEvmVerifierV1.sol";

struct ProofFixture {
    bytes proof;
    bytes publicValues;
    bytes32 vkey;
}

contract EvmTest is Test {
    using stdJson for string;

    ZkEvmVerifierV1 public evm;

    // Prepare pi, proof, vkey
    function loadFixture() public view returns (ProofFixture memory) {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/contracts/test/testdata/plonk-fixture.json");
        string memory json = vm.readFile(path);
        bytes memory jsonBytes = json.parseRaw(".");
        return abi.decode(jsonBytes, (ProofFixture));
    }

    // The vkey represents an ELF(sp1) app, it should be a constant when the contract is deployed.
    function setUp() public {
        console.logString("Setting up EvmVerifierTest");
        ProofFixture memory fixture = loadFixture();
        evm = new ZkEvmVerifierV1(fixture.vkey);
    }

    // Prove state success.
    function test_ValidProof() public view {
        ProofFixture memory fixture = loadFixture();
        evm.verifyPlonk(fixture.proof, fixture.publicValues);
    }

    // Prove state fail.
    function testFail_InValidProof() public view {
        ProofFixture memory fixture = loadFixture();
        // Create a fake proof.
        fixture.proof[31] = 0x00;
        evm.verifyPlonk(fixture.proof, fixture.publicValues);
    }
}