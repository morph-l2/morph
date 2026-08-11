# Archived production ABIs

These ABI-only artifacts are immutable compatibility inputs for the centralized
submitter cutover. They were frozen from the pre-Submitter source tree before
the current Rollup ABI was regenerated.

- `pre-submitter/Rollup.json` decodes the three production submission methods
  that still include `BatchSignatureInput`.
- `legacy-l1-staking/IL1Staking.json` preserves the retired public protocol
  interface, while `legacy-l1-staking/L1Staking.json` is the complete runtime
  interface used by the retirement tooling after the Solidity source is removed.

Do not regenerate or overwrite these files from the current contracts.
