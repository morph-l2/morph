// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1ERC20GatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1007,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1012,\"contract\":\"contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_array(t_uint256)1015_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC20GatewayStorageLayout = new(solc.StorageLayout)

var L1ERC20GatewayDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(L1ERC20GatewayStorageLayoutJSON), L1ERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC20Gateway"] = L1ERC20GatewayStorageLayout
	deployedBytecodes["L1ERC20Gateway"] = L1ERC20GatewayDeployedBin
}
