// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

// import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// import {CommonTest} from "./CommonTest.t.sol";
// import {Staking} from "../../L1/staking/Staking.sol";
// import {L1Sequencer} from "../../L1/staking/L1Sequencer.sol";
// import {Predeploys} from "../../libraries/constants/Predeploys.sol";
// import {L1CrossDomainMessenger} from "../../L1/L1CrossDomainMessenger.sol";
// import {L1MessageQueueWithGasPriceOracle} from "../../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";
// import {Rollup} from "../../L1/rollup/Rollup.sol";
// import {IRollup} from "../../L1/rollup/IRollup.sol";
// import {MockZkEvmVerifier} from "../../mock/MockZkEvmVerifier.sol";

// contract L1MessageBaseTest is CommonTest {
//     // staking config
//     event Registered(
//         address addr,
//         bytes32 tmKey,
//         bytes blsKey,
//         uint256 balance
//     );
//     event SequencerUpdated(
//         address[] sequencersAddr,
//         bytes[] sequencersBLS,
//         uint256 version
//     );
//     Staking staking;
//     uint256 public beginSeq = 10;
//     uint256 public version = 0;
//     address[] public sequencerAddrs;
//     bytes[] public sequencerBLSKeys;
//     uint256 public constant SEQUENCER_SIZE = 3;
//     uint256 public LOCK = 3;

//     // L1Sequencer config
//     L1Sequencer l1Sequencer;

//     address l2Sequencer = address(Predeploys.L2_SEQUENCER);

//     // Rollup config
//     Rollup rollup;
//     Rollup rollupImpl;
//     MockZkEvmVerifier verifier;

//     uint256 public PROOF_WINDOW = 100;
//     uint256 public MIN_DEPOSIT = 1000000000000000000; // 1 eth
//     uint256 public maxNumTxInChunk = 10;
//     uint64 public layer2ChainId = 53077;
//     uint32 public minGasLimit = 10000;

//     address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
//     bytes32 public stateRoot = bytes32(uint256(1));
//     IRollup.BatchData public batchData;
//     IRollup.BatchSignature public nilBatchSig;
//     address sequencerAddr = address(uint160(beginSeq));
//     uint256 public sequencerVersion;
//     uint256[] public sequencerIndex;
//     bytes public signature;

//     event UpdateSequencer(address indexed account, bool status);
//     event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);
//     event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);
//     event UpdateProver(address indexed account, bool status);
//     event UpdateVerifier(
//         address indexed oldVerifier,
//         address indexed newVerifier
//     );
//     event UpdateMaxNumTxInChunk(
//         uint256 oldMaxNumTxInChunk,
//         uint256 newMaxNumTxInChunk
//     );

//     // L1MessageQueueWithGasPriceOracle config
//     event QueueTransaction(
//         address indexed sender,
//         address indexed target,
//         uint256 value,
//         uint64 queueIndex,
//         uint256 gasLimit,
//         bytes data
//     );
//     L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracle;
//     uint256 l1MessageQueue_maxGasLimit = 100000000;
//     uint32 defaultGasLimit = 1000000;

//     // L1CrossDomainMessenger config
//     event SentMessage(
//         address indexed sender,
//         address indexed target,
//         uint256 value,
//         uint256 messageNonce,
//         uint256 gasLimit,
//         bytes message
//     );
//     event FailedRelayedMessage(bytes32 indexed messageHash);
//     event RelayedMessage(bytes32 indexed messageHash);

//     Staking stakingImpl;
//     L1Sequencer l1SequencerImpl;
//     L1CrossDomainMessenger l1CrossDomainMessenger;
//     L1CrossDomainMessenger l1CrossDomainMessengerImpl;

//     address l1FeeVault = address(3033);

//     function setUp() public virtual override {
//         super.setUp();
//         hevm.startPrank(multisig);

//         // deploy proxys
//         TransparentUpgradeableProxy rollupProxy = new TransparentUpgradeableProxy(
//                 address(emptyContract),
//                 address(multisig),
//                 new bytes(0)
//             );
//         TransparentUpgradeableProxy l1CrossDomainMessengerProxy = new TransparentUpgradeableProxy(
//                 address(emptyContract),
//                 address(multisig),
//                 new bytes(0)
//             );
//         TransparentUpgradeableProxy l1MessageQueueWithGasPriceOracleProxy = new TransparentUpgradeableProxy(
//                 address(emptyContract),
//                 address(multisig),
//                 new bytes(0)
//             );
//         TransparentUpgradeableProxy stakingProxy = new TransparentUpgradeableProxy(
//                 address(emptyContract),
//                 address(multisig),
//                 new bytes(0)
//             );
//         TransparentUpgradeableProxy l1SequencerProxy = new TransparentUpgradeableProxy(
//                 address(emptyContract),
//                 address(multisig),
//                 new bytes(0)
//             );

//         // deploy mock verifier
//         verifier = new MockZkEvmVerifier();

//         // deploy impls
//         rollupImpl = new Rollup(
//             layer2ChainId,
//             payable(address(l1CrossDomainMessengerProxy))
//         );
//         L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
//                 payable(address(l1CrossDomainMessengerProxy)),
//                 address(rollupProxy),
//                 address(alice)
//             );
//         l1CrossDomainMessengerImpl = new L1CrossDomainMessenger();
//         stakingImpl = new Staking();
//         l1SequencerImpl = new L1Sequencer(payable(l1CrossDomainMessengerProxy));

//         // upgrade and initialize
//         ITransparentUpgradeableProxy(address(rollupProxy)).upgradeToAndCall(
//             address(rollupImpl),
//             abi.encodeWithSelector(
//                 Rollup.initialize.selector,
//                 address(l1SequencerProxy),
//                 address(stakingProxy),
//                 address(l1MessageQueueWithGasPriceOracleProxy), // _messageQueue
//                 address(verifier), // _verifier
//                 maxNumTxInChunk, // _maxNumTxInChunk
//                 FINALIZATION_PERIOD_SECONDS, // _finalizationPeriodSeconds
//                 PROOF_WINDOW // _proofWindow
//             )
//         );
//         ITransparentUpgradeableProxy(
//             address(l1MessageQueueWithGasPriceOracleProxy)
//         ).upgradeToAndCall(
//                 address(l1MessageQueueWithGasPriceOracleImpl),
//                 abi.encodeWithSelector(
//                     L1MessageQueueWithGasPriceOracle.initialize.selector,
//                     l1MessageQueue_maxGasLimit // gasLimit
//                 )
//             );
//         ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxy))
//             .upgradeToAndCall(
//                 address(l1CrossDomainMessengerImpl),
//                 abi.encodeWithSelector(
//                     L1CrossDomainMessenger.initialize.selector,
//                     l1FeeVault, // feeVault
//                     address(rollupProxy), // rollup
//                     address(l1MessageQueueWithGasPriceOracleProxy) // messageQueue
//                 )
//             );
//         ITransparentUpgradeableProxy(address(stakingProxy)).upgradeToAndCall(
//             address(stakingImpl),
//             abi.encodeWithSelector(
//                 Staking.initialize.selector,
//                 address(alice),
//                 address(l1SequencerProxy),
//                 SEQUENCER_SIZE,
//                 MIN_DEPOSIT,
//                 LOCK
//             )
//         );

//         ITransparentUpgradeableProxy(address(l1SequencerProxy))
//             .upgradeToAndCall(
//                 address(l1SequencerImpl),
//                 abi.encodeWithSelector(
//                     L1Sequencer.initialize.selector,
//                     address(stakingProxy),
//                     address(rollupProxy)
//                 )
//             );

//         l1CrossDomainMessenger = L1CrossDomainMessenger(
//             payable(address(l1CrossDomainMessengerProxy))
//         );
//         rollup = Rollup(address(rollupProxy));
//         l1MessageQueueWithGasPriceOracle = L1MessageQueueWithGasPriceOracle(
//             address(l1MessageQueueWithGasPriceOracleProxy)
//         );
//         staking = Staking(address(stakingProxy));
//         l1Sequencer = L1Sequencer(payable(address(l1SequencerProxy)));

//         _changeAdmin(address(staking));
//         _changeAdmin(address(l1Sequencer));
//         _changeAdmin(address(rollup));
//         _changeAdmin(address(l1CrossDomainMessenger));
//         _changeAdmin(address(l1MessageQueueWithGasPriceOracle));

//         assertEq(
//             address(l1CrossDomainMessenger),
//             l1MessageQueueWithGasPriceOracle.messenger()
//         );

//         rollup.addProver(alice);
//         rollup.addProver(bob);
//         rollup.addChallenger(bob);
//         hevm.stopPrank();
//     }

//     function messageProve(
//         address from,
//         address to,
//         uint256 value,
//         uint256 nonce,
//         bytes memory message
//     ) public {
//         bytes32 _xDomainCalldataHash = keccak256(
//             _encodeXDomainCalldata(from, to, value, nonce, message)
//         );

//         // prove message
//         (, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
//             .getProveWithdrawalTransactionInputs(_xDomainCalldataHash);

//         uint256 withdrawalBatchIndex = 1;
//         hevm.mockCall(
//             address(l1CrossDomainMessenger.rollup()),
//             abi.encodeWithSelector(IRollup.withdrawalRoots.selector, wdRoot),
//             abi.encode(withdrawalBatchIndex)
//         );
//         l1CrossDomainMessenger.proveMessage(
//             from,
//             to,
//             value,
//             nonce,
//             message,
//             wdProof,
//             wdRoot
//         );

//         // warp finalization period
//         (, uint256 provenTime, ) = l1CrossDomainMessenger.provenWithdrawals(
//             _xDomainCalldataHash
//         );
//         hevm.warp(provenTime + FINALIZATION_PERIOD_SECONDS + 1);

//         // finalize batch
//         hevm.mockCall(
//             address(l1CrossDomainMessenger.rollup()),
//             abi.encodeWithSelector(
//                 IRollup.finalizedStateRoots.selector,
//                 withdrawalBatchIndex
//             ),
//             abi.encode(bytes32(uint256(1)))
//         );
//     }

//     function upgradeStorage(
//         address _messenger,
//         address _rollup,
//         address _enforcedTxGateway
//     ) public {
//         TransparentUpgradeableProxy l1MessageQueueWithGasPriceOracleProxy = TransparentUpgradeableProxy(
//                 payable(address(l1MessageQueueWithGasPriceOracle))
//             );
//         L1MessageQueueWithGasPriceOracle l1MessageQueueWithGasPriceOracleImpl = new L1MessageQueueWithGasPriceOracle(
//                 payable(_messenger), // _messenger
//                 address(_rollup), // _rollup
//                 address(_enforcedTxGateway) // _enforcedTxGateway
//             );
//         assertEq(_messenger, l1MessageQueueWithGasPriceOracleImpl.messenger());
//         assertEq(_rollup, l1MessageQueueWithGasPriceOracleImpl.rollup());
//         assertEq(
//             _enforcedTxGateway,
//             l1MessageQueueWithGasPriceOracleImpl.enforcedTxGateway()
//         );

//         hevm.prank(multisig);
//         proxyAdmin.upgrade(
//             ITransparentUpgradeableProxy(
//                 address(l1MessageQueueWithGasPriceOracleProxy)
//             ),
//             address(l1MessageQueueWithGasPriceOracleImpl)
//         );
//         assertEq(_messenger, l1MessageQueueWithGasPriceOracle.messenger());
//         assertEq(_rollup, l1MessageQueueWithGasPriceOracle.rollup());
//         assertEq(
//             _enforcedTxGateway,
//             l1MessageQueueWithGasPriceOracle.enforcedTxGateway()
//         );
//     }
// }
