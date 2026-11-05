use std::{
    fs,
    fs::File,
    io::BufReader,
    path::{Path, PathBuf},
};

use serde::de::DeserializeOwned;

pub fn resolve_block_input_files(paths: &[String]) -> Vec<PathBuf> {
    // Default: run all *.json under testdata/witnesses/
    if paths.is_empty() {
        let dir = default_block_inputs_dir();
        return list_json_files(&dir);
    }

    let mut out = Vec::new();
    for p in paths {
        let pb = PathBuf::from(p);
        if pb.is_dir() {
            out.extend(list_json_files(&pb));
        } else {
            out.push(pb);
        }
    }

    out.sort();
    out
}

fn default_block_inputs_dir() -> PathBuf {
    // bin/host (manifest dir) -> repo_root/testdata/block_inputs
    PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("../../testdata/block_inputs")
}

fn list_json_files(dir: &Path) -> Vec<PathBuf> {
    let mut files = Vec::new();
    if let Ok(rd) = fs::read_dir(dir) {
        for entry in rd.flatten() {
            let path = entry.path();
            if path
                .extension()
                .and_then(|e| e.to_str())
                .is_some_and(|e| e.eq_ignore_ascii_case("json"))
            {
                files.push(path);
            }
        }
    }
    files.sort();
    files
}

pub fn load_inputs<T: DeserializeOwned>(file_path: &str) -> Vec<T> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}
