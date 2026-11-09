use alloy::primitives::keccak256;
use alloy_primitives::B256;
use anyhow::{Context, bail};
pub mod evm;
pub mod execute;
pub mod utils;
#[cfg(all(feature = "network", not(feature = "local")))]
use std::time::Duration;
use std::time::Instant;

use evm::{EvmProofFixture, save_plonk_fixture};
use prover_executor_client::{types::input::ExecutorInput, verify};
use prover_utils::read_env_var;
#[cfg(feature = "local")]
use sp1_sdk::CpuProver;
use sp1_sdk::{
    Elf, HashableKey, ProveRequest, Prover, ProverClient, ProvingKey, SP1ProofWithPublicValues,
    SP1ProvingKey, SP1Stdin,
};
#[cfg(all(feature = "network", not(feature = "local")))]
use sp1_sdk::{NetworkProver, network::FulfillmentStrategy, network::NetworkMode};
use sp1_verifier::PlonkVerifier;

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
        let mut proof = self.generate_proof(stdin).await?;

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

    /// Requests a Plonk proof from the prover network, retrying failed attempts.
    ///
    /// Each attempt re-fetches auction params (base fee, max price per PGU) from the network, so
    /// transient rejections such as a stale base fee can succeed on retry. Unexecutable requests
    /// are not retried; after an auction timeout (no prover bid) the price buffer is raised.
    #[cfg(all(feature = "network", not(feature = "local")))]
    async fn generate_proof(
        &self,
        stdin: SP1Stdin,
    ) -> Result<SP1ProofWithPublicValues, anyhow::Error> {
        let max_retries: u32 = read_env_var("SP1_PROVE_RETRIES", 1);
        let auction_timeout = Duration::from_secs(read_env_var("SP1_AUCTION_TIMEOUT", 30));
        // Percentage applied to the network's default max_price_per_pgu. Raised after each
        // auction timeout (no prover bid) up to `price_buffer_max`.
        let mut price_buffer: u64 = read_env_var("SP1_PRICE_BUFFER", 120);
        let price_buffer_step: u64 = read_env_var("SP1_PRICE_BUFFER_STEP", 20);
        let price_buffer_max: u64 = read_env_var("SP1_PRICE_BUFFER_MAX", 200);
        let mut attempt = 0;
        loop {
            let rt = self
                .prover_client
                .prove(&self.pk, stdin.clone())
                .strategy(FulfillmentStrategy::Auction)
                .skip_simulation(true)
                .gas_limit(read_env_var("SP1_GAS_LIMIT", 10_000_000_000))
                .auction_timeout(auction_timeout)
                .max_price_per_pgu_buffer(price_buffer)
                .plonk()
                .timeout(Duration::from_secs(read_env_var("SP1_TIMEOUT", 1200)))
                .await;
            let e = match rt {
                Ok(proof) => return Ok(proof),
                Err(e) => e,
            };
            let kind = NetworkFailure::classify(&e);
            if !kind.retryable() || attempt >= max_retries {
                log::error!(
                    "network prove failed [{kind}], giving up after {} attempt(s): {e:#}",
                    attempt + 1
                );
                return Err(e).context(format!(
                    "network prove error [{kind}] after {} attempt(s)",
                    attempt + 1
                ));
            }
            attempt += 1;
            if kind == NetworkFailure::NoBid {
                price_buffer = (price_buffer + price_buffer_step).min(price_buffer_max);
            }
            log::warn!(
                "network prove failed [{kind}] (retry {attempt}/{max_retries}, \
                 price buffer {price_buffer}%): {e:#}"
            );
            tokio::time::sleep(Duration::from_secs(15 * u64::from(attempt))).await;
        }
    }

    /// Generates a Plonk proof with the local CPU prover.
    #[cfg(feature = "local")]
    async fn generate_proof(
        &self,
        stdin: SP1Stdin,
    ) -> Result<SP1ProofWithPublicValues, anyhow::Error> {
        self.prover_client.prove(&self.pk, stdin).plonk().await.context("local prove error")
    }
}

/// Coarse classification of a failed network proving attempt, used for retry decisions and logs.
#[cfg(all(feature = "network", not(feature = "local")))]
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
enum NetworkFailure {
    /// No prover bid on the request within the auction timeout (`Requested` state).
    NoBid,
    /// A prover was assigned but the proof was not delivered before the deadline.
    ProvingTimeout,
    /// The program cannot be executed with the given input; retrying will not help.
    Unexecutable,
    /// The assigned prover gave up on the request.
    Unfulfillable,
    /// The request failed during on-chain settlement.
    Reverted,
    /// The request reached its deadline on the network without a proof.
    Expired,
    /// The RPC server rejected or failed a call (e.g. base fee not accepted).
    Rpc,
    /// Anything else (simulation, transport, local errors).
    Other,
}

#[cfg(all(feature = "network", not(feature = "local")))]
impl NetworkFailure {
    fn classify(e: &anyhow::Error) -> Self {
        use sp1_sdk::network::Error as NetErr;
        for cause in e.chain() {
            if let Some(err) = cause.downcast_ref::<NetErr>() {
                return match err {
                    NetErr::RequestAuctionTimedOut { .. } => Self::NoBid,
                    NetErr::RequestTimedOut { .. } => Self::ProvingTimeout,
                    NetErr::RequestUnexecutable { .. } => Self::Unexecutable,
                    NetErr::RequestUnfulfillable { .. } => Self::Unfulfillable,
                    NetErr::RequestReverted { .. } => Self::Reverted,
                    NetErr::RequestExpired { .. } => Self::Expired,
                    NetErr::RpcError(_) => Self::Rpc,
                    _ => Self::Other,
                };
            }
            if cause.is::<tonic::Status>() {
                return Self::Rpc;
            }
        }
        Self::Other
    }

    fn retryable(self) -> bool {
        self != Self::Unexecutable
    }
}

#[cfg(all(feature = "network", not(feature = "local")))]
impl std::fmt::Display for NetworkFailure {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let s = match self {
            Self::NoBid => "no prover bid within auction timeout",
            Self::ProvingTimeout => "assigned but proving timed out",
            Self::Unexecutable => "unexecutable",
            Self::Unfulfillable => "unfulfillable",
            Self::Reverted => "reverted during settlement",
            Self::Expired => "expired",
            Self::Rpc => "rpc error",
            Self::Other => "other",
        };
        f.write_str(s)
    }
}

#[cfg(test)]
mod tests {

    use sp1_sdk::{
        Elf, HashableKey, Prover, ProverClient, ProvingKey, SP1ProofWithPublicValues, network::B256,
    };
    use sp1_verifier::PlonkVerifier;

    use crate::{
        BATCH_VERIFIER_ELF,
        evm::{EvmProofFixture, save_plonk_fixture},
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

    #[cfg(all(feature = "network", not(feature = "local")))]
    #[test]
    fn classify_network_failures() {
        use super::NetworkFailure;
        use sp1_sdk::network::Error as NetErr;

        let id = vec![0u8; 32];
        let case = |e: NetErr| NetworkFailure::classify(&anyhow::Error::from(e).context("ctx"));
        assert_eq!(
            case(NetErr::RequestAuctionTimedOut { request_id: id.clone() }),
            NetworkFailure::NoBid
        );
        assert_eq!(
            case(NetErr::RequestTimedOut { request_id: id.clone() }),
            NetworkFailure::ProvingTimeout
        );
        assert_eq!(
            case(NetErr::RequestUnexecutable { request_id: id.clone() }),
            NetworkFailure::Unexecutable
        );
        assert!(!NetworkFailure::Unexecutable.retryable());
        assert!(NetworkFailure::NoBid.retryable());

        // Same shape as the SDK's retry path: a raw tonic::Status inside anyhow, plus our context.
        // Fails if our tonic version differs from sp1-sdk's.
        let base_fee = anyhow::Error::from(tonic::Status::invalid_argument(
            "base fee 0.4085 PROVE is not accepted for Plonk proofs",
        ))
        .context("network prove error");
        assert_eq!(NetworkFailure::classify(&base_fee), NetworkFailure::Rpc);
        assert_eq!(NetworkFailure::classify(&anyhow::anyhow!("boom")), NetworkFailure::Other);
    }
}
