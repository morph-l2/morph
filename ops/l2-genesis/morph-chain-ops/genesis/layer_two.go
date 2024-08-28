package genesis

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rollup/rcfg"
	"github.com/morph-l2/go-ethereum/rollup/withdrawtrie"

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
	return db.Genesis(), withdrawRoot, nil
}
