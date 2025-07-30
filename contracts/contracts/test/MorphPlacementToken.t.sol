// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import {CommonTest} from "./base/CommonTest.t.sol";
import {IMorphPlacementToken} from "../l2/token/IMorphPlacementToken.sol";
import {MorphPlacementToken} from "../l2/token/MorphPlacementToken.sol";

contract MorphPlacementTokenTest is CommonTest {
    MorphPlacementToken public morphPlacementToken;
    TransparentUpgradeableProxy public morphPlacementTokenProxy;
    
    string public constant TOKEN_NAME = "Morph Placement Token";
    string public constant TOKEN_SYMBOL = "MPHP";
    uint256 public constant INITIAL_SUPPLY = 10000000000 ether;
    address public charlie = address(1000);
    
    function setUp() public virtual override {
        super.setUp();
        
        // Deploy implementation contract
        MorphPlacementToken implementation = new MorphPlacementToken();
        
        // Deploy proxy contract with empty contract as initial implementation
        morphPlacementTokenProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        

        // Initialize the proxy with the implementation，owner为alice
        hevm.startPrank(multisig);
        ITransparentUpgradeableProxy(address(morphPlacementTokenProxy)).upgradeToAndCall(
            address(implementation),
            abi.encodeCall(MorphPlacementToken.initialize, (TOKEN_NAME, TOKEN_SYMBOL, alice, INITIAL_SUPPLY))
        );
        hevm.stopPrank();
        
        morphPlacementToken = MorphPlacementToken(address(morphPlacementTokenProxy));
    }

    /************************
     * Initialization Tests *
     ************************/

    function test_initialize_succeeds() public {
        // Deploy a new implementation and proxy for testing
        MorphPlacementToken implementation = new MorphPlacementToken();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            address(multisig),
            new bytes(0)
        );
        
        hevm.startPrank(alice);
        MorphPlacementToken token = MorphPlacementToken(address(proxy));
        token.initialize("Test Token", "TEST", alice, INITIAL_SUPPLY);
        hevm.stopPrank();
        
        assertEq(token.name(), "Test Token");
        assertEq(token.symbol(), "TEST");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.balanceOf(alice), INITIAL_SUPPLY);
        assertEq(token.owner(), alice);
    }

    function test_initialize_initializeAgain_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("Initializable: contract is already initialized");
        morphPlacementToken.initialize("Test", "TEST", alice, 1000000 ether);
        hevm.stopPrank();
    }

    function test_initialize_zeroOwner_reverts() public {
        MorphPlacementToken implementation = new MorphPlacementToken();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            address(multisig),
            new bytes(0)
        );
        
        hevm.startPrank(alice);
        MorphPlacementToken token = MorphPlacementToken(address(proxy));
        hevm.expectRevert("mint to the zero address");
        token.initialize("Test", "TEST", address(0), 1000000 ether);
        hevm.stopPrank();
    }

    /************************
     * Basic ERC20 Tests *
     ************************/

    function test_name_succeeds() public {
        assertEq(morphPlacementToken.name(), TOKEN_NAME);
    }

    function test_symbol_succeeds() public {
        assertEq(morphPlacementToken.symbol(), TOKEN_SYMBOL);
    }

    function test_decimals_succeeds() public {
        assertEq(morphPlacementToken.decimals(), 18);
    }

    function test_totalSupply_succeeds() public {
        assertEq(morphPlacementToken.totalSupply(), INITIAL_SUPPLY);
    }

    function test_balanceOf_succeeds() public {
        assertEq(morphPlacementToken.balanceOf(alice), INITIAL_SUPPLY);
        assertEq(morphPlacementToken.balanceOf(multisig), 0);
    }

    /************************
     * Transfer Tests *
     ************************/

    function test_transfer_succeeds() public {
        uint256 transferAmount = 1000 ether;
        
        hevm.startPrank(alice);
        bool success = morphPlacementToken.transfer(multisig, transferAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.balanceOf(alice), INITIAL_SUPPLY - transferAmount);
        assertEq(morphPlacementToken.balanceOf(multisig), transferAmount);
    }

    function test_transfer_insufficientBalance_reverts() public {
        hevm.startPrank(bob);
        hevm.expectRevert("transfer amount exceeds balance");
        morphPlacementToken.transfer(charlie, 1000 ether);
        hevm.stopPrank();
    }

    function test_transfer_toZeroAddress_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("transfer to the zero address");
        morphPlacementToken.transfer(address(0), 1000 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_succeeds() public {
        uint256 transferAmount = 1000 ether;
        
        // First approve
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, transferAmount);
        hevm.stopPrank();
        
        // Then transferFrom
        hevm.startPrank(bob);
        bool success = morphPlacementToken.transferFrom(alice, charlie, transferAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.balanceOf(alice), INITIAL_SUPPLY - transferAmount);
        assertEq(morphPlacementToken.balanceOf(charlie), transferAmount);
        assertEq(morphPlacementToken.allowance(alice, bob), 0);
    }

    function test_transferFrom_insufficientAllowance_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("insufficient allowance");
        morphPlacementToken.transferFrom(multisig, bob, 1000 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_insufficientBalance_reverts() public {
        // Give bob some tokens
        hevm.startPrank(alice);
        morphPlacementToken.transfer(bob, 100 ether);
        hevm.stopPrank();
        
        // Approve more than balance
        hevm.startPrank(bob);
        morphPlacementToken.approve(charlie, 200 ether);
        hevm.stopPrank();
        
        // Try to transfer more than balance
        hevm.startPrank(charlie);
        hevm.expectRevert("transfer amount exceeds balance");
        morphPlacementToken.transferFrom(bob, alice, 200 ether);
        hevm.stopPrank();
    }

    /************************
     * Approval Tests *
     ************************/

    function test_approve_succeeds() public {
        uint256 approveAmount = 1000 ether;
        
        hevm.startPrank(alice);
        bool success = morphPlacementToken.approve(bob, approveAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.allowance(alice, bob), approveAmount);
    }

    function test_approve_zeroSpender_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("approve to the zero address");
        morphPlacementToken.approve(address(0), 1000 ether);
        hevm.stopPrank();
    }

    function test_increaseAllowance_succeeds() public {
        uint256 initialAmount = 1000 ether;
        uint256 increaseAmount = 500 ether;
        
        // First approve
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, initialAmount);
        hevm.stopPrank();
        
        // Then increase allowance
        hevm.startPrank(alice);
        bool success = morphPlacementToken.increaseAllowance(bob, increaseAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.allowance(alice, bob), initialAmount + increaseAmount);
    }

    function test_decreaseAllowance_succeeds() public {
        uint256 initialAmount = 1000 ether;
        uint256 decreaseAmount = 300 ether;
        
        // First approve
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, initialAmount);
        hevm.stopPrank();
        
        // Then decrease allowance
        hevm.startPrank(alice);
        bool success = morphPlacementToken.decreaseAllowance(bob, decreaseAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.allowance(alice, bob), initialAmount - decreaseAmount);
    }

    function test_decreaseAllowance_insufficientAllowance_reverts() public {
        uint256 initialAmount = 1000 ether;
        uint256 decreaseAmount = 1500 ether;
        
        // First approve
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, initialAmount);
        hevm.stopPrank();
        
        // Then try to decrease more than allowed
        hevm.startPrank(alice);
        hevm.expectRevert("decreased allowance below zero");
        morphPlacementToken.decreaseAllowance(bob, decreaseAmount);
        hevm.stopPrank();
    }

    /************************
     * Burn Tests *
     ************************/

    function test_burn_succeeds() public {
        uint256 burnAmount = 1000 ether;
        uint256 initialBalance = morphPlacementToken.balanceOf(alice);
        uint256 initialSupply = morphPlacementToken.totalSupply();
        
        hevm.startPrank(alice);
        morphPlacementToken.burn(burnAmount);
        hevm.stopPrank();
        
        assertEq(morphPlacementToken.balanceOf(alice), initialBalance - burnAmount);
        assertEq(morphPlacementToken.totalSupply(), initialSupply - burnAmount);
    }

    function test_burn_notOwner_reverts() public {
        hevm.startPrank(bob);
        hevm.expectRevert("Ownable: caller is not the owner");
        morphPlacementToken.burn(1000 ether);
        hevm.stopPrank();
    }

    function test_burn_zeroAmount_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("amount to burn is zero");
        morphPlacementToken.burn(0);
        hevm.stopPrank();
    }

    function test_burn_insufficientBalance_reverts() public {
        uint256 totalBalance = morphPlacementToken.balanceOf(alice);
        
        hevm.startPrank(alice);
        hevm.expectRevert("ERC20: burn amount exceeds balance");
        morphPlacementToken.burn(totalBalance + 1);
        hevm.stopPrank();
    }

    /************************
     * Pausable Tests *
     ************************/

    function test_setPause_true_succeeds() public {
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        assertTrue(morphPlacementToken.paused());
    }

    function test_setPause_false_succeeds() public {
        // First pause
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        assertTrue(morphPlacementToken.paused());
        
        // Then unpause
        morphPlacementToken.setPause(false);
        hevm.stopPrank();
        
        assertFalse(morphPlacementToken.paused());
    }

    function test_setPause_notOwner_reverts() public {
        hevm.startPrank(bob);
        hevm.expectRevert("Ownable: caller is not the owner");
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
    }

    function test_transfer_whenPaused_reverts() public {
        // Pause the contract
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        // Try to transfer
        hevm.startPrank(alice);
        hevm.expectRevert("Pausable: paused");
        morphPlacementToken.transfer(bob, 1000 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_whenPaused_reverts() public {
        // Pause the contract
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, 1000 ether);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        // Try to transferFrom
        hevm.startPrank(bob);
        hevm.expectRevert("Pausable: paused");
        morphPlacementToken.transferFrom(alice, charlie, 1000 ether);
        hevm.stopPrank();
    }

    function test_approve_whenPaused_reverts() public {
        // Pause the contract
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        // Try to approve
        hevm.startPrank(alice);
        hevm.expectRevert("Pausable: paused");
        morphPlacementToken.approve(bob, 1000 ether);
        hevm.stopPrank();
    }

    function test_burn_whenPaused_reverts() public {
        // Pause the contract
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        // Try to burn
        hevm.startPrank(alice);
        hevm.expectRevert("Pausable: paused");
        morphPlacementToken.burn(1000 ether);
        hevm.stopPrank();
    }

    /************************
     * Events Tests *
     ************************/

    function test_transfer_emitsTransferEvent() public {
        uint256 transferAmount = 1000 ether;
        
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, false, true);
        emit IERC20Upgradeable.Transfer(alice, bob, transferAmount);
        morphPlacementToken.transfer(bob, transferAmount);
        hevm.stopPrank();
    }

    function test_approve_emitsApprovalEvent() public {
        uint256 approveAmount = 1000 ether;
        
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, false, true);
        emit IERC20Upgradeable.Approval(alice, bob, approveAmount);
        morphPlacementToken.approve(bob, approveAmount);
        hevm.stopPrank();
    }

    function test_burn_emitsTransferEvent() public {
        uint256 burnAmount = 1000 ether;
        
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, false, true);
        emit IERC20Upgradeable.Transfer(alice, address(0), burnAmount);
        morphPlacementToken.burn(burnAmount);
        hevm.stopPrank();
    }

    /************************
     * Edge Cases Tests *
     ************************/

    function test_transfer_maxUint256_succeeds() public {
        uint256 maxAmount =1000 ether;
        
        // Give bob max tokens
        hevm.startPrank(alice);
        morphPlacementToken.transfer(bob, maxAmount);
        hevm.stopPrank();
        
        // Transfer back
        hevm.startPrank(bob);
        bool success = morphPlacementToken.transfer(alice, maxAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.balanceOf(bob), 0);
        assertEq(morphPlacementToken.balanceOf(alice), INITIAL_SUPPLY);
    }

    function test_approve_maxUint256_succeeds() public {
        uint256 maxAmount = type(uint256).max;
        
        hevm.startPrank(alice);
        bool success = morphPlacementToken.approve(bob, maxAmount);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.allowance(alice, bob), maxAmount);
    }

    function test_transferFrom_maxAllowance_succeeds() public {
        uint256 transferAmount = 1000 ether;
        
        // Approve max amount
        hevm.startPrank(alice);
        morphPlacementToken.approve(bob, type(uint256).max);
        hevm.stopPrank();
        
        // Transfer multiple times without reducing allowance
        hevm.startPrank(bob);
        morphPlacementToken.transferFrom(alice, charlie, transferAmount);
        morphPlacementToken.transferFrom(alice, charlie, transferAmount);
        hevm.stopPrank();
        
        // Allowance should still be max
        assertEq(morphPlacementToken.allowance(alice, bob), type(uint256).max);
    }

    /************************
     * Integration Tests *
     ************************/

    function test_completeTransferFlow_succeeds() public {
        uint256 transferAmount = 1000 ether;
        
        // 1. Transfer tokens to bob
        hevm.startPrank(alice);
        morphPlacementToken.transfer(bob, transferAmount);
        hevm.stopPrank();
        
        // 2. Bob approves charlie to spend his tokens
        hevm.startPrank(bob);
        morphPlacementToken.approve(charlie, transferAmount);
        hevm.stopPrank();
        
        // 3. Charlie transfers bob's tokens to alice
        hevm.startPrank(charlie);
        morphPlacementToken.transferFrom(bob, alice, transferAmount);
        hevm.stopPrank();
        
        // Verify final balances
        assertEq(morphPlacementToken.balanceOf(alice), INITIAL_SUPPLY);
        assertEq(morphPlacementToken.balanceOf(bob), 0);
        assertEq(morphPlacementToken.balanceOf(charlie), 0);
        assertEq(morphPlacementToken.allowance(bob, charlie), 0);
    }

    function test_pauseUnpauseFlow_succeeds() public {
        uint256 transferAmount = 1000 ether;
        
        // 1. Transfer tokens to bob
        hevm.startPrank(alice);
        morphPlacementToken.transfer(bob, transferAmount);
        hevm.stopPrank();
        
        // 2. Pause the contract
        hevm.startPrank(alice);
        morphPlacementToken.setPause(true);
        hevm.stopPrank();
        
        // 3. Verify transfers are blocked
        hevm.startPrank(bob);
        hevm.expectRevert("Pausable: paused");
        morphPlacementToken.transfer(charlie, 100 ether);
        hevm.stopPrank();
        
        // 4. Unpause the contract
        hevm.startPrank(alice);
        morphPlacementToken.setPause(false);
        hevm.stopPrank();
        
        // 5. Verify transfers work again
        hevm.startPrank(bob);
        bool success = morphPlacementToken.transfer(charlie, 100 ether);
        hevm.stopPrank();
        
        assertTrue(success);
        assertEq(morphPlacementToken.balanceOf(charlie), 100 ether);
    }
} 