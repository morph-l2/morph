package client

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
)

const chainlinkAggregatorV3ABI = `[
	{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},
	{"inputs":[],"name":"latestRoundData","outputs":[{"internalType":"uint80","name":"roundId","type":"uint80"},{"internalType":"int256","name":"answer","type":"int256"},{"internalType":"uint256","name":"startedAt","type":"uint256"},{"internalType":"uint256","name":"updatedAt","type":"uint256"},{"internalType":"uint80","name":"answeredInRound","type":"uint80"}],"stateMutability":"view","type":"function"}
]`

var parsedChainlinkAggregatorABI = mustParseChainlinkAggregatorABI()

// ChainlinkPriceFeed reads Chainlink AggregatorV3 feeds over RPC.
type ChainlinkPriceFeed struct {
	caller       bind.ContractCaller
	mu           sync.RWMutex
	tokenFeeds   map[uint16]common.Address
	ethUSDFeed   common.Address
	maxStaleness time.Duration
	log          log.Logger
}

// NewChainlinkPriceFeed creates a Chainlink price feed using an RPC endpoint.
func NewChainlinkPriceFeed(tokenFeedMap map[uint16]string, rpcURL string, ethUSDFeed common.Address, maxStaleness time.Duration) (*ChainlinkPriceFeed, error) {
	if rpcURL == "" {
		return nil, fmt.Errorf("chainlink price feed requires --chainlink-rpc")
	}

	caller, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect chainlink rpc: %w", err)
	}

	feed, err := NewChainlinkPriceFeedWithCaller(tokenFeedMap, caller, ethUSDFeed, maxStaleness)
	if err != nil {
		caller.Close()
		return nil, err
	}
	return feed, nil
}

// NewChainlinkPriceFeedWithCaller creates a Chainlink price feed with a caller.
// It is primarily useful for tests.
func NewChainlinkPriceFeedWithCaller(tokenFeedMap map[uint16]string, caller bind.ContractCaller, ethUSDFeed common.Address, maxStaleness time.Duration) (*ChainlinkPriceFeed, error) {
	if caller == nil {
		return nil, fmt.Errorf("chainlink price feed requires rpc caller")
	}
	if ethUSDFeed == (common.Address{}) {
		return nil, fmt.Errorf("chainlink price feed requires --chainlink-eth-usd-feed")
	}
	if maxStaleness <= 0 {
		return nil, fmt.Errorf("chainlink max staleness must be positive")
	}

	feeds := make(map[uint16]common.Address, len(tokenFeedMap))
	for tokenID, feedAddr := range tokenFeedMap {
		feedAddr = strings.TrimSpace(feedAddr)
		if !common.IsHexAddress(feedAddr) {
			return nil, fmt.Errorf("invalid chainlink feed address for token %d: %s", tokenID, feedAddr)
		}
		feeds[tokenID] = common.HexToAddress(feedAddr)
	}
	if len(feeds) == 0 {
		return nil, fmt.Errorf("chainlink price feed requires token mapping, please configure --token-mapping-chainlink")
	}

	return &ChainlinkPriceFeed{
		caller:       caller,
		tokenFeeds:   feeds,
		ethUSDFeed:   ethUSDFeed,
		maxStaleness: maxStaleness,
		log:          log.New("component", "chainlink_price_feed"),
	}, nil
}

// GetTokenPrice returns token price in USD from Chainlink.
func (c *ChainlinkPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	c.mu.RLock()
	feedAddress, exists := c.tokenFeeds[tokenID]
	ethUSDFeed := c.ethUSDFeed
	c.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("token ID %d not mapped to Chainlink feed", tokenID)
	}

	ethPrice, err := c.fetchFeedPrice(ctx, ethUSDFeed)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ETH/USD price from Chainlink: %w", err)
	}

	tokenPrice, err := c.fetchFeedPrice(ctx, feedAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token price from Chainlink for token %d: %w", tokenID, err)
	}

	c.log.Info("Fetched price from Chainlink",
		"source", "chainlink",
		"token_id", tokenID,
		"feed", feedAddress.Hex(),
		"token_price_usd", tokenPrice.String(),
		"eth_price_usd", ethPrice.String())

	return &TokenPrice{
		TokenID:       tokenID,
		Symbol:        feedAddress.Hex(),
		TokenPriceUSD: tokenPrice,
		EthPriceUSD:   ethPrice,
	}, nil
}

// GetBatchTokenPrices returns token prices in USD for multiple tokens.
func (c *ChainlinkPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	ethPrice, err := c.fetchFeedPrice(ctx, c.ethUSDFeed)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ETH/USD price from Chainlink: %w", err)
	}

	prices := make(map[uint16]*TokenPrice, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		c.mu.RLock()
		feedAddress, exists := c.tokenFeeds[tokenID]
		c.mu.RUnlock()
		if !exists {
			return nil, fmt.Errorf("token ID %d not mapped to Chainlink feed", tokenID)
		}

		tokenPrice, err := c.fetchFeedPrice(ctx, feedAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch token price from Chainlink for token %d: %w", tokenID, err)
		}

		prices[tokenID] = &TokenPrice{
			TokenID:       tokenID,
			Symbol:        feedAddress.Hex(),
			TokenPriceUSD: tokenPrice,
			EthPriceUSD:   new(big.Float).Copy(ethPrice),
		}
	}

	return prices, nil
}

func (c *ChainlinkPriceFeed) fetchFeedPrice(ctx context.Context, feedAddress common.Address) (*big.Float, error) {
	contract := bind.NewBoundContract(feedAddress, parsedChainlinkAggregatorABI, c.caller, nil, nil)

	var roundData []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &roundData, "latestRoundData"); err != nil {
		return nil, fmt.Errorf("latestRoundData call failed for feed %s: %w", feedAddress.Hex(), err)
	}

	roundID, answer, updatedAt, answeredInRound, err := parseChainlinkRoundData(roundData)
	if err != nil {
		return nil, fmt.Errorf("invalid latestRoundData response for feed %s: %w", feedAddress.Hex(), err)
	}
	if err := validateChainlinkRound(answer, updatedAt, roundID, answeredInRound, c.maxStaleness, time.Now()); err != nil {
		return nil, fmt.Errorf("invalid Chainlink round for feed %s: %w", feedAddress.Hex(), err)
	}

	var decimalsOut []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &decimalsOut, "decimals"); err != nil {
		return nil, fmt.Errorf("decimals call failed for feed %s: %w", feedAddress.Hex(), err)
	}
	decimals, err := parseChainlinkDecimals(decimalsOut)
	if err != nil {
		return nil, fmt.Errorf("invalid decimals response for feed %s: %w", feedAddress.Hex(), err)
	}

	return chainlinkAnswerToFloat(answer, decimals), nil
}

func parseChainlinkRoundData(values []interface{}) (roundID, answer, updatedAt, answeredInRound *big.Int, err error) {
	if len(values) != 5 {
		return nil, nil, nil, nil, fmt.Errorf("expected 5 values, got %d", len(values))
	}

	roundID, ok := values[0].(*big.Int)
	if !ok {
		return nil, nil, nil, nil, errors.New("roundId is not *big.Int")
	}
	answer, ok = values[1].(*big.Int)
	if !ok {
		return nil, nil, nil, nil, errors.New("answer is not *big.Int")
	}
	updatedAt, ok = values[3].(*big.Int)
	if !ok {
		return nil, nil, nil, nil, errors.New("updatedAt is not *big.Int")
	}
	answeredInRound, ok = values[4].(*big.Int)
	if !ok {
		return nil, nil, nil, nil, errors.New("answeredInRound is not *big.Int")
	}

	return roundID, answer, updatedAt, answeredInRound, nil
}

func parseChainlinkDecimals(values []interface{}) (uint8, error) {
	if len(values) != 1 {
		return 0, fmt.Errorf("expected 1 value, got %d", len(values))
	}

	switch decimals := values[0].(type) {
	case uint8:
		return decimals, nil
	case *big.Int:
		if !decimals.IsUint64() || decimals.Uint64() > 255 {
			return 0, fmt.Errorf("decimals out of uint8 range: %s", decimals.String())
		}
		return uint8(decimals.Uint64()), nil
	default:
		return 0, fmt.Errorf("decimals has unexpected type %T", values[0])
	}
}

func validateChainlinkRound(answer, updatedAt, roundID, answeredInRound *big.Int, maxStaleness time.Duration, now time.Time) error {
	if answer == nil || updatedAt == nil || roundID == nil || answeredInRound == nil {
		return errors.New("round data contains nil value")
	}
	if answer.Sign() <= 0 {
		return fmt.Errorf("answer must be positive, got %s", answer.String())
	}
	if updatedAt.Sign() <= 0 {
		return errors.New("updatedAt must be positive")
	}
	if answeredInRound.Cmp(roundID) < 0 {
		return fmt.Errorf("answeredInRound %s is older than roundId %s", answeredInRound.String(), roundID.String())
	}

	updated := time.Unix(updatedAt.Int64(), 0)
	if updated.After(now.Add(maxStaleness)) {
		return fmt.Errorf("updatedAt %s is too far in the future", updated.UTC().Format(time.RFC3339))
	}
	if now.Sub(updated) > maxStaleness {
		return fmt.Errorf("price is stale: updatedAt=%s maxStaleness=%s", updated.UTC().Format(time.RFC3339), maxStaleness)
	}

	return nil
}

func chainlinkAnswerToFloat(answer *big.Int, decimals uint8) *big.Float {
	price := new(big.Float).SetPrec(256).SetInt(answer)
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return price.Quo(price, new(big.Float).SetPrec(256).SetInt(scale))
}

func mustParseChainlinkAggregatorABI() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(chainlinkAggregatorV3ABI))
	if err != nil {
		panic(err)
	}
	return parsed
}
