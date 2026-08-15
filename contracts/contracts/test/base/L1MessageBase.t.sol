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
import {BatchHeaderCodecV0} from "../../libraries/codec/BatchHeaderCodecV0.sol";
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

    // Rollup storage slot for batchBlobVersionedHashes (forge inspect Rollup storage-layout)
    uint256 internal constant BATCH_BLOB_VERSIONED_HASHES_SLOT = 173;
    uint256 internal constant ROLLUP_DELAY_PERIOD_SLOT = 172;
    bytes32 internal constant ZERO_VERSIONED_HASH = 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;

    /// @dev Sets batchBlobVersionedHashes[batchIndex] so commitState/commitBatchWithProof can use stored hash in tests.
    function _setStoredBlobHash(uint256 batchIndex) internal {
        bytes32 slot = keccak256(abi.encode(batchIndex, BATCH_BLOB_VERSIONED_HASHES_SLOT));
        hevm.store(address(rollup), slot, ZERO_VERSIONED_HASH);
    }

    /// @dev Data hash for batch when numL1Messages=0: keccak256(8 bytes lastBlockNumber || 2 bytes numL1Messages).
    function _computeDataHash(uint64 lastBlockNumber, uint16 numL1Messages) internal pure returns (bytes32) {
        bytes memory data = new bytes(10);
        assembly {
            mstore(add(data, 0x20), shl(192, lastBlockNumber))
            mstore(add(data, 0x28), shl(240, numL1Messages))
        }
        return keccak256(data);
    }

    /// @dev Setup rollup delay and warp so commitBatchWithProof timing passes.
    function _setupDelayAndWarpForProof() internal {
        hevm.store(address(rollup), bytes32(ROLLUP_DELAY_PERIOD_SLOT), bytes32(uint256(3600)));
        hevm.warp(block.timestamp + 3601);
    }

    /// @dev Mock L1 message queue so getFirstUnfinalizedMessageEnqueueTime is not delayed (avoids "l1msg delay" when rollupDelay is used).
    function _mockMessageQueueNotDelayedForProof() internal {
        hevm.mockCall(
            address(l1MessageQueueWithGasPriceOracle),
            abi.encodeWithSignature("getFirstUnfinalizedMessageEnqueueTime()"),
            abi.encode(block.timestamp)
        );
    }

    /// @dev Mock verifier for commitBatchWithProof.
    function _mockVerifierForProof() internal {
        hevm.mockCall(
            rollup.verifier(),
            abi.encodeWithSignature("verifyAggregateProof(uint256,uint256,bytes,bytes32)"),
            abi.encode()
        );
    }

    /// @dev Build V0 batch header for commitBatchWithProof (blob = ZERO_VERSIONED_HASH).
    function _createBatchHeaderV0ForProof(
        uint256 batchIndex,
        uint64 l1MessagePopped,
        uint64 totalL1MessagePopped,
        bytes32 dataHash,
        bytes32 prevStateRoot,
        bytes32 postStateRoot,
        bytes32 withdrawalRoot,
        bytes32 sequencerSetVerifyHash,
        bytes32 parentBatchHash
    ) internal pure returns (bytes memory batchHeader) {
        batchHeader = new bytes(BatchHeaderCodecV0.BATCH_HEADER_LENGTH);
        bytes32 blobHash = ZERO_VERSIONED_HASH;
        assembly {
            let p := add(batchHeader, 0x20)
            mstore(p, 0)
            mstore(add(p, 1), shl(192, batchIndex))
            mstore(add(p, 9), shl(192, l1MessagePopped))
            mstore(add(p, 17), shl(192, totalL1MessagePopped))
            mstore(add(p, 25), dataHash)
            mstore(add(p, 57), blobHash)
            mstore(add(p, 89), prevStateRoot)
            mstore(add(p, 121), postStateRoot)
            mstore(add(p, 153), withdrawalRoot)
            mstore(add(p, 185), sequencerSetVerifyHash)
            mstore(add(p, 217), parentBatchHash)
        }
    }

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
                    finalizationPeriodSeconds, // _finalizationPeriodSeconds
                    proofWindow, // _proofWindow
                    proofRewardPercent // _proofRewardPercent
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
