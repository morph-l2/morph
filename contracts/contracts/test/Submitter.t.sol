// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ISubmitter} from "../l1/rollup/ISubmitter.sol";
import {Submitter} from "../l1/rollup/Submitter.sol";

contract DrainRollupMock {
    uint256 public lastCommittedBatchIndex;
    uint256 public lastFinalizedBatchIndex;

    function setBatchIndexes(uint256 committed, uint256 finalized) external {
        lastCommittedBatchIndex = committed;
        lastFinalizedBatchIndex = finalized;
    }

    function slash(ISubmitter submitterContract, address account) external returns (uint256) {
        return submitterContract.slash(account);
    }

    receive() external payable {}
}

contract RejectEther {
    receive() external payable {
        revert("reject ether");
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

    DrainRollupMock internal rollupMock;
    Submitter internal implementation;
    Submitter internal submitter;

    function setUp() public {
        vm.deal(alice, 100 ether);
        vm.deal(bob, 100 ether);
        rollupMock = new DrainRollupMock();
        implementation = new Submitter();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            proxyAdmin,
            abi.encodeCall(
                Submitter.initialize,
                (owner, address(rollupMock), MINIMUM_STAKE, CHALLENGE_DEPOSIT, REWARD_PERCENTAGE)
            )
        );
        submitter = Submitter(payable(address(proxy)));
    }

    function test_initializeAndImplementationProtection() public {
        assertEq(submitter.owner(), owner);
        assertEq(submitter.rollupContract(), address(rollupMock));
        assertEq(submitter.minimumStake(), MINIMUM_STAKE);
        assertEq(submitter.challengeDeposit(), CHALLENGE_DEPOSIT);
        assertEq(submitter.rewardPercentage(), REWARD_PERCENTAGE);

        vm.expectRevert("Initializable: contract is already initialized");
        implementation.initialize(owner, address(rollupMock), 1, 1, 1);
        vm.expectRevert("Initializable: contract is already initialized");
        submitter.initialize(owner, address(rollupMock), 1, 1, 1);
    }

    function test_registrationStakeAndDynamicMinimum() public {
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

        // A threshold increase must not prevent the account from exiting its existing stake.
        vm.prank(alice);
        submitter.withdraw();
        assertFalse(submitter.registered(alice));
        assertTrue(submitter.withdrawing(alice));
    }

    function test_withdrawWaitsForGlobalDrainAndRequiresOwnerToReauthorize() public {
        _registerAndStake(alice);
        rollupMock.setBatchIndexes(2, 1);

        vm.prank(alice);
        submitter.withdraw();
        assertFalse(submitter.isActive(alice));
        assertFalse(submitter.registered(alice));
        assertEq(submitter.stakeOf(alice), MINIMUM_STAKE);

        vm.prank(alice);
        vm.expectRevert("rollup not drained");
        submitter.claimWithdrawal(bob);

        rollupMock.setBatchIndexes(2, 2);
        uint256 beforeBalance = bob.balance;
        vm.prank(alice);
        submitter.claimWithdrawal(bob);
        assertEq(bob.balance - beforeBalance, MINIMUM_STAKE);
        assertEq(submitter.stakeOf(alice), 0);
        assertFalse(submitter.withdrawing(alice));

        vm.prank(alice);
        vm.expectRevert("not stakeable");
        submitter.stake{value: MINIMUM_STAKE}();
        vm.prank(owner);
        submitter.addSubmitter(alice);
    }

    function test_ownerRemovalUsesTheSameDrainGate() public {
        _registerAndStake(alice);
        rollupMock.setBatchIndexes(3, 2);

        vm.prank(owner);
        submitter.removeSubmitter(alice);
        assertFalse(submitter.registered(alice));
        assertTrue(submitter.withdrawing(alice));

        vm.prank(alice);
        vm.expectRevert("rollup not drained");
        submitter.claimWithdrawal(alice);
        assertEq(submitter.stakeOf(alice), MINIMUM_STAKE);

        rollupMock.setBatchIndexes(3, 3);
        vm.prank(alice);
        submitter.claimWithdrawal(alice);
        assertEq(submitter.stakeOf(alice), 0);
    }

    function test_exitingStakeRemainsSlashable() public {
        _registerAndStake(alice);
        rollupMock.setBatchIndexes(2, 1);
        vm.prank(alice);
        submitter.withdraw();

        uint256 rollupBalance = address(rollupMock).balance;
        uint256 reward = rollupMock.slash(submitter, alice);
        assertEq(reward, MINIMUM_STAKE / 4);
        assertEq(address(rollupMock).balance - rollupBalance, reward);
        assertEq(submitter.slashRemaining(), MINIMUM_STAKE - reward);
        assertEq(submitter.stakeOf(alice), 0);
        assertFalse(submitter.registered(alice));
        assertFalse(submitter.withdrawing(alice));

        vm.prank(owner);
        submitter.addSubmitter(alice);
        vm.prank(alice);
        submitter.stake{value: MINIMUM_STAKE}();
        assertTrue(submitter.isActive(alice));
    }

    function test_slashAccessZeroAddressAndRemainingClaim() public {
        _registerAndStake(alice);
        vm.prank(alice);
        vm.expectRevert("only rollup contract");
        submitter.slash(alice);

        assertEq(rollupMock.slash(submitter, address(0)), 0);
        rollupMock.slash(submitter, alice);

        uint256 beforeBalance = bob.balance;
        vm.prank(owner);
        submitter.claimSlashRemaining(bob);
        assertEq(bob.balance - beforeBalance, (MINIMUM_STAKE * 75) / 100);
    }

    function test_failedWithdrawalTransferPreservesStakeAndExit() public {
        _registerAndStake(alice);
        vm.prank(alice);
        submitter.withdraw();
        RejectEther receiver = new RejectEther();

        vm.prank(alice);
        vm.expectRevert("ETH transfer failed");
        submitter.claimWithdrawal(address(receiver));
        assertEq(submitter.stakeOf(alice), MINIMUM_STAKE);
        assertTrue(submitter.withdrawing(alice));
    }

    function _registerAndStake(address account) internal {
        vm.prank(owner);
        submitter.addSubmitter(account);
        vm.prank(account);
        submitter.stake{value: MINIMUM_STAKE}();
    }
}
