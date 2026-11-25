// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/// @title RateLimiter
/// @notice Token bucket rate limiter. The bucket holds at most `capacity` tokens and is refilled at `rate`
///         tokens per second, so the amount consumed during any period `T` never exceeds `capacity + rate * T`.
library RateLimiter {
    /**********
     * Errors *
     **********/

    /// @dev Thrown when the requested amount exceeds the bucket capacity, which can never succeed.
    /// @param capacity The capacity of the bucket.
    /// @param requested The requested amount.
    error ErrRateLimitCapacityExceeded(uint256 capacity, uint256 requested);

    /// @dev Thrown when the bucket does not hold enough tokens for now.
    /// @param minWaitInSeconds The minimum time to wait before the bucket can cover the request.
    /// @param available The amount of tokens currently available.
    error ErrRateLimitReached(uint256 minWaitInSeconds, uint256 available);

    /// @dev Thrown when the config is invalid.
    /// @param isEnabled Whether the rate limit is enabled.
    /// @param capacity The capacity of the bucket.
    /// @param rate The refill rate of the bucket.
    error ErrInvalidRateLimitConfig(bool isEnabled, uint128 capacity, uint128 rate);

    /***********
     * Structs *
     ***********/

    struct Config {
        // Whether the rate limit is enabled.
        bool isEnabled;
        // The maximum amount of tokens the bucket can hold.
        uint128 capacity;
        // The amount of tokens refilled per second.
        uint128 rate;
    }

    struct TokenBucket {
        // The amount of tokens in the bucket at `lastUpdated`.
        uint128 tokens;
        // The timestamp when the bucket was last updated.
        uint32 lastUpdated;
        // Whether the rate limit is enabled.
        bool isEnabled;
        // The maximum amount of tokens the bucket can hold.
        uint128 capacity;
        // The amount of tokens refilled per second.
        uint128 rate;
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev Consume tokens from the bucket, revert if the bucket cannot cover the amount.
    /// @param _bucket The token bucket.
    /// @param _amount The amount of tokens to consume.
    function consume(TokenBucket storage _bucket, uint256 _amount) internal {
        if (!_bucket.isEnabled || _amount == 0) return;

        uint256 _capacity = _bucket.capacity;
        if (_amount > _capacity) revert ErrRateLimitCapacityExceeded(_capacity, _amount);

        uint256 _rate = _bucket.rate;
        uint256 _tokens = _refill(_bucket.tokens, _capacity, _rate, _bucket.lastUpdated);
        if (_tokens < _amount) {
            uint256 _wait = _rate == 0 ? type(uint256).max : (_amount - _tokens + _rate - 1) / _rate;
            revert ErrRateLimitReached(_wait, _tokens);
        }

        _bucket.tokens = uint128(_tokens - _amount);
        _bucket.lastUpdated = uint32(block.timestamp);
    }

    /// @dev Give back tokens to the bucket when the guarded transfer did not happen.
    /// @param _bucket The token bucket.
    /// @param _amount The amount of tokens to give back.
    function refund(TokenBucket storage _bucket, uint256 _amount) internal {
        if (!_bucket.isEnabled || _amount == 0) return;

        uint256 _capacity = _bucket.capacity;
        uint256 _tokens = _refill(_bucket.tokens, _capacity, _bucket.rate, _bucket.lastUpdated) + _amount;
        _bucket.tokens = uint128(_tokens < _capacity ? _tokens : _capacity);
        _bucket.lastUpdated = uint32(block.timestamp);
    }

    /// @dev Update the config of the bucket. An enabled bucket is refilled with the old config and then
    ///      capped at the new capacity, so updating the config never grants extra tokens. A disabled bucket
    ///      starts full once enabled, and a disabled config must be all zero.
    /// @param _bucket The token bucket.
    /// @param _config The new config.
    function setConfig(TokenBucket storage _bucket, Config memory _config) internal {
        if (_config.isEnabled ? _config.rate > _config.capacity : (_config.capacity != 0 || _config.rate != 0)) {
            revert ErrInvalidRateLimitConfig(_config.isEnabled, _config.capacity, _config.rate);
        }

        uint256 _tokens = _bucket.isEnabled
            ? _refill(_bucket.tokens, _bucket.capacity, _bucket.rate, _bucket.lastUpdated)
            : _config.capacity;
        _bucket.tokens = uint128(_tokens < _config.capacity ? _tokens : _config.capacity);
        _bucket.lastUpdated = uint32(block.timestamp);
        _bucket.isEnabled = _config.isEnabled;
        _bucket.capacity = _config.capacity;
        _bucket.rate = _config.rate;
    }

    /// @dev Return the bucket refilled to the current block.
    /// @param _bucket The token bucket.
    function currentState(TokenBucket memory _bucket) internal view returns (TokenBucket memory) {
        _bucket.tokens = uint128(_refill(_bucket.tokens, _bucket.capacity, _bucket.rate, _bucket.lastUpdated));
        _bucket.lastUpdated = uint32(block.timestamp);
        return _bucket;
    }

    /*********************
     * Private Functions *
     *********************/

    /// @dev `_rate * elapsed` fits in uint256 since `_rate` is a uint128 and the timestamps are uint32.
    function _refill(
        uint256 _tokens,
        uint256 _capacity,
        uint256 _rate,
        uint256 _lastUpdated
    ) private view returns (uint256) {
        uint256 _refilled = _tokens + (block.timestamp - _lastUpdated) * _rate;
        return _refilled < _capacity ? _refilled : _capacity;
    }
}
