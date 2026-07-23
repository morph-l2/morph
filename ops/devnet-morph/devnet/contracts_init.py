"""
Phase 2b helper — initializes L1 contracts via cast send.

Replaces Initialize.s.sol because vm.startBroadcast(pk) in forge 1.5.1
does not correctly set msg.sender during simulation.

Each call is: cast send <proxy> "upgradeToAndCall(address,bytes)" <impl> <innerData>

Call after Phase 1 (Deploy.s.sol) and Phase 2a (InitializeImpls.s.sol).
"""

import logging
import subprocess
import os

log = logging.getLogger(__name__)

# L2 predeploy addresses
L2_MESSENGER              = "0x5300000000000000000000000000000000000007"
L2_ETH_GATEWAY            = "0x5300000000000000000000000000000000000006"
L2_STANDARD_ERC20_GATEWAY = "0x5300000000000000000000000000000000000008"
L2_CUSTOM_ERC20_GATEWAY   = "0x5300000000000000000000000000000000000016"
L2_WITHDRAW_LOCK_ERC20    = "0x5300000000000000000000000000000000000019"
L2_REVERSE_ERC20_GATEWAY  = "0x5300000000000000000000000000000000000018"
L2_WETH_GATEWAY           = "0x5300000000000000000000000000000000000010"
L2_ERC721_GATEWAY         = "0x5300000000000000000000000000000000000009"
L2_ERC1155_GATEWAY        = "0x530000000000000000000000000000000000000c"
L2_TOKEN_IMPLEMENTATION   = "0x530000000000000000000000000000000000000d"
L2_TOKEN_FACTORY          = "0x530000000000000000000000000000000000000e"


def _cast_call(contracts_dir, rpc_url, pk, to_addr, sig, *args):
    """Run cast send and raise on failure."""
    cmd = [
        "cast", "send", to_addr, sig, *[str(a) for a in args],
        "--rpc-url", rpc_url,
        "--private-key", pk,
    ]
    result = subprocess.run(cmd, capture_output=True, text=True, cwd=contracts_dir,
                            env={**os.environ})
    output = (result.stdout or "") + "\n" + (result.stderr or "")
    if result.returncode != 0:
        raise RuntimeError(f"cast send failed [{to_addr[:10]}... {sig}]\n{output[-500:]}")
    # Check for on-chain revert
    if "revert" in output.lower() and "status" in output:
        if '"status": 0' in output or '"status":"0x0"' in output:
            raise RuntimeError(f"cast send reverted [{to_addr[:10]}... {sig}]\n{output[-500:]}")
    log.debug(f"  OK: {to_addr[:10]}... {sig[:30]}")
    return output


def _calldata(contracts_dir, sig, *args):
    """Generate ABI-encoded calldata via cast calldata."""
    cmd = ["cast", "calldata", sig, *[str(a) for a in args]]
    result = subprocess.run(cmd, capture_output=True, text=True, cwd=contracts_dir,
                            env={**os.environ})
    if result.returncode != 0:
        raise RuntimeError(f"cast calldata failed [{sig}]\n{result.stderr}")
    return result.stdout.strip()


def initialize_contracts(
    contracts_dir: str,
    rpc_url: str,
    deployer_pk: str,
    config: dict,
    proxies: dict,
    impls: dict,
):
    pk = deployer_pk
    messenger = proxies["Proxy__L1CrossDomainMessenger"]
    msg_queue = proxies["Proxy__L1MessageQueueWithGasPriceOracle"]
    rollup    = proxies["Proxy__Rollup"]
    staking   = proxies["Proxy__L1Staking"]
    sequencer = proxies["Proxy__L1Sequencer"]
    router    = proxies["Proxy__L1GatewayRouter"]
    eth_gw    = proxies["Proxy__L1ETHGateway"]
    proxy_admin = proxies["Impl__ProxyAdmin"]

    # ===================================================================
    # Stage 1: Core contracts
    # ===================================================================
    log.info("Initializing core contracts...")

    # 1. L1CrossDomainMessenger: initialize(address,address,address)
    _cast_call(contracts_dir, rpc_url, pk, messenger,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1CrossDomainMessenger"],
               _calldata(contracts_dir, "initialize(address,address,address)",
                         config["l1FeeVaultRecipient"], rollup, msg_queue))

    # 2. L1MessageQueueWithGasPriceOracle: initialize(uint256,address)
    _cast_call(contracts_dir, rpc_url, pk, msg_queue,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1MessageQueueWithGasPriceOracle"],
               _calldata(contracts_dir, "initialize(uint256,address)",
                         str(config["l1MessageQueueMaxGasLimit"]),
                         impls["Impl__Whitelist"]))

    # 3. Rollup: initialize(address,address,address,uint256,uint256,uint256)
    _cast_call(contracts_dir, rpc_url, pk, rollup,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__Rollup"],
               _calldata(contracts_dir,
                         "initialize(address,address,address,uint256,uint256,uint256)",
                         staking, msg_queue,
                         impls["Impl__MultipleVersionRollupVerifier"],
                         str(config["finalizationPeriodSeconds"]),
                         str(config["rollupProofWindow"]),
                         str(config["proofRewardPercent"])))

    # 4. L1Staking: initialize(address,uint256,uint256,uint256,uint256,uint256,uint256)
    _cast_call(contracts_dir, rpc_url, pk, staking,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1Staking"],
               _calldata(contracts_dir,
                         "initialize(address,uint256,uint256,uint256,uint256,uint256,uint256)",
                         rollup,
                         str(int(config["stakingMinDeposit"] * 10**18)),
                         str(int(config["stakingChallengeDeposit"] * 10**18)),
                         str(config["stakingLockNumber"]),
                         str(config["stakingChallengerRewardPercentage"]),
                         str(config["stakingCrossChainGaslimitAdd"]),
                         str(config["stakingCrossChainGaslimitRemove"])))

    # 5. L1Sequencer: initialize(address)
    _cast_call(contracts_dir, rpc_url, pk, sequencer,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1Sequencer"],
               _calldata(contracts_dir, "initialize(address)",
                         config["contractAdmin"]))

    # 6. L1GatewayRouter: initialize(address,address)
    _cast_call(contracts_dir, rpc_url, pk, router,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1GatewayRouter"],
               _calldata(contracts_dir, "initialize(address,address)",
                         eth_gw, config["contractAdmin"]))

    # 7. L1ETHGateway: initialize(address,address,address)
    _cast_call(contracts_dir, rpc_url, pk, eth_gw,
               "upgradeToAndCall(address,bytes)",
               impls["Impl__L1ETHGateway"],
               _calldata(contracts_dir, "initialize(address,address,address)",
                         L2_ETH_GATEWAY, router, messenger))

    # ===================================================================
    # Stage 2: Gateway family
    # ===================================================================
    log.info("Initializing gateway contracts...")

    # 8. L1StandardERC20Gateway: initialize(address,address,address,address,address)
    _cast_call(contracts_dir, rpc_url, pk, proxies["Proxy__L1StandardERC20Gateway"],
               "upgradeToAndCall(address,bytes)", impls["Impl__L1StandardERC20Gateway"],
               _calldata(contracts_dir,
                         "initialize(address,address,address,address,address)",
                         L2_STANDARD_ERC20_GATEWAY, router, messenger,
                         L2_TOKEN_IMPLEMENTATION, L2_TOKEN_FACTORY))

    # 9. Other ERC20 gateways: initialize(address,address,address)
    erc20_defs = [
        ("Proxy__L1CustomERC20Gateway",       "Impl__L1CustomERC20Gateway",       L2_CUSTOM_ERC20_GATEWAY),
        ("Proxy__L1WithdrawLockERC20Gateway", "Impl__L1WithdrawLockERC20Gateway", L2_WITHDRAW_LOCK_ERC20),
        ("Proxy__L1ReverseCustomGateway",     "Impl__L1ReverseCustomGateway",     L2_REVERSE_ERC20_GATEWAY),
        ("Proxy__L1WETHGateway",              "Impl__L1WETHGateway",              L2_WETH_GATEWAY),
    ]
    for pkey, ikey, counterpart in erc20_defs:
        _cast_call(contracts_dir, rpc_url, pk, proxies[pkey],
                   "upgradeToAndCall(address,bytes)", impls[ikey],
                   _calldata(contracts_dir, "initialize(address,address,address)",
                             counterpart, router, messenger))

    # 8. L1ERC721Gateway: initialize(address,address)
    _cast_call(contracts_dir, rpc_url, pk, proxies["Proxy__L1ERC721Gateway"],
               "upgradeToAndCall(address,bytes)", impls["Impl__L1ERC721Gateway"],
               _calldata(contracts_dir, "initialize(address,address)",
                         L2_ERC721_GATEWAY, messenger))

    # 9. L1ERC1155Gateway
    _cast_call(contracts_dir, rpc_url, pk, proxies["Proxy__L1ERC1155Gateway"],
               "upgradeToAndCall(address,bytes)", impls["Impl__L1ERC1155Gateway"],
               _calldata(contracts_dir, "initialize(address,address)",
                         L2_ERC1155_GATEWAY, messenger))

    # 10. EnforcedTxGateway: initialize(address,address)
    _cast_call(contracts_dir, rpc_url, pk, proxies["Proxy__EnforcedTxGateway"],
               "upgradeToAndCall(address,bytes)", impls["Impl__EnforcedTxGateway"],
               _calldata(contracts_dir, "initialize(address,address)",
                         msg_queue, config["l1FeeVaultRecipient"]))

    # ===================================================================
    # Stage 3: Transfer all proxy admins to ProxyAdmin
    # ===================================================================
    log.info("Transferring admin ownership to ProxyAdmin...")

    all_proxies = [
        "Proxy__L1MessageQueueWithGasPriceOracle",
        "Proxy__Rollup",
        "Proxy__L1Staking",
        "Proxy__L1Sequencer",
        "Proxy__L1GatewayRouter",
        "Proxy__L1ETHGateway",
        "Proxy__L1StandardERC20Gateway",
        "Proxy__L1CustomERC20Gateway",
        "Proxy__L1WithdrawLockERC20Gateway",
        "Proxy__L1ReverseCustomGateway",
        "Proxy__L1WETHGateway",
        "Proxy__L1ERC721Gateway",
        "Proxy__L1ERC1155Gateway",
        "Proxy__EnforcedTxGateway",
        "Proxy__L1CrossDomainMessenger",   # last
    ]
    for pkey in all_proxies:
        _cast_call(contracts_dir, rpc_url, pk, proxies[pkey],
                   "changeAdmin(address)", proxy_admin)

    # ===================================================================
    # Stage 4: Post-init configuration (ContractInit)
    # ===================================================================
    log.info("Running post-init configuration (ContractInit)...")

    # 4a. GasPriceOracle: set L2 base fee (l2BaseFee stored as Gwei string)
    l2_base_fee_wei = str(int(float(config["l2BaseFee"]) * 1e9))
    _cast_call(contracts_dir, rpc_url, pk, msg_queue,
               "setL2BaseFee(uint256)", l2_base_fee_wei)

    # 4b. Whitelist: add L1Staking proxy to whitelist
    _cast_call(contracts_dir, rpc_url, pk, impls["Impl__Whitelist"],
               "updateWhitelistStatus(address[],bool)",
               f"[{staking}]", "true")

    # 4c. Rollup: import genesis batch + add challenger + initialize2/3
    _cast_call(contracts_dir, rpc_url, pk, rollup,
               "importGenesisBatch(bytes)", config["batchHeader"])
    _cast_call(contracts_dir, rpc_url, pk, rollup,
               "addChallenger(address)", config["rollupChallenger"])
    _cast_call(contracts_dir, rpc_url, pk, rollup,
               "initialize2(bytes32)",
               "0x0000000000000000000000000000000000000000000000000000000000000001")
    _cast_call(contracts_dir, rpc_url, pk, rollup,
               "initialize3(uint256)", str(config["rollupDelayPeriod"]))

    # 4d. L1Staking: whitelist all sequencer addresses
    addrs_str = "[" + ",".join(config["l2SequencerAddresses"]) + "]"
    _cast_call(contracts_dir, rpc_url, pk, staking,
               "updateWhitelist(address[],address[])", addrs_str, "[]")

    # 4e. L1GatewayRouter: set ERC20 gateway for WETH
    weth_addr = proxies.get("Impl__WETH", "0x0000000000000000000000000000000000000000")
    if weth_addr != "0x0000000000000000000000000000000000000000":
        _cast_call(contracts_dir, rpc_url, pk, router,
                   "setERC20Gateway(address[],address[])",
                   f"[{weth_addr}]", f"[{proxies['Proxy__L1WETHGateway']}]")

    log.info("Contract initialization complete.")
