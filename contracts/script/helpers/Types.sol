// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/// @notice Contract-name constants used by deployment and storage.
///         Mirrors the ProxyStorageName / ImplStorageName / ContractFactoryName
///         enums in the legacy Hardhat src/types.ts.
library Types {
    // --- Proxy storage names ---
    string constant PROXY_L1_CROSS_DOMAIN_MESSENGER    = "Proxy__L1CrossDomainMessenger";
    string constant PROXY_L1_MESSAGE_QUEUE              = "Proxy__L1MessageQueueWithGasPriceOracle";
    string constant PROXY_ROLLUP                        = "Proxy__Rollup";
    string constant PROXY_L1_STAKING                    = "Proxy__L1Staking";
    string constant PROXY_L1_SEQUENCER                  = "Proxy__L1Sequencer";
    string constant PROXY_L1_GATEWAY_ROUTER             = "Proxy__L1GatewayRouter";
    string constant PROXY_L1_STANDARD_ERC20_GATEWAY     = "Proxy__L1StandardERC20Gateway";
    string constant PROXY_L1_CUSTOM_ERC20_GATEWAY       = "Proxy__L1CustomERC20Gateway";
    string constant PROXY_L1_WITHDRAW_LOCK_ERC20_GATEWAY= "Proxy__L1WithdrawLockERC20Gateway";
    string constant PROXY_L1_REVERSE_CUSTOM_GATEWAY     = "Proxy__L1ReverseCustomGateway";
    string constant PROXY_L1_ETH_GATEWAY                = "Proxy__L1ETHGateway";
    string constant PROXY_L1_WETH_GATEWAY               = "Proxy__L1WETHGateway";
    string constant PROXY_L1_USDC_GATEWAY               = "Proxy__L1USDCGateway";
    string constant PROXY_L1_ERC721_GATEWAY             = "Proxy__L1ERC721Gateway";
    string constant PROXY_L1_ERC1155_GATEWAY            = "Proxy__L1ERC1155Gateway";
    string constant PROXY_ENFORCED_TX_GATEWAY           = "Proxy__EnforcedTxGateway";

    // --- Implementation storage names ---
    string constant IMPL_PROXY_ADMIN                     = "Impl__ProxyAdmin";
    string constant IMPL_EMPTY_CONTRACT                  = "Impl__EmptyContract";
    string constant IMPL_WETH                            = "Impl__WETH";
    string constant IMPL_L1_CROSS_DOMAIN_MESSENGER       = "Impl__L1CrossDomainMessenger";
    string constant IMPL_L1_MESSAGE_QUEUE                = "Impl__L1MessageQueueWithGasPriceOracle";
    string constant IMPL_ROLLUP                          = "Impl__Rollup";
    string constant IMPL_ZKEVM_VERIFIER_V1               = "Impl__ZkEvmVerifierV1";
    string constant IMPL_MULTIPLE_VERSION_ROLLUP_VERIFIER= "Impl__MultipleVersionRollupVerifier";
    string constant IMPL_L1_STAKING                      = "Impl__L1Staking";
    string constant IMPL_L1_SEQUENCER                    = "Impl__L1Sequencer";
    string constant IMPL_L1_GATEWAY_ROUTER               = "Impl__L1GatewayRouter";
    string constant IMPL_L1_STANDARD_ERC20_GATEWAY       = "Impl__L1StandardERC20Gateway";
    string constant IMPL_L1_CUSTOM_ERC20_GATEWAY         = "Impl__L1CustomERC20Gateway";
    string constant IMPL_L1_WITHDRAW_LOCK_ERC20_GATEWAY  = "Impl__L1WithdrawLockERC20Gateway";
    string constant IMPL_L1_REVERSE_CUSTOM_GATEWAY       = "Impl__L1ReverseCustomGateway";
    string constant IMPL_L1_ETH_GATEWAY                  = "Impl__L1ETHGateway";
    string constant IMPL_L1_WETH_GATEWAY                 = "Impl__L1WETHGateway";
    string constant IMPL_L1_USDC_GATEWAY                 = "Impl__L1USDCGateway";
    string constant IMPL_L1_ERC721_GATEWAY               = "Impl__L1ERC721Gateway";
    string constant IMPL_L1_ERC1155_GATEWAY              = "Impl__L1ERC1155Gateway";
    string constant IMPL_ENFORCED_TX_GATEWAY             = "Impl__EnforcedTxGateway";
}
