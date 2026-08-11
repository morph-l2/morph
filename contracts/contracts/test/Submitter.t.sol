// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ISubmitter} from "../l1/rollup/ISubmitter.sol";
import {Submitter} from "../l1/rollup/Submitter.sol";

contract LiabilityRollupMock {
    enum ReturnMode {
        Normal,
        RevertCall,
        Short,
        Long,
        BurnGas
    }

    mapping(address submitter => uint256 count) public counts;
    ReturnMode public returnMode;

    function setPendingBatchCount(address submitter, uint256 count) external {
        counts[submitter] = count;
    }

    function setReturnMode(ReturnMode mode) external {
        returnMode = mode;
    }

    function slash(ISubmitter submitterContract, address submitter) external returns (uint256) {
        return submitterContract.slash(submitter);
    }

    receive() external payable {}

    fallback() external payable {
        require(msg.sig == bytes4(keccak256("pendingBatchCount(address)")), "unknown selector");
        ReturnMode mode = returnMode;
        address submitter;
        assembly {
            submitter := calldataload(4)
        }
        uint256 count = counts[submitter];
        if (mode == ReturnMode.RevertCall) revert("getter reverted");
        if (mode == ReturnMode.BurnGas) {
            assembly {
                for {} 1 {} {}
            }
        }
        assembly {
            mstore(0, count)
        }
        if (mode == ReturnMode.Short) {
            assembly {
                return(0, 31)
            }
        }
        if (mode == ReturnMode.Long) {
            assembly {
                mstore(32, 1)
                return(0, 64)
            }
        }
        assembly {
            return(0, 32)
        }
    }
}

contract RejectEther {
    function stakeInto(Submitter target) external {
        target.stake{value: address(this).balance}();
    }

    function claimFrom(Submitter target, address receiver) external {
        target.claimCredit(receiver);
    }

    receive() external payable {
        revert("reject ether");
    }
}

contract BurnRefundGas {
    receive() external payable {
        assembly {
            for {} 1 {} {}
        }
    }
}

contract ReentrantSubmitterActor {
    Submitter internal immutable target;
    bool public reentrySucceeded;
    bytes32 public reentryErrorHash;

    constructor(Submitter target_) {
        target = target_;
    }

    function stakeInto() external {
        target.stake{value: address(this).balance}();
    }

    function withdrawToSelf() external {
        target.withdraw(address(this));
    }

    receive() external payable {
        try target.claimWithdrawal(address(this)) {
            reentrySucceeded = true;
        } catch (bytes memory reason) {
            reentryErrorHash = keccak256(reason);
        }
    }
}

contract SubmitterTest is Test {
    uint256 internal constant MINIMUM_STAKE = 1 ether;
    uint256 internal constant CHALLENGE_DEPOSIT = 2 ether;
    uint256 internal constant REWARD_PERCENTAGE = 25;

    address internal owner = address(0xA11CE);
    address internal proxyAdmin = address(0xA11);
    address internal alice = address(0xB0B);
    address internal bob = address(0xCAFE);

    LiabilityRollupMock internal rollupMock;
    Submitter internal implementation;
    Submitter internal submitter;

    function setUp() public {
        vm.deal(alice, 100 ether);
        vm.deal(bob, 100 ether);
        rollupMock = new LiabilityRollupMock();
        implementation = new Submitter();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            proxyAdmin,
            abi.encodeCall(
                Submitter.initialize, (owner, address(rollupMock), MINIMUM_STAKE, CHALLENGE_DEPOSIT, REWARD_PERCENTAGE)
            )
        );
        submitter = Submitter(payable(address(proxy)));
    }

    function test_initialize_proxyAndImplementationProtection() public {
        assertEq(submitter.owner(), owner);
        assertEq(submitter.rollupContract(), address(rollupMock));
        assertEq(submitter.minimumStake(), MINIMUM_STAKE);
        assertEq(submitter.challengeDeposit(), CHALLENGE_DEPOSIT);
        assertEq(submitter.rewardPercentage(), REWARD_PERCENTAGE);

        vm.expectRevert("Initializable: contract is already initialized");
        implementation.initialize(owner, address(rollupMock), 1, 1, 1);

        vm.prank(owner);
        vm.expectRevert("Initializable: contract is already initialized");
        submitter.initialize(owner, address(rollupMock), 1, 1, 1);
    }

    function test_initialize_rejectsInvalidParameters() public {
        _expectProxyInitializationRevert(address(0), address(rollupMock), 1, 1, 1, "invalid owner");
        _expectProxyInitializationRevert(owner, address(0), 1, 1, 1, "invalid rollup contract");
        _expectProxyInitializationRevert(owner, alice, 1, 1, 1, "invalid rollup contract");
        _expectProxyInitializationRevert(owner, address(rollupMock), 0, 1, 1, "invalid minimum stake");
        _expectProxyInitializationRevert(owner, address(rollupMock), 1, 0, 1, "invalid challenge deposit");
        _expectProxyInitializationRevert(owner, address(rollupMock), 1, 1, 0, "invalid reward percentage");
        _expectProxyInitializationRevert(owner, address(rollupMock), 1, 1, 101, "invalid reward percentage");
    }

    function test_addStakeAndDynamicMinimum() public {
        vm.prank(alice);
        vm.expectRevert("Ownable: caller is not the owner");
        submitter.addSubmitter(alice);

        vm.prank(owner);
        submitter.addSubmitter(alice);
        assertFalse(submitter.isActive(alice));

        vm.prank(alice);
        vm.expectRevert("below minimum stake");
        submitter.stake{value: MINIMUM_STAKE - 1}();

        vm.prank(alice);
        submitter.stake{value: MINIMUM_STAKE}();
        assertTrue(submitter.isActive(alice));

        vm.prank(owner);
        submitter.setMinimumStake(2 ether);
        assertFalse(submitter.isActive(alice));

        vm.prank(alice);
        vm.expectRevert("below minimum stake");
        submitter.stake{value: 0.5 ether}();
        vm.prank(alice);
        submitter.stake{value: 1 ether}();
        assertTrue(submitter.isActive(alice));

        vm.prank(owner);
        submitter.setMinimumStake(3 ether);
        assertFalse(submitter.isActive(alice));
        vm.prank(owner);
        submitter.setMinimumStake(2 ether);
        assertTrue(submitter.isActive(alice));
    }

    function test_multipleSubmittersCanBeActive() public {
        _registerAndStake(alice, MINIMUM_STAKE);
        _registerAndStake(bob, MINIMUM_STAKE + 1);
        assertTrue(submitter.isActive(alice));
        assertTrue(submitter.isActive(bob));
    }

    function test_withdraw_noLiabilityPaysAndRequiresFreshRegistration() public {
        _registerAndStake(alice, MINIMUM_STAKE);
        uint256 before = bob.balance;

        vm.prank(alice);
        submitter.withdraw(bob);

        assertEq(bob.balance - before, MINIMUM_STAKE);
        assertEq(submitter.stakeOf(alice), 0);
        assertFalse(submitter.exiting(alice));
        assertFalse(submitter.registered(alice));
        vm.prank(alice);
        vm.expectRevert("not stakeable");
        submitter.stake{value: MINIMUM_STAKE}();
    }

    function test_remove_pendingExitRejectsDuplicateAddUntilClaimed() public {
        vm.prank(owner);
        submitter.addSubmitter(alice);
        vm.prank(owner);
        vm.expectRevert("invalid submitter");
        submitter.addSubmitter(alice);

        vm.prank(alice);
        submitter.stake{value: MINIMUM_STAKE}();
        rollupMock.setPendingBatchCount(alice, 1);
        vm.prank(owner);
        submitter.removeSubmitter(alice);

        assertFalse(submitter.registered(alice));
        assertTrue(submitter.exiting(alice));
        assertEq(submitter.stakeOf(alice), MINIMUM_STAKE);
        vm.prank(owner);
        vm.expectRevert("invalid submitter");
        submitter.addSubmitter(alice);
        vm.prank(alice);
        vm.expectRevert("pending batch liability");
        submitter.claimWithdrawal(alice);

        rollupMock.setPendingBatchCount(alice, 0);
        vm.prank(alice);
        submitter.claimWithdrawal(alice);
        vm.prank(owner);
        submitter.addSubmitter(alice);
        assertTrue(submitter.registered(alice));
    }

    function test_exit_keepsStakeSlashable_untilAllBatchesTerminate() public {
        _registerAndStake(alice, MINIMUM_STAKE);
        rollupMock.setPendingBatchCount(alice, 2);

        vm.prank(alice);
        submitter.withdraw(alice);
        assertFalse(submitter.isActive(alice));
        assertTrue(submitter.exiting(alice));
        assertEq(submitter.stakeOf(alice), MINIMUM_STAKE);

        uint256 reward = rollupMock.slash(submitter, alice);
        assertEq(reward, MINIMUM_STAKE / 4);
        assertEq(submitter.stakeOf(alice), 0);
        assertTrue(submitter.exiting(alice));
        assertFalse(submitter.registered(alice));

        rollupMock.setPendingBatchCount(alice, 0);
        vm.prank(alice);
        submitter.claimWithdrawal(alice);
        assertFalse(submitter.exiting(alice));
        vm.prank(alice);
        vm.expectRevert("not stakeable");
        submitter.stake{value: MINIMUM_STAKE}();

        vm.prank(owner);
        submitter.addSubmitter(alice);
        vm.prank(alice);
        submitter.stake{value: MINIMUM_STAKE}();
        assertTrue(submitter.isActive(alice));
    }

    function test_withdraw_getterFailuresFailClosed() public {
        LiabilityRollupMock.ReturnMode[4] memory modes = [
            LiabilityRollupMock.ReturnMode.RevertCall,
            LiabilityRollupMock.ReturnMode.Short,
            LiabilityRollupMock.ReturnMode.Long,
            LiabilityRollupMock.ReturnMode.BurnGas
        ];
        for (uint256 i; i < modes.length; ++i) {
            address actor = address(uint160(0x1000 + i));
            vm.deal(actor, 10 ether);
            _registerAndStake(actor, MINIMUM_STAKE);
            rollupMock.setReturnMode(modes[i]);
            vm.prank(actor);
            submitter.withdraw(actor);
            assertFalse(submitter.isActive(actor));
            assertTrue(submitter.exiting(actor));
            assertEq(submitter.stakeOf(actor), MINIMUM_STAKE);
            rollupMock.setReturnMode(LiabilityRollupMock.ReturnMode.Normal);
            vm.prank(actor);
            submitter.claimWithdrawal(actor);
        }
    }

    function test_receiverFailureCreditsWithoutBlockingRemove() public {
        RejectEther actor = new RejectEther();
        vm.deal(address(actor), MINIMUM_STAKE);
        vm.prank(owner);
        submitter.addSubmitter(address(actor));
        actor.stakeInto(submitter);

        vm.prank(owner);
        submitter.removeSubmitter(address(actor));
        assertFalse(submitter.registered(address(actor)));
        assertFalse(submitter.exiting(address(actor)));
        assertEq(submitter.claimable(address(actor)), MINIMUM_STAKE);

        uint256 before = bob.balance;
        actor.claimFrom(submitter, bob);
        assertEq(bob.balance - before, MINIMUM_STAKE);
        assertEq(submitter.claimable(address(actor)), 0);
    }

    function test_withdraw_revertingAndGasBurningReceiversBecomeCredit() public {
        RejectEther reject = new RejectEther();
        BurnRefundGas burner = new BurnRefundGas();

        _registerAndStake(alice, MINIMUM_STAKE);
        vm.prank(alice);
        submitter.withdraw(address(reject));
        assertEq(submitter.claimable(alice), MINIMUM_STAKE);

        _registerAndStake(bob, MINIMUM_STAKE);
        vm.prank(bob);
        submitter.withdraw(address(burner));
        assertEq(submitter.claimable(bob), MINIMUM_STAKE);
    }

    function test_withdraw_receiverReentryIsBlocked() public {
        ReentrantSubmitterActor actor = new ReentrantSubmitterActor(submitter);
        vm.deal(address(actor), MINIMUM_STAKE);
        vm.prank(owner);
        submitter.addSubmitter(address(actor));
        actor.stakeInto();

        actor.withdrawToSelf();

        assertFalse(actor.reentrySucceeded());
        assertEq(
            actor.reentryErrorHash(),
            keccak256(abi.encodeWithSignature("Error(string)", "ReentrancyGuard: reentrant call"))
        );
        assertEq(address(actor).balance, MINIMUM_STAKE);
        assertEq(submitter.stakeOf(address(actor)), 0);
    }

    function test_claimCreditFailureRollsBackCredit() public {
        RejectEther reject = new RejectEther();
        _registerAndStake(alice, MINIMUM_STAKE);
        vm.prank(alice);
        submitter.withdraw(address(reject));

        vm.prank(alice);
        vm.expectRevert("ETH transfer failed");
        submitter.claimCredit(address(reject));
        assertEq(submitter.claimable(alice), MINIMUM_STAKE);
    }

    function test_slashPermissionsZeroAndOwnerAccounting() public {
        _registerAndStake(alice, 4 ether);

        vm.prank(owner);
        vm.expectRevert("only rollup contract");
        submitter.slash(alice);
        assertEq(rollupMock.slash(submitter, address(0)), 0);

        uint256 rollupBefore = address(rollupMock).balance;
        assertEq(rollupMock.slash(submitter, alice), 1 ether);
        assertEq(address(rollupMock).balance - rollupBefore, 1 ether);
        assertEq(submitter.slashRemaining(), 3 ether);
        assertEq(address(submitter).balance, 3 ether);

        uint256 ownerBefore = owner.balance;
        vm.prank(owner);
        submitter.claimSlashRemaining(owner);
        assertEq(owner.balance - ownerBefore, 3 ether);
        assertEq(submitter.slashRemaining(), 0);
    }

    function test_parameterUpdatesValidateBoundsAndPermissions() public {
        vm.prank(alice);
        vm.expectRevert("Ownable: caller is not the owner");
        submitter.updateChallengeDeposit(3 ether);

        vm.startPrank(owner);
        vm.expectRevert("invalid minimum stake");
        submitter.setMinimumStake(0);
        vm.expectRevert("invalid challenge deposit");
        submitter.updateChallengeDeposit(CHALLENGE_DEPOSIT);
        vm.expectRevert("invalid reward percentage");
        submitter.updateRewardPercentage(101);
        submitter.updateChallengeDeposit(3 ether);
        submitter.updateRewardPercentage(50);
        vm.stopPrank();

        assertEq(submitter.challengeDeposit(), 3 ether);
        assertEq(submitter.rewardPercentage(), 50);
    }

    function testFuzz_accountingConservation(uint96 aliceStake, uint96 bobStake) public {
        aliceStake = uint96(bound(aliceStake, MINIMUM_STAKE, 20 ether));
        bobStake = uint96(bound(bobStake, MINIMUM_STAKE, 20 ether));
        _registerAndStake(alice, aliceStake);
        _registerAndStake(bob, bobStake);

        RejectEther reject = new RejectEther();
        vm.prank(alice);
        submitter.withdraw(address(reject));
        rollupMock.slash(submitter, bob);

        uint256 tracked = submitter.stakeOf(alice) + submitter.stakeOf(bob);
        tracked += submitter.claimable(alice) + submitter.claimable(bob);
        tracked += submitter.slashRemaining();
        assertLe(tracked, address(submitter).balance);
    }

    function _registerAndStake(address actor, uint256 amount) internal {
        vm.prank(owner);
        submitter.addSubmitter(actor);
        vm.prank(actor);
        submitter.stake{value: amount}();
    }

    function _expectProxyInitializationRevert(
        address owner_,
        address rollup_,
        uint256 minimumStake_,
        uint256 deposit_,
        uint256 percentage_,
        string memory reason
    ) internal {
        Submitter freshImplementation = new Submitter();
        vm.expectRevert(bytes(reason));
        new TransparentUpgradeableProxy(
            address(freshImplementation),
            proxyAdmin,
            abi.encodeCall(Submitter.initialize, (owner_, rollup_, minimumStake_, deposit_, percentage_))
        );
    }
}
