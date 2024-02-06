const { ethers } = require('ethers');
const L1GasPriceOracle = require("../resource/GasPriceOracleABI.json");
const fs = require("fs");
require("dotenv").config({ path: ".env" });


/**
 * Update l2 gas price.
 * Use this command to execute: node .\scripts\authManage.js
 */
async function main() {
    const L2GasPriceOracleContract = new ethers.Contract(
        requireEnv("L2_GAS_PRICE_ORACLE"),
        L1GasPriceOracle.abi,
        createL2Signer(),
    )

    if (await L2GasPriceOracleContract.allowListEnabled == false) {
        // await L2GasPriceOracleContract.setAllowListEnabled(true);
    }

    let allowList = JSON.parse(fs.readFileSync("./resource/allowList.json"));
    console.log("allowList: " + allowList.user);
    console.log("val: " + allowList.val);
    // let result = await L2GasPriceOracleContract.setAllowList(allowList.user, allowList.val);
    // console.log("result.hash:" + result.hash);

}

/**
 * Use priveteKey and rpc create Signer for L2.
 * 
 * @returns 
 */
function createL2Signer() {
    let httpProvider = new ethers.providers.JsonRpcProvider(
        requireEnv("L2_RPC")
    );
    const signer = new ethers.Wallet(requireEnv("L2_GAS_ORACLE_PRIVATE_KEY"), httpProvider);
    return signer;
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

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});