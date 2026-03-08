package derivation

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/ethclient"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/node/types"
	"morph-l2/node/validator"
)

var (
	// RollupEventTopic is the CommitBatch event signature string.
	RollupEventTopic = "CommitBatch(uint256,bytes32)"
	// RollupEventTopicHash is the keccak256 hash of RollupEventTopic.
	// Used when filtering L1 logs for batch submissions.
	RollupEventTopicHash = crypto.Keccak256Hash([]byte(RollupEventTopic))
)

// BatchVerifyL2Client is the minimal L2 read interface required for batch verification.
// *types.RetryableClient satisfies this interface.
//
// HeaderByNumber is used for lightweight checks (root verification, PrevStateRoot).
// BlockByNumber is used for block-context verification where transaction data is needed.
type BatchVerifyL2Client interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*eth.Header, error)
	// BlockByNumber returns the full block including all transactions.
	// Used in verifyBlockContextHeaders to check NumTxs and NumL1Msgs.
	BlockByNumber(ctx context.Context, number *big.Int) (*eth.Block, error)
}

// BatchBlockContext holds per-block fields decoded from the BlockContexts calldata field.
// Only present for v1+ batches where block contexts are encoded directly in calldata.
// Blob-based (v0/legacy) batches cannot provide this without blob decoding.
//
// 60-byte wire layout (per block):
//
//	Number(8) | Timestamp(8) | BaseFee(32) | GasLimit(8) | NumTxs(2) | NumL1Msgs(2)
type BatchBlockContext struct {
	Number    uint64
	Timestamp uint64
	GasLimit  uint64
	BaseFee   *big.Int
	NumTxs    uint16 // total transactions in the block (L2 + L1 messages)
	NumL1Msgs uint16 // L1 message transactions in the block
}

// BatchRoots contains the key roots and block metadata parsed from a CommitBatch L1 transaction.
// All fields are populated via calldata parsing only — no blob decoding required.
//
// BlockContexts is non-nil only for v1+ batches where block contexts are encoded in calldata.
// For blob-based (v0/legacy) batches, BlockContexts is nil and only root verification is possible.
type BatchRoots struct {
	BatchIndex     uint64
	FirstBlockNum  uint64 // first L2 block number in the batch; 0 if unavailable (blob batches)
	LastBlockNum   uint64
	PrevStateRoot  common.Hash // stateRoot of the block just before this batch; zero if unavailable
	PostStateRoot  common.Hash
	WithdrawalRoot common.Hash
	NumL1Messages  uint16              // total L1 messages in the batch
	BlockContexts  []BatchBlockContext // per-block metadata; nil for blob-based batches
}

// BatchVerifier encapsulates the stateless logic for fetching and verifying L1 batch data.
// It exposes callable methods with no internal goroutines or scheduling.
// Scheduling is owned by the caller (e.g. BlockTagService).
type BatchVerifier struct {
	l1Client              *ethclient.Client
	l2EthClient           *ethclient.Client // owns the connection used by L2ToL1MessagePasser; closed via Close()
	l1BeaconClient        *L1BeaconClient   // nil if BeaconRpc not configured
	rollup                *bindings.Rollup
	rollupABI             *abi.ABI
	legacyRollupABI       *abi.ABI
	beforeMoveBlockCtxABI *abi.ABI
	RollupContractAddress common.Address

	// L2 contract for withdrawal root verification
	L2ToL1MessagePasser *bindings.L2ToL1MessagePasser

	// Optional: triggers challenge on state mismatch when enabled
	validator *validator.Validator

	// Upgrade transition config (fetched from geth at startup)
	baseHeight uint64
	switchTime uint64
	useZktrie  bool

	logger tmlog.Logger
}

// NewBatchVerifier creates a BatchVerifier using a subset of derivation Config.
// It connects to L1 and L2 but does not start any background goroutines.
// Call Close() when the BatchVerifier is no longer needed to release connections.
func NewBatchVerifier(ctx context.Context, cfg *Config, vt *validator.Validator, logger tmlog.Logger) (*BatchVerifier, error) {
	l1Client, err := ethclient.Dial(cfg.L1.Addr)
	if err != nil {
		return nil, fmt.Errorf("dial l1 node error: %w", err)
	}

	l2EthClient, err := ethclient.Dial(cfg.L2.EthAddr)
	if err != nil {
		l1Client.Close()
		return nil, fmt.Errorf("dial l2 eth node error: %w", err)
	}

	rollup, err := bindings.NewRollup(cfg.RollupContractAddress, l1Client)
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("create rollup binding error: %w", err)
	}

	msgPasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, l2EthClient)
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("create L2ToL1MessagePasser binding error: %w", err)
	}

	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("get rollup ABI: %w", err)
	}
	legacyRollupAbi, err := types.LegacyRollupMetaData.GetAbi()
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("get legacy rollup ABI: %w", err)
	}
	beforeMoveBlockCtxAbi, err := types.BeforeMoveBlockCtxABI.GetAbi()
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("get beforeMoveBlockCtx ABI: %w", err)
	}

	// Fetch upgrade transition config from geth (retries until geth is ready or ctx is cancelled)
	gethCfg, err := types.FetchGethConfigWithRetry(ctx, cfg.L2.EthAddr, logger)
	if err != nil {
		l1Client.Close()
		l2EthClient.Close()
		return nil, fmt.Errorf("failed to fetch geth config: %w", err)
	}
	logger.Info("BatchVerifier: geth config fetched",
		"switchTime", gethCfg.SwitchTime,
		"useZktrie", gethCfg.UseZktrie,
	)

	// L1 Beacon client for blob fetching (optional: only required for FetchBatchData)
	var l1BeaconClient *L1BeaconClient
	if cfg.BeaconRpc != "" {
		baseHttp := NewBasicHTTPClient(cfg.BeaconRpc, logger)
		l1BeaconClient = NewL1BeaconClient(baseHttp)
		logger.Info("BatchVerifier: L1 beacon client configured", "beaconRpc", cfg.BeaconRpc)
	} else {
		logger.Info("BatchVerifier: BeaconRpc not set, blob fetching disabled")
	}

	return &BatchVerifier{
		l1Client:              l1Client,
		l2EthClient:           l2EthClient,
		l1BeaconClient:        l1BeaconClient,
		rollup:                rollup,
		rollupABI:             rollupAbi,
		legacyRollupABI:       legacyRollupAbi,
		beforeMoveBlockCtxABI: beforeMoveBlockCtxAbi,
		RollupContractAddress: cfg.RollupContractAddress,
		L2ToL1MessagePasser:   msgPasser,
		validator:             vt,
		baseHeight:            cfg.BaseHeight,
		switchTime:            gethCfg.SwitchTime,
		useZktrie:             gethCfg.UseZktrie,
		logger:                logger.With("module", "batch_verifier"),
	}, nil
}

// Close releases the L1 and L2 RPC connections held by the BatchVerifier.
func (bv *BatchVerifier) Close() {
	bv.l1Client.Close()
	bv.l2EthClient.Close()
}

// FetchBatchRoots fetches state roots and block metadata from a CommitBatch L1 transaction.
// Only calldata is parsed — no blob fetching required.
//
// Populated fields depend on batch version:
//   - v0/legacy: only PostStateRoot, WithdrawalRoot, LastBlockNum (from BatchDataStore fallback)
//   - v1+:       all fields including PrevStateRoot, FirstBlockNum, NumL1Messages, BlockContexts
//
// For v0/legacy batches where LastBlockNumber is absent from calldata,
// it falls back to querying BatchDataStore on-chain.
func (bv *BatchVerifier) FetchBatchRoots(ctx context.Context, txHash common.Hash, batchIndex uint64) (*BatchRoots, error) {
	tx, pending, err := bv.l1Client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("get transaction %s: %w", txHash.Hex(), err)
	}
	if pending {
		return nil, fmt.Errorf("transaction %s is still pending", txHash.Hex())
	}

	batch, err := unpackCalldataWithABIs(bv.rollupABI, bv.legacyRollupABI, bv.beforeMoveBlockCtxABI, tx.Data())
	if err != nil {
		return nil, fmt.Errorf("unpack calldata for tx %s: %w", txHash.Hex(), err)
	}

	// Derive batchIndex from parentBatchHeader embedded in calldata
	parentBatchHeader := types.BatchHeaderBytes(batch.ParentBatchHeader)
	parentBatchIndex, err := parentBatchHeader.BatchIndex()
	if err != nil {
		return nil, fmt.Errorf("decode parent batch index: %w", err)
	}

	roots := &BatchRoots{
		BatchIndex:     parentBatchIndex + 1,
		LastBlockNum:   batch.LastBlockNumber,
		PrevStateRoot:  batch.PrevStateRoot,
		PostStateRoot:  batch.PostStateRoot,
		WithdrawalRoot: batch.WithdrawRoot,
		NumL1Messages:  batch.NumL1Messages,
	}

	// v0/legacy batches do not encode LastBlockNumber in calldata (it's in the blob).
	// Fall back to the on-chain BatchDataStore.
	if roots.LastBlockNum == 0 {
		batchData, err := bv.rollup.BatchDataStore(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(batchIndex))
		if err != nil {
			return nil, fmt.Errorf("query BatchDataStore for batchIndex %d: %w", batchIndex, err)
		}
		roots.LastBlockNum = batchData.BlockNumber.Uint64()
	}

	// Parse per-block contexts when encoded in calldata (v1+ batches).
	// Format: [numBlocks: 2 bytes][block0: 60 bytes][block1: 60 bytes]...
	// Each 60-byte context: Number(8)+Timestamp(8)+BaseFee(32)+GasLimit(8)+numTxs(2)+numL1Msgs(2)
	if len(batch.BlockContexts) >= 2 {
		blockContexts, firstBlockNum, err := parseBlockContexts(batch.BlockContexts)
		if err != nil {
			// Non-fatal: log and continue without block context verification
			bv.logger.Info("FetchBatchRoots: failed to parse BlockContexts, skipping block metadata",
				"batchIndex", roots.BatchIndex, "error", err)
		} else {
			roots.BlockContexts = blockContexts
			roots.FirstBlockNum = firstBlockNum
		}
	}

	return roots, nil
}

// parseBlockContexts decodes the BlockContexts calldata field into individual BatchBlockContext entries.
// Returns the slice of contexts and the first block's Number.
// Format: [numBlocks: 2 bytes] followed by numBlocks × 60-byte context records.
func parseBlockContexts(data []byte) ([]BatchBlockContext, uint64, error) {
	if len(data) < 2 {
		return nil, 0, fmt.Errorf("BlockContexts too short: %d bytes", len(data))
	}

	numBlocks := int(data[0])<<8 | int(data[1])
	if numBlocks == 0 {
		return nil, 0, fmt.Errorf("BlockContexts: numBlocks is zero")
	}
	expectedLen := 2 + numBlocks*60
	if len(data) < expectedLen {
		return nil, 0, fmt.Errorf("BlockContexts: need %d bytes for %d blocks, got %d", expectedLen, numBlocks, len(data))
	}

	contexts := make([]BatchBlockContext, numBlocks)
	for i := 0; i < numBlocks; i++ {
		off := 2 + i*60
		raw := data[off : off+60]

		var wb types.WrappedBlock
		txsNum, l1MsgNum, err := wb.DecodeBlockContext(raw)
		if err != nil {
			return nil, 0, fmt.Errorf("decode block context %d: %w", i, err)
		}
		contexts[i] = BatchBlockContext{
			Number:    wb.Number,
			Timestamp: wb.Timestamp,
			GasLimit:  wb.GasLimit,
			BaseFee:   wb.BaseFee,
			NumTxs:    txsNum,
			NumL1Msgs: l1MsgNum,
		}
	}

	return contexts, contexts[0].Number, nil
}

// FetchBatchData fetches a complete batch from L1, including blob decoding.
// It returns a *BatchInfo containing per-block metadata and the decoded L2 user transactions
// for each block (L1 messages are NOT included — they are injected separately by the sequencer).
//
// This extends FetchBatchRoots: where FetchBatchRoots only parses calldata,
// FetchBatchData also fetches blobs from the L1 beacon chain and decodes them via BatchInfo.ParseBatch.
//
// Requires l1BeaconClient to be configured (cfg.BeaconRpc must be set).
// For batches without blobs (no blobHashes on the L1 tx), ParseBatch still works
// with an empty Sidecar; blob-free batches encode txs directly in calldata.
func (bv *BatchVerifier) FetchBatchData(ctx context.Context, l1TxHash common.Hash) (*BatchInfo, error) {
	if bv.l1BeaconClient == nil {
		return nil, fmt.Errorf("FetchBatchData: l1BeaconClient not configured (set BeaconRpc in config)")
	}

	tx, pending, err := bv.l1Client.TransactionByHash(ctx, l1TxHash)
	if err != nil {
		return nil, fmt.Errorf("get L1 transaction %s: %w", l1TxHash.Hex(), err)
	}
	if pending {
		return nil, fmt.Errorf("L1 transaction %s is still pending", l1TxHash.Hex())
	}

	batch, err := unpackCalldataWithABIs(bv.rollupABI, bv.legacyRollupABI, bv.beforeMoveBlockCtxABI, tx.Data())
	if err != nil {
		return nil, fmt.Errorf("unpack calldata for tx %s: %w", l1TxHash.Hex(), err)
	}

	// If the L1 tx carries blob hashes, fetch and attach the blob sidecar.
	blobHashes := tx.BlobHashes()
	if len(blobHashes) > 0 {
		bv.logger.Info("FetchBatchData: tx has blobs, fetching from beacon",
			"txHash", l1TxHash.Hex(), "blobCount", len(blobHashes))

		// Get the L1 block to build indexed blob hashes (position of each blob in the block's sidecar).
		receipt, err := bv.l1Client.TransactionReceipt(ctx, l1TxHash)
		if err != nil {
			return nil, fmt.Errorf("get L1 tx receipt %s: %w", l1TxHash.Hex(), err)
		}
		l1BlockNum := receipt.BlockNumber

		l1Block, err := bv.l1Client.BlockByNumber(ctx, l1BlockNum)
		if err != nil {
			return nil, fmt.Errorf("get L1 block %d: %w", l1BlockNum.Uint64(), err)
		}
		indexedBlobHashes := dataAndHashesFromTxs(l1Block.Transactions(), tx)

		// Beacon chain lookup uses the L1 block timestamp to derive the slot number.
		blobSidecars, err := bv.l1BeaconClient.GetBlobSidecarsEnhanced(ctx, L1BlockRef{
			Time: l1Block.Time(),
		}, indexedBlobHashes)
		if err != nil {
			return nil, fmt.Errorf("fetch blob sidecars for tx %s (L1 block %d): %w",
				l1TxHash.Hex(), l1BlockNum.Uint64(), err)
		}

		// Match each blob hash to its sidecar and assemble the BlobTxSidecar.
		var blobTxSidecar eth.BlobTxSidecar
		matchedCount := 0
		for _, sidecar := range blobSidecars {
			var commitment kzg4844.Commitment
			copy(commitment[:], sidecar.KZGCommitment[:])
			versionedHash := KZGToVersionedHash(commitment)

			for _, expectedHash := range blobHashes {
				if !bytes.Equal(versionedHash[:], expectedHash[:]) {
					continue
				}
				matchedCount++
				b, err := hexutil.Decode(sidecar.Blob)
				if err != nil {
					return nil, fmt.Errorf("decode blob hex for tx %s: %w", l1TxHash.Hex(), err)
				}
				var blob Blob
				copy(blob[:], b)
				blobTxSidecar.Blobs = append(blobTxSidecar.Blobs, *blob.KZGBlob())
				blobTxSidecar.Commitments = append(blobTxSidecar.Commitments, commitment)
				blobTxSidecar.Proofs = append(blobTxSidecar.Proofs, kzg4844.Proof(sidecar.KZGProof))
				break
			}
		}
		if matchedCount == 0 {
			return nil, fmt.Errorf("FetchBatchData: no matching blob found for tx %s", l1TxHash.Hex())
		}
		bv.logger.Info("FetchBatchData: blobs matched", "matched", matchedCount, "expected", len(blobHashes))
		batch.Sidecar = blobTxSidecar
	}

	batchInfo := new(BatchInfo)
	if err := batchInfo.ParseBatch(batch); err != nil {
		return nil, fmt.Errorf("parse batch for tx %s: %w", l1TxHash.Hex(), err)
	}
	return batchInfo, nil
}

// VerifyBatch validates the L2 state for a completed batch against the L1 commitment.
//
// Verification steps (in order):
//  1. PostStateRoot: l2Block(LastBlockNum).stateRoot must equal L1-committed PostStateRoot.
//  2. WithdrawalRoot: L2ToL1MessagePasser.MessageRoot at LastBlockNum must equal WithdrawalRoot.
//  3. PrevStateRoot: l2Block(FirstBlockNum-1).stateRoot must equal L1-committed PrevStateRoot
//     (ensures the batch was applied to the correct prior state; skipped if FirstBlockNum unavailable).
//  4. BlockContexts: for each block context decoded from calldata (v1+ batches), verifies that
//     the actual L2 block header matches the committed Number, Timestamp, GasLimit, BaseFee,
//     NumTxs, and NumL1Msgs.
//  5. L2 user transactions: if batchData (from FetchBatchData) is provided, for each block the
//     decoded L2 user transactions from the blob are compared against the actual L2 block
//     transactions (excluding L1 message txs). This detects tx content divergence directly.
//
// Prerequisites:
//   - The L2 blocks must already exist locally (produced via Tendermint P2P).
//   - BatchRoots must have been obtained via FetchBatchRoots.
//   - batchData may be nil; if non-nil it must be the result of FetchBatchData for the same batch.
//
// On mismatch for step 1 or 2:
//   - If a Validator is configured and challenge is enabled, ChallengeState is triggered.
//
// Steps 3–5 return errors but do not trigger challenge (metadata/content checks, not fraud proofs).
//
// Validation is silently skipped during upgrade transitions (switchTime window).
func (bv *BatchVerifier) VerifyBatch(ctx context.Context, l2Client BatchVerifyL2Client, roots *BatchRoots, batchData *BatchInfo) error {
	// Blocks at or below the snapshot/genesis base height have no meaningful roots to verify
	if roots.LastBlockNum <= bv.baseHeight {
		bv.logger.Info("skipping verification: block at or below base height",
			"lastBlockNum", roots.LastBlockNum,
			"baseHeight", bv.baseHeight,
		)
		return nil
	}

	// ── Step 1 & 2: PostStateRoot + WithdrawalRoot ──────────────────────────────
	l2Header, err := l2Client.HeaderByNumber(ctx, new(big.Int).SetUint64(roots.LastBlockNum))
	if err != nil {
		return fmt.Errorf("get L2 header for block %d: %w", roots.LastBlockNum, err)
	}

	withdrawalRoot, err := bv.L2ToL1MessagePasser.MessageRoot(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: l2Header.Number,
	})
	if err != nil {
		return fmt.Errorf("get withdrawal root at block %d: %w", roots.LastBlockNum, err)
	}

	rootMismatch := l2Header.Root != roots.PostStateRoot
	withdrawalMismatch := !bytes.Equal(withdrawalRoot[:], roots.WithdrawalRoot.Bytes())

	if rootMismatch || withdrawalMismatch {
		// During an upgrade transition (ZK→MPT switch), skip to avoid false positives
		if bv.shouldSkipValidation(l2Header.Time) {
			bv.logger.Info("root validation skipped during upgrade transition",
				"batchIndex", roots.BatchIndex,
				"l1StateRoot", roots.PostStateRoot.Hex(),
				"l2StateRoot", l2Header.Root.Hex(),
				"blockTimestamp", l2Header.Time,
				"switchTime", bv.switchTime,
				"useZktrie", bv.useZktrie,
			)
			return nil
		}

		bv.logger.Error("batch verification failed: root mismatch",
			"batchIndex", roots.BatchIndex,
			"lastBlockNum", roots.LastBlockNum,
			"l1StateRoot", roots.PostStateRoot.Hex(),
			"l2StateRoot", l2Header.Root.Hex(),
			"l1WithdrawalRoot", roots.WithdrawalRoot.Hex(),
			"l2WithdrawalRoot", common.BytesToHash(withdrawalRoot[:]).Hex(),
			"rootMismatch", rootMismatch,
			"withdrawalMismatch", withdrawalMismatch,
		)

		// Trigger challenge if validator is configured and enabled
		if bv.validator != nil && bv.validator.ChallengeEnable() {
			if err := bv.validator.ChallengeState(roots.BatchIndex); err != nil {
				bv.logger.Error("challenge state failed", "batchIndex", roots.BatchIndex, "error", err)
			}
		}

		return fmt.Errorf("state mismatch for batch %d (block %d): stateRoot[L1=%s L2=%s] withdrawalRoot[L1=%s L2=%s]",
			roots.BatchIndex,
			roots.LastBlockNum,
			roots.PostStateRoot.Hex(),
			l2Header.Root.Hex(),
			roots.WithdrawalRoot.Hex(),
			common.BytesToHash(withdrawalRoot[:]).Hex(),
		)
	}

	// ── Step 3: PrevStateRoot (batch continuity) ────────────────────────────────
	// Verifies that this batch was applied on top of the correct prior L2 state.
	// Only checked when FirstBlockNum is known (v1+ batches) and is above the base height.
	if roots.PrevStateRoot != (common.Hash{}) && roots.FirstBlockNum > bv.baseHeight {
		prevBlockNum := roots.FirstBlockNum - 1
		prevHeader, err := l2Client.HeaderByNumber(ctx, new(big.Int).SetUint64(prevBlockNum))
		if err != nil {
			// Non-fatal: the previous block might not be available yet (e.g. syncing)
			bv.logger.Info("VerifyBatch: could not fetch prev block header for PrevStateRoot check",
				"batchIndex", roots.BatchIndex,
				"prevBlockNum", prevBlockNum,
				"error", err,
			)
		} else if prevHeader.Root != roots.PrevStateRoot {
			return fmt.Errorf("PrevStateRoot mismatch for batch %d (block %d): L1=%s L2=%s",
				roots.BatchIndex,
				prevBlockNum,
				roots.PrevStateRoot.Hex(),
				prevHeader.Root.Hex(),
			)
		}
	}

	// ── Step 4: Per-block metadata (v1+ batches with BlockContexts in calldata) ─
	// For blob-based batches BlockContexts is nil and this step is skipped.
	if len(roots.BlockContexts) > 0 {
		if err := bv.verifyBlockContextHeaders(ctx, l2Client, roots); err != nil {
			return err
		}
	}

	// ── Step 5: L2 user transaction content (requires blob data from FetchBatchData) ──
	// Compares decoded L2 user transactions from the blob with the actual L2 block transactions.
	// L1 message transactions (type 0x7E) are excluded from both sides of the comparison
	// because they are injected by the sequencer from the L1 queue and are NOT encoded in the blob.
	//
	// Note: PostStateRoot already guarantees correctness by implication, but this step provides
	// explicit per-transaction content validation and enables earlier, more targeted diagnostics.
	if batchData != nil {
		if err := bv.verifyBatchTransactions(ctx, l2Client, batchData); err != nil {
			return err
		}
	}

	return nil
}

// verifyBatchTransactions compares the L2 user transactions decoded from the blob
// against the actual transactions in each L2 block.
//
// L1 message transactions (L1MessageTxType = 0x7E) are excluded from the L2 block side
// because they are not encoded in the blob — they are fetched from the L1 message queue
// and injected at the front of the block by the sequencer at execution time.
//
// Per-block checks:
//   - User tx count: len(blobTxs) == len(l2Block non-L1-msg txs)
//   - User tx content: each tx's binary encoding must match byte-for-byte
func (bv *BatchVerifier) verifyBatchTransactions(ctx context.Context, l2Client BatchVerifyL2Client, batchData *BatchInfo) error {
	for _, bc := range batchData.blockContexts {
		if bc.Number <= bv.baseHeight {
			continue // skip blocks at or below snapshot height
		}

		block, err := l2Client.BlockByNumber(ctx, new(big.Int).SetUint64(bc.Number))
		if err != nil {
			return fmt.Errorf("block %d: get L2 block for tx verification: %w", bc.Number, err)
		}

		// Partition L2 block transactions into L1 message txs and L2 user txs.
		var l2UserTxs eth.Transactions
		var l1MsgCount uint16
		for _, tx := range block.Transactions() {
			if tx.IsL1MessageTx() {
				l1MsgCount++
			} else {
				l2UserTxs = append(l2UserTxs, tx)
			}
		}

		// Verify L1 message count against the blob-decoded block context (cross-validates
		// the blob-side l1MsgNum with the actual L2 block; complements Step 4's calldata check).
		if l1MsgCount != uint16(bc.l1MsgNum) {
			return fmt.Errorf("block %d: L1 message tx count mismatch: blob=%d L2=%d",
				bc.Number, bc.l1MsgNum, l1MsgCount)
		}

		// bc.SafeL2Data.Transactions holds the blob-decoded L2 user txs as RLP-encoded bytes.
		blobTxs := bc.SafeL2Data.Transactions
		if len(l2UserTxs) != len(blobTxs) {
			return fmt.Errorf("block %d: L2 user tx count mismatch: blob=%d L2=%d",
				bc.Number, len(blobTxs), len(l2UserTxs))
		}

		for i, l2Tx := range l2UserTxs {
			encodedL2Tx, err := l2Tx.MarshalBinary()
			if err != nil {
				return fmt.Errorf("block %d: tx[%d] marshal: %w", bc.Number, i, err)
			}
			if !bytes.Equal(encodedL2Tx, blobTxs[i]) {
				return fmt.Errorf("block %d: tx[%d] content mismatch: hash=%s",
					bc.Number, i, l2Tx.Hash().Hex())
			}
		}
	}
	return nil
}

// verifyBlockContextHeaders checks that each block context decoded from L1 calldata
// matches the actual L2 block.
//
// Per-block checks:
//   - Number, Timestamp, GasLimit, BaseFee — from block header
//   - NumTxs  — total transaction count (L2 txs + L1 message txs)
//   - NumL1Msgs — count of L1MessageTxType (0x7E) transactions; these are always the first
//     NumL1Msgs entries in the block, injected by the sequencer from the L1 message queue
//
// Uses BlockByNumber (one RPC per block) to obtain both header fields and transaction list,
// replacing the previous HeaderByNumber + TransactionCount two-call pattern.
//
// Note: transaction *content* and ordering are already guaranteed by PostStateRoot.
// These checks give explicit metadata consistency and more targeted early error detection.
func (bv *BatchVerifier) verifyBlockContextHeaders(ctx context.Context, l2Client BatchVerifyL2Client, roots *BatchRoots) error {
	for i, bc := range roots.BlockContexts {
		if bc.Number <= bv.baseHeight {
			continue // skip blocks at or below snapshot height
		}

		block, err := l2Client.BlockByNumber(ctx, new(big.Int).SetUint64(bc.Number))
		if err != nil {
			return fmt.Errorf("block %d (context %d): get L2 block: %w", bc.Number, i, err)
		}
		header := block.Header()

		// ── Header field checks ────────────────────────────────────────────────
		if header.Number.Uint64() != bc.Number {
			return fmt.Errorf("block %d: Number mismatch: L1=%d L2=%d",
				bc.Number, bc.Number, header.Number.Uint64())
		}
		if header.Time != bc.Timestamp {
			return fmt.Errorf("block %d: Timestamp mismatch: L1=%d L2=%d",
				bc.Number, bc.Timestamp, header.Time)
		}
		if header.GasLimit != bc.GasLimit {
			return fmt.Errorf("block %d: GasLimit mismatch: L1=%d L2=%d",
				bc.Number, bc.GasLimit, header.GasLimit)
		}
		// BaseFee: only check when L1-committed value is non-zero (pre-EIP-1559 blocks have nil)
		if bc.BaseFee != nil && bc.BaseFee.Sign() > 0 {
			if header.BaseFee == nil {
				return fmt.Errorf("block %d: BaseFee mismatch: L1=%s L2=nil",
					bc.Number, bc.BaseFee.String())
			}
			if header.BaseFee.Cmp(bc.BaseFee) != 0 {
				return fmt.Errorf("block %d: BaseFee mismatch: L1=%s L2=%s",
					bc.Number, bc.BaseFee.String(), header.BaseFee.String())
			}
		}

		// ── Transaction count checks ───────────────────────────────────────────
		txs := block.Transactions()

		// NumTxs: total transactions (L2 user txs + L1 message txs)
		if uint16(len(txs)) != bc.NumTxs {
			return fmt.Errorf("block %d: transaction count mismatch: L1=%d L2=%d",
				bc.Number, bc.NumTxs, len(txs))
		}

		// NumL1Msgs: L1 message transactions are type 0x7E (L1MessageTxType).
		// They are always injected at the front of the block before L2 user transactions.
		var l1MsgCount uint16
		for _, tx := range txs {
			if tx.IsL1MessageTx() {
				l1MsgCount++
			}
		}
		if l1MsgCount != bc.NumL1Msgs {
			return fmt.Errorf("block %d: L1 message count mismatch: L1=%d L2=%d",
				bc.Number, bc.NumL1Msgs, l1MsgCount)
		}
	}
	return nil
}

// shouldSkipValidation returns true when validation should be bypassed during the
// ZK→MPT upgrade window to avoid false-positive challenges.
// Skip conditions:
//   - Before switchTime and running MPT geth (useZktrie=false): old blocks, new geth
//   - After switchTime and running ZK geth (useZktrie=true): new blocks, old geth
func (bv *BatchVerifier) shouldSkipValidation(blockTimestamp uint64) bool {
	if bv.switchTime == 0 {
		return false
	}
	beforeSwitch := blockTimestamp < bv.switchTime
	return (beforeSwitch && !bv.useZktrie) || (!beforeSwitch && bv.useZktrie)
}

// unpackCalldataWithABIs unpacks a CommitBatch transaction's calldata across all known ABI versions.
// It is a package-level function so both Derivation and BatchVerifier can share it.
// Only calldata fields are populated; Sidecar (blob) is left empty.
func unpackCalldataWithABIs(rollupABI, legacyRollupABI, beforeMoveBlockCtxABI *abi.ABI, data []byte) (geth.RPCRollupBatch, error) {
	var batch geth.RPCRollupBatch
	if len(data) < 4 {
		return batch, fmt.Errorf("calldata too short: %d bytes", len(data))
	}

	switch {
	case bytes.Equal(beforeMoveBlockCtxABI.Methods["commitBatch"].ID, data[:4]):
		args, err := beforeMoveBlockCtxABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("beforeMoveBlockCtx commitBatch unpack: %w", err)
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

	case bytes.Equal(legacyRollupABI.Methods["commitBatch"].ID, data[:4]):
		args, err := legacyRollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("legacy commitBatch unpack: %w", err)
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

	case bytes.Equal(rollupABI.Methods["commitBatch"].ID, data[:4]):
		args, err := rollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("commitBatch unpack: %w", err)
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

	case bytes.Equal(rollupABI.Methods["commitBatchWithProof"].ID, data[:4]):
		args, err := rollupABI.Methods["commitBatchWithProof"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("commitBatchWithProof unpack: %w", err)
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

	default:
		return batch, types.ErrNotCommitBatchTx
	}

	return batch, nil
}
