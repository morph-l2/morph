// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {CommonTest} from "./CommonTest.t.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Whitelist} from "../../libraries/common/Whitelist.sol";
import {L1CrossDomainMessenger} from "../../L1/L1CrossDomainMessenger.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {L1Staking} from "../../L1/staking/L1Staking.sol";
import {Rollup} from "../../L1/rollup/Rollup.sol";
import {IRollup} from "../../L1/rollup/IRollup.sol";
import {MockZkEvmVerifier} from "../../mock/MockZkEvmVerifier.sol";

contract L1MessageBaseTest is CommonTest {
    // Staking config
    L1Staking l1Staking;
    uint256 public beginSeq = 10;
    // uint256 public version = 0;
    uint256 public LOCK = 3;

    // Rollup config
    Rollup rollup;
    Rollup rollupImpl;
    MockZkEvmVerifier verifier;

    uint256 public proofWindow = 100;
    uint256 public stakingValue = 1000000000000000000; // 1 eth
    uint256 public maxNumTxInChunk = 10;
    uint64 public layer2ChainId = 53077;
    uint32 public minGasLimit = 10000;

    address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
    bytes32 public stateRoot = bytes32(uint256(1));
    IRollup.BatchData public batchData;
    IRollup.BatchChallengeReward public nilBatchSig;
    address[] public sequencerSigned;
    bytes public signature;

    event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);
    event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);
    event UpdateVerifier(
        address indexed oldVerifier,
        address indexed newVerifier
    );
    event UpdateMaxNumTxInChunk(
        uint256 oldMaxNumTxInChunk,
        uint256 newMaxNumTxInChunk
    );

    // whitelist config
    Whitelist whitelistChecker;

    // L1MessageQueueWithGasPriceOracle config
    event QueueTransaction(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint64 queueIndex,
        uint256 gasLimit,
        bytes data
    );
    L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracle;
    uint256 l1MessageQueue_maxGasLimit = 100000000;
    uint32 defaultGasLimit = 1000000;

    // L1CrossDomainMessenger config
    event SentMessage(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 messageNonce,
        uint256 gasLimit,
        bytes message
    );
    event FailedRelayedMessage(bytes32 indexed messageHash);
    event RelayedMessage(bytes32 indexed messageHash);

    L1Staking l1StakingImpl;
    L1CrossDomainMessenger l1CrossDomainMessenger;
    L1CrossDomainMessenger l1CrossDomainMessengerImpl;

    address l1FeeVault = address(3033);

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

        // deploy mock verifier
        verifier = new MockZkEvmVerifier();

        // deploy impl
        rollupImpl = new Rollup(layer2ChainId);
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
        ITransparentUpgradeableProxy(
            address(l1MessageQueueWithGasPriceOracleProxy)
        ).upgradeToAndCall(
                address(l1MessageQueueWithGasPriceOracleImpl),
                abi.encodeCall(
                    L1MessageQueueWithGasPriceOracle.initialize,
                    (
                        l1MessageQueue_maxGasLimit, // gasLimit
                        address(whitelistChecker) // whitelistChecker
                    )
                )
            );
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxy))
            .upgradeToAndCall(
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
                    address(alice),
                    address(rollupProxy),
                    20,
                    stakingValue,
                    LOCK,
                    defaultGasLimit,
                    defaultGasLimit
                )
            )
        );

        l1CrossDomainMessenger = L1CrossDomainMessenger(
            payable(address(l1CrossDomainMessengerProxy))
        );
        rollup = Rollup(payable(address(rollupProxy)));
        l1MessageQueueWithGasPriceOracle = L1MessageQueueWithGasPriceOracle(
            address(l1MessageQueueWithGasPriceOracleProxy)
        );
        l1Staking = L1Staking(address(l1StakingProxy));

        _changeAdmin(address(l1Staking));
        _changeAdmin(address(rollup));
        _changeAdmin(address(l1CrossDomainMessenger));
        _changeAdmin(address(l1MessageQueueWithGasPriceOracle));

        assertEq(
            address(l1CrossDomainMessenger),
            l1MessageQueueWithGasPriceOracle.messenger()
        );

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
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );

        // prove message
        (, wdProof, wdRoot) = ffi.getProveWithdrawalTransactionInputs(
            _xDomainCalldataHash
        );

        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (wdRoot)),
            abi.encode(true)
        );

        return (wdProof, wdRoot);
    }

    function upgradeStorage(
        address _messenger,
        address _rollup,
        address _enforcedTxGateway
    ) public {
        TransparentUpgradeableProxy l1MessageQueueWithGasPriceOracleProxy = TransparentUpgradeableProxy(
                payable(address(l1MessageQueueWithGasPriceOracle))
            );
        L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
            payable(_messenger), // _messenger
            address(_rollup), // _rollup
            address(_enforcedTxGateway) // _enforcedTxGateway
        );
        assertEq(_messenger, l1MessageQueueWithGasPriceOracleImpl.messenger());
        assertEq(_rollup, l1MessageQueueWithGasPriceOracleImpl.rollup());
        assertEq(
            _enforcedTxGateway,
            l1MessageQueueWithGasPriceOracleImpl.enforcedTxGateway()
        );

        hevm.prank(multisig);
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(
                address(l1MessageQueueWithGasPriceOracleProxy)
            ),
            address(l1MessageQueueWithGasPriceOracleImpl)
        );
        assertEq(_messenger, l1MessageQueueWithGasPriceOracle.messenger());
        assertEq(_rollup, l1MessageQueueWithGasPriceOracle.rollup());
        assertEq(
            _enforcedTxGateway,
            l1MessageQueueWithGasPriceOracle.enforcedTxGateway()
        );
    }
}
