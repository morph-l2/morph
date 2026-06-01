package derivation

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmnode "github.com/tendermint/tendermint/node"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	nodecommon "morph-l2/node/common"
	"morph-l2/node/sync"
	"morph-l2/node/types"
)

var (
	RollupEventTopic     = "CommitBatch(uint256,bytes32)"
	RollupEventTopicHash = crypto.Keccak256Hash([]byte(RollupEventTopic))
)

type Derivation struct {
	ctx                   context.Context
	node                  *tmnode.Node
	syncer                *sync.Syncer
	l1Client              *ethclient.Client
	RollupContractAddress common.Address
	confirmations         rpc.BlockNumber
	l2Client              *types.RetryableClient
	logger                tmlog.Logger
	rollup                *bindings.Rollup
	metrics               *Metrics
	l1BeaconClient        *L1BeaconClient
	L2ToL1MessagePasser   *bindings.L2ToL1MessagePasser

	rollupABI             *abi.ABI
	legacyRollupABI       *abi.ABI // before remove skipMap
	beforeMoveBlockCtxABI *abi.ABI

	db Database

	cancel context.CancelFunc

	startHeight         uint64
	baseHeight          uint64
	fetchBlockRange     uint64
	pollInterval        time.Duration
	logProgressInterval time.Duration
	verifyMode          string // SPEC-005 section 4.2: "layer1" or "local" (default); bound at startup, never switches.
	reorgCheckDepth     uint64 // SPEC-005 section 4.7.6: how far back to scan for L1 hash divergence each poll.

	tagAdvancer *tagAdvancer

	isHaMode bool

	stop chan struct{}
}

type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
	ethereum.ChainReader
	ethereum.TransactionReader
}

// NewDerivationClient takes a shared l1Client owned by main.go. See
// sync.NewSyncer for rationale — every L1-touching component in this
// process shares one connection pool / retry / metrics surface.
func NewDerivationClient(ctx context.Context, cfg *Config, syncer *sync.Syncer, db Database, rollup *bindings.Rollup, l1Client *ethclient.Client, node *tmnode.Node, isHaMode bool, logger tmlog.Logger) (*Derivation, error) {
	if l1Client == nil {
		return nil, errors.New("l1Client cannot be nil")
	}
	aClient, err := authclient.DialContext(context.Background(), cfg.L2.EngineAddr, cfg.L2.JwtSecret)
	if err != nil {
		return nil, err
	}
	eClient, err := ethclient.Dial(cfg.L2.EthAddr)
	if err != nil {
		return nil, err
	}

	msgPasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, eClient)
	if err != nil {
		return nil, err
	}
	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	legacyRollupAbi, err := types.LegacyRollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	beforeMoveBlockCtxAbi, err := types.BeforeMoveBlockCtxABI.GetAbi()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	logger = logger.With("module", "derivation")
	metrics := PrometheusMetrics("morphnode")
	if cfg.MetricsServerEnable {
		go func() {
			_, err := metrics.Serve(cfg.MetricsHostname, cfg.MetricsPort)
			if err != nil {
				panic(fmt.Errorf("metrics server start error:%v", err))
			}
		}()
		logger.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
	}
	baseHttp := NewBasicHTTPClient(cfg.BeaconRpc, logger)
	l1BeaconClient := NewL1BeaconClient(baseHttp)

	l2Client := types.NewRetryableClient(aClient, eClient, logger)
	tagAdv := newTagAdvancer(l2Client, metrics, logger)

	return &Derivation{
		ctx:                   ctx,
		node:                  node,
		db:                    db,
		l1Client:              l1Client,
		syncer:                syncer,
		rollup:                rollup,
		rollupABI:             rollupAbi,
		legacyRollupABI:       legacyRollupAbi,
		beforeMoveBlockCtxABI: beforeMoveBlockCtxAbi,
		logger:                logger,
		RollupContractAddress: cfg.RollupContractAddress,
		confirmations:         cfg.L1.Confirmations,
		l2Client:              l2Client,
		cancel:                cancel,
		stop:                  make(chan struct{}),
		startHeight:           cfg.StartHeight,
		baseHeight:            cfg.BaseHeight,
		fetchBlockRange:       cfg.FetchBlockRange,
		pollInterval:          cfg.PollInterval,
		logProgressInterval:   cfg.LogProgressInterval,
		verifyMode:            cfg.VerifyMode,
		reorgCheckDepth:       cfg.ReorgCheckDepth,
		tagAdvancer:           tagAdv,
		metrics:               metrics,
		l1BeaconClient:        l1BeaconClient,
		L2ToL1MessagePasser:   msgPasser,
		isHaMode:              isHaMode,
	}, nil
}

func (d *Derivation) Start() {
	// Single-goroutine design: the SPEC-005 finalizer step runs at the end
	// of each derivationBlock iteration (see finalizer.go::finalizerTick).
	// Folded into the main loop so the cursor-rewind recovery paths
	// (handleL1Reorg + finalizerTick canonicality fail) don't race with
	// each other on the L1 cursor / tagAdvancer.
	go func() {
		d.syncer.Start()
		t := time.NewTicker(d.pollInterval)
		defer t.Stop()

		for {
			// don't wait for ticker during startup
			d.derivationBlock(d.ctx)
			d.finalizerTick()

			select {
			case <-d.ctx.Done():
				d.logger.Error("derivation node Unexpected exit")
				close(d.stop)
				return
			case <-t.C:
				continue
			}
		}
	}()
}

func (d *Derivation) Stop() {
	if d == nil {
		return
	}

	d.logger.Info("stopping derivation service")

	if d.cancel != nil {
		d.cancel()
	}
	<-d.stop
	d.logger.Info("derivation service is stopped")
}

func (d *Derivation) derivationBlock(ctx context.Context) {
	// SPEC-005 §4.7.6: check for an L1 reorg before processing any new logs.
	// The scan is a no-op when --derivation.confirmations=finalized (L1
	// finalized doesn't reorg by Ethereum consensus assumption) and
	// load-bearing when configured below finalized; the gate is intentionally
	// absent so behavior is uniform across configs.
	if reorgAt, err := d.detectReorg(ctx); err != nil {
		d.logger.Error("L1 reorg detection failed; skipping this poll", "err", err)
		return
	} else if reorgAt != nil {
		if err := d.handleL1Reorg(*reorgAt); err != nil {
			d.logger.Error("handle L1 reorg failed", "err", err)
		}
		// Don't process further this cycle: cursor was rewound, let the next
		// poll re-fetch from the new starting point. Avoids recording
		// potentially-still-unstable L1 hashes if the chain is mid-reorg.
		return
	}

	latestDerivation := d.db.ReadLatestDerivationL1Height()
	latest, err := d.getLatestConfirmedBlockNumber(d.ctx)
	if err != nil {
		d.logger.Error("get latest block number failed", "err", err)
		return
	}
	var start uint64
	if latestDerivation == nil {
		start = d.startHeight
	} else {
		start = *latestDerivation + 1
	}
	end := latest
	if latest < start {
		d.logger.Info("latest less than start", "latest", latest, "start", start)
		return
	} else if latest-start >= d.fetchBlockRange {
		end = start + d.fetchBlockRange
	}
	d.logger.Info("derivation start pull rollupData form l1", "startBlock", start, "end", end)
	logs, err := d.fetchRollupLog(ctx, start, end)
	if err != nil {
		d.logger.Error("eth_getLogs failed", "err", err)
		return
	}
	latestBatchIndex, err := d.rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		d.logger.Error("query rollup latestCommitted batch Index failed", "err", err)
		return
	}
	d.metrics.SetLatestBatchIndex(latestBatchIndex.Uint64())
	d.logger.Info("fetched rollup tx", "txNum", len(logs), "latestBatchIndex", latestBatchIndex)

	for _, lg := range logs {
		var (
			batchInfo  *BatchInfo
			lastHeader *eth.Header
		)
		switch d.verifyMode {
		case VerifyModeLocal:
			batchInfo, err = d.fetchBatchInfoOutline(ctx, lg.TxHash, lg.BlockNumber)
			if err != nil {
				if errors.Is(err, types.ErrNotCommitBatchTx) {
					continue
				}
				d.logger.Error("fetch batch info outline failed", "err", err)
				return
			}
			d.logger.Info("local verify fetched batch metadata",
				"batchIndex", batchInfo.batchIndex,
				"version", batchInfo.version,
				"parentTotalL1Popped", batchInfo.parentTotalL1MessagePopped,
				"expectedBlobs", len(batchInfo.blobHashes),
				"txNonce", batchInfo.nonce, "txHash", batchInfo.txHash,
				"l1BlockNumber", batchInfo.l1BlockNumber, "firstL2BlockNumber", batchInfo.firstBlockNumber, "lastL2BlockNumber", batchInfo.lastBlockNumber)
			// SPEC-005 §4.3 Path B entry-point: scenario A/B/C/D dispatch.
			// rebuildBlob needs every block in [first..last] locally. If
			// lastBlockNumber is missing, retry-poll within an observation
			// window. Two early-exit signals:
			//   - header arrives → P2P delivered, fall through to A/B
			//     (rebuildBlob); whether the head grew is irrelevant.
			//   - header still missing but L2 head grew → sequencer alive,
			//     skip this batch this poll (scenario D).
			// Window exhausted with neither signal → sequencer stopped,
			// fall through to L1 blob fill-gap (scenario C: pull real
			// blob + deriveForce skipping blocks already present locally).
			lastHdr, hdrErr := d.l2Client.HeaderByNumber(ctx, big.NewInt(int64(batchInfo.lastBlockNumber)))
			if hdrErr != nil && !errors.Is(hdrErr, ethereum.NotFound) {
				d.logger.Error("local verify: HeaderByNumber for lastBlock failed",
					"batchIndex", batchInfo.batchIndex,
					"lastBlockNumber", batchInfo.lastBlockNumber,
					"err", hdrErr)
				return
			}
			if lastHdr == nil {
				snapshot, err := d.l2Client.BlockNumber(ctx)
				if err != nil {
					d.logger.Error("local verify: BlockNumber snapshot failed",
						"batchIndex", batchInfo.batchIndex, "err", err)
					return
				}
				var grew bool
				for i := 0; i < 3; i++ {
					select {
					case <-ctx.Done():
						return
					case <-time.After(2 * time.Second):
					}
					hdr, err := d.l2Client.HeaderByNumber(ctx, big.NewInt(int64(batchInfo.lastBlockNumber)))
					if err != nil && !errors.Is(err, ethereum.NotFound) {
						d.logger.Error("local verify: HeaderByNumber retry failed",
							"batchIndex", batchInfo.batchIndex,
							"lastBlockNumber", batchInfo.lastBlockNumber,
							"err", err)
						return
					}
					if hdr != nil {
						lastHdr = hdr
						break
					}
					cur, err := d.l2Client.BlockNumber(ctx)
					if err != nil {
						d.logger.Error("local verify: BlockNumber observation failed",
							"batchIndex", batchInfo.batchIndex, "err", err)
						return
					}
					if cur > snapshot {
						grew = true
					}
				}
				if lastHdr == nil {
					if grew {
						d.logger.Info("local verify: lastBlock missing but L2 head growing; skipping batch this poll (scenario D)",
							"batchIndex", batchInfo.batchIndex,
							"lastBlockNumber", batchInfo.lastBlockNumber)
						continue
					}
					// Scenario C: sequencer stopped → L1 blob fill-gap.
					d.logger.Info("local verify: lastBlock missing and L2 head flat; fallback to L1 blob fill-gap (scenario C)",
						"batchIndex", batchInfo.batchIndex,
						"lastBlockNumber", batchInfo.lastBlockNumber)
					batchInfoFull, fetchErr := d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
					if fetchErr != nil {
						if errors.Is(fetchErr, types.ErrNotCommitBatchTx) {
							continue
						}
						d.logger.Error("local verify fill-gap: fetch real batch failed",
							"batchIndex", batchInfo.batchIndex, "error", fetchErr)
						return
					}

					// Quiesce blocksync + broadcast reactors so they don't race
					// with the derivation-driven reorg below. HA sequencers and
					// mock-mode (d.node == nil) skip — sequencers don't
					// auto-reorg, mock has no consensus reactors. Stop/Start
					// are idempotent via IsRunning checks, so retrying next
					// poll is safe if Stop fails here.
					if d.node != nil {
						if err = d.node.StopReactorsBeforeReorg(); err != nil {
							d.logger.Error("StopReactorsBeforeReorg failed; skipping reorg, will retry next poll",
								"batchIndex", batchInfo.batchIndex, "err", err)
							return
						}
					}
					localLatest, err := d.l2Client.BlockNumber(ctx)
					if err != nil {
						d.logger.Error("local verify fill-gap: read local latest failed",
							"batchIndex", batchInfo.batchIndex, "error", err)
						return
					}
					lastHeader, err = d.deriveForce(batchInfoFull, localLatest)
					if err != nil {
						d.logger.Error("local verify fill-gap: derive failed",
							"batchIndex", batchInfo.batchIndex, "error", err)
						return
					}

					// Restart reactors using the post-reorg head height so
					// blocksync rebuilds its pool from currentHeight+1 and
					// catches back up via P2P. If this fails the L2 chain is
					// still correctly reorged but reactors are degraded —
					// surface loudly so it's visible in monitoring; next poll
					// will retry only if a *new* mismatch appears.
					if d.node != nil {
						if err = d.node.StartReactorsAfterReorg(lastHeader.Number.Int64()); err != nil {
							d.logger.Error("StartReactorsAfterReorg failed; chain is reorged but reactors are degraded",
								"batchIndex", batchInfo.batchIndex,
								"postReorgHeight", lastHeader.Number.Int64(),
								"err", err)
						}
					}

					d.metrics.SetL2DeriveHeight(lastHeader.Number.Uint64())
					d.metrics.SetSyncedBatchIndex(batchInfo.batchIndex)
					break
				}
			}

			rebuilt, err := d.rebuildBlob(ctx, batchInfo)
			if err != nil {
				d.logger.Error("rebuildBlob failed", "err", err)
				return
			}
			lastHeader, err = d.fetchLocalLastHeader(ctx, batchInfo)
			if err != nil {
				d.logger.Error("local verify local last-header fetch failed", "batchIndex", batchInfo.batchIndex, "error", err)
				return
			}
			for i := range rebuilt {
				if rebuilt[i] != batchInfo.blobHashes[i] {
					// HA-mode invariant: blocks are committed via Raft consensus and
					// the L1 batch is built from those committed blocks, so the
					// rebuilt blob hash MUST equal the on-chain blob hash. A
					// mismatch here means the local L2 chain has diverged from what
					// the cluster committed — possible causes: corrupted DB, wrong
					// genesis, manual chain surgery, or a Raft / sequencer bug.
					// Auto-reorg is unsafe (would mask the real problem), so we
					// hard-stop derivation and require operator intervention.
					if d.isHaMode {
						d.logger.Error("HA node: blob hash mismatch detected — derivation halted, manual intervention required (this should never happen in a healthy cluster)",
							"batchIndex", batchInfo.batchIndex,
							"blobIndex", i,
							"expected", batchInfo.blobHashes[i].Hex(),
							"rebuilt", rebuilt[i].Hex(),
							"l1TxHash", lg.TxHash.Hex(),
							"l1BlockNumber", lg.BlockNumber)
						return
					}

					d.logger.Info("blob hash mismatch; triggering self-heal reorg",
						"batchIndex", batchInfo.batchIndex,
						"expected", batchInfo.blobHashes[i].Hex(),
						"rebuilt", rebuilt[i].Hex())

					batchInfoFull, fetchErr := d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
					if fetchErr != nil {
						d.logger.Error("local verify self-heal: fetch real batch failed",
							"batchIndex", batchInfo.batchIndex, "error", fetchErr)
						return
					}

					// Quiesce blocksync + broadcast reactors so they don't race
					// with the derivation-driven reorg below. HA sequencers and
					// mock-mode (d.node == nil) skip — sequencers don't
					// auto-reorg, mock has no consensus reactors. Stop/Start
					// are idempotent via IsRunning checks, so retrying next
					// poll is safe if Stop fails here.
					if d.node != nil {
						if err = d.node.StopReactorsBeforeReorg(); err != nil {
							d.logger.Error("StopReactorsBeforeReorg failed; skipping reorg, will retry next poll",
								"batchIndex", batchInfo.batchIndex, "err", err)
							return
						}
					}
					lastHeader, err = d.deriveForce(batchInfoFull, 0)
					if err != nil {
						d.logger.Error("local verify self-heal: derive failed",
							"batchIndex", batchInfo.batchIndex, "error", err)
						return
					}

					// Restart reactors using the post-reorg head height so
					// blocksync rebuilds its pool from currentHeight+1 and
					// catches back up via P2P. If this fails the L2 chain is
					// still correctly reorged but reactors are degraded —
					// surface loudly so it's visible in monitoring; next poll
					// will retry only if a *new* mismatch appears.
					if d.node != nil {
						if err = d.node.StartReactorsAfterReorg(lastHeader.Number.Int64()); err != nil {
							d.logger.Error("StartReactorsAfterReorg failed; chain is reorged but reactors are degraded",
								"batchIndex", batchInfo.batchIndex,
								"postReorgHeight", lastHeader.Number.Int64(),
								"err", err)
						}
					}
					break
				}
			}

			d.metrics.SetL2DeriveHeight(batchInfo.lastBlockNumber)
			d.metrics.SetSyncedBatchIndex(batchInfo.batchIndex)
		case VerifyModeLayer1:
			batchInfo, err = d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
			if err != nil {
				if errors.Is(err, types.ErrNotCommitBatchTx) {
					continue
				}
				d.logger.Error("fetch batch info failed", "txHash", lg.TxHash, "blockNumber", lg.BlockNumber, "error", err)
				return
			}
			d.logger.Info("fetch rollup transaction success", "txNonce", batchInfo.nonce, "txHash", batchInfo.txHash,
				"l1BlockNumber", batchInfo.l1BlockNumber, "firstL2BlockNumber", batchInfo.firstBlockNumber, "lastL2BlockNumber", batchInfo.lastBlockNumber)
			lastHeader, err = d.derive(batchInfo)
			if err != nil {
				d.logger.Error("derive blocks interrupt", "error", err)
				return
			}
			d.logger.Info("batch derivation complete", "batch_index", batchInfo.batchIndex, "currentBatchEndBlock", lastHeader.Number.Uint64())
			d.metrics.SetL2DeriveHeight(lastHeader.Number.Uint64())
			d.metrics.SetSyncedBatchIndex(batchInfo.batchIndex)
		default:
			// Unreachable: validateAndDefaultVerifyMode rejects unknown values
			// at startup and normalises empty to DefaultVerifyMode (local).
			// If we get here it's a programming error -- a new mode added to
			// the constant set without a switch arm. Fail loud rather than
			// silently fall through to stale semantics.
			d.logger.Error("unknown verifyMode reached derivationBlock; refusing to process batch", "verifyMode", d.verifyMode)
			return
		}

		if lastHeader.Number.Uint64() <= d.baseHeight {
			continue
		}
		if err := d.verifyBatchRoots(batchInfo, lastHeader); err != nil {
			// stateException only when the verifier produced a real mismatch
			// verdict (root or withdrawal root). Transient failures (e.g.
			// MessageRoot RPC error) just log and retry next poll.
			if errors.Is(err, ErrBatchVerifyDivergence) {
				d.metrics.SetBatchStatus(stateException)
			}
			d.logger.Error("batch roots verification failed", "batchIndex", batchInfo.batchIndex, "error", err)
			return
		}
		d.metrics.SetBatchStatus(stateNormal)
		d.metrics.SetL1SyncHeight(lg.BlockNumber)

		// SPEC-005 section 4.7.3: a verified batch (layer1 or local verify) advances safe.
		d.tagAdvancer.advanceSafe(d.ctx, batchInfo.batchIndex, lastHeader)
	}

	// SPEC-005 §4.7.6: record this poll's L1 block hashes so the next poll
	// can detect a reorg. Failure here must NOT advance the cursor -- a gap
	// in the recorded hashes would defeat detection across that gap.
	if err := d.recordL1Blocks(ctx, start, end); err != nil {
		d.logger.Error("recordL1Blocks failed; skipping cursor advance, will retry next poll", "err", err)
		return
	}

	d.db.WriteLatestDerivationL1Height(end)
	d.metrics.SetL1SyncHeight(end)
	d.logger.Info("write latest derivation l1 height success", "l1BlockNumber", end)
}

func (d *Derivation) fetchRollupLog(ctx context.Context, from, to uint64) ([]eth.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(from),
		ToBlock:   big.NewInt(0).SetUint64(to),
		Addresses: []common.Address{
			d.RollupContractAddress,
		},
		Topics: [][]common.Hash{
			{RollupEventTopicHash},
		},
	}
	return d.l1Client.FilterLogs(ctx, query)
}

func (d *Derivation) fetchRollupDataByTxHash(txHash common.Hash, blockNumber uint64) (*BatchInfo, error) {
	tx, pending, err := d.l1Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	if pending {
		return nil, errors.New("pending transaction")
	}
	batch, err := d.UnPackData(tx.Data())
	if err != nil {
		return nil, err
	}

	// Get block header to retrieve timestamp
	header, err := d.l1Client.HeaderByNumber(d.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	// Get transaction blob hashes
	blobHashes := tx.BlobHashes()
	if len(blobHashes) > 0 {
		d.logger.Info("Transaction contains blobs", "txHash", txHash, "blobCount", len(blobHashes))

		// Initialize indexedBlobHashes as nil
		var indexedBlobHashes []IndexedBlobHash

		// Only try to build IndexedBlobHash array if not forcing get all blobs
		// Try to get the block to build IndexedBlobHash array
		block, err := d.l1Client.BlockByNumber(d.ctx, big.NewInt(int64(blockNumber)))
		if err == nil {
			// Successfully got the block, now build IndexedBlobHash array
			d.logger.Info("Building IndexedBlobHash array from block", "blockNumber", blockNumber)
			indexedBlobHashes = dataAndHashesFromTxs(block.Transactions(), tx)
			d.logger.Info("Built IndexedBlobHash array", "count", len(indexedBlobHashes))
		} else {
			d.logger.Info("Failed to get block, will try fetching all blobs", "blockNumber", blockNumber, "error", err)
		}

		// Get all blobs corresponding to this timestamp
		blobSidecars, err := d.l1BeaconClient.GetBlobSidecarsEnhanced(d.ctx, L1BlockRef{
			Time: header.Time,
		}, indexedBlobHashes)
		if err != nil {
			return nil, fmt.Errorf("failed to get blobs, continuing processing:%v", err)
		}
		if len(blobSidecars) > 0 {
			// Index beacon sidecars by their KZG-derived versioned hash so we
			// can assemble the local sidecar in the exact order the L1 tx
			// declared its blobs. Multi-blob batches are decoded by
			// concatenating blob bodies in tx order; any reordering here
			// would corrupt the resulting zstd stream. The map key is
			// derived from the beacon-supplied commitment; verifyBlob below
			// re-derives the same hash from the actual blob bytes, so a
			// malicious beacon cannot forge an entry by lying about the
			// commitment.
			byHash := make(map[common.Hash]*BlobSidecar, len(blobSidecars))
			for _, sidecar := range blobSidecars {
				var commitment kzg4844.Commitment
				copy(commitment[:], sidecar.KZGCommitment[:])
				byHash[KZGToVersionedHash(commitment)] = sidecar
			}

			// Downstream (ParseBatch) only consumes Sidecar.Blobs and
			// Sidecar.Commitments; Proofs is intentionally left empty to
			// avoid an extra ~O(n) KZG op per blob per batch on every
			// sync. If a future consumer needs Proofs, compute them
			// lazily there or call kzg4844.ComputeBlobProof here.
			var blobTxSidecar eth.BlobTxSidecar
			for i, expectedHash := range blobHashes {
				sidecar, ok := byHash[expectedHash]
				if !ok {
					return nil, fmt.Errorf("blob %d (hash=%s) not found in beacon sidecars", i, expectedHash.Hex())
				}

				b, err := hexutil.Decode(sidecar.Blob)
				if err != nil {
					return nil, fmt.Errorf("failed to decode blob %d: %w", i, err)
				}
				// Reject malformed beacon responses up front. copy(blob[:], b)
				// silently:
				//   - zero-pads when len(b) < BlobSize (tail of the
				//     zero-initialized array stays zero)
				//   - truncates when len(b) > BlobSize (extra bytes dropped)
				// Either case would otherwise surface later as a confusing
				// blob-hash mismatch instead of a clear length error.
				if len(b) != BlobSize {
					return nil, fmt.Errorf("blob %d: unexpected length %d (want %d, hash=%s)", i, len(b), BlobSize, expectedHash.Hex())
				}
				var blob Blob
				copy(blob[:], b)

				if err := verifyBlob(&blob, expectedHash); err != nil {
					return nil, fmt.Errorf("blob %d: %w", i, err)
				}

				var commitment kzg4844.Commitment
				copy(commitment[:], sidecar.KZGCommitment[:])

				d.logger.Info("Matched blob", "txOrder", i, "beaconIndex", sidecar.Index, "hash", expectedHash.Hex())
				blobTxSidecar.Blobs = append(blobTxSidecar.Blobs, *blob.KZGBlob())
				blobTxSidecar.Commitments = append(blobTxSidecar.Commitments, commitment)
			}

			d.logger.Info("Blob matching results", "matched", len(blobTxSidecar.Blobs), "expected", len(blobHashes))
			batch.Sidecar = blobTxSidecar
		} else {
			return nil, fmt.Errorf("not matched blob,txHash:%v,blockNumber:%v", txHash, blockNumber)
		}
	}

	// Get L2 height
	l2Height, err := d.l2Client.BlockNumber(d.ctx)
	if err != nil {
		return nil, fmt.Errorf("query l2 block number error:%v", err)
	}
	rollupData, err := d.parseBatch(batch, l2Height)
	if err != nil {
		d.logger.Error("parse batch failed", "txNonce", tx.Nonce(), "txHash", txHash,
			"l1BlockNumber", blockNumber)
		return rollupData, fmt.Errorf("parse batch error:%v", err)
	}
	rollupData.l1BlockNumber = blockNumber
	rollupData.txHash = txHash
	rollupData.nonce = tx.Nonce()
	rollupData.blobHashes = tx.BlobHashes()
	return rollupData, nil
}

func (d *Derivation) UnPackData(data []byte) (geth.RPCRollupBatch, error) {
	var batch geth.RPCRollupBatch
	if bytes.Equal(d.beforeMoveBlockCtxABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.beforeMoveBlockCtxABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version           uint8     "json:\"version\""
			ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
			BlockContexts     []uint8   "json:\"blockContexts\""
			PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot     [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			BlockContexts:     rollupBatchData.BlockContexts,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else if bytes.Equal(d.legacyRollupABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.legacyRollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version                uint8     "json:\"version\""
			ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
			BlockContexts          []uint8   "json:\"blockContexts\""
			SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
			PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot          [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			BlockContexts:     rollupBatchData.BlockContexts,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else if bytes.Equal(d.rollupABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.rollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version           uint8     "json:\"version\""
			ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
			LastBlockNumber   uint64    "json:\"lastBlockNumber\""
			NumL1Messages     uint16    "json:\"numL1Messages\""
			PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot     [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			LastBlockNumber:   rollupBatchData.LastBlockNumber,
			NumL1Messages:     rollupBatchData.NumL1Messages,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else if bytes.Equal(d.rollupABI.Methods["commitBatchWithProof"].ID, data[:4]) {
		args, err := d.rollupABI.Methods["commitBatchWithProof"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("commitBatchWithProof Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version           uint8     "json:\"version\""
			ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
			LastBlockNumber   uint64    "json:\"lastBlockNumber\""
			NumL1Messages     uint16    "json:\"numL1Messages\""
			PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot     [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			LastBlockNumber:   rollupBatchData.LastBlockNumber,
			NumL1Messages:     rollupBatchData.NumL1Messages,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else {
		return batch, types.ErrNotCommitBatchTx
	}
	return batch, nil
}

func (d *Derivation) parseBatch(batch geth.RPCRollupBatch, l2Height uint64) (*BatchInfo, error) {
	batchInfo := new(BatchInfo)
	if err := batchInfo.ParseBatch(batch); err != nil {
		return nil, fmt.Errorf("parse batch error:%v", err)
	}
	if err := d.handleL1Message(batchInfo, batchInfo.parentTotalL1MessagePopped, l2Height); err != nil {
		return nil, fmt.Errorf("handle l1 message error:%v", err)
	}
	return batchInfo, nil
}

func (d *Derivation) handleL1Message(rollupData *BatchInfo, parentTotalL1MessagePopped, l2Height uint64) error {
	totalL1MessagePopped := parentTotalL1MessagePopped
	for bIndex, block := range rollupData.blockContexts {
		// This may happen to nodes started from snapshot, in which case we will no longer handle L1Msg
		if block.Number <= l2Height {
			totalL1MessagePopped += uint64(block.l1MsgNum)
			continue
		}
		var l1Transactions []*eth.Transaction
		l1Messages, err := d.getL1Message(totalL1MessagePopped, uint64(block.l1MsgNum))
		if err != nil {
			return fmt.Errorf("get l1 message error:%v", err)
		}
		if len(l1Messages) != int(block.l1MsgNum) {
			return fmt.Errorf("invalid l1 msg num,expect %v,have %v", block.l1MsgNum, len(l1Messages))
		}
		totalL1MessagePopped += uint64(block.l1MsgNum)
		if len(l1Messages) > 0 {
			for _, l1Message := range l1Messages {
				transaction := eth.NewTx(&l1Message.L1MessageTx)
				l1Transactions = append(l1Transactions, transaction)
			}
		}
		rollupData.blockContexts[bIndex].SafeL2Data.Transactions = append(encodeTransactions(l1Transactions), rollupData.blockContexts[bIndex].SafeL2Data.Transactions...)
	}

	return nil
}

func (d *Derivation) getL1Message(l1MessagePopped, l1MsgNum uint64) ([]types.L1Message, error) {
	if l1MsgNum == 0 {
		return nil, nil
	}
	return d.syncer.ReadL1MessagesInRange(l1MessagePopped, l1MessagePopped+l1MsgNum-1), nil
}

func (d *Derivation) derive(rollupData *BatchInfo) (*eth.Header, error) {
	var lastHeader *eth.Header
	for _, blockData := range rollupData.blockContexts {
		latestBlockNumber, err := d.l2Client.BlockNumber(context.Background())
		if err != nil {
			return nil, fmt.Errorf("get derivation geth block number error:%v", err)
		}
		if blockData.SafeL2Data.Number <= latestBlockNumber {
			d.logger.Info("new L2 Data block number less than latestBlockNumber", "safeL2DataNumber", blockData.SafeL2Data.Number, "latestBlockNumber", latestBlockNumber)
			lastHeader, err = d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(blockData.SafeL2Data.Number)))
			if err != nil {
				return nil, fmt.Errorf("query header by number error:%v", err)
			}
			continue
		}
		err = func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
			defer cancel()
			lastHeader, err = d.l2Client.NewSafeL2Block(ctx, blockData.SafeL2Data)
			if err != nil {
				d.logger.Error("new l2 block failed", "latestBlockNumber", latestBlockNumber, "error", err)
				return err
			}
			return nil
		}()
		if err != nil {
			return nil, fmt.Errorf("derivation error:%v", err)
		}
		d.logger.Info("new l2 block success", "blockNumber", blockData.Number)
	}

	return lastHeader, nil
}

// deriveForce writes the batch's blocks via NewSafeL2Block.
//
// skipNumber lets one implementation serve two SPEC-005 §4.3 Path B scenarios:
//
//   - skipNumber == 0 (scenario B, self-heal): every block is written; EL
//     SetCanonical reorgs the local fork onto the L1-canonical chain.
//   - skipNumber > 0 (scenario C, fill-gap): blocks with Number ≤ skipNumber
//     are skipped (already present locally, presumed valid via P2P), only
//     the missing tail is appended; no reorg of existing blocks.
//
// In both cases the parent of the first block we actually write must exist
// locally. For scenario B that's batch.firstBlockNumber-1 (above safe head).
// For scenario C with skipNumber == localLatestL2 that's localLatestL2 itself
// (necessarily ≥ firstBlockNumber-1 once skipNumber covers everything below).
func (d *Derivation) deriveForce(rollupData *BatchInfo, skipNumber uint64) (*eth.Header, error) {
	firstNum := rollupData.firstBlockNumber
	if firstNum == 0 {
		return nil, fmt.Errorf("invalid firstBlockNumber 0 for batch %d", rollupData.batchIndex)
	}

	// Anchor: parent of the first block we will WRITE must exist locally.
	// scenario B (skipNumber==0): firstNum-1.
	// scenario C: max(firstNum-1, skipNumber).
	parentNum := firstNum - 1
	if skipNumber > parentNum {
		parentNum = skipNumber
	}
	lastHeader, err := d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(parentNum)))
	if err != nil {
		return nil, fmt.Errorf("read parent header at %d: %w", parentNum, err)
	}
	if lastHeader == nil {
		return nil, fmt.Errorf("parent header at %d missing", parentNum)
	}

	for _, blockData := range rollupData.blockContexts {
		// Skip blocks already present locally (scenario C). For scenario B
		// skipNumber == 0 means this branch is never taken.
		if blockData.SafeL2Data.Number <= skipNumber {
			continue
		}

		// Pin the parent so SetCanonical reorgs from the local fork to the
		// L1-canonical chain. NewSafeL2Block executes the block internally
		// and fills the header with the resulting state/receipt roots —
		// the caller only knows block contents (txs + timestamp), not the
		// post-execution roots, so this is the right API for the rewrite.
		parentHash := lastHeader.Hash()
		safeData := *blockData.SafeL2Data
		safeData.ParentHash = &parentHash

		err = func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
			defer cancel()
			next, err := d.l2Client.NewSafeL2Block(ctx, &safeData)
			if err != nil {
				d.logger.Error("NewSafeL2Block failed",
					"batchIndex", rollupData.batchIndex,
					"blockNumber", safeData.Number,
					"parent", parentHash.Hex(),
					"error", err,
				)
				return err
			}
			if next == nil {
				return fmt.Errorf("header at %d missing after NewSafeL2Block", safeData.Number)
			}
			lastHeader = next
			return nil
		}()
		if err != nil {
			return nil, fmt.Errorf("apply block %d: %w", safeData.Number, err)
		}

		d.logger.Info("block written via NewSafeL2Block",
			"batchIndex", rollupData.batchIndex,
			"blockNumber", safeData.Number,
			"hash", lastHeader.Hash().Hex(),
		)
	}
	return lastHeader, nil
}

func (d *Derivation) getLatestConfirmedBlockNumber(ctx context.Context) (uint64, error) {
	return nodecommon.GetLatestConfirmedBlockNumber(ctx, d.l1Client, d.confirmations)
}
