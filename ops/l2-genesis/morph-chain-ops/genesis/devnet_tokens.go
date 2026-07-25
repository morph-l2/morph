package genesis

import (
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/vm"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/bindings/predeploys"
)

// DevnetTestToken defines a test token to be pre-registered in TokenRegistry for devnet
type DevnetTestToken struct {
	TokenID      uint16
	TokenAddress common.Address
	BalanceSlot  common.Hash
	Decimals     uint8
	Scale        *big.Int
}

// GetDevnetTestTokens returns the list of test tokens to pre-register in devnet
// Token 1: BTC - for testing high-value asset price queries (all data sources support)
// Token 2: ETH - for testing gas token benchmark and relative price calculation
// Token 3: BGB - for testing platform token and CEX-specific data sources
func GetDevnetTestTokens() []DevnetTestToken {
	return []DevnetTestToken{
		{
			TokenID:      1,
			TokenAddress: common.HexToAddress("0x0000000000000000000000000000000000000001"), // Mock BTC address
			BalanceSlot:  common.Hash{},                                                        // zero hash (no balance slot needed for mock)
			Decimals:     8,
			Scale:        big.NewInt(1e10), // 10^(18-8) for ETH decimals adjustment
		},
		{
			TokenID:      2,
			TokenAddress: common.HexToAddress("0x5300000000000000000000000000000000000011"), // L2WETH predeploy address
			BalanceSlot:  common.BigToHash(big.NewInt(3)),                                     // WETH standard balance slot
			Decimals:     18,
			Scale:        big.NewInt(1), // 1:1 ratio
		},
		{
			TokenID:      3,
			TokenAddress: common.HexToAddress("0x0000000000000000000000000000000000000003"), // Mock BGB address
			BalanceSlot:  common.Hash{},                                                        // zero hash
			Decimals:     18,
			Scale:        big.NewInt(1),
		},
	}
}

// SetDevnetTestTokens pre-registers test tokens in TokenRegistry storage for devnet
// This allows token-price-oracle to work out-of-the-box without manual token registration
func SetDevnetTestTokens(db vm.StateDB) error {
	contractAddr := predeploys.L2TokenRegistryAddr
	tokens := GetDevnetTestTokens()

	// Storage layout reference (from L2TokenRegistry.sol):
	// slot 151: mapping(uint16 => TokenInfo) tokenRegistry
	// slot 152: mapping(address => uint16) tokenRegistration
	// slot 153: mapping(uint16 => uint256) priceRatio
	// slot 156: EnumerableSet.UintSet supportedTokenSet

	tokenRegistrySlot := big.NewInt(151)
	tokenRegistrationSlot := big.NewInt(152)
	supportedTokenSetSlot := big.NewInt(156)

	log.Info("Pre-registering devnet test tokens in TokenRegistry", "count", len(tokens))

	for _, token := range tokens {
		// Set tokenRegistry[tokenID] = TokenInfo{...}
		// Storage location: keccak256(abi.encode(tokenID, 151))
		if err := setTokenInfo(db, contractAddr, tokenRegistrySlot, token); err != nil {
			return fmt.Errorf("failed to set tokenRegistry[%d]: %w", token.TokenID, err)
		}

		// Set tokenRegistration[tokenAddress] = tokenID
		// Storage location: keccak256(abi.encode(tokenAddress, 152))
		if err := setTokenRegistration(db, contractAddr, tokenRegistrationSlot, token.TokenAddress, token.TokenID); err != nil {
			return fmt.Errorf("failed to set tokenRegistration[%s]: %w", token.TokenAddress.Hex(), err)
		}

		log.Info("Pre-registered devnet token",
			"tokenID", token.TokenID,
			"address", token.TokenAddress.Hex(),
			"decimals", token.Decimals,
			"scale", token.Scale.String())
	}

	// Set supportedTokenSet (EnumerableSet.UintSet)
	if err := setSupportedTokenSet(db, contractAddr, supportedTokenSetSlot, tokens); err != nil {
		return fmt.Errorf("failed to set supportedTokenSet: %w", err)
	}

	log.Info("✓ Devnet test tokens pre-registered successfully", "tokenIDs", []uint16{1, 2, 3})
	return nil
}

// setTokenInfo sets TokenInfo struct in tokenRegistry mapping
func setTokenInfo(db vm.StateDB, contractAddr common.Address, registrySlot *big.Int, token DevnetTestToken) error {
	// Calculate base slot: keccak256(abi.encode(tokenID, registrySlot))
	tokenIDBytes := common.LeftPadBytes(big.NewInt(int64(token.TokenID)).Bytes(), 32)
	slotBytes := common.LeftPadBytes(registrySlot.Bytes(), 32)
	baseSlot := crypto.Keccak256Hash(append(tokenIDBytes, slotBytes...))

	// TokenInfo struct layout (compact storage):
	// slot+0: tokenAddress (address, 20 bytes) + first 12 bytes of balanceSlot
	// slot+1: remaining 20 bytes of balanceSlot
	// slot+2: isActive (bool, 1 byte) + decimals (uint8, 1 byte) in lowest 2 bytes
	// slot+3: scale (uint256, 32 bytes)

	// Slot+0: pack tokenAddress (20 bytes) into lowest bytes
	slot0Value := new(big.Int).SetBytes(token.TokenAddress.Bytes())
	db.SetState(contractAddr, baseSlot, common.BigToHash(slot0Value))

	// Slot+1: balanceSlot (bytes32)
	slot1Key := common.BigToHash(new(big.Int).Add(baseSlot.Big(), big.NewInt(1)))
	db.SetState(contractAddr, slot1Key, token.BalanceSlot)

	// Slot+2: isActive=false (0x00) + decimals (1 byte)
	// Pack as: [31 zeros][decimals][isActive=0]
	slot2Key := common.BigToHash(new(big.Int).Add(baseSlot.Big(), big.NewInt(2)))
	slot2Value := new(big.Int).SetUint64(uint64(token.Decimals) << 8) // decimals in second byte
	db.SetState(contractAddr, slot2Key, common.BigToHash(slot2Value))

	// Slot+3: scale (uint256)
	slot3Key := common.BigToHash(new(big.Int).Add(baseSlot.Big(), big.NewInt(3)))
	db.SetState(contractAddr, slot3Key, common.BigToHash(token.Scale))

	return nil
}

// setTokenRegistration sets reverse mapping: tokenRegistration[address] = tokenID
func setTokenRegistration(db vm.StateDB, contractAddr common.Address, registrationSlot *big.Int, tokenAddress common.Address, tokenID uint16) error {
	// Calculate storage location: keccak256(abi.encode(tokenAddress, registrationSlot))
	addrBytes := common.LeftPadBytes(tokenAddress.Bytes(), 32)
	slotBytes := common.LeftPadBytes(registrationSlot.Bytes(), 32)
	storageKey := crypto.Keccak256Hash(append(addrBytes, slotBytes...))

	// Set tokenID as uint16 (2 bytes) in storage
	tokenIDValue := new(big.Int).SetUint64(uint64(tokenID))
	db.SetState(contractAddr, storageKey, common.BigToHash(tokenIDValue))

	return nil
}

// setSupportedTokenSet sets EnumerableSet.UintSet for supported token IDs
func setSupportedTokenSet(db vm.StateDB, contractAddr common.Address, setBaseSlot *big.Int, tokens []DevnetTestToken) error {
	// EnumerableSet.UintSet layout:
	// struct UintSet {
	//     Set _inner; // slot 156
	// }
	// struct Set {
	//     bytes32[] _values;  // slot 156+0: array length at base slot, elements at keccak256(baseSlot)
	//     mapping(bytes32 => uint256) _indexes; // slot 156+1: mapping base
	// }

	// Set _values array length (number of tokens)
	lengthSlot := common.BigToHash(setBaseSlot)
	db.SetState(contractAddr, lengthSlot, common.BigToHash(big.NewInt(int64(len(tokens)))))

	// Calculate _values array storage location: keccak256(baseSlot)
	valuesBaseSlot := crypto.Keccak256Hash(lengthSlot.Bytes())

	// Set each token ID in _values array and _indexes mapping
	for i, token := range tokens {
		// Set _values[i] = tokenID (stored as bytes32/uint256)
		elemSlot := common.BigToHash(new(big.Int).Add(valuesBaseSlot.Big(), big.NewInt(int64(i))))
		tokenIDValue := new(big.Int).SetUint64(uint64(token.TokenID))
		db.SetState(contractAddr, elemSlot, common.BigToHash(tokenIDValue))

		// Set _indexes[tokenID] = i+1 (1-based index, 0 means not in set)
		// Storage location: keccak256(abi.encode(tokenID, setBaseSlot+1))
		indexesBaseSlot := new(big.Int).Add(setBaseSlot, big.NewInt(1))
		tokenIDBytes := common.LeftPadBytes(big.NewInt(int64(token.TokenID)).Bytes(), 32)
		indexSlotBytes := common.LeftPadBytes(indexesBaseSlot.Bytes(), 32)
		indexKey := crypto.Keccak256Hash(append(tokenIDBytes, indexSlotBytes...))
		indexValue := new(big.Int).SetInt64(int64(i + 1)) // 1-based
		db.SetState(contractAddr, indexKey, common.BigToHash(indexValue))
	}

	return nil
}
