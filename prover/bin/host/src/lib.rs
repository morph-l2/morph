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
#[cfg(feature = "network")]
use sp1_sdk::{network::NetworkMode, NetworkProver};
use sp1_sdk::{HashableKey, Prover, ProverClient, SP1ProvingKey, SP1Stdin, SP1VerifyingKey};
use sp1_verifier::PlonkVerifier;
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const BATCH_VERIFIER_ELF: &[u8] = include_bytes!("../../client/elf/verifier-client");

/// The maximum number of blocks that can be included in a single proof.
const MAX_PROVE_BLOCKS: usize = 4096;

/// A batch prover that can generate and verify proofs for a batch of EVM transactions.
pub struct BatchProver<C> {
    prover_client: C,
    pk: SP1ProvingKey,
    vk: SP1VerifyingKey,
}

#[cfg(feature = "network")]
pub type DefaultClient = NetworkProver;
#[cfg(feature = "local")]
pub type DefaultClient = CpuProver;

/// A batch prover that uses the default proving client based on the feature flag.
impl Default for BatchProver<DefaultClient> {
    fn default() -> Self {
        let prover_client = {
            #[cfg(feature = "network")]
            {
                ProverClient::builder()
                    .network_for(NetworkMode::Mainnet)
                    .rpc_url("https://rpc.mainnet.succinct.xyz")
                    .build()
            }
            #[cfg(feature = "local")]
            {
                ProverClient::builder().cpu().build()
            }
        };
        let (pk, vk) = prover_client.setup(BATCH_VERIFIER_ELF);
        log::info!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());
        Self { prover_client, pk, vk }
    }
}

impl BatchProver<DefaultClient> {
    /// Proves a batch of EVM transactions and verifies the proof.
    pub fn prove(
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
                .execute(BATCH_VERIFIER_ELF, &stdin)
                .run()
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
        let mut proof =
            self.prover_client.prove(&self.pk, &stdin).plonk().run().context("proving failed")?;
        let duration_mins = start.elapsed().as_secs() / 60;
        log::info!("Successfully generated proof!, time use: {} minutes", duration_mins);

        // Verify the proof.
        let plonk_proof = proof.bytes();
        let public_inputs = proof.public_values.to_vec();
        let plonk_vk = *sp1_verifier::PLONK_VK_BYTES;
        PlonkVerifier::verify(&plonk_proof, &public_inputs, &self.vk.bytes32(), plonk_vk)
            .context("failed to verify proof")?;
        log::info!("Successfully verified proof!");

        // Deserialize the public values.
        let pi_bytes = proof.public_values.read::<[u8; 32]>();
        log::info!(
            "pi_hash generated with sp1-vm prove: {}",
            alloy::hex::encode_prefixed(pi_bytes)
        );
        let fixture = EvmProofFixture {
            vkey: self.vk.bytes32().to_string(),
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
    use prover_executor_client::{
        types::{
            blob::{decode_transactions, get_origin_batch},
            input::BlobInfo,
        },
        BlobVerifier,
    };
    use prover_executor_host::blob::{encode_blob, populate_kzg};
    use prover_primitives::{MorphTxEnvelope, B256};
    #[test]
    fn test_blob() {
        //blob to txn
        let blob_bytes = load_zstd_blob();
        println!("blob_bytes len: {:?}", blob_bytes.len());

        let origin_batch = get_origin_batch(&blob_bytes).unwrap();
        println!("origin_batch len: {:?}", origin_batch.len());

        let mut block_contexts = origin_batch[0..600 * 60].to_vec();
        let txs_data = origin_batch[600 * 60..origin_batch.len()].to_vec();
        let tx_list: Vec<MorphTxEnvelope> = decode_transactions(txs_data.as_slice());
        println!("decoded tx_list_len: {:?}", tx_list.len());

        //txn to blob
        let mut tx_bytes: Vec<u8> = vec![];
        let x = tx_list.iter().flat_map(|tx| tx.rlp()).collect::<Vec<u8>>();
        tx_bytes.extend(x);
        assert!(tx_bytes == txs_data, "tx_bytes==txs_data");
        block_contexts.extend_from_slice(&tx_bytes);
        let blob = encode_blob(block_contexts).unwrap();
        let blob_info: BlobInfo = populate_kzg(&blob).unwrap();
        let (versioned_hash, batch_data) = BlobVerifier::verify(&blob_info, 600).unwrap();
        let versioned_hash_hex = alloy::hex::encode_prefixed(versioned_hash.as_slice());
        println!(
            "versioned_hash: {:?}, batch_data len: {:?}",
            versioned_hash_hex,
            batch_data.len()
        );
        assert!(
            versioned_hash_hex
                == "0x012bdf80720ba8d07c589d672e47d4b183ac861a2fcb6a5dad0e320a4f368f4f",
            "versioned_hash check"
        );

        assert!(batch_data.len() == origin_batch.len(), "batch_data.len() == origin_batch.len()");
    }

    pub fn load_zstd_blob() -> [u8; 131072] {
        use prover_primitives::alloy_primitives::hex;
        use std::{fs, path::Path};

        //https://etherscan.io/blob/0x012bdf80720ba8d07c589d672e47d4b183ac861a2fcb6a5dad0e320a4f368f4f?bid=6318849
        //https://explorer.morphl2.io/batches/47561
        let blob_data_path = Path::new("../../testdata/blob/mainnet_47561.data");
        let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
        let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();
        let mut array = [0u8; 131072];
        array.copy_from_slice(&hex_data);
        array
    }

    use sp1_sdk::{HashableKey, ProverClient, SP1ProofWithPublicValues};
    use sp1_verifier::PlonkVerifier;

    use crate::{
        evm::{save_plonk_fixture, EvmProofFixture},
        BATCH_VERIFIER_ELF,
    };

    #[test]
    pub fn verify_proof() {
        let proof_loaded = SP1ProofWithPublicValues::load("../../proof/artifact_mainnet").unwrap();

        let proof = proof_loaded.bytes();
        let public_inputs = proof_loaded.public_values.to_vec();

        let client = ProverClient::from_env();

        let (_pk, vk) = client.setup(BATCH_VERIFIER_ELF);
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
