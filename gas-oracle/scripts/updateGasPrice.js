const { ethers } = require('ethers');
const L1GasPriceOracle = require("../resource/GasPriceOracleABI.json");
require("dotenv").config({ path: ".env" });
const express = require('express')
const { register, collectDefaultMetrics, Gauge } = require('prom-client')
const logger = require("./logger")


/**
 * Update l2 gas price.
 * Use this command to execute: node .\scripts\updateGasPrice.js
 */
async function main() {
      
    // Create metric for prometheus
    const gauge_l2 = new Gauge({
        name: 'l1BaseFee_on_l2',
        help: 'l1BaseFee on l2'
    });
    register.registerMetric(gauge_l2);

    const gauge_l1 = new Gauge({
        name: 'l1BaseFee',
        help: 'l1BaseFee'
    });
    register.registerMetric(gauge_l1);


    const gauge_balance = new Gauge({
        name: 'gasOracleOwnerBalance',
        help: 'gasOracleOwnerBalance'
    });
    register.registerMetric(gauge_balance);


    // Update Logic
    const l2Signer = createL2Signer();
    const L2GasPriceOracleContract = new ethers.Contract(
        requireEnv("L2_GAS_PRICE_ORACLE"),
        L1GasPriceOracle.abi,
        l2Signer,
    )

    const l1Provider = createL1Provider();
    setInterval(async function () {
        try {
            //step1. get ethereum mainnet baseFee
            const l1BaseFee = (await l1Provider.getBlock()).baseFeePerGas;
            logger.main.info("current ethereum l1BaseFee is: " + ethers.utils.formatUnits(l1BaseFee, "wei"))

            if (l1BaseFee !== undefined) {
                gauge_l1.set(Number(ethers.utils.formatUnits(l1BaseFee, "gwei")));
            }

            //step2. set baseFee for l2
            const oracleBaseFee = await L2GasPriceOracleContract.l1BaseFee();
            logger.main.info("current l1BaseFee on l2 is: ", ethers.utils.formatUnits(oracleBaseFee, "wei"))

            if (oracleBaseFee !== undefined) {
                gauge_l2.set(Number(ethers.utils.formatUnits(oracleBaseFee, "gwei")));
            }

            const change = Math.abs(oracleBaseFee - l1BaseFee);
            if (change > requireEnv("GAS_THRESHOLD")) {
                await L2GasPriceOracleContract.setL1BaseFee(ethers.utils.formatUnits(l1BaseFee, "wei"));
                logger.main.info("successfully set l1BaseFee to: ", ethers.utils.formatUnits(await L2GasPriceOracleContract.l1BaseFee(), "wei"))
            } else {
                logger.main.info("l1BaseFee change value below threshold, change value = ", change)
            }


            //check balance
            const balance = await l2Signer.getBalance();
            if (balance !== undefined) {
                gauge_balance.set(Number(ethers.utils.formatEther(balance)));
            }
        } catch (e) {
            logger.main.error("update baseFee error", e)
        }
    }, requireEnv("INTERVAL"));





    //Prometheus client
    try {
        const app = express();
        const port = 6060;
        collectDefaultMetrics();

        app.get('/metrics', async (req, res) => {
            res.send(await register.metrics())
        })
        app.listen(port, () => {
            logger.main.info(`gasOracle prometheus client listening on port ${port}`)
        })
    } catch (e) {
        logger.main.error("prometheus client error", e)
    }

}

/**
 * Use priveteKey and rpc create Signer for L2.
 * 
 * @returns 
 */
function createL2Signer() {
    const httpProvider = new ethers.providers.JsonRpcProvider(
        requireEnv("L2_RPC")
    );
    const signer = new ethers.Wallet(requireEnv("L2_GAS_ORACLE_PRIVATE_KEY"), httpProvider);
    return signer;
}

/**
 * Use rpc create Signer for L1.
 * 
 * @returns 
 */
function createL1Provider() {
    const httpProvider = new ethers.providers.JsonRpcProvider(
        requireEnv("L1_RPC")
    );
    return httpProvider;
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