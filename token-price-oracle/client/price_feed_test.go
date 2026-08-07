package client

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"testing"
)

// stubFeed serves only the tokens present in prices and records every batch it was
// asked for, so tests can assert which tokens reached each feed.
type stubFeed struct {
	prices    map[uint16]float64
	err       error
	requested [][]uint16
}

func (s *stubFeed) GetTokenPrice(_ context.Context, tokenID uint16) (*TokenPrice, error) {
	prices, err := s.GetBatchTokenPrices(context.Background(), []uint16{tokenID})
	if err != nil {
		return nil, err
	}
	price, exists := prices[tokenID]
	if !exists {
		return nil, errors.New("token not served by this feed")
	}
	return price, nil
}

func (s *stubFeed) GetBatchTokenPrices(_ context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	s.requested = append(s.requested, append([]uint16(nil), tokenIDs...))
	if s.err != nil {
		return nil, s.err
	}

	out := make(map[uint16]*TokenPrice, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		price, exists := s.prices[tokenID]
		if !exists {
			continue
		}
		out[tokenID] = &TokenPrice{
			TokenID:       tokenID,
			TokenPriceUSD: big.NewFloat(price),
			EthPriceUSD:   big.NewFloat(3000),
		}
	}
	return out, nil
}

// TestFallbackMergesPartialCoverageAcrossFeeds covers the configuration this repo
// documents for devnet: Chainlink and Pyth map tokens 1 and 2 while the CEX feeds also
// cover token 3. Discarding a response for missing one token meant the oracle feeds
// failed every cycle and every token silently came from the CEX.
func TestFallbackMergesPartialCoverageAcrossFeeds(t *testing.T) {
	oracle := &stubFeed{prices: map[uint16]float64{1: 100, 2: 200}}
	cex := &stubFeed{prices: map[uint16]float64{1: 111, 2: 222, 3: 333}}

	feed := NewFallbackPriceFeed([]PriceFeed{oracle, cex}, []string{"oracle", "cex"})

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if len(prices) != 3 {
		t.Fatalf("resolved %d tokens, want 3", len(prices))
	}

	// Tokens 1 and 2 must come from the higher-priority feed, token 3 from the CEX.
	for tokenID, want := range map[uint16]float64{1: 100, 2: 200, 3: 333} {
		if got := prices[tokenID].TokenPriceUSD; got.Cmp(big.NewFloat(want)) != 0 {
			t.Errorf("token %d price = %s, want %v", tokenID, got.String(), want)
		}
	}

	// The second feed must only be asked for what the first could not resolve.
	if want := [][]uint16{{1, 2, 3}}; !reflect.DeepEqual(oracle.requested, want) {
		t.Errorf("oracle feed received %v, want %v", oracle.requested, want)
	}
	if want := [][]uint16{{3}}; !reflect.DeepEqual(cex.requested, want) {
		t.Errorf("cex feed received %v, want %v", cex.requested, want)
	}
}

func TestFallbackSkipsFeedsThatAlreadyResolvedEverything(t *testing.T) {
	first := &stubFeed{prices: map[uint16]float64{1: 100, 2: 200}}
	second := &stubFeed{prices: map[uint16]float64{1: 111, 2: 222}}

	feed := NewFallbackPriceFeed([]PriceFeed{first, second}, []string{"first", "second"})

	if _, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2}); err != nil {
		t.Fatal(err)
	}
	if len(second.requested) != 0 {
		t.Fatalf("second feed was queried %v, want no queries", second.requested)
	}
}

// TestFallbackReturnsPartialResultForUnservableToken pins the agreed semantics: a token
// no feed can price does not block the others, and the caller sees it by its absence.
func TestFallbackReturnsPartialResultForUnservableToken(t *testing.T) {
	first := &stubFeed{prices: map[uint16]float64{1: 100}}
	second := &stubFeed{prices: map[uint16]float64{1: 111}}

	feed := NewFallbackPriceFeed([]PriceFeed{first, second}, []string{"first", "second"})

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(prices) != 1 {
		t.Fatalf("resolved %d tokens, want 1", len(prices))
	}
	if _, exists := prices[2]; exists {
		t.Fatal("token 2 is served by no feed and must be absent from the result")
	}
}

// TestFallbackErrorsWhenNothingResolves is the single-feed regression: a CEX feed that
// returns an empty map with a nil error used to reach the updater as a successful
// cycle, which then advanced the last-successful-update timestamp.
func TestFallbackErrorsWhenNothingResolves(t *testing.T) {
	only := &stubFeed{prices: map[uint16]float64{}}

	feed := NewFallbackPriceFeed([]PriceFeed{only}, []string{"only"})

	if _, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2}); err == nil {
		t.Fatal("GetBatchTokenPrices succeeded on an empty result, want error")
	}
}

func TestFallbackContinuesPastFailingFeed(t *testing.T) {
	broken := &stubFeed{err: errors.New("rpc down")}
	healthy := &stubFeed{prices: map[uint16]float64{1: 100, 2: 200}}

	feed := NewFallbackPriceFeed([]PriceFeed{broken, healthy}, []string{"broken", "healthy"})

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(prices) != 2 {
		t.Fatalf("resolved %d tokens, want 2", len(prices))
	}
	// A failing feed must not shrink what the next feed is asked for.
	if want := [][]uint16{{1, 2}}; !reflect.DeepEqual(healthy.requested, want) {
		t.Errorf("healthy feed received %v, want %v", healthy.requested, want)
	}
}
