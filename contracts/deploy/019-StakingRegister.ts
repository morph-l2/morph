import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { ethers } from 'ethers'
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

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
    const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStroageName)
    const StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.Staking)

    const StakingProxyWithSigner = new ethers.Contract(
        StakingProxyAddress,
        StakingFactory.interface,
        signer,
    )

    const response = await StakingProxyWithSigner.register(tmKey, blsKey, 5000000, {
        gasLimit: 10000000,
        value: two * eth
    })
    console.log(`Transaction hash (on L1): ${response.hash}`)
    const receipt = await response.wait()
    console.log('Transaction in L1 height', receipt.blockNumber)

    await awaitCondition(
        async () => {
            const sequencerInfo = (await StakingProxyWithSigner.stakings(signer.address))
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