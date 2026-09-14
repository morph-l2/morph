package genesis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"morph-l2/bindings/hardhat"
	"morph-l2/morph-deployer/morph-chain-ops/genesis"
)

var Subcommands = cli.Commands{
	{
		Name:  "l2",
		Usage: "Generates an L2 genesis file and rollup config suitable for a deployed network",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "l1-rpc",
				Required: true,
				Usage:    "L1 RPC URL",
			},
			cli.StringFlag{
				Name:     "deploy-config",
				Required: true,
				Usage:    "Path to L2 deploy configuration JSON",
			},
			cli.StringFlag{
				Name:     "deployment-dir",
				Required: true,
				Usage:    "Path to the L1 deployment records JSON file",
			},
			cli.StringFlag{
				Name:     "outfile.l2",
				Required: true,
				Usage:    "Path to L2 genesis output file",
			},
			cli.StringFlag{
				Name:     "outfile.rollup",
				Required: true,
				Usage:    "Path to rollup output file",
			},
			cli.StringFlag{
				Name:  "outfile.genbatchheader",
				Usage: "Path to genesisBatchHeader output file",
			},
		},
		Action: func(ctx *cli.Context) error {
			deployConfig := ctx.String("deploy-config")
			config, err := genesis.NewDeployConfig(deployConfig)
			if err != nil {
				return err
			}

			if config.L1ChainID == 0 || config.L2ChainID == 0 {
				return fmt.Errorf("l1ChainID and l2ChainID must be positive")
			}
			if config.L1StartingBlockTag == nil || (config.L1StartingBlockTag.BlockHash == nil && config.L1StartingBlockTag.BlockNumber == nil) {
				return fmt.Errorf("l1StartingBlockTag must specify a block number or hash")
			}
			if err := validateOutputPaths(ctx); err != nil {
				return err
			}

			// The historical flag name accepts one deployment records file.
			depPath := ctx.String("deployment-dir")
			legacyStaking, err := validateDeploymentRecords(depPath, config.L1StakingProxy)
			if err != nil {
				return err
			}
			_, network := filepath.Split(depPath)
			network = strings.TrimSuffix(network, filepath.Ext(network))
			hh, err := hardhat.New4SingleDeployment(network, depPath)
			if err != nil {
				return err
			}

			// Read the appropriate deployment addresses from disk
			if err := config.GetDeployedAddresses(hh); err != nil {
				return err
			}

			if err := config.Check(); err != nil {
				return err
			}

			client, err := ethclient.Dial(ctx.String("l1-rpc"))
			if err != nil {
				return fmt.Errorf("cannot dial %s: %w", ctx.String("l1-rpc"), err)
			}

			defer client.Close()
			rpcContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			chainID, err := client.ChainID(rpcContext)
			if err != nil {
				return fmt.Errorf("error getting L1 chain ID: %w", err)
			}
			if chainID.Cmp(new(big.Int).SetUint64(config.L1ChainID)) != 0 {
				return fmt.Errorf("L1 RPC chain ID %s differs from l1ChainID %d", chainID, config.L1ChainID)
			}
			currentL1Header, err := client.HeaderByNumber(rpcContext, nil)
			if err != nil {
				return fmt.Errorf("error getting l1 current header: %w", err)
			}
			if err := verifyLegacyStakingDeployment(rpcContext, client, legacyStaking, currentL1Header); err != nil {
				return err
			}

			var l1StartBlock *types.Block
			if config.L1StartingBlockTag.BlockHash != nil {
				fmt.Printf("using L1StartingBlockTag.BlockHash: %s\n", config.L1StartingBlockTag.BlockHash)
				l1StartBlock, err = client.BlockByHash(rpcContext, *config.L1StartingBlockTag.BlockHash)
			} else if config.L1StartingBlockTag.BlockNumber != nil {
				fmt.Printf("using L1StartingBlockTag.BlockNumber: %d\n", config.L1StartingBlockTag.BlockNumber.Int64())
				l1StartBlock, err = client.BlockByNumber(rpcContext, big.NewInt(config.L1StartingBlockTag.BlockNumber.Int64()))
			}
			if err != nil {
				return fmt.Errorf("error getting l1 start block: %w", err)
			}
			fmt.Printf("The L1 Starting Block Hash: %s \n", l1StartBlock.Hash())

			// Build the developer L2 genesis block
			l2Genesis, withdrawRoot, err := genesis.BuildL2DeveloperGenesis(config, l1StartBlock, currentL1Header)
			if err != nil {
				return fmt.Errorf("error creating l2 developer genesis: %w", err)
			}

			l2GenesisBlock := l2Genesis.ToBlock(nil)

			var genesisBatchHeaderBytes []byte
			genBatchHeaderFile := ctx.String("outfile.genbatchheader")
			if len(genBatchHeaderFile) > 0 {
				var err error
				genesisBatchHeaderBytes, err = genesis.GenesisBatchHeader(l2GenesisBlock.Header())
				if err != nil {
					return err
				}
				fmt.Printf("generated genesis batch header bytes: %x \n", genesisBatchHeaderBytes)
				if err := writeGenesisFile(genBatchHeaderFile, hexutil.Bytes(genesisBatchHeaderBytes)); err != nil {
					return err
				}
			}

			rollupConfig, err := config.RollupConfig(l1StartBlock, l2GenesisBlock, withdrawRoot, genesisBatchHeaderBytes)
			if err != nil {
				return err
			}

			fmt.Printf("The L2 genesis state root: %s \n", l2GenesisBlock.Root().Hex())

			if err := writeGenesisFile(ctx.String("outfile.l2"), l2Genesis); err != nil {
				return err
			}
			return writeGenesisFile(ctx.String("outfile.rollup"), rollupConfig)
		},
	},
}

type deploymentRecord struct {
	Name    string
	Address common.Address
	Number  uint64
	Pending json.RawMessage
}

// Validate every record before the historical Hardhat loader uses unchecked
// type assertions. A configured address cannot replace the legacy deployment record.
func validateDeploymentRecords(path string, configuredStaking common.Address) (*deploymentRecord, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	var entries []map[string]json.RawMessage
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("invalid deployment records: %w", err)
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("deployment records must contain confirmed Proxy__L1Staking")
	}
	records := make([]deploymentRecord, len(entries))
	byName := make(map[string]*deploymentRecord, len(records))
	for index := range records {
		record := &records[index]
		entry := entries[index]
		if entry["name"] == nil || entry["address"] == nil {
			return nil, fmt.Errorf("deployment record %d requires a name and nonzero address", index)
		}
		if entry["number"] == nil || bytes.Equal(entry["number"], []byte("null")) {
			return nil, fmt.Errorf("deployment record %d requires a nonnegative block number", index)
		}
		// Match the exact keys consumed by Hardhat, without case-insensitive struct decoding.
		for key, target := range map[string]any{"name": &record.Name, "address": &record.Address, "number": &record.Number} {
			if err := json.Unmarshal(entry[key], target); err != nil {
				return nil, fmt.Errorf("invalid deployment records: record %d field %s: %w", index, key, err)
			}
		}
		record.Pending = entry["pending"]
		if strings.TrimSpace(record.Name) == "" || record.Address == (common.Address{}) {
			return nil, fmt.Errorf("deployment record %d requires a name and nonzero address", index)
		}
		// The Hardhat loader decodes numbers as float64 and then converts to uint.
		if record.Number > 1<<53-1 || record.Number > uint64(^uint(0)) {
			return nil, fmt.Errorf("deployment %s requires an exactly representable nonnegative block number", record.Name)
		}
		if record.Pending != nil && !bytes.Equal(record.Pending, []byte("true")) && !bytes.Equal(record.Pending, []byte("false")) {
			return nil, fmt.Errorf("deployment %s pending must be a boolean when provided", record.Name)
		}
		if byName[record.Name] != nil {
			return nil, fmt.Errorf("duplicate deployment %s", record.Name)
		}
		byName[record.Name] = record
	}
	staking := byName["Proxy__L1Staking"]
	if staking == nil {
		return nil, fmt.Errorf("confirmed Proxy__L1Staking deployment is required, including when l1StakingProxy is configured")
	}
	if bytes.Equal(staking.Pending, []byte("true")) {
		return nil, fmt.Errorf("Proxy__L1Staking deployment is pending; confirm the legacy deployment before generating genesis")
	}
	if staking.Address == common.HexToAddress("0x000000000000000000000000000000000000dEaD") {
		return nil, fmt.Errorf("Proxy__L1Staking must be the actual legacy deployment, not the dEaD placeholder")
	}
	if submitter := byName["Proxy__Submitter"]; submitter != nil && submitter.Address == staking.Address {
		return nil, fmt.Errorf("Proxy__L1Staking must not alias Proxy__Submitter")
	}
	if configuredStaking != (common.Address{}) && configuredStaking != staking.Address {
		return nil, fmt.Errorf("l1StakingProxy %s differs from confirmed Proxy__L1Staking %s", configuredStaking, staking.Address)
	}
	return staking, nil
}

func verifyLegacyStakingDeployment(ctx context.Context, client *ethclient.Client, staking *deploymentRecord, head *types.Header) error {
	if head == nil || head.Number == nil {
		return fmt.Errorf("L1 current header must contain a block number")
	}
	if new(big.Int).SetUint64(staking.Number).Cmp(head.Number) > 0 {
		return fmt.Errorf("Proxy__L1Staking deployment block %d exceeds L1 current block %s", staking.Number, head.Number)
	}
	code, err := client.CodeAt(ctx, staking.Address, head.Number)
	if err != nil {
		return fmt.Errorf("error checking Proxy__L1Staking code at L1 block %s: %w", head.Number, err)
	}
	if len(bytes.Trim(code, "\x00")) == 0 {
		return fmt.Errorf("Proxy__L1Staking %s has no contract code at L1 block %s", staking.Address, head.Number)
	}
	return nil
}

// Reject existing or repeated outputs before connecting to L1. The shell wrappers
// generate in a temporary directory and separately authorize replacement.
func validateOutputPaths(ctx *cli.Context) error {
	seen := make(map[string]bool)
	for _, name := range []string{"outfile.l2", "outfile.rollup", "outfile.genbatchheader"} {
		value := ctx.String(name)
		if value == "" {
			continue
		}
		path, err := filepath.Abs(value)
		if err != nil {
			return err
		}
		if seen[path] {
			return fmt.Errorf("output flags must use distinct paths: %s", path)
		}
		seen[path] = true
		if _, err := os.Lstat(path); err == nil {
			return fmt.Errorf("existing output is preserved: %s", path)
		} else if !os.IsNotExist(err) {
			return err
		}
		if info, err := os.Stat(filepath.Dir(path)); err != nil || !info.IsDir() {
			return fmt.Errorf("output directory must exist: %s", filepath.Dir(path))
		}
	}
	return nil
}

func writeGenesisFile(outfile string, input any) error {
	f, err := os.OpenFile(filepath.Clean(outfile), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(input)
}
