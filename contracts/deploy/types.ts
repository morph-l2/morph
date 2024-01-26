const ContractFactoryName = {
    // system
    DefaultProxy: 'Proxy',
    ProxyAdmin: 'ProxyAdmin',
    AddressManager: 'AddressManager',
    // messenger
    L1CrossDomainMessenger: 'L1CrossDomainMessenger',
    L1MessageQueue: 'L1MessageQueue',
    // rollup
    L2GasPriceOracle: 'L2GasPriceOracle',
    Rollup: 'Rollup',
    ZkEvmVerifierV1: 'ZkEvmVerifierV1',
    MultipleVersionRollupVerifier: 'MultipleVersionRollupVerifier',
    // staking
    Staking: 'Staking',
    L1Sequencer: 'L1Sequencer',
    // gateway
}

const ProxyStorageName = {
    // messenger
    L1CrossDomainMessengerProxyStroageName: 'Proxy__L1CrossDomainMessenger',
    L1MessageQueueProxyStroageName: 'Proxy__L1MessageQueue',
    // rollup
    L2GasPriceOracleProxyStorageName: 'Proxy__L2GasPriceOracle',
    RollupProxyStorageName: 'Proxy__Rollup',
    MultipleVersionRollupVerifierStorageName: 'Proxy__MultipleVersionRollupVerifier',
    // staking
    StakingProxyStroageName: 'Proxy__Staking',
    L1SequencerProxyStroageName: 'Proxy__L1Sequencer'
    // gateway
}

const ImplStorageName = {
    // system 
    ProxyAdmin: 'Impl__ProxyAdmin',
    AddressManager: 'Impl__AddressManager',
    // messenger
    L1CrossDomainMessengerStorageName: 'Impl__L1CrossDomainMessenger',
    L1MessageQueueStroageName: 'Impl__L1MessageQueue',
    // rollup
    L2GasPriceOracleStorageName: 'Impl__L2GasPriceOracle',
    RollupStorageName: 'Impl__Rollup',
    ZkEvmVerifierV1StorageName: 'Impl__ZkEvmVerifierV1',
    MultipleVersionRollupVerifierStorageName: 'Impl__MultipleVersionRollupVerifier',
    // staking
    StakingStorageName: 'Impl__Staking',
    L1SequencerStorageName: 'Impl__L1Sequencer'
    // gateway
}

export {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName
}