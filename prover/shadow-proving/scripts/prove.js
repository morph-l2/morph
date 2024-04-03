const { ethers } = require('ethers');
const { hexlify } = require("ethers/lib/utils");
const fs = require('fs');
const ZkEvmVerifierV1 = require("../bytecode/ZkEvmVerifierV1.json");
const ShadowRollup = require("../bytecode/ShadowRollup.json");


require("dotenv").config({ path: ".env" });

// This is a script for deploying your contracts. You can adapt it to deploy
// yours, or create new ones.
async function main() {

    ///prepare parameter
    let rollup_address = requireEnv("SHADOW_ROLLUP_ADDRESS");
    let privateKey = requireEnv("PRIVATE_KEY");

    ///prepare deployer
    let customHttpProvider = new ethers.providers.JsonRpcProvider(
        requireEnv("L1_RPC")
    );
    const signer = new ethers.Wallet(privateKey, customHttpProvider);
    console.log("signer.address: " + signer.address);

    let shadow_rollup = new ethers.Contract(rollup_address, ShadowRollup.abi, signer);


    await commitBatch(shadow_rollup, customHttpProvider);
    await proveState(shadow_rollup, customHttpProvider);



}

async function proveState(shadow_rollup, customHttpProvider) {

    let proveData = loadProveData();
    let _aggrProof = ethers.utils.arrayify(new Uint8Array(proveData.proof_data));
    let _kzgData = ethers.utils.arrayify(new Uint8Array(proveData.blob_kzg));

    let tx = await shadow_rollup.proveState(15, _aggrProof, _kzgData);
    await tx.wait();
    console.log("==============================");
    let receipt = await customHttpProvider.getTransactionReceipt(tx.hash);
    console.log("receipt: " + JSON.stringify(receipt));
}


async function commitBatch(shadow_rollup, customHttpProvider) {
    let batchData = {
        prevStateRoot: "0x000e99ef296bcca960ab82643bfb8798fe0e3fdd2cfdf63f36149ad21316ad21",
        postStateRoot: "0x0c331309ce13ebc35b680a146d02b05ccdaec2e4faedddf86c512ec271a1bb5e",
        withdrawalRoot: "0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757",
        dataHash: "0x85c4206f1433be4d12d2410ffecd6831e09439e52439e3b3f9ef7e0c26d160c7",
        blobVersionedHash: "0x01ece0bb19bccf011f86762399ee264fc44c0ec388553c543fcfc6c38d135f5b"
    };
    let tx = await shadow_rollup.commitBatch(15, batchData);
    await tx.wait();
    console.log("==============================");
    let receipt = await customHttpProvider.getTransactionReceipt(tx.hash);
    console.log("receipt: " + JSON.stringify(receipt));
}


function loadProveData() {
    const inputBuffer = fs.readFileSync('./prove.json');
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
