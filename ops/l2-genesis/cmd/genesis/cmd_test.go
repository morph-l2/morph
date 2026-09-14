package genesis

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/urfave/cli"

	genesisconfig "morph-l2/morph-deployer/morph-chain-ops/genesis"
)

func outputContext(t *testing.T, outputs ...string) *cli.Context {
	t.Helper()
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	for index, name := range []string{"outfile.l2", "outfile.rollup", "outfile.genbatchheader"} {
		value := ""
		if index < len(outputs) {
			value = outputs[index]
		}
		set.String(name, value, "")
	}
	return cli.NewContext(cli.NewApp(), set, nil)
}

func TestOutputPathsPreserveExistingFiles(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "genesis.json")
	if err := os.WriteFile(path, []byte("original"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := validateOutputPaths(outputContext(t, path)); err == nil || !strings.Contains(err.Error(), "preserved") {
		t.Fatalf("expected existing output rejection, got %v", err)
	}
	if err := writeGenesisFile(path, map[string]int{"chainId": 900}); err == nil {
		t.Fatal("writeGenesisFile replaced an existing file")
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "original" {
		t.Fatalf("existing output changed: %q, %v", data, err)
	}
}

func TestOutputPathsAreDistinctAndHaveExistingParents(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "genesis.json")
	if err := validateOutputPaths(outputContext(t, path, path)); err == nil {
		t.Fatal("duplicate output paths accepted")
	}
	if err := validateOutputPaths(outputContext(t, filepath.Join(directory, "missing", "genesis.json"))); err == nil {
		t.Fatal("missing output directory accepted")
	}
	if err := validateOutputPaths(outputContext(t, path, filepath.Join(directory, "rollup.json"))); err != nil {
		t.Fatal(err)
	}
}

func TestNetworkConfigurations(t *testing.T) {
	for _, network := range []string{"devnet", "qanet", "testnet", "holesky", "hoodi", "mainnet"} {
		t.Run(network, func(t *testing.T) {
			path := filepath.Join("..", "..", "deploy-config", network+"-deploy-config.json")
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var source map[string]json.RawMessage
			if err := json.Unmarshal(data, &source); err != nil {
				t.Fatal(err)
			}
			config, err := genesisconfig.NewDeployConfig(path)
			if err != nil {
				t.Fatal(err)
			}
			if config.L1ChainID == 0 || config.L2ChainID == 0 || config.L1StartingBlockTag == nil {
				t.Fatal("chain IDs and L1 starting block must be configured")
			}
			err = config.Check()
			if source["l2SequencerAddresses"] != nil && source["l2StakingAddresses"] == nil {
				// Historical templates lack current role and staking parameters.
				// Keep their recorded values; operators must supply a complete input.
				if err == nil {
					t.Fatal("incomplete historical configuration was accepted")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if len(config.L2StakingAddresses) != len(config.L2StakingTmKeys) || len(config.L2StakingAddresses) != len(config.L2StakingBlsKeys) {
				t.Fatal("staking addresses and consensus keys must have matching lengths")
			}
		})
	}
}

func TestCLIRejectsMissingFlagsAndInvalidIdentityBeforeReadingDeployments(t *testing.T) {
	app := cli.NewApp()
	app.Commands = []cli.Command{{Name: "genesis", Subcommands: Subcommands}}
	if err := app.Run([]string{"morph-genesis", "genesis", "l2"}); err == nil {
		t.Fatal("required input flags were accepted without values")
	}
	for _, test := range []struct {
		name, config, expected string
	}{
		{"zero_chain", `{"l1ChainID":0,"l2ChainID":53077,"l1StartingBlockTag":"earliest"}`, "must be positive"},
		{"missing_start", `{"l1ChainID":900,"l2ChainID":53077}`, "l1StartingBlockTag"},
	} {
		t.Run(test.name, func(t *testing.T) {
			directory := t.TempDir()
			input := filepath.Join(directory, "config.json")
			if err := os.WriteFile(input, []byte(test.config), 0o600); err != nil {
				t.Fatal(err)
			}
			err := app.Run([]string{"morph-genesis", "genesis", "l2", "--l1-rpc", "http://127.0.0.1:1",
				"--deploy-config", input, "--deployment-dir", filepath.Join(directory, "missing-deployments.json"),
				"--outfile.l2", filepath.Join(directory, "genesis.json"), "--outfile.rollup", filepath.Join(directory, "rollup.json")})
			if err == nil || !strings.Contains(err.Error(), test.expected) {
				t.Fatalf("expected %q before reading deployments, got %v", test.expected, err)
			}
		})
	}
}

func runCLIExpectingNoArtifacts(t *testing.T, config, deployments []byte, rpcURL string) error {
	t.Helper()
	directory := t.TempDir()
	configPath := filepath.Join(directory, "config.json")
	deploymentPath := filepath.Join(directory, "deployments.json")
	for path, content := range map[string][]byte{configPath: config, deploymentPath: deployments} {
		if content == nil {
			continue
		}
		if err := os.WriteFile(path, content, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	app := cli.NewApp()
	app.Commands = []cli.Command{{Name: "genesis", Subcommands: Subcommands}}
	arguments := []string{"morph-genesis", "genesis", "l2", "--l1-rpc", rpcURL,
		"--deploy-config", configPath, "--deployment-dir", deploymentPath}
	for _, name := range []string{"outfile.l2", "outfile.rollup", "outfile.genbatchheader"} {
		arguments = append(arguments, "--"+name, filepath.Join(directory, name+".json"))
	}
	err := app.Run(arguments)
	for _, name := range []string{"outfile.l2", "outfile.rollup", "outfile.genbatchheader"} {
		if _, statErr := os.Lstat(filepath.Join(directory, name+".json")); !os.IsNotExist(statErr) {
			t.Errorf("rejected invocation created %s: %v", name, statErr)
		}
	}
	for path, original := range map[string][]byte{configPath: config, deploymentPath: deployments} {
		if original == nil {
			continue
		}
		actual, readErr := os.ReadFile(path)
		if readErr != nil || !bytes.Equal(actual, original) {
			t.Errorf("rejected invocation changed input %s: %v", path, readErr)
		}
	}
	return err
}

func TestCLIRejectsInvalidLegacyDeploymentsBeforeRPC(t *testing.T) {
	var calls atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		http.Error(w, "RPC must not be called", http.StatusInternalServerError)
	}))
	defer server.Close()
	address := "0x0000000000000000000000000000000000000001"
	valid := fmt.Sprintf(`{"name":"Proxy__L1Staking","address":%q,"number":1}`, address)
	for _, test := range []struct {
		name, records, configAddress, expected string
	}{
		{"missing_file", "", address, "no such file"},
		{"missing_record", "[]", address, "confirmed Proxy__L1Staking"},
		{"submitter_only", "[" + strings.ReplaceAll(valid, "Proxy__L1Staking", "Proxy__Submitter") + "]", address, "confirmed Proxy__L1Staking"},
		{"pending", "[" + strings.ReplaceAll(valid, `"number":1`, `"number":1,"pending":true`) + "]", address, "pending"},
		{"submitter_alias", "[" + valid + "," + strings.ReplaceAll(valid, "Proxy__L1Staking", "Proxy__Submitter") + "]", address, "must not alias"},
		{"placeholder", "[" + strings.ReplaceAll(valid, address, "0x000000000000000000000000000000000000dEaD") + "]", "0x000000000000000000000000000000000000dEaD", "placeholder"},
		{"different_config", "[" + valid + "]", "0x0000000000000000000000000000000000000002", "differs from confirmed"},
		{"zero_address", "[" + strings.ReplaceAll(valid, address, "0x0000000000000000000000000000000000000000") + "]", address, "nonzero address"},
		{"not_array", `{}`, address, "invalid deployment records"},
		{"null_record", `[null]`, address, "requires a name"},
		{"case_variant_name", "[" + strings.ReplaceAll(valid, `"name"`, `"Name"`) + "]", address, "requires a name"},
		{"case_variant_number", "[" + strings.ReplaceAll(valid, `"number"`, `"Number"`) + "]", address, "block number"},
		{"invalid_address", "[" + strings.ReplaceAll(valid, address, "0x12") + "]", address, "invalid deployment records"},
		{"missing_number", "[" + strings.ReplaceAll(valid, `,"number":1`, "") + "]", address, "block number"},
		{"fractional_number", "[" + strings.ReplaceAll(valid, `"number":1`, `"number":1.5`) + "]", address, "invalid deployment records"},
		{"negative_number", "[" + strings.ReplaceAll(valid, `"number":1`, `"number":-1`) + "]", address, "invalid deployment records"},
		{"imprecise_number", "[" + strings.ReplaceAll(valid, `"number":1`, `"number":9007199254740993`) + "]", address, "exactly representable"},
		{"malformed_unrelated_record", "[" + valid + `,{"name":12,"address":"0x0000000000000000000000000000000000000002","number":1}]`, address, "invalid deployment records"},
		{"invalid_pending", "[" + strings.ReplaceAll(valid, `"number":1`, `"number":1,"pending":"false"`) + "]", address, "pending must be a boolean"},
		{"duplicate_name", "[" + valid + "," + valid + "]", address, "duplicate deployment"},
	} {
		t.Run(test.name, func(t *testing.T) {
			config := []byte(fmt.Sprintf(`{"l1ChainID":900,"l2ChainID":53077,"l1StartingBlockTag":"earliest","l1StakingProxy":%q}`, test.configAddress))
			var records []byte
			if test.records != "" {
				records = []byte(test.records)
			}
			err := runCLIExpectingNoArtifacts(t, config, records, server.URL)
			if err == nil || !strings.Contains(err.Error(), test.expected) {
				t.Fatalf("expected %q, got %v", test.expected, err)
			}
			if calls.Load() != 0 {
				t.Fatal("invalid deployment records reached RPC")
			}
		})
	}
}

func TestCLIRejectsPendingDeploymentsBeforeRPC(t *testing.T) {
	var calls atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		http.Error(w, "RPC must not be called", http.StatusInternalServerError)
	}))
	defer server.Close()
	configData, err := os.ReadFile(filepath.Join("..", "..", "deploy-config", "devnet-deploy-config.json"))
	if err != nil {
		t.Fatal(err)
	}
	fields := []struct{ name, field string }{
		{"Proxy__L1Staking", "l1StakingProxy"},
		{"Proxy__L1CrossDomainMessenger", "l1CrossDomainMessengerProxy"},
		{"Proxy__Rollup", "RollupProxy"},
		{"Proxy__L1GatewayRouter", "l1GatewayRouterProxy"},
		{"Proxy__L1StandardERC20Gateway", "l1StandardERC20GatewayProxy"},
		{"Proxy__L1CustomERC20Gateway", "l1CustomERC20GatewayProxy"},
		{"Proxy__L1ReverseCustomGateway", "l1ReverseCustomGatewayProxy"},
		{"Proxy__L1ETHGateway", "l1ETHGatewayProxy"},
		{"Proxy__L1ERC721Gateway", "l1ERC721GatewayProxy"},
		{"Proxy__L1ERC1155Gateway", "l1ERC1155GatewayProxy"},
		{"Proxy__L1WETHGateway", "l1WETHGatewayProxy"},
		{"Impl__WETH", "l1WETH"},
		{"Proxy__L1WithdrawLockERC20Gateway", "l1WithdrawLockERC20Gateway"},
		{"Impl__Submitter", ""},
	}
	for _, pendingName := range []string{"Proxy__Rollup", "Proxy__L1ERC1155Gateway", "Impl__Submitter"} {
		for _, prefilled := range []bool{false, true} {
			for _, broadcast := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/prefilled=%t/broadcast=%t", pendingName, prefilled, broadcast), func(t *testing.T) {
					var config map[string]any
					if err := json.Unmarshal(configData, &config); err != nil {
						t.Fatal(err)
					}
					var records []map[string]any
					for index, field := range fields {
						address := common.BigToAddress(big.NewInt(int64(index + 1))).Hex()
						record := map[string]any{"name": field.name, "address": address, "number": 1}
						if prefilled && field.field != "" {
							config[field.field] = address
						}
						if field.name == pendingName {
							record["number"], record["pending"] = 0, true
							record["deployer"], record["nonce"] = common.BigToAddress(big.NewInt(100)).Hex(), 12
							if broadcast {
								record["transactionHash"] = common.BigToHash(big.NewInt(1)).Hex()
							}
						}
						records = append(records, record)
					}
					input, err := json.Marshal(config)
					if err != nil {
						t.Fatal(err)
					}
					deployments, err := json.Marshal(records)
					if err != nil {
						t.Fatal(err)
					}
					err = runCLIExpectingNoArtifacts(t, input, deployments, server.URL)
					if expected := "deployment " + pendingName + " is pending"; err == nil || !strings.Contains(err.Error(), expected) {
						t.Fatalf("expected %q, got %v", expected, err)
					}
					if calls.Load() != 0 {
						t.Fatal("pending deployment reached RPC")
					}
				})
			}
		}
	}
}

func TestDeploymentRecordsAcceptConfirmedAddresses(t *testing.T) {
	address := common.BigToAddress(big.NewInt(1))
	for _, configured := range []common.Address{{}, address} {
		for _, pending := range []string{"", `,"pending":false`} {
			path := filepath.Join(t.TempDir(), "deployments.json")
			data := []byte(fmt.Sprintf(`[{"name":"Proxy__L1Staking","address":%q,"number":0%s},{"name":"Proxy__Rollup","address":%q,"number":1%s}]`,
				address, pending, common.BigToAddress(big.NewInt(2)), pending))
			if err := os.WriteFile(path, data, 0o600); err != nil {
				t.Fatal(err)
			}
			record, err := validateDeploymentRecords(path, configured)
			if err != nil || record.Address != address || record.Number != 0 {
				t.Fatalf("confirmed deployments were rejected: %v", err)
			}
		}
	}
}

func TestCLIRequiresLegacyContractCodeOnConnectedL1(t *testing.T) {
	config, err := os.ReadFile(filepath.Join("..", "..", "deploy-config", "devnet-deploy-config.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name, code, expected string
		number               uint64
	}{
		{"future_deployment_block", "0x6000", "exceeds L1 current block", 11},
		{"missing_code", "0x", "has no contract code", 1},
		{"zero_code", "0x0000", "has no contract code", 1},
		// Reaching the starting block request proves legacy code was checked first.
		{"confirmed_code", "0x6000", "error getting l1 start block", 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			var records []map[string]any
			for index, name := range []string{"Proxy__L1Staking", "Proxy__L1CrossDomainMessenger", "Proxy__Rollup", "Proxy__L1GatewayRouter",
				"Proxy__L1StandardERC20Gateway", "Proxy__L1CustomERC20Gateway", "Proxy__L1ReverseCustomGateway", "Proxy__L1ETHGateway",
				"Proxy__L1ERC721Gateway", "Proxy__L1ERC1155Gateway", "Proxy__L1WETHGateway", "Impl__WETH", "Proxy__L1WithdrawLockERC20Gateway"} {
				records = append(records, map[string]any{"name": name, "address": common.BigToAddress(big.NewInt(int64(index + 1))), "number": test.number})
			}
			deployments, err := json.Marshal(records)
			if err != nil {
				t.Fatal(err)
			}
			var codeCalls atomic.Int64
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var request struct {
					ID     json.RawMessage   `json:"id"`
					Method string            `json:"method"`
					Params []json.RawMessage `json:"params"`
				}
				if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
					t.Error(err)
					return
				}
				var result any
				switch request.Method {
				case "eth_chainId":
					result = "0x384"
				case "eth_getBlockByNumber":
					if len(request.Params) > 0 && string(request.Params[0]) == `"latest"` {
						result = &types.Header{Number: big.NewInt(10), Difficulty: big.NewInt(1), GasLimit: 30000000, Time: 1700000000}
					}
				case "eth_getCode":
					codeCalls.Add(1)
					if len(request.Params) != 2 || string(request.Params[0]) != `"0x0000000000000000000000000000000000000001"` || string(request.Params[1]) != `"0xa"` {
						t.Errorf("legacy code must be checked at the validated L1 head: %s", request.Params)
					}
					result = test.code
				default:
					t.Errorf("unexpected RPC method %s", request.Method)
				}
				if err := json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "id": request.ID, "result": result}); err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()
			err = runCLIExpectingNoArtifacts(t, config, deployments, server.URL)
			if err == nil || !strings.Contains(err.Error(), test.expected) {
				t.Fatalf("expected %q, got %v", test.expected, err)
			}
			wantCalls := int64(1)
			if test.number > 10 {
				wantCalls = 0
			}
			if codeCalls.Load() != wantCalls {
				t.Fatalf("expected %d legacy code requests, got %d", wantCalls, codeCalls.Load())
			}
		})
	}
}

func TestActiveConfigurationsIgnoreLegacyMetadata(t *testing.T) {
	for _, network := range []string{"devnet", "qanet", "hoodi", "mainnet"} {
		t.Run(network, func(t *testing.T) {
			path := filepath.Join("..", "..", "deploy-config", network+"-deploy-config.json")
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			// New checked-in inputs contain only current fields. Production decoding
			// remains permissive so saved inputs from older deployments still load.
			decoder := json.NewDecoder(bytes.NewReader(data))
			decoder.DisallowUnknownFields()
			var cleaned genesisconfig.DeployConfig
			if err := decoder.Decode(&cleaned); err != nil {
				t.Fatalf("active configuration contains an unknown field: %v", err)
			}

			var legacySource map[string]json.RawMessage
			if err := json.Unmarshal(data, &legacySource); err != nil {
				t.Fatal(err)
			}
			keys := []string{"maxTxPerBlock"}
			if network == "devnet" || network == "qanet" {
				keys = append(keys, "BLOCK_SIGNER_PRIVATE_KEY", "BLOCK_SIGNER_ADDRESS",
					"morphTokenName", "morphTokenSymbol", "morphTokenOwner",
					"morphTokenInitialSupply", "morphTokenDailyInflationRate")
				if network == "devnet" {
					keys = append(keys, "l2StakingPks")
				} else {
					keys = append(keys, "useMPT")
				}
			}
			for _, key := range keys {
				legacySource[key] = json.RawMessage(`{"ignoredLegacyValue":true}`)
			}
			legacyData, err := json.Marshal(legacySource)
			if err != nil {
				t.Fatal(err)
			}
			legacyPath := filepath.Join(t.TempDir(), "legacy-deploy-config.json")
			if err := os.WriteFile(legacyPath, legacyData, 0o600); err != nil {
				t.Fatal(err)
			}
			legacy, err := genesisconfig.NewDeployConfig(legacyPath)
			if err != nil {
				t.Fatalf("saved legacy input must still decode: %v", err)
			}
			if !reflect.DeepEqual(&cleaned, legacy) {
				t.Fatal("removed metadata changed the decoded deployment configuration")
			}
			if !bytes.Equal(deterministicGenesisArtifacts(t, &cleaned), deterministicGenesisArtifacts(t, legacy)) {
				t.Fatal("legacy metadata changed genesis, rollup configuration or genesis batch header")
			}
		})
	}
}

func deterministicGenesisArtifacts(t *testing.T, config *genesisconfig.DeployConfig) []byte {
	t.Helper()
	// Simulate one confirmed L1 deployment and fix only values normally obtained
	// from L1 or the wall clock. No external RPC or deployment is involved.
	for index, field := range []string{
		"L1StakingProxy", "L1CrossDomainMessengerProxy", "RollupProxy", "L1GatewayRouterProxy",
		"L1StandardERC20GatewayProxy", "L1CustomERC20GatewayProxy", "L1ReverseCustomGatewayProxy",
		"L1ETHGatewayProxy", "L1ERC721GatewayProxy", "L1ERC1155GatewayProxy",
		"L1WETHGatewayProxy", "L1WETH", "L1WithdrawLockERC20Gateway",
	} {
		reflect.ValueOf(config).Elem().FieldByName(field).Set(reflect.ValueOf(common.BigToAddress(big.NewInt(int64(index + 100)))))
	}
	if config.L2GenesisBlockTimestamp == 0 {
		config.L2GenesisBlockTimestamp = 1700000000
	}
	l1Header := &types.Header{Number: big.NewInt(1), Time: 1700000000, BaseFee: big.NewInt(1), GasLimit: 30000000}
	l1Block := types.NewBlockWithHeader(l1Header)
	generated, withdrawRoot, err := genesisconfig.BuildL2DeveloperGenesis(config, l1Block, l1Header)
	if err != nil {
		t.Fatal(err)
	}
	block := generated.ToBlock(nil)
	header, err := genesisconfig.GenesisBatchHeader(block.Header())
	if err != nil {
		t.Fatal(err)
	}
	rollup, err := config.RollupConfig(l1Block, block, withdrawRoot, header)
	if err != nil {
		t.Fatal(err)
	}
	artifacts, err := json.Marshal(map[string]any{"genesis": generated, "rollup": rollup, "header": header})
	if err != nil {
		t.Fatal(err)
	}
	return artifacts
}
