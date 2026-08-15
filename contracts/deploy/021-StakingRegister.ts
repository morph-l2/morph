import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { ethers } from 'ethers'
import {
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

const two = BigInt(2)
const gwei = BigInt(1e9)
const eth = gwei * gwei

export const StakingRegister = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    signer: any,
    tmKey: string,
    blsKey: string
): Promise<string> => {
    const L1StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StakingProxyStorageName)
    const L1StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)

    const StakingProxyWithSigner = new ethers.Contract(
        L1StakingProxyAddress,
        L1StakingFactory.interface,
        signer,
    )

    // just for devnet register value may set params to config
    const response = await StakingProxyWithSigner.register(tmKey, blsKey, {
        value: eth
    })
    console.log(`Transaction hash (on L1): ${response.hash}`)
    const receipt = await response.wait()
    console.log('Transaction in L1 height', receipt.blockNumber)

    await awaitCondition(
        async () => {
            const sequencerInfo = (await StakingProxyWithSigner.stakers(signer.address))
            if (sequencerInfo && sequencerInfo.tmKey === tmKey) {
                return true
            }
            return false
        },
        30000,
        1000
    )
    return ''
}

export default StakingRegister