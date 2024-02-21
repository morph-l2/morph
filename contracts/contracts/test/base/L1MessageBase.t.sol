// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {CommonTest} from "./CommonTest.t.sol";
import {L2GasPriceOracle} from "../../L1/rollup/L2GasPriceOracle.sol";
import {L1CrossDomainMessenger} from "../../L1/L1CrossDomainMessenger.sol";
import {L1MessageQueue} from "../../L1/rollup/L1MessageQueue.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {Rollup} from "../../L1/rollup/Rollup.sol";
import {IRollup} from "../../L1/rollup/IRollup.sol";
import {MockZkEvmVerifier} from "../../mock/MockZkEvmVerifier.sol";

contract L1MessageBaseTest is CommonTest {
    // L2GasPriceOracle config
    L2GasPriceOracle l2GasPriceOracle;
    L2GasPriceOracle l2GasPriceOracleImpl;
    uint64 txGas = 1;
    uint64 txGasContractCreation = 2;
    uint64 zeroGas = 1;
    uint64 nonZeroGas = 1;

    // Rollup config
    Rollup rollup;
    Rollup rollupImpl;
    MockZkEvmVerifier verifier;

    uint256 public PROOF_WINDOW = 100;
    uint256 public MIN_DEPOSIT = 1000000000000000000; // 1 eth
    uint256 public maxNumTxInChunk = 10;
    uint64 public layer2ChainId = 53077;
    uint32 public minGasLimit = 10000;

    // L1MessageQueue config
    event QueueTransaction(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint64 queueIndex,
        uint256 gasLimit,
        bytes data
    );

    L1MessageQueue l1MessageQueue;
    L1MessageQueue l1MessageQueueImpl;
    uint256 l1MessageQueue_maxGasLimit = 100000000;
    uint32 defaultGasLimit = 1000000;

    // L1MessageQueueWithGasPriceOracle config
    L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracle;

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

    L1CrossDomainMessenger l1CrossDomainMessenger;
    L1CrossDomainMessenger l1CrossDomainMessengerImpl;
    address l1FeeVault = address(3033);

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // deploy proxys
        TransparentUpgradeableProxy l2GasPriceOraclePorxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        TransparentUpgradeableProxy rollupProxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        TransparentUpgradeableProxy l1MessageQueueProxy = new TransparentUpgradeableProxy(
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

        // deploy mock verifier
        verifier = new MockZkEvmVerifier();
        // deploy impls
        l2GasPriceOracleImpl = new L2GasPriceOracle();
        rollupImpl = new Rollup(
            layer2ChainId,
            payable(address(l1CrossDomainMessengerProxy))
        );
        l1MessageQueueImpl = new L1MessageQueue(
            payable(address(l1CrossDomainMessengerProxy)),
            address(rollupProxy),
            address(alice)
        );
        L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
                payable(address(l1CrossDomainMessengerProxy)),
                address(rollupProxy),
                address(alice)
            );
        l1CrossDomainMessengerImpl = new L1CrossDomainMessenger();
        // upgrade and initialize
        ITransparentUpgradeableProxy(address(l2GasPriceOraclePorxy))
            .upgradeToAndCall(
                address(l2GasPriceOracleImpl),
                abi.encodeWithSelector(
                    L2GasPriceOracle.initialize.selector,
                    txGas,
                    txGasContractCreation,
                    zeroGas,
                    nonZeroGas
                )
            );

        ITransparentUpgradeableProxy(address(rollupProxy)).upgradeToAndCall(
            address(rollupImpl),
            abi.encodeWithSelector(
                Rollup.initialize.selector,
                address(l1MessageQueueWithGasPriceOracleProxy), // _messageQueue
                address(verifier), // _verifier
                maxNumTxInChunk, // _maxNumTxInChunk
                MIN_DEPOSIT, // _minDeposit
                FINALIZATION_PERIOD_SECONDS, // _finalizationPeriodSeconds
                PROOF_WINDOW // _proofWindow
            )
        );
        ITransparentUpgradeableProxy(address(l1MessageQueueProxy))
            .upgradeToAndCall(
                address(l1MessageQueueImpl),
                abi.encodeWithSelector(
                    L1MessageQueue.initialize.selector,
                    address(l2GasPriceOraclePorxy), // _gasOracle
                    l1MessageQueue_maxGasLimit // gasLimit
                )
            );
        ITransparentUpgradeableProxy(
            address(l1MessageQueueWithGasPriceOracleProxy)
        ).upgradeToAndCall(
                address(l1MessageQueueWithGasPriceOracleImpl),
                abi.encodeWithSelector(
                    L1MessageQueue.initialize.selector,
                    address(l2GasPriceOraclePorxy), // _gasOracle
                    l1MessageQueue_maxGasLimit // gasLimit
                )
            );
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxy))
            .upgradeToAndCall(
                address(l1CrossDomainMessengerImpl),
                abi.encodeWithSelector(
                    L1CrossDomainMessenger.initialize.selector,
                    l1FeeVault, // feeVault
                    address(rollupProxy), // rollup
                    address(l1MessageQueueWithGasPriceOracleProxy) // messageQueue
                )
            );

        rollup = Rollup(address(rollupProxy));
        l1CrossDomainMessenger = L1CrossDomainMessenger(
            payable(address(l1CrossDomainMessengerProxy))
        );
        l1MessageQueue = L1MessageQueue(address(l1MessageQueueProxy));
        l2GasPriceOracle = L2GasPriceOracle(address(l2GasPriceOraclePorxy));
        l1MessageQueueWithGasPriceOracle = L1MessageQueueWithGasPriceOracle(
            address(l1MessageQueueWithGasPriceOracleProxy)
        );

        _changeAdmin(address(rollup));
        _changeAdmin(address(l1CrossDomainMessenger));
        _changeAdmin(address(l1MessageQueue));
        _changeAdmin(address(l2GasPriceOracle));
        _changeAdmin(address(l1MessageQueueWithGasPriceOracle));

        assertEq(
            address(l1CrossDomainMessenger),
            l1MessageQueueWithGasPriceOracle.messenger()
        );
        assertEq(address(l1CrossDomainMessenger), l1MessageQueue.messenger());

        rollup.addSequencer(alice);
        rollup.addProver(alice);
        rollup.addProver(bob);
        rollup.addChallenger(bob);
        hevm.stopPrank();
    }

    function messageProve(
        address from,
        address to,
        uint256 value,
        uint256 nonce,
        bytes memory message
    ) public {
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );

        // prove message
        (, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(_xDomainCalldataHash);

        uint256 withdrawalBatchIndex = 1;
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(IRollup.withdrawalRoots.selector, wdRoot),
            abi.encode(withdrawalBatchIndex)
        );
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // warp finalization period
        (, uint256 provenTime, ) = l1CrossDomainMessenger.provenWithdrawals(
            _xDomainCalldataHash
        );
        hevm.warp(provenTime + FINALIZATION_PERIOD_SECONDS + 1);

        // finalize batch
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(
                IRollup.finalizedStateRoots.selector,
                withdrawalBatchIndex
            ),
            abi.encode(bytes32(uint256(1)))
        );
    }
}
