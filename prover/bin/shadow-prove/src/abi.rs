use alloy::sol;

// Codegen from ABI file to interact with the contract.
sol!(
    #[sol(rpc)]
    Rollup,
    "abi/Rollup.json"
);

sol!(
    #[sol(rpc)]
    ShadowRollup,
    "abi/ShadowRollup.json"
);
