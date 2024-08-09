// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title Predeploys
 * @notice Contains constant addresses for contracts that are pre-deployed to the L2 system.
 */
library Predeploys {
    /**
     * @notice Address of the L2_TO_L1_MESSAGE_PASSER predeploy.
     */
    address internal constant L2_TO_L1_MESSAGE_PASSER = 0x5300000000000000000000000000000000000001;

    /**
     * @notice Address of the L2_GATEWAY_ROUTER predeploy.
     */
    address internal constant L2_GATEWAY_ROUTER = 0x5300000000000000000000000000000000000002;

    /**
     * @notice Address of the Gov predeploy.
     */
    address internal constant GOV = 0x5300000000000000000000000000000000000004;

    /**
     * @notice Address of the L2_ETH_GATEWAY predeploy.
     */
    address internal constant L2_ETH_GATEWAY = 0x5300000000000000000000000000000000000006;

    /**
     * @notice Address of the L2_CROSS_DOMAIN_MESSENGER predeploy.
     */
    address internal constant L2_CROSS_DOMAIN_MESSENGER = 0x5300000000000000000000000000000000000007;

    /**
     * @notice Address of the L2_STANDARD_ERC20_GATEWAY predeploy.
     */
    address internal constant L2_STANDARD_ERC20_GATEWAY = 0x5300000000000000000000000000000000000008;

    /**
     * @notice Address of the L2_ERC721_GATEWAY predeploy.
     */
    address internal constant L2_ERC721_GATEWAY = 0x5300000000000000000000000000000000000009;

    /**
     * @notice Address of the L2_TX_FEE_VAULT predeploy.
     */
    address internal constant L2_TX_FEE_VAULT = 0x530000000000000000000000000000000000000a;

    /**
     * @notice Address of the PROXY_ADMIN predeploy.
     */
    address internal constant PROXY_ADMIN = 0x530000000000000000000000000000000000000b;

    /**
     * @notice Address of the L2_ERC1155_GATEWAY predeploy.
     */
    address internal constant L2_ERC1155_GATEWAY = 0x530000000000000000000000000000000000000c;

    /**
     * @notice Address of the MORPH_STANDARD_ERC20 predeploy.
     */
    address internal constant MORPH_STANDARD_ERC20 = 0x530000000000000000000000000000000000000D;

    /**
     * @notice Address of the MORPH_STANDARD_ERC20_FACTORY predeploy.
     */
    address internal constant MORPH_STANDARD_ERC20_FACTORY = 0x530000000000000000000000000000000000000e;

    /**
     * @notice Address of the GAS_PRICE_ORACLE predeploy. Includes fee information
     *         and helpers for computing the L1 portion of the transaction fee.
     */
    address internal constant GAS_PRICE_ORACLE = 0x530000000000000000000000000000000000000f;

    /**
     * @notice Address of the L2_WETH_GATEWAY predeploy.
     */
    address internal constant L2_WETH_GATEWAY = 0x5300000000000000000000000000000000000010;

    /**
     * @notice Address of the L2_WETH predeploy.
     */
    address internal constant L2_WETH = 0x5300000000000000000000000000000000000011;

    /**
     * @notice Address of the RECORD predeploy.
     */
    address internal constant RECORD = 0x5300000000000000000000000000000000000012;

    /**
     * @notice Address of the MORPH_TOKEN predeploy.
     */
    address internal constant MORPH_TOKEN = 0x5300000000000000000000000000000000000013;

    /**
     * @notice Address of the DISTRIBUTE predeploy.
     */
    address internal constant DISTRIBUTE = 0x5300000000000000000000000000000000000014;

    /**
     * @notice Address of the L2_STAKING predeploy.
     */
    address internal constant L2_STAKING = 0x5300000000000000000000000000000000000015;

    /**
     * @notice Address of the L2_CUSTOM_ERC20_GATEWAY predeploy.
     */
    address internal constant L2_CUSTOM_ERC20_GATEWAY = 0x5300000000000000000000000000000000000016;

    /**
     * @notice Address of the SEQUENCER predeploy.
     */
    address internal constant SEQUENCER = 0x5300000000000000000000000000000000000017;

    /**
     * @notice Address of the L2_REVERSE_ERC20_GATEWAY predeploy.
     */
    address internal constant L2_REVERSE_ERC20_GATEWAY = 0x5300000000000000000000000000000000000018;
}
