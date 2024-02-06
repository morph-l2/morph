package immutables

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/morph-l2/morph-deployer/op-chain-ops/deployer"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind/backends"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/rlp"
	"github.com/scroll-tech/go-ethereum/trie"
)

// ImmutableValues represents the values to be set in immutable code.
// The key is the name of the variable and the value is the value to set in
// immutable code.
type ImmutableValues map[string]any

// ImmutableConfig represents the immutable configuration for the L2 predeploy
// contracts.
type ImmutableConfig map[string]ImmutableValues

// Check does a sanity check that the specific values that
// Morph uses are set inside of the ImmutableConfig.
func (i ImmutableConfig) Check() error {
	if _, ok := i["GasPriceOracleOwner"]["owner"]; !ok {
		return errors.New("GasPriceOracleOwner owner not set")
	}
	if _, ok := i["L2CrossDomainMessenger"]["otherMessenger"]; !ok {
		return errors.New("L2CrossDomainMessenger otherMessenger not set")
	}
	if _, ok := i["L2ToL1MessagePasser"]["recAddr"]; !ok {
		return errors.New("L2ToL1MessagePasser otherMessenger not set")
	}
	if _, ok := i["L2Sequencer"]["otherSequencer"]; !ok {
		return errors.New("L2Sequencer otherSequencer not set")
	}
	if _, ok := i["L2TxFeeVault"]["recipient"]; !ok {
		return errors.New("L2TxFeeVault recipient not set")
	}
	return nil
}

// DeploymentResults represents the output of deploying each of the
// contracts so that the immutables can be set properly in the bytecode.
type DeploymentResults map[string]hexutil.Bytes

type SlotResults map[string]map[common.Hash]common.Hash

// BuildMorph will deploy the L2 predeploys so that their immutables are set
// correctly.
func BuildMorph(immutable ImmutableConfig, config *Config) (DeploymentResults, SlotResults, error) {
	if err := immutable.Check(); err != nil {
		return DeploymentResults{}, nil, err
	}

	deployments := []deployer.Constructor{
		{
			Name: "GasPriceOracle",
			Args: []interface{}{
				immutable["GasPriceOracleOwner"]["owner"],
			},
		},
		{
			Name: "L2CrossDomainMessenger",
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "L2TxFeeVault",
			Args: []interface{}{
				immutable["L2TxFeeVault"]["owner"],
				immutable["L2TxFeeVault"]["recipient"],
				immutable["L2TxFeeVault"]["minWithdrawalAmount"],
			},
		},
		{
			Name: "L2Sequencer",
			Args: []interface{}{
				immutable["L2Sequencer"]["otherSequencer"],
			},
		},
		{
			Name: "Gov",
		},
		{
			Name: "Submitter",
			Args: []interface{}{
				immutable["Submitter"]["rollupAddr"],
			},
		},
		{
			Name: "L2TokenImplementation",
		},
		{
			Name: "L2TokenFactory",
		},
	}
	return BuildL2(deployments, config)
}

// BuildL2 will deploy contracts to a simulated backend so that their immutables
// can be properly set. The bytecode returned in the results is suitable to be
// inserted into the state via state surgery.
func BuildL2(constructors []deployer.Constructor, config *Config) (DeploymentResults, SlotResults, error) {
	backend := deployer.NewBackend()
	deployments, err := deployer.Deploy(backend, constructors, l2Deployer)
	if err != nil {
		return nil, nil, err
	}
	results := make(DeploymentResults)
	var lastTx *types.Transaction
	for _, dep := range deployments {
		results[dep.Name] = dep.Bytecode

		switch dep.Name {
		case "L2Sequencer":
			if config == nil || len(config.L2SequencerAddresses) == 0 {
				continue
			}
			if len(config.L2SequencerAddresses) != len(config.L2SequencerTmKeys) ||
				len(config.L2SequencerAddresses) != len(config.L2SequencerBlsKeys) {
				return nil, nil, fmt.Errorf("wrong L2Sequencer infos config: inconsistent number")
			}
			infos := make([]bindings.TypesSequencerInfo, len(config.L2SequencerAddresses))
			for i, addr := range config.L2SequencerAddresses {
				infos[i] = bindings.TypesSequencerInfo{
					Addr:   addr,
					TmKey:  config.L2SequencerTmKeys[i],
					BlsKey: config.L2SequencerBlsKeys[i],
				}
			}
			opts, err := bind.NewKeyedTransactorWithChainID(deployer.TestKey, deployer.ChainID)
			if err != nil {
				return nil, nil, err
			}
			l2Sequencer, err := bindings.NewL2Sequencer(dep.Address, backend)
			if err != nil {
				return nil, nil, err
			}
			lastTx, err = l2Sequencer.Initialize(opts, infos)
			if err != nil {
				return nil, nil, err
			}
		default:
		}
	}
	slotResults := make(SlotResults)
	if lastTx != nil {
		backend.Commit()
		if _, err = bind.WaitMined(context.Background(), backend, lastTx); err != nil {
			return nil, nil, err
		}
		stateDB, err := backend.Blockchain().State()
		if err != nil {
			return nil, nil, err
		}
		for _, dep := range deployments {
			st := stateDB.StorageTrie(dep.Address)
			if st == nil {
				return nil, nil, fmt.Errorf("missing account %s in state, address: %s", dep.Name, dep.Address)
			}
			iter := trie.NewIterator(st.NodeIterator(nil))
			slotResults[dep.Name] = make(map[common.Hash]common.Hash)
			for iter.Next() {
				_, data, _, err := rlp.Split(iter.Value)
				if err != nil {
					return nil, nil, err
				}

				slotKey := common.BytesToHash(st.GetKey(iter.Key))
				slotValue := common.BytesToHash(data)
				slotResults[dep.Name][slotKey] = slotValue
			}
		}
	}

	return results, slotResults, nil
}

func l2Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (*types.Transaction, error) {
	var tx *types.Transaction
	var err error
	switch deployment.Name {
	case "GasPriceOracle":
		owner, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for Owner")
		}
		_, tx, _, err = bindings.DeployGasPriceOracle(opts, backend, owner)
	case "L2CrossDomainMessenger":
		_, tx, _, err = bindings.DeployL2CrossDomainMessenger(opts, backend)
	case "L2Sequencer":
		otherSequencer, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherSequencer")
		}
		_, tx, _, err = bindings.DeployL2Sequencer(opts, backend, otherSequencer)
	case "Gov":
		_, tx, _, err = bindings.DeployGov(opts, backend)
	case "Submitter":
		rollupAddr, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for rollup address")
		}
		_, tx, _, err = bindings.DeploySubmitter(opts, backend, rollupAddr)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		_, tx, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "L2TxFeeVault":
		owner, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for owner")
		}
		recipient, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for recipient")
		}
		minWithdrawalAmount, ok := deployment.Args[2].(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid type for minWithdrawalAmount")
		}
		_, tx, _, err = bindings.DeployL2TxFeeVault(opts, backend, owner, recipient, new(big.Int).SetUint64(minWithdrawalAmount))
	case "L2TokenImplementation":
		_, tx, _, err = bindings.DeployMorphStandardERC20(opts, backend)
	case "L2TokenFactory":
		_, tx, _, err = bindings.DeployMorphStandardERC20Factory(opts, backend, predeploys.L2TokenImplementationAddr)
	default:
		return tx, fmt.Errorf("unknown contract: %s", deployment.Name)
	}

	return tx, err
}
