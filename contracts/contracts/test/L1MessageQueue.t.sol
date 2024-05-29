// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {L1MessageQueueWithGasPriceOracle} from "../l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {IL1MessageQueueWithGasPriceOracle} from "../l1/rollup/IL1MessageQueueWithGasPriceOracle.sol";
import {L1CrossDomainMessenger} from "../l1/L1CrossDomainMessenger.sol";

contract L1MessageQueueTest is L1MessageBaseTest {
    /// @notice The intrinsic gas for transaction.
    uint256 INTRINSIC_GAS_TX = 21000;
    /// @notice The appropriate intrinsic gas for each byte.
    uint256 APPROPRIATE_INTRINSIC_GAS_PER_BYTE = 16;
    L1MessageQueueWithGasPriceOracle l1MessageQueue;

    function setUp() public virtual override {
        super.setUp();
        l1MessageQueue = l1MessageQueueWithGasPriceOracle;
    }

    function testConstructor() external {
        // Deploy L1MessageQueueWithGasPriceOracle
        // Verify it throws a custom error ErrZeroAddress() when the _messenger is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        l1MessageQueue = new L1MessageQueueWithGasPriceOracle(
                address(0),
                address(1),
                address(1)
            );

        // Verify it throws a custom error ErrZeroAddress() when the _rollup is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        l1MessageQueue = new L1MessageQueueWithGasPriceOracle(
                address(1),
                address(0),
                address(1)
            );   

        // Verify it throws a custom error ErrZeroAddress() when the _enforcedTxGateway is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        l1MessageQueue = new L1MessageQueueWithGasPriceOracle(
                address(1),
                address(1),
                address(0)
            );   
    }

    function testInitialize() external {
        // Verify initialize() initialized and the state variables are assigned successfully.
        assertEq(l1MessageQueueWithGasPriceOracle.maxGasLimit(), 100000000);
        assertEq(address(l1MessageQueueWithGasPriceOracle.whitelistChecker()), address(whitelistChecker));

        // Test modifier initializer, verify the initialize only can be called once.
        // Since the initiallize function is alreay called in L1MessageBaseTest.
        // Just call it again, it will trigger the error message. 
        // Use expectRevert to get the error message as expected.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1MessageQueueWithGasPriceOracle.initialize(l1MessageQueueMaxGasLimit, address(whitelistChecker));
    }

    function testGetCrossDomainMessage() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "message";
        uint256 value = 0;

        // Verify it throws error message "message index out of range" when the condition (1 < 0) returns false.
        // Get the return value of nextCrossDomainMessageIndex(), default is zero.
        uint256 nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        hevm.expectRevert("message index out of range");
        l1MessageQueueWithGasPriceOracle.getCrossDomainMessage(1);

        // Verify getCrossDomainMessage() executed successfully and returns correctly.
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
        bytes32 queueIndex = l1MessageQueueWithGasPriceOracle.getCrossDomainMessage(0);
        assertTrue(queueIndex != 0x0000000000000000000000000000000000000000000000000000000000000000);
    }

    function testEstimateCrossDomainMessageFee() external {
        hevm.startPrank(multisig);
        uint256 gasLimit = 100;
        
        // uint256 nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(1);

        // Verify the return value of estimateCrossDomainMessageFee() is equal to gasLimit * l2BaseFee.
        uint256 fee = l1MessageQueueWithGasPriceOracle
            .estimateCrossDomainMessageFee(multisig, gasLimit);
        assertEq(fee, gasLimit * 1);

        // Verify it returns 0
        address[] memory whiteList = new address[](1);
        whiteList[0] = address(multisig);
        whitelistChecker.updateWhitelistStatus(whiteList, true);
        assertEq(l1MessageQueueWithGasPriceOracle.estimateCrossDomainMessageFee(address(multisig), gasLimit), 0);

        hevm.stopPrank();
    }

    function testCalculateIntrinsicGasFee() external {
        hevm.startPrank(multisig);
        bytes memory _calldata = "";
        uint256 intrinsicGasFee = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _calldata
        );
        // Verify calculateIntrinsicGasFee() returns correctly when _calldata is empty.
        assertEq(
            intrinsicGasFee,
            INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE
        );
        _calldata = "0x00";
        intrinsicGasFee = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_calldata);
        // Verify calculateIntrinsicGasFee() returns correctly when _calldata isn't empty.
        assertEq(
            intrinsicGasFee,
            INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE
        );
        hevm.stopPrank();
    }


    function testSetL2BaseFee() external {
        // Verify the modifier onlyOwner works successfully.
        // It throws an error "Ownable: caller is not the owner" when msg.sender is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(1);

        // Update msg.sender is address(multisig).
        hevm.startPrank(multisig);
        // Emit the expected event.
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueueWithGasPriceOracle.UpdateL2BaseFee(0, 10);
        // Call the function that should emit the event.
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(10);
        // Verify the l2BaseFee is equal to 10.
        assertEq(l1MessageQueueWithGasPriceOracle.l2BaseFee(), 10);
        hevm.stopPrank();
    }

    function test_validateGasLimit() external {
        // store alice as messenger
        upgradeStorage(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());

        // append message
        hevm.prank(multisig);
        l1MessageQueue.updateMaxGasLimit(1);
        hevm.prank(alice);
        hevm.expectRevert("Gas limit must not exceed maxGasLimit");
        l1MessageQueue.appendCrossDomainMessage(alice, 3, "0x0");

        hevm.prank(multisig);
        l1MessageQueue.updateMaxGasLimit(100);
        hevm.prank(alice);
        hevm.expectRevert(
            "Insufficient gas limit, must be above intrinsic gas"
        );
        l1MessageQueue.appendCrossDomainMessage(alice, 3, "0x0");
    }

    function test_appendCrossDomainMessage() external {
        bytes memory _calldata = "0x0";

        // Verify the modifier onlyMessenger works correctly.
        // It throws an error "Only callable by the L1CrossDomainMessenger" when msg.sender isn't MESSENGER.
        hevm.expectRevert("Only callable by the L1CrossDomainMessenger");
        l1MessageQueueWithGasPriceOracle.appendCrossDomainMessage(
            address(alice),
            100,
            _calldata
        );

        // store alice as messenger
        upgradeStorage(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());
        address sender = AddressAliasHelper.applyL1ToL2Alias(address(alice));
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueue.QueueTransaction(
            sender,
            alice,
            0,
            0,
            gasLimit,
            _calldata
        );
        hevm.startPrank(alice);
        l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
        hevm.stopPrank();
    }

    function test_appendEnforcedTransaction() external {
        bytes memory _calldata = "0x0";
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");

        // Verify it throws an error "Only callable by the EnforcedTxGateway" when msg.sender isn't ENFORCED_TX_GATEWAAY.
        hevm.expectRevert("Only callable by the EnforcedTxGateway");
        l1MessageQueue.appendEnforcedTransaction(
            alice,
            bob,
            0,
            gasLimit,
            _calldata
        );

        // Verify it throws an error "only EOA" when msg.sender isn't an EOA address.
        hevm.prank(alice);
        hevm.expectRevert("only EOA");
        l1MessageQueue.appendEnforcedTransaction(
            address(this),
            bob,
            0,
            gasLimit,
            _calldata
        );


        hevm.prank(multisig);
        assertEq(alice, l1MessageQueue.ENFORCED_TX_GATEWAAY());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());

        // Verify the event QueueTransaction is emitted successfully as expected.
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueue.QueueTransaction(
            alice,
            bob,
            0,
            0,
            gasLimit,
            _calldata
        );
        hevm.prank(alice);
        l1MessageQueue.appendEnforcedTransaction(
            alice,
            bob,
            0,
            gasLimit,
            _calldata
        );
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
    }

    function test_pop_dropCrossDomainMessage() external {
        // store alice as messenger and rollup
        upgradeStorage(address(alice), address(alice), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());
        assertEq(alice, l1MessageQueue.ROLLUP_CONTRACT());
        bytes memory _calldata = "0x0";
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");

        // Verify it throws an error "Only callable by the rollup" when the msg.sender is not the ROLLUP_CONTRACT.
        hevm.prank(alice);
        l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        hevm.prank(bob);
        hevm.expectRevert("Only callable by the rollup");
        l1MessageQueue.popCrossDomainMessage(0, 1, 0x3ff);

        // Verify it throws an error "pop too many messages" when _count > 256.
        hevm.prank(alice);
        hevm.expectRevert("pop too many messages");
        l1MessageQueue.popCrossDomainMessage(0, 257, 0x3ff);

        // Verify it throws an error "start index mismatch" when pendingQueueIndex != _startIndex.
        hevm.prank(alice);
        hevm.expectRevert("start index mismatch");
        l1MessageQueue.popCrossDomainMessage(1, 2, 0x3ff);

        // Verify it throws an error "cannot drop pending message" when (_index < pendingQueueIndex) is false.
        hevm.prank(alice);
        l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        hevm.prank(alice);
        hevm.expectRevert("cannot drop pending message");
        l1MessageQueue.dropCrossDomainMessage(0);

        // append 10 message
        hevm.startPrank(alice);
        for (uint64 i = 0; i < 10; i++) {
            l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        }

        // Verify the event QueueTransaction is emitted successfully as expected.
        hevm.expectEmit(false, false, false, true);
        emit IL1MessageQueue.DequeueTransaction(0, 10, 0x3ff);

        // pop all 10 message
        l1MessageQueue.popCrossDomainMessage(0, 10, 0x3ff);
        for (uint64 i = 0; i < 10; i++) {
            assertTrue(l1MessageQueue.isMessageSkipped(i));
        }
        // drop all 10 message
        for (uint64 i = 0; i < 10; i++) {
            l1MessageQueue.dropCrossDomainMessage(i);
            assertTrue(l1MessageQueue.isMessageDropped(i));
        }
        hevm.stopPrank();
        
        // Verify it throws an error "start index mismatch" when pendingQueueIndex != _startIndex.
        hevm.prank(alice);
        hevm.expectRevert("message already dropped");
        l1MessageQueue.dropCrossDomainMessage(1);
    }

    function testUpdateMaxGasLimit() external{
        // Verify the modifier onlyOwner works successfully.
        // It throws an error "Ownable: caller is not the owner" when msg.sender is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l1MessageQueueWithGasPriceOracle.updateMaxGasLimit(1);

        // Verify the event is emitted successfully and the data is correct.
        hevm.expectEmit(false, false, false, true);
        emit IL1MessageQueue.UpdateMaxGasLimit(l1MessageQueueMaxGasLimit, 1);
        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.updateMaxGasLimit(1);
    }

    function testUpdateWhitelistChecker() external {
        // Verify the event is emitted successfully.
        hevm.expectEmit(true, true, false, false);
        emit IL1MessageQueueWithGasPriceOracle.UpdateWhitelistChecker(
            address(whitelistChecker),
            address(alice)
        );
        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.updateWhitelistChecker(address(alice));
        // Verify the whiteListChecker is updated successfully.
        assertEq(l1MessageQueueWithGasPriceOracle.whitelistChecker(), address(alice));
    }
}
