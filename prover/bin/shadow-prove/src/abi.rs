use alloy_sol_types::sol;

// Codegen from ABI file to interact with the contract.
sol!(
    #[sol(rpc)]
    Rollup,
    "abi/Rollup.json"
);

// Immutable production ABI from immediately before the Submitter cutover.
// Keep it in a namespace of its own because both ABI snapshots contain an
// `IRollup.BatchDataInput` internal type.
pub mod pre_submitter {
    use alloy_sol_types::sol;

    sol!(PreSubmitterRollup, "abi/pre_submitter/Rollup.json");
}
pub use pre_submitter::PreSubmitterRollup;

sol!(
    #[sol(rpc)]
    ShadowRollup,
    "abi/ShadowRollup.json"
);

sol!(
    #[sol(rpc)]
    SP1Verifier,
    "abi/SP1Verifier.json"
);
