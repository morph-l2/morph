use sp1_helper::{build_program_with_args, BuildArgs};

fn main() {
    build_program_with_args(
        "../client/shard",
        BuildArgs {
            ignore_rust_version: true,
            output_directory: "bin/client/shard/elf".to_string(),
            ..Default::default()
        },
    );
    build_program_with_args(
        "../client/agg",
        BuildArgs {
            ignore_rust_version: true,
            output_directory: "bin/client/agg/elf".to_string(),
            ..Default::default()
        },
    );
}
