// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

/* Testing utilities */
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import "../CommonTest.t.sol";
import {CrossDomainMessenger} from "../../universal/CrossDomainMessenger.sol";
import {ResourceMetering} from "../../L1/ResourceMetering.sol";
import {Rollup} from "../../L1/Rollup.sol";

// Free function for setting the prevBaseFee param in the MorphPortal.
function setPrevBaseFee(Vm _vm, address _portal, uint128 _prevBaseFee) {
    _vm.store(
        address(_portal),
        bytes32(uint256(1)),
        bytes32((block.number << 192) | _prevBaseFee)
    );
}

contract SetPrevBaseFee_Test is Portal_Initializer {
    function test_setPrevBaseFee_succeeds() external {
        setPrevBaseFee(vm, address(portal), 100 gwei);
        (uint128 prevBaseFee, , uint64 prevBlockNum) = portal.params();
        assertEq(uint256(prevBaseFee), 100 gwei);
        assertEq(uint256(prevBlockNum), block.number);
    }
}

// Tests for obtaining pure gas cost estimates for commonly used functions.
// The objective with these benchmarks is to strip down the actual test functions
// so that they are nothing more than the call we want measure the gas cost of.
// In order to achieve this we make no assertions, and handle everything else in the setUp()
// function.
contract GasBenchMark_MorphPortal is Portal_Initializer {
    // Reusable default values for a test withdrawal
    Types.WithdrawalTransaction _defaultTx;

    bytes32[_TREE_DEPTH] _withdrawalProof;
    bytes32 _withdrawalRoot;
    bytes32 _withdrawalHash;
    bytes batchHeader0Store;
    bytes32 stateRoot;

    // Use a constructor to set the storage vars above, so as to minimize the number of ffi calls.
    constructor() {
        super.setUp();
        _defaultTx = Types.WithdrawalTransaction({
            nonce: 0,
            sender: alice,
            target: bob,
            value: 100,
            gasLimit: 100_000,
            data: hex""
        });

        (_withdrawalHash, _withdrawalProof, _withdrawalRoot) = ffi
            .getProveWithdrawalTransactionInputs(_defaultTx);
    }

    // Get the system into a nice ready-to-use state.
    function setUp() public virtual override {
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS());
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();
        bytes memory batchHeader0 = new bytes(89);
        stateRoot = bytes32(uint256(1));
        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );
        batchHeader0Store = batchHeader0;
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);

        // Fund the portal so that we can withdraw ETH.
        vm.deal(address(portal), 0xFFFFFFFF);
    }

    function test_depositTransaction_benchmark() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        portal.depositTransaction{value: NON_ZERO_VALUE}(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    function test_depositTransaction_benchmark_1() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        setPrevBaseFee(vm, address(portal), 1 gwei);
        portal.depositTransaction{value: NON_ZERO_VALUE}(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    function test_proveWithdrawalTransaction_benchmark() external {
        bytes[] memory chunks = new bytes[](1);
        bytes memory chunk0;
        chunk0 = new bytes(1 + 60);
        chunk0[0] = bytes1(uint8(1));
        chunks[0] = chunk0;
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0Store,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            _withdrawalRoot,
            nilBatchSig
        );
        // update portal l1Messager to caller
        vm.startPrank(address(0));
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }
}

contract GasBenchMark_L1CrossDomainMessenger is Portal_Initializer {
    function test_sendMessage_benchmark_0() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 1 gwei);
        // The amount of data typically sent during a bridge deposit.
        bytes
            memory data = hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
        vm.resumeGasMetering();
        L1Messenger.sendMessage(bob, data, uint32(100));
    }

    function test_sendMessage_benchmark_1() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 10 gwei);
        // The amount of data typically sent during a bridge deposit.
        bytes
            memory data = hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
        vm.resumeGasMetering();
        L1Messenger.sendMessage(bob, data, uint32(100));
    }
}

contract GasBenchMark_L1StandardBridge_Deposit is Bridge_Initializer {
    function setUp() public virtual override {
        super.setUp();
        deal(address(L1Token), alice, 100000, true);
        vm.startPrank(alice, alice);
        L1Token.approve(address(L1Bridge), type(uint256).max);
    }

    function test_depositETH_benchmark_0() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 1 gwei);
        vm.resumeGasMetering();
        L1Bridge.depositETH{value: 500}(50000, hex"");
    }

    function test_depositETH_benchmark_1() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 10 gwei);
        vm.resumeGasMetering();
        L1Bridge.depositETH{value: 500}(50000, hex"");
    }

    function test_depositERC20_benchmark_0() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 1 gwei);
        vm.resumeGasMetering();
        L1Bridge.bridgeERC20({
            _localToken: address(L1Token),
            _remoteToken: address(L2Token),
            _amount: 100,
            _minGasLimit: 100_000,
            _extraData: hex""
        });
    }

    function test_depositERC20_benchmark_1() external {
        vm.pauseGasMetering();
        setPrevBaseFee(vm, address(portal), 10 gwei);
        vm.resumeGasMetering();
        L1Bridge.bridgeERC20({
            _localToken: address(L1Token),
            _remoteToken: address(L2Token),
            _amount: 100,
            _minGasLimit: 100_000,
            _extraData: hex""
        });
    }
}

contract GasBenchMark_L1StandardBridge_Finalize is Bridge_Initializer {
    function setUp() public virtual override {
        super.setUp();
        deal(address(L1Token), address(L1Bridge), 100, true);
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.startPrank(address(L1Bridge.messenger()));
        vm.deal(address(L1Bridge.messenger()), 100);
    }

    function test_finalizeETHWithdrawal_benchmark() external {
        // TODO: Make this more accurate. It is underestimating the cost because it pranks
        // the call coming from the messenger, which bypasses the portal
        // and oracle.
        L1Bridge.finalizeETHWithdrawal{value: 100}(alice, alice, 100, hex"");
    }
}
