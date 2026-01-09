# MPT Switch Test

Test the client switching logic of sequencer node and sentry node at MPT upgrade time.

## Architecture

```
Before upgrade:                      After upgrade (swap):
┌─────────────┐                     ┌─────────────┐
│  Sequencer  │──► ZK Geth          │  Sequencer  │──► MPT Geth
│    Node     │    (:8545)          │    Node     │    (:9545)
└─────────────┘                     └─────────────┘
                        ──────►
┌─────────────┐                     ┌─────────────┐
│   Sentry    │──► MPT Geth         │   Sentry    │──► ZK Geth
│    Node     │    (:9545)          │    Node     │    (:8545)
└─────────────┘                     └─────────────┘
```

**Key points:**
- Both nodes share the same two Geth instances
- When upgrade time is reached, both nodes swap Geth connections
- Sequencer: ZK Geth → MPT Geth
- Sentry: MPT Geth → ZK Geth

## Prerequisites

### Prepare Binaries

Place all binaries in the `bin` directory:

```bash
ops/mpt-switch-test/bin/
├── geth
├── morphnode
└── tendermint
```

Genesis file (`genesis-l2.json`) is already included in the directory.

## Usage

```bash
cd /Users/corey.zhang/workspace/morph/ops/mpt-switch-test

# 1. Start test environment (switch triggers after 60 seconds)
./test-mpt-switch-local.sh start 60

# 2. Monitor Sequencer switch logs
./test-mpt-switch-local.sh monitor sequencer

# 3. Monitor Sentry switch logs
./test-mpt-switch-local.sh monitor sentry

# 4. Check status
./test-mpt-switch-local.sh status

# 5. Stop services
./test-mpt-switch-local.sh stop

# 6. Clean data
./test-mpt-switch-local.sh clean
```

## Command List

| Command | Description |
|---------|-------------|
| `start [delay]` | Start test environment, delay is MPT switch delay in seconds (default 60) |
| `stop` | Stop all services |
| `clean` | Clean all test data |
| `status` | Check service status and block height |
| `monitor [target]` | Monitor logs (sequencer/sentry/all) |
| `logs [service]` | View logs (sequencer/sentry/zk-geth/mpt-geth/all) |

## Port Allocation

| Service | HTTP | Engine | P2P |
|---------|------|--------|-----|
| ZK Geth | 8545 | 8551 | 30303 |
| MPT Geth | 9545 | 9551 | 30304 |
| Sequencer Node | - | - | 26656 (RPC: 26657) |
| Sentry Node | - | - | 26756 (RPC: 26757) |

## Expected Logs

Both nodes should see similar switch logs:

```
MPT switch time reached, MUST wait for MPT node to sync
  mpt_time=<timestamp> current_time=<timestamp> target_block=<number>

Waiting for MPT node to sync...
  remote_block=<n> target_block=<m> blocks_behind=<diff>

Successfully switched to MPT client
  remote_block=<n> target_block=<m> wait_duration=<duration>
```

## File Structure

```
mpt-switch-test/
├── test-mpt-switch-local.sh  # Test script
├── verify-migration.sh       # Migration verification script
├── check-nodes.sh            # Node status check script
├── send-txs.sh               # Transaction sending script
├── README.md                 # This document
├── genesis-zk.json           # ZK mode genesis file
├── genesis-mpt.json          # MPT mode genesis file
├── bin/                      # Binary directory (needs to be placed manually)
│   ├── geth
│   ├── morphnode
│   ├── tendermint
│   └── migration-checker     # Built by verify-migration.sh
└── .testdata/                # Test data directory (generated on startup)
    ├── zk-geth/              # ZK Geth data
    ├── mpt-geth/             # MPT Geth data
    ├── sequencer-node/       # Sequencer Node data
    ├── sentry-node/          # Sentry Node data
    ├── jwt-secret.txt        # JWT secret
    └── *.log                 # Log files
```

## Migration Verification

After the MPT switch, you can verify that the migration was correct by comparing all account and storage data between ZK Geth and MPT Geth.

### Prerequisites

1. **go-ethereum on trie-checker branch**:
   ```bash
   cd /path/to/go-ethereum
   git checkout trie-checker
   ```

2. **Test environment running** with both nodes synced

### Verification Commands

```bash
# 1. Build migration-checker tool
./verify-migration.sh build

# 2. Check verification readiness
./verify-migration.sh status

# 3. Auto-verify at latest common block
./verify-migration.sh auto

# 4. Verify at specific block
./verify-migration.sh verify 0x10      # hex
./verify-migration.sh verify 100       # decimal

# 5. Paranoid mode (verify all node hashes)
./verify-migration.sh auto --paranoid
```

### What Migration-Checker Verifies

| Item | Description |
|------|-------------|
| Account count | Number of accounts must match |
| Nonce | Each account's nonce must be identical |
| Balance | Each account's balance must be identical |
| CodeHash | Contract code hashes must match |
| Storage | All storage slot values must match |

### Expected Output

```
[STEP] Running migration-checker...
  ZK DB:    .testdata/zk-geth/geth/chaindata
  MPT DB:   .testdata/mpt-geth/geth/chaindata
  ZK Root:  0x1234...
  MPT Root: 0x5678...

Accounts done: 1
Accounts done: 2
...
Accounts done: N

[SUCCESS] ==========================================
[SUCCESS]   Migration verification PASSED!
[SUCCESS]   All accounts and storage data match.
[SUCCESS] ==========================================
```
