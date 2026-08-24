import "@nomiclabs/hardhat-ethers";

import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";

import { getContractAddressByName } from "../src/deploy-utils";
import { ContractFactoryName, ProxyStorageName } from "../src/types";

export const SubmitterRegister = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    ownerSigner: any,
    submitter: string
): Promise<string> => {
    const proxyAddress = getContractAddressByName(path, ProxyStorageName.SubmitterProxyStorageName);
    if (!ethers.utils.isAddress(proxyAddress)) return "invalid Submitter proxy address";

    const ownerView = await hre.ethers.getContractAt(
        ContractFactoryName.Submitter,
        proxyAddress,
        ownerSigner
    );
    const ownerAddress = await ownerSigner.getAddress();
    if ((await ownerView.owner()).toLowerCase() !== ownerAddress.toLowerCase()) {
        return "registration signer is not Submitter owner";
    }

    if (!ethers.utils.isAddress(submitter)) return "invalid Submitter address";
    if (!(await ownerView.registered(submitter))) {
        await (await ownerView.addSubmitter(submitter)).wait();
    }

    const minimumStake = await ownerView.minimumStake();
    const currentStake = await ownerView.stakeOf(submitter);
    if (currentStake.lt(minimumStake)) {
        await (await ownerView.stake(submitter, { value: minimumStake.sub(currentStake) })).wait();
    }
    if (!(await ownerView.isActive(submitter))) return "Submitter did not become active";

    console.log(`Submitter ${submitter} registered and active with stake ${await ownerView.stakeOf(submitter)}`);
    return "";
};

export default SubmitterRegister;
