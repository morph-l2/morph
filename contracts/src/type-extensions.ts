import 'hardhat/types/runtime'
import 'hardhat/types/config'

declare module 'hardhat/types/runtime' {
    interface HardhatRuntimeEnvironment {
        deployConfig: {
            [key: string]: any
        }
        getDeployConfig(network: string): { [key: string]: any }
    }
}