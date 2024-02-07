const { expect } = require("chai");
const hre = require("hardhat");

const key = '0x3915789465ce806e8b40bb6409e42771e862456e3029ba9a9110e3f5b4aa6abf'

describe("L1GasPriceOracle", function () {
  let contract;
  let deployer;
  let nonOwner;
  beforeEach(async function () {
    const [signer, signer2] = await hre.ethers.getSigners();
    deployer = signer

    const L1GasPriceOracle = await hre.ethers.getContractFactory("L1GasPriceOracle");
    const oracle = await L1GasPriceOracle.deploy(deployer.address)
    contract = await oracle.deployed();

    nonOwner = signer2

  });  
  it("L1GasPriceOracle should be correctly initilized", async function(){
    expect(await contract.owner()).to.equal(deployer.address);
    expect(await contract.allowListEnabled()).to.equal(true)
    expect(await contract.gasPrice()).to.equal(0)
    expect(await contract.l1BaseFee()).to.equal(0)
    expect(await contract.overhead()).to.equal(0)
    expect(await contract.scalar()).to.equal(0)
  })

  it("Only owner can setAllowListEnabled", async function () {
    await contract.connect(deployer).setAllowListEnabled(false)
    expect(await contract.allowListEnabled()).to.equal(false)
    await expect(contract.connect(nonOwner).setAllowListEnabled(true)).to.be.revertedWith(
        'Ownable: caller is not the owner'
    );
  })

  it("Test onlyAllowed", async function(){
    await contract.setL1BaseFee(200000)
    expect(await contract.l1BaseFee()).to.equal(200000)

    await contract.setOverhead(2100)
    expect(await contract.overhead()).to.equal(2100)

    await contract.setScalar(1000000)
    expect(await contract.scalar()).to.equal(1000000)

    await expect(contract.connect(nonOwner).setL1BaseFee(1)).to.be.revertedWith(
        "not allowed"
    );
    await expect(contract.connect(nonOwner).setOverhead(1)).to.be.revertedWith(
        "not allowed"
    );
    await expect(contract.connect(nonOwner).setScalar(1)).to.be.revertedWith(
        "not allowed"
    );

    await contract.setAllowList([nonOwner.address], [true])
    await contract.connect(nonOwner).setL1BaseFee(300000)
    expect(await contract.l1BaseFee()).to.equal(300000)

    await contract.connect(nonOwner).setOverhead(2200)
    expect(await contract.overhead()).to.equal(2200)

    await contract.connect(nonOwner).setScalar(10000000)
    expect(await contract.scalar()).to.equal(10000000)

    await contract.setAllowList([nonOwner.address], [false])
    await expect(contract.connect(nonOwner).setScalar(1)).to.be.revertedWith(
        "not allowed"
    );

  })
})