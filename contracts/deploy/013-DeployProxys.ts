import "@nomiclabs/hardhat-web3"
import "@nomiclabs/hardhat-ethers"
import "@nomiclabs/hardhat-waffle"
import { Mutex } from 'async-mutex';

import { HardhatRuntimeEnvironment } from "hardhat/types"
import { storage, getContractAddressByName, assertContractVariableWithSigner } from "../src/deploy-utils"
import { ImplStorageName, ProxyStorageName, ContractFactoryName } from "../src/types"

export const deployContractProxyByStorageName = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string
): Promise<string> => {
    const emptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract)
    const ProxyFactoryName = ContractFactoryName.DefaultProxy

    const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
    const proxy = await ProxyFactory.deploy(emptyContractImplAddr, await deployer.getAddress(), "0x")
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", storageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash)
    // check params
    const IProxyContract = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, proxy.address)

    console.log(await IProxyContract.admin())
    await assertContractVariableWithSigner(IProxyContract, "admin", await deployer.getAddress())
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    const err = await storage(path, storageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != "") {
        return err
    }
    // return
    return ""
}

export const deployContractProxies = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    const L1CrossDomainMessengerStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStorageName
    const L1MessageQueueWithGasPriceOracleProxyStorageName =
        ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName

    const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
    const L1StakingProxyStorageName = ProxyStorageName.L1StakingProxyStorageName

    const L1GatewayRouterProxyStorageName = ProxyStorageName.L1GatewayRouterProxyStorageName
    const L1ETHGatewayProxyStorageName = ProxyStorageName.L1ETHGatewayProxyStorageName
    const L1StandardERC20GatewayProxyStorageName = ProxyStorageName.L1StandardERC20GatewayProxyStorageName
    const L1CustomERC20GatewayProxyStorageName = ProxyStorageName.L1CustomERC20GatewayProxyStorageName
    const L1WithdrawLockERC20GatewayProxyStorageName = ProxyStorageName.L1WithdrawLockERC20GatewayProxyStorageName
    const L1ReverseCustomGatewayProxyStorageName = ProxyStorageName.L1ReverseCustomGatewayProxyStorageName
    const L1ERC721GatewayProxyStorageName = ProxyStorageName.L1ERC721GatewayProxyStorageName
    const L1ERC1155GatewayProxyStorageName = ProxyStorageName.L1ERC1155GatewayProxyStorageName
    const EnforcedTxGatewayProxyStorageName = ProxyStorageName.EnforcedTxGatewayProxyStorageName
    const L1WETHGatewayProxyStorageName = ProxyStorageName.L1WETHGatewayProxyStorageName
    const L1USDCGatewayProxyStorageName = ProxyStorageName.L1USDCGatewayProxyStorageName

    const WETHFactoryName = ContractFactoryName.WETH
    const WETHImplStorageName = ImplStorageName.WETH

    let err = ""

    // ************************ token contracts deploy ************************
    if (config.l1WETHAddress == "") {
        // L1WETH deploy
        let Factory = await hre.ethers.getContractFactory(WETHFactoryName)
        let contract = await Factory.deploy()
        await contract.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            WETHImplStorageName,
            contract.address.toLocaleLowerCase(),
            contract.deployTransaction.hash
        )
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        err = await storage(path, WETHImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
    } else {
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        err = await storage(path, WETHImplStorageName, config.l1WETHAddress.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
    }

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessengerProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1CrossDomainMessengerStorageName)
    if (err != "") {
        return err
    }

    // L1MessageQueueWithGasPriceOracleProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ staking contracts deploy ************************
    // StakingProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1StakingProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // RollupProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, RollupProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ gateway contracts deploy ************************
    // L1GatewayRouterProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1GatewayRouterProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ETHGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1StandardERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1CustomERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1CustomERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1WithdrawLockERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1WithdrawLockERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ReverseCustomERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ReverseCustomGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ERC721GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC721GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ERC1155GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC1155GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // EnforcedTxGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, EnforcedTxGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1WETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1WETHGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1USDCGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1USDCGatewayProxyStorageName)
    if (err != "") {
        return err
    }
    // return nil
    return ""
}

export const deployContractProxyByStorageNameWithNonce = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string,
    nonce: number 
): Promise<string> => {
    try {
        const emptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract);
        const ProxyFactoryName = ContractFactoryName.DefaultProxy;

        const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName);
        // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
        const proxy = await ProxyFactory.deploy(emptyContractImplAddr, await deployer.getAddress(), "0x", {
            nonce: nonce,
        });
        await proxy.deployed()
        console.log(
            `%s=%s ; TX_HASH: %s`,
            storageName,
            proxy.address.toLocaleLowerCase(),
            proxy.deployTransaction.hash
        );

        // check params
        const IProxyContract = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, proxy.address);
        
        console.log(await IProxyContract.admin());
        await assertContractVariableWithSigner(IProxyContract, "admin", await deployer.getAddress());
        const blockNumber = await hre.ethers.provider.getBlockNumber();
        console.log("BLOCK_NUMBER: %s", blockNumber);
        const err = await storage(path, storageName, proxy.address.toLocaleLowerCase(), blockNumber || 0);
        if (err !== "") {
            throw new Error(err);
        }

        return "";
    } catch (error) {
        console.error(`Error deploying ${storageName}:`, error);
        return error.message || "Unknown error";
    }
};

export const deployContractProxiesConcurrently = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    const WETHFactoryName = ContractFactoryName.WETH;
    const WETHImplStorageName = ImplStorageName.WETH;

    try {
        if (config.l1WETHAddress == "") {
            const Factory = await hre.ethers.getContractFactory(WETHFactoryName);
            const contract = await Factory.deploy();
            await contract.deployed();
            console.log(
                "%s=%s ; TX_HASH: %s",
                WETHImplStorageName,
                contract.address.toLocaleLowerCase(),
                contract.deployTransaction.hash
            );

            const blockNumber = await hre.ethers.provider.getBlockNumber();
            console.log("BLOCK_NUMBER: %s", blockNumber);

            const err = await storage(path, WETHImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0);
            if (err != "") {
                return err;
            }
        } else {
            const blockNumber = await hre.ethers.provider.getBlockNumber();
            const err = await storage(path, WETHImplStorageName, config.l1WETHAddress.toLocaleLowerCase(), blockNumber || 0);
            if (err != "") {
                return err;
            }
        }

        console.log("Start to deploy proxies concurrently...")

        const proxyStorageNames = [
            ProxyStorageName.L1CrossDomainMessengerProxyStorageName,
            ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName,
            ProxyStorageName.L1StakingProxyStorageName,
            ProxyStorageName.RollupProxyStorageName,
            ProxyStorageName.L1GatewayRouterProxyStorageName,
            ProxyStorageName.L1ETHGatewayProxyStorageName,
            ProxyStorageName.L1StandardERC20GatewayProxyStorageName,
            ProxyStorageName.L1CustomERC20GatewayProxyStorageName,
            ProxyStorageName.L1WithdrawLockERC20GatewayProxyStorageName,
            ProxyStorageName.L1ReverseCustomGatewayProxyStorageName,
            ProxyStorageName.L1ERC721GatewayProxyStorageName,
            ProxyStorageName.L1ERC1155GatewayProxyStorageName,
            ProxyStorageName.EnforcedTxGatewayProxyStorageName,
            ProxyStorageName.L1WETHGatewayProxyStorageName,
            ProxyStorageName.L1USDCGatewayProxyStorageName,
        ];

        let nonce = await hre.ethers.provider.getTransactionCount(deployer.getAddress())
        const mutex = new Mutex();
        const results = await Promise.all(
            proxyStorageNames.map(async (storageName) => {
                console.log(`Starting deployment for: ${storageName}`);
                const release = await mutex.acquire(); // Acquire lock for getting nonce 
                const nonceToUse = nonce
                nonce++;  // Increment nonce for each deployment
                release();  // Release the lock
                
                const result = await deployContractProxyByStorageNameWithNonce(hre, path, deployer, storageName, nonceToUse);
                console.log(`Deployment completed for: ${storageName}`);
                return result; 
            })
        );

        const errors = results.filter((err) => err !== "");
        if (errors.length > 0) {
            return `Deployment failed with errors: ${errors.join(", ")}`;
        }
    } catch (error) {
        console.error("Error during deployment:", error);
        return `Deployment failed with error: ${error.message}`;
    }

    return "";
};


export default deployContractProxies