package genesis

import (
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/predeploys"
	"morph-l2/morph-deployer/morph-chain-ops/state"
)

func TestSetDevnetTestTokensStorageLayout(t *testing.T) {
	db := state.NewMemoryStateDB(nil)
	contractAddr := predeploys.L2TokenRegistryAddr
	db.CreateAccount(contractAddr)

	require.NoError(t, SetDevnetTestTokens(db))

	tokens := GetDevnetTestTokens()
	registrySlot := big.NewInt(151)
	registrationSlot := big.NewInt(152)
	supportedSetSlot := big.NewInt(156)

	require.Equal(t, common.BigToHash(big.NewInt(int64(len(tokens)))), db.GetState(contractAddr, common.BigToHash(supportedSetSlot)))

	valuesBaseSlot := crypto.Keccak256Hash(common.BigToHash(supportedSetSlot).Bytes())
	for i, token := range tokens {
		baseSlot := mappingSlot(new(big.Int).SetUint64(uint64(token.TokenID)), registrySlot)

		// These assertions mirror getTokenInfo(), including balance-slot decoding.
		require.Equal(t, common.BytesToHash(token.TokenAddress.Bytes()), db.GetState(contractAddr, baseSlot))
		storedBalanceSlot := db.GetState(contractAddr, offsetSlot(baseSlot, 1))
		hasBalanceSlot := storedBalanceSlot != (common.Hash{})
		require.Equal(t, token.NeedBalanceSlot, hasBalanceSlot)
		if hasBalanceSlot {
			actualBalanceSlot := common.BigToHash(new(big.Int).Sub(storedBalanceSlot.Big(), common.Big1))
			require.Equal(t, token.BalanceSlot, actualBalanceSlot)
		}

		statusAndDecimals := db.GetState(contractAddr, offsetSlot(baseSlot, 2)).Big().Uint64()
		require.Zero(t, statusAndDecimals&0xff, "tokens must start inactive")
		require.Equal(t, uint64(token.Decimals), statusAndDecimals>>8)
		require.Equal(t, common.BigToHash(token.Scale), db.GetState(contractAddr, offsetSlot(baseSlot, 3)))

		// This mirrors getTokenIdByAddress().
		reverseSlot := mappingSlot(new(big.Int).SetBytes(token.TokenAddress.Bytes()), registrationSlot)
		require.Equal(t, common.BigToHash(new(big.Int).SetUint64(uint64(token.TokenID))), db.GetState(contractAddr, reverseSlot))

		// These assertions mirror getSupportedIDList() and getSupportedTokenList().
		valueSlot := offsetSlot(valuesBaseSlot, int64(i))
		require.Equal(t, common.BigToHash(new(big.Int).SetUint64(uint64(token.TokenID))), db.GetState(contractAddr, valueSlot))
		indexSlot := mappingSlot(new(big.Int).SetUint64(uint64(token.TokenID)), new(big.Int).Add(supportedSetSlot, common.Big1))
		require.Equal(t, common.BigToHash(new(big.Int).SetInt64(int64(i+1))), db.GetState(contractAddr, indexSlot))
	}
}

// TestDevnetTestTokenDefinitions pins the two properties that cannot be derived from
// the encoding itself: WETH's balance slot is a real slot 0 rather than "no slot", and
// the placeholder addresses stay out of the precompile range.
func TestDevnetTestTokenDefinitions(t *testing.T) {
	byID := make(map[uint16]DevnetTestToken)
	for _, token := range GetDevnetTestTokens() {
		byID[token.TokenID] = token
	}

	weth := byID[2]
	require.Equal(t, predeploys.L2WETHAddr, weth.TokenAddress)
	require.True(t, weth.NeedBalanceSlot, "WrappedEther keeps _balances at slot 0, which still needs to be registered")
	require.Equal(t, common.Hash{}, weth.BalanceSlot)

	lowestNonPrecompile := big.NewInt(0xff)
	for _, token := range byID {
		require.Positive(t, new(big.Int).SetBytes(token.TokenAddress.Bytes()).Cmp(lowestNonPrecompile),
			"token %d address %s falls in the precompile range", token.TokenID, token.TokenAddress)

		// A scale of 10^(18-decimals) reads plausibly but double-applies the decimals
		// adjustment the oracle already makes, which collapses to 1 for an 18-decimal
		// token and truncates every priceRatio below ETH to zero.
		want := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token.Decimals)), nil)
		require.Zero(t, want.Cmp(token.Scale),
			"token %d scale must be 10^decimals, got %s", token.TokenID, token.Scale)
	}
}

// TestDevnetTestTokenPriceRatioKeepsPrecision walks the oracle's priceRatio formula for
// the cheapest pre-registered token. The registry stores priceRatio as a uint256, so a
// scale that leaves the ratio below 1 makes the token permanently unpriceable.
func TestDevnetTestTokenPriceRatioKeepsPrecision(t *testing.T) {
	// BGB near its spot price against ETH, the widest token/ETH gap in the set.
	tokenPriceUSD := big.NewFloat(1.5954)
	ethPriceUSD := big.NewFloat(1845.55)

	for _, token := range GetDevnetTestTokens() {
		ratio := new(big.Float).SetInt(token.Scale)
		ratio.Mul(ratio, tokenPriceUSD)
		ratio.Mul(ratio, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18-token.Decimals)), nil)))
		ratio.Quo(ratio, ethPriceUSD)

		truncated, _ := ratio.Int(nil)
		require.Positive(t, truncated.Sign(),
			"token %d priceRatio truncates to zero with scale %s", token.TokenID, token.Scale)
	}
}

func mappingSlot(key, slot *big.Int) common.Hash {
	keyBytes := common.LeftPadBytes(key.Bytes(), 32)
	slotBytes := common.LeftPadBytes(slot.Bytes(), 32)
	return crypto.Keccak256Hash(append(keyBytes, slotBytes...))
}

func offsetSlot(base common.Hash, offset int64) common.Hash {
	return common.BigToHash(new(big.Int).Add(base.Big(), big.NewInt(offset)))
}
