use sp1_helper::{build_program_with_args, BuildArgs};

fn main() {
    build_program_with_args(
        "../client",
        BuildArgs {
            ignore_rust_version: true,
            output_directory: "tests/keccak256/client/elf".to_string(),
            ..Default::default()
        },
    )
}
