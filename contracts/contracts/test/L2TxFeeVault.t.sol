// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";

import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";
import {L2TxFeeVault} from "../l2/system/L2TxFeeVault.sol";

contract L2TxFeeVaultTest is DSTestPlus {
    MockCrossDomainMessenger private messenger;
    L2TxFeeVault private vault;

    function setUp() public {
        messenger = new MockCrossDomainMessenger();
        vault = new L2TxFeeVault(address(this), address(1), 10 ether);
        vault.updateMessenger(address(messenger));
    }

    function test_owner_succeeds() public {
        assertEq(vault.owner(), address(this));
    }

    function test_transferOwnership_succeeds() public {
        address newOwner = address(100);

        vault.transferOwnership(newOwner);
        assertEq(vault.owner(), newOwner);

        hevm.prank(newOwner);
        vault.transferOwnership(address(this));
        assertEq(vault.owner(), address(this));
    }

    function test_renounceOwnership_succeeds() public {
        assertEq(vault.owner(), address(this));

        vault.renounceOwnership();
        assertEq(vault.owner(), address(0));
    }

    function test_withdraw_onlyOwner_reverts() public {
        hevm.deal(address(vault), 9 ether);
        hevm.expectRevert("caller is not the owner");
        hevm.prank(address(100));
        vault.withdraw();
    }

    function test_withdraw_belowMinimum_reverts() public {
        hevm.deal(address(vault), 9 ether);
        hevm.expectRevert("FeeVault: withdrawal amount must be greater than minimum withdrawal amount");
        vault.withdraw();
    }

    function test_withdraw_amountBelowMinimum_reverts(uint256 amount) public {
        amount = bound(amount, 0 ether, 10 ether - 1);
        hevm.deal(address(vault), 100 ether);
        hevm.expectRevert("FeeVault: withdrawal amount must be greater than minimum withdrawal amount");
        vault.withdraw(amount);
    }

    function test_withdraw_moreThanBalance_reverts(uint256 amount) public {
        hevm.assume(amount >= 10 ether);
        hevm.deal(address(vault), amount - 1);
        hevm.expectRevert("FeeVault: insufficient balance to withdraw");
        vault.withdraw(amount);
    }

    function test_withdrawOnce_succeeds() public {
        hevm.deal(address(vault), 11 ether);
        vault.withdraw();
        assertEq(address(messenger).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);
    }

    function test_withdrawAmountOnce_succeeds(uint256 amount) public {
        amount = bound(amount, 10 ether, 100 ether);

        hevm.deal(address(vault), 100 ether);
        vault.withdraw(amount);

        assertEq(address(messenger).balance, amount);
        assertEq(vault.totalProcessed(), amount);
        assertEq(address(vault).balance, 100 ether - amount);
    }

    function test_withdrawTwice_succeeds() public {
        hevm.deal(address(vault), 11 ether);
        vault.withdraw();
        assertEq(address(messenger).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);

        hevm.deal(address(vault), 22 ether);
        vault.withdraw();
        assertEq(address(messenger).balance, 33 ether);
        assertEq(vault.totalProcessed(), 33 ether);
    }

    function test_withdrawAmountTwice_succeeds(uint256 amount1, uint256 amount2) public {
        amount1 = bound(amount1, 10 ether, 100 ether);
        amount2 = bound(amount2, 10 ether, 100 ether);

        hevm.deal(address(vault), 200 ether);

        vault.withdraw(amount1);
        assertEq(address(messenger).balance, amount1);
        assertEq(vault.totalProcessed(), amount1);

        vault.withdraw(amount2);
        assertEq(address(messenger).balance, amount1 + amount2);
        assertEq(vault.totalProcessed(), amount1 + amount2);

        assertEq(address(vault).balance, 200 ether - amount1 - amount2);
    }

    function test_transfer_moreThanBalance_reverts(uint256 amount, address to) public {
        hevm.assume(to != address(0));
        hevm.assume(amount >= 10 ether);
        address[] memory allowedReceive = new address[](1);
        // set receive allowed account
        allowedReceive[0] = to;
        vault.updateReceiveAllowed(allowedReceive, true);
        hevm.deal(address(vault), amount - 1);
        hevm.expectRevert("FeeVault: insufficient balance to transfer");
        vault.transferTo(to, amount);
    }

    function test_allowed_transfer_reverts_caller_not_allowed(address to, address allowed) public {
        hevm.assume(to != address(0));
        hevm.assume(allowed != address(0));
        hevm.assume(allowed != address(this));

        // set allowed account
        hevm.deal(address(vault), 11 ether);
        hevm.expectRevert("FeeVault: caller is not allowed");
        hevm.prank(allowed);
        vault.transferTo(to);
    }

    function test_allowed_transfer_reverts_receiver_not_allowed(address to, address allowed) public {
        hevm.assume(to != address(0));
        hevm.assume(allowed != address(0));
        hevm.assume(allowed != address(this));
        address[] memory allowedAccounts = new address[](1);
        // set allowed account
        allowedAccounts[0] = allowed;
        vault.updateTransferAllowedStatus(allowedAccounts, true);
        hevm.deal(address(vault), 11 ether);
        hevm.expectRevert("FeeVault: recipient address not allowed");
        hevm.prank(allowed);
        vault.transferTo(to);
    }

    function test_owner_transfer_succeeds() public {
        address to = address(1024);
        address[] memory allowedReceive = new address[](1);
        // set receive allowed account
        allowedReceive[0] = to;
        vault.updateReceiveAllowed(allowedReceive, true);
        hevm.deal(address(vault), 11 ether);
        vault.transferTo(to);
        assertEq(address(to).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);
    }

    function test_allowed_transfer_succeeds(address allowed) public {
        address to = address(1024);
        hevm.assume(allowed != address(0));
        hevm.assume(address(to).balance == 0);
        address[] memory allowedAccounts = new address[](1);
        // set allowed account
        allowedAccounts[0] = allowed;
        vault.updateTransferAllowedStatus(allowedAccounts, true);

        address[] memory allowedReceive = new address[](1);
        // set receive allowed account
        allowedReceive[0] = to;
        vault.updateReceiveAllowed(allowedReceive, true);

        hevm.deal(address(vault), 11 ether);
        hevm.prank(allowed);
        vault.transferTo(to);
        assertEq(address(to).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);
    }
}
