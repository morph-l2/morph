// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title Predeploys
 * @notice Contains constant addresses for contracts that are pre-deployed to the L2 system.
 */
library Predeploys {
    /**
     * @notice Address of the L2ToL1MessagePasser predeploy.
     */
    address internal constant L2_TO_L1_MESSAGE_PASSER =
        0x5300000000000000000000000000000000000001;

    /**
     * @notice Address of the L2GatewayRouter predeploy.
     */
    address internal constant L2_Gateway_Router =
        0x5300000000000000000000000000000000000002;

    /**
     * @notice Address of the L2Sequencer predeploy.
     */
    address internal constant L2_SEQUENCER =
        0x5300000000000000000000000000000000000003;

    /**
     * @notice Address of the Gov predeploy.
     */
    address internal constant L2_GOV =
        0x5300000000000000000000000000000000000004;

    /**
     * @notice Address of the L2SUBMITTER predeploy.
     */
    address internal constant L2_SUBMITTER =
        0x5300000000000000000000000000000000000005;

    /**
     * @notice Address of the L2ETHGateWay predeploy.
     */
    address internal constant L2_ETH_Gateway =
        0x5300000000000000000000000000000000000006;

    /**
     * @notice Address of the L2CrossDomainMessenger predeploy.
     */
    address internal constant L2_CROSS_DOMAIN_MESSENGER =
        0x5300000000000000000000000000000000000007;

    /**
     * @notice Address of the L2StandardETHGateWay predeploy.
     */
    address internal constant L2_Standard_ERC20_Gateway =
        0x5300000000000000000000000000000000000008;

    /**
     * @notice Address of the L2ERC721Bridge predeploy.
     */
    address internal constant L2_ERC721_Gateway =
        0x5300000000000000000000000000000000000009;

    /**
     * @notice Address of the L2ERC721Bridge predeploy.
     */
    address internal constant L2TxFeeVault =
        0x530000000000000000000000000000000000001a;

    /**
     * @notice Address of the ProxyAdmin predeploy.
     */
    address internal constant PROXY_ADMIN =
        0x530000000000000000000000000000000000001B;

    /**
     * @notice Address of the GasPriceOracle predeploy. Includes fee information
     *         and helpers for computing the L1 portion of the transaction fee.
     */
    address internal constant GAS_PRICE_ORACLE =
        0x530000000000000000000000000000000000000f;
}
