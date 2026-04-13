package oracle_testing

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
)

func TestGetStorageAtL2TokenRegistry(t *testing.T) {
	TokenRegistryAddress = common.HexToAddress("0x5300000000000000000000000000000000000021")
	// Connect to Ethereum node
	rpcURL := "http://localhost:8545"
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Read allowListEnabled (slot 105, bool)
	allowListEnabledSlot := AllowListEnabledSlot
	allowListEnabledData, err := client.StorageAt(ctx, TokenRegistryAddress, allowListEnabledSlot, nil)
	if err != nil {
		t.Fatalf("Failed to read allowListEnabled: %v", err)
	}
	allowListEnabled := allowListEnabledData[31] != 0
	t.Logf("allowListEnabled: %v", allowListEnabled)

	for tokenID := uint16(1); tokenID <= 10; tokenID++ {
		t.Logf("---------- Querying tokenID: %d ----------", tokenID)

		// Calculate baseSlot
		baseSlot := CalculateUint16MappingSlot(tokenID, TokenRegistrySlot)
		t.Logf("TokenID %d base slot: %s", tokenID, baseSlot.Hex())

		// Create storage reader
		storageReader := func(slot common.Hash) ([]byte, error) {
			return client.StorageAt(ctx, TokenRegistryAddress, slot, nil)
		}

		// Parse TokenInfo
		tokenInfo, err := ParseTokenInfoFromStorage(storageReader, baseSlot)
		if err != nil {
			t.Logf("  [WARN] Failed to parse TokenInfo: %v", err)
			continue
		}

		t.Logf("TokenInfo for tokenID %d:", tokenID)
		t.Logf("  TokenAddress: %s", tokenInfo.TokenAddress.Hex())
		t.Logf("  BalanceSlot: %s", tokenInfo.BalanceSlot.Hex())
		t.Logf("  IsActive: %v", tokenInfo.IsActive)
		t.Logf("  Decimals: %d", tokenInfo.Decimals)
		t.Logf("  Scale: %s", tokenInfo.Scale.String())

		// Query tokenRegistration mapping (address => uint16)
		if tokenInfo.TokenAddress != (common.Address{}) {
			tokenRegistrationSlot := CalculateAddressMappingSlot(tokenInfo.TokenAddress, TokenRegistrationSlot)
			tokenRegistrationData, err := client.StorageAt(ctx, TokenRegistryAddress, tokenRegistrationSlot, nil)
			if err != nil {
				t.Logf("  [WARN] Failed to read tokenRegistration: %v", err)
			} else {
				registeredTokenID := uint16(tokenRegistrationData[30])<<8 | uint16(tokenRegistrationData[31])
				t.Logf("  TokenID for address %s: %d", tokenInfo.TokenAddress.Hex(), registeredTokenID)
				if registeredTokenID != tokenID {
					t.Errorf("  [ERR] Mismatch: expected tokenID %d, got %d", tokenID, registeredTokenID)
				}
			}
		}

		// Query priceRatio mapping (uint16 => uint256)
		priceRatioSlot := CalculateUint16MappingSlot(tokenID, PriceRatioSlot)
		priceRatioData, err := client.StorageAt(ctx, TokenRegistryAddress, priceRatioSlot, nil)
		if err != nil {
			t.Logf("  [WARN] Failed to read priceRatio: %v", err)
		} else {
			priceRatio := new(big.Int).SetBytes(priceRatioData)
			t.Logf("  PriceRatio: %s", priceRatio.String())
		}
	}

	// Test allowList mapping (address => bool)
	testAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	allowListSlot := CalculateAddressMappingSlot(testAddress, AllowListSlot)
	allowListData, err := client.StorageAt(ctx, TokenRegistryAddress, allowListSlot, nil)
	if err != nil {
		t.Fatalf("Failed to read allowList: %v", err)
	}
	isAllowed := allowListData[31] != 0
	t.Logf("Address %s in allowList: %v", testAddress.Hex(), isAllowed)
}

// TestCalculateStorageSlots tests the storage slot calculation functions
func TestCalculateStorageSlots(t *testing.T) {
	tokenID := uint16(1)
	baseSlot := CalculateUint16MappingSlot(tokenID, TokenRegistrySlot)
	t.Logf("Base slot for tokenID %d: %s", tokenID, baseSlot.Hex())

	// Test struct field slots
	tokenAddrSlot := CalculateStructFieldSlot(baseSlot, 0)
	t.Logf("TokenAddress slot: %s", tokenAddrSlot.Hex())

	balanceSlotField := CalculateStructFieldSlot(baseSlot, 1)
	t.Logf("BalanceSlot field slot: %s", balanceSlotField.Hex())

	isActiveSlot := CalculateStructFieldSlot(baseSlot, 2)
	t.Logf("IsActive slot: %s", isActiveSlot.Hex())

	scaleSlot := CalculateStructFieldSlot(baseSlot, 3)
	t.Logf("Scale slot: %s", scaleSlot.Hex())

	// Test address mapping slot
	testAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	addressMappingSlot := CalculateAddressMappingSlot(testAddress, TokenRegistrationSlot)
	t.Logf("Address mapping slot for %s: %s", testAddress.Hex(), addressMappingSlot.Hex())

	// Test priceRatio slot
	priceRatioSlot := CalculateUint16MappingSlot(tokenID, PriceRatioSlot)
	t.Logf("PriceRatio slot for tokenID %d: %s", tokenID, priceRatioSlot.Hex())
}

// TestCompareTokensFromJSON compares token configurations from JSON file with contract data
func TestCompareTokensFromJSON(t *testing.T) {
	TokenRegistryAddress = common.HexToAddress("0x5300000000000000000000000000000000000021")
	// Connect to Ethereum node
	rpcURL := "http://localhost:8545"
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Load token configurations from JSON file
	// Get absolute path relative to the test file location
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("Failed to get test file path")
	}
	testDir := filepath.Dir(testFile)
	// Navigate from scripts/oracle-testing to contracts/src/tokens/tokens.json
	tokensFilePath := filepath.Join(testDir, "..", "..", "src", "tokens", "tokens.json")
	absTokensPath, err := filepath.Abs(tokensFilePath)
	if err != nil {
		t.Fatalf("Failed to resolve absolute path: %v", err)
	}

	t.Logf("Loading tokens from: %s", absTokensPath)

	tokensData, err := os.ReadFile(absTokensPath)
	if err != nil {
		t.Fatalf("Failed to read tokens.json: %v", err)
	}

	var tokenConfigs []TokenConfig
	if err := json.Unmarshal(tokensData, &tokenConfigs); err != nil {
		t.Fatalf("Failed to parse tokens.json: %v", err)
	}

	t.Logf("Loaded %d token configurations from JSON file", len(tokenConfigs))

	// Create storage reader
	storageReader := func(slot common.Hash) ([]byte, error) {
		return client.StorageAt(ctx, TokenRegistryAddress, slot, nil)
	}

	// Compare each token
	mismatches := 0
	notRegistered := 0
	checked := 0

	for _, config := range tokenConfigs {
		t.Logf("\n---------- Comparing tokenID %d: %s (%s) ----------", config.TokenID, config.Name, config.Symbol)

		// Calculate baseSlot
		baseSlot := CalculateUint16MappingSlot(config.TokenID, TokenRegistrySlot)

		// Parse TokenInfo from contract
		contractTokenInfo, err := ParseTokenInfoFromStorage(storageReader, baseSlot)
		if err != nil {
			t.Logf("  [WARN] Failed to parse TokenInfo from contract: %v", err)
			notRegistered++
			continue
		}

		// Check if token is registered (non-zero address)
		if contractTokenInfo.TokenAddress == (common.Address{}) {
			t.Logf("  [SKIP] TokenID %d is not registered in contract", config.TokenID)
			notRegistered++
			continue
		}

		checked++
		hasMismatch := false

		// Compare decimals
		if contractTokenInfo.Decimals != config.Decimals {
			t.Errorf("  [MISMATCH] TokenID %d Decimals: contract=%d, JSON=%d", config.TokenID, contractTokenInfo.Decimals, config.Decimals)
			hasMismatch = true
		} else {
			t.Logf("  ✓ Decimals: %d", contractTokenInfo.Decimals)
		}

		// Compare balanceSlot
		expectedBalanceSlot := common.BigToHash(big.NewInt(int64(config.BalanceSlot)))
		if contractTokenInfo.BalanceSlot != expectedBalanceSlot {
			t.Errorf("  [MISMATCH] TokenID %d BalanceSlot: contract=%s, JSON=%s", config.TokenID, contractTokenInfo.BalanceSlot.Hex(), expectedBalanceSlot.Hex())
			hasMismatch = true
		} else {
			t.Logf("  ✓ BalanceSlot: %s", contractTokenInfo.BalanceSlot.Hex())
		}

		// Compare scale
		expectedScale, ok := new(big.Int).SetString(config.Scale, 10)
		if !ok {
			t.Errorf("  [ERROR] TokenID %d: Invalid scale in JSON: %s", config.TokenID, config.Scale)
			hasMismatch = true
		} else if contractTokenInfo.Scale.Cmp(expectedScale) != 0 {
			t.Errorf("  [MISMATCH] TokenID %d Scale: contract=%s, JSON=%s", config.TokenID, contractTokenInfo.Scale.String(), expectedScale.String())
			hasMismatch = true
		} else {
			t.Logf("  ✓ Scale: %s", contractTokenInfo.Scale.String())
		}

		// Read and compare priceRatio
		priceRatioSlot := CalculateUint16MappingSlot(config.TokenID, PriceRatioSlot)
		priceRatioData, err := client.StorageAt(ctx, TokenRegistryAddress, priceRatioSlot, nil)
		if err != nil {
			t.Logf("  [WARN] Failed to read priceRatio: %v", err)
		} else {
			contractPriceRatio := new(big.Int).SetBytes(priceRatioData)
			expectedPriceRatio, ok := new(big.Int).SetString(config.PriceRatio, 10)
			if !ok {
				t.Errorf("  [ERROR] TokenID %d: Invalid priceRatio in JSON: %s", config.TokenID, config.PriceRatio)
				hasMismatch = true
			} else if contractPriceRatio.Cmp(expectedPriceRatio) != 0 {
				t.Errorf("  [MISMATCH] TokenID %d PriceRatio: contract=%s, JSON=%s", config.TokenID, contractPriceRatio.String(), expectedPriceRatio.String())
				hasMismatch = true
			} else {
				t.Logf("  ✓ PriceRatio: %s", contractPriceRatio.String())
			}
		}

		// Log token address
		t.Logf("  TokenAddress: %s", contractTokenInfo.TokenAddress.Hex())

		if hasMismatch {
			mismatches++
		} else {
			t.Logf("  ✓ All fields match for tokenID %d", config.TokenID)
		}
	}

	// Summary
	t.Logf("\n==========================================")
	t.Logf("Comparison Summary:")
	t.Logf("==========================================")
	t.Logf("Total tokens in JSON: %d", len(tokenConfigs))
	t.Logf("Tokens checked: %d", checked)
	t.Logf("Tokens not registered: %d", notRegistered)
	t.Logf("Tokens with mismatches: %d", mismatches)
	t.Logf("Tokens matching: %d", checked-mismatches)
	t.Logf("==========================================")

	if mismatches > 0 {
		t.Errorf("Found %d token(s) with mismatches", mismatches)
	}
}
