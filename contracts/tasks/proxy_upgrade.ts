import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";

import fs from "fs";
import { ContractFactoryName } from "../src/types"
import { predeploys } from "../src"

// yarn hardhat upgradeProxy --proxyaddr 0x0165878a594ca255338adfa4d48449f69242eb8f --newimpladdr 0x9a9f2ccfde556a7e9ff0848998aa4a0cfd8863ae --network l1
task("upgradeProxy")
    .addParam("proxyaddr")
    .addParam("newimpladdr")
    .setAction(async (taskArgs, hre) => {
        const ProxyFactory = await hre.ethers.getContractFactory('Proxy')
        const proxy = ProxyFactory.attach(taskArgs.proxyaddr)
        console.log("before upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
        const res = await proxy.upgradeTo(taskArgs.newimpladdr)

        const receipt = await res.wait()
        console.log(`receipt status : ${receipt.status}`)
        console.log("after upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
    });

task("upgradeProxyWithPorxyAdmin")
    .addParam("proxyadminaddr")
    .addParam("proxyaddr")
    .addParam("newimpladdr")
    .setAction(async (taskArgs, hre) => {
        const ProxyAdminFactory = await hre.ethers.getContractAt('TransparentUpgradeableProxy', taskArgs.proxyadminaddr)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.proxyadminaddr)

        const ProxyFactory = await hre.ethers.getContractFactory('Proxy')
        const proxy = ProxyFactory.attach(taskArgs.proxyaddr)
        console.log("before upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))

        const res = await proxyAdmin.upgrade(taskArgs.proxyaddr, taskArgs.newimpladdr)
        const receipt = await res.wait()
        console.log(`receipt status : ${receipt.status}`)
        console.log("after upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
    });

// Proxy__L1StandardBridge??
// Proxy__L1CrossDomainMessenger??

/*
Proxy__MorphPortal

cast call 0xdc64a140aa3e981100a9beca4e685f962f0cf6c9 "GUARDIAN()(address)" --rpc-url http://127.0.0.1:9545
0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

cast call 0xdc64a140aa3e981100a9beca4e685f962f0cf6c9 "SYSTEM_CONFIG()(address)" --rpc-url http://127.0.0.1:9545
0x8A791620dd6260079BF849Dc5567aDC3F2FdC318

cast call 0xdc64a140aa3e981100a9beca4e685f962f0cf6c9 "l1Messenger()(address)" --rpc-url http://127.0.0.1:9545
0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9

cast call 0xdc64a140aa3e981100a9beca4e685f962f0cf6c9 "ROLLUP()(address)" --rpc-url http://127.0.0.1:9545
0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e

yarn hardhat upgradeMorphPortalImpl --guardian 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
--systemconfigproxyaddr 0x8A791620dd6260079BF849Dc5567aDC3F2FdC318 \
--rollupproxyaddr 0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e \
--l1cdmproxyaddr 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 --network l1

cast logs --from-block 0 --to-block `cast bn` --address 0x8a791620dd6260079bf849dc5567adc3f2fdc318 `cast sig-event "AdminChanged(address,address)"` --rpc-url http://127.0.0.1:9545

cast abi-decode --input "AnyFuncName(address,address)" 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000005fbdb2315678afecb367f032d93f642f64180aa3

yarn hardhat upgradeProxyWithPorxyAdmin --proxyadminaddr 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
--proxyaddr 0xdc64a140aa3e981100a9beca4e685f962f0cf6c9 --newimpladdr 0x4c5859f0F772848b2D91F1D83E2Fe57935348029 \
--network l1
*/
task("upgradeMorphPortalImpl")
    .addParam("guardian")
    .addParam("systemconfigproxyaddr")
    .addParam("rollupproxyaddr")
    .addParam("l1cdmproxyaddr")
    .setAction(async (taskArgs, hre) => {
        const guardian = taskArgs.guardian
        const systemconfigproxyaddr = taskArgs.systemconfigproxyaddr
        const rollupproxyaddr = taskArgs.rollupproxyaddr
        const l1cdmproxyaddr = taskArgs.l1cdmproxyaddr

        const MorphPortalFactory = await hre.ethers.getContractFactory("MorphPortal");
        const morphportalNewImpl = await MorphPortalFactory.deploy(
            guardian,
            true,
            systemconfigproxyaddr,
            rollupproxyaddr,
            l1cdmproxyaddr,
        );
        await morphportalNewImpl.deployed()
        console.log("new morphportalNewImpl contract address: ", morphportalNewImpl.address)
    })

/*
Proxy__MorphMintableERC20Factory

cast call 0xa513e6e4b8f2a923d98304ec87f64353c4d5c853 "BRIDGE()(address)" --rpc-url http://127.0.0.1:9545
0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0

yarn hardhat upgradeMorphMintableERC20FactoryImpl --l1stbproxyaddr 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0 --network l1 

cast logs --from-block 0 --to-block `cast bn` --address 0xa513e6e4b8f2a923d98304ec87f64353c4d5c853 `cast sig-event "AdminChanged(address,address)"` --rpc-url http://127.0.0.1:9545

cast abi-decode --input "AnyFuncName(address,address)" 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000005fbdb2315678afecb367f032d93f642f64180aa3

yarn hardhat upgradeProxyWithPorxyAdmin --proxyadminaddr 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
--proxyaddr 0xa513e6e4b8f2a923d98304ec87f64353c4d5c853 --newimpladdr 0x2bdCC0de6bE1f7D2ee689a0342D76F52E8EFABa3 \
--network l1
*/
task("upgradeMorphMintableERC20FactoryImpl")
    .addParam("l1stbproxyaddr")
    .setAction(async (taskArgs, hre) => {
        const l1stbproxyaddr = taskArgs.l1stbproxyaddr

        const MorphMintableERC20FactoryFactory = await hre.ethers.getContractFactory("MorphMintableERC20Factory");
        const MorphMintableERC20FactoryNewImpl = await MorphMintableERC20FactoryFactory.deploy(l1stbproxyaddr);
        await MorphMintableERC20FactoryNewImpl.deployed()
        console.log("new MorphMintableERC20FactoryNewImpl contract address: ", MorphMintableERC20FactoryNewImpl.address)
    })

/*
Proxy__L1ERC721Bridge

cast call 0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6 "MESSENGER()(address)" --rpc-url http://127.0.0.1:9545
0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9

cast call 0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6 "OTHER_BRIDGE()(address)" --rpc-url http://127.0.0.1:9545
0x4200000000000000000000000000000000000014

yarn hardhat upgradeL1ERC721BridgeImpl --l1cdmproxyaddr 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 --otherbridgeaddr 0x4200000000000000000000000000000000000014 --network l1

cast logs --from-block 0 --to-block `cast bn` --address 0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6 `cast sig-event "AdminChanged(address,address)"` --rpc-url http://127.0.0.1:9545

cast abi-decode --input "AnyFuncName(address,address)" 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000005fbdb2315678afecb367f032d93f642f64180aa3

yarn hardhat upgradeProxyWithPorxyAdmin --proxyadminaddr 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
--proxyaddr 0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6 --newimpladdr 0x7bc06c482DEAd17c0e297aFbC32f6e63d3846650 \
--network l1
*/
task("upgradeL1ERC721BridgeImpl")
    .addParam("l1cdmproxyaddr")
    .addParam("otherbridgeaddr")
    .setAction(async (taskArgs, hre) => {
        const l1cdmproxyaddr = taskArgs.l1cdmproxyaddr
        const otherbridgeaddr = taskArgs.otherbridgeaddr

        const L1ERC721BridgeFactory = await hre.ethers.getContractFactory("L1ERC721Bridge");
        const L1ERC721BridgeNewImpl = await L1ERC721BridgeFactory.deploy(l1cdmproxyaddr, otherbridgeaddr);
        await L1ERC721BridgeNewImpl.deployed()
        console.log("new L1ERC721BridgeNewImpl contract address: ", L1ERC721BridgeNewImpl.address)
    })

/*
Proxy__SystemConfig

yarn hardhat upgradeSystemConfigImpl --syscfgproxyaddr 0x8a791620dd6260079bf849dc5567adc3f2fdc318 --network l1

cast logs --from-block 0 --to-block `cast bn` --address 0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6 `cast sig-event "AdminChanged(address,address)"` --rpc-url http://127.0.0.1:9545

cast abi-decode --input "AnyFuncName(address,address)" 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000005fbdb2315678afecb367f032d93f642f64180aa3

yarn hardhat upgradeProxyWithPorxyAdmin --proxyadminaddr 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
--proxyaddr 0x8a791620dd6260079bf849dc5567adc3f2fdc318 --newimpladdr 0x4ed7c70f96b99c776995fb64377f0d4ab3b0e1c1 \
--network l1
*/
task("upgradeSystemConfigImpl")
    .addParam("syscfgproxyaddr")
    .setAction(async (taskArgs, hre) => {

        const syscfgFactory = await hre.ethers.getContractFactory('SystemConfig')
        const syscfg = syscfgFactory.attach(taskArgs.syscfgproxyaddr)

        const owner = await syscfg.connect(hre.waffle.provider).callStatic.owner({
            from: ethers.constants.AddressZero,
        })
        const overhead = await syscfg.connect(hre.waffle.provider).callStatic.overhead({
            from: ethers.constants.AddressZero,
        })
        const scalar = await syscfg.connect(hre.waffle.provider).callStatic.scalar({
            from: ethers.constants.AddressZero,
        })
        const batcherHash = await syscfg.connect(hre.waffle.provider).callStatic.batcherHash({
            from: ethers.constants.AddressZero,
        })
        const gasLimit = await syscfg.connect(hre.waffle.provider).callStatic.gasLimit({
            from: ethers.constants.AddressZero,
        })
        const p2pSequencerAddress = "0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"
        const resourceConfig = await syscfg.connect(hre.waffle.provider).callStatic.resourceConfig({
            from: ethers.constants.AddressZero,
        })

        const SyscfgNewImpl = await syscfgFactory.deploy(owner, overhead, scalar, batcherHash, gasLimit, p2pSequencerAddress, resourceConfig);
        await SyscfgNewImpl.deployed()
        console.log("new SyscfgNewImpl contract address: ", SyscfgNewImpl.address)
    })

/*
Proxy__SystemDictator

yarn hardhat upgradeSystemDictatorImpl --network l1

yarn hardhat upgradeProxy --proxyaddr 0x610178da211fef7d417bc0e6fed39f05609ad788 --newimpladdr 0xFD471836031dc5108809D173A067e8486B9047A3 --network l1
*/
task("upgradeSystemDictatorImpl")
    .setAction(async (taskArgs, hre) => {
        const SystemDictatorFactory = await hre.ethers.getContractFactory("SystemDictator");
        const SystemDictatorNewImpl = await SystemDictatorFactory.deploy();
        await SystemDictatorNewImpl.deployed()
        console.log("new SystemDictatorNewImpl contract address: ", SystemDictatorNewImpl.address)
    })

/*
Proxy__Rollup

cast call 0xb7f8bc63bbcad18155201308c8f3540b07f84f5e "layer2ChainId()(uint64)" --rpc-url http://127.0.0.1:9545
53077

cast call 0xb7f8bc63bbcad18155201308c8f3540b07f84f5e "MESSENGER()(address)" --rpc-url http://127.0.0.1:9545  
0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9

yarn hardhat upgradeRollupImpl --chainid 53077 --l1cdmproxyaddr 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 --network l1 

yarn hardhat upgradeProxy --proxyaddr 0xb7f8bc63bbcad18155201308c8f3540b07f84f5e --newimpladdr 0x1429859428C0aBc9C2C47C8Ee9FBaf82cFA0F20f --network l1
*/
task("upgradeRollupProxy")
    .addParam("proxyadminaddr")
    .addParam("rollupproxyaddr")
    .addParam("chainid")
    .setAction(async (taskArgs, hre) => {
        const ProxyAdminFactory = await hre.ethers.getContractFactory('ProxyAdmin')
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.proxyadminaddr)
        // upgrade
        const chainID = taskArgs.chainid
        const RollupFactory = await hre.ethers.getContractFactory("Rollup");
        const rollupNewImpl = await RollupFactory.deploy(chainID);
        await rollupNewImpl.deployed()
        console.log("new rollupNewImpl contract address: ", rollupNewImpl.address)
        if (!hre.ethers.utils.isAddress(taskArgs.rollupproxyaddr) || !hre.ethers.utils.isAddress(rollupNewImpl.address)) {
            console.log(`not address ${taskArgs.rollupproxyaddr} ${rollupNewImpl.address}`)
            return
        }
        const res = await proxyAdmin.upgrade(taskArgs.rollupproxyaddr, rollupNewImpl.address)
        const rec = await res.wait()
        console.log(`upgrade rollup ${rec.status === 1}`)
    })

/*
Proxy__GasOracle

EnableCurie

yarn hardhat gasOracleEnableCurie --gasoracleproxyaddr 0x530000000000000000000000000000000000000f --network l2
*/
task("gasOracleEnableCurie")
    .addParam("gasoracleproxyaddr")
    .setAction(async (taskArgs, hre) => {
        if (!hre.ethers.utils.isAddress(taskArgs.gasoracleproxyaddr)) {
            console.log(
                `GasOracle proxy address check failed ${taskArgs.gasoracleproxyaddr}`
            )
            return
        }

        let iGasOracle = await hre.ethers.getContractAt("GasPriceOracle", taskArgs.gasoracleproxyaddr)
        let isCurie = await iGasOracle.isCurie()

        if (isCurie) {
            console.log("Already set isCurie")
            return
        }

        // enable curie
        const res = await iGasOracle.enableCurie()
        const rec = await res.wait()
        isCurie = await iGasOracle.isCurie()
        console.log(`Enable curie ${isCurie === true ? "succeed" : "failed"}`)
    })

/*
Proxy__GasOracle

yarn hardhat gasOracleProxy-upgrade --l2config ../ops/l2-genesis/deploy-config/qanet-deploy-config.json  --network qanetl2
*/

task("gasOracleProxy-upgrade")
    .addParam("l2config")
    .setAction(async (taskArgs, hre) => {
        const data = fs.readFileSync(taskArgs.l2config);
        // @ts-ignore
        const l2Config = JSON.parse(data);

        const GasPriceOracleFactory = await hre.ethers.getContractFactory("GasPriceOracle")
        const gasOracleProxyNewImpl = await GasPriceOracleFactory.deploy(l2Config.gasPriceOracleOwner)
        await gasOracleProxyNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`GasPriceOracle new impl deploy at ${gasOracleProxyNewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)
        let res = await proxyAdmin.upgrade(predeploys.GasPriceOracle, gasOracleProxyNewImpl.address)
        let rec = await res.wait()
        console.log(`upgrade gasOracleProxy ${rec.status == 1 ? "success" : "failed"}`)
    })