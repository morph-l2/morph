// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Test} from "forge-std/Test.sol";

import {RateLimiter} from "../libraries/common/RateLimiter.sol";

contract RateLimiterHarness {
    using RateLimiter for RateLimiter.TokenBucket;

    RateLimiter.TokenBucket internal bucket;

    function consume(uint256 _amount) external {
        bucket.consume(_amount);
    }

    function refund(uint256 _amount) external {
        bucket.refund(_amount);
    }

    function setConfig(RateLimiter.Config memory _config) external {
        bucket.setConfig(_config);
    }

    function currentState() external view returns (RateLimiter.TokenBucket memory) {
        return RateLimiter.currentState(bucket);
    }

    function storedState() external view returns (RateLimiter.TokenBucket memory) {
        return bucket;
    }
}

contract RateLimiterTest is Test {
    uint128 internal constant CAPACITY = 300 ether;
    uint128 internal constant RATE = CAPACITY / 1 days;

    RateLimiterHarness internal limiter;

    function setUp() public {
        vm.warp(1_700_000_000);
        limiter = new RateLimiterHarness();
    }

    function test_consume_disabled_succeeds(uint256 _amount) external {
        limiter.consume(_amount);
        limiter.refund(_amount);

        _assertState(limiter.storedState(), 0, 0, false, 0, 0);
    }

    function test_consume_zeroAmount_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);

        // a zero amount neither refills nor touches the bucket
        vm.warp(block.timestamp + 100);
        limiter.consume(0);
        limiter.refund(0);
        _assertState(limiter.storedState(), 0, block.timestamp - 100, true, CAPACITY, RATE);
    }

    function test_consume_exceedCapacity_reverts() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));

        vm.expectRevert(
            abi.encodeWithSelector(RateLimiter.ErrRateLimitCapacityExceeded.selector, CAPACITY, uint256(CAPACITY) + 1)
        );
        limiter.consume(uint256(CAPACITY) + 1);
    }

    function test_consume_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));

        limiter.consume(100 ether);
        _assertState(limiter.storedState(), CAPACITY - 100 ether, block.timestamp, true, CAPACITY, RATE);

        // consume exactly the available tokens
        limiter.consume(CAPACITY - 100 ether);
        _assertState(limiter.storedState(), 0, block.timestamp, true, CAPACITY, RATE);

        // no refill within the same block, a single wei short needs to wait 1 second
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, 1, 0));
        limiter.consume(1);
    }

    function test_consume_rateLimitReached_reverts() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);

        // the wait time is rounded up
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, 11, 0));
        limiter.consume(uint256(RATE) * 10 + 1);

        vm.warp(block.timestamp + 10);
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, 1, uint256(RATE) * 10));
        limiter.consume(uint256(RATE) * 10 + 1);

        limiter.consume(uint256(RATE) * 10);
        _assertState(limiter.storedState(), 0, block.timestamp, true, CAPACITY, RATE);
    }

    function test_consume_refillCappedAtCapacity_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);

        vm.warp(block.timestamp + 1 days);
        assertEq(limiter.currentState().tokens, uint256(RATE) * 1 days);

        vm.warp(block.timestamp + 1 days);
        assertEq(limiter.currentState().tokens, CAPACITY);

        limiter.consume(CAPACITY);
        _assertState(limiter.storedState(), 0, block.timestamp, true, CAPACITY, RATE);
    }

    function test_consume_zeroRate_reverts() external {
        // a one-off allowance without refill
        limiter.setConfig(_config(true, CAPACITY, 0));
        limiter.consume(CAPACITY);

        vm.warp(block.timestamp + 365 days);
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, type(uint256).max, 0));
        limiter.consume(1);
    }

    function test_consume_throttled_reverts() external {
        limiter.setConfig(_config(true, 0, 0));

        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitCapacityExceeded.selector, 0, 1));
        limiter.consume(1);

        limiter.consume(0);
    }

    function test_refund_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));

        limiter.consume(100 ether);
        limiter.refund(50 ether);
        _assertState(limiter.storedState(), CAPACITY - 50 ether, block.timestamp, true, CAPACITY, RATE);

        // capped at the capacity
        limiter.refund(100 ether);
        _assertState(limiter.storedState(), CAPACITY, block.timestamp, true, CAPACITY, RATE);

        // refilled before the refund
        limiter.consume(CAPACITY);
        vm.warp(block.timestamp + 10);
        limiter.refund(1 ether);
        _assertState(limiter.storedState(), uint256(RATE) * 10 + 1 ether, block.timestamp, true, CAPACITY, RATE);
    }

    function test_setConfig_invalidConfig_reverts() external {
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrInvalidRateLimitConfig.selector, true, 1, 2));
        limiter.setConfig(_config(true, 1, 2));

        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrInvalidRateLimitConfig.selector, false, 1, 0));
        limiter.setConfig(_config(false, 1, 0));

        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrInvalidRateLimitConfig.selector, false, 0, 1));
        limiter.setConfig(_config(false, 0, 1));

        // the rate can be equal to the capacity, and a bucket can be throttled to zero
        limiter.setConfig(_config(true, CAPACITY, CAPACITY));
        limiter.setConfig(_config(true, 0, 0));
        _assertState(limiter.storedState(), 0, block.timestamp, true, 0, 0);
    }

    function test_setConfig_enable_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        _assertState(limiter.storedState(), CAPACITY, block.timestamp, true, CAPACITY, RATE);
    }

    function test_setConfig_update_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);
        vm.warp(block.timestamp + 100);

        // refilled with the old rate, a larger capacity does not grant extra tokens
        limiter.setConfig(_config(true, CAPACITY * 2, RATE * 2));
        _assertState(limiter.storedState(), uint256(RATE) * 100, block.timestamp, true, CAPACITY * 2, RATE * 2);

        // a smaller capacity caps the tokens
        limiter.setConfig(_config(true, RATE * 50, RATE));
        _assertState(limiter.storedState(), uint256(RATE) * 50, block.timestamp, true, RATE * 50, RATE);
    }

    function test_setConfig_disable_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);

        limiter.setConfig(_config(false, 0, 0));
        _assertState(limiter.storedState(), 0, block.timestamp, false, 0, 0);

        // enabled again with a full bucket
        vm.warp(block.timestamp + 1);
        limiter.setConfig(_config(true, CAPACITY, RATE));
        _assertState(limiter.storedState(), CAPACITY, block.timestamp, true, CAPACITY, RATE);
    }

    function test_currentState_succeeds() external {
        limiter.setConfig(_config(true, CAPACITY, RATE));
        limiter.consume(CAPACITY);
        uint256 _lastUpdated = block.timestamp;

        vm.warp(block.timestamp + 100);
        _assertState(limiter.currentState(), uint256(RATE) * 100, block.timestamp, true, CAPACITY, RATE);
        _assertState(limiter.storedState(), 0, _lastUpdated, true, CAPACITY, RATE);

        // matches what the next consume sees
        vm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, 1, uint256(RATE) * 100));
        limiter.consume(uint256(RATE) * 100 + 1);
        limiter.consume(uint256(RATE) * 100);
    }

    function test_consume_outflowBounded_succeeds(uint256 _seed, uint128 _capacity, uint128 _rate) external {
        _capacity = uint128(bound(_capacity, 1, type(uint128).max));
        _rate = uint128(bound(_rate, 0, _capacity));
        limiter.setConfig(_config(true, _capacity, _rate));

        uint256 _start = block.timestamp;
        uint256 _outflow;
        for (uint256 i = 0; i < 32; i++) {
            uint256 _random = uint256(keccak256(abi.encode(_seed, i)));
            vm.warp(block.timestamp + (_random % 1 hours));
            uint256 _amount = (_random >> 64) % (uint256(_capacity) + 1);
            try limiter.consume(_amount) {
                _outflow += _amount;
                if ((_random >> 192) % 4 == 0) {
                    limiter.refund(_amount);
                    _outflow -= _amount;
                }
            } catch {}

            assertLe(limiter.storedState().tokens, _capacity);
            assertLe(_outflow, uint256(_capacity) + uint256(_rate) * (block.timestamp - _start));
        }
    }

    function test_consumeAndRefund_sameBlock_succeeds(
        uint128 _capacity,
        uint128 _rate,
        uint256 _consumed,
        uint256 _elapsed,
        uint256 _amount
    ) external {
        _capacity = uint128(bound(_capacity, 1, type(uint128).max));
        _rate = uint128(bound(_rate, 0, _capacity));
        limiter.setConfig(_config(true, _capacity, _rate));
        limiter.consume(bound(_consumed, 0, _capacity));
        // `lastUpdated` is a uint32 timestamp
        vm.warp(block.timestamp + bound(_elapsed, 0, 3650 days));

        RateLimiter.TokenBucket memory _before = limiter.currentState();
        _amount = bound(_amount, 0, _before.tokens);
        limiter.consume(_amount);
        limiter.refund(_amount);

        RateLimiter.TokenBucket memory _after = limiter.currentState();
        assertEq(_after.tokens, _before.tokens);
        assertEq(_after.lastUpdated, _before.lastUpdated);
    }

    function _config(
        bool _isEnabled,
        uint128 _capacity,
        uint128 _rate
    ) internal pure returns (RateLimiter.Config memory) {
        return RateLimiter.Config({isEnabled: _isEnabled, capacity: _capacity, rate: _rate});
    }

    function _assertState(
        RateLimiter.TokenBucket memory _bucket,
        uint256 _tokens,
        uint256 _lastUpdated,
        bool _isEnabled,
        uint256 _capacity,
        uint256 _rate
    ) internal {
        assertEq(_bucket.tokens, _tokens);
        assertEq(_bucket.lastUpdated, _lastUpdated);
        assertEq(_bucket.isEnabled, _isEnabled);
        assertEq(_bucket.capacity, _capacity);
        assertEq(_bucket.rate, _rate);
    }
}
