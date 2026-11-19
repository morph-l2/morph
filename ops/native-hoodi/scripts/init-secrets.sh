#!/usr/bin/env bash
# init-secrets.sh — create env/secrets.env (mode 600) from the template: your deployer key + three freshly generated keys.
# The deployer key comes from the environment (DEPLOYER_PRIVATE_KEY), never from the command line. Prints only addresses.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${ROOT}/scripts/lib/env.sh"
OUT="${ROOT}/env/secrets.env"
[[ -n "${DEPLOYER_PRIVATE_KEY:-}" ]] || { echo "error: set DEPLOYER_PRIVATE_KEY in the environment first" >&2; exit 1; }
[[ ! -e "${OUT}" ]] || { echo "error: ${OUT} exists; move it away first (scripts/reset-env.sh --new-keys does that)" >&2; exit 1; }
python3 -I - "${ROOT}/env/hoodi.env.example" "${OUT}" <<'PY'
import json, os, re, subprocess, sys
src, dst = sys.argv[1:3]
def new_key():
    out = json.loads(subprocess.check_output(["cast", "wallet", "new", "--json"]))
    if isinstance(out, dict) and "data" in out: out = out["data"]
    if isinstance(out, list): out = out[0]
    return out["private_key"]
keys = {"DEPLOYER_PRIVATE_KEY": os.environ["DEPLOYER_PRIVATE_KEY"], "SEQUENCER_PRIVATE_KEY": new_key(),
        "BATCH_SUBMITTER_PRIVATE_KEY": new_key(), "L2_GAS_ORACLE_PRIVATE_KEY": new_key()}
text = open(src).read()
for k, v in keys.items():
    text, n = re.subn(r'^(export ' + k + r'=).*$', lambda m: m.group(1) + '"' + v + '"', text, flags=re.M)
    assert n == 1, k
fd = os.open(dst, os.O_WRONLY | os.O_CREAT | os.O_EXCL, 0o600)
with os.fdopen(fd, "w") as f: f.write(text)
PY
set -a; source "${OUT}"; set +a
for r in DEPLOYER SEQUENCER BATCH_SUBMITTER L2_GAS_ORACLE; do v="${r}_PRIVATE_KEY"; printf '%-16s %s\n' "${r}" "$(cast wallet address --private-key "${!v}")"; done
echo "written ${OUT} (mode 600). Fund the DEPLOYER address with ~1.5 Hoodi ETH, then follow README.md."
