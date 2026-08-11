import { expect } from "chai";
import { ethers } from "hardhat";

const EIP1967_ADMIN_SLOT = "0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103";

describe("Submitter deployment order", () => {
    it("installs and initializes Submitter before Rollup and admin transfer", async () => {
        const [proxyAdminSigner, submitterOwner, other] = await ethers.getSigners();
        const empty = await (await ethers.getContractFactory("EmptyContract")).deploy();
        const proxyFactory = await ethers.getContractFactory("TransparentUpgradeableProxy");
        const rollupProxy = await proxyFactory.deploy(empty.address, proxyAdminSigner.address, "0x");
        const submitterProxy = await proxyFactory.deploy(empty.address, proxyAdminSigner.address, "0x");
        const submitterImpl = await (await ethers.getContractFactory("Submitter")).deploy();
        const rollupImpl = await (await ethers.getContractFactory("Rollup")).deploy(53077);

        const proxyAdminView = await ethers.getContractAt(
            "ITransparentUpgradeableProxy",
            submitterProxy.address,
            proxyAdminSigner
        );
        const submitterFactory = await ethers.getContractFactory("Submitter");
        await proxyAdminView.upgradeToAndCall(
            submitterImpl.address,
            submitterFactory.interface.encodeFunctionData("initialize", [
                submitterOwner.address,
                rollupProxy.address,
                ethers.utils.parseEther("1"),
                ethers.utils.parseEther("1"),
                20,
            ])
        );
        expect(await proxyAdminView.implementation()).to.equal(submitterImpl.address);

        const submitter = await ethers.getContractAt("Submitter", submitterProxy.address, submitterOwner);
        expect(await submitter.owner()).to.equal(submitterOwner.address);
        expect(await submitter.rollupContract()).to.equal(rollupProxy.address);
        await expect(
            submitter.initialize(
                submitterOwner.address,
                rollupProxy.address,
                ethers.utils.parseEther("1"),
                ethers.utils.parseEther("1"),
                20
            )
        ).to.be.revertedWith("Initializable: contract is already initialized");

        const rollupProxyAdminView = await ethers.getContractAt(
            "ITransparentUpgradeableProxy",
            rollupProxy.address,
            proxyAdminSigner
        );
        const rollupFactory = await ethers.getContractFactory("Rollup");
        await rollupProxyAdminView.upgradeToAndCall(
            rollupImpl.address,
            rollupFactory.interface.encodeFunctionData("initialize", [
                submitterProxy.address,
                other.address,
                other.address,
                10,
                100,
                70,
            ])
        );
        const rollup = await ethers.getContractAt("Rollup", rollupProxy.address, submitterOwner);
        expect(await rollup.submitterContract()).to.equal(submitterProxy.address);

        const proxyAdmin = await (await ethers.getContractFactory("ProxyAdmin")).deploy();
        await proxyAdminView.changeAdmin(proxyAdmin.address);
        const rawAdmin = await ethers.provider.getStorageAt(submitterProxy.address, EIP1967_ADMIN_SLOT);
        expect(ethers.utils.getAddress(ethers.utils.hexDataSlice(rawAdmin, 12))).to.equal(proxyAdmin.address);
    });
});
