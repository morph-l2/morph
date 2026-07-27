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
		require.Equal(t, token.BalanceSlot != (common.Hash{}), hasBalanceSlot)
		actualBalanceSlot := storedBalanceSlot
		if hasBalanceSlot {
			actualBalanceSlot = common.BigToHash(new(big.Int).Sub(storedBalanceSlot.Big(), common.Big1))
		}
		require.Equal(t, token.BalanceSlot, actualBalanceSlot)

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

func mappingSlot(key, slot *big.Int) common.Hash {
	keyBytes := common.LeftPadBytes(key.Bytes(), 32)
	slotBytes := common.LeftPadBytes(slot.Bytes(), 32)
	return crypto.Keccak256Hash(append(keyBytes, slotBytes...))
}

func offsetSlot(base common.Hash, offset int64) common.Hash {
	return common.BigToHash(new(big.Int).Add(base.Big(), big.NewInt(offset)))
}
