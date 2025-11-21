package genesis

import (
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/core/vm"
	"github.com/morph-l2/go-ethereum/rollup/rcfg"
	"github.com/morph-l2/go-ethereum/rollup/withdrawtrie"

	"morph-l2/bindings/predeploys"
	"morph-l2/morph-deployer/morph-chain-ops/state"
)

// BuildL2DeveloperGenesis will build the developer Morph Genesis
// Block. Suitable for devnets.
func BuildL2DeveloperGenesis(config *DeployConfig, l1StartBlock *types.Block, curL1Header *types.Header) (*core.Genesis, common.Hash, error) {
	genspec, err := NewL2Genesis(config, l1StartBlock)
	if err != nil {
		return nil, common.Hash{}, err
	}

	db := state.NewMemoryStateDB(genspec)

	if config.FundDevAccounts {
		FundDevAccounts(db)
		SetPrecompileBalances(db)
	}

	storage, err := NewL2StorageConfig(config, curL1Header.BaseFee)
	if err != nil {
		return nil, common.Hash{}, err
	}

	immutable, imuConfig, err := NewL2ImmutableConfig(config)
	if err != nil {
		return nil, common.Hash{}, err
	}

	if err := SetL2Proxies(db); err != nil {
		return nil, common.Hash{}, err
	}

	SetL2CrossDomainMessengerBalances(db)

	if err := SetImplementations(db, storage, immutable, imuConfig); err != nil {
		return nil, common.Hash{}, err
	}

	withdrawRoot := withdrawtrie.ReadWTRSlot(rcfg.L2MessageQueueAddress, db)

	fmt.Println("get withdraw root:", withdrawRoot)

	// Verify L2TokenRegistry allowListEnabled configuration
	if err = VerifyL2TokenRegistryConfig(db); err != nil {
		return nil, common.Hash{}, fmt.Errorf("L2TokenRegistry verification failed: %w", err)
	}

	return db.Genesis(), withdrawRoot, nil
}

// VerifyL2TokenRegistryConfig verifies that L2TokenRegistry's allowListEnabled is set to true at slot 155
func VerifyL2TokenRegistryConfig(db vm.StateDB) error {
	contractAddr := predeploys.L2TokenRegistryAddr
	// AllowListEnabledSlot = 155
	allowListEnabledSlot := common.BigToHash(big.NewInt(155))

	// Read storage at slot 155
	storageValue := db.GetState(contractAddr, allowListEnabledSlot)

	// For bool type, true is represented as 0x01 (or any non-zero value in the last byte)
	// Check if the last byte is non-zero
	isEnabled := storageValue[31] != 0

	if !isEnabled {
		return fmt.Errorf("L2TokenRegistry allowListEnabled is not set to true at slot 155. Got: %s (value: %d)", storageValue.Hex(), storageValue[31])
	}

	fmt.Printf("✓ L2TokenRegistry allowListEnabled verified: true (slot 155 = %s, value = 0x%02x)\n", allowListEnabledSlot.Hex(), storageValue[31])
	return nil
}
