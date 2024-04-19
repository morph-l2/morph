// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const StakingStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var StakingStorageLayout = new(solc.StorageLayout)

var StakingDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(StakingStorageLayoutJSON), StakingStorageLayout); err != nil {
		panic(err)
	}

	layouts["Staking"] = StakingStorageLayout
	deployedBytecodes["Staking"] = StakingDeployedBin
}
