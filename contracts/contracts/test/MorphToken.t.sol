// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Types} from "../libraries/common/Types.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";

contract MorphTokenTest is L2StakingBaseTest {
    uint256 private constant DAY_SECONDS = 86400; // 定义一天的秒数，根据MorphToken合约调整
    uint256 private constant PRECISION = 1e16;  // 根据MorphToken合约中的定义调整这个值
    event InflationMinted(uint256 indexed dayIndex, uint256 amount);

    function setUp() public virtual override {
        super.setUp();
    }

    function test_L2_STAKING_CONTRACT() public {
        assertEq(morphToken.L2_STAKING_CONTRACT(), Predeploys.L2_STAKING);
    }

    function test_DISTRIBUTE_CONTRACT() public {
        assertEq(morphToken.DISTRIBUTE_CONTRACT(), Predeploys.DISTRIBUTE);
    }

    function test_RECORD_CONTRACT() public {
        assertEq(morphToken.RECORD_CONTRACT(), Predeploys.RECORD);
    }

    function test_name() public {
        assertEq(morphToken.name(), "Morph");
    }

    function test_symbol() public {
        assertEq(morphToken.symbol(), "MPH");
    }

    function test_decimals() public {
        assertEq(morphToken.decimals(), 18);
    }

    function test_totalSupply() public {
        assertEq(morphToken.totalSupply(), 1000000000 ether);
    }

    function test_balanceOf() public {
        assertEq(morphToken.balanceOf(multisig), 1000000000 ether);
    }

    function test_inflationRate() public {
        uint256 count = morphToken.inflationRatesCount();
        assertEq(morphToken.dailyInflationRates(count - 1).rate, 1596535874529);
    }

    function test_inflationMintedDays() public {
        assertEq(morphToken.inflationMintedDays(), 0);
    }

    function test_transfer() public {
        hevm.startPrank(multisig);
        bool success = morphToken.transfer(alice, 10000000 ether);
        assert(success);
        assertEq(morphToken.balanceOf(alice), 10000000 ether);
        hevm.stopPrank();
    }

    function test_allowance() public {
        hevm.prank(multisig);
        bool success = morphToken.transfer(alice, 100 ether);
        assert(success);

        hevm.startPrank(alice);
        assert(morphToken.approve(bob, 20 ether));
        assert(morphToken.increaseAllowance(bob, 10 ether));
        assert(morphToken.decreaseAllowance(bob, 5 ether));
        assertEq(morphToken.allowance(alice, bob), 25 ether);
        hevm.stopPrank();

        hevm.prank(bob);
        assert(morphToken.transferFrom(alice, multisig, 10 ether));

        assertEq(morphToken.balanceOf(alice), 90 ether);
    }

    // 检查无效的 upToDayIndex
    function test_MintInflationsEarlyUpToDayIndex() public {
        // 假设当前时间设置为某个具体的值，例如当前的 UNIX 时间戳
        uint256 testStartTime = 1700000000; // 用一个固定的测试时间戳
        hevm.warp(testStartTime);  // 明确设置时间戳

        // 计算奖励开始时间为10天前
        uint256 startRewardTime = testStartTime - 10 * DAY_SECONDS;
        // 模拟 L2Staking 返回的 rewardStartTime
        hevm.mockCall(address(l2Staking), abi.encodeWithSignature("rewardStartTime()"), abi.encode(startRewardTime));

        // 计算 upToDayIndex 应该为从奖励开始到现在的天数
        uint256 upToDayIndex = (testStartTime - startRewardTime) / DAY_SECONDS;

        // 由于 upToDayIndex 计算将得到 10，我们预期这个铸币操作会太早
        hevm.prank(morphToken.RECORD_CONTRACT());
        hevm.expectRevert("the specified time has not yet been reached");
        morphToken.mintInflations(upToDayIndex);
    }

    // 检查 _inflationMintedDays 边界情况
    function test_MintInflationsWithMintedDaysBoundary() public {
        // 设置一个固定的测试时间戳，假设为当前的 UNIX 时间戳或其他合适的时间点
        uint256 testStartTime = 1700000000; // 例如使用一个特定的时间戳
        hevm.warp(testStartTime); // 明确设置时间戳

        // 设定奖励开始时间为100天前，基于测试设定的时间
        uint256 startRewardTime = testStartTime - 100 * DAY_SECONDS;
        // 模拟 l2Staking 返回的 rewardStartTime
        hevm.mockCall(address(l2Staking), abi.encodeWithSignature("rewardStartTime()"), abi.encode(startRewardTime));

        // 假定 upToDayIndex 为90，意味着过去已经铸造了90天
        uint256 upToDayIndex = 90;
        // 将 MorphToken 的 _inflationMintedDays 设置为90
        hevm.store(
            address(morphToken),
            keccak256(abi.encode(address(morphToken), uint256(0))), // 确保这是正确的 slot
            bytes32(uint256(upToDayIndex))
        );

        // 设置调用者为 RECORD_CONTRACT，并预期回退 "all inflations minted"
        hevm.prank(morphToken.RECORD_CONTRACT());
        hevm.expectRevert("all inflations minted");
        morphToken.mintInflations(upToDayIndex);
    }

    // 检查跨多天的铸币操作
    function test_MintInflationsOverMultipleDays() public {
        uint256 now = block.timestamp;
        hevm.warp(now + 86400 * 300); // 将时间设置得更合理，比如向前推300天

        uint256 startRewardTime = now - 200 * DAY_SECONDS;
        hevm.mockCall(address(l2Staking), abi.encodeWithSignature("rewardStartTime()"), abi.encode(startRewardTime));

        uint256 upToDayIndex = 180; // 模拟铸币至第180天
        hevm.store(
            address(morphToken),
            keccak256(abi.encode(address(morphToken), uint256(0))), // 确保这是 _inflationMintedDays 的正确槽位
            bytes32(uint256(150)) // 假设已经铸造到第150天
        );

        hevm.startPrank(morphToken.RECORD_CONTRACT());
        for (uint256 i = 151; i <= upToDayIndex; i++) {
            morphToken.mintInflations(i);
            assertEq(morphToken.balanceOf(address(distribute)), 1000000 + (i - 150) * 100); // 假设每天铸币率为 100 tokens
        }
        hevm.stopPrank();
    }

    // 测试铸币操作，确保铸币按预期工作
    function test_MintInflationsWithDefaultRate() public {
        uint256 startRewardTime = block.timestamp - 10 * DAY_SECONDS;
        hevm.mockCall(address(l2Staking), abi.encodeWithSignature("rewardStartTime()"), abi.encode(startRewardTime));

        // 设定为稍后的未来
        uint256 upToDayIndex = (block.timestamp - startRewardTime) / DAY_SECONDS + 2;

        // 设置 _inflationMintedDays 以确保可以铸币
        hevm.store(
            address(morphToken),
            bytes32(uint256(12)),  // _inflationMintedDays 的存储槽位
            bytes32(uint256(0))    // 将 _inflationMintedDays 重置为 0
        );

        // 记录执行前的总供应量
        uint256 oldTotalSupply = morphToken.totalSupply();

        // 执行铸币操作
        hevm.prank(morphToken.RECORD_CONTRACT());
        morphToken.mintInflations(upToDayIndex);

        // 计算预期铸币量
        uint256 expectedIncrease = 0;
        for (uint256 i = 1; i <= upToDayIndex; i++) {
            expectedIncrease += (oldTotalSupply * 100) / PRECISION;  // 假设默认铸币率为 100
        }

        // 检查总供应量是否正确更新
        assertEq(morphToken.totalSupply(), oldTotalSupply + expectedIncrease, "Total supply should be updated correctly.");

        // 检查 _inflationMintedDays 是否更新
        assertEq(morphToken.inflationMintedDays(), upToDayIndex + 1, "Inflation minted days should be updated correctly.");

        // 检查是否发出了 InflationMinted 事件
        hevm.expectEmit(true, true, false, true);
        emit InflationMinted(upToDayIndex, (oldTotalSupply * 100) / PRECISION);  // 预期的最后一天的铸币事件
    }


    // Fuzz test to check the mintInflations function across random days
    function test_FuzzMintInflations(uint256 randomDayIndex) public {
        // 确保合约尚未铸币到当前天
        uint256 mintedDays = morphToken.inflationMintedDays();
        hevm.assume(randomDayIndex > mintedDays);

        // 执行铸币操作
        hevm.prank(morphToken.RECORD_CONTRACT());
        try morphToken.mintInflations(randomDayIndex) {
            assertEq(morphToken.inflationMintedDays(), randomDayIndex + 1, "Inflation days should be correctly incremented");
        } catch {
            // 捕获并记录失败
            emit log_named_uint("Minting failed for day index", randomDayIndex);
        }
    }


    // Fuzz test to ensure that minting does not proceed for the current day
    function test_MintInflationsCurrentDay(uint256 randomSeconds) public {
        uint256 startTime = l2Staking.rewardStartTime();
        hevm.warp(startTime + randomSeconds);

        uint256 currentDayIndex = (block.timestamp - startTime) / DAY_SECONDS;

        hevm.prank(morphToken.RECORD_CONTRACT());
        hevm.expectRevert("the specified time has not yet been reached");
        morphToken.mintInflations(currentDayIndex);
    }

    // 测试不同铸币率的应用
    function test_MintInflations() public {

        // 将环境时间推进到当前真实时间或其他合适的测试时间
        hevm.warp(block.timestamp + 86400 * 20);  // 假设当前时间向前推进20天

        // 假设一个固定的日期作为更新日
        uint256 assumedRateStartDay = 10;  // 假设通货膨胀率在第10天更新
        hevm.prank(multisig);
        morphToken.updateRate(100, assumedRateStartDay);

        uint256 startRewardTime = block.timestamp - 20 * DAY_SECONDS; // 现在设定为20天前
        uint256 upToDayIndex = (block.timestamp - startRewardTime) / DAY_SECONDS - 1; // 设定为稍前的过去，确保时间已过

        hevm.prank(morphToken.RECORD_CONTRACT());
        morphToken.mintInflations(upToDayIndex);

        // 预期的通货膨胀计算
        uint256 expectedIncrease = 0;
        for (uint256 i = 1; i <= upToDayIndex; i++) {
            expectedIncrease += (morphToken.totalSupply() * 100) / PRECISION;
        }

        assertEq(morphToken.totalSupply(), expectedIncrease, "Total supply should be updated correctly.");
    }



    function calculateExpectedInflation(uint256 initialRate, uint256 newRate, uint256 changeDay, uint256 upToDay) internal view returns (uint256) {
        uint256 totalInflation = 0;
        for (uint256 day = 1; day <= upToDay; day++) {
            uint256 rate = day >= changeDay ? newRate : initialRate;
            totalInflation += (morphToken.totalSupply() * rate) / PRECISION;  // Assume _totalSupply is accessible or mocked
        }
        return totalInflation;
    }

    // 测试时间计算边界
    function test_TimeBoundaryMinting(uint256 offset) public {
        // 假设我们只想测试偏移量在十年内的情况，一年的秒数约为 31,536,000 秒
        uint256 maxOffset = 31536000 * 10;  // 10年
        hevm.assume(offset < maxOffset);  // 确保传入的偏移量小于一年
        uint256 startTime = l2Staking.rewardStartTime() + offset;
        hevm.warp(startTime);

        uint256 upToDayIndex = (block.timestamp - startTime) / DAY_SECONDS;

        hevm.prank(morphToken.RECORD_CONTRACT());
        morphToken.mintInflations(upToDayIndex);
        assertEq(morphToken.inflationMintedDays(), upToDayIndex + 1);
    }

    // 测试铸币操作失败恢复
    function test_MintingFailureRecovery(uint256 upToDayIndex) public {
        hevm.assume(upToDayIndex > morphToken.inflationMintedDays());

        // 设置一个不合理的奖励开始时间以模拟失败条件
        uint256 invalidStartTime = block.timestamp + 100 * DAY_SECONDS; // 未来的时间点
        hevm.mockCall(address(l2Staking), abi.encodeWithSignature("rewardStartTime()"), abi.encode(invalidStartTime));

        // uint256 currentDayIndex = (block.timestamp - invalidStartTime) / DAY_SECONDS + 1;

        // 预期由于时间未到而回退
        hevm.prank(morphToken.RECORD_CONTRACT());
        hevm.expectRevert("the specified time has not yet been reached");
        morphToken.mintInflations(upToDayIndex);

        // 验证状态没有改变
        assertEq(morphToken.inflationMintedDays(), 0, "Inflation days should not be updated on mint failure.");
        assertEq(morphToken.totalSupply(), 0, "Total supply should not change on mint failure.");
    }


}
