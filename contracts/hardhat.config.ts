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
import "./tasks/staking_upgrade"
import "./src/plugin"
import * as process from "process";

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
const QA_URL = process.env.QA_RPC_URL || 'http://127.0.0.1:8545'
const SEPOLIA_URL = process.env.SEPOLIA_RPC_URL || 'http://127.0.0.1:8545'
const HOLESKY_URL = process.env.HOLESKY_RPC_URL || 'http://127.0.0.1:8545'

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
            chainId: 900,
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
            url: QA_URL,
            chainId: 900,
            gas: 'auto',
            gasPrice: 'auto',
            accounts: [DEPLOYER_PK]
        },
        sepolia: {
            url: SEPOLIA_URL,
            chainId: 11155111,
            gas: 'auto',
            gasPrice: 'auto',
            accounts: [DEPLOYER_PK]
        },
        holesky: {
            url: HOLESKY_URL,
            chainId: 17000,
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
                version: '0.8.24',
                settings: {
                    metadata: { bytecodeHash: 'none' },
                    optimizer: { enabled: true, runs: 10_000 },
                    evmVersion: 'cancun',
                },
            },
        ],
        // compileUsingRemoteVersion: 'v0.8.23+commit.69122e07',
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