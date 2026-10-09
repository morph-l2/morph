import { HardhatRuntimeEnvironment } from "hardhat/types";
import { assertContractVariable } from "../src/deploy-utils";
import { deployRecordedContract } from "../src/deployment-state";
import { ContractFactoryName, ImplStorageName } from "../src/types";

export const deployProxyAdmin = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any
): Promise<string> => {
    const contract = await deployRecordedContract(hre, path, deployer,
        ImplStorageName.ProxyAdmin, ContractFactoryName.ProxyAdmin, []);
    await assertContractVariable(contract, "owner", await deployer.getAddress());
    return "";
};

export default deployProxyAdmin;
