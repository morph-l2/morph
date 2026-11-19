#!/usr/bin/env bash
# check-no-secrets.sh — run before committing: fails if a tracked-or-untracked, non-ignored file in this directory
# contains something that looks like a private key (0x + 64 hex digits) or a secrets file.
set -uo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."
bad=0
# well-known PUBLIC values: Anvil/Hardhat account #0 dev key, EIP-1967 implementation / admin storage slots
ALLOW='ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80|360894a13ba1a3210667c828492db98dca3e2076cc373|b53127684a568b3173ae13b9f8a6016e243e63b6e8ee'
PAT='(^|[^0-9a-fA-F])(0x)?[0-9a-fA-F]{64}([^0-9a-fA-F]|$)'
hits() { grep -nE "${PAT}" "$1" | grep -vE "${ALLOW}"; }
while IFS= read -r f; do
  [[ -f "$f" ]] || continue
  if hits "$f" >/dev/null 2>&1; then
    echo "possible 32-byte hex value in: $f"
    hits "$f" | cut -c1-60 | sed 's/^/    /'
    bad=1
  fi
done < <(git ls-files -co --exclude-standard -- .)
case "$(git ls-files -co --exclude-standard -- env | tr '\n' ' ')" in *secrets*) echo "a secrets file is not ignored"; bad=1;; esac
[[ $bad -eq 0 ]] && echo "ok: nothing that looks like a private key"
exit $bad
