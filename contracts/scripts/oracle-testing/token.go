package oracle_testing

import (
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
)

// TokenConfig represents token configuration from JSON file
type TokenConfig struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    uint8  `json:"decimals"`
	TokenID     uint16 `json:"tokenID"`
	BalanceSlot uint64 `json:"balanceSlot"`
	Scale       string `json:"scale"`
	PriceRatio  string `json:"priceRatio"`
}

var TokenRegistryAddress = common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")

var (
	// TokenRegistrySlot is the storage slot for mapping(uint16 => TokenInfo)
	// TokenInfo struct layout:
	//   - tokenAddress: address (offset 0)
	//   - balanceSlot: bytes32 (offset 1)
	//   - isActive: bool (offset 2, byte 0)
	//   - decimals: uint8 (offset 2, byte 1)
	//   - scale: uint256 (offset 3)
	// Based on L2TokenRegistryStorageLayout: slot 151
	TokenRegistrySlot = common.BigToHash(big.NewInt(151))
	// TokenRegistrationSlot is the storage slot for mapping(address => uint16)
	// Based on L2TokenRegistryStorageLayout: slot 152
	TokenRegistrationSlot = common.BigToHash(big.NewInt(152))
	// PriceRatioSlot is the storage slot for mapping(uint16 => uint256)
	// Based on L2TokenRegistryStorageLayout: slot 153
	PriceRatioSlot = common.BigToHash(big.NewInt(153))
	// AllowListSlot is the storage slot for mapping(address => bool)
	// Based on L2TokenRegistryStorageLayout: slot 154
	AllowListSlot = common.BigToHash(big.NewInt(154))
	// AllowListEnabledSlot is the storage slot for bool allowListEnabled
	// Based on L2TokenRegistryStorageLayout: slot 155
	AllowListEnabledSlot = common.BigToHash(big.NewInt(155))
)

// TokenInfo represents the token information structure
type TokenInfo struct {
	TokenAddress common.Address
	BalanceSlot  common.Hash
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}

// CalculateUint16MappingSlot calculates the storage slot for a mapping key
// For mapping(key => value), the slot is: keccak256(abi.encode(key, mappingSlot))
func CalculateUint16MappingSlot(key uint16, mappingSlot common.Hash) common.Hash {
	// Convert key to 32 bytes (right-padded)
	keyBytes := make([]byte, 32)
	keyBytes[30] = byte(key >> 8) // high byte
	keyBytes[31] = byte(key)      // low byte

	// Convert mapping slot to 32 bytes (left-padded)
	slotBytes := mappingSlot.Bytes()
	paddedSlot := make([]byte, 32)
	copy(paddedSlot[32-len(slotBytes):], slotBytes)

	// Concatenate key and slot
	data := append(keyBytes, paddedSlot...)

	// Calculate keccak256 hash
	hash := crypto.Keccak256(data)

	return common.BytesToHash(hash)
}

// CalculateStructFieldSlot calculates the storage slot for a struct field within a mapping
// For a struct at baseSlot, fieldOffset is the offset within the struct
func CalculateStructFieldSlot(baseSlot common.Hash, fieldOffset uint64) common.Hash {
	// Add fieldOffset to baseSlot
	baseInt := new(big.Int).SetBytes(baseSlot[:])
	fieldInt := big.NewInt(int64(fieldOffset))
	result := new(big.Int).Add(baseInt, fieldInt)
	return common.BigToHash(result)
}

// CalculateAddressMappingSlot calculates the storage slot for a mapping key (address type)
// For mapping(address => value), the slot is: keccak256(abi.encode(key, mappingSlot))
func CalculateAddressMappingSlot(key common.Address, mappingSlot common.Hash) common.Hash {
	// Convert address to 32 bytes (left-padded)
	keyBytes := make([]byte, 32)
	copy(keyBytes[12:], key.Bytes())

	// Convert mapping slot to 32 bytes (left-padded)
	slotBytes := mappingSlot.Bytes()
	paddedSlot := make([]byte, 32)
	copy(paddedSlot[32-len(slotBytes):], slotBytes)

	// Concatenate key and slot
	data := append(keyBytes, paddedSlot...)

	// Calculate keccak256 hash
	hash := crypto.Keccak256(data)

	return common.BytesToHash(hash)
}

// ParseTokenInfoFromStorage parses TokenInfo from storage slots
// baseSlot is the base slot for the TokenInfo struct in the mapping
func ParseTokenInfoFromStorage(
	storageAt func(common.Hash) ([]byte, error),
	baseSlot common.Hash,
) (*TokenInfo, error) {
	// Read tokenAddress (offset 0)
	tokenAddrSlot := CalculateStructFieldSlot(baseSlot, 0)
	tokenAddrData, err := storageAt(tokenAddrSlot)
	if err != nil {
		return nil, err
	}
	tokenAddress := common.BytesToAddress(tokenAddrData[12:32])

	// Read balanceSlot (offset 1)
	balanceSlotField := CalculateStructFieldSlot(baseSlot, 1)
	balanceSlotData, err := storageAt(balanceSlotField)
	if err != nil {
		return nil, err
	}
	// Read isActive and decimals (offset 2, packed together)
	isActiveSlot := CalculateStructFieldSlot(baseSlot, 2)
	isActiveData, err := storageAt(isActiveSlot)
	if err != nil {
		return nil, err
	}
	isActive := isActiveData[31] != 0
	decimals := isActiveData[30]

	// Read scale (offset 3)
	scaleSlot := CalculateStructFieldSlot(baseSlot, 3)
	scaleData, err := storageAt(scaleSlot)
	if err != nil {
		return nil, err
	}
	scale := new(big.Int).SetBytes(scaleData)

	return &TokenInfo{
		TokenAddress: tokenAddress,
		BalanceSlot:  common.BytesToHash(balanceSlotData),
		IsActive:     isActive,
		Decimals:     decimals,
		Scale:        scale,
	}, nil
}
