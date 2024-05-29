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
        prevStateRoot: "0x1306f43ca4580b9b94bc83e79c23fbcaaff7f261557c7ea90c86c5b4c17584d2",
        postStateRoot: "0x1acb7b53ebd05988c1267d6cdebcc61f871d099cc2e983d1a2c371cb53fdd1fa",
        withdrawalRoot: "0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757",
        dataHash: "0xc8655a6aa50d4385480cc0b65f64f6b060a1b57ab68ee5a572516b946e9cdf70",
        blobVersionedHash: "0x0113416a8338bd500524d63315f3c3f6e9b96d3d5a3ad303451c4f4af31d5993"
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
