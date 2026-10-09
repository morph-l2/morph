import { HardhatRuntimeEnvironment } from "hardhat/types";
import { assertContractVariable } from "../src/deploy-utils";
import { deployRecordedContract } from "../src/deployment-state";
import { ContractFactoryName, ImplStorageName } from "../src/types";

export const deployZkEvmVerifierV1 = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, config: any
): Promise<string> => {
    const contract = await deployRecordedContract(hre, path, deployer,
        ImplStorageName.ZkEvmVerifierV1StorageName, ContractFactoryName.ZkEvmVerifierV1, [config.programVkey]);
    await assertContractVariable(contract, "programVkey", config.programVkey);
    return "";
};

export default deployZkEvmVerifierV1;
