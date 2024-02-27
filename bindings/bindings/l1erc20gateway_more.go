// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1ERC20GatewayStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var L1ERC20GatewayStorageLayout = new(solc.StorageLayout)

var L1ERC20GatewayDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(L1ERC20GatewayStorageLayoutJSON), L1ERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC20Gateway"] = L1ERC20GatewayStorageLayout
	deployedBytecodes["L1ERC20Gateway"] = L1ERC20GatewayDeployedBin
}
