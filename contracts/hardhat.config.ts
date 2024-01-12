import "@typechain/hardhat"
import "hardhat-abi-exporter"
import "@solidstate/hardhat-bytecode-exporter"
import "solidity-coverage"
import "@nomiclabs/hardhat-ethers"
import "@openzeppelin/hardhat-upgrades"
import "hardhat-preprocessor"
import fs from "fs"
import { TASK_COMPILE_SOLIDITY_GET_SOURCE_PATHS } from "hardhat/builtin-tasks/task-names";
import { subtask } from "hardhat/config";

import "./tasks/deploy"
import "./tasks/overflow_test"
import "./tasks/check"
import "./tasks/query"
import "./tasks/proxy_upgrade"
import "./src/plugin"

function getRemappings() {
  return fs
    .readFileSync("remappings.txt", "utf8")
    .split("\n")
    .filter(Boolean) // remove empty lines
    .map((line) => line.trim().split("="));
}

// prune forge style tests from hardhat paths
subtask(TASK_COMPILE_SOLIDITY_GET_SOURCE_PATHS).setAction(async (_, __, runSuper) => {
  const paths = await runSuper();
  return paths.filter((p: string) => !p.endsWith(".t.sol")).filter((p: string) => !p.includes("test/mocks"));
});

const DEPLOYER_PK = process.env.DEPLOYER_PRIVATE_KEY || '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'

module.exports = {
  defaultNetwork: 'hardhat',
  defender: {
    apiKey: "[apiKey]",
    apiSecret: "[apiSecret]",
  },
  abiExporter: {
    path: './abi',
    runOnCompile: true,
    clear: true,
  },
  bytecodeExporter: {
    path: './data',
    runOnCompile: true,
    clear: true,
    // flat: true,
    // only: [':ERC20$'],
  },
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    local: {
      url: "http://localhost:8545",
      chainId: 31337,
      gas: 'auto',
      gasPrice: 'auto',
    },
    l1: {
      url: "http://localhost:9545",
      chainId: 900,
      gas: 'auto',
      gasPrice: 'auto',
      accounts: [DEPLOYER_PK]
    },
    l2: {
      url: "http://localhost:8545",
      chainId: 53077,
      gas: 'auto',
      gasPrice: 'auto',
      accounts: [DEPLOYER_PK]
    },
    qanetl1: {
      url: "http://l2-qa-morph-l1-geth.bitkeep.tools",
      chainId: 900,
      gas: 'auto',
      gasPrice: 'auto',
      accounts: [DEPLOYER_PK]
    },
    testnetl1: {
      url: "http://10.11.63.110:8545",
      chainId: 900,
      gas: 'auto',
      gasPrice: 'auto',
      accounts: [DEPLOYER_PK]
    },
    sepolia: {
      url: "http://10.11.63.110:8545",
      chainId: 11155111,
      gas: 'auto',
      gasPrice: 'auto',
      accounts: [DEPLOYER_PK]
    }
  },
  foundry: {
    buildInfo: true,
  },
  solidity: {
    compilers: [
      {
        version: '0.8.16',
        settings: {
          optimizer: { enabled: true, runs: 10_000 },
        },
      },
      {
        version: '0.5.17',
        settings: {
          optimizer: { enabled: true, runs: 10_000 },
        },
      },
    ],
  },
  // gasReporter: {
  //     enabled: true,
  //     showMethodSig: true,
  //     maxMethodDiff: 10,
  // },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false,
  },
  preprocess: {
    eachLine: (hre) => ({
      transform: (line: string) => {
        if (line.match(/^\s*import /i)) {
          for (const [from, to] of getRemappings()) {
            if (line.includes(from)) {
              line = line.replace(from, to);
              break;
            }
          }
        }
        return line;
      },
    }),
  },
  paths: {
    sources: "./contracts",
    tests: "./integration-test",
    cache: "./cache-hardhat",
    artifacts: "./artifacts",
    deployConfig: './src/deploy-config',
  }
}