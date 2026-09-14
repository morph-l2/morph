import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "dotenv/config";
import { task } from "hardhat/config";
import { ethers } from "ethers";
import { deployProxyAdmin, deployEmptyContract, deployZkEvmVerifierV1, deployContractProxies,
    deployContractImpls, SubmitterInit, MessengerInit, RollupInit, GatewayInit, SequencerInit,
    AdminTransfer, SetFirstSequencer, ContractInit, SubmitterRegister } from "../deploy";
import { ensureDeploymentStorageWritable, ensureNoPendingTransactions, validateDeploymentRecords, validateDeploymentSigner } from "../src/deployment-state";
import { validateDeploymentConfig, registrationAddresses } from "../src/deployment-validation";
import { requireDeploymentAddresses, validateVerifierVersions, verifyDeployment, verifyRecordedConfiguration } from "../src/deployment-verification";
import { getContractAddressByName } from "../src/deploy-utils";
import { ContractFactoryName, ImplStorageName, ProxyStorageName } from "../src/types";

const step = async (name: string, action: () => Promise<string>) => {
    console.log(`\n${name}`);
    const error = await action();
    if (error) throw new Error(`${name}: ${error}`);
};

const concurrentCompatibility = (value: string) => {
    if (value !== "false" && value !== "true") throw new Error("concurrent must be true or false");
    if (value === "true") console.log("The concurrent option is accepted; transactions are confirmed and recorded sequentially to support retries");
};

const deploymentPrerequisites = [ImplStorageName.ProxyAdmin, ImplStorageName.EmptyContract,
    ImplStorageName.WETH, ImplStorageName.ZkEvmVerifierV1StorageName, ...Object.values(ProxyStorageName)];

task("deploy")
    .addParam("storagepath")
    .addOptionalParam("concurrent", "Compatibility option; deployment confirms and records transactions sequentially", "false")
    .setAction(async ({ storagepath, concurrent }, hre) => {
        concurrentCompatibility(concurrent);
        const config = hre.deployConfig;
        await validateDeploymentConfig(hre, config, "deploy");
        ensureDeploymentStorageWritable(storagepath);
        await validateDeploymentRecords(hre, storagepath);
        await validateVerifierVersions(hre, storagepath);
        await verifyRecordedConfiguration(hre, storagepath, config);
        const signer = await hre.ethers.provider.getSigner();
        await validateDeploymentSigner(hre, storagepath, signer);
        await step("Deploy ProxyAdmin", () => deployProxyAdmin(hre, storagepath, signer));
        await step("Deploy EmptyContract", () => deployEmptyContract(hre, storagepath, signer));
        await step("Deploy contract proxies", () => deployContractProxies(hre, storagepath, signer, config));
        await step("Deploy proof verifier", () => deployZkEvmVerifierV1(hre, storagepath, signer, config));
    });

task("initialize")
    .addParam("storagepath")
    .addOptionalParam("concurrent", "Compatibility option; deployment confirms and records transactions sequentially", "false")
    .setAction(async ({ storagepath, concurrent }, hre) => {
        concurrentCompatibility(concurrent);
        const config = hre.deployConfig;
        await validateDeploymentConfig(hre, config, "initialize");
        requireDeploymentAddresses(storagepath, deploymentPrerequisites);
        ensureDeploymentStorageWritable(storagepath);
        await validateDeploymentRecords(hre, storagepath);
        await validateVerifierVersions(hre, storagepath);
        await verifyRecordedConfiguration(hre, storagepath, config);
        const signer = await hre.ethers.provider.getSigner();
        await validateDeploymentSigner(hre, storagepath, signer);
        await step("Deploy contract implementations", () => deployContractImpls(hre, storagepath, signer, config));
        await step("Initialize Submitter", () => SubmitterInit(hre, storagepath, signer, config));
        await step("Initialize messenger", () => MessengerInit(hre, storagepath, signer, config));
        await step("Initialize Rollup", () => RollupInit(hre, storagepath, signer, config));
        await step("Initialize gateways", () => GatewayInit(hre, storagepath, signer, config));
        await step("Initialize L1Sequencer", () => SequencerInit(hre, storagepath, signer, config));
        await step("Transfer proxy administration", () => AdminTransfer(hre, storagepath, signer, config));
        await step("Set first sequencer", () => SetFirstSequencer(hre, storagepath, signer, config));
        await step("Set genesis batch and runtime parameters", () => ContractInit(hre, storagepath, signer, config));
        await verifyDeployment(hre, storagepath, config, false);
    });

task("register")
    .addParam("storagepath")
    .setAction(async ({ storagepath }, hre) => {
        const config = hre.deployConfig;
        await validateDeploymentConfig(hre, config, "register");
        const addresses = registrationAddresses(config);
        requireDeploymentAddresses(storagepath, [ProxyStorageName.SubmitterProxyStorageName]);
        await validateDeploymentRecords(hre, storagepath);
        const owner = await hre.ethers.provider.getSigner();
        await ensureNoPendingTransactions(hre, owner);
        const submitter = await hre.ethers.getContractAt(ContractFactoryName.Submitter,
            getContractAddressByName(storagepath, ProxyStorageName.SubmitterProxyStorageName), owner);
        if ((await submitter.owner()).toLowerCase() !== (await owner.getAddress()).toLowerCase()) {
            throw new Error("registration signer is not Submitter owner");
        }
        for (const address of addresses) {
            await step(`Register submitter ${address}`, () => SubmitterRegister(hre, storagepath, owner, address));
        }
    });

task("verify-deployment")
    .addParam("storagepath")
    .addFlag("runtime", "Verify a running network, allowing the fee service to update l2BaseFee")
    .setAction(async ({ storagepath, runtime }, hre) => {
        await verifyDeployment(hre, storagepath, hre.deployConfig, true, runtime);
    });

task("fund")
    .setAction(async (taskArgs, hre) => {
        console.log('\n---------------------------------- Fund Submitters ----------------------------------')
        const signer = await hre.ethers.getSigners()
        const batchSubmitterPkList: string[] = JSON.parse(process.env.batchSubmitterPks || "[]");
        for (let i = 0; i < batchSubmitterPkList.length; i++) {
            const submitter = new ethers.Wallet(batchSubmitterPkList[i], hre.ethers.provider)
            const tx = {
                to: submitter.address,
                value: ethers.utils.parseEther("100")
            }
            let balance = (await submitter.getBalance()).toString()

            if (balance.length < 20) {
                let receipt = await signer[0].sendTransaction(tx)
                await receipt.wait()
            }
            balance = (await submitter.getBalance()).toString()
            console.log(`${submitter.address} has balance: ${balance}`)
        }
    })
