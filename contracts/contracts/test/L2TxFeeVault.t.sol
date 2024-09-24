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

    function test_withdraw_zeroBalance_reverts(uint256 amount) public {
        hevm.assume(amount >= 10 ether);
        hevm.deal(address(vault), 0);
        hevm.expectRevert("FeeVault: insufficient balance to withdraw");
        vault.withdraw(amount);
    }

    function test_withdraw_zeroReceiptAddress_reverts() public {
        vault.updateRecipient(address(0));
        hevm.expectRevert("FeeVault: recipient address cannot be address(0)");
        vault.withdraw();
    }

    function test_withdrawOnce_succeeds() public {
        hevm.deal(address(vault), 11 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(address(vault).balance, vault.recipient(), address(this));
        vault.withdraw();
        assertEq(address(messenger).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);
    }

    function test_withdrawAmountOnce_succeeds(uint256 amount) public {
        amount = bound(amount, 10 ether, 100 ether);

        hevm.deal(address(vault), 100 ether);
        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(amount, vault.recipient(), address(this));
        vault.withdraw(amount);

        assertEq(address(messenger).balance, amount);
        assertEq(vault.totalProcessed(), amount);
        assertEq(address(vault).balance, 100 ether - amount);
    }

    function test_withdrawTwice_succeeds() public {
        hevm.deal(address(vault), 11 ether);
        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(11 ether, vault.recipient(), address(this));
        vault.withdraw();
        assertEq(address(messenger).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);

        hevm.deal(address(vault), 22 ether);
        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(address(vault).balance, vault.recipient(), address(this));
        vault.withdraw();
        assertEq(address(messenger).balance, 33 ether);
        assertEq(vault.totalProcessed(), 33 ether);
    }

    function test_withdrawAmountTwice_succeeds(uint256 amount1, uint256 amount2) public {
        amount1 = bound(amount1, 10 ether, 100 ether);
        amount2 = bound(amount2, 10 ether, 100 ether);

        hevm.deal(address(vault), 200 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(amount1, vault.recipient(), address(this));

        vault.withdraw(amount1);
        assertEq(address(messenger).balance, amount1);
        assertEq(vault.totalProcessed(), amount1);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(amount2, vault.recipient(), address(this));

        vault.withdraw(amount2);
        assertEq(address(messenger).balance, amount1 + amount2);
        assertEq(vault.totalProcessed(), amount1 + amount2);

        assertEq(address(vault).balance, 200 ether - amount1 - amount2);
    }

    function test_withdraw_minWithdrawAmountUpdated_succeeds() public {
        vault.updateMinWithdrawAmount(20 ether);
        hevm.deal(address(vault), 30 ether);

        hevm.expectRevert("FeeVault: withdrawal amount must be greater than minimum withdrawal amount");
        vault.withdraw(15 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(20 ether, vault.recipient(), address(this));

        vault.withdraw(20 ether);
        assertEq(address(messenger).balance, 20 ether);
        assertEq(vault.totalProcessed(), 20 ether);
    }

    function test_withdraw_recipientUpdated_succeeds() public {
        vault.updateRecipient(address(123));
        hevm.deal(address(vault), 50 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(30 ether, vault.recipient(), address(this));

        vault.withdraw(30 ether);
        assertEq(vault.recipient(), address(123));
        assertEq(address(messenger).balance, 30 ether);
        assertEq(vault.totalProcessed(), 30 ether);

        vault.updateRecipient(address(456));

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Withdrawal(20 ether, vault.recipient(), address(this));

        vault.withdraw(20 ether);
        assertEq(vault.recipient(), address(456));
        assertEq(address(messenger).balance, 50 ether);
        assertEq(vault.totalProcessed(), 50 ether);
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

    function test_transfer_zeroToAddress_reverts() public {
        hevm.deal(address(vault), 10 ether);
        hevm.expectRevert("FeeVault: recipient address cannot be address(0)");
        vault.transferTo(address(0), 10 ether);
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

    function test_transfer_notOwner_reverts() public {
        hevm.deal(address(vault), 11 ether);
        hevm.expectRevert("FeeVault: caller is not allowed");
        hevm.prank(address(1));
        vault.transferTo(address(1));
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

    function test_transfer_allowedReceiveIsFalse_reverts(address to, address allowed) public {
        hevm.assume(to != address(0));
        hevm.assume(allowed != address(0));
        hevm.assume(allowed != address(this));
        address[] memory allowedReceive = new address[](1);
        allowedReceive[0] = allowed;
        vault.updateReceiveAllowed(allowedReceive, false);
        hevm.deal(address(vault), 11 ether);
        hevm.expectRevert("FeeVault: recipient address not allowed");
        vault.transferTo(to);
    }

    function test_owner_transfer_succeeds() public {
        address to = address(1024);
        address[] memory allowedReceive = new address[](1);
        // set receive allowed account
        allowedReceive[0] = to;
        vault.updateReceiveAllowed(allowedReceive, true);
        hevm.deal(address(vault), 11 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Transfer(address(vault).balance, vault.recipient(), address(this));

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

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Transfer(address(vault).balance, vault.recipient(), allowed);

        hevm.prank(allowed);
        vault.transferTo(to);
        assertEq(address(to).balance, 11 ether);
        assertEq(vault.totalProcessed(), 11 ether);
    }

    function test_transfer_balanceCorrectAfterMultipleTransfers_succeeds() public {
        address to = address(1024);
        address[] memory allowedReceive = new address[](1);
        allowedReceive[0] = to;
        vault.updateReceiveAllowed(allowedReceive, true);

        hevm.deal(address(vault), 20 ether);
        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Transfer(10 ether, vault.recipient(), address(this));

        vault.transferTo(to, 10 ether);
        assertEq(address(to).balance, 10 ether);
        assertEq(address(vault).balance, 10 ether);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.Transfer(5 ether, vault.recipient(), address(this));

        vault.transferTo(to, 5 ether);
        assertEq(address(to).balance, 15 ether);
        assertEq(address(vault).balance, 5 ether);
    }

    function test_updateTransferAllowedStatus_notOwner_reverts() public {
        address[] memory allowedTransfer = new address[](1);
        allowedTransfer[0] = address(1);
        hevm.prank(address(123));
        hevm.expectRevert("caller is not the owner");
        vault.updateTransferAllowedStatus(allowedTransfer, true);
    }

    function test_updateTransferAllowedStatus_succeeds() public {
        address[] memory allowedTransfer = new address[](1);
        allowedTransfer[0] = address(123);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateTransferAllowed(allowedTransfer[0], true);

        vault.updateTransferAllowedStatus(allowedTransfer, true);
        assertTrue(vault.transferAllowed(address(123)));
    }

    function test_updateTransferAllowedStatus_withZeroAddress_succeeds() public {
        address[] memory allowedTransfer = new address[](1);
        allowedTransfer[0] = address(0);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateTransferAllowed(allowedTransfer[0], true);

        vault.updateTransferAllowedStatus(allowedTransfer, true);
        assertTrue(vault.transferAllowed(address(0)));
    }

    function test_updateReceiveAllowed_notOwner_reverts() public {
        address[] memory allowedReceive = new address[](1);
        allowedReceive[0] = address(1);
        hevm.prank(address(123));
        hevm.expectRevert("caller is not the owner");
        vault.updateReceiveAllowed(allowedReceive, true);
    }

    function test_updateReceiveAllowed_withZeroAddress_reverts() public {
        address[] memory allowedReceive = new address[](1);
        allowedReceive[0] = address(0);

        hevm.expectRevert("FeeVault: address cannot be address(0)");
        vault.updateReceiveAllowed(allowedReceive, true);
    }

    function test_updateReceiveAllowed_succeeds() public {
        address[] memory allowedReceive = new address[](1);
        allowedReceive[0] = address(123);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateReceiveAllowed(allowedReceive[0], true);

        vault.updateReceiveAllowed(allowedReceive, true);
        assertTrue(vault.receiveAllowed(address(123)));
        assertTrue(vault.isReceiveAllowed(address(123)));
    }

    function test_updateMessenger_notOwner_reverts() public {
        address oldMessenger = vault.messenger();
        address newMessenger = address(1);
        hevm.prank(address(123));
        hevm.expectRevert("caller is not the owner");
        vault.updateMessenger(newMessenger);
    }

    function test_updateMessenger_succeeds() public {
        address oldMessenger = vault.messenger();
        address newMessenger = address(1);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateMessenger(oldMessenger, newMessenger);

        vault.updateMessenger(newMessenger);
        assertEq(vault.messenger(), newMessenger);
    }

    function test_updateRecipient_notOwner_reverts() public {
        address oldRecipient = vault.messenger();
        address newRecipient = address(1);
        hevm.prank(address(123));
        hevm.expectRevert("caller is not the owner");
        vault.updateRecipient(newRecipient);
    }

    function test_updateRecipient_succeeds() public {
        address oldRecipient = vault.recipient();
        address newRecipient = address(1);

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateRecipient(oldRecipient, newRecipient);

        vault.updateRecipient(newRecipient);
        assertEq(vault.recipient(), newRecipient);
    }

    function test_updateMinWithdrawAmount_notOwner_reverts() public {
        uint256 oldAmount = vault.minWithdrawAmount();
        uint256 newAmount = 20 ether;
        hevm.prank(address(123));
        hevm.expectRevert("caller is not the owner");
        vault.updateMinWithdrawAmount(newAmount);
    }

    function test_updateMinWithdrawAmount_succeeds() public {
        uint256 oldAmount = vault.minWithdrawAmount();
        uint256 newAmount = 20 ether;

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateMinWithdrawAmount(oldAmount, newAmount);

        vault.updateMinWithdrawAmount(newAmount);
        assertEq(vault.minWithdrawAmount(), newAmount);
    }

    function test_updateMinWithdrawAmount_zeroMinWithdrawAmount_succeeds() public {
        uint256 oldAmount = vault.minWithdrawAmount();
        uint256 newAmount = 0 ether;

        hevm.expectEmit(true, true, true, true);
        emit L2TxFeeVault.UpdateMinWithdrawAmount(oldAmount, newAmount);

        vault.updateMinWithdrawAmount(newAmount);
        assertEq(vault.minWithdrawAmount(), newAmount);
    }
}
