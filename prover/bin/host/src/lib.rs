use anyhow::{bail, Context};
pub mod evm;
pub mod execute;
pub mod utils;
use evm::{save_plonk_fixture, EvmProofFixture};
use prover_executor_client::{types::input::ExecutorInput, verify};
use prover_primitives::{alloy_primitives::keccak256, B256};
use prover_utils::read_env_var;
#[cfg(feature = "local")]
use sp1_sdk::CpuProver;
use sp1_sdk::{
    network::FulfillmentStrategy, Elf, HashableKey, ProveRequest, Prover, ProverClient, ProvingKey,
    SP1ProvingKey, SP1Stdin,
};
#[cfg(all(feature = "network", not(feature = "local")))]
use sp1_sdk::{network::NetworkMode, NetworkProver};
use sp1_verifier::PlonkVerifier;
use std::time::{Duration, Instant};

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const BATCH_VERIFIER_ELF: &[u8] = include_bytes!("../../client/elf/verifier-client");

/// The maximum number of blocks that can be included in a single proof.
const MAX_PROVE_BLOCKS: usize = 4096;

/// A batch prover that can generate and verify proofs for a batch of EVM transactions.
pub struct BatchProver<C> {
    prover_client: C,
    pk: SP1ProvingKey,
}

#[cfg(feature = "local")]
pub type DefaultClient = CpuProver;
#[cfg(all(feature = "network", not(feature = "local")))]
pub type DefaultClient = NetworkProver;

// If both features are enabled (e.g. defaults + `--features local`), prefer `local`.
// If neither is enabled, fail fast with a clear message.
#[cfg(all(not(feature = "local"), not(feature = "network")))]
compile_error!("One of `local` or `network` features must be enabled for morph-prove.");

impl BatchProver<DefaultClient> {
    /// Creates a new BatchProver with the default proving client based on the feature flag.
    pub async fn new() -> Result<Self, anyhow::Error> {
        let prover_client = {
            #[cfg(all(feature = "network", not(feature = "local")))]
            {
                ProverClient::builder()
                    .network_for(NetworkMode::Mainnet)
                    .rpc_url(&read_env_var(
                        "SP1_RPC_URL",
                        "https://rpc.mainnet.succinct.xyz".to_owned(),
                    ))
                    .build()
                    .await
            }
            #[cfg(feature = "local")]
            {
                ProverClient::builder().cpu().build().await
            }
        };
        let pk = prover_client
            .setup(Elf::Static(BATCH_VERIFIER_ELF))
            .await
            .context("failed to setup prover")?;
        log::info!("Batch ELF Verification Key: {:?}", pk.verifying_key().bytes32());
        Ok(Self { prover_client, pk })
    }
}

impl BatchProver<DefaultClient> {
    /// Proves a batch of EVM transactions and verifies the proof.
    pub async fn prove(
        &self,
        input: &mut ExecutorInput,
        prove: bool,
    ) -> Result<Option<EvmProofFixture>, anyhow::Error> {
        let program_hash = keccak256(BATCH_VERIFIER_ELF);
        log::info!("Program Hash [view on Explorer]:");
        log::info!("{}", alloy::hex::encode_prefixed(program_hash));

        if input.block_inputs.len() > MAX_PROVE_BLOCKS {
            bail!(
                "check block_traces, blocks len = {} exceeds MAX_PROVE_BLOCKS = {}",
                input.block_inputs.len(),
                MAX_PROVE_BLOCKS
            );
        }

        // Execute in native and prepare input.
        let expected_hash = verify(input.clone()).context("native execution failed")?;
        log::info!(
            "pi_hash generated with native execution: {}",
            alloy::hex::encode_prefixed(expected_hash.as_slice())
        );

        // Execute the program in sp1-vm
        let mut stdin = SP1Stdin::new();
        stdin.write(&serde_json::to_string(&input)?);

        if read_env_var("DEVNET", false) {
            let (mut public_values, execution_report) = self
                .prover_client
                .execute(Elf::Static(BATCH_VERIFIER_ELF), stdin.clone())
                .await
                .context("sp1-vm execution failed")?;
            log::info!(
                "Program executed successfully, Number of cycles: {}",
                execution_report.total_instruction_count()
            );
            let pi_hash = public_values.read::<[u8; 32]>();
            let public_values = B256::from_slice(&pi_hash);

            log::info!(
                "pi_hash generated with sp1-vm execution: {}",
                alloy::hex::encode_prefixed(public_values.as_slice())
            );
            if pi_hash != expected_hash.as_slice() {
                bail!("pi_hash mismatch with expected hash");
            }
            log::info!("Values are correct!");
        }

        if !prove {
            log::info!("Execution completed, No prove request, skipping...");
            return Ok(None);
        }

        // Generate the proof
        log::info!("Start proving...");
        let start = Instant::now();
        let mut proof = self
            .prover_client
            .prove(&self.pk, stdin)
            .strategy(FulfillmentStrategy::Auction)
            .skip_simulation(true)
            .gas_limit(read_env_var("SP1_GAS_LIMIT", 10_000_000_000))
            .plonk()
            .timeout(Duration::from_secs(read_env_var("SP1_TIMEOUT", 1200)))
            .await
            .context("proving failed")?;
        let duration_mins = start.elapsed().as_secs() / 60;
        log::info!("Successfully generated proof!, time use: {} minutes", duration_mins);

        // Verify the proof.
        let plonk_proof = proof.bytes();
        let public_inputs = proof.public_values.to_vec();
        let plonk_vk = *sp1_verifier::PLONK_VK_BYTES;
        let vk = self.pk.verifying_key();
        PlonkVerifier::verify(&plonk_proof, &public_inputs, &vk.bytes32(), plonk_vk)
            .context("failed to verify proof")?;
        log::info!("Successfully verified proof!");

        // Deserialize the public values.
        let pi_bytes = proof.public_values.read::<[u8; 32]>();
        log::info!(
            "pi_hash generated with sp1-vm prove: {}",
            alloy::hex::encode_prefixed(pi_bytes)
        );
        let fixture = EvmProofFixture {
            vkey: vk.bytes32().to_string(),
            public_values: B256::from_slice(&pi_bytes).to_string(),
            proof: alloy::hex::encode_prefixed(proof.bytes()),
        };

        if read_env_var("SAVE_FIXTURE", false) {
            save_plonk_fixture(&fixture);
        }
        Ok(Some(fixture))
    }
}

#[cfg(test)]
mod tests {
    use prover_primitives::B256;

    use sp1_sdk::{Elf, HashableKey, Prover, ProverClient, ProvingKey, SP1ProofWithPublicValues};
    use sp1_verifier::PlonkVerifier;

    use crate::{
        evm::{save_plonk_fixture, EvmProofFixture},
        BATCH_VERIFIER_ELF,
    };

    #[tokio::test]
    pub async fn verify_proof() {
        let proof_loaded = SP1ProofWithPublicValues::load("../../proof/artifact_mainnet").unwrap();

        let proof = proof_loaded.bytes();
        let public_inputs = proof_loaded.public_values.to_vec();

        let client = ProverClient::from_env().await;

        let pk = client.setup(Elf::Static(BATCH_VERIFIER_ELF)).await.unwrap();
        let vk = pk.verifying_key();
        // Print the verification key.
        println!("Program Verification Key: {}", vk.bytes32());
        let plonk_vk = *sp1_verifier::PLONK_VK_BYTES;
        PlonkVerifier::verify(&proof, &public_inputs, &vk.bytes32(), plonk_vk)
            .expect("plonk verify failed");

        let fixture = EvmProofFixture {
            vkey: vk.bytes32().to_string(),
            public_values: B256::from_slice(&public_inputs).to_string(),
            proof: alloy::hex::encode_prefixed(proof),
        };

        save_plonk_fixture(&fixture);
        println!("save_plonk_fixture success!");
    }
}
