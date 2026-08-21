import "@nomiclabs/hardhat-ethers";

import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";

import { getContractAddressByName } from "../src/deploy-utils";
import { ContractFactoryName, ProxyStorageName } from "../src/types";

export const SubmitterRegister = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    ownerSigner: any,
    submitterSigner: any
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

    const account = await submitterSigner.getAddress();
    if (!(await ownerView.registered(account))) {
        await (await ownerView.addSubmitter(account)).wait();
    }

    const submitterView = ownerView.connect(submitterSigner);
    const minimumStake = await submitterView.minimumStake();
    const currentStake = await submitterView.stakeOf(account);
    if (currentStake.lt(minimumStake)) {
        await (await submitterView.stake({ value: minimumStake.sub(currentStake) })).wait();
    }
    if (!(await submitterView.isActive(account))) return "Submitter did not become active";

    console.log(`Submitter ${account} registered and active with stake ${await submitterView.stakeOf(account)}`);
    return "";
};

export default SubmitterRegister;
