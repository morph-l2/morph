require("@nomicfoundation/hardhat-toolbox");

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
      console.log(account.address);
  }
});


task("updateL1BaseFee")
  .addParam("addr")
  .setAction(async (taskArgs, hre) => {
    console.log(taskArgs)
    const [owner] = await hre.ethers.getSigners();
    const L1GasPriceOracle = await hre.ethers.getContractFactory("L1GasPriceOracle")
    const contract = await L1GasPriceOracle.attach(taskArgs.addr)
    await contract.connect(owner).setL1BaseFee(2000000)

    console.log("successfully set l1BaseFee to: ", await contract.l1BaseFee())
})

task("getL1BaseFee")
  .addParam("addr")
  .setAction(async (taskArgs, hre) => {
    console.log(taskArgs)
    const L1GasPriceOracle = await hre.ethers.getContractFactory("L1GasPriceOracle")
    const contract = await L1GasPriceOracle.attach(taskArgs.addr)
    console.log("l1BaseFee is: ", await contract.l1BaseFee())
})






/** @type import('hardhat/config').HardhatUserConfig */

module.exports = {
  solidity: "0.8.24",
  networks: {
    local: {
      url: `http://localhost:8545`,
      accounts: ['ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80']
    },
    devnet: {
      url: `http://localhost:8547`,
      accounts: ['ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80']
    }
  }
};
