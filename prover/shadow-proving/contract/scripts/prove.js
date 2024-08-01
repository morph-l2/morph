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


    console.log("=========commitBatch");
    await commitBatch(shadow_rollup, customHttpProvider);
    console.log("=========proveState");
    await proveState(shadow_rollup, customHttpProvider);



}

async function proveState(shadow_rollup, customHttpProvider) {

    let proveData = loadProveData();
    let _aggrProof = ethers.utils.arrayify(new Uint8Array(proveData.proof_data));
    let _kzgData = ethers.utils.arrayify(new Uint8Array(proveData.blob_kzg));
    let _pi_data = ethers.utils.arrayify(new Uint8Array(proveData.pi_data));
    console.log("_pi_data: " + _pi_data);

    let tx = await shadow_rollup.proveState(15, _aggrProof, _kzgData);
    await tx.wait();
    console.log("==============================");
    let receipt = await customHttpProvider.getTransactionReceipt(tx.hash);
    console.log("receipt: " + JSON.stringify(receipt));
}


async function commitBatch(shadow_rollup, customHttpProvider) {
    let batchData = {
        prevStateRoot: "0x226cc11d9dc85a98baac8109303ef74163778ca6ea1a5ef1b39f4d79ee49a5b4",
        postStateRoot: "0x226cc11d9dc85a98baac8109303ef74163778ca6ea1a5ef1b39f4d79ee49a5b4",
        withdrawalRoot: "0x3e00490d9ecd0777a41441e5bebaa848fb9e45627c296464f9302a83a908b8d2",
        dataHash: "0x25b849644b691771ff52cf20583681302539d86a61aeac4e7a479d5be01a89d6",
        blobVersionedHash: "0x015b4e3d3dcd64cc0eb6a5ad535d7a1844a8c4cdad366ec73557bcc533941370",
        sequencerSetVerifyHash:"0x5586abb0fb477a7951e0a249f28ee3ab81a9892d89b7bf54a887a924da4e93f0"
    };
    let tx = await shadow_rollup.commitBatch(15, batchData);
    await tx.wait();
    console.log("==============================");
    let receipt = await customHttpProvider.getTransactionReceipt(tx.hash);
    console.log("receipt: " + JSON.stringify(receipt));
}


function loadProveData() {
    const inputBuffer = fs.readFileSync('./testdata/local_proof.json');
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
