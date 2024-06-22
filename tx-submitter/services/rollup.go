package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/tx-submitter/iface"
	"github.com/morph-l2/tx-submitter/metrics"
	"github.com/morph-l2/tx-submitter/utils"

	"github.com/holiman/uint256"
	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/consensus/misc/eip4844"
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/crypto/bls12381"
	"github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/params"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/tendermint/tendermint/blssignatures"
)

const (
	txSlotSize     = 32 * 1024
	txMaxSize      = 4 * txSlotSize // 128KB
	minFinalizeNum = 2              // min finalize num from contract
	rotatorBuffer  = 30
)

type Rollup struct {
	ctx     context.Context
	metrics *metrics.Metrics

	l1RpcClient  *rpc.Client
	L1Client     iface.Client
	L2Clients    []iface.L2Client
	Rollup       iface.IRollup
	L1Sequencer  iface.IL1Sequencer
	L2Submitters []iface.IL2Submitter
	L2Sequencers []iface.IL2Sequencer

	chainId    *big.Int
	privKey    *ecdsa.PrivateKey
	rollupAddr common.Address
	abi        *abi.ABI

	pendingTxs *PendingTxs

	rollupFinalizeMu sync.Mutex

	cfg utils.Config
}

func NewRollup(
	ctx context.Context,
	metrics *metrics.Metrics,
	l1RpcClient *rpc.Client,
	l1 iface.Client,
	l2 []iface.L2Client,
	rollup iface.IRollup,
	l1Sequencer iface.IL1Sequencer,
	l2Submitters []iface.IL2Submitter,
	l2Sequencers []iface.IL2Sequencer,
	chainId *big.Int,
	priKey *ecdsa.PrivateKey,
	rollupAddr common.Address,
	abi *abi.ABI,
	cfg utils.Config,
) *Rollup {

	return &Rollup{
		ctx:     ctx,
		metrics: metrics,

		l1RpcClient: l1RpcClient,
		L1Client:    l1,
		L2Clients:   l2,
		// l1
		Rollup:      rollup,
		L1Sequencer: l1Sequencer,
		// l2
		L2Submitters: l2Submitters,
		L2Sequencers: l2Sequencers,

		privKey:    priKey,
		chainId:    chainId,
		rollupAddr: rollupAddr,
		abi:        abi,

		cfg: cfg,
	}
}

func (sr *Rollup) Start() {

	// metrics
	go utils.Loop(sr.ctx, 10*time.Second, func() {

		// get balacnce of wallet
		balance, err := sr.L1Client.BalanceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey), nil)
		if err != nil {
			log.Error("get wallet balance error", "error", err)
			return
		}
		// balance to eth
		balanceEth := new(big.Rat).SetFrac(balance, big.NewInt(params.Ether))

		// parse float64 from string
		balanceEthFloat, err := strconv.ParseFloat(balanceEth.FloatString(18), 64)
		if err != nil {
			log.Warn("parse balance to float error", "error", err)
			return
		}

		sr.metrics.SetWalletBalance(balanceEthFloat)

	})

	// init pendingtxs
	sr.pendingTxs = NewPendingTxs(sr.abi.Methods["commitBatch"].ID, sr.abi.Methods["finalizeBatch"].ID)
	txs, err := utils.ParseL1Mempool(sr.l1RpcClient, crypto.PubkeyToAddress(sr.privKey.PublicKey))
	if err != nil {
		log.Error("parse l1 mempool error", "error", err)
	} else {
		sr.pendingTxs.Recover(txs, sr.abi)
	}

	go utils.Loop(sr.ctx, 500*time.Millisecond, func() {
		sr.rollupFinalizeMu.Lock()
		defer sr.rollupFinalizeMu.Unlock()
		if err := sr.rollup(); err != nil {
			if utils.IsRpcErr(err) {
				sr.metrics.IncRpcErrors()
			}
			log.Error("rollup failed,wait for the next try", "error", err)
		}
	})

	if sr.cfg.Finalize {

		go utils.Loop(sr.ctx, time.Second*2, func() {
			sr.rollupFinalizeMu.Lock()
			defer sr.rollupFinalizeMu.Unlock()

			if err := sr.finalize(); err != nil {
				log.Error("finalize failed", "error", err)
			}
		})
	}

	var processtxMu sync.Mutex
	go utils.Loop(sr.ctx, time.Second*2, func() {
		processtxMu.Lock()
		defer processtxMu.Unlock()
		if err := sr.ProcessTx(); err != nil {
			log.Error("process tx err", "error", err)
		}
	})

}

func (sr *Rollup) ProcessTx() error {

	// case 1: in mempool
	//          -> check timeout
	// case 2: no in mempool
	// case 2.1: discarded
	// case 2.2: tx included -> success
	// case 2.3: tx included -> failed
	//          -> reset index to failed index

	// get all local txs
	txRecords := sr.pendingTxs.GetAll()
	if len(txRecords) == 0 {
		return nil
	}

	// query tx status
	for _, txRecord := range txRecords {

		rtx := txRecord.tx
		method := utils.ParseMethod(rtx, sr.abi)
		log.Info("process tx", "txHash", rtx.Hash().Hex(), "nonce", rtx.Nonce(), "method", method)
		// query tx
		_, ispending, err := sr.L1Client.TransactionByHash(context.Background(), txRecord.tx.Hash())
		if err != nil {
			if !utils.ErrStringMatch(err, ethereum.NotFound) {
				return fmt.Errorf("query tx  error:%w, tx: %s, nonce: %d", err, rtx.Hash().Hex(), rtx.Nonce())
			}
			sr.pendingTxs.IncQueryTimes(rtx.Hash()) // not found in mempool, increase query times
		} else {
			log.Info("query tx success", "txhash", rtx.Hash().Hex(), "pending", ispending)
		}

		// exist in mempool
		if ispending {
			if txRecord.sendTime+uint64(sr.cfg.TxTimeout.Seconds()) < uint64(time.Now().Unix()) {
				log.Info("tx timeout", "tx", rtx.Hash().Hex(), "nonce", rtx.Nonce(), "method", method)
				newtx, err := sr.ReSubmitTx(false, &rtx)
				if err != nil {
					log.Error("resubmit tx", "error", err, "tx", rtx.Hash().Hex(), "nonce", rtx.Nonce())
					return fmt.Errorf("resubmit tx error:%w", err)
				} else {
					log.Info("replace success", "old tx", rtx.Hash().Hex(), "new tx", newtx.Hash(), "nonce", rtx.Nonce())
					sr.pendingTxs.Remove(rtx.Hash())
					sr.pendingTxs.Add(*newtx)
				}
			}
		} else { // not in mempool
			receipt, err := sr.L1Client.TransactionReceipt(context.Background(), rtx.Hash())
			if err != nil {
				log.Error("query tx receipt error", "tx", rtx.Hash().Hex(), "nonce", rtx.Nonce(), "error", err)
				if !utils.ErrStringMatch(err, ethereum.NotFound) {
					return err
				}

				// sr.pendingTxs.txinfos
				if txRecord.queryTimes >= 5 {
					log.Warn("tx discarded",
						"hash", rtx.Hash().Hex(),
						"nonce", rtx.Nonce(),
						"query_times", txRecord.queryTimes,
					)
					replacedtx, err := sr.ReSubmitTx(true, &rtx)
					if err != nil {
						log.Error("resend discarded tx", "old_tx", rtx.Hash().Hex(), "nonce", rtx.Nonce(), "error", err)
						if utils.ErrStringMatch(err, core.ErrNonceTooLow) {
							log.Info("discarded tx removed",
								"hash", rtx.Hash().Hex(),
								"nonce", rtx.Nonce(),
								"method", method,
							)
							sr.pendingTxs.Remove(rtx.Hash())
							return nil
						}
						return fmt.Errorf("resend discarded tx: %w", err)
					} else {
						sr.pendingTxs.Remove(rtx.Hash())
					}
					sr.pendingTxs.Add(*replacedtx)
					log.Info("resend discarded tx", "old tx", rtx.Hash().Hex(), "new tx", replacedtx.Hash().Hex(), "nonce", replacedtx.Nonce())
				} else {
					log.Info("tx is not found, neither in mempool nor in block", "hash", rtx.Hash().Hex(), "nonce", rtx.Nonce(), "query_times", txRecord.queryTimes)
				}
			} else {
				logs := utils.ParseBusinessInfo(rtx, sr.abi)
				logs = append(logs,
					"block", receipt.BlockNumber,
					"hash", rtx.Hash().String(),
					"status", receipt.Status,
					"gas_used", receipt.GasUsed,
					"type", rtx.Type(),
					"nonce", rtx.Nonce(),
					"blob_fee_cap", rtx.BlobGasFeeCap(),
					"blob_gas", rtx.BlobGas(),
					"tx_size", rtx.Size(),
					"gas_limit", rtx.Gas(),
					"gas_price", rtx.GasPrice(),
				)

				log.Info("tx included",
					logs...,
				)

				if receipt.Status != types.ReceiptStatusSuccessful {
					// if tx is commitBatch
					if method == "commitBatch" {
						parentindex := utils.ParseParentBatchIndex(rtx.Data())
						index := parentindex + 1

						// prevent the SetFailedStatus operation from
						// happening between RemoveRollupRestriction
						// and SetPindex in the rollup function
						sr.rollupFinalizeMu.Lock()
						sr.pendingTxs.SetFailedStatus(index)
						sr.rollupFinalizeMu.Unlock()

					}

				} else {
					if method == "commitBatch" && sr.pendingTxs.failedIndex != nil {
						log.Info("fail revover", "failed_index", sr.pendingTxs.failedIndex)
						sr.pendingTxs.RemoveRollupRestriction()
					}
				}

				sr.pendingTxs.Remove(rtx.Hash())
				// set metrics
				fee := calcFee(receipt)
				if fee == 0 {
					log.Warn("fee is zero", "hash", rtx.Hash().Hex())
				}
				if method == "commitBatch" {
					sr.metrics.SetRollupCost(fee)
				} else if method == "finalizeBatch" {
					sr.metrics.SetFinalizeCost(fee)
				}
			}

		}

	}

	return nil

}

func (sr *Rollup) finalize() error {
	// get last finalized
	lastFinalized, err := sr.Rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last finalized error:%v", err)
	}
	// get last committed
	lastCommited, err := sr.Rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last committed error:%v", err)
	}

	target := big.NewInt(int64(sr.pendingTxs.pfinalize + 1))
	if target.Cmp(lastFinalized) <= 0 {
		target = new(big.Int).Add(lastFinalized, big.NewInt(1))
	}

	if target.Cmp(lastCommited) > 0 {
		log.Info("no need to finalize", "last_finalized", lastFinalized.Uint64(), "last_committed", lastCommited.Uint64())
		return nil
	}

	log.Info("finalize info",
		"lastFianlzied", lastFinalized,
		"lastCommited", lastCommited,
		"finalize_index", target,
	)

	// batch exist
	existed, err := sr.Rollup.BatchExist(nil, target)
	if err != nil {
		log.Error("query batch exist", "err", err)
		return err
	}
	if !existed {
		log.Warn("finalized batch not existed")
		return nil
	}

	// in challange window
	inWindow, err := sr.Rollup.BatchInsideChallengeWindow(nil, target)
	if err != nil {
		return fmt.Errorf("get batch inside challenge window error:%v", err)
	}
	if inWindow {
		log.Info("batch inside challenge window, wait")
		return nil
	}
	// finalize
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	// calldata
	calldata, err := sr.abi.Pack("finalizeBatch", target)
	if err != nil {
		return fmt.Errorf("pack finalizeBatch error:%v", err)
	}
	tip, feecap, _, err := sr.GetGasTipAndCap()
	if err != nil {
		log.Error("get gas tip and cap error", "business", "finalize")
		return fmt.Errorf("get gas tip and cap error:%v", err)
	}

	gasDefault := uint64(50_0000)
	gas, err := sr.EstimateGas(opts.From, sr.rollupAddr, calldata, feecap, tip)
	if err != nil {
		gas = gasDefault
	} else {
		gas = gas * 12 / 10 // add a buffer
	}

	var nonce uint64
	if sr.pendingTxs.pnonce != 0 {
		nonce = sr.pendingTxs.pnonce + 1
	} else {
		nonce, err = sr.L1Client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey))
		if err != nil {
			return fmt.Errorf("query layer1 nonce error:%v", err.Error())
		}
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   sr.chainId,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: feecap,
		Gas:       gas,
		To:        &sr.rollupAddr,
		Data:      calldata,
	})

	if uint64(tx.Size()) > txMaxSize {
		return core.ErrOversizedData
	}

	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return fmt.Errorf("sign tx error:%v", err)
	}

	log.Info("finalize tx info",
		"batch_index", target,
		"last_commited", lastCommited,
		"last_finalized", lastFinalized,
		"hash", signedTx.Hash().String(),
		"type", signedTx.Type(),
		"nonce", signedTx.Nonce(),
		"gas", signedTx.Gas(),
		"tip", signedTx.GasTipCap().String(),
		"fee_cap", signedTx.GasFeeCap().String(),
		"size", signedTx.Size(),
	)

	err = sr.SendTx(signedTx)
	if err != nil {
		log.Error("send finalize tx to mempool", "error", err.Error())
		if utils.ErrStringMatch(err, core.ErrNonceTooLow) {
			// adjust nonce
			n1, _, err := utils.ParseNonce(err.Error())
			if err != nil {
				return fmt.Errorf("parse nonce err: %w", err)
			}
			sr.pendingTxs.SetNonce(n1 - 1)
			log.Info("update pnonce", "nonce", n1-1)
		}
		return fmt.Errorf("send tx error:%v", err.Error())
	} else {
		log.Info("finalzie tx sent")

		sr.pendingTxs.SetNonce(signedTx.Nonce())
		sr.pendingTxs.SetPFinalize(target.Uint64())
		sr.pendingTxs.Add(*signedTx)
	}

	return nil

}

func (sr *Rollup) rollup() error {

	if !sr.cfg.PriorityRollup {
		// is the turn of the submitter
		currentSubmitter, start, end, err := sr.getCurrentSubmitter()
		if err != nil {
			return fmt.Errorf("get next submitter error:%v", err)
		}
		log.Info("rotator info",
			"turn", currentSubmitter.Hex(),
			"cur", sr.walletAddr(),
			"start", start,
			"end", end,
			"now", time.Now().Unix(),
		)

		if currentSubmitter.Hex() == sr.walletAddr() {
			left := int64(end) - time.Now().Unix()
			if left < rotatorBuffer {
				log.Info("rollup time not enough, wait next turn", "left", left)
				return nil
			}

			log.Info("start to rollup")
		} else {
			log.Info("wait for my turn")
			return nil
		}
	}

	if len(sr.pendingTxs.txinfos) > 11 {
		log.Info("too many txs in mempool, wait")
		return nil
	}

	var nonce uint64
	var batchIndex uint64
	var err error

	cindexBig, err := sr.Rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last committed batch index error:%v", err)
	}
	cindex := cindexBig.Uint64()

	if sr.pendingTxs.failedIndex != nil && cindex >= *sr.pendingTxs.failedIndex {
		sr.pendingTxs.RemoveRollupRestriction()
	}

	if sr.pendingTxs.failedIndex != nil {
		batchIndex = *sr.pendingTxs.failedIndex
	} else {
		if sr.pendingTxs.pindex != 0 {
			if cindex > sr.pendingTxs.pindex {
				batchIndex = cindex + 1
			} else {
				batchIndex = sr.pendingTxs.pindex + 1
			}

		} else {
			batchIndex = cindex + 1
		}
	}

	log.Info("batch info", "last_commit_batch", batchIndex-1, "batch_will_get", batchIndex)
	if sr.pendingTxs.ExistedIndex(batchIndex) {
		log.Info("batch index already commited", "index", batchIndex)
		return nil
	}

	if sr.pendingTxs.failedIndex != nil && batchIndex > *sr.pendingTxs.failedIndex {
		log.Warn("rollup rejected", "index", batchIndex)
		return nil
	}

	batch, err := GetRollupBatchByIndex(batchIndex, sr.L2Clients)
	if err != nil {
		return fmt.Errorf("get rollup batch by index err:%v", err)
	}

	// check if the batch is valid
	if batch == nil {
		log.Info("new batch not found, wait for the next turn")
		return nil
	}

	if len(batch.Signatures) == 0 {
		log.Info("length of batch signature is empty, wait for the next turn")
		return nil
	}

	var chunks [][]byte
	// var blobChunk []byte
	for _, chunk := range batch.Chunks {
		chunks = append(chunks, chunk)
	}
	signature, err := sr.aggregateSignatures(batch.Signatures)
	if err != nil {
		return err
	}
	rollupBatch := bindings.IRollupBatchData{
		Version:                uint8(batch.Version),
		ParentBatchHeader:      batch.ParentBatchHeader,
		Chunks:                 chunks,
		SkippedL1MessageBitmap: batch.SkippedL1MessageBitmap,
		PrevStateRoot:          batch.PrevStateRoot,
		PostStateRoot:          batch.PostStateRoot,
		WithdrawalRoot:         batch.WithdrawRoot,
		Signature:              *signature,
	}

	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	sequencerVersion, err := sr.L1Sequencer.NewestVersion(nil)
	if err != nil {
		return fmt.Errorf("get sequencer version error:%v", err)
	}

	// tip and cap
	tip, gasFeeCap, blobFee, err := sr.GetGasTipAndCap()
	if err != nil {
		return fmt.Errorf("get gas tip and cap error:%v", err)
	}

	// hotfix
	tip = new(big.Int).Mul(tip, big.NewInt(3))
	gasFeeCap = new(big.Int).Mul(gasFeeCap, big.NewInt(3))

	// calldata encode
	calldata, err := sr.abi.Pack("commitBatch", rollupBatch, sequencerVersion, []common.Address{}, signature.Signature)
	if err != nil {
		return fmt.Errorf("pack calldata error:%v", err)
	}

	gas, err := sr.EstimateGas(opts.From, sr.rollupAddr, calldata, gasFeeCap, tip)
	if err != nil {
		if sr.pendingTxs.HaveFailed() {
			log.Warn("estimate gas err, wait",
				"err", err,
				"update_index", cindex+1,
			)
			sr.pendingTxs.ResetFailedIndex(cindex + 1)
			return nil
		} else {
			msgcnt := utils.ParseL1MessageCnt(batch.Chunks)
			gas = RoughEstimateGas(msgcnt)
		}
	}

	// gas buffer
	gas = gas * sr.cfg.GasLimitBuffer / 100

	if sr.pendingTxs.pnonce != 0 {
		nonce = sr.pendingTxs.pnonce + 1
	} else {
		nonce, err = sr.L1Client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey))
		if err != nil {
			return fmt.Errorf("query layer1 nonce error:%v", err.Error())
		}
	}

	var tx *types.Transaction
	if len(batch.Sidecar.Blobs) > 0 {
		versionedHashes := make([]common.Hash, 0)
		for _, commit := range batch.Sidecar.Commitments {
			versionedHashes = append(versionedHashes, kZGToVersionedHash(commit))
		}
		// blob tx
		tx = types.NewTx(&types.BlobTx{
			ChainID:    uint256.MustFromBig(sr.chainId),
			Nonce:      nonce,
			GasTipCap:  uint256.MustFromBig(big.NewInt(tip.Int64())),
			GasFeeCap:  uint256.MustFromBig(big.NewInt(gasFeeCap.Int64())),
			Gas:        gas,
			To:         sr.rollupAddr,
			Data:       calldata,
			BlobFeeCap: uint256.MustFromBig(blobFee),
			BlobHashes: versionedHashes,
			Sidecar: &types.BlobTxSidecar{
				Blobs:       batch.Sidecar.Blobs,
				Commitments: batch.Sidecar.Commitments,
				Proofs:      batch.Sidecar.Proofs,
			},
		})

	} else {
		tx = types.NewTx(&types.DynamicFeeTx{
			ChainID:   sr.chainId,
			Nonce:     nonce,
			GasTipCap: tip,
			GasFeeCap: gasFeeCap,
			Gas:       gas,
			To:        &sr.rollupAddr,
			Data:      calldata,
		})
	}

	opts.Nonce = big.NewInt(int64(nonce))
	var signedTx *types.Transaction
	if tx.Type() == types.BlobTxType {
		signedTx, err = types.SignTx(tx, types.NewLondonSignerWithEIP4844(sr.chainId), sr.privKey)
		if err != nil {
			return fmt.Errorf("sign tx error:%v", err)
		}
	} else {
		signedTx, err = opts.Signer(opts.From, tx)
		if err != nil {
			return fmt.Errorf("sign tx error:%v", err)
		}
	}

	log.Info("rollup tx info",
		"batch_index", batchIndex,
		"hash", signedTx.Hash().String(),
		"type", signedTx.Type(),
		"nonce", signedTx.Nonce(),
		"gas", signedTx.Gas(),
		"tip", signedTx.GasTipCap().String(),
		"fee_cap", signedTx.GasFeeCap().String(),
		"blob_fee_cap", signedTx.BlobGasFeeCap(),
		"blob_gas", signedTx.BlobGas(),
		"size", signedTx.Size(),
		"blob_len", len(signedTx.BlobHashes()),
	)

	err = sr.SendTx(signedTx)
	if err != nil {
		log.Error("send tx to mempool", "error", err.Error())
		if utils.ErrStringMatch(err, core.ErrNonceTooLow) {
			// adjust nonce
			n1, _, err := utils.ParseNonce(err.Error())
			if err != nil {
				return fmt.Errorf("parse nonce err: %w", err)
			}
			sr.pendingTxs.SetNonce(n1 - 1)
			log.Info("update pnonce", "nonce", n1-1)
		}
		return fmt.Errorf("send tx error:%v", err.Error())
	} else {
		log.Info("rollup tx send to mempool succuess", "hash", signedTx.Hash().String())

		sr.pendingTxs.SetPindex(batchIndex)
		sr.pendingTxs.SetNonce(tx.Nonce())
		sr.pendingTxs.Add(*signedTx)
	}

	return nil
}

func (sr *Rollup) EstimateGas(from, to common.Address, data []byte, feecap *big.Int, tip *big.Int) (uint64, error) {

	gas, err := sr.L1Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasFeeCap: feecap,
		GasTipCap: tip,
		Data:      data,
	})
	return gas, err

}

func (sr *Rollup) aggregateSignatures(blsSignatures []eth.RPCBatchSignature) (*bindings.IRollupBatchSignature, error) {
	if len(blsSignatures) == 0 {
		return nil, fmt.Errorf("invalid batch signature")
	}
	signers := make([]*big.Int, len(blsSignatures))
	sigs := make([]blssignatures.Signature, 0)
	for i, bz := range blsSignatures {
		if len(bz.Signature) > 0 {
			sig, err := blssignatures.SignatureFromBytes(bz.Signature)
			if err != nil {
				return nil, err
			}
			sigs = append(sigs, sig)
			signers[i] = big.NewInt(int64(bz.Signer))
		}
	}
	aggregatedSig := blssignatures.AggregateSignatures(sigs)
	blsSignature := bls12381.NewG1().EncodePoint(aggregatedSig)
	rollupBatchSignature := bindings.IRollupBatchSignature{
		Version:   big.NewInt(int64(blsSignatures[0].Version)),
		Signers:   signers,
		Signature: blsSignature,
	}
	return &rollupBatchSignature, nil
}

func (sr *Rollup) GetGasTipAndCap() (*big.Int, *big.Int, *big.Int, error) {
	tip, err := sr.L1Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	head, err := sr.L1Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, nil, err
	}
	var gasFeeCap *big.Int
	if head.BaseFee != nil {
		gasFeeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		gasFeeCap = new(big.Int).Set(tip)
	}

	// calc blob fee cap
	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas)
	}
	blobFee = new(big.Int).Mul(blobFee, big.NewInt(200))
	blobFee = new(big.Int).Div(blobFee, big.NewInt(100))
	return tip, gasFeeCap, blobFee, nil
}

func (sr *Rollup) waitForReceipt(txHash common.Hash) (*types.Receipt, error) {
	t := time.NewTicker(time.Second)
	receipt := new(types.Receipt)
	var err error
	for range t.C {
		receipt, err = sr.L1Client.TransactionReceipt(context.Background(), txHash)
		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt != nil {
			t.Stop()
			break
		}
	}
	return receipt, nil
}

func (sr *Rollup) waitReceiptWithTimeout(time time.Duration, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(sr.ctx, time)
	defer cancel()
	return sr.waitReceiptWithCtx(ctx, txHash)
}

func (sr *Rollup) waitReceiptWithCtx(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	t := time.NewTicker(time.Second)
	receipt := new(types.Receipt)
	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("timeout")
		case <-t.C:
			receipt, err = sr.L1Client.TransactionReceipt(context.Background(), txHash)
			if errors.Is(err, ethereum.NotFound) {
				continue
			}
			if err != nil {
				return nil, err
			}
			if receipt != nil {
				t.Stop()
				return receipt, nil
			}
		}

	}

}

// Init is run before the submitter to check whether the submitter can be started
func (sr *Rollup) Init() error {

	isSequencer, err := sr.inSequencersSet()
	if err != nil {
		return err
	}

	if !isSequencer {
		return fmt.Errorf("this account is not sequencer")
	}

	return nil
}

func (sr *Rollup) walletAddr() string {
	return crypto.PubkeyToAddress(sr.privKey.PublicKey).Hex()
}

func GetRollupBatchByIndex(index uint64, clients []iface.L2Client) (*eth.RPCRollupBatch, error) {
	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query")
	}
	for _, client := range clients {
		batch, err := client.GetRollupBatchByIndex(context.Background(), index)
		if err != nil {
			log.Warn("failed to get batch", "error", err)
			continue
		}
		if batch != nil && len(batch.Signatures) > 0 {
			return batch, nil
		}
	}

	return nil, nil
}

func UpdateGasLimit(tx *types.Transaction) (*types.Transaction, error) {
	// add buffer to gas limit (*1.2)
	newGasLimit := tx.Gas() * 12 / 10

	var newTx *types.Transaction
	if tx.Type() == types.LegacyTxType {

		newTx = types.NewTx(&types.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: big.NewInt(tx.GasPrice().Int64()),
			Gas:      newGasLimit,
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		})
	} else if tx.Type() == types.DynamicFeeTxType {
		newTx = types.NewTx(&types.DynamicFeeTx{
			Nonce:     tx.Nonce(),
			GasTipCap: big.NewInt(tx.GasTipCap().Int64()),
			GasFeeCap: big.NewInt(tx.GasFeeCap().Int64()),
			Gas:       newGasLimit,
			To:        tx.To(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	} else if tx.Type() == types.BlobTxType {
		newTx = types.NewTx(&types.BlobTx{
			ChainID:    uint256.MustFromBig(tx.ChainId()),
			Nonce:      tx.Nonce(),
			GasTipCap:  uint256.MustFromBig(big.NewInt(tx.GasTipCap().Int64())),
			GasFeeCap:  uint256.MustFromBig(big.NewInt(tx.GasFeeCap().Int64())),
			Gas:        newGasLimit,
			To:         *tx.To(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       tx.Data(),
			BlobFeeCap: uint256.MustFromBig(tx.BlobGasFeeCap()),
			BlobHashes: tx.BlobHashes(),
			Sidecar:    tx.BlobTxSidecar(),
		})

	} else {
		return nil, fmt.Errorf("unknown tx type:%v", tx.Type())
	}
	return newTx, nil
}

// send tx to l1 with business logic check
func (r *Rollup) SendTx(tx *types.Transaction) error {

	// judge tx info is valid
	if tx == nil {
		return errors.New("nil tx")
	}

	err := sendTx(r.L1Client, r.cfg.TxFeeLimit, tx)
	if err != nil {
		return err
	}

	return nil

}

// send tx to l1 with business logic check
func sendTx(client iface.Client, txFeeLimit uint64, tx *types.Transaction) error {
	// fee limit
	if txFeeLimit > 0 {
		var fee *big.Int
		// calc tx gas fee
		if tx.Type() == types.BlobTxType {
			// blob fee
			blobfee := new(big.Int).Mul(tx.BlobGasFeeCap(), big.NewInt(int64(tx.BlobGas())))
			// calldata fee
			dfee := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas())))
			fee = new(big.Int).Add(blobfee, dfee)
		} else {
			// calldata fee
			fee = new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas())))
		}

		if fee.Cmp(big.NewInt(int64(txFeeLimit))) > 0 {
			return fmt.Errorf("%v:limit=%v,but got=%v", utils.ErrExceedFeeLimit, txFeeLimit, fee.String())
		}
	}

	return client.SendTransaction(context.Background(), tx)
}

func (sr *Rollup) ReSubmitTx(resend bool, tx *types.Transaction) (*types.Transaction, error) {
	if tx == nil {
		return nil, errors.New("nil tx")
	}

	// for sign
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return nil, fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	method := "replaced tx"
	if resend {
		method = "resubmitted tx"
	}

	// replaced tx info
	log.Info(method,
		"hash", tx.Hash().String(),
		"gas_fee_cap", tx.GasFeeCap().String(),
		"gas_tip", tx.GasTipCap().String(),
		"blob_fee_cap", tx.BlobGasFeeCap().String(),
		"gas", tx.Gas(),
		"nonce", tx.Nonce(),
	)

	tip, gasFeeCap, blobFeeCap, err := sr.GetGasTipAndCap()
	if err != nil {
		log.Error("get tip and cap", "err", err)
	}
	if !resend {
		// bump tip & feeCap
		bumpedFeeCap := calcThresholdValue(tx.GasFeeCap(), tx.Type() == types.BlobTxType)
		bumpedTip := calcThresholdValue(tx.GasTipCap(), tx.Type() == types.BlobTxType)

		// if bumpedTip > tip
		if bumpedTip.Cmp(tip) > 0 {
			tip = bumpedTip
		}

		if bumpedFeeCap.Cmp(gasFeeCap) > 0 {
			gasFeeCap = bumpedFeeCap
		}

		if tx.Type() == types.BlobTxType {
			bumpedBlobFeeCap := calcThresholdValue(tx.BlobGasFeeCap(), tx.Type() == types.BlobTxType)
			if bumpedBlobFeeCap.Cmp(blobFeeCap) > 0 {
				blobFeeCap = bumpedBlobFeeCap
			}
		}
	}

	var newTx *types.Transaction
	switch tx.Type() {
	case types.DynamicFeeTxType:
		newTx = types.NewTx(&types.DynamicFeeTx{
			To:        tx.To(),
			Nonce:     tx.Nonce(),
			GasFeeCap: gasFeeCap,
			GasTipCap: tip,
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	case types.BlobTxType:

		newTx = types.NewTx(&types.BlobTx{
			ChainID:    uint256.MustFromBig(tx.ChainId()),
			Nonce:      tx.Nonce(),
			GasTipCap:  uint256.MustFromBig(tip),
			GasFeeCap:  uint256.MustFromBig(gasFeeCap),
			Gas:        tx.Gas(),
			To:         *tx.To(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       tx.Data(),
			BlobFeeCap: uint256.MustFromBig(blobFeeCap),
			BlobHashes: tx.BlobHashes(),
			Sidecar:    tx.BlobTxSidecar(),
		})

	default:
		return nil, fmt.Errorf("replace unknown tx type:%v", tx.Type())

	}

	log.Info("new tx info",
		"tx_type", newTx.Type(),
		"gas_tip", tip.String(), //todo: convert to gwei
		"gas_fee_cap", gasFeeCap.String(), //todo: convert to gwei
		"blob_fee_cap", blobFeeCap.String(), //todo: convert to gwei
	)
	// sign tx
	opts.Nonce = big.NewInt(int64(newTx.Nonce()))
	newTx, err = opts.Signer(opts.From, newTx)
	if err != nil {
		return nil, fmt.Errorf("sign tx error:%w", err)
	}
	// send tx
	err = sr.SendTx(newTx)
	if err != nil {
		return nil, fmt.Errorf("send tx error:%w", err)
	}

	return newTx, nil
}

func (r *Rollup) getCurrentSubmitter() (*common.Address, uint64, uint64, error) {

	for _, l2Submitter := range r.L2Submitters {
		current, start, end, err := l2Submitter.GetCurrentSubmitter(nil)
		if err != nil {
			log.Warn("get current submitter error", "error", err)
			continue
		}
		return &current, start.Uint64(), end.Uint64(), nil

	}

	return nil, 0, 0, errors.New("failed to get current submitter")
}

func (r *Rollup) inSequencersSet() (bool, error) {

	for _, l2Sequencer := range r.L2Sequencers {
		isSequencer, _, err := l2Sequencer.InSequencersSet(
			nil,
			false,
			common.HexToAddress(r.walletAddr()),
		)
		if err != nil {
			log.Warn("get in sequencer set error", "error", err)
			continue
		}
		return isSequencer, nil
	}
	return false, errors.New("failed to get in sequencer set")
}
