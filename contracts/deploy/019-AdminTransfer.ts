import { HardhatRuntimeEnvironment } from "hardhat/types";
import { getContractAddressByName } from "../src/deploy-utils";
import { getDeploymentProxy } from "../src/deployment-state";
import { ImplStorageName, ProxyStorageName } from "../src/types";

export const AdminTransferByProxyStorageName = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, storageName: string
): Promise<string> => {
    const empty = getContractAddressByName(path, ImplStorageName.EmptyContract);
    const manager = getContractAddressByName(path, ImplStorageName.ProxyAdmin);
    const address = getContractAddressByName(path, storageName);
    const proxy = await getDeploymentProxy(hre, path, address, deployer);
    if (storageName !== ProxyStorageName.L1USDCGatewayProxyStorageName &&
        (await proxy.implementation()).toLowerCase() === empty.toLowerCase()) {
        throw new Error(`${storageName} has no implementation configured; proxy admin transfer stopped`);
    }
    const admin = await proxy.admin();
    if (admin.toLowerCase() === manager.toLowerCase()) return "";
    if (admin.toLowerCase() !== (await deployer.getAddress()).toLowerCase()) {
        throw new Error(`${storageName} admin ${admin} is neither the deployer nor the recorded ProxyAdmin`);
    }
    await (await proxy.changeAdmin(manager)).wait();
    if ((await proxy.admin()).toLowerCase() !== manager.toLowerCase()) {
        throw new Error(`${storageName} proxy admin transfer did not take effect`);
    }
    return "";
};

export const AdminTransfer = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, config: any
): Promise<string> => {
    for (const name of Object.values(ProxyStorageName)) {
        await AdminTransferByProxyStorageName(hre, path, deployer, name);
    }
    return "";
};

// Keep existing task arguments and confirm each admin transfer so retries can read the resulting state.
export const AdminTransferConcurrently = AdminTransfer;
