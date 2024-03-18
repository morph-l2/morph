const ContractFactoryName = {
    // system
    DefaultProxy: 'TransparentUpgradeableProxy',
    DefaultProxyInterface: 'ITransparentUpgradeableProxy',
    ProxyAdmin: 'ProxyAdmin',
    // empty contract
    EmptyContract: 'EmptyContract',
    // tokens
    WETH:'WrappedEther',
    // messenger
    L1CrossDomainMessenger: 'L1CrossDomainMessenger',
    L1MessageQueueWithGasPriceOracle: 'L1MessageQueueWithGasPriceOracle',
    // rollup
    Rollup: 'Rollup',
    ZkEvmVerifierV1: 'ZkEvmVerifierV1',
    MultipleVersionRollupVerifier: 'MultipleVersionRollupVerifier',
    // staking
    Staking: 'Staking',
    L1Sequencer: 'L1Sequencer',
    // gateway
    L1GatewayRouter: 'L1GatewayRouter',
    L1StandardERC20Gateway: 'L1StandardERC20Gateway',
    L1WETHGateway: 'L1WETHGateway',
    L1ETHGateway: 'L1ETHGateway',
    L1ERC721Gateway: 'L1ERC721Gateway',
    L1ERC1155Gateway: 'L1ERC1155Gateway',
    EnforcedTxGateway: "EnforcedTxGateway"
}

const ProxyStorageName = {
    // messenger
    L1CrossDomainMessengerProxyStorageName: 'Proxy__L1CrossDomainMessenger',
    L1MessageQueueWithGasPriceOracleProxyStorageName: 'Proxy__L1MessageQueueWithGasPriceOracle',
    // rollup
    RollupProxyStorageName: 'Proxy__Rollup',
    // staking
    StakingProxyStorageName: 'Proxy__Staking',
    L1SequencerProxyStorageName: 'Proxy__L1Sequencer',
    // gateway
    L1GatewayRouterProxyStorageName: 'Proxy__L1GatewayRouter',
    L1StandardERC20GatewayProxyStorageName: 'Proxy__L1StandardERC20Gateway',
    L1ETHGatewayProxyStorageName: 'Proxy__L1ETHGateway',
    L1WETHGatewayProxyStorageName: 'Proxy__L1WETHGateway',
    L1ERC721GatewayProxyStorageName: 'Proxy__L1ERC721Gateway',
    L1ERC1155GatewayProxyStorageName: 'Proxy__L1ERC1155Gateway',
    EnforcedTxGatewayProxyStorageName: "Proxy__EnforcedTxGateway"
}

const ImplStorageName = {
    // system 
    ProxyAdmin: 'Impl__ProxyAdmin',
    // empty contract
    EmptyContract: 'Impl__EmptyContract',
    // tokens
    WETH:'Impl__WETH',
    // messenger
    L1CrossDomainMessengerStorageName: 'Impl__L1CrossDomainMessenger',
    L1MessageQueueWithGasPriceOracle: 'Impl__L1MessageQueueWithGasPriceOracle',
    // rollup
    RollupStorageName: 'Impl__Rollup',
    ZkEvmVerifierV1StorageName: 'Impl__ZkEvmVerifierV1',
    MultipleVersionRollupVerifierStorageName: 'Impl__MultipleVersionRollupVerifier',
    // staking
    StakingStorageName: 'Impl__Staking',
    L1SequencerStorageName: 'Impl__L1Sequencer',
    // gateway
    L1GatewayRouterStorageName: 'Impl__L1GatewayRouter',
    L1StandardERC20GatewayStorageName: 'Impl__L1StandardERC20Gateway',
    L1ETHGatewayStorageName: 'Impl__L1ETHGateway',
    L1WETHGatewayStorageName: 'Impl__L1WETHGateway',
    L1ERC721GatewayStorageName: 'Impl__L1ERC721Gateway',
    L1ERC1155GatewayStorageName: 'Impl__L1ERC1155Gateway',
    EnforcedTxGatewayStorageName: "Impl__EnforcedTxGateway"
}

export {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName
}