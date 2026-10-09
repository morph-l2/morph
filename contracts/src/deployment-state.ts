import { ethers } from "ethers";
import fs from "fs";
import nodePath from "path";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { getContractAddressByName, readDeploymentRecords, storage } from "./deploy-utils";
import { ContractFactoryName, ImplStorageName, ProxyStorageName } from "./types";

export const ensureDeploymentStorageWritable = (path: string) => {
    fs.accessSync(nodePath.dirname(nodePath.resolve(path)), fs.constants.W_OK);
    if (fs.existsSync(path)) fs.accessSync(path, fs.constants.R_OK | fs.constants.W_OK);
};

export const ensureNoPendingTransactions = async (hre: HardhatRuntimeEnvironment, signer: any) => {
    const address = await signer.getAddress();
    const latest = await hre.ethers.provider.getTransactionCount(address, "latest");
    const pending = await hre.ethers.provider.getTransactionCount(address, "pending");
    if (pending > latest) {
        throw new Error(`Deployer has pending transactions: latest nonce=${latest}, pending nonce=${pending}; confirm the original transactions before retrying`);
    }
};

export const validateDeploymentSigner = async (hre: HardhatRuntimeEnvironment, path: string, signer: any) => {
    const address = await signer.getAddress();
    await ensureNoPendingTransactions(hre, signer);
    const manager = getContractAddressByName(path, ImplStorageName.ProxyAdmin);
    if (!manager) return;
    const contract = await hre.ethers.getContractAt(ContractFactoryName.ProxyAdmin, manager, signer);
    if ((await contract.owner()).toLowerCase() !== address.toLowerCase()) {
        throw new Error("Deployment signer is not the owner of the recorded ProxyAdmin");
    }
    const managedProxies = new Set(Object.values(ProxyStorageName));
    for (const record of readDeploymentRecords(path).filter(entry => managedProxies.has(entry.name))) {
        const admin = (await readProxyAddress(hre, record.address, "admin")).toLowerCase();
        if (admin !== address.toLowerCase() && admin !== manager.toLowerCase()) {
            throw new Error(`${record.name} proxy admin is neither the deployer nor the recorded ProxyAdmin`);
        }
    }
};

const slot = (name: string) => ethers.utils.hexZeroPad(
    ethers.BigNumber.from(ethers.utils.id(name)).sub(1).toHexString(), 32
);

export const readProxyAddress = async (hre: HardhatRuntimeEnvironment, address: string, field: "admin" | "implementation") =>
    ethers.utils.getAddress(ethers.utils.hexDataSlice(
        await hre.ethers.provider.getStorageAt(address, slot(`eip1967.proxy.${field}`)), 12
    ));

const confirmPendingDeployment = async (hre: HardhatRuntimeEnvironment, path: string, record: any) => {
    if (!record.pending) return;
    if (!ethers.utils.isHexString(record.transactionHash, 32)) {
        throw new Error(`${record.name} broadcast result was not recorded: deployer=${record.deployer} nonce=${record.nonce} expectedAddress=${record.address}; check the original transaction and record transactionHash before retrying; do not redeploy`);
    }
    let receipt = await hre.ethers.provider.getTransactionReceipt(record.transactionHash);
    if (!receipt) {
        try {
            receipt = await hre.ethers.provider.waitForTransaction(record.transactionHash, 1, 60000);
        } catch (_) {
            throw new Error(`${record.name} transaction ${record.transactionHash} is unconfirmed; preserve deployment records and confirm the original transaction before retrying`);
        }
    }
    if (!receipt || receipt.status !== 1 || receipt.contractAddress?.toLowerCase() !== record.address.toLowerCase()) {
        throw new Error(`${record.name} deployment transaction ${record.transactionHash} did not create the expected contract; preserve records and check the receipt`);
    }
    await storage(path, record.name, record.address, receipt.blockNumber, record.transactionHash, false);
};

export const validateDeploymentRecords = async (hre: HardhatRuntimeEnvironment, path: string) => {
    for (const record of readDeploymentRecords(path)) {
        await confirmPendingDeployment(hre, path, record);
        if (await hre.ethers.provider.getCode(record.address) === "0x") {
            throw new Error(`${record.name}: recorded address ${record.address} has no code on the current network; check the network and deployment file`);
        }
    }
};

export const deployRecordedContract = async (
    hre: HardhatRuntimeEnvironment, path: string, signer: any,
    name: string, factoryName: string, args: any[] = []
) => {
    const record = readDeploymentRecords(path).find(entry => entry.name === name);
    if (record) await confirmPendingDeployment(hre, path, record);
    const previous = record?.address;
    if (previous && ethers.utils.isAddress(previous)) {
        if (await hre.ethers.provider.getCode(previous) === "0x") {
            throw new Error(`${name}: ${previous} has no contract code`);
        }
        return hre.ethers.getContractAt(factoryName, previous, signer);
    }
    const factory = await hre.ethers.getContractFactory(factoryName, signer);
    const deployer = await signer.getAddress();
    await ensureNoPendingTransactions(hre, signer);
    const nonce = await signer.getTransactionCount("pending");
    const deployment = factory.getDeployTransaction(...args, { nonce });
    const gasLimit = (await signer.estimateGas(deployment)).mul(120).div(100);
    const feeData = await hre.ethers.provider.getFeeData();
    const fees = feeData.maxFeePerGas && feeData.maxPriorityFeePerGas
        ? { maxFeePerGas: feeData.maxFeePerGas, maxPriorityFeePerGas: feeData.maxPriorityFeePerGas }
        : { gasPrice: feeData.gasPrice || await hre.ethers.provider.getGasPrice() };
    const maximumFee = gasLimit.mul("maxFeePerGas" in fees ? fees.maxFeePerGas : fees.gasPrice);
    if ((await signer.getBalance()).lt(maximumFee.add(deployment.value || 0))) {
        throw new Error(`${name}: insufficient deployer balance for the deployment; no transaction was broadcast`);
    }
    const expected = ethers.utils.getContractAddress({ from: deployer, nonce });
    // Record sender, nonce, and expected address before broadcasting; a missing RPC response must not trigger another deployment.
    await storage(path, name, expected, 0, undefined, true, { deployer, nonce });
    const contract = await factory.deploy(...args, { nonce, gasLimit, ...fees });
    await storage(path, name, contract.address, 0, contract.deployTransaction.hash, true);
    const receipt = await contract.deployTransaction.wait();
    await storage(path, name, contract.address, receipt.blockNumber, receipt.transactionHash, false);
    console.log(`${name}=${contract.address}; TX_HASH: ${receipt.transactionHash}`);
    return contract;
};

// Read EIP-1967 proxy storage slots without requiring the caller to remain the proxy admin.
export const getDeploymentProxy = async (
    hre: HardhatRuntimeEnvironment, path: string, address: string, signer: any
): Promise<any> => {
    const direct = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, address, signer);
    const facade: any = {
        address,
        provider: direct.provider,
        interface: direct.interface,
        connect: () => facade,
        admin: () => readProxyAddress(hre, address, "admin"),
        implementation: () => readProxyAddress(hre, address, "implementation"),
        changeAdmin: (next: string) => direct.changeAdmin(next),
        upgradeToAndCall: async (implementation: string, initializer: string) => {
            const empty = getContractAddressByName(path, ImplStorageName.EmptyContract);
            const current = await facade.implementation();
            if (current.toLowerCase() !== empty.toLowerCase()) {
                throw new Error(`${address}: current implementation ${current} is not the recorded EmptyContract; initialization stopped`);
            }
            const admin = await facade.admin();
            const signerAddress = await signer.getAddress();
            let restoreAdmin = false;
            if (admin.toLowerCase() !== signerAddress.toLowerCase()) {
                const expectedAdmin = getContractAddressByName(path, ImplStorageName.ProxyAdmin);
                if (admin.toLowerCase() !== expectedAdmin.toLowerCase()) {
                    throw new Error(`${address}: proxy admin ${admin} does not match deployment records`);
                }
                const manager = await hre.ethers.getContractAt(ContractFactoryName.ProxyAdmin, admin, signer);
                if ((await manager.owner()).toLowerCase() !== signerAddress.toLowerCase()) {
                    throw new Error(`${address}: signer is not ProxyAdmin.owner()`);
                }
                // __Ownable_init uses msg.sender, so restore the deployer as proxy admin before initialize.
                // If a later transaction fails, the deployer remains admin and a retry can continue initialization.
                await (await manager.changeProxyAdmin(address, signerAddress)).wait();
                restoreAdmin = true;
            }
            const transaction = await direct.upgradeToAndCall(implementation, initializer);
            await transaction.wait();
            if (restoreAdmin) await (await direct.changeAdmin(admin)).wait();
            return transaction;
        },
    };
    return facade;
};
