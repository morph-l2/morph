// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {CommonTest} from "./CommonTest.t.sol";
import {Whitelist} from "../../libraries/common/Whitelist.sol";
import {L1CrossDomainMessenger} from "../../l1/L1CrossDomainMessenger.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {L1Staking} from "../../l1/staking/L1Staking.sol";
import {Rollup} from "../../l1/rollup/Rollup.sol";
import {IRollup} from "../../l1/rollup/IRollup.sol";
import {MockZkEvmVerifier} from "../../mock/MockZkEvmVerifier.sol";

contract L1MessageBaseTest is CommonTest {
    // Staking config
    L1Staking public l1Staking;
    L1Staking public l1StakingImpl;

    uint256 public constant STAKING_VALUE = 1e18; // 1 eth
    uint256 public constant CHALLENGE_DEPOSIT = 1e18; // 1 eth
    uint256 public constant LOCK_BLOCKS = 3;
    uint256 public rewardPercentage = 20;
    uint32 public defaultGasLimitAdd = 1000000;
    uint32 public defaultGasLimitRemove = 10000000;

    // Rollup config
    Rollup public rollup;
    Rollup public rollupImpl;
    MockZkEvmVerifier public verifier = new MockZkEvmVerifier();

    uint256 public proofWindow = 100;
    uint256 public maxNumTxInChunk = 10;
    uint64 public layer2ChainID = 53077;

    // whitelist config
    Whitelist public whitelistChecker;

    // L1MessageQueueWithGasPriceOracle config
    L1MessageQueueWithGasPriceOracle public l1MessageQueueWithGasPriceOracle;
    uint256 public l1MessageQueueMaxGasLimit = 100000000;
    uint32 public defaultGasLimit = 1000000;

    // L1CrossDomainMessenger config
    L1CrossDomainMessenger public l1CrossDomainMessenger;
    L1CrossDomainMessenger public l1CrossDomainMessengerImpl;

    address public l1FeeVault = address(3033);

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // deploy whitelist
        whitelistChecker = new Whitelist(address(multisig));

        // deploy proxy
        TransparentUpgradeableProxy rollupProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1CrossDomainMessengerProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1MessageQueueWithGasPriceOracleProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1StakingProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // deploy impl
        rollupImpl = new Rollup(layer2ChainID);
        L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
            payable(address(l1CrossDomainMessengerProxy)),
            address(rollupProxy),
            address(alice)
        );
        l1CrossDomainMessengerImpl = new L1CrossDomainMessenger();
        l1StakingImpl = new L1Staking(payable(l1CrossDomainMessengerProxy));

        // upgrade and initialize
        ITransparentUpgradeableProxy(address(rollupProxy)).upgradeToAndCall(
            address(rollupImpl),
            abi.encodeCall(
                Rollup.initialize,
                (
                    address(l1StakingProxy),
                    address(l1MessageQueueWithGasPriceOracleProxy), // _messageQueue
                    address(verifier), // _verifier
                    maxNumTxInChunk, // _maxNumTxInChunk
                    finalizationPeriodSeconds, // _finalizationPeriodSeconds
                    proofWindow // _proofWindow
                )
            )
        );
        ITransparentUpgradeableProxy(address(l1MessageQueueWithGasPriceOracleProxy)).upgradeToAndCall(
            address(l1MessageQueueWithGasPriceOracleImpl),
            abi.encodeCall(
                L1MessageQueueWithGasPriceOracle.initialize,
                (
                    l1MessageQueueMaxGasLimit, // gasLimit
                    address(whitelistChecker) // whitelistChecker
                )
            )
        );
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxy)).upgradeToAndCall(
            address(l1CrossDomainMessengerImpl),
            abi.encodeCall(
                L1CrossDomainMessenger.initialize,
                (
                    l1FeeVault, // feeVault
                    address(rollupProxy), // rollup
                    address(l1MessageQueueWithGasPriceOracleProxy) // messageQueue
                )
            )
        );
        ITransparentUpgradeableProxy(address(l1StakingProxy)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(rollupProxy),
                    STAKING_VALUE,
                    CHALLENGE_DEPOSIT,
                    LOCK_BLOCKS,
                    rewardPercentage,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );

        l1CrossDomainMessenger = L1CrossDomainMessenger(payable(address(l1CrossDomainMessengerProxy)));
        rollup = Rollup(payable(address(rollupProxy)));
        l1MessageQueueWithGasPriceOracle = L1MessageQueueWithGasPriceOracle(
            address(l1MessageQueueWithGasPriceOracleProxy)
        );
        l1Staking = L1Staking(address(l1StakingProxy));

        _changeAdmin(address(l1Staking));
        _changeAdmin(address(rollup));
        _changeAdmin(address(l1CrossDomainMessenger));
        _changeAdmin(address(l1MessageQueueWithGasPriceOracle));

        assertEq(address(l1CrossDomainMessenger), l1MessageQueueWithGasPriceOracle.MESSENGER());

        rollup.addChallenger(bob);
        hevm.stopPrank();
    }

    function messageProveAndRelayPrepare(
        address from,
        address to,
        uint256 value,
        uint256 nonce,
        bytes memory message
    ) public returns (bytes32[32] memory wdProof, bytes32 wdRoot) {
        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));

        // prove message
        (, wdProof, wdRoot) = ffi.getProveWithdrawalTransactionInputs(_xDomainCalldataHash);

        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (wdRoot)),
            abi.encode(true)
        );

        return (wdProof, wdRoot);
    }

    function upgradeStorage(address _messenger, address _rollup, address _enforcedTxGateway) public {
        TransparentUpgradeableProxy l1MessageQueueWithGasPriceOracleProxy = TransparentUpgradeableProxy(
            payable(address(l1MessageQueueWithGasPriceOracle))
        );
        L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
            payable(_messenger), // _messenger
            address(_rollup), // _rollup
            address(_enforcedTxGateway) // _enforcedTxGateway
        );
        assertEq(_messenger, l1MessageQueueWithGasPriceOracleImpl.MESSENGER());
        assertEq(_rollup, l1MessageQueueWithGasPriceOracleImpl.ROLLUP_CONTRACT());
        assertEq(_enforcedTxGateway, l1MessageQueueWithGasPriceOracleImpl.ENFORCED_TX_GATEWAAY());

        hevm.prank(multisig);
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(address(l1MessageQueueWithGasPriceOracleProxy)),
            address(l1MessageQueueWithGasPriceOracleImpl)
        );
        assertEq(_messenger, l1MessageQueueWithGasPriceOracle.MESSENGER());
        assertEq(_rollup, l1MessageQueueWithGasPriceOracle.ROLLUP_CONTRACT());
        assertEq(_enforcedTxGateway, l1MessageQueueWithGasPriceOracle.ENFORCED_TX_GATEWAAY());
    }
}
