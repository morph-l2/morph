import 'hardhat/types/runtime'
import 'hardhat/types/config'


export type DeployConfigSpec<
    TDeployConfig extends {
        [key: string]: any
    }
> = {
        [K in keyof TDeployConfig]: {
            type: 'address' | 'number' | 'string' | 'boolean'
            default?: any
        }
    }

declare module 'hardhat/types/config' {
    interface HardhatUserConfig {
        deployConfigSpec?: DeployConfigSpec<any>
    }

    interface HardhatConfig {
        deployConfigSpec?: DeployConfigSpec<any>
    }

    interface ProjectPathsUserConfig {
        deployConfig?: string
    }

    interface ProjectPathsConfig {
        deployConfig?: string
    }
}

declare module 'hardhat/types/runtime' {
    interface HardhatRuntimeEnvironment {
        deployConfig: {
            [key: string]: any
        }
        getDeployConfig(network: string): { [key: string]: any }
    }
}