// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2ERC20GatewayStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var L2ERC20GatewayStorageLayout = new(solc.StorageLayout)

var L2ERC20GatewayDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(L2ERC20GatewayStorageLayoutJSON), L2ERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC20Gateway"] = L2ERC20GatewayStorageLayout
	deployedBytecodes["L2ERC20Gateway"] = L2ERC20GatewayDeployedBin
}
