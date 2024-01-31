// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {L1CrossDomainMessenger} from "../L1/L1CrossDomainMessenger.sol";
import {L1ChugSplashProxy} from "../legacy/L1ChugSplashProxy.sol";
import {AddressManager} from "./AddressManager.sol";
import {Proxy} from "../libraries/proxy/Proxy.sol";
import {ProxyAdmin} from "../libraries/proxy/ProxyAdmin.sol";
import {Constants} from "../libraries/constants/Constants.sol";

/**
 * @title SystemDictator
 * @notice The SystemDictator is responsible for coordinating the deployment of a full Bedrock
 *         system. The SystemDictator is designed to support both fresh network deployments and
 *         upgrades to existing pre-Bedrock systems.
 */
contract SystemDictator is OwnableUpgradeable {
    //    /**
    //     * @notice Basic system configuration.
    //     */
    //    struct GlobalConfig {
    //        AddressManager addressManager;
    //        ProxyAdmin proxyAdmin;
    //        address controller;
    //        address finalOwner;
    //    }
    //
    //    /**
    //     * @notice Set of proxy addresses.
    //     */
    //    struct ProxyAddressConfig {
    //        address MorphPortalProxy;
    //        address l1CrossDomainMessengerProxy;
    //        address l1StandardBridgeProxy;
    //        address MorphMintableERC20FactoryProxy;
    //        address l1ERC721BridgeProxy;
    //        address systemConfigProxy;
    //    }
    //
    //    /**
    //     * @notice Set of implementation addresses.
    //     */
    //    struct ImplementationAddressConfig {
    //        L1CrossDomainMessenger l1CrossDomainMessengerImpl;
    //        L1StandardBridge l1StandardBridgeImpl;
    //        MorphMintableERC20Factory MorphMintableERC20FactoryImpl;
    //        L1ERC721Bridge l1ERC721BridgeImpl;
    //    }
    //
    //    /**
    //     * @notice Values for the system config contract.
    //     */
    //    struct SystemConfigConfig {
    //        address owner;
    //        uint256 overhead;
    //        uint256 scalar;
    //        bytes32 batcherHash;
    //        uint64 gasLimit;
    //        address unsafeBlockSigner;
    //    }
    //
    //    /**
    //     * @notice Combined system configuration.
    //     */
    //    struct DeployConfig {
    //        GlobalConfig globalConfig;
    //        ProxyAddressConfig proxyAddressConfig;
    //        ImplementationAddressConfig implementationAddressConfig;
    //    }
    //
    //    /**
    //     * @notice Step after which exit 1 can no longer be used.
    //     */
    //    uint8 public constant EXIT_1_NO_RETURN_STEP = 3;
    //
    //    /**
    //     * @notice Step where proxy ownership is transferred.
    //     */
    //    uint8 public constant PROXY_TRANSFER_STEP = 4;
    //
    //    /**
    //     * @notice System configuration.
    //     */
    //    DeployConfig public config;
    //
    //    /**
    //     * @notice Dynamic configuration for the MorphPortal. Determines
    //     *         if the system should be paused when initialized.
    //     */
    //    bool public MorphPortalDynamicConfig;
    //
    //    /**
    //     * @notice Current step;
    //     */
    //    uint8 public currentStep;
    //
    //    /**
    //     * @notice Whether or not the deployment is finalized.
    //     */
    //    bool public finalized;
    //
    //    /**
    //     * @notice Whether or not the deployment has been exited.
    //     */
    //    bool public exited;
    //
    //    /**
    //     * @notice Address of the old L1CrossDomainMessenger implementation.
    //     */
    //    address public oldL1CrossDomainMessenger;
    //
    //    /**
    //     * @notice Checks that the current step is the expected step, then bumps the current step.
    //     *
    //     * @param _step Current step.
    //     */
    //    modifier step(uint8 _step) {
    //        require(!finalized, "SystemDictator: already finalized");
    //        require(!exited, "SystemDictator: already exited");
    //        require(currentStep == _step, "SystemDictator: incorrect step");
    //        _;
    //        currentStep++;
    //    }
    //
    //    /**
    //     * @notice Constructor required to ensure that the implementation of the SystemDictator is
    //     *         initialized upon deployment.
    //     */
    //    constructor() {
    //        // Using this shorter variable as an alias for address(0) just prevents us from having to
    //        // to use a new line for every single parameter.
    //        address zero = address(0);
    //        initialize(
    //            DeployConfig(
    //                GlobalConfig(
    //                    AddressManager(zero),
    //                    ProxyAdmin(zero),
    //                    zero,
    //                    zero
    //                ),
    //                ProxyAddressConfig(zero, zero, zero, zero, zero, zero),
    //                ImplementationAddressConfig(
    //                    L1CrossDomainMessenger(payable(zero)),
    //                    L1StandardBridge(payable(zero)),
    //                    MorphMintableERC20Factory(zero),
    //                    L1ERC721Bridge(zero)
    //                )
    //            )
    //        );
    //    }
    //
    //    /**
    //     * @param _config System configuration.
    //     */
    //    function initialize(DeployConfig memory _config) public initializer {
    //        config = _config;
    //        currentStep = 1;
    //        __Ownable_init();
    //        _transferOwnership(config.globalConfig.controller);
    //    }
    //
    //    /**
    //     * @notice Configures the ProxyAdmin contract.
    //     */
    //    function step1() public onlyOwner step(1) {
    //        // Set the AddressManager in the ProxyAdmin.
    //        config.globalConfig.proxyAdmin.setAddressManager(
    //            config.globalConfig.addressManager
    //        );
    //
    //        // Set the L1CrossDomainMessenger to the RESOLVED proxy type.
    //        config.globalConfig.proxyAdmin.setProxyType(
    //            config.proxyAddressConfig.l1CrossDomainMessengerProxy,
    //            ProxyAdmin.ProxyType.RESOLVED
    //        );
    //
    //        // Set the implementation name for the L1CrossDomainMessenger.
    //        config.globalConfig.proxyAdmin.setImplementationName(
    //            config.proxyAddressConfig.l1CrossDomainMessengerProxy,
    //            "L1CrossDomainMessenger"
    //        );
    //
    //        // Set the L1StandardBridge to the CHUGSPLASH proxy type.
    //        config.globalConfig.proxyAdmin.setProxyType(
    //            config.proxyAddressConfig.l1StandardBridgeProxy,
    //            ProxyAdmin.ProxyType.CHUGSPLASH
    //        );
    //    }
    //
    //    /**
    //     * @notice Pauses the system by shutting down the L1CrossDomainMessenger and setting the
    //     *         deposit halt flag to tell the Sequencer's DTL to stop accepting deposits.
    //     */
    //    function step2() public onlyOwner step(2) {
    //        // Store the address of the old L1CrossDomainMessenger implementation. We will need this
    //        // address in the case that we have to exit early.
    //        oldL1CrossDomainMessenger = config
    //            .globalConfig
    //            .addressManager
    //            .getAddress("L1CrossDomainMessenger");
    //
    //        // Temporarily brick the L1CrossDomainMessenger by setting its implementation address to
    //        // address(0) which will cause the ResolvedDelegateProxy to revert. Better than pausing
    //        // the L1CrossDomainMessenger via pause() because it can be easily reverted.
    //        config.globalConfig.addressManager.setAddress(
    //            "L1CrossDomainMessenger",
    //            address(0)
    //        );
    //
    //        // Set the DTL shutoff block, which will tell the DTL to stop syncing new deposits from the
    //        // CanonicalTransactionChain. We do this by setting an address in the AddressManager
    //        // because the DTL already has a reference to the AddressManager and this way we don't also
    //        // need to give it a reference to the SystemDictator.
    //        config.globalConfig.addressManager.setAddress(
    //            "DTL_SHUTOFF_BLOCK",
    //            address(uint160(block.number))
    //        );
    //    }
    //
    //    /**
    //     * @notice Removes deprecated addresses from the AddressManager.
    //     */
    //    function step3() public onlyOwner step(EXIT_1_NO_RETURN_STEP) {
    //        // Remove all deprecated addresses from the AddressManager
    //        string[17] memory deprecated = [
    //            "CanonicalTransactionChain",
    //            "L2CrossDomainMessenger",
    //            "DecompressionPrecompileAddress",
    //            "Sequencer",
    //            "Proposer",
    //            "ChainStorageContainer-CTC-batches",
    //            "ChainStorageContainer-CTC-queue",
    //            "CanonicalTransactionChain",
    //            "StateCommitmentChain",
    //            "BondManager",
    //            "ExecutionManager",
    //            "FraudVerifier",
    //            "StateManagerFactory",
    //            "StateTransitionerFactory",
    //            "SafetyChecker",
    //            "L1MultiMessageRelayer",
    //            "BondManager"
    //        ];
    //
    //        for (uint256 i = 0; i < deprecated.length; i++) {
    //            config.globalConfig.addressManager.setAddress(
    //                deprecated[i],
    //                address(0)
    //            );
    //        }
    //    }
    //
    //    /**
    //     * @notice Transfers system ownership to the ProxyAdmin.
    //     */
    //    function step4() public onlyOwner step(PROXY_TRANSFER_STEP) {
    //        // Transfer ownership of the AddressManager to the ProxyAdmin.
    //        config.globalConfig.addressManager.transferOwnership(
    //            address(config.globalConfig.proxyAdmin)
    //        );
    //
    //        // Transfer ownership of the L1StandardBridge to the ProxyAdmin.
    //        L1ChugSplashProxy(
    //            payable(config.proxyAddressConfig.l1StandardBridgeProxy)
    //        ).setOwner(address(config.globalConfig.proxyAdmin));
    //
    //        // Transfer ownership of the L1ERC721Bridge to the ProxyAdmin.
    //        Proxy(payable(config.proxyAddressConfig.l1ERC721BridgeProxy))
    //            .changeAdmin(address(config.globalConfig.proxyAdmin));
    //    }
    //
    //    /**
    //     * @notice Upgrades and initializes proxy contracts.
    //     */
    //    function step5() public onlyOwner step(5) {
    //        // Upgrade the L1CrossDomainMessenger.
    //        config.globalConfig.proxyAdmin.upgrade(
    //            payable(config.proxyAddressConfig.l1CrossDomainMessengerProxy),
    //            address(
    //                config.implementationAddressConfig.l1CrossDomainMessengerImpl
    //            )
    //        );
    //
    //        // todo L1CrossDomainMessenger initialize
    //        // Try to initialize the L1CrossDomainMessenger, only fail if it's already been initialized.
    //        // try
    //        //     L1CrossDomainMessenger(
    //        //         payable(config.proxyAddressConfig.l1CrossDomainMessengerProxy)
    //        //     ).initialize()
    //        // {
    //        //     // L1CrossDomainMessenger is the one annoying edge case difference between existing
    //        //     // networks and fresh networks because in existing networks it'll already be
    //        //     // initialized but in fresh networks it won't be. Try/catch is the easiest and most
    //        //     // consistent way to handle this because initialized() is not exposed publicly.
    //        // } catch Error(string memory reason) {
    //        //     require(
    //        //         keccak256(abi.encodePacked(reason)) ==
    //        //             keccak256("Initializable: contract is already initialized"),
    //        //         string.concat(
    //        //             "SystemDictator: unexpected error initializing L1XDM: ",
    //        //             reason
    //        //         )
    //        //     );
    //        // } catch {
    //        //     revert(
    //        //         "SystemDictator: unexpected error initializing L1XDM (no reason)"
    //        //     );
    //        // }
    //
    //        // Upgrade the L1StandardBridge (no initializer).
    //        config.globalConfig.proxyAdmin.upgrade(
    //            payable(config.proxyAddressConfig.l1StandardBridgeProxy),
    //            address(config.implementationAddressConfig.l1StandardBridgeImpl)
    //        );
    //
    //        // Upgrade the MorphMintableERC20Factory (no initializer).
    //        config.globalConfig.proxyAdmin.upgrade(
    //            payable(config.proxyAddressConfig.MorphMintableERC20FactoryProxy),
    //            address(
    //                config.implementationAddressConfig.MorphMintableERC20FactoryImpl
    //            )
    //        );
    //
    //        // Upgrade the L1ERC721Bridge (no initializer).
    //        config.globalConfig.proxyAdmin.upgrade(
    //            payable(config.proxyAddressConfig.l1ERC721BridgeProxy),
    //            address(config.implementationAddressConfig.l1ERC721BridgeImpl)
    //        );
    //    }
    //
    //    /**
    //     * @notice Calls the first 2 steps of the migration process.
    //     */
    //    function phase1() external onlyOwner {
    //        step1();
    //        step2();
    //    }
    //
    //    /**
    //     * @notice Calls the remaining steps of the migration process, and finalizes.
    //     */
    //    function phase2() external onlyOwner {
    //        step3();
    //        step4();
    //        step5();
    //        finalize();
    //    }
    //
    //    /**
    //     * @notice Tranfers admin ownership to the final owner.
    //     */
    //    function finalize() public onlyOwner {
    //        // Transfer ownership of the ProxyAdmin to the final owner.
    //        config.globalConfig.proxyAdmin.transferOwnership(
    //            config.globalConfig.finalOwner
    //        );
    //
    //        // Optionally also transfer AddressManager and L1StandardBridge if we still own it. Might
    //        // happen if we're exiting early.
    //        if (currentStep <= PROXY_TRANSFER_STEP) {
    //            // Transfer ownership of the AddressManager to the final owner.
    //            config.globalConfig.addressManager.transferOwnership(
    //                address(config.globalConfig.finalOwner)
    //            );
    //
    //            // Transfer ownership of the L1StandardBridge to the final owner.
    //            L1ChugSplashProxy(
    //                payable(config.proxyAddressConfig.l1StandardBridgeProxy)
    //            ).setOwner(address(config.globalConfig.finalOwner));
    //
    //            // Transfer ownership of the L1ERC721Bridge to the final owner.
    //            Proxy(payable(config.proxyAddressConfig.l1ERC721BridgeProxy))
    //                .changeAdmin(address(config.globalConfig.finalOwner));
    //        }
    //
    //        // Mark the deployment as finalized.
    //        finalized = true;
    //    }
    //
    //    /**
    //     * @notice First exit point, can only be called before step 3 is executed.
    //     */
    //    function exit1() external onlyOwner {
    //        require(
    //            currentStep == EXIT_1_NO_RETURN_STEP,
    //            "SystemDictator: can only exit1 before step 3 is executed"
    //        );
    //
    //        // Reset the L1CrossDomainMessenger to the old implementation.
    //        config.globalConfig.addressManager.setAddress(
    //            "L1CrossDomainMessenger",
    //            oldL1CrossDomainMessenger
    //        );
    //
    //        // Unset the DTL shutoff block which will allow the DTL to sync again.
    //        config.globalConfig.addressManager.setAddress(
    //            "DTL_SHUTOFF_BLOCK",
    //            address(0)
    //        );
    //
    //        // Mark the deployment as exited.
    //        exited = true;
    //    }
}
