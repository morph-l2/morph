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
├── README.md                 # This document
├── genesis-l2.json           # L2 genesis file (included)
├── bin/                      # Binary directory (needs to be placed manually)
│   ├── geth
│   ├── morphnode
│   └── tendermint
└── .testdata/                # Test data directory (generated on startup)
    ├── zk-geth/              # ZK Geth data
    ├── mpt-geth/             # MPT Geth data
    ├── sequencer-node/       # Sequencer Node data
    ├── sentry-node/          # Sentry Node data
    ├── jwt-secret.txt        # JWT secret
    └── *.log                 # Log files
```
