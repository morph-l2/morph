// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Script} from "forge-std/Script.sol";
import {VmSafe} from "forge-std/Vm.sol";

/// @notice Reads a deploy-config JSON file and provides typed accessors.
///
/// Expected JSON schema (see script/config/l1.json for the devnet default):
///
///   string  contractAdmin
///   string  l1FeeVaultRecipient
///   string  l1WETHAddress         (empty = skip deploy)
///   string  l1USDCAddress         (empty = skip deploy)
///   uint    l1ChainID, l2ChainID
///   uint    l1MessageQueueMaxGasLimit
///   string  l2BaseFee             (decimal string, in Gwei)
///   bytes32 programVkey
///   uint    finalizationPeriodSeconds, rollupProofWindow, proofRewardPercent, rollupDelayPeriod
///   address rollupProposer, rollupChallenger
///   bytes   batchHeader
///   uint    stakingCrossChainGaslimitAdd, stakingCrossChainGaslimitRemove
///   uint    stakingLockNumber, stakingChallengeDeposit, stakingMinDeposit, stakingChallengerRewardPercentage
///   address[] l2SequencerAddresses
///   bytes32[] l2SequencerTmKeys
///   bytes[]   l2SequencerBlsKeys
library DeployConfig {
    using DeployConfig for DeployConfigCtx;

    struct DeployConfigCtx {
        VmSafe vm;
        string json;
    }

    /// @notice Load a deploy config from a JSON file path.
    function load(VmSafe vm, string memory path) internal returns (DeployConfigCtx memory ctx) {
        string memory raw = vm.readFile(path);
        ctx.vm = vm;
        ctx.json = raw;
    }

    // --- simple scalars ---

    function contractAdmin(DeployConfigCtx memory ctx) internal returns (address) {
        return ctx.vm.parseJsonAddress(ctx.json, ".contractAdmin");
    }

    function l1FeeVaultRecipient(DeployConfigCtx memory ctx) internal returns (address) {
        return ctx.vm.parseJsonAddress(ctx.json, ".l1FeeVaultRecipient");
    }

    function l1WETHAddress(DeployConfigCtx memory ctx) internal returns (string memory) {
        return ctx.vm.parseJsonString(ctx.json, ".l1WETHAddress");
    }

    function l1USDCAddress(DeployConfigCtx memory ctx) internal returns (string memory) {
        return ctx.vm.parseJsonString(ctx.json, ".l1USDCAddress");
    }

    function l1ChainID(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".l1ChainID");
    }

    function l2ChainID(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".l2ChainID");
    }

    function l1MessageQueueMaxGasLimit(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".l1MessageQueueMaxGasLimit");
    }

    function programVkey(DeployConfigCtx memory ctx) internal returns (bytes32) {
        return ctx.vm.parseJsonBytes32(ctx.json, ".programVkey");
    }

    function finalizationPeriodSeconds(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".finalizationPeriodSeconds");
    }

    function rollupProofWindow(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".rollupProofWindow");
    }

    function proofRewardPercent(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".proofRewardPercent");
    }

    function rollupDelayPeriod(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".rollupDelayPeriod");
    }

    function rollupProposer(DeployConfigCtx memory ctx) internal returns (address) {
        return ctx.vm.parseJsonAddress(ctx.json, ".rollupProposer");
    }

    function rollupChallenger(DeployConfigCtx memory ctx) internal returns (address) {
        return ctx.vm.parseJsonAddress(ctx.json, ".rollupChallenger");
    }

    function batchHeader(DeployConfigCtx memory ctx) internal returns (bytes memory) {
        return ctx.vm.parseJsonBytes(ctx.json, ".batchHeader");
    }

    // --- staking params ---

    function stakingCrossChainGaslimitAdd(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingCrossChainGaslimitAdd");
    }

    function stakingCrossChainGaslimitRemove(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingCrossChainGaslimitRemove");
    }

    function stakingLockNumber(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingLockNumber");
    }

    /// @notice staking challenge deposit in ETH (not wei) — multiply by 1 ether in caller.
    function stakingChallengeDeposit(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingChallengeDeposit");
    }

    /// @notice staking min deposit in ETH (not wei) — multiply by 1 ether in caller.
    function stakingMinDeposit(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingMinDeposit");
    }

    function stakingChallengerRewardPercentage(DeployConfigCtx memory ctx) internal returns (uint256) {
        return ctx.vm.parseJsonUint(ctx.json, ".stakingChallengerRewardPercentage");
    }

    // --- l2BaseFee (stored as decimal string in Gwei) ---

    /// @notice Returns l2BaseFee as a uint256 in wei. The JSON stores it as a decimal
    ///         string in Gwei (e.g. "0.1"). We parse it as a fixed-point number:
    ///         value * 1e9.
    function l2BaseFeeWei(DeployConfigCtx memory ctx) internal returns (uint256) {
        string memory s = ctx.vm.parseJsonString(ctx.json, ".l2BaseFee");
        // Parse "0.1" -> 1e8 wei.  This is only safe for simple decimals.
        return ctx.vm.parseUint(s) * 1 gwei;
    }

    // --- arrays ---

    function l2SequencerAddresses(DeployConfigCtx memory ctx) internal returns (address[] memory) {
        string[] memory raw = ctx.vm.parseJsonStringArray(ctx.json, ".l2SequencerAddresses");
        address[] memory out = new address[](raw.length);
        for (uint256 i = 0; i < raw.length; i++) {
            out[i] = ctx.vm.parseAddress(raw[i]);
        }
        return out;
    }

    function l2SequencerTmKeys(DeployConfigCtx memory ctx) internal returns (bytes32[] memory) {
        return ctx.vm.parseJsonBytes32Array(ctx.json, ".l2SequencerTmKeys");
    }

    function l2SequencerBlsKeys(DeployConfigCtx memory ctx) internal returns (bytes[] memory) {
        return ctx.vm.parseJsonBytesArray(ctx.json, ".l2SequencerBlsKeys");
    }
}
