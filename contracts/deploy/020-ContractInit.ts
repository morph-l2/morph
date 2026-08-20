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
    ImplStorageName,
} from "../src/types"

export const ContractInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log("ContractInit")
    // ------------------ gasPriceOracle init -----------------
    {
        const GasPriceOracleProxyAddress = getContractAddressByName(path, ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName)
        const GasPriceOracle = await hre.ethers.getContractAt(ContractFactoryName.L1MessageQueueWithGasPriceOracle, GasPriceOracleProxyAddress, deployer)
        // base fee
        const baseFeeStr = (config.l2BaseFee).toString()
        let res = await GasPriceOracle.setL2BaseFee(ethers.utils.parseUnits(baseFeeStr, "gwei"))
        let rec = await res.wait()
        console.log(`set base fee ${rec.status === 1} setL2BaseFee(${await GasPriceOracle.l2BaseFee()}) gwei`)

    }

    // ------------------ rollup init -----------------
    {
        const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
        const Rollup = await hre.ethers.getContractAt(ContractFactoryName.Rollup, RollupProxyAddress, deployer)
        // import genesis batch 
        const batchHeader: string = config.batchHeader

        // challenger
        const challenger: string = config.rollupChallenger
        const rollupDelayPeriod: number = config.rollupDelayPeriod

        if (!ethers.utils.isAddress(challenger)) {
            console.error('please check your address')
            return ''
        }
        if (rollupDelayPeriod==0){
            console.error('rollupDelayPeriod cannot set zero')
            return ''
        }
        let res = await Rollup.importGenesisBatch(batchHeader)
        let rec = await res.wait()
        console.log(`importGenesisBatch(%s) ${rec.status == 1 ? "success" : "failed"}`, batchHeader)
        res = await Rollup.addChallenger(challenger)
        rec = await res.wait()
        console.log(`addChallenger(%s) ${rec.status == 1 ? "success" : "failed"}`, challenger)

        res =await Rollup.initialize2("0x0000000000000000000000000000000000000000000000000000000000000001")
        rec = await res.wait()
        console.log(`initialize2(%s) ${rec.status == 1 ? "success" : "failed"}`)

        res = await Rollup.initialize3(rollupDelayPeriod)
        rec = await res.wait()
        console.log(`initialize3(%s) ${rec.status == 1 ? "success" : "failed"}`)
    }

    // ------------------ router init -----------------
    {
        const L1WETHAddress = getContractAddressByName(path, ImplStorageName.WETH)
        const L1WETHGatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1WETHGatewayProxyStorageName)

        const L1GatewayRouterProxyAddress = getContractAddressByName(path, ProxyStorageName.L1GatewayRouterProxyStorageName)
        const l1GatewayRouter = await hre.ethers.getContractAt(ContractFactoryName.L1GatewayRouter, L1GatewayRouterProxyAddress, deployer)

        // set token gateway
        const tokens = [L1WETHAddress]
        const gateways = [L1WETHGatewayProxyAddress]
        await l1GatewayRouter.setERC20Gateway(tokens, gateways)
        await awaitCondition(
            async () => {
                return (
                    (await l1GatewayRouter.getERC20Gateway(L1WETHAddress)).toLocaleLowerCase() === L1WETHGatewayProxyAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        console.log(`router set token gateway success`)

    }
    return ''
}

export default ContractInit
