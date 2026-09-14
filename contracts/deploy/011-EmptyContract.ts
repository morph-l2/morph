import { HardhatRuntimeEnvironment } from "hardhat/types";
import { assertContractVariable } from "../src/deploy-utils";
import { deployRecordedContract } from "../src/deployment-state";
import { ContractFactoryName, ImplStorageName } from "../src/types";

export const deployEmptyContract = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any
): Promise<string> => {
    const contract = await deployRecordedContract(hre, path, deployer,
        ImplStorageName.EmptyContract, ContractFactoryName.EmptyContract, []);

    return "";
};

export default deployEmptyContract;
