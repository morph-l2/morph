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
	if _, ok := i["L2Sequencer"]["otherSequencer"]; !ok {
		return errors.New("L2Sequencer otherSequencer not set")
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
		},
		{
			Name: "L2CrossDomainMessenger",
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "L2TxFeeVault",
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
		},
		{
			Name: "L2GatewayRouter",
		},
		{
			Name: "L2ETHGateway",
		},
		{
			Name: "L2StandardERC20Gateway",
		},
		{
			Name: "L2ERC721Gateway",
		},
		{
			Name: "L2ERC1155Gateway",
		},
		{
			Name: "MorphStandardERC20",
		},
		{
			Name: "MorphStandardERC20Factory",
		},
		{
			Name: "ProxyAdmin",
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
		case "Submitter":
			if config == nil || len(config.L2SequencerAddresses) == 0 {
				continue
			}
			if len(config.L2SequencerAddresses) != len(config.L2SequencerTmKeys) ||
				len(config.L2SequencerAddresses) != len(config.L2SequencerBlsKeys) {
				return nil, nil, fmt.Errorf("wrong L2Sequencer infos config: inconsistent number")
			}
			opts, err := bind.NewKeyedTransactorWithChainID(deployer.TestKey, deployer.ChainID)
			if err != nil {
				return nil, nil, err
			}
			submitter, err := bindings.NewSubmitter(dep.Address, backend)
			if err != nil {
				return nil, nil, err
			}
			lastTx, err = submitter.Initialize(opts, config.L2SequencerAddresses, big.NewInt(int64(config.L2GenesisBlockTimestamp)))
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
		_, tx, _, err = bindings.DeployGasPriceOracle(opts, backend, common.BigToAddress(common.Big1))
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
		_, tx, _, err = bindings.DeploySubmitter(opts, backend)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		_, tx, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "L2TxFeeVault":
		_, tx, _, err = bindings.DeployL2TxFeeVault(opts, backend, common.BigToAddress(common.Big1), common.BigToAddress(common.Big1), common.Big0)
	case "MorphStandardERC20":
		_, tx, _, err = bindings.DeployMorphStandardERC20(opts, backend)
	case "MorphStandardERC20Factory":
		_, tx, _, err = bindings.DeployMorphStandardERC20Factory(opts, backend, predeploys.MorphStandardERC20Addr)
	case "L2GatewayRouter":
		_, tx, _, err = bindings.DeployL2GatewayRouter(opts, backend)
	case "L2ETHGateway":
		_, tx, _, err = bindings.DeployL2ETHGateway(opts, backend)
	case "L2StandardERC20Gateway":
		_, tx, _, err = bindings.DeployL2StandardERC20Gateway(opts, backend)
	case "L2ERC721Gateway":
		_, tx, _, err = bindings.DeployL2ERC721Gateway(opts, backend)
	case "L2ERC1155Gateway":
		_, tx, _, err = bindings.DeployL2ERC1155Gateway(opts, backend)
	case "ProxyAdmin":
		_, tx, _, err = bindings.DeployProxyAdmin(opts, backend)
	default:
		return tx, fmt.Errorf("unknown contract: %s", deployment.Name)
	}
	return tx, err
}
