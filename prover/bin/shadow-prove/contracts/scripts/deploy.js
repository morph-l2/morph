const { ethers } = require('ethers');
const { hexlify } = require("ethers/lib/utils");
const fs = require('fs');
const ShadowRollup = require("../bytecode/ShadowRollup.json");
const ZkEvmVerifierV1 = require("../bytecode/ZkEvmVerifierV1.json");


require("dotenv").config({ path: ".env" });

// This is a script for deploying your contracts. You can adapt it to deploy
// yours, or create new ones.
async function main() {

    ///prepare parameter
    let l1_network_name = requireEnv("L1_NETWORK_NAME");
    let privateKey = requireEnv("PRIVATE_KEY");

    let config = loadConfig();
    let layer2ChainId = config[l1_network_name].layer2ChainId;
    console.log("l1_network_name: %s, layer2ChainId: %s", l1_network_name, layer2ChainId);

    ///prepare deployer
    let customHttpProvider = new ethers.providers.JsonRpcProvider(
        requireEnv("L1_RPC")
    );
    const signer = new ethers.Wallet(privateKey, customHttpProvider);
    console.log("signer.address: " + signer.address);
    let balance = await signer.getBalance();
    console.log("signer.balance: " + ethers.utils.formatEther(balance));


    ///deploy plonk_verifier
    const bytecode = hexlify(fs.readFileSync(`./contracts/libs/plonk-verifier/${l1_network_name}/plonk_verifier.bin`));
    const tx = await signer.sendTransaction({ data: bytecode });
    const receipt = await tx.wait();
    console.log("plonk_verifier address:", receipt.contractAddress);


    ///deploy ZkEvmVerifierV1
    const ZkEvmVerifierV1Factory = new ethers.ContractFactory(ZkEvmVerifierV1.abi, ZkEvmVerifierV1.bytecode, signer);
    zkEvmVerifier = await ZkEvmVerifierV1Factory.deploy(receipt.contractAddress);
    await zkEvmVerifier.deployed();
    console.log("zkEvmVerifier address:", zkEvmVerifier.address);


    ///deploy ShadowRollup
    let ShadowRollupFactory = new ethers.ContractFactory(ShadowRollup.abi, ShadowRollup.bytecode, signer);
    const shadowRollup = await ShadowRollupFactory.deploy(layer2ChainId, zkEvmVerifier.address);
    console.log("shadowRollup deploying...");

    await shadowRollup.deployed();
    console.log("shadowRollup address:", shadowRollup.address);
}


function loadConfig() {
    const inputBuffer = fs.readFileSync('./deploy-config.json');
    const inputString = inputBuffer.toString();
    return JSON.parse(inputString);
}

/**
 * Load environment variables 
 * 
 * @param {*} entry 
 * @returns 
 */
function requireEnv(entry) {
    if (process.env[entry]) {
        return process.env[entry]
    } else {
        throw new Error(`${entry} not defined in .env`)
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
