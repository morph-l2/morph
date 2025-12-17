use revm::context::{BlockEnv, CfgEnv, Evm, TxEnv};
use revm::database::State;
use revm::handler::instructions::EthInstructions;
use revm::handler::{EthFrame, EthPrecompiles};
use revm::inspector::NoOpInspector;
use revm::interpreter::interpreter::EthInterpreter;
use revm::primitives::{Address, U256};
use revm::{Context, Database, ExecuteEvm, MainBuilder, MainContext};

pub mod batch;
pub mod error;
/// The Ethereum EVM context type.
pub type EthEvmContext<DB> = Context<BlockEnv, TxEnv, CfgEnv, DB>;

pub struct EthEvm<DB: Database, I, PRECOMPILE = EthPrecompiles> {
    inner: Evm<
        EthEvmContext<DB>,
        I,
        EthInstructions<EthInterpreter, EthEvmContext<DB>>,
        PRECOMPILE,
        EthFrame,
    >,
    inspect: bool,
}

impl<DB: Database> EthEvm<DB, NoOpInspector> {
    /// Create a new [`MorphEvm`] instance.
    pub fn new(db: DB, block_env: BlockEnv, cfg_env: CfgEnv) -> Self {
        let ctx = Context::mainnet()
            .with_db(db)
            .with_block(block_env)
            .with_cfg(cfg_env)
            .with_tx(Default::default())
            .build_mainnet_with_inspector(NoOpInspector);

        Self { inner: ctx, inspect: false }
    }
}

#[test]
fn test_main_context() {
    let mut state = State::builder().build();

    let mut evm = EthEvm::new(state, BlockEnv::default(), CfgEnv::default());
    let tx = TxEnv {
        nonce: 0,
        gas_price: 1_000_000_000,
        gas_limit: 21_000,
        kind: revm::primitives::TxKind::Call(Address::default()),
        value: U256::from(1_000_000u64),
        data: revm::primitives::Bytes::new(),
        ..Default::default()
    };
    let rt = evm.inner.transact_one(tx);

    let state = evm.inner.finalize();

    let db = evm.inner.journaled_state.database.take_bundle();
}

