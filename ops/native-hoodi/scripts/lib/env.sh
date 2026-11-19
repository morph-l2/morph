# lib/env.sh — sourced by every script. Prepares the toolchain and fails early, with a clear message, if something is missing.
#
#   MORPH_ENV_FILE=/path/to/env.sh   optional: a file that puts Go / Rust / Node / Foundry on PATH (and sets any build flags
#                                    your machine needs). It is sourced first. Without it, the tools must already be on PATH.
#
# Runtime tools (every script):  node  cast  python3  curl  (anvil only for dryrun-fork.sh)
# Build tools (build-all.sh):    go (>= 1.24)  cargo  yarn  make  gcc
# This file only checks; it installs nothing.

if [[ -n "${MORPH_ENV_FILE:-}" ]]; then
  [[ -f "${MORPH_ENV_FILE}" ]] || { echo "error: MORPH_ENV_FILE=${MORPH_ENV_FILE} does not exist" >&2; exit 1; }
  # shellcheck disable=SC1090
  source "${MORPH_ENV_FILE}" >/dev/null 2>&1
fi

_missing=()
for _t in node cast python3 curl; do command -v "${_t}" >/dev/null 2>&1 || _missing+=("${_t}"); done
if [[ ${#_missing[@]} -gt 0 ]]; then
  echo "error: missing tools on PATH: ${_missing[*]}" >&2
  echo "  node >= 18, python3 >= 3.9, curl, and Foundry (cast, anvil; tested with 1.9.0: https://getfoundry.sh)." >&2
  echo "  If they are installed somewhere else, point MORPH_ENV_FILE at a file that sets PATH (see README.md)." >&2
  exit 1
fi
unset _missing _t

# build-all.sh calls this before it compiles anything
require_build_tools() {
  local m=() t gov
  for t in go cargo yarn make gcc; do command -v "${t}" >/dev/null 2>&1 || m+=("${t}"); done
  if [[ ${#m[@]} -gt 0 ]]; then
    echo "error: missing build tools on PATH: ${m[*]}" >&2
    echo "  needs Go >= 1.24 (tested 1.25.7), Rust (tested 1.95; needs OpenSSL headers + pkg-config), Node + yarn 1.x, make, gcc." >&2
    exit 1
  fi
  gov="$(go env GOVERSION 2>/dev/null | sed 's/^go//')"
  if ! printf '1.24.0\n%s\n' "${gov}" | sort -V -C; then
    echo "error: Go ${gov} is too old; Morph's go.work needs >= 1.24.0" >&2
    exit 1
  fi
}
